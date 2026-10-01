---
page_title: "xcsh_policy_based_routing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policy_based_routing examples."
---

# xcsh_policy_based_routing examples

<a id="canonical-2b37fbed9b408e1b36ae3feaeae2737ab01aa3c0205049bbd43a7fe1d7bd105d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29fe25d1069641113f03c1ef277a9428dac06913b7b9c317eb2a869990906ced"></a>

## Examples — Examples / bc29188ab947 / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- Examples

<a id="canonical-2a5bde9fa268ed511dc4c2887380cdffbbd41e963be9c8b051e338230fe8aa57"></a>

## Complete configurations — Examples / bc29188ab947 / 3

- [Resource](resources--policy_based_routing--examples--group-001.md#canonical-f4408521b7c4cd779b468c31f5c46b8e3264e447c3c12f626857400aa0737f83): valid configuration.

<a id="canonical-c0c79d1b4fe7427512e9feee2f23f68676355d8e87fc170daa0b4617f82362ca"></a>

## Next pages — Examples / bc29188ab947 / 4

- [Resource](resources--policy_based_routing--examples--group-001.md#canonical-f4408521b7c4cd779b468c31f5c46b8e3264e447c3c12f626857400aa0737f83)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)

<a id="canonical-f4408521b7c4cd779b468c31f5c46b8e3264e447c3c12f626857400aa0737f83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99b032d273a6ab7049549db34ceb0c96c277b65f83cdc2a5d45394e368ef76ff"></a>

## Resource — Resource / 7fb1a0fca3eb / 2

Breadcrumbs:

- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
- [Examples](resources--policy_based_routing--examples--group-001.md#canonical-2b37fbed9b408e1b36ae3feaeae2737ab01aa3c0205049bbd43a7fe1d7bd105d)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_policy_based_routing/resource.tf`; digest `sha256:a6f7aae564241cf8a8dc1fa822f0c450be182f4323bb106bc39515f008c22968`.

```terraform
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

<a id="canonical-4f3e3b0d223e057145974f05dc8eb333b6a9975baf1c3daa41e7feaf84e4b3e4"></a>

## Next pages — Resource / 7fb1a0fca3eb / 3

- [Examples](resources--policy_based_routing--examples--group-001.md#canonical-2b37fbed9b408e1b36ae3feaeae2737ab01aa3c0205049bbd43a7fe1d7bd105d)
- [xcsh_policy_based_routing](../resources/policy_based_routing.md#canonical-6f1c7cea107838781ae5ab4cd5ec1fa8f4de7488f244d42cc676348e55c31152)
