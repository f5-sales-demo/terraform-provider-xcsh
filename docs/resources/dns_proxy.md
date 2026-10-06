---
page_title: "xcsh_dns_proxy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy."
---

# xcsh_dns_proxy

<a id="canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_dns_proxy

Manages DNS Proxy in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

<a id="canonical-0210132203020131-0111003103222113-2111011130231303-0032112111101202-1103133200030202-3202032130321331-1300020213131032-2303212322003023"></a>

### Prerequisites for `xcsh_dns_proxy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3112110233001112-2112330010022121-0030213001032112-3003003330010230-2121111332232132-2300110332023130-1330031112102230-0221321032003202"></a>

### Minimal configuration for `xcsh_dns_proxy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSProxy Resource Example
# Manages DNS Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSProxy configuration
resource "xcsh_dns_proxy" "example" {
  name      = "example-dns-proxy"
  namespace = "system"
}
```

<a id="canonical-2321110033202303-3301131011101211-3232111001002123-2300221011102003-1103002212300003-0233200303120133-2212123123211102-1033310121113302"></a>

### Root configuration for `xcsh_dns_proxy`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3331023323303203-1120112101010320-3231000022332001-1210232110201201-2210020213032000-0111001222233131-3033221302202330-3222332000113103"></a>

### Explore this collection for `xcsh_dns_proxy`

- [Property reference](../guides/resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [Examples](../guides/resources--dns_proxy--examples--group-001.md#canonical-2133222100312232-1221112223202121-2002020000120023-2022033232120030-3110313330333121-1201030313001033-2320210131022111-0212023110003211)
- [Import](../guides/resources--dns_proxy--lifecycle--group-001.md#canonical-3303233331131231-1203112221022131-1013330320213321-3003331120010332-2132111030023201-0001023321302302-0313200222313111-0320223111222131)
- [Timeouts](../guides/resources--dns_proxy--lifecycle--group-001.md#canonical-1013001223032121-0233113100331033-3320320230331231-2103203102202323-1232031301230022-0022221101133303-0130002012022112-3112321122110200)
