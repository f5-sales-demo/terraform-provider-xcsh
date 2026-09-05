package provider

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

const (
	tokenTypeNormal int64 = 0
	tokenTypeJWT    int64 = 1
)

func tokenKindFromSpec(spec map[string]interface{}) (int64, bool) {
	value, ok := spec["type"]
	if !ok {
		return 0, false
	}

	switch typed := value.(type) {
	case float64:
		if typed == float64(tokenTypeNormal) {
			return tokenTypeNormal, true
		}
		if typed == float64(tokenTypeJWT) {
			return tokenTypeJWT, true
		}
		return -1, true
	case int64:
		return typed, true
	case int:
		return int64(typed), true
	case json.Number:
		parsed, err := typed.Int64()
		if err != nil {
			return -1, true
		}
		return parsed, true
	default:
		return -1, true
	}
}

func tokenCredential(resource *client.Token, fallbackKind int64) (credential string, content string, err error) {
	if resource == nil {
		return "", "", errors.New("token response is missing")
	}

	kind := fallbackKind
	if observed, ok := tokenKindFromSpec(resource.Spec); ok {
		kind = observed
	}
	if kind != tokenTypeNormal && kind != tokenTypeJWT {
		return "", "", errors.New("token response contains an unsupported token type")
	}

	if kind == tokenTypeJWT {
		content, _ = resource.Spec["content"].(string)
		if strings.TrimSpace(content) == "" {
			return "", "", errors.New("JWT token response is missing spec.content")
		}
		return content, content, nil
	}

	if resource.SystemMetadata == nil || strings.TrimSpace(resource.SystemMetadata.UID) == "" {
		return "", "", errors.New("NORMAL token response is missing system_metadata.uid")
	}
	return resource.SystemMetadata.UID, "", nil
}

func populateTokenCredentialState(data *TokenResourceModel, resource *client.Token) error {
	fallbackKind := tokenTypeNormal
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		fallbackKind = data.Type.ValueInt64()
	}
	credential, content, err := tokenCredential(resource, fallbackKind)
	if err != nil {
		return err
	}
	data.Uid = types.StringValue(credential)
	if content == "" {
		data.Content = types.StringNull()
	} else {
		data.Content = types.StringValue(content)
	}
	return nil
}
