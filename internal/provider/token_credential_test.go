package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestTokenCredentialNormalUsesSystemMetadataUID(t *testing.T) {
	t.Parallel()
	resource := &client.Token{
		Spec:           map[string]interface{}{"type": float64(tokenTypeNormal)},
		SystemMetadata: &client.TokenSystemMetadata{UID: "normal-credential"},
	}

	credential, content, err := tokenCredential(resource, tokenTypeJWT)
	if err != nil {
		t.Fatalf("tokenCredential returned error: %v", err)
	}
	if credential != "normal-credential" {
		t.Fatalf("credential = %q, want normal credential", credential)
	}
	if content != "" {
		t.Fatalf("content = %q, want empty for NORMAL token", content)
	}
}

func TestTokenCredentialJWTUsesSpecContent(t *testing.T) {
	t.Parallel()
	resource := &client.Token{
		Spec: map[string]interface{}{
			"type":    float64(tokenTypeJWT),
			"content": "jwt-credential",
		},
		SystemMetadata: &client.TokenSystemMetadata{UID: "object-uid"},
	}

	credential, content, err := tokenCredential(resource, tokenTypeNormal)
	if err != nil {
		t.Fatalf("tokenCredential returned error: %v", err)
	}
	if credential != "jwt-credential" || content != "jwt-credential" {
		t.Fatalf("JWT credential selection did not use spec.content")
	}
}

func TestTokenCredentialUsesConfiguredJWTKindWhenResponseOmitsType(t *testing.T) {
	t.Parallel()
	resource := &client.Token{
		Spec: map[string]interface{}{"content": "jwt-credential"},
	}

	credential, _, err := tokenCredential(resource, tokenTypeJWT)
	if err != nil {
		t.Fatalf("tokenCredential returned error: %v", err)
	}
	if credential != "jwt-credential" {
		t.Fatal("configured JWT fallback did not select spec.content")
	}
}

func TestTokenCredentialFailsClosedWithoutJWTContent(t *testing.T) {
	t.Parallel()
	resource := &client.Token{
		Spec:           map[string]interface{}{"type": float64(tokenTypeJWT)},
		SystemMetadata: &client.TokenSystemMetadata{UID: "must-not-be-used"},
	}

	_, _, err := tokenCredential(resource, tokenTypeNormal)
	if err == nil || !strings.Contains(err.Error(), "missing spec.content") {
		t.Fatalf("tokenCredential error = %v, want missing content diagnostic", err)
	}
	if strings.Contains(err.Error(), "must-not-be-used") {
		t.Fatal("token credential diagnostic exposed a credential")
	}
}

func TestTokenCredentialRejectsUnsupportedKindWithoutEchoingValues(t *testing.T) {
	t.Parallel()
	resource := &client.Token{
		Spec: map[string]interface{}{
			"type":    float64(7),
			"content": "must-not-appear",
		},
	}

	_, _, err := tokenCredential(resource, tokenTypeNormal)
	if err == nil || !strings.Contains(err.Error(), "unsupported token type") {
		t.Fatalf("tokenCredential error = %v, want unsupported type diagnostic", err)
	}
	if strings.Contains(err.Error(), "must-not-appear") {
		t.Fatal("token credential diagnostic exposed credential content")
	}
}

func TestTokenCredentialRejectsMalformedObservedKindInsteadOfFallingBack(t *testing.T) {
	t.Parallel()
	for name, kind := range map[string]interface{}{
		"string":     "1",
		"fractional": 1.5,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			resource := &client.Token{
				Spec: map[string]interface{}{
					"type":    kind,
					"content": "must-not-appear",
				},
			}

			_, _, err := tokenCredential(resource, tokenTypeJWT)
			if err == nil || !strings.Contains(err.Error(), "unsupported token type") {
				t.Fatalf("tokenCredential error = %v, want unsupported type diagnostic", err)
			}
			if strings.Contains(err.Error(), "must-not-appear") {
				t.Fatal("token credential diagnostic exposed credential content")
			}
		})
	}
}

func TestPopulateTokenCredentialStateSetsSensitiveOutputs(t *testing.T) {
	t.Parallel()
	data := TokenResourceModel{Type: types.Int64Value(tokenTypeJWT)}
	resource := &client.Token{
		Spec: map[string]interface{}{"content": "jwt-credential"},
	}

	if err := populateTokenCredentialState(&data, resource); err != nil {
		t.Fatalf("populateTokenCredentialState returned error: %v", err)
	}
	if data.Uid.ValueString() != "jwt-credential" {
		t.Fatal("uid did not expose the selected JWT credential")
	}
	if data.Content.ValueString() != "jwt-credential" {
		t.Fatal("content did not expose the JWT response field")
	}
}
