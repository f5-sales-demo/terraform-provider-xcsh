---
page_title: "xcsh_app_api_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group landing."
---

# xcsh_app_api_group landing

<a id="canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15d80a34db63aff3756d13f4d69710d84fc2b02596b64136d5afa64607db1e50"></a>

## xcsh_app_api_group — xcsh_app_api_group / dac273099e2a / 2

Breadcrumbs:

- xcsh_app_api_group

Manages app\_api\_group creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-c490742e87e2c872421b301deba024681a5edd5352bc0c8d1476a72421346532"></a>

## Prerequisites — xcsh_app_api_group / dac273099e2a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-db0938807671613d0ba484f08871b96e5d35cbd1245d43dc1f0d9f90d9d305bc"></a>

## Minimal configuration — xcsh_app_api_group / dac273099e2a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```

<a id="canonical-a0eee27ceb159e7d08652c76e82ec52093f3d48418437fec21e223afa521b6e2"></a>

## Root configuration — xcsh_app_api_group / dac273099e2a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3beea864950d27ba71e648522fb4a4dfe28c0b586b27e6ac2821c3c92a796d61"></a>

## Next pages — xcsh_app_api_group / dac273099e2a / 6

- [Property reference](../guides/resources--app_api_group--reference--group-001.md#canonical-da38fcbcc55ba288a857b791a3abc027289c19d58cfbd1d037ad6cb64dbc8584)
- [Examples](../guides/resources--app_api_group--examples--group-001.md#canonical-72d930a44e683dcfdf007e37a08cca2ce04cbe9821f13751dffb51b2b504ad90)
- [Import](../guides/resources--app_api_group--lifecycle--group-001.md#canonical-7625e7de94241581993f6fac64f9108ad8c04e60d4781ddee2a2322f367dc02b)
- [Timeouts](../guides/resources--app_api_group--lifecycle--group-001.md#canonical-92fbccdf25666d1e17906524cd56e06c5a4fb4dd7f35dbfc410184c89ea55ad3)
