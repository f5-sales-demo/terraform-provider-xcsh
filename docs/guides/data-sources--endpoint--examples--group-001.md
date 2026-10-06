---
page_title: "xcsh_endpoint examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint examples."
---

# xcsh_endpoint examples

<a id="canonical-2301003011121223-0213312301322023-1222030132322003-1032221331233132-3232021011010330-0121032221322231-2132203021132323-3131313023221333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- Examples

<a id="canonical-1210233301330100-1312112023011133-2221013002311212-1301311222320310-3232001122201231-2121012233121022-0323321111120130-0000301313232222"></a>

### Complete configurations for `xcsh_endpoint`

- [Data source](data-sources--endpoint--examples--group-001.md#canonical-0201022332111231-2211112003210301-2000031231200003-0300022131133233-0120131121000201-3212001123100330-2113220012031232-3021322031332022): valid configuration.

<a id="canonical-0201022332111231-2211112003210301-2000031231200003-0300022131133233-0120131121000201-3212001123100330-2113220012031232-3021322031332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Examples](data-sources--endpoint--examples--group-001.md#canonical-2301003011121223-0213312301322023-1222030132322003-1032221331233132-3232021011010330-0121032221322231-2132203021132323-3131313023221333)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_endpoint/data-source.tf`; digest `sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b`.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```
