---
page_title: "xcsh_site_mesh_group landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group landing."
---

# xcsh_site_mesh_group landing

<a id="canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21a6de1062e93aa31c86701055c2653f7eabde2ffeedc16c0c38886753422927"></a>

## xcsh_site_mesh_group — xcsh_site_mesh_group / 2588dc27075d / 2

Breadcrumbs:

- xcsh_site_mesh_group

Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

<a id="canonical-fdf511a1e0c6f91791c84d75b01ad3377a67f7f1d590b0d2b156a17689b6f70d"></a>

## Prerequisites — xcsh_site_mesh_group / 2588dc27075d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `site`.

- site: Sites to include in mesh connectivity

<a id="canonical-4bbc64fb73aa61bdc5cdbe92aa7b7a5c529f5a4786c49391646ad3a759056698"></a>

## Minimal configuration — xcsh_site_mesh_group / 2588dc27075d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteMeshGroup Resource Example
# Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SiteMeshGroup configuration
resource "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}
```

<a id="canonical-4dfdc38ffc5260181ae1aeb0b95d8fc9de4a2cf451d95d43a083f541c6a80631"></a>

## Root configuration — xcsh_site_mesh_group / 2588dc27075d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8003176952c1221f0423521a4b077960479cea34018e6c572cc72a66be4c9665"></a>

## Next pages — xcsh_site_mesh_group / 2588dc27075d / 6

- [Property reference](../guides/resources--site_mesh_group--reference--group-001.md#canonical-1669321b1fb130f863ad42d706ddb0457074e87b455cb3f938485fc05ca5a9d7)
- [Examples](../guides/resources--site_mesh_group--examples--group-001.md#canonical-fb5094d4a1ce4aa8a764458ae4ce2b2b0349dc4af1649ffdfb59f055cb68b77a)
- [Import](../guides/resources--site_mesh_group--lifecycle--group-001.md#canonical-7690561773a4054349b0e4d7e3705498650e27e358f3812098d58f52f36ebeb4)
- [Timeouts](../guides/resources--site_mesh_group--lifecycle--group-001.md#canonical-dc9d5a96ab3367043d38d6ec35224a15bada443db5d769012b9cbb3775746f59)
