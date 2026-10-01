---
page_title: "xcsh_app_firewall landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall landing."
---

# xcsh_app_firewall landing

<a id="canonical-7288942c4bc7dd68d1b199c1a6bbd0608786af8275ee244016370cf38b0812cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54d040f7f81dfed68ecf3c291fd6cab1e08ee47ab4a7b784d1ca3af2a1f0a7ab"></a>

## xcsh_app_firewall — xcsh_app_firewall / 781bffeb746e / 2

Breadcrumbs:

- xcsh_app_firewall

Manages Application Firewall in F5 Distributed Cloud.

<a id="canonical-0d339c89f92b55f19b1a22711340efa8bee0210453f445753bc298feecd99693"></a>

## Prerequisites — xcsh_app_firewall / 781bffeb746e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `service_policy`.

- service_policy: Fine-grained access control rules

<a id="canonical-ce1fc02c4e9f2bd366369b68630c9da86e97372648726ee8b50213abdbdf2f4b"></a>

## Minimal configuration — xcsh_app_firewall / 781bffeb746e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppFirewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppFirewall configuration
resource "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}
```

<a id="canonical-17b5c22d5cfe1b88a0eba46ef3adb8cc6a512f2fdba61bea6707cc2a0c62c32f"></a>

## Root configuration — xcsh_app_firewall / 781bffeb746e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ebda1f6739ae363d11181a47bbbf97dbdac276a0ef3e617050c575bd4f77328e"></a>

## Next pages — xcsh_app_firewall / 781bffeb746e / 6

- [Property reference](../guides/resources--app_firewall--reference--group-001.md#canonical-2b331ce5e4e05438b56d5edc81a6e4ab7d8db8e595334c3f9d5f968a1367202f)
- [Examples](../guides/resources--app_firewall--examples--group-001.md#canonical-ebf5a99de914d500da858d708980c50528e40c1e229ae5370715e64f911784cd)
- [Import](../guides/resources--app_firewall--lifecycle--group-001.md#canonical-cf84a5d1e51737470e7986531a2f260d38e7cf945ab9327ca16c05c87b0e5aeb)
- [Timeouts](../guides/resources--app_firewall--lifecycle--group-001.md#canonical-21eb372f25eb0e32d3e0872602771cd6f670cea80deb05e42753113e30578ee2)
