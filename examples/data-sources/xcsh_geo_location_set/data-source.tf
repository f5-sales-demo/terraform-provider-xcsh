# GeoLocationSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GeoLocationSet by name
data "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "staging"
}

output "geo_location_set_id" {
  value = data.xcsh_geo_location_set.example.id
}
