---
page_title: "xcsh_trusted_ca_list examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list examples."
---

# xcsh_trusted_ca_list examples

<a id="canonical-3011330123323123-2011120213211131-0222331311023223-2312020212011321-2112233113021100-0211101033231202-2301121002211210-3303202233210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122)
- Examples

<a id="canonical-0101313221123013-1033121210233031-0200130312122023-2313333222032233-2201122231220001-1131111311313212-0002333121332033-0012322222313300"></a>

### Complete configurations for `xcsh_trusted_ca_list`

- [Data source](data-sources--trusted_ca_list--examples--group-001.md#canonical-3212322202330102-2120312121312322-1122023321111303-2030002312232112-1123333303312102-3010220332113302-1000302221200323-1130000003121232): valid configuration.

<a id="canonical-3212322202330102-2120312121312322-1122023321111303-2030002312232112-1123333303312102-3010220332113302-1000302221200323-1130000003121232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122)
- [Examples](data-sources--trusted_ca_list--examples--group-001.md#canonical-3011330123323123-2011120213211131-0222331311023223-2312020212011321-2112233113021100-0211101033231202-2301121002211210-3303202233210020)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```
