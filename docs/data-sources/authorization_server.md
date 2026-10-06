---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server."
---

# xcsh_authorization_server

<a id="canonical-3021220012103201-0130032102230113-0221231210301022-2323112200300110-3212010013210200-3300200232211023-2113102021313101-1310130133021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_authorization_server

Reads Authorization Server information from F5 Distributed Cloud.

<a id="canonical-2200311330321203-3103000031013013-2120202023132321-1001011013103121-1210102302221010-2221210111312113-1333032332300202-3200203120121131"></a>

### Prerequisites for `xcsh_authorization_server`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1311000001303010-2101333302210231-1313222221030021-0300123023013033-3130231320223211-3233021123330320-0123231031111323-1131212313203032"></a>

### Minimal configuration for `xcsh_authorization_server`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```

<a id="canonical-3303023210120023-0113300130110220-0313213111302122-0333012003303023-0323010122313322-1000021301232111-0133032020123210-1203011020312311"></a>

### Root configuration for `xcsh_authorization_server`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1312302200222003-0313332211310211-2232312010212212-1231330120332211-0112223332221011-3021330323222133-0323233100133030-0101220110231310"></a>

### Explore this collection for `xcsh_authorization_server`

- [Property reference](../guides/data-sources--authorization_server--reference--group-001.md#canonical-1033321203032211-3120322231302331-2201330321022031-2230022113210210-1333130133303221-2233013123102231-1302133330012130-3023121200332031)
- [Examples](../guides/data-sources--authorization_server--examples--group-001.md#canonical-2033301013221130-0030022030303002-3031133301332223-3230301213310203-1301033002012133-1023011313333003-3130221223113223-2321031313000012)
