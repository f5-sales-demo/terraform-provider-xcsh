---
page_title: "xcsh_service_policy_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_set."
---

# xcsh_service_policy_set

<a id="canonical-1102322222212201-0323000120102213-0210100111122123-0013000010002101-1001031323000300-3302313333113330-3123230211030001-2322203121212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_service_policy_set

Reads an existing Service Policy Set from the requested namespace.

<a id="canonical-3312331200211123-3212123033332010-3203012202130120-0303001101013332-3112222022013003-3033332102002222-3231132313000001-1112030313123221"></a>

### Prerequisites for `xcsh_service_policy_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3123130203121123-0012122110213221-1130112121202320-3221110212020312-0122100323122310-3100122111332211-3320112110200123-2133010012000032"></a>

### Minimal configuration for `xcsh_service_policy_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

<a id="canonical-3210320100221300-2001210230102120-0020301033231231-3302030111333313-2233323213200002-0210021013120013-3233211130223132-1321023022303121"></a>

### Root configuration for `xcsh_service_policy_set`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1103223213011211-1023231331333131-2310013002133211-0312210203320030-0322210233221110-2013311023203211-0231233031133003-3032203121021303"></a>

### Explore this collection for `xcsh_service_policy_set`

- [Property reference](../guides/data-sources--service_policy_set--reference--group-001.md#canonical-3111031112022212-1110223232220233-3122113000331220-3030210330100312-1310312331103022-1302102001211000-3222222101332321-3231131202201203)
- [Examples](../guides/data-sources--service_policy_set--examples--group-001.md#canonical-1210220030301313-1112121100102103-1032120221121012-3133300310001112-1132200310032010-1032303010220113-3230331113010111-2023023333130212)
