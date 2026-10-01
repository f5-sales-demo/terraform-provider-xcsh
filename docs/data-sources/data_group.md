---
page_title: "xcsh_data_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_group landing."
---

# xcsh_data_group landing

<a id="canonical-2a09189b698920d6a460102d0b7a21e1e3e6e93e2691c7d484224253a3e63dc2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-466aa7c58ed861a073e792283d1e0473746b28379ae8e9ae41f52f780910240b"></a>

## xcsh_data_group — xcsh_data_group / 6fd816eb1c88 / 2

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-3d90ff6779978d518d69a94331db94f22d70a5fe83560581d8e60d47ec60dc89"></a>

## Prerequisites — xcsh_data_group / 6fd816eb1c88 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1c9687c2e53bedb3c9ada4b9f517cdf67f4812ad7a3b8fae003eec2e49749415"></a>

## Minimal configuration — xcsh_data_group / 6fd816eb1c88 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```

<a id="canonical-4f417d79e33c1516e3e531a74efdb5c8fa228e0865d6c9bc62f2bd3fa42144d2"></a>

## Root configuration — xcsh_data_group / 6fd816eb1c88 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e1d1957d2d298bf94bf7a0e5decc9c74717acea12ad3bc673b898491d4f2dcf3"></a>

## Next pages — xcsh_data_group / 6fd816eb1c88 / 6

- [Property reference](../guides/data-sources--data_group--reference--group-001.md#canonical-9af800ba0cbf5ecb49db6ffae6b8045770a7445bd1c47bd0c92945103279aa81)
- [Examples](../guides/data-sources--data_group--examples--group-001.md#canonical-30577a5df82559f5bd7a5798e84dcf0a6240212b5fc1fd6cbfe37d0f6b76bb0f)
