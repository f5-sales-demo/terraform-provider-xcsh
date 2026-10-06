terraform {
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh" }
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
