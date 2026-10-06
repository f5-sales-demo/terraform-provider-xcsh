---
page_title: "xcsh_swagger_object"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_swagger_object."
---

# xcsh_swagger_object

<a id="canonical-0113122122131303-3312011120012020-2331301002012223-3033302331132230-0311201302013200-2313211231330002-3123103203213300-2123202121003133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_swagger_object

Owns one immutable, content-verified Swagger object version. Content changes replace only the owned
version. Import uses namespace/name/version; latest and external presigned URLs are prohibited.
State contains the complete document.

<a id="canonical-0022001003321211-0100123123132032-1233230321332313-0020103132121311-3030021030103012-0312233302111333-1001201013000310-3213131321222002"></a>

### Prerequisites for `xcsh_swagger_object`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3201102022121021-1001233212200031-3112103303133022-3033302003332322-1213210300101212-1020021220332332-2312013302023233-0020232312311111"></a>

### Minimal configuration for `xcsh_swagger_object`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = { source = "f5-sales-demo/xcsh", version = ">= 15.3.0" }
  }
}

resource "xcsh_swagger_object" "example" {
  namespace = "demo"
  name      = "schema"
  content = jsonencode({
    openapi = "3.0.3"
    info    = { title = "Synthetic demo", version = "1" }
    paths   = {}
  })
}

output "swagger_path" {
  value = xcsh_swagger_object.example.path
}
```

<a id="canonical-0220010312103212-0131323131332201-1320322203211320-3001333003311100-0302232320231001-2220011112033032-1003332220330323-1213002022301000"></a>

### Root configuration for `xcsh_swagger_object`

Required root properties: `content`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1302210201302233-2020012303102310-0212031333011111-2020032003223131-3030020021303303-3233310100323113-3223032132202223-3121300033200211"></a>

### Explore this collection for `xcsh_swagger_object`

- [Property reference](../guides/resources--swagger_object--reference--group-001.md#canonical-1201120211111103-2221030103330020-2302333103300223-1010321323231300-1333012002311332-0113101012000220-0113100312031013-1131032312102331)
- [Examples](../guides/resources--swagger_object--examples--group-001.md#canonical-0310331101333000-0131302031302323-0231230320030331-0132320333022212-3100123310213201-2103312121102123-2033200122203232-3103210002013122)
- [Import](../guides/resources--swagger_object--lifecycle--group-001.md#canonical-0230231222233110-1330011230021332-1201331003033002-0023131311331000-2312110030201131-2200010032103120-1031100212120213-2023233312212201)
