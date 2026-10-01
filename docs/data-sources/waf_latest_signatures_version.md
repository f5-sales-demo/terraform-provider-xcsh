---
page_title: "xcsh_waf_latest_signatures_version landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_latest_signatures_version landing."
---

# xcsh_waf_latest_signatures_version landing

<a id="canonical-3220021120200003-2131100022002331-1323212033111120-3012322121033220-3000100023133103-1211320023113021-0322300310221012-2100112203232030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223001010113311-0000201031111202-1311300323211213-2330231222322102-2013010300103331-2231023232233112-0311110121101213-0033322032133301"></a>

## xcsh_waf_latest_signatures_version — xcsh_waf_latest_signatures_version / 332032100333 / 2

Breadcrumbs:

- xcsh_waf_latest_signatures_version

Resource retrieval operation.

<a id="canonical-2201130230222311-0210323202302221-2303131300122331-1110311012322122-1223321020102013-1132013001312231-3131232230220200-0211001013211312"></a>

## Prerequisites — xcsh_waf_latest_signatures_version / 332032100333 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1312232311211330-2332002303133313-0022130100030122-3131212131120331-2011231020313033-0202303201212123-2331311011222103-2110323113201010"></a>

## Minimal configuration — xcsh_waf_latest_signatures_version / 332032100333 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFLatestSignaturesVersion DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_latest_signatures_version" "example" {
}

output "waf_latest_signatures_version_result" {
  value = data.xcsh_waf_latest_signatures_version.example
}
```

<a id="canonical-0200033322020013-1010133231031013-2002111000220323-2000121123110231-2223121110303031-0201013310202133-0223130103031022-3323012001200222"></a>

## Root configuration — xcsh_waf_latest_signatures_version / 332032100333 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2212300233003002-0230220030031333-0102303302023132-1122320030332210-0312201000332230-3210123130133203-0311010012201311-3111231302202130"></a>

## Next pages — xcsh_waf_latest_signatures_version / 332032100333 / 6

- [Property reference](../guides/data-sources--waf_latest_signatures_version--reference--group-001.md#canonical-1010221311131332-3202023031130332-1201310131202103-2003100331322012-0003203310020231-1120033302003022-0230023002232313-3000132321103113)
- [Examples](../guides/data-sources--waf_latest_signatures_version--examples--group-001.md#canonical-2131332222021330-0302033333003232-1303001311302012-1102301302223003-3111333301021202-1020200000001103-2012120003202223-1032001220123110)
