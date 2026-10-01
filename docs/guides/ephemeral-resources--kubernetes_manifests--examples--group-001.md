---
page_title: "xcsh_kubernetes_manifests examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_kubernetes_manifests examples."
---

# xcsh_kubernetes_manifests examples

<a id="canonical-facff2d02f16b29739081fecca5d71af469e84f62929ac9ed28ace112c3855fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-413537a865b7b64340027fabf18ad34ad1b015bf4d69253f4580361c9f48b183"></a>

## Examples — Examples / ce5d0f420462 / 2

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-c83216e14c73f7838248127c78a0137201c23f4a90e3676b65e4448f86ede548)
- Examples

<a id="canonical-2b341ae7b5fae40b048280c59c69e5a878f05a780a5d2604f2094daa27d22261"></a>

## Complete configurations — Examples / ce5d0f420462 / 3

- [Ephemeral](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-6a033ef70ea342b1259f824b22d2c0d55a1b4ca11f8405365b515b62f62d4f87): valid configuration.

<a id="canonical-c018360ab08f180623339791073495e29ae719d917c43e7c6f4307c632e66b37"></a>

## Next pages — Examples / ce5d0f420462 / 4

- [Ephemeral](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-6a033ef70ea342b1259f824b22d2c0d55a1b4ca11f8405365b515b62f62d4f87)
- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-c83216e14c73f7838248127c78a0137201c23f4a90e3676b65e4448f86ede548)

<a id="canonical-6a033ef70ea342b1259f824b22d2c0d55a1b4ca11f8405365b515b62f62d4f87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6b3c79f5dc865baacccaa8953a6130908065e8039c412a810669e0f51ab2f0d"></a>

## Ephemeral — Ephemeral / c5b548f4ed40 / 2

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-c83216e14c73f7838248127c78a0137201c23f4a90e3676b65e4448f86ede548)
- [Examples](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-facff2d02f16b29739081fecca5d71af469e84f62929ac9ed28ace112c3855fe)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf`; digest `sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e`.

```terraform
# KubernetesManifests EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_kubernetes_manifests" "example" {
  site = "example-value"
}
```

<a id="canonical-614b0b2d227396551a93ac668a19a767ee7cab4d60136d7684ea268eb0b9a6fd"></a>

## Next pages — Ephemeral / c5b548f4ed40 / 3

- [Examples](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-facff2d02f16b29739081fecca5d71af469e84f62929ac9ed28ace112c3855fe)
- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-c83216e14c73f7838248127c78a0137201c23f4a90e3676b65e4448f86ede548)
