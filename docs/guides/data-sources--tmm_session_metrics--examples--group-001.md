---
page_title: "xcsh_tmm_session_metrics examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tmm_session_metrics examples."
---

# xcsh_tmm_session_metrics examples

<a id="canonical-1233211223130102-1300331332213330-1230330121133120-1002001001303333-0120223112101120-0033020020330023-1032001313021023-2023222120223131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-0001230110001322-3111312333111112-1132331032100110-2331112112213012-0233311131222231-0011030021132123-2021311322012101-1333301310111303)
- Examples

<a id="canonical-3012322220021023-1203001232322132-0330301333303111-1313131332201210-2211200032112131-1102012202311111-3330220131311030-1131221203321213"></a>

### Complete configurations for `xcsh_tmm_session_metrics`

- [Data source](data-sources--tmm_session_metrics--examples--group-001.md#canonical-2021203033233330-1201012121303000-3220311211300332-0310213011212013-3320303202032222-3320010012031321-1300330000321020-2332013121101232): valid configuration.

<a id="canonical-2021203033233330-1201012121303000-3220311211300332-0310213011212013-3320303202032222-3320010012031321-1300330000321020-2332013121101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md#canonical-0001230110001322-3111312333111112-1132331032100110-2331112112213012-0233311131222231-0011030021132123-2021311322012101-1333301310111303)
- [Examples](data-sources--tmm_session_metrics--examples--group-001.md#canonical-1233211223130102-1300331332213330-1230330121133120-1002001001303333-0120223112101120-0033020020330023-1032001313021023-2023222120223131)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tmm_session_metrics/data-source.tf`; digest `sha256:68c10b51b9e48b65d8a384096eef45f5c9fac00e135d661c65b0101b25c350c0`.

```terraform
# TmmSessionMetrics DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_tmm_session_metrics" "example" {
  namespace = "example-value"
}

output "tmm_session_metrics_result" {
  value = data.xcsh_tmm_session_metrics.example
}
```
