---
page_title: "xcsh_artifact_registry_token examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_artifact_registry_token examples."
---

# xcsh_artifact_registry_token examples

<a id="canonical-05ded51d3e607a3e15c2af840b498c3b1b531de1ea66c354eea7a579d3523c57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0168abe57e7fc07f2429b83036e8e59b7797acb1f187ff942b27a25db1299b7a"></a>

## Examples — Examples / 0d23821f12af / 2

Breadcrumbs:

- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-7bc402ebc64070ff4db3dbcdbe34d30a0f0786ef3344bee1fd37affdbc4fe9b1)
- Examples

<a id="canonical-25015010eb15ca327c4ecc9ec023b63f0aa888f82c63b8ae87539357cce495ae"></a>

## Complete configurations — Examples / 0d23821f12af / 3

- [Ephemeral](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-3d010c39560ceeb9f90293543567eba18d37d260ffb27f048dc69a0a1b4a7751): valid configuration.

<a id="canonical-092eddc119722b194bb4dd90475df7e92fad5beefeb0efc1105c32249e0c3873"></a>

## Next pages — Examples / 0d23821f12af / 4

- [Ephemeral](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-3d010c39560ceeb9f90293543567eba18d37d260ffb27f048dc69a0a1b4a7751)
- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-7bc402ebc64070ff4db3dbcdbe34d30a0f0786ef3344bee1fd37affdbc4fe9b1)

<a id="canonical-3d010c39560ceeb9f90293543567eba18d37d260ffb27f048dc69a0a1b4a7751"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02ff508dff587a7dfcb21c11c53bc08235a08977b1163c2e4be5a1df39f259c3"></a>

## Ephemeral — Ephemeral / 47d6db7ef9b9 / 2

Breadcrumbs:

- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-7bc402ebc64070ff4db3dbcdbe34d30a0f0786ef3344bee1fd37affdbc4fe9b1)
- [Examples](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-05ded51d3e607a3e15c2af840b498c3b1b531de1ea66c354eea7a579d3523c57)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_artifact_registry_token/ephemeral.tf`; digest `sha256:f91e2c9b88b3ed39e41f1e34edb4d5addb5289872f1ac18a74ddd0ff3f1aec6f`.

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

<a id="canonical-8adbfa994449f88931c7c74451ec5236eabc5e84edf8d5916161e5b53a84c264"></a>

## Next pages — Ephemeral / 47d6db7ef9b9 / 3

- [Examples](ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-05ded51d3e607a3e15c2af840b498c3b1b531de1ea66c354eea7a579d3523c57)
- [xcsh_artifact_registry_token](../ephemeral-resources/artifact_registry_token.md#canonical-7bc402ebc64070ff4db3dbcdbe34d30a0f0786ef3344bee1fd37affdbc4fe9b1)
