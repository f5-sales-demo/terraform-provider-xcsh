---
page_title: "xcsh_addon_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service examples."
---

# xcsh_addon_service examples

<a id="canonical-a0399fbaae69f1178cd71f7f0650d72facb92ceb4977c1fad637c238fe1acb8f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-deea3a6091510a8a99915bd810ee9da54d9f74281fa37fc523731066e3383192"></a>

## Examples — Examples / 6cd46a57fefd / 2

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)
- Examples

<a id="canonical-7b267c10c93e57d80125710b89b568c5a215819059b67aa336f08bbd146b3b2e"></a>

## Complete configurations — Examples / 6cd46a57fefd / 3

- [Data source](data-sources--addon_service--examples--group-001.md#canonical-9d15640f9394c5cb364de8b8e7834ab1536ab70e31e6b2ad05a4beb4ffbcb1e7): valid configuration.

<a id="canonical-9e06a0add50092a39cefd243b9763847da6522f3afa1fd6aaf48b5954f7a7475"></a>

## Next pages — Examples / 6cd46a57fefd / 4

- [Data source](data-sources--addon_service--examples--group-001.md#canonical-9d15640f9394c5cb364de8b8e7834ab1536ab70e31e6b2ad05a4beb4ffbcb1e7)
- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)

<a id="canonical-9d15640f9394c5cb364de8b8e7834ab1536ab70e31e6b2ad05a4beb4ffbcb1e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a4f7d4c211be0d0e5a72174869c2897670f832091a43c335904c9cc2b44fc12"></a>

## Data source — Data source / 192106ee397a / 2

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)
- [Examples](data-sources--addon_service--examples--group-001.md#canonical-a0399fbaae69f1178cd71f7f0650d72facb92ceb4977c1fad637c238fe1acb8f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service/data-source.tf`; digest `sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98`.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

<a id="canonical-b2692522f2c6b82c1baa2cb9fcad513baba9d30e646ad335c2d5525573605db4"></a>

## Next pages — Data source / 192106ee397a / 3

- [Examples](data-sources--addon_service--examples--group-001.md#canonical-a0399fbaae69f1178cd71f7f0650d72facb92ceb4977c1fad637c238fe1acb8f)
- [xcsh_addon_service](../data-sources/addon_service.md#canonical-bc6915e40d9e500717e2d5968bc07fc54c844773be24d2f94cca4e8c5e83fe99)
