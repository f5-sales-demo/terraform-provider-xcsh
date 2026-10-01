---
page_title: "xcsh_certificate landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate landing."
---

# xcsh_certificate landing

<a id="canonical-3012321000133203-2232232100110030-1133220132311111-1220230030331111-1113103323102333-2233030332231013-0322121312100331-0121101033322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011232203100103-1101001221333120-3001320122333100-0000101011003020-3221130300231333-0100112203132232-3113201000022320-0121322013210221"></a>

## xcsh_certificate — xcsh_certificate / 033223030013 / 2

Breadcrumbs:

- xcsh_certificate

Manages a Certificate resource in F5 Distributed Cloud for certificate. configuration.

<a id="canonical-1003131013010002-3202230222231311-0321000122012222-1003333100331323-2312030231100021-0210330232113310-1200320312012232-3333012331001200"></a>

## Prerequisites — xcsh_certificate / 033223030013 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3321122122111021-3322311121103031-1021102202100312-2010030000322031-3221001001013020-1210020013121232-0101133002331033-1100011231232022"></a>

## Minimal configuration — xcsh_certificate / 033223030013 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```

<a id="canonical-1311333322332213-3203111223303312-1211013130302212-1230002230322203-1321020303131312-2101113231221013-0120202013321033-2100020102002012"></a>

## Root configuration — xcsh_certificate / 033223030013 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0313110030200023-2010111121022000-0220133000300212-0120122030131332-1032012123323232-1313333210132231-1301020103322000-3133332311103333"></a>

## Next pages — xcsh_certificate / 033223030013 / 6

- [Property reference](../guides/data-sources--certificate--reference--group-001.md#canonical-0231301231300033-0330331110133021-2220121332302031-1112310222301003-2132000233121232-0233011201000301-3230232011333001-1002232100313000)
- [Examples](../guides/data-sources--certificate--examples--group-001.md#canonical-0332313110100301-2123222322221030-3221312133000120-2333200122131330-1111030233012332-3313211030323021-0130210110332303-0312311102101303)
