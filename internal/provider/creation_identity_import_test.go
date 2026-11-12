// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package provider_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Import has no original POST response and must not fabricate its receipt.
func checkImportedWithoutCreationReceipt(states []*terraform.InstanceState) error {
	if len(states) != 1 {
		return fmt.Errorf("expected one imported instance, got %d", len(states))
	}
	if states[0].Attributes["creation_identity_sha256"] != "" {
		return fmt.Errorf("import fabricated an original-creation receipt")
	}
	return nil
}
