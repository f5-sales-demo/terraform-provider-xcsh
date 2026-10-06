---
page_title: "xcsh_app_security_evidence examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_security_evidence examples."
---

# xcsh_app_security_evidence examples

<a id="canonical-1111331131013003-3213311303211110-2313013321313320-3311311023003023-2000200332111220-3022002302310032-1122310003121131-0331331311233212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-3103302000231201-2113100232313221-0102211103022300-2213011330310113-3301220213112210-1312012211213332-0301203231301211-3001330131022123)
- Examples

<a id="canonical-2223012030123202-2212102200022220-0201032132012311-2100233221131331-2232222221302131-0222003030010302-2201232233020022-2110002302203003"></a>

### Complete configurations for `xcsh_app_security_evidence`

- [Data source](data-sources--app_security_evidence--examples--group-001.md#canonical-1021122003121201-1300022010132211-1212102231110310-2122003313212300-1012310133112113-1113332321002312-2112112331331110-2221133121001122): valid configuration.

<a id="canonical-1021122003121201-1300022010132211-1212102231110310-2122003313212300-1012310133112113-1113332321002312-2112112331331110-2221133121001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_app_security_evidence](../data-sources/app_security_evidence.md#canonical-3103302000231201-2113100232313221-0102211103022300-2213011330310113-3301220213112210-1312012211213332-0301203231301211-3001330131022123)
- [Examples](data-sources--app_security_evidence--examples--group-001.md#canonical-1111331131013003-3213311303211110-2313013321313320-3311311023003023-2000200332111220-3022002302310032-1122310003121131-0331331311233212)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_security_evidence/data-source.tf`; digest `sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a`.

```terraform
# AppSecurityEvidence DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_app_security_evidence" "example" {
  namespace = "example-value"
}

output "app_security_evidence_result" {
  value = data.xcsh_app_security_evidence.example
}
```
