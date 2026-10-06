---
page_title: "xcsh_network_global_controller_sso_egress"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_controller_sso_egress."
---

# xcsh_network_global_controller_sso_egress

<a id="canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_global_controller_sso_egress

Global Controller SSO egress IPv4 addresses. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

<a id="canonical-0131000033221311-2112033233223211-0212001121122032-2233122003023122-1031232031303230-2312231220020220-0021211013111310-3021321323310332"></a>

### Prerequisites for `xcsh_network_global_controller_sso_egress`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3021231233013322-3123230231323011-1202020020121300-3133000321232220-2331213020013213-0132013023033122-1300302222322320-3211300303110222"></a>

### Minimal configuration for `xcsh_network_global_controller_sso_egress`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

<a id="canonical-3322320202131210-2133032100203223-3213313111333121-1010013310030023-3101211130031102-0012313200323022-0322330122000023-2102201302232110"></a>

### Root configuration for `xcsh_network_global_controller_sso_egress`

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2301323020301023-0231313333103131-2333120003132112-3131231101211322-3003201110013002-1120312330202213-2311220120122200-3121002221033333"></a>

### Explore this collection for `xcsh_network_global_controller_sso_egress`

- [Property reference](../guides/data-sources--network_global_controller_sso_egress--reference--group-001.md#canonical-3003322331121130-3222331001000333-2211002302211022-1222320311122100-3201133312312100-1030203313320022-0120013131302332-2021000013120000)
- [Examples](../guides/data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-1330212023033012-1303023012013222-0112221202020303-1011102330201000-3003322221102203-3310100311110131-1111331310003200-2210323333000232)
