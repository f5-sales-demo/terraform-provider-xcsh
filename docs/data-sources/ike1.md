---
page_title: "xcsh_ike1"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1."
---

# xcsh_ike1

<a id="canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_ike1

Reads Ike1 information from F5 Distributed Cloud.

<a id="canonical-0222203300201323-0313232230200033-3012311132133110-0000133232032203-3002211031102133-1223233031233030-3333222300220221-0211211112220300"></a>

### Prerequisites for `xcsh_ike1`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0331100033322312-1003212330223233-0331120321323333-0003320123231031-2031023013001031-3132220102022011-3110222002212202-1201223210223003"></a>

### Minimal configuration for `xcsh_ike1`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```

<a id="canonical-1213010323130112-0113201313012211-2031012303223103-0232112000210102-0033031122313301-1002321032332210-2323131223300321-0011312300120230"></a>

### Root configuration for `xcsh_ike1`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1231130103031120-1310022112002231-0022323331223032-2020131022201112-1023011103003221-3323210321303011-2123000320210111-1013110231330302"></a>

### Explore this collection for `xcsh_ike1`

- [Property reference](../guides/data-sources--ike1--reference--group-001.md#canonical-2113121021131213-0322330121000333-0232103301210212-3101112300210032-3102213101322310-0311031112112121-1311322002110322-0313010312330013)
- [Examples](../guides/data-sources--ike1--examples--group-001.md#canonical-3132332313110102-2301103110101323-3013231113332213-2101031223111032-3332331103131323-1221301031101331-1200200320022232-2002033021333210)
