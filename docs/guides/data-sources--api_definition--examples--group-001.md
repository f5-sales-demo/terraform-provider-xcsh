---
page_title: "xcsh_api_definition examples"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition examples."
---

# xcsh_api_definition examples

<a id="canonical-4f65221eb8214fce92c087c6dba0a758d6c4edc5500cdffac6f34a83f2e26a7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e69e33d2ea70619f16a7bc2e90a4f8d45bc94efef5ad0881109145b817cd347"></a>

## Examples — Examples / f6a61fe009dd / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- Examples

<a id="canonical-844faa3b67672a81357cb6e222bf9033a8eec2d174432ba5931505248ea21e63"></a>

## Complete configurations — Examples / f6a61fe009dd / 3

- [Data source](data-sources--api_definition--examples--group-001.md#canonical-4550306a027c59d59928741b7ee3a9c6c7e9a6a4082c862b2c20ec2766427db7): valid configuration.

<a id="canonical-934db33ea808f678a7fe5bf7cace8d24b16822ce5f991c9f89cb31a305886365"></a>

## Next pages — Examples / f6a61fe009dd / 4

- [Data source](data-sources--api_definition--examples--group-001.md#canonical-4550306a027c59d59928741b7ee3a9c6c7e9a6a4082c862b2c20ec2766427db7)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)

<a id="canonical-4550306a027c59d59928741b7ee3a9c6c7e9a6a4082c862b2c20ec2766427db7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d76ec110ee09ef3083606e3a0b3f7d9f1a9578e637cef5a36bfc025927efc78a"></a>

## Data source — Data source / c895d0c9769b / 2

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
- [Examples](data-sources--api_definition--examples--group-001.md#canonical-4f65221eb8214fce92c087c6dba0a758d6c4edc5500cdffac6f34a83f2e26a7d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_definition/data-source.tf`; digest `sha256:00117931ab83e7521384d3c585bd071aa24022def862d24108a8273097450108`.

```terraform
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```

<a id="canonical-7d8e6a745b78b848ca9b4867033e2cfc848bf54a9091a38bbfeb05c05832fe91"></a>

## Next pages — Data source / c895d0c9769b / 3

- [Examples](data-sources--api_definition--examples--group-001.md#canonical-4f65221eb8214fce92c087c6dba0a758d6c4edc5500cdffac6f34a83f2e26a7d)
- [xcsh_api_definition](../data-sources/api_definition.md#canonical-1df393999e0c5facf84460c20dfc66c998ea2afc9c81bd04c6d3cf585cad53ab)
