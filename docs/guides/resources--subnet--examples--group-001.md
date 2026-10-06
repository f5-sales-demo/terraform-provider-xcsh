---
page_title: "xcsh_subnet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet examples."
---

# xcsh_subnet examples

<a id="canonical-1213033300321211-3212302313120023-3023002000313103-0301300230331231-2033020103022203-2213121013303233-1020023113123100-3013010000012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- Examples

<a id="canonical-0002122310122011-2000202003233220-2003011123322211-0203213033131122-1100331322111021-2013031320233000-0303110112330321-1121313223010231"></a>

### Complete configurations for `xcsh_subnet`

- [Resource](resources--subnet--examples--group-001.md#canonical-0203011010313201-2233101131233311-2003031210233013-0312222212231323-3203110323201001-2231000231202022-2110212322310003-3131202203002022): valid configuration.

<a id="canonical-0203011010313201-2233101131233311-2003031210233013-0312222212231323-3203110323201001-2231000231202022-2110212322310003-3131202203002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_subnet](../resources/subnet.md#canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001)
- [Examples](resources--subnet--examples--group-001.md#canonical-1213033300321211-3212302313120023-3023002000313103-0301300230331231-2033020103022203-2213121013303233-1020023113123100-3013010000012213)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_subnet/resource.tf`; digest `sha256:4b17c21430cf56ad4295d650386f0dba8c0f28367f60769889a0c3628d334ae0`.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```
