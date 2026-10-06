---
page_title: "xcsh_cloud_region examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_region examples."
---

# xcsh_cloud_region examples

<a id="canonical-1113021303031313-2130103003133303-3102332131232032-0123302311230122-1221023131133130-0211322133230023-1313233232000301-0100320222200220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-1312131213131300-2032133001002132-0000012113013011-3311003200233300-1320113133322101-0123111332203012-2202233202001321-2113202320323303)
- Examples

<a id="canonical-3010321302022320-0203100222123120-1012322213113322-1333022113111122-0201131012200102-3033102031112230-0031103330021232-3232102121132000"></a>

### Complete configurations for `xcsh_cloud_region`

- [Data source](data-sources--cloud_region--examples--group-001.md#canonical-3101303132200022-0123003311322231-3330303011110000-1220020103221231-3030222230132110-3013130112303132-2010020233230323-2003131010100310): valid configuration.

<a id="canonical-3101303132200022-0123003311322231-3330303011110000-1220020103221231-3030222230132110-3013130112303132-2010020233230323-2003131010100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cloud_region](../data-sources/cloud_region.md#canonical-1312131213131300-2032133001002132-0000012113013011-3311003200233300-1320113133322101-0123111332203012-2202233202001321-2113202320323303)
- [Examples](data-sources--cloud_region--examples--group-001.md#canonical-1113021303031313-2130103003133303-3102332131232032-0123302311230122-1221023131133130-0211322133230023-1313233232000301-0100320222200220)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_region/data-source.tf`; digest `sha256:78c1e9c6e8676d021e9b37f2f6a06e7aa0e27e7d2ed71f73c7951ef98b928563`.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```
