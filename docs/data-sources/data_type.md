---
page_title: "xcsh_data_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type landing."
---

# xcsh_data_type landing

<a id="canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302111220301303-2212010233031020-0202111101303203-1011120010110032-0010102123300123-1331323100120222-0031120130002302-0012332133313200"></a>

## xcsh_data_type — xcsh_data_type / 332113221001 / 2

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-1122200221302302-2020102211322110-0312203111022230-0021212022102030-3020001301102010-1030131112020003-1001213002210212-1103302111003232"></a>

## Prerequisites — xcsh_data_type / 332113221001 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1110332113023020-2301231002122110-0201112321331232-1210203222120033-3310321221220001-3000120210133022-0133122100000302-0300023000023222"></a>

## Minimal configuration — xcsh_data_type / 332113221001 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataType by name
data "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}

output "data_type_id" {
  value = data.xcsh_data_type.example.id
}
```

<a id="canonical-1232030002302223-0031330112212312-2201112310232012-1111100010213030-0031212320122212-0220132312312222-3221323201213221-1221202202100313"></a>

## Root configuration — xcsh_data_type / 332113221001 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2013233323030200-3110131300121101-3001032300320131-0011233223300323-0220113002213220-2200300212210310-3210003022330210-2311211331120323"></a>

## Next pages — xcsh_data_type / 332113221001 / 6

- [Property reference](../guides/data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [Examples](../guides/data-sources--data_type--examples--group-001.md#canonical-0112113123333023-1303013303310031-2312012033233221-3103003032013203-3101002222312000-3233133102021111-3112213300131103-2110000000130202)
