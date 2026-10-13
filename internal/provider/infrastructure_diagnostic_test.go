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
	"strings"
	"testing"
)

func TestAccReadOnlyInfrastructureDiagnostic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.EnsureMockModeConfigured()
	if acctest.IsMockMode() {
		t.Setenv("XCSH_ACC_PRIVATE", t.TempDir())
	}
	c, err := acctest.GetTestClient()
	if err != nil {
		t.Fatal("infrastructure diagnostic client unavailable")
	}
	result := []map[string]interface{}{}
	for _, collection := range []string{"sites", "virtual_sites"} {
		row := map[string]interface{}{"collection": collection}
		path := "/api/config/namespaces/system/" + collection
		var raw map[string]json.RawMessage
		err := c.Get(context.Background(), path, &raw)
		row["read_success"] = err == nil
		if err != nil {
			var e *xcsherrors.XCSHError
			if errors.As(err, &e) {
				row["status"] = e.StatusCode
			}
			result = append(result, row)
			continue
		}
		var items []struct {
			Name      string          `json:"name"`
			Metadata  client.Metadata `json:"metadata"`
			Namespace string          `json:"namespace"`
		}
		row["items_present"] = raw["items"] != nil
		err = json.Unmarshal(raw["items"], &items)
		row["items_valid"] = err == nil
		row["count"] = len(items)
		candidates := []map[string]interface{}{}
		for _, item := range items {
			name := item.Name
			if name == "" {
				name = item.Metadata.Name
			}
			candidate := map[string]interface{}{"name_present": name != "", "name_has_separator": strings.Contains(name, "/"), "test_owned": strings.HasPrefix(name, acctest.TestResourcePrefix)}
			if name == "" || strings.ContainsAny(name, "/?# \t\r\n") {
				candidates = append(candidates, candidate)
				continue
			}
			var response struct {
				Metadata       client.Metadata `json:"metadata"`
				SystemMetadata struct {
					UID string `json:"uid"`
				} `json:"system_metadata"`
			}
			err := c.Get(context.Background(), path+"/"+name, &response)
			candidate["read_success"] = err == nil
			if err != nil {
				var e *xcsherrors.XCSHError
				if errors.As(err, &e) {
					candidate["status"] = e.StatusCode
				}
			}
			candidate["name_matches"] = response.Metadata.Name == name
			candidate["namespace_system"] = response.Metadata.Namespace == "system"
			candidate["namespace_empty"] = response.Metadata.Namespace == ""
			candidate["uid_present"] = response.Metadata.UID != "" || response.SystemMetadata.UID != ""
			candidates = append(candidates, candidate)
		}
		row["candidates"] = candidates
		result = append(result, row)
	}
	data, err := json.MarshalIndent(map[string]interface{}{"version": 1, "collections": result}, "", "  ")
	if err != nil {
		t.Fatal("infrastructure diagnostic encoding failed")
	}
	if os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/infrastructure-diagnostic.json", data, 0600) != nil {
		t.Fatal("private diagnostic persistence failed")
	}
}
