---
page_title: "xcsh_artifact_registry_token landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_artifact_registry_token landing."
---

# xcsh_artifact_registry_token landing

<a id="canonical-7bc402ebc64070ff4db3dbcdbe34d30a0f0786ef3344bee1fd37affdbc4fe9b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ed89e52be387c9c96a2ac92b2acbe98be5bd46b38ad1cd8effa131ba317f30f"></a>

## xcsh_artifact_registry_token — xcsh_artifact_registry_token / 300e3ecee56f / 2

Breadcrumbs:

- xcsh_artifact_registry_token

Authentication credential for access control.

<a id="canonical-41f666af8b9d48c12634913a526f520f0c02c8bc685917b0b3fd031e637b1457"></a>

## Prerequisites — xcsh_artifact_registry_token / 300e3ecee56f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-fdddf36efdebe4531c547a27a74b5f255500b2e277347bf1c2b85d6f175fc3aa"></a>

## Minimal configuration — xcsh_artifact_registry_token / 300e3ecee56f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ArtifactRegistryToken EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_artifact_registry_token" "example" {
  namespace = "example-value"
}
```

<a id="canonical-05f26d62c8936c7b04337ff3d5ab8b967f80dbeffba9022edea7e53f9eb0bd2c"></a>

## Root configuration — xcsh_artifact_registry_token / 300e3ecee56f / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0dc736cf6a002775016e246cdf27037913693ac21475f4c20fe32522c032e0a3"></a>

## Next pages — xcsh_artifact_registry_token / 300e3ecee56f / 6

- [Property reference](../guides/ephemeral-resources--artifact_registry_token--reference--group-001.md#canonical-a3dcfe8c172a2c5524817e37d40511bed73fd5a47e634a5bd12e0537f627528b)
- [Examples](../guides/ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-05ded51d3e607a3e15c2af840b498c3b1b531de1ea66c354eea7a579d3523c57)
- [Lifecycle](../guides/ephemeral-resources--artifact_registry_token--lifecycle--group-001.md#canonical-14f69f6067c03deafb334d9a4ceae9718e5ad0b7cc953075f928fd89e9944b00)
