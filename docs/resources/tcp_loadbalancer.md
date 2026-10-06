---
page_title: "xcsh_tcp_loadbalancer"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer."
---

# xcsh_tcp_loadbalancer

<a id="canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_tcp_loadbalancer

Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across
origin pools.

<a id="canonical-3012303130230330-2222001010311302-3022023322100201-1010102331303312-2023301333223313-3131331332223331-3011032200310230-2011032021303113"></a>

### Prerequisites for `xcsh_tcp_loadbalancer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`.

- origin_pool: Backend servers for TCP/UDP traffic

- healthcheck: Monitor origin server health

<a id="canonical-1003020231311212-0113020122312100-1031122322130313-2310131301322213-0230210100033002-2330030230213233-1223310112200322-2013222212323123"></a>

### Minimal configuration for `xcsh_tcp_loadbalancer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TCPLoadBalancer Resource Example
# Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TCPLoadBalancer configuration
resource "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}
```

<a id="canonical-1111011322033212-2023031113030211-0221020133223033-0122233233310110-0212201221033302-2121133102301103-1101101011101010-0211132000212203"></a>

### Root configuration for `xcsh_tcp_loadbalancer`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3232302022001112-3201233332301111-1301001123201221-0133322033121210-3010133231013330-2320300311122030-0321103223220010-0023122130102300"></a>

### Explore this collection for `xcsh_tcp_loadbalancer`

- [Property reference](../guides/resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [Examples](../guides/resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
- [Import](../guides/resources--tcp_loadbalancer--lifecycle--group-001.md#canonical-0010103130330001-3321333010133010-0123010301022122-0013201010103131-2103100023233113-3300122101333230-2102201103021130-2001121313201221)
- [Timeouts](../guides/resources--tcp_loadbalancer--lifecycle--group-001.md#canonical-0113230333002312-3121321100313030-2332210230213110-1001203203302201-2030120211111123-2000020211011020-2322332023011313-3212203323213223)
