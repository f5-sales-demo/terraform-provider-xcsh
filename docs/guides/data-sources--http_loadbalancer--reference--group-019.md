---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1121200303002003-1231323131233030-0223313122113033-0310223120311101-3013330213121112-3223101030231222-2213302113310223-0131211101133212"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-1301121110102333-3003232202003310-0323313222301011-2132310311110310-0333233010320332-1303111322213100-3200210013113231-2031111001023021"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2202322213122002-0012033032233213-0030122021322013-3223102103300330-1323312223001023-0112231202203101-3133223023122033-1220233132213220): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1112032312121213-0223300322233013-1131303210001112-2002031310303110-2310232330302000-1132101213302010-2123323323313131-3001322312032112): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-2030103111230230-3220223102021020-2123201320222030-3002330311130323-0013101011122133-0321021212020230-1212012033213331-0333130322103031): complete subsection reference.

<a id="canonical-2202322213122002-0012033032233213-0030122021322013-3223102103300330-1323312223001023-0112231202203101-3133223023122033-1220233132213220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1203100033001003-2023033130302013-2020032021213230-1033003332123122-2203313230120030-3322310200311122-2212202211232313-1233031023330101"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112032312121213-0223300322233013-1131303210001112-2002031310303110-2310232330302000-1132101213302010-2123323323313131-3001322312032112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0100231133331333-1003133303001111-1002222012133210-0030100022123222-3033310202333212-3311203113313332-1321121303231003-2122022313103200"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030103111230230-3220223102021020-2123201320222030-3002330311130323-0013101011122133-0321021212020230-1212012033213331-0333130322103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1100112311121310-1111011320313012-1102212131122030-0022023012111331-1212312231000301-3000202102120203-1320132202013023-1210111022200332"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103132123031212-1312011311132313-1101103331230311-0211210100021021-0223211311301221-2201203120330131-0323223220202031-2113302232301212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-3032332011020133-0202230002131001-2120102202310001-1023233133030213-2003100020033220-3320202030203220-3002212120000331-3022002120130101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011132011130333-0210213032201022-0233112013130221-1333101332310100-3311123321233200-3302001330230303-1003013201132323-0021222312201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0110022001120101-2313001121311312-1030320012233132-3103202132121130-1000310302113232-1032111101020033-3322221023021120-1233231021100002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102301213102300-3130211303123110-1211031100221301-0101130110032303-3233220303001111-3333311120020322-3000202300233331-2213001132011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.no_mtls

<a id="canonical-3102331103101200-1300200020133012-3212201023102022-1033211130112302-0111213322130103-0233311301011312-2202033231000203-3231303332020023"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311320301322302-1303220013302033-0332111200020023-1101002220032022-2331300211113301-0130223132030323-1230211323100313-3210102020111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-1321112301132100-3311122130222111-1030000333030100-2021301202103013-2233220213212003-3313332221010022-0010031220303131-2301302131022111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221203212221022-2312321133113202-1222303012100333-2123103030313231-1012021303103113-3232300033113221-1000231113031103-1202131222211332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.pass_through

<a id="canonical-0201223000223123-2113111120003131-1233012100023311-3210301320110223-1121212130032223-1212331021310210-1020112320000113-3020132023001321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.tls_config

<a id="canonical-1303230000123212-0021300330000120-1323323023203020-1103122301012131-0032202000033001-1210020010202233-0333213200223212-3023003032320310"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-2200312213221312-3212221122333302-2222331300223230-2131010120200320-1232331202213211-3131010301021010-0013011212110021-2101122311023230"></a>

### Direct properties for `https_auto_cert.tls_config`

- [custom_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3321303330313313-1023102100223210-0012100231302202-0313102310302120-0232233111333113-0013230032202120-3011301313332221-2320312110030121): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311200230002203-1203320312313000-0133330003131232-1221012211132131-3002320003012210-3111203202112020-2012321231102301-0200133112313231): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-3133012211221031-2320213030030231-3002213122221231-3223321103200220-3333232002031112-0033232130010001-3010101230033111-2331300100130002): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320333312023320-3223013123032113-2113232310323110-2013111100113321-1000101020100330-1322323010122013-2120003020300200-0132323201132013): complete subsection reference.

<a id="canonical-3321303330313313-1023102100223210-0012100231302202-0313102310302120-0232233111333113-0013230032202120-3011301313332221-2320312110030121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.custom_security

<a id="canonical-2121302111320221-2012230300301012-2311303023023203-1103002100023222-3020323101311201-2022130033101333-2302123212332100-2111033211131313"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2012300333031223-1312130013102103-1311221022232103-1122131212003333-1332000030322023-3113232111233210-2310002213031310-2023222010030320"></a>

### Direct properties for `https_auto_cert.tls_config.custom_security`

<a id="canonical-2312122002111323-1223301230013000-1302200210223223-1302222010310231-0223132032132222-3220322131023301-1011101121201123-1313303330210321"></a>

#### `https_auto_cert.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3003120113200102-3002111313221130-0133033001233200-3011111231303232-1213022101003011-2301203110110000-1123330323013310-2313110332000130"></a>

<a id="canonical-0010222111203131-1221021100232111-0301103102310323-1311103201123100-3233221331010112-1302323231311032-3331122110031212-3230101131120312"></a>

#### `https_auto_cert.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130201201321330-2311101220300322-0010320303322202-3020331330020103-0313320332031302-1332001113201232-0100322232130021-3033122221211221"></a>

<a id="canonical-3311110332330122-3100033222013100-2332223310331220-2032123122001213-2121221012121132-0231202110331111-1121021230201023-1300202001120030"></a>

#### `https_auto_cert.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3311200230002203-1203320312313000-0133330003131232-1221012211132131-3002320003012210-3111203202112020-2012321231102301-0200133112313231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.default_security

<a id="canonical-0110221032200233-0123030013302202-2222311311010320-2301021000313300-3102233231330121-3020302203200213-0310331222003021-2102200010332102"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133012211221031-2320213030030231-3002213122221231-3223321103200220-3333232002031112-0033232130010001-3010101230033111-2331300100130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.low_security

<a id="canonical-3031000220033311-1311330330220321-1132202101203223-1021112003302031-2211203020123203-2111113021312212-0033201203132121-2311112023003112"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320333312023320-3223013123032113-2113232310323110-2013111100113321-1000101020100330-1322323010122013-2120003020300200-0132323201132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132)
- https_auto_cert.tls_config.medium_security

<a id="canonical-3002200010300312-1320300313121032-3231232210323321-3013330020232213-2120002311330010-2202330323320321-3031312320231020-3221303220120130"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.use_mtls

<a id="canonical-1332230210300023-0212122101223322-2213010231110222-3003011321323021-1120110102101332-1210100313330022-2022320011011210-3333230221131233"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-2213131323122132-1211202321233130-1110203012033300-3100211033202121-0232033231020033-2101213332332133-2312110200203232-0013032121103233"></a>

### Direct properties for `https_auto_cert.use_mtls`

<a id="canonical-2120103111000201-1101032320311321-2220213112202003-1231211112010220-1302232113233012-0203123333211201-0123310021003032-0102101123212302"></a>

#### `https_auto_cert.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3022211310311000-2113002323211002-3101033102103010-3332313030200030-2222011331300210-0333231203022020-1133233222102312-3022332112023033): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-019.md#canonical-3031122111221032-3301130300001112-3312231022102031-0132103220103110-3030131110101030-1032113130131001-0002230301223303-1013210033303222): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-019.md#canonical-2110202012002322-1331000322220012-0322003302202201-1202301120221220-2203322233330010-3232030010201113-3002212331002320-1322000223330312): complete subsection reference.

<a id="canonical-1323300012312011-1213223312232322-2321231323120111-0232232021321030-2200313032022112-2332300113322011-0302000213333300-0123230131320220"></a>

<a id="canonical-1301210033301112-3301312300100020-3310031031220023-3231333030311200-1211222322210323-2112103001003230-2201230012120003-3030233122032023"></a>

#### `https_auto_cert.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-019.md#canonical-1323300021111131-0210300012100323-1030213301331200-2010322123033110-3012201130320312-3303210030320222-2213320012331332-1211233202213302): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-2320103323203311-2032012110110030-2033213110213001-3023133310030011-1310311100100221-3213202133022321-1103031113223332-1323011332023031): complete subsection reference.

<a id="canonical-3022211310311000-2113002323211002-3101033102103010-3332313030200030-2222011331300210-0333231203022020-1133233222102312-3022332112023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.crl

<a id="canonical-3102203323230331-2332331121111301-1031021010032123-2300331230310203-1222022330101121-0323133203101230-0222130213103103-0301011102032323"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032103131132202-0010003331100221-2332222332112022-1223133012202313-3031210200303233-1122300023201302-0010020211230101-0001020003301331"></a>

### Direct properties for `https_auto_cert.use_mtls.crl`

<a id="canonical-1233201310332222-2221100001120323-1110223122023300-1133322230022011-1210211211023130-3012333233123133-1121102021133312-3121012000020333"></a>

#### `https_auto_cert.use_mtls.crl.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2103303330320031-0300230330132101-0302013113033022-1033110220302020-3312020221310222-1330233211203310-1220031212013001-0330232300111100"></a>

<a id="canonical-0121103120302312-0312213130323222-0001321210102323-0203130223120303-2103113310220112-3300330202311130-3321200221302012-0300320233210111"></a>

#### `https_auto_cert.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0033223200330302-1110121331232322-1202010122133133-3310022323321310-2233331303213132-2031312313311312-0021323223030203-0222212332201322"></a>

<a id="canonical-3023302311010321-2101032123013111-2322233333003012-1332100002113222-2133111123330222-3120203221122232-3203232202112231-1300211113031002"></a>

#### `https_auto_cert.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3031122111221032-3301130300001112-3312231022102031-0132103220103110-3030131110101030-1032113130131001-0002230301223303-1013210033303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-0112313230133003-1022013212110230-3120323121133100-3202212220310113-0100231332310033-1323020012033203-1232113321333232-1233031323322132"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110202012002322-1331000322220012-0322003302202201-1202301120221220-2203322233330010-3232030010201113-3002212331002320-1322000223330312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-1100231002200303-3331321100303210-0000032312113220-2003210220131330-3213130311333000-3323201330013313-3013230002233120-2323130010333130"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0311112330100300-2021213103233211-0211001220311002-2010021320132112-3010013303222212-0102030202312231-0331112330213311-3203321000123100"></a>

### Direct properties for `https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-0110311233213233-0100322222011201-0223322012022021-0022011031333233-1313202010331221-2201123000232313-2202102021301302-1322032103010002"></a>

#### `https_auto_cert.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0123123101312031-3022300111202230-1230330311121213-3032110023123113-0301012213012032-3232302220332131-2302020133011210-1203013203313012"></a>

<a id="canonical-1332121112200111-0130030033300131-3022313212101133-0313221312221012-0000101012200323-0221002321323110-2011130302132003-1123220221230000"></a>

#### `https_auto_cert.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3302300032003133-1110221312331012-3000020301331212-3203012123210202-3100331312021021-1120330023322232-0023232332201321-0022002203311012"></a>

<a id="canonical-1112333130031022-1222221000321120-3121122233031020-3002130322011021-0100200322322301-3010120102203321-3301230332223031-3021201021221223"></a>

#### `https_auto_cert.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1323300021111131-0210300012100323-1030213301331200-2010322123033110-3012201130320312-3303210030320222-2213320012331332-1211233202213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0211231130230123-2232111000322000-2010130333011003-1120020001320111-3321011130302113-3022022003032023-0313003001331000-0333330313212130"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320103323203311-2032012110110030-2033213110213001-3023133310030011-1310311100100221-3213202133022321-1103031113223332-1323011332023031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3222033120312330-0220010302223110-3321013122203130-1222101301333322-1201332131110030-2011012200003302-0132223022330220-2020331330232323"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233101301210011-2131022021232032-1121330213103121-2311021301332330-0000000123012313-3330233301323120-1202201133220212-0003333213310231"></a>

### Direct properties for `https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-3003200023111201-0000122322231303-3320122301312123-3213032301203302-3012031333332300-3202110103221133-0102001013320202-2131013231303010"></a>

#### `https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3322121320310111-0021311031133102-3002213332023220-1023112102022302-3222311022023111-1222120222313303-1310032321220233-2003030001231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- js_challenge

<a id="canonical-1330313013020313-0000301213230310-2113112320001221-3102132110313013-3033321333122212-3213022133202312-1321331102121220-0332330321122003"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201123310322022-0120201111011021-0131311012133321-1303210002312131-1003221301313011-3223120033021300-1232033202113221-1201323120021031"></a>

### Direct properties for `js_challenge`

<a id="canonical-3102210100303112-1020303021310313-2133302230122021-0030133322212223-2112033032132323-0231022220122101-3223033211310120-1010211100002121"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2322220013032133-3301003321120221-1102202030101101-1210303000213001-1323100103213320-2220110321031000-3031100201103013-2223023021000313"></a>

<a id="canonical-2121210011120231-1233022001120302-1210121121222200-3113213013100011-1211023212132201-0002202032313222-1201112033130211-1003023200113121"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3032322122120211-2321313331012022-2313333033223322-0221212230003332-2122011202201200-2032300032312113-3110023321303313-0013121333103132"></a>

<a id="canonical-0333000123111333-1013211202023121-2032111012203131-1022032120331321-0322332020020230-2211233022021120-0012110222010211-0221333302011322"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- jwt_validation

<a id="canonical-2333203013123213-3021320330302333-0312333311010330-3330210123032031-2133202033212130-1313333232130123-2011322310210232-0122201112222231"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

<a id="canonical-1001103022130031-0101011122330022-2001222323012320-1303303003023022-3201023112313012-1321011131132302-1001111120021202-0030022332120110"></a>

### Direct properties for `jwt_validation`

- [action](data-sources--http_loadbalancer--reference--group-019.md#canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101): complete subsection reference.

- [authorization_server](data-sources--http_loadbalancer--reference--group-019.md#canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022): complete subsection reference.

- [jwks_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300): complete subsection reference.

- [mandatory_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-1003113313333323-2312032022332312-2213233012232010-3132202210222310-1022032321220123-3210013112210001-2012002222033322-0232120300110111): complete subsection reference.

- [reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201): complete subsection reference.

- [target](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321): complete subsection reference.

- [token_location](data-sources--http_loadbalancer--reference--group-019.md#canonical-2101012003211203-1333223011213101-1311003002102303-2000301122201222-1111130223313030-0211031013100112-2032203303233133-3303103031322301): complete subsection reference.

<a id="canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.action

<a id="canonical-2322133200003120-0020322322031123-0112011202201012-2132213001020232-1002201123203303-2222321031110231-3003001213332310-0320110230013333"></a>

Type: `"single"`. Computed.

Action

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

<a id="canonical-1100113233210203-1201203213113120-1100011112031102-3023030302323323-0203123200033111-0210301310023103-1002031100323321-1020000310112202"></a>

### Direct properties for `jwt_validation.action`

- [block](data-sources--http_loadbalancer--reference--group-019.md#canonical-3310212213010012-1310313200231023-1002132133301020-2010130313022021-2211022313021303-3012312000013030-2123233331213130-0213302023010323): complete subsection reference.

- [report](data-sources--http_loadbalancer--reference--group-019.md#canonical-1232021203130011-1200023112011311-2022031033310210-0111312331122321-2101232202012233-2322330321033131-2003100213012100-3022220031123311): complete subsection reference.

<a id="canonical-3310212213010012-1310313200231023-1002132133301020-2010130313022021-2211022313021303-3012312000013030-2123233331213130-0213302023010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.action](data-sources--http_loadbalancer--reference--group-019.md#canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101)
- jwt_validation.action.block

<a id="canonical-0123131313030303-1202000111022011-0331202100330322-3113133131203300-3311001000022211-3221101022103332-3213223331330330-1021203033022221"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232021203130011-1200023112011311-2022031033310210-0111312331122321-2101232202012233-2322330321033131-2003100213012100-3022220031123311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.action](data-sources--http_loadbalancer--reference--group-019.md#canonical-0330000101210233-3210012112023210-0100120200333133-1031102031030203-3332320022130222-1321213230203003-1303301123013331-1033300310312101)
- jwt_validation.action.report

<a id="canonical-3132020310313202-3130122132231021-0010221220113101-2002120201130103-1113003303002031-0100333021021022-0002302333323033-3121121311122203"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.authorization_server

<a id="canonical-3122030100112133-1022021311221223-3033231302322130-2221121221123100-3201310123322233-1230233020103222-0121211313003220-1103331221132212"></a>

Type: `"single"`. Computed.

Reference to Authorization Server object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3333030201112220-3210121122323011-1130300321231100-3110110111123101-3120310303100322-3230000102221012-2311333023333322-0200112000103232"></a>

### Direct properties for `jwt_validation.authorization_server`

- [authorization_servers](data-sources--http_loadbalancer--reference--group-019.md#canonical-0110023200312331-3013233310220320-2220002000211123-1120103200210113-0122233131131331-3202001121301301-1332120010120010-2011200031023032): complete subsection reference.

<a id="canonical-0110023200312331-3013233310220320-2220002000211123-1120103200210113-0122233131131331-3202001121301301-1332120010120010-2011200031023032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server.authorization_servers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.authorization_server](data-sources--http_loadbalancer--reference--group-019.md#canonical-3323202302020011-1013203000333112-1211331220023232-3232211311013021-0331110330133333-1013322200202110-2010313023231030-3213031121002022)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-2201121300001302-2112303131102201-3003320331222311-1310313301003301-3130311031231302-2310033233113021-2011221032310123-2333200121033211"></a>

Type: `"list"`. Computed.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2321003331030210-3021200223111102-1111101122023321-1133103000030331-2130203301223220-1112111010121323-0331021131132331-0213200013111022"></a>

### Direct properties for `jwt_validation.authorization_server.authorization_servers`

<a id="canonical-0132200223013223-1010312000010222-0102133020102302-3100210113031232-1011020323002202-0233133311233213-0122001100333133-0103212212022332"></a>

#### `jwt_validation.authorization_server.authorization_servers.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3003103020122123-2133133302303302-1031323330003330-2023010222111321-3023011100130100-0313303103301300-2233213302221301-2133320111132213"></a>

<a id="canonical-0003111120120220-2020102033012322-3221032222230033-3102103123021130-2120033010010332-3120011202212223-0222332313303132-3120000021102333"></a>

#### `jwt_validation.authorization_server.authorization_servers.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1030000130110022-0230330231220132-2013222302320313-0333210311130022-2033113013210220-2110221300111210-0010302210001312-1020001002332223"></a>

<a id="canonical-0122331333223330-1020213231012123-3231131301033130-0001100002231320-2210213030113132-0132332231012332-2110000203230030-0110303313102121"></a>

#### `jwt_validation.authorization_server.authorization_servers.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1221233210212303-1303230013013133-2330013321313313-2331033212023133-0023202113210300-1000220111211222-1323110320030023-2001232301101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.jwks_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.jwks_config

<a id="canonical-3213232130222003-2112003213211322-0333032013012131-1100122333211201-3113231012011332-0013333213102233-3110320122311313-1030003230031102"></a>

Type: `"single"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301321111011310-2001300112321220-0330003320131330-0003231031211230-1100101203223101-1002012311011201-0323202221210332-0322002022320211"></a>

### Direct properties for `jwt_validation.jwks_config`

<a id="canonical-2302031303122320-1001110321013020-1031113000030001-1101012333001112-1221031320112323-0312132103331023-3213311020131101-3032303302332120"></a>

#### `jwt_validation.jwks_config.cleartext` property

Type: `"string"`. Computed.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1003113313333323-2312032022332312-2213233012232010-3132202210222310-1022032321220123-3210013112210001-2012002222033322-0232120300110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.mandatory_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.mandatory_claims

<a id="canonical-3103213112223113-2213121033303313-3011313101232030-2313321131020211-3100001230122023-1112232012331233-2012102000122032-3122021031030220"></a>

Type: `"single"`. Computed.

Configurable Validation of mandatory Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2312022310200012-1000332231021101-1100310220032213-2310020020210320-1302031103330300-2130210113332021-1333300313201303-3010112202002100"></a>

### Direct properties for `jwt_validation.mandatory_claims`

<a id="canonical-0020323000000311-0212301312130023-0022021222132220-0301012103113333-1202002000230112-1122033001320112-0100230303200133-1233122010210131"></a>

#### `jwt_validation.mandatory_claims.claim_names` property

Type: `["list", "string"]`. Computed.

Claim Names. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.reserved_claims

<a id="canonical-2113301102103132-3331102333200000-3233002112130311-2112213330101002-0201133332323312-0120320021023000-3020011202003002-0313311003001001"></a>

Type: `"single"`. Computed.

Configurable Validation of reserved Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

<a id="canonical-2320333120120123-3132232321210122-1313033230133331-2200232102012222-2113301211302021-3010310021030323-3210023320330121-0032100233011033"></a>

### Direct properties for `jwt_validation.reserved_claims`

- [audience](data-sources--http_loadbalancer--reference--group-019.md#canonical-1032313201211312-2112212220231322-2000200202002013-1212001133013131-1202000121202010-3323133332031101-1220112120121330-2223121113123233): complete subsection reference.

- [audience_disable](data-sources--http_loadbalancer--reference--group-019.md#canonical-3031023031030233-2100312032300111-2330311112001331-2202321232311020-0311322330211103-1103300201201030-1320213302302223-3232300330013301): complete subsection reference.

<a id="canonical-3331331322323030-2013121300200012-2220021332213000-2133302013032232-0011013003132303-3301100203300120-0012303002002101-2311232132211102"></a>

<a id="canonical-3020133012231021-1022332102100231-2301300103033131-3213011002012100-1323022331331012-2221022312102122-1232121121202122-3230000201033320"></a>

#### `jwt_validation.reserved_claims.issuer` property

Type: `"string"`. Computed.

Exact Match. Exclusive with \[issuer\_disable\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [issuer_disable](data-sources--http_loadbalancer--reference--group-019.md#canonical-1302300002133211-1221002122012232-0213203122221131-2031102103003202-0032032303211222-3002110112011213-0213333030031032-1112101121302302): complete subsection reference.

- [validate_period_disable](data-sources--http_loadbalancer--reference--group-019.md#canonical-0203021202300321-0220001013120311-0212220000303131-2003002012310113-0112331120030000-0101213321200201-1330001203121112-3112220220220211): complete subsection reference.

- [validate_period_enable](data-sources--http_loadbalancer--reference--group-019.md#canonical-1102000310022031-1003100031223303-3232321121033022-3201311212221233-1130111221220012-1110120210031122-3122221313103221-1030021331313210): complete subsection reference.

<a id="canonical-1032313201211312-2112212220231322-2000200202002013-1212001133013131-1202000121202010-3323133332031101-1220112120121330-2223121113123233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- jwt_validation.reserved_claims.audience

<a id="canonical-0312033123222223-1202201231213222-2010110031312313-0311212130231201-3223112300110101-1012313013320001-1200133131112012-0321033113000232"></a>

Type: `"single"`. Computed.

Audiences

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3002123300322110-0010210233103311-0000131001120120-0121330100230030-3211120022213001-2033022303021303-3102122302203313-0213201100212312"></a>

### Direct properties for `jwt_validation.reserved_claims.audience`

<a id="canonical-0133200312312120-2221112031302123-3001211202000301-2122322212310301-2103311002210023-1112320113210002-3312023302210030-3002330020323300"></a>

#### `jwt_validation.reserved_claims.audience.audiences` property

Type: `["list", "string"]`. Computed.

Values. Configuration parameter for audiences

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3031023031030233-2100312032300111-2330311112001331-2202321232311020-0311322330211103-1103300201201030-1320213302302223-3232300330013301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.audience_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-2320300313120230-3221221333321213-2222133301320203-2313132320023333-1121201133312301-0312022031310210-3100013123303130-0201213130333201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for audience disable.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302300002133211-1221002122012232-0213203122221131-2031102103003202-0032032303211222-3002110112011213-0213333030031032-1112101121302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.issuer_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-2312100210122000-3232112213021000-1202113031111031-0211031010230213-3203221101211012-2021112212103202-3031002100220113-0321101032031120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for issuer disable.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203021202300321-0220001013120311-0212220000303131-2003002012310113-0112331120030000-0101213321200201-1330001203121112-3112220220220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-3332200122333312-1133321211323023-2211233323221321-1233112331032322-1233122222231023-1101001302320113-2333031203102311-0112030122133230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for validate period disable.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102000310022031-1003100031223303-3232321121033022-3201311212221233-1130111221220012-1110120210031122-3122221313103221-1030021331313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims.validate_period_enable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.reserved_claims](data-sources--http_loadbalancer--reference--group-019.md#canonical-0303133023120031-0231031101210213-1112203310123213-2311122021023032-2031322322111201-0112323012111131-2112103030122332-0312113221123201)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-2322223123323003-1002202222330303-1320033123203110-2120033033103332-0102300213203210-1112002112330030-3211311132211231-2120221212030321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for validate period enable.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.target

<a id="canonical-0000302000310333-2222312000231113-0130230103211100-1133030320323022-0112030303230010-2302110210223302-2230313130213212-3112101303211021"></a>

Type: `"single"`. Computed.

Define endpoints for which JWT token validation will be performed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

<a id="canonical-1020101211323131-1200333210303130-1203211001133120-0203321022222310-2113213013220320-3010230303230222-2202201001020130-2322330122313013"></a>

### Direct properties for `jwt_validation.target`

- [all_endpoint](data-sources--http_loadbalancer--reference--group-019.md#canonical-1101313102131101-1333202133213020-2311332211210103-3011122201000211-2010023320302030-3002102020231102-3012312132200012-1121000011031120): complete subsection reference.

- [api_groups](data-sources--http_loadbalancer--reference--group-019.md#canonical-1032120103002202-1110323013230013-1233012123030332-1223232312001302-2322021133310113-0101313331023123-0200230022131333-2223302021201030): complete subsection reference.

- [base_paths](data-sources--http_loadbalancer--reference--group-019.md#canonical-3032120320120311-3232303212221310-0121302102233123-1210203131301210-1322330222031020-0220223012120301-3100223111002002-0033301232101230): complete subsection reference.

<a id="canonical-1101313102131101-1333202133213020-2311332211210103-3011122201000211-2010023320302030-3002102020231102-3012312132200012-1121000011031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.all_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.target](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321)
- jwt_validation.target.all_endpoint

<a id="canonical-3101013330111012-0123223321330313-0222210103123300-0231202112133133-1210300130031120-0123210312300201-3323211220030110-3202010130302122"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032120103002202-1110323013230013-1233012123030332-1223232312001302-2322021133310113-0101313331023123-0200230022131333-2223302021201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.api_groups` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.target](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321)
- jwt_validation.target.api_groups

<a id="canonical-2121021221333122-3022322203310120-3023133020231130-0323330213103010-1311032231103010-3131203321030301-2320300131303331-3223011000232100"></a>

Type: `"single"`. Computed.

API Groups.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121200322012200-3212123031321203-1312301011011201-2121313020330211-3130211110301232-1321221302012011-2301113212313010-3032221120212312"></a>

### Direct properties for `jwt_validation.target.api_groups`

<a id="canonical-2012033222231020-3012310201003031-1012200002232310-0100121111312110-1313101012213200-0031301200223100-0102311231103312-1123121023102320"></a>

#### `jwt_validation.target.api_groups.api_groups` property

Type: `["list", "string"]`. Computed.

API Groups. Group or collection configuration

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3032120320120311-3232303212221310-0121302102233123-1210203131301210-1322330222031020-0220223012120301-3100223111002002-0033301232101230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.target.base_paths` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.target](data-sources--http_loadbalancer--reference--group-019.md#canonical-3320100111233100-2113302231232203-3303311321213110-0121302310123021-1231113212032011-3001021101020021-2023122021203001-2010102033133321)
- jwt_validation.target.base_paths

<a id="canonical-2320132010131203-3222010232113201-1121130313300123-2201030001110333-3320232021332300-0302111012313133-3110030013100230-3303101320120223"></a>

Type: `"single"`. Computed.

Base Paths.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013230132103000-2330333202321123-3331210310101000-3212120311023021-1002231332120313-0011321101203023-2032200300230122-0320330022121120"></a>

### Direct properties for `jwt_validation.target.base_paths`

<a id="canonical-2300012230003030-3020330310330222-1233003023120003-1033021021011030-0011103301232330-3111120213220222-0033320023300212-1101033331313331"></a>

#### `jwt_validation.target.base_paths.base_paths` property

Type: `["list", "string"]`. Computed.

Prefix Values. File system or URL path

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2101012003211203-1333223011213101-1311003002102303-2000301122201222-1111130223313030-0211031013100112-2032203303233133-3303103031322301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- jwt_validation.token_location

<a id="canonical-1313032200302112-0113112001220023-3101132310222310-3330211120221131-1111133202121130-2220110013012131-2200321232123123-1331011233100122"></a>

Type: `"single"`. Computed.

Configuration parameter for token location.

Additional upstream details:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

<a id="canonical-0321112200122332-2133300122232031-0031001102220032-0112322020013010-2000110213203132-3203102322303333-1232323102111120-0323023002223133"></a>

### Direct properties for `jwt_validation.token_location`

- [bearer_token](data-sources--http_loadbalancer--reference--group-019.md#canonical-1032111232021013-2223332323312312-0230231122101003-0210202333210322-0212233233233120-3120300221122332-3033230201120130-3220010301201222): complete subsection reference.

<a id="canonical-1032111232021013-2223332323312312-0230231122101003-0210202333210322-0212233233233120-3120300221122332-3033230201120130-3220010301201222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.token_location.bearer_token` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [jwt_validation](data-sources--http_loadbalancer--reference--group-019.md#canonical-0133002233023033-0213212130220323-1323123322310121-1331110232211203-2031102121220113-0003303230320233-3002110300000232-1203213002200130)
- [jwt_validation.token_location](data-sources--http_loadbalancer--reference--group-019.md#canonical-2101012003211203-1333223011213101-1311003002102303-2000301122201222-1111130223313030-0211031013100112-2032203303233133-3303103031322301)
- jwt_validation.token_location.bearer_token

<a id="canonical-2032233003301210-0002323112311332-1333300003321320-2021233211011022-0122110012132211-0211100012222213-2010300023122330-0302103221331130"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121023203221011-0201101021002021-0101212311232312-2121320103303333-0123032110032312-3313111220330000-2332202302121131-2130131301133210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- l7_ddos_action_block

<a id="canonical-0232103010002303-1023301101031232-1321322330211332-1311120022102110-3332103312332313-1003212031312331-1133322202000300-1112130233211030"></a>

Type: `["object", {}]`. Computed.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [l7_ddos_action_block](data-sources--http_loadbalancer--reference--group-019.md#canonical-0232103010002303-1023301101031232-1321322330211332-1311120022102110-3332103312332313-1003212031312331-1133322202000300-1112130233211030)
- [l7_ddos_action_default](data-sources--http_loadbalancer--reference--group-019.md#canonical-0230013202230202-2231233323102130-3311102130002131-3310333231302321-1223231322000033-1312301232101122-2020201123233332-3011101302103233)
- [l7_ddos_action_js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-0221100300320322-3123111133120131-3131320122210021-1323322022220102-3313222223213302-0331230113331000-3302212003030120-1201022300033111)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321021012132031-3001312230330301-0022003310222320-2211111231213322-2010212203031011-2021010223131312-3020200133123132-2023213310130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_default` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- l7_ddos_action_default

<a id="canonical-0230013202230202-2231233323102130-3311102130002131-3310333231302321-1223231322000033-1312301232101122-2020201123233332-3011101302103233"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301313020211313-2011033021131302-2123320013103200-0212212231232132-2313220232122330-1120020310001123-3110221330102102-2001223123013231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_action_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- l7_ddos_action_js_challenge

<a id="canonical-0221100300320322-3123111133120131-3131320122210021-1323322022220102-3313222223213302-0331230113331000-3302212003030120-1201022300033111"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2030201012131111-0330000330000113-3031330222131130-3321200211231000-0300323132321100-3111020302023222-2133132320223213-2130202121132023"></a>

### Direct properties for `l7_ddos_action_js_challenge`

<a id="canonical-0212200222310021-3110031020332111-1203200102312232-0000303322021333-3321332122102010-0213302101321222-0100301131230300-1130031003013211"></a>

#### `l7_ddos_action_js_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3120032003323303-2312122232220033-1010330210221101-2321222020210130-3030031100212323-0303300310103201-2002220030322103-3130121333201201"></a>

<a id="canonical-0202201312323022-3212031112213321-2100332231222302-1311213001323303-3002311121311033-3000313210213233-0323232330022003-3013310101200203"></a>

#### `l7_ddos_action_js_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0310023203031233-2231301000020122-1013311001323332-1013123121300322-3133321312300012-0332033200333122-1112111112310033-2122301021302023"></a>

<a id="canonical-2022321210113210-3131120202301113-3110312123021202-1030111303331022-0231320101210010-3122002232203131-3210000122333303-3333321212231221"></a>

#### `l7_ddos_action_js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- l7_ddos_protection

<a id="canonical-1113001110021101-1330202032023100-1033313231321330-0131002021033121-0300110001013303-3010203313023321-0200333110112221-3320333000013303"></a>

Type: `"single"`. Computed.

L7 DDoS protection is critical for safeguarding web applications, APIs, and services that are
exposed to the internet from sophisticated, volumetric, application-level threats. Configure
actions, thresholds and policies to apply during L7 DDoS attack. Defaults to \`map\[\]\`. Server
applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-clientside_action_choice": "[\"clientside_action_captcha_challenge\",\"clientside_action_js_challenge\",\"clientside_action_none\"]",
  "x-ves-oneof-field-ddos_policy_choice": "[\"ddos_policy_custom\",\"ddos_policy_none\"]",
  "x-ves-oneof-field-mitigation_action_choice": "[\"mitigation_block\",\"mitigation_captcha_challenge\",\"mitigation_js_challenge\"]",
  "x-ves-oneof-field-rps_threshold_choice": "[\"default_rps_threshold\",\"rps_threshold\"]"
}
```

<a id="canonical-3232023322211200-1333122020102101-2201200302120112-3223112031130212-2120023302322210-0011023330321111-2330013231100031-2023030230021201"></a>

### Direct properties for `l7_ddos_protection`

- [clientside_action_captcha_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-0221322221033021-3000332113030120-1123133132002223-3321032321210203-0222010332330130-2213011200330203-1100210001313222-0033021331132321): complete subsection reference.

- [clientside_action_js_challenge](data-sources--http_loadbalancer--reference--group-019.md#canonical-0030202102033233-0103213022003232-1102322212330102-1110201010233121-2200020011312332-0303332122030002-3122020113130100-0230221310010221): complete subsection reference.

- [clientside_action_none](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311001101232003-3311012301031333-0301201223203332-0321321301322300-3210332203113323-1310021303303111-0331300211210120-2323203112202020): complete subsection reference.

- [ddos_policy_custom](data-sources--http_loadbalancer--reference--group-020.md#canonical-0201100210222302-1320321202001120-2223231021210210-3231122230232110-3223220030301001-2211022303201010-2033203231003330-3232320103121031): complete subsection reference.

- [ddos_policy_none](data-sources--http_loadbalancer--reference--group-020.md#canonical-0311010100033111-3130301232000303-3123210031221223-2200133211303302-1023320231231332-2003020031323103-0310102103012133-1030013001300302): complete subsection reference.

- [default_rps_threshold](data-sources--http_loadbalancer--reference--group-020.md#canonical-2202000110133032-0220032020200212-1223231112002203-0201130110202112-2232221312311303-1312032332310122-3203310033223303-1103003102131201): complete subsection reference.

- [mitigation_block](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211023313203213-1030213203333230-2310010132231333-2013231321100132-3201202111303100-3220311000313013-0221210332333132-1002302313012102): complete subsection reference.

- [mitigation_captcha_challenge](data-sources--http_loadbalancer--reference--group-020.md#canonical-1031133210330210-0121031133213322-2003130133110221-3010213022121020-2230331320300312-2210230200312121-0333133221121311-3022102322112202): complete subsection reference.

- [mitigation_js_challenge](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131120333131103-3112111212112313-0222023003231031-3001211131031033-3313030101332323-0200222313320002-1101222122023332-3101023232110130): complete subsection reference.

<a id="canonical-2113102211020300-1231110132202103-0320113020223211-0130113133120230-0221103121233323-3033012213003201-3130023010321011-3321301030202320"></a>

<a id="canonical-2032220032020213-0322303222232003-3320221232123120-2332312220210331-2013221131300220-3333321220022033-0030000132103311-0323101323120333"></a>

#### `l7_ddos_protection.rps_threshold` property

Type: `"number"`. Computed.

Exclusive with \[default\_rps\_threshold\] Configure custom RPS threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 50000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "50000"
  }
}
```

<a id="canonical-0221322221033021-3000332113030120-1123133132002223-3321032321210203-0222010332330130-2213011200330203-1100210001313222-0033021331132321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.clientside_action_captcha_challenge

<a id="canonical-1223303301010313-2011221112220312-0023132330013031-1121132110112121-0230232000302310-1323011020021110-1021110033310012-0231022223010223"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122022323110023-3323001303132213-2000012232303211-2010101223222003-3301000310130210-2232320220333223-1222320012210221-2103221010112213"></a>

### Direct properties for `l7_ddos_protection.clientside_action_captcha_challenge`

<a id="canonical-2120210210032022-3323310331211302-0303303322210201-0110232310003010-3031300000230112-0011300000222133-2313101200003323-2220123302223211"></a>

#### `l7_ddos_protection.clientside_action_captcha_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2033202121220122-0203223302230222-0320112113032202-2312330210012101-3220111103001000-0203000323230312-1311033113010220-0032013113213323"></a>

<a id="canonical-3032223111102131-2132321303311020-1001132113222031-2122031120113230-2220122112330021-1303233321120012-2010023210001231-0321213323000232"></a>

#### `l7_ddos_protection.clientside_action_captcha_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0030202102033233-0103213022003232-1102322212330102-1110201010233121-2200020011312332-0303332122030002-3122020113130100-0230221310010221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.clientside_action_js_challenge

<a id="canonical-2000303021320022-3321312201203232-0103203103311101-0001110112200210-3033031221210123-2323110201121123-0312022011011211-2201203000320001"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203021030333331-2001020221132313-3120101022020233-0311233110331320-3111330030323021-0222301123200313-0100120130212333-2230133313330111"></a>

### Direct properties for `l7_ddos_protection.clientside_action_js_challenge`

<a id="canonical-2202121231013133-1011210033302003-2333313203320020-2121231202222230-2002330333330311-3121102222321200-1020333300032321-2310033020203101"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3002303323232020-1202103203013123-1101030322102210-0010321122032223-1331000111003103-1330003232323113-3212333232113332-1013330231112003"></a>

<a id="canonical-1321003330131132-2131023011330100-1122212301103102-1033031233221202-2033200232223022-0122210220201103-1130310112012312-3232323113311120"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2003322012322012-3130100310032302-2211003312133100-0300301212122320-1003002312133220-3011313010031222-2113320311022202-3012200030301102"></a>

<a id="canonical-0300122101111101-1010003221231300-0021322221122221-1312130012302333-0213210330002323-3131021121010113-3211101021330030-0123312222113102"></a>

#### `l7_ddos_protection.clientside_action_js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3311001101232003-3311012301031333-0301201223203332-0321321301322300-3210332203113323-1310021303303111-0331300211210120-2323203112202020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.clientside_action_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.clientside_action_none

<a id="canonical-0031300333321301-3320022013103221-3010221010113303-0303122232013030-2212321322110030-3202223322301201-0132330112100221-2213313211130232"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.
