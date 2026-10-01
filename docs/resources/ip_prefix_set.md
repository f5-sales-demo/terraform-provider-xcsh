---
page_title: "xcsh_ip_prefix_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set landing."
---

# xcsh_ip_prefix_set landing

<a id="canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b64232522b84f92e276469da2721b3476e568f98f777c1f254f125fc104b10aa"></a>

## xcsh_ip_prefix_set — xcsh_ip_prefix_set / 53f1c0c1062d / 2

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-6d5a12991a172f9ee723c0bb28ca23903bd931c5aac24ad6e092667faefb7ea8"></a>

## Prerequisites — xcsh_ip_prefix_set / 53f1c0c1062d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-96c4f3a22792a2792951a19461d5237403334a05fb6d431423e58606b556129b"></a>

## Minimal configuration — xcsh_ip_prefix_set / 53f1c0c1062d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

<a id="canonical-f0bfa0fec5c2d204803b1735087dc82b9ccdb657808ef6ec9bf7822984d147d5"></a>

## Root configuration — xcsh_ip_prefix_set / 53f1c0c1062d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d6f4e8602c2a0bf9ed59d726ee091b1ad912d196c3a5f469eedf4e10ed4ce864"></a>

## Next pages — xcsh_ip_prefix_set / 53f1c0c1062d / 6

- [Property reference](../guides/resources--ip_prefix_set--reference--group-001.md#canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9)
- [Examples](../guides/resources--ip_prefix_set--examples--group-001.md#canonical-663015d63e3a93116ac51fd9a7e4beaa6cb87aac18bd9ffbdc01f28309a6c55e)
- [Import](../guides/resources--ip_prefix_set--lifecycle--group-001.md#canonical-27139af5b9b634db45df1796498b1c8f426e5cd150ddac446c20e7bcf89ad306)
- [Timeouts](../guides/resources--ip_prefix_set--lifecycle--group-001.md#canonical-2baddc4b33b66191ebd6e79ad77556cf483aab47bad1398675f48723bc3f1db0)
