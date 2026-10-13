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

// Qualification-only probe: every write uses the run's ownership transport.
func TestAccAPIDefinitionDiagnostic(t *testing.T) {
	acctest.SkipIfNotAccTest(t)
	acctest.EnsureMockModeConfigured()
	if acctest.IsMockMode() {
		t.Setenv("XCSH_ACC_PRIVATE", t.TempDir())
	}
	c, err := acctest.GetTestClient()
	if err != nil {
		t.Fatal("diagnostic client unavailable")
	}
	ctx := context.Background()
	for _, variant := range []string{"empty", "relative", "absolute"} {
		name := acctest.RandomName("tf-acc-test-api-probe")
		spec := map[string]interface{}{}
		if variant != "empty" {
			content := `{"openapi":"3.0.3","info":{"title":"Acceptance API","version":"1.0.0"},"servers":[{"url":"https://example.com"}],"paths":{"/health":{"get":{"responses":{"200":{"description":"OK"}}}}}}`
			object, err := c.CreateSwaggerObject(ctx, "shared", name+"-spec", content)
			if err != nil {
				t.Fatal("diagnostic Swagger creation failed")
			}
			path, err := client.SwaggerIdentity("shared", name+"-spec", object.Metadata.Version)
			if err != nil {
				t.Fatal("diagnostic identity invalid")
			}
			if variant == "absolute" {
				path = strings.TrimRight(os.Getenv("XCSH_API_URL"), "/") + path
			}
			spec["swagger_specs"] = []string{path}
		}
		_, err := c.CreateAPIDefinition(ctx, &client.APIDefinition{Metadata: client.Metadata{Name: name, Namespace: "shared"}, Spec: spec})
		if err != nil {
			// Raw error remains inside the private workspace and never reaches artifacts.
			message := err.Error()
			var apiErr *xcsherrors.XCSHError
			if errors.As(err, &apiErr) {
				message = apiErr.Message
				if raw, ok := apiErr.Details["raw_response"].(string); ok {
					message = raw
				}
			}
			data, _ := json.Marshal(map[string]string{"variant": variant, "error": message})
			if os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/api-probe-"+variant+".json", data, 0600) != nil {
				t.Fatal("private diagnostic persistence failed")
			}
			t.Logf("[API_PROBE:%s rejected]", variant)
		} else {
			data, _ := json.Marshal(map[string]interface{}{"variant": variant, "accepted": true})
			if os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/api-probe-"+variant+".json", data, 0600) != nil {
				t.Fatal("private diagnostic persistence failed")
			}
			t.Logf("[API_PROBE:%s accepted]", variant)
		}
	}
}
