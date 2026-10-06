---
page_title: "xcsh_trusted_ca_list examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list examples."
---

# xcsh_trusted_ca_list examples

<a id="canonical-1313012232330110-3010303033130202-2023002102131213-1311200033113301-3030011313223302-0010020201221102-3110312100122103-0133111300230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-0222000113113203-2212232323102203-2101230202220010-2030322030013021-2133220023123121-2113200100110031-3003002313103330-2200302323113112)
- Examples

<a id="canonical-1232233331001322-1113213303011033-2011300021213301-2111130032312112-0202021300133023-2232112130330021-0133122310001300-0123302301111303"></a>

### Complete configurations for `xcsh_trusted_ca_list`

- [Resource](resources--trusted_ca_list--examples--group-001.md#canonical-2103021213110323-1232123021100101-0100111023203031-2120302130122313-1013031001112213-0112213033332111-0122031310202201-0001213323302301): valid configuration.

<a id="canonical-2103021213110323-1232123021100101-0100111023203031-2120302130122313-1013031001112213-0112213033332111-0122031310202201-0001213323302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-0222000113113203-2212232323102203-2101230202220010-2030322030013021-2133220023123121-2113200100110031-3003002313103330-2200302323113112)
- [Examples](resources--trusted_ca_list--examples--group-001.md#canonical-1313012232330110-3010303033130202-2023002102131213-1311200033113301-3030011313223302-0010020201221102-3110312100122103-0133111300230220)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_trusted_ca_list/resource.tf`; digest `sha256:98ef39907d2777d4837a31daa9bde5346e1139b6bb15f38e7c1bd69f44aac328`.

```terraform
# TrustedCAList Resource Example
# Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list management.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TrustedCAList configuration
resource "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}
```
