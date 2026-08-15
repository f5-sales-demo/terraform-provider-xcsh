// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSiteImage_UsesPublicImageDownloadContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("request method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/api/register/namespaces/system/get-image-download-url" {
			t.Errorf("request path = %q, want image-download action", r.URL.Path)
		}

		var request SiteImageDownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if request.Provider != "Kvm" {
			t.Errorf("provider = %q, want Kvm", request.Provider)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"image_download_url":"https://images.example.invalid/f5-ce.qcow2","image_md5_download_url":"https://images.example.invalid/f5-ce.qcow2.md5"}`))
	}))
	defer server.Close()

	response, err := NewClient(server.URL, "test-token").GetSiteImage(context.Background(), "Kvm")
	if err != nil {
		t.Fatalf("GetSiteImage() error = %v", err)
	}
	if response.ImageDownloadURL != "https://images.example.invalid/f5-ce.qcow2" {
		t.Errorf("ImageDownloadURL = %q", response.ImageDownloadURL)
	}
	if response.ImageMD5DownloadURL != "https://images.example.invalid/f5-ce.qcow2.md5" {
		t.Errorf("ImageMD5DownloadURL = %q", response.ImageMD5DownloadURL)
	}
}
