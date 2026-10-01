---
page_title: "xcsh_api_crawler examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler examples."
---

# xcsh_api_crawler examples

<a id="canonical-0da3f92e3e5ef92bbae93921afc86c6f674d415f4e126bbffe8924e640a7f6a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5deb9ee434f3fa466ff727bca8efe920db57ccf81b5e667ddfed470e017e5917"></a>

## Examples — Examples / ff1826899cb6 / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- Examples

<a id="canonical-05241206a795fd7e44d157407a8f4f804c55b5d08a6d4c3a75b89e11c57f3f7f"></a>

## Complete configurations — Examples / ff1826899cb6 / 3

- [Resource](resources--api_crawler--examples--group-001.md#canonical-7fe23e6a4d355332949292fe884091f78a0469249695b51fe5d7697dd296a9fb): valid configuration.

<a id="canonical-8e55aad8a5f986a24b5c9b9ea833242c7e29b1c6ec944b6d5e7553aa18d06e25"></a>

## Next pages — Examples / ff1826899cb6 / 4

- [Resource](resources--api_crawler--examples--group-001.md#canonical-7fe23e6a4d355332949292fe884091f78a0469249695b51fe5d7697dd296a9fb)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)

<a id="canonical-7fe23e6a4d355332949292fe884091f78a0469249695b51fe5d7697dd296a9fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8e24572b4ecfcea4429f35d3a68d86d5a53d9a8997f7ee253b40ff578f5e106"></a>

## Resource — Resource / d047c997b6ff / 2

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
- [Examples](resources--api_crawler--examples--group-001.md#canonical-0da3f92e3e5ef92bbae93921afc86c6f674d415f4e126bbffe8924e640a7f6a5)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_crawler/resource.tf`; digest `sha256:a73dd5a1c2c2dd7d614c3510b02a2ee9ecb351afd8ffe9c697a5142ace64fc30`.

```terraform
# APICrawler Resource Example
# Manages a API Crawler resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APICrawler configuration
resource "xcsh_api_crawler" "example" {
  name      = "example-api-crawler"
  namespace = "staging"
}
```

<a id="canonical-b4ef802667b374c09b155be636b51e6151c163895dd80a10647d271111dd5037"></a>

## Next pages — Resource / d047c997b6ff / 3

- [Examples](resources--api_crawler--examples--group-001.md#canonical-0da3f92e3e5ef92bbae93921afc86c6f674d415f4e126bbffe8924e640a7f6a5)
- [xcsh_api_crawler](../resources/api_crawler.md#canonical-df12282ae715803679eb5ae0b5b9e47fa5342e39e9169ab22f1a2c44e1da91e0)
