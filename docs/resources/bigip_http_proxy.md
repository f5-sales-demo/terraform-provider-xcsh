---
page_title: "xcsh_bigip_http_proxy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy."
---

# xcsh_bigip_http_proxy

<a id="canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bigip_http_proxy

Manages BIG-IP HTTP Proxy in a given namespace. If one already exists, it will give an error in F5
Distributed Cloud.

<a id="canonical-2121112130220303-2321333123102201-3331020133033113-0103221231201031-2022122211230222-0003031232000033-0113232221132212-0311121312132332"></a>

### Prerequisites for `xcsh_bigip_http_proxy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1021302023102320-2032223000121021-1320220123300110-1131110222023331-0312203112120201-1320300021120122-3123322013222000-3123300323323120"></a>

### Minimal configuration for `xcsh_bigip_http_proxy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPHTTPProxy Resource Example
# Manages BIG-IP HTTP Proxy in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BigIPHTTPProxy configuration
resource "xcsh_bigip_http_proxy" "example" {
  name      = "example-bigip-http-proxy"
  namespace = "staging"
}
```

<a id="canonical-2033131021221102-2130223333002332-0212333000203322-2213231221132210-2322100103323030-0021333022300000-0223100123322213-1232331213333301"></a>

### Root configuration for `xcsh_bigip_http_proxy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3023232331011211-0103011112120333-0301221301022103-0030111101331312-3332022323010313-3011022312232222-2112213122213020-0232233001113130"></a>

### Explore this collection for `xcsh_bigip_http_proxy`

- [Property reference](../guides/resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [Examples](../guides/resources--bigip_http_proxy--examples--group-001.md#canonical-2313230102320310-0012132010302323-1000322011120020-3200011232313003-3222013112301030-1323130302012001-0002211033301020-3013021303331210)
- [Import](../guides/resources--bigip_http_proxy--lifecycle--group-001.md#canonical-1322100321131123-3311223313323113-2230001002103133-3100231312000321-0222001120321210-2322202312220231-1131001322001222-2203201230031012)
- [Timeouts](../guides/resources--bigip_http_proxy--lifecycle--group-001.md#canonical-0032021232021322-1031002130311031-0030320223103110-3032103003210012-0231113002003300-0110201030322130-2013333120332300-1323202131121211)
