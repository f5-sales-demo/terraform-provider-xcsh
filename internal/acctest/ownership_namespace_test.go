// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestOwnedNamespaceCleanupUsesVerifiedEmptyCascade(t *testing.T) {
	for _, tc := range []struct {
		name, refs, response string
		want                 bool
	}{
		{"empty", `[]`, `{"items":[]}`, true},
		{"unrelated child", `[{"name":"unrelated","namespace":"fixture"}]`, `{"items":[]}`, false},
		{"missing references", `null`, `{"items":[]}`, false},
		{"empty omitted items", `[]`, `{}`, true},
		{"null response", `[]`, `null`, false},
		{"unknown response", `[]`, `{"error":"uncertain"}`, false},
		{"malformed items", `[]`, `{"items":"invalid"}`, false},
		{"failed response", `[]`, `{"items":[{"error_message":"failure"}]}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			j, err := CreateOwnershipJournal(filepath.Join(t.TempDir(), "journal.json"), "runone")
			if err != nil {
				t.Fatal(err)
			}
			owned := OwnedResource{Type: "xcsh_namespace", Name: "tf-acc-test-runone-ns"}
			if err = j.Intent(owned); err != nil {
				t.Fatal(err)
			}
			if err = j.Confirm(owned, "immutable"); err != nil {
				t.Fatal(err)
			}
			owned.UID = ownershipUIDDigest("immutable")
			deletes := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.Method == "GET" {
					json.NewEncoder(w).Encode(map[string]interface{}{"metadata": map[string]interface{}{"name": owned.Name}, "system_metadata": map[string]interface{}{"uid": "immutable"}, "referring_objects": json.RawMessage(tc.refs)})
					return
				}
				if r.Method == "POST" && r.URL.Path == "/api/web/namespaces/"+owned.Name+"/cascade_delete" {
					deletes++
					w.Write([]byte(tc.response))
					return
				}
				t.Error("unsafe namespace delete route")
				w.WriteHeader(501)
			}))
			defer server.Close()
			c := client.NewClient(server.URL, "synthetic")
			c.MaxRetries = 0
			err = (apiOwnershipBackend{c: c, run: j.RunID}).Delete(context.Background(), owned)
			if (err == nil) != tc.want {
				t.Fatalf("namespace cleanup decision mismatch: %v", err)
			}
			if tc.name == "unrelated child" || tc.name == "missing references" {
				if deletes != 0 {
					t.Fatal("cascade sent without empty verification")
				}
			}
			if tc.want && deletes != 1 {
				t.Fatal("cascade command unavailable")
			}
		})
	}
}

func TestOwnedNamespaceDeletionRaceAcceptsTypedAbsence(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		w.Write([]byte(`{"code":"NOT_FOUND"}`))
	}))
	defer server.Close()
	c := client.NewClient(server.URL, "synthetic")
	err := (apiOwnershipBackend{c: c}).Delete(context.Background(), OwnedResource{Type: "xcsh_namespace", Name: "tf-acc-test-runone-ns", UID: ownershipUIDDigest("immutable")})
	if !errors.Is(err, errOwnedAbsent) {
		t.Fatal("typed namespace absence rejected during deletion race")
	}
}

func TestNamespaceDeletionInProgressRequiresTypedExactDiagnostic(t *testing.T) {
	for _, tc := range []struct {
		status int
		body   string
		want   bool
	}{
		{400, `{"message":"Namespace synthetic already mark for deletion"}`, true},
		{400, `{"message":"Namespace synthetic already marked for deletion"}`, true},
		{400, `{"message":"Namespace synthetic already scheduled for deletion"}`, true},
		{400, `{"message":"Namespace deletion failed"}`, false},
		{403, `{"message":"Namespace synthetic already scheduled for deletion"}`, false},
		{400, `{"message":"Other object already scheduled for deletion"}`, false},
		{400, `invalid`, false},
	} {
		err := xcsherrors.NewAPIError(tc.status, []byte(tc.body), "synthetic", "POST")
		if namespaceDeletionInProgress(err) != tc.want {
			t.Fatal("uncertain namespace error treated as acknowledged deletion")
		}
	}
}
