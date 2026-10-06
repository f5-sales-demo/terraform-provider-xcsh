---
page_title: "xcsh_bgp examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp examples."
---

# xcsh_bgp examples

<a id="canonical-0023113332102111-1003103310223221-0311022030221000-2202300320022011-1203200001201133-2111200033013301-1230112300130222-1331122202233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- Examples

<a id="canonical-2102110000132213-1111003003201200-2103022131331302-3303112230001033-1213332110110311-1013123103020110-2002333211013311-2231112103233003"></a>

### Complete configurations for `xcsh_bgp`

- [Resource](resources--bgp--examples--group-001.md#canonical-2310210011323101-1213010202332001-0300113302132302-2013031322133103-3330212310300001-3122322121201332-1203200131333113-1313200031001223): valid configuration.

<a id="canonical-2310210011323101-1213010202332001-0300113302132302-2013031322133103-3330212310300001-3122322121201332-1203200131333113-1313200031001223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Examples](resources--bgp--examples--group-001.md#canonical-0023113332102111-1003103310223221-0311022030221000-2202300320022011-1203200001201133-2111200033013301-1230112300130222-1331122202233022)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp/resource.tf`; digest `sha256:7d85805409033bf09b03057cbf852697d9f9c443321c552ee56fb1009994fdd5`.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```
