// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"os"
	"path/filepath"
	"testing"
)

func ProvisionDisposablePrerequisites(ctx context.Context) error {
	c := client.NewClient(os.Getenv(EnvXCSHURL), os.Getenv(EnvXCSHToken))
	if err := attachOwnership(c); err != nil {
		return err
	}
	if err := provisionQuotaBoundCapabilities(ctx, c); err != nil {
		return err
	}
	name := CompactRandomName(20)
	_, err := c.CreateSegment(ctx, &client.Segment{Metadata: client.Metadata{Name: name, Namespace: "system"}, Spec: map[string]interface{}{}})
	if err != nil {
		if classifyCapabilityProbe(err) == capabilityDenied {
			return os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/segment-capability.json", []byte(`{"version":1,"type":"xcsh_segment","operation":"POST","status":403}`), 0600)
		}
		return err
	}
	if err := provisionMeshSelector(ctx, c); err != nil {
		return err
	}
	return os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/prerequisites.env", []byte(fmt.Sprintf("XCSH_TEST_SEGMENT_NAME=%s\n", name)), 0600)
}
func RequireDisposableSegment(t *testing.T) {
	t.Helper()
	path := os.Getenv("XCSH_ACC_PRIVATE") + "/segment-capability.json"
	data, err := os.ReadFile(path)
	if err == nil && string(data) == `{"version":1,"type":"xcsh_segment","operation":"POST","status":403}` {
		t.Log("[PREREQUISITE:operation_denied method=POST type=xcsh_segment status=403]")
		t.Skip("disposable segment creation denied")
	}
	t.Fatal("disposable segment prerequisite missing")
}

// provisionMeshSelector copies only the selector expression from verified
// existing lab infrastructure. The disposable selector is journal-owned.
func provisionMeshSelector(ctx context.Context, c *client.Client) error {
	ref, err := discoverMeshVirtualSite(ctx, c)
	if err != nil {
		if err.Error() == "existing infrastructure absent" || classifyCapabilityProbe(err) == capabilityDenied {
			return nil
		}
		return errors.New("mesh prerequisite discovery failed")
	}
	source, err := c.GetVirtualSite(ctx, ref.Namespace, ref.Name)
	if err != nil {
		return err
	}
	selector, ok := source.Spec["site_selector"].(map[string]interface{})
	if !ok {
		return errors.New("mesh selector malformed")
	}
	name := CompactRandomName(20)
	_, err = c.CreateVirtualSite(ctx, &client.VirtualSite{Metadata: client.Metadata{Name: name, Namespace: "system"}, Spec: map[string]interface{}{"site_type": "CUSTOMER_EDGE", "site_selector": selector}})
	if err != nil {
		return err
	}
	owned := readOnlySiteReference{Name: name, Namespace: "system"}
	count, err := verifyMeshSelectorExpansion(ctx, c, owned)
	if err != nil || count == 0 {
		return errors.New("disposable selector expansion failed")
	}
	owned.SelectedSites = count
	data, err := json.Marshal(owned)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "mesh-owned-reference.json"), data, 0600)
}
