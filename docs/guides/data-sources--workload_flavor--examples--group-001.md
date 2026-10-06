---
page_title: "xcsh_workload_flavor examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor examples."
---

# xcsh_workload_flavor examples

<a id="canonical-2221211011311201-1233330312213233-3212032023203223-0333133320100023-3323101111023020-3300022013322123-1011132000030232-2311330312202102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030)
- Examples

<a id="canonical-1013332110221213-2001313022312233-1232102311130221-2003132222012013-3133333210201030-1220021223113332-2001330022132022-1213023001331203"></a>

### Complete configurations for `xcsh_workload_flavor`

- [Data source](data-sources--workload_flavor--examples--group-001.md#canonical-3203031020011131-3202033303100303-1322213220233302-3302132331023112-3113123022323233-1123022212331032-0111222222310331-0012310120313013): valid configuration.

<a id="canonical-3203031020011131-3202033303100303-1322213220233302-3302132331023112-3113123022323233-1123022212331032-0111222222310331-0012310120313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md#canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030)
- [Examples](data-sources--workload_flavor--examples--group-001.md#canonical-2221211011311201-1233330312213233-3212032023203223-0333133320100023-3323101111023020-3300022013322123-1011132000030232-2311330312202102)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload_flavor/data-source.tf`; digest `sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173`.

```terraform
# WorkloadFlavor Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WorkloadFlavor by name
data "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}

output "workload_flavor_id" {
  value = data.xcsh_workload_flavor.example.id
}
```
