---
page_title: "xcsh_infraprotect_mitigation_ips examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_mitigation_ips examples."
---

# xcsh_infraprotect_mitigation_ips examples

<a id="canonical-2230332003132330-3003232213123221-1023033101300311-0101332232032132-1020000311020311-2102222232300101-2130223103131322-2223023103332200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-1223233033223000-1032001012013022-1301320302130033-2311101100202302-3130231002213230-3320300113130330-0113131111131313-3222111203302102)
- Examples

<a id="canonical-0230013233002000-1303120232303011-2123303100310031-3221120123231211-2301123302032330-1003122003010311-3331210203120223-3100303302212132"></a>

### Complete configurations for `xcsh_infraprotect_mitigation_ips`

- [Data source](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-2323222212111121-1020330310021221-2100022220100331-2103120000311113-3220321301201120-0113320021300110-2202332100312021-0113102330213302): valid configuration.

<a id="canonical-2323222212111121-1020330310021221-2100022220100331-2103120000311113-3220321301201120-0113320021300110-2202332100312021-0113102330213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](../data-sources/infraprotect_mitigation_ips.md#canonical-1223233033223000-1032001012013022-1301320302130033-2311101100202302-3130231002213230-3320300113130330-0113131111131313-3222111203302102)
- [Examples](data-sources--infraprotect_mitigation_ips--examples--group-001.md#canonical-2230332003132330-3003232213123221-1023033101300311-0101332232032132-1020000311020311-2102222232300101-2130223103131322-2223023103332200)
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
