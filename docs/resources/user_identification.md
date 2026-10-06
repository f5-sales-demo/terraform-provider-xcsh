---
page_title: "xcsh_user_identification"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification."
---

# xcsh_user_identification

<a id="canonical-0201231202111123-2212302023331001-3101130201301022-3011301201311121-1111130303322010-0031212223102323-3302121201312013-2312300311320112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_user_identification

Manages user\_identification creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-2300301231003002-2031001101013002-0101221111022010-0212101031121310-2222322003301000-1201021030221300-3102231120212110-0211011221013222"></a>

### Prerequisites for `xcsh_user_identification`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0023322231020012-2211003111330302-1201022233231003-2023032333212032-3033222110331103-3023311331220132-3322313120002333-1212313011202332"></a>

### Minimal configuration for `xcsh_user_identification`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

<a id="canonical-1300012121232032-0202223113310323-1033300011122302-1103010033012110-3103320030121201-2201030203332113-1220310333023313-3233321011310302"></a>

### Root configuration for `xcsh_user_identification`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1012230220022321-2033130030231211-0300313130002022-0002231113113033-3323231332300003-2333022213233323-0200333303331100-3330321320211322"></a>

### Explore this collection for `xcsh_user_identification`

- [Property reference](../guides/resources--user_identification--reference--group-001.md#canonical-0302203110230111-0101003133213200-0113030211023221-0210202013332211-1333132032000320-3231230023111133-2312332203222120-3011231133232311)
- [Examples](../guides/resources--user_identification--examples--group-001.md#canonical-1132012213120121-0213210121131211-3213202000202211-1310002332322212-0300103222133000-0303110110313011-2011101000122021-2233010301011232)
- [Import](../guides/resources--user_identification--lifecycle--group-001.md#canonical-0100101021131001-3320300133032202-1010122102032110-1002122121202103-1310010102233133-3000200232000333-3110030122330012-2300300223331230)
- [Timeouts](../guides/resources--user_identification--lifecycle--group-001.md#canonical-0200103110221213-0302102001110303-0202111202202202-3101310133020000-2122123020112312-0130113332312003-1123110121212323-2332330123202310)
