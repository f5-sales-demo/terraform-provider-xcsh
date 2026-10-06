---
page_title: "xcsh_sensitive_data_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy."
---

# xcsh_sensitive_data_policy

<a id="canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

<a id="canonical-0023202030021013-2003210231013030-0300203333332213-3123101222310303-2122032333223131-0113323213003002-0232313211321320-3003230110121212"></a>

### Prerequisites for `xcsh_sensitive_data_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-0001112103001132-3331101103120003-0210233112023321-3303100021211110-0213112233023332-0231032023101200-2332302230210000-3201321001233223"></a>

### Minimal configuration for `xcsh_sensitive_data_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SensitiveDataPolicy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SensitiveDataPolicy configuration
resource "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}
```

<a id="canonical-1110330112232111-0010321301213003-1001100301133022-3031201310110020-2233322022033322-1023020032112002-3300111231101202-1103210130023222"></a>

### Root configuration for `xcsh_sensitive_data_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1110230333012310-3222211201111211-0312331020110220-2033300322113112-1010020301210002-2320000200200233-3202113011112032-0320030112331213"></a>

### Explore this collection for `xcsh_sensitive_data_policy`

- [Property reference](../guides/resources--sensitive_data_policy--reference--group-001.md#canonical-0203133212301203-1132112323323302-1312032000023031-3213122123301323-0203001200123121-2232023331202233-1212100131131013-0000331232331001)
- [Examples](../guides/resources--sensitive_data_policy--examples--group-001.md#canonical-2122111122233010-1200132332303000-0300023113021202-0321320011112102-3001002020221313-0133111333131010-1101210322330033-2000103302332023)
- [Import](../guides/resources--sensitive_data_policy--lifecycle--group-001.md#canonical-1032002330023002-2100011203003323-1303201300020322-0230212210322211-3313213132101003-2112112213210012-0212310221123230-2001212033013002)
- [Timeouts](../guides/resources--sensitive_data_policy--lifecycle--group-001.md#canonical-1330012010011233-2223113022230230-3011323320120312-0021110123020331-0121331311212230-0033011033000000-1123231311232321-1301131203200020)
