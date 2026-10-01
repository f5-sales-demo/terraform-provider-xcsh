---
page_title: "xcsh_route landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route landing."
---

# xcsh_route landing

<a id="canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210010313031110-0012111102233303-0313330030111233-3210133121300331-3030103210210200-0013232123213000-3211212001110023-0131202130102333"></a>

## xcsh_route — xcsh_route / 001031212231 / 2

Breadcrumbs:

- xcsh_route

Manages route object in a given namespace. Route object is list of route rules. Each rule has match
condition to match incoming requests and actions to take on matching requests in F5 Distributed
Cloud.

<a id="canonical-2211100333330032-2221112323001103-1321120333020312-1121200122232301-0021021022230012-0003023022111203-2012110211022213-2102112211310213"></a>

## Prerequisites — xcsh_route / 001031212231 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3113210323332331-2330112030030322-1123232010121010-1123103223010330-2211201110030131-0130212112310000-1332103021300001-0220311322032101"></a>

## Minimal configuration — xcsh_route / 001031212231 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```

<a id="canonical-2132110333322013-0300112303111332-1030312302311102-2322312301011010-0230232330203000-3111202110111113-3023331200032223-3001232300321310"></a>

## Root configuration — xcsh_route / 001031212231 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2023022312110213-3312113200333320-1001200233032022-3202310122311231-0303301002111103-3103132113013020-2210201001132123-3220022200302120"></a>

## Next pages — xcsh_route / 001031212231 / 6

- [Property reference](../guides/resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [Examples](../guides/resources--route--examples--group-001.md#canonical-1222313231020302-3320201100121100-1033003333200232-3230303210001313-0012333022000220-2202211312132001-3130100332121202-2002322131033003)
- [Import](../guides/resources--route--lifecycle--group-001.md#canonical-1001200300032303-3102230123232020-1333122123132203-0301001023230121-1023222222113310-1130203222113323-1233012033130310-1032112013100301)
- [Timeouts](../guides/resources--route--lifecycle--group-001.md#canonical-3303212101323300-0330010021110113-1102300010010130-0331211123102103-3212101220112312-3212223310220222-0101232310032313-3203310012022003)
