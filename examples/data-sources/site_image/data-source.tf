# Example: retrieve the F5 Customer Edge QCOW2 image and its checksum URL.
#
# The URLs are sensitive and may be short-lived. Do not output them or put them
# in a tfvars file. Pass the image URL directly to the download/provisioning
# resource that consumes it.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "kvm" {
  platform = "Kvm"
}
