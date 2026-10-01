---
page_title: "xcsh_device_intelligence_summary reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_summary reference."
---

# xcsh_device_intelligence_summary reference

<a id="canonical-1c701a3bcd569781ea4b09b023243540a7da6f24b4b2c73c27e1755ab7ac2705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c104ae0569f522a34c97d3bdc846084c740a1402058d65fb9584a4495cc412ea"></a>

## Property reference — Property reference / 8d3148be8cea / 2

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
- Property reference

<a id="canonical-874a18160e528b37ab0840691d0835aaf004510e3b52a062a22391cf4da81add"></a>

## Direct properties — Property reference / 8d3148be8cea / 3

<a id="canonical-c1eba6bb3faef8ee0a36af2dc61c21586b73b1701f53c637f83219bf2e995d14"></a>

<a id="canonical-bd4a61a72f38dc7f36d1102fe79c5478acca1b6a093a98c11bc503586297498d"></a>

## end_time property — Property reference / 8d3148be8cea / 4

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

- [filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-34082f6eb6fd8b0300fb6097669fe0557ae03f4c0f17134205634b5570c4a14e): complete subsection reference.

<a id="canonical-0643032719eeb7a8ca8f222bbec9596ad1c0c81ece0f6ebba2691d2ec1f841aa"></a>

<a id="canonical-068f7d384f4513c84b77316838509fdd992dc658df20e52b82bd3cb00e9898b5"></a>

## high_risk_device_count property — Property reference / 8d3148be8cea / 5

Type: `"string"`. Computed.

Number of devices classified as high risk.

<a id="canonical-f684b1fe85a1f0d79f8f089adb5df18352ca4b40c903a05f2778d3502b9beded"></a>

<a id="canonical-b8da7e8e36de0b599591c93a28aaad463f57685a53b1c4fa42223c61a8469555"></a>

## high_risk_device_rate property — Property reference / 8d3148be8cea / 6

Type: `"number"`. Computed.

High-risk devices as a percentage of total devices.

<a id="canonical-5c2ae916eaf58896e1c04f78c4f46288ae5bffc9d2e51fe4345a0924de7dc372"></a>

<a id="canonical-319a08df6c17da7a82a56cbec3f63e4300b4b4c890e713a35277075230667010"></a>

## high_risk_txn_count property — Property reference / 8d3148be8cea / 7

Type: `"string"`. Computed.

Number of transactions classified as high risk.

<a id="canonical-1a30c30f78b9579a2cc574f94c19d751b7e618400efa1baa4e57289f04ce73e8"></a>

<a id="canonical-80a97000e9ee25e15d7688c525a94073ff8d10f7ae510e99f1eba529b3011ab8"></a>

## high_risk_txn_rate property — Property reference / 8d3148be8cea / 8

Type: `"number"`. Computed.

High-risk transactions as a percentage of total transactions.

<a id="canonical-a6a8f3187ec425c18bb3b5d8396ac4e427132c76758bf2d3d8b95011d8ee76fb"></a>

<a id="canonical-cbd85c52e59a48781c3f07152516ab8c37d03e0d5f83936fdd99c8e38ed5e4dd"></a>

## multi_acc_high_risk_device_count property — Property reference / 8d3148be8cea / 9

Type: `"string"`. Computed.

Number of high-risk devices associated with multiple accounts.

<a id="canonical-362fb484a44bf3c6a5829b2b968d6ca2e47fe86c7206988d0ac7f8d495938b17"></a>

<a id="canonical-7101b9ec674282487535813ddff1241ce9e36f62e0ae77efab13323cba0c29dd"></a>

## namespace property — Property reference / 8d3148be8cea / 10

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-3cfcb52acb5265e3e064535cf69e7cc32335674d8364e3000df88618e2f8f6ba"></a>

<a id="canonical-7629c2912ba8cf261484e7dc3b6fd613647ba029ec762d2086eedaf1b1662510"></a>

## start_time property — Property reference / 8d3148be8cea / 11

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

<a id="canonical-2f9d4f0694b82df8745cbf0a042b11cba123a0b40e2a7f7ee2e6617c75d78a9e"></a>

<a id="canonical-e3ee357e06fae5189694a48541550031b051714cc46f4b7ecd01ed726de517f5"></a>

## total_device_count property — Property reference / 8d3148be8cea / 12

Type: `"string"`. Computed.

Total Device Count. Total number of devices observed.

<a id="canonical-05a12de6c1980592e28b497480ac2161492e12f47bc536d8f67efcde662bb510"></a>

<a id="canonical-115f07ccb7b54c637365ede6e8b9b86b24cfe39aa927cdc5cbc2ad8b4938c7f7"></a>

## total_evaluated_txn_count property — Property reference / 8d3148be8cea / 13

Type: `"string"`. Computed.

Total Evaluated Transaction Count. Total number of transactions evaluated.

<a id="canonical-fce2636e52b15b3516c85e0a517a30ceb4e559c58cffb1e221e9647e52b7094d"></a>

<a id="canonical-91cd4651d53d46009c54118b4f2d2bbb13b7b76052dc4f5dd772df191e239a99"></a>

## total_multi_acc_device_count property — Property reference / 8d3148be8cea / 14

Type: `"string"`. Computed.

Total number of devices associated with multiple accounts.

<a id="canonical-356a32d2722a02fbc6bf718c968a9dedcb83a4916f52997a498be83ccc48e6e9"></a>

<a id="canonical-cdf58a89eda9514cf2db130565d24c2233135280beac38ac9d054bcb938eb16b"></a>

## total_txn_count property — Property reference / 8d3148be8cea / 15

Type: `"string"`. Computed.

Total Transaction Count. Total number of transactions.

<a id="canonical-e19f6ef4ad571043184e83d0b9ee11d5f09b57ba8139cd850ad05012f226b7e8"></a>

## All schema paths — Property reference / 8d3148be8cea / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_summary--reference--group-001.md#canonical-c1eba6bb3faef8ee0a36af2dc61c21586b73b1701f53c637f83219bf2e995d14) |
| `filters` | [filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-77a47add0b7d1d2ecefca06de54f405765697ba0a1cd44b393c5fdad68e318e6) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-5e84d783d0488559c7e73f6c955145c30bc46ec8ab746edaf66ae45744506030) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_summary--reference--group-001.md#canonical-72e9201d97cc6a8f35ffaf6dbd1591c9a84e14e695687d3938941ae1d7e3861c) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_summary--reference--group-001.md#canonical-48d495efb78d13e0b01e14aba684000d85c1be55e240beb6e0b49d915d09c803) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_summary--reference--group-001.md#canonical-55318f2a13e98e47e9abc39a98b085694a8b0e38bc04a33f8310ab3b3198fc6f) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_summary--reference--group-001.md#canonical-0448cee12c44150a38f2b2e7f1b098af9963498989d22f95f40b6447bc221140) |
| `high_risk_device_count` | [high_risk_device_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-0643032719eeb7a8ca8f222bbec9596ad1c0c81ece0f6ebba2691d2ec1f841aa) |
| `high_risk_device_rate` | [high_risk_device_rate](data-sources--device_intelligence_summary--reference--group-001.md#canonical-f684b1fe85a1f0d79f8f089adb5df18352ca4b40c903a05f2778d3502b9beded) |
| `high_risk_txn_count` | [high_risk_txn_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-5c2ae916eaf58896e1c04f78c4f46288ae5bffc9d2e51fe4345a0924de7dc372) |
| `high_risk_txn_rate` | [high_risk_txn_rate](data-sources--device_intelligence_summary--reference--group-001.md#canonical-1a30c30f78b9579a2cc574f94c19d751b7e618400efa1baa4e57289f04ce73e8) |
| `multi_acc_high_risk_device_count` | [multi_acc_high_risk_device_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-a6a8f3187ec425c18bb3b5d8396ac4e427132c76758bf2d3d8b95011d8ee76fb) |
| `namespace` | [namespace](data-sources--device_intelligence_summary--reference--group-001.md#canonical-362fb484a44bf3c6a5829b2b968d6ca2e47fe86c7206988d0ac7f8d495938b17) |
| `start_time` | [start_time](data-sources--device_intelligence_summary--reference--group-001.md#canonical-3cfcb52acb5265e3e064535cf69e7cc32335674d8364e3000df88618e2f8f6ba) |
| `total_device_count` | [total_device_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-2f9d4f0694b82df8745cbf0a042b11cba123a0b40e2a7f7ee2e6617c75d78a9e) |
| `total_evaluated_txn_count` | [total_evaluated_txn_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-05a12de6c1980592e28b497480ac2161492e12f47bc536d8f67efcde662bb510) |
| `total_multi_acc_device_count` | [total_multi_acc_device_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-fce2636e52b15b3516c85e0a517a30ceb4e559c58cffb1e221e9647e52b7094d) |
| `total_txn_count` | [total_txn_count](data-sources--device_intelligence_summary--reference--group-001.md#canonical-356a32d2722a02fbc6bf718c968a9dedcb83a4916f52997a498be83ccc48e6e9) |

<a id="canonical-3e3ddc4070b5229177acf40b4d96635ac1dd2627d69027c4cb77aac11c82e91c"></a>

## Next pages — Property reference / 8d3148be8cea / 17

- [filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-34082f6eb6fd8b0300fb6097669fe0557ae03f4c0f17134205634b5570c4a14e)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)

<a id="canonical-34082f6eb6fd8b0300fb6097669fe0557ae03f4c0f17134205634b5570c4a14e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3d43fb0cff3e34add719fb5f0e4e6f43c3620d68ffb676f5495aaf348a4ecef"></a>

## filters — filters / e6da13a209ee / 2

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
- [Property reference](data-sources--device_intelligence_summary--reference--group-001.md#canonical-1c701a3bcd569781ea4b09b023243540a7da6f24b4b2c73c27e1755ab7ac2705)
- filters

<a id="canonical-77a47add0b7d1d2ecefca06de54f405765697ba0a1cd44b393c5fdad68e318e6"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-afe4ce3edbdfd8c5412349c6327a15ce9b67ea1c4411b16c13f6aa75bce5ee3a"></a>

## Direct properties — filters / e6da13a209ee / 3

- [global_filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-2d6c270ce65d01b7a41db6c0e1434aa5b90578059eff1cb3791d10d35497034b): complete subsection reference.

<a id="canonical-0448cee12c44150a38f2b2e7f1b098af9963498989d22f95f40b6447bc221140"></a>

<a id="canonical-0ef67577fff1db2a74eec870d96e8dda3dcd612c6552b2fd7b3b0cd225a0399d"></a>

## region_filter property — filters / e6da13a209ee / 4

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

<a id="canonical-a0008ac0353f8ed40254c9ff4ef5dc2c9440bbcc34a0bdac5d82f16ef2c57633"></a>

## Next pages — filters / e6da13a209ee / 5

- [filters.global_filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-2d6c270ce65d01b7a41db6c0e1434aa5b90578059eff1cb3791d10d35497034b)
- [Property reference](data-sources--device_intelligence_summary--reference--group-001.md#canonical-1c701a3bcd569781ea4b09b023243540a7da6f24b4b2c73c27e1755ab7ac2705)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)

<a id="canonical-2d6c270ce65d01b7a41db6c0e1434aa5b90578059eff1cb3791d10d35497034b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e670f5f877839509a96eb0c938110b1a65536875b07b462b68b6c1984f53092"></a>

## filters.global_filters — filters.global_filters / a76e0c9ba565 / 2

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
- [Property reference](data-sources--device_intelligence_summary--reference--group-001.md#canonical-1c701a3bcd569781ea4b09b023243540a7da6f24b4b2c73c27e1755ab7ac2705)
- [filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-34082f6eb6fd8b0300fb6097669fe0557ae03f4c0f17134205634b5570c4a14e)
- filters.global_filters

<a id="canonical-5e84d783d0488559c7e73f6c955145c30bc46ec8ab746edaf66ae45744506030"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-f5c9a2e966fa9d3858931e4f8408e14777f023a6291ca6a86c8f3a1b7fe67824"></a>

## Direct properties — filters.global_filters / a76e0c9ba565 / 3

<a id="canonical-72e9201d97cc6a8f35ffaf6dbd1591c9a84e14e695687d3938941ae1d7e3861c"></a>

<a id="canonical-374ada8ccf31ad8d3363e9d0d6a7f561d115f305fd918c703eb7b01b7bab8fb4"></a>

## key property — filters.global_filters / a76e0c9ba565 / 4

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

<a id="canonical-48d495efb78d13e0b01e14aba684000d85c1be55e240beb6e0b49d915d09c803"></a>

<a id="canonical-abb68977891b72dd02309359be703b926afefa423546d8bc9526445ad144fbe0"></a>

## op property — filters.global_filters / a76e0c9ba565 / 5

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

<a id="canonical-55318f2a13e98e47e9abc39a98b085694a8b0e38bc04a33f8310ab3b3198fc6f"></a>

<a id="canonical-2bc876752c946f0385af8bf34f347bb9cde2eebe5a8734852b7f0ec4c5487fe9"></a>

## values property — filters.global_filters / a76e0c9ba565 / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-88a140df79b5c86dad57b819197a0b3452e1340330c8b771b00ed29827673446"></a>

## Next pages — filters.global_filters / a76e0c9ba565 / 7

- [filters](data-sources--device_intelligence_summary--reference--group-001.md#canonical-34082f6eb6fd8b0300fb6097669fe0557ae03f4c0f17134205634b5570c4a14e)
- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9)
