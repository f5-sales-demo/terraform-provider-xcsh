---
page_title: "xcsh_container_registry examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry examples."
---

# xcsh_container_registry examples

<a id="canonical-b02af979eb8bc9023d2ac42c427813e22c4f214bd2b8b65871fb12f69dd19488"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91359a54cbc0f886dcdb034a492d95f9030def318a14cd34b97a64c97f744857"></a>

## Examples — Examples / b9ed430b3812 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- Examples

<a id="canonical-55511e7bd749be78be1ac22eee61b808d26a9f1d31b61be3eaad3cbf0593b930"></a>

## Complete configurations — Examples / b9ed430b3812 / 3

- [Data source](data-sources--container_registry--examples--group-001.md#canonical-57137d725d4cf25a9c36ebe788d6fe7e00ce2047e35c61d85191f4b173288f19): valid configuration.

<a id="canonical-51de91312d3773bd5ee0f8921bd2df96eb894e743d4a03bf0f05932d75e58252"></a>

## Next pages — Examples / b9ed430b3812 / 4

- [Data source](data-sources--container_registry--examples--group-001.md#canonical-57137d725d4cf25a9c36ebe788d6fe7e00ce2047e35c61d85191f4b173288f19)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)

<a id="canonical-57137d725d4cf25a9c36ebe788d6fe7e00ce2047e35c61d85191f4b173288f19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9f5232a47ef32f65a3f3c8bdd9e5e126d95472eced40b48cd93a74ae8667918"></a>

## Data source — Data source / 78fcd6274553 / 2

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
- [Examples](data-sources--container_registry--examples--group-001.md#canonical-b02af979eb8bc9023d2ac42c427813e22c4f214bd2b8b65871fb12f69dd19488)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_container_registry/data-source.tf`; digest `sha256:235790f0ed1e99e7cff01491265f3f6071b78251c3ba60e8925b5d01fe4a9772`.

```terraform
# ContainerRegistry Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ContainerRegistry by name
data "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"
}

output "container_registry_id" {
  value = data.xcsh_container_registry.example.id
}
```

<a id="canonical-394e43138794b4bc1e1b89b5029b2827c632054a9121c10f3477b6c51a13ae8e"></a>

## Next pages — Data source / 78fcd6274553 / 3

- [Examples](data-sources--container_registry--examples--group-001.md#canonical-b02af979eb8bc9023d2ac42c427813e22c4f214bd2b8b65871fb12f69dd19488)
- [xcsh_container_registry](../data-sources/container_registry.md#canonical-17f6d44ebd4531a50d36aec8be1db646843481c3c17219cc6a239a366b16525f)
