---
page_title: "xcsh_http_loadbalancer"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer."
---

# xcsh_http_loadbalancer

<a id="canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_http_loadbalancer

Reads an HTTP Load Balancer that routes HTTP and HTTPS traffic with security controls.

<a id="canonical-1012320023103100-2021101201301330-1102131330311221-1123032031312213-3223221110303021-2213013011100312-1212310022320323-1002322101130230"></a>

### Prerequisites for `xcsh_http_loadbalancer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`, `app_firewall`, `certificate`, `rate_limiter`, `service_policy`, `bot_defense_policy`.

- origin_pool: Backend servers for traffic distribution

- app_firewall: WAF protection (requires WAAP subscription)

- healthcheck: Monitor backend availability

- certificate: TLS termination for HTTPS

- rate_limiter: Protect against traffic spikes

<a id="canonical-2021312201200231-2030132231200220-3123120130301302-3121012001022003-2100310221330203-0221313121002131-0112113103331301-1332012021023301"></a>

### Minimal configuration for `xcsh_http_loadbalancer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# HTTPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing HTTPLoadBalancer by name
data "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"
}

output "http_loadbalancer_id" {
  value = data.xcsh_http_loadbalancer.example.id
}
```

<a id="canonical-0031030113301023-0211331023211003-3232131200133010-1112013321133032-1311112222020022-1013110032331022-1311312310002313-0300221232031303"></a>

### Root configuration for `xcsh_http_loadbalancer`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2012113111302033-3221011212002332-2233001100103320-1133220023003031-1321322303222131-2130203202021331-1100111102101022-3122233101211133"></a>

### Explore this collection for `xcsh_http_loadbalancer`

- [Property reference](../guides/data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [Examples](../guides/data-sources--http_loadbalancer--examples--group-001.md#canonical-2012103322223222-2130303310322103-3112301200231103-1003312232213332-1311123122123311-3322111313213332-3221000203301021-1012200122123110)
