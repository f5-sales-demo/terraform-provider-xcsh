---
page_title: "xcsh_external_connector examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector examples."
---

# xcsh_external_connector examples

<a id="canonical-5b91edc54625be81d0068af801788d9ede1571781cfd3411cf395f1b0ca5ec31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f7e14ca64eaa85a1d8e0877877611add5729608b97749bffdb81c5a1420bd0"></a>

## Examples — Examples / 8e006acc0247 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- Examples

<a id="canonical-ff01317a423e72fd1d4dcefe5d8c35f6f5fa8fd81b040d236cdd32c47db4031d"></a>

## Complete configurations — Examples / 8e006acc0247 / 3

- [Data source](data-sources--external_connector--examples--group-001.md#canonical-2740ed018b1e265b9178908061da7e1d2341224c207dfe17129084b057d9a593): valid configuration.

<a id="canonical-ee8f2eea8439e95de58bb6c011cf4c665deac89ecc8411f2b77a04f91a44512c"></a>

## Next pages — Examples / 8e006acc0247 / 4

- [Data source](data-sources--external_connector--examples--group-001.md#canonical-2740ed018b1e265b9178908061da7e1d2341224c207dfe17129084b057d9a593)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)

<a id="canonical-2740ed018b1e265b9178908061da7e1d2341224c207dfe17129084b057d9a593"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa4b375ad7c182ef9b152da612ee1f8f1c14246ceaadb96b7a083260b6495731"></a>

## Data source — Data source / ccfe034d1982 / 2

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
- [Examples](data-sources--external_connector--examples--group-001.md#canonical-5b91edc54625be81d0068af801788d9ede1571781cfd3411cf395f1b0ca5ec31)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_external_connector/data-source.tf`; digest `sha256:8660d99e20bf0369dc735e047075e9bae757c75b8856770f0f1f3584ea3dc187`.

```terraform
# ExternalConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ExternalConnector by name
data "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}

output "external_connector_id" {
  value = data.xcsh_external_connector.example.id
}
```

<a id="canonical-ace8cbdc85a0e9742d3ba128414bad94379a98548510a76b8b3414c770994e9c"></a>

## Next pages — Data source / ccfe034d1982 / 3

- [Examples](data-sources--external_connector--examples--group-001.md#canonical-5b91edc54625be81d0068af801788d9ede1571781cfd3411cf395f1b0ca5ec31)
- [xcsh_external_connector](../data-sources/external_connector.md#canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16)
