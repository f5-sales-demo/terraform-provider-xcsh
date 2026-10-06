---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy."
---

# xcsh_enhanced_firewall_policy

<a id="canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Reads Enhanced Firewall Policy information from F5 Distributed Cloud.

<a id="canonical-1100201213220203-1303333013101031-2113332230211212-3230122011221213-1032332331303301-0120130013110300-1331311010330001-3200022103012110"></a>

### Prerequisites for `xcsh_enhanced_firewall_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1002330000003031-1202021221331321-2112121033032223-0303231100202101-3322331032320313-3233302322321220-3133023212233210-2100311200012020"></a>

### Minimal configuration for `xcsh_enhanced_firewall_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# EnhancedFirewallPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing EnhancedFirewallPolicy by name
data "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}

output "enhanced_firewall_policy_id" {
  value = data.xcsh_enhanced_firewall_policy.example.id
}
```

<a id="canonical-3233003310302001-3330231221210223-1022331330331131-2030100113112000-2320201122303031-2121012312132003-3320311033310303-0123001130320120"></a>

### Root configuration for `xcsh_enhanced_firewall_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2113013000311000-0033231112031323-1301002310132001-0123200233223312-0000203303210011-1112012331200010-0132322113323130-2110333122223333"></a>

### Explore this collection for `xcsh_enhanced_firewall_policy`

- [Property reference](../guides/data-sources--enhanced_firewall_policy--reference--group-001.md#canonical-1212330032022110-0300203031202310-3023323111320103-2002220303320001-1002031021230030-3103301111003102-1131210302121020-3222312110012110)
- [Examples](../guides/data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-1300302000322233-2111321030302220-3121311302320112-2002001221120313-1021122233212310-2103002013113012-1001001012133113-0222002222302222)
