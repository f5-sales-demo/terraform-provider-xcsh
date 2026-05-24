# Fast ACL Resource Example
# Manages new Fast ACL rule, has specification to match source IP, source port and action to apply. in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fast ACL configuration
resource "f5xc_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # FastAclRuleAction specifies possible action to be applied...
  action {
    # Configure action settings
  }
  # Policer Reference. Reference to policer object.
  policer_action {
    # Configure policer_action settings
  }
  # Reference. A policer direct reference.
  ref {
    # Configure ref settings
  }
}
