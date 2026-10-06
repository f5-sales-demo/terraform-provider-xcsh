package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

const SwaggerMaxContent = 5 * 1024 * 1024

// swaggerLabel validates the object-store identity before any network operation.
var swaggerLabel = regexp.MustCompile(`^[a-z](?:[a-z0-9-]*[a-z0-9])?$`)
var swaggerOpenAPI = regexp.MustCompile(`^3\.[01]\.\d+$`)
var swaggerVersion = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)

// SwaggerIdentity restricts every operation to one exact object version.
func SwaggerIdentity(namespace, name, version string) (string, error) {
	for _, value := range []string{namespace, name} {
		if len(value) > 63 || !swaggerLabel.MatchString(value) {
			return "", fmt.Errorf("swagger namespace and name must be DNS labels of at most 63 characters")
		}
	}
	prefix := "/api/object_store/namespaces/" + namespace + "/stored_objects/swagger/" + name
	if version == "" {
		return prefix, nil
	}
	if len(version) > 1024 || !swaggerVersion.MatchString(version) || strings.EqualFold(version, "latest") {
		return "", fmt.Errorf("swagger version must be exact; latest is prohibited")
	}
	return prefix + "/" + version, nil
}

// swaggerJSONValue rejects duplicate keys before accepting content or API envelopes.
func swaggerJSONValue(decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, ok := token.(json.Delim)
	if !ok {
		return nil
	}
	switch delimiter {
	case '{':
		seen := map[string]bool{}
		for decoder.More() {
			key, err := decoder.Token()
			if err != nil {
				return err
			}
			name, ok := key.(string)
			if !ok || seen[name] {
				return fmt.Errorf("duplicate or invalid JSON key")
			}
			seen[name] = true
			if err := swaggerJSONValue(decoder); err != nil {
				return err
			}
		}
	case '[':
		for decoder.More() {
			if err := swaggerJSONValue(decoder); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("invalid JSON delimiter")
	}
	_, err = decoder.Token()
	return err
}
func swaggerStrictJSON(content []byte) error {
	decoder := json.NewDecoder(strings.NewReader(string(content)))
	if err := swaggerJSONValue(decoder); err != nil {
		return err
	}
	if _, err := decoder.Token(); err != io.EOF {
		return fmt.Errorf("json must contain exactly one document")
	}
	return nil
}

// ValidateSwaggerContent preserves bytes while checking the supported JSON profile.
func ValidateSwaggerContent(content string) error {
	if len(content) == 0 || len(content) > SwaggerMaxContent || !utf8.ValidString(content) {
		return fmt.Errorf("swagger content must be nonempty UTF-8 of at most 5 MiB")
	}
	if err := swaggerStrictJSON([]byte(content)); err != nil {
		return fmt.Errorf("swagger content is invalid JSON: %w", err)
	}
	var document map[string]json.RawMessage
	if err := json.Unmarshal([]byte(content), &document); err != nil || document == nil {
		return fmt.Errorf("swagger content must be a JSON object")
	}
	var openapi, swagger string
	_ = json.Unmarshal(document["openapi"], &openapi)
	_ = json.Unmarshal(document["swagger"], &swagger)
	validProfile := (swaggerOpenAPI.MatchString(openapi) && swagger == "") || (swagger == "2.0" && openapi == "")
	if !validProfile {
		return fmt.Errorf("swagger content requires OpenAPI 3.0/3.1 or Swagger 2.0")
	}
	var info struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	}
	var paths map[string]json.RawMessage
	if json.Unmarshal(document["info"], &info) != nil || info.Title == "" || info.Version == "" || json.Unmarshal(document["paths"], &paths) != nil || paths == nil {
		return fmt.Errorf("swagger info.title, info.version and paths are required")
	}
	for path, raw := range paths {
		var endpoint map[string]json.RawMessage
		if !strings.HasPrefix(path, "/") || json.Unmarshal(raw, &endpoint) != nil || endpoint == nil {
			return fmt.Errorf("swagger paths must contain absolute API paths and object values")
		}
	}
	return nil
}

type SwaggerObject struct {
	Metadata struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
		Version   string `json:"version"`
	} `json:"metadata"`
	StringValue   *string `json:"string_value,omitempty"`
	BytesValue    *string `json:"bytes_value,omitempty"`
	PresignedURL  *string `json:"presigned_url,omitempty"`
	ContentFormat string  `json:"content_format"`
}

func (o *SwaggerObject) Content(namespace, name, version string) (string, error) {
	if o.Metadata.Namespace != namespace || o.Metadata.Name != name || o.Metadata.Version != version {
		return "", fmt.Errorf("swagger exact-version metadata identity mismatch")
	}
	if _, err := SwaggerIdentity(namespace, name, version); err != nil || version == "" {
		return "", fmt.Errorf("swagger response has invalid version")
	}
	if o.PresignedURL != nil || (o.StringValue == nil) == (o.BytesValue == nil) {
		return "", fmt.Errorf("swagger content is missing, ambiguous, or presigned")
	}
	var content string
	if o.StringValue != nil {
		content = *o.StringValue
	} else {
		raw, err := base64.StdEncoding.Strict().DecodeString(*o.BytesValue)
		if err != nil {
			return "", fmt.Errorf("swagger content has invalid base64")
		}
		content = string(raw)
	}
	if err := ValidateSwaggerContent(content); err != nil {
		return "", err
	}
	return content, nil
}

// SwaggerVersions is used only to reject out-of-state creation conflicts, never to select latest.
func (c *Client) SwaggerVersions(ctx context.Context, namespace, name string) ([]string, error) {
	prefix, err := SwaggerIdentity(namespace, name, "")
	if err != nil {
		return nil, err
	}
	var listing struct {
		Items []struct {
			Name     string `json:"name"`
			Versions []struct {
				Version string `json:"version"`
			} `json:"versions"`
		} `json:"items"`
	}
	query := url.Values{"name": {name}, "query_type": {"EXACT_MATCH"}, "latest_version_only": {"false"}}
	if err = c.GetOnce(ctx, strings.TrimSuffix(prefix, "/"+name)+"?"+query.Encode(), &listing); err != nil {
		return nil, err
	}
	if listing.Items == nil || len(listing.Items) > 1 {
		return nil, fmt.Errorf("swagger name-filtered listing is missing or ambiguous")
	}
	versions := []string{}
	for _, item := range listing.Items {
		if item.Name != name && item.Name != namespace+"/"+name {
			return nil, fmt.Errorf("swagger listing identity mismatch")
		}
		if len(item.Versions) == 0 || len(item.Versions) > 64 {
			return nil, fmt.Errorf("swagger listing version count invalid")
		}
		seen := map[string]bool{}
		for _, v := range item.Versions {
			if _, err := SwaggerIdentity(namespace, name, v.Version); err != nil || v.Version == "" || seen[v.Version] {
				return nil, fmt.Errorf("swagger listing versions invalid")
			}
			seen[v.Version] = true
			versions = append(versions, v.Version)
		}
	}
	return versions, nil
}

func (c *Client) CreateSwaggerObject(ctx context.Context, namespace, name, content string) (*SwaggerObject, error) {
	prefix, err := SwaggerIdentity(namespace, name, "")
	if err != nil {
		return nil, err
	}
	if err = ValidateSwaggerContent(content); err != nil {
		return nil, err
	}
	var result SwaggerObject
	// Version-issuing PUT cannot be replayed, even though PUT is normally idempotent.
	err = c.PutOnce(ctx, prefix, map[string]string{"namespace": namespace, "name": name, "object_type": "swagger", "string_value": content, "content_format": "json"}, &result)
	if err != nil {
		return nil, err
	}
	if _, err := SwaggerIdentity(namespace, name, result.Metadata.Version); err != nil || result.Metadata.Version == "" || result.Metadata.Name != name || result.Metadata.Namespace != namespace {
		return nil, fmt.Errorf("swagger upload did not return an exact matching identity")
	}
	return &result, nil
}

func (c *Client) GetSwaggerObject(ctx context.Context, namespace, name, version string) (*SwaggerObject, error) {
	p, err := SwaggerIdentity(namespace, name, version)
	if err != nil || version == "" {
		return nil, fmt.Errorf("swagger GET requires an exact version")
	}
	var result SwaggerObject
	if err = c.GetOnce(ctx, p, &result); err != nil {
		return nil, err
	}
	if _, err = result.Content(namespace, name, version); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteSwaggerObject(ctx context.Context, namespace, name, version string) error {
	p, err := SwaggerIdentity(namespace, name, version)
	if err != nil || version == "" {
		return fmt.Errorf("swagger DELETE requires an exact owned version")
	}
	_, err = c.doRequestWithRetry(ctx, http.MethodDelete, p, nil, false)
	return err
}
