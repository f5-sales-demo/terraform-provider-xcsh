---
page_title: "xcsh_segment"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment."
---

# xcsh_segment

<a id="canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

<a id="canonical-0123201331031031-1210023021033313-0310303031121003-3101002032300130-0123032213221121-2230331003110133-2032131220020312-0110313222331233"></a>

### Prerequisites for `xcsh_segment`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1130122332330212-2203313120221023-1233300001333033-1210312002023332-2023021213201112-3101010011310013-3032100233220120-3033313310132212"></a>

### Minimal configuration for `xcsh_segment`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

<a id="canonical-0312223312130121-1032001333221303-1111301133330230-0133123131203313-1013021130133023-3312130201011003-3202103032210332-2220221100030101"></a>

### Root configuration for `xcsh_segment`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0332321323021233-2002233120002031-2211323210331112-3032013103221322-3032012332311023-0100313020002321-3332203133211201-1002012030303132"></a>

### Explore this collection for `xcsh_segment`

- [Property reference](../guides/resources--segment--reference--group-001.md#canonical-0333310231313321-0120303030320002-3001033000112301-0230202002331312-0213020300113131-0011002021012330-3202203300010331-3312231333130010)
- [Examples](../guides/resources--segment--examples--group-001.md#canonical-0010012020032131-0303310301013201-0301311230200021-1232232112301310-2303121301123012-3012332033211130-2330330030312333-0313210121303333)
- [Import](../guides/resources--segment--lifecycle--group-001.md#canonical-1001002102021311-1203023103330002-2021312223200031-3213233012332321-2031323020033021-2230320332123133-0002031022101301-2032211323110003)
- [Timeouts](../guides/resources--segment--lifecycle--group-001.md#canonical-0202030120332322-0103222002111101-3123322313103321-0102031123023031-2102001323002030-2331311110220013-0222030302113113-2223010002210213)
