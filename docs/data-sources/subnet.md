---
page_title: "xcsh_subnet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet landing."
---

# xcsh_subnet landing

<a id="canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211003112101300-2021312222313120-0331230021230312-1131303231331102-0331130202130023-0203120000332011-3120122003320121-1102202221032110"></a>

## xcsh_subnet — xcsh_subnet / 221220131001 / 2

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

<a id="canonical-1312230102232132-2201310032122231-1103022131131321-2003100231203033-1201010200320221-3001130023312320-0110103103231030-3121000023213223"></a>

## Prerequisites — xcsh_subnet / 221220131001 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1213200002002311-2031300122030311-2033021232123023-2032221102313311-3010121300231302-1210030330123000-1131032320223322-2033123133303133"></a>

## Minimal configuration — xcsh_subnet / 221220131001 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Subnet by name
data "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}

output "subnet_id" {
  value = data.xcsh_subnet.example.id
}
```

<a id="canonical-2021033121103321-2231200312200230-0100202120322203-1212020112331303-2212110000310013-2321203122030201-3112321331023113-1123110331023332"></a>

## Root configuration — xcsh_subnet / 221220131001 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0202223002203121-3300302110230212-2210210021003200-1313323203130111-1013211022331033-1303003032023122-3312323201320203-3331301311121312"></a>

## Next pages — xcsh_subnet / 221220131001 / 6

- [Property reference](../guides/data-sources--subnet--reference--group-001.md#canonical-1102121302112331-0211122202223320-2203331110131003-2103332020023301-0010122012201110-0113221330320300-0120000001313021-1001213230013200)
- [Examples](../guides/data-sources--subnet--examples--group-001.md#canonical-1321312032030121-3122133013001130-0302032100203212-2123030011002321-3123332220300233-2113302010220023-2033110210311003-2002122223201211)
