---
page_title: "xcsh_authentication"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication."
---

# xcsh_authentication

<a id="canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_authentication

Reads Authentication information from F5 Distributed Cloud.

<a id="canonical-1213000100303323-3231321233222210-3232031203200201-0332321103200211-0120100322112233-3000312002001311-1213013332112202-3230201102310223"></a>

### Prerequisites for `xcsh_authentication`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1221331000101303-1200030331131132-0122220120313322-3231212221100121-0300133032003130-3233332022100002-1233301302002313-1103301203300310"></a>

### Minimal configuration for `xcsh_authentication`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

<a id="canonical-0023213333303333-1300302331303233-1110113002012233-3112312121220202-0231130212122033-3301230130222330-0100120232233310-0220130122031012"></a>

### Root configuration for `xcsh_authentication`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3310310010222120-0230131112312021-0102312023323331-0222320110202222-0113130130033301-3102022113010222-3233102010013022-3021232131222133"></a>

### Explore this collection for `xcsh_authentication`

- [Property reference](../guides/data-sources--authentication--reference--group-001.md#canonical-2002323222211220-0230023001010221-0122022322200120-0100003312210032-3002222000120032-3021211330022230-1032122111002303-1111303000311210)
- [Examples](../guides/data-sources--authentication--examples--group-001.md#canonical-3311100333131130-3203031113312032-0213122233111110-1002111313031302-0000011213332103-0001333330310320-1333121200120313-0112111020223000)
