---
page_title: "xcsh_workload"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload."
---

# xcsh_workload

<a id="canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_workload

Reads Workload information from F5 Distributed Cloud.

<a id="canonical-1320323331132320-3113121101211131-1110223122020231-2332223032101330-1302301000310000-1302203010111302-3222131123122331-0200101221112333"></a>

### Prerequisites for `xcsh_workload`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_k8s`.

- virtual_k8s: Namespace for workload deployment

<a id="canonical-2023222120322212-2100221333111113-0200122011111232-3110213003232033-2133203301133201-2123033323201013-1333101122120300-3203312223211331"></a>

### Minimal configuration for `xcsh_workload`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```

<a id="canonical-3011211310131032-3321100001200201-0131002231210321-1022233111212223-1233103331033332-2332033002303003-1200120123022021-2133112113020303"></a>

### Root configuration for `xcsh_workload`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1221010202223321-0121132023030312-2121102101000022-3010233111111033-2330331000301101-1032113311033002-3331333012113200-1233000333223013"></a>

### Explore this collection for `xcsh_workload`

- [Property reference](../guides/data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [Examples](../guides/data-sources--workload--examples--group-001.md#canonical-3110211130110313-2113120222301013-1332130101202131-2312310013002012-1223332113022010-0003103230022222-1213221321232131-0213012002030102)
