# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "staging"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
