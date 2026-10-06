---
page_title: "xcsh_namespace"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace."
---

# xcsh_namespace

<a id="canonical-2013030003332230-3323112220123010-1133000333310232-3133330332200222-0123031332332100-3300310323320232-0203323000033132-1310310302002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_namespace

Reads Namespace information from F5 Distributed Cloud.

<a id="canonical-0123211001311111-3101231010100200-2203213321201131-0201103123203131-3112133011121012-1322302321313331-2022210332011300-1233202111210301"></a>

### Prerequisites for `xcsh_namespace`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3123231023002103-2011030003030100-1331312113130013-3132313032103021-0103202322320213-2311011103101310-1203002301320003-1301300022022322"></a>

### Minimal configuration for `xcsh_namespace`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

<a id="canonical-2031033332020100-2113312113303200-1103311211131301-0120300331121203-3121223102131311-3231331300311333-2020320010123003-2202331012122111"></a>

### Root configuration for `xcsh_namespace`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0123133131312210-0031233111332321-0223100132112310-1011003132020320-0000321322033301-0322321103301310-3200332022000220-3011211111312110"></a>

### Explore this collection for `xcsh_namespace`

- [Property reference](../guides/data-sources--namespace--reference--group-001.md#canonical-1233101021033332-1222000332100203-2301103231011223-1201002301311201-3021202223213233-0113330233212033-2200103302332302-0032213221333113)
- [Examples](../guides/data-sources--namespace--examples--group-001.md#canonical-0103010212031000-3110032213333233-1300320330100230-0132203113231333-2303211220001201-2022030333123213-2113111320021123-0131231021231332)
