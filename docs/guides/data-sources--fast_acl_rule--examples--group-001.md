---
page_title: "xcsh_fast_acl_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule examples."
---

# xcsh_fast_acl_rule examples

<a id="canonical-3211232130101212-1130233301203323-3130113131323113-0011112033203332-2221012002203030-2101332011200103-0103233112100130-3302123121230230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- Examples

<a id="canonical-3130302112012002-3312200022000023-3111200203133011-1122120300000203-2331300231200332-2112120200233311-1002113332013202-2220012320313103"></a>

### Complete configurations for `xcsh_fast_acl_rule`

- [Data source](data-sources--fast_acl_rule--examples--group-001.md#canonical-1102032020313002-1130313020120112-2130000113222232-3010212201233021-0003220032230233-3111020133122120-2300210111300232-3101011300220210): valid configuration.

<a id="canonical-1102032020313002-1130313020120112-2130000113222232-3010212201233021-0003220032230233-3111020133122120-2300210111300232-3101011300220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-2322211232231123-2000222302002322-0230323221110333-3020201120001310-3010102231232032-3203221031003230-1300331222110232-2023131232010032)
- [Examples](data-sources--fast_acl_rule--examples--group-001.md#canonical-3211232130101212-1130233301203323-3130113131323113-0011112033203332-2221012002203030-2101332011200103-0103233112100130-3302123121230230)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl_rule/data-source.tf`; digest `sha256:d3a85aa2d2c4a98927959b827b07db2533ed58dbfde3c28ad82e3e2235d3a843`.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```
