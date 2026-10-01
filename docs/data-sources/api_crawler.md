---
page_title: "xcsh_api_crawler landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler landing."
---

# xcsh_api_crawler landing

<a id="canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9d0c995267d28a2017cd8da2586adf3d28fe348ce48064322198e2bbfbe36e7"></a>

## xcsh_api_crawler — xcsh_api_crawler / 2520eebd8ba7 / 2

Breadcrumbs:

- xcsh_api_crawler

Manages a API Crawler resource in F5 Distributed Cloud.

<a id="canonical-320227868d3e61e2ad3dced2f3b8a9c495ec6cc94775013e2d53b0d7c669b49d"></a>

## Prerequisites — xcsh_api_crawler / 2520eebd8ba7 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-69b5213da288b7bdcaf8df9061c42f5f2bca69e79caa08bd6a5f29efe50b0269"></a>

## Minimal configuration — xcsh_api_crawler / 2520eebd8ba7 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APICrawler Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APICrawler by name
data "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}

output "api_crawler_id" {
  value = data.xcsh_api_crawler.example.id
}
```

<a id="canonical-cd1098b370fc23e002bee930bfdc33a113442bcbe986cc3a3bb973a6c57dfb7b"></a>

## Root configuration — xcsh_api_crawler / 2520eebd8ba7 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-598954dfabc0c17cb4a521b77c11919fc6ddce16ea2f1cd2ca99604b63e4e07b"></a>

## Next pages — xcsh_api_crawler / 2520eebd8ba7 / 6

- [Property reference](../guides/data-sources--api_crawler--reference--group-001.md#canonical-28a3aac178dc4be3f000c830210b2c47249b673430fc2cf275ec9a1b9ece900b)
- [Examples](../guides/data-sources--api_crawler--examples--group-001.md#canonical-cff911ec166cbf5a694bf3603257154d58c94c5d4bb92d5158a71f07ff81a815)
