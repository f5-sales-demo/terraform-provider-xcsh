---
page_title: "xcsh_address_allocator examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator examples."
---

# xcsh_address_allocator examples

<a id="canonical-1312010000001103-1030323302011222-3330121220332231-1131222032331013-0330132110232100-0320310032101302-3103310132020301-3111132203120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-2303133311012021-3322130130031103-2120032332031012-3013200021213232-0100310121031123-3230030130121113-3213033330111232-3203303021311023)
- Examples

<a id="canonical-3302003222132322-2330010203212312-2111021210323130-3122203213123002-0201111303220010-2232220100110313-3220311110133122-0130321213222020"></a>

### Complete configurations for `xcsh_address_allocator`

- [Data source](data-sources--address_allocator--examples--group-001.md#canonical-3210130013211311-0020002222320233-0303230023102233-0321310111322301-3102130302120221-2332232303310002-1110020003210310-3232301222322001): valid configuration.

<a id="canonical-3210130013211311-0020002222320233-0303230023102233-0321310111322301-3102130302120221-2332232303310002-1110020003210310-3232301222322001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-2303133311012021-3322130130031103-2120032332031012-3013200021213232-0100310121031123-3230030130121113-3213033330111232-3203303021311023)
- [Examples](data-sources--address_allocator--examples--group-001.md#canonical-1312010000001103-1030323302011222-3330121220332231-1131222032331013-0330132110232100-0320310032101302-3103310132020301-3111132203120120)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_address_allocator/data-source.tf`; digest `sha256:8199b921632de8cfafabbbf9803b9c84ef8413a91ef1f5c0757142ce1c4bea0d`.

```terraform
# AddressAllocator Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddressAllocator by name
data "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"
}

output "address_allocator_id" {
  value = data.xcsh_address_allocator.example.id
}
```
