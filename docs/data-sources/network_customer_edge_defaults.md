---
page_title: "xcsh_network_customer_edge_defaults landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_defaults landing."
---

# xcsh_network_customer_edge_defaults landing

<a id="canonical-8c8aceca215153c5d6cbb0356d721adf1b2910c6b46462a618d6991aacfa7f50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b60c317f0f7f400b27583c9f2440dd15ccca8e34df73c9b388685f01b9c99a"></a>

## xcsh_network_customer_edge_defaults — xcsh_network_customer_edge_defaults / 394b47679358 / 2

Breadcrumbs:

- xcsh_network_customer_edge_defaults

Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

<a id="canonical-71ee489f53cb23edfcc31c48c543c9fe060226dd03af2a2099c437f8d333ae8f"></a>

## Prerequisites — xcsh_network_customer_edge_defaults / 394b47679358 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-275890c671e8ca80543695fec45a8439dbe2aea886d67c74b3d01ead32762fb2"></a>

## Minimal configuration — xcsh_network_customer_edge_defaults / 394b47679358 / 4

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

<a id="canonical-9bbf0c30dad858bd869df6e3c3edf4d8f4276a65cdd19f9b79d6b867799fcf05"></a>

## Root configuration — xcsh_network_customer_edge_defaults / 394b47679358 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3086729dc0bfdc3fcf8df8f47a924ac138dac956a1e4ea292e2be01a0b605f94"></a>

## Next pages — xcsh_network_customer_edge_defaults / 394b47679358 / 6

- [Property reference](../guides/data-sources--network_customer_edge_defaults--reference--group-001.md#canonical-9d1fe2e704b8f5c9223fc878a892d37925ab8d4550c772fa254b3648116a6364)
- [Examples](../guides/data-sources--network_customer_edge_defaults--examples--group-001.md#canonical-dad5520819a51e807e7cbd81ab2c891d213a43f84c729fcbc54363d4cf27c7dc)
