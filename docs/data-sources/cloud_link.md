---
page_title: "xcsh_cloud_link"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link."
---

# xcsh_cloud_link

<a id="canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_cloud_link

Reads Cloud Link information from F5 Distributed Cloud.

<a id="canonical-1232313121100213-1313001300102230-2310322310222230-0002231011021320-0221310231032221-1232211333101131-3320121200301033-3232320031121101"></a>

### Prerequisites for `xcsh_cloud_link`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2113333003111213-0113133012322312-2322131210030130-1201123003010313-3333222003112230-3032200232122311-2132300131313232-0233310200230002"></a>

### Minimal configuration for `xcsh_cloud_link`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudLink Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudLink by name
data "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}

output "cloud_link_id" {
  value = data.xcsh_cloud_link.example.id
}
```

<a id="canonical-2101232310300301-1130201031120030-0213221311211223-2203222010021131-1321131332030000-0300130113132132-2301102030210231-1020031100031201"></a>

### Root configuration for `xcsh_cloud_link`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2323301210133210-0303031030231302-3033213312303302-2023100033022113-2233233031222212-3223320221200011-1232232012301320-3021000020231020"></a>

### Explore this collection for `xcsh_cloud_link`

- [Property reference](../guides/data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [Examples](../guides/data-sources--cloud_link--examples--group-001.md#canonical-0132232132113111-3321222211103031-1212310102011310-1000030122323222-0011333223100213-3213303310122310-1013321333031321-0301211111333201)
