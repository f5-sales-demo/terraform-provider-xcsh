---
page_title: "xcsh_oidc_oauth_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery landing."
---

# xcsh_oidc_oauth_discovery landing

<a id="canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203202012213021-2203002003312120-3233022312123000-3333003210332113-0102031233332323-0223201211110011-1300312002002022-2303222023232020"></a>

## xcsh_oidc_oauth_discovery — xcsh_oidc_oauth_discovery / 112121313112 / 2

Breadcrumbs:

- xcsh_oidc_oauth_discovery

Resource creation operation.

<a id="canonical-1323113130312330-0311023221232022-0300032002033212-1002010030110321-3031310313102221-0122233023311231-3023010102112202-0112300332012210"></a>

## Prerequisites — xcsh_oidc_oauth_discovery / 112121313112 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1022112300223301-0320122030310002-3012302122210130-1323333011221130-0012112231220011-1201313223203121-1012333130022132-0013312002030330"></a>

## Minimal configuration — xcsh_oidc_oauth_discovery / 112121313112 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OIDCOauthDiscovery DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_oidc_oauth_discovery" "example" {
  namespace = "example-value"
}

output "oidc_oauth_discovery_result" {
  value = data.xcsh_oidc_oauth_discovery.example
}
```

<a id="canonical-1230003120231120-0303210331110200-0103203103010001-2211022020112212-1311002300331130-0010120202311011-3312302330333101-1301122102221300"></a>

## Root configuration — xcsh_oidc_oauth_discovery / 112121313112 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2012023313131203-2111133300011003-0321010100300313-2023211333023022-0220202312333030-3020012102313100-0032023032130232-3302221323013131"></a>

## Next pages — xcsh_oidc_oauth_discovery / 112121313112 / 6

- [Property reference](../guides/data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-2000200023102030-2300333232322203-1113221101011222-0232203102332013-0201033112023021-3021031133133122-2122002130231000-3102322223332022)
- [Examples](../guides/data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-3101302130120123-0223201112000011-3300033320013211-0000131221321003-1320311131121223-0001003132213223-1313133333323201-0121321211022220)
