---
page_title: "xcsh_infraprotect_synchronize_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_synchronize_configuration examples."
---

# xcsh_infraprotect_synchronize_configuration examples

<a id="canonical-1121022323313230-3323102033132023-1113121222120103-1221110200113103-3211332131310301-3200223020013221-0201323130230231-0113211002201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-3003320123032222-1321103312000302-3220231112323333-3132221121231130-0320003131013001-3130000003113133-2111213321103223-2223330322121111)
- Examples

<a id="canonical-2221320331111311-0130311220232131-3021233301200111-1001020113322030-3012001300101201-1333331333230002-1132220232112201-3212001121323302"></a>

### Complete configurations for `xcsh_infraprotect_synchronize_configuration`

- [Action](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-1013120113203013-0132301301023223-3203220333212301-2331233010311231-2121001222121010-0103302311222323-0221311222302123-0132211200320311): valid configuration.

<a id="canonical-1013120113203013-0132301301023223-3203220333212301-2331233010311231-2121001222121010-0103302311222323-0221311222302123-0132211200320311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-3003320123032222-1321103312000302-3220231112323333-3132221121231130-0320003131013001-3130000003113133-2111213321103223-2223330322121111)
- [Examples](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-1121022323313230-3323102033132023-1113121222120103-1221110200113103-3211332131310301-3200223020013221-0201323130230231-0113211002201030)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf`; digest `sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227`.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```
