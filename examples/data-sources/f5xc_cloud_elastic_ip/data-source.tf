# Cloud Elastic IP Data Source Example
# Retrieves information about an existing Cloud Elastic IP

# Look up an existing Cloud Elastic IP by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_elastic_ip_id" {
#   value = data.f5xc_cloud_elastic_ip.example.id
# }
