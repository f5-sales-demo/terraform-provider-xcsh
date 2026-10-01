---
page_title: "xcsh_api_definition landing"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition landing."
---

# xcsh_api_definition landing

<a id="canonical-0122213120002012-3231302001220300-2001233123112233-1130330020321211-3113132132002121-2133030200111003-0210010013030021-1211030320012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230331013122032-0320023120120131-1120332302001001-0332231203322013-2001233023310113-2333032303032033-3132023311223132-0002021210120311"></a>

## xcsh_api_definition — xcsh_api_definition / 122323021311 / 2

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

<a id="canonical-3023211211222233-2033211120302002-2303003223203001-2100003100013002-3000030333001302-0000021122101003-2130200201201300-0011220111222310"></a>

## Prerequisites — xcsh_api_definition / 122323021311 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

<a id="canonical-1321312003020100-2102311313102030-3322033331120222-1300013023011003-0213230300202213-3311210310330101-3030333123022332-3222030210323223"></a>

## Minimal configuration — xcsh_api_definition / 122323021311 / 4

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

<a id="canonical-3210300002003002-0002303121123011-0131030233321212-3132131203020321-0012022120203313-3101223320330213-2112110211002110-1311323223230331"></a>

## Root configuration — xcsh_api_definition / 122323021311 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2222022222012023-1101210003223313-2310210212310323-1121023110031322-1120222001113323-1031223200100113-0230210323020213-3002321012120112"></a>

## Next pages — xcsh_api_definition / 122323021311 / 6

- [Property reference](../guides/resources--api_definition--reference--group-001.md#canonical-1013013311212220-2100012121010012-2012032113201303-3011001133121233-0013201202222002-0113112031031113-2000220121302133-2321231322222211)
- [Examples](../guides/resources--api_definition--examples--group-001.md#canonical-1021200213031100-3331111211312013-1031323303301031-0322221202302021-2220021101122011-3011100220301120-2331131031230210-2001320110010210)
- [Import](../guides/resources--api_definition--lifecycle--group-001.md#canonical-2232021313301213-0120030010233332-0123231113030312-3033030122002113-3221001022032312-3200112300230200-3022031313321001-2123110232233111)
- [Timeouts](../guides/resources--api_definition--lifecycle--group-001.md#canonical-2002022223022103-3302230110102312-0030110120003121-1321201003010231-1013100221110231-2303332123222300-3223030011032000-3222210020121330)
