---
page_title: "xcsh_network_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy."
---

# xcsh_network_policy

<a id="canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_policy

Reads Network Policy information from F5 Distributed Cloud.

<a id="canonical-3320123132302311-2111213221231110-1021121233301232-0330222332222331-2012322012231123-2223130020013003-2200120332022211-2033212203021210"></a>

### Prerequisites for `xcsh_network_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3212301123132011-2032303033201222-3013010331010000-0313013112333123-2222010231030131-2301112033321013-1030231013123110-0332031332000122"></a>

### Minimal configuration for `xcsh_network_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicy by name
data "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}

output "network_policy_id" {
  value = data.xcsh_network_policy.example.id
}
```

<a id="canonical-3011230320013312-1012022000303222-1301033032000213-0221322023312003-3333312032121222-3003101222030200-3322312121023202-2313231001011011"></a>

### Root configuration for `xcsh_network_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3200223232021311-0002102020002102-3031211203201111-1302211203122212-1303123013223121-3211201212222333-1200301001002131-1133031203332320"></a>

### Explore this collection for `xcsh_network_policy`

- [Property reference](../guides/data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [Examples](../guides/data-sources--network_policy--examples--group-001.md#canonical-3101310111230003-2031330022332131-3120112210211023-0111230213001020-2002212100121322-2120101311133223-3111333310021231-2312101332310311)
