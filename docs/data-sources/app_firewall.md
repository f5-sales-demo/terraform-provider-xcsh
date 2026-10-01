---
page_title: "xcsh_app_firewall landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall landing."
---

# xcsh_app_firewall landing

<a id="canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81c5216e91bb5a0f7c70d3677b537bbfff582a3e4006c45646d5b621d4c95ae7"></a>

## xcsh_app_firewall — xcsh_app_firewall / 871bdc4291ce / 2

Breadcrumbs:

- xcsh_app_firewall

Manages Application Firewall in F5 Distributed Cloud.

<a id="canonical-41c70a36cb11f7f68b408c00367c9c83dbb2a578d00e11b5119f89cb720cb123"></a>

## Prerequisites — xcsh_app_firewall / 871bdc4291ce / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `service_policy`.

- service_policy: Fine-grained access control rules

<a id="canonical-8ea2918f58cb54e1e33ed4bf31a2fc865ff820af38773b5e7c3e13665c5c1538"></a>

## Minimal configuration — xcsh_app_firewall / 871bdc4291ce / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```

<a id="canonical-7ac89bdf8efe9a12c89987e6af46a07c5cedaa060ce4c61faf13cd97a4d97975"></a>

## Root configuration — xcsh_app_firewall / 871bdc4291ce / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-488b83b0bad09a7b015a17f47cd91f1de0a7c0f3436dd545088dcba8274c54b6"></a>

## Next pages — xcsh_app_firewall / 871bdc4291ce / 6

- [Property reference](../guides/data-sources--app_firewall--reference--group-001.md#canonical-eccb622733c64544d9dd1f7d76cd33f87ffce68fe0a448e2cdddeaebd2518395)
- [Examples](../guides/data-sources--app_firewall--examples--group-001.md#canonical-dd79e5b08d935c0b78b3339f9166535800c9d5021a648923f2d47769f1d7a913)
