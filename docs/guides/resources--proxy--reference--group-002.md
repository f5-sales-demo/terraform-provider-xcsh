---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-3010130202100113-3221230233033232-1222021013213332-2231021323111113-1231312230320332-1203220302203113-2221230300331132-0031321323103310"></a>

## Next pages — buffer_policy / 320321103313 / 6

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2200031212103131-3310131223321321-1013221221132310-2010222013231121-1313201321220130-0202321300102200-0211121130203102-2332023130201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303020121132001-2002333010210200-0223200230221131-0033033303222001-2231322303300133-2200032111133330-0223130302113132-2222021312212321"></a>

## dynamic_proxy.http_proxy.more_option.compression_params — compression_params / 102031113332 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.compression_params

<a id="canonical-0201033213231220-3001033300230203-3311012231233221-1001310312300023-3201210022303231-1331122112113032-2310333221020101-2313320231022123"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202223012201112-1322232312200111-3230221232321011-0331312201213111-2113220330220031-0010020300210002-1211320111003321-3003212323213122"></a>

## Direct properties — compression_params / 102031113332 / 3

<a id="canonical-0122313201213300-1103023220101303-0230023313000303-2232130120122300-2230302030122132-1210022313113323-0011210130110323-3231003203212022"></a>

<a id="canonical-1122210210100310-3230032121123230-3332300310223201-0133111023023331-0203322013110100-2120123232032111-1102002303033302-1030032303103213"></a>

## content_length property — compression_params / 102031113332 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
}
```

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

<a id="canonical-3111221020223312-1031303020101233-2313233320203312-3311303100331003-3100000212011003-1101300321200312-2312223000023103-0212232031022232"></a>

<a id="canonical-2122133122011202-3020133002301222-2310210033130300-3230333303300233-1211203100300113-2021211103210020-1022320122323103-1233311222303021"></a>

## content_type property — compression_params / 102031113332 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

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

<a id="canonical-3022032330103333-3222310003200332-1333130302122221-3330303013023103-1212232101120022-1232110212000313-2013003120233103-2310213311300032"></a>

<a id="canonical-3030303131310031-1130223212301333-0000002100313120-1002223310300321-3102220012002301-0211222322110112-2122203322210312-2013102321133202"></a>

## disable_on_etag_header property — compression_params / 102031113332 / 6

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

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

<a id="canonical-1300312220111301-0001012121333103-3333032212123002-1221031330212002-3232312011233033-2323222232102232-3221213011333023-2303231131030001"></a>

<a id="canonical-2121130100110310-0121233223321313-0221313101313310-3213323113122302-1021211311132300-3322033331311123-2232310332300333-0220113200020320"></a>

## remove_accept_encoding_header property — compression_params / 102031113332 / 7

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

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

<a id="canonical-2110001130120311-2301210232201232-1322132333223133-1300331020102121-0133332300220131-1120311311333222-2231311102330002-1023212111300333"></a>

## Next pages — compression_params / 102031113332 / 8

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3010323103131011-2202132221302001-3230101000320102-1213223032032320-1331201133331230-0321310100011302-0131321202201220-0211331012332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333223203312130-3311202123001201-0011023121313012-3223330001020022-1201030300130030-2000030133231322-2321100022123132-3230022331000330"></a>

## dynamic_proxy.http_proxy.more_option.disable_path_normalize — disable_path_normalize / 300031032203 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="canonical-1101100113013000-3120132013032211-0323202301031013-3223022122103200-1230221210303211-1021321311132113-2132132123011200-1221003200011230"></a>

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
disable_path_normalize = {}
```

<a id="canonical-2023103302113220-1330013021213323-3312321310311332-3231021022210023-3012000303311102-1131222310010231-1222303321013311-3121200130232102"></a>

## Direct properties — disable_path_normalize / 300031032203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133112332121320-2123302101300312-1331132331323000-1203013233111302-1013033031013230-2023231123023310-3001300021233233-3221022211032233"></a>

## Next pages — disable_path_normalize / 300031032203 / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2213331021323322-3030210022220213-2001313310032320-2310212103302212-2123120201001231-2223111022333000-1233101102313112-0132110222331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230302320012230-1120310012202220-2030331321223021-0031233203013330-0221121020101102-1320021203102222-1023033121113323-0231310232210233"></a>

## dynamic_proxy.http_proxy.more_option.enable_path_normalize — enable_path_normalize / 233223210133 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.enable_path_normalize

<a id="canonical-0202131213310023-2122123023333221-2331200130200303-0101010202222232-0121301303132332-0313332103321311-2110330023311122-0112120101332101"></a>

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
enable_path_normalize = {}
```

<a id="canonical-1222020220233210-1200032133233002-1320101023100211-1013010132033232-1221222323311003-3121013323303221-2000200101103112-2112000120211112"></a>

## Direct properties — enable_path_normalize / 233223210133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021031300130021-1221030133021011-3323213123121023-0302101030013312-0332232200221030-3213303202133210-1033332013333211-1033013131323001"></a>

## Next pages — enable_path_normalize / 233223210133 / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1300132130201131-0303023011210331-2213112112010120-1033210011003131-0103131233101201-1213132101102322-3112320303333013-0033212231102002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200121211021012-2332110331123211-1222233223231323-3302131212022230-0231312320021220-2122032101232200-3313302012033013-2231322102000130"></a>

## dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection — no_request_limit_per_connection / 131100003231 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-1231131032321020-1202100030212321-0233122210001010-3330121031203033-1013111022310321-1301113010030220-3021023033330321-1013011320020103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-3102122113201100-0321131130313030-1331032033120132-0031123133231010-0101101321302311-2102021313112113-0012001102123033-1132000100113333"></a>

## Direct properties — no_request_limit_per_connection / 131100003231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233212312031203-2301000031131100-3000202200303103-0320210200330010-2203222221130031-0322310001310111-0232032010323003-2223133102003321"></a>

## Next pages — no_request_limit_per_connection / 131100003231 / 4

- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010213231030122-0323333201032130-1030213303023320-2132021311221111-1321023020021120-2102211111022122-3210231102100003-1321121212310223"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add — request_cookies_to_add / 320003332002 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add

<a id="canonical-2022222312021313-2013221003321222-2013222203101023-2113020310032003-2030110132113101-3323330230011010-0121023003330101-3221103230332011"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-0102223022321233-1331211110212012-1111030110331102-0011131311001111-2210230223100020-2101231331303111-0021121121000002-0132000131313322"></a>

## Direct properties — request_cookies_to_add / 320003332002 / 3

<a id="canonical-0210220333113030-2111230002101012-1211131102330133-1303011123120312-1331121033302021-2100122032212322-3302213131223113-2321301003213010"></a>

<a id="canonical-1003232112002110-3103002331112132-3221331022311322-0331220332113220-3333032103100220-2211131321301013-3203233322023212-0100201101311333"></a>

## name property — request_cookies_to_add / 320003332002 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2301203233021130-1222003031120030-3120332323101333-2002001112222011-3310300111321110-1112122302303203-3331020303221210-3301313020331111"></a>

<a id="canonical-2033130321030103-3102230323331133-1312031011021010-3122121323103210-2011000233322201-1222033033333113-0123033031201331-2103022002010232"></a>

## overwrite property — request_cookies_to_add / 320003332002 / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301): complete subsection reference.

<a id="canonical-3232110311003230-0012032100322101-2211203202321130-2111330000232312-2223120323113212-0331301122021120-2112213122030232-0313101330003120"></a>

<a id="canonical-1100211012232330-0013201211311330-0033123113023012-0223123100002211-3001233101212331-0012100021100010-3212301032110202-3200121303011210"></a>

## value property — request_cookies_to_add / 320003332002 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-2212010332100202-3303320123333131-2300023300230302-1123130202021313-0112231031123332-1221122020100120-1001101313231110-3101003012031130"></a>

## Next pages — request_cookies_to_add / 320003332002 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311130202331003-0323233311000321-1233300021221021-0210230133322322-2230230320222231-2133130113333211-2322322202021033-2321310002201301"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value — secret_value / 100213102100 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0001301310201112-3222103103101101-2020122030230212-3313120311010010-3331210011130011-0132302031130302-2112301113003330-3303011221023112"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222233121101102-0022203222210121-3221313100203331-0021001223123121-3213032321131101-0302101232100121-0110323130230022-2312131032232103"></a>

## Direct properties — secret_value / 100213102100 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2003123130303002-0033123013233203-1333010122322230-0020001200312011-0120032321011102-0232230103330111-3313230132322113-3303312020021233): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-2233302133310312-1003211200232032-1301322322202333-1203333211210312-1130021211002320-1323232223300313-1223312230022030-3312332213130331): complete subsection reference.

<a id="canonical-3032212223202331-2312333321223033-2313332303212211-2013020301023133-2122023203011022-3231333212331012-0212103030012202-0223210332000310"></a>

## Next pages — secret_value / 100213102100 / 4

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2003123130303002-0033123013233203-1333010122322230-0020001200312011-0120032321011102-0232230103330111-3313230132322113-3303312020021233)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-2233302133310312-1003211200232032-1301322322202333-1203333211210312-1130021211002320-1323232223300313-1223312230022030-3312332213130331)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2003123130303002-0033123013233203-1333010122322230-0020001200312011-0120032321011102-0232230103330111-3313230132322113-3303312020021233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121320333223203-1112101213303230-3033223320002311-2000001313310111-0330232211332301-3010322010230230-3210320133000102-0102011302021332"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 113200102320 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3223110031103030-0322023113110213-3212113131221203-1221233231310202-2203122131232030-1323311101110031-0111133101330202-1320120031332222"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013220111003220-2112002213301031-1120223012111323-0101121320031000-1223110233000331-3311202231201121-1023210332323031-1320021233013313"></a>

## Direct properties — blindfold_secret_info / 113200102320 / 3

<a id="canonical-1310103132001013-0331220123000320-0232300020230321-3213032021103012-0311233033323311-1230223230321222-3012130333300232-0311233121231330"></a>

<a id="canonical-0210131320200000-0302121331121023-3011322101031100-3130032312000302-0123213232020121-3211330033132221-2133123020231032-0311021111001130"></a>

## decryption_provider property — blindfold_secret_info / 113200102320 / 4

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

<a id="canonical-3111021230131311-0101211132100312-1031031200021302-0233123001130010-3212000121122010-3211312231132133-2222233100312000-0020131122311011"></a>

<a id="canonical-1132231202121102-1213230122300233-3323012203302212-0120012321001030-3001302220133132-1023122032031211-0321100031003102-1010100321113311"></a>

## location property — blindfold_secret_info / 113200102320 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1132322330310030-1332020103201012-2233302120312300-2203032130001102-2133012312020332-0003310130201222-0230321312132221-2001101013113210"></a>

<a id="canonical-1120321311333331-3031232323312231-2011101100321113-1302301133111100-0002113310032320-0102231312222212-1101303030002002-2233230212333232"></a>

## store_provider property — blindfold_secret_info / 113200102320 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1203132110213223-0012012301112221-0200210322111123-2212101022011013-1112203212310300-3031311201320030-1321032021032110-1320231203001012"></a>

## Next pages — blindfold_secret_info / 113200102320 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2233302133310312-1003211200232032-1301322322202333-1203333211210312-1130021211002320-1323232223300313-1223312230022030-3312332213130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022000213011021-3122323100313001-3323033232022311-1002313102113111-0100121031010011-2331301212013203-2221023222112303-0033231013102220"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 003131110022 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2221300200230203-1103230202232212-1120100321231101-3103120231310311-3002322311120212-3230100002333023-0003131213111120-1131313112321032"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200321001203220-2000300230222202-2213020020331102-2333221323232313-3202133101011031-2212332211011330-1311113213031212-0311201123303101"></a>

## Direct properties — clear_secret_info / 003131110022 / 3

<a id="canonical-1213310200312330-3303111332132013-0220312223110131-3023102111320121-3212212202313033-3323330101121120-0221012313032322-2333023133211132"></a>

<a id="canonical-3303000020031013-1032013221232030-0111110331202021-0203303121222300-3010323111122301-0323130203031211-1323323313131102-1033103001133003"></a>

## provider_ref property — clear_secret_info / 003131110022 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2330132102030303-2122021321132333-0003131032302103-2231033323102211-2021301013300031-2133133120222330-1130300332313022-0333332010332121"></a>

<a id="canonical-0211023320232101-2131330100011321-1210333013301212-1120011300001302-0132123111211021-2122102202021302-2321130310303322-2013130313122122"></a>

## URL property — clear_secret_info / 003131110022 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-2133020121333113-2030200012203312-3001212301330202-0032110321231320-3213130213122102-1122312031131322-1210212223322031-3121030133302332"></a>

## Next pages — clear_secret_info / 003131110022 / 6

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303100311333213-3010122133221302-0103010320222030-3333221032021231-0212001100101131-2301000112130213-2220010032120011-0222232021211031"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add — request_headers_to_add / 301110011100 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add

<a id="canonical-2202131210002333-2223311202330033-3012001110202212-2302210331120322-0310302231110031-0222323111211312-1221113131000113-1001120301203112"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-2103010013031013-2202203122120012-1200332003213222-3100333031010222-2132200132011103-0230112200203031-3323100132321010-2302120212220310"></a>

## Direct properties — request_headers_to_add / 301110011100 / 3

<a id="canonical-2210013321203123-1223301312033202-1322303202011203-3101030303300022-2132232022120202-3331220333313313-0311013003021131-0132023130212330"></a>

<a id="canonical-2022212030301123-0303212221202300-1111202220302332-2323230230030333-1220000110033201-1020033332213303-2212302111213012-3300113120220230"></a>

## append property — request_headers_to_add / 301110011100 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303322301202131-3322330220302012-1330123000223011-2313230210003230-1122303022202211-3222010003121111-2303032132032333-2202032122322020"></a>

<a id="canonical-0213310011132123-0033321012301323-1011110213021203-0330331312310100-0013132031023112-0033313303302210-3210100113321223-3121023213130122"></a>

## name property — request_headers_to_add / 301110011100 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031): complete subsection reference.

<a id="canonical-3322230233023202-1113310222313103-0131222002322030-3231112222300332-3111103111310233-2102031301123003-0211112013032212-1201223323001133"></a>

<a id="canonical-1321232013303202-2130233302032122-3330332020301030-0210113222233320-0023132223312000-1033212130133203-1123000300013233-3012010323302100"></a>

## value property — request_headers_to_add / 301110011100 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-0320120131002213-2003312331123103-3321302101000033-3023131221000031-1133003123313203-1103303323003120-3013230133000212-1120021322000201"></a>

## Next pages — request_headers_to_add / 301110011100 / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222203223011220-3023200113121333-2110311310000012-3323322201132232-3133131220120200-3132230113322331-3001200102321231-3122021311132133"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value — secret_value / 003322222302 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-0102211321132033-3022303033333203-3220230332101230-1322303331211320-0210112321202113-0202113030232010-3020030230333320-3023122321131022"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122201222220111-1031133312310023-0212313310303022-3133120101220002-1022223200210110-3310231300230312-3231322011120110-0302111013223111"></a>

## Direct properties — secret_value / 003322222302 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2211130000033232-2003133020123311-3203103301100322-2031223301033121-1133223321112330-1100012022221111-1220030321300323-3232201323220102): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-0023100032221102-3000313313001020-0122312201033220-2231022212120022-1011130320311233-3320122110322133-2302122220211112-0003010130203001): complete subsection reference.

<a id="canonical-0223231103310200-1011322222233003-0101133323032203-3312102013112330-0022232302213133-2201132121121211-3013201000112120-1213212103332331"></a>

## Next pages — secret_value / 003322222302 / 4

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2211130000033232-2003133020123311-3203103301100322-2031223301033121-1133223321112330-1100012022221111-1220030321300323-3232201323220102)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-0023100032221102-3000313313001020-0122312201033220-2231022212120022-1011130320311233-3320122110322133-2302122220211112-0003010130203001)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2211130000033232-2003133020123311-3203103301100322-2031223301033121-1133223321112330-1100012022221111-1220030321300323-3232201323220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213113320003003-0010321021300133-3210130022111022-0330222113113313-1121301111322122-1131123013200300-1330202012022303-0021233222322300"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 113131313323 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1131302230310321-3001010113301103-0310322123021312-0132023220201031-0003302130123001-0312132210133300-2122322333300002-3230300101213121"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332033013133313-1133133121022220-3312132130310302-2223101002230222-2003010221212110-2100103113222130-1303133101120321-1030122321000210"></a>

## Direct properties — blindfold_secret_info / 113131313323 / 3

<a id="canonical-3313331002100310-3203222233132122-3022322110130231-1230222333223010-1031233332333101-0101221211000132-2003220122310110-2331122102322032"></a>

<a id="canonical-3012002100312022-2002200023001131-1012000021311113-1223022113012032-2002201113012321-3300112301212230-3312032100011010-3030323030233312"></a>

## decryption_provider property — blindfold_secret_info / 113131313323 / 4

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

<a id="canonical-1221001132103121-3012301331013310-1301111213032120-3333203212023022-3130102021220230-2130110101333013-0003233322012221-3023102221100110"></a>

<a id="canonical-3331102222332333-2002102020000311-1003303321203203-1301132212312132-3302211331112122-1220102233121103-2030011213022023-0012121230303301"></a>

## location property — blindfold_secret_info / 113131313323 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1030331203321001-1331233211101013-1101030123021312-1232322231311332-3032323020012102-3110301220001110-1310120203033232-0221332030322110"></a>

<a id="canonical-2121311311011030-1010313113130222-1013232220323003-2301233211200321-1002232203120103-0323231112323131-0031021002011112-1133103302311211"></a>

## store_provider property — blindfold_secret_info / 113131313323 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-3013333112013223-1311113003301111-1212130122200103-1220001321323213-2123103220301332-2222222020012333-2011130332203310-1010212311312321"></a>

## Next pages — blindfold_secret_info / 113131313323 / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0023100032221102-3000313313001020-0122312201033220-2231022212120022-1011130320311233-3320122110322133-2302122220211112-0003010130203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320101122213132-2300303310122323-3023101033231012-3203111002213202-2130210032103223-1211022231300031-1210211333002120-1000120330023001"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 100223001332 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3232232000112022-3203000010223311-2102021021002032-0003030131112211-1231223110332232-1033023113030313-1323210013331301-3123031311100031"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132111233230331-0021003100311130-3301231223012113-1110303003121122-2233331220322110-1330323032131022-3122213320212212-0101220312133311"></a>

## Direct properties — clear_secret_info / 100223001332 / 3

<a id="canonical-1020322220112012-1331103312102013-2112030213211033-0121011330123020-0221323101212101-0302120110123311-1012300221310331-1222002130133310"></a>

<a id="canonical-2310033322132203-3201233301323131-2023222120202213-3220030200311021-0220331131310110-0230111321200032-0203130210211331-0232231111303320"></a>

## provider_ref property — clear_secret_info / 100223001332 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3203022131030022-1023231022332103-2121110300010113-0320030311100123-0230110331113313-1233311230022033-2020031021001221-3311331130103320"></a>

<a id="canonical-2133003112323210-0120311112311011-0222112102113320-3232331221002313-3212330313211030-0131301012101103-3012301030210110-2113030132103223"></a>

## URL property — clear_secret_info / 100223001332 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-2111122013102200-3021202022031322-1230331201102213-0203030100323320-0022233332011303-3011012031232101-1113011102221311-0210030121030313"></a>

## Next pages — clear_secret_info / 100223001332 / 6

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121000110023231-0032220111311233-0103101132321000-0210120211131111-3111330300213322-0010012220113033-1301211033131122-2201330330132312"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add — response_cookies_to_add / 213200003102 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add

<a id="canonical-3021311133323301-3301313030222230-2103232331230230-2020232123201303-2201230232101011-2232333220112133-2101320021201103-3201231121111213"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-1300010233222230-3033032020312130-1003311310231222-0202123232123032-2202211230333102-3300222120222303-3121132301133303-2031020200321122"></a>

## Direct properties — response_cookies_to_add / 213200003102 / 3

<a id="canonical-2100313122110033-1010230312313003-0012110330301001-1001211032320223-3001103223230303-1100310111110023-2131103230213112-1023321333303223"></a>

<a id="canonical-1121102231130012-3110233223313303-1232300222001030-2000103232321210-0122133301310333-3211101003033020-0203310102320122-1312130122210013"></a>

## add_domain property — response_cookies_to_add / 213200003102 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0002121332213331-2321002022022100-0201121011130323-2010113200233231-3220312222223013-3121121033102331-3133120301322010-0201300032000020"></a>

<a id="canonical-1311132110202311-3010120230010312-2130110211332201-0201020312212032-0302112230313300-0301100022032110-0331202232011022-3221021113332330"></a>

## add_expiry property — response_cookies_to_add / 213200003102 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [add_httponly](resources--proxy--reference--group-002.md#canonical-3033213011100232-3233233212212022-1300303300033113-2330000333112330-1002000230212220-1131302000110101-3103222123201021-2003012322032232): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-002.md#canonical-2331122031133303-1013033330113203-3212131002231210-0321211300331221-3101323302000213-1302303001223023-0131100333021111-1232300230300233): complete subsection reference.

<a id="canonical-0222022313000211-3101310021233113-1133130002223022-1230133100303323-3020202302232110-1321231123231011-1322210021222123-0111220333020311"></a>

<a id="canonical-2333212122322323-3232110233322322-1210130022031013-3022211212100110-0223303110301001-1303020022133310-0200101220102313-2002123102322301"></a>

## add_path property — response_cookies_to_add / 213200003102 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [add_secure](resources--proxy--reference--group-002.md#canonical-3022132322222221-2113323112133200-1102231302131321-1233003303030320-2010310111120033-2031113023100231-1220013033133102-2220132120032000): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-002.md#canonical-2232322033322210-2123311102020333-1212111303033133-3202301213220233-1223012313320321-3111221333210203-1013333130213223-1123122212113030): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-002.md#canonical-0201000303101102-0323111232010312-0123131231100212-0101212232200003-0020103212102232-3321311322310313-3023110302223123-2232303021013310): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-002.md#canonical-0323112031212001-3031321031312312-3311323122100110-0113103323111012-3132001022332103-1133202320233230-1233020020000100-1332012132030130): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-002.md#canonical-0232130310110203-3213130300000330-0031101022100221-1111100010120002-2333232001103113-2313031033202112-0002130210011002-0131101301200320): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-002.md#canonical-2220010022130303-2223300332110113-3223232303213321-1123000032003130-1123010131112100-3032323221011212-3222222113122000-1211212032223132): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-002.md#canonical-0101303332312031-3212110122321000-1031230003032201-2220032323110012-0102312203231032-3123230321211000-2030020110130233-2011332003031133): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-002.md#canonical-0020202232330223-3120103211031311-1133322110201000-0203122103013021-1023032320013130-0313021012010030-1320323110120112-2011121012101233): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-002.md#canonical-2322330201013310-2122032230131032-1122220231333301-0033033113231232-3233123212011000-0101121312100323-0220201223202113-0023233221100102): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-002.md#canonical-0322020233022323-3301022331322223-1100330022133330-1010023101322120-0032330133302213-3312331313233201-2000332130131302-1001023102213301): complete subsection reference.

<a id="canonical-3220330222220321-3010233330132310-1231312230122011-0003133000101112-3002033203310132-2101233223333121-1110010031102100-1322202230321033"></a>

<a id="canonical-0221333222033331-3022100211222330-3030332111033003-2330002322310201-1223023212011012-2330010230323102-2001002320202320-0031010033310031"></a>

## max_age_value property — response_cookies_to_add / 213200003102 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

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

<a id="canonical-1123002130011300-3200103312031033-3023113101103230-2230323311312200-1212233331110002-0131223311212332-1023121112113021-3123221002112132"></a>

<a id="canonical-2131111200012300-2001222220121213-3022321122002210-2302230302122023-2303010301223001-1123222223231201-0112300323202233-3230333310221001"></a>

## name property — response_cookies_to_add / 213200003102 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0003320203011302-2012103030001301-1232323300233133-3312221233023313-1323111310231011-1322313110120300-0223223231333220-1003112000130010"></a>

<a id="canonical-2013113232023222-2311132302020311-3221203033230113-1233023010020120-0301022013222123-1211332300000031-0030110020110130-2300330303022211"></a>

## overwrite property — response_cookies_to_add / 213200003102 / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--proxy--reference--group-002.md#canonical-0200102201011213-2222230221032313-2311211112321022-3111131322023132-1303311311100312-0013320103303100-1033321320113313-1210233221123032): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-002.md#canonical-3202033321232232-1231123230032120-0011320312320010-0300221231211121-3200303303303331-0021020201102203-3111231121002330-1323013020310132): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-002.md#canonical-2001200323200002-0100021031222112-1113010212212110-1211123102320312-2032011110231213-1333033233201023-1330300201023123-0112001320101131): complete subsection reference.

- [secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333): complete subsection reference.

<a id="canonical-3033331112111032-0030310131023121-2333232311231221-0200000323033201-2310113121310201-0320013201302332-2101113012222202-1313103231001231"></a>

<a id="canonical-3122213113101313-1131330321120220-0310001200310112-0010033213320030-1031311202112320-1212112010013122-2311303010300221-2010032213221112"></a>

## value property — response_cookies_to_add / 213200003102 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-3323332213003100-1323311012322002-2002032322312002-0110230131213100-3202331032000232-0212302301310202-3303332032202320-3303102032011323"></a>

## Next pages — response_cookies_to_add / 213200003102 / 11

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-002.md#canonical-3033213011100232-3233233212212022-1300303300033113-2330000333112330-1002000230212220-1131302000110101-3103222123201021-2003012322032232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-002.md#canonical-2331122031133303-1013033330113203-3212131002231210-0321211300331221-3101323302000213-1302303001223023-0131100333021111-1232300230300233)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-002.md#canonical-3022132322222221-2113323112133200-1102231302131321-1233003303030320-2010310111120033-2031113023100231-1220013033133102-2220132120032000)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-002.md#canonical-2232322033322210-2123311102020333-1212111303033133-3202301213220233-1223012313320321-3111221333210203-1013333130213223-1123122212113030)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-002.md#canonical-0201000303101102-0323111232010312-0123131231100212-0101212232200003-0020103212102232-3321311322310313-3023110302223123-2232303021013310)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-002.md#canonical-0323112031212001-3031321031312312-3311323122100110-0113103323111012-3132001022332103-1133202320233230-1233020020000100-1332012132030130)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-002.md#canonical-0232130310110203-3213130300000330-0031101022100221-1111100010120002-2333232001103113-2313031033202112-0002130210011002-0131101301200320)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-002.md#canonical-2220010022130303-2223300332110113-3223232303213321-1123000032003130-1123010131112100-3032323221011212-3222222113122000-1211212032223132)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-002.md#canonical-0101303332312031-3212110122321000-1031230003032201-2220032323110012-0102312203231032-3123230321211000-2030020110130233-2011332003031133)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-002.md#canonical-0020202232330223-3120103211031311-1133322110201000-0203122103013021-1023032320013130-0313021012010030-1320323110120112-2011121012101233)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-002.md#canonical-2322330201013310-2122032230131032-1122220231333301-0033033113231232-3233123212011000-0101121312100323-0220201223202113-0023233221100102)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-002.md#canonical-0322020233022323-3301022331322223-1100330022133330-1010023101322120-0032330133302213-3312331313233201-2000332130131302-1001023102213301)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-002.md#canonical-0200102201011213-2222230221032313-2311211112321022-3111131322023132-1303311311100312-0013320103303100-1033321320113313-1210233221123032)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-002.md#canonical-3202033321232232-1231123230032120-0011320312320010-0300221231211121-3200303303303331-0021020201102203-3111231121002330-1323013020310132)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-002.md#canonical-2001200323200002-0100021031222112-1113010212212110-1211123102320312-2032011110231213-1333033233201023-1330300201023123-0112001320101131)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3033213011100232-3233233212212022-1300303300033113-2330000333112330-1002000230212220-1131302000110101-3103222123201021-2003012322032232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013101020220210-0222023313010021-0333032103211010-3001233212110112-0021232323212101-2302231231112221-2323130021232332-1120113223203123"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly — add_httponly / 101110203103 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-0132213222030030-2230100100112012-3203011102032002-2232230021210201-2001312131132112-2301301210312231-3023030011232221-0230100002102131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-0201312023300100-2331333000111031-0022122022031123-0221332302202022-0122120133301202-3121302333132132-1033212213131000-3033021010210301"></a>

## Direct properties — add_httponly / 101110203103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331020321031021-2213022002123000-0111001210130311-0201111220032232-0313021010120100-3302312101230000-3301032203003101-3120131030000031"></a>

## Next pages — add_httponly / 101110203103 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2331122031133303-1013033330113203-3212131002231210-0321211300331221-3101323302000213-1302303001223023-0131100333021111-1232300230300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232032231122220-2011130221302330-1213223212200201-0310331311031200-1312212233131313-1332111101211031-0133300201032200-0310311133022303"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned — add_partitioned / 301132030322 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-1011133030203203-1212123002112132-2221221121302222-1300211221112312-0032301202320100-0220300310020031-1301310010032030-3130102310311311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

<a id="canonical-2302020223002211-1022030303110111-1220203320322111-1110131001323222-3102333032011310-2333023311020103-0211013231010212-2301103033222202"></a>

## Direct properties — add_partitioned / 301132030322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022311310230112-1303210120221032-3012133233032201-0012132023310202-3100123300110020-2323120013311333-2112223220320300-2200323020130212"></a>

## Next pages — add_partitioned / 301132030322 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3022132322222221-2113323112133200-1102231302131321-1233003303030320-2010310111120033-2031113023100231-1220013033133102-2220132120032000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013130203212213-1131030210121320-3332302231131221-0120123031321213-1233201020300321-2131012011201122-2313013010333213-2230203012102120"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure — add_secure / 130323011310 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-2200113333202300-3201201220321103-3312032233210203-3133012003130021-0203013003322000-2331002030201312-2130210210200332-2100333220003102"></a>

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
add_secure = {}
```

<a id="canonical-0130131013300021-3010030003200011-2232202111010133-0101323331300000-1210123103023011-3303111322013130-0230323211311330-3122010300031100"></a>

## Direct properties — add_secure / 130323011310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000001203312110-3012002031203112-1100020221030031-2121121303301321-2210223112323122-2231200120103122-2133021020131113-2321122220013033"></a>

## Next pages — add_secure / 130323011310 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2232322033322210-2123311102020333-1212111303033133-3202301213220233-1223012313320321-3111221333210203-1013333130213223-1123122212113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110102313032332-3012113203123103-2212323311203130-0133022003303122-1331010322122212-2010330222213213-0210121332331122-3303222200221322"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain — ignore_domain / 111131201320 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1020321221322313-0032323310213223-0102000222211333-0121100202121322-3011122202112313-1120201023323030-0100211101011030-3113102312033012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

<a id="canonical-3223111123022030-2200201312312011-2012130310301231-2233033323132001-2012112011013202-1100100003323033-0312300022330323-0023121201023313"></a>

## Direct properties — ignore_domain / 111131201320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211122320122003-0331122123132220-2020130001311312-3232033021111100-2211321232231313-1313101003322330-2111330130021002-1232012022010103"></a>

## Next pages — ignore_domain / 111131201320 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0201000303101102-0323111232010312-0123131231100212-0101212232200003-0020103212102232-3321311322310313-3023110302223123-2232303021013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001001233022232-3123110211200013-2002022313232111-0021031112013120-0303321332000013-2133322123332132-0223001331222133-1030130323201322"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 012303030213 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-2321231121323100-1032111112203313-1223232230323303-0223030120210330-2010030100221332-1112012333022203-1032011031002111-1310010222002020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

<a id="canonical-0203011330302033-2103213323222210-1131013013331233-3103020303233013-2223032332211202-2232120202201233-3220033013001212-1033313101012223"></a>

## Direct properties — ignore_expiry / 012303030213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302230210221300-3133133220302010-3321213010330010-1031103111300320-2002122023132323-0012223002120223-0320210103112103-1100222323213210"></a>

## Next pages — ignore_expiry / 012303030213 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0323112031212001-3031321031312312-3311323122100110-0113103323111012-3132001022332103-1133202320233230-1233020020000100-1332012132030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213013330220032-0201333330223232-2003102323230033-2122113122103222-2130220111130211-1300002110101102-0112131321312123-1110113322103131"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 110123300311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-0222131313232122-1100212113122131-3012311313322110-2100020223332330-2031311031323132-1111113001102022-0002300001123013-2103021023212002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-0323231001010100-0122323133020323-3030332123313100-1011003033221030-1232121101020121-3021322001201113-0113002100333001-2231303220321100"></a>

## Direct properties — ignore_httponly / 110123300311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203232030230133-2322212233233311-1133330020212210-0311100001110323-3123113331013113-2002131123320113-2120301210311223-2310003232230022"></a>

## Next pages — ignore_httponly / 110123300311 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0232130310110203-3213130300000330-0031101022100221-1111100010120002-2333232001103113-2313031033202112-0002130210011002-0131101301200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133203223103011-3313000300221301-1101321022032030-0102203231132310-3001111320231322-2001112201020023-1312230300201011-2101310213222101"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 100222101200 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0211120332130021-0001213320131033-3312012332112221-1220202101331222-1301330223031010-2110122322320211-3133203313222332-0130222213201333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-1101103112100210-1201223001220001-2301103300130213-3313313332331230-0112003211333120-2231021221300201-3300233013233111-0133320113300110"></a>

## Direct properties — ignore_max_age / 100222101200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213220202220011-3013031110222023-0033021322301301-0211310012211320-3322330003233133-2021021130101312-1010201020211012-2200032100221031"></a>

## Next pages — ignore_max_age / 100222101200 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2220010022130303-2223300332110113-3223232303213321-1123000032003130-1123010131112100-3032323221011212-3222222113122000-1211212032223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012000332011211-0220203302321330-3110201102232203-0332003001130301-2322320212003100-3010112012013203-0113323121312200-0212003103220101"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 020311023203 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2312321101032211-2221330131000212-2123311023211202-2323331110022111-1002033003323001-3120131030122020-2210212132101220-3330313120103010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

<a id="canonical-0210200002111300-3011302113211220-2130121233213113-2012012031003212-2132223333012232-1020031232112222-2132132012312203-2322233302121011"></a>

## Direct properties — ignore_partitioned / 020311023203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202110010020231-3320102001120210-0203133022201001-2200303220301021-3311233231303131-1032110221222322-0010210332211220-2223120003000112"></a>

## Next pages — ignore_partitioned / 020311023203 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0101303332312031-3212110122321000-1031230003032201-2220032323110012-0102312203231032-3123230321211000-2030020110130233-2011332003031133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020231320003013-0133232102213011-3223333203122300-1023121231010123-1212001321320012-2301002301201331-2031220303333233-3210313002122002"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path — ignore_path / 000023300222 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-0023322030320302-1323031111323303-1023101111223010-1223111223310313-1011203100012011-3013131312203201-2200123131333312-1131222002113101"></a>

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
ignore_path = {}
```

<a id="canonical-3203333213030021-0021111210103020-2102111020300321-1122112112300032-0322010130200003-3211303313311121-0202301201201313-2231011031012032"></a>

## Direct properties — ignore_path / 000023300222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122300102020302-2232000020022112-1001230103332013-2102103321223231-1331132330012301-2122212010102233-0123320010113212-3312303032003020"></a>

## Next pages — ignore_path / 000023300222 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0020202232330223-3120103211031311-1133322110201000-0203122103013021-1023032320013130-0313021012010030-1320323110120112-2011121012101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323030200301132-1101201310003010-3213003013302122-3021130233100332-2100220223111102-0200231211202122-0213133101122101-3303131131113211"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 130121223003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3103132203323311-2130010101333120-3303002002132002-0112220003103333-2101201021013210-0330231122022032-3020012022210321-2200103020110221"></a>

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
ignore_samesite = {}
```

<a id="canonical-3111111003021133-1223233011110302-0310103203300003-3330130012230133-2303011201220020-2032133000020311-0332100001333002-3112332101203212"></a>

## Direct properties — ignore_samesite / 130121223003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021111002132031-0333330320113311-0010012123201113-3110330122103003-3133022230100112-2113301212321320-1100232033121113-1202311213121103"></a>

## Next pages — ignore_samesite / 130121223003 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2322330201013310-2122032230131032-1122220231333301-0033033113231232-3233123212011000-0101121312100323-0220201223202113-0023233221100102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123213301331100-2312332031001331-0130113022000321-2020230231223231-2112233011310312-2330032031101132-0102201212030223-0332333301332023"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure — ignore_secure / 223333100210 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-2113332000110031-3303031200131001-1323303213212123-1313023302311200-1112203333213000-3123102112132312-3130212022022133-3202332131223131"></a>

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
ignore_secure = {}
```

<a id="canonical-2120122311330330-0221121121011222-3111313002221012-0001200223220302-1021010133021103-0323202220033113-2233302122322123-2122113012001033"></a>

## Direct properties — ignore_secure / 223333100210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301122111233002-3211011320103303-2321030201113010-1322023303222230-0021322321122013-0210322313030003-0113011123211230-3203320321030031"></a>

## Next pages — ignore_secure / 223333100210 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0322020233022323-3301022331322223-1100330022133330-1010023101322120-0032330133302213-3312331313233201-2000332130131302-1001023102213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011202222102133-2022232031100223-2203032123313321-2212121331103032-1230333032020102-0020313013031011-0321320113110231-2323321333112111"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value — ignore_value / 202233230311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-1322201103113013-3210023332323021-3212102010002121-1131131300221211-3210220110323233-3310011301313210-3003101222311123-0102033032312102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

<a id="canonical-3301303103101300-3213023002023212-2000200012230220-1002100011110310-1030310223002013-0031210030330022-3110203110212132-3212303001122230"></a>

## Direct properties — ignore_value / 202233230311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030302010303020-1323130113131100-2302131310023203-1011131303120303-2123033030111232-0210122021303222-3312003322113000-2023010311003122"></a>

## Next pages — ignore_value / 202233230311 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0200102201011213-2222230221032313-2311211112321022-3111131322023132-1303311311100312-0013320103303100-1033321320113313-1210233221123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322111333302013-0202021231211033-1133003223301003-0203301321303021-0010212333133111-1311131102030112-0022020011320222-0230000223303002"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax — samesite_lax / 303212132201 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2111303103112302-2322230001000320-3010100021120311-3203323020120103-1100130330012200-1232301011211323-3313203020022220-1222203302130131"></a>

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
samesite_lax = {}
```

<a id="canonical-2301133312302302-3300100300033110-0211133201300111-1131330000100011-3102213220203212-0300331000212202-2200233120211003-0103011023223300"></a>

## Direct properties — samesite_lax / 303212132201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302213223311033-0012313023212300-0122221023322131-3110200010103212-3223122330110300-3231100012033320-3322031323003101-3303132030112221"></a>

## Next pages — samesite_lax / 303212132201 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3202033321232232-1231123230032120-0011320312320010-0300221231211121-3200303303303331-0021020201102203-3111231121002330-1323013020310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201233202321221-2212312121310130-2212221100221320-0232302201220200-0130033323223320-0112003230003221-1031331313321101-0212131233230211"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none — samesite_none / 233133302111 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-0302122300222310-2311210320302011-2033321103023233-3202000031012322-1232203313123031-1230003330233131-1131231203302211-0333111113012120"></a>

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
samesite_none = {}
```

<a id="canonical-1013123212231100-2220320230102122-2330230211212022-2001013322220113-2300302111210312-0223300021032001-3311330212231031-0233121100003120"></a>

## Direct properties — samesite_none / 233133302111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211030323110131-2223213031232302-0133323310213030-2222313030333110-3110302323202000-1011321112123012-1333013131011101-2302211032020211"></a>

## Next pages — samesite_none / 233133302111 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2001200323200002-0100021031222112-1113010212212110-1211123102320312-2032011110231213-1333033233201023-1330300201023123-0112001320101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220222203312213-3320123020212021-0001100320321120-2201233201112003-2320200020221122-2300310201120102-3011331031212223-2312111323021101"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict — samesite_strict / 211121011033 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-0303100222301012-0033123312320031-1212000003323222-0032132112033021-0122220103113113-1303323011211310-1121220131133323-0111203130230130"></a>

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
samesite_strict = {}
```

<a id="canonical-2323332132000132-0121113101322002-3122130001033113-1331300120031313-3132233223131221-0332111331231331-1231103023322320-2301123221132122"></a>

## Direct properties — samesite_strict / 211121011033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330130102013202-2323300311030032-2323012031233223-0022023223132213-2023101132222022-3311221200000302-0132110332010231-3031021011113003"></a>

## Next pages — samesite_strict / 211121011033 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310303303010100-1130222311130121-3121021233133032-2313123012003221-2122330232130221-1011120010221232-3201112323331113-2121330203132112"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value — secret_value / 022322103232 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-2300013201200212-3313122223010310-1003333012012002-1102003132021303-0203201000211030-1113113310033120-3130202012311220-3103221330012133"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031030133231323-0132230010130133-2223203011020301-0132013101312003-3312301323021020-0200311130303311-3331320023031123-3312221021201233"></a>

## Direct properties — secret_value / 022322103232 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0022310021300010-2002213123221021-1001012203233313-2130010100230112-0312302011203101-0121333223102332-3030033321101032-2033313220230311): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-1311101122321023-1021222121333111-0323323122313323-0232233202323111-0211331223233002-0000031332320020-2330323311321320-1033323213111111): complete subsection reference.

<a id="canonical-1213233320111023-0120310311110132-0230122013211033-0030300103330021-0302220321321231-1033003003203332-0331132031002111-2201111321230130"></a>

## Next pages — secret_value / 022322103232 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0022310021300010-2002213123221021-1001012203233313-2130010100230112-0312302011203101-0121333223102332-3030033321101032-2033313220230311)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-1311101122321023-1021222121333111-0323323122313323-0232233202323111-0211331223233002-0000031332320020-2330323311321320-1033323213111111)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0022310021300010-2002213123221021-1001012203233313-2130010100230112-0312302011203101-0121333223102332-3030033321101032-2033313220230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200000301322021-1130333131120231-3122100231123210-0130000130332132-3023320033302113-3200300133230300-3132111312100112-0003100223132322"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 300322030230 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2032201022302321-3220001201101001-3121100302131232-3221303122110200-3222103302023111-2033022111030223-1333302003213001-2002213102210113"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022312023131132-2230002201023132-3113000312303320-1021322202121312-1202133003001030-1330222300003120-2003223113132223-0330232222110201"></a>

## Direct properties — blindfold_secret_info / 300322030230 / 3

<a id="canonical-3200330221230330-0303003013011101-3201233120132000-0111201003030010-1000003030202301-0012101021320321-0032220321103120-2230312033022203"></a>

<a id="canonical-1231232023033333-1111021110032232-2323100213020231-2021121222230101-2312011301332320-3133330303200012-3000103232013211-3221211223313211"></a>

## decryption_provider property — blindfold_secret_info / 300322030230 / 4

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

<a id="canonical-2103222111203121-3021223030300201-0031032312213301-0130103131021110-3302312022020301-3311131113330132-3213320121332302-0120231100020033"></a>

<a id="canonical-3023121022013020-2323322032111233-3220311122030300-3123012031102021-1121023033212113-2031022232203021-1301203012211132-0000033300221210"></a>

## location property — blindfold_secret_info / 300322030230 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1200123303110021-3122121300022111-1030222021213031-2022330131022130-3101223113220112-0022033111113222-0012131000323011-2312333122021032"></a>

<a id="canonical-3113100003301203-3000322201211311-0302130003321022-2312212111030103-1111213131001131-2001110310311001-3030300221022103-3230220222003213"></a>

## store_provider property — blindfold_secret_info / 300322030230 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-3303003231222031-1300200101103311-1101201030001210-1030031030320302-3102133032300120-0031211211332123-0202003221301130-3312002121032210"></a>

## Next pages — blindfold_secret_info / 300322030230 / 7

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1311101122321023-1021222121333111-0323323122313323-0232233202323111-0211331223233002-0000031332320020-2330323311321320-1033323213111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311333331130010-3211001222322333-1330221202132301-1310211003001303-0123102201132130-2122110003312332-3321213303011331-1131112101330221"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 013223002033 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3330123310023133-3320003103012110-2203233331231002-3030333201000232-3210220232010132-1131023322023330-2332012230231230-3223230111200022"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130003221312013-1222311210132212-1102210230112023-2333001300210100-0112030313032033-3302101123212013-1203022113001031-1313020211220022"></a>

## Direct properties — clear_secret_info / 013223002033 / 3

<a id="canonical-0123003323021323-0110002311323103-2122021320231112-2310202322131033-0011003010013003-0232000013000311-2203232110333231-1311212023230013"></a>

<a id="canonical-1100033311203003-3201011030111201-2211013013123301-0332303023013221-0221321222220322-3210302200211132-3122303033200123-3213300113002110"></a>

## provider_ref property — clear_secret_info / 013223002033 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0201123322102021-0332103110120030-0303131202223122-3010101022030332-3022321301233220-2131221001201222-2111101201101301-0330033203113222"></a>

<a id="canonical-2110201031221022-0021103311110202-0121323332303320-1113103021202211-1332211001211211-3330110232310213-0130233313113332-3012000121213312"></a>

## URL property — clear_secret_info / 013223002033 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-0113301220313211-0323203201122203-2221220011113100-3221301113221033-1222022231021323-2120002000023333-0313020200233101-1000023122122130"></a>

## Next pages — clear_secret_info / 013223002033 / 6

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000122310002303-2000223311103021-0311120301130332-3123323001010230-0101210102023031-3002330313020213-1122103111012120-1333100120210331"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add — response_headers_to_add / 010200321100 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add

<a id="canonical-0202100121330003-1013012330213311-2032200112120101-3210302011102313-3010233130320012-1110001020300000-1310203033211122-0003221103123220"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-3011000032110133-3103331113123323-3110033133321103-3201111033220302-3122230331002222-1313000320012321-3322213300020121-3112232311211300"></a>

## Direct properties — response_headers_to_add / 010200321100 / 3

<a id="canonical-0313113011332211-1211103020323313-1103320000202200-2311201322321131-1013122100003131-0131131130221133-3220231302031120-0002020010113121"></a>

<a id="canonical-1033130201032221-1232132001021313-2022302010120333-3231223111320311-1022230203011200-0330020212011112-3021111220003321-0212002022122221"></a>

## append property — response_headers_to_add / 010200321100 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001332203111311-3201330300033222-2220330221012232-3003132020323121-0223302002212001-0122033233232232-3120001132233100-2133312213322230"></a>

<a id="canonical-0132330110202122-1133132222312231-3103130020322131-3212210321222221-1001220320302120-0113332011313122-3201231122122331-3030020231132122"></a>

## name property — response_headers_to_add / 010200321100 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112): complete subsection reference.

<a id="canonical-1200002120131302-2003033301113000-2301133331311001-1203311023020133-2333102200310132-0123113201112122-2022220132231233-3000203313233002"></a>

<a id="canonical-0122102220220313-2213100022301113-0112303001210200-0121333210013001-3101000101300023-0322300112200322-0300102200302112-1301223130210210"></a>

## value property — response_headers_to_add / 010200321100 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-0212203322231302-1313210010112213-1331320021113231-0330301003211310-1333003111231101-2023013013100303-2323032030220021-2032300222111020"></a>

## Next pages — response_headers_to_add / 010200321100 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111132111122203-3033101022301111-1302332300111033-3031310020203201-1110301232332313-2032331121203121-2321320310322321-2000302303220312"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value — secret_value / 330031030022 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-1323121332033300-1003030231121131-2322211133003010-2303320120232110-0103103102002033-0121102131133030-0230010113210300-0130230201111100"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013312331303300-2222330001023322-1213002013002210-2013013333102031-0323212222320220-0000223103031303-0113002331120130-0011300331333321"></a>

## Direct properties — secret_value / 330031030022 / 3

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3102132133120001-2322032200131131-3102200312123023-2111333031002203-1012032100010100-2010312231001232-1320103210121032-2211030310011200): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-1202133032020033-3232201002223003-2332303300323100-0232010323212231-0320330110012230-3300030113122103-3113013333231102-1103323223122200): complete subsection reference.

<a id="canonical-3233010101321322-3313013201232212-1221020120203000-1230332330220303-3330032020210020-3310130020212203-1202211202100000-2013322301301320"></a>

## Next pages — secret_value / 330031030022 / 4

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3102132133120001-2322032200131131-3102200312123023-2111333031002203-1012032100010100-2010312231001232-1320103210121032-2211030310011200)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-002.md#canonical-1202133032020033-3232201002223003-2332303300323100-0232010323212231-0320330110012230-3300030113122103-3113013333231102-1103323223122200)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3102132133120001-2322032200131131-3102200312123023-2111333031002203-1012032100010100-2010312231001232-1320103210121032-2211030310011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210003031013121-3031113132023300-1302201110202303-3113201211133200-0133001220013333-3101121103213103-1112312221311032-1020123122020323"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 110201212213 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2103102221213211-2203011112002232-3222310032320032-3233130132202023-0333202122011130-1303223230120223-3211010120120023-3222033001101313"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203331031132122-2302311322022331-3130031200023330-2023312331321013-1303103121211033-3203302322311230-2310010230222200-0130333302100303"></a>

## Direct properties — blindfold_secret_info / 110201212213 / 3

<a id="canonical-0111312023100012-0201030133111300-2131321030212110-2010221001021223-1110103121103301-0112321333220132-0221021002230111-2222133112121320"></a>

<a id="canonical-2301121110123202-3122021321113203-1011002110003330-1003121002302101-3003203120200313-2212131120112130-0033320011221133-3302033232120102"></a>

## decryption_provider property — blindfold_secret_info / 110201212213 / 4

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

<a id="canonical-3013020020103110-1121110332111103-3232000222113201-0222012213012033-2300303312103011-3203022203202232-3210012221013021-0321330200120230"></a>

<a id="canonical-2022023020011013-1312330020011112-0111121222222132-3213001130323322-3303110213311302-3133031002010332-2220320323132210-3023123203000312"></a>

## location property — blindfold_secret_info / 110201212213 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-1110231012003003-1011111013100011-0122121011100312-1132312303001033-2223121312223300-1203233002321003-1130103323002213-0020232102020112"></a>

<a id="canonical-3023320031332331-0032012313300120-2102023321300333-2213031203011001-2132120021101210-3322033221230212-1033022210321320-3303200001310212"></a>

## store_provider property — blindfold_secret_info / 110201212213 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1310013033313030-2021113333022030-1110330010301322-0202223032212102-2031012003200031-2201022110313021-0132002302220203-1232132033011303"></a>

## Next pages — blindfold_secret_info / 110201212213 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1202133032020033-3232201002223003-2332303300323100-0232010323212231-0320330110012230-3300030113122103-3113013333231102-1103323223122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120203220033320-0132020333003011-1322132332133121-0312111010321030-2011033010020101-0001313002202120-1212202312223330-1321321220113301"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 201023133202 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-002.md#canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0332313130210102-3232332311010310-0213130222230022-2030032130100231-1130020222200021-3001331331213003-1013133312021210-0102302031000313"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013220233322021-0330202021011022-1032111003300021-0312103323303100-1003011012311000-0021010331113220-3202212121100301-2322102001121110"></a>

## Direct properties — clear_secret_info / 201023133202 / 3

<a id="canonical-0001323233103000-2131102212210230-0021113103200010-1312233010221122-0030120202322203-1310111302111023-2220311202212032-2201232223022120"></a>

<a id="canonical-3031200111320213-3030301211011202-0030101232100101-3203233033123001-3123203203112220-0001211012132000-1112202311010310-0001133223210230"></a>

## provider_ref property — clear_secret_info / 201023133202 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1303101323232033-2311301322330003-3111330303000301-3111332002311010-1230120122211013-2203320132112002-0232002301012013-0110122233321010"></a>

<a id="canonical-0020031032321003-3111102021313202-2002332330213303-0022221003103230-3022223102121233-2023102223330020-3011310310000332-0001123011001333"></a>

## URL property — clear_secret_info / 201023133202 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-1000113303233033-3233122021123223-3100231232100302-2203020132213330-1222203332000132-1332232200131012-2310021230003221-0310310031123023"></a>

## Next pages — clear_secret_info / 201023133202 / 6

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211300100131300-3311111303123221-1030321233331213-3323311213121120-2221233311331232-3221212110030031-0311323303223112-3230200033102300"></a>

## dynamic_proxy.https_proxy — https_proxy / 321102210133 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.https_proxy

<a id="canonical-2100012333203313-0133032221230031-0122133000311232-3230212201200101-1112310123103300-1212322311331203-1321131033021110-2123132102220131"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https proxy.

Upstream description:

Parameters for dynamic HTTPS proxy.

Receipt-pinned upstream constraints:

```json
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
https_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333233301032321-3321101001002320-0300213200100000-0311203001332033-2121121311001221-1211131300102212-2221303203111200-3033212110322011"></a>

## Direct properties — https_proxy / 321102210133 / 3

- [more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202): complete subsection reference.

- [tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220): complete subsection reference.

<a id="canonical-1222321323022120-0100032211323001-1021110021003123-3303012330320123-2100222003333201-1330011323202122-2012010022111010-3220300102303220"></a>

## Next pages — https_proxy / 321102210133 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022031102201001-2102122233120302-2312120230300013-1101212133132002-3321200121312023-0131010121111320-3101333031233030-3310321101320022"></a>

## dynamic_proxy.https_proxy.more_option — more_option / 312010122321 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- dynamic_proxy.https_proxy.more_option

<a id="canonical-2233030200133233-3231230013030010-3020020220201332-2131311031203311-3130100232221111-2112323321112031-3220313030202200-2202302301103331"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123200103311002-2032032130233030-1232102031013322-3302300320001222-0133100220000332-1123300011212121-2231133222031311-1230201211021013"></a>

## Direct properties — more_option / 312010122321 / 3

- [buffer_policy](resources--proxy--reference--group-002.md#canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123): complete subsection reference.

- [compression_params](resources--proxy--reference--group-002.md#canonical-1222132113133101-3012113301330313-1001231201322310-3102232230330302-2311333002202122-3101002112021213-0301331130312113-1223113123330210): complete subsection reference.

<a id="canonical-0120301120131103-2132222303132210-2230010301111310-0101301232231221-1230223112031231-1320120223021012-1112221211020313-0203203121120311"></a>

<a id="canonical-0130103131212132-3201322200300301-1121101102301032-1112001220113120-1101122133331223-1232330332230111-1213311020000112-1333202322003221"></a>

## custom_errors property — more_option / 312010122321 / 4

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
    "metadata": {
      "confidence": 0.75,
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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-1313030123210300-2213213220000330-1032332300131232-3101000213121002-1102230112103203-1102032011203031-0000332210022113-1210232330302000"></a>

<a id="canonical-1102232333313101-3112102103301212-1322212020003013-1212103033011121-1301120333220203-0331011010332320-0320000013122323-0021232022332030"></a>

## disable_default_error_pages property — more_option / 312010122321 / 5

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_path_normalize](resources--proxy--reference--group-002.md#canonical-2012321031232122-3103310110123120-2212322223233021-3221100003023213-3103320011002121-1310230233220211-0201013323111120-2310322133221012): complete subsection reference.

- [enable_path_normalize](resources--proxy--reference--group-002.md#canonical-2011301323311002-0100322122311323-0033200002112312-2130133101033220-1100122111003300-0332213002323011-3222122013313333-2003100011133311): complete subsection reference.

<a id="canonical-2020003313312211-3310112200112201-2220101200202223-0231113211011101-3011202211231021-1223331101100130-0310300001331312-0311230313030001"></a>

<a id="canonical-3300220320312221-2210302032021223-3011210102020132-2311222332231330-0312031030123132-2113332100200303-2100211312132302-3330120001012003"></a>

## idle_timeout property — more_option / 312010122321 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-1133031030113001-3211122221310122-3100200031331231-0212030022132112-2020111131230310-0110123120311000-2310002321120023-0321130232003202"></a>

<a id="canonical-2231231310322212-2113021331331301-0131023233030330-3200112012211122-1323020133232302-3211233112333231-0010100033130012-2003130212220001"></a>

## max_request_header_size property — more_option / 312010122321 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-2111313221213020-2012021110002133-1202231210133100-0133331301032122-2023322331303132-0212101112133301-0322333123101210-3202233221210030"></a>

<a id="canonical-2213321322100022-0021321303033203-0011233002333333-1313001213331003-1223333220000202-0312210321021100-2330011211212131-0001120320202020"></a>

## max_requests_per_connection property — more_option / 312010122321 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-2210210100110313-0022103330230330-2121233301303310-0222013023232232-1203321013323021-0311212202110200-2133301313121203-1110330330320131): complete subsection reference.

- [request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203): complete subsection reference.

<a id="canonical-2233332213222302-2310300122112333-2122003112003030-2312203000103123-2033223300120300-3130203132211022-3311110033231000-2320300003231121"></a>

<a id="canonical-2233200031330121-2202203012200033-1022312212012333-0221122300001212-3023310310333210-0223021020121331-1131031031123213-0232033303110300"></a>

## request_cookies_to_remove property — more_option / 312010122321 / 9

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121): complete subsection reference.

<a id="canonical-0333222211023020-0202111230101111-0013132323101331-0333002103200233-3113331100111002-0330101021222100-1311231023010320-1030133022011333"></a>

<a id="canonical-1223131302020133-0200323011000021-0003302332021102-0333223330132102-0300020221333231-3310033023331203-2011123011330012-3300313320222130"></a>

## request_headers_to_remove property — more_option / 312010122321 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310): complete subsection reference.

<a id="canonical-1333121220313131-1002010123102033-0023123312002331-2320312030221110-1302103230333132-0313031310100112-1302321211021203-1102211302130302"></a>

<a id="canonical-1331020020033222-1200233121033330-2202203333113203-2230230300013300-2032332130223330-1313001123210100-3233322101002230-3200322021020032"></a>

## response_cookies_to_remove property — more_option / 312010122321 / 11

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310): complete subsection reference.

<a id="canonical-3200310321210311-1212323313231030-0033331230231131-1112300133012131-2202201310220203-2312320203032301-0232302010123030-1003233103031130"></a>

<a id="canonical-3133132332210011-0131121022132013-2103131000012203-1232112002021230-0030301302213331-0320000313233200-0213010032222213-2131033212111131"></a>

## response_headers_to_remove property — more_option / 312010122321 / 12

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013233101120321-0130031032201000-3022311000123221-3300312031121201-3101201232213310-2110302311323123-3202000113133303-3011023333321001"></a>

## Next pages — more_option / 312010122321 / 13

- [dynamic_proxy.https_proxy.more_option.buffer_policy](resources--proxy--reference--group-002.md#canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123)
- [dynamic_proxy.https_proxy.more_option.compression_params](resources--proxy--reference--group-002.md#canonical-1222132113133101-3012113301330313-1001231201322310-3102232230330302-2311333002202122-3101002112021213-0301331130312113-1223113123330210)
- [dynamic_proxy.https_proxy.more_option.disable_path_normalize](resources--proxy--reference--group-002.md#canonical-2012321031232122-3103310110123120-2212322223233021-3221100003023213-3103320011002121-1310230233220211-0201013323111120-2310322133221012)
- [dynamic_proxy.https_proxy.more_option.enable_path_normalize](resources--proxy--reference--group-002.md#canonical-2011301323311002-0100322122311323-0033200002112312-2130133101033220-1100122111003300-0332213002323011-3222122013313333-2003100011133311)
- [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](resources--proxy--reference--group-002.md#canonical-2210210100110313-0022103330230330-2121233301303310-0222013023232232-1203321013323021-0311212202110200-2133301313121203-1110330330320131)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323100112320331-1213200102322210-0013113223210102-0210112120311331-0133311023011230-2030021203230203-2320031231013320-0323201301302113"></a>

## dynamic_proxy.https_proxy.more_option.buffer_policy — buffer_policy / 101120331112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.buffer_policy

<a id="canonical-2201121022133030-0101122321201203-0310223122101231-3222113211021310-3120020001122032-0100322200322112-3123223311222113-0313321322310303"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

Receipt-pinned upstream constraints:

```json
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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133202101231020-1130122130310221-1103222113022201-2201300310012222-1030230000013120-0232012233013332-2102310200031233-2223212010132132"></a>

## Direct properties — buffer_policy / 101120331112 / 3

<a id="canonical-2002021103022022-3202332002103332-1030101320100133-2112023231303200-1202300313232133-0130201001313220-3102322112132010-2101302313221133"></a>

<a id="canonical-1200212223123232-3120332231233331-3131221200103202-0101021013231013-1301201223031320-2023331130101003-1110223310112203-1200030020222323"></a>

## disabled property — buffer_policy / 101120331112 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3110033021333123-0003222013110223-0111313201231221-2223310320111321-0130021202001230-1310013030002031-3131331210002230-2113211320000333"></a>

<a id="canonical-3320101231120313-2021133210031003-1303321210022203-1032311232312220-3231313211311303-0031213301102321-0233123203003323-1203211031220321"></a>

## max_request_bytes property — buffer_policy / 101120331112 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-0002310310103131-1302133032032310-0213202122212110-3023232033311131-0130323133232303-2022011110312222-1110002033000130-3102121221112112"></a>

## Next pages — buffer_policy / 101120331112 / 6

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1222132113133101-3012113301330313-1001231201322310-3102232230330302-2311333002202122-3101002112021213-0301331130312113-1223113123330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223031031202133-3120301111021021-0213113231321312-2022101301001011-1130001220233032-2323021010110100-1331211322030100-2300323121121333"></a>

## dynamic_proxy.https_proxy.more_option.compression_params — compression_params / 112101202213 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.compression_params

<a id="canonical-3220220331212302-3122301000220211-0130210223110023-2033213010213330-3212312320133211-1211331133022223-1322131311133131-0312333133322013"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132222222223010-3133122120121103-2103232212332132-2223212103112321-1022212030322303-2002100233033002-1003131120331300-2330220013200120"></a>

## Direct properties — compression_params / 112101202213 / 3

<a id="canonical-2112113302012012-3121020300113223-2321013212003013-3000322232122323-2110022112000032-1310333131231321-0002301202111200-2331102331333002"></a>

<a id="canonical-2133102121022210-3030333222203333-0012031213123100-0002023000132312-2211003000321023-3322121311221101-3221313020112002-1111133113101100"></a>

## content_length property — compression_params / 112101202213 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
}
```

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

<a id="canonical-0212032303300121-0020113001132301-1001030321000221-2311231200123021-0311033113030233-3022332010200220-3120123111023301-2332330020002312"></a>

<a id="canonical-0033011201113112-0121032013110230-3113130231132323-3321103300002020-3313300232230303-2123302323010110-3202011223111323-3312002123221102"></a>

## content_type property — compression_params / 112101202213 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

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

<a id="canonical-0102211203130010-2111323033102112-0020003110322101-1202311110230103-3121230220002222-2320200223101131-3321122032212003-1002110112013021"></a>

<a id="canonical-3332201101011210-2333112010331102-0031010032231223-1321223312333323-1320201211330003-1213012212200100-0311233200320313-0321103113101230"></a>

## disable_on_etag_header property — compression_params / 112101202213 / 6

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

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

<a id="canonical-0131212132302121-0123131023033013-1123023000110200-3303301002211331-3222130313123131-3200321012112303-1303232021222033-1322233131211330"></a>

<a id="canonical-0202030223230220-3101333203333123-3033003211313233-3322322312130303-3112122022333110-1002130120333021-2303121211012113-3210020033322133"></a>

## remove_accept_encoding_header property — compression_params / 112101202213 / 7

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

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

<a id="canonical-0300030131222031-1001003011033003-0100010313211012-2302332031003003-3003020120031332-3233002113303221-0311313321221100-3022102303302122"></a>

## Next pages — compression_params / 112101202213 / 8

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2012321031232122-3103310110123120-2212322223233021-3221100003023213-3103320011002121-1310230233220211-0201013323111120-2310322133221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320310202220300-1022120110110213-0330300121200130-2021132312332230-3333211011223132-3332212311023332-2003223230101200-2213203233032010"></a>

## dynamic_proxy.https_proxy.more_option.disable_path_normalize — disable_path_normalize / 212000122112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.disable_path_normalize

<a id="canonical-2333222322331033-3222333022323123-2202322233120303-1310321031232131-2000011133110311-3210001302331130-2033120301003120-1330001211113310"></a>

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
disable_path_normalize = {}
```

<a id="canonical-3022111113303221-2321021023332330-2302111112330230-1231122212202200-2002013120101133-2332123000120231-0201012323310311-3010131313101132"></a>

## Direct properties — disable_path_normalize / 212000122112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200203021203120-2203112113132003-0321322310221012-0310111023103103-0103000213130210-0010112032321321-3303033323333011-2103103330012030"></a>

## Next pages — disable_path_normalize / 212000122112 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2011301323311002-0100322122311323-0033200002112312-2130133101033220-1100122111003300-0332213002323011-3222122013313333-2003100011133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030110322211020-0300301123000222-2311031310021331-0122013113001230-3202321013111022-2110031022003303-0021332022011211-0030013012202310"></a>

## dynamic_proxy.https_proxy.more_option.enable_path_normalize — enable_path_normalize / 021321121202 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.enable_path_normalize

<a id="canonical-3333302232022233-2133110122231321-2211110303020302-0033121013122233-3333203300031203-2322123000101300-2133101101322123-1030302103100201"></a>

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
enable_path_normalize = {}
```

<a id="canonical-2013230220020322-0113220203300211-0231222331133012-3101330302023110-1222112331323021-2001023302212212-1233300110231211-3320211032331110"></a>

## Direct properties — enable_path_normalize / 021321121202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100203231033200-3321331300102320-3002130123321100-2022303313330321-3112210201101232-1020002203022023-1111012323330331-1003201331332302"></a>

## Next pages — enable_path_normalize / 021321121202 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2210210100110313-0022103330230330-2121233301303310-0222013023232232-1203321013323021-0311212202110200-2133301313121203-1110330330320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231331303011132-1023222112130003-0301103021320100-0020120032312200-1310112022020022-0200222102200112-1310230113003230-0121300320121003"></a>

## dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection — no_request_limit_per_connection / 232001021210 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection

<a id="canonical-2202220132301332-0233211122313101-3233330010112233-2221312112311012-0333322201002021-3313002231102022-2010101030303030-1013313030022010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-3010302103221110-0210321111221002-0232031110122200-2032012310221030-3301003321103103-0032221313331211-0220301000033333-1002121222302112"></a>

## Direct properties — no_request_limit_per_connection / 232001021210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122002332331010-1303332231002230-1332201311013003-0011313031021030-2123013231222332-3012322333033130-3022333313122203-0103103011133121"></a>

## Next pages — no_request_limit_per_connection / 232001021210 / 4

- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210322320020020-0112000233122230-1222201331133302-0033012021002010-1010111100211311-2321203122300322-2212231121323132-0323121121000123"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add — request_cookies_to_add / 322323022301 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add

<a id="canonical-2031210233122310-2303310112123101-0121203120323031-3013303230313120-2010100230212300-2003130301332310-0001120002120033-3212021031123220"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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

<a id="canonical-0202001123110010-0203111100231211-2331333110123333-2113131002001132-2101003221011313-1303120121201001-2133221201323230-2313112033323232"></a>

## Direct properties — request_cookies_to_add / 322323022301 / 3

<a id="canonical-0233121311213230-2320122132202220-3310111002302213-0030320122311312-1320111112110000-2203112313300130-1101100233322031-0211212311210330"></a>

<a id="canonical-3320130122320223-1233131323112301-1113212010112302-3011133122201111-1003023131322222-1223110232310322-3122011023011300-2020203002002320"></a>

## name property — request_cookies_to_add / 322323022301 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1202201130221202-2012020233221231-2230322000333212-2231222321330200-0231112110110110-2100202131103130-0122322320312311-2121213210231121"></a>

<a id="canonical-0312211100013202-0113102032130312-0121311133332300-2101332321332320-0232000033023011-3203020023020230-1003032211100220-3133200232333203"></a>

## overwrite property — request_cookies_to_add / 322323022301 / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000): complete subsection reference.

<a id="canonical-3301311032222000-2132003130312102-0222222001302200-2023201332221203-0203000113010010-2210121303320223-2123221123132133-3101311130220000"></a>

<a id="canonical-1121130030212121-2102323330212332-2212313300313012-3020123103220011-1322221200010021-2323012311011331-1001231313202103-2111011323233233"></a>

## value property — request_cookies_to_add / 322323022301 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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

<a id="canonical-1213121113223010-2302220033303103-1212012200230002-0103310332113313-1101002221231321-3303211332113211-0100012012133023-0022201003000100"></a>

## Next pages — request_cookies_to_add / 322323022301 / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
