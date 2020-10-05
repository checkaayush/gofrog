package internal

import (
	"github.com/checkaayush/gofrog/pkg/artifactory"
)

type fileStatisticsMaxHeap []artifactory.FileStatisticsResponse

func (h fileStatisticsMaxHeap) Len() int {
	return len(h)
}

func (h fileStatisticsMaxHeap) Less(i, j int) bool {
	return h[i].DownloadCount > h[j].DownloadCount
}

func (h fileStatisticsMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *fileStatisticsMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(artifactory.FileStatisticsResponse))
}

func (h *fileStatisticsMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
