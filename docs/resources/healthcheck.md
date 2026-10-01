---
page_title: "xcsh_healthcheck landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck landing."
---

# xcsh_healthcheck landing

<a id="canonical-f0e0b52fe26a915d9253dab7c548947e75b45cc54cc1365504f77e6116413610"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-748623dd5fff0d7b43b9bc52c6182c9c18d1c6bcb02d1557d41d577300e4fd8d"></a>

## xcsh_healthcheck — xcsh_healthcheck / 409e15f8cb98 / 2

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

<a id="canonical-18a463ffd624cb6e682324131f567a82ca61517da457246ffbd1b0dae6c2d810"></a>

## Prerequisites — xcsh_healthcheck / 409e15f8cb98 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3616f920e092fab3d5163ca3a31db970143093cb935af0afd93e75160c37e6fa"></a>

## Minimal configuration — xcsh_healthcheck / 409e15f8cb98 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Resource Example
# Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to determine if the given endpoint is healthy.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Healthcheck configuration
resource "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"

  healthy_threshold   = 1
  interval            = 1
  timeout             = 1
  unhealthy_threshold = 1
}
```

<a id="canonical-93ca2bc27c0fc11e481d949786dfb21192703edd2af128b61dd1101b01b478cb"></a>

## Root configuration — xcsh_healthcheck / 409e15f8cb98 / 5

Required root properties: `healthy_threshold`, `interval`, `name`, `namespace`, `timeout`, `unhealthy_threshold`. Full root flags and choices appear in the property reference.

<a id="canonical-1715920be30094e1a9f334f1591f0e0dd463d1141671373e464ff7d97d462d77"></a>

## Next pages — xcsh_healthcheck / 409e15f8cb98 / 6

- [Property reference](../guides/resources--healthcheck--reference--group-001.md#canonical-69cf6b124469edd6162003b0efb6990b1520a382f7ed148683df55cab5d76fe7)
- [Examples](../guides/resources--healthcheck--examples--group-001.md#canonical-b7cdb440e7554403280ee57b7ea19ebccc7158e032b753d3ebfc687c05a758d5)
- [Import](../guides/resources--healthcheck--lifecycle--group-001.md#canonical-a0298e08fc5c10a71d5c12cbb2fe59459804b4cc6d4c3d1a21133b4cd099bc6c)
- [Timeouts](../guides/resources--healthcheck--lifecycle--group-001.md#canonical-0766deb6b330da9e608a834ce20f0e9f7e61c23a2b605b7657308feda60bbcd4)
