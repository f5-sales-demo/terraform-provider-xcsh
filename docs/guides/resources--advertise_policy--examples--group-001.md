---
page_title: "xcsh_advertise_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy examples."
---

# xcsh_advertise_policy examples

<a id="canonical-1222032302121303-1310301213122232-0312232232033120-3310223121230221-0301130033022321-0320133021320102-3021001333220022-1013311233113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- Examples

<a id="canonical-2313123001031322-3221233023120133-3000033002202311-3021012032232331-2323200322101210-1123120110013132-0203123232131230-2113321130230113"></a>

### Complete configurations for `xcsh_advertise_policy`

- [Resource](resources--advertise_policy--examples--group-001.md#canonical-1130302312201231-2213332000223123-2201301302211121-1220330133313300-3102122020131231-2010302111013101-2113320200031320-0222303330222321): valid configuration.

<a id="canonical-1130302312201231-2213332000223123-2201301302211121-1220330133313300-3102122020131231-2010302111013101-2113320200031320-0222303330222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_advertise_policy](../resources/advertise_policy.md#canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000)
- [Examples](resources--advertise_policy--examples--group-001.md#canonical-1222032302121303-1310301213122232-0312232232033120-3310223121230221-0301130033022321-0320133021320102-3021001333220022-1013311233113223)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_advertise_policy/resource.tf`; digest `sha256:9e9301fb78d45973929d2d48eacb74419a673af009dabc983da6f509200329a9`.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```
