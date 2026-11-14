//go:build linux

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"testing"
)

func TestHeldOriginalBootstrapRefreshKeepsIdentityAndSuppressesCredentials(t *testing.T) {
	c, r, _, _ := bootstrapFixture(t)
	held, err := c.OriginalBootstrapTokenState(r.CreationIdentitySHA256)
	if err != nil {
		t.Fatal(err)
	}
	data := TokenResourceModel{ID: types.StringValue("example-token"), Uid: types.StringValue("synthetic.jwt.content"), Content: types.StringValue("synthetic.jwt.content")}
	if err := suppressHeldTokenBootstrap(&data, held, r.CreationIdentitySHA256); err != nil {
		t.Fatal(err)
	}
	if !data.Content.IsNull() || !data.Uid.IsNull() || data.ID.ValueString() != "example-token" {
		t.Fatal("held refresh exposed bootstrap or lost cleanup identity")
	}
	if err := suppressHeldTokenBootstrap(&data, held, "foreign"); err == nil {
		t.Fatal("foreign private marker accepted")
	}
}
func TestReleasedOriginalBootstrapMarkerNeedsDurableGrant(t *testing.T) {
	c, r, key, auth := bootstrapFixture(t)
	initial, err := c.OriginalBootstrapTokenState(r.CreationIdentitySHA256)
	if err != nil {
		t.Fatal(err)
	}
	writeBootstrapGrant(t, r, key, auth, "synthetic.jwt.content", nil)
	if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err != nil {
		t.Fatal(err)
	}
	released, err := c.OriginalBootstrapTokenState(r.CreationIdentitySHA256)
	if err != nil {
		t.Fatal(err)
	}
	if string(initial) == string(released) {
		t.Fatal("successful durable consumption did not update private state")
	}
	data := TokenResourceModel{Uid: types.StringValue("synthetic.jwt.content"), Content: types.StringValue("synthetic.jwt.content")}
	if err := suppressHeldTokenBootstrap(&data, released, r.CreationIdentitySHA256); err != nil {
		t.Fatal(err)
	}
	if data.Content.ValueString() != "synthetic.jwt.content" {
		t.Fatal("released original state cannot reconcile unchanged configuration")
	}
}
func TestLegacyTokenRefreshHasNoBootstrapMarker(t *testing.T) {
	data := TokenResourceModel{Content: types.StringValue("legacy")}
	if err := suppressHeldTokenBootstrap(&data, nil, ""); err != nil || data.Content.ValueString() != "legacy" {
		t.Fatal("legacy refresh changed")
	}
}
