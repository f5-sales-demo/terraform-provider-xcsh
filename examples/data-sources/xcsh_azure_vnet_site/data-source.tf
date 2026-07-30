# AzureVNETSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AzureVNETSite by name
data "xcsh_azure_vnet_site" "example" {
  name      = "example-azure-vnet-site"
  namespace = "staging"
}

output "azure_vnet_site_id" {
  value = data.xcsh_azure_vnet_site.example.id
}
