---
page_title: "xcsh_healthcheck landing"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck landing."
---

# xcsh_healthcheck landing

<a id="canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212100002012132-3221101132220021-1033132332130320-2310121013033230-1331000300030333-3312320033333112-1332333330230232-2231213012112032"></a>

## xcsh_healthcheck — xcsh_healthcheck / 310333021331 / 2

Breadcrumbs:

- xcsh_healthcheck

Manages a Healthcheck resource in F5 Distributed Cloud for healthcheck object defines method to
determine if the given endpoint is healthy. single healthcheck object can be referred to by one or
many cluster objects. configuration.

<a id="canonical-1313311312013310-3311113122330021-0031303201113203-2100213010230313-1033101312111230-2020132123311302-0122012233012230-0123013302221130"></a>

## Prerequisites — xcsh_healthcheck / 310333021331 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1000122322013301-2203012021012303-0322323100201020-3022133332102021-0233113003203322-3123031013213012-1010310030300120-1010312003020230"></a>

## Minimal configuration — xcsh_healthcheck / 310333021331 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Healthcheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Healthcheck by name
data "xcsh_healthcheck" "example" {
  name      = "example-healthcheck"
  namespace = "staging"
}

output "healthcheck_id" {
  value = data.xcsh_healthcheck.example.id
}
```

<a id="canonical-0231323122322110-2212101113221131-1203313220033302-1031200030312100-3232323330021020-2111131023311332-2221100211311222-0133320311033130"></a>

## Root configuration — xcsh_healthcheck / 310333021331 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0330302330332112-2013023213123121-3102330220303213-0211222313033211-1202301233231012-2102223032103212-3233020113312112-3100000323310032"></a>

## Next pages — xcsh_healthcheck / 310333021331 / 6

- [Property reference](../guides/data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- [Examples](../guides/data-sources--healthcheck--examples--group-001.md#canonical-3300231312022010-3123302121132032-1100300103132332-3130120233103111-1022032131210033-2333031103310022-3211103132313203-3201230321020110)
