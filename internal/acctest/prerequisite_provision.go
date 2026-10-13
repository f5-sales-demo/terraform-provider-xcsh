// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"os"
	"testing"
)

func ProvisionDisposablePrerequisites(ctx context.Context) error {
	c := client.NewClient(os.Getenv(EnvXCSHURL), os.Getenv(EnvXCSHToken))
	if err := attachOwnership(c); err != nil {
		return err
	}
	name := RandomName("tf-acc-test-segment")
	_, err := c.CreateSegment(ctx, &client.Segment{Metadata: client.Metadata{Name: name, Namespace: "system"}, Spec: map[string]interface{}{}})
	if err != nil {
		if classifyCapabilityProbe(err) == capabilityDenied {
			return os.WriteFile(os.Getenv("XCSH_ACC_PRIVATE")+"/segment-capability.json", []byte(`{"version":1,"type":"xcsh_segment","operation":"POST","status":403}`), 0600)
		}
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
