---
page_title: "xcsh_network_customer_edge_defaults examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_defaults examples."
---

# xcsh_network_customer_edge_defaults examples

<a id="canonical-3122311111020020-0121221101322000-1332133023312001-2223023020210131-0201032210033320-1030130221333023-3011100312033110-3033021330133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-2030202230323022-0201110111033011-3112302323000311-1231130201223133-0123022101003012-2310121012022212-0120311221210122-2230332213331100)
- Examples

<a id="canonical-3320200331030120-3310010202022132-0001300301031000-3003313322011321-2031300022012220-1022031200002023-0123001232311003-1011313101020202"></a>

### Complete configurations for `xcsh_network_customer_edge_defaults`

- [Data source](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-3001112300000133-3223020012320113-1332100113132213-1311103102113012-0213300302030302-0132301020322312-2000212301213233-0231010313312000): valid configuration.

<a id="canonical-3001112300000133-3223020012320113-1332100113132213-1311103102113012-0213300302030302-0132301020322312-2000212301213233-0231010313312000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-2030202230323022-0201110111033011-3112302323000311-1231130201223133-0123022101003012-2310121012022212-0120311221210122-2230332213331100)
- [Examples](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-3122311111020020-0121221101322000-1332133023312001-2223023020210131-0201032210033320-1030130221333023-3011100312033110-3033021330133130)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_customer_edge_defaults/data-source.tf`; digest `sha256:02dcb0c87e55fb1eb9417d76ac0888bfe3451a7f9f3e9a6cb8857b1ff2aff366`.

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

data "xcsh_network_customer_edge_defaults" "system_services" {}

output "customer_edge_default_egress" {
  value = {
    dns = {
      direction    = "egress"
      protocols    = ["udp", "tcp"]
      port         = 53
      destinations = data.xcsh_network_customer_edge_defaults.system_services.dns_servers
    }
    ntp = {
      direction    = "egress"
      protocols    = ["udp"]
      port         = 123
      destinations = data.xcsh_network_customer_edge_defaults.system_services.ntp_servers
    }
  }
}
```
