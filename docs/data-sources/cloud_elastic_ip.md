---
page_title: "xcsh_cloud_elastic_ip landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip landing."
---

# xcsh_cloud_elastic_ip landing

<a id="canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111302203223012-1131302212211102-3201003332310001-2311310001212121-1222330321011113-0020103323331103-2213123131203021-3112011331332111"></a>

## xcsh_cloud_elastic_ip — xcsh_cloud_elastic_ip / 320032000031 / 2

Breadcrumbs:

- xcsh_cloud_elastic_ip

Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5
Distributed Cloud.

<a id="canonical-2210332221100233-0130323231010020-0231121210200122-1212310320023021-0221203032101103-0202013222231323-1122030020211323-3001310023113303"></a>

## Prerequisites — xcsh_cloud_elastic_ip / 320032000031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1032002133210301-3030321200110321-3210122130020231-1331002312223123-1330012300023212-3222012301213112-2130011310222012-1313011323212130"></a>

## Minimal configuration — xcsh_cloud_elastic_ip / 320032000031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudElasticIP by name
data "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"
}

output "cloud_elastic_ip_id" {
  value = data.xcsh_cloud_elastic_ip.example.id
}
```

<a id="canonical-1232000201331302-2011200212231320-3010123211331200-2212013003223102-0022223013311103-1232131300302120-0000212022201230-2302301023212131"></a>

## Root configuration — xcsh_cloud_elastic_ip / 320032000031 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1021030310031222-2032012101320123-3032203020021331-1201300113213012-1110333301103131-3032302002121303-1323003310310101-1201001210131333"></a>

## Next pages — xcsh_cloud_elastic_ip / 320032000031 / 6

- [Property reference](../guides/data-sources--cloud_elastic_ip--reference--group-001.md#canonical-3233011131021232-3231202322001200-0221311312333302-0113233211312111-3032320102211210-0313320223303133-3011223131121200-0001033220313023)
- [Examples](../guides/data-sources--cloud_elastic_ip--examples--group-001.md#canonical-0013120111021103-0230021112331013-0312230200221322-2231002130302222-0203222212300212-0002203313103020-0032030222303213-2110302301220132)
