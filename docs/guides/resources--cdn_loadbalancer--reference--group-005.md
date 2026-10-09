---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3212030210213120-0221200210323221-1013300022102202-0223133000333031-1231302032110212-3121103301311030-0231122302113233-3011121321313212"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.name` property

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

<a id="canonical-3310132110221202-1130122133120123-1130112232332202-0220001000031311-3203001121311010-3320320110000201-1303111331123130-3230313303320023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2233033203230111-1001330232200020-0000200011212110-1320010300023222-2222312200031011-3011230333020120-2332033031102323-1131302000131021"></a>

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

<a id="canonical-3230210312313201-3213023022132221-3111100123210231-0111203113323011-3210101331232201-1120020303331233-1333022100201332-2222121033130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3133113331112333-1010123303023323-2312000303013001-3110233021113313-0101333102220321-0000023322121233-2001201332200302-3330111111313311"></a>

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

<a id="canonical-1200030320003030-3220331002321222-3321303320212011-3020133031313003-1332311200321313-3101032020110132-2203213112000321-0020023300001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-3001232121333231-2202213123123020-0001312000201312-3003003312230112-2311110112213211-0122203002331000-1312322011030232-1300332013203223"></a>

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

<a id="canonical-3131210303233232-1313133133010131-0322030321101220-1002100231111101-3221012102120201-0031123132202313-0310310311313002-1122201003310212"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item`

<a id="canonical-0123221132313211-2212001013120231-0310110331321031-3022020212333313-1033000133201110-0212233321130230-2013131333203211-0111230101201120"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.exact_values` property

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

<a id="canonical-3132223233332102-0313313300122230-3120313212311313-2101203321330120-1020130103303133-2223003312332232-1313300201221313-0123230013333302"></a>

<a id="canonical-1001213321013003-3220213221220011-1000030203011323-1010023133002332-1112101331120101-1120121320120102-0012232213322022-0213120313011331"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.regex_values` property

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

<a id="canonical-2313132033100032-0013013111023230-3330012333302322-0232213003230001-2132010010210200-1131130330033010-3302022311311202-3330312131211032"></a>

<a id="canonical-0203110023211320-2220010001300230-3022331321313201-1313031100002133-3310132011123111-2120111111200303-0133012301213001-0003222310122012"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.transformers` property

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

<a id="canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-2030211013311003-0223231211012002-3130222033312010-1132030121121011-3300332210031033-1122303330300221-2121211303133333-2233323303103212"></a>

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

<a id="canonical-1000102102220203-2011131232223222-1023223322331233-3303033013331302-1221202100102333-2322111002022111-3213232221113302-1330102001323231"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-3013330211330200-1332100111313312-3223301013331132-3300323101032322-0103112333233020-1331132121012333-1122211203320022-2133311030333323): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1332033201222120-0201131233311213-0313333331122123-2101110323101012-0130102222210233-1133132231102330-0231013213010211-2310330212100122): complete subsection reference.

<a id="canonical-3120212231021221-2000312312302110-1101303313332032-2321333021001300-3211330122123232-0233302031211121-1333123000230312-3110031323330220"></a>

<a id="canonical-1222031300112101-0032202311322103-2211232321231211-2003221113013213-0212230322312132-3222020113233001-1200023333200323-3310310321310030"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-1013101010232200-0100312233230123-3223332201121331-1301220113103123-3121303102203310-3300333110302200-1231223233113123-2030301232020332): complete subsection reference.

<a id="canonical-1003100203111231-2122202010011212-2311111313222011-3122302220002332-3210213121111213-2033313112102122-1032100123001231-3232230010032132"></a>

<a id="canonical-2220223200313331-0003013212320111-3101320000220223-2310313030002223-1001300321101231-3312331303203103-3103133101001002-2330211130302123"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.name` property

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

<a id="canonical-3013330211330200-1332100111313312-3223301013331132-3300323101032322-0103112333233020-1331132121012333-1122211203320022-2133311030333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-1112032023320000-0130121302301302-0323300221133010-0011121311030321-3120110100211000-1012011031223333-1022302203132203-2223231303202031"></a>

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

<a id="canonical-1332033201222120-0201131233311213-0313333331122123-2101110323101012-0130102222210233-1133132231102330-0231013213010211-2310330212100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0121332033002313-2302220300330132-2022200331031102-2231322003111330-3131203133012100-1020102132203333-0003033133001320-3301230010322300"></a>

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

<a id="canonical-1013101010232200-0100312233230123-3223332201121331-1301220113103123-3121303102203310-3300333110302200-1231223233113123-2030301232020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-005.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-3200113013101030-0122300101013323-1320022100030011-0132031300121030-0030133111002301-1000311211311000-1011311320102220-0112122101223322"></a>

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

<a id="canonical-2310330311110013-2003213101110321-1201222312030202-1131031303213022-2112302120333322-1113202023200120-0031013013220022-1211212131330300"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item`

<a id="canonical-2022221321001032-2310211023033112-2102131130310103-2120303130111202-3103102200122000-3013323222000211-2001120222322022-1001131313310231"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.exact_values` property

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

<a id="canonical-2321301030301031-2211003231100032-2230000112210301-2110132001100310-0011302110132210-2231302100020323-3001201313132330-2121331310311013"></a>

<a id="canonical-0331233302200203-0331321330313030-3233023223021220-3102101332003200-2213033210132033-3313230200030212-2302300210310003-1030311112133013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.regex_values` property

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

<a id="canonical-0120312213110122-3232233120221122-2032133002100002-2110223100210211-0201201320231022-2202203103013212-1132002112002312-1111122002323203"></a>

<a id="canonical-1332130130030220-1200331031003221-0200331001320100-2322233022011033-0010333121200031-0310230202321320-3121033311333110-0223112323223231"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.transformers` property

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

<a id="canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-3302313101313332-0333020133021212-3133012201102221-1123022113223210-1203132210101102-2223101310201120-0030222313300220-0322113103331100"></a>

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

<a id="canonical-0212022301133012-1212210331323003-2331130330011301-2301311131033023-1011121231202330-3132312121212131-3312122111320031-0120213333231332"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0102331231033001-0300002103030310-2023201023223020-2232020203031302-3232210013332121-3203321111003100-3203312110113331-2133210211021110): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0023021030200202-1012131123022010-3201200301003111-3310120133311312-0331101122100002-2301003032000133-1013003332112003-0321322031003012): complete subsection reference.

<a id="canonical-2020033230202133-2222203301002023-3313123113200130-0310133010032333-2221102213112213-2001320222011022-1313223332003113-2331122202222010"></a>

<a id="canonical-2330012103023131-1222130212011200-1030323133032213-0011212110323112-0132000023110231-1320002101212010-3323001210203130-1002011130312020"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-3330013301211313-2101332321331223-2100332033213202-1030122121020322-0231031123011100-0230232322210311-3303022030232331-1231211020022301): complete subsection reference.

<a id="canonical-0110103202003330-1003312120101033-0100332023200123-1330310001022312-3303201130300133-2133322012031111-2200002211213331-0111110232111001"></a>

<a id="canonical-0300330001123222-0333212311212233-0223112331133300-0311012113100221-0321030022102201-2201033111022130-0200320232332320-1021003223102011"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.name` property

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

<a id="canonical-0102331231033001-0300002103030310-2023201023223020-2232020203031302-3232210013332121-3203321111003100-3203312110113331-2133210211021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1022232301111311-3013102223111011-1001202103212222-2311113200223113-0120022011000010-3022311131030223-3210010102001333-1333000310102332"></a>

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

<a id="canonical-0023021030200202-1012131123022010-3201200301003111-3310120133311312-0331101122100002-2301003032000133-1013003332112003-0321322031003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-0231322232231100-0300001322332101-0003122003233031-0031020333312222-2220200312112031-0131010102321120-3000301332330203-1002112133003022"></a>

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

<a id="canonical-3330013301211313-2101332321331223-2100332033213202-1030122121020322-0231031123011100-0230232322210311-3303022030232331-1231211020022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-005.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3022302031223331-1001122021032332-0201011001022321-3112002303001302-2212012001321003-1011033012201232-3333201221001010-2011103202322032"></a>

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

<a id="canonical-0133102320002301-2313120213330220-0111001312220122-3200223130113120-0213303230312003-3311202210021330-0221223311131112-1021022311111013"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item`

<a id="canonical-0223231102231032-1303131213203002-3321311010000031-3200333023213000-3310333123232133-2212202320203020-2013111030123301-3312013000031332"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.exact_values` property

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

<a id="canonical-3221332231033013-2132032031031312-2311223220130021-0202303322230021-1221322103132323-1210113311000111-0222230020311032-3212112201001213"></a>

<a id="canonical-2302310332310130-2213120212100011-0301023332313112-3312010230332212-0103013322022223-2211303013303110-3132210311022010-0203000333320332"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.regex_values` property

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

<a id="canonical-0302203211131313-0301203332312320-2131200002300232-3311322233301323-0130013332013000-0001332112013313-3121312233310021-2020231020213203"></a>

<a id="canonical-1233111121232302-1210230021202311-1112012202003201-2022330233223122-1323100023001310-1101011003102020-3222003131120023-3300232132212101"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.transformers` property

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

<a id="canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-0310131303132120-0223020212201222-0100211023220010-1012022230331123-1333123121021222-0111020112032012-0112213012331213-3301101232312310"></a>

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

<a id="canonical-0211120333322311-1321300133331311-0001222301011302-3203113000120103-1020101231013033-3233330311112000-3202002122332320-2002003131330301"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-0330321131102030-0033002013122203-0231300323110020-0300231200221103-2110132213310113-3301002100223320-3000013121321011-0311221003032221): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-005.md#canonical-1133012211221310-1300322030210200-2110320231021302-0311110011322133-1200300132101123-2031120320031013-3212232032103023-0333000101101032): complete subsection reference.

<a id="canonical-1021010001200101-2133022101023222-1230220231010233-0133203131133133-2301301310002020-2102101302230033-3322311013331122-1200302232212221"></a>

<a id="canonical-3030003331233122-2232101113221121-3331302132233112-0202331033103301-3330320002301200-1032113321020021-0202013012231213-0021003110013031"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.invert_matcher` property

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

- [item](resources--cdn_loadbalancer--reference--group-005.md#canonical-1000202212012101-0103332322113231-2131321203103312-0103211110223131-3031212010322331-1203011223033103-2010010120200132-1232133231001110): complete subsection reference.

<a id="canonical-2230210103020003-0331123032212102-3021013323302123-2113130000303032-1011103103231210-0212111020100100-1232022213311132-1222123001113121"></a>

<a id="canonical-2221200030200333-2003113112023220-0310100011202330-0013112310212100-1033120120113130-0310023300201202-1001222221323133-1331320100031212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.key` property

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

<a id="canonical-0330321131102030-0033002013122203-0231300323110020-0300231200221103-2110132213310113-3301002100223320-3000013121321011-0311221003032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-0122330020113020-0320131011032230-0203203102230101-2002121233023130-0231032031013301-3000002222033323-1010313110133012-2013230210111311"></a>

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

<a id="canonical-1133012211221310-1300322030210200-2110320231021302-0311110011322133-1200300132101123-2031120320031013-3212232032103023-0333000101101032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-2210100102223012-0012030000033132-2302002133032332-3000102300131311-1120202233322003-3021130011033010-0213120332103321-1222132011012331"></a>

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

<a id="canonical-1000202212012101-0103332322113231-2131321203103312-0103211110223131-3031212010322331-1203011223033103-2010010120200132-1232133231001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-005.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-2201230101103102-3133120103301020-2231222023200132-3022323131311331-0010000222131301-0220023331121230-3233020302130020-1323200202030023"></a>

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

<a id="canonical-3312001233032211-1130203123020231-1211121113233331-1122100310202023-1333333201010012-0121331213100333-3301211120331300-2020111001312311"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item`

<a id="canonical-0332121001032213-1120101031332030-3001221211123201-3020112113310320-2201200321010110-2102111020122010-1210300333321012-3221131332313122"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-1003330012012102-3003123313210233-0021231313002213-0301031023333110-2232121121311101-2132112002300323-3311231203130131-0032023323201222"></a>

<a id="canonical-3012131230202133-1120202112020223-0030213233201210-2312000110213023-1011203112003323-2103333232302231-2130131123233030-0032031022222131"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-0111220311030032-3310031202010333-2331230210210011-2222310002030330-0233010122003123-0302322033022231-1213233013120021-0023313310220010"></a>

<a id="canonical-2110102310102200-1323321220113121-1221122023331011-0110300133110333-2221231001310033-0333122231320200-2022121111331111-0220100311002113"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.transformers` property

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

<a id="canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-0021013103213300-0232231032021332-0120122101130210-2130120003121012-0133222120112322-0100210000321331-2202300023221313-0132022203320221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101130301122132-2100313133112110-2202321301310103-1322230202233133-3132202110332100-0213321302110110-1031112230122030-0330011303232032"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-005.md#canonical-3131203222110020-2013332301110203-3003111101121002-1133332023223301-0233231233230211-0222101201000233-3120001000122223-3232112232212123): complete subsection reference.

<a id="canonical-3131203222110020-2013332301110203-3003111101121002-1133332023223301-0233231233230211-0222101201000233-3120001000122223-3232112232212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-3312131211310000-2322033231320103-0301201302313002-3321130013222222-3001202233220001-1031331203222012-0302033030102222-3331200121110000"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033101122132331-2022201130200130-3031230302113023-1032303203213033-2122330311131330-1300302331301111-2011221202000232-3310113203210111"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-3202130021213100-0332213111312130-3013203132210221-0002322222313003-0101110211113031-2302011301022222-1201330200120031-2311131213330003"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122100202013221-3121323200003002-0300332113222130-2201310203031010-0132330231003111-2302113002130311-0220032032211321-3222012211200220"></a>

<a id="canonical-0012322033312200-0033121210022231-1211112122100202-1333212100022121-1111302101322333-0301220203321032-3120111010021010-1210100021101321"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

Type: `"string"`. Optional, Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3203220232101010-2222003332210002-3222023000011321-0013112112322330-0223333210201120-3001230033131332-2011220223200130-3012303213102000"></a>

<a id="canonical-1021112023103023-2321100101011301-2303301222320222-1121110322010103-0121311111313110-2323323013001200-1213323230302001-0231330211202133"></a>

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1131100213132110-3310333232222331-2231310112312012-2311333102233132-3301112001213202-3003223102201221-0101320232303301-1030332212123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.ip_allowed_list

<a id="canonical-2113300122210002-2102300000023333-0331023230211021-0223322202311012-3201213122010312-0300301003132313-1030013233203331-1200121302032120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033210200301322-3131111102220022-0112222203200212-2230221201322320-0030100111000211-2031020320011122-2213011323213313-1020200033103000"></a>

### Direct properties for `api_rate_limit.ip_allowed_list`

<a id="canonical-1200223231301323-2332201332013020-0330100332012013-1120102211210003-0122333132211330-1300313213201003-0223231323221120-0011111203200113"></a>

#### `api_rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3313313301303203-1122102000003112-1022231311223221-1302133120330011-1231301301331313-1323332102200203-1110313222013213-0023131100032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-2211020100021201-2303322033023302-1320120123310132-3100313000103030-2212300132032222-0220322212111323-2300230131130123-0232003231100213"></a>

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
no_ip_allowed_list = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.server_url_rules

<a id="canonical-3020022312300323-0013213333320033-1033311032012200-3232300302011331-2210202121020323-0211000203231233-2300122221003000-3011012130232013"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130113223231313-1232032120232210-3120213121111030-1033213120032003-1313000322302322-2002201312020311-2203023112132210-1222213211231232"></a>

### Direct properties for `api_rate_limit.server_url_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-005.md#canonical-1123020333112122-2130131011232202-0220230110200101-0130122132022212-0102201333003233-3220222303131213-1000112121212301-1323210210003012): complete subsection reference.

<a id="canonical-3210002020300133-0133200100302131-1123023033112033-2130113333333101-2223123301032112-3122110100113003-3222303010032033-1101231111203010"></a>

<a id="canonical-1231332210221220-3201032202021320-0000103331201202-1230120303110012-2330311302311111-0110112033310311-2220210231331110-0311023233112132"></a>

#### `api_rate_limit.server_url_rules.api_group` property

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3132312122320133-1002321111011313-0103002220322023-1023213311000220-0030301100121213-0313122022221030-0033132322323032-0202131021200230"></a>

<a id="canonical-0110312120022200-3313033310101122-0002113010322201-1121033301330230-3002121320123213-3211201220132033-0300022130111003-1333101222012001"></a>

#### `api_rate_limit.server_url_rules.base_path` property

Type: `"string"`. Optional.

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

- [client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311): complete subsection reference.

- [inline_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021): complete subsection reference.

- [ref_rate_limiter](resources--cdn_loadbalancer--reference--group-006.md#canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202): complete subsection reference.

<a id="canonical-0221211213300122-1111333233102103-3021203101113132-0013002031331230-2013313300030231-0200231113011000-1302022231323021-3013302311233211"></a>

<a id="canonical-3021030113100112-2231332230222321-3001210220013123-0133301023031020-1212212201202121-1102010122332002-2310020011232333-2100232221333030"></a>

#### `api_rate_limit.server_url_rules.specific_domain` property

Type: `"string"`. Optional.

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

<a id="canonical-1123020333112122-2130131011232202-0220230110200101-0130122132022212-0102201333003233-3220222303131213-1000112121212301-1323210210003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-3323230123333111-2220023302313332-1203322110311002-2123330210032133-0212302020012321-1321011121131313-1233101233012020-2021210012031020"></a>

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

<a id="canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-3022123012300011-2103200333113130-0102301133332210-1211301022221310-0311233230030231-2302321330302021-3222033010131111-2200330020300202"></a>

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

<a id="canonical-1111320203131221-0032003221131003-3030112302030003-3022302323232121-2101301002303122-1022103210221300-2120131223331200-3233133102011000"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher`

- [any_client](resources--cdn_loadbalancer--reference--group-005.md#canonical-1001121311313002-0320330200133230-1323323031130323-2203320203133231-0311032312303320-0012313023130210-1330202011323221-1101331223002310): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-005.md#canonical-3211100132121310-2001123001112333-0210021011001302-2030301130323123-1232210201110233-0311333300232100-0311231200201322-2211320102020111): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3021213112110103-3303110222212011-3311210320211323-2033201233112103-1013012103023313-0330003332000011-3100011230000132-0121232220300223): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-005.md#canonical-0011031133302211-0202030230313323-0033223313130231-3010013233323210-1023223112002133-0201303202030302-1023011112031120-3133331110222023): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-1121321233302210-2201233030311302-0233012110112111-3301022303130020-3330230211332201-1122013113002133-3100021113102003-1302122332233003): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-3000100122201321-1031000322331033-0320100200312313-0310123303011001-0231212022301023-2030131301002101-0231000210030100-2010003030323113): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-006.md#canonical-0130003213232033-2202300221200002-2331102232130333-0210030322312120-3203313330213022-0310130203033112-1123331112033221-0112233032020000): complete subsection reference.

<a id="canonical-1001121311313002-0320330200133230-1323323031130323-2203320203133231-0311032312303320-0012313023130210-1330202011323221-1101331223002310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-3123110111101202-0002001213332012-3310023202211032-3123133332001222-2122220231330132-1111232231222211-1113311121013213-1231012332323121"></a>

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

<a id="canonical-3211100132121310-2001123001112333-0210021011001302-2030301130323123-1232210201110233-0311333300232100-0311231200201322-2211320102020111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-0100323203331322-2033223120312221-0210210332103201-3003303023201133-2000032210301022-0112013022110200-1030102100332132-2311321330121233"></a>

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

<a id="canonical-3021213112110103-3303110222212011-3311210320211323-2033201233112103-1013012103023313-0330003332000011-3100011230000132-0121232220300223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-3003310000130221-3211202233131031-3232223231231101-2111033313111333-2110233120230221-2200303001022122-0000211303102333-0022223212030001"></a>

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

<a id="canonical-3122112213112121-1330023112113230-2130031122303021-1202003233320233-1321332212131302-1030221220330132-3031000123313023-0111203031100033"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_list`

<a id="canonical-3213100002313123-0312121303011121-0332213023302222-0003113201231332-3210212233000231-3121011231030210-1300320003223321-1022001301222212"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_list.as_numbers` property

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

<a id="canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-2321223333223122-2213320101023223-0100110102303113-0303010202003112-3100101212020320-0100120322232320-3222001210023000-2223212321232100"></a>

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

<a id="canonical-3201310112022032-2023103133113001-2213131111331233-1033123332120001-2332220210110003-1000123322110231-0132301103012120-1210201203232301"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher`

- [asn_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-0020011133130001-2133103330222323-3200110323220330-3113332022021002-1101003110330122-1312000313210012-2322333200133203-3022332123223111): complete subsection reference.

<a id="canonical-0020011133130001-2133103330222323-3200110323220330-3113332022021002-1101003110330122-1312000313210012-2322333200133203-3022332123223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2011210033032032-1223022200233010-2211012313321331-0311120032133121-2122313023230323-3012231031300333-0003020002310111-1330011233111101"></a>

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

<a id="canonical-3111321032223010-0333313030012033-2220032111203312-2113033310023333-2111301122030223-1033212300120232-3230101321310133-3003300223201300"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-1110003321222232-1102313112131201-1130311223232010-3321321011003023-0011130232033033-2320130002130111-2220203001032232-1101012130010232"></a>

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

<a id="canonical-1023211130023111-0033131013311213-0230100232020031-3320332222011011-0311331212032231-2200223022002312-2312303021332111-2021322113110003"></a>

<a id="canonical-2231001010301100-2021301102001323-0112131203313222-0201023023333132-0301202031001102-1002013210210101-2313120132221303-3221302330312233"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.name` property

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

<a id="canonical-1103311032212223-1010322101001222-2120012323122302-0320002201233023-1101311333033223-0222000312100203-3232210302110212-2002203302032031"></a>

<a id="canonical-2100321122132122-0103320123331123-1332030202131220-0220302000213130-3322032203320223-1132200122303101-0133120202222301-2203211313010330"></a>

#### `api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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

<a id="canonical-3010213322220013-2003321221111332-0020330222030230-3013032310333032-3101221023030102-3300010032030022-2033021333303330-2211321023001020"></a>

<a id="canonical-1200011032231232-0210122322303022-3001003021312200-3131303301010011-2202312130021030-0301133001010021-2311203232102103-0132032232021103"></a>

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

<a id="canonical-2103003122331323-3331100032121200-1320321201333222-0330223123132303-1100300200213123-0032202020201331-1310311230233021-2031023022300332"></a>

<a id="canonical-0133300233112221-0130313022023231-1002011110110131-1021311100333220-1131100130123203-0232123120212311-0201333322203132-3312222330132312"></a>

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

<a id="canonical-0011031133302211-0202030230313323-0033223313130231-3010013233323210-1023223112002133-0201303202030302-1023011112031120-3133331110222023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-2330202113000011-1312121010112333-2233133100002331-0000023113122010-3121211212303101-0302102112030031-0331101102230120-2203223000213110"></a>

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

<a id="canonical-0123232023111100-2333031131322033-3112212032113333-3202123020313332-3200203200002300-1122100202113010-3012331030230310-3031233030330201"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.client_selector`

<a id="canonical-1211102011020310-0212112023201101-2122123300203101-0310301222021120-2213023313203312-0233233002132021-2330111030112203-1132200023110132"></a>

#### `api_rate_limit.server_url_rules.client_matcher.client_selector.expressions` property

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

<a id="canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-1133122320001010-0313103132200101-3032123312232212-1333331321120222-3002030031001001-2322200013010013-3133033312223001-3101030212321101"></a>

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

<a id="canonical-3323210331221133-1101331223210031-1121220020230021-1100031013301211-3111212132201003-2133201031302112-1030223010202012-2301111001010021"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher`

<a id="canonical-0022201132310221-3313123210033332-2033022113023313-2112300201203312-2210322133030112-1010310030122100-0110311013232332-0101333201100003"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123): complete subsection reference.

<a id="canonical-3313002220000222-1013023302210203-0331011102101332-1321220011231113-1201123222220003-0201321312303122-3000133220202331-0221333322011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2130111131232013-1220123201311202-3133000103231121-2322113333330323-0202110311211331-3113022121103001-1301223132322213-2023303310002223)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0101031020122023-1110211033210023-1110100213222020-0120230323100003-1110010332130031-1103232221320222-1232221333102133-1111131212001023"></a>

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

<a id="canonical-1133022102000122-0333103331323032-2130303321010330-0130201223023121-0333101132001020-0222010312221001-1333010123222100-3010111322223001"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-2233301113300221-2100122220032130-1112032112330033-2213003312310012-3110023012330330-1102210002202233-3202002113022323-2211212032131020"></a>

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

<a id="canonical-2010311313312031-3030200102030331-0310211123231023-2000313331113100-2233312003211230-2012221230121233-3233311221301333-0220222102101102"></a>

<a id="canonical-2303313220202022-0303230112112012-0210303231021210-0302313313220123-0132223020001312-3321030000103331-0120112110122203-3302010000102030"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-1001210130300100-0211013111110212-1300113100110233-0113103031001113-2312010121233213-3122101211121132-3122302322113101-1033123110302301"></a>

<a id="canonical-2320033030212331-3031122120320323-2113333332222111-0201313112320203-3302323013022022-1013322010120231-1323003201131213-3231232232002323"></a>

#### `api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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

<a id="canonical-2230322021002103-1103130033301003-0231212221211032-2112203312103220-2000132312120131-3200020211033212-0313001110332032-2122200210132320"></a>

<a id="canonical-2311313313232320-2232233221332023-2320220132212233-2001300011003123-3112300113123220-0002121030312313-3131130202033133-0113112120221301"></a>

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

<a id="canonical-1100302313323012-2303303132230020-0013002222321032-1012302203132030-2002330111030013-1033203311121022-3031101133233020-0322123110002200"></a>

<a id="canonical-2200033022132330-1013032221120002-1320220311223212-1023032300110220-3223032203020123-3021113222121221-3232023231110332-3012122013032130"></a>

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
