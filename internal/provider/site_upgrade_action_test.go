// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

func TestSiteUpgradeActionSchemasExposeOnlySMSv2Inputs(t *testing.T) {
	tests := []struct {
		name string
		new  func() action.Action
		want []string
	}{
		{name: "software", new: NewSiteUpgradeSwAction, want: []string{"site", "software_version"}},
		{name: "operating system", new: NewSiteUpgradeOSAction, want: []string{"os_version", "site"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var response action.SchemaResponse
			test.new().Schema(context.Background(), action.SchemaRequest{}, &response)
			got := make([]string, 0, len(response.Schema.Attributes))
			for name := range response.Schema.Attributes {
				got = append(got, name)
			}
			sort.Strings(got)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("schema attributes = %v, want %v", got, test.want)
			}
		})
	}
}

func TestSiteUpgradeActionsBindSystemNamespaceAndForceFalse(t *testing.T) {
	tests := []struct {
		name       string
		pathSuffix string
		model      any
		invoke     func(*client.Client, action.InvokeRequest, *action.InvokeResponse)
	}{
		{
			name: "software", pathSuffix: "upgrade_sw",
			model: SiteUpgradeSwActionModel{Site: types.StringValue("site/one"), SoftwareVersion: types.StringValue("crt-target")},
			invoke: func(c *client.Client, request action.InvokeRequest, response *action.InvokeResponse) {
				(&SiteUpgradeSwAction{client: c}).Invoke(context.Background(), request, response)
			},
		},
		{
			name: "operating system", pathSuffix: "upgrade_os",
			model: SiteUpgradeOSActionModel{Site: types.StringValue("site/one"), OSVersion: types.StringValue("9.2026.17")},
			invoke: func(c *client.Client, request action.InvokeRequest, response *action.InvokeResponse) {
				(&SiteUpgradeOSAction{client: c}).Invoke(context.Background(), request, response)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var gotBody map[string]any
			var gotRequestURI string
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				gotRequestURI = request.RequestURI
				if request.Method != http.MethodPost {
					t.Errorf("method = %s, want POST", request.Method)
				}
				if err := json.NewDecoder(request.Body).Decode(&gotBody); err != nil {
					t.Errorf("decode request: %v", err)
				}
				writer.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()

			var schemaResponse action.SchemaResponse
			actionUnderTest := actionForModel(test.model)
			actionUnderTest.Schema(context.Background(), action.SchemaRequest{}, &schemaResponse)
			request := action.InvokeRequest{Config: tfsdk.Config{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, test.model, schemaResponse.Schema.Type())}}
			var response action.InvokeResponse
			test.invoke(client.NewClient(server.URL, "token", client.WithMaxRetries(0)), request, &response)
			if response.Diagnostics.HasError() {
				t.Fatalf("invoke diagnostics: %v", response.Diagnostics)
			}
			wantURI := "/api/config/namespaces/system/sites/site%2Fone/" + test.pathSuffix
			if gotRequestURI != wantURI {
				t.Fatalf("request URI = %q, want %q", gotRequestURI, wantURI)
			}
			if gotBody["name"] != "site/one" || gotBody["namespace"] != "system" || gotBody["force"] != false {
				t.Fatalf("request body = %#v", gotBody)
			}
			if _, exists := gotBody["software_version"]; exists {
				t.Fatalf("request leaked public software_version key: %#v", gotBody)
			}
			if _, exists := gotBody["os_version"]; exists {
				t.Fatalf("request leaked public os_version key: %#v", gotBody)
			}
		})
	}
}

func TestSiteUpgradeActionAPIErrorsAreSanitized(t *testing.T) {
	const privateResponse = "private-upstream-detail"
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, privateResponse, http.StatusBadRequest)
	}))
	defer server.Close()
	actionUnderTest := &SiteUpgradeSwAction{client: client.NewClient(server.URL, "token", client.WithMaxRetries(0))}
	var schemaResponse action.SchemaResponse
	actionUnderTest.Schema(context.Background(), action.SchemaRequest{}, &schemaResponse)
	model := SiteUpgradeSwActionModel{Site: types.StringValue("site"), SoftwareVersion: types.StringValue("target")}
	request := action.InvokeRequest{Config: tfsdk.Config{Schema: schemaResponse.Schema, Raw: responseOperationRaw(t, model, schemaResponse.Schema.Type())}}
	var response action.InvokeResponse
	actionUnderTest.Invoke(context.Background(), request, &response)
	if !response.Diagnostics.HasError() {
		t.Fatal("API rejection did not fail closed")
	}
	for _, diagnostic := range response.Diagnostics.Errors() {
		if strings.Contains(diagnostic.Detail(), privateResponse) {
			t.Fatalf("diagnostic leaked API response: %q", diagnostic.Detail())
		}
	}
}

func actionForModel(model any) action.Action {
	switch model.(type) {
	case SiteUpgradeSwActionModel:
		return NewSiteUpgradeSwAction()
	case SiteUpgradeOSActionModel:
		return NewSiteUpgradeOSAction()
	default:
		panic("unsupported action model")
	}
}
