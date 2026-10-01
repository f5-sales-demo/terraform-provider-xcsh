---
page_title: "xcsh_site_upgrade_sw landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_sw landing."
---

# xcsh_site_upgrade_sw landing

<a id="canonical-0120130312312320-0223010000032113-2233123110100201-3130110122132200-2323212123211333-3120313220122233-0023002110221112-0320032112111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321200332321123-0333320321103113-3101320303231220-0121203332131223-0110331032301022-1013003230222131-3121000122133210-2021121203113300"></a>

## xcsh_site_upgrade_sw — xcsh_site_upgrade_sw / 033020332033 / 2

Breadcrumbs:

- xcsh_site_upgrade_sw

Request an in-place site software upgrade.

<a id="canonical-2210032002032112-2132213113122100-0111203321031023-0032133313200030-0111330210311200-3331120321110012-2231311102230130-2232002021133100"></a>

## Prerequisites — xcsh_site_upgrade_sw / 033020332033 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2033021033122311-3012322202002210-3232020022223021-3112001032123033-2223301113013013-2013102012320201-0121000102201320-0210212200120013"></a>

## Minimal configuration — xcsh_site_upgrade_sw / 033020332033 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteUpgradeSw Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts the upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_sw" "example" {
  config {
    site             = "example-value"
    software_version = "example-value"
  }
}
```

<a id="canonical-2323313123103201-3321220312131202-1322313002233323-0003111132310001-1003303231213010-0310003233212100-0312222200212323-1300010322221232"></a>

## Root configuration — xcsh_site_upgrade_sw / 033020332033 / 5

Required root properties: `site`, `software_version`. Full root flags and choices appear in the property reference.

<a id="canonical-2123230321132013-1202111030120131-2313123201031102-0223012321331230-0312311310133132-3322103001303103-3032122000303030-1231101200111020"></a>

## Next pages — xcsh_site_upgrade_sw / 033020332033 / 6

- [Property reference](../guides/actions--site_upgrade_sw--reference--group-001.md#canonical-0030032002321101-2000121030330322-0213020132122121-0103333103102130-2121013213001013-3311301322120002-1131123013013223-2002122103100032)
- [Examples](../guides/actions--site_upgrade_sw--examples--group-001.md#canonical-1133112323300000-3332003131302131-2300100101022322-3033010311201200-2203211331101220-1232230303131300-3331212323021122-2322203021032110)
- [Lifecycle](../guides/actions--site_upgrade_sw--lifecycle--group-001.md#canonical-0203222121000201-1320112033001313-0233310313211013-0210322212131220-0033002231212133-2033011102213110-3121102110131230-2201222003030233)
