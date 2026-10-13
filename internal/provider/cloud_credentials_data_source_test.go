// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider_test

import (
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/acctest"
	"testing"
)

func TestAccCloudCredentialsDataSource_basic(t *testing.T) {
	acctest.RequireExternalCredential(t, "cloud")

}
