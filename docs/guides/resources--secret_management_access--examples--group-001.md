---
page_title: "xcsh_secret_management_access examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access examples."
---

# xcsh_secret_management_access examples

<a id="canonical-958b1534377812bddb2136dcb193677d4717b4eb90874cc67fbfb236ca549673"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71290882eebe000e59f44c79af950ae5d389ba44e925daa70dde5bf303c652f0"></a>

## Examples — Examples / 4e0f0b8ce83e / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- Examples

<a id="canonical-50390f7d01df1e420f1472e4a59890fc0122d154ceb5d8db77d2cf6603d208b7"></a>

## Complete configurations — Examples / 4e0f0b8ce83e / 3

- [Resource](resources--secret_management_access--examples--group-001.md#canonical-33b225212396fd50f2735ac0ab47e65df6dc6bbc5d94bcb5210c714bbd6432fa): valid configuration.

<a id="canonical-882a9a3a75b0aad60ac2e988e9e4985ea1e6185c24c159f65449aee003d1180a"></a>

## Next pages — Examples / 4e0f0b8ce83e / 4

- [Resource](resources--secret_management_access--examples--group-001.md#canonical-33b225212396fd50f2735ac0ab47e65df6dc6bbc5d94bcb5210c714bbd6432fa)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)

<a id="canonical-33b225212396fd50f2735ac0ab47e65df6dc6bbc5d94bcb5210c714bbd6432fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ced58a73b65ea68f6b763e10f89af76cc3b4fb295efa1f39ada9bca328dec2"></a>

## Resource — Resource / bd1cc1987e1a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
- [Examples](resources--secret_management_access--examples--group-001.md#canonical-958b1534377812bddb2136dcb193677d4717b4eb90874cc67fbfb236ca549673)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_management_access/resource.tf`; digest `sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81`.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```

<a id="canonical-afda785695f30c754eb340ee90338e2a66fb2aaaa32522b26d31c5ed474e4fe3"></a>

## Next pages — Resource / bd1cc1987e1a / 3

- [Examples](resources--secret_management_access--examples--group-001.md#canonical-958b1534377812bddb2136dcb193677d4717b4eb90874cc67fbfb236ca549673)
- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8)
