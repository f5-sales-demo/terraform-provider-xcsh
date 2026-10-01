---
page_title: "xcsh_ike2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 examples."
---

# xcsh_ike2 examples

<a id="canonical-1586c8fb99e7237f1fab2a8ef86fd04ecbd9dcd11dbc823cef6bfe24275e7313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6091b65979ffc7a9c0138d88e9836c028c964c3ea0194e6aa532021fc7df74fd"></a>

## Examples — Examples / 24dce7d11d8d / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- Examples

<a id="canonical-f5be47b8503d8968e3729918d15552cd2fcc3c9426d7c566f755ecc6188d2b74"></a>

## Complete configurations — Examples / 24dce7d11d8d / 3

- [Resource](resources--ike2--examples--group-001.md#canonical-e52f4c4d3c894d2e5a9a270df50fbd1f78f9ac8b275f5c7976bf044a5aa7694a): valid configuration.

<a id="canonical-5c7c1a460365f469d5d926c2066eafad221b72b286e4b442254cae4ca8511f15"></a>

## Next pages — Examples / 24dce7d11d8d / 4

- [Resource](resources--ike2--examples--group-001.md#canonical-e52f4c4d3c894d2e5a9a270df50fbd1f78f9ac8b275f5c7976bf044a5aa7694a)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-e52f4c4d3c894d2e5a9a270df50fbd1f78f9ac8b275f5c7976bf044a5aa7694a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7724d299f0a88fea395d463cea318297fa4558e020fcc0052e6182b32119a50b"></a>

## Resource — Resource / b582deb92cd1 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Examples](resources--ike2--examples--group-001.md#canonical-1586c8fb99e7237f1fab2a8ef86fd04ecbd9dcd11dbc823cef6bfe24275e7313)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike2/resource.tf`; digest `sha256:9af0540e9ef3cddbd2cf33f981bb394c77fd860b251d272dd666ef02cd5bc017`.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

<a id="canonical-375cd231d33147bd85033c4d725e1ee0d0dc794dc1441baea9fd70e676abb91a"></a>

## Next pages — Resource / b582deb92cd1 / 3

- [Examples](resources--ike2--examples--group-001.md#canonical-1586c8fb99e7237f1fab2a8ef86fd04ecbd9dcd11dbc823cef6bfe24275e7313)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
