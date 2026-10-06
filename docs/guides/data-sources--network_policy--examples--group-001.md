---
page_title: "xcsh_network_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy examples."
---

# xcsh_network_policy examples

<a id="canonical-3101310111230003-2031330022332131-3120112210211023-0111230213001020-2002212100121322-2120101311133223-3111333310021231-2312101332310311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- Examples

<a id="canonical-3021201210220310-1323213012123320-2212133022210200-3110012033122101-3212103021130002-0311032031111220-2022112231113303-2022101202302211"></a>

### Complete configurations for `xcsh_network_policy`

- [Data source](data-sources--network_policy--examples--group-001.md#canonical-2132333112011321-3110031010120001-1012210013132100-1300112110232020-0212011303021021-1331012033020110-3220200233013013-1121211120210330): valid configuration.

<a id="canonical-2132333112011321-3110031010120001-1012210013132100-1300112110232020-0212011303021021-1331012033020110-3220200233013013-1121211120210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Examples](data-sources--network_policy--examples--group-001.md#canonical-3101310111230003-2031330022332131-3120112210211023-0111230213001020-2002212100121322-2120101311133223-3111333310021231-2312101332310311)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy/data-source.tf`; digest `sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01`.

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
