---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-1232013133302300-1300101221213231-3122310133003330-3212120233300022-1203321122133210-0132200233133123-0201201012313002-1231011220022022"></a>

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

<a id="canonical-3330122111302212-2131221203332012-3002211323333110-3103013020133100-1301023321011312-3310302232213301-1020311002112210-3302123123000323"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302): complete subsection reference.

<a id="canonical-3211320332202322-2122010322301020-0200333313031203-2113001101233021-1121203003102231-3112132222330103-2201302303222012-2002330310113013"></a>

<a id="canonical-0113012330312133-1323010001302113-0030101000132003-3331310220010001-1333301133033101-1021120301100111-3010323301002020-1320102230302120"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321): complete subsection reference.

<a id="canonical-3222132112301332-2201102220013330-2110030010211230-0223132213201322-2232301232112113-1213102101222321-1322233021100010-1203212201021021"></a>

<a id="canonical-2030102212200201-1300102302220131-1113123222012203-2130100220112122-2120301212131320-0303323231031111-2322230030001022-0131302213333103"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.name` property

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

<a id="canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3032032302123302-0113330330110111-3131110303311223-0231031021321101-1113111112102102-0313012133022022-0310321311101032-2013320103313330"></a>

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

<a id="canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-3321322012113230-1222111002110231-0032333313111000-3203313231320313-2133021311213320-2312332132313131-2212030231201000-0121233033103201"></a>

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

<a id="canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-2213313330003003-2303333221201323-0120323223313300-3321231331023320-1332202312012023-2031201122211011-1020302223133112-2112211102233020"></a>

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

<a id="canonical-3012320313322333-2330102100311212-3110032011110232-0110331222313221-2012130230222100-3310101103303332-1112123332233303-1013003211323320"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item`

<a id="canonical-1313332132000210-2123102233210031-2003323122210302-2101301321133030-2001233332003130-2013130203020101-1132311301322123-3133231332220001"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.exact_values` property

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

<a id="canonical-3111202232103230-2122212001131211-0003303132310123-0103210232333302-1300320202330210-0122230223311031-0030133320333200-0113300223221213"></a>

<a id="canonical-0323220203301213-0312120002111100-0322130000101331-1203220213033333-0221323001322200-2111323201233313-1013222302121220-0003231003131301"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.regex_values` property

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

<a id="canonical-3211312132002112-2012201321103111-1013211322330302-0131311330222022-1131323020022013-2233111322102332-2320120311201022-2012001320030220"></a>

<a id="canonical-1210212233003233-1213121013230132-2221000223330033-1223023220001131-0213013203122010-1110301302133031-0310012233001202-3302322331331313"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item.transformers` property

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

<a id="canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-1222322313132320-0121201111103001-0310330222302331-3300121002133013-0212320213332223-1010201113301323-3323331210333222-3003102122300323"></a>

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

<a id="canonical-3311033321031220-1213031301213231-0320301312103100-3321310302321003-1331132023231220-2111332123202211-3210231223113023-2210020301323020"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-007.md#canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-007.md#canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301): complete subsection reference.

<a id="canonical-1023302032031321-3220320320003102-0210020023313132-0020320303330203-0120103302122321-2210131113011322-3221122222221002-1011031230022000"></a>

<a id="canonical-1203203032022221-1303003020130320-2013221120001321-1130313110000013-1131212012110213-2033222222300320-2120110100021022-0122113011303112"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-007.md#canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332): complete subsection reference.

<a id="canonical-2032330003113313-1121003133303113-0110012011000311-0021033110102130-1133011003113213-3010211323001223-3202221212020122-0120223312303110"></a>

<a id="canonical-0120201132213330-1002100311223032-0122330121131002-1222321111332110-2222011331011233-3303323121223310-2120232131011000-3031002302300231"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.key` property

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

<a id="canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-2010110013031313-2222210322211232-3300323112312312-0123332133223110-1130112221200023-2300101130000212-1033310313033323-2032320122020113"></a>

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

<a id="canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-1111032032201011-0021322011313120-3130001213330213-1322321110131001-3032320210220032-3133012211223313-3313101013010022-1230132301102313"></a>

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

<a id="canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-3003321001211231-1231132103323022-0110323000030201-3132201321311111-2122203330002203-2100131030011211-3000133210310210-1331131203010213"></a>

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

<a id="canonical-2113230003020010-2133203300210203-3030011200010320-0223232100003332-0311131232310331-3323013200020001-3213201131333111-2133112022120011"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.query_params.item`

<a id="canonical-1212320103321131-1322000031212221-2202331200311201-2332011311111231-3313323321010230-1310122300013113-0133320220100031-2223313303023012"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-1221311003102213-3003120000333013-0033233211133003-1033002202212021-1100310010203330-2012020112011222-0002331031321021-0301121100232013"></a>

<a id="canonical-2021320202222203-2223231233003333-2102022302330021-0023331202032311-2121311323100013-3302311202101100-3113301201003211-3213023332222121"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-3311201032012023-2232321312203001-3010302311010033-3021231102320112-0300013320300022-3302022031233013-1002033000011222-1320300020200301"></a>

<a id="canonical-1103110321233320-3011003112332312-2321123131303220-1322301133202110-3123020203113033-0023200330100010-2023120130103003-1300303231103120"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.item.transformers` property

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

<a id="canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- api_rate_limit

<a id="canonical-0233031301102130-1130230220030210-3300120230310332-2002022322233223-2023100032330031-0223303022120030-3300222102202013-2011031031103200"></a>

Type: `"object"`. single nested block, Optional.

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

- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0233031301102130-1130230220030210-3300120230310332-2002022322233223-2023100032330031-0223303022120030-3300222102202013-2011031031103200)
- [disable_rate_limit](resources--http_loadbalancer--reference--group-017.md#canonical-1013223133201022-1313332221032232-1331200200123112-3000301121301112-0033001302203101-2312333301123320-2111320003320121-2311310003323022)
- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1212231232113102-0000210123322230-2321331212231011-1223210321030302-2020032100331231-0213102112123132-3222021323133333-1010333102020321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212101002021330-2311301211310331-0110120033203310-2002330003132102-3013302102213100-1320231030000021-1000213323323000-3301303122001111"></a>

### Direct properties for `api_rate_limit`

- [api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220): complete subsection reference.

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001): complete subsection reference.

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-009.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-009.md#canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-009.md#canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321): complete subsection reference.

- [server_url_rules](resources--http_loadbalancer--reference--group-009.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103): complete subsection reference.

<a id="canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.api_endpoint_rules

<a id="canonical-1132131113333103-2022130102320022-2230110113330212-0132031113100120-1231101222210233-3302112221010212-0100032312311122-1323013233103113"></a>

Type: `"object"`. list nested block, Optional.

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
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011212333003130-0132202111201102-3110103033003313-0301220102103110-0321112020021202-1202331201201212-1303330232303123-2203313211303021"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules`

- [any_domain](resources--http_loadbalancer--reference--group-007.md#canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--reference--group-007.md#canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103): complete subsection reference.

<a id="canonical-2331022231313010-2113003022132122-2210231312301110-3003323021102313-0012011003003223-2002113203031230-1031120331212231-1312000113102132"></a>

<a id="canonical-3211003130301301-0121023210021232-3000222030233133-3211220112022120-2111022101231333-1010003012013000-1113013310331220-2303311012022203"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Optional.

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

- [client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122): complete subsection reference.

<a id="canonical-3203121311310230-2100333322210103-3231301201012120-2322101111331110-0113333001101211-1133203230213233-3123210021321023-0022012232220120"></a>

<a id="canonical-0210130313110330-2130103222121211-1001230012003222-0223300221112033-2231113120231323-1322012221123130-0213320101321330-1131301223232001"></a>

#### `api_rate_limit.api_endpoint_rules.specific_domain` property

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

<a id="canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-1110303132222231-0210011010211110-1302312302110020-2010010132211001-0000103222331131-1231220010102200-0001112322312021-2123332133301000"></a>

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

<a id="canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-3133330222100221-3123200201332122-2301210032320032-2220232302020100-0130003323201012-3230201033330132-0010230123200332-1031122132002111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
api_endpoint_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103122013102103-3032010220333310-3203210322123122-2323132101213103-1302131212100132-3033331211022003-1132021212010300-3110213230310130"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.api_endpoint_method`

<a id="canonical-1332313221233021-0302301130011130-2210003302132300-0221303303112220-1333311101202123-1120020012010322-3321332311203013-1103223322130330"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.invert_matcher` property

Type: `"bool"`. Optional.

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

<a id="canonical-3003211102231313-2123223301120102-2200131103301203-3122013023013303-2332010312202003-3023002130322022-3001221323121323-2002010023100123"></a>

<a id="canonical-2230210332023122-2022023323310312-3031301020300222-1302102323103132-3300213023012301-0013000022132022-3233010332112201-0020132120301003"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.methods` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-3102022023020021-3321211122201030-1330012333122230-1002010222001122-2201312311031222-0131232321011013-3232103110111321-2000223122010220"></a>

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

<a id="canonical-0310121103222333-1321203010112201-1320031000022231-1331230311002231-2013331001000012-1003103212120331-1200301000323110-3223122321220311"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-007.md#canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-007.md#canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-007.md#canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-007.md#canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-007.md#canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-007.md#canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123): complete subsection reference.

<a id="canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-3323010020322300-0313023030332200-0222203212100021-0221102222200111-1113321102022103-1003311313201032-1000103221012213-0011103032313231"></a>

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

<a id="canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-1232103001103323-2123100133003223-1310020131111032-1202311103230302-3031220211232332-0220203121021033-2331102021013203-1220031123021212"></a>

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

<a id="canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0003333232231212-1232330320133012-3232102103300113-2113000113003231-1200203303123002-0103031223033233-2231230101310100-0203002312213022"></a>

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

<a id="canonical-2030113210101323-2313111130210122-3300013220333110-2301213131033001-1130201320212210-0230000302321220-2232011332130123-1320003023010202"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-1200121330130020-1002313201331323-2112331223232002-0230311131021021-0222333323333301-1211022100310010-3032000202221212-3032312102200032"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

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

<a id="canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2033113121211121-1130333112313221-3322012121000020-2302223201001213-0221013300130330-1333232121320320-0200101303301322-1303321123000000"></a>

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

<a id="canonical-3311032022131202-0101110000131111-1123003120311233-3200303103033031-0120322302011013-0300013131232033-0301120300222002-2200021221231022"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-007.md#canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323): complete subsection reference.

<a id="canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2203322122202032-1310023322220032-3103302022313223-1031311330322120-1221030232321123-2201101132112320-3012023311230213-2103333211021130"></a>

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

<a id="canonical-3110321203220223-0131320121132103-3231312301033003-1233023311110000-2010331310320310-2100300310011011-3300322010331010-1111131330212021"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3010012003333110-0123320332330311-1030330330111023-0133233322100212-0211123322002231-2313100310113110-1033312023213321-3210222101200120"></a>

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

<a id="canonical-3033013223013232-1123010333022010-0032312230211230-3301322113212230-3333301220222130-1323102312130330-0002211112313213-3010012301121201"></a>

<a id="canonical-1033203013002031-3202322323212102-0131301210303021-3233231223202103-0303101331323232-2302023310000001-3300201312220321-0133331313203232"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

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

<a id="canonical-3112221202003310-2030202111032232-2023221323223101-0020333131331022-3333022313311103-0021113113313012-0111231010001113-3312313332303013"></a>

<a id="canonical-2033302200300202-2303210121110100-1211113002010230-1203110032012202-3320133112230121-3023030113121211-1202212130330013-3302112300130231"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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

<a id="canonical-1231202131222002-2130110101321030-0012233222113321-0212213223113301-1110003323021212-1011122121120201-3320000303101120-0022122031113202"></a>

<a id="canonical-1232200131112302-2312203003022130-2033322132231023-3212031220222320-2031123313011002-3012201131222021-2221311203102300-3102333011100031"></a>

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

<a id="canonical-2020123320122213-1233123013332321-3301333201012200-3320011232311230-2032303122010221-3233301230231133-0233023330122102-2322001212233311"></a>

<a id="canonical-0213033212133330-1230331321321231-0233201232221001-1020112001120020-0321103312112101-1011330202212000-0022130333210210-2331211121110012"></a>

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

<a id="canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-3320313301030012-2133022223100122-2312030323133302-1331133321333313-3102020012111312-1321300000311311-0031131100323200-3033332311332133"></a>

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

<a id="canonical-3102200230130010-2330010033303202-1330103103102131-2031032112231221-2301023130321022-1311111111100100-1331111201110212-3333212223331231"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-3100120202202223-2202313020231022-0033330123012130-2122003320002323-3230212123223301-3013003230222202-0020100303101130-2121312302122200"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.client_selector.expressions` property

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

<a id="canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-3321023330020202-0120001131222133-0003103331110231-0310221223223232-3220020303120122-2112200130231330-3301301323201203-0210032120113320"></a>

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

<a id="canonical-0022123323311010-3312301032212112-2103310302032131-3303013302100132-0012012313231230-0031231023120313-1312110203003201-2001013233131311"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-3200001321311111-0113320233303132-3310000200320320-0213020013333103-2302120220230103-1011333031221300-2012011223201123-0131311102323212"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--http_loadbalancer--reference--group-007.md#canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132): complete subsection reference.

<a id="canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-3223122113221002-3203110203021002-1120312102330123-2111330011313133-2130210121001302-3320301011012210-1122001031303233-3213120023321111"></a>

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

<a id="canonical-0102220033101110-1202010120033230-2322212331110102-0031331122310331-0311320002002233-0110302312301331-3033132230330003-1311213030022021"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-1200223103333112-0111330330011203-3201122121332212-0233103332212231-3002313133032121-0303310210013011-0211020121203120-3032303210331102"></a>

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

<a id="canonical-3231021232232133-1223310010001113-3113030232101212-3332000210003021-0330032110220311-2200321020122101-1031213002021111-0201100111022120"></a>

<a id="canonical-1302002303133031-1133220133022003-1032013210210301-1211312111100010-2133232133030321-3201221120212103-1211203123133003-1102130130033301"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-0221132201233221-2121303130131103-3301023033122020-2210000133311302-3203332312233331-1203320132203211-0321131020103112-0322321023331223"></a>

<a id="canonical-0110331323201031-3303112130101230-2021122030232121-0131223132323101-0102333013123323-2231131033022223-3110123110310011-0212202031201032"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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

<a id="canonical-2113213213200302-1303120321221231-0320002220300203-3210100323132033-3201102333132302-1221031332203303-2122121021210331-0311232231312313"></a>

<a id="canonical-0100111121331132-2310013120113221-2001302021022131-0200023302333010-3120202220202332-2311031322200210-2330323133020200-3300332113022322"></a>

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

<a id="canonical-2133131203130330-3111013001213302-3330102323333023-3020112103333211-1130033310312022-2013233331213222-1300212323100122-3310023133320330"></a>

<a id="canonical-0222001211003212-2021313103100211-0321310301023132-3332131112230233-0023233123020111-1302213132210021-2321312203013012-1123322021031310"></a>

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

<a id="canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-1330322301103320-3002021223203303-3011033110031031-1213200322231031-2120210010003003-3323302313112232-1033011113032013-0220223330320030"></a>

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

<a id="canonical-2002331231313111-2133331120222011-0122323022322201-0020023132020122-2121301103000200-1220031011033211-0113332023031223-2201023233000000"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-0103320030100023-1032133201201222-3002302030331233-1021301201230301-3233231302102112-0100111331001311-2323111220023121-0013212211011222"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-2132131203111130-2122112133232030-2331323221012102-0210330110100132-2213032020210010-1121110033012210-0111233321121222-2301301223123023"></a>

<a id="canonical-1110310003023010-0230122233223010-0113230100303102-2131231113123030-0300122220203303-0133303102000123-1301100212021300-3232031122323322"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-0001121323103103-3300130121010021-0201000132233033-2000002010301303-1003323220132001-2121303301200131-0132000231200131-1310222100032211"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330110002323031-0022221022321223-3131033012221030-2331110203031100-2000023220131002-0021202301231022-3133213121020203-2300031120013331"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-3112103230121320-1030120301101101-3000030123011332-0203323222330031-3203333031003310-3103102131221202-1302221000310022-3313330003000300"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1020030210001303-0033221023301031-0310321013333311-3213230212231002-0311330133131000-3212331221002112-0311203332333312-2302321033001322"></a>

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

<a id="canonical-0211302001310313-0202223213210003-1310111020022232-3323212133313313-2111101222213000-1132130220003123-1100330212311202-0221200202102113"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0132321300203110-1203002233331023-0312211320210030-1311302233320132-0203013113230303-0213210003133131-2013022120013312-2130003200103030"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2110312103111032-0111300133313231-1122223111131000-2331332102112101-1102321110003123-3022323203213132-2231233120303023-2100313131112032"></a>

<a id="canonical-0121231020112002-2033323113110133-3000200131201023-3130213102232311-1322312212122010-0023101213223313-2201010222202322-2021013320233313"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1230110222023032-0202123322201032-1231333212321331-1030033123201112-2220320312011211-3102312132230211-2120002321310013-0111121033000203"></a>

<a id="canonical-1212110122101001-0300002330311301-0132011120211320-1122110111233000-3222212122200320-1220213032223113-0012212233012121-3312321300311121"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3032133021102233-1331030100031322-3030202220021113-1223322100201011-3110301231002132-0101010211022101-2100000313002220-0202120031122033"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323223100331201-2030110331220232-3100302003012123-1213110000013232-2300230131222131-3120012132201221-0110101310201221-3101233023211300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter`

- [ref_user_id](resources--http_loadbalancer--reference--group-007.md#canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003): complete subsection reference.

<a id="canonical-1311232223200320-3301220321321211-1112320123123300-0113023000002303-1012120320103000-1232021121221211-0002111333301032-3231202000323111"></a>

<a id="canonical-3332012121233030-1201333121111211-3313002113033302-2330003323232123-0023213330212102-1110123203322201-2100221132211112-3231023230110220"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2130211202000102-1122131022031000-1302113200201232-0220323001230131-0330332013020312-3032211201003303-1002000131331113-0110000231130312"></a>

<a id="canonical-2230211213011333-2203130221132121-2212222221321110-1213312130210133-1031133220013233-2023002210230022-3031233323330323-3313103100031221"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-007.md#canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321): complete subsection reference.

<a id="canonical-1322000210130123-0113311310103133-3112203131210121-0320302333110311-0113031113200002-3223121312302200-1113233233203222-3000010213202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-0232331112003121-0302311101311331-2331313120203301-1212132313310033-1202003333201312-2300103323213001-0033223300222010-0302110032012222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301110103323102-1232122011123301-0220223023223323-0301021000111232-0103333313111023-2223302300232011-2121310122311101-0002202100223123"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-1313010013102033-2130020321321311-1210233100113032-3103323032023110-2333010023121012-3230133132223220-0001201101231200-0110323322302131"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.name` property

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

<a id="canonical-0120213103130302-1201010301322301-1221020302323121-2211000313023332-2021213221122122-2231312103230020-1133021322211001-1311311033100013"></a>

<a id="canonical-1222302331203102-0333302210231033-2000133012013122-3203133033212131-3032222230210320-1322230231211222-2310300322232323-1330221032103320"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.namespace` property

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

<a id="canonical-1102030003221110-3032012132102022-3202301330132322-0210113301221213-2031110021122230-3221132132321111-2101000132101320-2202301212302032"></a>

<a id="canonical-0010230132013112-1012111231310300-3303112233302303-3112233223000201-2130022233332231-2000011103331212-3211020000310231-1113110133111012"></a>

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

<a id="canonical-3330031100020111-3030203233132221-1000321133200203-1321313131010003-3122213211133010-3300211303020132-1200031300300131-2210032020311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2203110110331322-3032012202233323-1113021333010211-1232210210220321-1332202303223133-0001323102023022-0221103020311310-3331000102013303"></a>

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
use_http_lb_user_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-3112213233313300-2121330033100010-0303132133030301-3320002023002222-2232220020303001-3300000201222122-2210013213011210-3033321100331201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021303010000330-3201110102231212-0231311322320232-2020300200311102-1132232230311013-0321223012302123-1302203001200010-3123030200102300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.ref_rate_limiter`

<a id="canonical-2100123312323022-2130220123200132-3201213220130011-3122103112103300-1032030132220232-3212003001102133-1203232212032221-0310033112223322"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.name` property

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

<a id="canonical-1302033100220222-2131010303211321-1211122221202311-1232112033101330-2012322002323201-3010130332322321-3301232332132111-2122333101223031"></a>

<a id="canonical-3131021130030200-1222223001000203-3003132131002010-2022022112111010-0220233022013311-1002122210013303-1210130100333113-2222112111021001"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.namespace` property

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

<a id="canonical-1222332210122300-3120310010333223-2032021230211232-3231123131001020-0111331013132320-1213103020100023-0213111231301032-1111312213330321"></a>

<a id="canonical-1223100302310010-0122030302102321-2232321032220311-2220211100323301-2003131321003230-3011321331020230-0002210230132012-3303001320323312"></a>

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

<a id="canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-1320020013211132-2031013002002013-1203201303313130-1213200320011033-2103201322111302-0133123010012232-3213222101200313-3302231313332121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012101011333320-2312001320032131-1123020300220002-1112121103001102-3020013131033233-3012130310233311-1320213033110110-0021113230011231"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-007.md#canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-008.md#canonical-1233331212230221-2313030123121321-0110131013202122-0110202121333133-2101222300301122-2202000111102312-1302211021231330-0133213302312221): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0131131201330232-0133320302200313-0033331320133212-0013033331301110-1231220121103031-0100333301213133-2133201011333021-0022313213110112): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-008.md#canonical-3011331233323332-2233203231030203-0223111130002031-3203010133313320-0233122301022001-3003032320313102-0013110210203333-3112100230212321): complete subsection reference.

<a id="canonical-3030212201122103-1222020231020231-0032300110132011-0021101210320033-3002230110100312-2303020103212110-1311331022320012-0131123220112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-007.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-007.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-1312232002030230-3031023203302003-2233331213111113-3230312233323132-0322031211203333-1313120203133101-1023232211001021-1132122212221201"></a>

Type: `"object"`. list nested block, Optional.

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
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020312320222103-0332033302031213-2313130333022222-3202213001323220-2332022203221102-0103113000321203-1112203311022013-1201001210131123"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2131023022232100-1313031230102113-1322331330011233-2003100010212013-1221031002330332-2130233033302100-1301100011232323-1232303001203211): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-2122101123213330-3233231220232002-1332013023321002-0223022233200113-3332002102113313-1330000032111110-3200220223030122-2200323201130011): complete subsection reference.

<a id="canonical-2013322200103021-2332012002310130-3302020130113310-1310032220221231-2133100321132131-0322310111003223-2333013332212003-2113103000213000"></a>
