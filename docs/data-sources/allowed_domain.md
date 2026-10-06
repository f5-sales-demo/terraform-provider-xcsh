---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_allowed_domain."
---

# xcsh_allowed_domain

<a id="canonical-3031012320322023-2113120332301001-0010123230112302-0303132031320332-1032231003302322-1130311111131220-1203223310030220-1032110013113220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_allowed_domain

Reads Allowed Domain information from F5 Distributed Cloud.

<a id="canonical-3200312313331333-2332333120332232-2333332001230021-3030102331333100-2321303320220111-2012321233202211-2233320011023120-3300031321132212"></a>

### Prerequisites for `xcsh_allowed_domain`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0301313322333111-2332130010101022-3201321113310021-0100121012313201-3133033030223000-2202121230221112-2213003131221012-3220023311202202"></a>

### Minimal configuration for `xcsh_allowed_domain`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
```

<a id="canonical-3200110101201122-0231330223212303-3233322321130330-3030113202313200-0103302221333122-3320211203332022-1032233122101320-1121102321131232"></a>

### Root configuration for `xcsh_allowed_domain`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0120120102313021-0100113121021121-1312311230122322-0033331113011311-0322021112111102-3011131133323220-0002023122000211-1120012232001322"></a>

### Explore this collection for `xcsh_allowed_domain`

- [Property reference](../guides/data-sources--allowed_domain--reference--group-001.md#canonical-2001103322233010-3101211230020031-2213110101011302-0313200303102033-3001012300210310-2320011103111122-1213233212110011-2230323220112210)
- [Examples](../guides/data-sources--allowed_domain--examples--group-001.md#canonical-1302323331312313-3102131232131300-1221010023033210-3232331202213030-1130131020111011-1011011203223203-0321101210111222-2303301203132121)
