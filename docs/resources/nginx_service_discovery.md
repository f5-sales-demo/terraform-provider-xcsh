---
page_title: "xcsh_nginx_service_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery landing."
---

# xcsh_nginx_service_discovery landing

<a id="canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333310123002120-2120231010010102-3130003121320002-3030120102230003-3333330111221131-2222222032131003-2200320112011202-0012313311211022"></a>

## xcsh_nginx_service_discovery — xcsh_nginx_service_discovery / 003010312022 / 2

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

<a id="canonical-1330202031230222-3113211330210202-1210231300300332-0102032330302023-3212001201230201-0121221210211011-3110011110222320-3312212233122021"></a>

## Prerequisites — xcsh_nginx_service_discovery / 003010312022 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2332331202222331-3313202310201203-0022131221121311-3313210111233221-3232321031230230-0213132233303003-3112003333030003-2010323011133321"></a>

## Minimal configuration — xcsh_nginx_service_discovery / 003010312022 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```

<a id="canonical-3021102220220110-0232012212203322-1002033213301310-3332221102313132-0302200132300031-0232133002321110-2132320323320123-3322323331003011"></a>

## Root configuration — xcsh_nginx_service_discovery / 003010312022 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3000031133103023-1210200303103021-2331231022313030-2133020113102013-3233203123103333-2313111111120203-0210001221023302-2303103023223220"></a>

## Next pages — xcsh_nginx_service_discovery / 003010312022 / 6

- [Property reference](../guides/resources--nginx_service_discovery--reference--group-001.md#canonical-1330100322001110-1113231031021202-3122022020201302-2211321310100221-0202312133002311-1211013321231330-0032333223202203-3203133113300021)
- [Examples](../guides/resources--nginx_service_discovery--examples--group-001.md#canonical-3001202301033101-0031000202120112-0333302030330103-1200011310233021-2031103233001231-1131103122222132-2013233222111311-0012112302231222)
- [Import](../guides/resources--nginx_service_discovery--lifecycle--group-001.md#canonical-3211211103013130-3312112131130211-0033312121310031-2022102113030313-0000333032102230-2313330101300213-0220332311230320-1232011101200030)
- [Timeouts](../guides/resources--nginx_service_discovery--lifecycle--group-001.md#canonical-2300321112132330-3003031320022123-1000330323231120-3021330100010233-0322022110203130-3211132303013033-2331011012020032-3220222320313113)
