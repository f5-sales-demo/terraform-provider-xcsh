// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"io"
	"net/http"
	"strings"
)

func publicContentDigest(content string) string {
	digest := sha256.Sum256([]byte(content))
	return hex.EncodeToString(digest[:])
}
func (b apiOwnershipBackend) swaggerVersion(ctx context.Context, r OwnedResource) (string, error) {
	if r.Version != "" {
		return r.Version, nil
	}
	versions, err := b.c.SwaggerVersions(ctx, r.Namespace, r.Name)
	if err != nil {
		if authoritativeOwnedAbsence(err) {
			return "", errOwnedAbsent
		}
		return "", err
	}
	if len(versions) == 0 {
		return "", errOwnedAbsent
	}
	if len(versions) != 1 {
		return "", errors.New("ambiguous owned Swagger versions")
	}
	return versions[0], nil
}
func (b apiOwnershipBackend) readSwagger(ctx context.Context, r OwnedResource) (client.Metadata, error) {
	version, err := b.swaggerVersion(ctx, r)
	if err != nil {
		return client.Metadata{}, err
	}
	object, err := b.c.GetSwaggerObject(ctx, r.Namespace, r.Name, version)
	if err != nil {
		if authoritativeOwnedAbsence(err) {
			return client.Metadata{}, errOwnedAbsent
		}
		return client.Metadata{}, err
	}
	content, err := object.Content(r.Namespace, r.Name, version)
	if err != nil || r.ContentSHA256 == "" || publicContentDigest(content) != r.ContentSHA256 {
		return client.Metadata{}, errors.New("owned Swagger content mismatch")
	}
	run := b.run
	if run == "" {
		run = strings.Split(strings.TrimPrefix(r.Name, TestResourcePrefix), "-")[0]
	}
	return client.Metadata{Name: r.Name, Namespace: r.Namespace, UID: version, Annotations: map[string]string{ownershipAnnotation: run}}, nil
}
func (b apiOwnershipBackend) deleteSwagger(ctx context.Context, r OwnedResource) error {
	if _, err := b.readSwagger(ctx, r); err != nil {
		return err
	}
	version, err := b.swaggerVersion(ctx, r)
	if err != nil {
		return err
	}
	err = b.c.DeleteSwaggerObject(ctx, r.Namespace, r.Name, version)
	if authoritativeOwnedAbsence(err) {
		return errOwnedAbsent
	}
	return err
}
func (t *ownershipTransport) createSwagger(req *http.Request) (*http.Response, error) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	_ = req.Body.Close()
	var input struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
		Content   string `json:"string_value"`
	}
	if json.Unmarshal(data, &input) != nil {
		return nil, errors.New("invalid Swagger request")
	}
	r := OwnedResource{Type: "xcsh_swagger_object", Namespace: input.Namespace, Name: input.Name, ContentSHA256: publicContentDigest(input.Content)}
	if err = t.j.validate(r); err != nil {
		return nil, err
	}
	if err = client.ValidateSwaggerContent(input.Content); err != nil {
		return nil, errors.New("invalid Swagger fixture")
	}
	versions, err := t.c.SwaggerVersions(req.Context(), r.Namespace, r.Name)
	if err != nil && !authoritativeOwnedAbsence(err) {
		return nil, errors.New("swagger absence verification failed")
	}
	if len(versions) != 0 {
		return nil, errors.New("swagger target has independently owned versions")
	}
	if err = t.j.Intent(r); err != nil {
		return nil, err
	}
	clone := req.Clone(req.Context())
	clone.Body = io.NopCloser(bytes.NewReader(data))
	clone.ContentLength = int64(len(data))
	clone.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(data)), nil }
	response, err := t.base.RoundTrip(clone)
	if err != nil {
		return response, err
	}
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response, nil
	}
	body, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		return nil, err
	}
	response.Body = io.NopCloser(bytes.NewReader(body))
	var object client.SwaggerObject
	if json.Unmarshal(body, &object) != nil || object.Metadata.Name != r.Name || object.Metadata.Namespace != r.Namespace || object.Metadata.Version == "" {
		return nil, errors.New("swagger creation readback missing exact identity")
	}
	r.Version = object.Metadata.Version
	metadata, err := (apiOwnershipBackend{c: t.c, run: t.j.RunID}).readSwagger(req.Context(), r)
	if err != nil {
		return nil, err
	}
	t.j.mu.Lock()
	for i, existing := range t.j.Resources {
		if ownedKey(existing) == ownedKey(r) {
			t.j.Resources[i].Version = r.Version
			t.j.Resources[i].ContentSHA256 = r.ContentSHA256
		}
	}
	err = t.j.save()
	t.j.mu.Unlock()
	if err != nil {
		return nil, err
	}
	if err = t.j.Confirm(r, metadata.UID); err != nil {
		return nil, err
	}
	return response, nil
}
