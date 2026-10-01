---
page_title: "xcsh_api_definition landing"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition landing."
---

# xcsh_api_definition landing

<a id="canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6262f757d73a14162106f01ac2d9f1d1de3bd681f6d3455c11e588b2fbe3d6ff"></a>

## xcsh_api_definition — xcsh_api_definition / 32599d6a2e47 / 2

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

<a id="canonical-fe0f80e5333ab2473c347f7a37f5721c46561e0b95ab2e13222bfb78036d57df"></a>

## Prerequisites — xcsh_api_definition / 32599d6a2e47 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

<a id="canonical-35968b4ad3ff9bb4325861046f436ce507c5877b143b137d71e0b44731bad2f0"></a>

## Minimal configuration — xcsh_api_definition / 32599d6a2e47 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```

<a id="canonical-537a27f66c939547f7a39fabae80231f6984474f887e03b639ef67d876b29946"></a>

## Root configuration — xcsh_api_definition / 32599d6a2e47 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2221ece52a46618f39480ca31034f901f9153e79a1ec387f70f01d9a22f2e8a4"></a>

## Next pages — xcsh_api_definition / 32599d6a2e47 / 6

- [Property reference](../guides/data-sources--api_definition--reference--group-001.md#canonical-fad0050df3f4cd1ee16f15782b363b846ffd70034b8a18de354e7a072ec26d5b)
- [Examples](../guides/data-sources--api_definition--examples--group-001.md#canonical-4f65221eb8214fce92c087c6dba0a758d6c4edc5500cdffac6f34a83f2e26a7d)
