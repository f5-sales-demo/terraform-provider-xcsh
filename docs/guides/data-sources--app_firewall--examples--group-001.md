---
page_title: "xcsh_app_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall examples."
---

# xcsh_app_firewall examples

<a id="canonical-3131132132112300-2031210311300023-1320230303032133-2101121211031120-0000302131110002-0122121020210203-3302311013131221-3301311322210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- Examples

<a id="canonical-1220113303201030-0130101003120213-0213110011030200-3030210020011111-2333322123200211-2000122203202110-3110013232311302-1212203031230022"></a>

### Complete configurations for `xcsh_app_firewall`

- [Data source](data-sources--app_firewall--examples--group-001.md#canonical-0122112313333132-2212103030220033-2313030100310020-2013011031101232-1333221300331121-3312123312233031-2111002030001000-0202231311001123): valid configuration.

<a id="canonical-0122112313333132-2212103030220033-2313030100310020-2013011031101232-1333221300331121-3312123312233031-2111002030001000-0202231311001123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Examples](data-sources--app_firewall--examples--group-001.md#canonical-3131132132112300-2031210311300023-1320230303032133-2101121211031120-0000302131110002-0122121020210203-3302311013131221-3301311322210103)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_firewall/data-source.tf`; digest `sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931`.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```
