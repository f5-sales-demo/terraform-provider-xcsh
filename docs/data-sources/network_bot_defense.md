---
page_title: "xcsh_network_bot_defense"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_bot_defense."
---

# xcsh_network_bot_defense

<a id="canonical-0033010203002120-1333301320122310-3023133010230310-1013120133033132-3201111033200121-2113002210321222-0020023233301212-2222022300321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_bot_defense

Bot Defense domains for an FQDN-aware firewall or proxy. Values are bundled from the pinned OpenAPI
release; this data source performs no network request. Ports and traffic direction are not encoded
in the manifest.

<a id="canonical-1312303210201330-0010121321321323-2223312012101200-0110113202030020-3121103201020023-2120132323322300-1132111023320331-0031230130101133"></a>

### Prerequisites for `xcsh_network_bot_defense`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3202220200031323-1103032120300102-2001011303022110-3200030000300303-2331323212330200-3222213021012011-1201100301110123-2331221310132122"></a>

### Minimal configuration for `xcsh_network_bot_defense`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```

<a id="canonical-2031002230333221-3211032331133000-0300233030103331-1122231322312130-0232210322112223-0322112011231111-2333103213100333-3012330223130201"></a>

### Root configuration for `xcsh_network_bot_defense`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-0331232221201322-0033022203311332-0301122310000210-1202300333132330-2322013102000101-2230302301011322-2131030122031033-3111100333023102"></a>

### Explore this collection for `xcsh_network_bot_defense`

- [Property reference](../guides/data-sources--network_bot_defense--reference--group-001.md#canonical-3221232120212230-2113110001332110-3020030310200102-2221330102200132-0233033133013112-0223001122321012-2032212130020112-2011211313321001)
- [Examples](../guides/data-sources--network_bot_defense--examples--group-001.md#canonical-0002130000312332-1110233230003121-2103233333033223-2312121023022123-2020010323101203-1010300303231210-0003103302003303-2112021022311230)
