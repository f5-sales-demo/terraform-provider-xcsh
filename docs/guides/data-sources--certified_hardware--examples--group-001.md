---
page_title: "xcsh_certified_hardware examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_certified_hardware examples."
---

# xcsh_certified_hardware examples

<a id="canonical-3211013330020133-2220002032011211-2213232320113221-0321132313332230-1131300330100223-1100020132102113-1120013311233100-2101331022200313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-3203100033102110-1200003111311001-2130133202021111-2000112331003330-0223333332201310-0221311220013120-2000203133203030-3221223311122122)
- Examples

<a id="canonical-1201023123020112-2310122222031333-2231033231033012-2222032010332033-3223322221113230-2310020301303000-2012210330322330-0213301222312311"></a>

### Complete configurations for `xcsh_certified_hardware`

- [Data source](data-sources--certified_hardware--examples--group-001.md#canonical-1001223122213112-2321321330020223-3001031021223130-0312000001213120-2202031212330020-1033003021022202-3021321100302202-2313303211212213): valid configuration.

<a id="canonical-1001223122213112-2321321330020223-3001031021223130-0312000001213120-2202031212330020-1033003021022202-3021321100302202-2313303211212213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_certified_hardware](../data-sources/certified_hardware.md#canonical-3203100033102110-1200003111311001-2130133202021111-2000112331003330-0223333332201310-0221311220013120-2000203133203030-3221223311122122)
- [Examples](data-sources--certified_hardware--examples--group-001.md#canonical-3211013330020133-2220002032011211-2213232320113221-0321132313332230-1131300330100223-1100020132102113-1120013311233100-2101331022200313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certified_hardware/data-source.tf`; digest `sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba`.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```
