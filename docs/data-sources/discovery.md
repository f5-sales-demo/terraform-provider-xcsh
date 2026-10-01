---
page_title: "xcsh_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery landing."
---

# xcsh_discovery landing

<a id="canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201222103332033-3232230013101103-2032021330231312-0012333231120031-3221133130113230-1203122313320003-2032021100302330-2111323211013111"></a>

## xcsh_discovery — xcsh_discovery / 032203223120 / 2

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

<a id="canonical-1021231233013031-3100001112310322-3223000011122312-1320213031132222-0223321132223320-0123323133233002-2232110322121312-2002101022203122"></a>

## Prerequisites — xcsh_discovery / 032203223120 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2303122133002000-0332112303022123-3313130200301203-2332001031121101-0322330203301113-3032300001112330-0332232203210103-1131330011231310"></a>

## Minimal configuration — xcsh_discovery / 032203223120 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

<a id="canonical-1311013220200103-1203021022121022-0233310030210311-1120221132012133-0031133111331032-2210300210210213-1100233211223212-1121023100033211"></a>

## Root configuration — xcsh_discovery / 032203223120 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0031201310211213-1310210313310233-0121122312232221-2200020211121220-1011111101001001-2230211331013202-0030121032033010-1302013111210303"></a>

## Next pages — xcsh_discovery / 032203223120 / 6

- [Property reference](../guides/data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [Examples](../guides/data-sources--discovery--examples--group-001.md#canonical-2233113310202003-3011212032322120-0321110101103330-1110220301021333-2113230300122223-0032111232321131-3031100223030032-3233311203131303)
