---
page_title: "xcsh_protected_domain"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain."
---

# xcsh_protected_domain

<a id="canonical-0332111003332231-3231333320302220-0013131312310303-3021021201222321-0302322122212110-0022133313313221-1010303303020331-2102022101100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protected_domain

Reads Protected Domain information from F5 Distributed Cloud.

<a id="canonical-0230023120333310-2021222020333111-1210220330230021-2012110320121012-1230123102301202-3322333011121313-2210311011222120-1113310233103003"></a>

### Prerequisites for `xcsh_protected_domain`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0132001100121122-2232230303313322-1333322212113011-0233021322330020-1012022013033211-1312121313320121-0030122203010232-2010313030320323"></a>

### Minimal configuration for `xcsh_protected_domain`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
```

<a id="canonical-3012211222032102-3201123000210002-0221023321100010-0120130202020011-2001033300313230-3321311010213332-3223021011323202-0330132123323311"></a>

### Root configuration for `xcsh_protected_domain`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1331321000321022-1330031302131020-3310030310032102-0003001003122021-0230213313020233-3012112302032213-3322012232111233-1302320133220012"></a>

### Explore this collection for `xcsh_protected_domain`

- [Property reference](../guides/data-sources--protected_domain--reference--group-001.md#canonical-0100212211301123-0131201333311010-0332012111101032-1203123003220200-0132003223232101-2021312112301101-0321113310100202-2032033113330123)
- [Examples](../guides/data-sources--protected_domain--examples--group-001.md#canonical-3030223231010012-3322000022222331-0212010302300020-3010113221023011-3121313131011111-0203010312301030-0122222022013100-1201023111230012)
