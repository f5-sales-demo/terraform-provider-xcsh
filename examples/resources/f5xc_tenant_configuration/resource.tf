# Tenant Configuration Resource Example
# Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification. configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tenant Configuration configuration
resource "f5xc_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # BasicConfiguration.
  basic_configuration {
    # Configure basic_configuration settings
  }
  # BruteForceDetectionSettings.
  brute_force_detection_settings {
    # Configure brute_force_detection_settings settings
  }
  # PasswordPolicy.
  password_policy {
    # Configure password_policy settings
  }
}
