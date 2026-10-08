---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2113110322302302-2000102202222022-1132023032031300-0000102100111122-0112112113103222-1210133210313131-0002012113212020-0133212011021302"></a>

## Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims`

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-0003011201021003-0213010021322221-3120231301210002-3021312233122213-1300323113122200-3010112330131222-2332331112213133-2330023000213022): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301): complete subsection reference.

<a id="canonical-0213131013102220-3023310302013302-3202300321321202-0332310103133033-3321110313013133-0110301223313023-0100122303113200-1202011211030323"></a>

<a id="canonical-0231221120003111-3331231203222120-3002103033001323-3231230131131333-0021210211020011-0320232001233120-1310233031301221-0100123303010022"></a>

### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.invert_matcher` property

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-0203322301301102-3102301233101300-1002313100310033-2121102301111000-0001321213001332-2221212100222123-2133100312013320-0300321121203300): complete subsection reference.

<a id="canonical-1202310110320132-0010033112012001-1303212111003101-1003223221131332-1133222300211303-3003123203300011-1123121230231303-3203123101330003"></a>

<a id="canonical-1010020201011112-0310033230210120-1311310212011221-0022113010021213-1322231111031330-3022133332210303-2203312210102130-1202120102022331"></a>

### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0003011201021003-0213010021322221-3120231301210002-3021312233122213-1300323113122200-3010112330131222-2332331112213133-2330023000213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3212032301221200-0200100100230321-2002120032231333-3100122231200121-0020323111121200-3030222203220012-1032022111113311-1200103232123020"></a>

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

<a id="canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2302001030010312-0323221100001213-1331113222231200-3101020203122023-3010332013100113-2311331303132210-0102303103310312-2013321321002111"></a>

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

<a id="canonical-0203322301301102-3102301233101300-1002313100310033-2121102301111000-0001321213001332-2221212100222123-2133100312013320-0300321121203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-3112231301211032-0033121231112203-0120100313132221-2113300303103220-1112230303101202-1310300212110202-0000320000102200-1211320223220031"></a>

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

<a id="canonical-3310231202311133-1023132300302102-2003003332003133-2130302231020322-1310330122232123-2023123131320230-1022322331100333-2120221021132200"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item`

<a id="canonical-3302311203010323-2223013023311300-3031333310310030-1031210122312312-0310112301320333-1220223120111323-3321021232120222-1202021320113001"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.exact_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0311133330223130-3302000313202231-2210113113223213-3321001313323300-0013302211011333-1213110220232033-3221111312001010-1132310213233020"></a>

<a id="canonical-2003321211011230-1233322203312331-3032113303122301-1222201121230132-0012221230201320-3212333230032220-2302222302021110-0310011112212212"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.regex_values` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2301311212300222-2310212312221233-1220032011200030-2203223132331010-3333213210132003-0232101131223010-0303033301010200-0212021322201211"></a>

<a id="canonical-0222110311021333-3223030312130231-3310021100012103-0312303330033033-0313312020021201-0200303312221202-0212002021213013-3101033321131203"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item.transformers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
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

<a id="canonical-2010122203110001-2111321202203330-2133200012221000-0031121111101120-2123031220022031-1301332011123330-0321321021202230-3123321230300003"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.asn_list`

<a id="canonical-2122103231223212-1322100213103030-3222230012110202-0112120110221213-0223121331130222-1233203322121203-1230201321331112-0011311300103332"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3301223000300111-3210131000323201-2111002322113221-3303312012330201-1001303331231130-3011103331230122-2211033311211121-0131133122103031"></a>

### Direct properties for `api_protection_rules.api_groups_rules.client_matcher.client_selector`

<a id="canonical-2221202220103133-2120220103230301-3122311222333011-3021302032302032-2112113103122121-0231212012100213-0310232022131001-2013132131212103"></a>

#### `api_protection_rules.api_groups_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0233302333330133-2102010103310301-1133100322031022-2330131332122013-2013230120001213-2223121131033130-2001210210220133-1102231331232130"></a>

<a id="canonical-0121302311222332-1132230031033032-3031001102303202-2200032130103131-2313012222112301-1103011302101223-1033233212233001-2233031323302111"></a>

#### `api_protection_rules.api_groups_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221): complete subsection reference.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-1221022321103011-2323212110312021-3301313311210313-2313023113311122-0220130312322000-3020003110301201-1211210030123330-1113023111131212): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-0112212123012302-2133330201021303-3203100112120223-2131131102302123-1330230312012001-2032230101300130-2113032230100012-3120131002213302): complete subsection reference.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-3221201132223132-0322013020201110-0020312210130003-2033223123010021-3301323211313200-2213110120203221-1033223223301323-3332323222322321): complete subsection reference.

<a id="canonical-3222132112301332-2201102220013330-2110030010211230-0223132213201322-2232301232112113-1213102101222321-1322233021100010-1203212201021021"></a>

<a id="canonical-2030102212200201-1300102302220131-1113123222012203-2130100220112122-2120301212131320-0303323231031111-2322230030001022-0131302213333103"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.jwt_claims.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
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
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
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
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-006.md#canonical-1000232202121331-0220321323200111-2211202201332332-0320003000223110-2322010033113302-2032133032000033-1020023231222310-0120011231031310)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [check_not_present](resources--http_loadbalancer--reference--group-006.md#canonical-2301301231130220-3022330120203122-3312122012121223-1011021100300311-0011212233230003-2230130122211123-0210310102112213-2332303323300312): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-006.md#canonical-2302200113210110-3323122133201313-2232122033020121-3210013030221001-3020120301002231-1133211000313033-0033131013321131-0211000020012301): complete subsection reference.

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

- [item](resources--http_loadbalancer--reference--group-006.md#canonical-0112023020121132-1033320220310300-0310233323221212-2021032103333032-0332013022221331-3111033112322011-1211013002322032-2223000113002332): complete subsection reference.

<a id="canonical-2032330003113313-1121003133303113-0110012011000311-0021033110102130-1133011003113213-3010211323001223-3202221212020122-0120223312303110"></a>

<a id="canonical-0120201132213330-1002100311223032-0122330121131002-1222321111332110-2222011331011233-3303323121223310-2120232131011000-3031002302300231"></a>

#### `api_protection_rules.api_groups_rules.request_matcher.query_params.key` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
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
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
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
- [api_protection_rules.api_groups_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3110002111102211-3332012221211002-3323110233102203-2312123130220123-2010223211321200-2013123123333003-2323220020311331-1322220031300221)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "custom_ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0233031301102130-1130230220030210-3300120230310332-2002022322233223-2023100032330031-0223303022120030-3300222102202013-2011031031103200)
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

- [api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220): complete subsection reference.

- [bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001): complete subsection reference.

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321): complete subsection reference.

- [server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103): complete subsection reference.

<a id="canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.api_endpoint_rules

<a id="canonical-1132131113333103-2022130102320022-2230110113330212-0132031113100120-1231101222210233-3302112221010212-0100032312311122-1323013233103113"></a>

Type: `"object"`. list nested block, Optional.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [any_domain](resources--http_loadbalancer--reference--group-006.md#canonical-1102221321330123-1112312231203321-2212320130100310-2020322220032221-2332230110121010-3112302312320212-3022321223202320-3021100213113332): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--reference--group-006.md#canonical-3211220233321223-0231100031032022-2012223000121203-1232010310013022-2210231321223033-3000133331210021-1023230032112210-2002333100210103): complete subsection reference.

<a id="canonical-2331022231313010-2113003022132122-2210231312301110-3003323021102313-0012011003003223-2002113203031230-1031120331212231-1312000113102132"></a>

<a id="canonical-3211003130301301-0121023210021232-3000222030233133-3211220112022120-2111022101231333-1010003012013000-1113013310331220-2303311012022203"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-1121233023131220-1003111220221011-0133310311113000-1021131213123031-3323030010330300-0200200213101021-1233030130032333-0013322100312031): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-007.md#canonical-3031121110211321-0222223200032133-1223103211312023-1012222201100010-0323202231130000-2323201330321110-1222232311021302-3202202311112033): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-2203030010313000-3130133321212122-1030213230333002-0013320321313202-2132022332230312-1233032110003311-1012112113123023-3202301300202122): complete subsection reference.

<a id="canonical-3203121311310230-2100333322210103-3231301201012120-2322101111331110-0113333001101211-1133203230213233-3123210021321023-0022012232220120"></a>

<a id="canonical-0210130313110330-2130103222121211-1001230012003222-0223300221112033-2231113120231323-1322012221123130-0213320101321330-1131301223232001"></a>

#### `api_rate_limit.api_endpoint_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-3102022023020021-3321211122201030-1330012333122230-1002010222001122-2201312311031222-0131232321011013-3232103110111321-2000223122010220"></a>

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

<a id="canonical-0310121103222333-1321203010112201-1320031000022231-1331230311002231-2013331001000012-1003103212120331-1200301000323110-3223122321220311"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-006.md#canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-006.md#canonical-3111231132022032-2012201213130111-3210122323121103-2130110002110211-0003213133322113-0102101110002231-1020231322230313-0033122012310213): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-006.md#canonical-0231131102210101-2133013103213130-0321113113303011-0210321100230110-0212110222110300-1203311112312300-3030332313100211-3001132332112302): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-006.md#canonical-3100320111032122-2112302021131110-2213020113220103-2123310222003200-1320103210121220-2203021031111200-0013233111012101-0233131002202122): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-006.md#canonical-1312202323020121-2330122112103213-2102331212330331-0311311101121312-0331221200021133-2111102011123223-0111002302103312-2030312132201000): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-006.md#canonical-1231130011002201-0201220210132022-1202011011322320-3222221230011211-1220312323300101-1113033232232300-1300130111323030-0113012032033021): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3201202030331112-3102102112311122-2223110300113033-2231001222303211-3001320300131233-3202223231022132-1213133331012100-3122020302033123): complete subsection reference.

<a id="canonical-0012221131313133-3123300110213130-3211112200313130-3000312112230301-2022023132233213-1201000232011321-2132222323102032-2211011231100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0003333232231212-1232330320133012-3232102103300113-2113000113003231-1200203303123002-0103031223033233-2231230101310100-0203002312213022"></a>

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

<a id="canonical-2030113210101323-2313111130210122-3300013220333110-2301213131033001-1130201320212210-0230000302321220-2232011332130123-1320003023010202"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-1200121330130020-1002313201331323-2112331223232002-0230311131021021-0222333323333301-1211022100310010-3032000202221212-3032312102200032"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2033113121211121-1130333112313221-3322012121000020-2302223201001213-0221013300130330-1333232121320320-0200101303301322-1303321123000000"></a>

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

<a id="canonical-3311032022131202-0101110000131111-1123003120311233-3200303103033031-0120322302011013-0300013131232033-0301120300222002-2200021221231022"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-006.md#canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323): complete subsection reference.

<a id="canonical-1200210320110211-0310131300222120-0012333132101230-3223331322111202-3300201212002110-0201012210223030-0202323130200002-1013322023301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-2133101103210301-1133300211033320-2220320131022112-0003022120230021-0102310230123222-0002201330021303-0302221113020123-0331103300033123)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
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

<a id="canonical-3102200230130010-2330010033303202-1330103103102131-2031032112231221-2301023130321022-1311111111100100-1331111201110212-3333212223331231"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-3100120202202223-2202313020231022-0033330123012130-2122003320002323-3230212123223301-3013003230222202-0020100303101130-2121312302122200"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.client_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-3321023330020202-0120001131222133-0003103331110231-0310221223223232-3220020303120122-2112200130231330-3301301323201203-0210032120113320"></a>

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

- [prefix_sets](resources--http_loadbalancer--reference--group-006.md#canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132): complete subsection reference.

<a id="canonical-2023121332212302-0021301111000301-3022222231220000-2210212020101232-3233312313100222-2210001112112231-1012213133000313-0310210223320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-3200322322220212-1113313312121132-0022020102332110-0003210012123332-0232201320020023-1331003101031110-1133111233110332-1032321212200130)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-0001121323103103-3300130121010021-0201000132233033-2000002010301303-1003323220132001-2121303301200131-0132000231200131-1310222100032211"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.api_endpoint_rules](resources--http_loadbalancer--reference--group-006.md#canonical-3212223132302311-3211023110012233-2310011110223112-2022032200210011-3212011230013320-3210011321120002-2230312110002121-3111300202333220)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0111201233103210-1210300122220103-1312220133231331-2022313310110133-2102101200102131-1032211032111002-2103033032111333-2222200312321022)
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
