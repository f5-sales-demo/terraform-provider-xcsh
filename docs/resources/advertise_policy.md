---
page_title: "xcsh_advertise_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy."
---

# xcsh_advertise_policy

<a id="canonical-0001000222330023-3030100013232123-3330120310010230-2130230202100101-2211301023201212-0322203103300021-2231103201203132-3101333221320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_advertise_policy

Manages an Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

<a id="canonical-3032303211231003-1033211123332312-1023123033001332-1121223313120313-1122122012121321-0331000233211002-1122233002123102-1201302121011233"></a>

### Prerequisites for `xcsh_advertise_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1030300330020323-2132221323221300-1222112102320200-0131331320302000-2111033221312322-2120313331023022-1231333010220232-0211323103232133"></a>

### Minimal configuration for `xcsh_advertise_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Resource Example
# Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AdvertisePolicy configuration
resource "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}
```

<a id="canonical-1313002113031111-2122232230101222-3332332003320033-1102223201322200-0320003331300303-3310300130313301-2303212122211013-0320013233103200"></a>

### Root configuration for `xcsh_advertise_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1213033300201312-2310332111211031-2203301322313222-0322230313321232-1200132221110010-0011101131120232-0030202023230301-0303201321330331"></a>

### Explore this collection for `xcsh_advertise_policy`

- [Property reference](../guides/resources--advertise_policy--reference--group-001.md#canonical-2311303001023102-3233233231231222-1130010332020001-3323201312312121-1032210233312021-1302023123223011-0023010011201300-2132121100111331)
- [Examples](../guides/resources--advertise_policy--examples--group-001.md#canonical-1222032302121303-1310301213122232-0312232232033120-3310223121230221-0301130033022321-0320133021320102-3021001333220022-1013311233113223)
- [Import](../guides/resources--advertise_policy--lifecycle--group-001.md#canonical-3220131202323231-2120121230133121-2320012030002312-0100000301233001-3023032020101302-0331332113212010-3212020332233100-1113302202011131)
- [Timeouts](../guides/resources--advertise_policy--lifecycle--group-001.md#canonical-1301203200233000-0211103303021221-3303210110033232-0010021033130333-1012322222211121-1000000300312030-3020323232333013-0100000233321220)
