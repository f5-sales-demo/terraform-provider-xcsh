---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0303313001011331-3013111201302210-3221001030022100-0212032220002003-0010322213222320-2011212120232211-0111122010022133-0033222102323100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0232103200323230-0302112110100212-1233103001313000-0200120001003020-1301023203301332-2201312102112323-0121303213332200-1001101202310332"></a>

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

<a id="canonical-1230001031221321-1231110213302320-3321302302332212-3212032331131000-2000311220013321-1223122322220331-2321103132210203-3211210131103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0233310330203330-1133122301000202-2103302330330022-3302320202333222-2302131131300333-1103233222323012-0100103201123311-3012102322310010"></a>

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

<a id="canonical-0211311231132120-3223313321010303-0011030101021130-1013312222223210-2321310020110312-0112110012000111-3312231300232311-1112130220332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-1023331123231202-1331222130023320-2210020101303122-2320212013223322-3031133012013202-0230030000202230-0122100130000331-0210113331322203"></a>

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

<a id="canonical-3211033102310210-2032223020300032-0111130223313002-0303100230102320-2103322303032102-1010001333230022-3202200333132310-0012213000132122"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0113100210213320-3010111212311230-0032101123132213-1123102131323120-3212101301300233-2201022101130210-1223333212313001-3031331230131221"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0313121021210021-2122203332113102-1232223113310000-0220333113123022-2202011023111222-2201221111232220-2320103101230232-0013300121322310"></a>

<a id="canonical-3330323120113222-0011210023232330-1320000313112031-1111310013112111-2132101233103200-3101312302033321-0300302333123301-2320233100133100"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2320133020331232-1023103330302333-3103033212013010-0100100010013132-2113321001002221-1031002021010302-2203321200320020-2213223303011332"></a>

<a id="canonical-1303220020133013-3330000110233300-3332012102221301-0232103100001000-3300022033000301-2312102002023301-3021313200021010-1313230130130002"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-1033001220002130-0030113313132331-0221122033122312-0020111230201223-2102112233223303-0013211101111022-0010220032002132-0003102122221302"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0130330330133003-1010323212130000-3322212101111122-2312220013332230-1313330002212101-2313031222231132-3133211322020103-1110031310323221"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2300203302330000-1113331120312103-3012203011332233-0231220202132013-0332000211122220-1032323002122312-3121232010322011-3322202120222012): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2000300230213113-2310133032200222-1032320100112111-1121202103203001-1310131122302031-2200001230003320-3132213322103310-1222303302121103): complete subsection reference.

<a id="canonical-3100220203110231-3331220031211202-0101212000300122-0203101102101100-1221231112332231-3312333033020122-2003022000223101-3020210013031131"></a>

<a id="canonical-0012322031223223-3100133013222002-3132320120202000-0102123230110230-2212321232332130-3230123111000023-3032200313031120-1322222201011212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1010221010113231-2210103210122133-0230121223002302-0113001011213211-1320032211031102-1003011332233332-3312131133133110-0202123112123232): complete subsection reference.

<a id="canonical-1233303031123223-2233002012221312-1211033012011330-2003322100002322-1010133222010301-2013003221312300-0312123331322033-2131033231231021"></a>

<a id="canonical-2001100012222130-0331123222331330-2222323203103020-2113231300322111-3020120331021310-0320103113232231-2203032212132100-0333113133232321"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2300203302330000-1113331120312103-3012203011332233-0231220202132013-0332000211122220-1032323002122312-3121232010322011-3322202120222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-0033232331300122-2300112210123211-0022302132021320-3332033301102230-3103012112232132-0102331332002301-2201332320032212-2210231302002021"></a>

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

<a id="canonical-2000300230213113-2310133032200222-1032320100112111-1121202103203001-1310131122302031-2200001230003320-3132213322103310-1222303302121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0032310121110333-0013011213330013-2302100012330203-2230021011111121-2332032033002020-2033313113213313-0110203010013211-3103321032203021"></a>

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

<a id="canonical-1010221010113231-2210103210122133-0230121223002302-0113001011213211-1320032211031102-1003011332233332-3312131133133110-0202123112123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-1311003221332203-2003220232312013-3223210020322103-3031220133021211-0122230020003023-0310203301132303-3132323133100312-3122012211000123"></a>

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

<a id="canonical-1002313111110121-1112112100312300-3321012213120100-0023110133222201-2311303132301203-0200100030011220-1223131031112211-0121220000122332"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item`

<a id="canonical-1323021200231111-1112021002112330-1202121112212011-3113110010231212-0000221233111300-1200330201220012-2033103101333000-0033331201002103"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2013232022312213-1133022203211113-1332231023032100-1020001131102123-2031302330221321-1121131331020203-1221333013221021-3221121122120012"></a>

<a id="canonical-3021010211130310-1230120332201330-1322313011302003-1201121003233202-3123330321131122-3302312200033130-2303231313011012-3002230101030211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2323212101000101-2333130001001100-1332213221023123-3300330303320112-1233023102103332-2321302321331230-3102131332033230-1210321223132031"></a>

<a id="canonical-3320131321033330-1000012301032123-1332100323322322-2211210323113300-3211233102021333-3222121133132320-1022100003011331-3233102110202101"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-3323231002312000-3230323202122303-2201113003033321-2011203322020130-2313222323322102-1133200213011021-0022330031021132-2022130002030022"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3311230003200211-3233321323101330-2101212301013133-2030022003213112-3013132030312101-1111213203202202-3132022010232311-2031032222332232"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0101331301111133-2211223010200301-1321223220203210-1332103201031321-1012201102102133-3110220001003103-2201130211223200-2133123213002233): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2100123100123021-1121003300111102-0002302310312001-2021003002133103-2321010011320231-0123300032230232-0133221211300032-1031120221003303): complete subsection reference.

<a id="canonical-2101003222332100-0031321112200020-3320001030233220-1302113313222220-0111023102112132-2010213131202022-3133301210023322-3133012303213132"></a>

<a id="canonical-1113201312130011-3020212021133333-2301020331312122-3231011022030222-3122222020303311-2303321030322201-0103320032100023-1110120322311013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0022301003133103-0302212001200223-0333231303101300-3000012113221001-3202311202030310-0230123023200131-3221021200300203-0331210010313022): complete subsection reference.

<a id="canonical-2322032333200201-0201033031301300-3320002223122201-3332231330231232-1121322100203202-2000102223310012-0323101110031312-0130233222030103"></a>

<a id="canonical-2020230221022311-1012123030132112-1320122112030031-0323032031023022-0101100002223233-0132302003023303-3111130232310020-3021320111023120"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0101331301111133-2211223010200301-1321223220203210-1332103201031321-1012201102102133-3110220001003103-2201130211223200-2133123213002233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3211231002103021-3330213223121221-2113010031100310-3023312120002322-1022321321201123-0132232110223220-2032302130312311-2301203113011221"></a>

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

<a id="canonical-2100123100123021-1121003300111102-0002302310312001-2021003002133103-2321010011320231-0123300032230232-0133221211300032-1031120221003303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-3233231322312223-2201321222103230-3222112112100200-1233111032313200-2113231100311012-2310102303222022-2231300112001310-1010301121313113"></a>

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

<a id="canonical-0022301003133103-0302212001200223-0333231303101300-3000012113221001-3202311202030310-0230123023200131-3221021200300203-0331210010313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-0211133122220023-2022213103003030-1332210010202120-1331031321002102-2021332103030201-0213030013122310-0212212012113231-0300333020302033"></a>

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

<a id="canonical-1131121021122030-1102332322003210-3032202031020320-2320031123233223-3323221231033320-1213333313330101-1212020223030023-1123030122312231"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item`

<a id="canonical-3311332301313122-2231202301203333-1133203231002300-2102020003303300-1122110333120221-1321203103231320-2333212023113320-3231020321011033"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3223210223313001-0000212122102102-0213213132331003-1231020012302122-0222113211123202-3130113110100030-3313002112002121-1130130321133101"></a>

<a id="canonical-2323322013002112-1130320212012223-1211202101312131-0103001312323010-2322020313111021-3101301211011113-3000331013333221-1323323220132123"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0212010111122313-1233023211233202-0132331320130311-1323201223010100-3221321200203221-0122311331000322-0002312023012301-3221132022201321"></a>

<a id="canonical-2122001212331111-0120302321020033-0221323010010203-0100211033331013-1103010003003300-2223213112300131-0330331102113211-2012321213013010"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-3031133220321110-1000121122002031-1212210332313202-2232211212230210-0311011133222332-0031022332322222-3030202223131202-1121101111232232"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2110021323003230-3023022112332121-1011120221013322-0301231023211313-1112223133121033-0223233301121120-3313203322313022-1300012101101012"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1320011110310313-1320210001131301-3300232211101031-3223111133202120-1221023313222120-2220211003102312-0332031233021320-0002000012132203): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1001312323012320-0021230103230122-1011100033123200-1211332213132112-0022112331233202-2322010222123030-0001201312013211-2220031301021222): complete subsection reference.

<a id="canonical-2000323012011030-1321200133122300-0111013003332133-3220030132230220-1130200211302010-0113012021232100-1202330223211212-1231320230101222"></a>

<a id="canonical-2021111230330202-2103110110011220-2201222320132002-0023223121022003-1213023012211311-3113333301103013-1212232331302212-2231220332231013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.invert_matcher` property

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

- [item](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2032111031231023-3323001331310000-0030332112320021-0120212231100033-0033032000130020-3203213300111123-1300033023231223-1221233000323130): complete subsection reference.

<a id="canonical-1302013032302213-1320213311100003-1313312113103020-3223133003300123-3023221130010113-3221010320222121-0002333303320230-0202211202302111"></a>

<a id="canonical-0231010312100123-2132333202033232-3230312000303312-2031323230110131-1221131221000020-1202110200013220-3330122000211312-2223203012130100"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.key` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1320011110310313-1320210001131301-3300232211101031-3223111133202120-1221023313222120-2220211003102312-0332031233021320-0002000012132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-1131011122020001-2332223030311222-3213233201231013-1213131221121203-3012320010132221-0020212010101213-1332011222010030-1022133112032310"></a>

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

<a id="canonical-1001312323012320-0021230103230122-1011100033123200-1211332213132112-0022112331233202-2322010222123030-0001201312013211-2220031301021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-3331102021310021-2021201000103021-1223020233313332-2232001232132131-0313233223300213-3211132233132220-0221211231200012-2132213023020131"></a>

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

<a id="canonical-2032111031231023-3323001331310000-0030332112320021-0120212231100033-0033032000130020-3203213300111123-1300033023231223-1221233000323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-1103113333002122-0033131231303131-2013033213101021-2212211012001300-0133302012233300-2131132313312332-1311320232330113-3233031233221023"></a>

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

<a id="canonical-2200101033200030-2033223003200230-1131220213102130-3321310302030032-0102001021313302-2120312322011212-0012013210331112-2133011310013001"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item`

<a id="canonical-2202332320121033-1032211110330233-1020011132230001-3120301223000122-1023003231203203-3112113103003232-1223321022123100-3033223200231022"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.exact_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2000011100033310-3201032133033201-1321010003330220-3233312101113121-2332303102120132-1221000301203213-3200322302031323-0101302021032031"></a>

<a id="canonical-1032110332022210-2032200032201331-2112202101111130-2132102031103110-2012321203322222-2212010012230300-1310331302321313-3103202012100200"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.regex_values` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300133213310201-0221203322300210-1000301200130013-3002231001101221-2323201231013323-0200133113230021-3013032010030310-3132302101323220"></a>

<a id="canonical-1101230000100013-0322232301310122-0113303230113321-2313003211212100-0010020213132222-2003030000033330-0230013233121122-0202332101111312"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.transformers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-3203200313103020-3121213313310212-3000230120202011-1230112232000303-0310233221120122-2210100231320222-0202311113121122-2132332023101132"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

<a id="canonical-2003332131330322-1331120120110312-1022233130020232-2023013213203122-3122100313001320-2023301300120103-2221330100020100-1300203322110131"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2313001303220112-0111323121220310-0310320000120011-0220012232331323-2100303120130213-1201010102123002-0323100303001333-0310210030330301): complete subsection reference.

<a id="canonical-2313001303220112-0111323121220310-0310320000120011-0220012232331323-2100303120130213-1201010102123002-0323100303001333-0310210030330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2113320301310110-3231010300201120-0032211333111301-0221220333200101-1233020330211303-1301212333303212-1123213201032011-1221023313300221"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2013033013111213-3212211322021201-1200201223210333-2320020023212332-3213210132232020-2210331211103210-3233323333330211-3321322122213022"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-0023330331323203-2320023113002103-0310223330322022-0132100311001110-3333220202322332-1312103010121210-3021002100130012-0031020101330020"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3031013110031200-0001000331120131-3233122221000132-0111333110330102-3023202032232301-1213222313311312-1301200110331300-0323311030000022"></a>

<a id="canonical-3131103113202113-0222201211031300-3010031132032100-1021330020331211-0113212022322303-2203133301222321-2212211311101203-1230301003330200"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300101201230302-0011032222101122-1002332000233213-3220101333213110-1233300010030310-2113013021112200-2102123012311103-1233013233212023"></a>

<a id="canonical-2103300230013203-1123110032100311-2122230021320303-2202133100112230-2120113013222012-1331022113221030-3212012303231000-3322222200100312"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2303033302312111-0231012102313300-0030302012321311-0331312000002132-3021100202110103-2120133133300232-0230222110332003-1102203220130300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.ip_allowed_list

<a id="canonical-0012001300022000-2130001230313211-2221323102232232-3100103032303302-1302032023330213-3233010022022320-2312023101300013-3001220003103223"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-1220222010122323-0333202130303321-1022323021032201-3132323113101111-1103131113200013-1102333212210020-0011101003333111-0121323021101233"></a>

### Direct properties for `api_rate_limit.ip_allowed_list`

<a id="canonical-2103220021203201-3223222011013010-2012232212022023-1311133030132321-0333113013221322-1203133012123213-1003233310301121-1233312131303112"></a>

#### `api_rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0231102032303012-3001232002310032-0303213312332003-1321010302001031-0232222203033013-0102231202310231-2103133110112032-0131120102211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-3122331200202103-1200222131323233-3310310030110112-3020000131210012-3300223210233110-1011232103103033-3213102211303303-3220000312300200"></a>

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

<a id="canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.server_url_rules

<a id="canonical-3320322303200332-3123220122301100-3321100303231010-1300133230333231-0222222022110002-0203301103020131-3331211322021012-1231102222110123"></a>

Type: `"list"`. Computed.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3000020213111220-1132112310312231-2223322222000202-1132013203132303-2102301310220010-1121011001111303-2213002002012232-3321302203210232"></a>

### Direct properties for `api_rate_limit.server_url_rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0300301130330202-0121200322320021-2203230032110133-1203021000102121-3232212110321103-1112100323312302-2302203002331130-0301101333130320): complete subsection reference.

<a id="canonical-0031311333032223-2102122201202111-1332210000000121-2211311121121202-0130232233313301-0110131202033032-1311130312121330-1021330230003213"></a>

<a id="canonical-0023133200323320-2123321000201113-0312120113103313-2131022231002111-2231003300231030-0302110121130002-1302212322322231-0032320221331201"></a>

#### `api_rate_limit.server_url_rules.api_group` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2000102231332121-1311221013212031-0011323303230113-1201233332330311-1231301312101330-1121313002000331-3031333130220211-2230013223031320"></a>

<a id="canonical-2100011202010303-0203121103030312-0011323130223311-2213112012321330-1101023301303302-3322103031333111-3101331232010203-3110223132000301"></a>

#### `api_rate_limit.server_url_rules.base_path` property

Type: `"string"`. Computed.

Base Path. Prefix of the request path.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022): complete subsection reference.

- [inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320): complete subsection reference.

- [ref_rate_limiter](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1121003122321003-1030113023222001-1120323002113122-1103033302313332-0220011200302103-2120112030110322-1320212132033201-2023201113201010): complete subsection reference.

- [request_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303): complete subsection reference.

<a id="canonical-2110202032232313-2103023100111131-0322013102221130-3110322123012222-1001200013032203-0022010211100023-1120321101333120-3203332220103103"></a>

<a id="canonical-3020300130301330-2321133331212000-3003012122010323-1112022312100333-0313333132300120-1213333103131031-3210012331203203-2310303130023223"></a>

#### `api_rate_limit.server_url_rules.specific_domain` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0300301130330202-0121200322320021-2203230032110133-1203021000102121-3232212110321103-1112100323312302-2302203002331130-0301101333130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-2323131230113001-2202010331013111-1110011100033121-3322022113132020-2003222301111323-2312220130120020-3321132210303311-1221312113130322"></a>

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

<a id="canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-0210221010103002-2210223002321202-1112210121213213-2031231330131202-0201133332312213-3201113322130130-3220000213211132-3321202102233222"></a>

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

<a id="canonical-2213121330302002-0230030223210332-2331311020203301-2221202311301131-2222130212012300-1233021123013221-3211221132232231-1202111300323301"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher`

- [any_client](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3100333201213100-0233210101312033-0013330320301030-1032013322202320-2221113000023020-2003231333113203-0103321211101312-2012020331330101): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2123231201203203-0210221320012000-1301302130201233-2112110130230231-0323131223030222-0231313311211331-3123323031110031-2121213111032222): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0212313230132222-1100122132011210-2331330302223220-0311333300303323-3221111020031331-1033002130113000-1133100323232332-1312122232023323): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1102311331212020-3211200320200222-2332222313232022-1303200130323321-0201321202321212-0130211201023332-0132232331302303-2110212200321220): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1111210112232110-3133000020031222-3110321003330330-2102221330302033-1120133301333103-1120330011032222-1323100133211013-1203121330010323): complete subsection reference.

- [ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3122022200133313-0102331311012322-3021212000002223-2000203333200202-1322331230332033-2221130003120212-3032002303233211-2303213201111032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1331332223312102-1132112313212300-1023223213002222-2333213112000033-3131223010112201-3111213233133211-0202302112202003-2021131013002322): complete subsection reference.

<a id="canonical-3100333201213100-0233210101312033-0013330320301030-1032013322202320-2221113000023020-2003231333113203-0103321211101312-2012020331330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-0011230331302121-2302111030123220-2010100332311213-0012100330331120-1231212303201331-1303210223100113-1003000021301312-0232323112100130"></a>

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

<a id="canonical-2123231201203203-0210221320012000-1301302130201233-2112110130230231-0323131223030222-0231313311211331-3123323031110031-2121213111032222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-0302120331313321-3012112201120331-3223333112121331-2302022013322013-0210111200002000-2011221123112030-0213002323000323-1032031331303011"></a>

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

<a id="canonical-0212313230132222-1100122132011210-2331330302223220-0311333300303323-3221111020031331-1033002130113000-1133100323232332-1312122232023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-1122332223222331-0102103011002112-0113003012122131-0201200221110100-0332121322321321-0330213233130230-2103133010313121-0132302101200233"></a>

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

<a id="canonical-1013102102222130-0211200132123203-2020100332123111-1202103000300323-1132233120132112-1231302012223202-1003113210023001-0112100023202303"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_list`

<a id="canonical-3201000120301201-1322203010031232-1332210022211110-2023023113302021-3101002101333220-0311322230012111-1331312021022332-2123323330223303"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-3333133203131103-1031131013301012-0113202331000312-0323120023112230-0022122011201123-3330003202312212-3211311322132201-2011111213003332"></a>

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

<a id="canonical-1313220031102222-0201021302022323-2213220330330023-0321213012321101-1032321102231003-3223330310233102-3012010203312002-0201102131121301"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher`

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2220121222022033-3210122112113001-3031102003233303-1012310013103003-2211013302211113-3130313322223303-2200001230030220-2123011101033110): complete subsection reference.

<a id="canonical-2220121222022033-3210122112113001-3031102003233303-1012310013103003-2211013302211113-3130313322223303-2200001230030220-2123011101033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1013030200122331-3022133322232030-1010103131020322-1131110213111200-3312221031022131-2001032331320331-2123233312121202-1102100222101222"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3033200330110130-2133130020230030-2023322102100203-2233110212031100-2203300321110201-1103212312210322-1010031020132101-2033032021131331"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-2132123033213203-0322213100312110-3031332330221003-0000222033301101-1031221223331333-2113311203022022-1023231123230003-1231323333102103"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1203120311113130-3212033001123200-2030201312231203-3001020013111123-3312221031001102-1110102122132232-1310031111000320-1000002110122032"></a>

<a id="canonical-3002121111303122-1220310331201122-1312111212322111-2020103022321321-3020022030313032-3110230222002311-1210330331300013-1011323330302323"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3122100131032102-0022102003013312-0313023201313320-3131110301332203-2210211002011122-0111133302100113-1221300211132233-0222201033131203"></a>

<a id="canonical-1201212310332313-1010331213232123-3202103100011311-1021233100110300-1101133331223322-3320331011010321-2212220212110220-1121011012032230"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2111121131302313-1323101312303303-0312211132300023-3103031220020130-0302310133011000-1332122013030210-3031011010032032-0221122321032012"></a>

<a id="canonical-2310323122132032-3320103333230211-0301313010120113-2000021230231310-0122131013113000-0132323011202313-0231011333232221-0131012132321201"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2231332030330333-1012010322022001-3311223233111102-2002022033301313-1100112312032201-3223101001312301-0110300322312003-0022313230313032"></a>

<a id="canonical-1322200231320003-1030320000120223-1220100333233302-1022301220312321-2120120133233103-3222103301032121-1123230202210330-3323120200103201"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1102311331212020-3211200320200222-2332222313232022-1303200130323321-0201321202321212-0130211201023332-0132232331302303-2110212200321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-1312031113123331-3221301101131122-2202003322201013-1321302202213210-2130302102322331-1331130310312030-1303032022221211-1222110201101212"></a>

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

<a id="canonical-2131130111103033-0302103313230013-3031013033020021-2103002011011111-2210212021321022-1111003011201111-1022122013030222-1222030010031200"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.client_selector`

<a id="canonical-3020321231322333-0123100220222000-1022231001301233-2100132111103101-2220313312303201-3210303321212110-0020002222230312-2300121012213033"></a>

#### `api_rate_limit.server_url_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-3000310002313233-3001331303011030-3011322130132023-0230102112302013-1011303122122003-2323100121003202-3223202012200123-0120132023121322"></a>

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

<a id="canonical-2022000333100222-1002130121203333-2221320031120312-0032330033203320-1332012210010223-3102121221202231-1320101211232112-1103212002330003"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher`

<a id="canonical-1220200011023323-0332210011023333-1301122312201222-3002210332311133-3233222332123330-3332112013000102-1220132110301013-3302311310011221"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0332010221101322-0322100113002212-2103310010322103-0222021000202312-3013222033311032-0112203133013130-2010310120112022-0010113011221321): complete subsection reference.

<a id="canonical-0332010221101322-0322100113002212-2103310010322103-0222021000202312-3013222033311032-0112203133013130-2010310120112022-0010113011221321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1321232022111010-3100033302000012-1301031030003033-2003331011103313-2220132002220222-2330320310333003-2233112203332301-3120120122223321"></a>

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1311322312212211-0330313300110111-1011233333312131-3011030212131310-1211103111000301-3212213323211213-0310302101201030-0310112133211232"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3313202002001101-0311303213323301-2132323333323300-1310321322113333-3212010321020032-0021202032332133-2310201023220233-1112003232113201"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2200333022333201-0233123023302100-0000130230223101-0012023322031333-1223202130220323-3123132201120222-1322132100232321-0201222032311011"></a>

<a id="canonical-2111310231211133-2031130013310011-1312313310013123-0012321012132132-0323223303011110-3022212220013210-3130310021102321-3312130121031030"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1013233212103223-1110132021220213-0001302111022313-2113203233111030-3202021321023012-1310002222131102-3101321303223301-0121031322233112"></a>

<a id="canonical-2010231223302223-2203103031113212-1001132331212333-2200112222111103-1123213112331322-1012011112001233-0001200222012313-0213112123011130"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3312331300003120-1013213231313033-2021130122100331-0332210001230133-3013011203201221-2132101111133312-2313211121202130-0132010131110030"></a>

<a id="canonical-0120133020230230-2011203203201003-2332133313311011-2332331023101212-1223333113230200-1320101031310233-0231301210200331-1200312323120321"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012132320313213-0102133131020110-1020030120100030-1310112302033013-0000002130013133-0131311012311223-2210123322011213-0000031022101101"></a>

<a id="canonical-0123133013321132-2202231122211222-2212031122130022-1122022230012030-1032100230112203-2231321020121123-0200311301003231-2231011322133312"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1111210112232110-3133000020031222-3110321003330330-2102221330302033-1120133301333103-1120330011032222-1323100133211013-1203121330010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-2121120311322110-2030300010233133-0031203330123032-2133330102112022-0000101121032132-1031231231222222-2223122113223302-1110322021103130"></a>

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
