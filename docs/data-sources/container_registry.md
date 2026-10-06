---
page_title: "xcsh_container_registry"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry."
---

# xcsh_container_registry

<a id="canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_container_registry

Reads Container Registry information from F5 Distributed Cloud.

<a id="canonical-1201002223312331-1323221123111010-1302033210131001-3003123303020220-2113303302323233-3111003030302001-2102031031122133-2213023003211220"></a>

### Prerequisites for `xcsh_container_registry`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1023123220121323-3212123130201331-0002323110230122-0230023322011013-1211123031121331-2011013022302000-1021123023310332-3021330130102331"></a>

### Minimal configuration for `xcsh_container_registry`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ContainerRegistry Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ContainerRegistry by name
data "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"
}

output "container_registry_id" {
  value = data.xcsh_container_registry.example.id
}
```

<a id="canonical-3122302132213021-2122123031313322-2033320223210312-2201100100321000-0131303321131230-2112033033322002-3221310320221101-3131003011231112"></a>

### Root configuration for `xcsh_container_registry`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3322220222311222-0310231223000200-3313010023023111-0221302123003131-0023001020203130-0220032312121011-2001113313031030-1212222132203022"></a>

### Explore this collection for `xcsh_container_registry`

- [Property reference](../guides/data-sources--container_registry--reference--group-001.md#canonical-3211202121023222-3310130210222030-3300301333003322-1210013130013311-0121323231231103-3102100321133312-2232203033100102-2132331202300121)
- [Examples](../guides/data-sources--container_registry--examples--group-001.md#canonical-2300022233211321-3223202330210002-0331022230100230-1002132001033202-0230103302011023-3102232023121120-1301332301023312-2131310121102020)
