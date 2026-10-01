---
page_title: "xcsh_site_mesh_group examples"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_site_mesh_group examples."
---

# xcsh_site_mesh_group examples

<a id="canonical-fb5094d4a1ce4aa8a764458ae4ce2b2b0349dc4af1649ffdfb59f055cb68b77a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37bfe0208467e53b3fe6dcc451c1b07a106d71901ffcf5209ea8df7156a5826a"></a>

## Examples — Examples / fda2b2f98468 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- Examples

<a id="canonical-cf515e071b29f1c99087653631994da0fb49600ab2753328d615cd175469cf87"></a>

## Complete configurations — Examples / fda2b2f98468 / 3

- [Resource](resources--site_mesh_group--examples--group-001.md#canonical-87f7de15d04a870e59c74fc17ac3c79bbbdd6a8c1b389d9b1470c58cfa537d73): valid configuration.

<a id="canonical-8d0b3780f67a9490f4f362441f8c8246908163f12d7d684af50a60d446d2219a"></a>

## Next pages — Examples / fda2b2f98468 / 4

- [Resource](resources--site_mesh_group--examples--group-001.md#canonical-87f7de15d04a870e59c74fc17ac3c79bbbdd6a8c1b389d9b1470c58cfa537d73)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)

<a id="canonical-87f7de15d04a870e59c74fc17ac3c79bbbdd6a8c1b389d9b1470c58cfa537d73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9b94a0e6704be061485bd8c3887e9d9877043a050311e0e17fa42b87997731c"></a>

## Resource — Resource / 2ef6a4264c82 / 2

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
- [Examples](resources--site_mesh_group--examples--group-001.md#canonical-fb5094d4a1ce4aa8a764458ae4ce2b2b0349dc4af1649ffdfb59f055cb68b77a)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_site_mesh_group/resource.tf`; digest `sha256:9fd9c5f06a1b22b1d6bf2d57189b47e2ae5e2ce09701525907fff4345875b45e`.

```terraform
# SiteMeshGroup Resource Example
# Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SiteMeshGroup configuration
resource "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}
```

<a id="canonical-c74c5507fd7ed0c4c64a53855509d4a0cc62ff831c5169a86e061516e9ab5951"></a>

## Next pages — Resource / 2ef6a4264c82 / 3

- [Examples](resources--site_mesh_group--examples--group-001.md#canonical-fb5094d4a1ce4aa8a764458ae4ce2b2b0349dc4af1649ffdfb59f055cb68b77a)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md#canonical-5445c68d577813aa505c1832591c89ba83436e7eb3c8c9b83203b87260e168b7)
