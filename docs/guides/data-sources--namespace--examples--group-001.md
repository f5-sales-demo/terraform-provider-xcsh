---
page_title: "xcsh_namespace examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace examples."
---

# xcsh_namespace examples

<a id="canonical-0103010212031000-3110032213333233-1300320330100230-0132203113231333-2303211220001201-2022030333123213-2113111320021123-0131231021231332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112)
- Examples

<a id="canonical-0103012202222321-2033111333112211-3110131331211211-1211110030102031-0132121223320002-3101022303130132-3332331110331031-3101230310210312"></a>

### Complete configurations for `xcsh_namespace`

- [Data source](data-sources--namespace--examples--group-001.md#canonical-0111132001320022-0023302101302010-3320002033232100-2210100313200330-2201112122230113-3012122113121202-1333011311020123-2012030121120300): valid configuration.

<a id="canonical-0111132001320022-0023302101302010-3320002033232100-2210100313200330-2201112122230113-3012122113121202-1333011311020123-2012030121120300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md#canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112)
- [Examples](data-sources--namespace--examples--group-001.md#canonical-0103010212031000-3110032213333233-1300320330100230-0132203113231333-2303211220001201-2022030333123213-2113111320021123-0131231021231332)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_namespace/data-source.tf`; digest `sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492`.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```
