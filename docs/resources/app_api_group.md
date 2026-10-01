---
page_title: "xcsh_app_api_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group landing."
---

# xcsh_app_api_group landing

<a id="canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111312000220310-3123120322333303-1311123101033310-3112211301003120-1033300223000211-2112231210010312-3111223322121012-0013312301321100"></a>

## xcsh_app_api_group — xcsh_app_api_group / 002121320222 / 2

Breadcrumbs:

- xcsh_app_api_group

Manages app\_api\_group creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-3010210013100232-2013320230201302-1002012303000131-3223220002101220-0122113231311103-1102233000302031-0110131222130210-0201031012110302"></a>

## Prerequisites — xcsh_app_api_group / 002121320222 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3123002103202000-1312130112010331-0023221020103300-2020130123211232-1131031130233101-0210113110033130-0133003121332100-3121310300112330"></a>

## Minimal configuration — xcsh_app_api_group / 002121320222 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```

<a id="canonical-2200323232021330-3223011121321331-0020121102301312-3220023230110200-2103330331102010-0120100313333230-0201320202032233-2211020123123202"></a>

## Root configuration — xcsh_app_api_group / 002121320222 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0323323222201210-2111003102132322-1301321210201102-0233231022103133-3202203000231120-1223021332122230-0220020130033021-0222132112311201"></a>

## Next pages — xcsh_app_api_group / 002121320222 / 6

- [Property reference](../guides/resources--app_api_group--reference--group-001.md#canonical-3122032033302330-3011112322022020-2220111323132101-2203222330000213-0220213001213111-2030332331013100-0313223112302312-1031233020112010)
- [Examples](../guides/resources--app_api_group--examples--group-001.md#canonical-1302312103002210-1032122003313033-3133000013320313-2200203030220230-3200103023322120-0201330103131101-3133332311012302-2311001022312100)
- [Import](../guides/resources--app_api_group--lifecycle--group-001.md#canonical-1312021132133132-2110021001112001-2121033312332230-1210332101002022-3120300010321200-3110132001313132-3202220203020233-0312133130000223)
- [Timeouts](../guides/resources--app_api_group--lifecycle--group-001.md#canonical-2102332330303133-0211121212310132-0113210012110210-3031111232001230-1122103323103131-1333031131233330-1001000120103020-2132221111223103)
