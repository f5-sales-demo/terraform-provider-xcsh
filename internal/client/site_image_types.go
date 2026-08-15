// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package client

import "context"

// SiteImageDownloadRequest is the public F5 XC image-download action request.
// The provider value is an API contract (for example, "Kvm"), rather than a
// Secure Mesh Site v2 configuration field.
type SiteImageDownloadRequest struct {
	Provider string `json:"provider"`
}

// SiteImageDownloadResponse holds the short-lived image URLs returned by F5
// XC. Keep both fields separate: callers use the image URL to provision the
// VM and the MD5 URL to verify the downloaded artifact.
type SiteImageDownloadResponse struct {
	ImageDownloadURL    string `json:"image_download_url"`
	ImageMD5DownloadURL string `json:"image_md5_download_url"`
}

// GetSiteImage requests the current Customer Edge image URLs for one platform.
// This is a public F5 XC action, not an object nested under securemesh_site_v2.
func (c *Client) GetSiteImage(ctx context.Context, provider string) (*SiteImageDownloadResponse, error) {
	var result SiteImageDownloadResponse
	err := c.Post(ctx, "/api/register/namespaces/system/get-image-download-url", SiteImageDownloadRequest{Provider: provider}, &result)
	return &result, err
}
