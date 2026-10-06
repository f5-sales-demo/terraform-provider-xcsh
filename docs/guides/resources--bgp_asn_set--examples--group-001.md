---
page_title: "xcsh_bgp_asn_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set examples."
---

# xcsh_bgp_asn_set examples

<a id="canonical-1311312101312312-0003113031020322-3000212100121303-1310231031111133-0303022310101211-1323001112323220-1210220310300300-1033220312010023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-3010010321231122-0133331023021212-0330320220102302-0102122320130321-1110011002112302-3010211202300010-0031103320123032-3230133030121103)
- Examples

<a id="canonical-3100211323323312-3323300020330032-3021122331033012-3332220033310231-2110103121332232-1111003310002113-0001203011133112-2333300022111210"></a>

### Complete configurations for `xcsh_bgp_asn_set`

- [Resource](resources--bgp_asn_set--examples--group-001.md#canonical-1220303231033322-1300223213110123-0311320313311132-0000220100013200-3111113020032331-3200100303032121-2202012123302023-0020113113101020): valid configuration.

<a id="canonical-1220303231033322-1300223213110123-0311320313311132-0000220100013200-3111113020032331-3200100303032121-2202012123302023-0020113113101020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_bgp_asn_set](../resources/bgp_asn_set.md#canonical-3010010321231122-0133331023021212-0330320220102302-0102122320130321-1110011002112302-3010211202300010-0031103320123032-3230133030121103)
- [Examples](resources--bgp_asn_set--examples--group-001.md#canonical-1311312101312312-0003113031020322-3000212100121303-1310231031111133-0303022310101211-1323001112323220-1210220310300300-1033220312010023)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_asn_set/resource.tf`; digest `sha256:3c2bb9b5fa3b8555e63a99053456d35fa40e857a94daa94de302421615af065a`.

```terraform
# BGPAsnSet Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPAsnSet configuration
resource "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"

  as_numbers = [1]
}
```
