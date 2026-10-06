---
page_title: "xcsh_crl"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl."
---

# xcsh_crl

<a id="canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_crl

Reads CRL information from F5 Distributed Cloud.

<a id="canonical-1031202131133002-3113131012111332-1321223010223312-0223111021312120-2233110232232302-1231231012302313-0111003210012302-0010132221102010"></a>

### Prerequisites for `xcsh_crl`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1132303101302230-1021211012111030-0120202133301230-0123023030300111-1232002003222202-0301030033220013-2201012331230220-3030210331233301"></a>

### Minimal configuration for `xcsh_crl`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

<a id="canonical-3103023111120032-3102230130132110-2310020320010010-0003121132123000-1023133113331303-0201033032032322-1101001200033210-0221332233132321"></a>

### Root configuration for `xcsh_crl`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2111100102030330-3032332211220231-2302122300212230-0111201032133111-3122113202311310-0200301033102102-3033103201323313-2113012201101002"></a>

### Explore this collection for `xcsh_crl`

- [Property reference](../guides/data-sources--crl--reference--group-001.md#canonical-3002300001121302-2031123030132100-2121300311223202-1012300002331221-2202200233021133-2322302101223023-1223032000220320-2322010130131013)
- [Examples](../guides/data-sources--crl--examples--group-001.md#canonical-2303331213223002-2112213123320220-3331203230311231-2101022333102302-0112033003011221-1031110223130221-0010123102223023-3303312211123230)
