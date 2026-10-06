---
page_title: "xcsh_secret_management_access"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access."
---

# xcsh_secret_management_access

<a id="canonical-0120220323020300-0010132221122301-2320223320011100-0233213331013202-1210021021231022-0233002013113103-1121001130200313-3213130202022123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_secret_management_access

Reads Secret Management Access information from F5 Distributed Cloud.

<a id="canonical-3001320010320131-3203031210320330-2020333303111320-3130330231133303-1212302301131312-0120300120022131-2331333200111121-0232030321332220"></a>

### Prerequisites for `xcsh_secret_management_access`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1030201120002010-3110001003120211-3120321330121212-1313120133002210-2122113013003012-0121100232001010-0223221211000102-0303013223022332"></a>

### Minimal configuration for `xcsh_secret_management_access`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```

<a id="canonical-0010201000303203-2102030203022212-2311211310233013-0020133120222022-3320311331221333-2330023121200103-2320220201211313-3321312303232213"></a>

### Root configuration for `xcsh_secret_management_access`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3312132233003322-2302321202311231-3002011323320003-1120200003132112-1022001131122321-0023103031211202-1220213333021020-1231020220333213"></a>

### Explore this collection for `xcsh_secret_management_access`

- [Property reference](../guides/data-sources--secret_management_access--reference--group-001.md#canonical-2312232123313113-0021013002332301-3020322212030101-3001020231232231-1321011213313220-1210113133302133-3321231013132103-0130012202121332)
- [Examples](../guides/data-sources--secret_management_access--examples--group-001.md#canonical-2233211101011210-0200323221322013-3111333211300333-1012012001332032-3313222213113030-0203213230312332-1221130132002332-0111031330020001)
