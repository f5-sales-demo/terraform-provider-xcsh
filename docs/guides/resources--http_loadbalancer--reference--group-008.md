---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3110132300130231-1001323012021233-1322002021013221-2222010013012322-0222103321332112-1030322002210110-1200022230012302-2210330102321322"></a>

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011): complete subsection reference.

<a id="canonical-0132130122112023-0001303101133310-3313223232121300-2023203133000321-1002030020230220-1130121023321113-1331102323332203-2023323231312321"></a>

<a id="canonical-3001101202131231-0110023320103220-2220020210200023-1330031100200021-2111303332203312-0113312133301301-3221312320310023-3131130210013001"></a>

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0223031032222220-2011222120210323-0313210222321312-3332020121102012-3023233220310222-3333231323122322-3031110121133130-0032123332230133"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0221221121110211-0011320032123021-3132211231112112-1031110122320121-2302012101123220-2031032020302320-0032120323300121-1110020102332120"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-3323300233200130-3222030201211321-2102131320212131-2203103333102010-1213303132000312-3132301002333323-3130123102102012-1122222003121301"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213103232011222-0022332000102330-0120032000211003-0113331123000202-3330222032132310-0310302300233131-2303222212011000-3300221323001312"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-1021212103122213-2110101122232131-1201331031002232-3210213002003300-2220212033131120-1133212213010032-1033032323322001-0203211012020313"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0130310310022322-0312321331113202-3112300130321311-1100122300220232-3201231332322113-2001102301211233-2322233020121011-3000122121033013"></a>

<a id="canonical-3302213221000303-3002001133003121-3222323312231013-0210131010333122-1032322000322021-1032130203132033-2201001113113333-0100030123021332"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3231101030222201-0003100232211013-2132030230202233-3111120130101223-0113103210130100-1331210113001201-2122302122332323-3302230331202303"></a>

<a id="canonical-1332222213233103-3202310121301312-2033102231100323-1310221222013030-2000322232220133-3233310311011123-0002302003212101-1021210023210012"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-1112032022303031-1310321332323132-3301000302203120-0220003021312212-3112123202311302-2320130023332300-0123211232001320-1203213232220003"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211031310010032-2130333323210311-3010331113031231-3211203030320113-1202211232323113-0201302313022322-3232111132100011-0312203333311032"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220): complete subsection reference.

<a id="canonical-0003311303201101-3001301131210220-0113133231023113-0100230323012101-3202102331123210-3131130302323002-0121012211100310-3203331201110020"></a>

<a id="canonical-2113003113023300-2333121211303002-2121303322202103-2213101220112000-0333102030120022-1311303121102033-1212320023130211-2031221011330013"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023): complete subsection reference.

<a id="canonical-0211313013102023-0210200002230113-2132112032011012-3200111020100310-2311202323323100-1003110023012231-1111100200313033-1020303131112123"></a>

<a id="canonical-2102103113310011-0220023030312103-1233331122133332-3222131220031311-3120031311222113-3221221321130110-2311121030132022-0201303022001120"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0002133200213133-2001303203103011-3121130030231223-0101333322221100-2210330112200201-1220020122131101-3311021202011110-1103333230223202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3103033330003201-2033103203322000-0110123113003200-1110120011121023-3322310311023002-1221122230002003-2031021332022101-0320223210203032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-1310102202200201-0012131313200012-2310132112111013-2313203023312123-2012111100002032-2230122032011323-3221101013332230-3023120021133201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332213103223311-1130212010212112-3103111002311032-1133113201211233-2330130223122032-2223300202023303-1330321032223133-3322311303133301"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers.item`

<a id="canonical-2133210111223122-0211131002310221-0100231003132032-1121203031210100-1002012323121002-2200101203201021-3202022203132200-1311233011122213"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1331210103033311-0231111001123001-0010123322232320-2100333310010120-0111213212313100-3022131011233301-3113212323213110-3122111033333110"></a>

<a id="canonical-2023113111020323-3002012212210320-0300213120111233-2312033212311020-2202122211312301-2111101332310100-3332330112230301-2333231001232320"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1331332311312310-0212023323302022-1103300200320031-3203022133313333-3103111102232130-0010123322311122-1030100302332030-0123231323232303"></a>

<a id="canonical-2030011030201203-3213103310201312-2320311011220113-3232003200201033-3113020033313010-1133322221131021-3310000221320302-3013020011122233"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0033131120213101-2021101213122313-1112301101110023-0020303203303030-2330012203010213-0231302000220110-0012000120012123-0022210112113322"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102303103120111-1033221103122221-0110023031300311-2111300010311303-1102310331312003-1210002120321220-3002213232122302-3112320320301033"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133): complete subsection reference.

<a id="canonical-2122231301331010-1320312333312003-1333001210032332-3133321021031021-1121310322232221-1102212011201031-1210332132002020-2122103210333102"></a>

<a id="canonical-0330200333130200-2232333110333122-3302202120023122-1131320000120122-2231031021103310-2220202222322202-2012023101012221-0211020213301323"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310): complete subsection reference.

<a id="canonical-3132211312321102-1113230313203121-0100021113222112-1033010200002002-3110030330231230-2012320131232123-1333002110200220-1322013302223211"></a>

<a id="canonical-0023321122032003-3132221333203213-0232302330321203-1233132332300332-2312011033023311-2101130213212313-3001232032203010-1201301001122002"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2213212220111130-2210120211112303-1130310210031133-2121103010330113-0332100101232323-2000333022101030-0313022222222000-0333312321011210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312021030100012-0322330333213110-0001221031311302-2110023221303221-1121020023310031-3001112312023122-2321222121023113-0331010320331021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-0111033122123022-1130321103113202-0223200230000033-2013111311033021-3323213123011032-2033001202310023-2111203212310123-3210111313320030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330122212003101-1332121032223100-1111011213000222-1210012220203032-1123023300301200-1213333233002013-3113123311120321-2013221123212110"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item`

<a id="canonical-2312202120212310-1202200320110121-2020332330333023-3011320030031232-1320013020130022-0102132033232012-0130032203311322-3232121101130111"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3203213223302010-0032123323231331-3102323133330002-2112103131130121-2321312120231221-1020303310122111-0200311213203031-1323013020200323"></a>

<a id="canonical-0231213313201313-2230303030122302-3020100022310111-3023220032020331-1023300332331001-1230020333033233-1003231312221220-2130221203202023"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1111203301223120-2132332121032212-1221022221300332-1330302310322310-1020122230012110-2333032321233302-2223233203022200-3010213001123030"></a>

<a id="canonical-2202111113000123-2100132322100112-3230220232300311-2321123031333331-0021111011023023-3213031302202101-0023212032332233-2112233001203331"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params

<a id="canonical-0320330213122122-1313120301101122-0000330133300220-1332310200101313-1303200103110131-2102031203002021-1330111302122001-3102101320132001"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112332003110330-0033023331230301-1100310133233211-2321032233131102-2133200232032312-1031012311300222-2303102133303020-0113221031121001"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102): complete subsection reference.

<a id="canonical-3021330220001013-1331033222021121-3102121302132200-0112203312110112-2221102232021211-2023301211321311-1131310210303222-3022201123003223"></a>

<a id="canonical-1301122131111121-3033033231233232-2110203110112132-0302022103001200-0222212311100103-1221130231210333-3312012013232310-0011030133303101"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Optional.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232): complete subsection reference.

<a id="canonical-3103003112002332-3010112013033031-2213210133023300-1223031322303030-2032012020130210-3221130312213030-3111010012323013-1023202312321113"></a>

<a id="canonical-1210013120021332-2310221011103010-0310333201233213-1332012122113013-0203131011303231-0213032321111130-2101302121100230-3100223130023313"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-0111010011022133-3123323033111131-0230120131203231-1233010222002033-1203123113023311-3111202132303132-2201301302301110-1330133332110123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3102033011333210-0311103333120023-0030100000101321-0102331030113201-0113233231232310-1212033103322112-3321121222123011-3130101220111121"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-2220000320223113-0302112200020230-3121020032101303-1122203032333332-1211301032133233-0101312312101122-3220021212201311-1123323001210111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001013132003023-3232021101232202-2320013211011133-1133120333101230-0220102022133323-3303133101001012-3330321002011220-2223111233232000"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item`

<a id="canonical-1313330222022100-3312122223233133-2323300210201222-0113021130333313-3333221322230110-1233231113002100-2213201021311000-1313032000133200"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1331100200123200-3031033231133032-3121032231302023-3100202200021330-0331332112321133-1000122203013232-0120233101033010-0112320112003010"></a>

<a id="canonical-2123313122123013-1310320131031301-1031331031202121-0120112013111331-1001033303220313-1321023233022013-3220031011003213-3131203300111300"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0120313202221230-3131012213211000-1303232203300113-2020313300110121-1211211332033013-1203202221023021-2111130112113102-2113102011213210"></a>

<a id="canonical-0233311202011130-1021103322302132-1121033111220210-1010022112101022-0303031020001013-3220220231133000-0003132221230323-3120031002112023"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-0123312320102111-3213033002332023-0232203123102221-1001130102221023-2113023301230111-1233313013310122-3100113001110231-2301001312200003"></a>

Type: `"object"`. single nested block, Optional.

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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

Terraform syntax:

```terraform
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231010331002021-1220201223030103-2022123033023120-3333333201322020-3300330032201230-2331303322002223-2021212003323113-1123222322100211"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules`

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033): complete subsection reference.

<a id="canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-3210011231220003-3320121223332000-3302121222323300-0320131302113210-3110323330132221-1022123010110213-2231120123020032-0312230313123121"></a>

Type: `"object"`. list nested block, Optional.

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112113103223303-1312021011130332-0033130010200032-2022110023223322-2131200131110002-1312200311000000-0301330313010123-1031233122031320"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`

- [any_domain](resources--http_loadbalancer--reference--group-008.md#canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131): complete subsection reference.

- [any_url](resources--http_loadbalancer--reference--group-008.md#canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-008.md#canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-008.md#canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201): complete subsection reference.

<a id="canonical-1330133330101211-2113321332210330-3321303312302210-1232030312033201-2232010301013322-0113322231123010-3030011121031102-3033123021132323"></a>

<a id="canonical-3022332121031311-1032012031333310-0110030033001320-1130220333322203-1301020010211332-1110300221310001-1030222310122211-2332130212001211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222): complete subsection reference.

<a id="canonical-2113330011220321-2003310231110000-1003000031333203-3112102001233311-3101310013212322-1202202111330133-2100210231230203-3121300021220302"></a>

<a id="canonical-3000310023303321-3221030022033313-2033133002331201-2100330100210102-1220331110021213-1103301012311131-0230200001212002-0211330203200312"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.specific_domain` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-0133111030321200-1213122130100333-1131321000131330-0302332132012332-3031101323310320-3313201212300100-0312003030132330-0220221113211122"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-3321200132022202-3131231010032230-2313331022202330-3022022123133200-3012220230012121-1021110331022320-3213211131023332-0102013322030322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_url = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-3230030233302333-0213010233312000-3001233322002120-2300212233212310-2030212203233213-2031102103222111-3302311120213111-1210233320330212"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

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

Terraform syntax:

```terraform
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203233203012211-3011302202030310-2103013122122000-2121103200332313-1212013132032001-0223313322131212-2333133313033213-1312230302331100"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint`

<a id="canonical-3032121031003022-1102221111012322-3300231132001003-0030111303103100-0313000002311112-1221312101332331-0133120123110300-3212010330210220"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1321111203032120-1312202221100022-3213210112233010-2011323011111133-0023133312210003-0313313203221030-1202222223312003-1020103030113301"></a>

<a id="canonical-0133132313113331-0202303001313012-0203020221002333-2013232212202220-3212312301023300-2012212333212312-3001033220030133-3230311220211210"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-3203002232121121-2223200201110121-2001233012033212-2020101010301032-3300220121121300-2213303003122112-3213031311010012-2030333301313220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030032233132000-2030012021313311-1100202120320200-0020032223200013-0201223311130123-3203210321113310-2301202333101233-3310021000302330"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups`

<a id="canonical-0013112121233013-0112330220303020-3203300332333320-2200322123022103-3022020202232100-0010302121022320-2120133221101230-2023021332323203"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups.api_groups` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0313112301223312-3013231202200230-0302022303222131-3031221132223213-1212220001301313-2000303310231210-0122313233131012-0202023200303113"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303212231110200-0010323032012301-2300130213103202-1030002212230123-1232110023132213-1212211210320312-3313011130230223-0001232101133301"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-008.md#canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-008.md#canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-008.md#canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-008.md#canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-009.md#canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-009.md#canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-009.md#canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220): complete subsection reference.

<a id="canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2122303312133030-0311200301201012-3300202320103110-2210120103232031-0303330202101111-3331031110002031-1201023021100231-1023321331032203"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-0013300312001301-2302301331003330-2003110011133013-0233332312032230-1020131121002112-3330223122203030-2133010310003221-2111230001210210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-2113330312123003-3202321033220022-0203323130222003-1130311221332331-2122220011212103-3131121310313121-1023030200133032-2002023001232311"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202332221133222-2100220103202131-1111310102233231-3122012220302101-0300120012100133-2011313101123320-2000232001013012-0032221232212121"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list`

<a id="canonical-0110002022202321-2110101332320121-1031320102021331-3222131023013212-1013102310301023-0000132030111101-2200211133123010-3333022120102032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-0312220303220013-2030123102021332-1031220202130320-1113122000123121-3002123002301302-0332003313212222-1210022010023023-1010331112011121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303122322002010-3212303200212122-1322111231032311-2202313112201133-1301211033222011-1111201331311211-0330303300132012-2231111213113013"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-008.md#canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201): complete subsection reference.

<a id="canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2031332202110200-3210321322222121-3333101320000332-0202301113311222-0122101033121021-2003321212022200-3011123003130123-1003330331302200"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220331311302233-3133322033301020-3312111200332030-1313333210010130-0213303301221311-3120100233310230-0133232230201103-1103011303002302"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3023303322122121-0333313023033203-3132311331023311-1233303331331330-2132201202012001-1322002201002321-2002232031311103-0303323021202132"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1013023200132221-0021010311121333-2230310103112030-2303100021333021-1023103021330301-2321132112130312-2213222113333200-1130030220312111"></a>

<a id="canonical-3120201212323010-2122130132022021-0102132112030013-0223111111211310-2303333220202023-0120123332220103-0031120312102320-1111321223212221"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213120113311302-0110231210120211-3331233033112000-0313133022232032-1010332310232023-3301221003203003-1310233232300310-3101032320101333"></a>

<a id="canonical-2313002122303133-3111110313011310-1132311212103031-1013012213013020-2233013120221210-0320333330103312-3311100330011301-0100031112321031"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0030330110113131-3032130230302102-3111202232212021-1210122021112110-1000233001210112-0323122322230020-3022121322202212-2200210022301100"></a>

<a id="canonical-0032111212210213-0121320230301200-3323132333211303-2001300232310012-3303222113022320-3322301331300033-1211032230301323-2201131010311100"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0133122021221021-3120313110121001-0223200022312211-0222212310200233-2110220112311130-1210110213302301-2330003030133122-3000112033011111"></a>

<a id="canonical-0023102131231331-2003213021000331-0220103222233030-2000233023110003-3322021220120031-2322023221230023-1123312320003032-1222312333020032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-1332002202232331-1012330000010133-3310103332023221-2032020323332301-3123130000121223-3003323021233020-3223133321302001-0100331112233102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112211322222100-2223213203331120-1022032332313123-2003102311130310-1130131203222210-2200133132100223-0001102032013213-2331230111130020"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector`

<a id="canonical-3021100201132312-0122311133130223-3011323113010323-0020000303210231-0120120131122322-3103210301000133-0130021032113213-3233020132312302"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-3333122010313100-3120120131332123-2233321323230230-2030230130130012-0022331122132223-2022133302313030-0331121330202121-3022230201213113"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102101011201203-0021011330300233-1133231032010012-1211331330302233-2112022232013223-2323030133032333-1230323022033022-0102021123012201"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`

<a id="canonical-1002231312030322-3311131333200230-1020110032102021-2330233333130032-3023210123120323-0332011131213231-3231131010331311-2211021220320233"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-008.md#canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111): complete subsection reference.

<a id="canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2013323000223021-2120210302002010-1033120102210120-0213011210100302-0120313323210331-2323200233033110-0332303002331332-2032332033101221"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321211201211121-1300002233110111-0133003100120232-2013322021131021-3100003220133022-0012323100223320-1132221333300203-0113330112023133"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-1031113320111020-2022132222013330-3313011013300330-1321022310302211-2102212230120122-3311201202330231-3010220222231233-2322100200132321"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202220001130022-0202022200100031-0330312213030123-1213221231311033-1233001223133302-3330100323121131-1332120211212320-1232223103200212"></a>
