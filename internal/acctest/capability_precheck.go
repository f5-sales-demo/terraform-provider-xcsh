// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

var capabilityProbeCache sync.Map
var capabilityProbeMu sync.Mutex

type capabilityDecision int

const (
	capabilityInconclusive capabilityDecision = iota
	capabilityAllowed
	capabilityDenied
)

func classifyCapabilityProbe(err error) capabilityDecision {
	if err == nil {
		return capabilityAllowed
	}
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) {
		return capabilityInconclusive
	}
	if apiErr.StatusCode == http.StatusForbidden {
		return capabilityDenied
	}
	return capabilityInconclusive
}
func executeCapabilityProbe(ctx context.Context, c interface {
	Post(context.Context, string, interface{}, interface{}) error
}, path string) error {
	var response map[string]interface{}
	// The invalid name cannot create an object. Only an explicit 403 is a skip;
	// validation errors and successful reads never prove mutation support.
	return c.Post(ctx, path, map[string]interface{}{"metadata": map[string]interface{}{"name": "-"}}, &response)
}
func precheckLiveCapability(t *testing.T) {
	t.Helper()
	if schemaPath := os.Getenv("XCSH_ACC_SCHEMA"); schemaPath != "" {
		schema, err := LoadFixtureSchema(schemaPath)
		if err != nil {
			t.Fatal("installed fixture schema unavailable")
		}
		for _, name := range acceptanceTestFixtures[t.Name()] {
			if len(schema.Validate(name, acceptanceFixtures[name]())) > 0 {
				t.Fatal("positive fixture fails installed schema validation")
			}
		}
	}
	requirements, err := TestOperationRequirements(t.Name())
	if err != nil {
		t.Fatalf("generated operation inventory missing for this test")
	}
	c, err := GetTestClient()
	if err != nil {
		t.Fatal("capability client unavailable")
	}
	capabilityProbeMu.Lock()
	defer capabilityProbeMu.Unlock()
	for _, requirement := range requirements {
		switch requirement.Type {
		case "xcsh_aws_tgw_site", "xcsh_nfv_service":
			t.Logf("[PREREQUISITE:cloud_appliance_out_of_scope method=POST type=%s]", requirement.Type)
			t.Skip("new cloud appliance provisioning is excluded from this run")
		}

		// Resolve references within the disposable namespace only after the namespace
		// create operation has been checked. Never guess mutation rights from GET.
		path := requirement.Path
		if requirement.Namespace == "disposable" {
			// The namespace's POST denial is sufficient. Child permissions are tested by
			// the actual operation once the prerequisite exists.
			continue
		}
		if requirement.Method != "POST" {
			continue
		}
		key := requirement.Method + " " + path
		var decision capabilityDecision
		if cached, ok := capabilityProbeCache.Load(key); ok {
			decision = cached.(capabilityDecision)
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			err := executeCapabilityProbe(ctx, c, path)
			cancel()
			decision = classifyCapabilityProbe(err)
			if decision == capabilityDenied {
				capabilityProbeCache.Store(key, decision)
			}
		}
		if decision == capabilityDenied {
			t.Logf("[PREREQUISITE:operation_denied method=POST type=%s status=403]", requirement.Type)
			t.Skip("operation-specific prerequisite returned HTTP 403")
		}
		t.Logf("[PREREQUISITE:tested method=%s type=%s]", requirement.Method, requirement.Type)
	}
}
