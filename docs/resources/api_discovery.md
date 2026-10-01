---
page_title: "xcsh_api_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_discovery landing."
---

# xcsh_api_discovery landing

<a id="canonical-3333330030203101-0213011320123210-0101011221311003-1110030033021003-0023100303333012-1213112122012213-3330220223223100-0213222121332310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003100000331333-0002132232212213-0032302200332221-2033221133300032-0012303110312301-0113211110311012-0103012111210000-0101023211202211"></a>

## xcsh_api_discovery — xcsh_api_discovery / 320312212121 / 2

Breadcrumbs:

- xcsh_api_discovery

Manages API discovery creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-3202200032120203-1202303223220001-3013202112310131-1112232213002320-1130221313302201-2023322201212033-3330011122232022-1302102302300122"></a>

## Prerequisites — xcsh_api_discovery / 320312212121 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2100332110223330-3201102231101103-3330320120312330-3231000212010022-0003130330000232-3122033220331210-1202110002130311-2200023313221233"></a>

## Minimal configuration — xcsh_api_discovery / 320312212121 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

<a id="canonical-0303102133313233-2001303222310230-0223222213230100-3020100210010133-3120012020031222-3022323001123333-1201300122103001-0310113121331220"></a>

## Root configuration — xcsh_api_discovery / 320312212121 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2213221013003133-2001211033023302-1012313223031011-3120123100220103-3210112220012202-1202102321111012-3100012202321133-1220330302331013"></a>

## Next pages — xcsh_api_discovery / 320312212121 / 6

- [Property reference](../guides/resources--api_discovery--reference--group-001.md#canonical-1113112330221303-3231101122212110-2001031202133133-0003330200133112-1300003330331131-1120302331131301-1010000310313122-3333223113222003)
- [Examples](../guides/resources--api_discovery--examples--group-001.md#canonical-2103330003101230-1200233231321132-1332020330220230-0211002003332330-1301223301010230-1123100002232003-0133303013301232-2212101303032210)
- [Import](../guides/resources--api_discovery--lifecycle--group-001.md#canonical-1311311330302132-1233323123001231-3203311022310121-1030302022212313-3111233121311110-0133101031311132-3322101000100233-3333123330132131)
- [Timeouts](../guides/resources--api_discovery--lifecycle--group-001.md#canonical-0013002101201022-1220320001222012-0023000302230130-1111121032320001-2003312101331203-1100101232112130-2011222030302000-1232100132102231)
