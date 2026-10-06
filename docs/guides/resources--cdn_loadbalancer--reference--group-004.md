---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-0133201313223223-3313320132211210-1012110321232121-2221112020301202-0222330100213103-0322321103030222-2002232103223330-0132102231013312"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1130100100032021-1120123101300103-3211033012003000-0113312023013020-2302113123312331-0223232032020330-1302000313122121-3301133111312300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1302000012223331-2220013102000322-1121000010112110-3223100201002213-1200012020103303-1313310132200010-3012333010232111-3130120203101311): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1100100202232013-0232301222100320-2030012302021021-1331100130032220-1032330020223133-1030233121002322-1312230133302102-2220023200112222): complete subsection reference.

<a id="canonical-0130011310302223-0320220232010322-0022103110323033-0101203221101303-3131032122202121-1030122322200323-0111313111112003-1321020033312111"></a>

<a id="canonical-1021301010112022-2033113000200102-1001022313312320-1120332303020310-1130231131330101-3031011122033001-1220000310300021-1123300001133030"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1320212322133301-2203122232230300-1230320130033311-0021221020121000-2122332332333011-2233231112111312-2102210133012121-3202333103023123): complete subsection reference.

<a id="canonical-2322000022330311-1202033122133302-1021323031323322-3202233111232213-3200033312120331-1331130133112110-2220002322311133-0330100000022233"></a>

<a id="canonical-2210210032030033-3121310103231223-1121330022211202-2222102001033011-1213310110223032-0200022201300302-3300221331330131-0032032010020301"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

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

<a id="canonical-1302000012223331-2220013102000322-1121000010112110-3223100201002213-1200012020103303-1313310132200010-3012333010232111-3130120203101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-2320000301132221-2213321220021231-3211123302100132-0110120331030030-1101030121120322-3013231210200001-1231121212102000-0131211310320300"></a>

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

<a id="canonical-1100100202232013-0232301222100320-2030012302021021-1331100130032220-1032330020223133-1030233121002322-1312230133302102-2220023200112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-2131332231011200-2000022303033022-3131232120021222-0230021303101002-1020121022021010-3321021112111130-1020232010303310-1303203120013220"></a>

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

<a id="canonical-1320212322133301-2203122232230300-1230320130033311-0021221020121000-2122332332333011-2233231112111312-2102210133012121-3202333103023123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-0111012023032301-0113321022133311-2013302101202313-0203112200021100-2232332221301030-2122121210133301-2323033120331033-2333333013131021"></a>

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

<a id="canonical-2112021310003313-3032200303110302-3103231220323023-1001331113330203-3112212101233312-0123200020130233-3220012113010123-1030332003010021"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.headers.item`

<a id="canonical-3011321230010201-2111201100220311-2010210012023233-2330331002020111-2131223230302230-3001101200310031-3132202202113013-0212120133122332"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0023323330133020-3223011012331212-0032213130232032-2120020021320112-2131103111030221-2213023320133323-0011322030112122-0333312000230100"></a>

<a id="canonical-3231021011131030-2021331030103012-2232022301100001-1322312102230133-1330313022113101-1210310013330221-1131211310010023-2010332232312221"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1102210002313023-2113212002330032-0212101123113230-2101012020210203-1023130320000113-0102011311222202-0201031210212202-1003132330113003"></a>

<a id="canonical-0330021101101031-3303002003231212-0103022012320333-0300002320332101-1000300012211221-3011122101020000-1111302013210213-0311103102210201"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022222021122001-3002003010030312-1020011133230112-3013003301012123-0223333311032301-1332111121323103-1331231131301233-0000112001032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-1103202032222031-3001033000223031-3220121013222201-2131030000032202-3322203200132232-2122303223120120-2131002003110111-2031120021323101"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-2003221213323000-3120002102301130-2100222313311333-2133320013103030-3000130302312002-0232230103133123-3011202303312112-3231311023303313"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims`

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-2002100233213123-2212030303301300-0202113130323021-0313032010132303-0001321213013321-1320202230301131-1322200220222313-2011102021311301): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0211011322300031-2023320112021032-3010012120022121-2003123133301232-0313321313022300-0311000122013222-3311120021222103-2133001100312303): complete subsection reference.

<a id="canonical-2232231110210333-2131032233013220-3032232111031032-1130320332232201-2320312112000322-3322200223312030-2323310111210313-2030222033221320"></a>

<a id="canonical-2313201220013133-1230110312312021-1233222332211003-1010330122110032-0010010213311210-0330013012013203-1310333123130202-2121200020203121"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-0012322212301233-0102320232122212-2113330221011013-2031023110202130-2312110000320203-2131232233222013-0020011221013311-3121311131003113): complete subsection reference.

<a id="canonical-0031331231312001-1102202012000221-1110300031333202-0013003021311122-1032300331220000-2031111303133220-0023103203331332-3001100010112002"></a>

<a id="canonical-1311113113032231-3311303101101232-1220130012103102-1011211021320221-1130013213103100-3130320010320203-2130312022323311-0320022011231133"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

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

<a id="canonical-2002100233213123-2212030303301300-0202113130323021-0313032010132303-0001321213013321-1320202230301131-1322200220222313-2011102021311301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-2022222021122001-3002003010030312-1020011133230112-3013003301012123-0223333311032301-1332111121323103-1331231131301233-0000112001032221)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3010302233303010-3113133211030212-2222011231013101-0121002032313202-2113013010333313-3021112013111333-2033230030002201-1113301323201120"></a>

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

<a id="canonical-0211011322300031-2023320112021032-3010012120022121-2003123133301232-0313321313022300-0311000122013222-3311120021222103-2133001100312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-2022222021122001-3002003010030312-1020011133230112-3013003301012123-0223333311032301-1332111121323103-1331231131301233-0000112001032221)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-1130102001333112-0033011012233331-2202033330111220-2223211030100020-0133023133030000-3300022221020010-3011121003100201-2322331000302213"></a>

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

<a id="canonical-0012322212301233-0102320232122212-2113330221011013-2031023110202130-2312110000320203-2131232233222013-0020011221013311-3121311131003113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-2022222021122001-3002003010030312-1020011133230112-3013003301012123-0223333311032301-1332111121323103-1331231131301233-0000112001032221)
- api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-1332023113232223-1110023200022220-0133321003203000-2002331131130120-1030011031201301-0221030302320102-3131111310330233-3302122312313103"></a>

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

<a id="canonical-1023230333233203-3232101121210320-3110233322022121-0032333220020232-3201021200113323-1123220113031302-1100230322330021-3102032323000033"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item`

<a id="canonical-1121013021103233-0031030221112013-2033310223310213-3320133233021001-0323112013021103-1133120212232103-3330311302201231-1222110131303321"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0330113220200332-2231302332202333-3302322210303231-1231102332312020-1323022320313221-2220232022233321-1212101133020112-3301000233010231"></a>

<a id="canonical-2221030013120221-2032101232033303-1111023303000002-3113232121330232-1302000302203221-2102011100232011-0020211313203111-2232301000012331"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0120213231313302-0200223303323312-1331311130033012-3123312102332320-0021031032201322-3013221020313222-0032333211333313-1212320100032131"></a>

<a id="canonical-1112001031130223-0202013100023302-3323222003112122-2011120311200022-2110101131030022-3331221300031223-1322212120013213-1301033131330332"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params

<a id="canonical-0031021303111233-3001202133211120-1111310023301313-0022302032001103-2222231023222323-3100122021223120-3201120130332113-3130022232132021"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-2100020013131020-0322031121332231-0002201102111033-3122020032323333-3013311012203103-0211220132011021-0113203222330220-3022212223111210"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1312132302112301-1320100132311210-1133011333322220-3221002302202331-1303200031102123-1233122013110332-0321033102031201-0010202222211311): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1011002212102223-3233011111330113-2232102132122020-1202230301210102-1120031302000123-0130111002131103-0332232301200021-3122333132331322): complete subsection reference.

<a id="canonical-3301320211230322-3210033012101311-0111000301033111-3220312230012012-0113231321213003-3222331120322301-1102123121130030-0130311323232003"></a>

<a id="canonical-2332220032211012-3030002003123002-2021121102110001-0323022313220313-2213101110133101-3033022220221122-3231213222130313-1013230303133303"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-3013210210302021-2300300201220203-0100113023013310-2332112123001013-3331131322023222-3013303330121320-0210100311201123-2331130233313303): complete subsection reference.

<a id="canonical-0011311330010121-2102231220011312-0223121203012330-3111232201221123-1002022000020333-2131033022310110-1131203120211021-3113233023012003"></a>

<a id="canonical-3121212032320323-3102010200310320-1300010300213023-3011303300021121-0020101123231233-3033130102311003-0223302301333322-1010201222003102"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.key` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1312132302112301-1320100132311210-1133011333322220-3221002302202331-1303200031102123-1233122013110332-0321033102031201-0010202222211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-2102210010111302-1012012023332033-0332002220321000-1001200120323030-1323200010213323-2131320221021133-0210132102322132-0123323112022223"></a>

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

<a id="canonical-1011002212102223-3233011111330113-2232102132122020-1202230301210102-1120031302000123-0130111002131103-0332232301200021-3122333132331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3021102102330013-3120122202122031-1321332001021330-0111020200221200-1222111333111310-1002211203311303-3221322330323202-1312111220100123"></a>

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

<a id="canonical-3013210210302021-2300300201220203-0100113023013310-2332112123001013-3331131322023222-3013303330121320-0210100311201123-2331130233313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113)
- api_rate_limit.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-1230230330211210-1002002323130012-2103222210233033-3232113001121001-3203301213010032-1033001230011212-1121310322021021-3132122101113132"></a>

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

<a id="canonical-1030323103020032-0231100103031002-2232013203131103-0102333132303033-3332112130200111-3130313211202011-2112301321022131-1201230132200130"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item`

<a id="canonical-3202033132323000-0202201202333010-0220112100103303-0001020110000201-3213222232300112-0203130332221130-3133023212102330-3110103310123320"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2001232001313003-1001302132103011-1032113320033212-1013333330322322-2130213102103113-0123231011211011-0100220122001002-1030133011322021"></a>

<a id="canonical-0212330303133131-1112013112110112-0101230121023311-1031132313221033-0320310031332133-3321213321011023-1010131303233332-1011023211031320"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3133030331320313-0021013222022313-0000203323333223-3312113102232232-2020100302230020-1310222000310330-1201330322101323-2131210323301012"></a>

<a id="canonical-3311001313201320-1302020330131112-3123200331213223-2300330102201330-3212100332030020-0113012102133313-1113121022022210-1210210212232102"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-1321320021223310-2220311302230202-2012320121211121-1230313313200022-3313001313032332-2320222120211101-0010331023000133-2313001023030123"></a>

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

<a id="canonical-3013102210101312-2000332030101023-2003010132312213-0323112120121202-3022013210310302-3322130012131232-3323001211031103-0121300310332131"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules`

- [bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130): complete subsection reference.

<a id="canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-1001331312332113-0011300020312213-1112302310000121-2201312012012310-0130233231321012-0203021210221110-3100223022310011-0233322121321112"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-1101312103123313-1233133111101211-1302232331222212-3203102301230300-0311302111121300-1000203000122210-0232212223001320-1120300113232210"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-1201320111203120-3332223101112022-2012321103220100-2221223332233033-2010031201113313-2320230303033310-1331101222111303-1201330033323221): complete subsection reference.

- [any_url](resources--cdn_loadbalancer--reference--group-004.md#canonical-2201201000113012-2321221131010120-1120111132033020-1033012110031120-2231131022100311-0331001321033210-3233203000313113-3132222020213012): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-004.md#canonical-1220122231011231-2200231221123101-1301123301132322-0312031201012033-1030131101320211-2321000032101302-0201121312022022-2103332113230300): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-004.md#canonical-2113212020210033-3010221101301030-0003101203132101-2111230231331133-2002330113220123-0313120011022233-3031223010200303-2131300133100023): complete subsection reference.

<a id="canonical-2203232033301301-1310232213320032-1311123030012310-0113102211201022-3302332020313130-0120211120021021-0200131311001111-2321020302123300"></a>

<a id="canonical-0223103222103200-1331132310022331-0301211312112031-3001323303112211-0011022310331031-0230100331103103-1010023102012031-3131230233100110"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.base_path` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003): complete subsection reference.

<a id="canonical-3102200022003203-3032333321233132-2203133100300230-2322010133020013-1003233030332133-0332310313031010-0021133312313223-2000300212220332"></a>

<a id="canonical-1201332100111022-1210100231012312-2200223003331022-2010121031210102-3123000200010130-1110320330013212-3122013120223332-2030020332202212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.specific_domain` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-1201320111203120-3332223101112022-2012321103220100-2221223332233033-2010031201113313-2320230303033310-1331101222111303-1201330033323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-0123133031231102-0030112023210223-1203013320332020-3233033220103102-2101200010031021-3002232022233003-1103130300202102-3230030310322010"></a>

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

<a id="canonical-2201201000113012-2321221131010120-1120111132033020-1033012110031120-2231131022100311-0331001321033210-3233203000313113-3132222020213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-2310232231100210-1112212220020102-3330122102320131-1021121330203010-3200130123303223-1122003030232320-0200303001231101-3023003022233101"></a>

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

<a id="canonical-1220122231011231-2200231221123101-1301123301132322-0312031201012033-1030131101320211-2321000032101302-0201121312022022-2103332113230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-2310212331102203-3103332012002220-1013212220033031-2002222320103202-3213233011021210-2313201200311111-3201120112312112-2321030331200202"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

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

<a id="canonical-3130000110220212-0103232020022233-2030203320001321-0101322230310330-3030232102303231-1113203232132021-3033212203010220-1311232302302313"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint`

<a id="canonical-1223121031132322-2023223001121111-1332322322311300-0223032110321031-1121321312123112-0113302020130210-1300221233311232-3320333232200023"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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

<a id="canonical-1000300312111002-0311132220221020-3233022233333231-3213331030113211-0031200232220001-0231100203200323-3100003012101012-1130212301210331"></a>

<a id="canonical-2233120020303023-2310310321001331-1301011203002203-2000330123023203-1132220112112130-0101311331100233-2221222032023132-0102100332112031"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2113212020210033-3010221101301030-0003101203132101-2111230231331133-2002330113220123-0313120011022233-3031223010200303-2131300133100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-3112223002103303-2201313332223213-2133212323023131-1313322201211113-2010020203010132-1103302120110001-3221200323021111-0330311221323322"></a>

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

<a id="canonical-0233111003232233-0312302003202121-1211032212322100-0201211221033320-1223121220331212-1110121012322223-2312012222102200-0031021011102001"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups`

<a id="canonical-0023112100132021-1201100201031132-1202123000033323-0130331131112331-3333212213132203-1000020231231010-2303123123330111-3001300123021123"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups.api_groups` property

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

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

<a id="canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0020100232132112-3133012333033112-0220301131221011-3230103330301021-2111031131002221-1002321030221332-1020202131131332-1323220330130223"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

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

<a id="canonical-0203122303112331-3002021033330121-2003332320132133-2333002120000310-2321320133231113-3322223112102311-2021110312201302-2003113100333321"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher`

- [any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-2302331200031202-1210002300221010-2132213322310030-1321220130123132-1201020202232103-1210120201001211-0320002010102133-1031101133103133): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-1130320030100033-1211303132000120-0223231303301010-0033211233210222-2311131230013122-0131312331020230-3313303001233130-0312103021202330): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2110333231132032-3211332313121302-2013022022213000-0212200023123101-0103323232220020-3222100300300332-2301122303130302-1303012203121123): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-004.md#canonical-2120310313303211-3202011103102020-2010110230302233-0202320101300021-3023311312011032-1221331303122013-0200021332101210-2330330001202330): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-3201021030112331-0103312100311221-2330200122322220-3200133211322310-0303103102303303-0011220101003200-1123122120033131-2031113233010001): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2003112110333113-2231021110032122-0123132011102103-1223210331201310-2133130221021232-0212310113131300-1322010011002020-1201310103133011): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2112131213110233-0233201110033223-2013001111300332-3231011123323020-0000330233311011-2131331110201130-0102200023031130-0133223332110331): complete subsection reference.

<a id="canonical-2302331200031202-1210002300221010-2132213322310030-1321220130123132-1201020202232103-1210120201001211-0320002010102133-1031101133103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2110223122032030-1121303203030102-1001331220003330-0111331232033030-0002200110020232-1003100102100311-1211302312332002-0313012021330100"></a>

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

<a id="canonical-1130320030100033-1211303132000120-0223231303301010-0033211233210222-2311131230013122-0131312331020230-3313303001233130-0312103021202330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-0301303133112223-3133300212220223-1313110100002123-1200230220220133-1210100333232111-3302230223010210-0101033111003013-3123130223133310"></a>

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

<a id="canonical-2110333231132032-3211332313121302-2013022022213000-0212200023123101-0103323232220020-3222100300300332-2301122303130302-1303012203121123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-1121120333311012-2121122310320322-0203033132233211-1001201313022130-0001322301331203-0220001231321000-1321000113110130-2031113222320313"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0232303333113023-2321331002103212-3300121323101333-0212312333222003-2201033023013021-1300302003210111-2000102002212302-1112112123201202"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list`

<a id="canonical-3303330022202003-0321202022320220-3113021133203321-2200333101102203-0320202000213330-0203123211230123-1011020210222311-1301323331132213"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

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

<a id="canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-0100031103232123-1211020112133133-1031103311231203-0310133003020213-0013003303101003-0232302112232003-0201111011220200-2320231031021131"></a>

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

<a id="canonical-3300310312302331-0302221001210121-0203002233200231-2232013301112331-0203113321202213-2021230213320030-2322020122001323-3222012202022211"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher`

- [asn_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-0103022321200201-2110322302102221-3310133203120012-1122320300232311-1232131313320010-2210323021230110-3211130100332101-3330030212323320): complete subsection reference.

<a id="canonical-0103022321200201-2110322302102221-3310133203120012-1122320300232311-1232131313320010-2210323021230110-3211130100332101-3330030212323320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2312110231301001-0022003010110333-0103222000320230-0301200330301003-1031200011202211-1113032230212123-3202211023033223-0012003211112010"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0121130130022332-0333302131333200-0231222111123123-3122100010300132-0132020303030231-2200002200230322-2122101331123111-3010113232023211"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-2313110313231101-3000023003133212-1031221220202023-3022102133300303-0310113222033223-1302331100210300-0330030203222022-0320012321122132"></a>

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

<a id="canonical-3200120313003321-3320032320322210-2003030232000233-0101010212012101-1013120221232330-1020130013210000-1012132032130300-1302230003222313"></a>

<a id="canonical-0211021302000221-1301320111010112-0320130201221311-1133322023003023-1013120113200033-0010201330103101-3321012111100300-2232113230223303"></a>

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

<a id="canonical-1110033101030122-0032032011331020-1020200120133321-1200012213120100-1322101221113000-0133102023121313-0030113100220132-3102321102120103"></a>

<a id="canonical-1122103032123313-0322303121202113-3113210210211001-0130132112012333-2320100310222022-2021230010102233-2301112132031132-0002011200312002"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
  }
}
```

<a id="canonical-0002203330302320-3002121123032222-2033203123020030-2003101302320020-1011212033001003-2322222212120022-0120032102100101-1211030313333213"></a>

<a id="canonical-0121102302301323-0103212213222123-3311121201113210-2221330122131222-3011012222132301-1113213103211021-2022001020102203-2111300211302001"></a>

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

<a id="canonical-1020102223023003-1031110113201311-2032210310221022-1013310111331131-0200232302130212-1223132010311323-0121023100222101-2200230330012003"></a>

<a id="canonical-2300312221022130-0201303132103000-2312300303122130-2303331000233002-0323130012203330-1111231312212100-3332003330333300-2031303013022101"></a>

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

<a id="canonical-2120310313303211-3202011103102020-2010110230302233-0202320101300021-3023311312011032-1221331303122013-0200021332101210-2330330001202330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-1321311201121333-3303013200000223-1010332232002221-2012003213313331-3333121202023220-1310100200132202-1122012012331301-2003333032313212"></a>

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

<a id="canonical-1322011102022023-2311100121332232-0302100002012200-0113133111303000-3111312020130301-1030333200211123-2121232213110203-2213200010023203"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector`

<a id="canonical-1203320230110210-0011010033303210-2230330101310200-1332011231030222-1022111202110122-0322030210210020-1322221102301121-0212210223002110"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-3013033103112012-1230023202321211-3201111020100102-3021330101312200-2003331130213021-0023102321100320-1202322333130220-1331212122232312"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0001311312332132-1232000211201111-1312233110223000-2110002102332320-1313222001101331-0121211320322030-2003013132203231-1322220033223121"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher`

<a id="canonical-3111303003302000-3212321232211000-1232123303012201-0003222121323322-3312032131200022-3131130130222313-1221112322231203-0213011311310313"></a>

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-1233030103120332-3100201120113232-3300021103020120-1331012131312311-0223233130012210-1303312112102230-1020012002223002-1102313102113313): complete subsection reference.

<a id="canonical-1233030103120332-3100201120113232-3300021103020120-1331012131312311-0223233130012210-1303312112102230-1020012002223002-1102313102113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2013102030323000-2020011230213322-1202102121213220-1333011130012113-3020033121231320-0113100230110330-3110301321232303-3033001233033211"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3232311230223002-3020111031223121-3010111322001301-3222013120112212-1210303201023223-0000232132031310-0013310303231131-3330002032332003"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3020300132112321-0131003210301320-1232223133130200-3123032013212120-2223331330332210-0231333220222231-2023123123313023-3100220303210301"></a>

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

<a id="canonical-2213300300103001-2131113223333231-2120100213003220-1313113113103330-3323133331212333-2122122310200213-3020011200202001-0022013222123333"></a>

<a id="canonical-2102122212022102-3101001031000333-1320231211233320-0031101223132321-1133223310030311-0331221212010221-2211113312020102-2231030113012100"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-3121003303131200-3312221033301002-3233012021203100-0130332021032332-1212110102320123-3000010110022033-1322023021303332-3213012200333210"></a>

<a id="canonical-1113303112222330-1033310200123200-3333231030132332-0130313132230223-2110111111031011-3001130212203212-0111031233231211-0210122100300220"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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
  }
}
```

<a id="canonical-2210112102021320-2101032011100222-2311332330023012-2212230000100323-3321330201210010-2323203201002122-2001010111133313-2212202322132100"></a>

<a id="canonical-1322301023131320-1313231300001122-3220310302303302-0303220213211300-3120230231001222-1332020222121111-0223013310300220-0022011311022023"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-0322120002012012-1010123321112302-3323200313202322-0312222333121011-3211113202230122-1331311023123120-1012030113321013-0233102302211202"></a>

<a id="canonical-2000030303312130-3231002121133203-0030313220021310-3300022321033222-0022332111323323-2322323103211331-0330130221230021-0202132220030110"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-3201021030112331-0103312100311221-2330200122322220-3200133211322310-0303103102303303-0011220101003200-1123122120033131-2031113233010001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-3212320033030133-1113311210220320-2000211213100011-0222132313212220-3231130333101213-3100313320021123-3000130220331001-3301332102201102"></a>

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

<a id="canonical-1201103301200231-3312233201030213-3231233200023303-2133113000123331-1321023133230132-2302203121331010-0130312220223302-2001331231103101"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list`

<a id="canonical-0111231033312022-1110231030121023-2103200033100323-0333031203212011-0111201220121031-1312223003031110-1000222013301012-3211331223211310"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

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

<a id="canonical-0212001233102113-1002233101310312-2010130010301321-0130201130100130-3313132103202331-1230030033332112-3202020212122310-2213003020201123"></a>

<a id="canonical-3003012301220323-3322121112322131-0230200012320012-2232021221113023-2321211330023331-2330001121001213-0230332230102133-2202201130200033"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-2003112110333113-2231021110032122-0123132011102103-1223210331201310-2133130221021232-0212310113131300-1322010011002020-1201310103133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-1112310322100010-2123100003000230-1120020013231223-2322200132232301-1013010232120302-1121201212201000-3131233311322010-2322020012001201"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-0233312323111131-0000203131001111-2113100311303103-0131013230002120-3121213021102132-0112202121311010-3020100133212023-3301013033001000"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list`

<a id="canonical-2321013202310213-0123231203013323-3023230320220031-1020330013200310-3112311211033031-3320303010101131-2033301131021213-2133012113311101"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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

<a id="canonical-2112131213110233-0233201110033223-2013001111300332-3231011123323020-0000330233311011-2131331110201130-0102200023031130-0133223332110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0133002331210023-2111010111223100-3132212210021022-0223033332010122-3003112121012021-3132010331011032-3322011322331120-0013111002132302"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2300313222212303-2031001223310200-2211211302201021-2330311122221000-1121132101113212-1311020111220312-0133322210120220-3303113233220322"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-1010321200003212-3020231132210322-2312121013222310-1133021013020101-2120230311331100-3231233203200322-1003303113103312-2022130221032210"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

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

<a id="canonical-2110300001012331-2303213312011303-1202232031101231-2333312011200300-2010033002202200-3103023212220030-2212223233202111-1103221021233012"></a>

<a id="canonical-0020302313231302-2020303030210201-3210221111212103-0013221102100022-2231132201003310-1220300011123012-2332132030333223-0200030321320321"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1130213100332221-0221011111132003-2233100202031121-1022222233100303-3113003131302203-1220233203223032-0230213323300232-0302312110232000"></a>
