---
page_title: "xcsh_network_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_log_receiver landing."
---

# xcsh_network_global_log_receiver landing

<a id="canonical-251b976c87f6429c83905ed8e699743bc66f5d388ccb287b6ab1dc588898d9a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-072a4bf94e8c9bf7b3f26a858e566bf463afd8746eb28bf870a0c3a98f64fe23"></a>

## xcsh_network_global_log_receiver — xcsh_network_global_log_receiver / 5dabd7b0d55f / 2

Breadcrumbs:

- xcsh_network_global_log_receiver

Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses.
Values are bundled from the pinned OpenAPI release; this data source performs no network request.
Ports and traffic direction are not encoded in the manifest.

<a id="canonical-c18b14fd1a0b27815177cd003982d5dbba5e7bbffad73029701a33aaee36cd24"></a>

## Prerequisites — xcsh_network_global_log_receiver / 5dabd7b0d55f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-bd6ac7f16e3119c46f715e95b422db9c95c7f98237913bbdaf333ba2ab108872"></a>

## Minimal configuration — xcsh_network_global_log_receiver / 5dabd7b0d55f / 4

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

data "xcsh_network_global_log_receiver" "receivers" {}

# This example chooses TLS syslog on TCP 6514. The manifest supplies only
# destinations; choose the port required by the configured log receiver.
output "tls_syslog_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 6514
    destinations = data.xcsh_network_global_log_receiver.receivers.cidr_blocks
  }
}
```

<a id="canonical-5ec7a6c928b11170728edaa0f15acbf15ef9e4cb55ba464015d4a136cc6b79c3"></a>

## Root configuration — xcsh_network_global_log_receiver / 5dabd7b0d55f / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-9103fc73b15e7a4ffea7a2337df2f3b6ef03c6ad009ebd7e4f7d43ce850203d7"></a>

## Next pages — xcsh_network_global_log_receiver / 5dabd7b0d55f / 6

- [Property reference](../guides/data-sources--network_global_log_receiver--reference--group-001.md#canonical-72b74013a2b858f443b06606f41d24e6214625ccc7f7916d9f905b8d9e85df4f)
- [Examples](../guides/data-sources--network_global_log_receiver--examples--group-001.md#canonical-62c4b37833b41dd8f8d601a10f553355c910f426180fe77dc01774450591215d)
