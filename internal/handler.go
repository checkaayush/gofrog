package internal

import (
	"github.com/checkaayush/gofrog/pkg/artifactory"
)

// Handler handles requests to interact with Artifactory API
type Handler struct {
	rtClient *artifactory.Client
}

// NewHandler creates a new instance of handler for Artifactory API interaction
func NewHandler(rtClient *artifactory.Client) *Handler {
	return &Handler{rtClient: rtClient}
}
