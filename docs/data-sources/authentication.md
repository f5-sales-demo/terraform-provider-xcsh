---
page_title: "xcsh_authentication landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication landing."
---

# xcsh_authentication landing

<a id="canonical-fc74a07ec6743be123535a3b99a60acfe7e3d06088347cb014d5bfa9986a7413"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67010cfbede6faa4ee3638213ee538251843a5afc0d82075671fe5a2ec852d2b"></a>

## xcsh_authentication — xcsh_authentication / 0ba4101b3b27 / 2

Breadcrumbs:

- xcsh_authentication

Manages a Authentication resource in F5 Distributed Cloud.

<a id="canonical-69f404736033d75e1aa18dfaed9a9419307ce0dceff8a4026fc720b753c63c34"></a>

## Prerequisites — xcsh_authentication / 0ba4101b3b27 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0b9ffcff70cbdcef545c21afd6d99a222d72668ff1b1cabc1062ebf42871a346"></a>

## Minimal configuration — xcsh_authentication / 0ba4101b3b27 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

<a id="canonical-f4d04a982c756d8912d8befd2ae148aa1771c3f1d229712aef4841cac9b9da9f"></a>

## Root configuration — xcsh_authentication / 0ba4101b3b27 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-79a1a6caf9768564a6145e7fcb54e80db85bd5923fd5e7248d21615460462376"></a>

## Next pages — xcsh_authentication / 0ba4101b3b27 / 6

- [Property reference](../guides/data-sources--authentication--reference--group-001.md#canonical-82eea9682c2c11291a2ba818100f690ec2a8060ec997c2ac4e6950b355cc0d64)
- [Examples](../guides/data-sources--authentication--examples--group-001.md#canonical-f543f75ce3357d8e276af5544257737200167f9301ffcd387f66063716548ac0)
