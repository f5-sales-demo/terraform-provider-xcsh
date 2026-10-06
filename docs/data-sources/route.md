---
page_title: "xcsh_route"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route."
---

# xcsh_route

<a id="canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_route

Reads route rules that match incoming requests and select the corresponding actions.

<a id="canonical-3302301023130133-0101112301203320-0020112320001020-1231100302330310-0313300113222113-3201002011021033-0000131011110202-1111011330233031"></a>

### Prerequisites for `xcsh_route`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1321002331222231-2033311001211100-3000023201202203-1103223212110003-1103320132101213-3333111332022301-3310031112010222-3002122201132030"></a>

### Minimal configuration for `xcsh_route`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```

<a id="canonical-0320033100121100-2310323022102011-3303331133122131-0002133101212221-1000120302311113-3003011203030001-2122331200121312-1202310202031202"></a>

### Root configuration for `xcsh_route`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0002221131221203-1113231020003323-0023113332232133-2012302213221110-2323330121312001-0003232030332011-2002322002011221-1232100300003202"></a>

### Explore this collection for `xcsh_route`

- [Property reference](../guides/data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [Examples](../guides/data-sources--route--examples--group-001.md#canonical-2103210131020030-3332100111120032-0121012330312101-3320321023201002-0230023100223111-1033211232100203-1122111100333010-3310033010101013)
