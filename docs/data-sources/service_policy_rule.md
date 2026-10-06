---
page_title: "xcsh_service_policy_rule"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule."
---

# xcsh_service_policy_rule

<a id="canonical-0203201000321103-2322122120203010-1333333333000031-3122311331020310-2002331222132033-2302323200122121-1300322033001000-2033110022321133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_service_policy_rule

Reads Service Policy Rule information from F5 Distributed Cloud.

<a id="canonical-3320121111201111-3213313111120301-0032033011331331-1320213201101231-1302111122203101-1322002210023213-2122302110112112-3022303032303122"></a>

### Prerequisites for `xcsh_service_policy_rule`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1032330001122111-2103201001012313-2201332001232013-1322302110301302-2212233233332230-1203320212000020-3322002103312001-1132020112222302"></a>

### Minimal configuration for `xcsh_service_policy_rule`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicyRule by name
data "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"
}

output "service_policy_rule_id" {
  value = data.xcsh_service_policy_rule.example.id
}
```

<a id="canonical-1101103201032023-0211332103210301-1213213121031131-1110130003210221-2321302033001033-1302312310012112-0021210021200333-1123333123323101"></a>

### Root configuration for `xcsh_service_policy_rule`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3211131022013210-1001020013132332-3223132302211121-3302020010322332-1232133133331313-3313131001222102-2201013321131023-3102130211112123"></a>

### Explore this collection for `xcsh_service_policy_rule`

- [Property reference](../guides/data-sources--service_policy_rule--reference--group-001.md#canonical-3331311020213130-0211011231012003-3002021121221003-3111211213002110-0101120220330331-1320033313332300-0113322023123130-0211211110021101)
- [Examples](../guides/data-sources--service_policy_rule--examples--group-001.md#canonical-1223112310330201-1133220313031233-2011110022130103-3103332022313100-1011012222033323-0131131222332313-3103032031112112-0002010031221231)
