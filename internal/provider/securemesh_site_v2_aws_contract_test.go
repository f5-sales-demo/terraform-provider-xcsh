// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestValidateSecuremeshSiteV2AWSContract(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name string
		data SecuremeshSiteV2ResourceModel
	}{
		{
			name: "requires system namespace",
			data: SecuremeshSiteV2ResourceModel{
				Namespace: types.StringValue("tenant-a"),
				AWS:       &SecuremeshSiteV2AWSModel{},
			},
		},
		{
			name: "requires nodes",
			data: SecuremeshSiteV2ResourceModel{
				Namespace: types.StringValue("system"),
				AWS: &SecuremeshSiteV2AWSModel{
					NotManaged: &SecuremeshSiteV2AWSNotManagedModel{NodeList: types.ListNull(types.StringType)},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var resp resource.ValidateConfigResponse
			validateSecuremeshSiteV2AWSContract(ctx, test.data, &resp)
			if !resp.Diagnostics.HasError() {
				t.Fatal("expected AWS SMSv2 contract validation error")
			}
		})
	}
}
