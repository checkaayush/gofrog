package artifactory

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_ListArtifactsByRepo(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
	}
	type args struct {
		ctx  context.Context
		repo string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Artifact
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.ListArtifactsByRepo(tt.args.ctx, tt.args.repo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListArtifactsByRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListArtifactsByRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetFileStatistics(t *testing.T) {
	type fields struct {
		config     *Config
		httpClient *http.Client
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FileStatisticsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				config:     tt.fields.config,
				httpClient: tt.fields.httpClient,
			}
			got, err := c.GetFileStatistics(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetFileStatistics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetFileStatistics() = %v, want %v", got, tt.want)
			}
		})
	}
}
