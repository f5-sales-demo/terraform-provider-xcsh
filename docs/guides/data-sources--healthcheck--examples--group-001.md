---
page_title: "xcsh_healthcheck examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck examples."
---

# xcsh_healthcheck examples

<a id="canonical-f0b76284dbc9978e50c137bedc62f4d54a39d90fbf353d0ae54dede3e1b39214"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44b80982491c0cb540e772cee47d031c35b77681574e00c5bf548391b1b050d4"></a>

## Examples — Examples / e9ee35ffefc9 / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- Examples

<a id="canonical-93441eeabf27ca4f1dba3d30622a5db3d39b8c8382a6d4052549bc322b613775"></a>

## Complete configurations — Examples / e9ee35ffefc9 / 3

- [Data source](data-sources--healthcheck--examples--group-001.md#canonical-9f514ab3c0e90c8868ae010cbe3721055bd04b53d6793dc40a0b272d81a95ab8): valid configuration.

<a id="canonical-9eaa807db2cffe40e5094807cbde7349517e62dd3c9bee5733cbcc1cd874b004"></a>

## Next pages — Examples / e9ee35ffefc9 / 4

- [Data source](data-sources--healthcheck--examples--group-001.md#canonical-9f514ab3c0e90c8868ae010cbe3721055bd04b53d6793dc40a0b272d81a95ab8)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)

<a id="canonical-9f514ab3c0e90c8868ae010cbe3721055bd04b53d6793dc40a0b272d81a95ab8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13ce9d1109ceaf67fa8be9f55d1f10c638d0b44b04323895117fa6a99b319683"></a>

## Data source — Data source / c22f1346193b / 2

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
- [Examples](data-sources--healthcheck--examples--group-001.md#canonical-f0b76284dbc9978e50c137bedc62f4d54a39d90fbf353d0ae54dede3e1b39214)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_healthcheck/data-source.tf`; digest `sha256:35072c285c20d3995104b4e1308ca686f5cc2a63a3a1e8c4cc3c6d11f32af3c6`.

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

<a id="canonical-dd461d565ff647ca051b6a39aff5081d6ebd9ad2eabafd7eebdbb99020fa77d2"></a>

## Next pages — Data source / c22f1346193b / 3

- [Examples](data-sources--healthcheck--examples--group-001.md#canonical-f0b76284dbc9978e50c137bedc62f4d54a39d90fbf353d0ae54dede3e1b39214)
- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-725ac46fa2736f6e2d5f694e53ef2171f3c08fc8b2d65790732141dbfc7b2da6)
