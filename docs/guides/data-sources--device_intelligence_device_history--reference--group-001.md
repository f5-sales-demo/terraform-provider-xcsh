---
page_title: "xcsh_device_intelligence_device_history reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_device_history reference."
---

# xcsh_device_intelligence_device_history reference

<a id="canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-502ec42fd1d1aea42f45ccd3d1d73f5c1d7a74817b3beb49ffc772fb506dd8bb"></a>

## Property reference — Property reference / d67d97d57b86 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- Property reference

<a id="canonical-9d637191cf85cc11dd5fe4dab66eb6ac369e8533fb3a45a867d2c904b2fb6d14"></a>

## Direct properties — Property reference / d67d97d57b86 / 3

<a id="canonical-e88c6fc7c233deef1f8ea57bfd52f7631905286e0624f98f932d231f172369de"></a>

<a id="canonical-2cbc8cc91e6149ff444e337396377ef9138d9d53870d4e327da15404815e767a"></a>

## device_id property — Property reference / d67d97d57b86 / 4

Type: `"string"`. Required.

DeviceID. Device identifier.

<a id="canonical-94c0b51a25f5ed49699024b1d1a3e8232a9276bef5c8eb122c6b5c246ded4a4e"></a>

<a id="canonical-0fd2bea22c3dc04b7512ac921bf34506645405ad6d59a8dac2ded346be609041"></a>

## end_time property — Property reference / d67d97d57b86 / 5

Type: `"string"`. Optional.

End Time. End time of the query period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-f871ca52d9f8d479219c2274cd66894d5b2b7ea5e4b0abef9e423ba23326a6e8): complete subsection reference.

<a id="canonical-2154d8780cb7c35f8022331a4c09e95c7c581f9a799d55a2eac2afdff50315b1"></a>

<a id="canonical-ff8c3fbfc5423e24f9cf22ad898b442c4fb83be67975c01818e29ab27bd1113f"></a>

## namespace property — Property reference / d67d97d57b86 / 6

Type: `"string"`. Required.

Namespace. Namespace name.

- [pagination](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-5096ac408562e63fb712f4ea657c93744f087059cab63aad872bdedd84d6d882): complete subsection reference.

- [records](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-136fa5b94208bfb4bada8e02d9190a2f084bd260021aae16df54948b7bcc36c8): complete subsection reference.

- [sort](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-0cf8c266ce469763005532bbdb0d67c9a796b72b9210af7276214b6cc9192f8d): complete subsection reference.

<a id="canonical-6986d59721b4b530d15b068dafb2c41efded45bc8a26d4fd3cbaca2f361b20d5"></a>

<a id="canonical-c3945cf373ff8632e2373ddb9f2538087817469bce5fac47ebf0236dc63c2341"></a>

## start_time property — Property reference / d67d97d57b86 / 7

Type: `"string"`. Optional.

Start Time. Start time of the query period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-43cfb3f0254534bd6ee0ecad18c4d4bd68758bcc6ce59d04cbbf55a29b98ecfc"></a>

## All schema paths — Property reference / d67d97d57b86 / 8

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `device_id` | [device_id](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-e88c6fc7c233deef1f8ea57bfd52f7631905286e0624f98f932d231f172369de) |
| `end_time` | [end_time](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-94c0b51a25f5ed49699024b1d1a3e8232a9276bef5c8eb122c6b5c246ded4a4e) |
| `filters` | [filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-7525b30aded385e7a20fd289463683b38d6c2750de2b700040f8f4f29592b05a) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-bcc1737abc58ae1bfe2a492dcd94b282fc3875d9491c6c67afd639a6598bae66) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-f17626e4968894ed133983523c1c53d9308ebc34dd0dddc2dac98ecd05a34ed0) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-c3cf96f340fe5d5caf9705a261143fc7b25574e0e1aec2e406ad22c8cf2b0a6f) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-68dcaf890e4a2e31c094e1f950d960a325a2c0dbf330abfb911eed5d310ba920) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-ee0408a949c97331137925d9079699498bf5ced06f7dfb3182e732efb0796c81) |
| `namespace` | [namespace](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-2154d8780cb7c35f8022331a4c09e95c7c581f9a799d55a2eac2afdff50315b1) |
| `pagination` | [pagination](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-394eea03357db1c70975186b77740f145734724933aaca6d59ececdeb790fbef) |
| `pagination.page_number` | [pagination.page_number](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-84be7bd4a5abd92a87ccb08b24dd4d2d40ef4cf64e4e739929aa7c97c79bc6bf) |
| `pagination.page_size` | [pagination.page_size](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-a6f5da96a4df8fad00329344df5925598f6ca74c4381779c45e50c33706b8072) |
| `records` | [records](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-eddd023af4fb6f2ce1730fcff95b627d6b794e90ccb459d00f222e551e55fc78) |
| `records.action_taken` | [records.action_taken](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-47aab07492271e38913cf81be1363b0c87faf621c8ebedea41bc8442c6d93463) |
| `records.detected_signals` | [records.detected_signals](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-b881238805a1cd227af6009e1cce5e744daeb64bac3c7c0ee4d983ea3237bc60) |
| `records.endpoint_label` | [records.endpoint_label](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-2b5ed81123c6390be7931de1ae2f7a389f342cd132e4b97ac31964a8ac291861) |
| `records.risk_score` | [records.risk_score](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-b992cff49c7dee31a24a13396018093af1c41177f7c335658e8db50b881d6d37) |
| `records.timestamp` | [records.timestamp](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-d11991738715e205d9f4cedb1fa4f7bfc0ad1a52ff5352a628da8df19725bf18) |
| `records.txn_id` | [records.txn_id](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-8f10f86be296ef34d323e5230ef55af90edcdba9657bfd842a5eb225ce3e3ac2) |
| `records.url` | [records.url](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-c8dd3a81a5e7af06c5c45ca5fd8bf1b7e5c4cbe92714a847708448bad93017ed) |
| `sort` | [sort](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-9ed088da774e262927043b3c28eef066fd7795c3d80b389d2cb609acb54d9ff0) |
| `sort.key` | [sort.key](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-93f307455a844a04532fb2d0fbde46639915c5929e81d5789ef3986de4e896b3) |
| `sort.order` | [sort.order](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-2b1a0c35a6ad02a7fcb48151d4f9281489887ca5edcbe1f0060c0acb07871993) |
| `start_time` | [start_time](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-6986d59721b4b530d15b068dafb2c41efded45bc8a26d4fd3cbaca2f361b20d5) |

<a id="canonical-6988970b7703dac38da99680aae7d72911af7b80172ee5892711e18e8cb10d2c"></a>

## Next pages — Property reference / d67d97d57b86 / 9

- [filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-f871ca52d9f8d479219c2274cd66894d5b2b7ea5e4b0abef9e423ba23326a6e8)
- [pagination](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-5096ac408562e63fb712f4ea657c93744f087059cab63aad872bdedd84d6d882)
- [records](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-136fa5b94208bfb4bada8e02d9190a2f084bd260021aae16df54948b7bcc36c8)
- [sort](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-0cf8c266ce469763005532bbdb0d67c9a796b72b9210af7276214b6cc9192f8d)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-f871ca52d9f8d479219c2274cd66894d5b2b7ea5e4b0abef9e423ba23326a6e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-039c047f0fb877c1ba4cdc4baf16364381f41bd60b2d32ae05b6d0b17fa1d223"></a>

## filters — filters / aafd6f69467c / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- filters

<a id="canonical-7525b30aded385e7a20fd289463683b38d6c2750de2b700040f8f4f29592b05a"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-6e9991452373bedf157f4076a67be8951dfabbef530bee8f12d87d7c29ff0f5e"></a>

## Direct properties — filters / aafd6f69467c / 3

- [global_filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-3ea48734ddbf43973fac2a2feb33def5af95f53a8a04edcc1271e31268508168): complete subsection reference.

<a id="canonical-ee0408a949c97331137925d9079699498bf5ced06f7dfb3182e732efb0796c81"></a>

<a id="canonical-9f3cff4d7a441b9aedfd8ea9775ccaf99be16111307043bb224f88020f6e6ec9"></a>

## region_filter property — filters / aafd6f69467c / 4

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

<a id="canonical-abcc460c399e2170f8c5117d038481636a5caad6518c6a6c9be2a5e66e53f6d4"></a>

## Next pages — filters / aafd6f69467c / 5

- [filters.global_filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-3ea48734ddbf43973fac2a2feb33def5af95f53a8a04edcc1271e31268508168)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-3ea48734ddbf43973fac2a2feb33def5af95f53a8a04edcc1271e31268508168"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36292129d34662b44b8cdb2b83b93b39346e2fc3e33b32926d606ca19e6ed0ef"></a>

## filters.global_filters — filters.global_filters / 73f523307fb2 / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-f871ca52d9f8d479219c2274cd66894d5b2b7ea5e4b0abef9e423ba23326a6e8)
- filters.global_filters

<a id="canonical-bcc1737abc58ae1bfe2a492dcd94b282fc3875d9491c6c67afd639a6598bae66"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-d075de59524938470a039677484528d5ec05af23b7ad4bfe456ca7805c082397"></a>

## Direct properties — filters.global_filters / 73f523307fb2 / 3

<a id="canonical-f17626e4968894ed133983523c1c53d9308ebc34dd0dddc2dac98ecd05a34ed0"></a>

<a id="canonical-8b1b8c2b6250451cb34450cd55c6a1939cfae4b276b0e7329988f45ddc62de8e"></a>

## key property — filters.global_filters / 73f523307fb2 / 4

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

<a id="canonical-c3cf96f340fe5d5caf9705a261143fc7b25574e0e1aec2e406ad22c8cf2b0a6f"></a>

<a id="canonical-86bb870799623e7aac9926e8364db408a151bed3c7b4a92b007eb09dde165dcc"></a>

## op property — filters.global_filters / 73f523307fb2 / 5

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

<a id="canonical-68dcaf890e4a2e31c094e1f950d960a325a2c0dbf330abfb911eed5d310ba920"></a>

<a id="canonical-26b0a78d58eecbc5413a67f0629f0d60f2c451d4c360a5e5ede0edd891d6502e"></a>

## values property — filters.global_filters / 73f523307fb2 / 6

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-9448b39ff5b288a2c052b243d67179d8f7150b0980df721acc91737fa93413f6"></a>

## Next pages — filters.global_filters / 73f523307fb2 / 7

- [filters](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-f871ca52d9f8d479219c2274cd66894d5b2b7ea5e4b0abef9e423ba23326a6e8)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-5096ac408562e63fb712f4ea657c93744f087059cab63aad872bdedd84d6d882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-540be225e5b2dd2be407d17f54719453b7f2a8a9dba10f6eaf7bcc58b31a4392"></a>

## pagination — pagination / c4bad0fe70ff / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- pagination

<a id="canonical-394eea03357db1c70975186b77740f145734724933aaca6d59ececdeb790fbef"></a>

Type: `"single"`. Optional.

Pagination for Request with number and size.

<a id="canonical-3eba36a0c6367af521887442a5cf4fea91eb656e057a3cf315d7101f12fddc8b"></a>

## Direct properties — pagination / c4bad0fe70ff / 3

<a id="canonical-84be7bd4a5abd92a87ccb08b24dd4d2d40ef4cf64e4e739929aa7c97c79bc6bf"></a>

<a id="canonical-b3df32f50531575b64c6692a1be8d91fb1de5f1e698e2c3dcc96d4ff844f0087"></a>

## page_number property — pagination / c4bad0fe70ff / 4

Type: `"number"`. Optional.

Configuration parameter for page number.

<a id="canonical-a6f5da96a4df8fad00329344df5925598f6ca74c4381779c45e50c33706b8072"></a>

<a id="canonical-14fc869d26c7466fa800c5c92cad4035f3782d631a34344eab3faaf125597a5d"></a>

## page_size property — pagination / c4bad0fe70ff / 5

Type: `"number"`. Optional.

Page Size. Size or capacity specification

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 500),
}
```

<a id="canonical-6aad3b466709b4dccbfe28afefd4be2a1bae61d9a83c6b7707d5d8fe339ea42d"></a>

## Next pages — pagination / c4bad0fe70ff / 6

- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-136fa5b94208bfb4bada8e02d9190a2f084bd260021aae16df54948b7bcc36c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d01c00ab3a721742f560c375eb9a978d83b65d92020ab60d3dd44d08ff91de15"></a>

## records — records / d6118c470d7a / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- records

<a id="canonical-eddd023af4fb6f2ce1730fcff95b627d6b794e90ccb459d00f222e551e55fc78"></a>

Type: `"list"`. Computed.

Records. List of activity records.

<a id="canonical-43e865d5b0aa4bd87a7858bb842a67c7af6237241f4c557684532e5c97a87476"></a>

## Direct properties — records / d6118c470d7a / 3

<a id="canonical-47aab07492271e38913cf81be1363b0c87faf621c8ebedea41bc8442c6d93463"></a>

<a id="canonical-46e5730417a4b6a1a907ffe6829b5739c998acef1093def733a9903ce3fe07c3"></a>

## action_taken property — records / d6118c470d7a / 4

Type: `"string"`. Computed.

Action taken in response to the risk evaluation for this event.

<a id="canonical-b881238805a1cd227af6009e1cce5e744daeb64bac3c7c0ee4d983ea3237bc60"></a>

<a id="canonical-124c6c3253d7e59a828d7c541fc0dac6e05097a8cc2ffe1b11b834b2f49dfea3"></a>

## detected_signals property — records / d6118c470d7a / 5

Type: `["list", "string"]`. Computed.

Risk or integrity signals observed for this event.

<a id="canonical-2b5ed81123c6390be7931de1ae2f7a389f342cd132e4b97ac31964a8ac291861"></a>

<a id="canonical-0a7e0a6e71df0a83e6eee955665609fb670805ef3d775a985202499ec4863642"></a>

## endpoint_label property — records / d6118c470d7a / 6

Type: `"string"`. Computed.

Human-readable label for the endpoint or action.

<a id="canonical-b992cff49c7dee31a24a13396018093af1c41177f7c335658e8db50b881d6d37"></a>

<a id="canonical-b17524159d6144b36b65e546e8575a376d9533d17ec121e0566a1c9a928b622f"></a>

## risk_score property — records / d6118c470d7a / 7

Type: `"number"`. Computed.

Risk score assigned to this event (0–100).

<a id="canonical-d11991738715e205d9f4cedb1fa4f7bfc0ad1a52ff5352a628da8df19725bf18"></a>

<a id="canonical-e34c1c39a245c72d3df815f4f3c3a0f6ba5854e124897c1b603f1b22eaa8b363"></a>

## timestamp property — records / d6118c470d7a / 8

Type: `"string"`. Computed.

Event timestamp in RFC 3339 format with milliseconds (UTC).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="canonical-8f10f86be296ef34d323e5230ef55af90edcdba9657bfd842a5eb225ce3e3ac2"></a>

<a id="canonical-12f459b05ebf1a1fd0571539e411fe9c7ea0ee8a0e43d3b6ede3cfd6ec7c922c"></a>

## txn_id property — records / d6118c470d7a / 9

Type: `"string"`. Computed.

Identifier of the associated transaction.

<a id="canonical-c8dd3a81a5e7af06c5c45ca5fd8bf1b7e5c4cbe92714a847708448bad93017ed"></a>

<a id="canonical-b87f81d82733a7afa215d9fceacdbd8533c2bfd4019f2adb319ab4d85b222a53"></a>

## url property — records / d6118c470d7a / 10

Type: `"string"`. Computed.

Endpoint URL or host associated with the event.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`),
    ""),
}
```

<a id="canonical-0c8fe8b298ead6e2008e9c89d2606edb8c3e6b85c49d8144f3b5ba591d574013"></a>

## Next pages — records / d6118c470d7a / 11

- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)

<a id="canonical-0cf8c266ce469763005532bbdb0d67c9a796b72b9210af7276214b6cc9192f8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f72b99e59f01a0de756790876233b5038ecd9a15ed3ac68b6b5271768f350cc"></a>

## sort — sort / 33e88ba0bc0d / 2

Breadcrumbs:

- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- sort

<a id="canonical-9ed088da774e262927043b3c28eef066fd7795c3d80b389d2cb609acb54d9ff0"></a>

Type: `"single"`. Optional.

Sort Option. Query Result Sort Option.

<a id="canonical-c7d2d051319bef216ef621a5743331c87c0ac13c5fc55547746895c19bb128e8"></a>

## Direct properties — sort / 33e88ba0bc0d / 3

<a id="canonical-93f307455a844a04532fb2d0fbde46639915c5929e81d5789ef3986de4e896b3"></a>

<a id="canonical-b4ba67dec1e50d39a14cd787d27ce0730ef5c091c7a86e95965b9fa98504a4fa"></a>

## key property — sort / 33e88ba0bc0d / 4

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

<a id="canonical-2b1a0c35a6ad02a7fcb48151d4f9281489887ca5edcbe1f0060c0acb07871993"></a>

<a id="canonical-c4c0c956562c545f944422eb832fdc1283ba735c7e154ca3ba5d1479e72a1836"></a>

## order property — sort / 33e88ba0bc0d / 5

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

<a id="canonical-8ada0fd4b805ca57898a5364c4b4c4772b6523b742be121be6b21c3836ccd1cf"></a>

## Next pages — sort / 33e88ba0bc0d / 6

- [Property reference](data-sources--device_intelligence_device_history--reference--group-001.md#canonical-57a6e4c4c70f6a59052929b2dd18df6408ce5754db6d69b2b048613e92dd4d8c)
- [xcsh_device_intelligence_device_history](../data-sources/device_intelligence_device_history.md#canonical-4f287dc4ac3600af7828b70c0aa416a821d47f3785e18cec7baf65231b298079)
