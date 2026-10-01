---
page_title: "xcsh_network_regional_edges landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_regional_edges landing."
---

# xcsh_network_regional_edges landing

<a id="canonical-1023001310210202-1103332033010032-1201022013330111-2231303022312333-3020121223333020-2332032002032323-0213232310320002-2311311301002021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001323133013232-0033223223321203-1121202202222100-1023331311302110-2210333203131300-3103313103020311-1020321113013201-3102313201223201"></a>

## xcsh_network_regional_edges — xcsh_network_regional_edges / 332211112300 / 2

Breadcrumbs:

- xcsh_network_regional_edges

Regional Edge IPv4 networks for origin ingress allowlists. Values are bundled from the pinned
OpenAPI release; this data source performs no network request. Ports and traffic direction are not
encoded in the manifest.

<a id="canonical-0210130030000031-2003310210230103-1310303220021112-3331220211210323-0013213033013112-2323100203112220-1333312201000301-0021330032332211"></a>

## Prerequisites — xcsh_network_regional_edges / 332211112300 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3203032220331321-1220310330112030-2202122131013312-1002222232303120-1123230133001211-1031203223030112-1330303132001203-2300130103331013"></a>

## Minimal configuration — xcsh_network_regional_edges / 332211112300 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```

<a id="canonical-2210211113122030-0110030231031310-0000112311021323-2301013200032310-2131013120112300-1222332112023103-3203010000111032-2030002301322302"></a>

## Root configuration — xcsh_network_regional_edges / 332211112300 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2200111311320121-3001210100132331-3200113322121202-1223313002102320-3232022312302203-2000233202102311-0112200102102303-0111113312233202"></a>

## Next pages — xcsh_network_regional_edges / 332211112300 / 6

- [Property reference](../guides/data-sources--network_regional_edges--reference--group-001.md#canonical-3200132122123211-1311200021302313-1330223222012333-2301330001203121-2032303231120232-2032111121003012-2020302202111221-0022103200132122)
- [Examples](../guides/data-sources--network_regional_edges--examples--group-001.md#canonical-1221302330021112-3331231200012232-2213221100101223-3011300131133211-3312233111013301-1200030210201233-1212033110031130-2103232233000132)
