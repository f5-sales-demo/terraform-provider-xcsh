---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1221320221131021-1110321300020103-1113100121002030-0112311320101002-2112021030302002-3132333301313100-0110211323032123-2032132212130111"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 323321222102 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0113020303211220-1332132312111122-1301012222223012-2211200003001330-0122331332012230-1223322331220311-1123011031012021-3200331220121021"></a>

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

<a id="canonical-0021012003331033-1003203220133132-0330112233003222-1321032322223120-3121012332213000-2210322331101101-3112000032320101-3002113231032030"></a>

## Direct properties — clear_secret_info / 323321222102 / 3

<a id="canonical-3111300220203102-3031202013232321-2312111032100332-3001100302231212-1101020300111320-0302322211132302-1203010011220012-3301030131300131"></a>

<a id="canonical-0322122200203200-1030203310012002-2333332322333222-2221032023132122-3301010110323312-1303123330223313-3010220032010030-2020313002211203"></a>

## provider_ref property — clear_secret_info / 323321222102 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0021121303003233-3200122100110013-1133113020133311-0001003012301202-1102113310312132-2313320320203010-0233100132023220-2331203032301301"></a>

<a id="canonical-1013131323321323-1322002212102232-0101310020110203-0320230313120110-3120213022010201-3121131301313102-3301222130002211-0300212021210030"></a>

## URL property — clear_secret_info / 323321222102 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0321222230022113-1312202203203123-2020011221133013-1000001330132121-2131332213312122-0213223223222212-3200202011310013-0000103023221120"></a>

## Next pages — clear_secret_info / 323321222102 / 6

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100030200122331-0121031300121213-2210001323032010-1233103230312212-1323013332030131-3213312332331111-2113001232223011-3112333013111000"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add — request_headers_to_add / 122022120212 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add

<a id="canonical-3312031003003121-2202113112200322-0311030122313202-3130310302231112-0031002230332131-1120122303100123-1331331323031201-2212133220310223"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0231123121010212-1031200332231302-1231220301123321-3212332223201331-1111132022200131-2201112323320233-3020011012011001-3333322201233310"></a>

## Direct properties — request_headers_to_add / 122022120212 / 3

<a id="canonical-2212000203223010-2011131120233121-3202112123223212-1202111112133130-2231013131032122-0310131023332023-3010222311220312-2102212020231131"></a>

<a id="canonical-3303133111101000-0232121123332233-2222012022310031-2030132203223333-2021312110103230-3313311301200130-0021322011032122-2110300212333103"></a>

## append property — request_headers_to_add / 122022120212 / 4

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

<a id="canonical-2210312100202211-2013002111300220-0020103200032232-2200022311110202-1210002020220311-1012010313120200-2002231012021120-3013201032032113"></a>

<a id="canonical-0020323123202022-3230130012121311-2003232133101212-3320300112230301-3023222213233121-2213220110013232-2030232320103030-3210020002102332"></a>

## name property — request_headers_to_add / 122022120212 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223): complete subsection reference.

<a id="canonical-3101202010010102-2013122031202021-0320233111012113-2202211220032300-3021322310300301-0213202323222031-1011311031121332-1321233313221013"></a>

<a id="canonical-1102012323020111-3130200123013303-3203012332331100-3212231310222020-3023020100101123-1202130212021311-3020133012333323-1302233223023012"></a>

## value property — request_headers_to_add / 122022120212 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0203012333030023-3313202322131310-3121313200321021-2213122302121322-0011221003012003-2233231301110130-2111121023031123-2311332120010320"></a>

## Next pages — request_headers_to_add / 122022120212 / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312203131033322-2322313302303311-1102003021012221-1033022022212113-3102223223310322-3122032003222133-1110002002332220-3302103113031031"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value — secret_value / 301302112223 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-1201101302100032-2312131313221330-0211230131000012-2121132320121313-2133303012002113-0222332313332133-1212332302103120-1002132021133021"></a>

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

<a id="canonical-2212113101230322-1100212231131320-2123133130332003-3301102330001032-2002001322033123-3003023302013022-3020333022323330-3200113221123230"></a>

## Direct properties — secret_value / 301302112223 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2031112112320023-2113300013121202-2023001332113201-1002000123330210-0121032211312110-3032133300101303-3300000213021300-2312330303231233): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-3112222122223320-0030311322322011-2101000223213320-3210330331031200-3311323031220032-3302233113130331-3312120333001332-3213031331112311): complete subsection reference.

<a id="canonical-0323210321333003-2032010330221333-2033102322012132-1111302110320223-1312102212011302-3213303311233131-2211313022230203-3331101103230102"></a>

## Next pages — secret_value / 301302112223 / 4

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2031112112320023-2113300013121202-2023001332113201-1002000123330210-0121032211312110-3032133300101303-3300000213021300-2312330303231233)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-3112222122223320-0030311322322011-2101000223213320-3210330331031200-3311323031220032-3302233113130331-3312120333001332-3213031331112311)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2031112112320023-2113300013121202-2023001332113201-1002000123330210-0121032211312110-3032133300101303-3300000213021300-2312330303231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130000321120022-0102232222012233-2223300120020133-3230203013033121-1311121220101321-3300220130023001-3120230233301021-2123303101210203"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 220111211323 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1002102230330012-3122222012010311-0222330200330320-3210312211000010-0010202023332122-3330033020220023-0010301020033212-2303031023232022"></a>

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

<a id="canonical-2012023021212101-0213322123300010-0223302010010103-2013322231101300-3230202202231303-3201121110311020-2220003301330132-1010120111310132"></a>

## Direct properties — blindfold_secret_info / 220111211323 / 3

<a id="canonical-0322132111223111-2202133330222211-1221230231001330-3231222132300210-3010112231231102-2123131010101110-2202002103210011-3222132012300221"></a>

<a id="canonical-1111121203021312-3300113213020030-0202210132233123-3123203010310332-0301202330030312-2031333010301212-1003122230231331-2233302233330003"></a>

## decryption_provider property — blindfold_secret_info / 220111211323 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2320313302110102-2210321030231001-2101100211020111-2230323322223010-0310213100130211-1000210020132333-0322013331220231-3120313220311312"></a>

<a id="canonical-0002232200321111-3332033130230100-1011223003310121-0013203202322220-0223111001030112-1311302313003021-0210232003122120-1233100100202120"></a>

## location property — blindfold_secret_info / 220111211323 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1123323133232202-3211333100130330-0020002032203111-1322101201000131-1220020331212222-3032331030312213-0321111332111123-2331200003212000"></a>

<a id="canonical-3130032131200221-2003200101302223-0233032100000103-0122122212230131-1103323002210222-1220010301302122-3331201232223100-1101033323023030"></a>

## store_provider property — blindfold_secret_info / 220111211323 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2102113020120312-3131313202331113-1220033321322112-2100200312211130-0203113112231330-1322100321012123-0102321133333030-2231212123230022"></a>

## Next pages — blindfold_secret_info / 220111211323 / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3112222122223320-0030311322322011-2101000223213320-3210330331031200-3311323031220032-3302233113130331-3312120333001332-3213031331112311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123000302230123-1131211201003131-3013230132202331-3331200103323130-3021123210203311-2213010322023010-2112130312232102-3200311201312322"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 000012210100 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1023202223030322-0313112312020121-2212013320301123-2000311231112222-0221023111221332-1020202100200000-3021231222311020-3111212131220202"></a>

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

<a id="canonical-2013002212302320-0110301222331231-2011131033222212-2132112032302212-2233022312332330-0103312233303033-1120102322001333-1210320203012302"></a>

## Direct properties — clear_secret_info / 000012210100 / 3

<a id="canonical-3123032203120132-1120012321001000-0003211232211102-3312131213333323-3032311131330303-2023032022123231-0003031300032202-1123200033130021"></a>

<a id="canonical-1022002313100200-1102333312232101-3003101102233300-1033331111323121-3222322330033302-1130100333122021-2233230010031233-1200022322313223"></a>

## provider_ref property — clear_secret_info / 000012210100 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2122312013300023-1221200303212211-1102101332313232-1332332111322121-2203113320303313-2012103022033212-1020333120132211-1133323213322121"></a>

<a id="canonical-1211332200031201-0323001010301123-2322121013023223-0021300313322023-1223001221000001-1130032222131110-1322322010202020-0031020231132020"></a>

## URL property — clear_secret_info / 000012210100 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2003232110211110-2013332001113103-1120120002123113-0202300010111233-3120022103003311-1122213213101133-1210331331020321-1312332101002212"></a>

## Next pages — clear_secret_info / 000012210100 / 6

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022230331121330-3002221311131022-2331323102301103-3213010000313102-0202022003310032-1132132232130222-0020210200110002-0331001331121201"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add — response_cookies_to_add / 002131120332 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add

<a id="canonical-0203120121330102-2310212113121002-1120030322022220-0021020203203213-0212332201212022-2331232103202013-0313320031023130-1313210203310003"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013013102233211-3312321330201301-1101132320010210-3012130132311001-1002000030310221-0331110131211231-3210021231231323-3301033220022200"></a>

## Direct properties — response_cookies_to_add / 002131120332 / 3

<a id="canonical-2022112002030110-1222131032112301-1233132032302020-0120322100003302-2312102032222000-1332312211010230-3323001133230011-2232010333322202"></a>

<a id="canonical-3001033010332220-1123123232030030-2300022332310233-3332033113322013-1112113321133020-1312110323021110-2021102301222131-3020013011212001"></a>

## add_domain property — response_cookies_to_add / 002131120332 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1212032033300313-1301132213333302-3030222031003203-1322202300112311-2022221101100221-1110322100031020-3322033223130021-0122232100120202"></a>

<a id="canonical-1211203123030113-3120121231311331-0302221131301313-0232120000210311-2223222100320131-3213233310133310-0230102022223012-2111013111133032"></a>

## add_expiry property — response_cookies_to_add / 002131120332 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [add_httponly](data-sources--proxy--reference--group-003.md#canonical-3122230200320020-0203101011212112-1122120213302300-2310030121010020-2113103310313322-0030213121220120-1130321033312301-0201011231131013): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-003.md#canonical-2230311202132222-0220021001101113-1313002110233213-1303322302232111-3323121233121200-1313302123003032-0002313033020311-2001213003322312): complete subsection reference.

<a id="canonical-1210103000201213-3220112232320100-3330002123110310-2300011203210321-2210123033311210-3303213100123121-3000000213330323-1112320221323100"></a>

<a id="canonical-0310101220220211-1322311110131023-3231032200302032-0112122301111220-2021111331312002-2321100102013103-1332001122220200-2022310020103331"></a>

## add_path property — response_cookies_to_add / 002131120332 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [add_secure](data-sources--proxy--reference--group-003.md#canonical-1122110102003113-0201231213102303-0020300201220112-1222102303312001-2100031202323000-0000032132220301-3311200133123333-1121002013201130): complete subsection reference.

- [ignore_domain](data-sources--proxy--reference--group-003.md#canonical-1011223022130232-2213000332120032-2222200123303221-2122101103123003-2303123221003011-2300302132220023-1100123022122301-2333001300020220): complete subsection reference.

- [ignore_expiry](data-sources--proxy--reference--group-003.md#canonical-2100131121013311-2320001110331133-3102001101312223-1130310331113211-2112313320212311-0312012120323111-0020121300103212-3220033122233032): complete subsection reference.

- [ignore_httponly](data-sources--proxy--reference--group-003.md#canonical-3332320023202320-0023212233213011-2232311332000001-1311001012012110-2131000311011132-3000031312120200-2211003110213123-2031323332113112): complete subsection reference.

- [ignore_max_age](data-sources--proxy--reference--group-003.md#canonical-1021003112130310-1300012012002110-3321101000010322-3333030103303322-0330001321102311-3333111211302130-0310000210303330-1020232122301032): complete subsection reference.

- [ignore_partitioned](data-sources--proxy--reference--group-003.md#canonical-3120201003310332-1302111023203023-1301200121132200-3201021200113313-1210132121030221-3102312000200222-2032230323130000-0111220213101202): complete subsection reference.

- [ignore_path](data-sources--proxy--reference--group-003.md#canonical-2101230011122320-2101010310122300-0011120122303032-1111032230002223-2020201231332123-0031102301003313-3301111122131121-0030212001123201): complete subsection reference.

- [ignore_samesite](data-sources--proxy--reference--group-003.md#canonical-0301322122323102-1001130011132022-1213212213313111-0102220011102322-3210103332132032-1313111101213331-2322230101310332-3220312210231223): complete subsection reference.

- [ignore_secure](data-sources--proxy--reference--group-003.md#canonical-1030330302332322-0233031021220000-0100332101130121-1233000123120220-0203102112111100-2323322101123323-2003300203100223-2101002213311112): complete subsection reference.

- [ignore_value](data-sources--proxy--reference--group-003.md#canonical-1332223103223212-3322102032333023-3332322110323033-0103203131103223-0121033323102203-1303320032031201-3203102002320120-2303233220233220): complete subsection reference.

<a id="canonical-1212020221200133-2323012123101123-3213132102310110-2233230111231022-3102103021301310-2002221321223101-2021020021202002-1110002211310103"></a>

<a id="canonical-0123100000221231-2232301122022313-1212230202123012-2201211011322220-2013320111201223-0120213322231230-0200331323013001-2132221323211111"></a>

## max_age_value property — response_cookies_to_add / 002131120332 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0110202313010001-1121133213333113-0101222132333212-3200022023312210-2310131213301221-1331200000123030-0323333323032232-3223210130201322"></a>

<a id="canonical-3013102233201132-1001331012333123-2013111010301313-2321030133200031-3112220032100122-2313032002110201-1211033121233333-0322311303123032"></a>

## name property — response_cookies_to_add / 002131120332 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3000122130200330-3231321103321023-3331223223203131-1120233000131120-3333013101301112-0210113230232133-3200330101230123-0323223310210211"></a>

<a id="canonical-3232320120221212-1030121131222300-1000202211311210-1320020011212123-1203212211000120-0001230222200301-0331330113203330-1331233230032112"></a>

## overwrite property — response_cookies_to_add / 002131120332 / 9

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

- [samesite_lax](data-sources--proxy--reference--group-003.md#canonical-3102032203001213-3032011002021310-1032011123113110-0102132012023003-2232232132122312-1101322232111312-3321312303222022-2232233320002123): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-003.md#canonical-2212222330313101-2322111231313030-2332023121212222-3030001220232220-1130332111302120-3311322001200001-2122312100001220-0102023303131222): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-003.md#canonical-1321202301011021-2333122122030010-2212121313020312-0133303302201213-0023000333131221-0313003231300312-0011313331231210-0120322232202212): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130): complete subsection reference.

<a id="canonical-3032110323023033-2002113132223120-1200213300220020-0023233022202212-0012023221111023-0130313101113113-3001233302302130-2331033323302301"></a>

<a id="canonical-3333011220011311-0301300230110020-0110310310221021-0101021121021003-2010331103021233-0110132333000030-1031230032212011-2103013211203131"></a>

## value property — response_cookies_to_add / 002131120332 / 10

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1211123111102133-0200033311003332-0133330010112313-3003130021311102-1102200303010121-1033011333223303-0121202332310110-0003132033233111"></a>

## Next pages — response_cookies_to_add / 002131120332 / 11

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--reference--group-003.md#canonical-3122230200320020-0203101011212112-1122120213302300-2310030121010020-2113103310313322-0030213121220120-1130321033312301-0201011231131013)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--reference--group-003.md#canonical-2230311202132222-0220021001101113-1313002110233213-1303322302232111-3323121233121200-1313302123003032-0002313033020311-2001213003322312)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--reference--group-003.md#canonical-1122110102003113-0201231213102303-0020300201220112-1222102303312001-2100031202323000-0000032132220301-3311200133123333-1121002013201130)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--reference--group-003.md#canonical-1011223022130232-2213000332120032-2222200123303221-2122101103123003-2303123221003011-2300302132220023-1100123022122301-2333001300020220)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--reference--group-003.md#canonical-2100131121013311-2320001110331133-3102001101312223-1130310331113211-2112313320212311-0312012120323111-0020121300103212-3220033122233032)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--reference--group-003.md#canonical-3332320023202320-0023212233213011-2232311332000001-1311001012012110-2131000311011132-3000031312120200-2211003110213123-2031323332113112)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--reference--group-003.md#canonical-1021003112130310-1300012012002110-3321101000010322-3333030103303322-0330001321102311-3333111211302130-0310000210303330-1020232122301032)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--reference--group-003.md#canonical-3120201003310332-1302111023203023-1301200121132200-3201021200113313-1210132121030221-3102312000200222-2032230323130000-0111220213101202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--reference--group-003.md#canonical-2101230011122320-2101010310122300-0011120122303032-1111032230002223-2020201231332123-0031102301003313-3301111122131121-0030212001123201)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--reference--group-003.md#canonical-0301322122323102-1001130011132022-1213212213313111-0102220011102322-3210103332132032-1313111101213331-2322230101310332-3220312210231223)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--reference--group-003.md#canonical-1030330302332322-0233031021220000-0100332101130121-1233000123120220-0203102112111100-2323322101123323-2003300203100223-2101002213311112)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--reference--group-003.md#canonical-1332223103223212-3322102032333023-3332322110323033-0103203131103223-0121033323102203-1303320032031201-3203102002320120-2303233220233220)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--reference--group-003.md#canonical-3102032203001213-3032011002021310-1032011123113110-0102132012023003-2232232132122312-1101322232111312-3321312303222022-2232233320002123)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--reference--group-003.md#canonical-2212222330313101-2322111231313030-2332023121212222-3030001220232220-1130332111302120-3311322001200001-2122312100001220-0102023303131222)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--reference--group-003.md#canonical-1321202301011021-2333122122030010-2212121313020312-0133303302201213-0023000333131221-0313003231300312-0011313331231210-0120322232202212)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3122230200320020-0203101011212112-1122120213302300-2310030121010020-2113103310313322-0030213121220120-1130321033312301-0201011231131013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213213211322031-1203100210131022-1210220011020130-3321010232031032-0233313000323313-2020310023300102-0203312032122101-1202133200103322"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly — add_httponly / 132332121132 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-1030022032310100-2200333123033030-1020121311301202-3232002113321223-3210202201221232-1012003202213320-3233013010301320-1210031031113033"></a>

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

<a id="canonical-0120201212331330-2303203001112231-3023201313133233-2200321211000231-3013021301013222-3200310133321110-0132033233223213-1302023121010333"></a>

## Direct properties — add_httponly / 132332121132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012212311222110-1220021010323012-3202120112213111-0030221123213112-0331130213232313-3000211112112320-0010100011231313-1200113332013010"></a>

## Next pages — add_httponly / 132332121132 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2230311202132222-0220021001101113-1313002110233213-1303322302232111-3323121233121200-1313302123003032-0002313033020311-2001213003322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233201110033133-1033000012311201-2210101021113032-3003021330302033-0223302022230002-1102101301232101-0200213110132303-3133331002023211"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned — add_partitioned / 331301200330 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-2011313122003311-0223031331331003-2301310011200011-3112222330233212-1131120132013220-3123003001333333-3132230233110333-0311330131203110"></a>

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

<a id="canonical-2323323010021213-1100313023102120-1000132232230320-0011131221302130-1330203300010331-0202332131031022-0312313313313013-0201301133112011"></a>

## Direct properties — add_partitioned / 331301200330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331103210113112-0121212211020231-2131021121003303-3011330201021010-1012212230023322-2331020222111001-0331220123010021-0123200212031213"></a>

## Next pages — add_partitioned / 331301200330 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1122110102003113-0201231213102303-0020300201220112-1222102303312001-2100031202323000-0000032132220301-3311200133123333-1121002013201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223113110121223-2020112013120300-2222221312121011-0103111112103122-2333021032322323-3321012101200223-2121220230332301-1031011023223022"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure — add_secure / 303220222112 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-0321121222223030-1311333233112223-2103133110020100-0201112112220123-2223311032212113-2120131010310131-1012112301223313-3212002002312110"></a>

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

<a id="canonical-3332111011023303-0213103002303022-1210132313231330-0122111111031103-1022202022021330-2331313103013121-0000333031021011-2122022221212022"></a>

## Direct properties — add_secure / 303220222112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020102333232231-0300232222232320-2320211103231111-1201030223323310-1220002130021021-0331122203302310-0021311231203310-2200222213212112"></a>

## Next pages — add_secure / 303220222112 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1011223022130232-2213000332120032-2222200123303221-2122101103123003-2303123221003011-2300302132220023-1100123022122301-2333001300020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300011201020003-3311130321203311-2131002310203311-3132322132112213-3202112113000022-0031020002133003-1311320230212200-3212202323312011"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain — ignore_domain / 132120212020 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1231030112203133-1010332211102021-3210120233303311-3320121112233023-2123300013011131-2002103011020201-1232121131310112-1320221000011130"></a>

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

<a id="canonical-1213222023333232-2213230331113313-2211011301330022-0231313303330313-0312321320131211-1200310200213333-3310201013030120-0303121313233332"></a>

## Direct properties — ignore_domain / 132120212020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331212203300212-1220023002101010-0230022222030002-1211011302123011-3313100322131020-1020010112010120-2222102202013002-2101110220111001"></a>

## Next pages — ignore_domain / 132120212020 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2100131121013311-2320001110331133-3102001101312223-1130310331113211-2112313320212311-0312012120323111-0020121300103212-3220033122233032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020110201123222-1221132210220101-1202312321002202-2031303013113132-2022031012113303-0011103102122321-1001221002212232-2101100002111220"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 001012103022 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-0233113013210303-2013112211312213-3113312111130111-2202032233310023-3213131133321033-2113332020222301-1331220220232222-0033110133120232"></a>

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

<a id="canonical-1232300120130211-1023001010001001-3003232133110201-2001121233200122-1330333022213201-2110323110123330-2001200103222020-1120332103202112"></a>

## Direct properties — ignore_expiry / 001012103022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101313132330132-0021223130331301-1020020020211130-3021323331032333-0231200003013032-0000032202220031-2320022110311020-0213002333303100"></a>

## Next pages — ignore_expiry / 001012103022 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3332320023202320-0023212233213011-2232311332000001-1311001012012110-2131000311011132-3000031312120200-2211003110213123-2031323332113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122012023002321-3130203223002230-2131121132022133-1333303200313022-3102031311303220-2211013010303022-3200300300132000-0300333211102201"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 020232330132 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-0230123113032310-2322121012122312-3031113000122300-2221222231011322-1311103002101030-1013101312023002-1230013310330002-1331021203130223"></a>

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

<a id="canonical-3220333321131231-0121222222021323-2300013333003330-1201232213113002-1033102132320313-2212323003213231-1002110232321123-0212032011100000"></a>

## Direct properties — ignore_httponly / 020232330132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132202001103310-2321331313111100-3131130323133110-2131213100101233-2003322023013331-1122303112022332-1223123312323030-0101323030033331"></a>

## Next pages — ignore_httponly / 020232330132 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1021003112130310-1300012012002110-3321101000010322-3333030103303322-0330001321102311-3333111211302130-0310000210303330-1020232122301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313010033320023-3310201211312120-3112123120313232-2001232311302322-2231323130303122-0202031302011002-2120112200331100-1310330000312323"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 212121113333 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0131313331203112-2100101212131003-2020203023333111-0310100020310331-0310210320213022-2313111203333013-3003102002231012-0231303203003233"></a>

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

<a id="canonical-2003211120123322-3200331311120320-0031111232010220-0210301013121023-1012132233332232-0131211233123130-3020121230213323-3003310220121300"></a>

## Direct properties — ignore_max_age / 212121113333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012221221302130-1333100031031113-2333332103223203-3203233200121311-3031131310033323-2013030132220223-1223221111020113-1333323232200201"></a>

## Next pages — ignore_max_age / 212121113333 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3120201003310332-1302111023203023-1301200121132200-3201021200113313-1210132121030221-3102312000200222-2032230323130000-0111220213101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111012222320112-2213330122103222-3213103103303103-1130201112130001-0233023312132020-2302133220132010-3202201223323231-3102200222310201"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 220213232222 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-0231110023133002-3110330031210010-0320233231333020-0302020122033101-0330232221310331-0103321113031111-1131131223313020-2310010123333320"></a>

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

<a id="canonical-3110122212102300-2212003002103230-2220021010103322-1212222210021230-1001213021022112-2221202021203012-2231011003222212-2211320202302100"></a>

## Direct properties — ignore_partitioned / 220213232222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333202320133032-2333330022112213-1112330012103330-3332132310211320-0202303231033100-3203322222311033-3101321222232132-0120111112132322"></a>

## Next pages — ignore_partitioned / 220213232222 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2101230011122320-2101010310122300-0011120122303032-1111032230002223-2020201231332123-0031102301003313-3301111122131121-0030212001123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020113322122330-3223103332011122-0313230201330213-3022332011230011-3221211100231212-3330001030032320-0201202333133333-0000123013333300"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path — ignore_path / 020130203112 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-2030032232231213-3130313302001010-2300312000022300-1132120112323211-0231113112103320-0302023001012301-0322012212132303-3013222333100010"></a>

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

<a id="canonical-0002010322013113-2123100303331011-2001211330203221-0113130031323230-3132103322201221-1123001030012123-1233132003201202-3212332113212123"></a>

## Direct properties — ignore_path / 020130203112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033323001212122-0133111231302022-3313022302033230-2210301201121112-2011212323020211-3332210120120103-2031220130021210-2132000301111202"></a>

## Next pages — ignore_path / 020130203112 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0301322122323102-1001130011132022-1213212213313111-0102220011102322-3210103332132032-1313111101213331-2322230101310332-3220312210231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320032011000020-3232302030322301-3333022303010323-0000201210000122-0220120012201310-2123112321320321-1001211323322101-3230330312310332"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 133131002032 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-2203303222021322-0201021023033333-2123103311320302-1331302322030200-3301022300222132-0103101123022023-1001321330113020-3321332123013221"></a>

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

<a id="canonical-1330331333310131-1303331013302001-2132022111122030-0002233203332100-1313210031330321-2220310200110101-2312001312012221-2002023202100331"></a>

## Direct properties — ignore_samesite / 133131002032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133020102310122-1210310332003120-0220333112000111-1313333223102212-0323022320001120-1203022331131103-0033313301302003-2332303230233003"></a>

## Next pages — ignore_samesite / 133131002032 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1030330302332322-0233031021220000-0100332101130121-1233000123120220-0203102112111100-2323322101123323-2003300203100223-2101002213311112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033031011333112-0122300202212031-0213332033231103-3312303111231212-1133230301230000-1311210133301111-3202021322001333-2223312212120232"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure — ignore_secure / 323110002110 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-2200022321221103-2123330123332320-2311233131232032-2023203213210230-2101311320111103-2332003031300032-2100012103012022-2203222011323222"></a>

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

<a id="canonical-2010310111113132-0300100131121102-2322023303121100-0321201100221000-2113333031100232-1320203213310133-2031320020033300-0022330223332031"></a>

## Direct properties — ignore_secure / 323110002110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122302022120013-0010330112311013-2230030231012303-3003032232203101-3000210102322000-0103122102132030-3330010031221200-1233010331000212"></a>

## Next pages — ignore_secure / 323110002110 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1332223103223212-3322102032333023-3332322110323033-0103203131103223-0121033323102203-1303320032031201-3203102002320120-2303233220233220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311032201031211-0311330133221110-2203020211011033-0201102220230030-1101100033313020-1211001311323301-1022102323103021-0230000312002201"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value — ignore_value / 002003301211 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-3020130010203313-2210332222101212-1322222213313300-2101011023020002-2202022313232223-2223301103300232-2232102331233001-1020331111212232"></a>

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

<a id="canonical-2223233313031100-0220202023012022-3223012110113211-1331202122320101-0031212002030223-0210133131121111-0113230301302003-1102000002302030"></a>

## Direct properties — ignore_value / 002003301211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023302233030220-0312201122030131-0131231000222212-1210101133311121-3010112101111203-3031120312022303-2030312331000031-0231121323013131"></a>

## Next pages — ignore_value / 002003301211 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3102032203001213-3032011002021310-1032011123113110-0102132012023003-2232232132122312-1101322232111312-3321312303222022-2232233320002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331222330023200-0011031300012030-0132131110030012-1000120023113333-1132330211123032-0001121333300231-3020121213313203-1111000310012021"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax — samesite_lax / 300130221111 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2222110232022130-1222222010113003-0200310333010113-1220212012200300-0123132111000303-1001223013233230-1210320030233102-1211002312322133"></a>

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

<a id="canonical-3112123311322323-2120211022103312-1020102130030130-3320221022322330-0002120223332110-2103103331303220-2101120232212122-2013321233123211"></a>

## Direct properties — samesite_lax / 300130221111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311333033221313-3301010323232002-3010211332312321-1222002111301010-2133321200332333-1120110021033130-2110300011300301-3020130030310322"></a>

## Next pages — samesite_lax / 300130221111 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2212222330313101-2322111231313030-2332023121212222-3030001220232220-1130332111302120-3311322001200001-2122312100001220-0102023303131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113201133201111-1032331030013030-0012232211302230-1332323330132310-0220101202012001-0212330132320122-1210133213110301-3012001312122000"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none — samesite_none / 003302001012 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-0312001230223020-1002031033023122-1100130220123213-1310030333202222-1003211123331333-0331112021100212-1002312132332031-3203220321233001"></a>

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

<a id="canonical-1231320111003002-1033112223023313-2001120002300222-1310013300101300-3123131022333100-2001000121303213-2303200203200302-1230031302213113"></a>

## Direct properties — samesite_none / 003302001012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131021023201230-0010023212212302-3312031320220230-2310231110131131-3212020012310330-3313022330023010-1031002230312111-2023200121122013"></a>

## Next pages — samesite_none / 003302001012 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1321202301011021-2333122122030010-2212121313020312-0133303302201213-0023000333131221-0313003231300312-0011313331231210-0120322232202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212131200103120-0100333012322322-2231013203033310-3122213232022023-2300011331133013-3312012033300022-1133321122203331-1212210310122022"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict — samesite_strict / 130310020312 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1320223211213201-1223210201030313-1102123301201132-2100323030100020-0232022333320012-3200033230230033-0032330000330111-3213232300121022"></a>

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

<a id="canonical-1323121123011232-0001311002011133-1200001022202030-0221033333323233-3013321132003020-2212120110222333-1130332311032110-1213212131311131"></a>

## Direct properties — samesite_strict / 130310020312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131300121300312-3333300101212303-3201200203001130-3222011230303223-3312022021200100-2110112132120102-2101001101000200-1102112221230130"></a>

## Next pages — samesite_strict / 130310020312 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333200003120203-2222111013010110-0201021122233300-2313201133000020-1003222333313000-0000333130013211-0122210020322323-0201003223123221"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value — secret_value / 230201330332 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-3303231130222320-3223020332320100-0313213221120110-0122233002133301-2332120003203031-2332103202211133-2311010202321121-1202030103302313"></a>

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

<a id="canonical-1133133221233301-3031320023022103-1100033100231130-2020000233302013-0213132133000213-0011132001222331-3003231023200002-0131331303103322"></a>

## Direct properties — secret_value / 230201330332 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-1320221103311302-1333211323122013-2321233100010221-1021310323032121-0232220333021210-1323100023120001-1220033213233333-0312032322111111): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-0111102122200200-0231211332230122-0002101102313301-2021122122203210-0203231033302312-3222322203312130-2020310123200302-2303323322131300): complete subsection reference.

<a id="canonical-2003001312230010-1110131222031200-3321030111301322-1333130232033321-3130022323201222-0022232202003112-1212013330203122-1201323230001123"></a>

## Next pages — secret_value / 230201330332 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-1320221103311302-1333211323122013-2321233100010221-1021310323032121-0232220333021210-1323100023120001-1220033213233333-0312032322111111)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-0111102122200200-0231211332230122-0002101102313301-2021122122203210-0203231033302312-3222322203312130-2020310123200302-2303323322131300)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1320221103311302-1333211323122013-2321233100010221-1021310323032121-0232220333021210-1323100023120001-1220033213233333-0312032322111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210320330331021-1003331132133230-1221211130333203-3303111221031201-0333301210003313-1002333301122022-3131321010230021-3231032320332002"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 010020033222 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3301333212312201-3212102303333011-1300233102130222-2032333312031302-1323322320032323-1032203333321020-3111021223230103-0133010300100113"></a>

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

<a id="canonical-1032010333203031-0311210232233323-3113033322010220-0323133011033301-3222010211301030-1003031331031202-3121111112313123-2331130313011021"></a>

## Direct properties — blindfold_secret_info / 010020033222 / 3

<a id="canonical-0111120030232022-1323011110111102-3033202332200303-2212102123010221-1211132233012111-0203330322001021-0302132221231323-3030113031020103"></a>

<a id="canonical-1302030020123110-1332130233101023-3120032112331213-1203330011220303-3123331122230302-3130311323012302-1332022113020333-1020021210331023"></a>

## decryption_provider property — blindfold_secret_info / 010020033222 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3020112322000122-1122021322330203-1320311223332311-0331220232120013-2213313102123132-2000202211132200-1120321122023313-0013101230310330"></a>

<a id="canonical-3132031000013211-1033130112322232-2313311300113310-3230031321221332-0013220330000012-3201300213212222-3322121112230102-2333002203023013"></a>

## location property — blindfold_secret_info / 010020033222 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232222122232311-1303120010233130-2230232120103321-2023120003321320-1200211100322123-1133121230202310-3211203200201212-2010131310302021"></a>

<a id="canonical-0003132222023302-2013110202201001-3020132001002120-2302230322031213-0321331322111101-0320110211131103-3102100013032222-0133300133310022"></a>

## store_provider property — blindfold_secret_info / 010020033222 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2310331301133230-3011012131122202-3232133220120312-0213202301323311-2113323221330322-2033032033232010-2130300320020120-0030030110200222"></a>

## Next pages — blindfold_secret_info / 010020033222 / 7

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0111102122200200-0231211332230122-0002101102313301-2021122122203210-0203231033302312-3222322203312130-2020310123200302-2303323322131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313212111230123-0012020021112203-3330032121110313-0202030022220110-3113302320232002-1110213333020311-0130110333111211-2030303113301112"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 020001123111 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-003.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2200302302110211-1030300101310031-3322320121003133-0213223013303100-2221221310320301-2322100123021320-2320301032302133-0031200210212013"></a>

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

<a id="canonical-0202112332131303-2011312302013211-1201103301222330-0122003333030010-3300110233332000-2222210003323123-3332321011113130-0101132112112220"></a>

## Direct properties — clear_secret_info / 020001123111 / 3

<a id="canonical-0033313300123220-0303310302311121-3233300133011331-2303023003323011-2313313003332201-1000333331122312-1303033221011200-0321333230020320"></a>

<a id="canonical-2030323322202223-2202211000033221-2003331003212133-0132230223211231-1303202210201332-1330312311233013-0312032222232012-2003022232132120"></a>

## provider_ref property — clear_secret_info / 020001123111 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2021332301232030-3332130002023033-0213332130212210-3323202321013330-1003020310333300-0203131021231232-1330130101230111-2132011002112003"></a>

<a id="canonical-0223032102312130-1310111321032232-0031232133001303-2312113213201101-3312103332223113-0002032211130330-1310133203112123-2100312221330010"></a>

## URL property — clear_secret_info / 020001123111 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3202323232302220-1120120212212132-1122022333131230-1201110313213021-0223103222112122-0311221201001330-1013310200303313-0133000313313211"></a>

## Next pages — clear_secret_info / 020001123111 / 6

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320002031233322-3212323023103123-1201220013122320-3131322103003331-2313233133301133-3021000213131132-0110000321113331-3221320213222103"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add — response_headers_to_add / 321113021113 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add

<a id="canonical-2123101123332311-3031202301323130-0110300312210222-2331121320202100-1032103031232223-1221221122131211-3001213003030130-0201323233223000"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0023001121011230-1311320110201330-1223133010020220-2133331221002233-3202322132000033-1223221310122030-0011123230123231-2312120222312023"></a>

## Direct properties — response_headers_to_add / 321113021113 / 3

<a id="canonical-1203032111112212-1020332110123231-3130231133112021-0131313031133202-1321301012112002-3212120301020021-3231122210222203-1212212233102003"></a>

<a id="canonical-1313200022301201-0323002103303123-3103133013202020-2110202003312022-3033213033333002-2230113230133323-3220120002332103-2022213221112201"></a>

## append property — response_headers_to_add / 321113021113 / 4

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

<a id="canonical-3011221200301132-0212202310111112-0303101213203330-2211100330222323-1201320103233113-0310121302112332-1220031203302303-2032112113121230"></a>

<a id="canonical-0330201310020132-1033031111213321-0323201103122202-0200103321100100-1130311333211233-1130322303132110-3123130312003031-1010031033322112"></a>

## name property — response_headers_to_add / 321113021113 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331): complete subsection reference.

<a id="canonical-1012013020301023-2200130313023231-0212323002033231-3023001233120132-1120300321332100-3011223232231003-3001301013222003-1022111320101103"></a>

<a id="canonical-2121021112220231-1323313232222212-1012102002102123-0111131023321111-1120332223210033-0003230332102322-3101122320322332-3022033332110002"></a>

## value property — response_headers_to_add / 321113021113 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0212033303110211-3100101010332032-2132203013320232-3330222310102021-0130030312233211-1200011313313303-1030230223221130-0131323212232333"></a>

## Next pages — response_headers_to_add / 321113021113 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330312122121231-1122020131121123-0322120031132222-1322101221033213-2101032203330111-1033101203211330-3322032110313121-3003031011203333"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value — secret_value / 331111332233 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-0331200100122020-0301121230130102-2003231012100132-0112200302302112-0323233221121232-2003030300322010-3302211031220210-0311310312003221"></a>

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

<a id="canonical-0113010100212312-2110202122313032-0200332223012002-2230013330320311-2021330112210011-3011301210110202-1102123122332120-2203302102121201"></a>

## Direct properties — secret_value / 331111332233 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2003012131112120-3303100100311231-1030111010120121-1300333220012030-0023121100122323-3130010332130313-2303100311223230-0003232123222222): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-1102122120002100-1201303223210221-0233320231232230-1321212120103212-2130002101011111-0010031200301122-2301221310332231-3000312110121100): complete subsection reference.

<a id="canonical-3132203201020012-3232102010132313-1132232121001103-0131021030110123-3033330231130133-3220131212232102-1333300222331010-3313312332212011"></a>

## Next pages — secret_value / 331111332233 / 4

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2003012131112120-3303100100311231-1030111010120121-1300333220012030-0023121100122323-3130010332130313-2303100311223230-0003232123222222)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-1102122120002100-1201303223210221-0233320231232230-1321212120103212-2130002101011111-0010031200301122-2301221310332231-3000312110121100)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2003012131112120-3303100100311231-1030111010120121-1300333220012030-0023121100122323-3130010332130313-2303100311223230-0003232123222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333111233310223-0201123230003121-3203203211230003-1202231121001103-0201302310212022-0333222100311213-1321123022323323-3101012220111132"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 003302030011 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3033000212132313-1011020111322012-0131230311002101-0221211232220222-0200312013302312-0302011011221003-3111030312100301-3121130012003023"></a>

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

<a id="canonical-0322320301130013-1313232223332132-1200201111311111-1310200000131100-2321312301022333-2123011001211132-3331333221010301-2313203321010311"></a>

## Direct properties — blindfold_secret_info / 003302030011 / 3

<a id="canonical-1333311221313012-0030120122211311-1133310333000323-1330123010022211-1113302230030213-1130310111220103-3212022001202221-0002211203133121"></a>

<a id="canonical-1330100201231311-2131110000323221-2030303233331022-0010233203123230-1201222211231023-3122101333232033-0211000010133033-1213032130112020"></a>

## decryption_provider property — blindfold_secret_info / 003302030011 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3310030111011022-3321011130112133-3113132022130011-3233200231011133-1220312210201110-0032212231001233-2112111300221030-3110021021200320"></a>

<a id="canonical-0012030212111201-2132002333033010-0102320121120331-1300002112312101-1203130232202030-1033033311230302-0303203102230030-2313212101201220"></a>

## location property — blindfold_secret_info / 003302030011 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0222300231300331-0132020303031330-3131022120033021-2013110032131232-1332220232013012-1211113200120330-0010022120030132-0001313023101200"></a>

<a id="canonical-2330100113010313-3122202303200321-1123113200002032-1231010022033210-3230301113303032-1320011101131201-1210201333230231-2010121203120130"></a>

## store_provider property — blindfold_secret_info / 003302030011 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0233320000010101-2310121321003100-1102202110031311-0133132232200320-0221230110310011-3302023133120331-0021302133332302-3020032302301033"></a>

## Next pages — blindfold_secret_info / 003302030011 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1102122120002100-1201303223210221-0233320231232230-1321212120103212-2130002101011111-0010031200301122-2301221310332231-3000312110121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013233103103001-1302231202320212-0222131000033210-3002333202133130-0020021311001021-0203131010010131-3231013301020120-2330322000123022"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 213110002200 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--reference--group-003.md#canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0012100123010332-3102003203112311-2012322311223210-0021133121103213-1221210203230022-1023023030322011-0121323120312313-1300132303101320"></a>

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

<a id="canonical-2021222031133103-1110322300202330-3303223013130130-0021233021013320-1301023031201310-2212220031010323-0131232211132233-2103010020320033"></a>

## Direct properties — clear_secret_info / 213110002200 / 3

<a id="canonical-1333200201101002-1033310131011013-0321333022133222-0010322133221113-1202013000131213-1333203313123211-0333322113023222-1221302001133332"></a>

<a id="canonical-2333210010301030-1102011132210101-3223332320103032-3030101032303000-2122030313211332-2110330001023003-3330103201010200-0132021032003111"></a>

## provider_ref property — clear_secret_info / 213110002200 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2032211333012200-0220231303113202-2311000022011121-1102123302100200-1000332131033000-1231001213202310-0332003312123230-2013230032302311"></a>

<a id="canonical-1221102101101111-2300111000302113-2123332302303312-1131011303331013-3001230232003003-3303023133302003-3020230230131123-0312331030102301"></a>

## URL property — clear_secret_info / 213110002200 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1113203011113023-2231020023123010-2322133323322102-3122032201223102-0020210022222122-3323000121313332-3232003101310133-1111000203330022"></a>

## Next pages — clear_secret_info / 213110002200 / 6

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022303120223303-0122313013020221-2213103331111301-0321130331200111-2030100012213331-0013110312002332-0312010230201312-0122130230023002"></a>

## dynamic_proxy.https_proxy.tls_params — tls_params / 203312130122 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- dynamic_proxy.https_proxy.tls_params

<a id="canonical-2300121300020131-3132302333231323-3132222033011130-0022032313222002-2010322122031022-0022320002210033-3111311330003131-1101302211212013"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-1111101113133110-1012321131202103-1320001112303112-3111122320311003-0211131022120322-0231001013121333-1032321031301023-0302310311022330"></a>

## Direct properties — tls_params / 203312130122 / 3

- [no_mtls](data-sources--proxy--reference--group-003.md#canonical-1023112302312003-1312011202020002-0322211203230223-3231321001022232-2221312333131312-3220322311022112-2120203313101113-0311322212001220): complete subsection reference.

- [tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001): complete subsection reference.

- [tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001): complete subsection reference.

- [use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131): complete subsection reference.

<a id="canonical-1203223301223120-3323311203013101-3123331202133311-3300230033011011-0300220221320030-1010010111112121-2301013121233211-3121100213311311"></a>

## Next pages — tls_params / 203312130122 / 4

- [dynamic_proxy.https_proxy.tls_params.no_mtls](data-sources--proxy--reference--group-003.md#canonical-1023112302312003-1312011202020002-0322211203230223-3231321001022232-2221312333131312-3220322311022112-2120203313101113-0311322212001220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1023112302312003-1312011202020002-0322211203230223-3231321001022232-2221312333131312-3220322311022112-2120203313101113-0311322212001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023012033103221-0233333001321303-3203001121000131-2111300231300121-0003302033012112-3331312301012221-3323302023122303-0112311010100313"></a>

## dynamic_proxy.https_proxy.tls_params.no_mtls — no_mtls / 203221332321 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- dynamic_proxy.https_proxy.tls_params.no_mtls

<a id="canonical-1002223300332130-1230101312003133-3303313330101023-2021123122231123-3101112032202332-3300220121122133-2231212301330201-2333313211300021"></a>

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

<a id="canonical-3303103222000202-2000023300020130-1313323103123001-0301111321123210-1110211111001322-1003032132111031-1112020032231123-1210102323123313"></a>

## Direct properties — no_mtls / 203221332321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201321130020013-1313130213020212-3332220023311000-3221302201003031-2121313303233231-1220323212131300-3311321012232130-0212212202113321"></a>

## Next pages — no_mtls / 203221332321 / 4

- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103301230022313-1012103302203333-2001200110030302-0021003033010223-3322013133012311-0021113010203321-2021323231002001-0200233013221131"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates — tls_certificates / 003200121221 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- dynamic_proxy.https_proxy.tls_params.tls_certificates

<a id="canonical-1002312003201232-2023232030333123-2231313101130321-2323101203222030-2113231220210200-3033303331023323-1112300223332130-0231230210233312"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2101202231331212-0032003122132333-1312311330321210-2223323212310303-1133202221303103-3023331013010120-3221302122100121-0300103323302330"></a>

## Direct properties — tls_certificates / 003200121221 / 3

<a id="canonical-1300132222232231-1223212123003021-0111212022130003-3323320311212010-3300001320121301-3100232323012031-1010023213011301-1113232112020030"></a>

<a id="canonical-2023023323333010-0120013132333330-1010131132033201-2021132311221021-1221331131302310-2100332012110031-1223130321121030-3112023012223210"></a>

## certificate_url property — tls_certificates / 003200121221 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-0231003113201323-1020120131300322-3230321110203300-1102022300311322-0110022231221221-0131330230232000-3023032332311103-3032202212033123): complete subsection reference.

<a id="canonical-2230333312130010-3022222130132212-0231222010012012-0021310101330113-2121333113203120-0132322321331000-1220303031313233-3221122132103232"></a>

<a id="canonical-2122203303310203-3102022213210300-0211230103333133-1110201012302102-3323203132020330-0223220121202113-0310211121320011-1200130030122210"></a>

## description_spec property — tls_certificates / 003200121221 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--reference--group-003.md#canonical-2101033220232001-1100001001011103-2233202122213332-0232301311211332-3123231020112122-1133123101332221-0102211102111112-0223212113212131): complete subsection reference.

- [private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100): complete subsection reference.

- [use_system_defaults](data-sources--proxy--reference--group-003.md#canonical-0210120212131231-3021303200313100-3212202103233020-1011221011112310-2110003203323012-1310103203022122-2230011320133221-2011331231201332): complete subsection reference.

<a id="canonical-2312213323113132-3330112333131200-1233033131222032-1210321131033330-3330211030113131-3300110023203000-0211312100131132-1202112122331032"></a>

## Next pages — tls_certificates / 003200121221 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](data-sources--proxy--reference--group-003.md#canonical-0231003113201323-1020120131300322-3230321110203300-1102022300311322-0110022231221221-0131330230232000-3023032332311103-3032202212033123)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](data-sources--proxy--reference--group-003.md#canonical-2101033220232001-1100001001011103-2233202122213332-0232301311211332-3123231020112122-1133123101332221-0102211102111112-0223212113212131)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](data-sources--proxy--reference--group-003.md#canonical-0210120212131231-3021303200313100-3212202103233020-1011221011112310-2110003203323012-1310103203022122-2230011320133221-2011331231201332)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0231003113201323-1020120131300322-3230321110203300-1102022300311322-0110022231221221-0131330230232000-3023032332311103-3032202212033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303220230100201-3312203323213120-2203121330302331-2212231220010132-2313130201221232-1230302213003001-1330001200132110-1300001033330111"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 310103021130 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

<a id="canonical-3133222303131002-2122230320100212-1012211110031022-3310322230231332-2203323012000233-3113332121322211-0201233033003031-0111031300323231"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2130320231011132-3202012011233010-3012222012201231-1201012313121332-1300031320001030-0232330003200001-3130301303100222-1312112332113001"></a>

## Direct properties — custom_hash_algorithms / 310103021130 / 3

<a id="canonical-2023110202032210-1030213310201112-3122032123133210-2011232302131231-1023103201122233-1311100200330100-3131203213111303-2310332331120323"></a>

<a id="canonical-1302303111003131-3311312321133323-1000232031001103-3010110302120200-0121001111323200-1312110123110121-1122123013311023-3131300110130212"></a>

## hash_algorithms property — custom_hash_algorithms / 310103021130 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0113133012021131-1130211320230233-0021031222202101-2302333222011012-1321213031002022-3132213330211101-3213200022031112-1233233103003200"></a>

## Next pages — custom_hash_algorithms / 310103021130 / 5

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2101033220232001-1100001001011103-2233202122213332-0232301311211332-3123231020112122-1133123101332221-0102211102111112-0223212113212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132331020110123-2303110110133023-1031003113313111-0310103023331101-2220111000011010-2101210011102311-2331112230311301-0231102022202032"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 100132123230 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-3112022122101323-0202021220102300-2301030102212030-1322202212310302-0232120133103111-0223301011113111-1013322303020021-0122230123032013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-2232120200233311-0333110033130300-2221231011310232-2320231030333313-2301230230320122-2210303213301022-0023110032112332-3220300211323032"></a>

## Direct properties — disable_ocsp_stapling / 100132123230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122311203301112-2213002213130310-2000301033120211-0331011300300200-1232033321213111-3000011103133233-1230003122222011-0321300303300021"></a>

## Next pages — disable_ocsp_stapling / 100132123230 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222210212201000-3303123022131020-1233312310132322-0202123311312210-2020022330202233-3233002331302312-2031002212001203-3313110301003110"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key — private_key / 300300200301 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

<a id="canonical-2201122231301012-0023331312313012-3111012303001312-3313003131030003-2202110322321101-1013333300130223-1002122210212113-3323121010100111"></a>

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

<a id="canonical-1311331010002011-2301212033203311-2021331113123330-2022003210330131-1112332332301010-1111333223030223-0123201300133302-2121033031203121"></a>

## Direct properties — private_key / 300300200301 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-1302122310302002-0120221303120332-3233300233112211-1113333001111131-0232313021110311-0020131020103302-2030301211033310-2122113320302021): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-3303200310023122-1032010022210020-3021121023320000-0112130121302001-3221302313312310-3322301321312233-2031200011331012-1200210233221132): complete subsection reference.

<a id="canonical-0221330310300211-3003130303112013-2320132301333012-2002231311011232-3310023032220022-0033200332020313-1121211221021223-2020201120012200"></a>

## Next pages — private_key / 300300200301 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-1302122310302002-0120221303120332-3233300233112211-1113333001111131-0232313021110311-0020131020103302-2030301211033310-2122113320302021)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-3303200310023122-1032010022210020-3021121023320000-0112130121302001-3221302313312310-3322301321312233-2031200011331012-1200210233221132)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1302122310302002-0120221303120332-3233300233112211-1113333001111131-0232313021110311-0020131020103302-2030301211033310-2122113320302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123202101000031-0010000101323130-1133131211000332-2232331232201322-1030103221211000-2133032220031122-2301210233300132-2030301233122101"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 013010311231 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0220201112112132-0230303230010213-2201101010300111-0110301120022122-1013012213023133-2132122232312123-1032211112121132-3022210313110030"></a>

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

<a id="canonical-0211301212003110-0301023020200310-2221133122213223-1023320302313120-2200110332333110-0321001330002001-1132203223300032-3233101033312001"></a>

## Direct properties — blindfold_secret_info / 013010311231 / 3

<a id="canonical-0113332310103101-2201013210101201-0322100201210010-0130031001220132-1101031000312111-0101223011302321-0100312211232230-1120120103220211"></a>

<a id="canonical-2130303302331103-3310330131223301-0221033320232133-3310222301201000-2322331120221200-3333330322122001-1220122032333232-2031203200200202"></a>

## decryption_provider property — blindfold_secret_info / 013010311231 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3311220132230002-2320322001203333-1233002013210001-2132331211212302-2000010110112303-0021123101023301-1103200010131302-1112001132333203"></a>

<a id="canonical-2310021302002101-1113130012212230-3111330121203131-3300120222110003-3102303132131113-1312321030330331-3012333013031210-0103222231031013"></a>

## location property — blindfold_secret_info / 013010311231 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1211132311312320-3000011013212332-2031301121321132-0331222330001310-1321220321232222-0333220001221023-1301002123310200-2311231002100201"></a>

<a id="canonical-1322102121122020-1112202031302022-1202032110000123-1221322013011233-0023331003030201-1311011303102002-0133330313230031-0311330231302030"></a>

## store_provider property — blindfold_secret_info / 013010311231 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2123300313300311-1001310120233131-1113333113000122-1111323310003110-1211221303213301-3303310033003102-2312023103131210-0301011211001112"></a>

## Next pages — blindfold_secret_info / 013010311231 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3303200310023122-1032010022210020-3021121023320000-0112130121302001-3221302313312310-3322301321312233-2031200011331012-1200210233221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320313333313031-2231331023223112-3302100202322013-1123202130001330-3323122323211113-2113331013310103-2321001022322022-3001002220032230"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 310100231231 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-2232203133202201-3333221331213311-3021233210331122-3300223213000300-0221232032101112-1030222132330303-3320132222023033-1210122333011033"></a>

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

<a id="canonical-2213210001122112-1011122312301220-3021331220212122-1211232001003100-3122002121333000-1010203300033123-2210122101221322-2013110223002012"></a>

## Direct properties — clear_secret_info / 310100231231 / 3

<a id="canonical-3320300320001002-3130031302021133-1313031321123203-2322022303002203-2332123220023202-0222101023033101-2213230201032202-1112112133200033"></a>

<a id="canonical-3302203221211222-0021313123002012-0331332302320310-2113212310020131-1103132203113102-2211010103212022-0023211201130131-3010302010330031"></a>

## provider_ref property — clear_secret_info / 310100231231 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1020100020033023-1120223332312312-3100031003333000-1023013321312232-1023031311213011-2030022002003200-3132232223321303-0311003013111311"></a>

<a id="canonical-0331103300212231-1003110000210312-2123303033030210-2002301021000331-0103022320323202-3020021111313021-1221333330002322-1331213121111102"></a>

## URL property — clear_secret_info / 310100231231 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021023322223100-0231332213121221-3332202300031111-2222311122111220-1313330112032120-1031230223010231-0223331232132020-2032223221020231"></a>

## Next pages — clear_secret_info / 310100231231 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0210120212131231-3021303200313100-3212202103233020-1011221011112310-2110003203323012-1310103203022122-2230011320133221-2011331231201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323202030223001-0220120132303002-1010001220021223-2332222130230020-3011011001111210-0101120212001013-3200032212020333-3230210210002202"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults — use_system_defaults / 331020303332 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults

<a id="canonical-2232201200302112-0320111133320302-1310300220032100-1203112203332311-1200031222102001-3102333300013303-2123120102021220-3231013120212222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-0220133210020121-0100003113113322-1031301233310111-3023320031123323-2132032300211303-0312112101230230-0232311012222130-1030030022202021"></a>

## Direct properties — use_system_defaults / 331020303332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203133333030031-2300211111020131-1012122322201203-0203222032230223-3322311010032032-0113102303021032-0210001202223333-2210112223000221"></a>

## Next pages — use_system_defaults / 331020303332 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011202122021232-0220302232120030-0322131001120301-1313300221111130-2220300300332310-1113030010222032-2323120033002030-1213101212000132"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config — tls_config / 002222323121 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- dynamic_proxy.https_proxy.tls_params.tls_config

<a id="canonical-2233112202122010-3010112132002021-3210133102201233-0102331302303331-1203210212030220-2213011200032022-2112200013302032-2331033311323313"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-2313120230210130-1311200021123323-2010110003323323-1132321102023110-1101133211112003-3031111023303303-2311020220210332-3232222111102122"></a>

## Direct properties — tls_config / 002222323121 / 3

- [custom_security](data-sources--proxy--reference--group-003.md#canonical-2211120102100322-3102312113211200-2310020331210123-2110111320121202-0002013212020332-3223301032211222-3023203212010123-1023312211211213): complete subsection reference.

- [default_security](data-sources--proxy--reference--group-003.md#canonical-2003120012313300-2200013003312231-0231010021103102-0222300032223222-1320123321121121-3202002011130101-2011312223313133-3132102111003023): complete subsection reference.

- [low_security](data-sources--proxy--reference--group-003.md#canonical-1322232001111230-3033322212200012-0223101130132111-3222202020301011-0031003132111201-3030013311200103-2131102320001111-3330323312313210): complete subsection reference.

- [medium_security](data-sources--proxy--reference--group-003.md#canonical-2210312203012321-3301212233210203-1331333021200202-1033303202302320-1121210033321312-0103311220213133-1122313020201213-2211110021222103): complete subsection reference.

<a id="canonical-1331200112120210-3231322002103320-2113322120013103-0212112302001110-1310003202013323-3113310021110320-0133212211200112-3201202232223230"></a>

## Next pages — tls_config / 002222323121 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](data-sources--proxy--reference--group-003.md#canonical-2211120102100322-3102312113211200-2310020331210123-2110111320121202-0002013212020332-3223301032211222-3023203212010123-1023312211211213)
- [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](data-sources--proxy--reference--group-003.md#canonical-2003120012313300-2200013003312231-0231010021103102-0222300032223222-1320123321121121-3202002011130101-2011312223313133-3132102111003023)
- [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](data-sources--proxy--reference--group-003.md#canonical-1322232001111230-3033322212200012-0223101130132111-3222202020301011-0031003132111201-3030013311200103-2131102320001111-3330323312313210)
- [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](data-sources--proxy--reference--group-003.md#canonical-2210312203012321-3301212233210203-1331333021200202-1033303202302320-1121210033321312-0103311220213133-1122313020201213-2211110021222103)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2211120102100322-3102312113211200-2310020331210123-2110111320121202-0002013212020332-3223301032211222-3023203212010123-1023312211211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312223310123333-3132232221302133-1330033022231223-3012113110323220-1320133021320201-0111133201221310-3131130313301033-3212320022123220"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.custom_security — custom_security / 100122121210 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- dynamic_proxy.https_proxy.tls_params.tls_config.custom_security

<a id="canonical-0101030312021031-2020030211033233-0322120320230323-2111320223003101-1311201213323310-2221133022000013-1211312111311020-1130111332220323"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0021203212322311-1203001102222310-1102003033111303-0111301323312130-0222310103103120-1233001323120032-2120110022033223-2302110102213011"></a>

## Direct properties — custom_security / 100122121210 / 3

<a id="canonical-2003332020112002-3210312220301022-3001002301303232-0131013312213112-1302033021123101-1000322131321013-1123301100302321-0113312212001110"></a>

<a id="canonical-1032133211212032-3223103011123201-0021311122130310-1223201131101220-1201320001320102-2111212201011331-3310323022102332-1010312230200200"></a>

## cipher_suites property — custom_security / 100122121210 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1210233313320332-1121130020313022-2022331102103000-3301030303210330-3222112231001010-0010233033112312-2121020022111110-3013031130330231"></a>

<a id="canonical-2233131133100333-2110200100010303-3210332322232200-2020300321013223-1231222103322311-1000213001020120-2330002030320021-2321312223100201"></a>

## max_version property — custom_security / 100122121210 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022111030033021-0030322110011203-2121133310121210-0321323120330203-0212333011320102-1201123131033233-3113303223023031-0231300120112233"></a>

<a id="canonical-3120122113000322-0211221200021310-2310201103331021-1202300123023023-2101211300311231-2121003111121023-1221230032210302-1012131012131031"></a>

## min_version property — custom_security / 100122121210 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322321202020222-3123122032220221-3132212101321213-3001213310203210-1322213133310002-1212101003310001-2022003021221100-2222111230101023"></a>

## Next pages — custom_security / 100122121210 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2003120012313300-2200013003312231-0231010021103102-0222300032223222-1320123321121121-3202002011130101-2011312223313133-3132102111003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131210232210002-2101011110130230-0302011110200202-0223131233022012-2111023003321010-1211231002313130-0320131212001201-3223032330200331"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.default_security — default_security / 112010122021 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- dynamic_proxy.https_proxy.tls_params.tls_config.default_security

<a id="canonical-3223030231100320-2012111120031111-2301211201031211-3332323200303003-0230100011031211-2233131202203130-2220231033020220-2013213212202322"></a>

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

<a id="canonical-2213332112231221-1101010212111111-0312200121303212-1331122001000130-2313300132320033-1323202123312132-3111110103201303-2211101313102020"></a>

## Direct properties — default_security / 112010122021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223330221233311-0010231232003210-2320210201112200-3103112023003003-1231021222000022-1003200322021130-2033102100201011-0012302211031322"></a>

## Next pages — default_security / 112010122021 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1322232001111230-3033322212200012-0223101130132111-3222202020301011-0031003132111201-3030013311200103-2131102320001111-3330323312313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323320302231133-3301231110303320-2121001223321223-3110202212332132-1320033311333311-0311032030112303-3201333130130331-1211322231300011"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.low_security — low_security / 121011021323 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- dynamic_proxy.https_proxy.tls_params.tls_config.low_security

<a id="canonical-1220303220311330-1101122010200201-3000233111013203-3201023330031113-0301031121020101-1213213110223333-3322012113303323-3021021202113121"></a>

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

<a id="canonical-3010313213001303-1000111032031030-3330012133211120-3132011313312131-0222321220312213-1210332120133310-0212230210023302-3203021212100221"></a>

## Direct properties — low_security / 121011021323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212122221330232-0022030331030231-2300100022030230-0020331103113323-3100320220303212-1310130221123122-1223222321101312-1122102333020220"></a>

## Next pages — low_security / 121011021323 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2210312203012321-3301212233210203-1331333021200202-1033303202302320-1121210033321312-0103311220213133-1122313020201213-2211110021222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213110212110322-3200312322120320-1200033311130201-2332300113333111-2110213200011001-0111133212303320-1230010010201120-0320311230230100"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.medium_security — medium_security / 331133133031 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- dynamic_proxy.https_proxy.tls_params.tls_config.medium_security

<a id="canonical-2131330032020130-3311112022321122-3311023213102130-0232202032303231-0032030210120103-0030021122101131-3122300233323213-1303300121302012"></a>

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

<a id="canonical-1033330332222303-3112001320210213-2111211001332312-2001202233303100-2301210112001021-1012332003320002-3001102011110103-1002223010211233"></a>

## Direct properties — medium_security / 331133133031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233033030011131-0020200232200322-2321031000333331-1021011231311012-1010203101233121-1010212231113300-0330310001010011-3021213222311023"></a>

## Next pages — medium_security / 331133133031 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132011332221321-1332333303223013-2320030300331101-1321130300201330-1302323113322300-0123023033113203-0113000131222220-3022212310331230"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls — use_mtls / 023321120020 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- dynamic_proxy.https_proxy.tls_params.use_mtls

<a id="canonical-2000211023303131-0012230310021300-0022112001322322-2332011321010212-2100130113212330-1020212222020300-0021120300200211-3333001223320132"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-3223210031202121-0221331202322120-2203321031002231-0031221123033333-2302300212322320-0213020102012013-0103111220132333-1222130211000031"></a>

## Direct properties — use_mtls / 023321120020 / 3

<a id="canonical-1103113010020312-3300002120333031-2211001101000133-0312213313003010-2011223103131322-0111321133213011-3021310013032100-0321210202313300"></a>

<a id="canonical-0112312203010103-2312133020202100-1020022101113313-3232303311323031-3320031102032012-1322332121333132-2202232000120130-2031202212111201"></a>

## client_certificate_optional property — use_mtls / 023321120020 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--proxy--reference--group-003.md#canonical-3232321001130202-1301031200001033-3220030003032210-0313323223012122-3330110213230122-1201302333323310-3001302213000023-1021112101033001): complete subsection reference.

- [no_crl](data-sources--proxy--reference--group-003.md#canonical-2302132330013031-1101202110332220-2320320230323133-3210130122131121-3303232022221020-2223302103232311-2233222021332233-1212230303313020): complete subsection reference.

- [trusted_ca](data-sources--proxy--reference--group-003.md#canonical-0113322102331021-0102312331132203-1131012310302113-3003012200201322-0012303213001132-2323010231111133-2130300201322033-1031010133201001): complete subsection reference.

<a id="canonical-3210023202232112-0212033313200000-3021111130211213-3022111022021333-0321132231022130-2330322320200102-0101321223003003-0013330323121332"></a>

<a id="canonical-2201213201212223-3110100011033033-0233310033320323-2303203022221331-1311333223101131-1023120201133301-2302131302112333-1132021133221210"></a>

## trusted_ca_url property — use_mtls / 023321120020 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--proxy--reference--group-004.md#canonical-2010122003302033-1301211033330203-3022022121101331-3100102323332312-3002231211130301-2332311313323130-1232011003111010-0213113320211312): complete subsection reference.

- [xfcc_options](data-sources--proxy--reference--group-004.md#canonical-3121122030011130-1120332230321011-1210123321213212-1333110133301201-0312133122322311-0000300211111022-1133132110103131-1011321303323000): complete subsection reference.

<a id="canonical-3222312000103100-0020310122232323-0311100333120121-1130311100313233-3331200031001011-3033021022212211-3110031012123013-0200023033131321"></a>

## Next pages — use_mtls / 023321120020 / 6

- [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](data-sources--proxy--reference--group-003.md#canonical-3232321001130202-1301031200001033-3220030003032210-0313323223012122-3330110213230122-1201302333323310-3001302213000023-1021112101033001)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](data-sources--proxy--reference--group-003.md#canonical-2302132330013031-1101202110332220-2320320230323133-3210130122131121-3303232022221020-2223302103232311-2233222021332233-1212230303313020)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](data-sources--proxy--reference--group-003.md#canonical-0113322102331021-0102312331132203-1131012310302113-3003012200201322-0012303213001132-2323010231111133-2130300201322033-1031010133201001)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](data-sources--proxy--reference--group-004.md#canonical-2010122003302033-1301211033330203-3022022121101331-3100102323332312-3002231211130301-2332311313323130-1232011003111010-0213113320211312)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](data-sources--proxy--reference--group-004.md#canonical-3121122030011130-1120332230321011-1210123321213212-1333110133301201-0312133122322311-0000300211111022-1133132110103131-1011321303323000)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3232321001130202-1301031200001033-3220030003032210-0313323223012122-3330110213230122-1201302333323310-3001302213000023-1021112101033001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203322011110232-2033312110102311-2302033122033312-1210212001313113-0023201123033023-2000211012322112-1021210213100320-1133323303201113"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.crl — crl / 012012033300 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- dynamic_proxy.https_proxy.tls_params.use_mtls.crl

<a id="canonical-1303123232010222-2232220211003013-2023320332332330-3123120313001031-3332300223103232-2311121201101030-1020002203002022-2103022013322123"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2033231112112220-3130211130122022-2110130123002013-3202120020330330-3010031123132311-1130101220112211-2111033321310112-3130332101300231"></a>

## Direct properties — crl / 012012033300 / 3

<a id="canonical-0133030130112123-1212102000133321-0311213121323301-2323033000122222-3211023210313102-0203200033231223-1231313331030133-0210213233223112"></a>

<a id="canonical-2200012223100311-2302120320122121-2020013023112332-0033002113100120-2302331232132002-0123121122201102-3332020021301112-0022320230311323"></a>

## name property — crl / 012012033300 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3223210211221312-1120002302130021-2320323020311011-0132201003303232-2313133110133002-2001030003300213-2213202113230030-0330122322000131"></a>

<a id="canonical-0312211323122322-1300011000202203-1123002122211332-3333003210133020-2331003121221023-0032203233303103-2020130223330111-3120022330012230"></a>

## namespace property — crl / 012012033300 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0000323110331110-3232231213201033-3223100323220300-2313023310031130-0210100303322330-3100133100233321-0212122000312301-3123020001112103"></a>

<a id="canonical-0033121001000210-2201210202213210-1201123332211312-0332031133132223-3203031120231110-1020313012130021-0300022203022110-1212003302010311"></a>

## tenant property — crl / 012012033300 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0033301122223302-2203311202231001-1320233003022110-1203011001323022-0113320120023310-2100002020302330-3303323311230121-2033203202133223"></a>

## Next pages — crl / 012012033300 / 7

- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2302132330013031-1101202110332220-2320320230323133-3210130122131121-3303232022221020-2223302103232311-2233222021332233-1212230303313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233303022032030-2113202020113202-3101010320133112-2020203003233033-2200331221202001-3333100100312320-2200211213133130-2122130321101012"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl — no_crl / 332301211123 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl

<a id="canonical-0131113120331002-0010031001001200-0110003003011223-2120211010221030-1103213121000003-1333012311113002-1220312222011012-1300101230203200"></a>

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

<a id="canonical-0132113200112131-0313133111103222-0002133121201123-3031313321001320-0111311012332032-2222333200233120-3300230210321310-3233313012321133"></a>

## Direct properties — no_crl / 332301211123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112223012323211-3003300123333203-1111010222231023-0101223013013002-2033032212313122-1002033033303232-2110012011321100-0230222233001100"></a>

## Next pages — no_crl / 332301211123 / 4

- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0113322102331021-0102312331132203-1131012310302113-3003012200201322-0012303213001132-2323010231111133-2130300201322033-1031010133201001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332030232013021-0333113120132130-3232300332203002-3212300101211333-1301123010133333-3331302200233222-3320201213033310-0110113322131032"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca — trusted_ca / 012111021333 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca

<a id="canonical-1200001031022130-1320001220323302-0323102312131312-2330301113213313-0320302332203300-0223300300320000-1200110221322102-2031032201210113"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0230231023320011-1020032202212203-2000120110103022-1102022023013220-0331010221211103-2021113101032111-3201020213322121-1330213233013131"></a>

## Direct properties — trusted_ca / 012111021333 / 3

<a id="canonical-1103202323220033-2101211321111112-2121113333300023-2110222103032012-3203100213120033-3011113322000003-3020302310021231-0021210332111221"></a>

<a id="canonical-2133212200030023-2113233102210120-1303030231130132-1013320230222312-2102332221220020-0312220230101131-1020002122113102-3030220232333013"></a>

## name property — trusted_ca / 012111021333 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1303332203210303-2023211021000011-2331111123020320-0030312303133223-3221030101122311-1003211021323331-2003121100232102-3233200110013313"></a>

<a id="canonical-0000321133031331-0332331013111021-2102220010110202-2101300301032231-2003230232323212-2200323121130321-2000221111100311-2112311211132002"></a>

## namespace property — trusted_ca / 012111021333 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2013220022101112-0002121100302233-0112301331332211-2022211231011220-3303123211320132-2232012112110332-1302111223120003-2112111123321012"></a>

<a id="canonical-3120222230132320-2300122321101200-0130332113020310-1220323211333023-1001303212211023-1213201313133333-0223211220312221-1313033303120002"></a>

## tenant property — trusted_ca / 012111021333 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
