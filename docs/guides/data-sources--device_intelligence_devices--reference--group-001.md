---
page_title: "xcsh_device_intelligence_devices reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_devices reference."
---

# xcsh_device_intelligence_devices reference

<a id="canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-899315cb73860797fa7ae17f982b747dd67a7ba902df88f486d6cbf4bb3f2981"></a>

## Property reference — Property reference / 4a6c82e940ab / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- Property reference

<a id="canonical-2fbe52d0436c5b1cecded1d98eed17b2a161d6b332c806306d69e46fa3c4434f"></a>

## Direct properties — Property reference / 4a6c82e940ab / 3

- [devices](data-sources--device_intelligence_devices--reference--group-001.md#canonical-4b8960a4783fe0710db7ce7338a239369fa1c1a93d6e7f1b0230986f89984767): complete subsection reference.

<a id="canonical-f04df33985c05b176e0adfbbc339aa0ae2ba758fd2cce85f9b929bca6e3f9ecc"></a>

<a id="canonical-62172c68797d0f71fae0c5331ad888ad07b4c6a2706d4e983b5e8904648ad94a"></a>

## end_time property — Property reference / 4a6c82e940ab / 4

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-20f0057dcd2aa959befb04a898fda74a5ea95d3ada0d6e9ceadc9b1473625edc): complete subsection reference.

<a id="canonical-c10f84097c83b60a838ad2a937a6d4abef5b585cdefd3898a831ce1ed5a923c9"></a>

<a id="canonical-970a1ebea852630386137cc76e320d32016d2d1752bd777bb649838932b2e8a7"></a>

## namespace property — Property reference / 4a6c82e940ab / 5

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](data-sources--device_intelligence_devices--reference--group-001.md#canonical-fed58ede44e42711931b6d84fb4bfe4f077956fa969c6307770a119bf250d173): complete subsection reference.

- [sort](data-sources--device_intelligence_devices--reference--group-001.md#canonical-045e9d56ece2f375658737c8e12314e0ffd3bb34ac516200aed8d0bcf09e3f4a): complete subsection reference.

<a id="canonical-7ebe56501cc563d4abdd3454c22e929c2f095fe8a06603047760a62b77072853"></a>

<a id="canonical-1a2f115c6cd0b55c7fbf6572bc8f5c6c9afa8ab0443ba252a4de6f39c1988aed"></a>

## start_time property — Property reference / 4a6c82e940ab / 6

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-9bbd6122d7467442df2370153010ffbb6f6918d95df56b1defb16c782e3b28a8"></a>

## All schema paths — Property reference / 4a6c82e940ab / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `devices` | [devices](data-sources--device_intelligence_devices--reference--group-001.md#canonical-d904e611f62a0704236b85775ac265cdb82b862fc37fb8c69996ae181c132919) |
| `devices.action_taken` | [devices.action_taken](data-sources--device_intelligence_devices--reference--group-001.md#canonical-575c1da0616f7d2abac8fae16cf2e38a9c8f93459a7eb1dd1f3726af382a3705) |
| `devices.confidence` | [devices.confidence](data-sources--device_intelligence_devices--reference--group-001.md#canonical-a7d926211df46c6f1ca138db3fb23513eaa33ac7bb34540d6fb646c6208301dc) |
| `devices.device_id` | [devices.device_id](data-sources--device_intelligence_devices--reference--group-001.md#canonical-99def501dd4480edc695a9f5f93f5d51f960a0b8b3fceab3f09e1183f1a01c74) |
| `devices.high_risk_txn_count` | [devices.high_risk_txn_count](data-sources--device_intelligence_devices--reference--group-001.md#canonical-089f336dd5a7c7d3e1bbbd8faec2c9f0859202d81d0bc3d8827c5562a1f907db) |
| `devices.latest_txn_id` | [devices.latest_txn_id](data-sources--device_intelligence_devices--reference--group-001.md#canonical-61d4b9dc95d1b9ec92e677c5117d37f2a878eea90857ea8ee3f7a52831e73009) |
| `devices.linked_accounts` | [devices.linked_accounts](data-sources--device_intelligence_devices--reference--group-001.md#canonical-78bb2aead81a8515b823baed82e51b19d9d6c61d804b2d6b2e3a714a0fe58152) |
| `devices.risk_score` | [devices.risk_score](data-sources--device_intelligence_devices--reference--group-001.md#canonical-99172783876c28e0112b67b151eb4243d7115e9d98fd00a30ddb66fcdf258a9f) |
| `devices.risk_signals` | [devices.risk_signals](data-sources--device_intelligence_devices--reference--group-001.md#canonical-fef10b03f45c51c277d2a0447b5b6869b96fc2b1186302ef234d2fa48fa48ae8) |
| `end_time` | [end_time](data-sources--device_intelligence_devices--reference--group-001.md#canonical-f04df33985c05b176e0adfbbc339aa0ae2ba758fd2cce85f9b929bca6e3f9ecc) |
| `filters` | [filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-8bee034792b8d5e32ba2cd2b32771366edfc1050c3cf6a77eb9915ba97b0780c) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-739895a8e94a9033d5cc4fbfca6da7d66b42ae685fd6ca2611607d98187b6fb1) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_devices--reference--group-001.md#canonical-97d1e82431789db4da7e95d59f5683352c3c94836ad0f73419d35c418df29836) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_devices--reference--group-001.md#canonical-2cec43c385ac96a70da8bc6d9e8699757e6800cc0f93f049c03b6286d2579d17) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_devices--reference--group-001.md#canonical-0a2b7045564c2888d80efba7c3bbfa358f17358af747267d1daa570b428e48f9) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_devices--reference--group-001.md#canonical-627be91e8f3d149c34bfd4fe1453788d92a0878b5c288128632e6aa8d9cd1767) |
| `namespace` | [namespace](data-sources--device_intelligence_devices--reference--group-001.md#canonical-c10f84097c83b60a838ad2a937a6d4abef5b585cdefd3898a831ce1ed5a923c9) |
| `pagination` | [pagination](data-sources--device_intelligence_devices--reference--group-001.md#canonical-85b149e4c9152c90b35aeb4c77ef495d86dd80aba623b3fcbd47c57b956d1056) |
| `pagination.page_number` | [pagination.page_number](data-sources--device_intelligence_devices--reference--group-001.md#canonical-a7d3e4d9deccf58cc27535fb1f3e76d9fcc1734df1d176a6f524c6e2265c0776) |
| `pagination.page_size` | [pagination.page_size](data-sources--device_intelligence_devices--reference--group-001.md#canonical-4d4940a7ef1686ae28d680199ca05afe1cf27c9bb3a12944b6e752d73cc8d628) |
| `sort` | [sort](data-sources--device_intelligence_devices--reference--group-001.md#canonical-9f067831b91906e369b6a77d076cb95bf40c47420cb921423fc3228ddf389c2d) |
| `sort.key` | [sort.key](data-sources--device_intelligence_devices--reference--group-001.md#canonical-a1e940fe37bc268439397fe6c59f2480d8c74b98d798d79ef783504869f1c41d) |
| `sort.order` | [sort.order](data-sources--device_intelligence_devices--reference--group-001.md#canonical-26d69df62366c8731a0dc51c98f7564b43eb251aaaf1ddb81e8d01a3bc7b1b13) |
| `start_time` | [start_time](data-sources--device_intelligence_devices--reference--group-001.md#canonical-7ebe56501cc563d4abdd3454c22e929c2f095fe8a06603047760a62b77072853) |

<a id="canonical-eef6a44cc539d9fae264812662a8a44828e3d0352a4e952859f1ce7d1ade5666"></a>

## Next pages — Property reference / 4a6c82e940ab / 8

- [devices](data-sources--device_intelligence_devices--reference--group-001.md#canonical-4b8960a4783fe0710db7ce7338a239369fa1c1a93d6e7f1b0230986f89984767)
- [filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-20f0057dcd2aa959befb04a898fda74a5ea95d3ada0d6e9ceadc9b1473625edc)
- [pagination](data-sources--device_intelligence_devices--reference--group-001.md#canonical-fed58ede44e42711931b6d84fb4bfe4f077956fa969c6307770a119bf250d173)
- [sort](data-sources--device_intelligence_devices--reference--group-001.md#canonical-045e9d56ece2f375658737c8e12314e0ffd3bb34ac516200aed8d0bcf09e3f4a)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-4b8960a4783fe0710db7ce7338a239369fa1c1a93d6e7f1b0230986f89984767"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d67e8c00cf1f819137037c35f98429caef4cdd9b31c5d0fbd1e07cbfe6b3b102"></a>

## devices — devices / 90abb3a0346b / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- devices

<a id="canonical-d904e611f62a0704236b85775ac265cdb82b862fc37fb8c69996ae181c132919"></a>

Type: `"list"`. Computed.

Devices. List of devices for this page.

<a id="canonical-9c9a004bb0f9025a8bb608c592f5ba7d23a88425632d75fabcfa61cabd7e20ff"></a>

## Direct properties — devices / 90abb3a0346b / 3

<a id="canonical-575c1da0616f7d2abac8fae16cf2e38a9c8f93459a7eb1dd1f3726af382a3705"></a>

<a id="canonical-e36c1e85224fe3e9e9c352c95f04187574abb15a0b9bf52de285ccff6af7e31b"></a>

## action_taken property — devices / 90abb3a0346b / 4

Type: `"string"`. Computed.

Action taken or recommended based on the risk assessment.

<a id="canonical-a7d926211df46c6f1ca138db3fb23513eaa33ac7bb34540d6fb646c6208301dc"></a>

<a id="canonical-3bc74b991866a74711b1b7282c3761d487c54685d8745a4fca47e9bf6d5bebe8"></a>

## confidence property — devices / 90abb3a0346b / 5

Type: `"number"`. Computed.

Confidence level of the risk assessment (0–100).

<a id="canonical-99def501dd4480edc695a9f5f93f5d51f960a0b8b3fceab3f09e1183f1a01c74"></a>

<a id="canonical-75570c7bdee9d490c65e76e71211e0cb53f3941829e225c059d3bd40b04dd689"></a>

## device_id property — devices / 90abb3a0346b / 6

Type: `"string"`. Computed.

Device ID. Unique identifier for the device.

<a id="canonical-089f336dd5a7c7d3e1bbbd8faec2c9f0859202d81d0bc3d8827c5562a1f907db"></a>

<a id="canonical-7bbbed954c6026dc3a2894536de2d32539f34be020333f2c6128b45d65b4f475"></a>

## high_risk_txn_count property — devices / 90abb3a0346b / 7

Type: `"string"`. Computed.

Number of high-risk transactions linked to the device.

<a id="canonical-61d4b9dc95d1b9ec92e677c5117d37f2a878eea90857ea8ee3f7a52831e73009"></a>

<a id="canonical-64fdfb9aa89f0cf8d7db0f1f852c156fea008bea86152ba5daa5e79142007d08"></a>

## latest_txn_id property — devices / 90abb3a0346b / 8

Type: `"string"`. Computed.

Identifier of the most recent transaction associated with the device.

<a id="canonical-78bb2aead81a8515b823baed82e51b19d9d6c61d804b2d6b2e3a714a0fe58152"></a>

<a id="canonical-5573fe0c1dfb44457d9208818c3083df9450cb05e6effb62314b98b1581488e5"></a>

## linked_accounts property — devices / 90abb3a0346b / 9

Type: `"string"`. Computed.

Number of distinct accounts linked to this device.

<a id="canonical-99172783876c28e0112b67b151eb4243d7115e9d98fd00a30ddb66fcdf258a9f"></a>

<a id="canonical-cb95792801ec8b069b38b135d09644f2e05c64c8a4c4caaf37960b1f2ba8d8e8"></a>

## risk_score property — devices / 90abb3a0346b / 10

Type: `"number"`. Computed.

Overall risk score for the device (0–100).

<a id="canonical-fef10b03f45c51c277d2a0447b5b6869b96fc2b1186302ef234d2fa48fa48ae8"></a>

<a id="canonical-803199a605ddf1d50228536a5d287a8a5395524ae471573168b4105a6a106a54"></a>

## risk_signals property — devices / 90abb3a0346b / 11

Type: `["list", "string"]`. Computed.

List of risk signals detected for this device.

<a id="canonical-771d7a514fcb94ca86b2491d4528dca1ca920b88c599478ef182f8934693231d"></a>

## Next pages — devices / 90abb3a0346b / 12

- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-20f0057dcd2aa959befb04a898fda74a5ea95d3ada0d6e9ceadc9b1473625edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33a877b57491b59efda1ddc57a4901271136d4ae761155583fcc06b50cb1ccf1"></a>

## filters — filters / 2f30309e115d / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- filters

<a id="canonical-8bee034792b8d5e32ba2cd2b32771366edfc1050c3cf6a77eb9915ba97b0780c"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-e1dabb5a8cc96a00184c4b957f84b7a64732ac9cabe8ffc1a913cb7157219404"></a>

## Direct properties — filters / 2f30309e115d / 3

- [global_filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-a76af3ae5bbd5cd86b06f560acd272b1b9584a34ac85a723d42aae5ad76516d7): complete subsection reference.

<a id="canonical-627be91e8f3d149c34bfd4fe1453788d92a0878b5c288128632e6aa8d9cd1767"></a>

<a id="canonical-28938837f9ca8afff6ed32d43d334ac02617a677b738a22de213eab6e9289d5f"></a>

## region_filter property — filters / 2f30309e115d / 4

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

<a id="canonical-bd6a9f16042d476c1b05c7141f752a9f3d5fd2ce5315765580275c4623b1100b"></a>

## Next pages — filters / 2f30309e115d / 5

- [filters.global_filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-a76af3ae5bbd5cd86b06f560acd272b1b9584a34ac85a723d42aae5ad76516d7)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-a76af3ae5bbd5cd86b06f560acd272b1b9584a34ac85a723d42aae5ad76516d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f29ad9843ae1eb84a41ba579e930b5869add84d8b8c27e5781eafad58fcf842"></a>

## filters.global_filters — filters.global_filters / 4296d27da72a / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-20f0057dcd2aa959befb04a898fda74a5ea95d3ada0d6e9ceadc9b1473625edc)
- filters.global_filters

<a id="canonical-739895a8e94a9033d5cc4fbfca6da7d66b42ae685fd6ca2611607d98187b6fb1"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-b3fa28e90c2dd66b4057f9b4e843e3601f06efa4adc579634c7268e8b4af586e"></a>

## Direct properties — filters.global_filters / 4296d27da72a / 3

<a id="canonical-97d1e82431789db4da7e95d59f5683352c3c94836ad0f73419d35c418df29836"></a>

<a id="canonical-bc5565e1f17c986f7fafd4559b8b107b15c897ce6f841ddbebe00e509c5a6df6"></a>

## key property — filters.global_filters / 4296d27da72a / 4

Type: `"string"`. Optional.

\[Enum:
TIMESTAMP|USERNAME|CLIENT\_TOKEN|IP\_ADDRESS|ASN|AS\_ORGANIZATION|COUNTRY|METHOD|HOST|PATH|URL|REFERER|TRAFFIC\_CHANNEL|IS\_ATTACK|BOT\_REASON|TRAFFIC\_TYPE|THREAT\_TYPE|SDK\_VERSION|ACTION\_TAKEN|COOKIE\_AGE|BOT\_COOKIE|USER\_AGENT|USER\_AGENT\_OS\_FAMILY|USER\_AGENT\_FAMILY|BROWSER\_FINGERPRINT|USER\_FINGERPRINT|HEADER\_FINGERPRINT|DEVICE\_ID|FLOW|AGENT|APPLICATION\_NAME|PROTECTED\_APPLICATION|RESPONSE\_CODE|SERVER\_RESPONSE\_CODE|TRANSACTION\_RESULT|MOBILE\_TRANSACTION\_INSIGHT|WEB\_TRANSACTION\_INSIGHT|TRIGGERED\_RULE|FLOW\_CATEGORY|FLOW\_LABEL|ENDPOINT\_NAME|ENDPOINT\_LABEL|BOT\_ENDPOINT\_POLICY|KNOWN\_BOT\_NAME|KNOWN\_BOT\_CATEGORY|KNOWN\_BOT\_PROVIDER|KNOWN\_BOT\_CATEGORY\_TYPE|KNOWN\_BOT\_MITIGATION|ABSOLUTE|PERCENTAGE|TREND|ENDPOINT\_POLICY\]
Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible
values are \`TIMESTAMP\`, \`USERNAME\`, \`CLIENT\_TOKEN\`, \`IP\_ADDRESS\`, \`ASN\`,
\`AS\_ORGANIZATION\`, \`COUNTRY\`, \`METHOD\`, \`HOST\`, \`PATH\`, \`URL\`, \`REFERER\`,
\`TRAFFIC\_CHANNEL\`, \`IS\_ATTACK\`, \`BOT\_REASON\`, \`TRAFFIC\_TYPE\`, \`THREAT\_TYPE\`,
\`SDK\_VERSION\`, \`ACTION\_TAKEN\`, \`COOKIE\_AGE\`, \`BOT\_COOKIE\`, \`USER\_AGENT\`,
\`USER\_AGENT\_OS\_FAMILY\`, \`USER\_AGENT\_FAMILY\`, \`BROWSER\_FINGERPRINT\`,
\`USER\_FINGERPRINT\`, \`HEADER\_FINGERPRINT\`, \`DEVICE\_ID\`, \`FLOW\`, \`AGENT\`,
\`APPLICATION\_NAME\`, \`PROTECTED\_APPLICATION\`, \`RESPONSE\_CODE\`, \`SERVER\_RESPONSE\_CODE\`,
\`TRANSACTION\_RESULT\`, \`MOBILE\_TRANSACTION\_INSIGHT\`, \`WEB\_TRANSACTION\_INSIGHT\`,
\`TRIGGERED\_RULE\`, \`FLOW\_CATEGORY\`, \`FLOW\_LABEL\`, \`ENDPOINT\_NAME\`, \`ENDPOINT\_LABEL\`,
\`BOT\_ENDPOINT\_POLICY\`, \`KNOWN\_BOT\_NAME\`, \`KNOWN\_BOT\_CATEGORY\`, \`KNOWN\_BOT\_PROVIDER\`,
\`KNOWN\_BOT\_CATEGORY\_TYPE\`, \`KNOWN\_BOT\_MITIGATION\`, \`ABSOLUTE\`, \`PERCENTAGE\`, \`TREND\`,
\`ENDPOINT\_POLICY\`. Defaults to \`TIMESTAMP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TIMESTAMP",
    "USERNAME",
    "CLIENT_TOKEN",
    "IP_ADDRESS",
    "ASN",
    "AS_ORGANIZATION",
    "COUNTRY",
    "METHOD",
    "HOST",
    "PATH",
    "URL",
    "REFERER",
    "TRAFFIC_CHANNEL",
    "IS_ATTACK",
    "BOT_REASON",
    "TRAFFIC_TYPE",
    "THREAT_TYPE",
    "SDK_VERSION",
    "ACTION_TAKEN",
    "COOKIE_AGE",
    "BOT_COOKIE",
    "USER_AGENT",
    "USER_AGENT_OS_FAMILY",
    "USER_AGENT_FAMILY",
    "BROWSER_FINGERPRINT",
    "USER_FINGERPRINT",
    "HEADER_FINGERPRINT",
    "DEVICE_ID",
    "FLOW",
    "AGENT",
    "APPLICATION_NAME",
    "PROTECTED_APPLICATION",
    "RESPONSE_CODE",
    "SERVER_RESPONSE_CODE",
    "TRANSACTION_RESULT",
    "MOBILE_TRANSACTION_INSIGHT",
    "WEB_TRANSACTION_INSIGHT",
    "TRIGGERED_RULE",
    "FLOW_CATEGORY",
    "FLOW_LABEL",
    "ENDPOINT_NAME",
    "ENDPOINT_LABEL",
    "BOT_ENDPOINT_POLICY",
    "KNOWN_BOT_NAME",
    "KNOWN_BOT_CATEGORY",
    "KNOWN_BOT_PROVIDER",
    "KNOWN_BOT_CATEGORY_TYPE",
    "KNOWN_BOT_MITIGATION",
    "ABSOLUTE",
    "PERCENTAGE",
    "TREND",
    "ENDPOINT_POLICY"),
}
```

<a id="canonical-2cec43c385ac96a70da8bc6d9e8699757e6800cc0f93f049c03b6286d2579d17"></a>

<a id="canonical-ee9fa612331992b689575f951e6cf26e2f2eea789632c8063a1c8096d894b2b8"></a>

## op property — filters.global_filters / 4296d27da72a / 5

Type: `"string"`. Optional.

\[Enum:
IN|NOT\_IN|MATCHES\_REGEX|DOES\_NOT\_MATCH\_REGEX|INCLUDES|DOES\_NOT\_INCLUDE|STARTS\_WITH|ENDS\_WITH\]
Operator for query filter - IN: Filter Operator Specifies that query result includes filter values -
NOT\_IN: Filter Operator Specifies that query result excludes filter values - MATCHES\_REGEX: Filter
Operator Specifies that query result matches filter regex - DOES\_NOT\_MATCH\_REGEX: Filter..
Possible values are \`IN\`, \`NOT\_IN\`, \`MATCHES\_REGEX\`, \`DOES\_NOT\_MATCH\_REGEX\`,
\`INCLUDES\`, \`DOES\_NOT\_INCLUDE\`, \`STARTS\_WITH\`, \`ENDS\_WITH\`. Defaults to \`IN\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IN",
    "NOT_IN",
    "MATCHES_REGEX",
    "DOES_NOT_MATCH_REGEX",
    "INCLUDES",
    "DOES_NOT_INCLUDE",
    "STARTS_WITH",
    "ENDS_WITH"),
}
```

<a id="canonical-0a2b7045564c2888d80efba7c3bbfa358f17358af747267d1daa570b428e48f9"></a>

<a id="canonical-189dce8e79ca35923d5551beaec29445e267f16fc4fc910568134dff7042473f"></a>

## values property — filters.global_filters / 4296d27da72a / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-45d7947dddeee18da210666b58acdfb77f56a347f5628b5343747171514bf6a8"></a>

## Next pages — filters.global_filters / 4296d27da72a / 7

- [filters](data-sources--device_intelligence_devices--reference--group-001.md#canonical-20f0057dcd2aa959befb04a898fda74a5ea95d3ada0d6e9ceadc9b1473625edc)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-fed58ede44e42711931b6d84fb4bfe4f077956fa969c6307770a119bf250d173"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1732a87a8589de95f041aace9d932c681daf0a46956431caa3df1b0aff0e4e23"></a>

## pagination — pagination / 97ccee6bca34 / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- pagination

<a id="canonical-85b149e4c9152c90b35aeb4c77ef495d86dd80aba623b3fcbd47c57b956d1056"></a>

Type: `"single"`. Optional.

Pagination for Request with number and size.

<a id="canonical-d8f703d8d775c3267761068a9f7d9a2f67dc6e566b37597f1232a632ec50566c"></a>

## Direct properties — pagination / 97ccee6bca34 / 3

<a id="canonical-a7d3e4d9deccf58cc27535fb1f3e76d9fcc1734df1d176a6f524c6e2265c0776"></a>

<a id="canonical-800321007cd85bfa6f9c062aae30a59a975f9b130d64e7880b57a79544da96a4"></a>

## page_number property — pagination / 97ccee6bca34 / 4

Type: `"number"`. Optional.

Configuration parameter for page number.

<a id="canonical-4d4940a7ef1686ae28d680199ca05afe1cf27c9bb3a12944b6e752d73cc8d628"></a>

<a id="canonical-d89d9c2f1bdeb37a8602f601424b27d658a5efbab8bae9e553c4b86511ed61d1"></a>

## page_size property — pagination / 97ccee6bca34 / 5

Type: `"number"`. Optional.

Page Size. Size or capacity specification

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-b148331e9e097d01349fee0a3e27c4bb85cfdb2a33f2bedf0de8fed741c9c178"></a>

## Next pages — pagination / 97ccee6bca34 / 6

- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)

<a id="canonical-045e9d56ece2f375658737c8e12314e0ffd3bb34ac516200aed8d0bcf09e3f4a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83953625268a05e9e230e1f1b923baa6a2889041c3fe5642122d24cf97aed9f9"></a>

## sort — sort / 32fe1d2c5a06 / 2

Breadcrumbs:

- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- sort

<a id="canonical-9f067831b91906e369b6a77d076cb95bf40c47420cb921423fc3228ddf389c2d"></a>

Type: `"single"`. Optional.

Sort Option. Query Result Sort Option.

<a id="canonical-fca0bf17a63beded314236bc20a1db453f5bf563768bbf268d2e68dbcbb355f2"></a>

## Direct properties — sort / 32fe1d2c5a06 / 3

<a id="canonical-a1e940fe37bc268439397fe6c59f2480d8c74b98d798d79ef783504869f1c41d"></a>

<a id="canonical-fb657c43be7bafed11057ca57db4850238aeccdf0e2443b544c1f59d48e39106"></a>

## key property — sort / 32fe1d2c5a06 / 4

Type: `"string"`. Optional.

\[Enum:
TIMESTAMP|USERNAME|CLIENT\_TOKEN|IP\_ADDRESS|ASN|AS\_ORGANIZATION|COUNTRY|METHOD|HOST|PATH|URL|REFERER|TRAFFIC\_CHANNEL|IS\_ATTACK|BOT\_REASON|TRAFFIC\_TYPE|THREAT\_TYPE|SDK\_VERSION|ACTION\_TAKEN|COOKIE\_AGE|BOT\_COOKIE|USER\_AGENT|USER\_AGENT\_OS\_FAMILY|USER\_AGENT\_FAMILY|BROWSER\_FINGERPRINT|USER\_FINGERPRINT|HEADER\_FINGERPRINT|DEVICE\_ID|FLOW|AGENT|APPLICATION\_NAME|PROTECTED\_APPLICATION|RESPONSE\_CODE|SERVER\_RESPONSE\_CODE|TRANSACTION\_RESULT|MOBILE\_TRANSACTION\_INSIGHT|WEB\_TRANSACTION\_INSIGHT|TRIGGERED\_RULE|FLOW\_CATEGORY|FLOW\_LABEL|ENDPOINT\_NAME|ENDPOINT\_LABEL|BOT\_ENDPOINT\_POLICY|KNOWN\_BOT\_NAME|KNOWN\_BOT\_CATEGORY|KNOWN\_BOT\_PROVIDER|KNOWN\_BOT\_CATEGORY\_TYPE|KNOWN\_BOT\_MITIGATION|ABSOLUTE|PERCENTAGE|TREND|ENDPOINT\_POLICY\]
Key for query filter - TIMESTAMP: Timestamp Filter Key Use Timestamp as key to query. Possible
values are \`TIMESTAMP\`, \`USERNAME\`, \`CLIENT\_TOKEN\`, \`IP\_ADDRESS\`, \`ASN\`,
\`AS\_ORGANIZATION\`, \`COUNTRY\`, \`METHOD\`, \`HOST\`, \`PATH\`, \`URL\`, \`REFERER\`,
\`TRAFFIC\_CHANNEL\`, \`IS\_ATTACK\`, \`BOT\_REASON\`, \`TRAFFIC\_TYPE\`, \`THREAT\_TYPE\`,
\`SDK\_VERSION\`, \`ACTION\_TAKEN\`, \`COOKIE\_AGE\`, \`BOT\_COOKIE\`, \`USER\_AGENT\`,
\`USER\_AGENT\_OS\_FAMILY\`, \`USER\_AGENT\_FAMILY\`, \`BROWSER\_FINGERPRINT\`,
\`USER\_FINGERPRINT\`, \`HEADER\_FINGERPRINT\`, \`DEVICE\_ID\`, \`FLOW\`, \`AGENT\`,
\`APPLICATION\_NAME\`, \`PROTECTED\_APPLICATION\`, \`RESPONSE\_CODE\`, \`SERVER\_RESPONSE\_CODE\`,
\`TRANSACTION\_RESULT\`, \`MOBILE\_TRANSACTION\_INSIGHT\`, \`WEB\_TRANSACTION\_INSIGHT\`,
\`TRIGGERED\_RULE\`, \`FLOW\_CATEGORY\`, \`FLOW\_LABEL\`, \`ENDPOINT\_NAME\`, \`ENDPOINT\_LABEL\`,
\`BOT\_ENDPOINT\_POLICY\`, \`KNOWN\_BOT\_NAME\`, \`KNOWN\_BOT\_CATEGORY\`, \`KNOWN\_BOT\_PROVIDER\`,
\`KNOWN\_BOT\_CATEGORY\_TYPE\`, \`KNOWN\_BOT\_MITIGATION\`, \`ABSOLUTE\`, \`PERCENTAGE\`, \`TREND\`,
\`ENDPOINT\_POLICY\`. Defaults to \`TIMESTAMP\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TIMESTAMP",
    "USERNAME",
    "CLIENT_TOKEN",
    "IP_ADDRESS",
    "ASN",
    "AS_ORGANIZATION",
    "COUNTRY",
    "METHOD",
    "HOST",
    "PATH",
    "URL",
    "REFERER",
    "TRAFFIC_CHANNEL",
    "IS_ATTACK",
    "BOT_REASON",
    "TRAFFIC_TYPE",
    "THREAT_TYPE",
    "SDK_VERSION",
    "ACTION_TAKEN",
    "COOKIE_AGE",
    "BOT_COOKIE",
    "USER_AGENT",
    "USER_AGENT_OS_FAMILY",
    "USER_AGENT_FAMILY",
    "BROWSER_FINGERPRINT",
    "USER_FINGERPRINT",
    "HEADER_FINGERPRINT",
    "DEVICE_ID",
    "FLOW",
    "AGENT",
    "APPLICATION_NAME",
    "PROTECTED_APPLICATION",
    "RESPONSE_CODE",
    "SERVER_RESPONSE_CODE",
    "TRANSACTION_RESULT",
    "MOBILE_TRANSACTION_INSIGHT",
    "WEB_TRANSACTION_INSIGHT",
    "TRIGGERED_RULE",
    "FLOW_CATEGORY",
    "FLOW_LABEL",
    "ENDPOINT_NAME",
    "ENDPOINT_LABEL",
    "BOT_ENDPOINT_POLICY",
    "KNOWN_BOT_NAME",
    "KNOWN_BOT_CATEGORY",
    "KNOWN_BOT_PROVIDER",
    "KNOWN_BOT_CATEGORY_TYPE",
    "KNOWN_BOT_MITIGATION",
    "ABSOLUTE",
    "PERCENTAGE",
    "TREND",
    "ENDPOINT_POLICY"),
}
```

<a id="canonical-26d69df62366c8731a0dc51c98f7564b43eb251aaaf1ddb81e8d01a3bc7b1b13"></a>

<a id="canonical-2d03f7c0e217615793ea741364f0769bd6971de62c871e9a45faaaf39b876f66"></a>

## order property — sort / 32fe1d2c5a06 / 5

Type: `"string"`. Optional.

\[Enum: DESCENDING|ASCENDING\] Sort algorithm Sort in descending order Sort in ascending order.
Possible values are \`DESCENDING\`, \`ASCENDING\`. Defaults to \`DESCENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DESCENDING",
    "ASCENDING"),
}
```

<a id="canonical-8b9e8fa68bcc47c1a36ef6ed701ffe549b054f4b9aea304f49cf8a58e4eeace4"></a>

## Next pages — sort / 32fe1d2c5a06 / 6

- [Property reference](data-sources--device_intelligence_devices--reference--group-001.md#canonical-44a62deff463fb871382d9c68671f83b580d2862aa9293e18cb445f06d1be873)
- [xcsh_device_intelligence_devices](../data-sources/device_intelligence_devices.md#canonical-7911c305a3aa551402a0597c54c40adaa717422bb3cfe729a620174f3c16433f)
