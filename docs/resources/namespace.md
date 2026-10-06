---
page_title: "xcsh_namespace"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace."
---

# xcsh_namespace

<a id="canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the namespace in F5 Distributed Cloud.

<a id="canonical-2231221230203010-1311333333322110-0032130330230110-1330111132132021-1130132311121323-1111032133312313-3313012200211223-3322010123323013"></a>

### Prerequisites for `xcsh_namespace`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0100100113332201-1223313000112230-3212011230221312-1020032111201331-1131311101330002-3000232121220113-3331222123113222-3113311003311023"></a>

### Minimal configuration for `xcsh_namespace`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```

<a id="canonical-2310120113110331-0201102330023332-3001012112010211-2132003303023222-3102010003302211-3301031212131200-0322220023310013-1200030210132311"></a>

### Root configuration for `xcsh_namespace`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3331211021320322-3030031312321122-2300010021211213-0103002023201222-0302302320131021-2110132333302123-3012313323302122-3022123023333223"></a>

### Explore this collection for `xcsh_namespace`

- [Property reference](../guides/resources--namespace--reference--group-001.md#canonical-0021301131013303-1333120321233220-3113022223110333-0100200232103001-3111031002202300-3231113102202033-3022100133111131-1310313032202230)
- [Examples](../guides/resources--namespace--examples--group-001.md#canonical-0123202330031033-0303010003032300-1110000210011033-0332010112301132-1111030222330030-0220310211111100-2310320200233000-2302132023320230)
- [Import](../guides/resources--namespace--lifecycle--group-001.md#canonical-0320000300030320-2021002123203112-0303313311033303-1233033211003313-1002122203320111-2330321121202110-2211110330123011-3002101211000321)
- [Timeouts](../guides/resources--namespace--lifecycle--group-001.md#canonical-0223221203320003-2000300231233123-3232211333113211-1112120132223333-1223220000033010-3221002110303210-0103320300120112-0222300021311101)
