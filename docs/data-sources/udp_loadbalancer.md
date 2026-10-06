---
page_title: "xcsh_udp_loadbalancer"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer."
---

# xcsh_udp_loadbalancer

<a id="canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_udp_loadbalancer

Reads UDP Loadbalancer information from F5 Distributed Cloud.

<a id="canonical-2301303021013312-3101303022030300-2113112132123312-1202121101223212-0220301320101003-0013023032303120-3013211310330213-2312032023110220"></a>

### Prerequisites for `xcsh_udp_loadbalancer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1332321222132201-0203202231131311-3022013030313130-3100020110033300-0232003020331012-3330132120122333-3203000311020002-3301220103010220"></a>

### Minimal configuration for `xcsh_udp_loadbalancer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UDPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UDPLoadBalancer by name
data "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}

output "udp_loadbalancer_id" {
  value = data.xcsh_udp_loadbalancer.example.id
}
```

<a id="canonical-1023221211030332-1030202023210122-0123001323303132-3302112233330312-3010332001010332-3103232021330212-2233023022002020-2303210002011020"></a>

### Root configuration for `xcsh_udp_loadbalancer`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2201112130212001-3222003122223223-3023112231302202-3031233131121223-3120003231311321-2321322023303333-3022031032122202-2102131211231200"></a>

### Explore this collection for `xcsh_udp_loadbalancer`

- [Property reference](../guides/data-sources--udp_loadbalancer--reference--group-001.md#canonical-2100233330113013-1011020101321013-1113020011303120-3332020200220223-1130023001010003-0211031020201213-0121011023231220-3020101330032130)
- [Examples](../guides/data-sources--udp_loadbalancer--examples--group-001.md#canonical-3320030331211200-1233322332302002-3111203123123123-1000011112201332-1323220022102231-1010130310033302-2113232220113311-3310022022322132)
