// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"os"
	"testing"
)

func TestAccMeshDiagnostic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.EnsureMockModeConfigured()
	if acctest.IsMockMode() {
		t.Setenv("XCSH_ACC_PRIVATE", t.TempDir())
	}
	acctest.RequireReadOnlyVirtualSite(t)
	c, err := acctest.GetTestClient()
	if err != nil {
		t.Fatal("mesh probe client unavailable")
	}
	ref := map[string]interface{}{"name": acctest.ReadOnlyVirtualSiteName(), "namespace": "system"}
	results := []map[string]interface{}{}
	for _, variant := range []string{"base", "markers", "data_plane"} {
		spec := map[string]interface{}{"virtual_site": []map[string]interface{}{ref}, "full_mesh": map[string]interface{}{"control_and_data_plane_mesh": map[string]interface{}{}}}
		if variant == "markers" {
			spec["bfd_disabled"] = map[string]interface{}{}
			spec["disable_re_fallback"] = map[string]interface{}{}
		}
		if variant == "data_plane" {
			spec["full_mesh"] = map[string]interface{}{"data_plane_mesh": map[string]interface{}{}}
		}
		_, err := c.CreateSiteMeshGroup(context.Background(), &client.SiteMeshGroup{Metadata: client.Metadata{Name: acctest.RandomName("tf-acc-test-meshprobe"), Namespace: "system"}, Spec: spec})
		row := map[string]interface{}{"variant": variant, "accepted": err == nil}
		if err != nil {
			message := err.Error()
			var e *xcsherrors.XCSHError
			if errors.As(err, &e) {
				message = e.Message
				if raw, ok := e.Details["raw_response"].(string); ok {
					message = raw
				}
			}
			row["error"] = message
		}
		results = append(results, row)
	}
	data, err := json.Marshal(results)
	if err != nil {
		t.Fatal("mesh probe encoding failed")
	}
	if os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/mesh-probe.json", data, 0600) != nil {
		t.Fatal("mesh probe persistence failed")
	}
}
