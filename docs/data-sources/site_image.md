---
page_title: "xcsh_site_image landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_image landing."
---

# xcsh_site_image landing

<a id="canonical-3131222302011312-0030101332231322-0021111121002033-3033310200203330-1301310123023111-0121222100300001-3102203221303023-0331320320202102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113223223320113-3333021211112332-3320323003002000-0021221112233322-3130200020002303-2233113101200303-2230201123011102-1300133032031103"></a>

## xcsh_site_image — xcsh_site_image / 332200113113 / 2

Breadcrumbs:

- xcsh_site_image

Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No
caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution
does not imply successful boot.

<a id="canonical-3202330023121122-0330320300212121-2210213221303200-0333110132203022-3023220112333311-0111013203213322-0320122200021211-0010022111121311"></a>

## Prerequisites — xcsh_site_image / 332200113113 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1021023113323101-1123121123022201-0022222322331122-0031322101131130-3031303100002123-3313211003120212-1231203320202300-3221013231323000"></a>

## Minimal configuration — xcsh_site_image / 332200113113 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

<a id="canonical-2120202223321221-1000122003012332-1033131212220101-0333031210111031-2213311213122303-0223203233213021-0101200201121203-1113130310210113"></a>

## Root configuration — xcsh_site_image / 332200113113 / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-0330021223103133-3102221321102110-3033130110211233-3212012333012330-2000103102033212-3321231011220132-2330031202233232-3031322022023011"></a>

## Next pages — xcsh_site_image / 332200113113 / 6

- [Property reference](../guides/data-sources--site_image--reference--group-001.md#canonical-3230203203002131-2312110021103212-3200131121210123-1302203223102322-2302331310010222-3030113311121323-3030201010300123-3222123123130310)
- [Examples](../guides/data-sources--site_image--examples--group-001.md#canonical-1012123311312022-0323311010031330-1232003023023331-0010122132312101-1231331122323201-0003202220330030-2222032332303213-3312032213033003)
