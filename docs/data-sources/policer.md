---
page_title: "xcsh_policer"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer."
---

# xcsh_policer

<a id="canonical-3330032030013021-1312303333232323-1331010000013021-1132031333330232-2302330331211311-0310120132213301-2033021113022132-2232102121103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_policer

Reads Policer information from F5 Distributed Cloud.

<a id="canonical-2202321330030212-1031111133311222-3122211110020221-0303010310031013-0122113101123133-0023222313313232-2013231130012220-3001202000133330"></a>

### Prerequisites for `xcsh_policer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2202011000213020-2003100302113122-1202201022010102-3013002002112121-0312222311010330-3330013313210103-3200222322211011-2231310110102013"></a>

### Minimal configuration for `xcsh_policer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

<a id="canonical-3312131120202220-2032130003121130-3331303233132023-3022213012311323-1022213210213133-2220122122113010-0330223211321033-3120233121133003"></a>

### Root configuration for `xcsh_policer`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3003022230331110-3220101220003133-2113330222000110-1131002312333220-0232001100213130-0012121211320113-0111321321230113-0212113031111320"></a>

### Explore this collection for `xcsh_policer`

- [Property reference](../guides/data-sources--policer--reference--group-001.md#canonical-3320132331212333-3010302313330232-0221223202012002-1222203102211013-2121323323003011-1111232130221313-0133131033222112-3002101300233013)
- [Examples](../guides/data-sources--policer--examples--group-001.md#canonical-0132331122033320-3221103230001203-3301233112320033-2130310033311020-2213223002233022-3000022212233110-3200121231300001-0013133033121331)
