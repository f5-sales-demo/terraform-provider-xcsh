package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestCreationIdentityReceipt(t *testing.T) {
	original := creationIdentityReceipt("original-immutable-uid")
	if original.IsNull() || original.ValueString() == "original-immutable-uid" || len(original.ValueString()) != 71 {
		t.Fatal("missing sanitized receipt")
	}
	if err := verifyCreationIdentity(original, "original-immutable-uid"); err != nil {
		t.Fatal(err)
	}
	for _, uid := range []string{"replacement", ""} {
		if verifyCreationIdentity(original, uid) == nil {
			t.Fatal("changed original identity admitted")
		}
	}
	if !creationIdentityReceipt("").IsNull() {
		t.Fatal("absent POST UID gained ownership")
	}
	if err := verifyCreationIdentity(types.StringNull(), "live-legacy-uid"); err != nil {
		t.Fatal(err)
	}
}
