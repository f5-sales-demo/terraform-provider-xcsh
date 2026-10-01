---
page_title: "xcsh_api_definition landing"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition landing."
---

# xcsh_api_definition landing

<a id="canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202120233131113-3113032201100112-0201001233000122-3002312133013101-3132032331122001-3312310310111130-0101321120202302-3323320331123333"></a>

## xcsh_api_definition — xcsh_api_definition / 122202321013 / 2

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

<a id="canonical-3332003320003211-0303032223021013-0330031013331322-0313331113020130-1012111201320023-2111222302320103-0202022333231320-0003123111133133"></a>

## Prerequisites — xcsh_api_definition / 122202321013 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

<a id="canonical-0311211220231022-3103333321232310-0302112012010010-1233100312303211-0013301120131323-0110032301031331-1301320023101013-0301232231023300"></a>

## Minimal configuration — xcsh_api_definition / 122202321013 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```

<a id="canonical-1103132202133312-1230210321111013-3313220321332223-2232200002030133-1221201010131033-2020133200032312-0321323312133120-1312230221211012"></a>

## Root configuration — xcsh_api_definition / 122202321013 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0202020132303211-0222101212012033-0321102000302203-0100031033210001-3321011103321321-2201323003201333-1300330001312122-0202330232202210"></a>

## Next pages — xcsh_api_definition / 122202321013 / 6

- [Property reference](../guides/data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- [Examples](../guides/data-sources--api_definition--examples--group-001.md#canonical-1033121102020132-2320020110333032-2102300020133012-3123220022131120-3112301032313011-1100003031333322-3012330310222003-3302320212221331)
