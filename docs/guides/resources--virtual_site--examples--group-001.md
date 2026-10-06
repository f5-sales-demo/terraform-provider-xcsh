---
page_title: "xcsh_virtual_site examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_virtual_site examples."
---

# xcsh_virtual_site examples

<a id="canonical-2012211220121122-0020110113010220-0003310321031112-3022120221101003-0102001113103033-3213021033123003-3333023021232031-2133232021100201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
- Examples

<a id="canonical-1322001020313311-1112231002123012-2320320220030303-2003111123231333-2333032321121223-3012021031011301-3212222222212013-1302231320310200"></a>

### Complete configurations for `xcsh_virtual_site`

- [Resource](resources--virtual_site--examples--group-001.md#canonical-0212333203312330-3320301102212300-3032103331311011-3321002330020023-2003221033111101-3000301001100322-1220212313313310-1101030333323000): valid configuration.

<a id="canonical-0212333203312330-3320301102212300-3032103331311011-3321002330020023-2003221033111101-3000301001100322-1220212313313310-1101030333323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_virtual_site](../resources/virtual_site.md#canonical-0021200012330233-0110101013310110-3130112023213001-0301030013233201-0302311302331132-3122030310203321-2022313333210003-2333313130301011)
- [Examples](resources--virtual_site--examples--group-001.md#canonical-2012211220121122-0020110113010220-0003310321031112-3022120221101003-0102001113103033-3213021033123003-3333023021232031-2133232021100201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_site/resource.tf`; digest `sha256:d4f2ae5a53db544456cc0750c3ea7b5a8e1e23a8aad03279e065c9ce6ac7f872`.

```terraform
# VirtualSite Resource Example
# Manages virtual site object in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualSite configuration
resource "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}
```
