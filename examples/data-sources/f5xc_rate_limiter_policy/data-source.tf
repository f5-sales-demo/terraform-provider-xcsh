# Rate Limiter Policy Data Source Example
# Retrieves information about an existing Rate Limiter Policy

# Look up an existing Rate Limiter Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "rate_limiter_policy_id" {
#   value = data.f5xc_rate_limiter_policy.example.id
# }
