---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-2212011320322023-1101022131130001-0130300302013222-0330232202110332-2200130332111302-2202221323012132-2210301212211213-2020022130110130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.compression_params` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.compression_params

<a id="canonical-0313033112003111-2131131313202310-0310031032122100-2303211320132223-3331103232121323-0320331222212100-2023100122002211-0131301333121303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3212132110233033-0213123333110203-1110320231032210-2112200122223101-3110130131313131-0321011020223101-0001123313012020-0223200322300013"></a>

### Direct properties for `http_proxy.more_option.compression_params`

<a id="canonical-1322130002332112-0230232123110323-0232012131331333-0100113123000233-3223220121133023-3221332000303330-0222120232200220-0010211030001202"></a>

#### `http_proxy.more_option.compression_params.content_length` property

Type: `"number"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1031020310231023-2322010012221120-0122303310130022-1030010331301120-1133032310101000-1023311202121200-0111110210231122-1010130120110311"></a>

<a id="canonical-2033002312102123-2300303032220002-3022333230323123-1111300123123211-1323000121201023-3110002113113201-2110010001320213-0323123232030232"></a>

#### `http_proxy.more_option.compression_params.content_type` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-3032003032003313-1021313033000233-3022022031131121-1130121201010010-2120001210120200-1103332211201001-3122211111330320-3021010222233131"></a>

<a id="canonical-1100121010012212-0210132323110132-2212322113303222-1301021101111000-0033032132201011-1130202213010312-3131320213200103-2231113233323202"></a>

#### `http_proxy.more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Computed.

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

<a id="canonical-2232202130012220-3333231323331001-2021221313101122-3323331312111021-0303031013333310-1011320002012113-1122202301333201-2202113313321213"></a>

<a id="canonical-0013102031322120-1002121313330321-2233103121323201-2031002002111230-1131113033330112-2333202210121033-1303220111311121-2012100003301011"></a>

#### `http_proxy.more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Computed.

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

<a id="canonical-1003210011230331-2323213330220012-2132303121131103-2001310000033121-2320121032032302-0201100310313031-0202303321302310-3110013120231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.disable_path_normalize

<a id="canonical-1212020001203000-1002120131000330-1130002202233320-0220111300223101-3020222232023222-0120311112213310-2332113122313223-3112102232123113"></a>

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

<a id="canonical-1222133101111030-2231302230300323-1002110332333200-1212330122133323-2100222233003123-1011133001010013-3121103331320101-0331233320102230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.enable_path_normalize

<a id="canonical-3020131031130323-1003333333013013-0031312111203230-0310332113001133-2202023111001233-2121033023033110-3300311321320120-1112233200102232"></a>

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

<a id="canonical-1202323110231213-1113003323031203-2323321001333030-3110301331112323-0100221331123110-1311231300211313-0300021121133003-2122123200122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-3110332221023003-0010012021211113-2023301200130303-3222120211012021-0331212112013313-3113232233010100-2210230223231130-0210331222132133"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.request_cookies_to_add

<a id="canonical-1113011033320100-3100121032123121-3233031121013102-1031130132232213-3202123333300222-2301201301130112-2323212312033203-3233220231201101"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2000300201330212-0333013122201131-2023322220011003-1313301303120212-2212110131231321-3302203110022121-3312103200122201-1220033310113000"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add`

<a id="canonical-3033000322330000-1213232200133320-1211122301003003-2012113300231112-0010211301200130-2020310332110233-0000222021201331-1021102222002211"></a>

#### `http_proxy.more_option.request_cookies_to_add.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1010232220302003-3220202200031323-2333002213123010-1111103100323102-0023211032023033-3032300020132311-1321032231030323-0311103111213130"></a>

<a id="canonical-1331133222123213-2332303303301101-3233303022101021-0233322101313021-0303302231221003-1312230122032021-2220112011301103-1010312333211110"></a>

#### `http_proxy.more_option.request_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

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

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-3203332320020120-3133031132312222-1302112213013230-1100111123331300-1320310210110012-1022120032303310-2330003333232203-2211320331011323): complete subsection reference.

<a id="canonical-3330013233111301-1202200232133333-3200333322223231-2103012312102013-0033022330200313-1303221313023312-1311330233111200-3022222221010100"></a>

<a id="canonical-3131230033000310-0331232323333223-0313132202223233-0330011011313020-0220212120311323-0202102311002020-1030102000100013-1233003231313231"></a>

#### `http_proxy.more_option.request_cookies_to_add.value` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3203332320020120-3133031132312222-1302112213013230-1100111123331300-1320310210110012-1022120032303310-2330003333232203-2211320331011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213)
- http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0203033032303303-0123120312203130-2013232002310203-3222023000210323-1131022313020101-0121200002133200-0231130033132022-0330102113212001"></a>

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

<a id="canonical-0030230123330313-3033110232133201-3220110132110130-0031002322211002-1323121201300322-3123221310030101-1010322303302013-0330200323112332"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-3222330112013202-1201101123310031-0031012002100210-3113323030203100-3331101213221231-3010231230131003-2003313100131110-3000201003030133): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-3112120022323310-0030201012233223-2203223000211320-2020033120300223-0301231003331213-1112322002332030-3120220122312221-2131303003110212): complete subsection reference.

<a id="canonical-3222330112013202-1201101123310031-0031012002100210-3113323030203100-3331101213221231-3010231230131003-2003313100131110-3000201003030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213)
- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-3203332320020120-3133031132312222-1302112213013230-1100111123331300-1320310210110012-1022120032303310-2330003333232203-2211320331011323)
- http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0200033012003323-0013322130300100-2011011013101202-2030222302131113-1322231221003333-2331001230300201-0300301033133003-0210032112021113"></a>

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

<a id="canonical-1220112111131012-2230102211221223-3131031200230111-3313212200110133-3121130233121210-0101302121213332-2323301232130001-1203203213210232"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0130001331303131-3333310333203120-1322212022213130-0303102001333301-1031330002211220-0111003230132022-0301301111000313-2022102201300002"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1011332010303003-1213312121211321-2213202322201313-1303123330021003-3212003320002323-2323232112221330-2100213033212203-2003231333001123"></a>

<a id="canonical-1000011202130133-0322303010023210-1001200210020002-1302120203002020-0021102013013123-3301030232203223-3030000132003113-0331300120133232"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-1022310330123320-3103020123313330-0022110212010001-3013022123332200-1202230030102300-1200112313001113-3303022131122333-0233003220000301"></a>

<a id="canonical-3122000102312012-1030101132033300-1121231320110320-0232022213320300-3102022012003012-0210102120221110-2300123201111233-1000231310201222"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3112120022323310-0030201012233223-2203223000211320-2020033120300223-0301231003331213-1112322002332030-3120220122312221-2131303003110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213)
- [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-3203332320020120-3133031132312222-1302112213013230-1100111123331300-1320310210110012-1022120032303310-2330003333232203-2211320331011323)
- http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0000210223011120-1130233221012022-0103123200122123-0120213303032301-1023223003121300-1112333313013233-3021322210032112-2301231022302133"></a>

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

<a id="canonical-0133210223030231-2202210230121133-1032222210200110-0330332021102122-0031321230021130-1113333023021000-3002003212010103-1101110130213232"></a>

### Direct properties for `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0001030300322200-0313002133010203-2231111202212031-3232101310222221-2333220111102221-3310310123131130-0133012201133201-1001220332131100"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3030303211211211-3310101212230233-1323330200201221-0110023032201002-2110303130301030-3121300313010021-1302203002233002-1231000133023210"></a>

<a id="canonical-1232233113312221-1102101102110331-0113122221313322-0022033303213231-2302030020132130-0130320010330213-2103321101101312-1320020131033031"></a>

#### `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-2213012201232033-0321330202023321-2330031333303321-2113210102231203-3232031210002002-2201101202033323-0031022310111323-3332302313022022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.request_headers_to_add

<a id="canonical-0012103111211323-0333112200113131-1032002301331130-2023212131232012-1232300131123303-1311330302210322-0321330323130231-3131131131011032"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1122200332001311-2031200013321331-0310103013012211-0111121201031332-0030111033203022-3303033322233100-2211021333220101-0332232121202312"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add`

<a id="canonical-1323003012311023-0210113211212233-0222203202210333-0112122001121002-2100000100321232-1302112301112200-2322323132321312-3212313202332000"></a>

#### `http_proxy.more_option.request_headers_to_add.append` property

Type: `"bool"`. Computed.

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

<a id="canonical-1131320000020230-0000123322100013-0320000230000302-2230133212233030-1322301012111220-3033132320232201-0132110100021333-0223102033200231"></a>

<a id="canonical-0332122230331230-0120223230012010-3123032103133021-1232000123321333-2121213301012312-3022023130001012-3212110323031212-2132110230002220"></a>

#### `http_proxy.more_option.request_headers_to_add.name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-2233023120020123-2320103321331302-0220020102012332-2011131023113021-3022301330332321-0112220323010011-2312121333223331-2002300021010112): complete subsection reference.

<a id="canonical-3213312202203033-3012220023103210-1231122133122220-2010023103100002-3312303002000020-2332023010202030-0213132213022211-3022211123321212"></a>

<a id="canonical-2303000231221332-0032031131000131-0302311200121212-1333231010312021-2332231032221010-1010123132300002-0013101113010112-0011012210103013"></a>

#### `http_proxy.more_option.request_headers_to_add.value` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-2233023120020123-2320103321331302-0220020102012332-2011131023113021-3022301330332321-0112220323010011-2312121333223331-2002300021010112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-2213012201232033-0321330202023321-2330031333303321-2113210102231203-3232031210002002-2201101202033323-0031022310111323-3332302313022022)
- http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-0012133331210231-1033120113223323-1020300211202120-1313201130120200-2212211100022100-1201223311031002-1100100013211013-3032120110200030"></a>

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

<a id="canonical-0130321320302122-0020310000011300-3030021022212310-0120020121213132-1210002002100010-2323312132322010-3222012302211223-1103203113200301"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-2322103023021230-2333021202300121-0100132113123123-3312330022332213-2321300223223031-2332132020130222-2103032003310212-3200011301132213): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-3220010110020022-0123211300311233-3020120222322330-0032210211000110-1230032202330032-2130320021123320-3203001231130332-0100100113013323): complete subsection reference.

<a id="canonical-2322103023021230-2333021202300121-0100132113123123-3312330022332213-2321300223223031-2332132020130222-2103032003310212-3200011301132213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-2213012201232033-0321330202023321-2330031333303321-2113210102231203-3232031210002002-2201101202033323-0031022310111323-3332302313022022)
- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-2233023120020123-2320103321331302-0220020102012332-2011131023113021-3022301330332321-0112220323010011-2312121333223331-2002300021010112)
- http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0100003012311202-0013222012310310-3020211233233001-0033110201312022-3120010111013032-1321122101132033-2112321030011213-0222013311131203"></a>

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

<a id="canonical-1233321332010030-0231033103120132-1122322121233130-1223001313323321-0110201001310202-1223020120111012-2113312022313302-1131310133131133"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0011311022333030-1202323122211012-0320220320201000-0011121233233031-2213112221312031-0212311102201333-3320230013102301-0323221020330010"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0300011331133010-3202023310003323-0322221003210101-0313121030321020-2033333332223003-1230010332301202-1231320203121001-2220031321221101"></a>

<a id="canonical-0112121331001031-1312231221322000-3013212010203333-1300330330303032-0300022101131121-2012202122030002-3010200320212001-0130213132102213"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3320023322012332-1111321211222023-0003202323023233-2211230333100020-1111021033023222-2002331032302212-2110331002313000-1033330223131223"></a>

<a id="canonical-2221333011022133-1120132130111111-1000210303110233-3122112200011112-0213102020110213-0301301331033020-3300333330232331-2312031332102332"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3220010110020022-0123211300311233-3020120222322330-0032210211000110-1230032202330032-2130320021123320-3203001231130332-0100100113013323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-2213012201232033-0321330202023321-2330031333303321-2113210102231203-3232031210002002-2201101202033323-0031022310111323-3332302313022022)
- [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-2233023120020123-2320103321331302-0220020102012332-2011131023113021-3022301330332321-0112220323010011-2312121333223331-2002300021010112)
- http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0201211033211200-3200310322011212-3011133230330231-2230323013132020-2303221112211222-1020302302121231-1003103031301023-2322313032302302"></a>

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

<a id="canonical-2123003202321031-2203121113103002-3200021333320223-2012131211001203-1303333012113331-3200311033310131-1010022131130330-3031313130232201"></a>

### Direct properties for `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0010212122022320-2132131302032203-2013122030213213-3023132233231111-3112211022120030-2211233031122122-3112212300100202-3001310133111102"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1233320022020013-3202233331231313-0301313303133223-0130320311003332-3312302023222312-2232023011020130-0311011320030333-0121131223101220"></a>

<a id="canonical-0100022212203200-0022000313202120-3023032302033123-1210032300203203-0133132013103300-1030213000302101-1330032323211310-1213001022313011"></a>

#### `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.response_cookies_to_add

<a id="canonical-1200123032030133-1311012231221230-0103320121233100-3311001301023020-2102110110301231-1013312113122003-3203323300313112-1012323010231313"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320332310302101-0002330303111321-3201131303001222-2133300102333220-3003303122111322-0323103331101210-1330002212132030-2203032330332120"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add`

<a id="canonical-2113231012001031-1130301311113232-0122110331213320-3232210320202200-2223321022132333-2100313030023203-3333231020002011-3230332320333203"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Computed.

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

<a id="canonical-0210030111020122-3202120112000020-2010131310221301-1102231220321010-1332102121213102-0210123110211011-3111111122001211-2121032230033322"></a>

<a id="canonical-2103331210110111-3232021100210101-2210101332300020-0121301132111211-0013303031011000-0100130010130300-2123112130023102-3202200321011001"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Computed.

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

- [add_httponly](data-sources--proxy--reference--group-004.md#canonical-0030232003220300-3213011023101101-1121020013132121-2021120311201030-2230031201100031-1100222110003202-1103220122023020-0200210132033022): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-004.md#canonical-2311213122101130-1131333233303123-1000133320121213-3322232121120331-1133113202101321-0000033130030002-0331331021322112-3132312221320002): complete subsection reference.

<a id="canonical-0313033301121032-3003322310301321-2223112230312313-1132203121031112-0233302012132303-3230233300310011-0111130133003333-0320330211301323"></a>

<a id="canonical-1313301130031231-2032201301111200-2103331012221000-2221132332021022-2033200002032021-0201023313202213-1033220303310220-2102121301302130"></a>

#### `http_proxy.more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Computed.

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

- [add_secure](data-sources--proxy--reference--group-004.md#canonical-1103030202320220-3213030030021211-3212323322233110-3021221132122003-2002002022303202-1110331312333230-1203302020112223-2023303321001232): complete subsection reference.

- [ignore_domain](data-sources--proxy--reference--group-004.md#canonical-2013330303122211-3312303302313230-0131223031310230-3111003021112222-1212003011011200-2232313100201122-0221332110003112-3222312121323231): complete subsection reference.

- [ignore_expiry](data-sources--proxy--reference--group-004.md#canonical-3303313210121012-2213311203123310-2003110312220201-3100301101221102-2120221030123233-3202201101001110-3201231221120122-2202231300030110): complete subsection reference.

- [ignore_httponly](data-sources--proxy--reference--group-004.md#canonical-2103022130032101-0131332231320333-0123033230021200-2001113012023012-0022223102112032-3002130322030121-3033323113313323-3103321211123010): complete subsection reference.

- [ignore_max_age](data-sources--proxy--reference--group-004.md#canonical-0021020310013002-2120133201010013-2121120133112331-1131221333123132-3120103202331032-3001301232313011-3123121220121121-3222210032001202): complete subsection reference.

- [ignore_partitioned](data-sources--proxy--reference--group-004.md#canonical-3033130001220320-2221132002112103-3303032102203111-3121323000103032-0222001323102133-3232113020200302-1023330030011132-2311033020300023): complete subsection reference.

- [ignore_path](data-sources--proxy--reference--group-004.md#canonical-1321211302212311-3121011003123031-1203023103011010-1023030010213300-0313212001020311-2102331011002321-2332322301013312-2233023103230330): complete subsection reference.

- [ignore_samesite](data-sources--proxy--reference--group-004.md#canonical-0003313213211331-0131201030032031-3200003013023330-0330222032232313-3211122111021212-3020302033131222-1333233200310012-1322301031030232): complete subsection reference.

- [ignore_secure](data-sources--proxy--reference--group-004.md#canonical-1210313321122210-1000300110312100-2322311223200210-3302003010001220-2221221300031123-1002122121303210-0123332010030033-0012332011213302): complete subsection reference.

- [ignore_value](data-sources--proxy--reference--group-004.md#canonical-3020330330330330-2123200323322210-3323132312133232-2030303013202102-1313323221221011-2013101323233031-3330323213230102-0111322210131211): complete subsection reference.

<a id="canonical-3230100330101132-1031010002311301-3103300123310201-2123313031101131-3230002331032101-0223003130132233-1321000112101333-0330003010220032"></a>

<a id="canonical-3111211333030023-3121021312003301-3311310320100031-0331003312132012-0210211121132233-1231313321020131-2002122123333010-0123311332011001"></a>

#### `http_proxy.more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Computed.

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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-2203223121231210-1220102232310303-3213302321022123-0033132330132322-2020230210132332-3110302032101121-0011022201331223-2001233103123101"></a>

<a id="canonical-1221001232030322-3323023322010100-3230011320122111-0311332203310213-3320002123131233-3331333022130333-0101313111131312-3322002332212002"></a>

#### `http_proxy.more_option.response_cookies_to_add.name` property

Type: `"string"`. Computed.

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

<a id="canonical-3032333202033300-2131313322330000-1230030000203213-2131013303231223-0223130010002332-1303032330112011-3002131301120320-2232211231202103"></a>

<a id="canonical-2233021113131110-1101321001212230-1133122220230332-0302211103231032-0033133010210123-0022201331300331-3011030323313222-0133210302320213"></a>

#### `http_proxy.more_option.response_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

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

- [samesite_lax](data-sources--proxy--reference--group-004.md#canonical-3300232120300000-1111133031123223-1110103330112220-3032001323111213-3002313330211023-3021310202230032-3331021223212302-1221000112201133): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-004.md#canonical-1120203032212110-3230321133102230-2321213233110312-3313031321210122-0102021131220031-0302121131021031-1001323030013112-1302030320333100): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-004.md#canonical-3030120112022012-1133321221113122-0033023332021112-3010330323222011-0203233130313100-0202021003002310-0310330022212112-0021330122213131): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-0030103320211303-1032001020222132-2120221300121030-3112100101313210-0323003333232010-2322202231303202-3301302013322211-0320103012123331): complete subsection reference.

<a id="canonical-0222030122213021-2312121333100133-2030013121322131-0323120211101332-0122310123110303-1000031130120031-1230211233002210-1312112102030211"></a>

<a id="canonical-1113233130103000-1231213312201302-2322232103210222-2101220230012330-3123212123222103-3020112310202332-1033121032202200-0000033130103012"></a>

#### `http_proxy.more_option.response_cookies_to_add.value` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0030232003220300-3213011023101101-1121020013132121-2021120311201030-2230031201100031-1100222110003202-1103220122023020-0200210132033022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-1331002022203331-1033333310320033-2001201311210010-0321032103111211-1233220311111223-1300102223123111-1020030100011002-1201021123003332"></a>

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

<a id="canonical-2311213122101130-1131333233303123-1000133320121213-3322232121120331-1133113202101321-0000033130030002-0331331021322112-3132312221320002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-0020020021322301-2323112330110333-3132031320131301-3030112203220102-0031201010023003-3110233333200022-0221330323311020-2120023213311032"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103030202320220-3213030030021211-3212323322233110-3021221132122003-2002002022303202-1110331312333230-1203302020112223-2023303321001232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-3011221200103301-1222210003201331-0311102332120120-0332010031121332-2330031030002322-2102302233131003-1131101110021221-0003032211030222"></a>

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

<a id="canonical-2013330303122211-3312303302313230-0131223031310230-3111003021112222-1212003011011200-2232313100201122-0221332110003112-3222312121323231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1111131023013203-0112033121001130-1301212212221212-1302103223012233-2220332321020130-1233203031130330-2201310022101321-1101322122120332"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303313210121012-2213311203123310-2003110312220201-3100301101221102-2120221030123233-3202201101001110-3201231221120122-2202231300030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-2112200201133031-0033032031113300-0113111013203012-2111302312221220-2330102030030021-1223330101202320-0331322222200213-2011011320112132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103022130032101-0131332231320333-0123033230021200-2001113012023012-0022223102112032-3002130322030121-3033323113313323-3103321211123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-2320132221330003-0011011232132110-3031301320032313-3333311000203302-0000320111201111-1010301020322231-1111232332231212-2302032133120021"></a>

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

<a id="canonical-0021020310013002-2120133201010013-2121120133112331-1131221333123132-3120103202331032-3001301232313011-3123121220121121-3222210032001202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1211321020230120-2322013003030213-1322230202122333-2213203210220112-1330333021111023-2202032130213133-3002103000002213-1221302302203103"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033130001220320-2221132002112103-3303032102203111-3121323000103032-0222001323102133-3232113020200302-1023330030011132-2311033020300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-0103010131332030-3110223133113203-3011011322120231-0011122321333322-2300033132222003-2030311020323002-2011233321303232-2203320330121123"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321211302212311-3121011003123031-1203023103011010-1023030010213300-0313212001020311-2102331011002321-2332322301013312-2233023103230330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-1320001222320213-0310131312002323-2222203332202312-0013003110011120-3321122321101011-0100100131303023-2331223323102011-2011132232032322"></a>

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

<a id="canonical-0003313213211331-0131201030032031-3200003013023330-0330222032232313-3211122111021212-3020302033131222-1333233200310012-1322301031030232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3330101012320310-0132221021103101-3103312213311320-3222131131203223-1102033211003220-1033122103300230-3010010113330131-0030133201101112"></a>

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

<a id="canonical-1210313321122210-1000300110312100-2322311223200210-3302003010001220-2221221300031123-1002122121303210-0123332010030033-0012332011213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-0300001231320332-2132013332301022-3033233003030100-0022222002331003-2012322020211110-0211333113301122-2331010211103100-1210011332222202"></a>

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

<a id="canonical-3020330330330330-2123200323322210-3323132312133232-2030303013202102-1313323221221011-2013101323233031-3330323213230102-0111322210131211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-3021032121332123-1023233222033123-2100321020233211-2102231102321303-1312022012013301-0320031202300021-1320331030223123-0332011332013112"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300232120300000-1111133031123223-1110103330112220-3032001323111213-3002313330211023-3021310202230032-3331021223212302-1221000112201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-0100331312131203-1211220132311333-2102020011232021-1102302030201201-3012203021230301-3133320302123232-0211213222332223-2111310332320202"></a>

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

<a id="canonical-1120203032212110-3230321133102230-2321213233110312-3313031321210122-0102021131220031-0302121131021031-1001323030013112-1302030320333100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-3211112320230211-3202201132110202-3021223013012010-1210232013220030-0331113120020233-1233200131210313-2132121033122230-0030010120303301"></a>

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

<a id="canonical-3030120112022012-1133321221113122-0033023332021112-3010330323222011-0203233130313100-0202021003002310-0310330022212112-0021330122213131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-3310301322003302-0020103110330300-0002202332212232-0301012102233002-0031202130313030-1132323320033023-3201020211003101-0020202033232210"></a>

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

<a id="canonical-0030103320211303-1032001020222132-2120221300121030-3112100101313210-0323003333232010-2322202231303202-3301302013322211-0320103012123331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-0330122331133332-1132112222233101-2120123103120200-2220202012320331-1300301020223333-3101320200202103-0110232323000202-0203212201200230"></a>

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

<a id="canonical-0230221100002100-1203320222020312-3313103300030010-0033320130101313-2220003302300211-1031100332021330-3212012013123103-3222031002122000"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-2310230120112333-1312032023222122-3303100102313302-3021200033230332-1321201330011103-0131320203322322-0211100312103121-0201310031001002): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-1300313212213211-1303020212213332-2103313000010022-0221003231210323-0320002211000000-3001230000012013-0223211033323121-1330022003230212): complete subsection reference.

<a id="canonical-2310230120112333-1312032023222122-3303100102313302-3021200033230332-1321201330011103-0131320203322322-0211100312103121-0201310031001002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0030103320211303-1032001020222132-2120221300121030-3112100101313210-0323003333232010-2322202231303202-3301302013322211-0320103012123331)
- http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1112302211001001-0313120311033130-1202210231033323-0032230231010330-3320321031310221-2233113231020202-0103222123001200-2303202121232020"></a>

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

<a id="canonical-3030331303122232-0332102103200213-3212131031232302-2210201103112130-1110202001031323-2320033122311200-1022200203121022-2110300211130221"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1000213030013100-3000111232000301-2320321212023200-3023312202232030-1132021012120203-1101301011123211-3303331331201231-1000133210232110"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0010102132120333-1231311102311111-0313200000120000-2103013111100103-0320222022101032-3010120112113002-3310312011012300-2332333210130032"></a>

<a id="canonical-0003111021031200-3301211232020201-3003303302132002-1130132302323200-3023023222230212-1301302023231322-0020320202120121-2322200322333021"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3212230120220110-3110221231111303-0212023231133103-0231010023213301-2212220011022113-2102110112023203-2302030201132033-1212033101201310"></a>

<a id="canonical-1013103100113033-0001301022320103-0002100322021223-0320110032230210-1111311001231221-2003102133203210-0001220102202331-3031012232021022"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-1300313212213211-1303020212213332-2103313000010022-0221003231210323-0320002211000000-3001230000012013-0223211033323121-1330022003230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322)
- [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0030103320211303-1032001020222132-2120221300121030-3112100101313210-0323003333232010-2322202231303202-3301302013322211-0320103012123331)
- http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2330320131133003-2203130001113011-3331301101120330-1222133301311212-2120110300213032-2333113003012100-3110101212103220-3102123322033130"></a>

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

<a id="canonical-1231010231331211-2000332300132023-2232130301013300-3323130302020312-1322020122231302-0002003131323112-1313010322110102-0010101013220220"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-1032001023333022-3103110001101321-2131211302311222-1333111211031212-2312012122212122-0100010302221220-0300300122000212-3120201130313012"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2033131202313003-3203211110032230-1333221332333201-0010332323103122-2101012111301311-2200300112100133-0221223020311032-2023131020023100"></a>

<a id="canonical-1200000332311332-1310113123130320-1013223123313100-3120120000321031-0000132012221010-1112030332021001-0302030201200310-0030111103023300"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-3220133020330022-1232303102011310-1113033131110013-3100220200031212-3021203311201313-2122021321110320-0110232121302032-2311132322032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.response_headers_to_add

<a id="canonical-2300103121122010-3001301033010231-0133213333231212-3101121230301122-1330123230021300-1310022212311120-2031131303111210-2132110220313001"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2113032000230201-1103013033012232-2122003222033123-1231013101111302-1210021033022130-0333320210120030-0220300313321333-2331102230102123"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add`

<a id="canonical-3313012113311301-3213110112302230-3010210021003001-2221030330301032-2102223031312100-1013101222022111-2321001031223111-3133330130001210"></a>

#### `http_proxy.more_option.response_headers_to_add.append` property

Type: `"bool"`. Computed.

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

<a id="canonical-0123322303020310-3320022311222001-0301022222110323-3030223320110021-0300321323303231-1220331100010012-3013113023031210-2232202200010201"></a>

<a id="canonical-2011022310213012-2003012012002133-1330232131322120-1212103232123202-0221021020223321-0303322010022002-2221310232322031-2323121333320203"></a>

#### `http_proxy.more_option.response_headers_to_add.name` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](data-sources--proxy--reference--group-004.md#canonical-0023313330203210-1020302000012220-1300212120112013-3310321023331120-0120032003002310-2233233121032232-1123021133232123-3210303003030011): complete subsection reference.

<a id="canonical-0303232221131002-0213001331122131-0021203013102120-0120320123102102-0003111033132013-1020023101323313-0221200300313330-1310122120101130"></a>

<a id="canonical-2310312331011320-1103032102200102-0012023030113032-2023213102030132-1021001103032203-0021222210110301-0112212132212012-1113121322313211"></a>

#### `http_proxy.more_option.response_headers_to_add.value` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0023313330203210-1020302000012220-1300212120112013-3310321023331120-0120032003002310-2233233121032232-1123021133232123-3210303003030011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-3220133020330022-1232303102011310-1113033131110013-3100220200031212-3021203311201313-2122021321110320-0110232121302032-2311132322032113)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-3001022123232121-1120011210200220-3202200010120130-2323123310200100-1121112020111222-1113031101111323-1113113003020333-0301323320033313"></a>

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

<a id="canonical-3210320310133030-3333032203130022-3222203212233330-1320121113112230-2223123113110200-2023313103112000-0022003320322022-2113030322332000"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-004.md#canonical-3302221112102031-2303021103332021-3332122033010312-1012220233002200-3003213120110120-0323023331333232-0202320033030011-1330301023030112): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-004.md#canonical-1020031133122212-2031321223302010-2031022221111201-0321212100201121-0301101000301301-3113112212222312-0233311110200221-1100021200032033): complete subsection reference.

<a id="canonical-3302221112102031-2303021103332021-3332122033010312-1012220233002200-3003213120110120-0323023331333232-0202320033030011-1330301023030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-3220133020330022-1232303102011310-1113033131110013-3100220200031212-3021203311201313-2122021321110320-0110232121302032-2311132322032113)
- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0023313330203210-1020302000012220-1300212120112013-3310321023331120-0120032003002310-2233233121032232-1123021133232123-3210303003030011)
- http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3100211230013112-2013021001132102-3111023312320020-3301231130112110-0321010130310212-3213222101220330-3021200332103211-3301023011023211"></a>

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

<a id="canonical-1030131321313010-3033122331220031-0232200202230031-1210002333323023-0122113312113330-3222221232301212-3320132231103011-3000231312100133"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2312233123033323-3110311302111022-0112101010001302-3211202123311121-1121110332111011-0232020112011131-2321102331031210-1103313233302102"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1012312321132231-2011330231303332-0302311023202302-2003103101131310-0211123003223113-3031001031102033-3333101021122012-0132323023223230"></a>

<a id="canonical-3313003301002333-2002211022312210-2323312112023311-1001223212023332-3210130123323210-2332213033320333-3102111020203320-3300201102103130"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-1122031120132133-2032103321300312-0000022001221321-0230212332312020-3021322010011012-2311130302321231-3330113233011133-2003013221131123"></a>

<a id="canonical-1220130100001133-3322102101322003-0021212020032321-0233002212121300-1103102032013322-3121102030302331-2202121030101302-1002000013220022"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-1020031133122212-2031321223302010-2031022221111201-0321212100201121-0301101000301301-3113112212222312-0233311110200221-1100021200032033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- [http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-3220133020330022-1232303102011310-1113033131110013-3100220200031212-3021203311201313-2122021321110320-0110232121302032-2311132322032113)
- [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-004.md#canonical-0023313330203210-1020302000012220-1300212120112013-3310321023331120-0120032003002310-2233233121032232-1123021133232123-3210303003030011)
- http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0103002101130030-1201031220223010-0123311110333220-3220331011023301-2132123320310303-1012222322223112-0230210210102112-0123121210100211"></a>

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

<a id="canonical-1321013330032100-0303133101300102-3121033011332021-2011033322201322-1203213020012110-1100022232022210-2123022133213300-0020322220233032"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0000111202001022-2113022311110011-1322333012230023-2100231011330231-2232011131100112-2113330202232212-3111203302121030-1013211202320030"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0220200130102331-2332133020310220-2101021320201012-1311202132313001-3201110130032221-1231321132130231-3033100233131003-0233331130232231"></a>

<a id="canonical-0122120201230130-2132121033322313-0003122032103221-1211311000211001-1122120113112303-3323210200321302-2301230232203110-3311132112032111"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-0311103232311212-0122021311132121-3312130333133131-0323020121313100-3121300110200320-2101201223212331-0111013021211303-2001132132123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_forward_proxy_policy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- no_forward_proxy_policy

<a id="canonical-2310011101002013-3322323123332321-0121203101211110-3230111130002230-1130201221322013-0333330001310331-3020301030030112-3222302010330102"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-1030230220130100-0131133033331303-0222210130321320-0031101211030302-3010333010202023-1003003201301120-2030130231203121-0133333310330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_interception` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- no_interception

<a id="canonical-1013023110032033-3302100211223003-1013000321132320-3030301322333203-1220320210021310-3300313321011210-2333330211031221-1231213103320203"></a>

Type: `["object", {}]`. Computed.

\[OneOf: no\_interception, tls\_intercept; Default: no\_interception\] Configuration parameter for
no interception.

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

- [no_interception](data-sources--proxy--reference--group-004.md#canonical-1013023110032033-3302100211223003-1013000321132320-3030301322333203-1220320210021310-3300313321011210-2333330211031221-1231213103320203)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-2121232312132030-1021002130200310-0121113113303233-3010003311301312-3031223213223011-2223320220020023-0033302022021110-1101201020202012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
