---
page_title: "xcsh_nfv_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service examples."
---

# xcsh_nfv_service examples

<a id="canonical-17846793b391435f50864456d6d6124f1ec9432c80f12fa51c56820307807d96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1464324bbf00e5c1003f50416a7c56b65630c436cb278365dc0d533b0fea25e1"></a>

## Examples — Examples / 43e13801a83b / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- Examples

<a id="canonical-63ed103011f28a076647d76324212c3b16a607d428f166d7bc6a11cbc75285c2"></a>

## Complete configurations — Examples / 43e13801a83b / 3

- [Data source](data-sources--nfv_service--examples--group-001.md#canonical-bbf2e440b5f2765c2e18bf7932d1c6a87abe61331f34f4ebc2b422fd4acf66ca): valid configuration.

<a id="canonical-2ef84703ebabdf3fb4d25e46b5e650174224b9a0823195ee257bca4bd7eaea0f"></a>

## Next pages — Examples / 43e13801a83b / 4

- [Data source](data-sources--nfv_service--examples--group-001.md#canonical-bbf2e440b5f2765c2e18bf7932d1c6a87abe61331f34f4ebc2b422fd4acf66ca)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)

<a id="canonical-bbf2e440b5f2765c2e18bf7932d1c6a87abe61331f34f4ebc2b422fd4acf66ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ecee6d51f67e0d5a4962d3863710ed4d0feba4fdb7eb1222897bd74b7fc0acde"></a>

## Data source — Data source / 11869a2a0de1 / 2

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
- [Examples](data-sources--nfv_service--examples--group-001.md#canonical-17846793b391435f50864456d6d6124f1ec9432c80f12fa51c56820307807d96)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nfv_service/data-source.tf`; digest `sha256:e1bcf01b10f5271939ad8fda84b9d63610e3af5498aba69a8e2b34b2c60983a2`.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```

<a id="canonical-4e5997c419255e0da2cd3dca42bc977a8e12017717e8da1aafce8ed35bb0817c"></a>

## Next pages — Data source / 11869a2a0de1 / 3

- [Examples](data-sources--nfv_service--examples--group-001.md#canonical-17846793b391435f50864456d6d6124f1ec9432c80f12fa51c56820307807d96)
- [xcsh_nfv_service](../data-sources/nfv_service.md#canonical-a6e13d1ea9c2aa4e6f35b818ca53ad6df10cbcafdfb9884a3a265a277ed400f0)
