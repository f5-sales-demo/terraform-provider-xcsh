---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3213131012011200-0231333010203210-3020113112001211-0210102213010120-0020332003333221-0010310023332321-1022110032022222-0300032020303201"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3321230312101012-0332103111301101-2103003032200222-0100133001233223-2110332331120012-2210311310302100-1000133120332130-3130000132111110"></a>

<a id="canonical-3302202332103120-3110331020203012-3330301213123020-3223323223103333-1203112112113023-1323023311211031-0012301310022113-2200001010033123"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1021121300320102-0100211112110100-3231211212102323-3122223231031102-0311311101233202-3012302203102021-1210313110223131-2331031302110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-025.md#canonical-3203002113012103-0331110000131012-1032002320020112-1220220322320001-2310222100221303-2303112033001330-2300102120103132-3033222030222111)
- [routes.simple_route.advanced_options.response_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-025.md#canonical-0312203022301201-3021320100230120-2201200300313122-3200021120121033-2031213033211101-3000013011102303-3230211320312320-2333311023333322)
- routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0200210212223323-3232301203212102-1030133210130302-1011311011003333-1120131220203011-1032122033330102-0312210222101230-1133120203031113"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2212010212210201-0022223010210021-0310012303102031-1232121010331002-3102112310120113-0010322002320310-2220220210000330-3113101132221333"></a>

### Direct properties for `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0313222212302111-3231011132031031-0323130033213331-0230100100132312-0330100201132202-2032001203322020-1220010101333121-3002230023011301"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0310311230320200-1100232003010101-2210032333300113-3010100001103031-2222001001301320-1220203222032223-2012230111113030-0200003232103310"></a>

<a id="canonical-1032112001223030-1211020212202002-1202310312101300-1301123223322302-3301331010323233-2122012003332333-2111013110310030-0012021201002223"></a>

#### `routes.simple_route.advanced_options.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1202213103230100-0130021013021303-3333333123022220-3221102231223031-1100222100120321-1231101313331303-1012212321223322-1311032132122310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retract_cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- routes.simple_route.advanced_options.retract_cluster

<a id="canonical-3110221120110203-2233121103120220-1111200323230112-0121131321201002-2113331023221011-3123230003000212-3331211012202013-0102320002031222"></a>

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

<a id="canonical-3010333000322231-1332121313210110-1032031303020122-3011002010100021-3130012013120201-1002131312131023-0232133122202300-0033100331231300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retry_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- routes.simple_route.advanced_options.retry_policy

<a id="canonical-1033120002123312-2002111132232131-2210321200331201-3221321122002111-2331021323011310-0211133012321303-0310120212122220-1121332130300120"></a>

Type: `"single"`. Computed.

Retry policy configuration for route destination.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211311302320113-0333221120112000-0101230220113010-3023112003120202-2311331111011200-0202020320132221-2011213002230010-1223230233003131"></a>

### Direct properties for `routes.simple_route.advanced_options.retry_policy`

- [back_off](data-sources--http_loadbalancer--reference--group-026.md#canonical-0101311201212012-0132301311031020-0301323222111111-0023001032001322-0200120110320122-2311120023032101-0303231201313012-2232231323200132): complete subsection reference.

<a id="canonical-3200310021013313-1230203200122131-0023201212103002-3320222301320032-1131100013313310-2122033300320010-1223313312021113-0021213001110132"></a>

<a id="canonical-2001010311333120-2101013132313231-3311303323231210-2102003213203020-1133212323301213-1130223321203010-3121020130021020-0300030122232121"></a>

#### `routes.simple_route.advanced_options.retry_policy.num_retries` property

Type: `"number"`. Computed.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-1122112123332012-3231322011132203-3312322202313110-0032203320301310-1000033123023001-3200303112232112-2222110312110213-2123311032313131"></a>

<a id="canonical-0110210303202000-3022322122023001-1101020100000112-1031233012023030-0030111121013321-3122122113102231-3213023221203201-3313323302213013"></a>

#### `routes.simple_route.advanced_options.retry_policy.per_try_timeout` property

Type: `"number"`. Computed.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0112003100122310-1320300202202330-2201012132330201-0032033111323130-1110132133020031-3033200323322102-2023030333120322-3121222110031121"></a>

<a id="canonical-3322121001333210-0133201033300310-3230110230321223-3133330203032101-0130103111111202-1100021210322020-0020332200022101-2231102203221330"></a>

#### `routes.simple_route.advanced_options.retry_policy.retriable_status_codes` property

Type: `["list", "number"]`. Computed.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

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

<a id="canonical-2203200230130011-0300301013130302-1120020210110012-1320001022301331-1123100211102100-0203330303203300-1330130010013130-2210132333233200"></a>

<a id="canonical-3121013331022200-1211122113231003-3122231321211330-2322010300213101-2001210120232133-1033101333221311-0120001212200002-0222203120022023"></a>

#### `routes.simple_route.advanced_options.retry_policy.retry_condition` property

Type: `["list", "string"]`. Computed.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Additional upstream details:

For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0101311201212012-0132301311031020-0301323222111111-0023001032001322-0200120110320122-2311120023032101-0303231201313012-2232231323200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.retry_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-3010333000322231-1332121313210110-1032031303020122-3011002010100021-3130012013120201-1002131312131023-0232133122202300-0033100331231300)
- routes.simple_route.advanced_options.retry_policy.back_off

<a id="canonical-2331100323220332-2223321210123110-0110102013132312-0133030333122100-3230123010031201-0312220221321100-1030230313320301-2203312302230330"></a>

Type: `"single"`. Computed.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230131101033310-0003031032100112-1030322203210103-3200331012211100-0230231133202333-0232002011322332-2333302210202332-3303121122333323"></a>

### Direct properties for `routes.simple_route.advanced_options.retry_policy.back_off`

<a id="canonical-1202033232112013-0012330101231013-1020033302322132-2231102210012313-3012212110220011-0112330130012101-0020131210223103-1113313211321111"></a>

#### `routes.simple_route.advanced_options.retry_policy.back_off.base_interval` property

Type: `"number"`. Computed.

Specifies the base interval between retries in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-1033223013312301-1023323011110202-3321123313031330-3022130301210200-3203323021310302-1232200111122123-3300302200002012-0303331021313202"></a>

<a id="canonical-0030013320003303-0020213101302301-3011202213200010-2203313133313103-1323311113200012-2120010333221321-2230212022100020-1320302320032233"></a>

#### `routes.simple_route.advanced_options.retry_policy.back_off.max_interval` property

Type: `"number"`. Computed.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- routes.simple_route.advanced_options.specific_hash_policy

<a id="canonical-0013331021002312-1003002312013110-0011011212113221-0003301123332223-0032212212211320-0000013302330133-3230020223121001-2211233033311233"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Additional upstream details:

List of hash policy rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3023003010130012-0023122322301223-0112301311301112-3231311312121331-1001302211101202-2000100322001123-0112312102131313-0203020102122002"></a>

### Direct properties for `routes.simple_route.advanced_options.specific_hash_policy`

- [hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010): complete subsection reference.

<a id="canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy

<a id="canonical-3002122322220111-0210003201122023-3013112212312030-0221100130213211-3310211201103033-3122113101003231-1312020303100000-1121132033112133"></a>

Type: `"list"`. Computed.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221100020113120-2111201023103211-2120003322122131-3323230121321022-2323322300331223-0013112300310132-3133301320123233-1002033230100322"></a>

### Direct properties for `routes.simple_route.advanced_options.specific_hash_policy.hash_policy`

- [cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102): complete subsection reference.

<a id="canonical-3330312120302321-3311123330033002-3102333312311000-0220021100111003-2230221101010233-2103103012333013-0330002023002300-0021132232132031"></a>

<a id="canonical-2031232203200302-3212021113020300-2230100221230112-1333131201300312-3103231200101022-3320012020232311-1131020323303122-2311332122001313"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.header_name` property

Type: `"string"`. Computed.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0322333110231311-2231112311021332-3300001123033322-1100023102111311-0302231130021020-3200222231013303-2111102200321222-2222333330021231"></a>

<a id="canonical-0111020203321000-2202032233233223-3333322130321101-1300333222331212-1001132100211020-1133232331012213-0022300223313221-0103231303013310"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.source_ip` property

Type: `"bool"`. Computed.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303103210311200-1032233330021201-2330203301300010-0331312231213310-1223202210013323-1211300011030100-0132130223100222-0232231110322020"></a>

<a id="canonical-0103121230322002-1022120131020010-3322030110221131-1110313022020023-1201000003331313-2313120301132121-0330221311321113-1200210001213100"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.terminal` property

Type: `"bool"`. Computed.

Terminal. Specify if its a terminal policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie

<a id="canonical-0211022223133232-2122303301111333-0023012102123330-0313210202201233-1323000211130222-3221111032131322-0100121310210202-0110030221102300"></a>

Type: `"single"`. Computed.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

<a id="canonical-3101120200303031-0033033332233000-0332213122301222-3201001101133310-0210003112132102-2111100020000220-0031002323013313-3111232103333002"></a>

### Direct properties for `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie`

- [add_httponly](data-sources--http_loadbalancer--reference--group-026.md#canonical-1123301132230010-0132120032320031-2103123100310031-3312312220211303-0211331132212332-3101130221031003-1003131313023033-1031023122322211): complete subsection reference.

- [add_secure](data-sources--http_loadbalancer--reference--group-026.md#canonical-0300303310212302-3323100311200331-1000101103133101-0201000010031001-3313321230210002-2331111002223020-3211020120313331-3012113021001023): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-026.md#canonical-3000100011232033-0013332303310231-1032223032123330-3302010320223231-0032012020101011-2203322003202101-0221011212002000-1120133230110321): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-026.md#canonical-2030303133213031-1022033102032021-3000013032210113-2113002131302003-3123001201321130-3122000110130333-3201232223323313-2100130101202213): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-026.md#canonical-2212013122223101-3231122333211011-0323220033112113-2222133302020222-3310103202202203-0022223112231321-1010101033132112-3021002112030112): complete subsection reference.

<a id="canonical-1231211131300101-2330221302003202-0010312322220231-0020323012321233-2201133002233031-0201333231311120-3033023122232330-1303110213102022"></a>

<a id="canonical-3210313021220311-1310320122131323-1323311010113303-1101230330022122-0010032322231323-2102202332030030-1002213132011312-3010030103223310"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.name` property

Type: `"string"`. Computed.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0212233003222312-2120201213133202-3123321230202310-2013013313021302-2130132213103032-1301002223211020-0132323311122031-2330330113103030"></a>

<a id="canonical-0112201122013020-3110203321022123-0311223203311300-3313230222231201-0132230310120333-1203202011131100-3222233230312302-1110310320210202"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.path` property

Type: `"string"`. Computed.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](data-sources--http_loadbalancer--reference--group-026.md#canonical-2002120333200013-0013131133202321-1100002210022010-3202120002022010-2111332322133301-1220310120001213-2233222113303131-2301302031100010): complete subsection reference.

- [samesite_none](data-sources--http_loadbalancer--reference--group-026.md#canonical-0323113120313133-2333030132023321-3210202233202313-1201021231331202-2332300301123313-3100223303021001-3023322021230220-1022120233110302): complete subsection reference.

- [samesite_strict](data-sources--http_loadbalancer--reference--group-026.md#canonical-2113002332233101-1102110012210010-0020223122203200-2023001000313201-0010211032330211-1321020132002100-3230030013011003-2020131031111331): complete subsection reference.

<a id="canonical-2001030300320212-2320121312010232-2021011023310021-1323221210323032-0121321233230111-1222012221120311-2133030313001233-2131221221312113"></a>

<a id="canonical-2023303313213313-3130003011010213-1213022203132120-2023110201111311-2022033330220210-1202001320003011-2110332021030223-1300203312133223"></a>

#### `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ttl` property

Type: `"number"`. Computed.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123301132230010-0132120032320031-2103123100310031-3312312220211303-0211331132212332-3101130221031003-1003131313023033-1031023122322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_httponly

<a id="canonical-2333111012012032-0012111222221230-3331032211121223-3203101222221100-2110230313312202-3300120312310222-1113102331322223-1213012112230002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-0300303310212302-3323100311200331-1000101103133101-0201000010031001-3313321230210002-2331111002223020-3211020120313331-3012113021001023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.add_secure

<a id="canonical-2131220111102010-3001113111203322-3030303202122032-1012332133200331-2312210322313232-1011112330311113-2311221232001233-1110232111301300"></a>

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

<a id="canonical-3000100011232033-0013332303310231-1032223032123330-3302010320223231-0032012020101011-2203322003202101-0221011212002000-1120133230110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_httponly

<a id="canonical-2021231212023211-2110203133220213-1031011203332223-0000321321332211-0031101230301000-0312100132032200-3031221033021310-2200203233032010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-2030303133213031-1022033102032021-3000013032210113-2113002131302003-3123001201321130-3122000110130333-3201232223323313-2100130101202213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_samesite

<a id="canonical-3022211332023302-1000232133030210-2013230022103013-1333211100212313-2203110121120303-0020223221211020-1213120331131103-1011203210103101"></a>

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

<a id="canonical-2212013122223101-3231122333211011-0323220033112113-2222133302020222-3310103202202203-0022223112231321-1010101033132112-3021002112030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.ignore_secure

<a id="canonical-0333022303302101-3132123033311112-1300001120300322-0121130230100013-2101203011101001-1111123321011331-0232231130111133-1121301121200032"></a>

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

<a id="canonical-2002120333200013-0013131133202321-1100002210022010-3202120002022010-2111332322133301-1220310120001213-2233222113303131-2301302031100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_lax

<a id="canonical-1120033333321102-0202213333220021-0031203320132113-0332010323121033-2002332003222121-0102303033333111-1001112032120120-0113102232232131"></a>

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

<a id="canonical-0323113120313133-2333030132023321-3210202233202313-1201021231331202-2332300301123313-3100223303021001-3023322021230220-1022120233110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_none

<a id="canonical-3033222032111130-1231230230230230-1102201201000022-2022222132333012-1223101322101030-2213313323331301-2221231233322301-3232001011013230"></a>

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

<a id="canonical-2113002332233101-1102110012210010-0020223122203200-2023001000313201-0010211032330211-1321020132002100-3230030013011003-2020131031111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0102213113001301-2132232011232310-2301301201331013-0033113222132023-2230311110230131-2212002213230233-0131303133313113-3102312200120103)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-0113102013212231-3013322211230211-3333332000233321-3032232210112001-0223011222212313-2220102302131333-2231232232033132-1331011331133010)
- [routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie](data-sources--http_loadbalancer--reference--group-026.md#canonical-0201113312002012-1011223323113103-3020312320221212-3200001031223000-3311120021231121-0211000101013231-2132320023102103-1103032311323102)
- routes.simple_route.advanced_options.specific_hash_policy.hash_policy.cookie.samesite_strict

<a id="canonical-3222203221020302-3130323000323233-2032001020122310-0111021110130331-3032123000233223-2201323232133020-0003111131030232-1210001120231031"></a>

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

<a id="canonical-0232020012220210-2131122322331311-1233003120122212-1022232102000111-1213330333301002-1011030121012203-1321112203003132-3220321033222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- routes.simple_route.advanced_options.waf_exclusion_policy

<a id="canonical-2122311030001221-3231010131131232-1220200333310132-0322010002113122-3000103000320233-1333031110203223-1203211211003103-0123233221321013"></a>

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

<a id="canonical-3200031112002320-0221310322001312-1030002101130120-1223201120303122-2201131012101223-2020212002130221-1331020113010112-3000222112133110"></a>

### Direct properties for `routes.simple_route.advanced_options.waf_exclusion_policy`

<a id="canonical-3133202303200320-3103020022333312-3332013232010031-0001032202101010-2220001000303311-1222311101320021-1320331322332132-3323320320230003"></a>

#### `routes.simple_route.advanced_options.waf_exclusion_policy.name` property

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

<a id="canonical-1111023101212000-1001033013013303-2203002031023003-0310202013211002-3011211300100332-3321103230303031-2110010131220202-0301130001132100"></a>

<a id="canonical-0112302030331121-2100200013313030-0111032201130012-3222302112121230-1232120203201020-2010100223223023-3220133023331220-0100123102133203"></a>

#### `routes.simple_route.advanced_options.waf_exclusion_policy.namespace` property

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

<a id="canonical-1003213121333303-2310323311031000-3230311321102213-0202020120323320-3331030223120201-3113331102002100-3322332321020123-3122032201302313"></a>

<a id="canonical-2310030302033111-3001112320121202-2032120030103123-2132032020002013-2202022121203031-3130331000100332-3332322021300022-2312201332031230"></a>

#### `routes.simple_route.advanced_options.waf_exclusion_policy.tenant` property

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

<a id="canonical-3122203013233322-1210003000003103-2130312222213310-3023233211321112-2212212021300103-2210113220210333-3333320101033322-2003103001132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.web_socket_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-025.md#canonical-2103300321231223-3313311003122203-3310320230321101-3210221022312202-1113110131120131-3210203313221100-2202300131232101-0111320332221033)
- routes.simple_route.advanced_options.web_socket_config

<a id="canonical-3030202000013111-2223101210313102-2122133203322232-1303312003333211-3220102213331110-1011122110332202-1230221333223212-1330132321102312"></a>

Type: `"single"`. Computed.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3212103312213310-1010332221010023-0230323310030032-1220010021233111-2021333333122321-1301221032122130-1222200221031223-0102303221311222"></a>

### Direct properties for `routes.simple_route.advanced_options.web_socket_config`

<a id="canonical-0001001122130121-2330223310223203-2221232333332231-2010301022033100-1002111133101321-0203302011300103-0030322221221200-3201220231230000"></a>

#### `routes.simple_route.advanced_options.web_socket_config.use_websocket` property

Type: `"bool"`. Computed.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0331102102230323-3130302213120331-1200231110320102-3002023220020221-2111211231201100-0023220000221000-0113201001132021-2012120011223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.auto_host_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.auto_host_rewrite

<a id="canonical-2122100312200230-0330123002320321-2000101133313102-3211110003211323-3223000020033100-1110202310022310-2311132022223032-0113021200222121"></a>

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

<a id="canonical-1121020232022113-1322033113012323-0302022300002110-1230102303002111-0233231333313100-0121321301012231-0320333200122110-3233102022221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.caching_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.caching_disable

<a id="canonical-1321013131101210-2020131032232303-2232130012123220-2123011013032332-1012313103011002-0233312122233303-3130212032132121-2013101032231020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching disable.

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

<a id="canonical-1320003112002213-0231032121033323-0320310001322232-1232123000331010-1123222302232210-0031331330010020-1121100100102113-2012212311333032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.caching_inherit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.caching_inherit

<a id="canonical-1010231203310203-3120232321322231-3013322320012221-0300320300233110-1202322033221232-1201111232203320-1002201020302023-3233310203122323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for caching inherit.

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

<a id="canonical-2323121023310332-1023222300110202-3113030001323113-3021233110020010-3212232210031100-1120032211310333-3132203231221323-0232222221022310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.disable_host_rewrite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.disable_host_rewrite

<a id="canonical-3120230221302203-3333002301102210-0123232222001223-3222231120333032-0313232100310313-1102120131321121-0130032002030311-3302203201001002"></a>

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

<a id="canonical-3132021111231011-1213220233020212-2221133003002012-1223202120310012-3321200330113130-0123012133021312-2201300122111202-0121311211131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.headers

<a id="canonical-3103011103030130-3331301121223213-3131203330011000-2301002030031232-0310021301031322-2123100021112131-1110030010201203-2020112301122132"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

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

<a id="canonical-0100101213221111-1101210033231330-2233013033232220-1033111012231121-3101310020101200-0100103012113302-2211001103223230-1010202220212321"></a>

### Direct properties for `routes.simple_route.headers`

<a id="canonical-2211123133113013-0313323321101011-1320033132333210-0131112101001232-2300101023101223-3020020223130322-3203020002333332-0210021020111011"></a>

#### `routes.simple_route.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3332200003132220-2033001233312123-3303033112001010-3013022230122303-2200323121210101-3121121021233321-3102223003222333-0232010102103310"></a>

<a id="canonical-0322132121101223-1102230213120232-0032203022122111-2313011310301212-0112200223321001-0100100102101200-1112022331032211-3232210333301233"></a>

#### `routes.simple_route.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121201221130132-1113321201311201-0231330001330332-1123210120313233-2231333023101221-0330300313110202-2301222311033132-0103010210213212"></a>

<a id="canonical-1320012002210312-3110200302202331-2322031010320311-2113021220202100-0223110033033311-2232122102300313-0001330122113332-3022232213103232"></a>

#### `routes.simple_route.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3212021002213112-2033013331132000-1200113131033222-1323032201233133-3321013330312110-0110012330130022-1312223303003002-3000311130321023"></a>

<a id="canonical-1023011112332200-3101112001213211-0133003023212333-3112021023020002-0230213310210100-2002320303013230-2023302201113303-0202011123010210"></a>

#### `routes.simple_route.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2033130213131131-0231222211132013-0223212110110132-1210103213102300-1112003122301020-3001231331101011-1022233101220131-1332320130332112"></a>

<a id="canonical-2110231233331123-2122103312313313-0333233122133312-0233203213112010-2303120210101002-3121321200030323-0122200311033232-0312331331022021"></a>

#### `routes.simple_route.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0221302322012013-2300012121111101-1301101000322101-2020313032312332-1113101210233201-0102020310330001-1100100122322201-0311331230031133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.incoming_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.incoming_port

<a id="canonical-1032323302023002-2332303303320101-1220202223330021-1201001000333101-0023113000103032-3303023123312221-2201332330231031-0103132001133322"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-0301120123113323-3301132011210102-2312102100010323-0100333032132300-3202213002120110-2221220112132302-1121022302011133-2331233030021010"></a>

### Direct properties for `routes.simple_route.incoming_port`

- [no_port_match](data-sources--http_loadbalancer--reference--group-026.md#canonical-1030021200212010-0012031012122021-0130110102032032-0000002130030112-2303221022231111-1130303332132003-1320113310131101-1221313103200110): complete subsection reference.

<a id="canonical-2031302113320321-0310132213233021-3320013223200001-2003321010120012-3010122312313132-3230121320221311-2000230201123232-1211032122032011"></a>

<a id="canonical-1311300200023023-2021301212301200-3130202102111311-0102031311021103-0030201132200230-0330202120330222-3201120001202302-2230300001012020"></a>

#### `routes.simple_route.incoming_port.port` property

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3120220310001321-1123120010123021-2011100220221321-1000131302120210-1202200202221212-0011002113023100-3030031121330210-1013221031131011"></a>

<a id="canonical-1233320001311312-3012122011201333-0322310101021111-2220313303001031-3201321130021033-2020103003303113-3221003210211032-3231133212322303"></a>

#### `routes.simple_route.incoming_port.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-1030021200212010-0012031012122021-0130110102032032-0000002130030112-2303221022231111-1130303332132003-1320113310131101-1221313103200110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.incoming_port](data-sources--http_loadbalancer--reference--group-026.md#canonical-0221302322012013-2300012121111101-1301101000322101-2020313032312332-1113101210233201-0102020310330001-1100100122322201-0311331230031133)
- routes.simple_route.incoming_port.no_port_match

<a id="canonical-3333032101020202-0332330303131013-3201130331330200-1030132202331212-0323313211000220-3103221001230000-2321333120123102-0320020103103002"></a>

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

<a id="canonical-0003132220301321-1220330033212333-2222322002322300-0313013021302201-2130003013111320-1313000211100023-1031121223320111-2301123233231131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.origin_pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.origin_pools

<a id="canonical-3313113331012001-0310110311320320-1313021020221220-2023211332220033-3020032213200010-2100310112123232-1202331231130132-3331301332023203"></a>

Type: `"list"`. Computed.

Origin Pools. Origin Pools for this route.

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

<a id="canonical-1230103310023230-1202213212121200-2103223320313332-3111223032111120-3313022130102100-0003302232213123-3311330332113203-0321030012312003"></a>

### Direct properties for `routes.simple_route.origin_pools`

- [cluster](data-sources--http_loadbalancer--reference--group-026.md#canonical-3230220212001003-2120010321100213-2031330121002102-0210231200221000-3031302302131321-2310000310122003-2132132012201010-2311021001321122): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-026.md#canonical-2122010200001030-0131322020123030-0311202111013203-2220311130232023-0023032331033332-0231023310011001-3212112301222032-0232202220331203): complete subsection reference.

- [pool](data-sources--http_loadbalancer--reference--group-026.md#canonical-1010331333023020-1030001323032332-1212311011100133-2101122100222210-0211220123033310-3131301112131001-0031011211200201-3311210132300303): complete subsection reference.

<a id="canonical-3022011100301203-3133323032110301-0002133202121300-0311032111312321-0232020113210211-1232101230000111-3232202012321103-0121111021331110"></a>

<a id="canonical-2212132010232210-1221123213323230-0011110112101130-3333010213122111-1232333303123333-0111020303213010-0323001121130130-1030113033232023"></a>

#### `routes.simple_route.origin_pools.priority` property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0222230122010313-2222310013013331-0200221013332022-2211002110122001-2012102201320003-1213322112321333-1002030221001302-1023322313130333"></a>

<a id="canonical-0311322210113201-3001213113132001-3122300203001212-3022221001020022-3121011001110300-2123020132313102-2200130021212312-0132002332322032"></a>

#### `routes.simple_route.origin_pools.weight` property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3230220212001003-2120010321100213-2031330121002102-0210231200221000-3031302302131321-2310000310122003-2132132012201010-2311021001321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.origin_pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-026.md#canonical-0003132220301321-1220330033212333-2222322002322300-0313013021302201-2130003013111320-1313000211100023-1031121223320111-2301123233231131)
- routes.simple_route.origin_pools.cluster

<a id="canonical-0331323013133233-1131212102021130-0301332133300233-2310200300313122-2131030033032011-1330211331002311-0132201013123313-3200201330133330"></a>

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

<a id="canonical-1203233201211230-0123030023131023-0313102101003013-3030323321112203-0023111133233213-1112301021320000-1121031121033110-1030113331200000"></a>

### Direct properties for `routes.simple_route.origin_pools.cluster`

<a id="canonical-0211111220211133-3000212332112212-0301012202210333-2213320133001102-2223033233222323-2211101130102110-3313121021010301-1200231110012321"></a>

#### `routes.simple_route.origin_pools.cluster.name` property

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

<a id="canonical-2220220331233211-3312333311132321-3312033213132111-1132311302100203-3301310230111130-2001030102222201-1011322111032300-2313312100011110"></a>

<a id="canonical-1230021232302222-3003023120222200-2102303323213123-3021212321211011-3133120323031112-3022003010103211-2132133003302303-0132313013332313"></a>

#### `routes.simple_route.origin_pools.cluster.namespace` property

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

<a id="canonical-3030210232300330-3010323133112211-2310133203201202-2222002302123121-2012122331030222-3223103103223010-3202302203221202-3300112122320200"></a>

<a id="canonical-2221120231122101-3033131132030330-1113110032211333-0020011122031103-1020010322310331-0313131113212213-0002221112122321-0032303111120102"></a>

#### `routes.simple_route.origin_pools.cluster.tenant` property

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

<a id="canonical-2122010200001030-0131322020123030-0311202111013203-2220311130232023-0023032331033332-0231023310011001-3212112301222032-0232202220331203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.origin_pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-026.md#canonical-0003132220301321-1220330033212333-2222322002322300-0313013021302201-2130003013111320-1313000211100023-1031121223320111-2301123233231131)
- routes.simple_route.origin_pools.endpoint_subsets

<a id="canonical-3011013130123232-2103200302330030-2001112113011211-0313333022321322-3220203330111123-3203201302212200-3000021231133303-0113231010013220"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010331333023020-1030001323032332-1212311011100133-2101122100222210-0211220123033310-3131301112131001-0031011211200201-3311210132300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.origin_pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-026.md#canonical-0003132220301321-1220330033212333-2222322002322300-0313013021302201-2130003013111320-1313000211100023-1031121223320111-2301123233231131)
- routes.simple_route.origin_pools.pool

<a id="canonical-3312300200011110-1110331122202110-0330233212032303-1232332211023003-3133102132000023-0030233300002322-1031110101213131-2320320222110003"></a>

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

<a id="canonical-2200120321131122-1313113122220320-0023001211013323-2231222023322123-1113223200012002-2103122231312132-3131022302030330-0010223311210001"></a>

### Direct properties for `routes.simple_route.origin_pools.pool`

<a id="canonical-3211221003202330-2122230000331313-0213031023210132-0011312333132001-0031033111020133-2023320323320232-1130200302330220-3123113221103123"></a>

#### `routes.simple_route.origin_pools.pool.name` property

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

<a id="canonical-3121302032022130-2013023232100113-0233001020332222-1110213230123222-1301023100133313-2301201113231303-2002232212000033-0011300110211310"></a>

<a id="canonical-0201101212103003-2200022001121000-1201103211301001-2102302020100113-1003022220211223-0132333002203212-3311132003011133-0030311201033103"></a>

#### `routes.simple_route.origin_pools.pool.namespace` property

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

<a id="canonical-2102220112122113-2300020331300303-0121221200001303-1322212010112321-2311313330330013-1030112113121121-2021301330313321-3110202021123020"></a>

<a id="canonical-3210211133121120-3013031231121212-3232313001221121-3233210233211312-2220321021313310-1000232213000100-1131333033201103-1223122001113011"></a>

#### `routes.simple_route.origin_pools.pool.tenant` property

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

<a id="canonical-1310003223002313-2031303232323101-0223020331212312-1121103021001223-1321123103101310-0210203301020320-1310203332100001-1302312231333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.path

<a id="canonical-1333030112000332-3133000311032202-2133232231303330-1011331203132222-3010011113222223-1301230312121221-0220313233322312-1013220003030200"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-3322012101100020-0003023233020213-3021132301013221-2301303223103131-0000220030013132-3221310130233000-0021001210111013-1203123233312022"></a>

### Direct properties for `routes.simple_route.path`

<a id="canonical-2030123320210001-0111311010013011-0022310333321220-2103030212031302-0001233212012011-3033133113012133-1131110200032220-2203213301123110"></a>

#### `routes.simple_route.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2302313001011213-3222222103230013-1010221201331002-3222101010021220-2110102310211300-2221201103101233-0321222013302211-2030000022213320"></a>

<a id="canonical-2220321312323011-3133002003030130-3201211201232003-0000122012201030-2333322223233203-2330223232100022-2233230313200323-2223323131102003"></a>

#### `routes.simple_route.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0003023103010202-2202112222332312-0220231212021020-3320113003203131-2203211331232203-3000101212011303-2320100010202030-1012000112120030"></a>

<a id="canonical-3230102201031202-2313132033230133-1001312120001301-0100000121312310-3132021030102013-3001131020313030-2302032120023022-0103112131003100"></a>

#### `routes.simple_route.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3013022233310333-3211123033103002-2011301011131203-1221021311100202-1220322312201131-1232212212011223-1330231331301232-3032220031230320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- routes.simple_route.query_params

<a id="canonical-3303031300310322-3031120210022120-1222120022030013-3201231212130322-1202102032301002-3031210023303322-1332011002133231-1231323331022323"></a>

Type: `"single"`. Computed.

Handling of incoming query parameters in simple route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

<a id="canonical-2110233031233001-2320203001323230-3132012121231110-3002002012020012-3321122211330113-3100101133023010-0012200023212111-0212302222300100"></a>

### Direct properties for `routes.simple_route.query_params`

- [remove_all_params](data-sources--http_loadbalancer--reference--group-026.md#canonical-2030312303031013-2002211010012013-1213313120011231-3020030031020330-2121331000020121-1232010302221210-3120331003302323-1023201300221130): complete subsection reference.

<a id="canonical-1213321100010331-1102003003232311-0300021232110203-1011021211031131-3030231010301021-2233303331332012-1332230030012213-0332020311002320"></a>

<a id="canonical-3021230120200022-3010310321101020-3101001300011023-3101303102233020-2121011331120333-0122000220330112-1211123031322300-2233003113021122"></a>

#### `routes.simple_route.query_params.replace_params` property

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [retain_all_params](data-sources--http_loadbalancer--reference--group-026.md#canonical-2131313030331302-3323223333230302-2313010010122221-1032313311221010-2003013332313320-3233203001323021-1230322223212312-3302113032331022): complete subsection reference.

<a id="canonical-2030312303031013-2002211010012013-1213313120011231-3020030031020330-2121331000020121-1232010302221210-3120331003302323-1023201300221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.query_params.remove_all_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-026.md#canonical-3013022233310333-3211123033103002-2011301011131203-1221021311100202-1220322312201131-1232212212011223-1330231331301232-3032220031230320)
- routes.simple_route.query_params.remove_all_params

<a id="canonical-2302330030300232-1200320113101133-1303311202110201-3013101302302310-0212230102013302-1211112320223130-1032312322030310-1203112302001320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-2131313030331302-3323223333230302-2313010010122221-1032313311221010-2003013332313320-3233203001323021-1230322223212312-3302113032331022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.query_params.retain_all_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [routes](data-sources--http_loadbalancer--reference--group-024.md#canonical-3232213310021101-3001022200303021-3022000231122212-3301110112333122-3031302113101300-1300212103000230-3011232013130010-0033032303113000)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-025.md#canonical-1013212301331233-0220132333230332-3120203333230311-0203012113113321-3321231220112311-1122003101103203-0102302112023322-1122000203023103)
- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-026.md#canonical-3013022233310333-3211123033103002-2011301011131203-1221021311100202-1220322312201131-1232212212011223-1330231331301232-3032220031230320)
- routes.simple_route.query_params.retain_all_params

<a id="canonical-2001233212333021-1223201133023130-3122332302012330-1011131021121102-1013113131303120-2123331112030313-3233213121030103-0122120322130113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- sensitive_data_disclosure_rules

<a id="canonical-2013210220231001-2131302311333232-2121100303010033-3312102323101300-2312112000021011-0201232211321213-0011020332200223-3303300222222022"></a>

Type: `"single"`. Computed.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020120021323121-1332301232232312-2211310011000103-3012213211010101-0030112021223121-2300323310302311-3113110332100000-0322023323203132"></a>

### Direct properties for `sensitive_data_disclosure_rules`

- [sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-026.md#canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321): complete subsection reference.

<a id="canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules.sensitive_data_types_in_response` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="canonical-0332132210120301-1211131212233303-0012123300002310-3111012212201332-0112032333003000-2021120031302213-0012033211013030-1033220201231030"></a>

Type: `"list"`. Computed.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333201313220111-1222123003110133-1000231000223202-1312012003211200-0300320310331132-2003231312023111-1120332100102121-3133303221021300"></a>

### Direct properties for `sensitive_data_disclosure_rules.sensitive_data_types_in_response`

- [api_endpoint](data-sources--http_loadbalancer--reference--group-026.md#canonical-0223211112203132-3213101003131222-1230332300321100-1333011021323120-0103233313332121-2333321230102002-0320003223022021-3313002203010002): complete subsection reference.

- [body](data-sources--http_loadbalancer--reference--group-026.md#canonical-3010010003013121-2203101333021322-3221113010222013-2201013301203101-1302201201223013-2332231003202311-1033201232210021-0101111031000100): complete subsection reference.

- [mask](data-sources--http_loadbalancer--reference--group-026.md#canonical-2110210000233021-3202211001331333-3202321202101303-2012232111020000-0322030133310021-0000223313222233-0110001210032312-2221020122103030): complete subsection reference.

- [report](data-sources--http_loadbalancer--reference--group-026.md#canonical-0121212022011311-1020323112123032-0001300123033013-0300122331212331-2003130031020333-3202030123232132-0030202232133230-3100322000210101): complete subsection reference.

<a id="canonical-0223211112203132-3213101003131222-1230332300321100-1333011021323120-0103233313332121-2333321230102002-0320003223022021-3313002203010002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-026.md#canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint

<a id="canonical-0130333113302011-0133113010323230-1013201012331031-2003033301133122-2313031301133101-3110230213003123-1000010000031223-1323021332121231"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2002111211200313-0222210212003020-0122310011221113-0123221231002233-0030211301320100-0112310000213303-1322331121030332-0313121300031320"></a>

### Direct properties for `sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint`

<a id="canonical-0211023001121332-0002303133331301-2201130122113001-3101002101112333-3212102122230300-3011022002130120-1201022331132200-1333320331121223"></a>

#### `sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint.methods` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2200223032223313-3132013320332310-3101203111010202-0023120111102200-3101202211122333-1301322313033303-1103213101313030-0203003302101103"></a>

<a id="canonical-1123303210100031-2312033102012212-3012232133120121-3112123211133131-2103100013311223-0332020003202032-3200031333301302-1120231000312100"></a>

#### `sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint.path` property

Type: `"string"`. Computed.

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

<a id="canonical-3010010003013121-2203101333021322-3221113010222013-2201013301203101-1302201201223013-2332231003202311-1033201232210021-0101111031000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules.sensitive_data_types_in_response.body` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-026.md#canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="canonical-2102230013303200-1112112111303002-3323011302023112-3031100220321013-2123020012203122-3003033311221102-3213231132121123-2230310212311132"></a>

Type: `"single"`. Computed.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1300012300130302-0320322323030021-0013312220011210-0133123232202300-2112310103222103-3201233232020020-0121233110120202-3002032131322120"></a>

### Direct properties for `sensitive_data_disclosure_rules.sensitive_data_types_in_response.body`

<a id="canonical-1122130121213220-1301111130333013-2203220200011301-0230200310332313-0121333322212033-0223201200001213-3333002133330303-1020130122021321"></a>

#### `sensitive_data_disclosure_rules.sensitive_data_types_in_response.body.fields` property

Type: `["list", "string"]`. Computed.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

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
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2110210000233021-3202211001331333-3202321202101303-2012232111020000-0322030133310021-0000223313222233-0110001210032312-2221020122103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-026.md#canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask

<a id="canonical-0121300100112023-1212203332031212-1302110122211102-0110003222012110-0110201032032022-3021222320323111-2333323112302302-0111012103230212"></a>

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

<a id="canonical-0121212022011311-1020323112123032-0001300123033013-0300122331212331-2003130031020333-3202030123232132-0030202232133230-3100322000210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_disclosure_rules.sensitive_data_types_in_response.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2312300322011001-1323320300202210-0001110221031030-0320300213213203-0013111103111133-0212131010131121-2111200121200102-3312112221121301)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](data-sources--http_loadbalancer--reference--group-026.md#canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.report

<a id="canonical-0130303321130130-0131213221102110-3102110031013300-0123231311110303-2200232032320101-1122232000123223-0103010020323302-3120210132200223"></a>

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

<a id="canonical-1111113201333121-3123100001123201-0011013100212320-1131120220102201-2312121211133221-0310211131231233-2223203221032301-2301031302111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- sensitive_data_policy

<a id="canonical-1233012233330121-3030022033230201-2333203001311122-0122203010231303-0110030001013320-3113122201121323-2222321313230113-3223013123232202"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Additional upstream details:

Settings for data type policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123100130132300-1003223011233301-0022001333202020-3230013312121102-3102103112031333-0301321113033021-2321120332113232-0311120223332202"></a>

### Direct properties for `sensitive_data_policy`

- [sensitive_data_policy_ref](data-sources--http_loadbalancer--reference--group-026.md#canonical-3003201012332132-1212002333321231-3321132110012302-1012331011031122-0021313212111131-1022131333313103-3211320220223300-3203021112013000): complete subsection reference.

<a id="canonical-3003201012332132-1212002333321231-3321132110012302-1012331011031122-0021313212111131-1022131333313103-3211320220223300-3203021112013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sensitive_data_policy.sensitive_data_policy_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [sensitive_data_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-1111113201333121-3123100001123201-0011013100212320-1131120220102201-2312121211133221-0310211131231233-2223203221032301-2301031302111331)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-0332203311221231-1032020011213032-3231032120010010-0213222100332100-2331321000122030-2012011023122032-0201021133233332-2212001120032100"></a>

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

<a id="canonical-1133313000322301-1133033131132332-3322331230221113-2212132110303030-3023001011321223-2133113313033311-2011010023303223-2231100103032021"></a>

### Direct properties for `sensitive_data_policy.sensitive_data_policy_ref`

<a id="canonical-2330221323100211-1223323030220310-2232303001123033-2123333222221322-3233032033213013-3011012100123003-0102103300201122-3021120213111211"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.name` property

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

<a id="canonical-2110322013222330-1210303130220302-3121023223131031-3120002030002231-1021100331110233-1212210103231103-1232310023010321-1323030101112302"></a>

<a id="canonical-2003030130001030-2223213202210313-2323331200121333-1122210332103123-3201130133003031-1310323321331023-1213323103100133-2102222013002310"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.namespace` property

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

<a id="canonical-3113021331000110-0111103303002123-3220131331120122-3003121232211001-0002101202110020-3320321220133013-2331000232312323-3330012312013310"></a>

<a id="canonical-3232001110231100-1111002111100000-3202231203130101-1002110132133201-0100013102231133-2210022333012213-1232331020300332-0303011223130332"></a>

#### `sensitive_data_policy.sensitive_data_policy_ref.tenant` property

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

<a id="canonical-3032010102013301-1011233111333022-1301232320231300-1133103120030100-2321013303210022-2213102120021010-1000322312230310-2230332300333132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_policies_from_namespace` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- service_policies_from_namespace

<a id="canonical-2210322011133032-1130200230121211-3102212123010321-3312101211130233-3323101010013002-1002011302333221-3302300230113112-1120220333133013"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- single_lb_app

<a id="canonical-0031200011131311-2233112323100033-0111001011133120-2221230201121322-2021123130030211-2222030012320123-0323231202112012-3300012012003232"></a>

Type: `"single"`. Computed.

Specific settings for Machine learning analysis on this HTTP LB, independently from other LBs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_choice": "[\"disable_discovery\",\"enable_discovery\"]",
  "x-ves-oneof-field-malicious_user_detection_choice": "[\"disable_malicious_user_detection\",\"enable_malicious_user_detection\"]"
}
```

<a id="canonical-1003230132121320-1302301321032200-2322031332210010-0010030213331212-3313322122202300-3331311132113221-1033301001122222-2302131203321213"></a>

### Direct properties for `single_lb_app`

- [disable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-1233231130000331-3230133122312003-1022320003101330-3113102300133321-2103321303203301-1213312000301003-3122202031201222-1032012113230001): complete subsection reference.

- [disable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-026.md#canonical-2313033210011111-1031113023322233-1020323131021001-2230100220110130-3211303302311132-3320110032210202-0202112020103323-0212023102001130): complete subsection reference.

- [enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023): complete subsection reference.

- [enable_malicious_user_detection](data-sources--http_loadbalancer--reference--group-026.md#canonical-1100132111033032-0200211233102031-1111121300100222-3322202230312133-1300122030231312-1122033010221013-2113300203300320-0223212123033222): complete subsection reference.

<a id="canonical-1233231130000331-3230133122312003-1022320003101330-3113102300133321-2103321303203301-1213312000301003-3122202031201222-1032012113230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.disable_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- single_lb_app.disable_discovery

<a id="canonical-0120323100232313-1332221300312303-3003030232101101-0122332200322330-1330130301131002-3212223013002320-0113111122101023-0203001011333133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable discovery.

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

<a id="canonical-2313033210011111-1031113023322233-1020323131021001-2230100220110130-3211303302311132-3320110032210202-0202112020103323-0212023102001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- single_lb_app.disable_malicious_user_detection

<a id="canonical-2331312030222330-1330330230101031-1311233222301210-0300230201332321-3211323303030311-1113312232213312-0122113331032203-2130300202132002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable malicious user detection.

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

<a id="canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- single_lb_app.enable_discovery

<a id="canonical-1013021130000121-0233132313220031-1133322300231001-1000020322032023-0332222300031030-3133003313033322-2300120321313332-2312333210123102"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

<a id="canonical-1022310020233023-0323302220101023-1221022300330302-0113120111220003-0331332312102310-3011213113233232-0031012110101121-2121123010121101"></a>

### Direct properties for `single_lb_app.enable_discovery`

- [api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133): complete subsection reference.

- [api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-026.md#canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003): complete subsection reference.

- [custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-3021213101102110-3311000133033333-1312001032222023-3101120301131022-2301132020230320-2210102022133120-0111213212021231-0233023213121020): complete subsection reference.

- [default_api_auth_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2311033321231111-1010211201112222-3312200133120130-2132331121120320-3202200120011133-1103132101211302-1202130000213212-3113121032031012): complete subsection reference.

- [disable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-026.md#canonical-2020211223231133-0120313322101133-0103220121132330-2112021003301313-3310301322323020-1110230103332103-1032220122130210-0210313133202323): complete subsection reference.

- [discovered_api_settings](data-sources--http_loadbalancer--reference--group-026.md#canonical-1213131303010033-2330331212203112-1303100001300320-3303111000231300-3030330021332232-0301300111200321-2021233231001033-3001332132003011): complete subsection reference.

- [enable_learn_from_redirect_traffic](data-sources--http_loadbalancer--reference--group-026.md#canonical-1310231131213123-0322312303003203-3002301033102331-1011031210033232-2300012011301220-2011232231023222-0111310000110112-3321031332313021): complete subsection reference.

<a id="canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.api_crawler

<a id="canonical-0101101133213100-0101031012121332-0333232313210303-2311302212200102-3131002010002330-1231233023110300-0121230110031013-2000310123220221"></a>

Type: `"single"`. Computed.

API Crawling. API Crawler message.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

<a id="canonical-3321020022113322-1323213301222130-0122122113102221-2013332121211303-3202013121203312-1233102003201102-1232031331211132-3213312011111211"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler`

- [api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211): complete subsection reference.

- [disable_api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-2023320021030302-3230031133311200-3010002213201212-3323330300011223-0100311233322322-0102012320023011-0221233323321302-3123023022131310): complete subsection reference.

<a id="canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config

<a id="canonical-1332312312130000-1203213021211000-3332330110032230-2312310102223322-1330303003033330-0020202013323300-0332203233312223-1303310230303000"></a>

Type: `"single"`. Computed.

Crawler Configure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103303002122221-0211022133111213-1120202110123112-2303023131221323-3120330203032120-1110320011101032-1223122203312310-1211112131311131"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config`

- [domains](data-sources--http_loadbalancer--reference--group-026.md#canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313): complete subsection reference.

<a id="canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-3131210111030012-1312332320213001-2213221131323121-0213213313011002-3320023002132130-1012130102021301-1302232300223202-0320100230131330"></a>

Type: `"list"`. Computed.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-3201233011301033-1200213202020123-3010212333231230-0320002122323232-2001323121001323-3322212022211330-1101213132030333-3303102011132210"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-2230310323120013-2123132221221222-2131133023231022-2001231001221311-2223113313303101-3230011132313210-0101332321213122-0032000202013021"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Computed.

Select the domain to execute API Crawling with given credentials.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133110123002012-1230323111010302-3120213320000221-0300020012133023-2312322310021330-1110332103231100-2211102022233002-1231322011332030): complete subsection reference.

<a id="canonical-0133110123002012-1230323111010302-3120213320000221-0300020012133023-2312322310021330-1110332103231100-2211102022233002-1231322011332030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-026.md#canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-2231221131303032-3320322302023100-2210003200212231-1023010230003300-0201211210022323-0223302121010331-3212112001321113-1102322313120023"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0121320133113220-3013223033132123-1303311030023022-1010232321203212-3300010033202222-2223013023100221-3121112230121132-0300111333102021"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130001122202212-3301313221322222-2332100120012330-2330100100331130-2232223323002123-3030012203222311-3123313300123123-3010220020113110): complete subsection reference.

<a id="canonical-0012212130331111-0220203303303230-2202211303213222-2232130213100202-2310002223322230-0213310130103221-1121003302301020-2123123102030022"></a>

<a id="canonical-0111113223331112-3032220111310103-0302003011123323-3202013213331120-3331223100032102-1233222121123000-3120232010133200-0020303322302222"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3130001122202212-3301313221322222-2332100120012330-2330100100331130-2232223323002123-3030012203222311-3123313300123123-3010220020113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-026.md#canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133110123002012-1230323111010302-3120213320000221-0300020012133023-2312322310021330-1110332103231100-2211102022233002-1231322011332030)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-1031222123102031-3003012211222113-3013103332021222-3033213222311000-0202030113202200-0110021323201132-1212330311020021-0201100330323212"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1220303233022011-2011332120121133-1003303331212103-3331332001123313-0221121031312111-1101333332121300-2230230100331201-0233211022322313"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-026.md#canonical-0100321103222301-3011231313112332-0201122100133123-3113111133233132-0201231311311102-0010223332230000-1202201131330123-1102012002122100): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-026.md#canonical-3012330012231122-3311233313203231-0300301031133211-3001330001002031-2103102201202012-1012032130023232-2121322233001221-3322001312122331): complete subsection reference.

<a id="canonical-0100321103222301-3011231313112332-0201122100133123-3113111133233132-0201231311311102-0010223332230000-1202201131330123-1102012002122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-026.md#canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133110123002012-1230323111010302-3120213320000221-0300020012133023-2312322310021330-1110332103231100-2211102022233002-1231322011332030)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130001122202212-3301313221322222-2332100120012330-2330100100331130-2232223323002123-3030012203222311-3123313300123123-3010220020113110)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-0110321101100210-1331311012113311-0222032003132230-3310211022133132-1300121313322103-1321110013011212-1231111323331330-3303321130102301"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1232110311200130-3233132122211333-1232122010023313-0103001013202210-3013211120212211-3130010030310110-0132100223332323-3310311203310222"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-2302002303020223-3120203220222233-3030122211300211-0212113010331121-0213002311123100-1133201033231200-1313121321002101-0222233232222131"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0033120013200220-0202222302210233-0311232210100223-3100320212332121-3002112302001331-1331233012100323-3231011231211301-3222023021130130"></a>

<a id="canonical-1220323322312002-1202202121120132-2132031102221120-1102022023113001-1020300130100233-0213231132021131-2220033233000013-2003021020122010"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3310321201300201-3110231032130022-2002031010323200-3130002121121301-1231331211231030-0320230300333220-1033030023033101-1133213110102011"></a>

<a id="canonical-0121302033120221-3321300110032121-0103121001210322-1103201200123203-0310320022121221-3332331201101123-2031120221313331-2033001001101000"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3012330012231122-3311233313203231-0300301031133211-3001330001002031-2103102201202012-1012032130023232-2121322233001221-3322001312122331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133120322312222-2312223333031212-1032032132030213-3301300112323232-2302012312301022-3110302310020121-2131323233322010-2212313212231211)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains](data-sources--http_loadbalancer--reference--group-026.md#canonical-2012013101130033-3013103112133302-2021021120313310-0212120003320102-3120212233321000-1230031111203212-2033123222231203-1302231123110313)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--http_loadbalancer--reference--group-026.md#canonical-0133110123002012-1230323111010302-3120213320000221-0300020012133023-2312322310021330-1110332103231100-2211102022233002-1231322011332030)
- [single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130001122202212-3301313221322222-2332100120012330-2330100100331130-2232223323002123-3030012203222311-3123313300123123-3010220020113110)
- single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-3221120321320130-1020120221011002-0300333330232022-3310022101312003-1002330313002110-0022202122311303-2001031233210212-1001121333002103"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011122220213132-3323303320213222-0333011100023300-0211010210122232-3013323321332101-3123312120123033-0212101232232321-2123331031331123"></a>

### Direct properties for `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-2021131330332210-0202303112133013-0330112022213303-1003231032023312-0103120111202232-0120301123020030-1212020011333303-1333210320123030"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1333231012202321-1013113321233030-3303223213323202-0233030322110033-2230310111103101-3223212002311231-2010222320311332-3323100220311133"></a>

<a id="canonical-3322211001201300-2201103101121022-0000210123010130-0101101312200201-0310211122312111-3033003031302313-1013312130123120-0220220232220100"></a>

#### `single_lb_app.enable_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2023320021030302-3230031133311200-3010002213201212-3323330300011223-0100311233322322-0102012320023011-0221233323321302-3123023022131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_crawler](data-sources--http_loadbalancer--reference--group-026.md#canonical-1122300132122203-1113231131101200-0021200301202132-2030211130001200-3100101010012101-1120303001023123-3223212121001201-1011322201031133)
- single_lb_app.enable_discovery.api_crawler.disable_api_crawler

<a id="canonical-2003211123332122-3313113311123032-2300230213231201-1132303332320320-1123003021333333-0300222321101031-3030330223300333-1303331120201200"></a>

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

<a id="canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.api_discovery_from_code_scan

<a id="canonical-2032020321001330-1101222202300301-2012212002112131-3010330210311323-2102322030222000-1322222132302012-3103320100100033-2011322330102233"></a>

Type: `"single"`. Computed.

Select codebase and Repositories.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220302312131331-2223010103102002-2312003323001320-3123032113310211-3022131200320311-1012022002320201-2022200012203103-3301112300031010"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan`

- [code_base_integrations](data-sources--http_loadbalancer--reference--group-026.md#canonical-0131010222303121-2003130330221120-0213033313020022-3113001032013220-2031112100110222-2220032130113032-2031022010010301-1203332311333223): complete subsection reference.

<a id="canonical-0131010222303121-2003130330221120-0213033313020022-3113001032013220-2031112100110222-2220032130113032-2031022010010301-1203332311333223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-026.md#canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-3130212221132211-1302212312213303-1102232112123300-0012001230101233-3021112110320210-0202101223210232-3001232001311233-0122232202220200"></a>

Type: `"list"`. Computed.

Configuration parameter for codebase integrations.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3220031112223020-1320231023213202-2030330103130331-1212303101020303-3321222312002213-1113133331102323-0221200302333010-0111133020003302"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](data-sources--http_loadbalancer--reference--group-026.md#canonical-0310120133101002-2013021012021131-3330302001032223-0123112313012210-1123122231310021-3331011313303230-3331110001100231-2001312221130020): complete subsection reference.

- [code_base_integration](data-sources--http_loadbalancer--reference--group-026.md#canonical-0010222013000222-0121232032031030-3312331030030301-3323300211132321-0203021303201311-0022232022032033-0033220113200011-3032122102131310): complete subsection reference.

- [selected_repos](data-sources--http_loadbalancer--reference--group-026.md#canonical-2331001311312010-1310031003233132-3110032121120130-3123310321110231-2022211132330320-3112132000002312-3303020212201111-3311101320013212): complete subsection reference.

<a id="canonical-0310120133101002-2013021012021131-3330302001032223-0123112313012210-1123122231310021-3331011313303230-3331110001100231-2001312221130020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-026.md#canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-026.md#canonical-0131010222303121-2003130330221120-0213033313020022-3113001032013220-2031112100110222-2220032130113032-2031022010010301-1203332311333223)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-3203110223213101-1021333233202320-2300021003023103-3010102113121332-1312322231221122-3311330010301030-2321232032330131-2322320232222332"></a>

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

<a id="canonical-0010222013000222-0121232032031030-3312331030030301-3323300211132321-0203021303201311-0022232022032033-0033220113200011-3032122102131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-026.md#canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-026.md#canonical-0131010222303121-2003130330221120-0213033313020022-3113001032013220-2031112100110222-2220032130113032-2031022010010301-1203332311333223)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-3330010123021010-0030022100323020-2011313232102010-1310301230012322-0320233223020103-1221133321200333-0100221323202031-2013112131120210"></a>

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

<a id="canonical-3332330010001233-0311113100010203-2202023220111222-1230012023210103-1300122123013020-1230022330223320-3012202113110201-0331021220200331"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration`

<a id="canonical-3123002002211223-0130231323222111-2101330030132100-2030332101310310-2311100023231322-3211021232100131-3331030231130313-3323331330031322"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.name` property

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

<a id="canonical-2312102101113101-0231220333003012-3231332000022310-2022333321121220-2230221133323231-3012020123323223-0223211113230213-1322330220103323"></a>

<a id="canonical-0222102212232021-3322002303033313-1020323310001303-2013323200132032-1230302130123211-3011110012021312-1212020213221221-1320110322303032"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.namespace` property

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

<a id="canonical-0110033333012102-3022130023230103-2113113112310103-2001021010120320-2223332121212312-0211112213232100-0103333021201310-0003300013321001"></a>

<a id="canonical-3321133110002101-0033332231020031-3022013101110232-2201201020133222-1123220112011011-0232130223023101-1003110110122311-2000303333012003"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration.tenant` property

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

<a id="canonical-2331001311312010-1310031003233132-3110032121120130-3123310321110231-2022211132330320-3112132000002312-3303020212201111-3311101320013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-026.md#canonical-0322203133112201-2312001131022121-2222200121331110-0120230232011113-1101122003000222-1131120232112101-3311000003213012-1212300110221003)
- [single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-026.md#canonical-0131010222303121-2003130330221120-0213033313020022-3113001032013220-2031112100110222-2220032130113032-2031022010010301-1203332311333223)
- single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-0010301022332103-3013101011310102-3021100212121133-3330111232230210-0000302112020122-1121320210310103-3312023200211222-0031023212103123"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021311313223203-1100213210210111-0033231230103132-0032000002013022-3010323333333031-0112101003200310-1323013202321321-3122020012103103"></a>

### Direct properties for `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos`

<a id="canonical-0200123022003233-3103103131011022-0331310231300013-1220231221331021-2322320323321232-0003201303013212-0311211321312221-0312220201101232"></a>

#### `single_lb_app.enable_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos.api_code_repo` property

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3021213101102110-3311000133033333-1312001032222023-3101120301131022-2301132020230320-2210102022133120-0111213212021231-0233023213121020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.custom_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.custom_api_auth_discovery

<a id="canonical-3020203331301223-1303031301333230-1300011332013231-1223031130332312-1301313221020221-3111103001133213-0200102201330010-3212023123130000"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303001203103211-0003031131313231-3033222121003222-2233013301300213-0101120220202233-3300103320222012-1222133030121002-0222213321201321"></a>

### Direct properties for `single_lb_app.enable_discovery.custom_api_auth_discovery`

- [api_discovery_ref](data-sources--http_loadbalancer--reference--group-026.md#canonical-2222112212130021-1323023112111223-3313230211003003-0033303100031311-2003320301313103-0313330313323111-3210122233201103-3300233100220303): complete subsection reference.

<a id="canonical-2222112212130021-1323023112111223-3313230211003003-0033303100031311-2003320301313103-0313330313323111-3210122233201103-3300233100220303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- [single_lb_app.enable_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-3021213101102110-3311000133033333-1312001032222023-3101120301131022-2301132020230320-2210102022133120-0111213212021231-0233023213121020)
- single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-3102202012002031-3001220333312130-2010022201311011-2302220211131223-3330013012320001-2130323122001023-0022122121312010-0333310330321320"></a>

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

<a id="canonical-2320222001011022-2321213011020002-0232333133220121-1023013120220101-3222103032203122-3332131303113110-0132121021301030-1013022001131333"></a>

### Direct properties for `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref`

<a id="canonical-0030103122122011-2110100222011122-0000020212000221-2303120222330223-3330013231022300-3333210310221232-0202110312011103-0300003333330021"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.name` property

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

<a id="canonical-1332000322101332-3302000032030103-2321231220233311-0310132033233121-1013223321100331-0221122012032110-3023012031321113-0200210311221230"></a>

<a id="canonical-3131322100312121-1230212223032022-2131313232003122-1320131112210233-1011113101002112-3131313302321210-3221002321321210-3121320223100012"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.namespace` property

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

<a id="canonical-0031100203312211-0001302121332113-3200210003320032-1300212300230203-3213333101200003-1030030033103100-2030332030112301-0333331022131311"></a>

<a id="canonical-0303312321223131-1010220010021033-2332013020130211-3210010022121000-3002103011320020-2332021021311300-1020131212122330-3103231010212031"></a>

#### `single_lb_app.enable_discovery.custom_api_auth_discovery.api_discovery_ref.tenant` property

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

<a id="canonical-2311033321231111-1010211201112222-3312200133120130-2132331121120320-3202200120011133-1103132101211302-1202130000213212-3113121032031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.default_api_auth_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.default_api_auth_discovery

<a id="canonical-2022301110312121-3202120310231321-0101322030002232-3002310233123111-0131200220211030-1110200200220012-0030021013100022-2210001331333301"></a>

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

<a id="canonical-2020211223231133-0120313322101133-0103220121132330-2112021003301313-3310301322323020-1110230103332103-1032220122130210-0210313133202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.disable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.disable_learn_from_redirect_traffic

<a id="canonical-3301120331103303-2023203120020120-1120020232130222-1032333131232201-2203013030223033-2232313131002210-1120320011200032-0230031020230000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-1213131303010033-2330331212203112-1303100001300320-3303111000231300-3030330021332232-0301300111200321-2021233231001033-3001332132003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.discovered_api_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.discovered_api_settings

<a id="canonical-1220012213000222-2012331003022113-0012011332123331-2220301123200331-1103221333311333-3101031033233001-1013200201311123-2122221001000000"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1020030102303000-2132103323013330-3100303313311301-2030100112133133-0210121022220211-0130132032231123-2211110331020002-0223300021102101"></a>

### Direct properties for `single_lb_app.enable_discovery.discovered_api_settings`

<a id="canonical-2320232211212121-3033233033233013-2322000113223223-2221131332023213-0313310032102130-1020033333122330-0220003223223031-1022120112313122"></a>

#### `single_lb_app.enable_discovery.discovered_api_settings.purge_duration_for_inactive_discovered_apis` property

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-1310231131213123-0322312303003203-3002301033102331-1011031210033232-2300012011301220-2011232231023222-0111310000110112-3321031332313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- [single_lb_app.enable_discovery](data-sources--http_loadbalancer--reference--group-026.md#canonical-2010220101022201-0331031001221111-2122032101321300-1012113030001121-3122321210012202-0002210102301013-2313130231232210-0220211103202023)
- single_lb_app.enable_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1312333102121303-1113311000221012-1230002010113111-3301202303111103-3121202330000220-1130312130032201-2311212120022001-1313222300020331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-1100132111033032-0200211233102031-1111121300100222-3322202230312133-1300122030231312-1122033010221013-2113300203300320-0223212123033222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `single_lb_app.enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [single_lb_app](data-sources--http_loadbalancer--reference--group-026.md#canonical-2233210203213223-1321222013111023-1010000021130322-1303110211330010-2012321032020330-1330323032211222-0222320333131221-0013232101213310)
- single_lb_app.enable_malicious_user_detection

<a id="canonical-2222330012023222-2222102022203320-1220232313112300-2030320222220123-2310332000333223-1311110120021023-1002222221201223-2100322031020302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-1123002022123300-1112121112130100-2221312020130023-3123011020012032-0303121031103103-2302330113032012-2232201101111100-3332313211000001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- slow_ddos_mitigation

<a id="canonical-3121031121321100-2012213003230023-3003203101312213-3231321020102033-3120333312001022-3223321321011001-2030230302212100-1332201302111300"></a>

Type: `"single"`. Computed.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Additional upstream details:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-026.md#canonical-3121031121321100-2012213003230023-3003203101312213-3231321020102033-3120333312001022-3223321321011001-2030230302212100-1332201302111300)
- [system_default_timeouts](data-sources--http_loadbalancer--reference--group-026.md#canonical-2311332221223231-3111310312131011-0212202323322201-0012320321330321-1322303231130231-3223132013131331-2311231333201222-1301213121123022)

Select alternatives according to the provider validators above.

<a id="canonical-2220210032133022-2133130333133110-2330113123110230-1231122113131123-2131101212313110-1331222232000113-0331232223100130-3321033011133033"></a>

### Direct properties for `slow_ddos_mitigation`

- [disable_request_timeout](data-sources--http_loadbalancer--reference--group-026.md#canonical-1233000100013121-3020120210323111-0003301133330201-1100110211111231-2120230111010023-0120130000033030-0200232130311223-0100102033111221): complete subsection reference.

<a id="canonical-1021311321213120-2230302021111131-2113310131212011-1300011313232322-2033121210020203-0023122212031232-0320012301102233-1122130013122320"></a>

<a id="canonical-2221022011303231-0313222313221132-3230011003131021-2131113122133131-1221032030233321-3323223223331231-2103300212231303-1231201303000110"></a>

#### `slow_ddos_mitigation.request_headers_timeout` property

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Additional upstream details:

The default value is 10000 milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-0212030202323000-3321021213323200-2133120221130111-1003223202223313-3120122111301200-2320012203132123-0103221002113021-0232011320212203"></a>

<a id="canonical-3100223213132021-3312003301312233-0233312321321210-3221130300213031-1200310302103231-0030031311121333-1300211003200133-0231311100112021"></a>

#### `slow_ddos_mitigation.request_timeout` property

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-1233000100013121-3020120210323111-0003301133330201-1100110211111231-2120230111010023-0120130000033030-0200232130311223-0100102033111221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slow_ddos_mitigation.disable_request_timeout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [slow_ddos_mitigation](data-sources--http_loadbalancer--reference--group-026.md#canonical-1123002022123300-1112121112130100-2221312020130023-3123011020012032-0303121031103103-2302330113032012-2232201101111100-3332313211000001)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-2211332212221103-1213303323121020-2213120320213300-3220121203301031-0131100300313131-3312132302322222-1332012331313121-1103113132103330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-1132130232001202-3132322001023103-2220102200330221-3102023123230101-2011032011032233-0011310023211220-1012121130020101-2232213232312220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `source_ip_stickiness` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- source_ip_stickiness

<a id="canonical-2232300333201231-1322210302201100-2133311321132322-2102213131331030-1211231100230011-0131302130333003-1112000021020133-3102331223113200"></a>

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

<a id="canonical-1002113332100220-0210021202232223-2330212033110033-3003023131320231-2021130022231113-0331333211332210-3013003301000123-3121110021322331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `system_default_timeouts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- system_default_timeouts

<a id="canonical-2311332221223231-3111310312131011-0212202323322201-0012320321330321-1322303231130231-3223132013131331-2311231333201222-1301213121123022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system default timeouts.

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

<a id="canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- trusted_clients

<a id="canonical-2032331132022032-2231310211103230-3302330023003212-1101212121100231-2223333013123130-1113110113333121-3200222103201233-1333213113330211"></a>

Type: `"list"`. Computed.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3110000122233321-1113221023203111-3002211201211310-1323120123323332-3233321130211230-0231331211313012-1332021111113302-2000301301120001"></a>

### Direct properties for `trusted_clients`

<a id="canonical-1310303333132331-0200111330000201-0232210120231111-1102332231211132-1300320301013330-1123022310323030-1320221321120311-1301330122312223"></a>

#### `trusted_clients.actions` property

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0122323213133320-3022010312331022-0320203111111310-2131001122221013-2220322022313123-2230211230020033-0103101023113302-0212000101210031"></a>

<a id="canonical-0302023000201133-1320232130331202-0003202130001202-2120003203002011-2231101101101111-3122021131222012-2231320301222323-3201313232123313"></a>

#### `trusted_clients.as_number` property

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--http_loadbalancer--reference--group-026.md#canonical-1221311101021013-0201131300003012-3332032223302300-0023123202202211-0100001223200202-0030132012103000-3212210111133320-1231122001011010): complete subsection reference.

<a id="canonical-2300311303320132-3331122000311322-2021220221232201-1332333231130222-1031332301010030-0132013012320300-1333331213332222-2013303330111333"></a>

<a id="canonical-0212032232103202-2211300121000201-0322233022323110-2231102130323133-1323123222121223-1232123320331322-2212021301120210-2031000112012332"></a>

#### `trusted_clients.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--http_loadbalancer--reference--group-026.md#canonical-3230030020131223-0131220302231003-1202033331313212-3132011202113220-1110322103023033-0001110313000232-2023300010012323-3102231132112100): complete subsection reference.

<a id="canonical-0223311132300312-1010012133222210-0322111300223102-3010033333231130-1132022101101130-0003013111222232-1110302211201023-0003013223120032"></a>

<a id="canonical-0211301201001032-2222321210012021-2331333133101112-0202303310013310-0120210102313322-1233110112011223-2130103110121200-0302203133331201"></a>

#### `trusted_clients.ip_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0211120123233323-2033331110312031-3032232033120302-2020222123330302-2002013223311130-3103230220323130-0212021023002132-2312301001102310"></a>

<a id="canonical-1102303120011303-0122022032322203-1330102322323311-0222013011313110-1232301301002123-0022220111012312-1230200010312113-3032310131202223"></a>

#### `trusted_clients.ipv6_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-027.md#canonical-3202020021333101-0232122300001313-1310212212321110-3311231212132311-2313112132132232-3202311332122013-3132131301122321-3232311311103133): complete subsection reference.

- [skip_processing](data-sources--http_loadbalancer--reference--group-027.md#canonical-3112332033131101-3011121131303030-2322233022200023-1033011221031321-1221020032320031-2131032122111013-2221303131132220-2113010210201012): complete subsection reference.

<a id="canonical-2230211303302220-1111021312333302-1212332102310213-0300211210031212-0333002221222001-3122301102201210-3211131100122032-3301032102301230"></a>

<a id="canonical-0330133221003233-0312101111330033-0212320103302303-0233033321033331-2221110220330200-2212211310113000-3323120211010001-3001021220212113"></a>

#### `trusted_clients.user_identifier` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](data-sources--http_loadbalancer--reference--group-027.md#canonical-0002202033200000-2031310200300103-2022023232313133-1301002112320111-0320301311221130-1231221131023322-0201122332032300-3310012101133310): complete subsection reference.

<a id="canonical-1221311101021013-0201131300003012-3332032223302300-0023123202202211-0100001223200202-0030132012103000-3212210111133320-1231122001011010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- trusted_clients.bot_skip_processing

<a id="canonical-1001103030201221-2123201330131303-3302230122103223-3120220311120301-1003311300011011-0002000213231012-2201230332013121-1103033213022300"></a>

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

<a id="canonical-3230030020131223-0131220302231003-1202033331313212-3132011202113220-1110322103023033-0001110313000232-2023300010012323-3102231132112100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- trusted_clients.http_header

<a id="canonical-2023331012000133-0211310002303211-2310213130303132-3313012010322023-1033111202012332-2113312231200331-1300311213230033-2332200211021300"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220313100103201-1121022210113300-3302323211110310-0200033332021132-0231221322133110-0001213212323113-3312330131121220-0102103321310233"></a>

### Direct properties for `trusted_clients.http_header`

- [headers](data-sources--http_loadbalancer--reference--group-026.md#canonical-2123202123101322-0222221032113031-0301131313120111-0333302200030310-1123033130300033-3113103022022032-2000223321203001-1003220332330103): complete subsection reference.

<a id="canonical-2123202123101322-0222221032113031-0301131313120111-0333302200030310-1123033130300033-3113103022022032-2000223321203001-1003220332330103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- [trusted_clients.http_header](data-sources--http_loadbalancer--reference--group-026.md#canonical-3230030020131223-0131220302231003-1202033331313212-3132011202113220-1110322103023033-0001110313000232-2023300010012323-3102231132112100)
- trusted_clients.http_header.headers

<a id="canonical-3232032330232231-2132330232200103-2220233120230333-0320100003301020-3130022000131330-2301230303211331-2132020323003112-1200200222330300"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```
