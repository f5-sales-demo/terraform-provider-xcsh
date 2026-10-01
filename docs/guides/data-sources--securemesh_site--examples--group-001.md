---
page_title: "xcsh_securemesh_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site examples."
---

# xcsh_securemesh_site examples

<a id="canonical-e9fdcf3f99896d5d2d8fc3260c18902073b98d1b6b84881b2fcdb33ae6a33ade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b442d08a548f10e23be788c8a9084a224655035d3962df24be218e21a258ee24"></a>

## Examples — Examples / c6a49f6bcf74 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- Examples

<a id="canonical-b9e1161ae6f998227c53d1fb08518c246bef89e5885f8a64a2da43e10e1e92ef"></a>

## Complete configurations — Examples / c6a49f6bcf74 / 3

- [Data source](data-sources--securemesh_site--examples--group-001.md#canonical-efdf017d807db965d9f2cb7189e7e15a4f3b3c1bfea4839684c8db5ee25c2544): valid configuration.

<a id="canonical-f883942fff713dc8affd945c17280b94a7fc2c965096e68b7f581caf82511c73"></a>

## Next pages — Examples / c6a49f6bcf74 / 4

- [Data source](data-sources--securemesh_site--examples--group-001.md#canonical-efdf017d807db965d9f2cb7189e7e15a4f3b3c1bfea4839684c8db5ee25c2544)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)

<a id="canonical-efdf017d807db965d9f2cb7189e7e15a4f3b3c1bfea4839684c8db5ee25c2544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-813e0c3e4cc8490673dc43e0c379849f6d874ebd3398c77d9b7e5fe3f9bf0d78"></a>

## Data source — Data source / 95e0226b0ae6 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
- [Examples](data-sources--securemesh_site--examples--group-001.md#canonical-e9fdcf3f99896d5d2d8fc3260c18902073b98d1b6b84881b2fcdb33ae6a33ade)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site/data-source.tf`; digest `sha256:476e0c14fbf185d65694b51a11030a2462b91f28a9d067a59c797b6328e138a4`.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

<a id="canonical-73d7924c715f64258642178df7ac46c83c17b617dc30f4cec6b3fd4ab190363d"></a>

## Next pages — Data source / 95e0226b0ae6 / 3

- [Examples](data-sources--securemesh_site--examples--group-001.md#canonical-e9fdcf3f99896d5d2d8fc3260c18902073b98d1b6b84881b2fcdb33ae6a33ade)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-5f91b1b3404e561dd8f463d8833156b2b57b09182e422126fe076a2265156cec)
