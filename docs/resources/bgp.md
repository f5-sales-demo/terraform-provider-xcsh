---
page_title: "xcsh_bgp"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp."
---

# xcsh_bgp

<a id="canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

<a id="canonical-1211330030112230-1231122203201103-1012122300002311-0122113022101232-1001021031031331-1011233321233311-0133030200111301-3200332231313323"></a>

### Prerequisites for `xcsh_bgp`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0030202202110321-2033303110313112-0132210300303102-2301121212302030-1212231313333023-0220110303310233-3302332112133122-3013322111331103"></a>

### Minimal configuration for `xcsh_bgp`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

<a id="canonical-3231020001131212-1132302333302031-2023203200300123-2302001313031332-3123013111233133-2130231131203330-2322022113023230-0013123213113112"></a>

### Root configuration for `xcsh_bgp`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1213223323012233-3030300301123133-2322210331103132-3110313301203120-1200011012303212-0023033100122312-1133021021113213-2101223220033233"></a>

### Explore this collection for `xcsh_bgp`

- [Property reference](../guides/resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [Examples](../guides/resources--bgp--examples--group-001.md#canonical-0023113332102111-1003103310223221-0311022030221000-2202300320022011-1203200001201133-2111200033013301-1230112300130222-1331122202233022)
- [Import](../guides/resources--bgp--lifecycle--group-001.md#canonical-0100132031003210-2120031120132100-3313102103301211-1202001221130302-0312112022100313-1001222310002020-0031122332111031-1210300301223030)
- [Timeouts](../guides/resources--bgp--lifecycle--group-001.md#canonical-3033003321102320-3300233211111010-1012110202313213-1331003010312013-2222110103031002-3210202002100022-2210213012322232-1011311221133302)
