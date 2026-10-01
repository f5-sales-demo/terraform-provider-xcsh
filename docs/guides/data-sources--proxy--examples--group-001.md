---
page_title: "xcsh_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy examples."
---

# xcsh_proxy examples

<a id="canonical-d14acff9d989657492459e395475a529ffb9792ff80fe7c7e6c110e70623badf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9157028c8e6ca9ed5525484d8201af79b7015a81a5c7c4d871911761781a76e6"></a>

## Examples — Examples / 4a1e8667c3ad / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- Examples

<a id="canonical-0a61d65f79ddce06486f08a29c3bb4fee60d25543f420f6c794f7d284a19bd1a"></a>

## Complete configurations — Examples / 4a1e8667c3ad / 3

- [Data source](data-sources--proxy--examples--group-001.md#canonical-c5bef1507a0a5b6ef5666df47c2aed046d3dfecbaf6bcdfce332178f49462a50): valid configuration.

<a id="canonical-4535e1054773a4fab14c7f4a04cf541235b5cd491f5e067951a3996c7c7fc2e6"></a>

## Next pages — Examples / 4a1e8667c3ad / 4

- [Data source](data-sources--proxy--examples--group-001.md#canonical-c5bef1507a0a5b6ef5666df47c2aed046d3dfecbaf6bcdfce332178f49462a50)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)

<a id="canonical-c5bef1507a0a5b6ef5666df47c2aed046d3dfecbaf6bcdfce332178f49462a50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34a65346847e20be75bd111cdd90ea8968a64731a7e4baa272759f96e9d72637"></a>

## Data source — Data source / 41f174af8ff6 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
- [Examples](data-sources--proxy--examples--group-001.md#canonical-d14acff9d989657492459e395475a529ffb9792ff80fe7c7e6c110e70623badf)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_proxy/data-source.tf`; digest `sha256:d2bc93690268dd4557ac75b114b1e4c362c8f8d843db7cc8d8dd1623ef5dbfe6`.

```terraform
# Proxy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Proxy by name
data "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}

output "proxy_id" {
  value = data.xcsh_proxy.example.id
}
```

<a id="canonical-5ca29503a135817670da12a0c86e5cb018512a343fea421b69be0e9a8cc5d297"></a>

## Next pages — Data source / 41f174af8ff6 / 3

- [Examples](data-sources--proxy--examples--group-001.md#canonical-d14acff9d989657492459e395475a529ffb9792ff80fe7c7e6c110e70623badf)
- [xcsh_proxy](../data-sources/proxy.md#canonical-9aca84c8091d75518a6b3463e1bacffff1cd7efc46037dfe20ab0b065e1dce0f)
