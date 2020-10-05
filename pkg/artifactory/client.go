package artifactory

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	aqlSearchEndpoint      = "/artifactory/api/search/aql"
	fileStatisticsEndpoint = "/artifactory/api/storage"
)

// Client represents client for interaction with Artifactory API
type Client struct {
	config     *Config
	httpClient *http.Client
}

// NewClient returns an instance of Artifactory API client
func NewClient(config *Config) (*Client, error) {
	if config.BaseURL == "" {
		return nil, fmt.Errorf("no base URL specified")
	}
	config.BaseURL = strings.TrimRight(config.BaseURL, "/")

	if config.User == "" || config.Password == "" {
		return nil, fmt.Errorf("invalid credentials provided")
	}

	return newClient(config, http.DefaultClient), nil
}

func newClient(config *Config, httpClient *http.Client) *Client {
	return &Client{config, httpClient}
}

// ListArtifactsByRepo searches for artifacts for given repository
func (c *Client) ListArtifactsByRepo(ctx context.Context, repo string) ([]Artifact, error) {
	url := fmt.Sprintf("%s%s", c.config.BaseURL, aqlSearchEndpoint)
	searchAQL := fmt.Sprintf(`items.find({"repo":{"$eq":"%s"}})`, repo)

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(searchAQL))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.config.User, c.config.Password)
	req.Header.Set("Content-Type", "text/plain")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", res.Status)
	}

	var a AQLSearchResponse
	err = json.NewDecoder(res.Body).Decode(&a)
	if err != nil {
		return nil, err
	}

	return a.Results, nil
}

// GetFileStatistics returns file statistics like download count for file at given path
func (c *Client) GetFileStatistics(ctx context.Context, path string) (*FileStatisticsResponse, error) {
	url := fmt.Sprintf("%s%s/%s?stats", c.config.BaseURL, fileStatisticsEndpoint, path)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.config.User, c.config.Password)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: %s", res.Status)
	}

	var f FileStatisticsResponse
	err = json.NewDecoder(res.Body).Decode(&f)
	if err != nil {
		return nil, err
	}

	return &f, nil
}
