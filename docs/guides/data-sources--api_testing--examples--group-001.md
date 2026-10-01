---
page_title: "xcsh_api_testing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing examples."
---

# xcsh_api_testing examples

<a id="canonical-beb504e58ed4f5cdc2da1cb373b5d71c3498aaee616e4b912c35bcb206ab55b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1b0c023cd6dfe23ba89ab4b250fc3c3e189178febd24b3b4d4844d03db22f59"></a>

## Examples — Examples / 969ee37a29d5 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- Examples

<a id="canonical-51e20bbb6b65fa52df903d45e1108d03c407ff9d7f43801dc0575d89b8b78ba5"></a>

## Complete configurations — Examples / 969ee37a29d5 / 3

- [Data source](data-sources--api_testing--examples--group-001.md#canonical-fc065b5c4ca4ab2eabefa01b1fddc8411e2aacaaa96c7a222dc07cf819b69518): valid configuration.

<a id="canonical-3a256d30a528dc8e95d0ab04606796d4e067aa40a3600ad05f134e54a3bc3175"></a>

## Next pages — Examples / 969ee37a29d5 / 4

- [Data source](data-sources--api_testing--examples--group-001.md#canonical-fc065b5c4ca4ab2eabefa01b1fddc8411e2aacaaa96c7a222dc07cf819b69518)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)

<a id="canonical-fc065b5c4ca4ab2eabefa01b1fddc8411e2aacaaa96c7a222dc07cf819b69518"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a1f860bdad024ece6f04def72cd3c106a7fca8ccd99f29200c95bf119dcf709"></a>

## Data source — Data source / c7986ab00775 / 2

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
- [Examples](data-sources--api_testing--examples--group-001.md#canonical-beb504e58ed4f5cdc2da1cb373b5d71c3498aaee616e4b912c35bcb206ab55b3)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_testing/data-source.tf`; digest `sha256:509ae39ca1a6610bcc317fbc3a5f0883f2fcba290d264181d678caabac1f11fd`.

```terraform
# APITesting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APITesting by name
data "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}

output "api_testing_id" {
  value = data.xcsh_api_testing.example.id
}
```

<a id="canonical-32eb21c21139e30b166f2c231cc23b6ee0a6ee25d376e2fc62be75940b8bf1e6"></a>

## Next pages — Data source / c7986ab00775 / 3

- [Examples](data-sources--api_testing--examples--group-001.md#canonical-beb504e58ed4f5cdc2da1cb373b5d71c3498aaee616e4b912c35bcb206ab55b3)
- [xcsh_api_testing](../data-sources/api_testing.md#canonical-53728795593d298c8eed58f811191d2ef0f7236e7e77cb78260547474431a0ac)
