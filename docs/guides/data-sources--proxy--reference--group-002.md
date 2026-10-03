---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1230023013312221-2100203202313011-1133130321112210-3310130213121101-3323301132311120-2031133111113021-3100313201231312-1033202030102133"></a>

## content_type property — compression_params / 323110123301 / 5

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

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

<a id="canonical-1302031213300313-1013113230021212-3331032202031310-2001023212300000-2310210233111123-0312010313110220-0331322010220030-3122202102103033"></a>

<a id="canonical-0310111200022311-0113230130100030-2322230330121120-3330301102022011-1122120301012030-1123003011113313-2120201011222133-0300320010313212"></a>

## disable_on_etag_header property — compression_params / 323110123301 / 6

Type: `"bool"`. Computed.

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

<a id="canonical-2111101123033202-0011333333221022-0030311000103111-3203303233133113-2331220323230102-1121303311210103-1201223120331231-2110002012303222"></a>

<a id="canonical-2210222202012010-1112010303002001-3332102102011102-3233031011103331-0300123313012223-0023203221230102-2230102123332203-1111231032132021"></a>

## remove_accept_encoding_header property — compression_params / 323110123301 / 7

Type: `"bool"`. Computed.

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

<a id="canonical-0231232002003023-1312213011200200-3022032333113030-3230102021101113-0030133313202230-2000211102033023-3133100312201220-3323100301212212"></a>

## Next pages — compression_params / 323110123301 / 8

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3203323212312201-1231032110031320-2020103200220313-3212233112202321-1030133022313212-3212101120003222-0012012010330012-2000103003213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002032202222102-3210230320303321-2313313212212223-0100320022313000-1331123031210233-1133030100321102-2231123030112310-1330323112032320"></a>

## dynamic_proxy.http_proxy.more_option.disable_path_normalize — disable_path_normalize / 031103130320 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.disable_path_normalize

<a id="canonical-1113233131312122-3130102222100323-3231320311313003-1310221021330023-3030222122111001-2101022103200131-0013322303312130-3131013210211301"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2212122301133330-1122101002301212-1320211210312233-0022213000023322-3131021212321130-0132001020111222-3213332130133212-3010331101132123"></a>

## Direct properties — disable_path_normalize / 031103130320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330133322130102-1312023100023311-2013311003221303-2113210330133210-2311313201220220-2320230202312011-0120032301203032-1333313233003111"></a>

## Next pages — disable_path_normalize / 031103130320 / 4

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3200333303133021-3231310312203010-2233202233001200-0113213233012032-0020231311133133-1130212020030322-2121010020313133-2111331200210323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320001103132132-0222020012301101-3020311203202221-0112300120212311-3220002302210012-0200102110200320-2112322033121303-0310121301022000"></a>

## dynamic_proxy.http_proxy.more_option.enable_path_normalize — enable_path_normalize / 210221011323 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.enable_path_normalize

<a id="canonical-1013200221322220-2303203122231100-2220211203312111-3131012311313011-3032010320121222-3032020223110100-1022222233312332-1212000331030330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1311103020330022-1230132022213001-2223232132232102-2320020323300312-3001323230020232-1030123100110122-2301031133320321-1103032231021231"></a>

## Direct properties — enable_path_normalize / 210221011323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300310131202213-0210202211322131-3221120103313203-2121023212111111-3103103032301203-3111300020213122-2300031210110333-1100302212210123"></a>

## Next pages — enable_path_normalize / 210221011323 / 4

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1022212012133233-2230231101300000-2201203102103101-2233110103201202-3020303221310220-2210131311112023-3312113102322221-2313213012032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112022213123213-1120113100130232-2302123103220102-0031332000020022-0222010102300212-3031231130021113-1023132002032200-3320301333100020"></a>

## dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection — no_request_limit_per_connection / 322230002301 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection

<a id="canonical-2000323201012320-3100103022020310-3123303320001201-1030021200210023-1030001230021201-2030100323302131-3122122130100233-0310212333023031"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0100010031221010-2103211122120011-0213033232301220-0021022020033211-2102033102322022-0313233012131311-0201030022130022-3311103200332013"></a>

## Direct properties — no_request_limit_per_connection / 322230002301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121233122133023-0332330230330023-0222323000233311-1022021310313023-2210221123131130-3121323030020121-3132110320212112-3320111302211033"></a>

## Next pages — no_request_limit_per_connection / 322230002301 / 4

- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331011021232230-1130000230210100-3003131013233301-2133202213100111-1103310220013332-1202322030201033-2030330032221131-3231132320023220"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add — request_cookies_to_add / 222000103212 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add

<a id="canonical-2130320311301030-1232110232001001-2303031323322313-3101213313002222-2320230031302332-1233131000130212-3312300322001333-2131311311101133"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

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

<a id="canonical-1211113003301010-1032202003323222-1122020010012323-2200233203232013-3112202301013313-0320110122012202-2001121011113221-3301322212131012"></a>

## Direct properties — request_cookies_to_add / 222000103212 / 3

<a id="canonical-1232112232231320-2103131202200113-1020211021103210-0201230001230332-1032322021233103-2300222032120032-3230322313032102-0001123012122012"></a>

<a id="canonical-0132331123333102-0212232321021001-3320313103011331-1311012303330211-0320010032113031-3222112113201003-1102121321223001-1031103211120122"></a>

## name property — request_cookies_to_add / 222000103212 / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-3311101103220112-1002132233203333-1203021102013101-0202222313221000-0232120033101211-3030330310102031-3233121212313222-3333023122011032"></a>

<a id="canonical-2200303023123231-0302203302233321-1100201211211222-3013132233301002-0013012000123011-1110211003010021-2122231232210230-2233213033323010"></a>

## overwrite property — request_cookies_to_add / 222000103212 / 5

Type: `"bool"`. Computed.

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032): complete subsection reference.

<a id="canonical-3023030233021220-0202311202300332-1033012333012113-1103322301333322-3320000023031230-0202323203200020-3323023101023000-2112333212010202"></a>

<a id="canonical-3123311023003012-2122002123110311-0223331312132000-3332023002213303-2120220222232102-1110221222320222-2112111303003030-3300301212120020"></a>

## value property — request_cookies_to_add / 222000103212 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

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

<a id="canonical-3202332102103010-1200222032022312-0200020130332210-1221322133011020-2333112121233002-0002310121112322-2012033213012121-3010112220130332"></a>

## Next pages — request_cookies_to_add / 222000103212 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203122313300323-0113302203013313-0223131303102103-1213313031133332-0100031300112320-1112332022101012-3030211112102131-2101311111022202"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value — secret_value / 302131230030 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0333230303230001-1233300210233331-3020210231321132-1031311122212001-1220232330130111-2000032233002001-1213012030001221-1033331012303011"></a>

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

<a id="canonical-1003022233300212-2202123321021100-3021323200112233-3130021221112101-3100112111013101-0022211023311300-1022200320333203-2021322032000202"></a>

## Direct properties — secret_value / 302131230030 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-0313013020021030-2002113111010212-2210121310232331-1100011000010003-3223202131130123-0021001311223120-3102322223312111-0100120033123322): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0232103032301203-0201202110301122-2213321000022112-2323020312201033-0023333220302310-3003330233110230-0003031003202111-3032121231102013): complete subsection reference.

<a id="canonical-0032021331220100-3310013203122200-1110011121002300-3110120133111001-2310301203112322-0223231221103203-2102303210311333-1130313221331220"></a>

## Next pages — secret_value / 302131230030 / 4

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-0313013020021030-2002113111010212-2210121310232331-1100011000010003-3223202131130123-0021001311223120-3102322223312111-0100120033123322)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0232103032301203-0201202110301122-2213321000022112-2323020312201033-0023333220302310-3003330233110230-0003031003202111-3032121231102013)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0313013020021030-2002113111010212-2210121310232331-1100011000010003-3223202131130123-0021001311223120-3102322223312111-0100120033123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212223032212133-2103033011231100-1010231203113111-0121112033012311-1223200220021330-1131001210331221-0202003312230212-0123312211220330"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 311332212331 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1211303312220022-1312022323120323-3011211303333311-3132111010121232-0113300330031023-3223321133003011-2110301210112110-1110313202320011"></a>

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

<a id="canonical-1222320222302022-2231102323310012-2230212113331201-2333113030020001-2121103323033332-1201232001203103-3333003030000312-1102320333202023"></a>

## Direct properties — blindfold_secret_info / 311332212331 / 3

<a id="canonical-3303311232111323-3222113001012120-3032223032120130-3212102210010203-0220203220211303-0131100132233333-0212100021322100-1312302103002131"></a>

<a id="canonical-1101113233311032-1232032100303013-3032131320120201-1212022110233210-1313221323131031-1212302231102113-0313020120211312-2310201320332031"></a>

## decryption_provider property — blindfold_secret_info / 311332212331 / 4

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

<a id="canonical-1223030023101231-3112022210032323-0102122112113111-0200221231203023-0320230011231121-3100300030120021-2031302313331323-2200130002221321"></a>

<a id="canonical-3322322332002032-2333130030001200-1222210000220123-1120001301311121-1112030313131030-3203200220322011-2023011013332001-1232312232221030"></a>

## location property — blindfold_secret_info / 311332212331 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-3210113310033300-2322002110333102-0022212201033331-0102133331203003-1010023310233312-2303210113322233-0131101332023020-2031030211000200"></a>

<a id="canonical-1221323202211202-1122323132331022-1212302033322323-1112210132221212-1311003020012002-0123302230113113-1030311230321112-1300223202132033"></a>

## store_provider property — blindfold_secret_info / 311332212331 / 6

Type: `"string"`. Computed.

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

<a id="canonical-2101311100310031-2111320021000111-2103002001332303-3121331210230103-1010230203320032-3311320123331210-2003212101223001-0020121211330333"></a>

## Next pages — blindfold_secret_info / 311332212331 / 7

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0232103032301203-0201202110301122-2213321000022112-2323020312201033-0023333220302310-3003330233110230-0003031003202111-3032121231102013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020002203000100-3230131211112002-2000330230023300-0130011131212211-2313030301003001-3230330013221213-1232333000213112-3230201333321221"></a>

## dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 010211101130 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2310323101130232-1133301231011132-2032103022123020-2022112230133001-3003200011223031-3212120120233020-2132021223301100-2011033121212132"></a>

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

<a id="canonical-3001102000003300-1302203323133123-1201110321012011-3033201101023023-0221100223311033-0110323331103222-3322032203231021-0023231031031210"></a>

## Direct properties — clear_secret_info / 010211101130 / 3

<a id="canonical-0312331223033230-0211003030101323-3131213131002022-2222333230102330-3313231300230201-3103333121201112-1301202210332330-2131103103310111"></a>

<a id="canonical-0113021232020320-2333333323121032-1201230302033102-1010002311223003-0102100102320230-1310320231220332-1332012212001333-2330023123312331"></a>

## provider_ref property — clear_secret_info / 010211101130 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2323300032213010-3121211031212221-3120330123120222-1100002232200201-3002331033103313-2013110012232003-1122023222312003-1001002320200312"></a>

<a id="canonical-2323113222132103-2232021013030003-3002231302103122-3301303012020021-2230323020211003-1112322200231220-2030010200002232-3032111321013223"></a>

## URL property — clear_secret_info / 010211101130 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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

<a id="canonical-3123231033011303-2210303032303132-1223123320103031-1103013210223321-1223031333201033-2303222003010313-3010201133332322-1020103232201013"></a>

## Next pages — clear_secret_info / 010211101130 / 6

- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010203021031300-0000110310313022-0310012202210313-1310112232213122-0323301230221202-3302212320320011-0122011222100133-3002320321333033"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add — request_headers_to_add / 101100311101 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add

<a id="canonical-3033212330002301-2223112230232320-3303200103112333-1112230022013013-0313213311013121-0302131013331302-3222000332122222-3101001032101222"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

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

<a id="canonical-0322133100322131-0130221031130100-3023331010331200-2120030023331332-0012323331302223-3300113230332000-1131100100003123-2131303322302002"></a>

## Direct properties — request_headers_to_add / 101100311101 / 3

<a id="canonical-2321122302230302-2013133233200101-1133331021120212-2230323313212122-0031213222121201-1231321303101310-3003003030102231-0113021320311011"></a>

<a id="canonical-3100323010123323-2320301020220223-2130023020201220-2111132121203033-3232310130310321-1113332021210331-0323310312102312-0130113101331202"></a>

## append property — request_headers_to_add / 101100311101 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-2320132322211121-2102333302210233-3323020111200302-2221023103003132-0223110202030021-0310133113131321-0003022223030302-3333323203303331"></a>

<a id="canonical-2121103002103120-3200302231020103-2032220212013101-1313131100131100-3230311012111302-0030312011123301-3302303232031323-0020123320123223"></a>

## name property — request_headers_to_add / 101100311101 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012): complete subsection reference.

<a id="canonical-1223303110101311-2313013223303000-1121030211303130-1203110311322321-3003021220110333-3110301222221223-0131331131332302-0021122331201302"></a>

<a id="canonical-3211323130113033-0031301020231221-3233202002030322-1022103300332133-1003302001112232-1002101203213320-2111133221020001-0132231301302303"></a>

## value property — request_headers_to_add / 101100311101 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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

<a id="canonical-3332023122022231-0223132100102133-0321033311012010-1003001030000323-1203020333311123-0121332332023320-2021021323100112-2121030021020200"></a>

## Next pages — request_headers_to_add / 101100311101 / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013231030322001-2232221302232233-2031011130332311-1313123331310332-3221022010321221-2103121122130323-3012323303200122-2313033021311011"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value — secret_value / 333220100120 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-2130230302102212-1132001110323302-0203313122101123-3030122312133130-0132330203001310-1201103130313330-0002021201121002-3313230321123000"></a>

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

<a id="canonical-0211300003302231-0131103332113031-0212322232223310-0203220201233033-2221322030023333-3333130032010020-3131131201123030-1312213103112032"></a>

## Direct properties — secret_value / 333220100120 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1030023133302332-3102110121220301-3012111312210130-3203020223200132-2001033331112200-2131221333003032-0032013322131222-2112002033123313): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-2100310300023120-2120103332133310-3132022220322302-1202113210122131-1001330031131210-1121012021303110-1132130300022231-2112202122302212): complete subsection reference.

<a id="canonical-0223031302213311-0020223130032123-1313313131112330-0331232003210100-0330203302122001-0033220210122120-0031122313101233-3111010030201120"></a>

## Next pages — secret_value / 333220100120 / 4

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1030023133302332-3102110121220301-3012111312210130-3203020223200132-2001033331112200-2131221333003032-0032013322131222-2112002033123313)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-2100310300023120-2120103332133310-3132022220322302-1202113210122131-1001330031131210-1121012021303110-1132130300022231-2112202122302212)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1030023133302332-3102110121220301-3012111312210130-3203020223200132-2001033331112200-2131221333003032-0032013322131222-2112002033123313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222210000210113-2002000122332003-0013213121233220-2201300032110120-3322213321032200-2300320331111233-1221312030032022-0112122013301301"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 321300003203 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0230110211222232-0300211231021233-1333102222010201-2313310000200202-1003333101033102-3313230113030221-0101313331212221-3001112321010023"></a>

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

<a id="canonical-2201031230302121-3120123130003313-0120210322133320-2110000320212130-2033003103013021-1332203120222221-3312310310301320-2111030133123020"></a>

## Direct properties — blindfold_secret_info / 321300003203 / 3

<a id="canonical-3001023131031010-2231330200312133-2322101230312331-3011320223022223-0000030033333020-0120001031010132-3320333303121212-0132312230200200"></a>

<a id="canonical-3302013123012200-0113200212322123-0221220310030111-0030303200002132-2011122113223122-3110221033223133-0321121131223302-3102002200302103"></a>

## decryption_provider property — blindfold_secret_info / 321300003203 / 4

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

<a id="canonical-1311311131030100-1210002313201333-0112333221213020-0210112201211020-3321312032333331-2012130112011321-2033102121330320-1100202011331101"></a>

<a id="canonical-0101332033100011-2020321303000330-1130210232111223-2301112230021032-3021333112203023-1033233222313011-2111003120131030-3000202000030330"></a>

## location property — blindfold_secret_info / 321300003203 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-0202003202012110-0213223010031310-2213300200100313-3310302010322121-1231232130010303-3201110010122131-0002121023213222-0003210021231212"></a>

<a id="canonical-0300310313332030-0103113302233032-0200222311222320-0302321102132010-1323202210232332-3010300112213011-2122220100022001-0031310103303321"></a>

## store_provider property — blindfold_secret_info / 321300003203 / 6

Type: `"string"`. Computed.

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

<a id="canonical-3010003031302333-0020131211223002-1103333331300130-0211212221323213-1130321212122021-2313213213331003-3201221002223100-2233333230110031"></a>

## Next pages — blindfold_secret_info / 321300003203 / 7

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2100310300023120-2120103332133310-3132022220322302-1202113210122131-1001330031131210-1121012021303110-1132130300022231-2112202122302212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032021120112302-1132211213020220-2221331020023103-3131232113331131-3132330032320320-2321300221303031-3013130000220310-2111333100301330"></a>

## dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 020202003202 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012)
- dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1330030230210000-3221210120301102-2103012100032121-2020030333101001-0230203232011213-1222330311320303-3112130302111012-3303322031123201"></a>

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

<a id="canonical-2230013111023232-0220033320211331-3233221000201010-3200122010210030-2220211021321112-1230303303021011-3332011132202032-0201112212313311"></a>

## Direct properties — clear_secret_info / 020202003202 / 3

<a id="canonical-1312023111303332-0313231320121023-0031221020031331-0332001120210012-1312122113220000-0013030221311233-3103203123002232-3112200113321200"></a>

<a id="canonical-3112212320232133-3011313022333312-1102313200321230-1122310211003330-0110311201013311-2223210200111112-3112221203132332-2133032011012231"></a>

## provider_ref property — clear_secret_info / 020202003202 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3330131200223123-0333131311021013-0111121223230110-2200201312111333-1113331120203120-2033302312131222-0311321303202233-2033320011212321"></a>

<a id="canonical-0122201010103122-0330012123232012-3230221220300110-3033031023001002-2303202110220100-2011112131012030-2012101312010200-2200310332321023"></a>

## URL property — clear_secret_info / 020202003202 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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

<a id="canonical-2121201322213221-0220313012311023-0020200213112123-2303123101011302-3232001020110333-3011332232222132-0223220222101022-2321030320000120"></a>

## Next pages — clear_secret_info / 020202003202 / 6

- [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111132200222230-3120023132201022-3000123213032032-3101230110331330-3010203300133011-1113230233132322-2131211320023333-0101132333000200"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add — response_cookies_to_add / 003130103022 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add

<a id="canonical-2303221113023010-3131012220121131-0330120301220120-3100110331223003-0303233132033021-0112220111212022-1232323201333112-0213302133110011"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

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

<a id="canonical-1001100321301310-1122101000220111-3013331110030311-2110211013302010-0112333013222120-1123323111032133-0011331103231032-1033230023033222"></a>

## Direct properties — response_cookies_to_add / 003130103022 / 3

<a id="canonical-3023013221131033-3202201310301300-2332010032102102-1233031200300201-1023132003100011-1100022102322232-2121201331101322-2132221033322223"></a>

<a id="canonical-0301131103110103-3123302230100322-0202330230313201-3203320001332320-2131220111113310-0022313201120212-1203310313001020-3312211033012103"></a>

## add_domain property — response_cookies_to_add / 003130103022 / 4

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

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

<a id="canonical-1033211023122220-0131120210001303-2322021230222232-0303111203013313-1231321031332130-3131313023012223-1203312000221322-3100331313110031"></a>

<a id="canonical-0102013311012221-3013312103011123-1000002022102012-1131023211121113-1311222200123123-1121331331102313-3122331231310032-0330032323033312"></a>

## add_expiry property — response_cookies_to_add / 003130103022 / 5

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

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

- [add_httponly](data-sources--proxy--reference--group-002.md#canonical-1330320023303103-2302022130231131-2333232331213210-0002003331120312-0201333020223020-1232133321122300-1122113013020030-1220003331100012): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-002.md#canonical-3001022031232111-0122001320123233-1310332333101030-1320233302212010-3012310113033010-0122210321033202-0312311303112102-1112021130023003): complete subsection reference.

<a id="canonical-2003222133030332-1123332022110010-3112233213021133-0123312131313012-3121212120131112-3002201333103213-2232000001100111-0020003212303110"></a>

<a id="canonical-0321232330100031-1133121323023203-1132220323020122-0301022032301123-3321010312012320-3232012022101011-3202111103031233-0121302233133030"></a>

## add_path property — response_cookies_to_add / 003130103022 / 6

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

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

- [add_secure](data-sources--proxy--reference--group-002.md#canonical-0011102001032030-2331301112031120-1300022220021201-3320121023213032-2302001023033003-1322213002221113-1223023120131120-1012201320030211): complete subsection reference.

- [ignore_domain](data-sources--proxy--reference--group-002.md#canonical-2000320130332032-1233020232300330-2002133203003311-3312200230233212-2122220303123211-3320330322222230-1300010131133120-3210223023123231): complete subsection reference.

- [ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-1210112000100300-0332300323233332-2212132223111012-0310310211302021-0110320131032113-1220113113132301-2310132120112302-3033222131021333): complete subsection reference.

- [ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-3232020031003300-0031310323013133-0022131022233013-1231210333020221-3323231302203112-1133002233213031-2022111003222120-0013203333303132): complete subsection reference.

- [ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-3203111331101100-1133310211331120-1103032202331311-1213221210300333-3131001031221321-2200003330023222-0033023003200020-1120202002030101): complete subsection reference.

- [ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-0233030130130312-3322213313223010-2321000120031233-1003130313222231-1020300333211303-3031311032310112-3212202220120120-3221011311100230): complete subsection reference.

- [ignore_path](data-sources--proxy--reference--group-002.md#canonical-1113133120202223-1022303120003311-3112222232223321-3232211120002030-0103200122122123-3022000201310203-0110121311103213-1003231022022322): complete subsection reference.

- [ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-2103030113132301-0301222332011111-1121233230131310-2123221130233021-0332220000130213-1221012000111321-2300122220133232-2233223111131321): complete subsection reference.

- [ignore_secure](data-sources--proxy--reference--group-002.md#canonical-3233220130010331-2032012212000020-2313232310022133-3133230003311210-0000310201313221-2220112312012011-0012121021113322-2312030230120230): complete subsection reference.

- [ignore_value](data-sources--proxy--reference--group-002.md#canonical-0022102133201032-3321113032333101-3331132223220132-1302203111101301-3032120221020231-3000321310112230-3220320323312202-2200033122032220): complete subsection reference.

<a id="canonical-2003031311121321-1223131020202121-0110213100303031-1331131133103000-1301122311021202-1103101100111030-3001022331020023-2331302301222022"></a>

<a id="canonical-0031202020210103-2013323211232101-3333100210020113-0030111000032020-1230220010123300-2100300022233123-2020333131201331-0231130322311002"></a>

## max_age_value property — response_cookies_to_add / 003130103022 / 7

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

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

<a id="canonical-3313201223020301-0203303212322000-1113201022203120-1130012021020013-0123312312002312-0130012321230022-0231121311001222-1100023122333101"></a>

<a id="canonical-0313000000101211-3331211333233132-2210333032320132-0133130201312230-1220030301313331-0222131102013113-1200013230313201-2312122223222122"></a>

## name property — response_cookies_to_add / 003130103022 / 8

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-0110222213003303-2020220123133310-2210311110223312-2013032112032001-1002110300330130-3323031320230111-2320202113021320-3300222311203012"></a>

<a id="canonical-1101122333023023-2232300123310132-1203021110311121-1133210111220210-3112122022312123-2030223210230103-2223221213223013-2322212213311031"></a>

## overwrite property — response_cookies_to_add / 003130103022 / 9

Type: `"bool"`. Computed.

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

- [samesite_lax](data-sources--proxy--reference--group-002.md#canonical-1321101233023113-2230303111330010-1322303301031233-3302120010112303-2213220032222020-0212033123332010-2201112211201312-2330013331202030): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-002.md#canonical-3000332003321310-2222200123021310-0022123022030213-2300033332121010-1232020131322111-0023111020201203-1013103313120120-1213331301010331): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-002.md#canonical-2302100211010213-1221110132232310-2212203011100023-1232220102103322-3120330203030100-3002113110022323-2221303300000213-0230130110100330): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203): complete subsection reference.

<a id="canonical-1230232323110101-1102031022210022-1021312321323213-0300001323213303-2112213330203121-2302321323030303-0220131203011102-1121012011203323"></a>

<a id="canonical-0122332221101300-1113211232110210-2321032311023311-3332113103300003-1133110220330100-0110112131330220-2011303003033311-3303112220332103"></a>

## value property — response_cookies_to_add / 003130103022 / 10

Type: `"string"`. Computed.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

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

<a id="canonical-3322333223130311-2210011220112003-1200221013131132-2101210303303102-3320001123333010-3230010223133202-0213303332230201-0013030332133202"></a>

## Next pages — response_cookies_to_add / 003130103022 / 11

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-002.md#canonical-1330320023303103-2302022130231131-2333232331213210-0002003331120312-0201333020223020-1232133321122300-1122113013020030-1220003331100012)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-002.md#canonical-3001022031232111-0122001320123233-1310332333101030-1320233302212010-3012310113033010-0122210321033202-0312311303112102-1112021130023003)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-002.md#canonical-0011102001032030-2331301112031120-1300022220021201-3320121023213032-2302001023033003-1322213002221113-1223023120131120-1012201320030211)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-002.md#canonical-2000320130332032-1233020232300330-2002133203003311-3312200230233212-2122220303123211-3320330322222230-1300010131133120-3210223023123231)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-1210112000100300-0332300323233332-2212132223111012-0310310211302021-0110320131032113-1220113113132301-2310132120112302-3033222131021333)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-3232020031003300-0031310323013133-0022131022233013-1231210333020221-3323231302203112-1133002233213031-2022111003222120-0013203333303132)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-3203111331101100-1133310211331120-1103032202331311-1213221210300333-3131001031221321-2200003330023222-0033023003200020-1120202002030101)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-0233030130130312-3322213313223010-2321000120031233-1003130313222231-1020300333211303-3031311032310112-3212202220120120-3221011311100230)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-002.md#canonical-1113133120202223-1022303120003311-3112222232223321-3232211120002030-0103200122122123-3022000201310203-0110121311103213-1003231022022322)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-2103030113132301-0301222332011111-1121233230131310-2123221130233021-0332220000130213-1221012000111321-2300122220133232-2233223111131321)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-002.md#canonical-3233220130010331-2032012212000020-2313232310022133-3133230003311210-0000310201313221-2220112312012011-0012121021113322-2312030230120230)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-002.md#canonical-0022102133201032-3321113032333101-3331132223220132-1302203111101301-3032120221020231-3000321310112230-3220320323312202-2200033122032220)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-002.md#canonical-1321101233023113-2230303111330010-1322303301031233-3302120010112303-2213220032222020-0212033123332010-2201112211201312-2330013331202030)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-002.md#canonical-3000332003321310-2222200123021310-0022123022030213-2300033332121010-1232020131322111-0023111020201203-1013103313120120-1213331301010331)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-002.md#canonical-2302100211010213-1221110132232310-2212203011100023-1232220102103322-3120330203030100-3002113110022323-2221303300000213-0230130110100330)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1330320023303103-2302022130231131-2333232331213210-0002003331120312-0201333020223020-1232133321122300-1122113013020030-1220003331100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002113012232333-2133031333101333-0123233122203311-2121032100233111-2331233003022113-3002013103011333-3110200012132131-2210323112012202"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly — add_httponly / 303312012301 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-0310122130232112-2112303020323010-3031023021201200-1130110131122123-2201132012111011-3113013103333323-0001203003201012-3021230300123112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3301033221003222-1012302201033013-0301131112001203-0311121230231310-2033020130023113-3321021333131031-2321301201323031-0321212311231231"></a>

## Direct properties — add_httponly / 303312012301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201303302110231-1202332032031100-2112231031200030-0223232103223021-0032213133131020-2033312002012000-2321110232111200-0330013103231333"></a>

## Next pages — add_httponly / 303312012301 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3001022031232111-0122001320123233-1310332333101030-1320233302212010-3012310113033010-0122210321033202-0312311303112102-1112021130023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133110113000130-2121122313130201-0310212312312300-2323312323023122-0232323220123011-3012313002111201-0132011020232223-0320110222022232"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned — add_partitioned / 132313323220 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-0110302311211213-3001303232113032-0003032011123132-0133302231331331-1310310101301212-0000213221332023-1032033122213310-2313230122302223"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1320202020331203-1222002332123300-2133100112323300-3300312333313213-3332012233003001-2212130300120100-2300122201212113-2310122311300030"></a>

## Direct properties — add_partitioned / 132313323220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131013203220102-2010233200133121-2100020001010323-1223030200211100-2203213002303131-1230312323330113-2000030310122100-1131330130211020"></a>

## Next pages — add_partitioned / 132313323220 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0011102001032030-2331301112031120-1300022220021201-3320121023213032-2302001023033003-1322213002221113-1223023120131120-1012201320030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101101121030302-3221021011210203-2110023020103023-1222123331110311-2321131220311200-3330031311033032-0023200220210023-3012003313131021"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure — add_secure / 311312331000 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-0220131211131002-0332312121332013-1330233133101112-3220330332333000-3313302101130201-2300300120112221-3100012203003200-2331120132030221"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1101203332122033-2311011201002121-1033200132011013-1232312032132320-1102221212213010-2030003331011113-3020201111310231-3012001311220220"></a>

## Direct properties — add_secure / 311312331000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122020230323300-3022102023023002-1102200210102011-0003002012300110-1001032010122012-3032002021233302-1223332303012232-1321221131310021"></a>

## Next pages — add_secure / 311312331000 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2000320130332032-1233020232300330-2002133203003311-3312200230233212-2122220303123211-3320330322222230-1300010131133120-3210223023123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013032303231010-0122203210330221-0031113332310200-3210233213302031-3122212312103031-0313003301021221-3212131221101123-3123201013110003"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain — ignore_domain / 212112220332 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-2101021113223202-0123220102122120-0123221031303212-1312220102021220-1202223020210220-0002020000313002-3131103022312331-2012233001131102"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0113003301131021-1110220311230100-2332222301130112-1000233212003213-2223000311033203-0301032103131320-0120020031311302-0320320230111320"></a>

## Direct properties — ignore_domain / 212112220332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013001113231221-2001211301102001-1113031120202333-3332223133203330-1021011210033030-0302033223011231-2131012132132310-3303003110132310"></a>

## Next pages — ignore_domain / 212112220332 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1210112000100300-0332300323233332-2212132223111012-0310310211302021-0110320131032113-1220113113132301-2310132120112302-3033222131021333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223001031002330-2003331112013203-0233200303033013-1200021021031303-3020122331312213-2101310203303310-2320110201321020-2112131020113112"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 212201003012 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-0032021002112131-0003113112020131-3120301302320101-3213313323300300-2112223001122211-0221323123012330-0011020113213020-1120030031203032"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0311132330203333-0220120212101321-3202122002202310-0312231000132133-0311110132300310-3231032003303210-1002303333330201-1131020203220211"></a>

## Direct properties — ignore_expiry / 212201003012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002110212101032-1133013223101133-2323230201000121-3030112312223200-1101331331133221-1213230213103103-2220303303301220-2222031123002133"></a>

## Next pages — ignore_expiry / 212201003012 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3232020031003300-0031310323013133-0022131022233013-1231210333020221-3323231302203112-1133002233213031-2022111003222120-0013203333303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101310232213213-2202333320133331-2210230131313020-1132021232320212-1132102110202203-0100223011330332-0221013112210100-1233001313201321"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 210023332203 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-2313311020230312-2220012210312121-3313322021232321-3200112231012013-1301123033111102-3232203120130100-2332301213231102-0301232300021000"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2113021123112311-1120212111012003-1003302102233000-3203221311000111-3300131211030333-3110012232003111-1133111231310211-0112033001020331"></a>

## Direct properties — ignore_httponly / 210023332203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113332203330013-1122210213131021-1012020331021230-2022212320220122-2102331133200221-0321321310310030-0030201030130131-3031203333000333"></a>

## Next pages — ignore_httponly / 210023332203 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3203111331101100-1133310211331120-1103032202331311-1213221210300333-3131001031221321-2200003330023222-0033023003200020-1120202002030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002232011100123-2230110030322120-1320030102020330-0332301100101232-0111013313021222-2130011300333231-1003002131011200-1230313311120131"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 003123230113 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1321012033301332-3302212312232303-0122211230023002-0121111130121222-2101103012213111-1110031012203111-3213200211133000-3120021013232121"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3102233331003123-1103000010002133-1311301210311320-3202322130121122-1222303023200232-3102321323212313-1320122020321123-0111221002220000"></a>

## Direct properties — ignore_max_age / 003123230113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113310332003233-2231213203102003-1311023132222103-0222331321303010-2232213130103003-1002133100203213-3313131002310020-2100100021223032"></a>

## Next pages — ignore_max_age / 003123230113 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0233030130130312-3322213313223010-2321000120031233-1003130313222231-1020300333211303-3031311032310112-3212202220120120-3221011311100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000103121323031-1110031002311033-2032013313201003-1102101132222000-2122000012121211-3213212323220202-2011002002301203-3123021012222221"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 333032333100 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2230013130210020-0003131033111010-2310230123223133-0231213112133200-2022021213120021-1330003213230231-2122120301321300-0012032221232201"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2312121201321101-0220002000231231-1111200200310210-2223020213333000-2002330012001001-1212123300332303-2133220232121311-1302302111321122"></a>

## Direct properties — ignore_partitioned / 333032333100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022021313213303-0323133032033120-1013110013313312-0023130313322232-2131013010122031-1000003111221211-3021210121110100-0000303213102030"></a>

## Next pages — ignore_partitioned / 333032333100 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1113133120202223-1022303120003311-3112222232223321-3232211120002030-0103200122122123-3022000201310203-0110121311103213-1003231022022322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212221112130230-3021120103300200-1310020320332210-3013322201302013-0123003110323110-2131201023333212-1020122100001120-2230302022112001"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path — ignore_path / 032000330013 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-0201110230001312-3212013101310213-2320110011202212-1101100113212332-2211121302011023-1021120111000321-3212010100200332-2031301200300102"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1013331012211221-2032302233222120-3203000332330002-3033322321322013-3002000300232031-1132111002032213-3220321112021332-2113202133331303"></a>

## Direct properties — ignore_path / 032000330013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203021011222230-0120012201131023-3111123032031033-3300003212232110-2020321223211022-1300003131230130-2212203330010232-0000110232321301"></a>

## Next pages — ignore_path / 032000330013 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2103030113132301-0301222332011111-1121233230131310-2123221130233021-0332220000130213-1221012000111321-2300122220133232-2233223111131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300111000020031-3011102022113002-1001221213323102-3313311111231322-3203221203300232-3103102312203321-1010323130112122-1111130133010331"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 001201230020 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-2330110303223023-3313302133211022-0233030220312310-1103112330203223-1221312131230132-1221100210112131-1003123023100332-0312111312203230"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2313211102232220-3220120222200210-1122233230222223-2220231321022022-3011032011023113-2013231132232303-2030222321112330-1022100030200131"></a>

## Direct properties — ignore_samesite / 001201230020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203010020112221-3001111313321220-2222000023002010-0232133033332231-2013201001200132-0210331221030203-0021303023033021-3203300330202002"></a>

## Next pages — ignore_samesite / 001201230020 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3233220130010331-2032012212000020-2313232310022133-3133230003311210-0000310201313221-2220112312012011-0012121021113322-2312030230120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223232121213003-0112002103030211-2000013021023023-2332201002120121-0123333211330100-2032221302021212-1121202300112233-0031003221021211"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure — ignore_secure / 322031211030 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-3303201032311210-2020021000200200-2220232002301232-0211113322233021-0230201133112211-1010131022001332-2132330011231011-1010101312013003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2110330102311221-3202311103201222-3102301122302011-1000130131011330-2320110302323311-2330103110201013-0312020002320031-3310300321313030"></a>

## Direct properties — ignore_secure / 322031211030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231031020233201-3100310020311311-1310230112313101-2103201220300231-0103331332303132-3022203212000313-1000232220210030-3033133013022223"></a>

## Next pages — ignore_secure / 322031211030 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0022102133201032-3321113032333101-3331132223220132-1302203111101301-3032120221020231-3000321310112230-3220320323312202-2200033122032220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121320323133202-1332311330102210-1113201113332130-0212132100320000-1332131230013313-1311223023221323-0312113001322301-0202011200230221"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value — ignore_value / 321302110223 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-2301101020300023-0111213200000021-1323000000033301-3310331102012233-0030213212133212-1232010331201013-3131013233000003-2131332310003122"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2212221230203023-1313030021133320-0213330301313000-3100012332123212-2020303232210202-0001213111001211-1000303132312133-0030100131130300"></a>

## Direct properties — ignore_value / 321302110223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230213313223212-1323223033023301-1202033300101010-3120131210031222-3023101322220200-3320102020011333-1212303221130303-0033023211020332"></a>

## Next pages — ignore_value / 321302110223 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1321101233023113-2230303111330010-1322303301031233-3302120010112303-2213220032222020-0212033123332010-2201112211201312-2330013331202030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000033022021010-0032030200231123-3102033331021313-2112210100003002-2030100223032211-3220031211133111-3002321111132120-1233201220323133"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax — samesite_lax / 131203303113 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2202132320013103-3010012212231032-2101223300130022-1001213131103112-2231202333201030-3121003213311320-2320300021323001-0303013133230313"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0031201133232223-0110222111123310-0331122021131331-0202111222122222-3322123010022003-2130311301020331-3321110103000312-0231230030310331"></a>

## Direct properties — samesite_lax / 131203303113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023330110123000-3213332013300030-2121123101212310-0123101000221000-2110121220323233-0332313001312330-2302333000310031-2203131233300210"></a>

## Next pages — samesite_lax / 131203303113 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3000332003321310-2222200123021310-0022123022030213-2300033332121010-1232020131322111-0023111020201203-1013103313120120-1213331301010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313112103211203-0021312303022120-1302231020200220-0331330100321322-3001123131030221-1201110231223111-0101130312133123-2132030120032023"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none — samesite_none / 120103020033 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-1023311230313012-0120313221202230-3332212131221231-0112320210311300-3300003001002203-0023023123010100-3001311022130311-1301220013212001"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2303020121101333-3220201313121003-3123303303002032-0212323321012110-2202313301121030-2012030330300032-1101111031103231-0221133313013323"></a>

## Direct properties — samesite_none / 120103020033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321123233123323-0203012031001303-3310120123332111-0332323031000003-2010222301032113-3033230003023100-0303100220032113-2221002301310121"></a>

## Next pages — samesite_none / 120103020033 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2302100211010213-1221110132232310-2212203011100023-1232220102103322-3120330203030100-3002113110022323-2221303300000213-0230130110100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330202122033012-3301210113121323-2112312303332313-1313002020330223-0201311022122111-3003330220123110-3203312132210202-2213230320330231"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict — samesite_strict / 210120213012 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-2322131320020103-2300130103020322-3113133131332023-0031120230120110-3223233320103111-3231103232021121-3203332032130100-2133221212003303"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0010202211030310-0032031032020120-0200123102133330-1331333003132223-0112103102103311-0232322311320030-2101003213330111-1222332101023020"></a>

## Direct properties — samesite_strict / 210120213012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030002111231220-2030312231021203-1112312223032321-3223110113201120-3223332300113113-0223130023110332-1011311202313003-2232103001030301"></a>

## Next pages — samesite_strict / 210120213012 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020123233023101-2230130011033221-3033131313202111-0310303102022031-0323203012133232-1322333020322201-2222332222012010-2333233300111032"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value — secret_value / 023131030011 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-3012002322003201-0231300023000001-0302103313033300-3213113012213023-1223311310222202-1201302023030212-2202102002320333-2111121210023111"></a>

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

<a id="canonical-3013032331302212-3202230111012321-0210011011103000-3000121321130210-1033201311122201-0302322112130031-3033023222003223-2123220222203110"></a>

## Direct properties — secret_value / 023131030011 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3200001130303132-3013213021102132-1123000331211133-0313021202000333-1213101310333201-1221333111130233-2121323122111112-0100333023201101): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0103203201123021-2323131121130323-1110013233332102-2112103023202201-1232030123210122-3301032300123331-2100001310120210-3330011230230312): complete subsection reference.

<a id="canonical-2232201133110123-1000103110220310-1101033301201332-2230201220330223-2122313213110230-3303122033012203-0321311121010011-1233331330113302"></a>

## Next pages — secret_value / 023131030011 / 4

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3200001130303132-3013213021102132-1123000331211133-0313021202000333-1213101310333201-1221333111130233-2121323122111112-0100333023201101)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0103203201123021-2323131121130323-1110013233332102-2112103023202201-1232030123210122-3301032300123331-2100001310120210-3330011230230312)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3200001130303132-3013213021102132-1123000331211133-0313021202000333-1213101310333201-1221333111130233-2121323122111112-0100333023201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303032113112131-1110022103132000-0010200303231023-2012020221013312-2320302230021300-2232121331303033-3031000310310331-0001200020123332"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 132022023203 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0100220301033201-3112201120003031-3212023010131120-1011112212222101-2010320023022232-2223030002300112-3131303121213203-3031131200011312"></a>

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

<a id="canonical-0321012103223010-2000021310101223-0301113103123031-1003133102131330-1010001011000010-1010102223111122-0101202110221323-0120333212113221"></a>

## Direct properties — blindfold_secret_info / 132022023203 / 3

<a id="canonical-0321222212300231-2212221100131023-1233200021032232-0102231033011212-0113122232020030-2320020003123132-3022200320233231-0011312231213003"></a>

<a id="canonical-1200210003023220-1232002233221331-1303102103110202-2033322131311201-1113310302010300-1030222022310101-2301233103331230-3032133201101121"></a>

## decryption_provider property — blindfold_secret_info / 132022023203 / 4

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

<a id="canonical-1000033231102332-3331101220100303-2231220202321201-2312310233111211-0002023333111332-1133101321323022-3312313133310231-1201112101020100"></a>

<a id="canonical-2310231113212333-3112100020203132-1000130113121132-0023023030220112-2112200001031012-0301232020033222-1133302232232120-2131012030301230"></a>

## location property — blindfold_secret_info / 132022023203 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-3222223232013300-1011122202303010-1113102100300303-2010201110333222-1223221013011323-3220023221321200-2123131310033301-0302233330000111"></a>

<a id="canonical-3313102333303313-2212230021120020-2311212130303133-2010103001013313-2110203302131323-0021102230332333-2320333100130200-1233200133313210"></a>

## store_provider property — blindfold_secret_info / 132022023203 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1111001233333012-2210132220110130-1202331111210132-3011203032211221-0223102213003202-0211101213110223-3203021131203310-0002221311110132"></a>

## Next pages — blindfold_secret_info / 132022023203 / 7

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0103203201123021-2323131121130323-1110013233332102-2112103023202201-1232030123210122-3301032300123331-2100001310120210-3330011230230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221111023121311-3000021110301222-1213122100011310-0213012202120103-1110221232013101-3023312130213122-3233230333130011-1120231311301111"></a>

## dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 221030323222 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203)
- dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3211102222212302-0033302310130320-1031323022131132-0121131100222101-1222031021211330-1312201131101231-2211000200302110-1221000122110021"></a>

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

<a id="canonical-2212203133233031-3201320230333111-0230203133322333-3000213133233203-2320003222110222-2210202131323301-2032203010133032-2220100033032303"></a>

## Direct properties — clear_secret_info / 221030323222 / 3

<a id="canonical-3303201130223220-0203303223302211-3001021132200222-1222203131300213-1222111000010232-1103003211020300-2132223000223002-2121332230303121"></a>

<a id="canonical-2201332201131232-3312132113332033-0320331133201230-2322203231310230-0313201320321230-3201012113031323-2033313321321331-1122210112203012"></a>

## provider_ref property — clear_secret_info / 221030323222 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0331003312223320-1111131030123311-0200321322003221-1321210120132220-2131022001331130-3003000113011312-2211001203111333-3223133323310012"></a>

<a id="canonical-3301122311011331-2023132222221312-0222101113230101-1223002211003223-0222002321233332-2322102230033011-2021202132120202-2211133012313320"></a>

## URL property — clear_secret_info / 221030323222 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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

<a id="canonical-3103310022030112-3212033202122112-1132332313130232-0322320232103200-3132033330031012-1223131311110320-2332112123333120-1321013132112312"></a>

## Next pages — clear_secret_info / 221030323222 / 6

- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220102322121130-1123232112111100-3201112232302010-1131212032033000-1331212233103012-0022101302021311-3111210220302000-1132312033200230"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add — response_headers_to_add / 103210221030 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add

<a id="canonical-3312212330130212-1201120222102111-0201011303032130-2220131322201202-1311213021322313-1302101033302113-3111130332210110-3103131033031213"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

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

<a id="canonical-2331002131113232-0022020012132333-2012202121113323-1300000201312032-1231130211212201-0312132333013023-1202330201110132-0332130123331003"></a>

## Direct properties — response_headers_to_add / 103210221030 / 3

<a id="canonical-1323013212323222-0133321011100133-2333213202320102-1223333010032010-2223003322201213-3032033112221002-2303013121220003-2300013210310010"></a>

<a id="canonical-1200133121312032-0000102231022222-2120130013011130-2302232122330323-0133312023033132-3222111102120220-3000020110132230-0310101321023310"></a>

## append property — response_headers_to_add / 103210221030 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-1012103112123222-1301212121322131-2311113300033002-1002302330012130-3233130120211102-2020022212321211-1302220202233010-3010312200232002"></a>

<a id="canonical-1320310123321230-2100232322103030-2300333200320001-3033332210030000-0230012021210010-3001303222213233-3100131020000211-1110120100210223"></a>

## name property — response_headers_to_add / 103210221030 / 5

Type: `"string"`. Computed.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123): complete subsection reference.

<a id="canonical-1230313323202122-2020220200321131-3322123220303122-1232202302011313-2301030330311020-3233312120032133-0111213333132301-1022121123202203"></a>

<a id="canonical-0311331313102213-3111220022230330-1312201113222120-0112121111323003-3120010120220300-1022022120113200-0303021012022213-1133133320001003"></a>

## value property — response_headers_to_add / 103210221030 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

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

<a id="canonical-1111100200003323-3321301222033130-2220020011220200-0300222030330212-2200010033303001-2100003113111200-2012121000322123-1200212221120030"></a>

## Next pages — response_headers_to_add / 103210221030 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200203133033313-1020002223001101-3001032313013322-0003002222332023-2322313320132312-1121220033013210-2200302113001130-0322121202131032"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value — secret_value / 221212233101 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-0202312230102310-0202111122000122-2000221003011320-3123021103110030-2213030221113132-1013111103002002-1320332303203212-3110312302301201"></a>

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

<a id="canonical-2000203330201223-1311332131300313-2033130103232211-2331302212220002-1212130032213202-0231012130132302-3221110013332233-0113130320022330"></a>

## Direct properties — secret_value / 221212233101 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3131013011323133-2030001313001023-2033103301222120-0001023000020223-3322030121211210-2223310202330202-3112213313113220-3110303030312102): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-3003011133003003-0230110223002131-0000130130221012-0320203010112020-2222332312333231-2013021021033332-3322130120200021-2101013222223203): complete subsection reference.

<a id="canonical-0110202310203310-1012301320331222-1111312003011010-3301102113331230-2201310111211120-3001011310300221-3301010010210123-3312231203212330"></a>

## Next pages — secret_value / 221212233101 / 4

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3131013011323133-2030001313001023-2033103301222120-0001023000020223-3322030121211210-2223310202330202-3112213313113220-3110303030312102)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-3003011133003003-0230110223002131-0000130130221012-0320203010112020-2222332312333231-2013021021033332-3322130120200021-2101013222223203)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3131013011323133-2030001313001023-2033103301222120-0001023000020223-3322030121211210-2223310202330202-3112213313113220-3110303030312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002002011211031-3313131311010221-0022223322121302-0110020201012333-1233010111130002-0102303230220222-3232232011320222-3020001020311100"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 031222331010 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1121220313233030-1120102232203323-1332201110330110-2032322310121320-1230120122031222-2301011113313200-0212100131332320-3322222011332332"></a>

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

<a id="canonical-1223110221002102-1200333220212330-3022011302302322-2323333310121332-0232310233032300-1220222000201121-3120200303313233-2021133123130202"></a>

## Direct properties — blindfold_secret_info / 031222331010 / 3

<a id="canonical-2203122130223323-3003210121230023-1021220133033131-3320003312111331-3330030010033020-0113112013231021-3232011211001010-3133000233233300"></a>

<a id="canonical-0000100133130013-3333032303303331-1322313111202202-2323231310110332-3301232211332012-2211003023012232-3021020303322133-0301302222222311"></a>

## decryption_provider property — blindfold_secret_info / 031222331010 / 4

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

<a id="canonical-0101123010030322-0033102211302131-2000200321033131-1322131201111022-3032102003220133-3011213101231032-2002232201130011-3212312321220301"></a>

<a id="canonical-0201220211033122-0231121322201132-1101313102230003-1233000233103310-3332002010212221-1210120233121003-2310330110001302-2331301333033132"></a>

## location property — blindfold_secret_info / 031222331010 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-3133303311122322-1310223103032302-2013010233303330-1312322111100110-3223200101223111-1020300112103021-0210103310023021-0121221010030221"></a>

<a id="canonical-3130201221201333-3221230111033113-2203101023333011-1312122202202112-3130330021201133-3312120023302332-1122122002301131-0320120122321322"></a>

## store_provider property — blindfold_secret_info / 031222331010 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1030332112121031-0102230113213330-2231112111011130-0210022330232032-0222033202212212-2331323020003030-0102010301233010-1322130023133021"></a>

## Next pages — blindfold_secret_info / 031222331010 / 7

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3003011133003003-0230110223002131-0000130130221012-0320203010112020-2222332312333231-2013021021033332-3322130120200021-2101013222223203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210021123312020-1131202230031220-0010302112333213-3103001333220320-1120211223212203-0221132323111210-3100301202113003-2321122110023132"></a>

## dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 000232133122 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123)
- dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2321312023321100-1110303211202212-1330033333322010-2022310312130231-0322110112100033-1300203012310301-2131310320122331-3230100003121011"></a>

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

<a id="canonical-1120121311213220-1331022100011200-3020030100101301-0111113220212323-2011131233022333-3002120233232320-3030011030323323-3200131002031222"></a>

## Direct properties — clear_secret_info / 000232133122 / 3

<a id="canonical-3102201322001013-3203201013031202-3003333023302012-3131310322031333-2103002121302113-3322232313022210-0300102131110032-1133032300032321"></a>

<a id="canonical-3011102112032221-3120121312213112-0131221121210332-2110312332300122-1101231310012111-2002111210332320-1022312220133110-1222132021303222"></a>

## provider_ref property — clear_secret_info / 000232133122 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202110312310200-1031310233010010-3110311022210110-2133120010113003-0132021131000010-2320003022322121-1013121223031113-0120221031331221"></a>

<a id="canonical-3032220131013001-2313230023301003-0312100031012011-3203100312110302-0223232203221132-1122102203212101-3123030002130032-1212001000002301"></a>

## URL property — clear_secret_info / 000232133122 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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

<a id="canonical-3231120022123131-3013113130333303-1112010011223121-1311330132002221-0221222020331113-3320003023002132-3301010202013111-0330303013333230"></a>

## Next pages — clear_secret_info / 000232133122 / 6

- [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211100010310322-3013313121213021-1320103312202010-0322202323222300-1112323221122012-2310233300011230-1100113031301132-2220003201013202"></a>

## dynamic_proxy.https_proxy — https_proxy / 031222133103 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.https_proxy

<a id="canonical-1222311003032122-1302001312111013-2110112322002123-3320333003001322-0310220221331213-3032032120010010-1010200111113120-0300001233302320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1320301321021133-1011333312032300-0102232211232012-1032210020033000-3123223221231100-0222102010133303-0121200113203120-0133030123022202"></a>

## Direct properties — https_proxy / 031222133103 / 3

- [more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121): complete subsection reference.

- [tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133): complete subsection reference.

<a id="canonical-1320311220022101-3320321222212301-3130332012132323-1233013213231020-3212233120210003-3120013110011213-2222122303212111-0122332103213101"></a>

## Next pages — https_proxy / 031222133103 / 4

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203023111122002-0233123333310313-2203312103212201-2012000313122011-3031113100013101-2121113231031210-2320100002210033-1102120323121133"></a>

## dynamic_proxy.https_proxy.more_option — more_option / 120022231013 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- dynamic_proxy.https_proxy.more_option

<a id="canonical-0330330032102302-0232000313232301-3130033233130020-3331220030003003-0300111231313010-0220033130201110-3321033000213100-3331323221103331"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

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

<a id="canonical-1301312121203032-0121110022103212-0010121221132232-0113312020000001-0131121220002113-1210131231020001-1020130303110223-1113122130233320"></a>

## Direct properties — more_option / 120022231013 / 3

- [buffer_policy](data-sources--proxy--reference--group-002.md#canonical-1303102010212200-2123321202223330-2000222103103030-2333113231120301-2311212220101110-3310112001323202-1011333131023301-0003122130300303): complete subsection reference.

- [compression_params](data-sources--proxy--reference--group-002.md#canonical-1300300133233300-1223221223133013-1022301301030100-3330231313221320-0113302330220102-2031110231311202-2122210033231312-2221213333100110): complete subsection reference.

<a id="canonical-1302320102022003-1333333000110020-1012331230132103-1222010333210330-2202321101220122-3232101221031000-0102132312131302-2111301003330301"></a>

<a id="canonical-2030233311221333-3033022301111301-1131020030133202-1122310212221323-0023013132023313-1332002321030022-1230120122320301-3200300222102222"></a>

## custom_errors property — more_option / 120022231013 / 4

Type: `["map", "string"]`. Computed.

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

<a id="canonical-0223222221302010-0121102012211230-0121112100120322-0123230212103030-0002101113013302-0030330321232012-3001300102033103-1133102033002212"></a>

<a id="canonical-0330111222120001-1121011220011021-3332102321000210-3120101230211103-3310032110222220-3100223020302300-2202333200312133-1322103301331221"></a>

## disable_default_error_pages property — more_option / 120022231013 / 5

Type: `"bool"`. Computed.

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

- [disable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-2021330300201200-1321031132112312-3210210020302020-1201320002200120-0222123032233313-0030333330021332-0301022221032211-1131210120113230): complete subsection reference.

- [enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-2223113323111313-1020333300201031-0322032313321010-1012200231231223-0123200303221002-3233003213023203-3303020003213301-3011033222002303): complete subsection reference.

<a id="canonical-0312330002201332-3100003203032203-0101221121320130-1221121131122020-2012003110103213-3312121132100001-0321022011323020-0223310333331300"></a>

<a id="canonical-3113123030012111-3232013030322100-1231133232031021-1211202020010013-1321020230022000-2320211232002223-0121021310112200-0313222200232230"></a>

## idle_timeout property — more_option / 120022231013 / 6

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

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

<a id="canonical-0002213032113133-1200220100003330-0231031223001001-0330113303132221-0202112113332311-1112120120223211-1213301131311022-2111121131311332"></a>

<a id="canonical-1002201033111103-1322011332311200-3030303213030112-0223031123312312-1201211111011212-1221010120310301-3321023030112113-0020213203221231"></a>

## max_request_header_size property — more_option / 120022231013 / 7

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

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

<a id="canonical-2332110313100232-1131222210031023-0303012112310203-0323102202300320-3303200030222001-1001023231031120-3011000223203020-1032200300223022"></a>

<a id="canonical-3103301222021123-1311002313233330-1221222130111000-1212312111101001-0033322032303130-3223123022001103-3200320003331321-1110131123313002"></a>

## max_requests_per_connection property — more_option / 120022231013 / 8

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-2221323221232001-2003200310122001-0121102203111132-2312131321132231-1022000002200020-3333232001203113-3320123302202300-0300313200012032): complete subsection reference.

- [request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210): complete subsection reference.

<a id="canonical-0013021232203110-0121011220101202-0003133232002101-1303313031223013-0032001011132200-2033312112211033-2131120213333130-1011020010110310"></a>

<a id="canonical-2103211310323102-0322102333031230-1221100021231123-3032331000020211-1032120131032320-1333113311331011-0200221230223212-1100113223222032"></a>

## request_cookies_to_remove property — more_option / 120022231013 / 9

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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

- [request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230): complete subsection reference.

<a id="canonical-1212333223303331-3022333132132112-1203003331023323-3032033020311210-2331301132021322-3211312001213031-3303322313130032-0303311302102031"></a>

<a id="canonical-3111123323002311-2102021002011102-2303300310220230-3103331021200003-0232131221331200-1210230030321130-0332311133221312-0310132311230010"></a>

## request_headers_to_remove property — more_option / 120022231013 / 10

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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

- [response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332): complete subsection reference.

<a id="canonical-0313202333230232-0103223103102030-2202010300303120-1313132013333101-2332322122312001-1122100232002210-1311122231003100-3313300213211212"></a>

<a id="canonical-1301110100030200-3013121121302020-0032330101133120-2000130311321121-1130102020031010-0130002312323110-2202100110313332-2331213020112202"></a>

## response_cookies_to_remove property — more_option / 120022231013 / 11

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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

- [response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032): complete subsection reference.

<a id="canonical-2232102300210212-3233132000220001-2013323301012302-0321031023311030-2000212133332333-2103031122121122-1330211330231213-3230132013332311"></a>

<a id="canonical-3331301200313120-1231330320030203-2100320212012022-0101200220333021-2022302120322113-0203233132002212-0202301221033032-1230320112313122"></a>

## response_headers_to_remove property — more_option / 120022231013 / 12

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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

<a id="canonical-0123222031202213-3020130100011321-3121010112133310-2301133300102331-1200202303223110-0121100131123113-1211130323310100-1032332111330223"></a>

## Next pages — more_option / 120022231013 / 13

- [dynamic_proxy.https_proxy.more_option.buffer_policy](data-sources--proxy--reference--group-002.md#canonical-1303102010212200-2123321202223330-2000222103103030-2333113231120301-2311212220101110-3310112001323202-1011333131023301-0003122130300303)
- [dynamic_proxy.https_proxy.more_option.compression_params](data-sources--proxy--reference--group-002.md#canonical-1300300133233300-1223221223133013-1022301301030100-3330231313221320-0113302330220102-2031110231311202-2122210033231312-2221213333100110)
- [dynamic_proxy.https_proxy.more_option.disable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-2021330300201200-1321031132112312-3210210020302020-1201320002200120-0222123032233313-0030333330021332-0301022221032211-1131210120113230)
- [dynamic_proxy.https_proxy.more_option.enable_path_normalize](data-sources--proxy--reference--group-002.md#canonical-2223113323111313-1020333300201031-0322032313321010-1012200231231223-0123200303221002-3233003213023203-3303020003213301-3011033222002303)
- [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--reference--group-002.md#canonical-2221323221232001-2003200310122001-0121102203111132-2312131321132231-1022000002200020-3333232001203113-3320123302202300-0300313200012032)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1303102010212200-2123321202223330-2000222103103030-2333113231120301-2311212220101110-3310112001323202-1011333131023301-0003122130300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231213131333233-1313311113113002-0310021000300122-1131331121202113-0021001113001302-2033200110211112-1022323322321122-0310211323103322"></a>

## dynamic_proxy.https_proxy.more_option.buffer_policy — buffer_policy / 221011223221 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.buffer_policy

<a id="canonical-1221233103020020-0003012303203331-0212230222222101-2003133212022001-0121012210022112-0103213300230112-1011033033220303-1003103313210123"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0222131121220300-2013002302011301-3200122133112300-1320331212031220-1122210022133002-0302033322323033-2013222212222003-0101231202011220"></a>

## Direct properties — buffer_policy / 221011223221 / 3

<a id="canonical-1203121332032200-1223221101111131-1310112222102202-3133233200132133-2313110030313103-3011123213332230-2130123013213113-0010002233313120"></a>

<a id="canonical-1033221120222213-2113211301013222-3200032233123210-0001010322113200-2120101221010100-2003120231300022-3131302203031303-3003103223310322"></a>

## disabled property — buffer_policy / 221011223221 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-1021131103002220-0300022023301210-3200010011320302-3120120332300130-0211122032211102-0113003022022230-1001322131012303-2131230101131102"></a>

<a id="canonical-2210222333112103-0123113112023221-0031211100113022-2212210023112013-2233133323022201-0132023122010001-3211113133022103-3111000133303033"></a>

## max_request_bytes property — buffer_policy / 221011223221 / 5

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

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

<a id="canonical-1010301231010322-1333202001313133-0310221323032033-2112300300322323-1320200300202331-1022301113300202-3201133312201033-3303313202331232"></a>

## Next pages — buffer_policy / 221011223221 / 6

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1300300133233300-1223221223133013-1022301301030100-3330231313221320-0113302330220102-2031110231311202-2122210033231312-2221213333100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003222131223103-2113032033303120-1202312131303021-0302221323210113-0101313001013202-3300033002031030-1101322121212220-3200113323123023"></a>

## dynamic_proxy.https_proxy.more_option.compression_params — compression_params / 030310023123 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.compression_params

<a id="canonical-3131232203221303-1301131211000030-0000300231122333-3320212213012330-2103033113211120-1023230022101131-3010100320233032-1203112331212132"></a>

Type: `"single"`. Computed.

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101222333130201-0323031013231101-2013031123231113-1210203311021123-2130332031112202-3110132220213103-1331231312312212-1103222210301113"></a>

## Direct properties — compression_params / 030310023123 / 3

<a id="canonical-1311303312130301-1023002133002212-3010220210302301-2103223022001132-0022032231332033-3112010132021102-0031312001202101-1232331232003220"></a>

<a id="canonical-1332112112121113-0310310122023123-1020110120223311-1322322233203210-0002000133103120-3201020331002101-0313203001312303-2231133313110130"></a>

## content_length property — compression_params / 030310023123 / 4

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

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

<a id="canonical-0202323112211111-1310033312203211-1000031212200310-2313003310113303-0022301010333100-3013000101013112-0123330211231033-2331123222300032"></a>

<a id="canonical-2310230012100123-0310321330133210-1301102303100003-0023131232311110-0301033201020331-3323203021030103-1301301321010311-1013131203332121"></a>

## content_type property — compression_params / 030310023123 / 5

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

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

<a id="canonical-0203013101331223-3022223011232200-0101313112222101-1011230231132233-0312201121110103-3322130232202100-1301001333012223-1310233011221310"></a>

<a id="canonical-0323321100320312-3210120321120331-1223310123332011-3102123002332001-1113213310012210-1111022202132313-3210112002300003-3302003330200033"></a>

## disable_on_etag_header property — compression_params / 030310023123 / 6

Type: `"bool"`. Computed.

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

<a id="canonical-1010221101013213-2032320220222230-1102100011002203-3311200031013012-2230101231012133-1033130330112113-1202231112020231-1000111302331112"></a>

<a id="canonical-0321122232222013-2003321313002300-0203121112201221-0302131101010112-2000023003311100-3200222312323011-3311123010203032-3123103002022212"></a>

## remove_accept_encoding_header property — compression_params / 030310023123 / 7

Type: `"bool"`. Computed.

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

<a id="canonical-0302122103230113-3101122312113333-0123301101213030-2300210102101132-0302112113011222-0301102220201031-0110000330322220-1013113133102030"></a>

## Next pages — compression_params / 030310023123 / 8

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2021330300201200-1321031132112312-3210210020302020-1201320002200120-0222123032233313-0030333330021332-0301022221032211-1131210120113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210302323030230-1322303201323210-1121222033122122-3131213121210113-2300233321321002-3310023330001203-2120122331131113-2330210023211330"></a>

## dynamic_proxy.https_proxy.more_option.disable_path_normalize — disable_path_normalize / 132130120002 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.disable_path_normalize

<a id="canonical-3103332100222101-1112002010100030-1022130110103320-0300221131200130-0222300223101232-0223103123020310-2123021211020331-3213302213300322"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0203031132130211-1002233031003210-3121200000133203-2102020020233223-3010012021310220-1230121022100310-2210121103002102-1113011001203121"></a>

## Direct properties — disable_path_normalize / 132130120002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132323323131301-1313230323232032-1231311023303001-3212112131300110-3301210130123011-2132101111230201-0102301311321312-2111122221202310"></a>

## Next pages — disable_path_normalize / 132130120002 / 4

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2223113323111313-1020333300201031-0322032313321010-1012200231231223-0123200303221002-3233003213023203-3303020003213301-3011033222002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301000322120221-2312111111032303-1131032321013122-0231012000131202-0212110022300031-3330221220213120-0112033223031033-0311122020302123"></a>

## dynamic_proxy.https_proxy.more_option.enable_path_normalize — enable_path_normalize / 312101201021 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.enable_path_normalize

<a id="canonical-2303121322131302-0230000310301101-2320103130300030-0030232112012122-1302210012013113-2012231111301211-0121003020113310-3320323211012003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3221021133112013-3322010001202302-1010023131320010-1322210320333133-0122000131123023-2213013230132322-2310120020012012-0121110100221132"></a>

## Direct properties — enable_path_normalize / 312101201021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010323333220011-2132233233101313-1222333213102220-0212330230332032-2013010103233313-3310020010002022-3001310101132020-3101003002133323"></a>

## Next pages — enable_path_normalize / 312101201021 / 4

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2221323221232001-2003200310122001-0121102203111132-2312131321132231-1022000002200020-3333232001203113-3320123302202300-0300313200012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120201012332122-3033101132303111-1311220100020111-0032120213211111-3232233121230220-2223200232301232-0231020311312000-0100221111011211"></a>

## dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection — no_request_limit_per_connection / 212100122022 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection

<a id="canonical-0311321200323311-3132122110012022-3221232200013330-3110022100232103-3012002031333200-1312122310023310-3112313221121133-2333323311033011"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2012310213123102-2223231021120200-3033213121131032-0331001101102302-0002230303033200-3303011022123033-2031323103002220-2002311011300322"></a>

## Direct properties — no_request_limit_per_connection / 212100122022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011103101100313-2210021332121320-1002302321212232-2130110030020110-2203322112131021-2002210010123312-2231330203233212-3231231103121201"></a>

## Next pages — no_request_limit_per_connection / 212100122022 / 4

- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110323233012113-1122113021223100-1113321230232302-2132303233023011-0320333133131210-3130120132113132-0133000001321000-3330000203122010"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add — request_cookies_to_add / 013011302331 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add

<a id="canonical-1120003130332102-3121303130133111-0023133331300321-0002012231203222-1101000221120303-2323012301233003-1123031022032000-1310010103221012"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

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

<a id="canonical-1112132001010210-2021123303021011-3203321220331203-2103202122200113-2130301302132303-1311223203031231-3031103221130313-0212330201002202"></a>

## Direct properties — request_cookies_to_add / 013011302331 / 3

<a id="canonical-1100023031122013-1331003032000011-2011322302303301-3203312230121030-2233200223203200-0300320203032202-1222001122113000-2020031023003013"></a>

<a id="canonical-2032230023011302-1231131032313212-3132000303033032-3201303312132103-2323311133331131-2130200310100323-0221100221210313-3022021023213200"></a>

## name property — request_cookies_to_add / 013011302331 / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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

<a id="canonical-1130010121100103-3011233012203111-2110212110322131-0010312120010233-2321110312001212-3333001021001011-1020331130031320-1300002110101210"></a>

<a id="canonical-0210033232330122-3123110302131233-1200330103321303-2112102321200233-1223322031032333-0112002211000201-0020131211301103-3123212113103201"></a>

## overwrite property — request_cookies_to_add / 013011302331 / 5

Type: `"bool"`. Computed.

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003): complete subsection reference.

<a id="canonical-3021100312103311-2102033330200311-0013322122301123-0122110333211210-3323112320311102-3310332223011132-1000001312233221-0231020102213213"></a>

<a id="canonical-2212201133200000-3110303303323003-2333211310331031-0112021103032022-0130032032031233-3131012202211210-2300303000220221-1230322031131131"></a>

## value property — request_cookies_to_add / 013011302331 / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

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

<a id="canonical-1313123200311133-3131320012200333-2120113303300320-3312301000111330-1302302011023013-1213011333313321-2001321120131121-3103023101203033"></a>

## Next pages — request_cookies_to_add / 013011302331 / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320030221123132-3131133020211110-2012121100221123-0113221330313123-1101031133131220-2211310221231202-2323231211033012-2112010123313310"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value — secret_value / 133211302032 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0133113203220032-3033213123102000-1013000023302223-3103331001200322-0100311102203213-3321203303331111-2132103233303122-1210221212212101"></a>

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

<a id="canonical-0221302002330022-2313221121310232-0110231113132203-3233023103231020-2202212122221012-1312310012101022-2102323122101321-0130230331132301"></a>

## Direct properties — secret_value / 133211302032 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1301330002000203-2311011011030020-3202222132233330-0210220323030111-0201130301212230-0013202310131120-1102320300210211-2030312223223132): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0002213211333012-0301110220233033-1022321002103003-3230303322201203-1102022101031132-1230012130231310-0021333120201223-3131321021013010): complete subsection reference.

<a id="canonical-2022201020311220-2212212033331121-0112101321020120-0030220002110132-2230302201223302-0202123013011012-1210011022323023-1332000032233031"></a>

## Next pages — secret_value / 133211302032 / 4

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1301330002000203-2311011011030020-3202222132233330-0210220323030111-0201130301212230-0013202310131120-1102320300210211-2030312223223132)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0002213211333012-0301110220233033-1022321002103003-3230303322201203-1102022101031132-1230012130231310-0021333120201223-3131321021013010)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1301330002000203-2311011011030020-3202222132233330-0210220323030111-0201130301212230-0013202310131120-1102320300210211-2030312223223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132022301130302-0021322123120020-3300333132210211-1120222202230033-2221302131213010-3130202311131001-0130300122322022-1333201310131030"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 122123103230 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-1231001223112233-1321000010203011-1032331011123031-2031100232232332-1110311110000232-1313233231000023-1030103103223213-1032030231000311"></a>

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

<a id="canonical-2230000133333211-3311233113100013-0013321320112322-0211201302002313-0203130133303210-0100033123300020-3023201313120120-2101132000232002"></a>

## Direct properties — blindfold_secret_info / 122123103230 / 3

<a id="canonical-3103002003010201-3021132302230011-3300010130223332-2101111230202210-0111103311123100-0111311311020023-2103010130023331-1220221010320302"></a>

<a id="canonical-1220201332330311-2310002323311310-3223000031230100-2230300020222233-3132222201222323-2132123330313310-2301113121221332-3202311310002222"></a>

## decryption_provider property — blindfold_secret_info / 122123103230 / 4

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

<a id="canonical-3132202111311010-3311010320310033-1122322303312122-1203012001322312-0210313310102133-3130203121113012-0013131301013223-2022311320301310"></a>

<a id="canonical-0332113300132000-2132211012023001-3331211110333111-0203313333020012-3300021112012010-3321320121122102-1112213330211100-2000123330101323"></a>

## location property — blindfold_secret_info / 122123103230 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-3001003132301223-1232013310232003-0133100211221012-1131300331221312-1110222122002303-2011332001312321-3233001331213330-0323313322212331"></a>

<a id="canonical-1333131000101330-2032010110203123-1011313121111203-3233131212221311-0213300001320103-2332233233210333-0312320201013320-2122222223022333"></a>

## store_provider property — blindfold_secret_info / 122123103230 / 6

Type: `"string"`. Computed.

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

<a id="canonical-1113222312133133-0301120220222030-2001212322223222-3211002223332313-0203312010313032-1331333233332203-0312202322013122-3323102103312331"></a>

## Next pages — blindfold_secret_info / 122123103230 / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0002213211333012-0301110220233033-1022321002103003-3230303322201203-1102022101031132-1230012130231310-0021333120201223-3131321021013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
