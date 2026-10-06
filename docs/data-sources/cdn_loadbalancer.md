---
page_title: "xcsh_cdn_loadbalancer"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer."
---

# xcsh_cdn_loadbalancer

<a id="canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cdn_loadbalancer

Reads CDN Loadbalancer information from F5 Distributed Cloud.

<a id="canonical-2133230233223102-3022332321032303-2320031230230013-3231313031033321-0011001232030001-2002001002031013-0311011310311122-1013300121123103"></a>

### Prerequisites for `xcsh_cdn_loadbalancer`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cdn_origin_pool`.

- cdn_origin_pool: Origin servers for CDN content

<a id="canonical-1103011120313023-2131212232312311-1313012022211103-3103330311221230-1112002010321313-2023233300023132-1212331320321311-2232032122330300"></a>

### Minimal configuration for `xcsh_cdn_loadbalancer`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNLoadBalancer by name
data "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"
}

output "cdn_loadbalancer_id" {
  value = data.xcsh_cdn_loadbalancer.example.id
}
```

<a id="canonical-3101012120301301-2000333222121221-0031120330222112-0232312333030210-0022201030311111-2310133032022111-0030203300300313-0122103222020021"></a>

### Root configuration for `xcsh_cdn_loadbalancer`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2103220203221301-3133113200130132-0323232132230223-3032013020322003-2033321133000130-0310331300002303-0212300302310311-2222012022122333"></a>

### Explore this collection for `xcsh_cdn_loadbalancer`

- [Property reference](../guides/data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [Examples](../guides/data-sources--cdn_loadbalancer--examples--group-001.md#canonical-3110121132323313-0011302232310122-0300300101301303-0101333112121220-3330120023223101-1130020331101223-0203301010301122-3330120123222302)
