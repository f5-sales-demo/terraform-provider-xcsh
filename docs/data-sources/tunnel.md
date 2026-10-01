---
page_title: "xcsh_tunnel landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel landing."
---

# xcsh_tunnel landing

<a id="canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000332332012022-2230202200201120-0212031300223022-1130230022032023-3002223231122111-3331211022100033-0101312201223103-2211001000011130"></a>

## xcsh_tunnel — xcsh_tunnel / 033122213132 / 2

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-0301213020020033-3231101000313000-0123212210021330-0302230320323033-0033220223103332-2223021331102322-1112302032013023-2000002111101021"></a>

## Prerequisites — xcsh_tunnel / 033122213132 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1211211202213330-3202312031212113-3031211111033121-0223013122231013-1023331113203101-1132031203231303-0122320220100101-3022210330202133"></a>

## Minimal configuration — xcsh_tunnel / 033122213132 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

<a id="canonical-2210020221300002-1132321130031301-1013203302001001-0222121330222022-2321220003110312-2300200330031223-0013110202131120-0023300201323112"></a>

## Root configuration — xcsh_tunnel / 033122213132 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1313320010110113-2310103020010231-1021303102011120-3030330320032200-2211310101103032-0312333131023003-3303222220130001-2032302032123222"></a>

## Next pages — xcsh_tunnel / 033122213132 / 6

- [Property reference](../guides/data-sources--tunnel--reference--group-001.md#canonical-0310222123103303-2002032320320210-3211232033310102-3032120133333111-3213323112202231-0000133003113031-1030333103133032-2321311122323322)
- [Examples](../guides/data-sources--tunnel--examples--group-001.md#canonical-3300330231121311-2222032231103133-0121230003001020-2233233231213002-1010021013211133-0322111310201012-0020010233021132-1133232313203310)
