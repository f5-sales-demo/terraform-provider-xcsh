---
page_title: "xcsh_api_testing landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing landing."
---

# xcsh_api_testing landing

<a id="canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033102311002231-0212131312132223-0312233100030312-1231012333001220-2013301101012303-2033323121120303-0102200223212103-3000003201232023"></a>

## xcsh_api_testing — xcsh_api_testing / 322222302231 / 2

Breadcrumbs:

- xcsh_api_testing

Manages a API Testing resource in F5 Distributed Cloud.

<a id="canonical-0001331130301010-2212031133230332-2321112123013030-0210123113101321-2102103111223022-3201323331221123-3230000330203233-2113010212301313"></a>

## Prerequisites — xcsh_api_testing / 322222302231 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3230012021322313-1111011332003131-0133332302133031-0223103102210100-3201133110023202-0132303000320230-0333033332001201-2222123202121231"></a>

## Minimal configuration — xcsh_api_testing / 322222302231 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```

<a id="canonical-0211333101022000-1321322303321000-1103110201032332-1033022112020003-3122031101323133-0023320113321312-0231221200220101-3230303330322332"></a>

## Root configuration — xcsh_api_testing / 322222302231 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1130111200000320-2300203111101203-1111011203030021-3300132313123023-1203133210022231-3111330333301112-2121023020321111-3332102110321103"></a>

## Next pages — xcsh_api_testing / 322222302231 / 6

- [Property reference](../guides/data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [Examples](../guides/data-sources--api_testing--examples--group-001.md#canonical-2332231100103211-2032311033113031-3002312201302303-1303231131130130-0310212022223232-1201123210232101-0230031123302302-0012222311112303)
