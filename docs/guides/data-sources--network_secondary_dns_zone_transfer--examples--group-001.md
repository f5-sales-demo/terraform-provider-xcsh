---
page_title: "xcsh_network_secondary_dns_zone_transfer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_secondary_dns_zone_transfer examples."
---

# xcsh_network_secondary_dns_zone_transfer examples

<a id="canonical-2330121313012102-0110302012330233-3333233112001102-1011223330131202-1111002221302211-0201120022000131-0130201201223323-0222223311310233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-0211000112320000-3032222023300022-3233300332112210-3202010232211322-0100212322311011-2201301123313213-3323022110111321-3110203223320302)
- Examples

<a id="canonical-2323203230100001-3000321112111201-3001103311102233-3103030330313212-0231230123333213-0232331223211132-0013332333011111-0330232322212032"></a>

### Complete configurations for `xcsh_network_secondary_dns_zone_transfer`

- [Data source](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-2310232330102123-1301221133333330-3130113312200000-2230201220022322-0012113313220132-0230202113121010-2102133212132013-3023221133023301): valid configuration.

<a id="canonical-2310232330102123-1301221133333330-3130113312200000-2230201220022322-0012113313220132-0230202113121010-2102133212132013-3023221133023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-0211000112320000-3032222023300022-3233300332112210-3202010232211322-0100212322311011-2201301123313213-3323022110111321-3110203223320302)
- [Examples](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-2330121313012102-0110302012330233-3333233112001102-1011223330131202-1111002221302211-0201120022000131-0130201201223323-0222223311310233)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_secondary_dns_zone_transfer/data-source.tf`; digest `sha256:da49a1e4ce345336b0b68a09a44132872466f1849e8370d3d88d79c58a018362`.

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

data "xcsh_network_secondary_dns_zone_transfer" "authoritative_dns" {}

# The manifest combines transfer and notify sources, so both explicit DNS
# rules use the same published allowlist.
output "secondary_dns_rules" {
  value = [
    {
      direction = "ingress"
      protocol  = "tcp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
    {
      direction = "ingress"
      protocol  = "udp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
  ]
}
```
