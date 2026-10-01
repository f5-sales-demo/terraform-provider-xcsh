---
page_title: "xcsh_api_crawler landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_crawler landing."
---

# xcsh_api_crawler landing

<a id="canonical-3133010202200222-3213011120000312-1321322311223200-2311232132101333-2211031002320321-3221011221222302-0233012202301010-3201312221013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311221122320200-0101231313302111-0203121002100212-0030121010320031-3220310023130222-2231313332032303-0312211000312130-2101111001113232"></a>

## xcsh_api_crawler — xcsh_api_crawler / 031022211212 / 2

Breadcrumbs:

- xcsh_api_crawler

Manages a API Crawler resource in F5 Distributed Cloud.

<a id="canonical-0112212132010213-2320112133301022-2033001132100333-0120113230111132-2130213302002203-3120010133131011-0121000203220222-1012203121121330"></a>

## Prerequisites — xcsh_api_crawler / 031022211212 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2121030131031033-3213220202001323-0023031213211010-2210333122300103-0200320122203021-3302032021100033-0022321200233203-3010222123303013"></a>

## Minimal configuration — xcsh_api_crawler / 031022211212 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0100020232323100-3231021330101120-2300320022131131-2331113323200113-0210333303131322-3110303323131303-0310203133132031-3102313021003103"></a>

## Root configuration — xcsh_api_crawler / 031022211212 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3201032100310200-3113002121223302-1023320313022113-2310122131323300-0102012321013202-3312233222130131-2310212203200212-0203130132123201"></a>

## Next pages — xcsh_api_crawler / 031022211212 / 6

- [Property reference](../guides/resources--api_crawler--reference--group-001.md#canonical-3010222002133020-2211321021131000-2301002232300112-1123303232312202-2000011002203003-2003112030230322-2103211310110013-1020232220123011)
- [Examples](../guides/resources--api_crawler--examples--group-001.md#canonical-0031220333210232-0332113233210223-2322322103210201-2233302012301233-1213103110011133-1032010212232333-3332202102103212-1000221333122211)
- [Import](../guides/resources--api_crawler--lifecycle--group-001.md#canonical-1111111322222323-0131002203332011-1333010222221202-2202312231020213-1321333331222033-2031323033033130-1332032021020003-0300103012010011)
- [Timeouts](../guides/resources--api_crawler--lifecycle--group-001.md#canonical-1111102132232132-2032111302213101-2323321003100130-1312231123312000-2310132200101222-1323203201330301-0333313111032120-3123222021313030)
