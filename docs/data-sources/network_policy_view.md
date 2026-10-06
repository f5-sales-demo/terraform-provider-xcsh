---
page_title: "xcsh_network_policy_view"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view."
---

# xcsh_network_policy_view

<a id="canonical-1101130230110012-3022101000100330-0002130003202312-1101110301302313-0202110313103012-0333330312131122-0312220113032211-3303322222220120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_policy_view

Reads Network Policy View information from F5 Distributed Cloud.

<a id="canonical-2023023020333313-0010002310030122-0231212121133312-2310131021223012-3230220323232000-3331311302130311-1123231123131002-0133231100200232"></a>

### Prerequisites for `xcsh_network_policy_view`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1133220011322331-1321110023013333-0200121111333201-0013232311321322-0332011030222211-3011112111211033-1220301333320210-1222101120200301"></a>

### Minimal configuration for `xcsh_network_policy_view`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```

<a id="canonical-1200023010133313-2133032202203031-3230123100101321-3131010213330232-3203313222121022-2021003202101133-2330030222312302-0001233213201032"></a>

### Root configuration for `xcsh_network_policy_view`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2201032002300223-3212203000303023-3333123300203321-1301033102221032-3230113033211332-0330002022132031-2003013331020201-1030120210321322"></a>

### Explore this collection for `xcsh_network_policy_view`

- [Property reference](../guides/data-sources--network_policy_view--reference--group-001.md#canonical-1022223021101230-1222100103330000-1000012100203123-1033200021310031-2322131202002013-3031332320233313-3032231023013330-2300002312132121)
- [Examples](../guides/data-sources--network_policy_view--examples--group-001.md#canonical-0330213012123133-3101332210132313-1201320303023213-1223310233311020-0001230131310102-0030332112310201-1202303102201012-0011210012111031)
