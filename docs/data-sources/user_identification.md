---
page_title: "xcsh_user_identification"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification."
---

# xcsh_user_identification

<a id="canonical-2201200133133200-1323022002031211-0333333233002132-2202230300203100-0300223221213321-3023000331000320-1020113111130300-2213210002320321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_user_identification

Reads User Identification information from F5 Distributed Cloud.

<a id="canonical-1011111102121321-2013033211231312-0330023032221003-2213033301212120-2003132010122003-0011020111221120-0322130113310121-3012303031123021"></a>

### Prerequisites for `xcsh_user_identification`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3112320113331221-0310032010311132-3033203113331333-1303021213303230-0030232200111033-0322020322212333-3221011331012001-0212101332312301"></a>

### Minimal configuration for `xcsh_user_identification`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UserIdentification by name
data "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}

output "user_identification_id" {
  value = data.xcsh_user_identification.example.id
}
```

<a id="canonical-2211331100232233-1100310112130101-0332003222210201-1103332233130032-1200133131132300-0300202313331320-0312102111021102-0311130313021331"></a>

### Root configuration for `xcsh_user_identification`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1133021332021323-2233210103210030-2122133310312300-2331101211112311-2110310211330310-2132110112200100-3110101013213103-3111213220313110"></a>

### Explore this collection for `xcsh_user_identification`

- [Property reference](../guides/data-sources--user_identification--reference--group-001.md#canonical-3310121033102132-2303003001031311-1133201000013200-3221030012100122-3312001102312332-2301122032131033-0221223313103222-0332010121330113)
- [Examples](../guides/data-sources--user_identification--examples--group-001.md#canonical-2332023102210113-1231211101202322-2303231032012131-0320023202113003-3233222222302313-1020333202322011-0000013021301000-1003303030200123)
