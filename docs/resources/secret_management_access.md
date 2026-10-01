---
page_title: "xcsh_secret_management_access landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access landing."
---

# xcsh_secret_management_access landing

<a id="canonical-6be9f80a3a7a0206a0d3b04c5f1b83e522746b76ffdaf726eacb4ae910005ac8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-931ff2d141fffb6df62f36a04df013ef38658fb2bf98f3d60c99ae6c13bcd0cd"></a>

## xcsh_secret_management_access — xcsh_secret_management_access / 5784fb68f524 / 2

Breadcrumbs:

- xcsh_secret_management_access

Manages secret\_management\_access creates a new object in storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-e2ca7ca89e5e8fd14135d0530a96b154b89738b17dc1dacd1813277d2bfa165a"></a>

## Prerequisites — xcsh_secret_management_access / 5784fb68f524 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-799d9be9641a62c426898005e36571ce4d3f67eedc2c3ee4a4713ac918dd1cbe"></a>

## Minimal configuration — xcsh_secret_management_access / 5784fb68f524 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-772f36c22dd61a649f4624f95b9137a2269e33f43ce2950df6e4caa4a2b8cafa"></a>

## Root configuration — xcsh_secret_management_access / 5784fb68f524 / 5

Required root properties: `name`, `namespace`, `provider_name`. Full root flags and choices appear in the property reference.

<a id="canonical-ca0f02b9b929fc5b6c32e82d5359f13b3da308621121b1da103a421e3df3f888"></a>

## Next pages — xcsh_secret_management_access / 5784fb68f524 / 6

- [Property reference](../guides/resources--secret_management_access--reference--group-001.md#canonical-0e23e078c4f840b4addb6fd267dbf9e580ec8235242861dc0cc540ed30d8d09b)
- [Examples](../guides/resources--secret_management_access--examples--group-001.md#canonical-958b1534377812bddb2136dcb193677d4717b4eb90874cc67fbfb236ca549673)
- [Import](../guides/resources--secret_management_access--lifecycle--group-001.md#canonical-752fa880cd8b5458063a3d87f20a3eba3db97f52747797d68c805a93d1750fec)
- [Timeouts](../guides/resources--secret_management_access--lifecycle--group-001.md#canonical-c8bb39ba32845092f3e9acc7dc0664a707bcb217c29d0e5238061bd3e7bd54f8)
