# Adopt the exact XC-owned SLI child discovered after a KVM CE registers.

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 10.0.0"
    }
  }
}

resource "xcsh_smsv2_kvm_runtime_interface" "sli" {
  namespace      = "system"
  site           = "onprem-example-kvm"
  interface_name = "ves-io-securemesh-site-v2-onprem-example-kvm-network-onprem-ce-01-12345-ens4-0"
  expected_mac   = "52:54:00:20:00:11"
  hostname       = "onprem-ce-01-12345"
  device         = "ens4"
  ipv4_cidr      = "10.201.0.11/24"
}
