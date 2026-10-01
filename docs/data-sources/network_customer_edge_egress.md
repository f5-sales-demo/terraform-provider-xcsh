---
page_title: "xcsh_network_customer_edge_egress landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_egress landing."
---

# xcsh_network_customer_edge_egress landing

<a id="canonical-3220221103330020-3110221000233302-0211101212002003-1012001301133302-1131222331211321-2333221222011322-0330222303230333-2322220210010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130030120210132-1110133211303022-0022103310121003-2313333220313220-1112202100111223-3200133031330112-0332112133301100-2212210013303030"></a>

## xcsh_network_customer_edge_egress — xcsh_network_customer_edge_egress / 030213233033 / 2

Breadcrumbs:

- xcsh_network_customer_edge_egress

Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are
intentionally excluded. Values are bundled from the pinned OpenAPI release; this data source
performs no network request. Ports and traffic direction are not encoded in the manifest.

<a id="canonical-2010133323202210-3031300103220313-1110213310000212-0212223102211032-2121023222122333-1201202022023110-1113021002313213-3210133123133031"></a>

## Prerequisites — xcsh_network_customer_edge_egress / 030213233033 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1101102020220111-3233021131331200-3003030123001030-1332322010322121-3323011121000033-0220201303213232-1221110210301330-3121113200330310"></a>

## Minimal configuration — xcsh_network_customer_edge_egress / 030213233033 / 4

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

data "xcsh_network_customer_edge_egress" "secure_mesh_v2" {}

# Use the domains with an FQDN-aware control. The legacy CE branch is not
# included in this data source.
output "secure_mesh_v2_https_egress" {
  value = {
    direction              = "egress"
    protocol               = "tcp"
    port                   = 443
    registration_addresses = data.xcsh_network_customer_edge_egress.secure_mesh_v2.registration_addresses
    domains                = data.xcsh_network_customer_edge_egress.secure_mesh_v2.domains
  }
}
```

<a id="canonical-2222011310002330-0321303212120121-2022010121311222-2223103100222211-0303112221202221-3002023033212230-3211230022110001-2130123031220110"></a>

## Root configuration — xcsh_network_customer_edge_egress / 030213233033 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0110203020032232-0212322113132011-0112313232321020-1103230032202333-1020222211311330-1211321020131300-1003031312231123-1000101213232121"></a>

## Next pages — xcsh_network_customer_edge_egress / 030213233033 / 6

- [Property reference](../guides/data-sources--network_customer_edge_egress--reference--group-001.md#canonical-3330330333211022-3210001223230133-2222131201231022-2130313303130101-2303311233122033-0012211213111320-0130220120201232-0201122200232020)
- [Examples](../guides/data-sources--network_customer_edge_egress--examples--group-001.md#canonical-1333101203321301-3201210223302323-2122211312232210-3013232221331311-2120221213012230-0300222231222110-1301202011101010-2223303330333200)
