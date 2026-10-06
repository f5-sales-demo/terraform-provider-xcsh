---
page_title: "xcsh_ike2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 examples."
---

# xcsh_ike2 examples

<a id="canonical-1130101033033222-2212233311300131-1210202130100120-3023323030232133-0213103222002211-1303133001221200-2221113230210131-3112213100212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- Examples

<a id="canonical-0203113102203320-0120122022323311-0333030233230201-2303231011133321-0023330133331221-0333121333013201-0323312322030212-2302100120323101"></a>

### Complete configurations for `xcsh_ike2`

- [Data source](data-sources--ike2--examples--group-001.md#canonical-0032223302130332-1133100133332220-3111102102021312-3033312130120213-3121021201332330-2333213110213301-0312310200330000-3020013312200213): valid configuration.

<a id="canonical-0032223302130332-1133100133332220-3111102102021312-3033312130120213-3121021201332330-2333213110213301-0312310200330000-3020013312200213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-1332033132212110-1332133210322000-0321030302123031-3010310020103003-2122221201132213-2201311321232022-0331322230010032-2233221102111032)
- [Examples](data-sources--ike2--examples--group-001.md#canonical-1130101033033222-2212233311300131-1210202130100120-3023323030232133-0213103222002211-1303133001221200-2221113230210131-3112213100212123)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike2/data-source.tf`; digest `sha256:db87600cf054124d776876313565955caf46f37f46627a7f890a5aaf280c673b`.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```
