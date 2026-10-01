---
page_title: "xcsh_segment examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment examples."
---

# xcsh_segment examples

<a id="canonical-0418839d33d311e131d6c8096eb96c74b36716c6c6f8f95cbcf0cdbf37919cff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4295f0bf024e99049c370a19774443c0f26082f8ac1ff0bb169cd32198a0fb"></a>

## Examples — Examples / 407903a48bab / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- Examples

<a id="canonical-7798ee08941d01ff41250a549b3ee99c00a53aed3bcccf6871b2a93c56a29ec3"></a>

## Complete configurations — Examples / 407903a48bab / 3

- [Resource](resources--segment--examples--group-001.md#canonical-2d6ec087ee1a3d55619d5415cffcf7b031e21383bafe36048a870aa8e3edc250): valid configuration.

<a id="canonical-a0b7b8b1e564ca1576c21811000c7b6830e06999b93cf8a5d7cd7fb4665c05eb"></a>

## Next pages — Examples / 407903a48bab / 4

- [Resource](resources--segment--examples--group-001.md#canonical-2d6ec087ee1a3d55619d5415cffcf7b031e21383bafe36048a870aa8e3edc250)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)

<a id="canonical-2d6ec087ee1a3d55619d5415cffcf7b031e21383bafe36048a870aa8e3edc250"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16fbb856d396abe41f3fe8d8ec91165d4abea1ff0af91fa05fab2a578d0c65ea"></a>

## Resource — Resource / 8fa50e025f03 / 2

Breadcrumbs:

- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
- [Examples](resources--segment--examples--group-001.md#canonical-0418839d33d311e131d6c8096eb96c74b36716c6c6f8f95cbcf0cdbf37919cff)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_segment/resource.tf`; digest `sha256:d784ed069273afd06c3680c3f92fabd8acd51be0aa4d921e6bfa8799ea0d6232`.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

<a id="canonical-8ec058651187cd900bf347981a140b4109371795323909ed6d07beeca2ca72cf"></a>

## Next pages — Resource / 8fa50e025f03 / 3

- [Examples](resources--segment--examples--group-001.md#canonical-0418839d33d311e131d6c8096eb96c74b36716c6c6f8f95cbcf0cdbf37919cff)
- [xcsh_segment](../resources/segment.md#canonical-115b348cc2da26907c0d3a2200fee1c45353481468d739f4370fd6bf8ca9a410)
