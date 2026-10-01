---
page_title: "xcsh_bot_peer_threat_types landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_peer_threat_types landing."
---

# xcsh_bot_peer_threat_types landing

<a id="canonical-2013010020311013-0321123032202111-3000312123121121-3301311011300221-1213110110313232-0030311231210211-3033301232310333-0010002333121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131022002022233-2130203231131112-0321203132121212-3020021021112223-1110221321110202-2110320310010100-3211111200310110-0303212102010311"></a>

## xcsh_bot_peer_threat_types — xcsh_bot_peer_threat_types / 032132033312 / 2

Breadcrumbs:

- xcsh_bot_peer_threat_types

Resource creation operation.

<a id="canonical-2302123121312022-2301001213302311-0321121223321302-2311320221330100-3231232013033111-2331332212213121-3321322300001321-3232102311211020"></a>

## Prerequisites — xcsh_bot_peer_threat_types / 032132033312 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3200322122021332-0220330313313301-2313130302303322-3203011022123000-3033131131201003-3012333211001203-3310302010123223-3020030322012212"></a>

## Minimal configuration — xcsh_bot_peer_threat_types / 032132033312 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotPeerThreatTypes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_threat_types" "example" {
  namespace = "example-value"
}

output "bot_peer_threat_types_result" {
  value = data.xcsh_bot_peer_threat_types.example
}
```

<a id="canonical-3130010313200023-2001232232200312-0231303122311022-1120022323233022-1211223002122332-0012310212220212-0033132320233132-2022321232333023"></a>

## Root configuration — xcsh_bot_peer_threat_types / 032132033312 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0103321012300232-0311132030110322-1102113323320220-2321012020121120-1333301130103321-3113230012210302-2023112101032232-1232011211233003"></a>

## Next pages — xcsh_bot_peer_threat_types / 032132033312 / 6

- [Property reference](../guides/data-sources--bot_peer_threat_types--reference--group-001.md#canonical-3332100122232332-3333313223020231-0231003300003111-0203331223221022-1130001020133313-3333032112000312-2231220001303011-2221222230201332)
- [Examples](../guides/data-sources--bot_peer_threat_types--examples--group-001.md#canonical-1103132310203222-1120313020330113-0011321100311123-1333222112230011-3120013033102012-3030123330103203-1122100100321000-2210030210131031)
