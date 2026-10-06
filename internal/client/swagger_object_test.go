package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSwaggerExactRoutesAndInvalidContent(t *testing.T) {
	for _, v := range []string{"latest", "LATEST", "../v1", "v1/other", "v1?x", "v1#x"} {
		if _, e := SwaggerIdentity("demo", "fixture", v); e == nil {
			t.Errorf("accepted version %q", v)
		}
	}
	for _, c := range []string{"", "{}", `{"openapi":"3.0.3","openapi":"3.0.3","info":{"title":"x","version":"1"},"paths":{}}`, `{"openapi":"3.0.3","info":{"title":"x","version":"1"},"paths":{"relative":{}}}`, `{"swagger":"2.0","openapi":"3.0.3","info":{"title":"x","version":"1"},"paths":{}}`} {
		if ValidateSwaggerContent(c) == nil {
			t.Error("accepted malformed profile")
		}
	}
	calls := 0
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != "DELETE" || r.URL.Path != "/api/object_store/namespaces/demo/stored_objects/swagger/fixture/v1" {
			t.Errorf("wrong delete %s %s", r.Method, r.URL.Path)
		}
		fmt.Fprint(w, "{}")
	}))
	defer s.Close()
	c := NewClient(s.URL, "fixture", WithMaxRetries(0))
	if err := c.DeleteSwaggerObject(context.Background(), "demo", "fixture", "latest"); err == nil {
		t.Fatal("latest delete accepted")
	}
	if calls != 0 {
		t.Fatal("invalid version sent")
	}
	if err := c.DeleteSwaggerObject(context.Background(), "demo", "fixture", "v1"); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatal(calls)
	}
}
