---
page_title: "xcsh_workload_flavor examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor examples."
---

# xcsh_workload_flavor examples

<a id="canonical-0021201222321131-3310011210200020-1122213013302121-3312000030010212-3122000022200221-2210211023201200-0113202133200222-2323301131323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001)
- Examples

<a id="canonical-0221220032223022-1103202123310201-3302121003220122-3112010231332300-1110302030322330-1032313320332100-2223031132023302-2113223303103022"></a>

### Complete configurations for `xcsh_workload_flavor`

- [Resource](resources--workload_flavor--examples--group-001.md#canonical-3122300202300023-2031012101111123-0330001332122130-3331033120333023-3202103021033303-3031120222000320-1212110121311311-0223303331032102): valid configuration.

<a id="canonical-3122300202300023-2031012101111123-0330001332122130-3331033120333023-3202103021033303-3031120222000320-1212110121311311-0223303331032102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-1031333230130321-2120200203322003-3020210212312201-1302323133331112-0230023310233322-3012020133012201-2023100110323022-2002102111310001)
- [Examples](resources--workload_flavor--examples--group-001.md#canonical-0021201222321131-3310011210200020-1122213013302121-3312000030010212-3122000022200221-2210211023201200-0113202133200222-2323301131323210)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_workload_flavor/resource.tf`; digest `sha256:3a521cc20a27b0aff22bd7904a8747719375d23df70f4cf67625e0e522097422`.

```terraform
# WorkloadFlavor Resource Example
# Manages workload_flavor in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}
```
