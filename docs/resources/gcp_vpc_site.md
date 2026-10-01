---
page_title: "xcsh_gcp_vpc_site landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site landing."
---

# xcsh_gcp_vpc_site landing

<a id="canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03891e98ca0bd8763d92c3ac4973f2af014ff5557c1f300139333ce4f87c7c0f"></a>

## xcsh_gcp_vpc_site — xcsh_gcp_vpc_site / 99ff11abf6b8 / 2

Breadcrumbs:

- xcsh_gcp_vpc_site

Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud
VPC environments.

<a id="canonical-2dceeb587d2b85ec3dcbeeb3e91b850bbfeb89781449f2aafe9f3c52d42bbd1c"></a>

## Prerequisites — xcsh_gcp_vpc_site / 99ff11abf6b8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cloud_credentials`.

- cloud_credentials: GCP authentication for deployment

<a id="canonical-7643fe4eaf517c1ffd0a8a3560ca77ac205c5084e530c4142c4dcbdec19b5649"></a>

## Minimal configuration — xcsh_gcp_vpc_site / 99ff11abf6b8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GCPVPCSite Resource Example
# Manages a GCP VPC Site resource in F5 Distributed Cloud for deploying F5 sites within Google Cloud VPC environments.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GCPVPCSite configuration
resource "xcsh_gcp_vpc_site" "example" {
  name      = "example-gcp-vpc-site"
  namespace = "staging"

  gcp_region    = "example-value"
  instance_type = "example-value"
  ssh_key       = "example-value"
}
```

<a id="canonical-c9874df86c4a774e88d27e5838adf63cf421c7dbbcfb65d9d22ce53bd86ebcef"></a>

## Root configuration — xcsh_gcp_vpc_site / 99ff11abf6b8 / 5

Required root properties: `gcp_region`, `instance_type`, `name`, `namespace`, `ssh_key`. Full root flags and choices appear in the property reference.

<a id="canonical-1513b45a638827737173e7014d781ae409823c82554bb4a33ced254553086dd4"></a>

## Next pages — xcsh_gcp_vpc_site / 99ff11abf6b8 / 6

- [Property reference](../guides/resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [Examples](../guides/resources--gcp_vpc_site--examples--group-001.md#canonical-9c3b5148b7bb2f407b76bff289c776f5e6ffb922e40edef4bd02be0a5ddabf62)
- [Import](../guides/resources--gcp_vpc_site--lifecycle--group-001.md#canonical-4746106dcbc16a51f3678d4f5c5aab485c8a1e0a35a6baba5c187788d8a60fd6)
- [Timeouts](../guides/resources--gcp_vpc_site--lifecycle--group-001.md#canonical-6c20293d826cf2d7e58bcf1ae0e1b45e6e06547f76ba125ddd63e4c9a4a2bbaa)
