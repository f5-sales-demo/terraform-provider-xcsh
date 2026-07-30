# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
