---
page_title: "xcsh_proxy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy examples."
---

# xcsh_proxy examples

<a id="canonical-1312000312222230-1120313213032320-3301302120322211-0333033013030102-3131013101001131-1002001223123002-1133320233001021-0211303231120120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- Examples

<a id="canonical-2111103032011123-3302012002231021-3200033022232131-1113333113201213-2300233303312030-0300020103300103-2311032211312232-1222031201333203"></a>

### Complete configurations for `xcsh_proxy`

- [Resource](resources--proxy--examples--group-001.md#canonical-1131121133130333-1031102020110210-0003123302322333-0011132131332003-1313211113301313-1021103110320002-3333022302000102-2301322011030023): valid configuration.

<a id="canonical-1131121133130333-1031102020110210-0003123302322333-0011132131332003-1313211113301313-1021103110320002-3333022302000102-2301322011030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Examples](resources--proxy--examples--group-001.md#canonical-1312000312222230-1120313213032320-3301302120322211-0333033013030102-3131013101001131-1002001223123002-1133320233001021-0211303231120120)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_proxy/resource.tf`; digest `sha256:8d3ef9470714fab754f0f7631e40f5f93a8f5acfb3c6def25dff74cba83c2ef7`.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```
