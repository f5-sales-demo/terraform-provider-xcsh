---
page_title: "xcsh_network_policy_view examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view examples."
---

# xcsh_network_policy_view examples

<a id="canonical-0330213012123133-3101332210132313-1201320303023213-1223310233311020-0001230131310102-0030332112310201-1202303102201012-0011210012111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- Examples

<a id="canonical-3202132111003000-0132203131020032-2003333030202222-0110311312002013-2123123101012231-1333003012111302-2231133302121121-3111113001310110"></a>

### Complete configurations for `xcsh_network_policy_view`

- [Data source](data-sources--network_policy_view--examples--group-001.md#canonical-1121223333211212-3310103032120301-0011323123031111-1330000112112300-0130211020023132-3101102212301313-3001322302212313-3300202112323030): valid configuration.

<a id="canonical-1121223333211212-3310103032120301-0011323123031111-1330000112112300-0130211020023132-3101102212301313-3001322302212313-3300202112323030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md#canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120)
- [Examples](data-sources--network_policy_view--examples--group-001.md#canonical-0330213012123133-3101332210132313-1201320303023213-1223310233311020-0001230131310102-0030332112310201-1202303102201012-0011210012111031)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_view/data-source.tf`; digest `sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a`.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```
