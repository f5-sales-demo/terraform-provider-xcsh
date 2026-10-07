---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3310122132030102-1203120120323111-2132201013301221-0100331310020010-0001300021101210-2203001130113222-2032311310131010-0110121221112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params

<a id="canonical-1110303313130302-1020111002321211-2030122310300303-1201310100201113-3103022230122111-2303323330113222-0231303032132113-0310012103132130"></a>

Type: `"list"`. Computed.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2221200300223213-2022102222110320-3011120100022220-3322221321021332-1112032132121031-2012331202200100-0001323110013310-0312300303001232"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.query_params`

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-2220011112330022-0310301332013030-2100300201032233-2033211023320131-3031223120003133-0000203211313132-1103330320321132-2010000032123111): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3120203323203001-0322331323021110-3301203220221012-2330320013032203-1030100123022021-1013101322321330-2112331301113311-0123010232113112): complete subsection reference.

<a id="canonical-3330303032303013-0001031200231101-2110101200023132-0113133301311302-3001231032322112-2013032230230013-1012230121212031-1032032033321333"></a>

<a id="canonical-1203101310101020-2011212311030311-2330223132333231-1320010102313300-2100333032303112-1300313213211010-2211321231021323-2300201332111012"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001030100231200-2210201303321103-0213302111002000-3332130321021122-0230201001213000-1113313311233103-1302223120310232-2103003120301332): complete subsection reference.

<a id="canonical-1033331022200120-3012113011123111-0033112313213330-0122331232000300-1312330032302000-1302310111030102-1222302110132212-2122300212200033"></a>

<a id="canonical-2221010121332130-0233213031113310-0201231033330102-1023321001312013-2202113313210031-3030110311323331-2023132130322302-2323003001102023"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.key` property

Type: `"string"`. Computed.

A case-sensitive HTTP query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2220011112330022-0310301332013030-2100300201032233-2033211023320131-3031223120003133-0000203211313132-1103330320321132-2010000032123111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310122132030102-1203120120323111-2132201013301221-0100331310020010-0001300021101210-2203001130113222-2032311310131010-0110121221112121)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-2231023303133320-1222012322201312-0312020130113230-1310100210020120-1013220333100003-0103032001132021-1301110102232212-1301112002013022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-3120203323203001-0322331323021110-3301203220221012-2330320013032203-1030100123022021-1013101322321330-2112331301113311-0123010232113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310122132030102-1203120120323111-2132201013301221-0100331310020010-0001300021101210-2203001130113222-2032311310131010-0110121221112121)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3023022020030203-0011303223133101-0202103212111312-3020200002013213-0122300003021130-2231330311220210-1010300323023103-0303103201013121"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1001030100231200-2210201303321103-0213302111002000-3332130321021122-0230201001213000-1113313311233103-1302223120310232-2103003120301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-3021123223221003-3113310000200133-3031311313303133-1013210032233121-3012013303102311-3323202213200301-0223130322200003-3330321310301100)
- [api_protection_rules.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-3133201101101310-2031330001023301-1033110232303330-1110030121203131-1023221110103232-0202031312003110-1300230213333230-3322121113333200)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310122132030102-1203120120323111-2132201013301221-0100331310020010-0001300021101210-2203001130113222-2032311310131010-0110121221112121)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-1211203130111233-0210132231113000-1113110331101222-3232201002330000-2212111021021003-3013223212222013-2221112230000032-2321130003312300"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0130033201233012-3231021121013132-1201031203323320-0223110031003210-3133111131130201-2322303022033023-3130102302032230-0323122002102310"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item`

<a id="canonical-0111111031101120-0330133123201211-3003021100102231-1103331333101302-3111302111310121-2303333230212312-1210111013013002-3313331203222320"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2032110201030101-1222303123221120-3333122213003322-3032331232232110-1332330203030131-2222031202003330-1233120122302112-3031231200211300"></a>

<a id="canonical-2212023223103303-1203211023202323-1320301010221012-0010323330131021-2303013230121210-3123203002233323-3133110213133112-2032102230130032"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3210331201303102-2333001132313110-2030202322013031-0321321003232000-3111212311010312-2001101210003202-3232013332302323-3111003032131311"></a>

<a id="canonical-3332102330113111-2213333321023013-1222223312000201-3113313132313113-2001011033323122-1302112213300201-3113321022110231-3003121211332330"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- api_protection_rules.api_groups_rules

<a id="canonical-2011211112310032-0230210010011310-1211131232323112-3313201213013223-0031121020132323-2331222211011032-2333031111023210-3230312100013210"></a>

Type: `"list"`. Computed.

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-2012011121232203-1301320310312230-0211133010312230-2301032012223022-3331320102222333-0100333222012311-3301321202221201-0113111123201002"></a>

### Direct properties for `api_protection_rules.api_groups_rules`

- [action](data-sources--http_loadbalancer--reference--group-006.md#canonical-3020123231322122-3002200301232223-1112021032231130-2323213133100010-2223010022222211-1220022221011102-3312102130121023-2021031100210333): complete subsection reference.

- [any_domain](data-sources--http_loadbalancer--reference--group-006.md#canonical-1212101130321320-0030111032110202-3020132121000332-3020102001233210-3313020011322233-3003130031032312-3330220223200322-1321001201111302): complete subsection reference.

<a id="canonical-1122203323002322-3210312302002203-3310201020220030-0020120113331322-3130033330231312-1023311220031200-0213032320032301-1231320113211011"></a>

<a id="canonical-0303301122001130-1201032303001023-2021122010101211-1221230121320123-1301212322133021-1012010022012333-0122112213033203-2230122002000111"></a>

#### `api_protection_rules.api_groups_rules.api_group` property

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Additional upstream details:

Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3131033331210230-2100203012020033-0033032030310200-1302030020103110-2120313212203121-2321203112031232-0123210000001230-0111130301112220"></a>

<a id="canonical-1302002031133301-3331100033120120-0012113212330222-3210133310333303-3113231332002111-3331303333102113-3332321120223011-1223233103003213"></a>

#### `api_protection_rules.api_groups_rules.base_path` property

Type: `"string"`. Computed.

Base Path. Prefix of the request path. For example: /v1.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-006.md#canonical-2033211201003302-2031102213013202-0223233020221023-3301020223213110-3200131131232010-0332311132201303-1010000020311312-3321200020330333): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202): complete subsection reference.

<a id="canonical-2201211310100003-2011321222010223-2203301201330222-1211123123111120-1320130301330022-1111330012212330-3132031232200002-3121102033300312"></a>

<a id="canonical-2030120000121330-0110303030213201-3130302301111132-3313110231013131-3130123330110310-1301233132312021-1201311100103123-0122221230102332"></a>

#### `api_protection_rules.api_groups_rules.specific_domain` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3020123231322122-3002200301232223-1112021032231130-2323213133100010-2223010022222211-1220022221011102-3312102130121023-2021031100210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- api_protection_rules.api_groups_rules.action

<a id="canonical-1001121120002112-2123310012113023-1303210321301112-0032130330133010-3300131330133320-3020320130102032-1102232132200120-2010231210003322"></a>

Type: `"single"`. Computed.

The action to take if the input request matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

<a id="canonical-3200212000101200-1010322331210200-0001022202333031-0222031021223221-1332011210301011-2232012320031202-2001002321021320-1313203022331000"></a>

### Direct properties for `api_protection_rules.api_groups_rules.action`

- [allow](data-sources--http_loadbalancer--reference--group-006.md#canonical-3033033303100311-1001000330001220-2122131210122220-3113330301000230-0321330231320300-1101001233213002-2012331233211003-0231121130332200): complete subsection reference.

- [deny](data-sources--http_loadbalancer--reference--group-006.md#canonical-2023112031201323-2220022112310302-3120010302302030-1031310332230311-3320222001230302-0012101112320300-0331223021032000-3211200103233302): complete subsection reference.

<a id="canonical-3033033303100311-1001000330001220-2122131210122220-3113330301000230-0321330231320300-1101001233213002-2012331233211003-0231121130332200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action.allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-006.md#canonical-3020123231322122-3002200301232223-1112021032231130-2323213133100010-2223010022222211-1220022221011102-3312102130121023-2021031100210333)
- api_protection_rules.api_groups_rules.action.allow

<a id="canonical-2321312222203221-2333121333301211-2121132310220031-3331310001103132-1030201232103232-0131100213202123-3003232302210212-2113213331021110"></a>

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

<a id="canonical-2023112031201323-2220022112310302-3120010302302030-1031310332230311-3320222001230302-0012101112320300-0331223021032000-3211200103233302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action.deny` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.action](data-sources--http_loadbalancer--reference--group-006.md#canonical-3020123231322122-3002200301232223-1112021032231130-2323213133100010-2223010022222211-1220022221011102-3312102130121023-2021031100210333)
- api_protection_rules.api_groups_rules.action.deny

<a id="canonical-1002121302030323-3332232011213211-3010333021233000-3210020002013231-3231010232011110-0211023210010103-1302001023223103-3023112122010303"></a>

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

<a id="canonical-1212101130321320-0030111032110202-3020132121000332-3020102001233210-3313020011322233-3003130031032312-3330220223200322-1321001201111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- api_protection_rules.api_groups_rules.any_domain

<a id="canonical-2023131131313031-3012303333313322-1023232321202321-0203330332023121-1132000110001111-0023002213133303-3131231233203033-2102113212300222"></a>

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

<a id="canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- api_protection_rules.api_groups_rules.client_matcher

<a id="canonical-0332111021332230-1120132100000100-0010122020132201-0101100011310323-0231331213211221-1322213212332010-2021313021321021-0002201300100020"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-2221321230211020-3131020231003201-0113213121010002-0132103102033323-3101300201000013-1133330002231320-2313202112331001-0302012031333132"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher`

- [any_client](data-sources--http_loadbalancer--reference--group-006.md#canonical-2301022122030020-1233123030123230-0022033220321212-1303131321201111-0321120113331002-1113132021032300-0233000332101331-3012101122122332): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-006.md#canonical-1231033101130322-2112311211102100-3003220102103013-2212011023010022-1320031103100313-2012302102321202-0002231202103102-0222230201221031): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-0030020201332022-0032100032331102-0303001320322102-3300330301130013-1323133333032211-3132113023313101-3021321200311110-1121030222010011): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0320330320001233-1130023200330130-2222221020130233-2132200110232123-0303221013103333-0212330230312031-3001002113212212-3203000023112211): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-006.md#canonical-1011201121200230-1032001333021220-3313331333331012-3320221030032232-1202120220023310-3202311012120201-3230012031021121-0122212020011013): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3230103330333033-3322012100123200-1033002123321202-3021031102133112-0101210323133312-3330202200011213-1333010330031101-0333000310020020): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-1311231003231230-2003011223012030-3022200023230301-2310221200313030-3012010220111323-3331001230220202-0101102131013210-2303222131312301): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-1212102232200221-3021110312231213-0232201101012303-1131321000000213-3222202323302201-1002022322330033-1330111232102323-2111132031113110): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-2122312223130021-1231130133111031-1031221202103202-3212331223210133-3211022032110231-2331103232021232-1023003103221011-1030312210211320): complete subsection reference.

<a id="canonical-2301022122030020-1233123030123230-0022033220321212-1303131321201111-0321120113331002-1113132021032300-0233000332101331-3012101122122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.any_client

<a id="canonical-2333002221233112-0033133120320021-2030011003333301-3211023210031122-3002220113332011-3301013121130232-2321001301030102-2032021032321021"></a>

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

<a id="canonical-1231033101130322-2112311211102100-3003220102103013-2212011023010022-1320031103100313-2012302102321202-0002231202103102-0222230201221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.any_ip

<a id="canonical-0011200300123020-0202111031131131-2001133211300123-0022003331232203-1023122011211102-2231232002002223-3023323130232323-0323003022030101"></a>

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

<a id="canonical-0030020201332022-0032100032331102-0303001320322102-3300330301130013-1323133333032211-3132113023313101-3021321200311110-1121030222010011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.asn_list

<a id="canonical-1300021023233130-1030131020312122-2201330333110302-2311003220123010-2021222313020212-1202010122130311-3313020202213102-2330112113231100"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-2113033021003211-3113020012221022-0332101102321103-2332203000212210-3303333201303221-1322303321221203-3213033323230331-1113001133230220"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_list`

<a id="canonical-3000030213002301-0203020330101022-1003202213231112-1330223331232030-0211313223132323-2100131302200230-3031222003103112-2012122131123312"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0320330320001233-1130023200330130-2222221020130233-2132200110232123-0303221013103333-0212330230312031-3001002113212212-3203000023112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="canonical-0031320110321033-2330133101203202-0201120123300030-3133332012103230-1313100020202311-1113203130010112-2011123210031201-3322123022000330"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-2231020322023001-2101033303002111-1131213102021231-3023330010120001-0233031320002332-3022020103130022-3313312130131331-0013012013021212"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-3100230301020221-1112100111323303-2021322232321211-2303312003323212-3211013013311212-3211110012122230-3201123300312332-3021333132110002): complete subsection reference.

<a id="canonical-3100230301020221-1112100111323303-2021322232321211-2303312003323212-3211013013311212-3211110012122230-3201123300312332-3021333132110002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0320330320001233-1130023200330130-2222221020130233-2132200110232123-0303221013103333-0212330230312031-3001002113212212-3203000023112211)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1312022100010133-2130021002232111-0123113303120312-0000033202312122-2221131310020030-0301023233311123-2331313332323213-0202300201002230"></a>

Type: `"list"`. Computed.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3130131223013021-0023320211213332-3333211212101033-2112022300023000-0203331301212300-0120301303200020-1301212312021310-1213020311333031"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-2302330220322021-2100033110002110-1133322121010220-0203300223200333-2331020230201121-1231100330120212-2313110001231021-3312120311102200"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3332120113030233-2101231223103320-3322022310210220-0011221201322123-2002302023110112-3110333100303333-3210210220032210-2330213210300133"></a>

<a id="canonical-3331321202311201-2100013033032233-2120213030321033-2000231120110332-2211003232022121-1021303310321300-1210131231321121-3000001000013323"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3100213001312323-0120323023230302-1333032022131332-2121333222130012-1310103233033102-1122023000131000-2100032001133111-2020333122302131"></a>

<a id="canonical-0333320113110232-2130233130121223-3331022122230212-1330001323310300-3310301010230023-2012202131230003-3111220210220203-2320212223332132"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
  }
}
```

<a id="canonical-0230032223221032-1012222001020031-2100233102021032-0323313000123032-0333021200301303-0332331300311320-1102020013033122-0021302213122220"></a>

<a id="canonical-3220033012023230-0232301023002323-1020033302112012-1302203003313333-3103013301000212-1000320133213223-0233303011011110-3320001132020102"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3202233113010220-0232022311021001-2321013032213310-1201033231330312-1322203301331130-1111210011211132-3101100320022002-2012100123212010"></a>

<a id="canonical-3231300202110212-3321202212002031-2130332002202023-0300321021321202-2222121222301022-2322130311110223-1010323230110222-1033033212301322"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1011201121200230-1032001333021220-3313331333331012-3320221030032232-1202120220023310-3202311012120201-3230012031021121-0122212020011013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.client_selector

<a id="canonical-0310321311102020-2210233101030022-1022322123000111-3013011202022311-3300323013113111-2021130233012110-3102213312330231-3213103221301012"></a>

Type: `"single"`. Computed.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-1031302012211333-2212202212003332-2003120012022221-0220020003010322-0000103232132131-1330011021132310-2121012223122213-2020313223000322"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.client_selector`

<a id="canonical-1323223123003110-0101031110122231-1301302233133332-2220103330323020-1031230213111032-0132021001212033-3022021022033302-1331231121012021"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.client_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3230103330333033-3322012100123200-1033002123321202-3021031102133112-0101210323133312-3330202200011213-1333010330031101-0333000310020020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="canonical-1313300012232022-0130232311113023-1222102233123101-3031210330212201-0111102333012013-3033032232302120-3023303132203012-3322323001303231"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-1110332022302012-3103313203120302-0313212320032213-0210102312121202-3333132233322133-2202332103121123-3130030330031200-2231103121203013"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_matcher`

<a id="canonical-1121030021110022-2011003012310220-1112023213333302-1210213131321301-2132331202002221-3222332221320211-0023100330031210-3020221311112132"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.invert_matcher` property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-1220232101102032-2213122212313113-1331000310311033-1132132303111121-3220033132121310-2301202010221300-3030200320033232-2113320223003301): complete subsection reference.

<a id="canonical-1220232101102032-2213122212313113-1331000310311033-1132132303111121-3220033132121310-2301202010221300-3030200320033232-2113320223003301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3230103330333033-3322012100123200-1033002123321202-3021031102133112-0101210323133312-3330202200011213-1333010330031101-0333000310020020)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1100103113033202-1321010012300333-3220323301310020-1322311230221000-1202212011130002-2010312120112322-3232122320013130-2223131021313312"></a>

Type: `"list"`. Computed.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-3301303021012213-2300332112022311-1010123320003221-2131323210333301-2033301130002100-1022121130213001-0222133232002202-2020113212013300"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-2100331021012223-3120102032221113-3212302332301322-2123102203223322-2030312211021333-2022312301230100-2010212120132130-2001302111230102"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030210221332111-3033332111222102-2333311000330212-1300322102011003-1031333322212132-0031100011133000-3101310012012213-3022321032300213"></a>

<a id="canonical-3212010112303303-1030301130020321-3200223211000313-0200331001013000-2000333223032023-2300001321310232-3330221011020222-1001122320232101"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3223131000103023-1103131322301331-1131130103110301-1020000123312030-3300030121033000-3112213233222301-0011111200100023-2203000020032210"></a>

<a id="canonical-1013123300212023-2120301100200121-2110213020201120-3021030111231030-2001211132331211-3310323022210021-1020221023123113-0033120111113200"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
  }
}
```

<a id="canonical-2301223110330210-1020232223100120-0332111131202321-0310332212112231-1100011310321212-2133212032222230-3333021133101201-0213320120203133"></a>

<a id="canonical-1001122320133231-2301021211312012-2031322231313020-2111232030013233-2003301003100303-3000022103110331-0312233133220031-2030011103320121"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1022332210210011-1123213321111332-3120201210102233-3112303221000033-3001201031133033-2012333313322303-0033003331110120-0201133133102332"></a>

<a id="canonical-1021221303131223-1130313321033301-3031221220032203-2201330000021121-0112331113330301-1233302230010332-0232013031020311-0320122031300011"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1311231003231230-2003011223012030-3022200023230301-2310221200313030-3012010220111323-3331001230220202-0101102131013210-2303222131312301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list

<a id="canonical-1331011332011201-2003132213132300-3200221022001030-1122211211332332-0123312131303031-1023202102002020-2332033002312100-0212103213120013"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-3331302310220012-1000102030223111-0101100010002220-1131023102201030-1211032221000202-2223230020300210-1310301301001002-0321303032331323"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list`

<a id="canonical-1203111031221010-1013011032333320-1310302033330022-1132211300303310-1011013303131313-2013310322033123-1302020011321103-2112330100201113"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-0210203312020312-3213203001010200-3302100330030221-3221101132033120-2221200100303313-1012200121333120-1023232201311110-3130330331001220"></a>

<a id="canonical-2101203211010320-2001033231011310-2220233130120222-2133013122331303-0200203313302202-0331000030121300-2210003133011221-2033332233332202"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212102232200221-3021110312231213-0232201101012303-1131321000000213-3222202323302201-1002022322330033-1330111232102323-2111132031113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list

<a id="canonical-0212022121310103-0310032220221231-1223021233302123-1223110100003210-1332133233203103-3310222330302010-0310120232321032-0012310103111103"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-1222031310230330-3231112120333300-3102002132321130-3113232121203011-0201013103232233-0103313111211230-1012230200012110-3011203111330233"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list`

<a id="canonical-2120310212110131-3201212233001101-2332301322302030-3222202310333000-3021000003330330-1100333123211212-0031130121133220-1300233202102020"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2122312223130021-1231130133111031-1031221202103202-3212331223210133-3211022032110231-2331103232021232-1023003103221011-1030312210211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0302222103133021-0331310203313102-1221302121121113-0101103223313130-2201232311021012-0020130210232332-2331101232032212-2133223112321122)
- api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1023122110122330-1033322122202323-0110121122310131-3312103220101020-0301211020223322-1332123301210003-1121030032211200-2032022113113212"></a>

Type: `"single"`. Computed.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-3001020310020332-3200102122331222-3220213101030211-0320302131123021-1300300323020011-0130132223130101-3322220023033012-3230311232003231"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0020233002232302-2211030100011331-3303133023001002-3120110200132232-2103221021112313-2113331220200212-0202103132010223-1030103230323322"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0031220223103233-1303113332310021-2213132233112312-3331033033012101-1303202132323220-1333111123203032-3002222031220333-2231020301202303"></a>

<a id="canonical-2332223210110200-0100012111002202-0020233332123010-2322222320100212-3330231112300110-2303112320230331-1103320133111133-3331033000031202"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312021232322112-0332133103301212-1032010231321322-3003120301222023-0001331212221012-2123220033020311-2203112302322200-2122302000211003"></a>

<a id="canonical-3313030300131021-3001323200312130-3001320311131222-1000331230021320-1113210300213200-2012232303031223-3330220313010222-2120322122220313"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033211201003302-2031102213013202-0223233020221023-3301020223213110-3200131131232010-0332311132201303-1010000020311312-3321200020330333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- api_protection_rules.api_groups_rules.metadata

<a id="canonical-1113233332220020-3201313210033231-1010022100332100-0123330330201232-2302102200210220-0131010331132233-3201202131121231-2232120303302302"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0133200210031113-2300033133203101-0320103111301021-1222000013000200-3033031232012202-3211323111300313-0102031110003102-0202132213133112"></a>

### Direct properties for `api_protection_rules.api_groups_rules.metadata`

<a id="canonical-1001220212101133-1200132331223321-2322033132112121-3322112323132323-0101233102232101-0231312133123120-3010213211001200-3123321300201223"></a>

#### `api_protection_rules.api_groups_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3320333131310321-1001103323313003-3013220132011321-0233013332033211-3001013332203022-1012230122320121-2322010133021013-3233121030033310"></a>

<a id="canonical-2302200132133021-3012010010022022-3102223333033101-0033030320120000-1322200100303022-0223100320320000-3313330011033032-1200002213110313"></a>

#### `api_protection_rules.api_groups_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- api_protection_rules.api_groups_rules.request_matcher

<a id="canonical-3321112021300032-0012302001201312-3231031100101211-3220012103303101-2232321211322310-0231100012023032-3203010302221021-0013010110022011"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Additional upstream details:

Request conditions for matching a rule.

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

<a id="canonical-0312220012221110-2230022320123332-1230320102031332-3002311021003213-1301203020002300-1303121322300202-1103203321201112-3332011303130111"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher`

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110): complete subsection reference.

<a id="canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-0130012030310310-0013022120113222-1303223113322201-1022120323212111-1032203033002311-3020021010010231-2232133112331223-1013320001220223"></a>

Type: `"list"`. Computed.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3232230322110303-0333032221130010-3320320312132011-0010130232033023-0232323010113130-3332002321332300-3110123001220012-2221222123320313"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1102313302033322-0310212323022232-2300220211002031-0332310200331002-0313333222322113-0200200220131232-1122330022112210-0103001111133123): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-0032220033223233-0133311310310112-3012213011302120-0110111311323220-2312012120321131-2130202121100320-1312133003111313-0312033302023023): complete subsection reference.

<a id="canonical-3313031232103303-2330320013001212-3300003102301322-0302330102330101-0033022103232222-0233302201001311-2212210331123300-2313321012230223"></a>

<a id="canonical-2120322320203030-2300021223102301-1232210232123322-1111130000131133-2302300011331001-1211213012013323-3300030303210211-3001130100122032"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-2300110132202210-0330210020033233-1012331012012003-2330313121330222-1321220203330300-2322100121000201-0031302132100121-2210333223302001): complete subsection reference.

<a id="canonical-2032211323312000-0323300230302020-0233201032220302-3002210103213101-3121120120021312-1203202030122032-2130333220333100-2320230120333302"></a>

<a id="canonical-1311303232230020-0133021132023103-2121210110210111-1230003233211130-0202221232212201-3322330000210221-3230311121003210-1333103103211220"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1102313302033322-0310212323022232-2300220211002031-0332310200331002-0313333222322113-0200200220131232-1122330022112210-0103001111133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2033210222013021-3232302212131000-0030033230202213-2223030111300133-3030311301111223-0221020321132020-3332232223212023-2101232100223010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0032220033223233-0133311310310112-3012213011302120-0110111311323220-2312012120321131-2130202121100320-1312133003111313-0312033302023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0331202322011130-0212020210302221-0212032321301232-0010013013322330-0122112322030101-1221132121021030-1020123302301310-0332012100123031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2300110132202210-0330210020033233-1012331012012003-2330313121330222-1321220203330300-2322100121000201-0031302132100121-2210333223302001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-2130022102013322-3211010102003101-1033102101132333-0322132332122320-3002103022131233-2230000322022213-2210103323301300-0023230130033211"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-1111320010231012-0300133121331200-1010313201103203-3320310200133333-2120311002131123-0123101122103302-3333233330221220-1323131203003223"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0003002313200212-1132302120001302-2113021110331203-2122321332330101-0322231032000031-0312101203131102-0223310233001223-1302233003011023"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103302231220101-3223230203011023-2330110102230023-2303033011131301-1021011020231333-1012302020021310-1123202231130221-1022222311233033"></a>

<a id="canonical-3002311301000103-0232232310322000-2002121203130303-0122120131013132-3221213033203101-3322220130000003-2000011302032322-1011202222103201"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0000001233113300-0120013120320001-1023103301223200-0023210322123013-3010003001312323-1102311233012131-1333012201011020-2221233021211332"></a>

<a id="canonical-3332121003032323-1011302232203133-2121012210121033-3001123000321302-1303322332003110-0230321123130111-2132030313323010-1131102202313021"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-2112111133300102-1213212221021223-2330101001232320-3322103110200320-3120122312030021-3302313221103113-3022122130310022-3000122320023331"></a>

Type: `"list"`. Computed.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1310011103310210-3321100323200033-2031000331222023-1022131130323333-0222223033132222-0020213130133132-0210121310202131-0023312300233110"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010312230120123-0133020133011130-2210132232100331-0112112210333312-2123223121003200-0231322212201131-3302203310021120-0233203303211120): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1201223320223213-1020212300312111-1201023132213033-3010002321323301-2131301223313012-3003033122000032-0103311123002303-1131303202113112): complete subsection reference.

<a id="canonical-2132110233331201-3022012130012300-0203102113310030-2220202300321213-2320021311213100-2011210021202110-0130202101111020-3333331221001213"></a>

<a id="canonical-2113311222322023-1332000102230333-3102100320013312-1303222000120000-0020211210121032-3110211211122012-2011011301033233-3020120031313300"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1331221100033320-3012300222231031-2133321111231220-2230310330213012-3232302220102303-1301123200012200-2110103201321232-0221130323121310): complete subsection reference.

<a id="canonical-0232233020221102-3001322001310133-3030132233110222-3130312300032120-1033113321223323-1112211300123331-2301111311010132-3032133020220311"></a>

<a id="canonical-2110232222020300-3001011212102210-1010123101010111-0330330312221301-3123330333133233-3023001220033321-3333223331103231-1123111111030103"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1010312230120123-0133020133011130-2210132232100331-0112112210333312-2123223121003200-0231322212201131-3302203310021120-0233203303211120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-3132211233331323-3310210020330011-3202333323301332-3323223231312021-0132011212323322-0300202122321313-1200301110023001-1300030012321201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1201223320223213-1020212300312111-1201023132213033-3010002321323301-2131301223313012-3003033122000032-0103311123002303-1131303202113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-0003232111030210-3331303221210320-3300011203101310-2310221031101023-3303023202301000-1333330302111230-0210001103031101-1133220033203323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1331221100033320-3012300222231031-2133321111231220-2230310330213012-3232302220102303-1301123200012200-2110103201321232-0221130323121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-2133103311130013-2330110020330231-3220031323102121-1232100102030330-1113023111321232-2231002123221131-0300200103120033-1001131132320333"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-2110002210201333-1233321102132331-0312323330301313-3112120301233330-1332021100031322-1012133022132331-1332123321000103-1231200232233311"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.headers.item`

<a id="canonical-0221032223003120-1032132333010121-3203232330020122-3300313033201022-3231320132123130-2323031201000200-2011321210121120-3113321000112021"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3123300103030021-0122223322133300-0101331032333201-0230133000301212-2230130111233003-0133210000323002-3031223333313101-2321023302231122"></a>

<a id="canonical-0320211030231211-2021222033023000-3022311011030000-2011301123330033-3132302121120011-1220121201122133-2120230030021123-3200231311300020"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2021000330333220-2303221232133131-3132333102231110-0030203233220010-0321323112321010-0322121031031111-1001132231010332-0203001301310202"></a>

<a id="canonical-3013021013032023-0133331022121022-2230211021102223-1230012132112213-1303013120102133-1112121011201131-2300202321131201-1203311220002102"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-3033131130133133-2201212300312123-1220122212012131-0100313021022131-3003230200113303-3231103211202321-1032231101333132-2330130111121212"></a>

Type: `"list"`. Computed.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1020212330311223-0211332231303033-3211113323003303-3130221013130231-2313023030102000-2312313200011002-2220332301312103-1022233012101320"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-0330123223201101-1220231000021332-0011300131231300-3032333121231212-2303100103130302-0231013123231313-2221222211021022-2031021233033223): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-3130201230012223-2201333132032111-1321230022132312-0131331211021312-1230222002020122-2131131202001123-3200020103031303-1012231132312120): complete subsection reference.

<a id="canonical-2030311123320131-0103132100031020-1120112322113010-0021131100023232-3123302033331202-0320122233130202-2031213003302312-2213330101120221"></a>
