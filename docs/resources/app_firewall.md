---
page_title: "xcsh_app_firewall landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall landing."
---

# xcsh_app_firewall landing

<a id="canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110310010003313-3320013133323112-2032303303300221-0133311230222301-3200203232101322-2310221323132010-3101302203223302-2201330022132223"></a>

## xcsh_app_firewall — xcsh_app_firewall / 322313101232 / 2

Breadcrumbs:

- xcsh_app_firewall

Manages Application Firewall in F5 Distributed Cloud.

<a id="canonical-0031030321302021-3321022311113301-2123012202021301-0103100032332220-2332320002010010-1103331010111311-0323300221203332-3230312121122103"></a>

## Prerequisites — xcsh_app_firewall / 322313101232 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `service_policy`.

- service_policy: Fine-grained access control rules

<a id="canonical-3032013330000230-1032213302233103-1212031221231220-1203003021312220-1232211303130212-1020130212323220-2311000201032223-3123313302331023"></a>

## Minimal configuration — xcsh_app_firewall / 322313101232 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppFirewall Resource Example
# Manages Application Firewall in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppFirewall configuration
resource "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}
```

<a id="canonical-0113231130020231-1130333201232020-2200322322101232-3303223123203030-1222110102330233-3123221201233222-1213001330300222-0030120230030233"></a>

## Root configuration — xcsh_app_firewall / 322313101232 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3223312201331213-0321223203120331-0101012001221013-2323233321133123-3122300213122200-3233033212011300-1100301113112331-1033131303022032"></a>

## Next pages — xcsh_app_firewall / 322313101232 / 6

- [Property reference](../guides/resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [Examples](../guides/resources--app_firewall--examples--group-001.md#canonical-3223331122212131-3221011031110000-3122201120311300-2021200030110011-0220321000300132-0202212232110313-0013011132121033-2101011320103031)
- [Import](../guides/resources--app_firewall--lifecycle--group-001.md#canonical-3033201022113101-3211011303131013-0032132120121103-0122023302120031-0320321330332110-1122232103021330-2201123000113020-1323003211223223)
- [Timeouts](../guides/resources--app_firewall--lifecycle--group-001.md#canonical-0201322303130233-0211322300320302-3103320020130212-0002131301303112-3312130030322220-0031322300113210-0213110301010332-0300111320323202)
