---
page_title: "xcsh_healthcheck landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck landing."
---

# xcsh_healthcheck landing

<a id="canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e640219ee945ea094f7be738b46473ec7d03033ff6e0ffd67effcb2ead9c658e"></a>

## xcsh_healthcheck — xcsh_healthcheck / 68c827d3f27d / 2

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

<a id="canonical-77d761f4f55daf090dce15e3909c4b374f47656c8879bd721a1af1ac1b1f2a5c"></a>

## Prerequisites — xcsh_healthcheck / 68c827d3f27d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-406ba1f1a31891b33aed0848ca7fe4892f5c38fadb3479c644d0cc1844d8322c"></a>

## Minimal configuration — xcsh_healthcheck / 68c827d3f27d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Healthcheck by name
data "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"
}

output "healthcheck_id" {
  value = data.xcsh_healthcheck.example.id
}
```

<a id="canonical-2dedae94a6457a5d63de83f24d80cd90eeefc2489574bd7ea9425d6a1fe353dc"></a>

## Root configuration — xcsh_healthcheck / 68c827d3f27d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3ccbcf96872e76d9d2f28ce725ab73e562c6fb4692ace4e6ef217d96d003bd0e"></a>

## Next pages — xcsh_healthcheck / 68c827d3f27d / 6

- [Property reference](../guides/data-sources--healthcheck--reference--group-001.md#canonical-af83443e8f0c2ca01493595db5613ed42ea9fa1d54f93db769449e4758dd43b6)
- [Examples](../guides/data-sources--healthcheck--examples--group-001.md#canonical-f0b76284dbc9978e50c137bedc62f4d54a39d90fbf353d0ae54dede3e1b39214)
