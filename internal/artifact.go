package internal

import (
	"container/heap"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	rt "github.com/checkaayush/gofrog/pkg/artifactory"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var numWorkers = 10

type artifactDownloadStats struct {
	Path         string `json:"path"`
	DownloadLink string `json:"downloadLink"`
	Downloads    int    `json:"downloads"`
}

type popularArtifactsResponse struct {
	Results []artifactDownloadStats `json:"results"`
	Error   string                  `json:"error"`
}

// GetMostPopularArtifacts returns most popular artifacts for given repo
func (h *Handler) GetMostPopularArtifacts(c echo.Context) (err error) {
	repo := c.QueryParam("repo")
	count, err := strconv.Atoi(c.QueryParam("count"))

	mostPopular := []artifactDownloadStats{}
	resp := popularArtifactsResponse{Results: mostPopular}
	if repo == "" || err != nil || count <= 0 {
		resp.Error = "Invalid request parameters"
		return c.JSON(http.StatusBadRequest, resp)
	}

	// Search artifacts in given repo
	artifacts, err := h.rtClient.ListArtifactsByRepo(context.Background(), repo)
	if err != nil {
		log.Infof("No artifacts found for repo %s. Error: %s", repo, err.Error())
		resp.Error = "No artifacts found for given repo"
		return c.JSON(http.StatusNotFound, resp)
	}
	log.Infof("Found %d artifacts", len(artifacts))

	// Start background workers for processing artifacts
	rtCh := make(chan rt.Artifact, len(artifacts))
	statsCh := make(chan *rt.FileStatisticsResponse, len(artifacts))
	for w := 0; w < numWorkers; w++ {
		go getArtifactStatistics(h.rtClient, rtCh, statsCh)
	}

	// Producer: Send artifacts to rtCh channel for further processing
	for _, artifact := range artifacts {
		rtCh <- artifact
	}
	close(rtCh)

	// Consumer: Receive fetched file statistics from statsCh channel
	stats := []rt.FileStatisticsResponse{}
	for i := 0; i < len(artifacts); i++ {
		stat := <-statsCh
		if stat != nil {
			stats = append(stats, *stat)
		}
	}

	if len(stats) == 0 {
		resp.Error = "Failed to fetch artifact download statistics"
		return c.JSON(http.StatusNotFound, resp)
	}

	// Prepare most popular artifacts using file statistics max-heap
	hp := fileStatisticsMaxHeap(stats)
	heap.Init(&hp)
	for i := 0; i < count; i++ {
		if hp.Len() <= 0 {
			break
		}

		stat := heap.Pop(&hp).(rt.FileStatisticsResponse)
		dlStat := artifactDownloadStats{
			Path:         getArtifactPathFromURI(stat.URI),
			DownloadLink: stat.URI,
			Downloads:    stat.DownloadCount,
		}
		mostPopular = append(mostPopular, dlStat)
	}

	res := popularArtifactsResponse{Results: mostPopular}
	return c.JSON(http.StatusOK, res)
}

func getArtifactStatistics(client *rt.Client, rtCh <-chan rt.Artifact, statsCh chan<- *rt.FileStatisticsResponse) {
	for artifact := range rtCh {
		endpoint := fmt.Sprintf("%s/%s/%s", artifact.Repo, artifact.Path, artifact.Name)

		stats, err := client.GetFileStatistics(context.Background(), endpoint)
		if err != nil {
			log.Debugf("Endpoint: %s, Error: %s", endpoint, err.Error())
		}

		statsCh <- stats
	}
}

func getArtifactPathFromURI(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	return strings.TrimPrefix(u.Path, "/artifactory/")
}
