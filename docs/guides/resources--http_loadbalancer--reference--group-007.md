---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2012110031022323-1000013013122301-2310313132021330-2210102301222322-2131123123033320-3302322232221202-0302312023003223-0012211203230121"></a>

## name property — cookie_matchers / 311033320113 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2101332331002000-3230010013200100-0112003122210113-1221311003012311-2213020300210322-2110021212000313-1332013030233021-0322103103000323"></a>

## Next pages — cookie_matchers / 311033320113 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-007.md#canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011032333000010-2110010223133203-0102021211233332-0213101231232310-0030221201211112-3301123000002213-3021201311331021-1003332013031112"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 032022100321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0223031032222220-2011222120210323-0313210222321312-3332020121102012-3023233220310222-3333231323122322-3031110121133130-0032123332230133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-1301211311110030-0201103322320321-2322123032210231-0310030003031223-1111101122121131-1000220202113033-0221011312012101-3003231303012133"></a>

## Direct properties — check_not_present / 032022100321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221032011231332-1001322333233102-1113030312322332-0113022331211330-1323301320021212-0122303032001001-2331310132011010-1220300312021131"></a>

## Next pages — check_not_present / 032022100321 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123113320102020-3032111203121301-3202312023312232-3220303033201223-2011030113122012-0001111230122023-1013232321322210-3230300130332203"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present — check_present / 100030200100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0221221121110211-0011320032123021-3132211231112112-1031110122320121-2302012101123220-2031032020302320-0032120323300121-1110020102332120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-2213332321331323-1003203111101102-2213133020213303-3132002133222301-0031313233122002-3320032301200232-1103031301012231-3201001022021103"></a>

## Direct properties — check_present / 100030200100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231133001130310-0100222300131332-2300322333300102-2212231133030033-0112030222331121-2123320300222230-2311022102210101-0132002122132313"></a>

## Next pages — check_present / 100030200100 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1012000023300311-2100310300202101-0330200133301120-2132232320103013-2222300100123211-2232201001100122-1233230020202012-1012030213200011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213103232011222-0022332000102330-0120032000211003-0113331123000202-3330222032132310-0310302300233131-2303222212011000-3300221323001312"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item — item / 200210113210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-3323300233200130-3222030201211321-2102131320212131-2203103333102010-1213303132000312-3132301002333323-3130123102102012-1122222003121301"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-3302213221000303-3002001133003121-3222323312231013-0210131010333122-1032322000322021-1032130203132033-2201001113113333-0100030123021332"></a>

## Direct properties — item / 200210113210 / 3

<a id="canonical-1021212103122213-2110101122232131-1201331031002232-3210213002003300-2220212033131120-1133212213010032-1033032323322001-0203211012020313"></a>

<a id="canonical-1332222213233103-3202310121301312-2033102231100323-1310221222013030-2000322232220133-3233310311011123-0002302003212101-1021210023210012"></a>

## exact_values property — item / 200210113210 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-0130310310022322-0312321331113202-3112300130321311-1100122300220232-3201231332322113-2001102301211233-2322233020121011-3000122121033013"></a>

<a id="canonical-1131012313001003-2320002221131320-3231321222303021-2312221013302012-2010102322330010-3231201320203101-1233032011000131-3001033013011121"></a>

## regex_values property — item / 200210113210 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3231101030222201-0003100232211013-2132030230202233-3111120130101223-0113103210130100-1331210113001201-2122302122332323-3302230331202303"></a>

<a id="canonical-1330301031312313-3221000312111313-1131002313133230-3323001012202230-0312223233021200-1023213230202030-3330311112113030-3102132223232011"></a>

## transformers property — item / 200210113210 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-1033211222222222-0233232302312012-3100011231131132-3000212011020031-2203033003012022-2011223221200110-0202312121021033-0200032310321330"></a>

## Next pages — item / 200210113210 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211031310010032-2130333323210311-3010331113031231-3211203030320113-1202211232323113-0201302313022322-3232111132100011-0312203333311032"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers — headers / 133022222231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-1112032022303031-1310321332323132-3301000302203120-0220003021312212-3112123202311302-2320130023332300-0123211232001320-1203213232220003"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

<a id="canonical-2113003113023300-2333121211303002-2121303322202103-2213101220112000-0333102030120022-1311303121102033-1212320023130211-2031221011330013"></a>

## Direct properties — headers / 133022222231 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220): complete subsection reference.

<a id="canonical-0003311303201101-3001301131210220-0113133231023113-0100230323012101-3202102331123210-3131130302323002-0121012211100310-3203331201110020"></a>

<a id="canonical-2102103113310011-0220023030312103-1233331122133332-3222131220031311-3120031311222113-3221221321130110-2311121030132022-0201303022001120"></a>

## invert_matcher property — headers / 133022222231 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023): complete subsection reference.

<a id="canonical-0211313013102023-0210200002230113-2132112032011012-3200111020100310-2311202323323100-1003110023012231-1111100200313033-1020303131112123"></a>

<a id="canonical-2303332302201313-2013200021331320-3203113113103110-2231300230022230-2322201310302103-0121031333133001-2331012131202200-1021003301330113"></a>

## name property — headers / 133022222231 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1111203122110121-1302322201030211-2111131200121101-2012230031003010-0121202321010313-3113202132030303-2201313112330331-1332013020112123"></a>

## Next pages — headers / 133022222231 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-007.md#canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2110322001203333-3210020311020010-3132330023321212-0020011103201020-2012321031321121-3333022200211103-0001113303232322-2131012131203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102230101221100-2120101221001303-1022233333130231-2211232023201030-1220030302011232-2013322133323032-1112311103013331-2313222102302000"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present — check_not_present / 101132313321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0002133200213133-2001303203103011-3121130030231223-0101333322221100-2210330112200201-1220020122131101-3311021202011110-1103333230223202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-2000100101011221-1131331033102323-3110200203022012-1101221200322130-1312110202212232-3330333331100331-2132122300330112-2230122311322012"></a>

## Direct properties — check_not_present / 101132313321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313330320123000-0110001121011303-1033113300013103-1030121003210100-1131230311113332-0312213121200323-0032000332011223-0301011323113120"></a>

## Next pages — check_not_present / 101132313321 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0131011301333221-1213032033122212-3133122030003223-3330230002221201-0010033011321103-0222021111101122-1333312230221231-1320212320131220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121201320322313-3313033100021210-2003003101110233-1220210312333101-1030332202323312-2300003010010110-0033220131320222-3312000002222022"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present — check_present / 213032131301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3103033330003201-2033103203322000-0110123113003200-1110120011121023-3322310311023002-1221122230002003-2031021332022101-0320223210203032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-3203130120110022-2003320010222003-3032001120212320-3331310131333200-0120122023210203-3130032122332330-3302100211200220-2213012130000201"></a>

## Direct properties — check_present / 213032131301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132333323230113-0213201123211122-0023113320022010-2022210300211230-3123212023122013-3110102020201012-1013232133300212-2011100020011001"></a>

## Next pages — check_present / 213032131301 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1110320202331212-3101201330223333-2012013200101223-0212203002330323-0320122032313113-2301211211232203-0202030211022332-0311030010103023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332213103223311-1130212010212112-3103111002311032-1133113201211233-2330130223122032-2223300202023303-1330321032223133-3322311303133301"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.item — item / 011110022311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-1310102202200201-0012131313200012-2310132112111013-2313203023312123-2012111100002032-2230122032011323-3221101013332230-3023120021133201"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2023113111020323-3002012212210320-0300213120111233-2312033212311020-2202122211312301-2111101332310100-3332330112230301-2333231001232320"></a>

## Direct properties — item / 011110022311 / 3

<a id="canonical-2133210111223122-0211131002310221-0100231003132032-1121203031210100-1002012323121002-2200101203201021-3202022203132200-1311233011122213"></a>

<a id="canonical-2030011030201203-3213103310201312-2320311011220113-3232003200201033-3113020033313010-1133322221131021-3310000221320302-3013020011122233"></a>

## exact_values property — item / 011110022311 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1331210103033311-0231111001123001-0010123322232320-2100333310010120-0111213212313100-3022131011233301-3113212323213110-3122111033333110"></a>

<a id="canonical-1232321210231113-0012221311120321-2000202200121021-3022213313000111-0323230020232102-2222000030000003-0311222010110203-3223201030333331"></a>

## regex_values property — item / 011110022311 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1331332311312310-0212023323302022-1103300200320031-3203022133313333-3103111102232130-0010123322311122-1030100302332030-0123231323232303"></a>

<a id="canonical-3302030331302202-1211212122010131-1120113313033321-0031310303031322-0203022110100310-1011110302011203-1021030133323302-2333222302032212"></a>

## transformers property — item / 011110022311 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3320330233133330-1332132212113230-0310133020130213-0031210200310310-2232010331303212-0232030110202123-0101302130132113-0320101002211330"></a>

## Next pages — item / 011110022311 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102303103120111-1033221103122221-0110023031300311-2111300010311303-1102310331312003-1210002120321220-3002213232122302-3112320320301033"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims — jwt_claims / 213311310102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0033131120213101-2021101213122313-1112301101110023-0020303203303030-2330012203010213-0231302000220110-0012000120012123-0022210112113322"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330200333130200-2232333110333122-3302202120023122-1131320000120122-2231031021103310-2220202222322202-2012023101012221-0211020213301323"></a>

## Direct properties — jwt_claims / 213311310102 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133): complete subsection reference.

<a id="canonical-2122231301331010-1320312333312003-1333001210032332-3133321021031021-1121310322232221-1102212011201031-1210332132002020-2122103210333102"></a>

<a id="canonical-0023321122032003-3132221333203213-0232302330321203-1233132332300332-2312011033023311-2101130213212313-3001232032203010-1201301001122002"></a>

## invert_matcher property — jwt_claims / 213311310102 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310): complete subsection reference.

<a id="canonical-3132211312321102-1113230313203121-0100021113222112-1033010200002002-3110030330231230-2012320131232123-1333002110200220-1322013302223211"></a>

<a id="canonical-1232101002312201-3320013130100022-0021130032220023-0030003231013113-1030011032313330-1231132103232232-1312231121301003-3220233212131020"></a>

## name property — jwt_claims / 213311310102 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2033233102322303-2200123200132023-3312322031012221-2002012310233011-2013303301032221-3032230003312223-3110032030333012-3322030223022302"></a>

## Next pages — jwt_claims / 213311310102 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-007.md#canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2202013113121212-0312302301312321-2203331102213021-3321211020110121-1301132230300211-1031011303212011-1100112222210020-3132032331330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203012301210101-3121311102320220-2130321212322312-0311123002123002-0311133222132323-1300222303133311-3300233301132112-0011023033220310"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 031002330120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2213212220111130-2210120211112303-1130310210031133-2121103010330113-0332100101232323-2000333022101030-0313022222222000-0333312321011210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-2212020312000223-0222332320131232-0201202210211121-3021130121323112-0101322013031222-3133123130221100-1032311321223103-1221101013113133"></a>

## Direct properties — check_not_present / 031002330120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303111021230132-2222210313302312-3212123300101022-2312022221312003-0301313232013113-1223103032233130-2332101020032012-3203101133230203"></a>

## Next pages — check_not_present / 031002330120 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2012002332011020-0333231311122122-2012110033333113-0120300002132332-0203321002111313-2130200102000022-0200023110302321-2323001020122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102020211011330-3211201322302033-2112301013022003-0021330333012130-2031202231122302-1110203323320101-2312133321323200-2132213133321123"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present — check_present / 120202233013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312021030100012-0322330333213110-0001221031311302-2110023221303221-1121020023310031-3001112312023122-2321222121023113-0331010320331021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-3233020020311012-1212023200002023-2023212123203311-0113323220313031-0301132131132010-1300003301123323-0132302303213213-3232132102101322"></a>

## Direct properties — check_present / 120202233013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002021310012332-3031232332322202-0313013232213113-0032311211200223-1321303230213001-1302002321011200-1233200100332313-0302100211132221"></a>

## Next pages — check_present / 120202233013 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2120111123131100-2123300020302112-2331201211201300-3223211132210032-3100021311223202-2012020210031032-1230012021112223-2200322103311310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330122212003101-1332121032223100-1111011213000222-1210012220203032-1123023300301200-1213333233002013-3113123311120321-2013221123212110"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item — item / 210030031213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-0111033122123022-1130321103113202-0223200230000033-2013111311033021-3323213123011032-2033001202310023-2111203212310123-3210111313320030"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-0231213313201313-2230303030122302-3020100022310111-3023220032020331-1023300332331001-1230020333033233-1003231312221220-2130221203202023"></a>

## Direct properties — item / 210030031213 / 3

<a id="canonical-2312202120212310-1202200320110121-2020332330333023-3011320030031232-1320013020130022-0102132033232012-0130032203311322-3232121101130111"></a>

<a id="canonical-2202111113000123-2100132322100112-3230220232300311-2321123031333331-0021111011023023-3213031302202101-0023212032332233-2112233001203331"></a>

## exact_values property — item / 210030031213 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3203213223302010-0032123323231331-3102323133330002-2112103131130121-2321312120231221-1020303310122111-0200311213203031-1323013020200323"></a>

<a id="canonical-1223220223321121-0123201121233100-1230222313032011-2212230112033132-3230132233131301-3131003301203010-0332030210330232-1120323202100200"></a>

## regex_values property — item / 210030031213 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1111203301223120-2132332121032212-1221022221300332-1330302310322310-1020122230012110-2333032321233302-2223233203022200-3010213001123030"></a>

<a id="canonical-1313101122112222-3211301320121310-1322323222113121-2220131003323302-0001232133101120-1301230023102312-2210130131133320-1320032021112112"></a>

## transformers property — item / 210030031213 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0122001303032033-3313101331203310-3301221112312232-2200002313201010-2111303222131133-1310021123321211-3331103321223033-1003131112313202"></a>

## Next pages — item / 210030031213 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112332003110330-0033023331230301-1100310133233211-2321032233131102-2133200232032312-1031012311300222-2303102133303020-0113221031121001"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params — query_params / 030310312212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params

<a id="canonical-0320330213122122-1313120301101122-0000330133300220-1332310200101313-1303200103110131-2102031203002021-1330111302122001-3102101320132001"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301122131111121-3033033231233232-2110203110112132-0302022103001200-0222212311100103-1221130231210333-3312012013232310-0011030133303101"></a>

## Direct properties — query_params / 030310312212 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102): complete subsection reference.

<a id="canonical-3021330220001013-1331033222021121-3102121302132200-0112203312110112-2221102232021211-2023301211321311-1131310210303222-3022201123003223"></a>

<a id="canonical-1210013120021332-2310221011103010-0310333201233213-1332012122113013-0203131011303231-0213032321111130-2101302121100230-3100223130023313"></a>

## invert_matcher property — query_params / 030310312212 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232): complete subsection reference.

<a id="canonical-3103003112002332-3010112013033031-2213210133023300-1223031322303030-2032012020130210-3221130312213030-3111010012323013-1023202312321113"></a>

<a id="canonical-2132120030203313-3202221133132030-0133210211331203-0221303101322031-2111213331203321-3322131022323101-0022232310220220-0132122020020100"></a>

## key property — query_params / 030310312212 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-2112231310331310-0201310320112212-2112332121121220-2013232030321111-2032012122101301-1120233233301010-0220212031202013-2222003333123033"></a>

## Next pages — query_params / 030310312212 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-007.md#canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2000312303011131-1310130332222303-0331322333203013-1230121012023223-3311022133012103-2121021032023033-3132320321213311-3011310032330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011220131213323-1201312322122203-1312230202103033-3111023221322120-0032103330331332-0333211200200320-3320033100210120-2032310312102023"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present — check_not_present / 123013300131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-0111010011022133-3123323033111131-0230120131203231-1233010222002033-1203123113023311-3111202132303132-2201301302301110-1330133332110123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-3320002012131212-0100222022322031-3213012232032220-1302230322022222-3030000211002313-1132103100113233-0310210110012001-1321020223211123"></a>

## Direct properties — check_not_present / 123013300131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012222221012020-3111320233232130-3220132302220332-2202020000220213-0102003311012131-0300311303111311-2001033013231233-2023202233302102"></a>

## Next pages — check_not_present / 123013300131 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2200220323213322-2000010132222101-2221223131230120-0012101310002203-2111131300000130-1311012332213010-1021120223102320-2220102221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000233112311002-1200002313302033-1323302121013012-3233221133031313-1231122001301101-3120000211100130-1302023213102311-3012112301101132"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present — check_present / 302312303223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3102033011333210-0311103333120023-0030100000101321-0102331030113201-0113233231232310-1212033103322112-3321121222123011-3130101220111121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-2102312312212133-3202330221110132-3103112102111312-2320002100002301-3231222123003130-2323201333020020-0013311110032211-1130323023031233"></a>

## Direct properties — check_present / 302312303223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232113232121222-2101321203233103-2033121131100103-0220230020013020-1310123020311032-2012113212200300-3212113302311133-3221203002011220"></a>

## Next pages — check_present / 302312303223 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0302201210130133-3213120002023132-1131210033000211-1302122002210033-2333012132111202-2011200200332122-3002022131222223-0231101230312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001013132003023-3232021101232202-2320013211011133-1133120333101230-0220102022133323-3303133101001012-3330321002011220-2223111233232000"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.query_params.item — item / 002322120222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-2220000320223113-0302112200020230-3121020032101303-1122203032333332-1211301032133233-0101312312101122-3220021212201311-1123323001210111"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2123313122123013-1310320131031301-1031331031202121-0120112013111331-1001033303220313-1321023233022013-3220031011003213-3131203300111300"></a>

## Direct properties — item / 002322120222 / 3

<a id="canonical-1313330222022100-3312122223233133-2323300210201222-0113021130333313-3333221322230110-1233231113002100-2213201021311000-1313032000133200"></a>

<a id="canonical-0233311202011130-1021103322302132-1121033111220210-1010022112101022-0303031020001013-3220220231133000-0003132221230323-3120031002112023"></a>

## exact_values property — item / 002322120222 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1331100200123200-3031033231133032-3121032231302023-3100202200021330-0331332112321133-1000122203013232-0120233101033010-0112320112003010"></a>

<a id="canonical-3311230133210321-2033100222100131-3033333012123213-0130223200321231-1101132113123320-3112001100100013-0231331112131022-2023200102302110"></a>

## regex_values property — item / 002322120222 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-0120313202221230-3131012213211000-1303232203300113-2020313300110121-1211211332033013-1203202221023021-2111130112113102-2113102011213210"></a>

<a id="canonical-1011211321021030-1203310132033233-1203112202031321-3233130002201113-0000313303102312-3000122100021003-3033032332331210-2110120031230102"></a>

## transformers property — item / 002322120222 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0233233121311301-1102312130311100-1133211230233313-0001123200002030-2322111201132032-2211223302330012-0201121220230313-0021231003330301"></a>

## Next pages — item / 002322120222 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231010331002021-1220201223030103-2022123033023120-3333333201322020-3300330032201230-2331303322002223-2021212003323113-1123222322100211"></a>

## api_rate_limit.bypass_rate_limiting_rules — bypass_rate_limiting_rules / 110203120001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-0123312320102111-3213033002332023-0232203123102221-1001130102221023-2113023301230111-1233313013310122-3100113001110231-2301001312200003"></a>

Type: `"object"`. single nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

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

<a id="canonical-3110032313120330-2002122101031302-1001000322221113-0111100210120332-3230011301130100-3102303203233020-3223203223011223-2303323121033210"></a>

## Direct properties — bypass_rate_limiting_rules / 110203120001 / 3

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033): complete subsection reference.

<a id="canonical-2322113113213231-1300331033022002-1300302110211332-2200101013222221-0233121111302032-3031030131301301-3023001210002010-1311111133031031"></a>

## Next pages — bypass_rate_limiting_rules / 110203120001 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112113103223303-1312021011130332-0033130010200032-2022110023223322-2131200131110002-1312200311000000-0301330313010123-1031233122031320"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules — bypass_rate_limiting_rules / 010022321023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-3210011231220003-3320121223332000-3302121222323300-0320131302113210-3110323330132221-1022123010110213-2231120123020032-0312230313123121"></a>

Type: `"object"`. list nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_endpoint"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_groups"),
  validators.ConflictingListObjectAttributes("any_url",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_groups"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_groups",
    "base_path")}
```

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

<a id="canonical-3022332121031311-1032012031333310-0110030033001320-1130220333322203-1301020010211332-1110300221310001-1030222310122211-2332130212001211"></a>

## Direct properties — bypass_rate_limiting_rules / 010022321023 / 3

- [any_domain](resources--http_loadbalancer--reference--group-007.md#canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131): complete subsection reference.

- [any_url](resources--http_loadbalancer--reference--group-007.md#canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-007.md#canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-007.md#canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201): complete subsection reference.

<a id="canonical-1330133330101211-2113321332210330-3321303312302210-1232030312033201-2232010301013322-0113322231123010-3030011121031102-3033123021132323"></a>

<a id="canonical-3000310023303321-3221030022033313-2033133002331201-2100330100210102-1220331110021213-1103301012311131-0230200001212002-0211330203200312"></a>

## base_path property — bypass_rate_limiting_rules / 010022321023 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222): complete subsection reference.

<a id="canonical-2113330011220321-2003310231110000-1003000031333203-3112102001233311-3101310013212322-1202202111330133-2100210231230203-3121300021220302"></a>

<a id="canonical-1201101121101011-1300322120323222-0310022300002302-1311202032233001-2230323030301232-3323310201020310-0203300012122230-0102120111223100"></a>

## specific_domain property — bypass_rate_limiting_rules / 010022321023 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-2021313212331202-1303320020223031-0022030010122213-0200010310210231-3301122002100132-2000101301000301-3023322123333331-0101301202003221"></a>

## Next pages — bypass_rate_limiting_rules / 010022321023 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](resources--http_loadbalancer--reference--group-007.md#canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](resources--http_loadbalancer--reference--group-007.md#canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](resources--http_loadbalancer--reference--group-007.md#canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](resources--http_loadbalancer--reference--group-007.md#canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0132120010232023-2120333332011212-3022310321203131-1202230000000123-2210313112230121-0000313103000112-0013311222103313-0033021121121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311313123323323-1122110213230112-2202003211101213-0202232012310200-1302000332112111-1122021021231022-2222111213200310-2302130212322232"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain — any_domain / 311301132222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-0133111030321200-1213122130100333-1131321000131330-0302332132012332-3031101323310320-3313201212300100-0312003030132330-0220221113211122"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-1322021012121012-0233320102320022-1200220102100212-0220003223022001-1032012030103102-0130103130322212-0300123232220213-3103031331321200"></a>

## Direct properties — any_domain / 311301132222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120020311323133-0213233331002312-0012032202300233-2122302011113210-3020022003311113-3031001111212113-2302130010300213-1323032203011122"></a>

## Next pages — any_domain / 311301132222 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103121101322112-2100332012000200-0012003012122131-0033010133230322-3021013003201133-3020103323322013-1201301211323101-0220113202302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032310323310322-3111232101320232-3120002212100113-0231032133012203-2211301010121201-0100213022033312-0113103120320032-1301130022310313"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url — any_url / 121010020010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-3321200132022202-3131231010032230-2313331022202330-3022022123133200-3012220230012121-1021110331022320-3213211131023332-0102013322030322"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-2102311113312202-0020211022301032-1301132220310003-1002320122123221-0213322331312210-0231302123001323-0213213323130032-2131010022203212"></a>

## Direct properties — any_url / 121010020010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102122123031013-1131001211002003-2212220322222220-2331212322300320-2003133322201212-0033030010330103-0003022123120220-2023031132303210"></a>

## Next pages — any_url / 121010020010 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2110031310101110-1303103222112222-1320122301213130-0011200231101121-2302201210300100-2322110233030333-2120311323010233-3231113023031103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203233203012211-3011302202030310-2103013122122000-2121103200332313-1212013132032001-0223313322131212-2333133313033213-1312230302331100"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint — api_endpoint / 102120020012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-3230030233302333-0213010233312000-3001233322002120-2300212233212310-2030212203233213-2031102103222111-3302311120213111-1210233320330212"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
```

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

<a id="canonical-0133132313113331-0202303001313012-0203020221002333-2013232212202220-3212312301023300-2012212333212312-3001033220030133-3230311220211210"></a>

## Direct properties — api_endpoint / 102120020012 / 3

<a id="canonical-3032121031003022-1102221111012322-3300231132001003-0030111303103100-0313000002311112-1221312101332331-0133120123110300-3212010330210220"></a>

<a id="canonical-1313003322113322-2232111111300301-1222320031013220-3000321023202213-0133331002103303-0323211300130333-2003320230332113-0213100023121133"></a>

## methods property — api_endpoint / 102120020012 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1321111203032120-1312202221100022-3213210112233010-2011323011111133-0023133312210003-0313313203221030-1202222223312003-1020103030113301"></a>

<a id="canonical-2111103311010210-3222121312031331-3133110320022002-1022101013110201-2200310201320002-3131330303312033-0202113331032203-0011212121022201"></a>

## path property — api_endpoint / 102120020012 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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

<a id="canonical-2013032011200312-1100313310221123-1322022022330313-0120121130322101-0113320212231232-0132320302203103-0301301102311102-2020131300311101"></a>

## Next pages — api_endpoint / 102120020012 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200113020100230-1001333030320020-2031211032320122-1230230223121200-3021001330030102-3221120020020100-3200001032100131-3330002000333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030032233132000-2030012021313311-1100202120320200-0020032223200013-0201223311130123-3203210321113310-2301202333101233-3310021000302330"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups — api_groups / 000313022232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-3203002232121121-2223200201110121-2001233012033212-2020101010301032-3300220121121300-2213303003122112-3213031311010012-2030333301313220"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
```

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

<a id="canonical-1130103023130131-2123211133120323-1201202101001211-3233313331333310-3231310122223323-2200221103010303-0333311231103031-1010032102002000"></a>

## Direct properties — api_groups / 000313022232 / 3

<a id="canonical-0013112121233013-0112330220303020-3203300332333320-2200322123022103-3022020202232100-0010302121022320-2120133221101230-2023021332323203"></a>

<a id="canonical-0121203122002112-3323100000233123-3322200020233202-2123313100303233-2331013122032021-3023011333030330-1213112000013220-2232102223211132"></a>

## api_groups property — api_groups / 000313022232 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-3123323203210133-1213323020011022-1120133012023201-3222332031313301-0023230223131020-0320202130111111-2112321020020100-2130332312113002"></a>

## Next pages — api_groups / 000313022232 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303212231110200-0010323032012301-2300130213103202-1030002212230123-1232110023132213-1212211210320312-3313011130230223-0001232101133301"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher — client_matcher / 303212210000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0313112301223312-3013231202200230-0302022303222131-3031221132223213-1212220001301313-2000303310231210-0122313233131012-0202023200303113"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

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

<a id="canonical-0013300201300123-0330023132120103-3013232122020211-0203331131020023-0032301131131022-1202131212012111-3233231320133122-3200113322022232"></a>

## Direct properties — client_matcher / 303212210000 / 3

- [any_client](resources--http_loadbalancer--reference--group-007.md#canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-007.md#canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-007.md#canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-007.md#canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-007.md#canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-007.md#canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220): complete subsection reference.

<a id="canonical-1210203231213121-3123312231220220-2000010321103222-2332122121300301-0110100230302123-0020312303223331-0132203233331122-1331000103231302"></a>

## Next pages — client_matcher / 303212210000 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-007.md#canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-007.md#canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-007.md#canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-007.md#canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-007.md#canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-007.md#canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2333010110323301-3130300032202001-3020212211033302-3200322232000223-0223330222020103-0322211201320330-0220120330030111-1303123033210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233121231022303-0021223200202100-0022131103202303-2132222211103100-3332322101220331-2330011112322133-0221002232112333-1321002211002312"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client — any_client / 203320000200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2122303312133030-0311200301201012-3300202320103110-2210120103232031-0303330202101111-3331031110002031-1201023021100231-1023321331032203"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-3223313310123010-2200212130323311-0303023001312231-1312111233013003-3203323202302112-0232132213123022-1202203032102023-0001321322332211"></a>

## Direct properties — any_client / 203320000200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133023013223322-3311022222131202-3112021312023210-0011031000212112-3103013012320131-3200022231223223-2113103021002023-2002020100032333"></a>

## Next pages — any_client / 203320000200 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1313032013200331-1303110023313222-1201330133202132-1111003201020323-3220033221302323-2122331002111013-0332321330110212-3122301001302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232320320012133-2320001301310131-3003031030031301-3210301301033012-3203321001213002-0322221132031110-3033000102220020-3020303103312202"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip — any_ip / 311202303101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-0013300312001301-2302301331003330-2003110011133013-0233332312032230-1020131121002112-3330223122203030-2133010310003221-2111230001210210"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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

<a id="canonical-0120020212301002-1203002013220201-1230200201033313-0122223003322110-2130000312122101-1220212101233221-2231021332112311-0033110211331010"></a>

## Direct properties — any_ip / 311202303101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321011013201012-3032211330311112-2133031231301311-1113011232021332-1103010200303302-3303213331301121-3221030003211313-2222020012231001"></a>

## Next pages — any_ip / 311202303101 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3103323302220132-1003100003223021-0223121213302210-3222013302313130-0200313320322222-0003120333232123-0201323021310312-3320113320330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202332221133222-2100220103202131-1111310102233231-3122012220302101-0300120012100133-2011313101123320-2000232001013012-0032221232212121"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list — asn_list / 021330111322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-2113330312123003-3202321033220022-0203323130222003-1130311221332331-2122220011212103-3131121310313121-1023030200133032-2002023001232311"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
```

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

<a id="canonical-2132123000101300-0333012120013232-3231333111220012-1002123321310233-2100132320022222-1322021302301102-3011203130221012-3333120123311113"></a>

## Direct properties — asn_list / 021330111322 / 3

<a id="canonical-0110002022202321-2110101332320121-1031320102021331-3222131023013212-1013102310301023-0000132030111101-2200211133123010-3333022120102032"></a>

<a id="canonical-2033210213230223-0131302232220120-0232110213202110-1300202213330103-1022130222110021-1230230130222003-0111022023220030-3233323111203232"></a>

## as_numbers property — asn_list / 021330111322 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-1103202330013202-3332222030132101-2100030001203012-0222121202301330-2201321203111131-0300321202302223-1223320231123013-1220212311111302"></a>

## Next pages — asn_list / 021330111322 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303122322002010-3212303200212122-1322111231032311-2202313112201133-1301211033222011-1111201331311211-0330303300132012-2231111213113013"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher — asn_matcher / 002012203102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-0312220303220013-2030123102021332-1031220202130320-1113122000123121-3002123002301302-0332003313212222-1210022010023023-1010331112011121"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
```

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

<a id="canonical-2300033301301030-0123020021022002-0101021022213201-3011123330302012-3111111022103220-3102112010021020-2000030002322123-1102103210333203"></a>

## Direct properties — asn_matcher / 002012203102 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-007.md#canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201): complete subsection reference.

<a id="canonical-1020321130120301-1132230210220023-2112221021220023-0130011023032101-1032120201030301-1030213213120211-2221023111310302-2130301330111302"></a>

## Next pages — asn_matcher / 002012203102 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-007.md#canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0300233232013012-1313310113122010-1002110223320210-1201121002333210-0012131002021002-2023121313300101-0133300103102101-3013120320132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220331311302233-3133322033301020-3312111200332030-1313333210010130-0213303301221311-3120100233310230-0133232230201103-1103011303002302"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 010132312003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2031332202110200-3210321322222121-3333101320000332-0202301113311222-0122101033121021-2003321212022200-3011123003130123-1003330331302200"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120201212323010-2122130132022021-0102132112030013-0223111111211310-2303333220202023-0120123332220103-0031120312102320-1111321223212221"></a>

## Direct properties — asn_sets / 010132312003 / 3

<a id="canonical-3023303322122121-0333313023033203-3132311331023311-1233303331331330-2132201202012001-1322002201002321-2002232031311103-0303323021202132"></a>

<a id="canonical-2313002122303133-3111110313011310-1132311212103031-1013012213013020-2233013120221210-0320333330103312-3311100330011301-0100031112321031"></a>

## kind property — asn_sets / 010132312003 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1013023200132221-0021010311121333-2230310103112030-2303100021333021-1023103021330301-2321132112130312-2213222113333200-1130030220312111"></a>

<a id="canonical-0032111212210213-0121320230301200-3323132333211303-2001300232310012-3303222113022320-3322301331300033-1211032230301323-2201131010311100"></a>

## name property — asn_sets / 010132312003 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0213120113311302-0110231210120211-3331233033112000-0313133022232032-1010332310232023-3301221003203003-1310233232300310-3101032320101333"></a>

<a id="canonical-0023102131231331-2003213021000331-0220103222233030-2000233023110003-3322021220120031-2322023221230023-1123312320003032-1222312333020032"></a>

## namespace property — asn_sets / 010132312003 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-0030330110113131-3032130230302102-3111202232212021-1210122021112110-1000233001210112-0323122322230020-3022121322202212-2200210022301100"></a>

<a id="canonical-1000032000110023-3220233200001032-1301110300230133-1012001321303220-0103201331121312-0010000002311301-2232130321012031-2021332331113001"></a>

## tenant property — asn_sets / 010132312003 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0133122021221021-3120313110121001-0223200022312211-0222212310200233-2110220112311130-1210110213302301-2330003030133122-3000112033011111"></a>

<a id="canonical-0323111020002131-3001310111120231-2332211321113002-1311002112231103-3231032032201103-3200031313332231-2211313223021211-2003001301221011"></a>

## uid property — asn_sets / 010132312003 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0223302003120203-2020312302313233-0010233013333222-0222313113133310-0320230021312230-2213103010302022-0102202113223212-3023020210100331"></a>

## Next pages — asn_sets / 010132312003 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2310212103013113-1100303220012111-1332102213102101-2312202131121300-1230013103032011-3320023310303221-2002113131100121-0013210210133211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3113312003231312-0131331320123120-2112122033223023-3200031103311113-1301033132002210-3220200030300103-2132303312313212-0212003033102003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112211322222100-2223213203331120-1022032332313123-2003102311130310-1130131203222210-2200133132100223-0001102032013213-2331230111130020"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector — client_selector / 032131122332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-1332002202232331-1012330000010133-3310103332023221-2032020323332301-3123130000121223-3003323021233020-3223133321302001-0100331112233102"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
```

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

<a id="canonical-3003002110023230-1210111132212010-2032130000203231-2322031102300312-0220102100333223-3200221303120212-0320112110122320-0203221121013023"></a>

## Direct properties — client_selector / 032131122332 / 3

<a id="canonical-3021100201132312-0122311133130223-3011323113010323-0020000303210231-0120120131122322-3103210301000133-0130021032113213-3233020132312302"></a>

<a id="canonical-2323310333203313-3320320002120003-2030121303020130-0211232002032223-3112010020130001-0300003200301000-2230230230320332-3100011310032220"></a>

## expressions property — client_selector / 032131122332 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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

<a id="canonical-1023233131300303-0133133130310233-2020113122122330-1032321032000300-1213023021333331-1333211302021201-0101110100223313-0032333003002000"></a>

## Next pages — client_selector / 032131122332 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102101011201203-0021011330300233-1133231032010012-1211331330302233-2112022232013223-2323030133032333-1230323022033022-0102021123012201"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher — ip_matcher / 031112230020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-3333122010313100-3120120131332123-2233321323230230-2030230130130012-0022331122132223-2022133302313030-0331121330202121-3022230201213113"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
```

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

<a id="canonical-3121002321021322-3012220031220013-2222311211031330-3200201030232222-3321232130222302-3320331002010203-2120212331021301-2231223302313110"></a>

## Direct properties — ip_matcher / 031112230020 / 3

<a id="canonical-1002231312030322-3311131333200230-1020110032102021-2330233333130032-3023210123120323-0332011131213231-3231131010331311-2211021220320233"></a>

<a id="canonical-3112213333031211-1110110222021310-3203300131112233-0233232023001023-0131210203012110-0303320323212132-0131300201221322-3310113220301010"></a>

## invert_matcher property — ip_matcher / 031112230020 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-007.md#canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111): complete subsection reference.

<a id="canonical-2011233110123310-2000103103312230-1122021332031132-2200331103233101-1031333310232321-2112311033300300-2031232313301012-3301123122313133"></a>

## Next pages — ip_matcher / 031112230020 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-007.md#canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3313202030031003-1123103301111002-0002111323113100-1022023213020212-3231321131131320-2222210011100203-2313232332132003-2111232003313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321211201211121-1300002233110111-0133003100120232-2013322021131021-3100003220133022-0012323100223320-1132221333300203-0113330112023133"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 121031301103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2013323000223021-2120210302002010-1033120102210120-0213011210100302-0120313323210331-2323200233033110-0332303002331332-2032332033101221"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323011220101201-2112220022310331-3302030110032002-2002132123233002-0233310301113111-3010133021232011-2332220213011233-3121021331210132"></a>

## Direct properties — prefix_sets / 121031301103 / 3

<a id="canonical-1031113320111020-2022132222013330-3313011013300330-1321022310302211-2102212230120122-3311201202330231-3010220222231233-2322100200132321"></a>

<a id="canonical-3112010330210302-0313221023132223-0130122033101200-1023202120002003-3001200102023302-1020001013312130-2212130000210300-0321030333030230"></a>

## kind property — prefix_sets / 121031301103 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2202220001130022-0202022200100031-0330312213030123-1213221231311033-1233001223133302-3330100323121131-1332120211212320-1232223103200212"></a>

<a id="canonical-3221011223013323-1021220001211231-0102301031002233-0210032201211301-0331110100310321-1312132311301131-3321222211310013-0120212310310121"></a>

## name property — prefix_sets / 121031301103 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1122231112133130-1003001300221001-2030232220002331-3022221121332101-2010211020000313-2132230300310022-0233230310113333-0010223020330003"></a>

<a id="canonical-0103023121313130-1013013023203100-1223303103303020-1102320332313233-1310110232112222-2111202310232122-0212221200220132-1032003323001130"></a>

## namespace property — prefix_sets / 121031301103 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-2032322312011002-1201210300012222-2302230020210011-3033313312000131-0010100331233002-1001003331203032-0000020000202013-0033111002122122"></a>

<a id="canonical-0111023111113213-0100000203300202-3020012322100002-3103311123201030-0132323221220120-3221010313203201-2011032000033302-0112302002330001"></a>

## tenant property — prefix_sets / 121031301103 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1102001221022131-0132213131000330-2213121230102311-3003220222030020-2122112010010301-3101211133133231-0230303101321030-2211232002330200"></a>

<a id="canonical-3123232121123200-2231210310300303-3201211333222010-0010210123320121-3201211012001333-1331011222303213-2112321213223200-0212030122321100"></a>

## uid property — prefix_sets / 121031301103 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0300312032200103-2133130330021203-1130213320330302-0321001312022100-1333202213103330-0320000023112031-1321213322132120-2111300100322123"></a>

## Next pages — prefix_sets / 121031301103 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2112210201211101-0212311022213221-1022100111213003-0100111233033032-2110021210200322-3102121001111103-1331231323211323-2223320033022133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200113333332102-3331022332002031-1120103010102301-2121331130312211-2212212010121333-2200210300321022-1003200210000330-3113200103222123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103313100230223-3102111002100312-1310003033221001-1233101010022222-0210223103322010-0000303200201320-2111113321233321-0321231112332023"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list — ip_prefix_list / 332232222213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-0202312213312000-3203203322023010-0011230313132331-3112023010323212-2113132222003010-1320121310320300-1030031333320102-1303123120302212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330222020103032-0002120022221100-0312033323032200-2320321010120221-3223322121033301-1330031313003222-2111220233112220-1323121013313222"></a>

## Direct properties — ip_prefix_list / 332232222213 / 3

<a id="canonical-2202103202101002-3320000021333022-2213333103201221-3121120321212102-2303312203113213-3000322320103323-3231223121021133-3020303201332212"></a>

<a id="canonical-3110113303321000-2313231210203112-2113220023022321-1232211033031212-2233210101202101-1021002031112111-1010030102332003-0213033010211003"></a>

## invert_match property — ip_prefix_list / 332232222213 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-2213330313001113-0111112013131331-1111020311100030-0031030013011202-3321333333210302-2220222121121211-2330122131303212-0111201011311203"></a>

<a id="canonical-0021331311101113-2133132101212312-2213110100103322-1333222302003031-3211332230332230-2102221002031233-2302300102303330-2232002211021133"></a>

## ip_prefixes property — ip_prefix_list / 332232222213 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-3310122321300222-0021120030001310-0113012130212133-2130012132023332-0220232333221111-3213223233111011-3130211321022301-1100201032220230"></a>

## Next pages — ip_prefix_list / 332232222213 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3312120300100220-2330131103032312-0202010330213232-3020322331303310-0212132301011221-1033003031230310-3321101123123011-1122312120111220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000223310121233-1213232012311322-1101323223120323-2233310223102301-2222002313123210-0122133303312030-3320232102201202-1231331103201311"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 212210103203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-0302332011311321-0221033331302301-0221021130101110-0033201112323320-3023013313331103-0333333322032211-3022120332322131-2232331331212000"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
```

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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311010120200033-0120101130211001-0323121330332302-3101231003200210-1230031220313021-3012102121023302-2131103220100013-1310310002212011"></a>

## Direct properties — ip_threat_category_list / 212210103203 / 3

<a id="canonical-1133220333300300-1012001201111020-2000203021233221-3313322102331220-0221021303111203-2200331323010032-0201132211231201-2230121010020211"></a>

<a id="canonical-1100211203120211-3031232102131003-2322211231202133-2312332132012030-1303132322102022-0220330022222231-3302021301020320-1020331200013130"></a>

## ip_threat_categories property — ip_threat_category_list / 212210103203 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-0023133300012122-0113031232102332-2211001301203221-3323113212320122-1020112212133233-0331312101230213-2011021233211332-2002020100133201"></a>

## Next pages — ip_threat_category_list / 212210103203 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3031203222211302-2303123322301120-3302130021100031-1000101233131220-3101201020012210-1301302133100302-2200220030010101-2113133013012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330210233110210-1330020321223101-0100100312100112-1230203030022200-1212110001100133-3300132320133032-1120022033011222-2131312311213222"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 011013122112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0031120021232233-0233123101022221-2112113211301033-3201302022332302-1121122132102213-0113031020213320-2001333022103103-0303230330013103"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301321001010223-3030211210030312-3220321322310322-1102110231123023-0020302202303020-3120113011133310-0011322000030012-0232023330021000"></a>

## Direct properties — tls_fingerprint_matcher / 011013122112 / 3

<a id="canonical-0320213132131201-3131223031311303-3112212312100200-3122111213112330-3231331003231131-2102213231132323-3002000223020203-0011230233121120"></a>

<a id="canonical-1221320033033021-0331210312201213-0300313033223323-3130330231201321-2030102320133210-0332022321103103-0330201210333023-0320131311121001"></a>

## classes property — tls_fingerprint_matcher / 011013122112 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1021101222131230-3302130312002103-0332220303232302-3000331102323210-1203301333212300-2012121110002130-1030023011120132-1230323230100013"></a>

<a id="canonical-1130121332330111-0213012000220202-3301211103021320-3011100213032133-3310320213202020-0233333130000031-0101123023012322-1020301113102121"></a>

## exact_values property — tls_fingerprint_matcher / 011013122112 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3022201232220002-1131320312131203-3022011332132331-0003213332110130-2121230200223320-1033211003011222-0100300132220302-3201333201323100"></a>

<a id="canonical-2033321321313223-1233112220033103-0011101203131200-3203310201121212-3213233310032210-1123320102023111-0011010111311103-1211003110033210"></a>

## excluded_values property — tls_fingerprint_matcher / 011013122112 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-3011301301102221-2012113002220113-0231233221103220-3330013203023321-3100120112313131-0233020001332322-0222003111031333-3123303222320033"></a>

## Next pages — tls_fingerprint_matcher / 011013122112 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3113033230331123-2320113030220232-0121002220222021-3011011031030000-0212023201333231-2002022010203012-0330010321102110-2022123020332223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203232312002020-0320310221331103-1010212312221211-0313130112222123-1110133222003130-2120103231333123-2311010002200133-0202110001023112"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher — request_matcher / 102110330322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-1332010230112030-1013220102023001-2131300200110331-2020002313312011-1022021101231311-1232101230032222-0203330321332031-1331230121003211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233023010232133-2003200130022301-0103101323233231-1011102202220321-1000021210302230-3221012122331012-1001112322000220-0122111002313221"></a>

## Direct properties — request_matcher / 102110330322 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011): complete subsection reference.

<a id="canonical-0032210321213321-1111213020222011-2010101000202233-3303232232033230-3331221211201220-3221103313000033-0133232321010023-0313321122321200"></a>

## Next pages — request_matcher / 102110330322 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022221320003013-2311013202102213-1230300000220033-2101213022033001-0211220233030103-2302002220312123-1300221123010313-3213122101100020"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers — cookie_matchers / 301131130322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-3131223312031330-2131301312303203-2103203112011103-0103130213121212-3020203110122103-2210323330120212-0023211113003303-0201233200102301"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122230213320030-3002112313031232-1231110332211330-0010021213211301-1112313012122330-0210032020011320-2033203102312301-1003112003222012"></a>

## Direct properties — cookie_matchers / 301131130322 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232): complete subsection reference.

<a id="canonical-0330123202222313-1212123103203130-3032023303022212-2101001233132022-2201130322210202-1012311100003211-0112211011310330-0320121322230202"></a>

<a id="canonical-1133110203110023-1002333300222303-0000103323103321-2303133120102200-0032020333330300-2013111120331003-2123302032030133-0210112110223230"></a>

## invert_matcher property — cookie_matchers / 301131130322 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231): complete subsection reference.

<a id="canonical-3022320002033011-1210210311211011-0220012000200201-1011302021021110-0231202230001111-2133313211011232-1213222033020031-0003331330010222"></a>

<a id="canonical-1011301001023203-0111233101011321-2012001000330133-2201020031302320-3032231301122310-1303133222133320-1030203303133233-1100111310003222"></a>

## name property — cookie_matchers / 301131130322 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0031220331033002-2032330212120211-3132212313221030-0211000112331120-0020122030110031-0231030020223003-2320102212310103-2230200302211133"></a>

## Next pages — cookie_matchers / 301131130322 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-007.md#canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0332001211220202-2222302130003203-0200000303223110-2320212201333033-2003221000321000-1111332320123003-2223120102320032-2021301300331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010030122232100-0021022333202302-2212201032130110-1133010111232023-2113232233200303-3020221321322101-0211232223131330-0130201221201010"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 330303021133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2031121112020331-3302223131021322-0023321220220321-2100131103122313-3013232300101001-3122332122022333-1303103223202323-1113221312333232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-2020203023133210-0002330230231130-2032311101103131-1203030012032121-1221110011032312-3032222212033032-2103203113211103-3033031201333201"></a>

## Direct properties — check_not_present / 330303021133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200220211023020-1130302322033133-0301000120322032-2303201301111022-1303332001233301-3032303333300231-3130323230312130-3302133101102130"></a>

## Next pages — check_not_present / 330303021133 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013311121031001-0323310112313032-0200311210032213-2003222203010331-2122132302232200-3003203030221220-2310303333330010-2333120002303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112313132322313-3100123030312230-0302313123221323-0033000231122030-1003301103320313-2113313302312011-2323303022100032-3123110211220321"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present — check_present / 100303001223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3101031320010310-3103310000233320-3120022022311012-0033322331002211-0203230132132112-2013333021123021-3123203000102100-1112000231102220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-0303213301130121-1311223200313112-2233301230130022-1022133102333023-3301112003000213-0223130101122321-0302032332130210-2131001003211103"></a>

## Direct properties — check_present / 100303001223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111232122221032-0102023103132023-2232223000131002-0203122133120331-1303001112212132-3323003222320120-3333122311312122-0101301013302210"></a>

## Next pages — check_present / 100303001223 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0310011312213322-1120200011320002-1111312033130310-3102020121101100-3002323002130121-1130010313020013-3130121101032131-0020313110132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233101322133310-2201300001132012-0322100202112002-2211202303303121-2102033002101013-1023203212211320-1013122023221032-0002203211232333"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item — item / 321302321110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-0322112323232031-0311132313001003-3031222022213013-0003021002023003-1002211230002032-0210320122023312-0000133323000221-3213200012022310"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-3301211220230310-2312302100001333-2201233122310231-0331232331211030-0002032002312013-3211220030322233-2303303032010131-0301133302112202"></a>

## Direct properties — item / 321302321110 / 3

<a id="canonical-0211131030230313-2002320012220223-0130113020030010-1131003320022312-2123221021030233-2201212122331212-0232013302132022-0311202303102220"></a>

<a id="canonical-0232123220112311-2202331321332013-3133200130322200-3210012103310131-0231212200202302-0033001200311101-3030233011333110-3211001010013303"></a>

## exact_values property — item / 321302321110 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2132310332120021-1202210100203312-0330302021221322-2213013303312323-1313012130213022-1333122012213122-2103023023313003-1123101311123303"></a>

<a id="canonical-3200011103133103-0102133323311113-1232032312100003-0333200110030211-0110023013320303-1010210121201313-3020320031121211-1101110213212001"></a>

## regex_values property — item / 321302321110 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3023203023101113-3133111313211022-1201101232023301-2321233120232012-1113010020233013-2012310001020120-3012120021033032-2031001221300332"></a>

<a id="canonical-2223303203233012-3302100121020320-1111301310201323-1330202001210213-1112210011112122-1230322213133310-3303002031032121-0332301101223023"></a>

## transformers property — item / 321302321110 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0110112030331022-3301122001333321-1230031200002113-2012113300310233-3003123100202333-0010213100010312-1301131133333312-1210123333033202"></a>

## Next pages — item / 321302321110 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3301222333113101-1111130001300320-3002000110212131-1002211101201232-1103310322231013-1132232110313103-1133100303230111-0213102222113032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101132012220211-2101003013321211-1202230210012120-1011131013103322-1220331000200311-2021203330003232-3030120003011132-1322121030103121"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers — headers / 331230313200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-0100022210022131-3130131232331323-2322121320013101-0303232103132330-3102201322120011-2303302320202210-3221223231202122-1230303101203233"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

<a id="canonical-2100211323031123-3013100003111123-3000123033331033-1010202102133310-2021202200113231-1022013133203231-0321233220023320-2211300332121201"></a>

## Direct properties — headers / 331230313200 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222): complete subsection reference.

<a id="canonical-0212000221221212-3311231333313022-3123321331310120-3203013101333330-2020123320123010-3001223333222022-0200010220010121-3211222000003312"></a>

<a id="canonical-2130223112330010-2001132020200133-1113120030101011-3213301301300300-0301211032032330-0210213213122303-3032202312130300-3331302303113212"></a>

## invert_matcher property — headers / 331230313200 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313): complete subsection reference.

<a id="canonical-0121100023210323-2122023002111211-0103102233031111-0213000222000332-0123232333031111-0022222223023222-3023010201032102-0301102000313003"></a>

<a id="canonical-3200201102112031-2303232221122322-2121122222023001-2333010230223210-2000110303133310-2112100333311200-1310322320211322-3123120003101133"></a>

## name property — headers / 331230313200 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3330221333110031-2101302300001222-3313033131133100-3010020213312111-0123301010303123-2210011330110030-2113200103122100-3131032212111003"></a>

## Next pages — headers / 331230313200 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-007.md#canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1122130132031003-3223101202211302-2201131001300001-1122323213110303-0030221023012002-3321022031222123-1213320211233011-1013123122303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330323101131212-0132332033101130-2132002022120013-0001311233013310-3123221302333230-3313233110210021-1333331033103333-1022123202130123"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present — check_not_present / 022132211331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-2110023000333300-2111303303020220-2012003311332112-0100323131021013-1233303112002022-0011312123233222-1312021310221330-1030302232110312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Upstream description:

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

<a id="canonical-1333232033320230-2201312221321002-1032132331333001-2021322223311211-2212103130020100-2312302101213033-0002030310022302-2100201210322112"></a>

## Direct properties — check_not_present / 022132211331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223101130013210-3120223202203322-2233303130033303-3231322223030003-0232133232121322-0001213323301321-2021130332200220-0301232330012320"></a>

## Next pages — check_not_present / 022132211331 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0320220010023331-2302202111322120-0120310011031302-0232113023120201-3001011012322033-0301313233320222-3103310230023331-0323010203001222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011321303202332-0023101203322122-1310213030021323-0232222222213221-2102131110003231-2333032011323102-3002103310113302-0010310213131221"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present — check_present / 022031203311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0230101211001311-1020011021210311-1201223130201220-0212020100010331-1312212131210120-1231233200303203-3322032220330012-0231101212023210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

Upstream description:

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

<a id="canonical-0031120202110230-3321202301011201-3332131012202011-3310022122003300-1112123211231002-1320020122211301-0010010313201230-1331132020000000"></a>

## Direct properties — check_present / 022031203311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133101123212131-2312002232213121-0001023030102033-2022233230312223-1303200020330331-0133122230103022-2031212201123321-1203033031200113"></a>

## Next pages — check_present / 022031203311 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2312310310233132-3312120113120301-2102131130101312-0113220232131123-0233110022113003-2010102233201332-3031202331220220-3013210133102313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
