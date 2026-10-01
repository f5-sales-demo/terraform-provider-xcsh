---
page_title: "xcsh_address_allocator landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator landing."
---

# xcsh_address_allocator landing

<a id="canonical-2303133311012021-3322130130031103-2120032332031012-3013200021213232-0100310121031123-3230030130121113-3213033330111232-3203303021311023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020333320212020-3132023131321233-1110013311230020-1102200021312230-3113223332313210-2203123212332032-2230223313030311-0022033200201301"></a>

## xcsh_address_allocator — xcsh_address_allocator / 303202003323 / 2

Breadcrumbs:

- xcsh_address_allocator

Manages Address Allocator will create an address allocator object in 'system' namespace of the user
in F5 Distributed Cloud.

<a id="canonical-2031300323000332-2012303120323032-3301121011333002-1103020021302220-2202010100303302-0312231131013011-2233233013223212-3203011110021301"></a>

## Prerequisites — xcsh_address_allocator / 303202003323 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1300200120303211-0323030003120333-1301303012101020-1131032032322001-0133210013030123-3120020030130123-3201213130323300-3012232130011302"></a>

## Minimal configuration — xcsh_address_allocator / 303202003323 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1331212102113112-2301013103212220-3330120031333222-1231013022103023-2203010201022103-3331203102313220-2331323321110322-2102221120211123"></a>

## Root configuration — xcsh_address_allocator / 303202003323 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1333313223300022-3121221330211300-3230112110320230-2110311323130131-3023223130131230-3032000120000001-3131100320220313-2021110311031202"></a>

## Next pages — xcsh_address_allocator / 303202003323 / 6

- [Property reference](../guides/data-sources--address_allocator--reference--group-001.md#canonical-0012311232330222-1032233333322133-2322023003332231-0123033020123120-0101033200130133-1131010113203203-2210031202200112-2212010003230021)
- [Examples](../guides/data-sources--address_allocator--examples--group-001.md#canonical-1312010000001103-1030323302011222-3330121220332231-1131222032331013-0330132110232100-0320310032101302-3103310132020301-3111132203120120)
