terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

locals {
  schema_content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema-${substr(sha256(local.schema_content), 0, 32)}"
  content   = local.schema_content
  lifecycle {
    create_before_destroy = true
  }
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
