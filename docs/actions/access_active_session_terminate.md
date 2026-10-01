---
page_title: "xcsh_access_active_session_terminate landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session_terminate landing."
---

# xcsh_access_active_session_terminate landing

<a id="canonical-2f194631960b3886621a0ca55d7034950298c86bd87a8b01dbfa9bfbf6957dc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f8b01c6c30200fb57dda13cf078709d2f887887eb5973eeeb4c50b75a10e9ad"></a>

## xcsh_access_active_session_terminate — xcsh_access_active_session_terminate / c82fca1f594e / 2

Breadcrumbs:

- xcsh_access_active_session_terminate

Resource deletion operation.

<a id="canonical-b6036ba203ed1f855c3d07588a8cb6e20d46e289c3d2c19e3da969c868e257aa"></a>

## Prerequisites — xcsh_access_active_session_terminate / c82fca1f594e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-44331457557760af89d2ca1417b97acb70c1bf9a489c7538fe7c5348ed516250"></a>

## Minimal configuration — xcsh_access_active_session_terminate / c82fca1f594e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```

<a id="canonical-73eec4c515116c60576b097b5e27fdb2541ab89453c73924133eed504823aa94"></a>

## Root configuration — xcsh_access_active_session_terminate / c82fca1f594e / 5

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-eb843714fd3a35d7e061144876fe48352f605456cc38b53963c5feb08a8a3769"></a>

## Next pages — xcsh_access_active_session_terminate / c82fca1f594e / 6

- [Property reference](../guides/actions--access_active_session_terminate--reference--group-001.md#canonical-864e20de9897d401a5520422020dd1f8ad4a591c625080ef2b314f6ffe92d62a)
- [Examples](../guides/actions--access_active_session_terminate--examples--group-001.md#canonical-5f7f94defd00ef83888bc10e66540b97d13fdba27ff8c1198ed4890ac7763246)
- [Lifecycle](../guides/actions--access_active_session_terminate--lifecycle--group-001.md#canonical-0981081a99815666a3e5014cfdcb353f87640941c4722d5a14d88fe370f20297)
