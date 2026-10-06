---
page_title: "xcsh_proxy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy."
---

# xcsh_proxy

<a id="canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_proxy

Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.
configuration.

<a id="canonical-3001030221130200-0130303122222220-3020000332221103-2313230133301022-0102220010211110-1301220020122322-3322233111110022-2031001301002022"></a>

### Prerequisites for `xcsh_proxy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0333300320213012-1101121011102313-3210233111102103-1323332322111031-3330021130302312-3010003302200111-0033001121213322-1013300122131300"></a>

### Minimal configuration for `xcsh_proxy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Proxy Resource Example
# Manages a Proxy resource in F5 Distributed Cloud for tcp loadbalancer create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Proxy configuration
resource "xcsh_proxy" "example" {
  name      = "example-proxy"
  namespace = "staging"
}
```

<a id="canonical-0311122301032013-2320310201111302-0232012210230100-0020012011230122-3012130121121000-0223121220003320-1131102213011121-0220322010303023"></a>

### Root configuration for `xcsh_proxy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2102332200110332-2013331132332232-0213300101002301-2323202203312101-2021110100111303-0012121100312112-2011312312112102-1321022203311122"></a>

### Explore this collection for `xcsh_proxy`

- [Property reference](../guides/resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [Examples](../guides/resources--proxy--examples--group-001.md#canonical-1312000312222230-1120313213032320-3301302120322211-0333033013030102-3131013101001131-1002001223123002-1133320233001021-0211303231120120)
- [Import](../guides/resources--proxy--lifecycle--group-001.md#canonical-2233020331221110-3220321131302032-0111012300331101-3323011320113313-0213103313030023-3211320021321122-1013320201322113-2220013020312123)
- [Timeouts](../guides/resources--proxy--lifecycle--group-001.md#canonical-2113131330220223-2001220323122123-0121230010120202-3110010332021132-0003333100000123-1011010130232203-1023002230231330-2233320000121132)
