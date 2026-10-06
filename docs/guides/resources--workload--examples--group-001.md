---
page_title: "xcsh_workload examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload examples."
---

# xcsh_workload examples

<a id="canonical-3201002123330130-0130300023013311-3322300132010333-1032200112231302-3120201212103100-0310100313110131-1230202231003130-0112021013220320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- Examples

<a id="canonical-0300012132131000-0000103020022311-2230001211322303-0011200210030321-2001201312323321-0131122230211103-3221232122022010-1010123102212312"></a>

### Complete configurations for `xcsh_workload`

- [Resource](resources--workload--examples--group-001.md#canonical-3202301020033123-1000311020333023-2200330011313010-3211330331312122-3231102022332233-0023313320023323-2321332001221132-0001132011323320): valid configuration.

<a id="canonical-3202301020033123-1000311020333023-2200330011313010-3211330331312122-3231102022332233-0023313320023323-2321332001221132-0001132011323320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Examples](resources--workload--examples--group-001.md#canonical-3201002123330130-0130300023013311-3322300132010333-1032200112231302-3120201212103100-0310100313110131-1230202231003130-0112021013220320)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload/resource.tf`; digest `sha256:95bc586cf100a6ef22147aba592992dc22645f4b290f9226bf45165f3414e58c`.

```terraform
# Workload Resource Example
# Manages a Workload resource in F5 Distributed Cloud for workload.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Workload configuration
resource "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}
```
