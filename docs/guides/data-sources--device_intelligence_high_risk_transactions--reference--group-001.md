---
page_title: "xcsh_device_intelligence_high_risk_transactions reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_high_risk_transactions reference."
---

# xcsh_device_intelligence_high_risk_transactions reference

<a id="canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- Property reference

<a id="canonical-1303021211313322-3310311032010010-0301001320322122-1210100233302033-1320333212222211-3123130123330313-2222300110133001-0010111203032213"></a>

### Direct properties for `xcsh_device_intelligence_high_risk_transactions`

<a id="canonical-1123223300303123-0333002112001013-2031302321021200-3230012000203000-0101210320031321-1333233022101022-0220311233102130-3113113021010110"></a>

#### `end_time` property

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

- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-0120311103201323-2121130021000132-1100232131301313-3310103321303130-0301013313302333-3012120221111210-2320103212303230-3320322231310011): complete subsection reference.

<a id="canonical-3111110130122101-2032100111003133-1011221111203313-0320322222031302-2101332003301200-1202002311331321-0111133222302131-2311121012200211"></a>

<a id="canonical-0230300331330100-3323003131023032-3313103001123010-1210313333003012-2202001032002132-0111200321010003-0322201230133300-1021303001200130"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace. Namespace name.

<a id="canonical-1101221233331330-2020201120012202-2120023230110331-3120000013030212-1332123020331203-2313213312031001-3132032332013100-2213202231101001"></a>

<a id="canonical-0212113333220321-2010220112132012-1231123211013020-1020132333320111-3313232323212122-1011322102230311-3210331310012300-0300222201303313"></a>

#### `start_time` property

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

- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1302300303113012-0021110302030233-3333102210111132-1313202330311232-3220230301100221-3131210133121012-0223331023211203-0301303003330001): complete subsection reference.

<a id="canonical-3030012103231300-3201311313001013-2000003021030001-1010222301300133-3011021330023321-2102222233200120-0100221010331031-0132201230203330"></a>

### All schema paths for `xcsh_device_intelligence_high_risk_transactions`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1123223300303123-0333002112001013-2031302321021200-3230012000203000-0101210320031321-1333233022101022-0220311233102130-3113113021010110) |
| `filters` | [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-3022312302130113-3102112233233133-0221120031110312-3303022131312123-2011033213303133-2013222020130332-3220321313210012-1333120321033010) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-3300222003302303-2133320133112023-1010213020232013-3332220210012210-2330100321023122-0000012123212121-2001031121123103-1233331133012010) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-0301003130220123-0013310213111111-2022332302303232-1123222001222311-0032112213011020-2323120200111232-2310011211313120-3001011301011301) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-3200312203322233-0303011112220321-2202230230021202-0110222313220212-2211323031333333-0200030320332231-0121113021320002-3133113001032200) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2221213322300001-0113032220032121-3012222011013312-1313020012000133-1112100112021122-1121200022020212-2132303000302121-1100013023020212) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-3321330000211321-1213201001200203-2320131201221210-0333110120021032-2230202033211033-0222223300011021-3330302123110131-2223101103003210) |
| `namespace` | [namespace](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-3111110130122101-2032100111003133-1011221111203313-0320322222031302-2101332003301200-1202002311331321-0111133222302131-2311121012200211) |
| `start_time` | [start_time](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1101221233331330-2020201120012202-2120023230110331-3120000013030212-1332123020331203-2313213312031001-3132032332013100-2213202231101001) |
| `time_series_results` | [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2013210322222131-1131203030321032-0102210033113112-2013033010132220-0332320121312133-3320211010112332-1110033010001021-3022230030330222) |
| `time_series_results.series_key` | [time_series_results.series_key](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2221102113200133-3113320233013310-3113013112311110-1231122123023002-2213213100131300-0113203110102331-2100021121120123-2211102203122333) |
| `time_series_results.time_series` | [time_series_results.time_series](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2122122000112211-2131323031330302-1010012033212022-1330331120203112-0213000000200111-1113102010322210-3021312321132210-2213333233110222) |
| `time_series_results.time_series.timestamp` | [time_series_results.time_series.timestamp](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2311330000310330-1123010030132013-0233301113232002-2033133032101220-2323201232033312-2102032032102123-1122320033230233-0132022333013022) |
| `time_series_results.time_series.value` | [time_series_results.time_series.value](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1312313012002332-2130213112123303-2123033102003200-3202331001320002-1131110023212110-3112202212013011-1120112123303300-0131130303112021) |

<a id="canonical-0120311103201323-2121130021000132-1100232131301313-3310103321303130-0301013313302333-3012120221111210-2320103212303230-3320322231310011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222)
- filters

<a id="canonical-3022312302130113-3102112233233133-0221120031110312-3303022131312123-2011033213303133-2013222020130332-3220321313210012-1333120321033010"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-3110211211002131-3122201033122203-2020022310021013-1200332221130303-1203212321322322-2112200123030201-1003213202320110-0201213020201103"></a>

### Direct properties for `filters`

- [global_filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1031110220102322-1132022331000100-2321101123013321-3300210203002202-3031033010003200-2001000000021103-0323130000102223-1013330033032233): complete subsection reference.

<a id="canonical-3321330000211321-1213201001200203-2320131201221210-0333110120021032-2230202033211033-0222223300011021-3330302123110131-2223101103003210"></a>

<a id="canonical-1003010211232300-3320010133231002-2310002210331223-0011210312213233-0231032213133021-3100112123000203-1210210333330203-0121230033201103"></a>

#### `filters.region_filter` property

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

<a id="canonical-1031110220102322-1132022331000100-2321101123013321-3300210203002202-3031033010003200-2001000000021103-0323130000102223-1013330033032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters.global_filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222)
- [filters](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-0120311103201323-2121130021000132-1100232131301313-3310103321303130-0301013313302333-3012120221111210-2320103212303230-3320322231310011)
- filters.global_filters

<a id="canonical-3300222003302303-2133320133112023-1010213020232013-3332220210012210-2330100321023122-0000012123212121-2001031121123103-1233331133012010"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-1013322231113012-3033320002112303-0033032313033220-0301202310022112-1332223332031101-1102300012113330-3113000110013010-0331022022022211"></a>

### Direct properties for `filters.global_filters`

<a id="canonical-0301003130220123-0013310213111111-2022332302303232-1123222001222311-0032112213011020-2323120200111232-2310011211313120-3001011301011301"></a>

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

<a id="canonical-3200312203322233-0303011112220321-2202230230021202-0110222313220212-2211323031333333-0200030320332231-0121113021320002-3133113001032200"></a>

<a id="canonical-0302102320031312-3030131311321001-1310220000221221-3333023003222212-1210000123031113-0010030200120022-3031102120212322-3132331131010000"></a>

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

<a id="canonical-2221213322300001-0113032220032121-3012222011013312-1313020012000133-1112100112021122-1121200022020212-2132303000302121-1100013023020212"></a>

<a id="canonical-1023330000003000-3310223022213212-1113320031111101-2033022002201302-1003222313321330-2221022103333022-0302200202232130-2221122033131000"></a>

#### `filters.global_filters.values` property

Type: `["list", "string"]`. Optional.

Values. An unordered list of filter strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 64),
}
```

<a id="canonical-1302300303113012-0021110302030233-3333102210111132-1313202330311232-3220230301100221-3131210133121012-0223331023211203-0301303003330001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `time_series_results` properties

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222)
- time_series_results

<a id="canonical-2013210322222131-1131203030321032-0102210033113112-2013033010132220-0332320121312133-3320211010112332-1110033010001021-3022230030330222"></a>

Type: `"list"`. Computed.

Collection of time series grouped by a series key.

<a id="canonical-2310220002002123-1110111120230300-0332022311031313-3322100033231333-3000333201032131-3023300013320201-2310133011130022-1300111023310231"></a>

### Direct properties for `time_series_results`

<a id="canonical-2221102113200133-3113320233013310-3113013112311110-1231122123023002-2213213100131300-0113203110102331-2100021121120123-2211102203122333"></a>

#### `time_series_results.series_key` property

Type: `"string"`. Computed.

Identifier for the time series (e.g., 'total', 'high\_risk').

- [time_series](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1103330333220221-1330301211031103-2300031012001033-1321100232023320-3102010201013032-3301230000320123-0013032102112133-0131033330033331): complete subsection reference.

<a id="canonical-1103330333220221-1330301211031103-2300031012001033-1321100232023320-3102010201013032-3301230000320123-0013032102112133-0131033330033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `time_series_results.time_series` properties

Breadcrumbs:

- [xcsh_device_intelligence_high_risk_transactions](../data-sources/device_intelligence_high_risk_transactions.md#canonical-2323323100323133-2031230011120020-3320233202222001-1321322211033211-2221211100303331-0220103011323320-1220023022213033-0133011203321103)
- [Property reference](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-2030120313301233-0221233133100221-0201300032331331-1120303000311221-2203033322111232-1001320313023121-2003030130133011-0112310120123222)
- [time_series_results](data-sources--device_intelligence_high_risk_transactions--reference--group-001.md#canonical-1302300303113012-0021110302030233-3333102210111132-1313202330311232-3220230301100221-3131210133121012-0223331023211203-0301303003330001)
- time_series_results.time_series

<a id="canonical-2122122000112211-2131323031330302-1010012033212022-1330331120203112-0213000000200111-1113102010322210-3021312321132210-2213333233110222"></a>

Type: `"list"`. Computed.

Sequence of timestamped values for this series.

<a id="canonical-3213030223230212-2032023013012000-0103320303000223-1031000101013303-3330230010023031-2211011112303010-1331311320223333-3103210231021110"></a>

### Direct properties for `time_series_results.time_series`

<a id="canonical-2311330000310330-1123010030132013-0233301113232002-2033133032101220-2323201232033312-2102032032102123-1122320033230233-0132022333013022"></a>

#### `time_series_results.time_series.timestamp` property

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

<a id="canonical-1312313012002332-2130213112123303-2123033102003200-3202331001320002-1131110023212110-3112202212013011-1120112123303300-0131130303112021"></a>

<a id="canonical-0312130203330221-1131232010001113-3132221212222121-2030123120322113-3202322023020103-0122001113003112-1103122133312331-2113222212320330"></a>

#### `time_series_results.time_series.value` property

Type: `"string"`. Computed.

Value. Value observed at the given timestamp.
