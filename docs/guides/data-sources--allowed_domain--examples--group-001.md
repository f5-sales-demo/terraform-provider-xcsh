---
page_title: "xcsh_allowed_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain examples."
---

# xcsh_allowed_domain examples

<a id="canonical-72efddb7d276e7706910b3e4eef629cc5c74854545163ae33946456ab3c63799"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9a195dbfb81ead0e7c3246712bb2c50d7d3f4e445020ebba6a31cf8be9d7f30"></a>

## Examples — Examples / 5df5b25bd5fa / 2

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)
- Examples

<a id="canonical-8d94bad44c7c62789ab43f2493090e93ca206ffa7252f3e52558393d8bad0186"></a>

## Complete configurations — Examples / 5df5b25bd5fa / 3

- [Data source](data-sources--allowed_domain--examples--group-001.md#canonical-1000396a45322cb06be578c38595bbf99620c01b01c2b2479d813c95c85be904): valid configuration.

<a id="canonical-e33834bfbd85cca6c3d6628601e404fd55fcfcf6e142beb7fc042f3467b0f5e0"></a>

## Next pages — Examples / 5df5b25bd5fa / 4

- [Data source](data-sources--allowed_domain--examples--group-001.md#canonical-1000396a45322cb06be578c38595bbf99620c01b01c2b2479d813c95c85be904)
- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)

<a id="canonical-1000396a45322cb06be578c38595bbf99620c01b01c2b2479d813c95c85be904"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07d5e8cb91121cde8be371ccdc7fe9cd387ee79f1f0636790dd3c79b8df5104f"></a>

## Data source — Data source / b1fda7d73e97 / 2

Breadcrumbs:

- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)
- [Examples](data-sources--allowed_domain--examples--group-001.md#canonical-72efddb7d276e7706910b3e4eef629cc5c74854545163ae33946456ab3c63799)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_allowed_domain/data-source.tf`; digest `sha256:e22d7f6871a55bc1c9dd4658347afbf10c80e91173e9c9082eacc38fb5713006`.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```

<a id="canonical-eab3a49f367db37ac6009dbf101b9733387d5f8378e414db4f944aff5068122a"></a>

## Next pages — Data source / b1fda7d73e97 / 3

- [Examples](data-sources--allowed_domain--examples--group-001.md#canonical-72efddb7d276e7706910b3e4eef629cc5c74854545163ae33946456ab3c63799)
- [xcsh_allowed_domain](../data-sources/allowed_domain.md#canonical-cd1b8e8b9763ec41046ec5b23378de3e4eb43cba5cd5576863af43284e5075e8)
