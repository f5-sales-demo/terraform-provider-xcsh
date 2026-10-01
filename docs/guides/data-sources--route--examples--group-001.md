---
page_title: "xcsh_route examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route examples."
---

# xcsh_route examples

<a id="canonical-9391d20cfe41560e191bcd91f8e4b8422c2d0ad54f96e4235a550fc4f43c4447"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-733b9689b2b71663cd03b60dda15ed3a6f5cfa4b785412c940d18366bc006930"></a>

## Examples — Examples / 36d00b2ff976 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- Examples

<a id="canonical-591e9710aad3d081b6d00f790433973bb073808a0867bd5226aec87cb9d07234"></a>

## Complete configurations — Examples / 36d00b2ff976 / 3

- [Data source](data-sources--route--examples--group-001.md#canonical-950546bcf78198b84e6f06e6ccb47cc8b525a87f61097f57a0e96840a3a10907): valid configuration.

<a id="canonical-223ba035f1ab028345c1d06e04c1de80777f7bc37b243192c1cf06ffe424739f"></a>

## Next pages — Examples / 36d00b2ff976 / 4

- [Data source](data-sources--route--examples--group-001.md#canonical-950546bcf78198b84e6f06e6ccb47cc8b525a87f61097f57a0e96840a3a10907)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)

<a id="canonical-950546bcf78198b84e6f06e6ccb47cc8b525a87f61097f57a0e96840a3a10907"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cf12efb7d9d30a9d09178094a196dc19be6d888c3bb1ad9d94f22b9200e5f87"></a>

## Data source — Data source / 5cc60e3f503d / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
- [Examples](data-sources--route--examples--group-001.md#canonical-9391d20cfe41560e191bcd91f8e4b8422c2d0ad54f96e4235a550fc4f43c4447)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_route/data-source.tf`; digest `sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f`.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```

<a id="canonical-e61e02c2642a981dc565a33fa455382124a5937cb98d0a3e408f026359ee772d"></a>

## Next pages — Data source / 5cc60e3f503d / 3

- [Examples](data-sources--route--examples--group-001.md#canonical-9391d20cfe41560e191bcd91f8e4b8422c2d0ad54f96e4235a550fc4f43c4447)
- [xcsh_route](../data-sources/route.md#canonical-ea42e1ea4cf100a20de1cf54455610a98c198155f1e50696dd3051675e64bcc7)
