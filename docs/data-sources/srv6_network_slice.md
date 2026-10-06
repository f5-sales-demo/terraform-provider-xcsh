---
page_title: "xcsh_srv6_network_slice"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice."
---

# xcsh_srv6_network_slice

<a id="canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_srv6_network_slice

Reads Srv6 Network Slice information from F5 Distributed Cloud.

<a id="canonical-1320022322323312-1113022232133203-1120010122320130-1213011220320301-3013310230300001-1123031211121121-2230312330331212-0202200320323213"></a>

### Prerequisites for `xcsh_srv6_network_slice`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2300200312311313-0011001011321310-0310323113123323-2321303012110113-2303100203002001-3300320223130333-3010233321330102-1323322032002000"></a>

### Minimal configuration for `xcsh_srv6_network_slice`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```

<a id="canonical-0330200133123211-0000101313202011-0023220201221223-2003032212333301-0131321211031102-2020130302030011-0321201303023300-1212013312201202"></a>

### Root configuration for `xcsh_srv6_network_slice`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3100103332203232-3132330302031000-1100122003033032-1113201002211031-0211032030133111-3102233031221311-1033001211022212-3333323120132122"></a>

### Explore this collection for `xcsh_srv6_network_slice`

- [Property reference](../guides/data-sources--srv6_network_slice--reference--group-001.md#canonical-3032223200202320-2101231011221211-2032122223212321-3111010311120030-0120223232313121-1201110023330122-3333233103131300-1301103122200231)
- [Examples](../guides/data-sources--srv6_network_slice--examples--group-001.md#canonical-2302231300112313-0022112123132022-3101033201332331-3132003300223310-2230202223132110-1302301203313020-2001202320332301-2031322030010122)
