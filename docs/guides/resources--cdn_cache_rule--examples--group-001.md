---
page_title: "xcsh_cdn_cache_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule examples."
---

# xcsh_cdn_cache_rule examples

<a id="canonical-26f48ea8b5f91d0dad499291c518db3babf530eb629d1bb63ef4b4191f7fdfa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-614171ce7b916245b184c93cd2818287bdf76593a8436fe24501c59f551c62c5"></a>

## Examples — Examples / fe55e881d186 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- Examples

<a id="canonical-bcba6fc9714fc1a69268c4fab77428ac6691cab6e4f04d5b48f3dbe49972e1b5"></a>

## Complete configurations — Examples / fe55e881d186 / 3

- [Resource](resources--cdn_cache_rule--examples--group-001.md#canonical-2626359a798835f3525da59361be30ea9b8cd6b20aaf2e50c1b6cf0175595b49): valid configuration.

<a id="canonical-1a04d57dcd555c1066825ef88f2b3b4a1a93dd6b47b910416121c02e1c060109"></a>

## Next pages — Examples / fe55e881d186 / 4

- [Resource](resources--cdn_cache_rule--examples--group-001.md#canonical-2626359a798835f3525da59361be30ea9b8cd6b20aaf2e50c1b6cf0175595b49)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)

<a id="canonical-2626359a798835f3525da59361be30ea9b8cd6b20aaf2e50c1b6cf0175595b49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b493292ce37a38465305371151de245ecfed4f761ed361bc74aa52d7d27cc001"></a>

## Resource — Resource / 7055a6c62ae0 / 2

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
- [Examples](resources--cdn_cache_rule--examples--group-001.md#canonical-26f48ea8b5f91d0dad499291c518db3babf530eb629d1bb63ef4b4191f7fdfa5)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_cache_rule/resource.tf`; digest `sha256:a602b3888b8d6105a3cf057267d0d93b04e1fb4a8ec89a85c41a5bb720a38a3e`.

```terraform
# CDNCacheRule Resource Example
# Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNCacheRule configuration
resource "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}
```

<a id="canonical-d9e8f18c4d2e07cbdaa3d4bd057d868eba73f4ebb0d0b1ad42188ee8ddb27e6f"></a>

## Next pages — Resource / 7055a6c62ae0 / 3

- [Examples](resources--cdn_cache_rule--examples--group-001.md#canonical-26f48ea8b5f91d0dad499291c518db3babf530eb629d1bb63ef4b4191f7fdfa5)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-e0e92aa98813cd7b91f796af1402650a131c06ef78bf9801126f371cdb70a081)
