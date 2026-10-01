---
page_title: "xcsh_api_crawler examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler examples."
---

# xcsh_api_crawler examples

<a id="canonical-cff911ec166cbf5a694bf3603257154d58c94c5d4bb92d5158a71f07ff81a815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfcf540f5609b706f8a57ce7fe145c40793ecf3e1ec7478a07b5089a7363a775"></a>

## Examples — Examples / 3f0d3b13018d / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- Examples

<a id="canonical-47823340c0c30125c571584e5a4448b55d2be730a47991fa30cbb0286eeb3f7f"></a>

## Complete configurations — Examples / 3f0d3b13018d / 3

- [Data source](data-sources--api_crawler--examples--group-001.md#canonical-eedee31e7745d3b9f96bbdb6992e1cbf5d918a4b2110d0cfe9400e74f28c0830): valid configuration.

<a id="canonical-247778e1e21e78feed2330693fc139474dcdd223f7942576f2564bbf1f044647"></a>

## Next pages — Examples / 3f0d3b13018d / 4

- [Data source](data-sources--api_crawler--examples--group-001.md#canonical-eedee31e7745d3b9f96bbdb6992e1cbf5d918a4b2110d0cfe9400e74f28c0830)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)

<a id="canonical-eedee31e7745d3b9f96bbdb6992e1cbf5d918a4b2110d0cfe9400e74f28c0830"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78946fcc07846c97ca651e5228e8a0285cc2c9eec4b85dd6665bc3e307799caa"></a>

## Data source — Data source / 8988cd88b0d5 / 2

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
- [Examples](data-sources--api_crawler--examples--group-001.md#canonical-cff911ec166cbf5a694bf3603257154d58c94c5d4bb92d5158a71f07ff81a815)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_crawler/data-source.tf`; digest `sha256:3bf584a1f7a1b8c51ad4462cae33ab926ce0db3d2feff9eb685959704de3f510`.

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

<a id="canonical-699b7cf50546f7fedac27b159b7a36652b3db313d70abc3610b23b895cf09bbd"></a>

## Next pages — Data source / 8988cd88b0d5 / 3

- [Examples](data-sources--api_crawler--examples--group-001.md#canonical-cff911ec166cbf5a694bf3603257154d58c94c5d4bb92d5158a71f07ff81a815)
- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-cedfb8c689c5d10465fe169549e2cc3c250d6d8c6030df19c6e083bfbfa76aa5)
