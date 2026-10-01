---
page_title: "xcsh_app_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type landing."
---

# xcsh_app_type landing

<a id="canonical-2200113311021331-0200020232011230-0011223212102321-1001021222113033-0011022023131301-0311022330302220-2111320322332313-1010123013021130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102123032201233-2223222321210332-2311212203232333-3113011033231101-3320311130221120-2312301103021132-0111120311123133-2103012031232312"></a>

## xcsh_app_type — xcsh_app_type / 300313323222 / 2

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-3313231032223120-0003002121013132-0032202223111211-2033010322003213-3011122331310100-1131032332202303-1002313022202302-0312210200200313"></a>

## Prerequisites — xcsh_app_type / 300313323222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1230121002221123-0203121223022010-1000030232003201-0122210112130203-0322111301032200-1113303023002233-0223320120022230-3112303310203213"></a>

## Minimal configuration — xcsh_app_type / 300313323222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```

<a id="canonical-1022032333332200-2303311022323001-1210110202330230-1320003231133021-1220223103010102-3131020112222331-3121310012323311-0102002313230023"></a>

## Root configuration — xcsh_app_type / 300313323222 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3021003031112232-0100031113331232-1322013320023131-2023220122010121-2011130330301323-2203001032310310-1211121001230200-2032233330133230"></a>

## Next pages — xcsh_app_type / 300313323222 / 6

- [Property reference](../guides/data-sources--app_type--reference--group-001.md#canonical-3221103121003133-1230232300121110-2122110130133103-3202001113101201-3311021023133001-3112111111202133-2003201302132001-0322123211030221)
- [Examples](../guides/data-sources--app_type--examples--group-001.md#canonical-2102211203220302-1231033312102010-0032320200033323-2012000223133230-3112033310312301-3331003232031021-0230112012023120-3302011322211012)
