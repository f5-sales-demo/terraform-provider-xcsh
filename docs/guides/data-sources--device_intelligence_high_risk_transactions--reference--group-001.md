---
page_title: "xcsh_device_intelligence_high_risk_transactions reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions reference."
---

# xcsh_device_intelligence_high_risk_transactions reference

<a id="canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73265dfaf4d4e10431078e9a6442fc8f78fe6aa5db71bf37aac147c1045633a7"></a>

## Property reference — Property reference / 27c3316cd3c2 / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- Property reference

<a id="canonical-2cc3df10fb0dd2cef74c16c464dff0c6a204e09e158391033a86c7f049cc181c"></a>

## Direct properties — Property reference / 27c3316cd3c2 / 3

<a id="canonical-5baf0cdb3f0960478dcb9260ec1808c0119383797fbca44a28d6f49cd75c9114"></a>

<a id="canonical-265ffa3984a167866d6e51c8487bfe15f7bbb99a45e92b35e4f741b030aa1cf7"></a>

## end_time property — Property reference / 27c3316cd3c2 / 4

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

- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-18d5387b9970901e50b9dc77f44f9cdc311f7cbfc6629564b84e6cecf8eadd05): complete subsection reference.

<a id="canonical-d551c6918e4150df45a558f738eaa37291f83c60620b5f79157eac9db5646825"></a>

<a id="canonical-cc193b70e1d77047800c930144ab1c1fc527c2f992aaf81810a44f4d1e86c8fc"></a>

## namespace property — Property reference / 27c3316cd3c2 / 5

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-51a6ff7c888581a2982ec53dd80073267e6c8f63b79f6341de3be1d0a78ad441"></a>

<a id="canonical-328ec09dd0c47e17c867516f57a74a8ec5beef9186fb4fdc4d50605daf9a04ca"></a>

## start_time property — Property reference / 27c3316cd3c2 / 6

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

- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-72c335c60953232fff4a455e778bcd6ee8b31429dd91f6462bf4b96331cc3f01): complete subsection reference.

<a id="canonical-06f7bc2e5d76ae00af517adf158e6c0a1fc8a7ee8b589375733883b3f40b480c"></a>

## All schema paths — Property reference / 27c3316cd3c2 / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-5baf0cdb3f0960478dcb9260ec1808c0119383797fbca44a28d6f49cd75c9114) |
| `filters` | [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-cadb2717d25afbdf2960d536f329dd9b853e7cdf87a8873ee8e779067f6393c4) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-f0a83cb39fe1f58b449c8b87fea241a4bc4392da0019b999813596d36ff5f184) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-310dca1b07d275558afb2cee5ba81ab50e5a7148bb62056eb4165dd8c1171171) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-e0da3eaf33156a39a2b2c26214ab7a26a5ecdfff20338fad195c9e02df5c13a0) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-a99fac01173a8399c6a851f67720601f5641625a5980a2269ecc0c99501cb226) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-f9f0097967841823b8761a643f51824eac88f94f2aaf0149fcc9b51dab4530e4) |
| `namespace` | [namespace](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-d551c6918e4150df45a558f738eaa37291f83c60620b5f79157eac9db5646825) |
| `start_time` | [start_time](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-51a6ff7c888581a2982ec53dd80073267e6c8f63b79f6341de3be1d0a78ad441) |
| `time_series_results` | [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8793aa9d5d8cce4e1290f5d6873c47a83ee19d9ff89445be543c4049cab0cf2a) |
| `time_series_results.series_key` | [time_series_results.series_key](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-a949781fd7e2f1f4d71d6d546d69b2c2a79d0770178d44bd9025961ba54a36bf) |
| `time_series_results.time_series` | [time_series_results.time_series](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-9a6805a59decdf324418f98a7cf588d62700081557484ea4c9db97a4a7fef52a) |
| `time_series_results.time_series.timestamp` | [time_series_results.time_series.timestamp](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-b5f00d3c5b10c7872fc57b828f7ce468bb86e3f69238e49b5ae0fb2f1e2bf1ca) |
| `time_series_results.time_series.value` | [time_series_results.time_series.value](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-76dc60be9c9d66f39b3d20e0e2f41e025d50b994d68a61c55859bcf01d733589) |

<a id="canonical-f38d22f90f6817325259396e926eeb01f52578e00ff055d026209aa2603dd1f8"></a>

## Next pages — Property reference / 27c3316cd3c2 / 8

- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-18d5387b9970901e50b9dc77f44f9cdc311f7cbfc6629564b84e6cecf8eadd05)
- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-72c335c60953232fff4a455e778bcd6ee8b31429dd91f6462bf4b96331cc3f01)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)

<a id="canonical-18d5387b9970901e50b9dc77f44f9cdc311f7cbfc6629564b84e6cecf8eadd05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d496509dda84f6a3882b424760fa9733639b9eba9681b321439e2e14219c8853"></a>

## filters — filters / 0dd782519570 / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- filters

<a id="canonical-cadb2717d25afbdf2960d536f329dd9b853e7cdf87a8873ee8e779067f6393c4"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-43125bb0f811fb42b40a4f6b059369ef2d3a77c9d059b0236493ff2319b0f853"></a>

## Direct properties — filters / 0dd782519570 / 3

- [global_filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-4d5284ba5e2bd010b945b1f9f09230a2cd3c40e0810002533b7004ab47f0f3af): complete subsection reference.

<a id="canonical-f9f0097967841823b8761a643f51824eac88f94f2aaf0149fcc9b51dab4530e4"></a>

<a id="canonical-080e2f04c045352a2351984eed7dd9c930317b2ef18cf6c165c5c2a07e0036d1"></a>

## region_filter property — filters / 0dd782519570 / 4

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

<a id="canonical-50df5f2e3e90370bae63712f3208feca0109183ab062b54da23c5a6126e3e83b"></a>

## Next pages — filters / 0dd782519570 / 5

- [filters.global_filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-4d5284ba5e2bd010b945b1f9f09230a2cd3c40e0810002533b7004ab47f0f3af)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)

<a id="canonical-4d5284ba5e2bd010b945b1f9f09230a2cd3c40e0810002533b7004ab47f0f3af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ead5c6cfe025b30f3b73e8318b42967eafe35152c065fcd70141c43d28a2a5"></a>

## filters.global_filters — filters.global_filters / 2d8a495a4bd4 / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-18d5387b9970901e50b9dc77f44f9cdc311f7cbfc6629564b84e6cecf8eadd05)
- filters.global_filters

<a id="canonical-f0a83cb39fe1f58b449c8b87fea241a4bc4392da0019b999813596d36ff5f184"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-324b8376cc775e4174a00a69ff2c3aa66401b3570432060acd4989badef5d100"></a>

## Direct properties — filters.global_filters / 2d8a495a4bd4 / 3

<a id="canonical-310dca1b07d275558afb2cee5ba81ab50e5a7148bb62056eb4165dd8c1171171"></a>

<a id="canonical-4bf000c0f4aca9e657e0d5518f28287243ab7e7ca9293fca32822b9ca968f740"></a>

## key property — filters.global_filters / 2d8a495a4bd4 / 4

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

<a id="canonical-e0da3eaf33156a39a2b2c26214ab7a26a5ecdfff20338fad195c9e02df5c13a0"></a>

<a id="canonical-f80787df5846831261f0e329acf2600bc00b8bd3245e2a91210b6cc78eb5ac9c"></a>

## op property — filters.global_filters / 2d8a495a4bd4 / 5

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

<a id="canonical-a99fac01173a8399c6a851f67720601f5641625a5980a2269ecc0c99501cb226"></a>

<a id="canonical-c039116a178a7309601c74aa782f3ea41765adb7fc176ab04baf681495f62430"></a>

## values property — filters.global_filters / 2d8a495a4bd4 / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-3021e01039d0824970c2dc5b3b9e84c4bbed44d6b1b6523c797b96e14f22d203"></a>

## Next pages — filters.global_filters / 2d8a495a4bd4 / 7

- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-18d5387b9970901e50b9dc77f44f9cdc311f7cbfc6629564b84e6cecf8eadd05)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)

<a id="canonical-72c335c60953232fff4a455e778bcd6ee8b31429dd91f6462bf4b96331cc3f01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4a0209b54558b303e2b5377fa40fb7fc0fe139dcbc07e21b47c570a7054bd2d"></a>

## time_series_results — time_series_results / 8e42f59820cf / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- time_series_results

<a id="canonical-8793aa9d5d8cce4e1290f5d6873c47a83ee19d9ff89445be543c4049cab0cf2a"></a>

Type: `"list"`. Computed.

Collection of time series grouped by a series key.

<a id="canonical-008f1fc05b550455adba453c1504379f364ec5315ecce55b221863fc55188b99"></a>

## Direct properties — time_series_results / 8e42f59820cf / 3

<a id="canonical-a949781fd7e2f1f4d71d6d546d69b2c2a79d0770178d44bd9025961ba54a36bf"></a>

<a id="canonical-418cdc6ff198b61a866df76d785ec04bff53ccf82d3e891f4a9cfa7cfc043b58"></a>

## series_key property — time_series_results / 8e42f59820cf / 4

Type: `"string"`. Computed.

Identifier for the time series (e.g., 'total', 'high\_risk').

- [time_series](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-53f3fa297cc65353b034604f7942e2f8d21211cef1b00e1b0739259f1d3fc3fd): complete subsection reference.

<a id="canonical-d75ef562e3f3f9e6f6b49d5059454812e4d34ad129e0b6125d8cbe863c4a6625"></a>

## Next pages — time_series_results / 8e42f59820cf / 5

- [time_series_results.time_series](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-53f3fa297cc65353b034604f7942e2f8d21211cef1b00e1b0739259f1d3fc3fd)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)

<a id="canonical-53f3fa297cc65353b034604f7942e2f8d21211cef1b00e1b0739259f1d3fc3fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e732bb268e2c718013e3302b4d0111f3fcb042cda5156cc47dd78affd392d254"></a>

## time_series_results.time_series — time_series_results.time_series / b93a28a6ff3d / 2

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-8c637c6f29bdf42921c0ef7d58cc0d69a33fa56e41e372d98331c7c516d186ea)
- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-72c335c60953232fff4a455e778bcd6ee8b31429dd91f6462bf4b96331cc3f01)
- time_series_results.time_series

<a id="canonical-9a6805a59decdf324418f98a7cf588d62700081557484ea4c9db97a4a7fef52a"></a>

Type: `"list"`. Computed.

Sequence of timestamped values for this series.

<a id="canonical-36723f295db84057dea66a998c6d8e97e2e8b2131a0570d65369fdbd97aa6e3c"></a>

## Direct properties — time_series_results.time_series / b93a28a6ff3d / 3

<a id="canonical-b5f00d3c5b10c7872fc57b828f7ce468bb86e3f69238e49b5ae0fb2f1e2bf1ca"></a>

<a id="canonical-b4d4a4caa234eddec654f3f23d732323854fe39dc83d28602011410a7f72432d"></a>

## timestamp property — time_series_results.time_series / b93a28a6ff3d / 4

Type: `"string"`. Computed.

Timestamp (epoch seconds). Unix epoch timestamp in seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="canonical-76dc60be9c9d66f39b3d20e0e2f41e025d50b994d68a61c55859bcf01d733589"></a>

<a id="canonical-273ffc412c04916ad3cfa70c53e460e4e03bcfaf7836196297a329a39f2938e0"></a>

## value property — time_series_results.time_series / b93a28a6ff3d / 5

Type: `"string"`. Computed.

Value. Value observed at the given timestamp.

<a id="canonical-7e5ba7ed3df557d6f61c7a13c5f07bf676e72701c27dd673223ab671e8b892f5"></a>

## Next pages — time_series_results.time_series / b93a28a6ff3d / 6

- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-72c335c60953232fff4a455e778bcd6ee8b31429dd91f6462bf4b96331cc3f01)
- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-bbed0edf8db05608f8be2a8179ea53e5a9950cfd284c5ef8682ca9cf1f163e53)
