---
page_title: "xcsh_workload_flavor"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor."
---

# xcsh_workload_flavor

<a id="canonical-2030203000323030-0232123232021133-3020001332010013-1013132311011132-1021330332030103-1003300002310323-0333232112231113-1231333120223030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_workload_flavor

Reads Workload Flavor information from F5 Distributed Cloud.

<a id="canonical-3131202331202331-1131321112232220-1330000310223033-1202001111211012-0022320232011331-2221020033012203-1003303302200303-2323120233101303"></a>

### Prerequisites for `xcsh_workload_flavor`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2100201131102031-3003021021200010-2231221201201203-3000313030211333-3111133221300301-0312031000003320-1202222000223321-2210321032231033"></a>

### Minimal configuration for `xcsh_workload_flavor`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3013302033001332-2122320021200232-2233302121330312-0133230133322210-1302031230323002-2110203333200111-0032121020222223-1132122331311311"></a>

### Root configuration for `xcsh_workload_flavor`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1203122213002302-3010112133022312-1302100130133011-2330321210022103-2211121031110201-1113203100300300-1122101301132300-3202223302003200"></a>

### Explore this collection for `xcsh_workload_flavor`

- [Property reference](../guides/data-sources--workload_flavor--reference--group-001.md#canonical-2101210323120012-1332103220321021-0330203102320020-3302111001233322-1031112300210030-1300131022030232-2232013130123221-0310120231323221)
- [Examples](../guides/data-sources--workload_flavor--examples--group-001.md#canonical-2221211011311201-1233330312213233-3212032023203223-0333133320100023-3323101111023020-3300022013322123-1011132000030232-2311330312202102)
