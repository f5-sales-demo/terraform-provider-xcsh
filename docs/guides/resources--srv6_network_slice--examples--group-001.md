---
page_title: "xcsh_srv6_network_slice examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice examples."
---

# xcsh_srv6_network_slice examples

<a id="canonical-228c5eca67933ce19c2353ffab9baa21c19c683b96c643cd0f23c62c39acc9e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0494adfb577b5bb6e839791dbbe298507e6011ab9ae7d82df8efc997822e7e5"></a>

## Examples — Examples / d0b88c89257f / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
- Examples

<a id="canonical-716b8dbfc73d93461ae3c818744c94be1885a2c672eab73a3745a325bd4b47f8"></a>

## Complete configurations — Examples / d0b88c89257f / 3

- [Resource](resources--srv6_network_slice--examples--group-001.md#canonical-7d558ca171678561f84d9e12155a546c041c781360f09b04c9421e1d94cb52ee): valid configuration.

<a id="canonical-80f4ef4d4e561161f6fb9c36e04246049c8d21045c89dfff1520e20a49e56e48"></a>

## Next pages — Examples / d0b88c89257f / 4

- [Resource](resources--srv6_network_slice--examples--group-001.md#canonical-7d558ca171678561f84d9e12155a546c041c781360f09b04c9421e1d94cb52ee)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)

<a id="canonical-7d558ca171678561f84d9e12155a546c041c781360f09b04c9421e1d94cb52ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b9fd4b395079cc64dd3df801a03a38354a7fe23d95ef40a5ff91e936197268a"></a>

## Resource — Resource / e0b5ba05fdeb / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
- [Examples](resources--srv6_network_slice--examples--group-001.md#canonical-228c5eca67933ce19c2353ffab9baa21c19c683b96c643cd0f23c62c39acc9e1)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_srv6_network_slice/resource.tf`; digest `sha256:28b76f9b2c2ac199134ab8ed152b0dff8fdb07e8a894ab2aa2291993080df8d0`.

```terraform
# Srv6NetworkSlice Resource Example
# Manages srv6_network_slice creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Srv6NetworkSlice configuration
resource "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"

  sid_prefixes = ["example-value"]
}
```

<a id="canonical-1aa7023651d84a4b7c3f7c6b1a389a50853e11eb77869ae2652361684d54c020"></a>

## Next pages — Resource / e0b5ba05fdeb / 3

- [Examples](resources--srv6_network_slice--examples--group-001.md#canonical-228c5eca67933ce19c2353ffab9baa21c19c683b96c643cd0f23c62c39acc9e1)
- [xcsh_srv6_network_slice](../resources/srv6_network_slice.md#canonical-431435cc9f00e58094c1dd963b8aa9cec4251ea8dcbf3313b046ca1921402d3e)
