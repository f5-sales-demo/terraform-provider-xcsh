package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func imageResolverFixture() KVMImageResolverContract {
	return KVMImageResolverContract{
		ConfigurationList: "/api/config/namespaces/system/securemesh_site_v2s",
		ConfigurationGet:  "/api/config/namespaces/system/securemesh_site_v2s/{site_name}",
		SiteList:          "/api/config/namespaces/system/sites",
		Query:             "/api/maurice/software_os_version",
	}
}

func TestResolveKVMImageObservedOwnership(t *testing.T) {
	for _, mode := range []string{"valid", "missing configuration", "duplicate configuration", "missing uid", "non-kvm", "missing owner", "duplicate owner", "wrong owner", "null owner", "missing image", "foreign image", "missing error", "item error", "empty name", "bad checksum", "wrong host", "http url", "credentials url", "fragment url", "query error", "recreated configuration", "replaced site"} {
		t.Run(mode, func(t *testing.T) {
			contract := imageResolverFixture()
			calls, posts, configLists, siteLists := 0, 0, 0, 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				w.Header().Set("Content-Type", "application/json")
				var result any
				switch r.URL.Path {
				case contract.ConfigurationList:
					configLists++
					item := map[string]any{"name": "example-kvm", "uid": "configuration-uid"}
					if mode == "missing uid" {
						delete(item, "uid")
					}
					if mode == "recreated configuration" && configLists > 1 {
						item["uid"] = "changed-uid"
					}
					items := []any{item}
					if mode == "missing configuration" {
						items = nil
					}
					if mode == "duplicate configuration" {
						items = append(items, item)
					}
					result = map[string]any{"items": items}
				case strings.ReplaceAll(contract.ConfigurationGet, "{site_name}", "example-kvm"):
					result = map[string]any{"spec": map[string]any{"kvm": map[string]any{}}}
					if mode == "non-kvm" {
						result = map[string]any{"spec": map[string]any{"aws": map[string]any{}}}
					}
				case contract.SiteList:
					siteLists++
					owner := map[string]any{"kind": "securemesh_site_v2", "uid": "configuration-uid"}
					item := map[string]any{"uid": "site-uid", "owner_view": owner}
					if mode == "replaced site" && siteLists > 1 {
						item["uid"] = "changed-site"
					}
					if mode == "wrong owner" {
						owner["uid"] = "foreign-uid"
					}
					if mode == "null owner" {
						item["owner_view"] = nil
					}
					items := []any{item}
					if mode == "missing owner" {
						items = nil
					}
					if mode == "duplicate owner" {
						items = append(items, item)
					}
					result = map[string]any{"items": items}
				case contract.Query:
					posts++
					if r.Method != http.MethodPost {
						t.Error("image query must POST")
					}
					var body map[string][]string
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil || len(body) != 1 || len(body["uids"]) != 1 || body["uids"][0] != "site-uid" {
						t.Error("query did not use the exact observed Site UID")
					}
					if mode == "query error" {
						w.WriteHeader(http.StatusInternalServerError)
						result = map[string]any{"message": "private-backend-diagnostic"}
						break
					}
					item := map[string]any{"download_image_link": "https://downloads.volterra.io/example.qcow2", "copy_image_name": "example.qcow2", "image_md5_sum": strings.Repeat("a", 32), "error_description": ""}
					switch mode {
					case "missing error":
						delete(item, "error_description")
					case "item error":
						item["error_description"] = "private-backend-diagnostic"
					case "empty name":
						item["copy_image_name"] = " "
					case "bad checksum":
						item["image_md5_sum"] = "malformed"
					case "wrong host":
						item["download_image_link"] = "https://example.com/example.qcow2"
					case "http url":
						item["download_image_link"] = "http://downloads.volterra.io/example.qcow2"
					case "credentials url":
						item["download_image_link"] = "https://user:secret@downloads.volterra.io/example.qcow2"
					case "fragment url":
						item["download_image_link"] = "https://downloads.volterra.io/example.qcow2#fragment"
					}
					images := map[string]any{"site-uid": item}
					if mode == "missing image" {
						delete(images, "site-uid")
					}
					if mode == "foreign image" {
						images["foreign-uid"] = item
					}
					result = map[string]any{"images": images}
				default:
					t.Error("unexpected endpoint (legacy endpoint and downloads are forbidden)")
					w.WriteHeader(404)
					return
				}
				if r.URL.Path != contract.Query && r.Method != http.MethodGet {
					t.Error("identity lookup must GET")
				}
				_ = json.NewEncoder(w).Encode(result)
			}))
			defer server.Close()
			api := NewClient(server.URL, "test-token", WithMaxRetries(0))
			got, err := api.ResolveKVMImage(context.Background(), "example-kvm", contract)
			if mode == "valid" {
				if err != nil || got.ImageName != "example.qcow2" || got.MD5 != strings.Repeat("a", 32) || calls != 7 || posts != 1 {
					t.Fatalf("valid resolution failed: calls=%d posts=%d error=%v", calls, posts, err)
				}
			} else {
				if err == nil {
					t.Fatal("invalid response accepted")
				}
				if got != (KVMImage{}) {
					t.Fatal("failure retained image metadata")
				}
				for _, secret := range []string{"private-backend-diagnostic", "configuration-uid", "site-uid", "https://", "test-token"} {
					if strings.Contains(err.Error(), secret) {
						t.Error("diagnostic leaked response data")
					}
				}
				if posts > 1 {
					t.Error("image query was replayed")
				}
			}
		})
	}
}

func TestResolveKVMImageInvalidInputMakesNoCalls(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { calls++; w.WriteHeader(500) }))
	defer server.Close()
	api := NewClient(server.URL, "test-token", WithMaxRetries(0))
	for _, name := range []string{"", "../foreign", "name?other=true", "example/other", " name", "EXAMPLE"} {
		if _, err := api.ResolveKVMImage(context.Background(), name, imageResolverFixture()); err == nil {
			t.Error("invalid input accepted")
		}
	}
	if calls != 0 {
		t.Fatalf("invalid inputs made %d network calls", calls)
	}
}
