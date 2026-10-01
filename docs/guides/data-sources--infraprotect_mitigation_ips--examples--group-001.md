---
page_title: "xcsh_infraprotect_mitigation_ips examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_mitigation_ips examples."
---

# xcsh_infraprotect_mitigation_ips examples

<a id="canonical-acf837bcc3ba76e94b3d1c3511fae39e4803523592aaec119cad377aab2d3fa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c1ef0807362ecc59bcd0d0de961bb65b16f23bc43683135fd92362bd0cf299e"></a>

## Examples — Examples / a9c38f251b4c / 2

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
- Examples

<a id="canonical-ab921825b41067050c9d2ec9085b87738a6df7ffa289018fef933562de15ac09"></a>

## Complete configurations — Examples / a9c38f251b4c / 3

- [Data source](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-bbaa655948f34269902a843d93600d57e8e7185817e09c14a2f90d89174bc9f2): valid configuration.

<a id="canonical-a3b85a3eadfd0c59881fa55bdcf28df682d8e86bc714accd77a68ec6d3bbfdf9"></a>

## Next pages — Examples / a9c38f251b4c / 4

- [Data source](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-bbaa655948f34269902a843d93600d57e8e7185817e09c14a2f90d89174bc9f2)
- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)

<a id="canonical-bbaa655948f34269902a843d93600d57e8e7185817e09c14a2f90d89174bc9f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70050bfce50434561ff73ebca47f0a0ad7f423c8be8e38dcab52382971c1c9df"></a>

## Data source — Data source / 665fefb88833 / 2

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
- [Examples](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-acf837bcc3ba76e94b3d1c3511fae39e4803523592aaec119cad377aab2d3fa0)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf`; digest `sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d`.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

<a id="canonical-1a6275899fd2ebcb0b9da4ebdcde45471dec32d0a2c7bd7338378213081b7659"></a>

## Next pages — Data source / 665fefb88833 / 3

- [Examples](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-acf837bcc3ba76e94b3d1c3511fae39e4803523592aaec119cad377aab2d3fa0)
- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-6bbcfac04e0461ca71e3270fb54508b2dcb429ecf8c1773c17755777ea563c92)
