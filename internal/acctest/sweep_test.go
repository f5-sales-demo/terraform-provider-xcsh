// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestIsTestResourceAcceptsOnlyCanonicalPrefix(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "tf-acc-test-resource", want: true},
		{name: "tf-test-resource", want: false},
		{name: "production-resource", want: false},
		{name: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isTestResource(test.name); got != test.want {
				t.Fatalf("isTestResource(%q) = %t, want %t", test.name, got, test.want)
			}
		})
	}
}

func TestRandomNameEnforcesCanonicalSweepPrefix(t *testing.T) {
	for _, prefix := range []string{"tf-acc-test", "tf-acc-test-healthcheck"} {
		name := RandomName(prefix)
		if !isTestResource(name) {
			t.Fatalf("RandomName(%q) produced unsweepable name %q", prefix, name)
		}
	}

	deferred := func() (recovered any) {
		defer func() { recovered = recover() }()
		_ = RandomName("untracked")
		return nil
	}()
	if deferred == nil {
		t.Fatal("RandomName accepted a prefix that the sweeper cannot recognize")
	}
}

// TestMain handles setup and teardown for the acceptance test suite,
// including test resource sweepers.
//
// Run sweepers with:
//
//	TF_ACC=1 go test ./internal/acctest -v -sweep=all -timeout 30m
//
// Or to sweep specific resources:
//
//	TF_ACC=1 go test ./internal/acctest -v -sweep=xcsh_namespace -timeout 30m
func TestMain(m *testing.M) {
	resource.TestMain(m)
}
