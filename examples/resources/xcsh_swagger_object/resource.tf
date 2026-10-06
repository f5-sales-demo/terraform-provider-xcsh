terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema"
  content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
