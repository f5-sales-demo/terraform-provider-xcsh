---
page_title: "xcsh_app_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type examples."
---

# xcsh_app_type examples

<a id="canonical-2102211203220302-1231033312102010-0032320200033323-2012000223133230-3112033310312301-3331003232031021-0230112012023120-3302011322211012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- Examples

<a id="canonical-3300203312133331-3013103211330201-2133201210101121-0301331022311222-1202023123201303-1320312102111302-3012212001023200-0121313301020320"></a>

### Complete configurations for `xcsh_app_type`

- [Data source](data-sources--app_type--examples--group-001.md#canonical-2120332221223021-0023301011011233-2130011301010102-0211320120213212-0002010313231021-2230011131203330-0102221100001033-1300033301010133): valid configuration.

<a id="canonical-2120332221223021-0023301011011233-2130011301010102-0211320120213212-0002010313231021-2230011131203330-0102221100001033-1300033301010133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md#canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130)
- [Examples](data-sources--app_type--examples--group-001.md#canonical-2102211203220302-1231033312102010-0032320200033323-2012000223133230-3112033310312301-3331003232031021-0230112012023120-3302011322211012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_type/data-source.tf`; digest `sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933`.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```
