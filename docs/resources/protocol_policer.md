---
page_title: "xcsh_protocol_policer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer landing."
---

# xcsh_protocol_policer landing

<a id="canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c50b6badd594111220c27a615f8db1e526cb34badb8c8c478459b35e159c2814"></a>

## xcsh_protocol_policer — xcsh_protocol_policer / ba7fa9b0c89f / 2

Breadcrumbs:

- xcsh_protocol_policer

Manages protocol\_policer object, protocol\_policer object contains list of L4 protocol match
condition and corresponding traffic rate limits in F5 Distributed Cloud.

<a id="canonical-b14f26009c165ddf53eee83b62265dc88c12dae61f5d82268f0f5d0438a5e324"></a>

## Prerequisites — xcsh_protocol_policer / ba7fa9b0c89f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-eb3a3db84ece82898ff3e8674213bf120df9b650502ae54380c716a86af2ea5b"></a>

## Minimal configuration — xcsh_protocol_policer / ba7fa9b0c89f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Resource Example
# Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolPolicer configuration
resource "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}
```

<a id="canonical-15da196c69cdf9a71bb01875176ed5c82453f617e2f570b825e98787e75510b1"></a>

## Root configuration — xcsh_protocol_policer / ba7fa9b0c89f / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-c1ca894e2e6c25907b83bbb32f7941036cff07b2e5303056eedfcaf674d1d0ec"></a>

## Next pages — xcsh_protocol_policer / ba7fa9b0c89f / 6

- [Property reference](../guides/resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [Examples](../guides/resources--protocol_policer--examples--group-001.md#canonical-d8eecb4e2fce187ac1c664867ddead2a9474a00809319f7a26614fa8502ffd78)
- [Import](../guides/resources--protocol_policer--lifecycle--group-001.md#canonical-e96f6cd6f58c70630fe1d81be05a3b47c0a0a245b47314da42e1b729fd21fcdc)
- [Timeouts](../guides/resources--protocol_policer--lifecycle--group-001.md#canonical-0dd40932e49981ca58f04ad535e2dd87d1527151aee29bbb904cf3a7ea3fa9a1)
