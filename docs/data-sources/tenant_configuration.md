---
page_title: "xcsh_tenant_configuration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration landing."
---

# xcsh_tenant_configuration landing

<a id="canonical-1230121131133133-0211313013212213-1101132331100023-2203020001230102-0301100132002332-3002232111300122-2302311100323013-2003322300110201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213221223120313-2231030210223213-0030220021000032-2223113331120330-1103111003100103-0332003223101020-2033001213002302-1020111030032100"></a>

## xcsh_tenant_configuration — xcsh_tenant_configuration / 300103332202 / 2

Breadcrumbs:

- xcsh_tenant_configuration

Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration
specification. configuration.

<a id="canonical-3200322121012330-1111220133201100-0230201302110212-3103130333212131-2201312130021023-3132101222230102-0331022111320013-1002321013230203"></a>

## Prerequisites — xcsh_tenant_configuration / 300103332202 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1302133000310323-0110030121221011-3300130231021101-2221021311022133-1211212233220211-1021330330323301-2021133022311323-3230021231212303"></a>

## Minimal configuration — xcsh_tenant_configuration / 300103332202 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TenantConfiguration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TenantConfiguration by name
data "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}

output "tenant_configuration_id" {
  value = data.xcsh_tenant_configuration.example.id
}
```

<a id="canonical-2121022123102030-1231022000301030-1131220222013323-2100213133211112-0133331211310230-2131212331301103-0331101213233302-1111310003132031"></a>

## Root configuration — xcsh_tenant_configuration / 300103332202 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1203232131313020-0123003220213133-1210113211133010-3301103202130211-2003301122022223-3020113312013202-1323103033000332-3100101202111103"></a>

## Next pages — xcsh_tenant_configuration / 300103332202 / 6

- [Property reference](../guides/data-sources--tenant_configuration--reference--group-001.md#canonical-3122132313001203-3223033303101001-1300213201000103-1022102223020331-0330111321322330-2333230000221030-2133000302200313-2230013023120232)
- [Examples](../guides/data-sources--tenant_configuration--examples--group-001.md#canonical-3233212102312300-1110101220100003-2133212022310023-0100210120032012-3222033131310000-3321000113321031-3333220321102300-3232122211111221)
