package provider

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// creationIdentityReceipt records only the immutable UID in the original POST response.
// Legacy and imported objects have no original creation receipt.
func creationIdentityReceipt(uid string) types.String {
	if uid == "" {
		return types.StringNull()
	}
	return types.StringValue(fmt.Sprintf("sha256:%x", sha256.Sum256([]byte(uid))))
}

func verifyCreationIdentity(original types.String, uid string) error {
	if original.IsNull() {
		return nil
	}
	if original.IsUnknown() || uid == "" || !original.Equal(creationIdentityReceipt(uid)) {
		return errors.New("immutable identity differs from the original creation receipt; reconcile ownership before mutation")
	}
	return nil
}
