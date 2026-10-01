---
page_title: "xcsh_ike2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 examples."
---

# xcsh_ike2 examples

<a id="canonical-5c44f3eaa6bf5c1d6489c418cbeccb9f274ea0a5737c1a60a95ec91dd69d099b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-235d28f81868aef53f32fb21b3b457f90bf1ff693f67f1e13bdba326b2418ed1"></a>

## Examples — Examples / 10998c2a1632 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- Examples

<a id="canonical-74685e3eadeeebcf6e2c87cebff2163580d7cf1b315a49721c4913a892022421"></a>

## Complete configurations — Examples / 10998c2a1632 / 3

- [Data source](data-sources--ike2--examples--group-001.md#canonical-0eaf273e5f41ffa8d5492276cfd9c627d9261fbcbf9d49f136d20f00c81f6827): valid configuration.

<a id="canonical-583752420b9b3903e18192900c6d01e5a11c51896c702c9e8ca759edc0200575"></a>

## Next pages — Examples / 10998c2a1632 / 4

- [Data source](data-sources--ike2--examples--group-001.md#canonical-0eaf273e5f41ffa8d5492276cfd9c627d9261fbcbf9d49f136d20f00c81f6827)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-0eaf273e5f41ffa8d5492276cfd9c627d9261fbcbf9d49f136d20f00c81f6827"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25d400b2b93aacf5f3c41d3391021d28c5168e55eba49204c4a3126096882a65"></a>

## Data source — Data source / 315942002406 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Examples](data-sources--ike2--examples--group-001.md#canonical-5c44f3eaa6bf5c1d6489c418cbeccb9f274ea0a5737c1a60a95ec91dd69d099b)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike2/data-source.tf`; digest `sha256:db87600cf054124d776876313565955caf46f37f46627a7f890a5aaf280c673b`.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

<a id="canonical-4191a60df00f07ecfebe5cb82232592a713cdfabaea762c54e6631d0d5d4c931"></a>

## Next pages — Data source / 315942002406 / 3

- [Examples](data-sources--ike2--examples--group-001.md#canonical-5c44f3eaa6bf5c1d6489c418cbeccb9f274ea0a5737c1a60a95ec91dd69d099b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
