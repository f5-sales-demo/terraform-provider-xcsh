---
page_title: "xcsh_forwarding_class landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class landing."
---

# xcsh_forwarding_class landing

<a id="canonical-0323010303031321-2222112303012320-2022001320033212-3211310231002211-0310131010121230-1201102203203131-0230233130001111-3201321012320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102101013132312-1231301101110330-1210033331200113-3310000220120210-1031313100033012-3130021102120032-3211012012102110-1222112202122211"></a>

## xcsh_forwarding_class — xcsh_forwarding_class / 300100323130 / 2

Breadcrumbs:

- xcsh_forwarding_class

Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users
in system namespace. configuration.

<a id="canonical-3030202001300330-1111023201012130-2303032311022011-0213023013022023-0121320331000213-1303020230033201-0133330230130233-3121331330331230"></a>

## Prerequisites — xcsh_forwarding_class / 300100323130 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1220133222123001-0130003221003000-2200111113003202-2201320120203132-3210233101330132-1203012221011321-3022123031031001-3111213020103332"></a>

## Minimal configuration — xcsh_forwarding_class / 300100323130 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ForwardingClass Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ForwardingClass by name
data "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}

output "forwarding_class_id" {
  value = data.xcsh_forwarding_class.example.id
}
```

<a id="canonical-1200132331232230-0001030213202110-1122033300011230-3322233201221213-2003331301232212-0222303011133310-2220213023132310-3330231320310201"></a>

## Root configuration — xcsh_forwarding_class / 300100323130 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1310133100220002-0103012230201330-1123313322300110-2232303102121000-2222010330312222-3002202130012103-1221131033110011-2330323011300222"></a>

## Next pages — xcsh_forwarding_class / 300100323130 / 6

- [Property reference](../guides/data-sources--forwarding_class--reference--group-001.md#canonical-1232301301333312-0233001311220232-0313221003110113-2220222123020101-3332223221102212-0221300232013330-0022330020133203-2230200222233213)
- [Examples](../guides/data-sources--forwarding_class--examples--group-001.md#canonical-1311111311300112-2102033310212222-0300103200221301-3230203003323211-2012012101313220-3103132200331032-1112102303131321-0100100130030332)
