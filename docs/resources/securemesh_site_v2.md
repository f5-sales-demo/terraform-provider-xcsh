---
page_title: "xcsh_securemesh_site_v2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 landing."
---

# xcsh_securemesh_site_v2 landing

<a id="canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8192b38c307d7a9225d066e8853b2d0de7e690b3c4071b159e8b353ab72bc19f"></a>

## xcsh_securemesh_site_v2 — xcsh_securemesh_site_v2 / 9f2c6dc86e23 / 2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

<a id="canonical-5108bc2fc29d696f74554039d0f794ce95bc6f342c376931f8f2f6a378ff0f08"></a>

## Prerequisites — xcsh_securemesh_site_v2 / 9f2c6dc86e23 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d8f690d31b227c21b8cb6335bc2f0c3b1707e91890899a7f9572c5fa75249f92"></a>

## Minimal configuration — xcsh_securemesh_site_v2 / 9f2c6dc86e23 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Resource Example
# Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with security and networking controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSiteV2 configuration
resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}
```

<a id="canonical-6e32edbcb4d3f58513587fa0283e1f6ed7aed2690aa7d1b5b1527b07a08158c3"></a>

## Root configuration — xcsh_securemesh_site_v2 / 9f2c6dc86e23 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-5f1fd7975bef20fb7f0624caf79bbc6b4d77944208020fb4f0b4ed5596489505"></a>

## Next pages — xcsh_securemesh_site_v2 / 9f2c6dc86e23 / 6

- [Property reference](../guides/resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [Examples](../guides/resources--securemesh_site_v2--examples--group-001.md#canonical-73c06db8e151ec7d688d6a0f99293e7afad26d835f5b001586581f6dcf25331f)
- [Import](../guides/resources--securemesh_site_v2--lifecycle--group-001.md#canonical-86e169906c4be566312046ed41f201624fa38d143863572cf34f247231a26dad)
- [Timeouts](../guides/resources--securemesh_site_v2--lifecycle--group-001.md#canonical-310ebeba746e606eb66633a081c6c0006e9d4488acaf9ff3b64c3247e4a2eeb7)
