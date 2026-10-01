---
page_title: "xcsh_protocol_policer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer landing."
---

# xcsh_protocol_policer landing

<a id="canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e612021899d3dd8b69a1dc0c047506eb7aa49cd0b11075671ad68c7dac5164d0"></a>

## xcsh_protocol_policer — xcsh_protocol_policer / 091d92a3fec6 / 2

Breadcrumbs:

- xcsh_protocol_policer

Manages protocol\_policer object, protocol\_policer object contains list of L4 protocol match
condition and corresponding traffic rate limits in F5 Distributed Cloud.

<a id="canonical-c4275b15a475c405d90d0ba6042ae4230cb623d79132948b68f17e925dcf0b59"></a>

## Prerequisites — xcsh_protocol_policer / 091d92a3fec6 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-5354e2d97ae00c88a5ce40d1178f23ee85c974640bd10805c1baedc1d229888f"></a>

## Minimal configuration — xcsh_protocol_policer / 091d92a3fec6 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

<a id="canonical-4636c6383f0684699d65afbbb074d616d79450a90fd60b758fc32ff1cb63b277"></a>

## Root configuration — xcsh_protocol_policer / 091d92a3fec6 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-298521c671414686fa31abdb6a6d0f9818ad2d415a0a304dcd9f27a95c181b27"></a>

## Next pages — xcsh_protocol_policer / 091d92a3fec6 / 6

- [Property reference](../guides/data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [Examples](../guides/data-sources--protocol_policer--examples--group-001.md#canonical-67a0ff220c34dd46fc92c2f52ddfe0d0c4c17068f875b188345d25d5061dbbab)
