---
page_title: "xcsh_endpoint"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint."
---

# xcsh_endpoint

<a id="canonical-1030332030133033-3031112013120203-2020030113002032-1121120230133230-0032331030330322-0110231131310002-2031000003132100-1320211123110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-2222222300113321-0312320312313233-3220202033210230-3203301033332312-3002012122200032-3011320011120312-0310303033133120-2211303130013133"></a>

### Prerequisites for `xcsh_endpoint`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-0111330033131223-1120031102300031-1213020322300122-2201023322101202-3101001210102233-1313122233010323-2112223032133133-3103023023130211"></a>

### Minimal configuration for `xcsh_endpoint`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

<a id="canonical-2030023011321331-2222333133022111-3130022023203313-2203121033203033-2302000220122303-3103002220132211-1130011330233330-2102023121121331"></a>

### Root configuration for `xcsh_endpoint`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0313312033101101-2103233302233220-2003033113233323-0202223310200030-0030320202312201-1122133210322001-2020233011021331-0021231231120133"></a>

### Explore this collection for `xcsh_endpoint`

- [Property reference](../guides/resources--endpoint--reference--group-001.md#canonical-2202001221303013-2331112001103210-3300011012323210-1313330313021303-3322000202012112-3333303230220023-1210111320223001-3012001121233131)
- [Examples](../guides/resources--endpoint--examples--group-001.md#canonical-1121101120010011-2333021113122233-0010112232231322-0031100220232023-2001101311210332-1010331101033020-2312023010330321-0223130201200122)
- [Import](../guides/resources--endpoint--lifecycle--group-001.md#canonical-2011110030121120-1212131101212003-1032010213223012-0101131313103302-0231112212111322-0212322233113233-0230120303132200-2220122112220322)
- [Timeouts](../guides/resources--endpoint--lifecycle--group-001.md#canonical-2131130210121013-3302210031131331-2303203333002020-2130313320320220-0312002211222311-0113302201330021-1221020202100021-0230321302113303)
