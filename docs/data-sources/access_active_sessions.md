---
page_title: "xcsh_access_active_sessions landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions landing."
---

# xcsh_access_active_sessions landing

<a id="canonical-d80936c849771550fc4fe0718c8a494dc68ce51087b0c6d03d5c8f5446ecad9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5ed5d5714038cfd150b6c8a71f1dea5d59b733953c655e7e798aaab9daacc6d"></a>

## xcsh_access_active_sessions — xcsh_access_active_sessions / 67540dd82140 / 2

Breadcrumbs:

- xcsh_access_active_sessions

Resource retrieval operation.

<a id="canonical-1c8d2292f94c49ced477f707112911c1818f6d21e43e0692dc7ff212d783b2d3"></a>

## Prerequisites — xcsh_access_active_sessions / 67540dd82140 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4203aa30fdd0ad1b03d1acb68e93d08efb5bb4cfe8c7f78d9330fa8bda539c62"></a>

## Minimal configuration — xcsh_access_active_sessions / 67540dd82140 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```

<a id="canonical-054b0c16bf6f1b09f797f3b307a5a423305a8792ca8619e45e5133d2d58c3508"></a>

## Root configuration — xcsh_access_active_sessions / 67540dd82140 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-53df68738f25c43cb7656a900eef0088d1f47cd64d8c568edf78de79af87b408"></a>

## Next pages — xcsh_access_active_sessions / 67540dd82140 / 6

- [Property reference](../guides/data-sources--access_active_sessions--reference--group-001.md#canonical-ea3f8a31392a8fec51cb6b0a32c4892b70baf6ac0574b3877fc23d6b78865c9b)
- [Examples](../guides/data-sources--access_active_sessions--examples--group-001.md#canonical-9c6a94d6ca9e39822b716f3896258ace46b22249c601fe3eb13a110372c48134)
