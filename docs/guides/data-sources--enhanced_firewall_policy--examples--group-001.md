---
page_title: "xcsh_enhanced_firewall_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy examples."
---

# xcsh_enhanced_firewall_policy examples

<a id="canonical-1300302000322233-2111321030302220-3121311302320112-2002001221120313-1021122233212310-2103002013113012-1001001012133113-0222002222302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- Examples

<a id="canonical-1303202222120000-3323012020312132-0331202032121120-0231002103220020-3133103020121113-0000131122313302-3232020030202201-2100332322002201"></a>

### Complete configurations for `xcsh_enhanced_firewall_policy`

- [Data source](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-3212301120123002-2120002222312122-0030221131010231-2202311220223233-2003202123012223-3303322102000200-3203311323100031-2131223212103132): valid configuration.

<a id="canonical-3212301120123002-2120002222312122-0030221131010231-2202311220223233-2003202123012223-3303322102000200-3203311323100031-2131223212103132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md#canonical-1311231100333113-0230113312212103-1202110301210321-3202021333320231-3003331110212213-2102231103013021-0332332202022123-1300223100313200)
- [Examples](data-sources--enhanced_firewall_policy--examples--group-001.md#canonical-1300302000322233-2111321030302220-3121311302320112-2002001221120313-1021122233212310-2103002013113012-1001001012133113-0222002222302222)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_enhanced_firewall_policy/data-source.tf`; digest `sha256:4fad9180d1e11ef8fde0a09a6c59611ed259b2e0dc9e6684bd5a7a00db211172`.

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
