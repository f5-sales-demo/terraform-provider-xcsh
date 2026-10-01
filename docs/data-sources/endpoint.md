---
page_title: "xcsh_endpoint landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint landing."
---

# xcsh_endpoint landing

<a id="canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322310012111103-3321002012233003-3031021102020021-2222211110122220-3113332230330010-2000000232321032-2221201200101322-0301211022332002"></a>

## xcsh_endpoint — xcsh_endpoint / 210000302100 / 2

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-3220331102211223-2312221113031232-0222233002321022-1003201020010101-2032021010102121-2213033021323103-0233032031002331-2312112322232132"></a>

## Prerequisites — xcsh_endpoint / 210000302100 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-1231213222230133-3113103331323313-1233201330010121-1133323101103122-1332103013112211-0130030031122011-1003012220121210-3231322022101330"></a>

## Minimal configuration — xcsh_endpoint / 210000302100 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```

<a id="canonical-1131031022322033-3131003021112001-3200032100021313-1122033213323330-0123022020301310-1332120131102032-2330020110210011-0200022223013203"></a>

## Root configuration — xcsh_endpoint / 210000302100 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3330302222020113-1203213220132022-1021031210310000-0202211003312111-2233110331200212-0033212013330010-2111322102331122-1000312323301130"></a>

## Next pages — xcsh_endpoint / 210000302100 / 6

- [Property reference](../guides/data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [Examples](../guides/data-sources--endpoint--examples--group-001.md#canonical-2301003011121223-0213312301322023-1222030132322003-1032221331233132-3232021011010330-0121032221322231-2132203021132323-3131313023221333)
