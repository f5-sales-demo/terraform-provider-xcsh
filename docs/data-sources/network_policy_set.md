---
page_title: "xcsh_network_policy_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_set."
---

# xcsh_network_policy_set

<a id="canonical-3022111020001102-2220213113331210-1332233320322003-0003223303012213-1120233210001232-0202220312332012-2312220103333300-0203313111321331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_policy_set

Reads an existing Network Policy Set in the requested namespace.

<a id="canonical-1011113323113201-3112032302330122-3121020013301121-2210013122312022-1322332113000333-3022012132121003-3321020332212010-1131101132330133"></a>

### Prerequisites for `xcsh_network_policy_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2002030201333133-2212330221322020-3031033212223331-1003101003232203-0020001122203321-1301013111012011-3030100222111313-1022002303021201"></a>

### Minimal configuration for `xcsh_network_policy_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```

<a id="canonical-0113230031323332-3001110230033322-2132201101213033-3231333001113003-1223032013202211-2030222311310111-0100223332103303-3233311012313012"></a>

### Root configuration for `xcsh_network_policy_set`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0103101203310201-2000301032233210-0212220130103322-3132001212112123-3010311331322121-1131133223022033-0302102301010310-2233031330033131"></a>

### Explore this collection for `xcsh_network_policy_set`

- [Property reference](../guides/data-sources--network_policy_set--reference--group-001.md#canonical-3202333230333132-0113330223032322-2031313232020132-3230321223333031-0123100330113200-2131003230131301-1122203022233100-1003132220303022)
- [Examples](../guides/data-sources--network_policy_set--examples--group-001.md#canonical-1232113121123000-3300302111221013-1210121120012010-3032211003322111-3112000330100221-2332033101132211-0232032222131202-2010201203203112)
