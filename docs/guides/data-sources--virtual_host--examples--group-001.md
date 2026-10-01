---
page_title: "xcsh_virtual_host examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host examples."
---

# xcsh_virtual_host examples

<a id="canonical-1007d5b46c405c4dac828dcf0244f383fa0b2d87d812c2533acf250ab93f0856"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-363b70a769396565b155206961456c7804091ce0e2188042d73bdde613936ba0"></a>

## Examples — Examples / a1921444b860 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- Examples

<a id="canonical-dbf8acc1f18deae949aa7b51b0e5ddb87d208fee48fa943a5693c7327cfa272f"></a>

## Complete configurations — Examples / a1921444b860 / 3

- [Data source](data-sources--virtual_host--examples--group-001.md#canonical-f4f015266e79eda02eb19eb0c34839a2ca186d0ebb3e2a8a9f2c9f6d435879b3): valid configuration.

<a id="canonical-891ec05edf49bc4492807124fbc48fb077a64c1de1d7e69ed1c898d84d813516"></a>

## Next pages — Examples / a1921444b860 / 4

- [Data source](data-sources--virtual_host--examples--group-001.md#canonical-f4f015266e79eda02eb19eb0c34839a2ca186d0ebb3e2a8a9f2c9f6d435879b3)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)

<a id="canonical-f4f015266e79eda02eb19eb0c34839a2ca186d0ebb3e2a8a9f2c9f6d435879b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d9aae3f0fd16a0d60dd40f2236825997d033c6786123e495b835d3289edfcb1"></a>

## Data source — Data source / f3825ac6c666 / 2

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
- [Examples](data-sources--virtual_host--examples--group-001.md#canonical-1007d5b46c405c4dac828dcf0244f383fa0b2d87d812c2533acf250ab93f0856)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_host/data-source.tf`; digest `sha256:23ea7a52b77c427f7e6b79f3895a5a23c6caa53b8b9e855b6dcaac10831f3a72`.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```

<a id="canonical-5211872a5039f6e445b220b216d49df2b9947824b4e7723e28927f1187c8695f"></a>

## Next pages — Data source / f3825ac6c666 / 3

- [Examples](data-sources--virtual_host--examples--group-001.md#canonical-1007d5b46c405c4dac828dcf0244f383fa0b2d87d812c2533acf250ab93f0856)
- [xcsh_virtual_host](../data-sources/virtual_host.md#canonical-c365aed4cfa0fbfdfe78adcb3521d491ecfe7a11e97f7277bd85e506856cbee8)
