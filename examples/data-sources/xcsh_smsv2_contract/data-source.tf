# Read the immutable clean-break SMSv2 AWS TGW Connect contract compiled into
# the provider. Gate infrastructure mutation on the published capabilities.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_contract" "current" {}

output "smsv2_contract" {
  value = {
    id                  = data.xcsh_smsv2_contract.current.contract_id
    version             = data.xcsh_smsv2_contract.current.contract_version
    api_release         = data.xcsh_smsv2_contract.current.api_release_tag
    telemetry_schema_id = data.xcsh_smsv2_contract.current.telemetry_schema_id
    capabilities        = data.xcsh_smsv2_contract.current.capabilities
    f5xc_authorities    = data.xcsh_smsv2_contract.current.f5xc_authorities
    aws_authorities     = data.xcsh_smsv2_contract.current.aws_authorities
  }
}
