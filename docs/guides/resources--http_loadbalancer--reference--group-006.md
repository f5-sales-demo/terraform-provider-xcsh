---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0222110311021333-3223030312130231-3310021100012103-0312303330033033-0313312020021201-0200303312221202-0212002021213013-3101033321131203"></a>

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

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

<a id="canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params

<a id="canonical-1023132001332122-0113230213001032-0002223023002321-0032132123330233-3022101201132020-3221310211130123-1000320033023210-0222012001132122"></a>

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012120032030323-0210213020201220-0123030322030320-3300302311321222-0303233002333020-3222201103322210-2303200022003130-2102201310103113"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-3011220122222220-1301122320210331-2213302020311031-1211133011333120-2121203003121002-3022002102301322-0102101101011301-0310222231300312): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-2023121032333321-0233223332301213-3330322320230301-2110333020031332-3322122223202121-3133323020311310-3313131120030212-1331030112331012): complete subsection reference.

<a id="canonical-2021210020301220-0101221313121232-0202200311101031-2332120133100313-1121301021202113-0223220033022012-3112332100322013-2302000203211013"></a>

<a id="canonical-1103312121312122-0000112102233322-2313333330332132-3133131302213033-1220322112010320-0231021111222303-0012210133103121-3301313012320233"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-1302013023113101-1311012201231101-3200232200130033-0033312220111233-3112201120300322-0010121230021030-0002300313230232-2203312103220012): complete subsection reference.

<a id="canonical-1220012002213031-2021022320033023-3223203221132202-2312312311330312-0311230032331012-3301013132030331-2120111000120331-2223322301230123"></a>

<a id="canonical-1123102120220330-3233003321333200-1012321101130012-0000033311320201-0333211013011121-0332012033332121-0301101000310110-2132310121312131"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.key` property

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

<a id="canonical-3011220122222220-1301122320210331-2213302020311031-1211133011333120-2121203003121002-3022002102301322-0102101101011301-0310222231300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-1322203200231100-2322303120212022-1011113111332202-1320133203132101-2230333323231231-2203012123031231-3322020220223020-1131013310122123"></a>

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

<a id="canonical-2023121032333321-0233223332301213-3330322320230301-2110333020031332-3322122223202121-3133323020311310-3313131120030212-1331030112331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3003233333332201-1303101031131321-3112320010222232-1212312230311223-2221110331013220-0222332311130313-3033313130212101-1131103133122301"></a>

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

<a id="canonical-1302013023113101-1311012201231101-3200232200130033-0033312220111233-3112201120300322-0010121230021030-0002300313230232-2203312103220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-0211222202323310-1030023102123130-0032003101030030-0222020201031103-2133111220333211-2023001223333001-2000102310202331-1301113033222200"></a>

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

<a id="canonical-1301201120113032-1131200221123322-3303103321110212-0130300300032003-1130200321322021-3232230131321210-3022002000002202-1213022121332323"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item`

<a id="canonical-3211002111132202-2030321212002201-1131301103231230-3222110311203301-1012113210022200-0023332000323200-0033200012030020-3213310010220131"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.exact_values` property

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

<a id="canonical-1101313110111133-1303313203210231-2201233123010221-2112002332031232-3320103232333222-2212031232210220-2131211203121213-3211001133322023"></a>

<a id="canonical-2301221210020133-3111101002001020-1102023233112031-3222212030202023-3010233231221002-2323122021213213-1212130220321003-0123031031022111"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.regex_values` property

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

<a id="canonical-0221000333231023-2122101233111102-3312310133313201-1111102031100310-0001331001111200-0232112032123022-2301333320121012-0332032023133210"></a>

<a id="canonical-0000110032303102-1312033201301203-2223303113213033-0112123130123320-3331100230031002-3213232121233233-3001312030223111-3210321323113030"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.query_params.item.transformers` property

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

<a id="canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- api_protection_rules.api_groups_rules

<a id="canonical-0301120020103121-2203112221132003-3111131222031312-2103330211133333-2113012021302113-2233022000213032-3212313032122202-0113101022330113"></a>

Type: `"object"`. list nested block, Optional.

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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_groups_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223020322113031-2033210323222222-3210113100222212-3110313221100121-0031212001330120-1031121103200301-2230232322111132-2303100122013320"></a>

### Direct properties for `api_protection_rules.api_groups_rules`

- [action](resources--http_loadbalancer--reference--group-006.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-3213023311121323-1113021222023213-1131132223030021-3332000100022232-1320331303331232-1113022133331012-0302020130233200-0320011212213120): complete subsection reference.

<a id="canonical-0011230112122322-1330113212102121-0313133321120331-3211102330221112-1210001220013230-1120002320302103-1031003022002203-0030120033300112"></a>

<a id="canonical-3331301033202332-2201223333301220-2000002333032113-2111032010101220-0030122130312033-2313001221202122-0101323130313313-3222023200323300"></a>

#### `api_protection_rules.api_groups_rules.api_group` property

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

<a id="canonical-1213023212133200-3022323200001200-3323001303100130-0211232000303032-1300101233020312-3200233121313013-2111320303022211-2111000001230110"></a>

<a id="canonical-1301220303331332-3201212233303233-2310132020202222-2002213313323011-2321003211333310-3100333333300330-1332032323032202-2302010101032030"></a>

#### `api_protection_rules.api_groups_rules.base_path` property

Type: `"string"`. Optional.

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

- [client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-006.md#canonical-3122101313113213-0230122323313020-2213032333110001-0122003110111023-2020313332102030-1103032103222301-1211233231220112-0322000110200233): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012): complete subsection reference.

<a id="canonical-0100103332212103-0311112022302222-1323101121010112-1102230310310032-2303301121230013-3313302301220133-0122213311000011-0322010220213011"></a>

<a id="canonical-3312002021321320-1301032212230013-0033002003001000-1101232031100003-2313030011100033-1032111320033321-3233013201221230-0021202031000013"></a>

#### `api_protection_rules.api_groups_rules.specific_domain` property

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

<a id="canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.action

<a id="canonical-0103330233213001-3113331100312310-2211003030121312-0312032010131011-3312023331132233-0020011002220121-0303032100133333-3310000030233233"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310221332122111-0132331210022022-0220323322312133-0132121212102220-1122031000030023-2011303101123223-2100031001110211-2121211322212033"></a>

### Direct properties for `api_protection_rules.api_groups_rules.action`

- [allow](resources--http_loadbalancer--reference--group-006.md#canonical-2311201103303202-3333101313100311-0010213131331302-0101102133122112-1122303302033132-1102103121303302-0133232333303222-2020001130301001): complete subsection reference.

- [deny](resources--http_loadbalancer--reference--group-006.md#canonical-1011103131332301-0003222000202323-3000303101231300-1321103311102030-1212231223203111-3130301030303131-0320132010233301-3312103032031012): complete subsection reference.

<a id="canonical-2311201103303202-3333101313100311-0010213131331302-0101102133122112-1122303302033132-1102103121303302-0133232333303222-2020001130301001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action.allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-006.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- api_protection_rules.api_groups_rules.action.allow

<a id="canonical-2011133102202232-2123011010200010-1221110323001211-2223031320332333-3001122303213313-0020011013003213-3220110210001311-2222201113330301"></a>

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
allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011103131332301-0003222000202323-3000303101231300-1321103311102030-1212231223203111-3130301030303131-0320132010233301-3312103032031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.action.deny` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-006.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- api_protection_rules.api_groups_rules.action.deny

<a id="canonical-2112220020020000-1322322031302330-0320101312301300-3112133132120321-1012100123233202-2211122202021212-3312311300202100-2202003033031101"></a>

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
deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213023311121323-1113021222023213-1131132223030021-3332000100022232-1320331303331232-1113022133331012-0302020130233200-0320011212213120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.any_domain

<a id="canonical-2020023203023022-0310111202201011-0113311230231212-3012303312001121-0303130300030221-2311221300222031-2310120010231323-3001300012332201"></a>

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

<a id="canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.client_matcher

<a id="canonical-2211032301110320-3023013122132232-1121121302211111-3211012100030200-3110320302123030-3002012103201032-2323023231111323-0102211131001211"></a>

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

<a id="canonical-0332131301232311-3210100332010211-0330013101122321-1020212231010221-1310130012000320-3113020031011133-1021112332123030-2002203020302020"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-006.md#canonical-1300013322303103-1331110120302230-0102233212323100-2132323233020331-3002002303113302-1301031100120032-0212000131331212-0220333203032233): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-0320011203103032-3303133012331003-1021132330033103-0031233301102301-2322210323303312-2201100111231023-2121210332222222-2112120223122101): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-2321121232320223-0232023121030202-0011333130323002-1001123203010011-2102210000232100-1223013002232000-2010332002103322-1112322322201133): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-1330201310021011-1330101002323003-2232000112312001-2232012231022311-0101130333300112-0313203200300011-2212333322030032-3122131111200321): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-2003011313330030-1202302233223210-0012211101230003-0211300333033030-3320133032010330-0130020033330030-0233023230233212-2003221002130210): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-0123200231003121-2230301310320322-1132110230301000-3020301013222211-1010312322020020-1233321132113202-3231120013222233-2312310101110302): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0020130311132331-1120230211123201-2203233322212130-2002331231003300-0300110323300011-3231313303001130-2202322012200333-3212013201120332): complete subsection reference.

<a id="canonical-1300013322303103-1331110120302230-0102233212323100-2132323233020331-3002002303113302-1301031100120032-0212000131331212-0220333203032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.any_client

<a id="canonical-0013223200300220-1000031131103100-3012130322321223-1121103022030213-0122222013221022-0213300033021111-1212330211001223-2122222230133113"></a>

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

<a id="canonical-0320011203103032-3303133012331003-1021132330033103-0031233301102301-2322210323303312-2201100111231023-2121210332222222-2112120223122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.any_ip

<a id="canonical-3300010111221233-3221020022321032-1210222002120022-0221020203322032-2311322001320122-0301011130311303-0322103030221033-1211201222200210"></a>

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

<a id="canonical-2321121232320223-0232023121030202-0011333130323002-1001123203010011-2102210000232100-1223013002232000-2010332002103322-1112322322201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.asn_list

<a id="canonical-1112011222122203-2102223323033100-0232223132103302-2323101310011231-1212000102110102-2311103213202012-3202030012110031-2330323222300112"></a>

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

<a id="canonical-2010122203110001-2111321202203330-2133200012221000-0031121111101120-2123031220022031-1301332011123330-0321321021202230-3123321230300003"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_list`

<a id="canonical-2122103231223212-1322100213103030-3222230012110202-0112120110221213-0223121331130222-1233203322121203-1230201321331112-0011311300103332"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_list.as_numbers` property

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

<a id="canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="canonical-0020112121133213-2332121121202111-3303223212122130-1321012303120011-1121301010120302-0320011222211032-0202121031110220-0102032120131313"></a>

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

<a id="canonical-2222103132202300-1102200123301132-0013301320100103-2020312203213331-0213131032002303-3230231222302123-3110333020221333-1202203133122221"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-1222201100210311-0021323021012223-1003303023211321-3012200213021021-3220210211223312-2233003310203331-3003010313001312-1332222120003030): complete subsection reference.

<a id="canonical-1222201100210311-0021323021012223-1003303023211321-3012200213021021-3220210211223312-2233003310203331-3003010313001312-1332222120003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1210321202300300-1011212333331022-3303332223201223-0022001332223300-0131103223112201-0210232310111212-1021233103312002-1121121033031010"></a>

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231132321010201-3321121231120300-1212133222201013-0132123223113202-3031330030103222-0232231302103033-1330211121321222-3030200321010103"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3300202012102323-2121213001022000-0112130221213210-0000200032102203-2322132300231133-3000201101222003-0212211100021232-2023013312123032"></a>

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

<a id="canonical-2333323023003102-2133200120332232-0102112303023121-0132321321123231-3233210321003113-3230333313003231-1032023121013110-0033201321102022"></a>

<a id="canonical-2003112202000122-0010233033122332-2302213212130232-0232223113231201-1310320123111110-0231013032333113-1322212210231022-2000022312111211"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.name` property

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

<a id="canonical-3310300313010003-2033333003102021-1211122133012210-0201033121100333-1311100212301031-2212310132120300-0222131012332331-3103211201030330"></a>

<a id="canonical-0320113002001102-3102200131221322-1231222210033003-0111203200301321-1102122232223202-2231321210003113-0102000332032312-1332322230023233"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets.namespace` property

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

<a id="canonical-1301021213211310-2013331312320003-2123321321312330-3033002233120010-3332030213302133-2332122121010312-2131111003100022-1221111120021202"></a>

<a id="canonical-0103012312203100-1030211221120303-2013012020323031-0013203320302102-3312223132330131-3120303013230000-3002320332320131-2302211322311230"></a>

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

<a id="canonical-0323301301002000-2131212231332210-3313010232232101-3030200312101103-0002303330032230-0131020111002001-1002322002212100-0331122331112103"></a>

<a id="canonical-3131202200230003-0311102123220020-0021213233211200-1332313221312013-1231313033302302-3111113200231111-1333320013031202-2132222312120212"></a>

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

<a id="canonical-1330201310021011-1330101002323003-2232000112312001-2232012231022311-0101130333300112-0313203200300011-2212333322030032-3122131111200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.client_selector

<a id="canonical-2001022122331123-2322333311301203-3200231131103332-3103303133033210-1100032232033211-3010102330313310-3200022332111210-1132102102010122"></a>

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

<a id="canonical-3301223000300111-3210131000323201-2111002322113221-3303312012330201-1001303331231130-3011103331230122-2211033311211121-0131133122103031"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.client_selector`

<a id="canonical-2221202220103133-2120220103230301-3122311222333011-3021302032302032-2112113103122121-0231212012100213-0310232022131001-2013132131212103"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.client_selector.expressions` property

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

<a id="canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="canonical-3213320202222233-0021102220021023-1322031121030033-2120222321221131-3332122311201221-2212311223022203-2100322003111110-3132231010310311"></a>

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

<a id="canonical-3221323101223300-0230031131230010-0023130101300312-3130302312203001-2212230002321120-0022023100120202-3012023212130233-2313213301102111"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_matcher`

<a id="canonical-2001211002121330-1231232303023032-1100220220131200-1120302230200313-1013230111202100-0311230122201120-1121330103031321-0321232202002023"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.invert_matcher` property

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

- [prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-1323201312113011-1233011112101002-0321030213030330-0003000122121332-2321113010213302-2320200121220211-2313322310332300-0113303230111031): complete subsection reference.

<a id="canonical-1323201312113011-1233011112101002-0321030213030330-0003000122121332-2321113010213302-2320200121220211-2313322310332300-0113303230111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2123033322312121-0211001223301300-2310113312020003-2011311031312202-3223032332010033-2033001333330022-1201232320230303-2312120001133013"></a>

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120123002323131-2002131220020003-0211022113213011-0010021201020333-0010210320301102-0203000202002112-1021331223330211-1222102132132023"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-1301212202222011-0021022333122002-1320310300003131-1330120213220020-1323233010213213-3002000113330011-0310021213322331-0131033221122000"></a>

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

<a id="canonical-1333122023013112-0233111220113113-0101003231022020-2130223002320313-1011010120303133-1230310312033301-2133322330231112-2013232222032032"></a>

<a id="canonical-3103302013113120-1111013020320022-1133113203230323-2111200323131311-1201111032333033-0012201203320330-1021002132112313-1220131302013031"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-3111112232300310-3203100113013212-0100201321232223-0123023100032300-2210033032131323-1020233130120230-0211112103221201-0033101003212031"></a>

<a id="canonical-2023021321330213-0013120200230332-1311113133101003-2202211300120131-2131212132000010-3023123233132003-3321031222333232-2002230311120313"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

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

<a id="canonical-3303310321203201-2120202101332032-0113301333233132-1101330123103110-2301222133121321-0102301313022132-0221002111030232-1033032131121022"></a>

<a id="canonical-1002210233122023-3200101323111310-0322330013002322-1023102320332022-0212031111221001-2300020121232112-1122103303231330-3222213011221121"></a>

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

<a id="canonical-3003003030220112-0322101320212323-0213001002003120-0030013001032230-2012231202321021-1001202212100123-0002213220232001-3310232021303112"></a>

<a id="canonical-1113233030230323-1201330121323222-3302320302133332-3202101013113230-1302113333013333-2013212123031111-1332320122212113-0201011300020012"></a>

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

<a id="canonical-2003011313330030-1202302233223210-0012211101230003-0211300333033030-3320133032010330-0130020033330030-0233023230233212-2003221002130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list

<a id="canonical-2021030233002330-2320231023221001-3120303321222301-2123010033331011-2202032222111121-2310002132212231-1022013301133013-0323132212322012"></a>

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

<a id="canonical-3210121001210010-1210222011000221-2023132302133303-3331021322201202-2323021030012022-0120233313211212-0001022000112320-0101000202022002"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list`

<a id="canonical-3130232031223300-1010203301320322-3000300231210113-0210330303332210-3022111211033211-0110302112021123-1112113002133101-0122323330332233"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list.invert_match` property

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

<a id="canonical-3320111021313210-2001002210211002-3323122122321321-2033020333032310-3030232300332032-3101012120132002-3220123100320012-3011212130321223"></a>

<a id="canonical-0233031320021133-3130033112202013-0300221131032130-2210321132121331-0101123333321220-0220230322203230-3210113223012300-3303033101010131"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list.ip_prefixes` property

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

<a id="canonical-0123200231003121-2230301310320322-1132110230301000-3020301013222211-1010312322020020-1233321132113202-3231120013222233-2312310101110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list

<a id="canonical-3021103122302301-3223332333330312-1021223133323310-1123221312032002-2120033112013313-0100331332003030-0022311032320303-3330132021301330"></a>

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

<a id="canonical-1203220120100121-0000201313131211-1311323210012232-3323230320130333-2311230202101032-2002032111030203-2013223332002012-2002020200213320"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list`

<a id="canonical-1023101331302200-1013132331133332-2103020321032031-1322100211021123-0332220020010101-0310201213022003-1333120330101322-2301200132021210"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0020130311132331-1120230211123201-2203233322212130-2002331231003300-0300110323300011-3231313303001130-2202322012200333-3212013201120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3002123121211111-3100310330111301-1122010203333001-2133120112230322-2020003201021131-1233321133011302-1222022210311232-1330211330321231"></a>

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

<a id="canonical-3002000031313012-2032221133131220-3311230310123211-0332321223322312-2032210302311103-3113223333322231-3122330312002312-1201103110102222"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-2321123032230201-2121233131121132-1132020123101311-3301320330031010-0020330113300113-1210101033023111-3200210103033023-1210003000332031"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.classes` property

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1221213202111100-0031001330213111-2100333333103131-2332201113102131-3111020301030310-3211013313103020-3002231100320202-1223100020102032"></a>

<a id="canonical-2332231033021132-3330203000200120-2321112003033200-0113012323131110-2000111211222120-3102302311231332-0011220111003223-1111021002121123"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

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

<a id="canonical-0131311020233320-3301011020000333-1320222203233131-1100311023031131-3110031223131211-0230020013212330-0112002202100132-1112311123012321"></a>

<a id="canonical-3012131130322002-3202111000011231-2312002232302310-2110202322011303-3332103133111031-1120201312232212-1302302322131300-3313303302311133"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

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

<a id="canonical-3122101313113213-0230122323313020-2213032333110001-0122003110111023-2020313332102030-1103032103222301-1211233231220112-0322000110200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.metadata

<a id="canonical-3221120032313213-3302302130220213-3002323201001110-3330300103321210-3132321311213011-1232210203331303-3312012301311223-2102012231010012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021300102220020-0331030111112031-1103001330130001-3232221223210003-2122011012220233-1033110033133202-2222113110212331-2323031230021032"></a>

### Direct properties for `api_protection_rules.api_groups_rules.metadata`

<a id="canonical-1031230011103231-0131231213330220-0211120131013320-3032320230223003-1200301133030000-3321003221203002-3230023011103100-3202312023301032"></a>

#### `api_protection_rules.api_groups_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-0233302333330133-2102010103310301-1133100322031022-2330131332122013-2013230120001213-2223121131033130-2001210210220133-1102231331232130"></a>

<a id="canonical-0121302311222332-1132230031033032-3031001102303202-2200032130103131-2313012222112301-1103011302101223-1033233212233001-2233031323302111"></a>

#### `api_protection_rules.api_groups_rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.request_matcher

<a id="canonical-3202112031121121-3132032220312211-3332123222323102-0012021131101310-1312201213230111-2020002302300001-2222223020312001-3203002300232010"></a>

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

<a id="canonical-3233223022133213-3021020302121320-3100000322121303-0300223012200101-2132332101330121-0130100133312332-2103120231323330-3230320211332201"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-007.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-007.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221): complete subsection reference.

<a id="canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-3333031233221103-2021321121212022-1323030023001111-0102211033222212-2223311220201000-3030212222000130-3111003302033000-0032100210232221"></a>

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221333310301132-0312122121301122-2302131233132133-2103001133203233-0112030133301133-2301222101033030-3320130100321012-2313123320031201"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111003230333002-0132232133331110-1233131331121333-0010101301032030-3232112101303302-2000223001323313-1013020000111011-3310322011321003): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1031102332213300-1332223223321002-1100010010221113-3121200103330330-1322300120333322-1222100211300112-0000011022200100-2002101331011220): complete subsection reference.

<a id="canonical-1032110210113020-1213133303020332-3000122311032023-1300230112112210-0321321111012223-0033220200013311-1011021330213212-2121111233322213"></a>

<a id="canonical-0212332030001222-3002201332310131-0002210023001011-2130203122000031-2303012010113220-0110200230300312-3121122231301203-1212313203002130"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-3233103230131132-3310103322011132-3331212013102300-3133013122223210-3213130111220121-3123301321123213-3033033131202312-1103122132100332): complete subsection reference.

<a id="canonical-2320103133002300-3232003202120320-2102000331101202-2331011003321330-2031202333133220-0101110220133120-3300320311301032-1233320031201211"></a>

<a id="canonical-1201302200313013-0322112210133330-2122212223301200-0330331012312023-3322220232301002-3012113020233130-3111100033210023-1322020303123030"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.name` property

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

<a id="canonical-0111003230333002-0132232133331110-1233131331121333-0010101301032030-3232112101303302-2000223001323313-1013020000111011-3310322011321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-1312010102132022-0313321111231333-1221002001323230-1131312200231101-1301332223322111-1313102213213123-1121223030012321-0022032103320113"></a>

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

<a id="canonical-1031102332213300-1332223223321002-1100010010221113-3121200103330330-1322300120333322-1222100211300112-0000011022200100-2002101331011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-1012231221120102-3313120131231100-2000113301201213-2311321212311012-3120023313030331-1010023203231032-1000322000022002-0210032100002311"></a>

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

<a id="canonical-3233103230131132-3310103322011132-3331212013102300-3133013122223210-3213130111220121-3123301321123213-3033033131202312-1103122132100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-006.md#canonical-2330223010032220-2103331312331301-3013322330020123-1101220300222333-1023300323201133-0203302332130032-2200000120333102-1132221332331012)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-1033213322230012-3000111311012130-0113201110020312-2302130131233021-2102121102022102-1330311313323222-1010310223111301-0302320012003322"></a>

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

<a id="canonical-1312102221321021-2123032300101320-1230011013103233-0202122012001103-1210020210023123-3221130321123101-1020210331223100-3221213031103322"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item`

<a id="canonical-2231232212231013-3230223020200013-0301200233120221-3021311300133201-2320120221113131-1222033111231013-3333223031131120-0332211103130312"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.exact_values` property

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

<a id="canonical-1113233223001223-2103130223321220-1231323122231012-0011031013031303-1113331011201333-0211100203202100-1111220330210132-1211111300111130"></a>

<a id="canonical-2310122211332000-0210120330011332-0012301000312203-1011102122233001-1011122231222030-3333200120030033-3122101213131212-0221200320331300"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.regex_values` property

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

<a id="canonical-0302010123023133-1130023131211201-3033321133121322-1210220120113230-3013001121201132-3101200102013210-0230310013231313-0112212233303321"></a>

<a id="canonical-0120220131300331-1311210121301301-2010122021333030-1331201322230020-3110210020000030-1300223200303310-3022132130103021-2131331232320233"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item.transformers` property

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

<a id="canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-2321212010023130-2110111210222313-1030122320033202-2300220331101030-2331201310221331-2310223301312321-1203130202010210-0221021022131020"></a>

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012202231000211-2313323100102231-0101121210102012-3322022102231012-3101221332013123-3210022130000113-1322202323310220-1101222231000203"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0111031002110111-3222121311112200-2220301230201120-0020333233222101-3213030323222201-1101231321333300-1211332011132132-3023112320200120): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-3230303120123113-0020202120332113-2121230020230322-0230012020203032-0101130301203023-0010310023130213-2200331031010222-2202310221130030): complete subsection reference.

<a id="canonical-2231111002330320-0220322000302300-1021233230312332-3220030232022230-0301033323311010-0322013311210330-0001230222222103-0322211202013333"></a>

<a id="canonical-3201002011131331-2003330323111020-0000023122023132-3330322130310001-3120221302010223-1033301000320001-1001021022213101-1020213113012132"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-2112120021310311-3101013232321221-0123222010003011-0302201210223320-0033220030331032-2112220112130022-3221213001330030-3201023033003032): complete subsection reference.

<a id="canonical-0332301132133300-2322011000330331-2031111203113330-2130103101021112-0320333102033011-1313303232300113-0021110121221020-2233223311202102"></a>

<a id="canonical-0320012312330133-2121302220222102-3330013203231230-2303010021033232-1120310020001101-1031133221211303-0032220232211321-2133033333023203"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.name` property

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

<a id="canonical-0111031002110111-3222121311112200-2220301230201120-0020333233222101-3213030323222201-1101231321333300-1211332011132132-3023112320200120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-2022002231102101-2232202200110002-2232323121222302-1130110311200302-1121223103102023-3132322101303213-2200110023221220-3110012121300101"></a>

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

<a id="canonical-3230303120123113-0020202120332113-2121230020230322-0230012020203032-0101130301203023-0010310023130213-2200331031010222-2202310221130030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-1102212111033222-1132220212012232-2030233330211003-0230320100001230-1111311303312023-1331301323120313-2313333031202022-3320010020223122"></a>

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

<a id="canonical-2112120021310311-3101013232321221-0123222010003011-0302201210223320-0033220030331032-2112220112130022-3221213001330030-3201023033003032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_groups_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules.api_groups_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-006.md#canonical-1102113023001102-3003123110221310-3002231200030103-3302100303220112-2000122003203002-3100301023002011-0000112102303003-0030220020333230)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-2303131220031203-1201031120032212-3202133320101033-2012301012211311-2323220331330022-3201213311300130-3230102231232021-1303203312100030"></a>

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

<a id="canonical-3322030211002310-0223310312200023-2221230102212310-1200012003300101-3201331001231210-1202111322021212-1300210320301120-2122223013210003"></a>

### Direct properties for `api_protection_rules.api_groups_rules.request_matcher.headers.item`

<a id="canonical-2103121322010200-3002030011101211-2100011321211330-0121303310022101-0303010003313120-3102303203002121-0111221333010211-0321021212311010"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.exact_values` property

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

<a id="canonical-1122101233200020-3032201223120311-1132310220030100-2122011031030300-2002233130002231-2311212231201303-2020202103320032-3122130230020210"></a>

<a id="canonical-2013230320113012-0230212131130230-3303132212230003-0013130002130113-2121313202301011-2030230212213003-0313213010300220-2003231000231023"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.regex_values` property

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

<a id="canonical-3013001102233130-3102220111203120-2302130122211213-3010022220033322-0120030221312321-2033223021301312-0212223032120301-0101031100133003"></a>

<a id="canonical-2100221021211012-3100113213213123-2303312113133301-2230202122120320-2203003023001311-0020232130133200-2201122120301122-1200133113302002"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.headers.item.transformers` property

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
