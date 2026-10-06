---
page_title: "xcsh_enhanced_firewall_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy examples."
---

# xcsh_enhanced_firewall_policy examples

<a id="canonical-0300121330133313-2033001313033321-1102113313303112-3122233131133213-3012030223103201-1312010030303003-3313002333002221-0002221303031112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- Examples

<a id="canonical-3001131232020021-3003333233132203-1331023033223023-2011021330102301-2202330202321101-0010212132221210-0020203302311131-0112030033002133"></a>

### Complete configurations for `xcsh_enhanced_firewall_policy`

- [Resource](resources--enhanced_firewall_policy--examples--group-001.md#canonical-1110202103311302-3120003312311211-0300112112111000-2313110111313033-2330220202320211-2232211200331203-2121322111231302-3032332113231121): valid configuration.

<a id="canonical-1110202103311302-3120003312311211-0300112112111000-2313110111313033-2330220202320211-2232211200331203-2121322111231302-3032332113231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Examples](resources--enhanced_firewall_policy--examples--group-001.md#canonical-0300121330133313-2033001313033321-1102113313303112-3122233131133213-3012030223103201-1312010030303003-3313002333002221-0002221303031112)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_enhanced_firewall_policy/resource.tf`; digest `sha256:de754ad7eba55318affca16381d010b7282e358ea583d9fad1b36b6208539006`.

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
