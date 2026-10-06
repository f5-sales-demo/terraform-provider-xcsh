---
page_title: "xcsh_crl"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl."
---

# xcsh_crl

<a id="canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

<a id="canonical-2022313011021003-1000303110032101-1132233321001122-3032000010322103-1010123222103033-1121330031222223-2203233003102223-1121301231222111"></a>

### Prerequisites for `xcsh_crl`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1301203302013022-2120012013223232-1132213000010323-1322021002103032-3120223100232112-1221231313132100-2320133330133223-1210003001013221"></a>

### Minimal configuration for `xcsh_crl`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

<a id="canonical-0322023221021003-1330033333303100-1313013310300032-3002212212113010-0003023331300010-3020231020322220-0221113220100323-2032132333022223"></a>

### Root configuration for `xcsh_crl`

Required root properties: `name`, `namespace`, `refresh_interval`, `server_address`, `server_port`, `timeout`. Full root flags and choices appear in the property reference.

<a id="canonical-0221133130230332-2000312320320011-1212033123131302-1322300230001211-1301130102033111-1211001331132303-3311012101110002-3331023232323220"></a>

### Explore this collection for `xcsh_crl`

- [Property reference](../guides/resources--crl--reference--group-001.md#canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300)
- [Examples](../guides/resources--crl--examples--group-001.md#canonical-3310232032200130-0113023102123231-0210011231122312-1033220100003200-0021120321310011-2301211233010312-1100231101320111-2312112221212132)
- [Import](../guides/resources--crl--lifecycle--group-001.md#canonical-3311000203101232-2031203322230023-1030231121212302-1002211312332133-0101301200101300-0011223301313212-1032313213320100-1103210210212030)
- [Timeouts](../guides/resources--crl--lifecycle--group-001.md#canonical-3012021231000130-0022121322021330-2011232313312113-1331011310310123-1310001300010133-0001110031121222-2232202011122331-0310112201202021)
