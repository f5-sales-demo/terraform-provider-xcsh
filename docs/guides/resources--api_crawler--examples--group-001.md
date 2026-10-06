---
page_title: "xcsh_api_crawler examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler examples."
---

# xcsh_api_crawler examples

<a id="canonical-0031220333210232-0332113233210223-2322322103210201-2233302012301233-1213103110011133-1032010212232333-3332202102103212-1000221333122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- Examples

<a id="canonical-1131322321323210-0310330333221012-1233331302132330-2220323332210200-3123111330303320-0123113212121331-3133323110130032-0001133211210113"></a>

### Complete configurations for `xcsh_api_crawler`

- [Resource](resources--api_crawler--examples--group-001.md#canonical-1333320203321222-1031031111030302-2110210221023332-2020100021013313-2022001012210210-2112211123110133-3211311312211331-3102211222213323): valid configuration.

<a id="canonical-1333320203321222-1031031111030302-2110210221023332-2020100021013313-2022001012210210-2112211123110133-3211311312211331-3102211222213323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_api_crawler](../resources/api_crawler.md#canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200)
- [Examples](resources--api_crawler--examples--group-001.md#canonical-0031220333210232-0332113233210223-2322322103210201-2233302012301233-1213103110011133-1032010212232333-3332202102103212-1000221333122211)
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
