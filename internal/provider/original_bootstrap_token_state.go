package provider

import (
	"encoding/hex"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

type originalBootstrapTokenState struct {
	SchemaVersion          int    `json:"schema_version"`
	Kind                   string `json:"kind"`
	State                  string `json:"state"`
	CreationIdentitySHA256 string `json:"creation_identity_sha256"`
	ScopeSHA256            string `json:"scope_sha256"`
	ScopeFileSHA256        string `json:"scope_file_sha256"`
	Nonce                  string `json:"nonce"`
}

// OriginalBootstrapTokenState records a held token only after the POST has
// created its immutable identity. Released state requires durable grant consumption.
func (c *OriginalCreationCapture) OriginalBootstrapTokenState(identity string) ([]byte, error) {
	if c == nil {
		return nil, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.bootstrap == nil {
		return nil, nil
	}
	identity = strings.TrimPrefix(identity, "sha256:")
	if !captureHash(identity) || c.receipt.Target.TypeName != "xcsh_token" {
		return nil, bootstrapError()
	}
	if c.creationReceipt != nil && c.creationReceipt.CreationIdentitySHA256 != identity {
		return nil, bootstrapError()
	}
	state := "held"
	if c.bootstrapReleased {
		state = "released"
	}
	return json.Marshal(originalBootstrapTokenState{1, "original-bootstrap-token", state, identity, c.receipt.ScopeSHA256, c.receipt.ScopeFileSHA256, c.receipt.Nonce})
}

// A failed Create can retain identity but must not expose bootstrap on a later
// refresh. Metadata/ID remain available for owned destroy without registration.
func suppressHeldTokenBootstrap(data *TokenResourceModel, raw []byte, identity string) error {
	if len(raw) == 0 {
		return nil
	}
	var marker originalBootstrapTokenState
	if err := bootstrapDecode(raw, &marker); err != nil {
		return err
	}
	if marker.SchemaVersion != 1 || marker.Kind != "original-bootstrap-token" ||
		(marker.State != "held" && marker.State != "released") ||
		marker.CreationIdentitySHA256 != strings.TrimPrefix(identity, "sha256:") ||
		!captureHash(marker.CreationIdentitySHA256) || !captureHash(marker.ScopeSHA256) || !captureHash(marker.ScopeFileSHA256) ||
		len(marker.Nonce) != 32 {
		return bootstrapError()
	}
	nonce, err := hex.DecodeString(marker.Nonce)
	if err != nil || len(nonce) != 16 || strings.ToLower(marker.Nonce) != marker.Nonce {
		return bootstrapError()
	}
	if marker.State == "held" {
		data.Uid = types.StringNull()
		data.Content = types.StringNull()
	}
	return nil
}
