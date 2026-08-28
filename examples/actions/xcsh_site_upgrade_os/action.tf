# SiteUpgradeOS Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts this upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_os" "example" {
  config {
    name      = "example-value"
    namespace = "example-value"
    version   = "example-value"
  }
}
