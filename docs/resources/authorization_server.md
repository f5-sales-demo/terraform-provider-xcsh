---
page_title: "xcsh_authorization_server landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server landing."
---

# xcsh_authorization_server landing

<a id="canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0474398117221bbd85b707675b50758c494d1cc0ae2576a436fc832d8978ac9"></a>

## xcsh_authorization_server — xcsh_authorization_server / 920fc7e670f9 / 2

Breadcrumbs:

- xcsh_authorization_server

Manages authorization\_server creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-f316e6b40875dbdbdb43c5b8fbf126b4a226b94897919b69031ed9e6660ea9e5"></a>

## Prerequisites — xcsh_authorization_server / 920fc7e670f9 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-457c7124aab6d14d4305f5ac38cc5e7d7439779fce8571bb29ae06d3fe12b3e5"></a>

## Minimal configuration — xcsh_authorization_server / 920fc7e670f9 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```

<a id="canonical-133cf0c1983df20c2161f63fd719471b24853195cc5e9f9174d93c402074778b"></a>

## Root configuration — xcsh_authorization_server / 920fc7e670f9 / 5

Required root properties: `jwks_uri`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-26d5734a3167c1d37b2ebcc80252ffec985865687792506e8ee86ec46dc06d65"></a>

## Next pages — xcsh_authorization_server / 920fc7e670f9 / 6

- [Property reference](../guides/resources--authorization_server--reference--group-001.md#canonical-2554adc63fdb57eae149d6c2409f7a163791404f3563f7d1e34c7446f838913d)
- [Examples](../guides/resources--authorization_server--examples--group-001.md#canonical-1bb99cc587193065495177dec67c564e855be1be071a164a9bfb77b8a6432f7c)
- [Import](../guides/resources--authorization_server--lifecycle--group-001.md#canonical-f57139de248b0d90ceec4021b5a904e910d59733455de40fb819d20747e60d1f)
- [Timeouts](../guides/resources--authorization_server--lifecycle--group-001.md#canonical-431e51cd300c88ee0bf392d51f7a6c66e9491d2c2b9357f4ba22cda2cd6ac16d)
