---
page_title: "xcsh_access_active_session examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session examples."
---

# xcsh_access_active_session examples

<a id="canonical-1fae4c0e495ede3a8cc200a71a1345278a488be35029641220310fefd9671a9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2293656ab67eb2e413053f97a64f7b23829d591514222dd4f688a8264889740d"></a>

## Examples — Examples / 30be9e9a4e77 / 2

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
- Examples

<a id="canonical-d458382c55caf4f955f0cebf6e924044a3606491b904083a60e74fcc0316433c"></a>

## Complete configurations — Examples / 30be9e9a4e77 / 3

- [Data source](data-sources--access_active_session--examples--group-001.md#canonical-5a9216fd259763d66099385b1c17c27e28fc5d02e79ac1745de2a9a469e50181): valid configuration.

<a id="canonical-5e253b764ce05d30f9051980c10316aa2419e7211023b66b91b002fe5d6d8070"></a>

## Next pages — Examples / 30be9e9a4e77 / 4

- [Data source](data-sources--access_active_session--examples--group-001.md#canonical-5a9216fd259763d66099385b1c17c27e28fc5d02e79ac1745de2a9a469e50181)
- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)

<a id="canonical-5a9216fd259763d66099385b1c17c27e28fc5d02e79ac1745de2a9a469e50181"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-255839263d3541808ea62b3928bdbff97b0173eade255ccd3a1e71019a596062"></a>

## Data source — Data source / 309dbf54c3fe / 2

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
- [Examples](data-sources--access_active_session--examples--group-001.md#canonical-1fae4c0e495ede3a8cc200a71a1345278a488be35029641220310fefd9671a9e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_session/data-source.tf`; digest `sha256:d3411f6908febc7c8a9bf8b482b2e026429a95e16a75b366bc2a33e4f32326f7`.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```

<a id="canonical-87b96154ea7bed98e6fa2191329d12f8981978de29e57712932f4ce13a13d70b"></a>

## Next pages — Data source / 309dbf54c3fe / 3

- [Examples](data-sources--access_active_session--examples--group-001.md#canonical-1fae4c0e495ede3a8cc200a71a1345278a488be35029641220310fefd9671a9e)
- [xcsh_access_active_session](../data-sources/access_active_session.md#canonical-4033739913ba1c6085f12486a3986623b110d98373ef67466a8b65704eed74cb)
