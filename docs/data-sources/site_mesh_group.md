---
page_title: "xcsh_site_mesh_group landing"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group landing."
---

# xcsh_site_mesh_group landing

<a id="canonical-06f1ce29c6f54c5a003416d320f4de9faa10a30dfd60441f817101f8dba26e22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcd51bda7060cdecaa08535f06e86a4547e62df5f2ca6826dd87d7f7162ac92e"></a>

## xcsh_site_mesh_group — xcsh_site_mesh_group / 648bfe7b24ea / 2

Breadcrumbs:

- xcsh_site_mesh_group

Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

<a id="canonical-2cc33d3f198ea874250a77db0bd722db93249604a445b8a559ac68b7b86c02bf"></a>

## Prerequisites — xcsh_site_mesh_group / 648bfe7b24ea / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `site`.

- site: Sites to include in mesh connectivity

<a id="canonical-1743ee116f47f5298f4125782e73a3812146c18fffe41862870feec1f2e4706b"></a>

## Minimal configuration — xcsh_site_mesh_group / 648bfe7b24ea / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteMeshGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SiteMeshGroup by name
data "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}

output "site_mesh_group_id" {
  value = data.xcsh_site_mesh_group.example.id
}
```

<a id="canonical-b288576ab93374ee3435b9ac2d211fd321a3c2d8937167ba8eb5f8bbe6473b47"></a>

## Root configuration — xcsh_site_mesh_group / 648bfe7b24ea / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-036f28fbd12aec530d7c0294b02d7b61700f86d9f7a0af7a6448ebd0aa3b8c2d"></a>

## Next pages — xcsh_site_mesh_group / 648bfe7b24ea / 6

- [Property reference](../guides/data-sources--site_mesh_group--reference--group-001.md#canonical-c98fa646f5fd6d8327e7c22d463f059277517a5dc88daa59ed45a3712907cd05)
- [Examples](../guides/data-sources--site_mesh_group--examples--group-001.md#canonical-d4c40650ecf86c580d3877f661b8ae24c251d4c0650148c0b1cc19fe8c3b4076)
