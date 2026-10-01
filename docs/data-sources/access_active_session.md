---
page_title: "xcsh_access_active_session landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session landing."
---

# xcsh_access_active_session landing

<a id="canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-177fed81c4b02e57e79071acee3e7baa8adbf2094e8fc442ce53ec2cea8b5453"></a>

## xcsh_access_active_session — xcsh_access_active_session / 80a748d1428b / 2

Breadcrumbs:

- xcsh_access_active_session

Resource retrieval operation.

<a id="canonical-5ff8590703c05e5425ea2d0178846976a4193aefbccd6f5fdbcd2eea77843e31"></a>

## Prerequisites — xcsh_access_active_session / 80a748d1428b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-01895b2d2aa5eb8461303ae8c2ba781b0d48ae935a1d669ef8224ceda773b783"></a>

## Minimal configuration — xcsh_access_active_session / 80a748d1428b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```

<a id="canonical-2a3547a5e329e1c2d8b5dcdd33f978ce8d89a4d98aeb44a45e8fab3dd9940fac"></a>

## Root configuration — xcsh_access_active_session / 80a748d1428b / 5

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2eebb26a37ed22ca88054cb7914f5c1bf54c124606d81785d06b8c80e35235b4"></a>

## Next pages — xcsh_access_active_session / 80a748d1428b / 6

- [Property reference](../guides/data-sources--access_active_session--reference--group-001.md#canonical-f9145d41000cc5377b88dd55e2dfd831d018e430e7e6108f72457e70d726428d)
- [Examples](../guides/data-sources--access_active_session--examples--group-001.md#canonical-1fae4c0e495ede3a8cc200a71a1345278a488be35029641220310fefd9671a9e)
