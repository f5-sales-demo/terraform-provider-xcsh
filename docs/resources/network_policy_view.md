---
page_title: "xcsh_network_policy_view landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view landing."
---

# xcsh_network_policy_view landing

<a id="canonical-83b0c3b4c91e37648a731cb08b936fee9da871201453ba08d68bc65bd7e90d13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2067e9e951c436480f3e384a14b1ed8961a3f60b2171ef9816d2e820d50f8d5d"></a>

## xcsh_network_policy_view — xcsh_network_policy_view / ce5b4c94299d / 2

Breadcrumbs:

- xcsh_network_policy_view

Manages a Network Policy View resource in F5 Distributed Cloud for network policy view
specification. configuration.

<a id="canonical-3d319acc29f8a2fccaf6c4e7c6f35f0d8dfc0e8c2a16e673577dcf8e322897d7"></a>

## Prerequisites — xcsh_network_policy_view / ce5b4c94299d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d95a521a9c5551874142f77f7714249892257111a1b65328be79b4d04fe6bcf0"></a>

## Minimal configuration — xcsh_network_policy_view / ce5b4c94299d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyView Resource Example
# Manages a Network Policy View resource in F5 Distributed Cloud for network policy view specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyView configuration
resource "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}
```

<a id="canonical-518a7613089118e5c961ee2759c6441b402a3384e27ef8555f80e242012cf23d"></a>

## Root configuration — xcsh_network_policy_view / ce5b4c94299d / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-6b7dad4e522a95811f1bf17fe3aad9ab8b54139f38eb46cadddd6b12fb464630"></a>

## Next pages — xcsh_network_policy_view / ce5b4c94299d / 6

- [Property reference](../guides/resources--network_policy_view--reference--group-001.md#canonical-34f78dc44e87217dc96dab524ff95a89aa9ae54b480db3fdc19f7688a701f3d1)
- [Examples](../guides/resources--network_policy_view--examples--group-001.md#canonical-15d10fdee37558d71631ffa0a9e48b246c466ebfd04d568290f1d59bec87a79f)
- [Import](../guides/resources--network_policy_view--lifecycle--group-001.md#canonical-3f9ddcf404b2e6ab64ca1a996e42047688e42eee4229f3d3a769d1f67ae44e0b)
- [Timeouts](../guides/resources--network_policy_view--lifecycle--group-001.md#canonical-547ea016533a54df735aad194220a1bfc6c65d4a302a8d6fb4360da1c482153c)
