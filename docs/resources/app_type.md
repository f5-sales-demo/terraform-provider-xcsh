---
page_title: "xcsh_app_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type landing."
---

# xcsh_app_type landing

<a id="canonical-a8407eb1efb2ca1dacc63939149c3ba5bda3fedc3204bcc1b84ff7225a913692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a5039dbbba34bd197689d5ad1b41d3a3e2f8f6a8d4da649df3f8dc059a0062e"></a>

## xcsh_app_type — xcsh_app_type / 53d5728980ea / 2

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-092819d7cb0415810e1af5c2336187a5124fcf6a4109ad0903115ed7bd3453c7"></a>

## Prerequisites — xcsh_app_type / 53d5728980ea / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c20756fb67b22f67a589eeb248217b5ecb3eb734266a254c2ff44ae388298bb2"></a>

## Minimal configuration — xcsh_app_type / 53d5728980ea / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

<a id="canonical-fb6c705e5f6a8f55794cb12a7b534e20f50c3d45d9041ef8ffd7ffe8949f939a"></a>

## Root configuration — xcsh_app_type / 53d5728980ea / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1adb0d8b2e04f2e448d3440c8e58683fb22012e00c6db4cefdd95ea24de8167d"></a>

## Next pages — xcsh_app_type / 53d5728980ea / 6

- [Property reference](../guides/resources--app_type--reference--group-001.md#canonical-841debb259db374e58022de533dff86777952c5a097831b4bcdff0ae78059c78)
- [Examples](../guides/resources--app_type--examples--group-001.md#canonical-f7b9298394e83b8e4a094f5b90535372a175dda62f0fa6b17e5bb40ccf6b6b6d)
- [Import](../guides/resources--app_type--lifecycle--group-001.md#canonical-ab258ec5a93ee1df71192e632b797ccc4b4b30bfbb21b13b9c0678be62a3aa05)
- [Timeouts](../guides/resources--app_type--lifecycle--group-001.md#canonical-41f605c431033c6ade1a121ab4c7178617f1f7adec16124a4da7e946b8e889e9)
