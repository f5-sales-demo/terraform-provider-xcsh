---
page_title: "xcsh_api_crawler examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler examples."
---

# xcsh_api_crawler examples

<a id="canonical-3033332101013230-0112123023331122-1221102333031200-0302111301111031-1120302110301131-1023232102311101-1120221301330013-3333200122200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- Examples

<a id="canonical-2333303311100033-1112002123130012-3320221113303213-3332011011301000-1321033230330332-0132301310132022-0013231100202122-1303120322131311"></a>

### Complete configurations for `xcsh_api_crawler`

- [Data source](data-sources--api_crawler--examples--group-001.md#canonical-3232313232030132-1313101131032321-3321122323312312-2121023201302333-1131210120221023-0201010031003033-3221100000321310-3302203000200300): valid configuration.

<a id="canonical-3232313232030132-1313101131032321-3321122323312312-2121023201302333-1131210120221023-0201010031003033-3221100000321310-3302203000200300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_api_crawler](../data-sources/api_crawler.md#canonical-3032313323203012-2021301131010010-1211333201122111-1021320230300330-0211003112312030-1200030031330121-3012320020032333-2333221312222211)
- [Examples](data-sources--api_crawler--examples--group-001.md#canonical-3033332101013230-0112123023331122-1221102333031200-0302111301111031-1120302110301131-1023232102311101-1120221301330013-3333200122200111)
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
