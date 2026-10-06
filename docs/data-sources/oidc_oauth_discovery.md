---
page_title: "xcsh_oidc_oauth_discovery"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_oidc_oauth_discovery."
---

# xcsh_oidc_oauth_discovery

<a id="canonical-1212212133022200-1202310302030030-1123103212311111-3102001300103310-0102100100032330-3310033310100102-1213021002223010-3132011230221122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_oidc_oauth_discovery

Reads OIDC/OAuth discovery information from F5 Distributed Cloud.

<a id="canonical-2203202012213021-2203002003312120-3233022312123000-3333003210332113-0102031233332323-0223201211110011-1300312002002022-2303222023232020"></a>

### Prerequisites for `xcsh_oidc_oauth_discovery`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1323113130312330-0311023221232022-0300032002033212-1002010030110321-3031310313102221-0122233023311231-3023010102112202-0112300332012210"></a>

### Minimal configuration for `xcsh_oidc_oauth_discovery`

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

<a id="canonical-1022112300223301-0320122030310002-3012302122210130-1323333011221130-0012112231220011-1201313223203121-1012333130022132-0013312002030330"></a>

### Root configuration for `xcsh_oidc_oauth_discovery`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1230003120231120-0303210331110200-0103203103010001-2211022020112212-1311002300331130-0010120202311011-3312302330333101-1301122102221300"></a>

### Explore this collection for `xcsh_oidc_oauth_discovery`

- [Property reference](../guides/data-sources--oidc_oauth_discovery--reference--group-001.md#canonical-2000200023102030-2300333232322203-1113221101011222-0232203102332013-0201033112023021-3021031133133122-2122002130231000-3102322223332022)
- [Examples](../guides/data-sources--oidc_oauth_discovery--examples--group-001.md#canonical-3101302130120123-0223201112000011-3300033320013211-0000131221321003-1320311131121223-0001003132213223-1313133333323201-0121321211022220)
