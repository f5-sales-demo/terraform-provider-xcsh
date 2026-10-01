---
page_title: "xcsh_cminstance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance landing."
---

# xcsh_cminstance landing

<a id="canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021232203221320-2123233021122002-2231110001030313-2202121112300001-1000300122320130-2010023313200212-2211301001031023-1130022023230101"></a>

## xcsh_cminstance — xcsh_cminstance / 203003131223 / 2

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-2232300321033101-2332111011323113-1310103021002032-1033303220121102-2313101112313322-2100131023113222-1212120112223210-1132031011021321"></a>

## Prerequisites — xcsh_cminstance / 203003131223 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2211321021011010-2200111032321031-0313210022231331-0200322303123123-0112102102111230-3023031321301010-2030010331300020-2013112333101322"></a>

## Minimal configuration — xcsh_cminstance / 203003131223 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

<a id="canonical-0321031001211123-0002210323232303-0332203021302202-3331213100112021-3131203023103330-3120323201021132-1220220020203221-3030132021033301"></a>

## Root configuration — xcsh_cminstance / 203003131223 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1112023331100013-2132100123332303-0213132222020030-1220122330110102-3301230301210213-3020232210331311-1322220102311010-0130031102221003"></a>

## Next pages — xcsh_cminstance / 203003131223 / 6

- [Property reference](../guides/data-sources--cminstance--reference--group-001.md#canonical-3021233103221211-3000330212303311-3123011101210133-2233111111213112-2200323123322113-3330123132032201-0000230310101123-2230201321203202)
- [Examples](../guides/data-sources--cminstance--examples--group-001.md#canonical-2213332301233130-3302331022010303-2123033122303330-2001220110232311-3121111202200321-3232321201300311-0122300113032220-2021013313103203)
