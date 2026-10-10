---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2210002000001212-0010203202310332-1010121210231301-3211000110202332-0022230020101320-1312133232033210-1220213012122321-0223023202033001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.compression_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.compression_params

<a id="canonical-1121131012300332-2021323232113211-0313320003030013-0323330333223223-2332101130220121-1033321203333320-0031313131030120-0102312030331010"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Receipt-pinned upstream constraints:

```json
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310101201233022-3012122022210300-2233012231032323-0211001231333130-2223022020230320-3322321323331031-2133201232123323-2021003221331211"></a>

### Direct properties for `more_option.compression_params`

<a id="canonical-2232222112120022-1210233201132300-2011232023322030-1333032011301300-2220023022001332-2330223031101322-3010113121012111-3202220023233122"></a>

#### `more_option.compression_params.content_length` property

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-1222121131321102-0310221332100002-2112113033112111-2302221103312121-1120110032022002-2323333200210131-2123110011033322-3210330233013111"></a>

<a id="canonical-2001321110120231-1012022031231210-2123022110210103-1331113232023201-3311230333131110-3303331130110210-2200321000213310-1232311302011000"></a>

#### `more_option.compression_params.content_type` property

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Additional upstream details:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3322132223311003-1013301032111003-1001203323301010-3313011021110132-0112231202332102-0221203300231221-0101112022231331-1033023331223303"></a>

<a id="canonical-0022100322230120-0222301011100323-3020233001100313-3101113310313300-3330211011213331-0222310200332110-0003133312102203-1233013003310103"></a>

#### `more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3003330301022113-0311323333232302-0130213220123233-2033201213121321-1002300201323222-3320133113223212-1223200200023102-3330201111220002"></a>

<a id="canonical-3102211330023033-0103330121010230-3103131112301132-0322320013210130-1111103001333103-0212013120300120-2002321021002310-1332111221310100"></a>

#### `more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1033132001213130-2001323231130303-3302133101202300-2012320223111211-2113022313122123-3231031321132200-0022333302032210-1001120012321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.disable_path_normalize

<a id="canonical-0031312320132020-3012321002233212-2012033220311202-0220320331220201-0133312310033010-3001301122000121-1330321221230022-3210300033320232"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.enable_path_normalize

<a id="canonical-2223132031010201-3001001330102303-1100302130200330-1131133312013011-1121233200131030-1233301013023331-2121112011101023-1102131102221132"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023102010030230-2332223133022232-3123113033013210-3231031103313320-3102130133132220-1233310023222222-3310320102211232-1232211101323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.no_request_limit_per_connection

<a id="canonical-1111012202212331-3101223002301022-1111133110302213-0300013131301101-3321310113223220-3110000330131212-0030333130303223-0233231320003111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.request_cookies_to_add

<a id="canonical-3032323122310110-3121030123113131-3023320332332001-3202231032213332-1120301133031303-2213133132222320-1230322332200312-3121212312110220"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213121122032113-3101211223101220-3131022022331021-0200130013210121-0322011302323111-0002023010330011-1120322021101022-3103210111330220"></a>

### Direct properties for `more_option.request_cookies_to_add`

<a id="canonical-3223031230130313-2102013213320213-0230313030120222-2010321020000032-1232311222103111-0110132331011332-0302012332130110-0022013131223010"></a>

#### `more_option.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0330001022311333-1112122103232003-2032332310213200-2003213210223022-3020300311111220-1112331331331130-2112213303223311-0130113123332002"></a>

<a id="canonical-3202203222112113-2222100003202221-0121223222212233-3303003001131333-0211213023230312-1002131011212111-2121110210121022-2323232303203330"></a>

#### `more_option.request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002): complete subsection reference.

<a id="canonical-0213331232201033-1313310132332030-1322013201332100-2001333030222210-3222212212313223-1332333113230120-1322011212110013-2112211320101332"></a>

<a id="canonical-0031321321102120-0311123230011103-1313313133131221-0102020211222201-3003131102032131-3213120232032310-2003102102202322-2302322333302203"></a>

#### `more_option.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- more_option.request_cookies_to_add.secret_value

<a id="canonical-0221121301010031-3113320022011033-2132220322220332-3301332130321120-1200230002022010-2230220333110030-3301003313012221-2133032013022121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230330021021022-1332203313120100-3133122120010212-1302112213120311-2100022331100232-0103301330203220-0112030323010202-2330113201201022"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3212223011311230-1101101021112121-0000001323101121-3121010323012221-1122033033202333-0113322212312222-0130101301111303-1012021321001103): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-1130230013202222-3321303302113110-0031022130203202-3021000132131211-1232111221120003-3111130031312020-2000132202000331-1213100002032323): complete subsection reference.

<a id="canonical-3212223011311230-1101101021112121-0000001323101121-3121010323012221-1122033033202333-0113322212312222-0130101301111303-1012021321001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3022301301330213-3130223201112012-3221013302021031-0322101202130203-2132132022011133-1131203233132223-0333212000221302-2111122122033121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120332122010202-2121223311213013-1013112202212310-3222102301123103-0103330001012003-1312233100121012-1230322311300333-3123033121032220"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3111210021221031-1333013030220011-2203233303330030-2221133010001302-1323213223311233-1103310232221313-1310131303113323-3212320013221302"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2231133030003110-1323321222321031-2031033330103012-2111312313131300-1203213220032302-2212230023201232-1233113210200320-2013032022033000"></a>

<a id="canonical-2222301032331100-0333300032212013-2211113323232010-1111220211231313-3211013230133130-0111220032210020-0122131323003031-0330233130012001"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1310232213102210-3223031111321233-0020103120301303-1333001210012223-2323320323311300-0333003023022223-3133021032113203-2001230312210112"></a>

<a id="canonical-3113213232301203-2001312222112132-2233303213301010-3311121233201223-3301002122302223-0010122222302300-1201311002001130-3002313202300130"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-1130230013202222-3321303302113110-0031022130203202-3021000132131211-1232111221120003-3111130031312020-2000132202000331-1213100002032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2121322231223222-1010122033302110-3022113010122321-3010323202233031-1001330112223313-2201300230003332-2302130233231133-2310330130230232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013332300332002-3322233323003231-3331023321130030-2011222310100212-3021120131123131-3311320313301111-3003031213212130-0120022013012120"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3232021011321100-1120123303101023-2110102012131222-3222312323213100-1220100232132021-2100310122021321-1323313113300320-0230321331000123"></a>

#### `more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2110120230133213-1123313130111121-3113302322323220-3002013100112031-3002320031310001-3001332232032232-2310312301322203-1213100212331310"></a>

<a id="canonical-2010322033121001-3332213210030210-1023201300223010-3332132123200210-0130132021332322-0122012132220021-1330313221301103-2300301310301103"></a>

#### `more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.request_headers_to_add

<a id="canonical-2130203021120103-0323302203001320-0202002010230312-0313110122330201-2011033231233121-3000321130232300-1133032213022220-3130203220310303"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200321211200331-3133103132130011-1033200100021331-0113301023030102-0211302030220312-1201113212102212-1211013100133303-0213200310202021"></a>

### Direct properties for `more_option.request_headers_to_add`

<a id="canonical-0020212320213223-0320210010000203-1113302231302003-3201131003122113-2110013313100012-2310202012222131-3112201221221130-1210023301320113"></a>

#### `more_option.request_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223110022033123-2212001121200000-1320022222302310-0210001331302030-2020332012212331-2102103103133332-1312310002201120-0313011322323232"></a>

<a id="canonical-1212031330310002-1323032323023212-2112333011033221-2323232000312223-3230311312202303-3200002302032132-1213230201332103-1112103203223130"></a>

#### `more_option.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132): complete subsection reference.

<a id="canonical-1110312211102232-3310202211222101-1131310302321120-2210123012203213-1002322102201300-1013100102331202-0133033013212033-2310000202323203"></a>

<a id="canonical-1211021133301013-1203333213111012-0201213200332200-0120310130020113-1130303303310002-0231030211312021-1003201130220120-0020201333232332"></a>

#### `more_option.request_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- more_option.request_headers_to_add.secret_value

<a id="canonical-3021003333212310-0212103131121332-0310032320203132-2300313111111310-2131230233223102-2301011101313202-1313112032232233-0200303231303300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300210310212123-2233222020333031-3011231230312201-0120221323130010-0210111322312111-3012023132313302-1131130300223130-0023232122122003"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-2010103132210122-2112320332010301-2312113203113220-2020202102020110-3230111012202102-2013032220012301-2101232330313301-0103000202321321): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3132322013300010-2321001232232230-2321232232031202-0223213322231303-1010330010213313-3322200112111211-2310313203111211-1103102212011201): complete subsection reference.

<a id="canonical-2010103132210122-2112320332010301-2312113203113220-2020202102020110-3230111012202102-2013032220012301-2101232330313301-0103000202321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3011200322113002-2233131122133300-2233021121312112-2020310331002321-2233121212301323-3233212232131113-3313321031331223-0313022021203303"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210131302333200-0322030000113303-2333311001000220-3313032312330003-3113132003120030-3230312320022110-3103023102113332-2020101122310020"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3132133012011212-3333201101223033-1333200011013200-2211112103220322-1312332230210223-3123021012320221-1103233230221022-0000120120023122"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0003322130212200-0112132022121220-2310300120211131-0002100130321313-3203023101031311-3003322031012211-3112103032213222-2211202233033113"></a>

<a id="canonical-1103112203110333-0303332300013231-1202333031013320-3102110133211032-3111203221331023-1311023120231112-2031013103030321-2311333210031003"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2000213321302213-3001121301010313-2111102023221122-1101120110300223-2312011102023000-1023120101300300-0101101103202332-0030122003200200"></a>

<a id="canonical-3201113002300202-2233133121131130-2030030032111312-0212323030003033-1333132333322003-0002301222011103-1103123221032300-1331211120200100"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-3132322013300010-2321001232232230-2321232232031202-0223213322231303-1010330010213313-3322200112111211-2310313203111211-1103102212011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2133110200023110-2120222330031220-0301100112232023-3201102122222321-0110322323222232-1202120200122332-1303112221121031-2330310211231230"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310330121210132-0010212300220133-2121302110111130-1232032132203312-2021223030121301-1330120033230203-1312021102123120-3230200000102133"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-2011322302010102-0303102020100213-0032320013131201-2231012211021131-3103233102113033-3103221102001020-3302002011230212-2033223322213121"></a>

#### `more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1203311033303300-1012122233031113-1312310012103010-1300331312100102-0331221001103313-0101300212020310-0122132211301113-0002103211312103"></a>

<a id="canonical-3132120211310012-0230232233000030-2030303221233202-2000033213210211-0030313322102203-0212132212211213-3010120031233331-1231200131133103"></a>

#### `more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.response_cookies_to_add

<a id="canonical-2210232100010011-1111213113122000-0233322020322012-2233232312233130-2131020222212133-1023312011122300-2300021132030302-0231100333212010"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110303323220233-2231211211333003-0020033013110001-2201312031021310-3130312212102122-0003120111313321-3311120121010320-2123220223311032"></a>

### Direct properties for `more_option.response_cookies_to_add`

<a id="canonical-0130232302120322-2223013233211002-3232031032002011-1320311111012221-0003301212020101-1121233322232321-1212102232120200-2013200120230011"></a>

#### `more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0200300200321301-1021031133310101-0010323230003021-1020033320110112-3312313122230021-1102113210322113-0001201110022232-2331130200131112"></a>

<a id="canonical-1131120132203003-0131101002023213-2322210021013102-0230121002002020-0321332011323231-0232302021200222-0323130120220031-2100003123212223"></a>

#### `more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-1002223003312222-3323120332331013-2110013313012200-0310311031032200-3120310230321321-2220220120200110-2331000210131301-2200301302230222): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-022.md#canonical-3032202112123222-1203330201103322-2201301210111202-2103213021030130-0312313201233030-0113100302023230-1320020002033131-3220030032003230): complete subsection reference.

<a id="canonical-1232102210233232-2012221103202222-0012133213010020-1212301033221013-3222312111212023-0303220001233332-0121020333132100-0303010021103311"></a>

<a id="canonical-1100233112330313-1322130103120331-1220311123332133-0310032003303002-3233201023303221-1110031211013331-3130122211213112-3121321030303330"></a>

#### `more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_secure](resources--http_loadbalancer--reference--group-022.md#canonical-3122323312222233-0123020120213002-2320121333022110-1012310032330111-3110330322133130-2201003000130212-1113120210201203-1101300320213223): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-022.md#canonical-0202333311301131-2012001021033211-3321233012130301-1030202123133202-1031123031220323-0001212200123130-2103021111331113-1323233233032022): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-022.md#canonical-3103210013332130-0130102020120301-1322322331232231-0001030203212102-2303200111200013-2031032121300230-1022231000230322-3112321032220321): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-022.md#canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-022.md#canonical-1230032003021303-2222103313010210-3122331222222303-2003030321011123-3020302131020021-0213101112230101-0231203012321100-3011222032123121): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-022.md#canonical-0020113202110313-3213001200332002-0210220031101122-1111021332233111-1200001221122212-2211013021230023-2100312112131202-2110003213010331): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-022.md#canonical-2203100133020033-3302330110301230-2033023031033333-3202331011111133-3131213212320220-0023330110030133-3312131320203330-0100021001110101): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-022.md#canonical-3102202303211003-2113120021232103-2130232022121211-2212031113323321-3122201000122330-0112011030100001-3100031111123032-1032210030321010): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-022.md#canonical-0002020303200233-3101302111100323-3200302300131111-0013102123133200-3010102210020310-2123032023311011-3323300311033131-3330012021231230): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-022.md#canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110): complete subsection reference.

<a id="canonical-1221020110032102-2121230213101211-1003033010300232-1011013000303013-2332333311310100-2320201202103221-3302032312133102-2013231331212110"></a>

<a id="canonical-0130001202300022-3010322013112103-2011300031220211-2020012130113033-3121022013033331-3230230121201213-1200000320131320-0030311102331123"></a>

#### `more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-1221221012100330-3301013232301110-3111033133030120-1211233022122112-2231332112103233-1003013323030330-0012230332221330-1311133020023031"></a>

<a id="canonical-0302103101323003-1012130233233110-1110102130301200-2331230010233222-2230223110213123-3013311211301311-0102213212122301-0203003020320123"></a>

#### `more_option.response_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0230221012220033-1001000131103220-1322330232031102-1120112321012320-3123022213001100-1330023002100303-3231222003312013-0322100310303210"></a>

<a id="canonical-2123000130213023-2123320023301303-3112331002221123-0000113120033021-0213120211223130-3123121133132111-0230330300310313-1201203131001323"></a>

#### `more_option.response_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-022.md#canonical-2330303333211221-2202033113332220-1120033202332121-0320332101312012-2322102332223022-1322012331213210-0001230321211032-1111320222330013): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-022.md#canonical-1332120002013122-1133230013132220-3110131310111110-3221333313023120-1213021213132300-2020010202123201-1232130222012003-1213113011012213): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-022.md#canonical-3200311123303121-3223302333130102-2013211221321310-3331031100310202-1011011132210011-1311002320131212-2320111013112103-0311130332202121): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301): complete subsection reference.

<a id="canonical-0200102003030300-3133312330100212-0222100230323310-3032133001220023-0101003202010012-0122011133300122-2313322012311120-1132301323221132"></a>

<a id="canonical-3100031201002300-1022131011312021-0021121213323233-2020030103331101-0233101031111100-2033132030230130-0102321130131300-2300221100330331"></a>

#### `more_option.response_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1002223003312222-3323120332331013-2110013313012200-0310311031032200-3120310230321321-2220220120200110-2331000210131301-2200301302230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_httponly

<a id="canonical-2012231221230031-1033312213213311-1011332301011313-1000333300222102-0333312201323233-2212112013202102-3132010332021033-0130032302221213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032202112123222-1203330201103322-2201301210111202-2103213021030130-0312313201233030-0113100302023230-1320020002033131-3220030032003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_partitioned

<a id="canonical-3331113321300300-1231033130101032-2012230122101233-1213011313213312-2230000321323301-3030030212000033-0211033132320033-3210102312130211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122323312222233-0123020120213002-2320121333022110-1012310032330111-3110330322133130-2201003000130212-1113120210201203-1101300320213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_secure

<a id="canonical-1302010002030101-3221023300011123-1111301000332113-0122212111312223-2231211220312113-2012222123110200-0122330122303020-0021300011301230"></a>

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
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202333311301131-2012001021033211-3321233012130301-1030202123133202-1031123031220323-0001212200123130-2103021111331113-1323233233032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_domain

<a id="canonical-2031120102133221-1330300213221023-3122133100201330-3102223311210032-0301321203002022-2300223321102303-3313302020133220-2021132103313131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103210013332130-0130102020120301-1322322331232231-0001030203212102-2303200111200013-2031032121300230-1022231000230322-3112321032220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-0313102123210110-0230222123323330-0222301022112231-0120123023032000-3120221031001100-0131201201111203-0331022000120110-2210130102230122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-1320132312121302-1302210133321031-1220333311001332-3133010020113231-0122222023112200-3200130003223233-1313223101023021-2031013030210303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230032003021303-2222103313010210-3122331222222303-2003030321011123-3020302131020021-0213101112230101-0231203012321100-3011222032123121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1110023003123313-0231110112022103-1113103332000112-3121002030331232-1303222013032013-0001200022133313-1023312310030322-0223022002003323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020113202110313-3213001200332002-0210220031101122-1111021332233111-1200001221122212-2211013021230023-2100312112131202-2110003213010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-1002210120110013-0121223220223223-0221100210312310-2013210233312213-1022321133101022-3312032020211011-3333231222010020-3313032323231333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203100133020033-3302330110301230-2033023031033333-3202331011111133-3131213212320220-0023330110030133-3312131320203330-0100021001110101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_path

<a id="canonical-0211300202122213-2111120002313001-0120113102300100-2332011121101122-2020133231000010-2311323322022232-2111222110000320-0222303133133230"></a>

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
ignore_path = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102202303211003-2113120021232103-2130232022121211-2212031113323321-3122201000122330-0112011030100001-3100031111123032-1032210030321010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3030113033021231-2331301211302011-2202133021013201-3112123001132010-3312112233123333-3102110332022131-3210102222033111-2002121212231020"></a>

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
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002020303200233-3101302111100323-3200302300131111-0013102123133200-3010102210020310-2123032023311011-3323300311033131-3330012021231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_secure

<a id="canonical-0321323310001023-3011033333322330-2123131221013320-0002332020333000-2331102100333322-2120203133232032-1202203101133122-2332323002030313"></a>

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
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_value

<a id="canonical-3303310332300023-2330121133120201-2303303032220230-0033333301212021-0301200210131103-0311102002332310-3130312122021030-0310102312132031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330303333211221-2202033113332220-1120033202332121-0320332101312012-2322102332223022-1322012331213210-0001230321211032-1111320222330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2221220202231312-1222332333213110-0000203132332230-1323003010302323-2001111023122201-0303203220303032-0302331001033020-1203131302312030"></a>

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
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332120002013122-1133230013132220-3110131310111110-3221333313023120-1213021213132300-2020010202123201-1232130222012003-1213113011012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_none

<a id="canonical-0030223113031120-0323122102230123-2122303023023120-2103302110032103-0132201112313110-3110011321311023-1111312010213003-2102121032321120"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200311123303121-3223302333130102-2013211221321310-3331031100310202-1011011132210011-1311002320131212-2320111013112103-0311130332202121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_strict

<a id="canonical-3031031131230122-0133032102123202-2133220322213101-3120131322101112-0022130223021200-2030001101100222-3111213133121203-2122100330310222"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.secret_value

<a id="canonical-2120210132333100-3201302311113121-0121212310132213-2222320233211302-3221113013201033-3202232012313031-1121212033113012-0203200320032121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321221100200112-0000211331020011-2030202002201021-2010031311123121-3231312330202311-3111321110322103-2311212233001100-2321011030100012"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3021231033011023-1011212223030023-2030200113003131-3322130011212001-1111133222233230-3112130301022221-0112233203021013-3000022003112333): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-1123201200031101-2233032321030232-1300131011211233-1133230233030101-0232031323331322-0020221323211110-3023231133130300-2112201123133020): complete subsection reference.

<a id="canonical-3021231033011023-1011212223030023-2030200113003131-3322130011212001-1111133222233230-3112130301022221-0112233203021013-3000022003112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3102233100020032-0311321100223313-0200333101111312-1113130223221332-1231132233103133-3300200312302123-3202310323123000-2333013232200201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122300302201101-1210322211032123-1102310022132212-2003120002332231-2133312012020321-0210130210320011-2010022222023133-3331321332120313"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2012010011030123-1000000122000012-3001110202022333-3321022113032110-1033113100000010-1133013021220123-1230001331112223-1232322202033001"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0201203233232220-0120323232113000-0221002322020123-0013112303332231-3300021332332312-2110233303101202-0132123230032030-2230212121011211"></a>

<a id="canonical-1213301023323302-2313101010112322-1312030312310000-1130211012122003-2222123100322011-2232302322021201-0031322100200300-3301112031130202"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2232212101231013-1231330012122003-3301330122100220-1102121023233022-2100131001230300-1011302320333312-0031312232200102-1311103101221321"></a>

<a id="canonical-3023221320310332-1003310331211031-0031132000313233-2120023001222131-3113200110213102-1003033030132000-3231010233012320-0130231020102320"></a>

#### `more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-1123201200031101-2233032321030232-1300131011211233-1133230233030101-0232031323331322-0020221323211110-3023231133130300-2112201123133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2200200122310020-0110310232232021-1222333032121113-2132113310203121-3231301322122303-2021200233110202-0132211223101330-3021123303122012"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313230223121022-0122310233030320-2123322112321320-0323311130031133-0032021123303233-2010331321203300-0123012232011032-0110011203313230"></a>

### Direct properties for `more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0103101013031023-2223221001221333-2320310332223023-0310103113213321-0001132023021032-0012000302101220-1110213011222113-1202223013221311"></a>

#### `more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3021311113323321-0302310310303310-3120333223231113-3310133230023021-0100100332123212-2122201132300121-1302123311221211-3302030223320122"></a>

<a id="canonical-0323101230210203-2101132033331112-3023012200310231-2022303213320022-0023201121321231-3322130300103221-2010111331202112-2201130020002131"></a>

#### `more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.response_headers_to_add

<a id="canonical-1301103200200321-0000011321231201-3311302210031330-2023130120232313-3330022223021122-3120120123122323-3031001001031312-1330002332211003"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302222212011212-2032201101300213-2103130300212131-0310331212201000-1102112221303310-3131033133131103-0031201133121111-1020311022333002"></a>

### Direct properties for `more_option.response_headers_to_add`

<a id="canonical-0203310022313333-2203022002300223-0323300122003123-2301111133311122-2200103322330230-3331132133323123-0221000323101013-1100202200010023"></a>

#### `more_option.response_headers_to_add.append` property

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3013113212021310-2110330122331033-2010003233130013-2033121123123012-0303312331030102-3021012112212011-2310231110102231-1310121110200321"></a>

<a id="canonical-3132201233221033-2223303322023230-3133230212100112-1312013311231000-0221022302131033-3201030311302011-0122223132101001-2132222330123133"></a>

#### `more_option.response_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231): complete subsection reference.

<a id="canonical-0121233013012010-1022322011320020-3212313112132130-2111210020030302-1012223320220102-3322013023202033-2310103130331101-1110231233321231"></a>

<a id="canonical-2001012103120310-3230023031130223-3101131321022223-0001223011203110-2013100203203220-1100333130103033-0101120020020120-0121103332302200"></a>

#### `more_option.response_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- more_option.response_headers_to_add.secret_value

<a id="canonical-1131232202121221-3231323222231302-2103322102032323-0133333312313200-2323132212030032-2332330130303111-0220021002023311-2100301332323122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313120003102231-3222103300332101-1101100310300302-1202031332013111-2322333221322233-0310330231033021-0301030322322132-3113301213022022"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3030120032231000-2100231032120312-1320230213332233-2212100020132121-3221000312202102-1022120113330231-1330223333003232-2321101101211320): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3021201211021111-2000322133100203-1322031332223201-0321212023100121-3232211031301120-3100222010222211-2210332102032010-0321331313321322): complete subsection reference.

<a id="canonical-3030120032231000-2100231032120312-1320230213332233-2212100020132121-3221000312202102-1022120113330231-1330223333003232-2321101101211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2202300123231233-0232013101031001-3123212121000011-0020131022111121-2203302300013212-0122023313132112-2313121123111223-1321230322131013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113222102120003-0001131131233002-2223211033011230-1011230321333222-0201302300313002-1022131333313010-2013223131122312-0233101333230202"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0310231002322303-1003221333031200-0113132313111300-0112030302100021-1133000103310101-1111122321011333-0212002011031103-3213232130312203"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-2213311001122111-2331321303030202-2201330000103321-0121311012320330-2131220123132011-1112001123232023-0213033031221330-3103213101021110"></a>

<a id="canonical-2230223200133113-2123310303103211-0333011201130103-2301223021332000-3331100323123332-3200300311030312-0020212213211020-0113212330113110"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3333103000232120-1311231301113020-2133203001120212-3211011232233301-3113333110303232-2313200323022232-2002200312121212-0320101011330211"></a>

<a id="canonical-1131111313200202-0132301032033231-2220232221302012-0321130031131202-1101102300333033-1101310220200103-1133002310332211-3223130212222200"></a>

#### `more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-3021201211021111-2000322133100203-1322031332223201-0321212023100121-3232211031301120-3100222010222211-2210332102032010-0321331313321322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0122121332321113-3332021300302123-0211022203233203-0230202310210031-1231131213011003-1211031133000202-1131013100030221-0022323321311100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111310030221113-3000103300023033-3200212202030300-3232332223210330-2002122013230113-3220020003212122-2021232301000102-1213033212020223"></a>

### Direct properties for `more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1113010312122311-2122030130220232-3331320100200100-3312033111322101-0202211322200012-1131003033120020-3010122010111000-1220020332312103"></a>

#### `more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3212013113003213-0022020202012330-2202312201331313-2002020121101202-1332033123102003-1010122003200113-0221301223301213-2201023112233103"></a>

<a id="canonical-3012201100122031-0202030132030011-1311112030013112-1032020323232031-3011333012132102-0011003302103221-0231221021313330-3221211220302011"></a>

#### `more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0022320220230323-0223203133103320-3032023020320010-3000011312022000-3332213021010220-3323011033022102-3132320201321311-0310230130132011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `multi_lb_app` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- multi_lb_app

<a id="canonical-1113010133303031-1030321102031003-3030322233333020-0132110322111130-3302030002322023-3231011110103303-3222330333330311-1020023031330321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

OneOf alternatives in this subsection:

- [multi_lb_app](resources--http_loadbalancer--reference--group-022.md#canonical-1113010133303031-1030321102031003-3030322233333020-0132110322111130-3302030002322023-3231011110103303-3222330333330311-1020023031330321)
- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1201323130312033-1332212010203222-1112131130000000-3313131233123132-0112201202300002-2122020103032101-3012032203220001-1123233022300133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
multi_lb_app = {}
```

This is an empty object or choice marker. It has no direct properties.
