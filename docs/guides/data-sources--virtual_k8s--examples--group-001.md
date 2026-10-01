---
page_title: "xcsh_virtual_k8s examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s examples."
---

# xcsh_virtual_k8s examples

<a id="canonical-c39a52aefa52ab272de791fa63cb96987fe1981f36e2b0473eae32b77dc43ad8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8c829b47061d6a1d2806da7fb63f62126582667a3c720e18a12154a612e4c52"></a>

## Examples — Examples / 6fafd9246bc1 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- Examples

<a id="canonical-0a288e69ac4b333931e10be6d40be899025a291e4c704b092b44a88b26422d6d"></a>

## Complete configurations — Examples / 6fafd9246bc1 / 3

- [Data source](data-sources--virtual_k8s--examples--group-001.md#canonical-68d9a6c8fc4e031cb67494c5c4df77b7526b96fc2cb52a18a805280d47c1261b): valid configuration.

<a id="canonical-e09b8f15e8b18c9cb5370acfa8275a1671fe798519bdd96f1b88382bed774750"></a>

## Next pages — Examples / 6fafd9246bc1 / 4

- [Data source](data-sources--virtual_k8s--examples--group-001.md#canonical-68d9a6c8fc4e031cb67494c5c4df77b7526b96fc2cb52a18a805280d47c1261b)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)

<a id="canonical-68d9a6c8fc4e031cb67494c5c4df77b7526b96fc2cb52a18a805280d47c1261b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93ae5051fc5680ec641d16c76fb44a389730b5684ba3b46f7d7689acfd1fe68a"></a>

## Data source — Data source / c55de738a2ec / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- [Examples](data-sources--virtual_k8s--examples--group-001.md#canonical-c39a52aefa52ab272de791fa63cb96987fe1981f36e2b0473eae32b77dc43ad8)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_k8s/data-source.tf`; digest `sha256:85c7894f9985a87a12ec00003bcdb34e2727cf22575e06597a586ec7fcc0019d`.

```terraform
# VirtualK8S Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualK8S by name
data "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}

output "virtual_k8s_id" {
  value = data.xcsh_virtual_k8s.example.id
}
```

<a id="canonical-f993ff00e7a7fae4407d4dd06c6b8d80d8d111aa6e297fde7305dd9260f9a570"></a>

## Next pages — Data source / c55de738a2ec / 3

- [Examples](data-sources--virtual_k8s--examples--group-001.md#canonical-c39a52aefa52ab272de791fa63cb96987fe1981f36e2b0473eae32b77dc43ad8)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
