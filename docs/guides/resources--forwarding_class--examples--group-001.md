---
page_title: "xcsh_forwarding_class examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class examples."
---

# xcsh_forwarding_class examples

<a id="canonical-5955f4f6761160e682dae3b90ae4fc0c70c41e30a8dcd6aaaefdab6915e072ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8273cd1b9efc5734feae5afce2d97cf70c0d6e67cf5f8125b9c34fd0382f497d"></a>

## Examples — Examples / ed975540287a / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- Examples

<a id="canonical-47e62bc3df80e75565b2da77898c6bf4cca72e7579780e88697dc8f481f5e2bc"></a>

## Complete configurations — Examples / ed975540287a / 3

- [Resource](resources--forwarding_class--examples--group-001.md#canonical-cfcd74f41d1a80dbba587e6e122e193b1d7461bf080242b28b9ab55ba4553054): valid configuration.

<a id="canonical-7b0103d58077f76a3c8d37013977d713428e0eccf66615d2c9cc8a74ec533f50"></a>

## Next pages — Examples / ed975540287a / 4

- [Resource](resources--forwarding_class--examples--group-001.md#canonical-cfcd74f41d1a80dbba587e6e122e193b1d7461bf080242b28b9ab55ba4553054)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)

<a id="canonical-cfcd74f41d1a80dbba587e6e122e193b1d7461bf080242b28b9ab55ba4553054"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b6b10d5166cffce852ba2dbc730db754432d7e3936510c52baf85d8e3e8dcbc7"></a>

## Resource — Resource / 3710435ab0c2 / 2

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
- [Examples](resources--forwarding_class--examples--group-001.md#canonical-5955f4f6761160e682dae3b90ae4fc0c70c41e30a8dcd6aaaefdab6915e072ea)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forwarding_class/resource.tf`; digest `sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8`.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```

<a id="canonical-4ccb75006a130cb99c66c941bf0f330ec6a24a1cab36af1d6d166eaefc298375"></a>

## Next pages — Resource / 3710435ab0c2 / 3

- [Examples](resources--forwarding_class--examples--group-001.md#canonical-5955f4f6761160e682dae3b90ae4fc0c70c41e30a8dcd6aaaefdab6915e072ea)
- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-b532b65420c12831270a1c2064a45d87a3cc38ad0f9908b904eac37340a37726)
