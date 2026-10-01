---
page_title: "xcsh_cluster examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster examples."
---

# xcsh_cluster examples

<a id="canonical-2d44b777a62a6d2fa2ade46bd69fd423ae33771c7c995f33d89a240f913d68b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70ce131fa89c3f14e87bde6b5d9f5d7b4111e4649c07905a3288ce46be31d848"></a>

## Examples — Examples / 4f8d049d230a / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- Examples

<a id="canonical-2e58df67da6fb9771da1bcabe14f7c035e186fb94fee10322e16b7815647ab1f"></a>

## Complete configurations — Examples / 4f8d049d230a / 3

- [Data source](data-sources--cluster--examples--group-001.md#canonical-d79e1c9ec8b9f8c0668a6bc30fff0c345457b7a8418ba3eb18e69652a9275749): valid configuration.

<a id="canonical-0735601a6e8025c4c19ab7becc1248b5453274404efc3a50486edb298836e0fb"></a>

## Next pages — Examples / 4f8d049d230a / 4

- [Data source](data-sources--cluster--examples--group-001.md#canonical-d79e1c9ec8b9f8c0668a6bc30fff0c345457b7a8418ba3eb18e69652a9275749)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)

<a id="canonical-d79e1c9ec8b9f8c0668a6bc30fff0c345457b7a8418ba3eb18e69652a9275749"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31c7ad749747f2832c61435bd33ec25d714c07f63c97d96104173bd1d4115071"></a>

## Data source — Data source / 3d1bdd223491 / 2

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
- [Examples](data-sources--cluster--examples--group-001.md#canonical-2d44b777a62a6d2fa2ade46bd69fd423ae33771c7c995f33d89a240f913d68b1)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cluster/data-source.tf`; digest `sha256:657c300c6f1b904145d0c5ed8c458dbdc20cbe652caf3ef6b0a40bc6ebb5e4ee`.

```terraform
# Cluster Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cluster by name
data "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"
}

output "cluster_id" {
  value = data.xcsh_cluster.example.id
}
```

<a id="canonical-bf45b95003bb4733f0eacfc264017fa4c2dcc61e2032939128d996db38300af1"></a>

## Next pages — Data source / 3d1bdd223491 / 3

- [Examples](data-sources--cluster--examples--group-001.md#canonical-2d44b777a62a6d2fa2ade46bd69fd423ae33771c7c995f33d89a240f913d68b1)
- [xcsh_cluster](../data-sources/cluster.md#canonical-2e92d3096d97ff338e7244a89218e5e2a2ba4877a459c64fdf9727d96290d0c7)
