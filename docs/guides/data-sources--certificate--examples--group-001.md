---
page_title: "xcsh_certificate examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate examples."
---

# xcsh_certificate examples

<a id="canonical-0332313110100301-2123222322221030-3221312133000120-2333200122131330-1111030233012332-3313211030323021-0130210110332303-0312311102101303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- Examples

<a id="canonical-1211330223300012-3330121203103121-3022332000313330-1223122201032020-3221231011133013-0332011300223210-2332102200311010-3132232222222121"></a>

### Complete configurations for `xcsh_certificate`

- [Data source](data-sources--certificate--examples--group-001.md#canonical-0230322312131310-1100223021111133-2013003200312120-3001030011211123-2210201231002213-3101333023233123-2120323312113233-3103212213321302): valid configuration.

<a id="canonical-0230322312131310-1100223021111133-2013003200312120-3001030011211123-2210201231002213-3101333023233123-2120323312113233-3103212213321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010)
- [Examples](data-sources--certificate--examples--group-001.md#canonical-0332313110100301-2123222322221030-3221312133000120-2333200122131330-1111030233012332-3313211030323021-0130210110332303-0312311102101303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate/data-source.tf`; digest `sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e`.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```
