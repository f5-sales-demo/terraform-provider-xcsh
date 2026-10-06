---
page_title: "xcsh_malicious_user_mitigation"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation."
---

# xcsh_malicious_user_mitigation

<a id="canonical-1101131202021321-3221321112031101-3221121031010113-2211012203220000-0302111121132132-3121301122233133-1013203002033303-3311201201113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_malicious_user_mitigation

Manages malicious\_user\_mitigation creates a new object in the storage backend for
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-1321321011200123-2002232201023130-1021013031213222-3201311313022232-0310322233112302-1330131102122112-2232021201021300-1312122203120012"></a>

### Prerequisites for `xcsh_malicious_user_mitigation`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1301301323231103-3333301112111132-0213220111032303-3000012101030001-3103302100101122-0200121023210010-1112202221101103-0211233011000230"></a>

### Minimal configuration for `xcsh_malicious_user_mitigation`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MaliciousUserMitigation Resource Example
# Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MaliciousUserMitigation configuration
resource "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}
```

<a id="canonical-0311021323022310-1110322011212330-3012110221233121-0123312221200330-2322031011221122-3320013103103232-1211023200213220-2111201131201331"></a>

### Root configuration for `xcsh_malicious_user_mitigation`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2001120120003201-2122031220311112-2012210011232011-3003323333002223-0022122001223103-2130200311111103-2220322020012111-2101033213221232"></a>

### Explore this collection for `xcsh_malicious_user_mitigation`

- [Property reference](../guides/resources--malicious_user_mitigation--reference--group-001.md#canonical-1110001002333302-2300113210021312-0300220333020331-3110232013111231-2133223211023113-1021322132102302-1212032231131231-3211212320213101)
- [Examples](../guides/resources--malicious_user_mitigation--examples--group-001.md#canonical-0122333121121121-2133233012323031-2100322101123321-3101010121133222-1002313333123011-2132000233033313-0320000203211130-1132123110011023)
- [Import](../guides/resources--malicious_user_mitigation--lifecycle--group-001.md#canonical-0201333231030120-0021033003300132-2301313031132121-0303001233310032-0322010201302101-1333031132113002-1310121202003303-3013211320321122)
- [Timeouts](../guides/resources--malicious_user_mitigation--lifecycle--group-001.md#canonical-2003311330130020-0321230120033220-0313222211013010-2030201021331333-3210202020023022-3223320132011101-1013032130200110-2222303103223111)
