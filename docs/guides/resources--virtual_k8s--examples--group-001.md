---
page_title: "xcsh_virtual_k8s examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s examples."
---

# xcsh_virtual_k8s examples

<a id="canonical-b9dfe44a63b9ab731f1bf8981d28b8e26f5cfe8a38945fcd155244c144832a31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cdc9b7910fe68227a6c66f47f29e535a06fd1a3137660f7e71da156e0d9dfbb"></a>

## Examples — Examples / 59c73098f4e3 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- Examples

<a id="canonical-27a8d2b465c7855dcdb4060ac20e9e0fcd41579f454125835985832adc45c805"></a>

## Complete configurations — Examples / 59c73098f4e3 / 3

- [Resource](resources--virtual_k8s--examples--group-001.md#canonical-e93c85527fc70edc2f3cea912c4af913a9fde3122314a2cd1f59c0534b095d5c): valid configuration.

<a id="canonical-135b4247ff1c8d8162d77d4cdf09fc8c3cc898dac007acc9e65129f7735c3da6"></a>

## Next pages — Examples / 59c73098f4e3 / 4

- [Resource](resources--virtual_k8s--examples--group-001.md#canonical-e93c85527fc70edc2f3cea912c4af913a9fde3122314a2cd1f59c0534b095d5c)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)

<a id="canonical-e93c85527fc70edc2f3cea912c4af913a9fde3122314a2cd1f59c0534b095d5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0e84878476f2e3dcaf4a0eabd528370df2bccf7177a88760830ec13b3c96328"></a>

## Resource — Resource / 1e7ac14bd0d3 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
- [Examples](resources--virtual_k8s--examples--group-001.md#canonical-b9dfe44a63b9ab731f1bf8981d28b8e26f5cfe8a38945fcd155244c144832a31)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_virtual_k8s/resource.tf`; digest `sha256:a5065caf12c5cd73a6083207375f30b6129bce86fd82163c8845c9e0ab44c400`.

```terraform
# VirtualK8S Resource Example
# Manages virtual_k8s will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualK8S configuration
resource "xcsh_virtual_k8s" "example" {
  name      = "example-virtual-k8s"
  namespace = "staging"
}
```

<a id="canonical-2f5208237bb73aab69df06de1e2e33098b784c5a134d9bfd0d14e4d4fa6cb8da"></a>

## Next pages — Resource / 1e7ac14bd0d3 / 3

- [Examples](resources--virtual_k8s--examples--group-001.md#canonical-b9dfe44a63b9ab731f1bf8981d28b8e26f5cfe8a38945fcd155244c144832a31)
- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-ad9f5a1cdc816cd36fa876353050dc4ae0112a25623374e5d93fc619c8087b90)
