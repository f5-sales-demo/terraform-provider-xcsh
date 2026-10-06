---
page_title: "xcsh_protected_application"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application."
---

# xcsh_protected_application

<a id="canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_protected_application

Reads Protected Application information from F5 Distributed Cloud.

<a id="canonical-2033332023323113-2120213110211223-1233023130232121-0002330122000300-1321332121201303-2233001330330210-0220100111202103-2133112003300230"></a>

### Prerequisites for `xcsh_protected_application`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0132103123100031-3013022210312021-3113313222020333-1112000213220223-1231220000101231-0322012130030032-0101323222011301-1321320010013232"></a>

### Minimal configuration for `xcsh_protected_application`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```

<a id="canonical-2201323013131022-3302012120001011-1013101011300333-1132033022320212-3003321100013232-0032222121013223-2101320200120030-0103230321201012"></a>

### Root configuration for `xcsh_protected_application`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0232323223100120-2310323123330011-2112122033030021-3112322003212231-2331022032310200-0101030133302223-3300332310111000-1110320311012312"></a>

### Explore this collection for `xcsh_protected_application`

- [Property reference](../guides/data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Examples](../guides/data-sources--protected_application--examples--group-001.md#canonical-2201033120231331-3003011110101323-2002010000323102-3000312102110323-2131311201000033-2021230001203130-1330202023130023-2223332012123102)
