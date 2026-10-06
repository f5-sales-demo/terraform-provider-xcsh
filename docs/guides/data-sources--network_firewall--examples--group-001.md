---
page_title: "xcsh_network_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_firewall examples."
---

# xcsh_network_firewall examples

<a id="canonical-1212323312020232-0200032322020210-1210213001133311-3210113331033111-1131201322032200-2233231033030332-1022333021112211-3100311203333131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- Examples

<a id="canonical-3031211012113222-3100130013010232-2301022200133221-3130023323220213-0223121003020222-0220323002223101-2301332113213330-3300022221021313"></a>

### Complete configurations for `xcsh_network_firewall`

- [Data source](data-sources--network_firewall--examples--group-001.md#canonical-3300332103213333-1310213100231123-2021322103001210-2022331312022013-2332313010303132-1230222130330123-2121010332210230-0311321221221113): valid configuration.

<a id="canonical-3300332103213333-1310213100231123-2021322103001210-2022331312022013-2332313010303132-1230222130330123-2121010332210230-0311321221221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_firewall](../data-sources/network_firewall.md#canonical-2123120223131230-3313320221210020-1323113310210220-3023322120003333-2331320233121110-1003002203212011-3023303311023211-2310331120221112)
- [Examples](data-sources--network_firewall--examples--group-001.md#canonical-1212323312020232-0200032322020210-1210213001133311-3210113331033111-1131201322032200-2233231033030332-1022333021112211-3100311203333131)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_firewall/data-source.tf`; digest `sha256:5641721d230eef959a292217d55745f5f8be26fc7ce1feff6b83cbca8e44625c`.

```terraform
# NetworkFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkFirewall by name
data "xcsh_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

output "network_firewall_id" {
  value = data.xcsh_network_firewall.example.id
}
```
