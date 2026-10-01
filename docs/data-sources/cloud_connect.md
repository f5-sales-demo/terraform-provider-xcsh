---
page_title: "xcsh_cloud_connect landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect landing."
---

# xcsh_cloud_connect landing

<a id="canonical-2011032310310033-3311323122002031-2213302012003111-3033233110012013-2021203231210000-3221202000202111-0001331133100330-3301210132312000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210031232332020-3103231002201323-1020221102132122-3302030311001003-1110221010312233-1322310321200123-3020310120223223-2130301111113101"></a>

## xcsh_cloud_connect — xcsh_cloud_connect / 012322212112 / 2

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

<a id="canonical-1112012320302232-2202210303210231-3223313211031023-2323102123023211-1312100030103200-3131203202101331-2321000013131323-2303131133223113"></a>

## Prerequisites — xcsh_cloud_connect / 012322212112 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1320102132011031-0130232331031020-2030310032310123-3212323200021300-2113323233322112-2223212220301123-2013230012023122-1302213130111313"></a>

## Minimal configuration — xcsh_cloud_connect / 012322212112 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudConnect Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudConnect by name
data "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}

output "cloud_connect_id" {
  value = data.xcsh_cloud_connect.example.id
}
```

<a id="canonical-1332030032221121-1013100122323303-2203001223000320-0101330032122100-0133011001221110-2320112313132220-0322033322013102-3321202311021121"></a>

## Root configuration — xcsh_cloud_connect / 012322212112 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1220310100232130-2032011003303212-3101300120022023-1220331211312131-0310312200322222-1020303221220203-0011221030322310-2311313332210111"></a>

## Next pages — xcsh_cloud_connect / 012322212112 / 6

- [Property reference](../guides/data-sources--cloud_connect--reference--group-001.md#canonical-0222032211111313-0003010022101312-0310023330122333-0201122213032311-0122213101110201-2002020111110210-1221320313132123-3300313332231200)
- [Examples](../guides/data-sources--cloud_connect--examples--group-001.md#canonical-0032101013232201-0133023230030221-0133202133332200-2113320033202103-3000331130231101-0120212200131012-2110110121212222-1031033003011033)
