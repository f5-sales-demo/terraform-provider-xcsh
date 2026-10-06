---
page_title: "xcsh_enhanced_firewall_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy."
---

# xcsh_enhanced_firewall_policy

<a id="canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_enhanced_firewall_policy

Manages an Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy
specification. configuration.

<a id="canonical-0321003301132132-0013110203021102-1223110222123030-2203020123322210-0201320210220002-1333311003303110-2120102123022123-2232030010302300"></a>

### Prerequisites for `xcsh_enhanced_firewall_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1012130223132300-1213031022302103-1121313102220201-1210003303112101-0131311200002303-2123031210010312-3300233300033321-2331331103210133"></a>

### Minimal configuration for `xcsh_enhanced_firewall_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# EnhancedFirewallPolicy Resource Example
# Manages a Enhanced Firewall Policy resource in F5 Distributed Cloud for enhanced firewall policy specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic EnhancedFirewallPolicy configuration
resource "xcsh_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "staging"
}
```

<a id="canonical-2103303013233103-2302231233210123-2310003303232120-3022320233121323-3121012201002302-1233102132133112-3223233133302213-3103131300221220"></a>

### Root configuration for `xcsh_enhanced_firewall_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1322213300011313-0003331310010033-2012331120232022-3130132230320033-1303321110103101-0130212202322100-3302331120113013-3021122200013020"></a>

### Explore this collection for `xcsh_enhanced_firewall_policy`

- [Property reference](../guides/resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [Examples](../guides/resources--enhanced_firewall_policy--examples--group-001.md#canonical-0300121330133313-2033001313033321-1102113313303112-3122233131133213-3012030223103201-1312010030303003-3313002333002221-0002221303031112)
- [Import](../guides/resources--enhanced_firewall_policy--lifecycle--group-001.md#canonical-2221333202332223-3112120322020302-3333133330010331-0101031033323203-3213323110231223-1232301211322232-0223233122212010-1322311123303200)
- [Timeouts](../guides/resources--enhanced_firewall_policy--lifecycle--group-001.md#canonical-2012131022321103-3221101033202120-0212002002103111-1022312030110320-3233033120100013-0223011130002023-2132302330220001-0013313110103220)
