---
page_title: "xcsh_virtual_host"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host."
---

# xcsh_virtual_host

<a id="canonical-3003121122323110-3033220033233331-3332132022313023-0311020131102101-3230333213220101-3221133313021313-2331201132110012-2011123023323220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_virtual_host

Reads Virtual Host information from F5 Distributed Cloud.

<a id="canonical-2122230022331011-1003232330101211-1300231320003323-1133123210003130-1221200123200301-0113110020223101-0310233012111011-2121103313031133"></a>

### Prerequisites for `xcsh_virtual_host`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3101100330222212-1102220331202023-2211121103312310-0132021100200332-2123302031002132-3323303102233110-2213212313211001-1333300030132113"></a>

### Minimal configuration for `xcsh_virtual_host`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualHost Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualHost by name
data "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"
}

output "virtual_host_id" {
  value = data.xcsh_virtual_host.example.id
}
```

<a id="canonical-1011011101002332-2303220002002231-1023303113132123-0331122121330021-0220010302231210-0011301011031302-0302120220223302-3323231103332231"></a>

### Root configuration for `xcsh_virtual_host`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2200233023331210-3122232120110233-0103330112011022-2120211133030210-0133012021321323-1220201323221003-2310230132021030-0300322232113230"></a>

### Explore this collection for `xcsh_virtual_host`

- [Property reference](../guides/data-sources--virtual_host--reference--group-001.md#canonical-1331002312103212-0111200321323210-2012213313103303-2203323023011130-2212023000123013-2012220022323101-1032110322132213-0110330311221131)
- [Examples](../guides/data-sources--virtual_host--examples--group-001.md#canonical-0100001331112310-1230100011301031-2230200220313033-0002101033032003-3322002302312013-3120010230021103-0322303302110022-2321033300201112)
