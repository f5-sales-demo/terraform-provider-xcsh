---
page_title: "xcsh_device_intelligence_multi_account_devices reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices reference."
---

# xcsh_device_intelligence_multi_account_devices reference

<a id="canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f130684b9ec61e48614c448b5fdf5fa2fd9da6bc68b5171b5d2522ec27716dc2"></a>

## Property reference — Property reference / 994dd8f88d1d / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- Property reference

<a id="canonical-97182d50e19adbabdb1072fa593ca7b5631c6d97f8b9af550be6b955b02838aa"></a>

## Direct properties — Property reference / 994dd8f88d1d / 3

<a id="canonical-46d2cf0fe37a7a80daa8bb2607eeda90bad8e8d20fcc8367c6ad3c53ac7ba079"></a>

<a id="canonical-45ee9239aded0e1c129268c86f1dc78396d632beb6f3ba1608cac9820b7c8cc8"></a>

## end_time property — Property reference / 994dd8f88d1d / 4

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

- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-523bba4129bfa6e9eccde14b0feb4ec4ae8b1e287dade2f02a3014b7d7a912af): complete subsection reference.

- [multi_account_devices](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-ce57511943d180c2f86e7b9a48985361838a4cc928bcc07954ae3bafe380aab5): complete subsection reference.

<a id="canonical-20add95e42060a38a2d28f59fc58d984de5489051bb6d74a766dbd2d051f7151"></a>

<a id="canonical-5bb84a6ea6921c77a7335145b56c41ea13173bfb883fa3874d7a6164dc0f6082"></a>

## namespace property — Property reference / 994dd8f88d1d / 5

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-675056472a5243c2fdd35aa5e739932a63fb551cf9219be9fdf17f4680521d24"></a>

<a id="canonical-4347a8e481d6b0a30ff395ae24441491adba01c3a089fbbef9b9d38e57feccd3"></a>

## start_time property — Property reference / 994dd8f88d1d / 6

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

<a id="canonical-9bd7bf9e395583cc48b91c86910230b1482a9b05c47975b3814481199a4f01cc"></a>

## All schema paths — Property reference / 994dd8f88d1d / 7

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-46d2cf0fe37a7a80daa8bb2607eeda90bad8e8d20fcc8367c6ad3c53ac7ba079) |
| `filters` | [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-7f6caab7120512e55c625a4ebd2fb6be0a7b3020dd2c8ff248f44c27b7d1f172) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-de4f64d02810fc1124f7c54647a64e4bf742c0b70cd9082ce0215a69bf56c0f3) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-b304e4e52cc272c25efdc703e5e7da52f66a4a52010621ab5d306c04a5fbaf49) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-70157a13949d74e42fce0c68d566c17141adf518d7ce8976d2751917fe9e34c5) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-d432143911b213bacb3881fe1c0ac54e3273f280f66bb8821241192bab9a1a9d) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-e6b469c2f6d0d7e9314ba1539b3400d14a4f0efa4888e75fa1476a7f33aa80fd) |
| `multi_account_devices` | [multi_account_devices](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-a185ad02d73d7f33801e9455dcc4bd4520c367da984bb6f64e3630e6d1e684ae) |
| `multi_account_devices.account_range` | [multi_account_devices.account_range](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-6316b26c4d00ca0fcaf5a3332de85c61873ddb960c03a2c87820446c7372f642) |
| `multi_account_devices.device_count` | [multi_account_devices.device_count](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-d2b087aba0c46e3dfec34bab55320876b0ba5a5837a07ada989003f15db89b98) |
| `namespace` | [namespace](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-20add95e42060a38a2d28f59fc58d984de5489051bb6d74a766dbd2d051f7151) |
| `start_time` | [start_time](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-675056472a5243c2fdd35aa5e739932a63fb551cf9219be9fdf17f4680521d24) |

<a id="canonical-9f6195d6231dd25ee35de7c4c986ce7454ee8f99c0e1318ecd9e63caf8ae71d3"></a>

## Next pages — Property reference / 994dd8f88d1d / 8

- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-523bba4129bfa6e9eccde14b0feb4ec4ae8b1e287dade2f02a3014b7d7a912af)
- [multi_account_devices](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-ce57511943d180c2f86e7b9a48985361838a4cc928bcc07954ae3bafe380aab5)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)

<a id="canonical-523bba4129bfa6e9eccde14b0feb4ec4ae8b1e287dade2f02a3014b7d7a912af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eacf95c5fdf3b8051ddcf175d1741fe816f4a0229e2a3d386a855b952b1450a"></a>

## filters — filters / 21403ab4e8c7 / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- filters

<a id="canonical-7f6caab7120512e55c625a4ebd2fb6be0a7b3020dd2c8ff248f44c27b7d1f172"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-7a6d1f0f572a09357f33a3ba40d489033c7ddcd3199a7dd3545341e9d0329543"></a>

## Direct properties — filters / 21403ab4e8c7 / 3

- [global_filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-5bf5e8fe808f4ad3056d3fe6f9535c9cdad0b23d8a83656610ca9bbb1e6c162d): complete subsection reference.

<a id="canonical-e6b469c2f6d0d7e9314ba1539b3400d14a4f0efa4888e75fa1476a7f33aa80fd"></a>

<a id="canonical-f43d587db17c4a55d1a87c34838e62c9edfcf48dbdd866c2b8ca895f75d8aa1b"></a>

## region_filter property — filters / 21403ab4e8c7 / 4

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

<a id="canonical-def1c77436f2d3931908e1f4e1458a0f9beeba20e7f567e99fd4be21bdbcf5c5"></a>

## Next pages — filters / 21403ab4e8c7 / 5

- [filters.global_filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-5bf5e8fe808f4ad3056d3fe6f9535c9cdad0b23d8a83656610ca9bbb1e6c162d)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)

<a id="canonical-5bf5e8fe808f4ad3056d3fe6f9535c9cdad0b23d8a83656610ca9bbb1e6c162d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f11d444cf1a8b65d496ac4d703be9b3ee6e10f11d86c0a4abb72f80b9e34e68"></a>

## filters.global_filters — filters.global_filters / 2b5b72051718 / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-523bba4129bfa6e9eccde14b0feb4ec4ae8b1e287dade2f02a3014b7d7a912af)
- filters.global_filters

<a id="canonical-de4f64d02810fc1124f7c54647a64e4bf742c0b70cd9082ce0215a69bf56c0f3"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-84f19e91032a1546609fe604ed278058de8fda523bb0cf0f234e4f87fe0dcb0f"></a>

## Direct properties — filters.global_filters / 2b5b72051718 / 3

<a id="canonical-b304e4e52cc272c25efdc703e5e7da52f66a4a52010621ab5d306c04a5fbaf49"></a>

<a id="canonical-b841a67b905a3efe7e1b5baf9bd31031448856dec9331ca186b5ee1f0c41880b"></a>

## key property — filters.global_filters / 2b5b72051718 / 4

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

<a id="canonical-70157a13949d74e42fce0c68d566c17141adf518d7ce8976d2751917fe9e34c5"></a>

<a id="canonical-20022193e3a012569a721ed3dbadeb2574288772675cde15a70781bb2a86f93f"></a>

## op property — filters.global_filters / 2b5b72051718 / 5

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

<a id="canonical-d432143911b213bacb3881fe1c0ac54e3273f280f66bb8821241192bab9a1a9d"></a>

<a id="canonical-cbedf7d28e7e85c11f9185577ddcca32115f8c1c00eade07b09b9a32f556bb88"></a>

## values property — filters.global_filters / 2b5b72051718 / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-28e53216168ade3b12f19b793a6845cae06a2fe429a07c14b6f3b2607769f3e3"></a>

## Next pages — filters.global_filters / 2b5b72051718 / 7

- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-523bba4129bfa6e9eccde14b0feb4ec4ae8b1e287dade2f02a3014b7d7a912af)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)

<a id="canonical-ce57511943d180c2f86e7b9a48985361838a4cc928bcc07954ae3bafe380aab5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b588d285621e29278b3a440c76e6cca43a55c392274da8c2fd3a60688ed06612"></a>

## multi_account_devices — multi_account_devices / 866677ab0409 / 2

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- multi_account_devices

<a id="canonical-a185ad02d73d7f33801e9455dcc4bd4520c367da984bb6f64e3630e6d1e684ae"></a>

Type: `"list"`. Computed.

Distribution of devices by account range buckets.

<a id="canonical-842c788fdf831feb8e7960675d07b3926b6282857e66f5fd498b4c9eae4b103a"></a>

## Direct properties — multi_account_devices / 866677ab0409 / 3

<a id="canonical-6316b26c4d00ca0fcaf5a3332de85c61873ddb960c03a2c87820446c7372f642"></a>

<a id="canonical-79259d96fe9ad30ad926103e5e26d3165742db6d121d6275d7466d665fbe42c6"></a>

## account_range property — multi_account_devices / 866677ab0409 / 4

Type: `"string"`. Computed.

Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').

<a id="canonical-d2b087aba0c46e3dfec34bab55320876b0ba5a5837a07ada989003f15db89b98"></a>

<a id="canonical-e4f20fa975312ac0f2fd2338ed7b9fa29387e14516a5da1f61bb876c9f0c4531"></a>

## device_count property — multi_account_devices / 866677ab0409 / 5

Type: `"string"`. Computed.

Number of devices in this account range.

<a id="canonical-b9cbae86db26bbc6a581addddf2aaf5b97bca294c8eeab81f43dfa8a6635bb65"></a>

## Next pages — multi_account_devices / 866677ab0409 / 6

- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0297391df16a3a9dd9f80b68cd4db055feaadb98c0563c5575f355edfa4a9064)
- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-d7032758bccd44c0b77f242e2c10eff4929329ac7188194b5b4d1630cd87bd40)
