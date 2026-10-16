terraform {
  required_version = ">= 1.14.0"
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 15.8.3"
    }
  }
}

# Install the current published release containing xcsh_secret_policy.
# Keep the namespace independently owned. The application pod must carry app=demo-tls.
resource "xcsh_secret_policy" "application" {
  name                     = "demo-wingman-policy"
  namespace                = "demo-app"
  allow_f5xc               = true
  decrypt_cache_timeout    = "30s"
  rule_combining_algorithm = "FIRST_MATCH"
  rules = [
    { name = "allow-application", action = "ALLOW", client_selector = ["app=demo-tls"] },
    { name = "deny-all", action = "DENY" }
  ]
}
