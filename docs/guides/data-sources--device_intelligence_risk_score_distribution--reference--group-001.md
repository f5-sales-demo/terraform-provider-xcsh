---
page_title: "xcsh_device_intelligence_risk_score_distribution reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_risk_score_distribution reference."
---

# xcsh_device_intelligence_risk_score_distribution reference

<a id="canonical-0133122332303212-2021022130012330-2033203113001123-2122022002301131-1113132203202220-2330133310103203-3120012013102030-2301000100213231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- Property reference

<a id="canonical-3123003120212212-2200223031112032-0103330332320112-0332010213001230-0201113113311011-2311211131021202-0010200211112111-2002210010102020"></a>

### Direct properties for `xcsh_device_intelligence_risk_score_distribution`

<a id="canonical-2010011333101100-3200020022213322-2301003120022030-3003201101021031-1220323022331102-1033120302332121-2120220300313300-1113200322220322"></a>

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

- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3100331320123331-0233133313322333-0132210020330022-3310010231110200-0132113120123222-2331023012321020-3031303002332012-1313221122310211): complete subsection reference.

<a id="canonical-3100011331312102-0322101131231302-2333300131123001-0011032320023333-3221003232212312-2110022233312203-2020220130311203-0331310022032222"></a>

<a id="canonical-3231112333022310-2113031010223210-1033003203231300-3232311303332213-0003230012111213-3111123330020210-1031310300200233-2032303031120231"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace. Namespace name.

- [risk_score_distribution](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3013013212003322-2122223133210121-1302130312133021-0330132120321103-1033110031033023-1221330002310030-2111012231022010-0313313031320032): complete subsection reference.

<a id="canonical-1123113011232120-2231100210020320-3000210310033000-1013022323323232-3120010102302020-1301011330121033-2022232000200120-3110302300333231"></a>

<a id="canonical-3303020213230003-0223121100110103-1320212300103211-0310132232103213-2013200311010301-1210002322013123-0201222021000232-2032110313233032"></a>

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

<a id="canonical-3232211110333222-1211320032020111-1000210301032020-0011213111020222-1331232033213021-1220121100223130-1102200303330101-1233220100010120"></a>

### All schema paths for `xcsh_device_intelligence_risk_score_distribution`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `end_time` | [end_time](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-2010011333101100-3200020022213322-2301003120022030-3003201101021031-1220323022331102-1033120302332121-2120220300313300-1113200322220322) |
| `filters` | [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0003302030130222-3132320302232121-0011212113100320-0212133033223133-1013030002000102-2110132121110233-3012301323133100-3111332031222322) |
| `filters.global_filters` | [filters.global_filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1232322103321011-3203232300131011-2332102123110223-3021203200113023-1101223022033101-0231322222203300-1223122131113101-2213310321132023) |
| `filters.global_filters.key` | [filters.global_filters.key](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3300201000233023-1132100210033310-0113313333012210-1103120110020220-0122111233331031-3210223231200322-3133311222331002-0310203030201311) |
| `filters.global_filters.op` | [filters.global_filters.op](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3031033111310033-0022021121132302-1221030103210230-1102132332011103-0100313103313133-1211302312201201-1133310300113321-2113021332220021) |
| `filters.global_filters.values` | [filters.global_filters.values](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0002003200100012-1322131010012001-1123022131121202-2232031003311110-3001010020302131-3013110020110320-3013003313102231-3301311201200311) |
| `filters.region_filter` | [filters.region_filter](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-2331330101003121-2002031232102333-1200020101031210-0103121122021211-3102113201121120-2110313203133022-0331033030111103-2202000100123122) |
| `namespace` | [namespace](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3100011331312102-0322101131231302-2333300131123001-0011032320023333-3221003232212312-2110022233312203-2020220130311203-0331310022032222) |
| `risk_score_distribution` | [risk_score_distribution](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3310321013300132-3301223231230020-3202312013311102-2302332232122101-0303121033031333-0303331322123201-1020301023030301-3303223221300031) |
| `risk_score_distribution.device_count` | [risk_score_distribution.device_count](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1002110201203303-0010002000021033-0200032212201213-2213010230310232-1323020121330121-1230023201121031-0232112322300330-2131230003230200) |
| `risk_score_distribution.risk_rank` | [risk_score_distribution.risk_rank](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0003000311232313-0212120323113332-2311233300101112-3212231111301311-2131233013310022-1031232323121102-2033220202120320-3102100013200210) |
| `start_time` | [start_time](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-1123113011232120-2231100210020320-3000210310033000-1013022323323232-3120010102302020-1301011330121033-2022232000200120-3110302300333231) |

<a id="canonical-3100331320123331-0233133313322333-0132210020330022-3310010231110200-0132113120123222-2331023012321020-3031303002332012-1313221122310211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0133122332303212-2021022130012330-2033203113001123-2122022002301131-1113132203202220-2330133310103203-3120012013102030-2301000100213231)
- filters

<a id="canonical-0003302030130222-3132320302232121-0011212113100320-0212133033223133-1013030002000102-2110132121110233-3012301323133100-3111332031222322"></a>

Type: `"single"`. Optional.

Global Filters. Query Global Filters.

<a id="canonical-2230311121213020-2003121031302323-3321212222321201-1202221220003213-2100311312110212-3203120102323030-3231200300210333-3022133202311311"></a>

### Direct properties for `filters`

- [global_filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3313021023003120-3110032203321233-2303232213102000-0113222312213012-3323110321121020-1212121012023203-0111301230110003-2100021110133131): complete subsection reference.

<a id="canonical-2331330101003121-2002031232102333-1200020101031210-0103121122021211-3102113201121120-2110313203133022-0331033030111103-2202000100123122"></a>

<a id="canonical-1032111111211301-0230020031113313-2311031023202130-1122233203031121-1321130011032232-1231131311020333-3300313203313122-3302221030032212"></a>

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

<a id="canonical-3313021023003120-3110032203321233-2303232213102000-0113222312213012-3323110321121020-1212121012023203-0111301230110003-2100021110133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `filters.global_filters` properties

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0133122332303212-2021022130012330-2033203113001123-2122022002301131-1113132203202220-2330133310103203-3120012013102030-2301000100213231)
- [filters](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-3100331320123331-0233133313322333-0132210020330022-3310010231110200-0132113120123222-2331023012321020-3031303002332012-1313221122310211)
- filters.global_filters

<a id="canonical-1232322103321011-3203232300131011-2332102123110223-3021203200113023-1101223022033101-0231322222203300-1223122131113101-2213310321132023"></a>

Type: `"list"`. Optional.

Global Filters. List of global filters.

<a id="canonical-1302120130122332-3331222232221032-2132003110220211-1123022313002332-2212203311230331-3110122130233120-1120022032213001-2330020322102203"></a>

### Direct properties for `filters.global_filters`

<a id="canonical-3300201000233023-1132100210033310-0113313333012210-1103120110020220-0122111233331031-3210223231200322-3133311222331002-0310203030201311"></a>

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

<a id="canonical-3031033111310033-0022021121132302-1221030103210230-1102132332011103-0100313103313133-1211302312201201-1133310300113321-2113021332220021"></a>

<a id="canonical-1231012013013021-0003030100301310-3100320221130330-1330320022020213-0202301033302013-1203322030030000-1103330232002120-0131132322333010"></a>

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

<a id="canonical-0002003200100012-1322131010012001-1123022131121202-2232031003311110-3001010020302131-3013110020110320-3013003313102231-3301311201200311"></a>

<a id="canonical-1233103210222220-0323011232120310-3121321003332301-2302032001330321-2201312123010213-2113210312031110-0112031203201011-2100313202113320"></a>

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

<a id="canonical-3013013212003322-2122223133210121-1302130312133021-0330132120321103-1033110031033023-1221330002310030-2111012231022010-0313313031320032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `risk_score_distribution` properties

Breadcrumbs:

- [xcsh_device_intelligence_risk_score_distribution](../data-sources/device_intelligence_risk_score_distribution.md#canonical-0220221212202100-3300002033312230-0123021230311320-1120110032303212-1000013303201203-1310101222000102-0330103220003033-0002223100110032)
- [Property reference](data-sources--device_intelligence_risk_score_distribution--reference--group-001.md#canonical-0133122332303212-2021022130012330-2033203113001123-2122022002301131-1113132203202220-2330133310103203-3120012013102030-2301000100213231)
- risk_score_distribution

<a id="canonical-3310321013300132-3301223231230020-3202312013311102-2302332232122101-0303121033031333-0303331322123201-1020301023030301-3303223221300031"></a>

Type: `"list"`. Computed.

Distribution of devices across risk ranks.

<a id="canonical-3032122311232232-2003110203112031-3211111021002320-3110231032022230-3233330021230111-2300133021333211-3112031301122212-1223212133311212"></a>

### Direct properties for `risk_score_distribution`

<a id="canonical-1002110201203303-0010002000021033-0200032212201213-2213010230310232-1323020121330121-1230023201121031-0232112322300330-2131230003230200"></a>

#### `risk_score_distribution.device_count` property

Type: `"string"`. Computed.

Device Count. Number of devices in this risk rank.

<a id="canonical-0003000311232313-0212120323113332-2311233300101112-3212231111301311-2131233013310022-1031232323121102-2033220202120320-3102100013200210"></a>

<a id="canonical-3200021231002131-2233212033013122-2021330321022112-3003322322111221-3000012122301112-3101202230200220-0331120030131212-3010101001211221"></a>

#### `risk_score_distribution.risk_rank` property

Type: `"string"`. Computed.

Risk rank label (e.g., low, medium, high).
