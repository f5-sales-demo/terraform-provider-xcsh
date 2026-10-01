---
page_title: "xcsh_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery landing."
---

# xcsh_discovery landing

<a id="canonical-a1ba1cd74df843e69061f566e3d6b3bf2ffdbb044874803148f03446cc9ffbea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1a93f8feeb074538e27cb7606fed60de97dc5ec636b7e038e250cbc95ee51d5"></a>

## xcsh_discovery — xcsh_discovery / 33de663a3ad8 / 2

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

<a id="canonical-49b6f1cdd0056d3aeb0056b6789cd7aa2be5eaf81bedfbc2ae53a6768244a8da"></a>

## Prerequisites — xcsh_discovery / 33de663a3ad8 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b369f0803e5b329bf7720c63be04d6513af23c57cec015bc3eba39135df05b74"></a>

## Minimal configuration — xcsh_discovery / 33de663a3ad8 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

<a id="canonical-751e88136324a64a2fd0c93558a5e19f0d7d5f4ea4c2492750be5ae6592d03e5"></a>

## Root configuration — xcsh_discovery / 33de663a3ad8 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0d87496774937d2f196b6ba9a022566845551041ac97d1e20c64e3c4721d5933"></a>

## Next pages — xcsh_discovery / 33de663a3ad8 / 6

- [Property reference](../guides/data-sources--discovery--reference--group-001.md#canonical-89b3e09314ef5fec837c9034f68d6bb66048b1e8e127c21bebfbd751a57f9668)
- [Examples](../guides/data-sources--discovery--examples--group-001.md#canonical-af5f4883c598ee98395114fc54a3127f97b306ab0e56ee5dcd42b30eefd63773)
