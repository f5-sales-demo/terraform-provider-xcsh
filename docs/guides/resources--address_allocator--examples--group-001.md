---
page_title: "xcsh_address_allocator examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator examples."
---

# xcsh_address_allocator examples

<a id="canonical-2122233302132220-3232230300022102-0223303330221223-1202332103132120-1130001332322133-0300033120330001-1310232003003020-1303302301211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
- Examples

<a id="canonical-2300130322020212-1002123233322012-0313213313012311-1233000001332013-0021312211131210-2110222133130132-3300022320101102-3133123030201203"></a>

### Complete configurations for `xcsh_address_allocator`

- [Resource](resources--address_allocator--examples--group-001.md#canonical-2200201023213310-0330310001123331-2033001310110202-2023100221002013-1313323312211120-0002122100113112-1300221101123201-0230210303031011): valid configuration.

<a id="canonical-2200201023213310-0330310001123331-2033001310110202-2023100221002013-1313323312211120-0002122100113112-1300221101123201-0230210303031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322)
- [Examples](resources--address_allocator--examples--group-001.md#canonical-2122233302132220-3232230300022102-0223303330221223-1202332103132120-1130001332322133-0300033120330001-1310232003003020-1303302301211102)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_address_allocator/resource.tf`; digest `sha256:0d0c2e8aebf1264b1b9bd98f9f1564545390bccd08c3989c8f49aed2a844faa2`.

```terraform
# AddressAllocator Resource Example
# Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AddressAllocator configuration
resource "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"

  address_pool = ["example-value"]
}
```
