---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3023212030210311-2221020310013000-2330012200320220-1203213223221222-0321202010302311-2023202123211131-0332331013110221-3310113202122331"></a>

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-2330200333121301-2101131203123312-1001003110000223-3010011212110132-2033212313120002-3330303020331012-0013331100021321-0223132332112011): complete subsection reference.

<a id="canonical-0103110310302212-3113303211131322-0020031333132231-2132011112001012-0202100101113030-3111131133211333-0302122130221010-3023021220330102"></a>

<a id="canonical-3130211023330233-0323330000122103-1233223312231201-1303212023330331-1222002212013300-3122232232233311-0301300232300012-0323111320122300"></a>

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0330123223201101-1220231000021332-0011300131231300-3032333121231212-2303100103130302-0231013123231313-2221222211021022-2031021233033223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2200302230320021-3230230211022322-2022031103130331-3232110200022331-2001102331011031-1022300030320212-1113312021301032-1212130222303031"></a>

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

<a id="canonical-3130201230012223-2201333132032111-1321230022132312-0131331211021312-1230222002020122-2131131202001123-3200020103031303-1012231132312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-0313321131121022-1110230110300030-2301022200133002-1303323003110222-0303011103222030-0101001013033020-2221003011213322-1011233032312320"></a>

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

<a id="canonical-2330200333121301-2101131203123312-1001003110000223-3010011212110132-2033212313120002-3330303020331012-0013331100021321-0223132332112011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-0033212322330101-3033303323100103-0032012111321220-3231320130001130-1012221200000212-3120012220202011-3233102220100130-1223013303330132"></a>

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

<a id="canonical-3113101122203333-2232322113110322-0012232121232002-3202011231120300-2221222202130212-0100013210223012-3333131002221321-3302312003320121"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item`

<a id="canonical-2100013212003103-3213032323003111-2102203012101122-3110302332013000-2010301323002212-1133020330313310-1232133030113100-3021311312111121"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3020033213123232-0300300303201100-2312313303123211-3022003212331311-3031230222213300-2103111303333302-3033123110313302-2111213120122000"></a>

<a id="canonical-2122323113133101-3113221002003331-1110123102010210-3131022032021011-0213203103030211-1213321211111132-3201002023201113-1221133100331233"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3102033122312131-0202311121010032-1100130022101110-0130211023013101-2002311310233232-3311232000122130-0000012301322200-1102202221100230"></a>

<a id="canonical-3010113021313331-3300102110310011-2033031211002113-0321313102102330-3031131030331023-1021110230031212-0031220022111121-3003202132223221"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-1313130202021101-2133201310021310-2010010002022223-0330301200332123-3000321031100120-1230020200100323-1133100031333330-2113221012020311"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2023323033032320-1300021213221133-3222322121222123-0221223220200232-2132232220122322-2120223333330010-0003132231002102-1302113130003003"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.query_params`

- [check_not_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-3203000210000203-1202122322331301-0010022211002113-1031011113300022-3031331210003013-3000100230112131-2030321320112111-0321001123011033): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-1213323101301133-3313330210323102-3301311101333310-2112311123330030-2320211300002033-0202231133110011-3031003301010012-2031110133313322): complete subsection reference.

<a id="canonical-3211331022102222-3202001312230311-2100002100233120-2101122220031102-1123233213310011-3322131110231021-1110313221011122-2023101132220310"></a>

<a id="canonical-2101221131311000-1230122023023201-0011002202121030-0310302221011002-1112133331332301-3121221112031012-2013211300021030-1323223001303001"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-3230012303111032-1021330130001222-1213001323133232-0310322122020113-0331022003111230-1002032020311330-1201322022301211-1132230213123201): complete subsection reference.

<a id="canonical-1231231332101101-1222103120313211-1322013031102303-2221322231113003-2220232032233312-2202113111311012-1210102322201322-1110100211222030"></a>

<a id="canonical-2201032102330323-2322231213013111-2322210331301322-0303200000012133-0223112113122011-1102232330310031-1201301023333033-2200203203033302"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.key` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3203000210000203-1202122322331301-0010022211002113-1031011113300022-3031331210003013-3000100230112131-2030321320112111-0321001123011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-1121123032321030-3301023122002133-1200213022121203-2133022113313212-1010231023102113-1002030023123311-1332110020103112-2320211133110111"></a>

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

<a id="canonical-1213323101301133-3313330210323102-3301311101333310-2112311123330030-2320211300002033-0202231133110011-3031003301010012-2031110133313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-0213310013321223-2130032333310330-0030102011330310-2030213222101210-0010221233213122-1011213100013021-1102023111003021-3230231303232002"></a>

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

<a id="canonical-3230012303111032-1021330130001222-1213001323133232-0310322122020113-0331022003111230-1002032020311330-1201322022301211-1132230213123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-2031332111223222-3020213003222132-0132331100100110-0201133131032211-3100311133033320-3332201023212023-0232021322130002-0313113103033300"></a>

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

<a id="canonical-3212322320300100-3113110231312021-0020032000112321-1322132331331221-3120012013321001-3002003031102111-1113230200013331-1032011130232020"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.query_params.item`

<a id="canonical-0221010101010322-3230101203300313-2013211200103211-3233223321001012-2200103112102111-3132021200333322-1233303212003232-2202300220213123"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.exact_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1011103332320022-0303221032131031-3210230120001101-1211121011033300-3031311330213100-3010121122123303-1030020000331121-3302002202022030"></a>

<a id="canonical-3130202012200032-3322203100100202-2321111202230030-0232323100012232-1331012310123311-1321002313213021-1002011201131201-3131123213322233"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.regex_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2321233012212221-3203211312331023-2022132230120201-0203113320101100-2212013231131012-1302210003220121-1213132001331311-1331031301111331"></a>

<a id="canonical-2311111023202232-1031000323212133-2302333121030021-2112322310213121-0213200122213000-0021320331013223-2110123030133300-3123202013221331"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.transformers` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- api_rate_limit

<a id="canonical-1002023121010233-2202310032231001-1033211001111301-0230103012300001-1333032021233330-3310221010100120-1020010332012323-3212222311123321"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-1002023121010233-2202310032231001-1033211001111301-0230103012300001-1333032021233330-3310221010100120-1020010332012323-3212222311123321)
- [disable_rate_limit](data-sources--http_loadbalancer--reference--group-017.md#canonical-2302102000331233-3112332122210222-1012101130112301-1003310221300013-3103112033310210-3213311023132331-2133321121033023-2003221013233103)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-0111113122010303-2301222300220321-1211003332003113-1321232202233231-0110133202230113-1313110222123003-1231230132033110-0121210312001113)

Select alternatives according to the provider validators above.

<a id="canonical-0032010310322011-0131313210312131-1103222132030011-0001223101303121-0320322031001223-3302123311213332-2200311202200231-3121213101031021"></a>

### Direct properties for `api_rate_limit`

- [api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202): complete subsection reference.

- [bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112): complete subsection reference.

- [custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3013011132131110-1031303333020113-2201331231223013-0222022331001113-0210133030331301-0013322311301220-2310330013302312-0102233021031233): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3111133023102333-1120102130112313-3103100331031323-1110010110020202-1210323033311321-0221232303123301-1130212210120030-1332100011133320): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3012211103120212-2201223321132021-2021010120120230-0112133020310233-2120203202111131-3111213210002021-3101102323132030-1220030133230312): complete subsection reference.

- [server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201): complete subsection reference.

<a id="canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.api_endpoint_rules

<a id="canonical-0122223033122321-0013012013330031-2123131100300323-1131010022012233-2121223010232230-2121101031032212-1302300202123033-1003132003002123"></a>

Type: `"list"`. Computed.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-2321232103300303-1323313322332101-0210321001032101-0112331132122122-0132222102130201-0122130323331232-2332323131032132-2300200310013230"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-007.md#canonical-0131032323123312-1311303032130013-1220212230233012-2032221002002001-2210231022021123-3122132310322110-1200311001032320-0310111002030330): complete subsection reference.

- [api_endpoint_method](data-sources--http_loadbalancer--reference--group-007.md#canonical-0120011232003031-3130331002010212-1321200133001303-0121212001033121-3321001302301333-0302203321303322-3121030200303213-0330330012100322): complete subsection reference.

<a id="canonical-0003112212013023-1231232203103132-3133323002210030-3131203123023201-2313203122112213-2300301032010132-1233000231022310-1020322000010113"></a>

<a id="canonical-3002232122132003-3201231023201332-1320311132132031-1033032233112003-2233021203312022-2120202100001211-1001213213130223-0331013110011012"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--reference--group-007.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--reference--group-007.md#canonical-3312223301223300-0100320231121331-2200103031212013-2133203231122320-3200220223312311-1322011023113302-0321210220333333-0013201011200123): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300): complete subsection reference.

<a id="canonical-3333110001032203-0212003321310202-1010303221222201-1332113301010300-1213302121120310-3021131232122230-0302121101133033-0001310301010312"></a>

<a id="canonical-2313000123203131-0022101101121030-0112202000302320-3323032301223030-1021131213333210-0022231332110030-3300233110133223-0201330210123312"></a>

#### `api_rate_limit.api_endpoint_rules.specific_domain` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0131032323123312-1311303032130013-1220212230233012-2032221002002001-2210231022021123-3122132310322110-1200311001032320-0310111002030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-2102221102313031-2230022010102313-1202310221023203-3013211330012013-0331121211030221-1201311011021232-0202213030323310-3123323003303031"></a>

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

<a id="canonical-0120011232003031-3130331002010212-1321200133001303-0121212001033121-3321001302301333-0302203321303322-3121030200303213-0330330012100322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-3211233130021311-0302321123031231-3003201222320113-0103021100131332-0320323023301133-1211330023330123-3330233231123031-2001100112130020"></a>

Type: `"single"`. Computed.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-3223030303112203-0012323131332033-3220132320230301-0020222032201121-2332312023002100-0010102201200332-3100022110310020-3200122132012000"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.api_endpoint_method`

<a id="canonical-2320233213300102-2220031030220010-2002010311033210-0002221331320300-0102101111331300-0020313010323112-2222300303233212-0000101130132220"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.invert_matcher` property

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-1123112301111222-1323302210322302-2222102122121113-0210213302123023-3233101233002003-2232203330130022-0300222300321013-2001133112100101"></a>

<a id="canonical-1312021231323332-1301320221312120-0232000133332332-2011200211201310-2033012000101233-1303102231111311-0113122003003023-2012110233131323"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-1231232003220032-0133331130110313-0222302202133231-1111032010102121-0131111032200322-0011323001330200-0003223330200023-1100230031121331"></a>

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

<a id="canonical-0032011321201021-0010120012202033-1311131133101231-0121033020223213-2303002012331012-2333101330020133-3212231233033223-2131201103131210"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher`

- [any_client](data-sources--http_loadbalancer--reference--group-007.md#canonical-1032011112101203-2130202330201002-2201131232232301-2321200320332103-1103213000033330-2300212302133203-1201122011331321-3212312133230031): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-007.md#canonical-1123023232132220-2011112031202212-2132112320331111-2212103312332022-2320022210130330-3302313102321231-0320331131001112-2230133001020233): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-2220010113322103-2303200312130120-2013201012023200-2132311102321203-3023021232321123-1010000322002131-1300101201102313-3030131102233110): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-007.md#canonical-1231020233021330-1100132112121213-3133223330101231-0213212332020011-0210001210131101-2232020012022012-1223013211020113-1302301102101122): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-1010101230032120-1332030020203310-0310303231312013-0313113302011012-3301123232320113-1222332331202232-3310202312201332-2322133301313021): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-007.md#canonical-3202230111330110-1103220101312101-0133122333132123-0132002123120000-1031320313310022-1321201102030001-2022113003101022-0130230103232032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-3121101201010002-0110102323110222-0122303223121031-2202001231013321-2000132332121310-0111312010022200-3030211033013112-1220311323120003): complete subsection reference.

<a id="canonical-1032011112101203-2130202330201002-2201131232232301-2321200320332103-1103213000033330-2300212302133203-1201122011331321-3212312133230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-1113032230013132-1303021002300001-3022022132202211-1002032022022032-1203131203210122-3111012301100322-2112321121002120-3030113333333311"></a>

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

<a id="canonical-1123023232132220-2011112031202212-2132112320331111-2212103312332022-2320022210130330-3302313102321231-0320331131001112-2230133001020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-2020331312220303-2112010102311200-2223003331213013-1223323233223321-0130003022122330-1333312320313321-3201210201320323-1001111101130301"></a>

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

<a id="canonical-2220010113322103-2303200312130120-2013201012023200-2132311102321203-3023021232321123-1010000322002131-1300101201102313-3030131102233110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0000230101202332-1213233221102311-2003030230120230-1220131331011221-0121233103200310-1321220302020002-1301023111311203-0201033120332332"></a>

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

<a id="canonical-0333102203130100-2013132213221123-3103200013313330-0130121013030311-2201203132132001-2202021311111022-0123111200321232-3313110200303213"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-3302133103233233-1213131003211223-0332233123121300-3311023223310210-1013331131102012-2031023211001022-1301332330222202-1111130001120221"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-3323310001133032-2230312322320222-2212321133321332-0002330001321223-3101032300321330-0030200102112322-0102010300322123-2032120211301101"></a>

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

<a id="canonical-3033300121030322-0323312000131200-3202223122022122-3113130021121201-1311113100302321-0022332130131222-3110221310231100-2212212020231313"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-3212013222131220-1232133311311113-1331310333133013-2302202001133013-0231000113321201-2203222221320032-0202103223110230-0310003021233122): complete subsection reference.

<a id="canonical-3212013222131220-1232133311311113-1331310333133013-2302202001133013-0231000113321201-2203222221320032-0202103223110230-0310003021233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-3233333231210211-1123111220011130-1002022021221321-2311102202230212-0233113323003002-1120131322312102-3201123021222301-2011112030300012"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0102120300200100-3322200202101200-2123212213202023-3023221123203001-0232030120033231-2323301101200023-0203323322123310-3011123003003233"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-1133323322102120-1331021300311323-2030210110020121-2012010330200032-3333002131211001-1020101311033232-3001333210133222-1132021212022300"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2233032110013301-3130202303300111-0223001120031130-3300203002032321-0120012121233311-3210300322223213-0222331122031102-2130211101010233"></a>

<a id="canonical-2030201132233213-1130332030231330-2111232011113202-2031013331020323-2020010112020130-0203333121020203-3210000231000220-3322222222222220"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1322121322311232-0131201001131311-0311031301311121-0103313331231030-2331210303130332-2323031001231103-2321113002331133-2200100233122100"></a>

<a id="canonical-2100133110111332-0323232130011113-0230223023313322-3232110001333220-1310302333130012-2212123023101013-2031221002203322-2300220102002101"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3033220032321222-3002311030123230-0031002131323300-3022220211333233-0332211210103022-3301031000003033-0131120333301330-0022002320230011"></a>

<a id="canonical-3031210330002330-1110001210133331-0110001313212101-3033220332332020-1013201211000010-2302133202333220-1123333311310223-0121212101311323"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2021122013003033-3233032123201110-1031113003133103-0112123212121200-1211333300200112-1310021011003120-1032003002302332-2311320002230112"></a>

<a id="canonical-1212212200101011-2132321122032023-3131022300132330-2321311012233132-3211323201102021-0233121300123220-1200233123023220-1332120023001022"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1231020233021330-1100132112121213-3133223330101231-0213212332020011-0210001210131101-2232020012022012-1223013211020113-1302301102101122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-2231100330101010-0000012002322332-2020333221320011-2102101203013222-2313133200330033-3023213320103030-1211003233300101-0220310012223303"></a>

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

<a id="canonical-2310300021232110-0130113320321010-0132210123023232-1010121231312333-1020203103133302-3223200132331201-3010210100313210-1013232130101230"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-2233322313022133-1121102010323132-0221211133011213-1212202203131201-3010322003033101-1312011211331122-3133133020010333-2221311330321333"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-1211101301031223-2032223202230130-0111033202222132-3100113100023201-2100303110323023-3120133132303330-0013132323100032-3003021111012030"></a>

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

<a id="canonical-3302111310321003-1132300023333113-2121231302210332-2303330320012002-2130112000312012-3330310001130322-2010321310202323-3222102023321230"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-1300321330102321-1022203133233312-2100110101000030-3200022220223033-0012330332200120-1311031002232113-2312032131330231-0321120121211130"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-007.md#canonical-1010010231003131-1100312201102222-2330212030011333-2111230333230330-2200130112301013-1132032002120121-2310303121322133-1322310330200023): complete subsection reference.

<a id="canonical-1010010231003131-1100312201102222-2330212030011333-2111230333230330-2200130112301013-1132032002120121-2310303121322133-1322310330200023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2302213321301223-0332321112323200-0113222222022320-1122322223301213-0330120221223202-3320333123103332-3212010210120201-2332022110232233"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0023012000003023-0323331032123100-0001000333101103-1222310322333101-1332100330121130-0011010231322101-3211131033101121-0322210210230020"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-2001003113100103-2200230203233133-0311212120320031-2231330301110301-1331222203101221-3223200233131023-2330213301111231-3232011203212031"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2013100020112322-1133202131332222-3200003230310131-0132101223302232-2322230330022103-1010013311221032-0203223300103222-3010102031121213"></a>

<a id="canonical-0221330302300333-3300100013330312-3232033310110302-2320310310320221-3120000021121022-1013023310222302-1012312211133010-2213200202023230"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2311300100332131-1021311132130222-2210200113323311-3313120200312311-3330330122002333-3333032313201230-3201021211200303-1212010033311121"></a>

<a id="canonical-1000120023120133-3231323213313130-2021232002212130-2320003013322120-1102020202333313-0112233301103023-3112013301023222-1333333232001203"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3212312333332212-0303311201230312-2013120132122022-2010021012123323-1021003100213023-0010123212313120-2311213321030300-0210310023331020"></a>

<a id="canonical-2012221120233123-3221020000300233-3003133211303031-3001011130201213-0233112230111022-2202301012000113-1230220020020102-0331302201321321"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0200013001120301-0231320132012330-3003001001203213-3120122121202122-3002311330232201-3231320100202330-2201102220223120-0231013133333112"></a>

<a id="canonical-0031020113031011-1120100031013103-3200021203330101-0212201223122223-2131321221121013-3000010201000030-0020303220001313-1322101101012021"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1010101230032120-1332030020203310-0310303231312013-0313113302011012-3301123232320113-1222332331202232-3310202312201332-2322133301313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-3303321331003211-3301103000213110-1332203031022320-3022233230113033-0302223021030313-1320132120222232-0013320000111203-1202021200100031"></a>

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

<a id="canonical-1332321033212312-2333300030012120-3230200112012120-2303222012000321-0002333302223132-1101220300123132-1123200100100022-2113331321103301"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-0113121122013110-1123033320322120-3121001003323110-1210030331333122-1022202210112213-2010111102303012-1120230330232003-0000110113112231"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-1101223331102312-0231133012312110-0213133113031021-1331232011201213-0011013310322012-1010003023322220-0332333222321220-2322032103201223"></a>

<a id="canonical-2032320210032333-2212002021022223-0111002321223231-3122321031112202-1201231322322201-2332322003112022-0202301111031320-0221303001012330"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3202230111330110-1103220101312101-0133122333132123-0132002123120000-1031320313310022-1321201102030001-2022113003101022-0130230103232032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-2320313231001003-1210211111021223-1311301030300302-2200102331211333-3013103130310121-3012102001031111-3113212322100010-2200122002331230"></a>

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

<a id="canonical-3330022000210133-3101003321232333-2303321100312300-2120231022233001-1321033303020233-0312010011312302-1120003221113021-3032220232332033"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-0131201310303322-2102231101012133-3300000122101003-1310032222203023-2302230123301202-3021203021320213-3331223313130203-1002030121310100"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3121101201010002-0110102323110222-0122303223121031-2202001231013321-2000132332121310-0111312010022200-3030211033013112-1220311323120003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-2011123120323103-2331221313112233-1312212201010023-3132021022302100-1202000010021010-2222010010211301-1131310133301332-1203133132000321"></a>

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

<a id="canonical-1001123321120332-3211022301211320-3010012302210213-2220010311032021-1132331013133130-1120021000112023-1102230310321002-3002302011322122"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-2302202231123311-2232031000213010-0223030101113103-1133220022222030-1321020310132302-0212300100300131-2133102203103201-2312002310022302"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3210310322303021-0123110013301110-1032001102213020-3302300311222011-3200102020133201-0010203001113320-0301111022033230-0101303012000310"></a>

<a id="canonical-0320113330011323-0310213112221301-2010022303132110-0313303332303202-1323122021330233-2112003313300203-1220231313123113-2301220010221002"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0132301321203131-1230303202210220-3232322120331000-0201033120013231-3000111212110203-2130120311111203-3321021313323100-0001022320212230"></a>

<a id="canonical-2213002103032322-0321202121330030-3123022021331232-0301231332333021-1122313113001222-2201322000113323-1033223201121210-1113201020330322"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3101102021110000-1132112123313200-3231300331121113-1321103033230323-2002333322321111-2231023000313023-1321222103032222-1223103202010320"></a>

Type: `"single"`. Computed.

Configuration parameter for inline rate limiter.

Additional upstream details:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

<a id="canonical-0012103021322102-2003230212000322-3200112203332102-3013223212332110-1211213130201121-2222131021221003-1322033012221033-1301000321030313"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter`

- [ref_user_id](data-sources--http_loadbalancer--reference--group-007.md#canonical-3101320012233311-3102220101011220-2210033131133032-2112123321312233-0021210032210201-0302211202333120-0021220122312233-2322221330233210): complete subsection reference.

<a id="canonical-1100311300330133-0120331011031011-0310010121001311-2303022133021331-2303200123202012-1323132120212331-0201310133123332-0100000313022302"></a>

<a id="canonical-3212313223112112-2232332033213223-0033032212110301-2311030101022101-2023120211002203-2303111122302322-2020310131022100-1300332023201130"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-1032131030202303-1030230003133022-0101332112200100-0303201113133211-1103302102101022-1232102100331310-1133302331203013-0201020033212122"></a>

<a id="canonical-3103330310111213-1231010130310321-3000321303111132-0311002233101230-2122213033003312-1111202310210103-0123122310112020-0233000312202112"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.unit` property

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-007.md#canonical-1021330133002321-3013030230000202-2121113211201030-3223010132103323-3120101212231120-3231122003202032-0000122003312312-0232330201011332): complete subsection reference.

<a id="canonical-3101320012233311-3102220101011220-2210033131133032-2112123321312233-0021210032210201-0302211202333120-0021220122312233-2322221330233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-007.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-1321320201211031-0322221313320302-3232330103131031-2220003023013133-0333210200210110-1100012313132303-0311303013101311-2211030303311033"></a>

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

<a id="canonical-1022010200211310-1230333002131232-0132011203230203-3002210333012001-2312332110303230-2302221011312302-2002031332100031-2122113012111213"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-1122331202032113-1010001311210321-1121002130200320-0323223202003230-3312022111030003-0130121121131022-1012220321030232-3222131301221002"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1003222231002230-3202013030111130-1021022123020133-3132021212020313-2003210323223322-1223223013210232-0021032120130303-1331202131010221"></a>

<a id="canonical-1013103002023302-1302111130011020-3203333111133030-1221211301112013-0022102202132030-1002223010200112-0011211330023220-0010312302203302"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.namespace` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0311100201121202-0322303021220200-2213320001321210-1012303012331103-3211322131331233-1220130002013100-2330301231213133-2200121103103222"></a>

<a id="canonical-1231310103330233-3001100211222033-0112210212233303-1113231301200133-1331320200033301-2313332301212211-1333231333333233-1100210103200213"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1021330133002321-3013030230000202-2121113211201030-3223010132103323-3120101212231120-3231122003202032-0000122003312312-0232330201011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-007.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2021332132333023-2003121012232101-1331012133202220-0310010101322211-1110302231021032-2322102111130220-1030201232221312-2133012310230031"></a>

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

<a id="canonical-3312223301223300-0100320231121331-2200103031212013-2133203231122320-3200220223312311-1322011023113302-0321210220333333-0013201011200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-0031201023112111-2313112300033000-0021222010333331-3313312020233223-1231321320200220-3020123211200210-3111123021131321-0313301202102011"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-1103210001111322-1103233212032120-3013110032123000-0303323331201221-0101030213333123-2003212121232223-0301331110321203-3101102200213111"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.ref_rate_limiter`

<a id="canonical-3112301123233322-0223010032220032-0132323133033123-3011203130321123-2331303332312133-0133133303331101-2002102311212302-0302320120232310"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230210201322023-1030112313231031-1313312201120320-0322121301120213-2101101200331323-2122333321021002-1012132313322123-0002310100010321"></a>

<a id="canonical-2022020011000302-1312301230221111-0322313002122102-0311000122213222-3301021320000130-3110311023032110-0322312020200023-2220103220213232"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.namespace` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1212000330233321-2210001022010302-0101133331220333-3120332323013121-1030100123321021-0122120211233323-2132313001002103-1220332213311210"></a>

<a id="canonical-1311113102320010-1002301223200003-0332330110320110-0001313221023132-1200331222221001-2233210220121112-3120330110200133-0131322103002121"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.tenant` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-2132122232022331-1001301233323123-1001310002013103-1213122310203100-1101112132302013-3032002110000122-1120201130103203-2103032100200022"></a>

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

<a id="canonical-1222330311010323-1111020133313123-1022232021232001-0013023223203120-1310112112211230-1131202121102103-2031011303111020-0320120021013320"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher`

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-007.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-008.md#canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-008.md#canonical-1210330031123323-3321033310123200-3323221033120330-0302113212111112-1231012323312022-2333330011003321-0211002010211131-1233302200112310): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-008.md#canonical-3102022311222332-0111210032013011-3001010132122001-2311201033202222-0200210203210312-1110200201313233-0102223200212201-3130023112023102): complete subsection reference.

<a id="canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-007.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-3332231002010300-3112202131223003-0111321012323030-3202021002032113-2310103311122133-3120200132110200-1233201313102123-3030013233321333"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3031111103031300-0012210121113122-2021223122122023-1323030020121020-0311131320011321-3322202200011333-3120130013202133-0203223032213001"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-3310231102302311-0213330221333201-1121101220021221-3201012321331203-0002133303320111-0020312313312312-3030103011020111-2113332232333313): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-008.md#canonical-3221211201111101-3212123333300020-0002222300300310-2331133003231201-1220031310103213-1320333200313223-0110032121121220-1012003011101301): complete subsection reference.

<a id="canonical-0022302111331023-0231332322220321-3200123122212330-1300101012133132-1121011322231323-0013113200102000-3232033302022330-0210232213000012"></a>

<a id="canonical-0200012101010302-0110323123312330-3200132112022210-0231001132131300-1122030213222212-1223132133020123-0103333111313030-2333112110211233"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](data-sources--http_loadbalancer--reference--group-008.md#canonical-1230212311222221-3222232133221330-0210200113023121-1233210133201313-2130200110211021-1013113303110113-2030220201103211-0111200012132332): complete subsection reference.

<a id="canonical-0303012221210330-0332100022123331-0100033002221332-3122202122032022-2331220300211100-3112213120300130-3322322301133330-1312000201002303"></a>

<a id="canonical-3131121303100212-3002233003120033-0002222203203332-1021322013201233-2230133121333231-2322030013101333-2023301332312210-2323011131020003"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.name` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
