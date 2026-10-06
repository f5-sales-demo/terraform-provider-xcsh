---
page_title: "xcsh_registration_approval"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration_approval."
---

# xcsh_registration_approval

<a id="canonical-1331103323121310-0002130210000313-2020131122020200-0233121113323131-3323011213331230-0033000123022130-3131333113210323-3013000301112322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_registration_approval

Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.
configuration.

<a id="canonical-0023300130322000-1230323323202130-1000321031231323-1001100210212020-1330313202132330-2200212313021033-1020113110112230-2111230100300002"></a>

### Prerequisites for `xcsh_registration_approval`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1312211201003030-2210303212331010-1130332201001030-1331113120130120-1202303031202001-1001011112310201-2230321122231132-0030111320201230"></a>

### Minimal configuration for `xcsh_registration_approval`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RegistrationApproval Resource Example
# Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RegistrationApproval configuration
resource "xcsh_registration_approval" "example" {
  name      = "example-registration-approval"
  namespace = "staging"

  cluster_size = 1
}
```

<a id="canonical-3023330000110301-0132112322012002-0203111102031003-3202032231312131-2203130032131202-1331000202111231-0330333320202122-2023102312202311"></a>

### Root configuration for `xcsh_registration_approval`

Required root properties: `cluster_size`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3323203000121122-0332012112212213-3203020303312232-2332330312121321-0000201112023211-1230220021123010-3011011203201032-3120300010332012"></a>

### Explore this collection for `xcsh_registration_approval`

- [Property reference](../guides/resources--registration_approval--reference--group-001.md#canonical-0221110020100131-3031023122311011-2222002100222120-2132023011212021-1212033310020221-0302211021123030-3031320113201232-1302122130102300)
- [Examples](../guides/resources--registration_approval--examples--group-001.md#canonical-0120212003331230-0231120030031133-2330122110012231-1321111321200023-2022120032210303-3130002001000100-2300132013030121-2023021330011321)
- [Import](../guides/resources--registration_approval--lifecycle--group-001.md#canonical-2312110032322001-1213023023130212-2003013021212001-3113300022210023-3121221302330121-1211133310201312-2020312303011231-0120033200220303)
