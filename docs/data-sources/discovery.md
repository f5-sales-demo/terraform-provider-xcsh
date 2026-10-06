---
page_title: "xcsh_discovery"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery."
---

# xcsh_discovery

<a id="canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_discovery

Reads Discovery information from F5 Distributed Cloud.

<a id="canonical-2201222103332033-3232230013101103-2032021330231312-0012333231120031-3221133130113230-1203122313320003-2032021100302330-2111323211013111"></a>

### Prerequisites for `xcsh_discovery`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1021231233013031-3100001112310322-3223000011122312-1320213031132222-0223321132223320-0123323133233002-2232110322121312-2002101022203122"></a>

### Minimal configuration for `xcsh_discovery`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

<a id="canonical-2303122133002000-0332112303022123-3313130200301203-2332001031121101-0322330203301113-3032300001112330-0332232203210103-1131330011231310"></a>

### Root configuration for `xcsh_discovery`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1311013220200103-1203021022121022-0233310030210311-1120221132012133-0031133111331032-2210300210210213-1100233211223212-1121023100033211"></a>

### Explore this collection for `xcsh_discovery`

- [Property reference](../guides/data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [Examples](../guides/data-sources--discovery--examples--group-001.md#canonical-2233113310202003-3011212032322120-0321110101103330-1110220301021333-2113230300122223-0032111232321131-3031100223030032-3233311203131303)
