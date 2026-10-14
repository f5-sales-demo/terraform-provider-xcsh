// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package mocks

import (
	"context"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"testing"
)

func TestSwaggerMockExactVersionLifecycle(t *testing.T) {
	s := NewServer()
	defer s.Close()
	c := client.NewClient(s.URL(), "test")
	content := `{"openapi":"3.0.3","info":{"title":"Fixture","version":"1"},"paths":{}}`
	versions, err := c.SwaggerVersions(context.Background(), "system", "fixture")
	if err != nil || len(versions) != 0 {
		t.Fatalf("empty inventory %v %v", versions, err)
	}
	object, err := c.CreateSwaggerObject(context.Background(), "system", "fixture", content)
	if err != nil {
		t.Fatal(err)
	}
	exact, err := c.GetSwaggerObject(context.Background(), "system", "fixture", object.Metadata.Version)
	if err != nil {
		t.Fatal(err)
	}
	got, err := exact.Content("system", "fixture", object.Metadata.Version)
	if err != nil || got != content {
		t.Fatal("content changed")
	}
	if err = c.DeleteSwaggerObject(context.Background(), "system", "fixture", object.Metadata.Version); err != nil {
		t.Fatal(err)
	}
	versions, err = c.SwaggerVersions(context.Background(), "system", "fixture")
	if err != nil || len(versions) != 0 {
		t.Fatal("version remains")
	}
}
