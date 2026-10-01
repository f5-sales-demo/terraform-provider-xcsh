---
page_title: "xcsh_nginx_service_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery landing."
---

# xcsh_nginx_service_discovery landing

<a id="canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310222203223020-0201123212211000-3302132323312212-3231111012221032-0332223112130021-3112210331002013-2111320002100120-0321322223213221"></a>

## xcsh_nginx_service_discovery — xcsh_nginx_service_discovery / 222002003020 / 2

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

<a id="canonical-3123210100000322-3321002230222321-2123321331102002-0121201311011022-0102022110232300-2020300303032330-2120210102210313-1223320333320200"></a>

## Prerequisites — xcsh_nginx_service_discovery / 222002003020 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0210132310111121-2001313122313003-3211202012021123-2102302031103102-2330210221223013-0213101233212211-0223210220201021-0111023100003213"></a>

## Minimal configuration — xcsh_nginx_service_discovery / 222002003020 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```

<a id="canonical-1303301131132230-2300331310302023-2223120231110030-0012113012232130-1223320103232301-0322201121101200-2032330331233113-0333001123322130"></a>

## Root configuration — xcsh_nginx_service_discovery / 222002003020 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3002103220100110-1312311310321210-2001323300022321-1320003031332023-3120122100320201-0002021211310020-3230330102122121-3203101031131220"></a>

## Next pages — xcsh_nginx_service_discovery / 222002003020 / 6

- [Property reference](../guides/data-sources--nginx_service_discovery--reference--group-001.md#canonical-2331133322323211-2212201203300210-0302021312100033-3110301002321021-2121133003033222-2312100130121302-1001330211000330-1211201201023212)
- [Examples](../guides/data-sources--nginx_service_discovery--examples--group-001.md#canonical-1312122322200101-0103213331231131-3012332303232320-3111010023031103-1130020230030100-0101103120101222-1112023222022322-2320032102130001)
