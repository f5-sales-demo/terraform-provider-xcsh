---
page_title: "xcsh_segment landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment landing."
---

# xcsh_segment landing

<a id="canonical-7a8b795fee80e2e6cc7579e56e701793aec6fb3c8a0be25bcde59beed1d43650"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f2438a8d358508b19f338a0c6215bd9c37cb4639153fd755d0b918c7e35008"></a>

## xcsh_segment — xcsh_segment / 681d7b2d58c0 / 2

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

<a id="canonical-2a74ccd42be16964a07c90776814f47a7dc3edf51002ccdaab40224a5f308ef5"></a>

## Prerequisites — xcsh_segment / 681d7b2d58c0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4b60fa0a410dbb1a6a2cdc1f1a79de44a00f410ba40428a15b12aaa4a8bac4ef"></a>

## Minimal configuration — xcsh_segment / 681d7b2d58c0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```

<a id="canonical-ac7c343e2e238f4f1c945ba6ab7e1a7d6e3307d4d9a061c54cc36808b4a59386"></a>

## Root configuration — xcsh_segment / 681d7b2d58c0 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-5a3decc658872e5412d10d6099c5e41a98e56ac25d51d3ded19caa6db830519e"></a>

## Next pages — xcsh_segment / 681d7b2d58c0 / 6

- [Property reference](../guides/data-sources--segment--reference--group-001.md#canonical-542cdf7a24561b3f7a476758c7d287b3eb9814b546e6df99ab410ec927020b6f)
- [Examples](../guides/data-sources--segment--examples--group-001.md#canonical-248c349c9eaccfb1eb432c30ddd926ca75a9811b28d5ce7c70003f7f597195f5)
