---
page_title: "xcsh_external_connector landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector landing."
---

# xcsh_external_connector landing

<a id="canonical-3086c8cf07670672fcc669a57a540097dca1a3c427b2fe0c893d9acee8740f16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e2df4748cf2d66e07f65a088907defabb768f4bc355162fbb63a6b665e94a54"></a>

## xcsh_external_connector — xcsh_external_connector / d80d4afc5989 / 2

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

<a id="canonical-55e1d4672d1e12e8ad396faf0cf0af21b9ea9b10d0a73b89848bddb758032fb2"></a>

## Prerequisites — xcsh_external_connector / d80d4afc5989 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-b76ce85245aaa74bfaf5623b931f3eb09604348e7621570ad7c6b65008a51ae5"></a>

## Minimal configuration — xcsh_external_connector / d80d4afc5989 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-e05b655a574859ce7bd2f3571ace17877a621635e88b6c5d4ae9d20e6717398b"></a>

## Root configuration — xcsh_external_connector / d80d4afc5989 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8de446d0fdba3053312a6f3a06f4af6533c5f74f8e41de030a6764cbf688d6ca"></a>

## Next pages — xcsh_external_connector / d80d4afc5989 / 6

- [Property reference](../guides/data-sources--external_connector--reference--group-001.md#canonical-6f077a223c3703599a4a3e823820f48c4074bd2d98d5da48dc00ffcce261f1fe)
- [Examples](../guides/data-sources--external_connector--examples--group-001.md#canonical-5b91edc54625be81d0068af801788d9ede1571781cfd3411cf395f1b0ca5ec31)
