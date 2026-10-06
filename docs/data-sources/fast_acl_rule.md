---
page_title: "xcsh_fast_acl_rule"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule."
---

# xcsh_fast_acl_rule

<a id="canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_fast_acl_rule

Reads a Fast ACL rule with source IP and port matching and an associated action.

<a id="canonical-3003102221322103-1301202021333011-0230110020031113-3113300200211123-3321203103133120-1013130232023232-2201012010301102-3001333212223213"></a>

### Prerequisites for `xcsh_fast_acl_rule`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2322010131311222-0121221312210132-3132320101110201-3131232113101331-0220113112012330-1312120222021220-3102322202131013-1002120211130322"></a>

### Minimal configuration for `xcsh_fast_acl_rule`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```

<a id="canonical-3201022321330102-1031121321013320-3210320311132222-1132332332302030-3121321311232203-2002120012332101-0321103332033200-2332111221200110"></a>

### Root configuration for `xcsh_fast_acl_rule`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0212113122101310-3010223203121001-0133002021010332-1333103322033303-0031000321313312-1020203131212203-1332313032000120-3101202211002221"></a>

### Explore this collection for `xcsh_fast_acl_rule`

- [Property reference](../guides/data-sources--fast_acl_rule--reference--group-001.md#canonical-0130033020112332-1320113013031122-3323102231201130-2302010233220010-0332222213110322-3330103313001312-0313202331310301-0031300023203130)
- [Examples](../guides/data-sources--fast_acl_rule--examples--group-001.md#canonical-3211232130101212-1130233301203323-3130113131323113-0011112033203332-2221012002203030-2101332011200103-0103233112100130-3302123121230230)
