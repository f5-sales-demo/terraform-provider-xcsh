---
page_title: "xcsh_network_customer_edge_defaults examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_defaults examples."
---

# xcsh_network_customer_edge_defaults examples

<a id="canonical-dad5520819a51e807e7cbd81ab2c891d213a43f84c729fcbc54363d4cf27c7dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f883d318f412229e01c31340c3dfa1798dc0a1a84a36008b1b06ed4345dd1222"></a>

## Examples — Examples / 8147ea326928 / 2

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-8c8aceca215153c5d6cbb0356d721adf1b2910c6b46462a618d6991aacfa7f50)
- Examples

<a id="canonical-d52412e970653dce38e82c2e9c9511d48bc0a8c46e1c2c713e79d432d86620df"></a>

## Complete configurations — Examples / 8147ea326928 / 3

- [Data source](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-c15b001feb206e177e4177a7754d25c627c323321ec48eb6809b19ef2d137d80): valid configuration.

<a id="canonical-f27933d40d9441611c45afe24aa76a534f157ef0e26d321082e3e450d7e3e8eb"></a>

## Next pages — Examples / 8147ea326928 / 4

- [Data source](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-c15b001feb206e177e4177a7754d25c627c323321ec48eb6809b19ef2d137d80)
- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-8c8aceca215153c5d6cbb0356d721adf1b2910c6b46462a618d6991aacfa7f50)

<a id="canonical-c15b001feb206e177e4177a7754d25c627c323321ec48eb6809b19ef2d137d80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6f6952fae3ffa9a0aafb53d0fb6f61a38efbbd4aea12fe24c8eae67deb70767"></a>

## Data source — Data source / a60575b85515 / 2

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-8c8aceca215153c5d6cbb0356d721adf1b2910c6b46462a618d6991aacfa7f50)
- [Examples](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-dad5520819a51e807e7cbd81ab2c891d213a43f84c729fcbc54363d4cf27c7dc)
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

<a id="canonical-06e180ccaf7c1cadc8b9d1a904394569f8c82e350920adedf67e3b22b503a1d9"></a>

## Next pages — Data source / a60575b85515 / 3

- [Examples](data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-dad5520819a51e807e7cbd81ab2c891d213a43f84c729fcbc54363d4cf27c7dc)
- [xcsh_network_customer_edge_defaults](../data-sources/network_customer_edge_defaults.md#canonical-8c8aceca215153c5d6cbb0356d721adf1b2910c6b46462a618d6991aacfa7f50)
