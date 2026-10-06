---
page_title: "xcsh_virtual_network examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network examples."
---

# xcsh_virtual_network examples

<a id="canonical-0101031222132031-3220122033303210-1331311212220132-3302211031002300-0121302012331202-2313202133033133-2132210103310120-3230232030302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- Examples

<a id="canonical-0200320230111212-1303220110121231-1030103002233331-2132222221313213-3011130302230000-2300302003212002-0011323200230330-1010303121331323"></a>

### Complete configurations for `xcsh_virtual_network`

- [Resource](resources--virtual_network--examples--group-001.md#canonical-1322123120312230-3132110232130120-0302230100113022-0331311030201303-3002210133220221-3332123213020313-3120112033301002-2020030121000232): valid configuration.

<a id="canonical-1322123120312230-3132110232130120-0302230100113022-0331311030201303-3002210133220221-3332123213020313-3120112033301002-2020030121000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_virtual_network](../resources/virtual_network.md#canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001)
- [Examples](resources--virtual_network--examples--group-001.md#canonical-0101031222132031-3220122033303210-1331311212220132-3302211031002300-0121302012331202-2313202133033133-2132210103310120-3230232030302201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_network/resource.tf`; digest `sha256:24ff8ad3e56ba4ece18f3beaaabdcc55805f2e25c74acd17556111d54902f473`.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```
