---
page_title: "xcsh_api_definition landing"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition landing."
---

# xcsh_api_definition landing

<a id="canonical-1a9d8086edc81a3081bdb5af5cf08e65d779e0999f32054324107309653381b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acf4768e382d861d58fb20413eb63e8781bcbd17bf3b338fde2f5ade02264635"></a>

## xcsh_api_definition — xcsh_api_definition / d98cc36bb275 / 2

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

<a id="canonical-cb965aaf8f958c82b30eb8c1900d01c2c033f0720025a4439c82187005a15ab4"></a>

## Prerequisites — xcsh_api_definition / d98cc36bb275 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

<a id="canonical-79d8321092d7748cfa3fd62a701cb14327b308a7f5934f11ccfdb2beea324eeb"></a>

## Minimal configuration — xcsh_api_definition / d98cc36bb275 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```

<a id="canonical-e4c020c202cd96c51d32fe66de763239062988f7d1af8f279652509475eebb3d"></a>

## Root configuration — xcsh_api_definition / d98cc36bb275 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-aa2aa18b51903af7b4926d3b592d437a58a815fb4dae04172c93b227c2e46616"></a>

## Next pages — xcsh_api_definition / d98cc36bb275 / 6

- [Property reference](../guides/resources--api_definition--reference--group-001.md#canonical-471f59a89019910686397873c505f66f07862a821758d35780a19c9fb9b7aaa5)
- [Examples](../guides/resources--api_definition--examples--group-001.md#canonical-49827350fd565d874def3c4d3aa62c89a8251685c5428c58bd74db2481e14124)
- [Import](../guides/resources--api_definition--lifecycle--group-001.md#canonical-ae277c6718304bfe1bb57336cf31a097e904a3b6e05b0b20ca377e419b52ebd5)
- [Timeouts](../guides/resources--api_definition--lifecycle--group-001.md#canonical-822ab293f2b144b60c5180d97984312d4742952db3f9bab0eb305380ea90867c)
