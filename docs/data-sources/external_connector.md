---
page_title: "xcsh_external_connector landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector landing."
---

# xcsh_external_connector landing

<a id="canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332023133101310-2030330231121232-0013331211220020-2021001331323322-2323131220331023-3003111101120233-2323120322122312-1211322110221110"></a>

## xcsh_external_connector — xcsh_external_connector / 333011212021 / 2

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

<a id="canonical-1111320131101213-0231013201023220-2231032112332233-0030330022330201-2321322221230100-3100221303232021-2010202331312313-1120000302332302"></a>

## Prerequisites — xcsh_external_connector / 333011212021 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2313123032201102-1011222222131023-3322331112020323-2103013303322300-2112001003102032-1312020111130022-3113301223121100-0020221101223211"></a>

## Minimal configuration — xcsh_external_connector / 333011212021 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ExternalConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ExternalConnector by name
data "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}

output "external_connector_id" {
  value = data.xcsh_external_connector.example.id
}
```

<a id="canonical-3200112312111122-1113102011213032-1323310233031113-0122303201132013-1322120201120311-3220202312301131-1022322131020032-1213011303212023"></a>

## Root configuration — xcsh_external_connector / 333011212021 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2031321010123100-3331232203001103-0301022212330322-0012331022331211-0303301133131033-2032100131320003-0022121312103023-3312202031123022"></a>

## Next pages — xcsh_external_connector / 333011212021 / 6

- [Property reference](../guides/data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [Examples](../guides/data-sources--external_connector--examples--group-001.md#canonical-1123210132313011-1012021123322001-3100001220223320-0001132020312132-3132011113011320-0130333103100101-3033032111330123-0030221132300301)
