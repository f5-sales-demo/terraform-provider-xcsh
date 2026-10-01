---
page_title: "xcsh_virtual_host landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host landing."
---

# xcsh_virtual_host landing

<a id="canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020232232102113-1001111321020212-2320313200023223-3201202121020322-2132200233322013-0001223333302122-0300320312332130-0320032222230122"></a>

## xcsh_virtual_host — xcsh_virtual_host / 001313131331 / 2

Breadcrumbs:

- xcsh_virtual_host

Manages virtual host in a given namespace in F5 Distributed Cloud.

<a id="canonical-3211203202111233-3011030030322001-2123001010323233-3310030331021311-1221231223132202-0311123013312130-3003102112012301-3331302200333010"></a>

## Prerequisites — xcsh_virtual_host / 001313131331 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2312020303022121-3131110221103022-3010131031330000-2013101223220233-2310012231023103-2100111102013202-0221231133203110-0002210330010212"></a>

## Minimal configuration — xcsh_virtual_host / 001313131331 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualHost Resource Example
# Manages virtual host in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}
```

<a id="canonical-0103122023330011-2322120012022001-1010030232212330-1302032333311103-0212112320233222-2301012101212012-2123231221121322-0311330311331332"></a>

## Root configuration — xcsh_virtual_host / 001313131331 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0301001210312233-2100002012312210-2210123232220210-2300221100032223-0211032313001010-3320230301332322-3032021300121020-0323331100322331"></a>

## Next pages — xcsh_virtual_host / 001313131331 / 6

- [Property reference](../guides/resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [Examples](../guides/resources--virtual_host--examples--group-001.md#canonical-1210132032103233-2030210232320032-1132323223302032-0101302113222301-0232320020302021-2032322012000331-2220323303330101-1302002211023223)
- [Import](../guides/resources--virtual_host--lifecycle--group-001.md#canonical-3010122311030023-3301131312110032-1003122023332231-3031220211222311-0010113103131003-3110200330233031-1023022301030303-3300223123110301)
- [Timeouts](../guides/resources--virtual_host--lifecycle--group-001.md#canonical-2130013001012310-3332011203200100-0233011030320013-2230032200323033-2021023221323203-1021200002002121-3032000123132221-0310221131031010)
