---
page_title: "xcsh_ip_prefix_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set."
---

# xcsh_ip_prefix_set

<a id="canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_ip_prefix_set

Reads IP Prefix Set information from F5 Distributed Cloud.

<a id="canonical-2232002300112200-0032021020211222-2001330123100201-1010010002021121-1301310201322313-1210011300213112-2333233020202100-1101123020320303"></a>

### Prerequisites for `xcsh_ip_prefix_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0101210200233220-0201011122023311-0100200102222330-0013320312122302-1031001332331003-0203131332112311-0023213323231210-2020010102223303"></a>

### Minimal configuration for `xcsh_ip_prefix_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

<a id="canonical-2203312000331320-3223021222110110-2322212222022013-3100130001312300-2302120221330210-3310100210203233-1002323011110320-2221021321202131"></a>

### Root configuration for `xcsh_ip_prefix_set`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0033322201110100-3303121312101223-0210330323010011-2100210222312102-3232302210333013-0110032012302133-1300223221001321-0332220231331223"></a>

### Explore this collection for `xcsh_ip_prefix_set`

- [Property reference](../guides/data-sources--ip_prefix_set--reference--group-001.md#canonical-0210223101032303-3303221220210133-3110331001212210-2021220232323132-2011212301131311-0200220122100200-1103212333012112-1201121010111102)
- [Examples](../guides/data-sources--ip_prefix_set--examples--group-001.md#canonical-1101000022102203-0000201123101023-1320212232000013-0331231023232211-2210323020013101-0001101321233310-0201121321201012-1033323211201001)
