---
page_title: "xcsh_app_api_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group landing."
---

# xcsh_app_api_group landing

<a id="canonical-3011310102212300-2112210220330313-1330322333011033-1111313021020310-3121110131333221-1200033222203013-2121123032213301-2023132013121010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321311200031123-0202120131331110-3021322213313110-1111323021132032-3211032002320200-2203130000331002-0211131023032113-2200233223230220"></a>

## xcsh_app_api_group — xcsh_app_api_group / 011312131012 / 2

Breadcrumbs:

- xcsh_app_api_group

Manages app\_api\_group creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-1122100210113132-3322200312320212-2211332001332131-0220030100121013-1101113323313211-3113201020201302-2021003332311223-0332333002300002"></a>

## Prerequisites — xcsh_app_api_group / 011312131012 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2203220103313312-3101211011333332-3322003321300122-0022102121131011-0313021323312113-1310121032012233-2021020310123103-1220231002030021"></a>

## Minimal configuration — xcsh_app_api_group / 011312131012 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

<a id="canonical-1102303303022232-3302110132023122-1213313110021120-1232220010330121-2333221130212130-0100220211323110-1213013000303210-3211031212110310"></a>

## Root configuration — xcsh_app_api_group / 011312131012 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2310001331000313-1130332232222120-0333323131323220-1320003223233111-3323222313111300-2030101303101133-1022121213123121-2320303023000330"></a>

## Next pages — xcsh_app_api_group / 011312131012 / 6

- [Property reference](../guides/data-sources--app_api_group--reference--group-001.md#canonical-3131003230120102-0133232130312100-3201332223003222-1312202223000112-1310030333311011-3211213212221220-1032213012023303-0012130201000232)
- [Examples](../guides/data-sources--app_api_group--examples--group-001.md#canonical-0302100133100312-2121120001033302-0022331022132312-3310031221011031-0332100330131023-1202303321310212-2322110002012103-3331001030301032)
