---
page_title: "xcsh_device_intelligence_multi_account_devices reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_multi_account_devices reference."
---

# xcsh_device_intelligence_multi_account_devices reference

<a id="canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- Property reference

<a id="canonical-3301030012201023-2132301201321020-1201103010102023-1133313311332202-3331213122122330-1220231101130123-1131021102023230-0213130112313002"></a>

### Direct properties for `xcsh_device_intelligence_multi_account_devices`

<a id="canonical-1012310230330033-3203132213222000-3122222023230212-0013323231222100-2322312032203102-0033303020031213-3012223103301103-2230132322001321"></a>

#### `end_time` property

Type: `"string"`. Optional.

End time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
end\_time will be evaluated to start\_time+10m If start\_time is not specified, then the end\_time
will be evaluated to &lt;current time&gt;.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1102032323221001-0221233322123221-3230303132011023-0033322310323010-2232202301320220-1331223132023300-0222030001102313-3113222101022233): complete subsection reference.

- [multi_account_devices](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-3032111311010121-1003310120003002-3320123213232122-1020212011031201-2003202210303021-0220233030001321-1110223203232233-3203200022222311): complete subsection reference.

<a id="canonical-0200223131211132-1002001200220320-2202310220331121-3330112031212010-3132111020210011-0123231231131022-1312123123310231-0011013313011101"></a>

<a id="canonical-2113012002311100-3201212231232223-3123010013023322-1121033022132311-1203013012312113-3320232122331111-0023321223211111-2300022003202222"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-1213110011121013-0222110210033002-3331310311222211-3213032121030222-1203332311110130-3321020121233221-3331330113331012-2000110201310210"></a>

<a id="canonical-1011323221020321-2231323100320130-0102210212203020-1233013130132003-2112311203022332-2312330323220112-0020302230212002-0023133020303020"></a>

#### `start_time` property

Type: `"string"`. Optional.

Start time of the query period Format: unix\_timestamp|RFC 3339 Optional: If not specified, then the
start\_time will be evaluated to end\_time-10m If end\_time is not specified, then the start\_time
will be evaluated to &lt;current time&gt;-10m.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-1123232010221232-2212210201301313-2213030311011011-2311123010013222-0103011303233323-2020033322032013-1031132212011210-3130003312002002"></a>

### All schema paths for `xcsh_device_intelligence_multi_account_devices`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1012310230330033-3203132213222000-3122222023230212-0013323231222100-2322312032203102-0033303020031213-3012223103301103-2230132322001321) |
| `filters` | [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1333123022222313-0102001101023211-1130120211221032-2331023323122332-0022132303000200-3131023020333302-1020331010300213-2313310133011302) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-3132103312103100-0220010033300101-0210331330111012-1013221210321023-3313100230002313-0030312100200230-3200020111221221-2333111230003303) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-2303001032103211-0230300213023002-1132333130130003-3211321331221102-3312122210221102-0001001202012223-1131030012300010-2211332322331021) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1300011113220103-2110213113103210-0233303200301220-3111121230011301-1001223133110120-3113303220211312-3102131101210113-3332213203103011) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-3110030201100321-0101230201032322-3023032020013332-0130002230111032-0302130333022000-3312122323202002-0102100101210223-2223212201222131) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-3212231012213002-3312310031133221-0301102322011103-2123031000003101-1022103300323322-1020202032131133-2201101312221333-0303222220003331) |
| `multi_account_devices` | [multi_account_devices](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-2201201122310002-3113033113330303-2000013221101111-3130301023311011-0200300312133122-2120102323123312-1032031203003212-3101321220102232) |
| `multi_account_devices.account_range` | [multi_account_devices.account_range](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1203011223021230-1031000030220033-3022331122030303-0231322011301201-2013033131232112-0030000322023020-1320020010101230-1303130233121002) |
| `multi_account_devices.device_count` | [multi_account_devices.device_count](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-3102230020132223-2200301012320331-3332300310232223-1111030200201312-2300232211221120-0313220013223122-2120210000033301-1131232021232120) |
| `namespace` | [namespace](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0200223131211132-1002001200220320-2202310220331121-3330112031212010-3132111020210011-0123231231131022-1312123123310231-0011013313011101) |
| `start_time` | [start_time](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1213110011121013-0222110210033002-3331310311222211-3213032121030222-1203332311110130-3321020121233221-3331330113331012-2000110201310210) |

<a id="canonical-1102032323221001-0221233322123221-3230303132011023-0033322310323010-2232202301320220-1331223132023300-0222030001102313-3113222101022233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210)
- filters

<a id="canonical-1333123022222313-0102001101023211-1130120211221032-2331023323122332-0022132303000200-3131023020333302-1020331010300213-2313310133011302"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-0232223033211130-1133313303232000-1101313130330113-1131011310013332-2001123310220002-0221320222033103-2012222011112321-1102230110110022"></a>

### Direct properties for `filters`

- [global_filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1123331132203332-2000203310223103-0011123103333212-3321110311302130-3122310023020331-2022200312111212-0100302221232323-0132123001120231): complete subsection reference.

<a id="canonical-3212231012213002-3312310031133221-0301102322011103-2123031000003101-1022103300323322-1020202032131133-2201101312221333-0303222220003331"></a>

<a id="canonical-1322123101330033-1113022200210311-1333030322032322-1000311020210003-0330133131303103-0121212213313103-1110110310013221-3100030221111003"></a>

#### `filters.region_filter` property

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASIA","CA","EU","US"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

<a id="canonical-1123331132203332-2000203310223103-0011123103333212-3321110311302130-3122310023020331-2022200312111212-0100302221232323-0132123001120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters.global_filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210)
- [filters](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-1102032323221001-0221233322123221-3230303132011023-0033322310323010-2232202301320220-1331223132023300-0222030001102313-3113222101022233)
- filters.global_filters

<a id="canonical-3132103312103100-0220010033300101-0210331330111012-1013221210321023-3313100230002313-0030312100200230-3200020111221221-2333111230003303"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-0333010131101010-3033012220231211-3110211222301031-1300032332212303-3232123201003301-0131201230002210-2223231302332000-2321320310321220"></a>

### Direct properties for `filters.global_filters`

<a id="canonical-2303001032103211-0230300213023002-1132333130130003-3211321331221102-3312122210221102-0001001202012223-1131030012300010-2211332322331021"></a>

#### `filters.global_filters.key` property

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ABSOLUTE","ACTION_TAKEN","AGENT","APPLICATION_NAME","ASN","AS_ORGANIZATION","BOT_COOKIE","BOT_ENDPOINT_POLICY","BOT_REASON","BROWSER_FINGERPRINT","CLIENT_TOKEN","COOKIE_AGE","COUNTRY","DEVICE_ID","ENDPOINT_LABEL","ENDPOINT_NAME","ENDPOINT_POLICY","FLOW","FLOW_CATEGORY","FLOW_LABEL","HEADER_FINGERPRINT","HOST","IP_ADDRESS","IS_ATTACK","KNOWN_BOT_CATEGORY","KNOWN_BOT_CATEGORY_TYPE","KNOWN_BOT_MITIGATION","KNOWN_BOT_NAME","KNOWN_BOT_PROVIDER","METHOD","MOBILE_TRANSACTION_INSIGHT","PATH","PERCENTAGE","PROTECTED_APPLICATION","REFERER","RESPONSE_CODE","SDK_VERSION","SERVER_RESPONSE_CODE","THREAT_TYPE","TIMESTAMP","TRAFFIC_CHANNEL","TRAFFIC_TYPE","TRANSACTION_RESULT","TREND","TRIGGERED_RULE","URL","USERNAME","USER_AGENT","USER_AGENT_FAMILY","USER_AGENT_OS_FAMILY","USER_FINGERPRINT","WEB_TRANSACTION_INSIGHT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-1300011113220103-2110213113103210-0233303200301220-3111121230011301-1001223133110120-3113303220211312-3102131101210113-3332213203103011"></a>

<a id="canonical-2010330121322101-0003022201111012-1200213332120010-3231021320001120-3132203331221102-0323230030330033-0203103210332013-3332003130230033"></a>

#### `filters.global_filters.op` property

Type: `"string"`. Optional.

\[Enum:
IN|NOT\_IN|MATCHES\_REGEX|DOES\_NOT\_MATCH\_REGEX|INCLUDES|DOES\_NOT\_INCLUDE|STARTS\_WITH|ENDS\_WITH\]
Operator for query filter - IN: Filter Operator Specifies that query result includes filter values -
NOT\_IN: Filter Operator Specifies that query result excludes filter values - MATCHES\_REGEX: Filter
Operator Specifies that query result matches filter regular expression - DOES\_NOT\_MATCH\_REGEX: Filter..
Possible values are \`IN\`, \`NOT\_IN\`, \`MATCHES\_REGEX\`, \`DOES\_NOT\_MATCH\_REGEX\`,
\`INCLUDES\`, \`DOES\_NOT\_INCLUDE\`, \`STARTS\_WITH\`, \`ENDS\_WITH\`. Defaults to \`IN\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DOES_NOT_INCLUDE","DOES_NOT_MATCH_REGEX","ENDS_WITH","IN","INCLUDES","MATCHES_REGEX","NOT_IN","STARTS_WITH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3110030201100321-0101230201032322-3023032020013332-0130002230111032-0302130333022000-3312122323202002-0102100101210223-2223212201222131"></a>

<a id="canonical-2320100122121323-2100112203323332-1332012311232233-2123310301000301-1010202011123132-3021030301302201-2012231132320133-0030100120200023"></a>

#### `filters.global_filters.values` property

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-3032111311010121-1003310120003002-3320123213232122-1020212011031201-2003202210303021-0220233030001321-1110223203232233-3203200022222311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `multi_account_devices` properties

Breadcrumbs:

- [xcsh_device_intelligence_multi_account_devices](../data-sources/device_intelligence_multi_account_devices.md#canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000)
- [Property reference](data-sources--device_intelligence_multi_account_devices--reference--group-001.md#canonical-0002211303210131-3301122203222131-3121332000231220-3031103123001111-3332222231232120-3000111203301111-1311330311113231-3322102221001210)
- multi_account_devices

<a id="canonical-2201201122310002-3113033113330303-2000013221101111-3130301023311011-0200300312133122-2120102323123312-1032031203003212-3101321220102232"></a>

Type: `"list"`. Computed.

Distribution of devices by account range buckets.

<a id="canonical-2311202031022011-1202013202210213-2023032210100030-1312321230302210-0322111130032102-0213103122203002-3331032212001220-2032310012120102"></a>

### Direct properties for `multi_account_devices`

<a id="canonical-1203011223021230-1031000030220033-3022331122030303-0231322011301201-2013033131232112-0030000322023020-1320020010101230-1303130233121002"></a>

#### `multi_account_devices.account_range` property

Type: `"string"`. Computed.

Bucket representing the number of accounts linked to a device (e.g., '1', '2-3', '10+').

<a id="canonical-3102230020132223-2200301012320331-3332300310232223-1111030200201312-2300232211221120-0313220013223122-2120210000033301-1131232021232120"></a>

<a id="canonical-2010023013202033-3133200301333223-2032132112001213-1131001323032102-1223120220022011-1332121233113331-1021202310302132-2232102301000322"></a>

#### `multi_account_devices.device_count` property

Type: `"string"`. Computed.

Number of devices in this account range.
