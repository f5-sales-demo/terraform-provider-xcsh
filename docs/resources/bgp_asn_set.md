---
page_title: "xcsh_bgp_asn_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set."
---

# xcsh_bgp_asn_set

<a id="canonical-3010010321231122-0133331023021212-0330320220102302-0102122320130321-1110011002112302-3010211202300010-0031103320123032-3230133030121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_bgp_asn_set

Manages bgp\_asn\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-0313313201130223-1113311212210111-1123332120112101-1322021212232111-0000333232201311-2023200301121003-1302020122020121-2020320223110230"></a>

### Prerequisites for `xcsh_bgp_asn_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2212222210201221-0010301022212300-0333031230220213-0231231011220303-1120110131011333-1200200011002110-1320312123013030-1310031211203113"></a>

### Minimal configuration for `xcsh_bgp_asn_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Resource Example
# Manages bgp_asn_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPAsnSet configuration
resource "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"

  as_numbers = [1]
}
```

<a id="canonical-0311010310112213-0312031332012203-3200122230321310-1211003030312121-3012133012020330-1022020323202200-1210312002113211-0230002201032101"></a>

### Root configuration for `xcsh_bgp_asn_set`

Required root properties: `as_numbers`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3032322210300033-3221333221133323-0120023010132323-1130011130332122-2330022301313223-0200123221021231-3123010300331122-3221203303202011"></a>

### Explore this collection for `xcsh_bgp_asn_set`

- [Property reference](../guides/resources--bgp_asn_set--reference--group-001.md#canonical-2221001120333021-0023132212300332-0300031102333013-0012022101331031-3232222323021231-0110023232233332-0120221100112302-0102233011221200)
- [Examples](../guides/resources--bgp_asn_set--examples--group-001.md#canonical-1311312101312312-0003113031020322-3000212100121303-1310231031111133-0303022310101211-1323001112323220-1210220310300300-1033220312010023)
- [Import](../guides/resources--bgp_asn_set--lifecycle--group-001.md#canonical-1010101200311130-1022032330233022-1101101220110023-1033122000033231-3222131220110102-0000203222212233-0033113220302332-3130102330112100)
- [Timeouts](../guides/resources--bgp_asn_set--lifecycle--group-001.md#canonical-2231101132322311-0031203302331011-1103202230200221-2020032032103322-0002203010133332-0322333211232230-0022233031100312-2333111302123033)
