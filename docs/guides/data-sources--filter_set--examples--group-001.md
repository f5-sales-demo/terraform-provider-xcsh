---
page_title: "xcsh_filter_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set examples."
---

# xcsh_filter_set examples

<a id="canonical-1221330213020113-0002233103033230-1202120022111310-1110203311220112-1101333133103123-2101032321203030-1310100203212111-0210320120222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- Examples

<a id="canonical-2212302130330202-3020222123211100-1222201303223030-0223132311300013-1300313232303102-1112121020132113-3022323113320323-3131210111021230"></a>

### Complete configurations for `xcsh_filter_set`

- [Data source](data-sources--filter_set--examples--group-001.md#canonical-2120310123232131-1010111201313002-2222112022000303-1312202013021330-2121202223020110-3313022212330131-1112230210131311-2011010312313322): valid configuration.

<a id="canonical-2120310123232131-1010111201313002-2222112022000303-1312202013021330-2121202223020110-3313022212330131-1112230210131311-2011010312313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md#canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001)
- [Examples](data-sources--filter_set--examples--group-001.md#canonical-1221330213020113-0002233103033230-1202120022111310-1110203311220112-1101333133103123-2101032321203030-1310100203212111-0210320120222021)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_filter_set/data-source.tf`; digest `sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d`.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```
