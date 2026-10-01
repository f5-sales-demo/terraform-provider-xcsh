---
page_title: "xcsh_cloud_connect examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect examples."
---

# xcsh_cloud_connect examples

<a id="canonical-1b4515988a7c81c8322c907aa5b22dee9666424be27bf8fc5d830756c34a7f99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34e602a5a1af9f3b6911a12cfdd23308762953c57d9ed6d4f195daed17ac4d90"></a>

## Examples — Examples / f6cb5010d926 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- Examples

<a id="canonical-a657c297b78b120f4d1623fee247df7aac39f5aea3aa9440c0cc500441940e55"></a>

## Complete configurations — Examples / f6cb5010d926 / 3

- [Resource](resources--cloud_connect--examples--group-001.md#canonical-b3a3d9e092f6122d43b3e36427d227d35802b1e34d7ecbad7574b3071e613cfd): valid configuration.

<a id="canonical-86e477fe992eff685f0a85814d5e4c227a2fafb7e66c131179fadbdaeaf25fbc"></a>

## Next pages — Examples / f6cb5010d926 / 4

- [Resource](resources--cloud_connect--examples--group-001.md#canonical-b3a3d9e092f6122d43b3e36427d227d35802b1e34d7ecbad7574b3071e613cfd)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)

<a id="canonical-b3a3d9e092f6122d43b3e36427d227d35802b1e34d7ecbad7574b3071e613cfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac30ee2b28f62e43944f8d6cdf8d2ee150800dfd4375a501d55b2cd0c4ceaff6"></a>

## Resource — Resource / 4b6b254f2af8 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
- [Examples](resources--cloud_connect--examples--group-001.md#canonical-1b4515988a7c81c8322c907aa5b22dee9666424be27bf8fc5d830756c34a7f99)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_connect/resource.tf`; digest `sha256:1c68a2114f267127cdbf735730210ab219b41bb6ec8f0bf42263953d188b9f46`.

```terraform
# CloudConnect Resource Example
# Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudConnect configuration
resource "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}
```

<a id="canonical-e42e25fc592468af147230deab07a241d9574297768f503c91c2dc028bb126c5"></a>

## Next pages — Resource / 4b6b254f2af8 / 3

- [Examples](resources--cloud_connect--examples--group-001.md#canonical-1b4515988a7c81c8322c907aa5b22dee9666424be27bf8fc5d830756c34a7f99)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f)
