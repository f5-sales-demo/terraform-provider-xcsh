package client

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// KVMImageResolverContract is emitted only after the immutable API image
// contract passes the generator's exact semantic and provenance validation.
type KVMImageResolverContract struct {
	ConfigurationList string
	ConfigurationGet  string
	SiteList          string
	Query             string
}

// KVMImage contains no object UID or raw backend diagnostic. DownloadURL is
// sensitive Terraform state. The consumer must verify MD5 after downloading;
// resolving metadata does not claim an artifact download or boot succeeded.
type KVMImage struct {
	DownloadURL string
	ImageName   string
	MD5         string
}

var imageSiteName = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?$`)
var imageMD5 = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

type imageIdentity struct{ configuration, site string }

func (c *Client) imageIdentity(ctx context.Context, name string, contract KVMImageResolverContract) (imageIdentity, error) {
	var configurations struct {
		Items []struct{ Name, UID string } `json:"items"`
	}
	if err := c.Get(ctx, contract.ConfigurationList, &configurations); err != nil {
		return imageIdentity{}, fmt.Errorf("KVM image configuration list failed")
	}
	identity, matches := imageIdentity{}, 0
	for _, item := range configurations.Items {
		if item.Name == name {
			matches++
			identity.configuration = item.UID
		}
	}
	if matches != 1 || strings.TrimSpace(identity.configuration) == "" {
		return imageIdentity{}, fmt.Errorf("KVM image requires exactly one configuration with a nonempty UID")
	}
	var configuration struct {
		Spec map[string]any `json:"spec"`
	}
	if err := c.Get(ctx, strings.ReplaceAll(contract.ConfigurationGet, "{site_name}", url.PathEscape(name)), &configuration); err != nil {
		return imageIdentity{}, fmt.Errorf("KVM image configuration read failed")
	}
	if _, ok := configuration.Spec["kvm"].(map[string]any); !ok {
		return imageIdentity{}, fmt.Errorf("KVM image requires a KVM configuration")
	}
	var sites struct {
		Items []struct {
			UID   string                      `json:"uid"`
			Owner *struct{ Kind, UID string } `json:"owner_view"`
		} `json:"items"`
	}
	if err := c.Get(ctx, contract.SiteList, &sites); err != nil {
		return imageIdentity{}, fmt.Errorf("KVM image Site list failed")
	}
	matches = 0
	for _, item := range sites.Items {
		if item.Owner != nil && item.Owner.Kind == "securemesh_site_v2" && item.Owner.UID == identity.configuration {
			matches++
			identity.site = item.UID
		}
	}
	if matches != 1 || strings.TrimSpace(identity.site) == "" {
		return imageIdentity{}, fmt.Errorf("KVM image requires exactly one Site owned by the selected configuration")
	}
	return identity, nil
}

// ResolveKVMImage uses the current console's owner-view join. Never substitute a
// configuration UID, caller UID, provider-only request, or static image default.
// Rechecking identity after the query rejects concurrent deletion/recreation.
func (c *Client) ResolveKVMImage(ctx context.Context, siteName string, contract KVMImageResolverContract) (KVMImage, error) {
	if !imageSiteName.MatchString(siteName) {
		return KVMImage{}, fmt.Errorf("KVM image requires a valid site name")
	}
	if contract.ConfigurationList == "" || contract.ConfigurationGet == "" || contract.SiteList == "" || contract.Query == "" {
		return KVMImage{}, fmt.Errorf("KVM image resolution contract is incomplete")
	}
	before, err := c.imageIdentity(ctx, siteName, contract)
	if err != nil {
		return KVMImage{}, err
	}
	var response struct {
		Images map[string]struct {
			DownloadURL string  `json:"download_image_link"`
			ImageName   string  `json:"copy_image_name"`
			MD5         string  `json:"image_md5_sum"`
			Error       *string `json:"error_description"`
		} `json:"images"`
	}
	if err := c.Post(ctx, contract.Query, map[string]any{"uids": []string{before.site}}, &response); err != nil {
		// The API error may contain object identifiers or signed URLs.
		return KVMImage{}, fmt.Errorf("KVM image query failed")
	}
	item, exists := response.Images[before.site]
	if !exists || len(response.Images) != 1 {
		return KVMImage{}, fmt.Errorf("KVM image response must contain only the requested Site mapping")
	}
	if item.Error == nil || *item.Error != "" {
		return KVMImage{}, fmt.Errorf("KVM image response has a missing or nonempty error status")
	}
	if strings.TrimSpace(item.ImageName) == "" || !imageMD5.MatchString(item.MD5) {
		return KVMImage{}, fmt.Errorf("KVM image response has incomplete image name or checksum")
	}
	link, err := url.ParseRequestURI(item.DownloadURL)
	if err != nil || link.Scheme != "https" || link.Host != "downloads.volterra.io" || link.User != nil || link.Fragment != "" || strings.Contains(item.DownloadURL, "#") || link.Path == "" || link.Path == "/" {
		return KVMImage{}, fmt.Errorf("KVM image response has an unsupported download URL")
	}
	after, err := c.imageIdentity(ctx, siteName, contract)
	if err != nil {
		return KVMImage{}, err
	}
	if before != after {
		return KVMImage{}, fmt.Errorf("KVM image ownership changed during resolution")
	}
	return KVMImage{DownloadURL: item.DownloadURL, ImageName: item.ImageName, MD5: strings.ToLower(item.MD5)}, nil
}
