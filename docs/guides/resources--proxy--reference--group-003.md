---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-0301100002130013-1012123133131013-3110010232322231-3202301312002013-2010303021022220-1031101300302012-1302132033012002-0101003322000101"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value — secret_value / 322122331311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-2012032101103300-1323100203330222-0200112130223020-2013112200332301-3231132133123001-3133023101103033-0210133130020133-1331303100213203"></a>

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

<a id="canonical-1103131011133230-0102303100033311-0102120132201223-0113222233133301-3111121301211021-3231032021200123-3001232201312203-0232223002031320"></a>

## Direct properties — secret_value / 322122331311 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0131201303010331-2131112100321201-0111200122301000-2302302023001302-2012302030011111-1322110221100122-2010013001323123-0332022032332131): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-2212311010210312-3001323222221101-1032000012131230-2310303323220330-0312213122313330-3300131220211130-0311331021321313-0103011210120000): complete subsection reference.

<a id="canonical-0222202222302001-1120321230210103-2110323303302223-2100211220012210-2023122120300113-2013131333003213-2201300003001323-3210333101230300"></a>

## Next pages — secret_value / 322122331311 / 4

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0131201303010331-2131112100321201-0111200122301000-2302302023001302-2012302030011111-1322110221100122-2010013001323123-0332022032332131)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-2212311010210312-3001323222221101-1032000012131230-2310303323220330-0312213122313330-3300131220211130-0311331021321313-0103011210120000)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0131201303010331-2131112100321201-0111200122301000-2302302023001302-2012302030011111-1322110221100122-2010013001323123-0332022032332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210333232223313-3120223320112332-3322330212111323-1331132332213002-2320121112003332-1002310223122102-2021023021031331-2031320002230303"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 213302100130 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0032010031122032-0202012313311032-3100202003322221-2122221313301333-3133332213222030-2323212031323322-2233210310101122-2003321232101100"></a>

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

<a id="canonical-1301131230033112-2312130001010322-1200012032022113-3102223002221200-3120213030100222-3002030112312001-1200133213132333-3321233111212230"></a>

## Direct properties — blindfold_secret_info / 213302100130 / 3

<a id="canonical-3122102131003202-0121230121032320-2101111100300120-3021202322122002-1333211111312322-0012321301120022-1233121320011110-0231302011032010"></a>

<a id="canonical-3332033012023222-3131312033113213-0211221033221123-1213011011023310-1320332321333222-2133322102230331-1112311133331112-0230012122123023"></a>

## decryption_provider property — blindfold_secret_info / 213302100130 / 4

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

<a id="canonical-1232310100202231-3212032220133120-3011112001022022-2022302132000233-1330101323101220-0033010003030121-0123003202231103-3033322302202211"></a>

<a id="canonical-2332033332212312-3203133113230130-0210321220130313-0233131032222013-1201223020122103-2011333201101200-3103320233131323-1312323130200202"></a>

## location property — blindfold_secret_info / 213302100130 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-3021100321202232-0331330102230010-1313201122300013-0111131103301112-0332203231010312-2303100211322132-3210303323031223-0211030330312030"></a>

<a id="canonical-1223311301301201-3330121210110101-0033103120023010-3123320020130022-0001002133222000-1300222211220223-3123323122130121-2020321121221233"></a>

## store_provider property — blindfold_secret_info / 213302100130 / 6

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

<a id="canonical-1200213200110332-0323220310332102-2203323300201220-0033211123230211-0122312310311210-3023012211033122-3331011331200221-3112022212330331"></a>

## Next pages — blindfold_secret_info / 213302100130 / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2212311010210312-3001323222221101-1032000012131230-2310303323220330-0312213122313330-3300131220211130-0311331021321313-0103011210120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221022100311112-0333132111211322-3221330301332313-1030001021010000-1123321003232202-1213113100303011-3223312311123121-3031121211330132"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 321200032221 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3220130301232111-2023321103132200-2000201002201310-2120023030020110-3101302323310310-0301221131200022-3212310233210100-2113102203313331"></a>

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

<a id="canonical-1003213313131110-1313111232220300-1021333030210203-3030302011331312-0231310321012223-2113230130320100-3121133113113310-1222202212020203"></a>

## Direct properties — clear_secret_info / 321200032221 / 3

<a id="canonical-1010222013311110-3013312213202221-0010103320222100-2021001133331221-1212111221030030-3310111033102323-3131221122320202-0120213103031011"></a>

<a id="canonical-1220201031121222-2123223023202221-0110202220231022-1310300112131033-1232202321023313-0211222110130113-3313132211103000-3011013313212210"></a>

## provider_ref property — clear_secret_info / 321200032221 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2231312211132020-1312013002030110-2332023221121312-3133210232322231-0100332300110131-0033231133322021-0203323230101133-3000000102211231"></a>

<a id="canonical-3110002333003311-0122122230010301-3220011333000101-3131323322032002-1212022031003213-0201033322131132-0222032120112222-2222101320232201"></a>

## URL property — clear_secret_info / 321200032221 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-0200020123022103-3212112123101120-1012323313312323-2203300012211132-0230200101131221-2113000230133213-1321320130321302-1200221320203303"></a>

## Next pages — clear_secret_info / 321200032221 / 6

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303203110103322-1213031132130120-1013321001133330-2210022311333032-1102213010200303-3210311212311132-2103200121123010-3111020120110211"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add — request_headers_to_add / 001311203313 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add

<a id="canonical-3032001302313331-2033020223001322-1322001120103212-1113303122333121-3102203031330223-1320332022312233-2132031002333320-1113133130000102"></a>

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

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221102133333002-1211312331313302-0210030213301231-2121011222021310-3023311123213200-2010101223030231-0021323112202330-0031011003100331"></a>

## Direct properties — request_headers_to_add / 001311203313 / 3

<a id="canonical-3233301321211031-2013200220102310-1312112023102103-3011011010013121-1023013130332311-2223311313330303-1122333330323220-0311300031123102"></a>

<a id="canonical-1223113211100311-2210333300320312-0222211000202312-0203000331200022-0123033132002311-1100230223312302-2112022010103012-1203301010112132"></a>

## append property — request_headers_to_add / 001311203313 / 4

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

<a id="canonical-3120321211211202-2211101113213230-1322010210302122-3333113012021100-2032222311002333-0233100220320010-2312120201020121-1013232303321021"></a>

<a id="canonical-1301311311222023-1121123133200310-2321233010030213-0203203311020232-0133320011101223-3301202201130133-0231221010002100-1020223222132023"></a>

## name property — request_headers_to_add / 001311203313 / 5

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

- [secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133): complete subsection reference.

<a id="canonical-0032021220222300-1223123002201013-0310321012201013-1302001133331103-3103201101323013-0211331220000001-2221332013231033-1203223223303021"></a>

<a id="canonical-2112001223211213-0131000223203120-0213002101000101-1012022032210032-1132331131033120-1331313201113221-0000001203011303-2000030121120012"></a>

## value property — request_headers_to_add / 001311203313 / 6

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

<a id="canonical-1320002103111232-0200320100020322-3323102210023322-1003000102130121-0032133023111222-1223132111001310-2222011011001203-3231032301023332"></a>

## Next pages — request_headers_to_add / 001311203313 / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333233123300302-3312102313331323-2202131310101330-0032232201033131-0113021102230333-0233030232001321-3320333233310132-2212023111303231"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value — secret_value / 333032300011 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-1303010212333201-0312112222010130-3120112023110313-0111300333032303-1303110323331133-2112203122322323-2322020222113001-2310023120202323"></a>

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

<a id="canonical-2102103001030201-2320133032230020-2002301220122032-2012200021323132-0032110331110311-1020303021010120-2320310230032211-3231221120030132"></a>

## Direct properties — secret_value / 333032300011 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-3000332013332130-3203012020121100-3013131023233003-3000321012022221-1022211012121102-2100131301022002-0310322101102330-1121322120133011): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-3013211221333301-0102123101313322-2212001003211300-0333303021031220-2120231322223102-0022310211133032-1202121330320100-3120232322000131): complete subsection reference.

<a id="canonical-2333002200310302-0223230232001122-0223033012032001-0233112310221010-0013313311300302-3130133201000330-2103211132313000-2013323300122333"></a>

## Next pages — secret_value / 333032300011 / 4

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-3000332013332130-3203012020121100-3013131023233003-3000321012022221-1022211012121102-2100131301022002-0310322101102330-1121322120133011)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-3013211221333301-0102123101313322-2212001003211300-0333303021031220-2120231322223102-0022310211133032-1202121330320100-3120232322000131)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3000332013332130-3203012020121100-3013131023233003-3000321012022221-1022211012121102-2100131301022002-0310322101102330-1121322120133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302003021131021-2110312023131213-1220321220222212-0030212233020131-3232331002003000-3011220331233330-1031210000020113-2020032130232033"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 021220101331 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3132331131232130-2101121322323000-2133310112310031-1001212003223330-3321013212313120-3301331230321203-3223030323201323-3230003121001321"></a>

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

<a id="canonical-2311321122100221-2003233332003133-1101211130311213-2213233032211102-2231012021301221-2103210232001100-0100231130012311-3000203330212022"></a>

## Direct properties — blindfold_secret_info / 021220101331 / 3

<a id="canonical-3131320232011121-1103213112011320-1023012021310212-1102312211233320-1320313313330212-2300213220301013-1321123012033100-3113330233303332"></a>

<a id="canonical-2220321312311231-3231202032010310-0203333331300130-3220201120000300-3221030201302311-2232321311102022-3223010321031100-1122331213221033"></a>

## decryption_provider property — blindfold_secret_info / 021220101331 / 4

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

<a id="canonical-3330110111133132-3331010003313132-0103303201211200-3323213211210211-3110312122020332-0210213222032011-3311222113131121-0200220302123312"></a>

<a id="canonical-0020331012312313-1013310303332032-3122330133102031-0210330321201230-0020001220033013-3130000020331000-2303232012200120-1213023022112112"></a>

## location property — blindfold_secret_info / 021220101331 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-2320033032000231-2212113200133310-0302322232232002-0312301120311223-2020310102030123-1302223033112003-1113331311021233-1300321130020012"></a>

<a id="canonical-1212030113133111-0331002131300001-0011010103220032-0001002222203122-3320211310312102-0213310020203033-3213020310321322-2031232101003332"></a>

## store_provider property — blindfold_secret_info / 021220101331 / 6

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

<a id="canonical-0300200300133032-0221131223122311-3001023201231110-2101010013100012-3011322101022323-1321223122020323-1230012221032132-2122101232221211"></a>

## Next pages — blindfold_secret_info / 021220101331 / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3013211221333301-0102123101313322-2212001003211300-0333303021031220-2120231322223102-0022310211133032-1202121330320100-3120232322000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223130110220203-1231330123222313-2021311301030130-3221122133303011-1221333310111110-1023232103122002-0102013013132133-0233011223221220"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 313021310121 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0310112231113133-2020130021131213-2322210232013012-0221312302133303-1121013010113222-2030022331030323-1131330111233220-2003023313023022"></a>

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

<a id="canonical-0300132000333312-3002003200023201-3102022301002312-0022023123123331-3122033111301312-3203032312333011-2020320102031330-0102313031332100"></a>

## Direct properties — clear_secret_info / 313021310121 / 3

<a id="canonical-2303020312220032-2030321103210031-3221323111000222-1111301302223122-3120332101103212-1331311203033300-2233030323311112-1322300302210013"></a>

<a id="canonical-2230011301203221-1011311320131130-0031020110223320-3221030301121132-0302302111120003-1301220112330221-2323012333213203-3212011103223320"></a>

## provider_ref property — clear_secret_info / 313021310121 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2103001302111231-3120010233233213-0010232311133301-1323102122312231-1032132331201232-1332231320131312-3101133301323112-1330013111301221"></a>

<a id="canonical-1012121031322111-0210031011013002-3021221221301120-0321132230231100-2103031310321003-1120010202200000-0113010013022101-3131032103222220"></a>

## URL property — clear_secret_info / 313021310121 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-1321222130301321-3130121013321021-2300233011202301-0121021033111200-0231120323210220-1311012212012212-2012003321021201-1231111133130123"></a>

## Next pages — clear_secret_info / 313021310121 / 6

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103323303202300-2023231123200223-0013330032130102-3112330202003312-3100110203111002-2003202302333033-0102001022232313-0312222133110232"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add — response_cookies_to_add / 130310100033 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add

<a id="canonical-0131201032122021-3010202001203131-3333211331121103-2312023020312012-2001020033333010-1212022102233300-1323220210123323-0022013312202311"></a>

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

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320001321232232-1100321331310131-3330303202000131-3100021200222322-0023021302020102-3230321223023031-2211312300211013-2012322222132330"></a>

## Direct properties — response_cookies_to_add / 130310100033 / 3

<a id="canonical-3312132021303220-3012010312122123-3121230220022103-0313223031003232-0020013312033112-3122201031312331-3201321203300333-3302021323313103"></a>

<a id="canonical-3133030113202000-1121113003112012-3122331100010231-2220302200133313-2103122010131333-1320012002003101-1311122211032012-3221210331032013"></a>

## add_domain property — response_cookies_to_add / 130310100033 / 4

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

<a id="canonical-1332221222231313-2221030223213213-3013110020203232-0100103301222001-1301120023222130-2203303112110112-2223023212202101-1200331131302302"></a>

<a id="canonical-2222000320321313-1002133100103101-2132032102332222-1133013202113221-2111202330213312-0100100003221221-2231301023323020-3110220222300100"></a>

## add_expiry property — response_cookies_to_add / 130310100033 / 5

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

- [add_httponly](resources--proxy--reference--group-003.md#canonical-0322010232101122-0220323003101213-3323131000321030-0300010230213023-1331022000133201-3201200302000110-3021322321022201-1202123022311221): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-003.md#canonical-2212200310330103-1032232331211021-2113122202302111-3011332310020302-0322000231121330-2200301223120322-0302331030022301-3123112112020102): complete subsection reference.

<a id="canonical-0321121023130330-0101333112232023-2302301202323130-2333210020122303-1010111313310300-1133230111330113-0033203203302001-0200211333102031"></a>

<a id="canonical-0302321303033020-0222331323012020-3210121020101303-2102230110212211-3132321103331010-0313131333101020-0310201331032112-1300003103122213"></a>

## add_path property — response_cookies_to_add / 130310100033 / 6

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

- [add_secure](resources--proxy--reference--group-003.md#canonical-0001010013121313-3303320320020010-2101110003333123-2232132113101033-2302203033311322-1110222122010213-2232121111202200-1322020110121301): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-003.md#canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-003.md#canonical-1102103003333103-0031133300111212-2030023201032333-0123220310012322-1300131203233230-3212111012333230-3013201033020311-2032132232031022): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-003.md#canonical-3321312012221102-2011213032302220-0111223111010211-3032230011320113-3310202223120232-0132133130131112-0130021113003300-1303222003311323): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-003.md#canonical-2200202100213111-0132121020222001-1203333223222122-0220200332022213-2003210022021220-3311003320233333-2233131101110012-0103013203133013): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-003.md#canonical-3010103012002231-1230233022121112-0132103120103221-2101131130231201-0332012312311031-1022022021313131-2313101101103111-0213323330103331): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-003.md#canonical-2033011031112123-1001200301010122-2102002131220130-2300002033112303-0122231133323312-0213001310121011-1122322312200210-0300020030310111): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-003.md#canonical-1211213001210301-2322033012322331-1103132223010023-0321332020332102-3330322330101330-2133023313300333-1213101111220203-1310233003200121): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-003.md#canonical-2213321311302100-3220123332103320-0211203033210101-2323003122302332-3320323021310000-3331321200102210-2303200132132310-0300032002002013): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-003.md#canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200): complete subsection reference.

<a id="canonical-0330131020302331-0032121212210102-0012320311313033-1210032030322203-3021012300033202-2311021312302011-0010110133021010-0220130013221313"></a>

<a id="canonical-1232020301203221-2132332212312030-1321013120123021-3313230100333201-2323103132230110-3211130033320221-0133112300312300-3311202021231211"></a>

## max_age_value property — response_cookies_to_add / 130310100033 / 7

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

<a id="canonical-1133300320123320-1022133111203132-2131111313122011-1013023322113102-3122111001012002-3312321033010320-0323011201131113-1011312133323020"></a>

<a id="canonical-2113321303123322-3010330311210210-3310200031122232-2302322332233021-2031313331302301-3133210222311303-3320113322230123-3232222020221331"></a>

## name property — response_cookies_to_add / 130310100033 / 8

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

<a id="canonical-2310321330021330-0020320321130101-3223002321200100-2031300320233031-2202210323210322-3113223122132121-0032323111202323-1123201020300312"></a>

<a id="canonical-1221002321122030-1101200020010230-1002130333102030-2213330131020321-3012122200210010-1110220211110010-1012232032231201-3203011311030201"></a>

## overwrite property — response_cookies_to_add / 130310100033 / 9

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

- [samesite_lax](resources--proxy--reference--group-003.md#canonical-0130331210231000-1311132112001330-2100133002022130-0003332021222230-3032210131202201-1203330013213010-2013000302222323-1131300123132013): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-003.md#canonical-0211210313001301-3213002022100221-1031002001112022-0232231100321333-2030212220310303-2320203300312102-0221301031213023-0102223331033320): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-003.md#canonical-3211313331201302-2103100000320112-3123223001012122-1003210203003032-2000002323123311-1202010022011332-0101100012211022-0002011001123110): complete subsection reference.

- [secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011): complete subsection reference.

<a id="canonical-1323110321001120-2110022133101211-3102123123300333-3111123130313331-2023020331210132-0230010210203323-2210322131222213-1022112330132010"></a>

<a id="canonical-1221233312301303-0232000300121321-0133330302233322-3013313022230033-1113102002112213-1212221102223103-3230010210332130-3001100032023112"></a>

## value property — response_cookies_to_add / 130310100033 / 10

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

<a id="canonical-3210300321312303-0200312012221211-3011101033313001-0002201313220032-3033202330010123-0103030233022100-1200130021003033-2103003001331002"></a>

## Next pages — response_cookies_to_add / 130310100033 / 11

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-003.md#canonical-0322010232101122-0220323003101213-3323131000321030-0300010230213023-1331022000133201-3201200302000110-3021322321022201-1202123022311221)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-003.md#canonical-2212200310330103-1032232331211021-2113122202302111-3011332310020302-0322000231121330-2200301223120322-0302331030022301-3123112112020102)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-003.md#canonical-0001010013121313-3303320320020010-2101110003333123-2232132113101033-2302203033311322-1110222122010213-2232121111202200-1322020110121301)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-003.md#canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-003.md#canonical-1102103003333103-0031133300111212-2030023201032333-0123220310012322-1300131203233230-3212111012333230-3013201033020311-2032132232031022)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-003.md#canonical-3321312012221102-2011213032302220-0111223111010211-3032230011320113-3310202223120232-0132133130131112-0130021113003300-1303222003311323)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-003.md#canonical-2200202100213111-0132121020222001-1203333223222122-0220200332022213-2003210022021220-3311003320233333-2233131101110012-0103013203133013)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-003.md#canonical-3010103012002231-1230233022121112-0132103120103221-2101131130231201-0332012312311031-1022022021313131-2313101101103111-0213323330103331)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-003.md#canonical-2033011031112123-1001200301010122-2102002131220130-2300002033112303-0122231133323312-0213001310121011-1122322312200210-0300020030310111)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-003.md#canonical-1211213001210301-2322033012322331-1103132223010023-0321332020332102-3330322330101330-2133023313300333-1213101111220203-1310233003200121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-003.md#canonical-2213321311302100-3220123332103320-0211203033210101-2323003122302332-3320323021310000-3331321200102210-2303200132132310-0300032002002013)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-003.md#canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-003.md#canonical-0130331210231000-1311132112001330-2100133002022130-0003332021222230-3032210131202201-1203330013213010-2013000302222323-1131300123132013)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-003.md#canonical-0211210313001301-3213002022100221-1031002001112022-0232231100321333-2030212220310303-2320203300312102-0221301031213023-0102223331033320)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-003.md#canonical-3211313331201302-2103100000320112-3123223001012122-1003210203003032-2000002323123311-1202010022011332-0101100012211022-0002011001123110)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0322010232101122-0220323003101213-3323131000321030-0300010230213023-1331022000133201-3201200302000110-3021322321022201-1202123022311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113131210203222-0203323301202220-2320001222131101-3311311332021231-1013331311113020-3230200232311012-2031320102313203-0222210013311300"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly — add_httponly / 313203221003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-3331321132111110-1303201230300302-0133132303130120-1301301310003332-3030312203032331-1230323220232211-3201201113312202-3302003011000303"></a>

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

<a id="canonical-2130023200323002-2123120223320313-3113120022333303-0313300000021032-0030233220332123-1120303201120210-2031121320323312-3220330202213130"></a>

## Direct properties — add_httponly / 313203221003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223331022301230-2031202303320200-2101010312130330-1001320003330211-3133033133001212-2212132301223222-0120003033112330-0132020032010202"></a>

## Next pages — add_httponly / 313203221003 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2212200310330103-1032232331211021-2113122202302111-3011332310020302-0322000231121330-2200301223120322-0302331030022301-3123112112020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203222100001222-1311013302133112-3233231312003020-2202231120112113-3032113112222232-1311233022110023-0131203322311121-2213101121303233"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned — add_partitioned / 330031133313 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-1022300231122233-2300033323102210-3210210032322033-0103000321223020-0003332020021001-2033302332220231-1233303333300203-3330021000233031"></a>

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

<a id="canonical-3000211330113020-1331010133121200-1201112032022120-3223311003212210-2001223321212001-3122333123230332-0020033133232300-1120012201102001"></a>

## Direct properties — add_partitioned / 330031133313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121223111112022-1122000300110133-1223011031330002-0320001221013013-2023100332130003-1002101112030031-2221033302222322-1320012100201102"></a>

## Next pages — add_partitioned / 330031133313 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0001010013121313-3303320320020010-2101110003333123-2232132113101033-2302203033311322-1110222122010213-2232121111202200-1322020110121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102220012300212-3102330230222320-1002223202010000-0010132332231301-2321123111230312-3310123302332132-3212210133302301-3110313303313032"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure — add_secure / 002033022032 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-0320101201012130-3320000121222322-0031313133130131-0102313220030212-2221321203303301-2221132310233213-0120211032210332-0310102022311012"></a>

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

<a id="canonical-3201111100101322-0212233110110100-0011323223002303-0303012103321111-1231200032310333-3110211301310333-3331012332301233-3110232303222103"></a>

## Direct properties — add_secure / 002033022032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130132002200313-3133200100310011-3113330122021000-3130312111131302-3113010233110001-2230302123020212-1220111001323001-0210321303120113"></a>

## Next pages — add_secure / 002033022032 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320211333000321-0300000113012212-0021102021011011-3012220103233210-2001232030312320-3202312030021323-1130210002001323-1202323200020013"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain — ignore_domain / 233231100012 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1102102133022220-2301331301132113-2122123132110202-0101023032010320-2101011033101130-1320110001122221-2033332310203100-2131112031122131"></a>

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

<a id="canonical-3121001212131310-2303132231023202-2102013311120103-0013033021131332-3211113231020031-2220121212223320-2322312020002023-1033310102002303"></a>

## Direct properties — ignore_domain / 233231100012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012233122012001-1331001031010303-2120230222023312-0103010032333131-3133321122013033-3103021003301112-1311310312323331-0321321312333220"></a>

## Next pages — ignore_domain / 233231100012 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1102103003333103-0031133300111212-2030023201032333-0123220310012322-1300131203233230-3212111012333230-3013201033020311-2032132232031022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200031023011231-2333303010310113-0301130231131211-3133032031320031-2212132231321132-1031320311321123-0333300111320113-2112210300120323"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 211130233010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-1311232102322203-1001312301000031-2020313310322201-3002221320223020-0130032101211000-2010202011303310-3121132222001120-2120230312230221"></a>

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

<a id="canonical-1312203103003333-1301123302301330-0031231032031331-0013321112211133-1102121202223333-0023301323302123-3323202122332000-3110311012021311"></a>

## Direct properties — ignore_expiry / 211130233010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112003132313032-3030031301110102-0012020333103220-1330320210210123-2333001011313310-3312210223332233-2321300331203022-0121002232322233"></a>

## Next pages — ignore_expiry / 211130233010 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3321312012221102-2011213032302220-0111223111010211-3032230011320113-3310202223120232-0132133130131112-0130021113003300-1303222003311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231213101321030-1100323020000102-1033121331010330-2301302223323212-0113300320102330-0311320123321202-0210313331011011-2213311311001312"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 112210311312 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-0331322312001122-3332031002303031-1012002121213213-0302133020212111-1221110222013022-2002032021221011-2211023131100310-1221102211323200"></a>

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

<a id="canonical-0320002202331321-2300103033312020-0110010030002121-1202130302022223-3331003302332300-2221300333231002-1122032022033300-1201323012122121"></a>

## Direct properties — ignore_httponly / 112210311312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131230222112022-2212221023312113-2310112021023121-0130133321121031-1322231302333330-1230112302020211-1033021033020210-1112311010112323"></a>

## Next pages — ignore_httponly / 112210311312 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2200202100213111-0132121020222001-1203333223222122-0220200332022213-2003210022021220-3311003320233333-2233131101110012-0103013203133013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002013320221013-3113203210313111-2330122313130303-2113211222300330-2220232023323011-3011032322221211-1302333220302210-2232100132200011"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 322012210021 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0021033021202111-0202011120112302-0320213102111313-3001303133012031-2331211301200103-0311322321202030-1120312230312322-3321202312011030"></a>

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

<a id="canonical-1320133100100000-3132132322102020-1321022123120321-3302103111132303-1011113023113321-3313100023132012-3121320010220312-0201210332200231"></a>

## Direct properties — ignore_max_age / 322012210021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300011201031132-0220012301331211-3120302110100102-1301303301231101-2001110333102311-3130223103230120-1120312232301132-0023212223332202"></a>

## Next pages — ignore_max_age / 322012210021 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3010103012002231-1230233022121112-0132103120103221-2101131130231201-0332012312311031-1022022021313131-2313101101103111-0213323330103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322320200310200-2103333313000220-3013000323011320-1200202203203310-3230110222210232-3212232232002003-2200120020121203-3211130130103101"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 013132003030 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2330223031212303-2210033121111302-2330203200131033-2013001021133011-2211132331013130-1311213321021200-0303332103030133-1220121321013321"></a>

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

<a id="canonical-2130200310100230-0310320130000012-0122231312203222-3213312020331231-3103330312210331-3031110323200301-2130132021311131-1231121101122122"></a>

## Direct properties — ignore_partitioned / 013132003030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012323331013201-0220222000020113-2232033121110022-1033032131220200-3222133322321203-1313302202221203-2231220102300021-0022231200011031"></a>

## Next pages — ignore_partitioned / 013132003030 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2033011031112123-1001200301010122-2102002131220130-2300002033112303-0122231133323312-0213001310121011-1122322312200210-0300020030310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011323221133222-0023102000102300-1012021133100200-1031323301300030-3232220212301132-1300011300320210-0302002332100022-1123032102010030"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path — ignore_path / 232130101112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-2210222201222220-3322123223010022-2320221122133010-0022221332023210-1221230022310211-3313301302313232-1011223331131213-2120030132122021"></a>

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

<a id="canonical-3220220332200201-1312323132023100-0131330103010002-1210032012220103-2222001203220002-2311322312033022-3320223132210131-3203002131211021"></a>

## Direct properties — ignore_path / 232130101112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032112133013230-0101303102023223-0000211031032013-0331012212321303-0321303013113332-0023331322031020-1211212032333210-1202112103000032"></a>

## Next pages — ignore_path / 232130101112 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1211213001210301-2322033012322331-1103132223010023-0321332020332102-3330322330101330-2133023313300333-1213101111220203-1310233003200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221230023123310-1100101002233120-2131230313213312-0103320000233221-3133002213033002-3300330110121233-0101232233212121-2133022003203000"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 212123311133 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-1132103301333331-2100322212221230-3303203330321333-3232023010201303-1230300230322233-2000300022011223-1302133020123210-3201302300202320"></a>

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

<a id="canonical-3211133233100121-0001012230113220-1111030103113203-3012333231300212-0310032001202330-3211322200332003-2310132111120331-2131000212221311"></a>

## Direct properties — ignore_samesite / 212123311133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032122201220002-1132322302022112-0223222031000323-2300200121033331-2123312202133221-1202220121300231-3310102113302103-2113101201210231"></a>

## Next pages — ignore_samesite / 212123311133 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2213321311302100-3220123332103320-0211203033210101-2323003122302332-3320323021310000-3331321200102210-2303200132132310-0300032002002013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323323312000312-1230023111201032-0332002032121300-3200220322223012-1310102012133103-0002332131303110-0303131311201232-0031222022333320"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure — ignore_secure / 001013103201 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-3022331022100322-2320213332033311-2211032211200223-1201013302021000-2200012201012321-3220210333333301-0112310322101133-3121201211011300"></a>

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

<a id="canonical-1113223310022010-1133021102002202-2013132010211213-3003211130020110-0231120033111111-3000113111020110-2330103222031200-3233211330122300"></a>

## Direct properties — ignore_secure / 001013103201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102130131203313-1123111222221312-1111321223130311-3323210131110033-1212033203030100-2002202213321030-2230000302001213-3220330010003103"></a>

## Next pages — ignore_secure / 001013103201 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033231032031112-2000313230320020-0030013230212331-2013332010331122-1003113032220003-0303021332033333-1133202201331200-3123312301100012"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value — ignore_value / 320133012213 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-0201122322233323-0301231210021300-0301220222232332-3310031130132213-3103120322000122-2021323310123102-0303333313212213-0303022313321112"></a>

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

<a id="canonical-1303320320101222-3230120003112010-3011100112331313-1111312132110123-2032302202220210-2310331303213131-3330002100233010-1013121013013133"></a>

## Direct properties — ignore_value / 320133012213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330320120023110-2131302312202001-3323320130132100-0313210310303200-2132220302201313-0020222102313211-2131000011010301-0221321022202313"></a>

## Next pages — ignore_value / 320133012213 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0130331210231000-1311132112001330-2100133002022130-0003332021222230-3032210131202201-1203330013213010-2013000302222323-1131300123132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000121101011231-2113000022202221-0111010013122201-2102003013100203-0231212320233220-1331303112333003-1232000221212233-1010111121010030"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax — samesite_lax / 002323000010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-0201011101112200-2113213311121101-2133131131202020-1300100130302132-3303122121112022-2311211012332301-0032301221111112-2301033133023030"></a>

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

<a id="canonical-1113322020332121-0210132012013202-1133211230323300-2001313302322111-0213322112000003-1213321112301312-1110321201010221-3031323110331203"></a>

## Direct properties — samesite_lax / 002323000010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331130302300233-0011012301222310-1122301103102000-2202133030203020-3101330322211330-3002312033002212-1300203111002333-0210123122201202"></a>

## Next pages — samesite_lax / 002323000010 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0211210313001301-3213002022100221-1031002001112022-0232231100321333-2030212220310303-2320203300312102-0221301031213023-0102223331033320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010111303013030-2311222330300221-2230001131332001-2311322203233122-3201003022021130-1333303103323220-3103201212222123-2331320120001101"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none — samesite_none / 220330233013 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-3222132023011210-1110103201003321-0222320123101312-2322113123302302-0222031222120111-0203301211003210-0100033103201312-2102122012300110"></a>

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

<a id="canonical-0303321313132331-3031321313132222-2233232322003121-0232203121211013-2200100032333311-3001312110033310-2210013000121133-3331302333011310"></a>

## Direct properties — samesite_none / 220330233013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220120122222203-0311013010301020-1020113213303112-1200122012313201-0223130232220010-1121232310332331-3210020332211333-2103010221032022"></a>

## Next pages — samesite_none / 220330233013 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3211313331201302-2103100000320112-3123223001012122-1003210203003032-2000002323123311-1202010022011332-0101100012211022-0002011001123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130021212211310-1121222311302020-0312103302332301-1220312131231332-2201123110001030-2131012012031223-2110202101030120-1231300203011133"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict — samesite_strict / 322302310320 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-0132131101102101-1101030130221002-0012310102200113-0330011100223211-2320212321000130-3010002330101332-0203102331210211-1203120123031020"></a>

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

<a id="canonical-3112133301330120-1022321310310202-1101222101011303-0031113110131031-1313132200021030-3303203011103000-1032022202102020-2013021010010321"></a>

## Direct properties — samesite_strict / 322302310320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220000112200311-1201033022031001-1203030033302201-2201330212300033-3331012300100031-1101211102133313-0030333212013000-2321012330302101"></a>

## Next pages — samesite_strict / 322302310320 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013001232221300-2222012213102002-0332101103333022-3033030233332311-2121211233100031-2003033333213011-2123312231101121-3220220222311111"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value — secret_value / 031320103233 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-0112032002023223-3312231333031222-3133332123303123-1030201130032211-1213023132333002-2123122111221122-3010203201213311-2001302223102101"></a>

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

<a id="canonical-3202023013320232-3032211002001203-1031311001223110-1030201022021222-0000200021222020-3031022103222001-0013201001231211-3000330231303023"></a>

## Direct properties — secret_value / 031320103233 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0001030020322030-2121031322213202-1300302123121010-1101210232321230-3201101132303330-0120332302103200-3000203021310202-2200103022130323): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-2303122020111321-2122001200332131-0132233322023201-3121112022133121-1101103220330003-0032120233223310-0031032132001122-1211131330323120): complete subsection reference.

<a id="canonical-0310003022213130-0003212221010200-1102233211100232-3030202130100330-2223320101103203-3032132320000230-3103002132313200-1321203121313322"></a>

## Next pages — secret_value / 031320103233 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0001030020322030-2121031322213202-1300302123121010-1101210232321230-3201101132303330-0120332302103200-3000203021310202-2200103022130323)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-2303122020111321-2122001200332131-0132233322023201-3121112022133121-1101103220330003-0032120233223310-0031032132001122-1211131330323120)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0001030020322030-2121031322213202-1300302123121010-1101210232321230-3201101132303330-0120332302103200-3000203021310202-2200103022130323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110013012131230-2102212312012101-1102032130022103-1301230001010033-0320210101312303-3102323322020100-2231102113233232-1212131132322232"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 030103223010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0023033123100330-3111121320022312-2103302222000220-3230123020213020-2132310322100001-1113111222123103-2301322031030011-0132203213012112"></a>

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

<a id="canonical-3101300213220333-3100132301323002-2311132012023230-1122233303332332-3103023000113310-3020201331331012-0011312032303101-2000121223333323"></a>

## Direct properties — blindfold_secret_info / 030103223010 / 3

<a id="canonical-1100232022331223-0002321011302200-3212300113210210-2123003310033223-1133303133133012-0032100100002330-3032332312113113-2331020320030012"></a>

<a id="canonical-2000311130200222-3130232012221302-3210222220032112-1123033303202132-0012130012110110-1312023133112220-2312321112101323-1002210111210211"></a>

## decryption_provider property — blindfold_secret_info / 030103223010 / 4

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

<a id="canonical-3333233213211113-2002103012133211-2112200233100331-2102103113113031-2122022310320210-0301332001310023-3333202003101233-2322210220020002"></a>

<a id="canonical-2320123000322021-0133303323331132-3030113321210330-1313313210113030-3033133011332222-3200131311131012-1102332223330323-0022132311103030"></a>

## location property — blindfold_secret_info / 030103223010 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-0222013012303013-0301332323122310-2031323023210102-2000320302310113-1123103121011001-1013103012101122-1120023203311203-0310131330332233"></a>

<a id="canonical-3201213333100131-2032220322022000-1122021120233033-1320202011212101-1010201221313313-0010223310212203-2212003223230310-2322031100133011"></a>

## store_provider property — blindfold_secret_info / 030103223010 / 6

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

<a id="canonical-1332303230321300-3300323301010230-1131223322132003-3032323231232133-0332103312032200-0123002102320203-0003231133220011-1033120103103101"></a>

## Next pages — blindfold_secret_info / 030103223010 / 7

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2303122020111321-2122001200332131-0132233322023201-3121112022133121-1101103220330003-0032120233223310-0031032132001122-1211131330323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222200002100310-2103013032001223-3033000222221301-3010103110212322-3213112300312013-3003313101022312-0000323020000030-0103110121210331"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 121000101203 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0330322112003110-1202120010330112-2200021201333301-3310233132223113-2223220313310003-2003101132010233-2123121333332100-0210100231002100"></a>

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

<a id="canonical-0000120123112233-2002311020222323-1211221031113012-0221031000233131-1221300333322303-3133303200000321-1001200101323131-3202312203121130"></a>

## Direct properties — clear_secret_info / 121000101203 / 3

<a id="canonical-2211032110332013-0032012131230101-2210320323232101-3230000221302302-0200303333301120-0332313103101101-0032132220100012-0112030031312133"></a>

<a id="canonical-1333211301213233-0301233003133011-3032201331020102-2321021320010032-0231222020001321-3011320331032111-1113010310030002-1103222032103211"></a>

## provider_ref property — clear_secret_info / 121000101203 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0102200001102300-2020103012130203-3330122112101200-1100011133120231-3020110320030213-1310202000033023-1002022133102210-2222123313223031"></a>

<a id="canonical-0231103120122222-3320211111011000-0121003313222320-1220310323132131-2133312221131130-1210223010132201-1030201230111132-1001230130300112"></a>

## URL property — clear_secret_info / 121000101203 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-2301001331101213-1223123010102011-1330213113021103-0332021131010331-3111023211333110-2201102001233233-1321210102300131-2211031322101011"></a>

## Next pages — clear_secret_info / 121000101203 / 6

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130330001312322-0031202212203001-2120311331121220-1332023333203112-0023131121122102-1230302300131232-0333322210320003-2301223312010230"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add — response_headers_to_add / 200332122220 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add

<a id="canonical-1133210313300331-3002323303321202-2131123113011113-0110002013030120-2100112013231330-3113313302020303-0332122111010121-0110203300323321"></a>

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

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320213333032212-2300002332303231-3221311022131101-0132112201311002-0220303231120103-1202113113300102-2100122133210130-1223032323033210"></a>

## Direct properties — response_headers_to_add / 200332122220 / 3

<a id="canonical-3111120033103020-2333023022131332-0111132031133213-2303123230121321-3132012311132323-3311330101020022-2012302330232330-0100301302203310"></a>

<a id="canonical-3033010223122332-1101301011301102-0312112220132323-3230021022011310-2203202113011103-0020030123320310-0232333322100112-1001121320300102"></a>

## append property — response_headers_to_add / 200332122220 / 4

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

<a id="canonical-3001231000010331-3331012131330110-0332113320031220-2330222222303001-0022010023001213-0102312320332223-0323131223301203-2210101131130103"></a>

<a id="canonical-3232232033132333-2301000102122313-2121233012102300-3022223211103203-3310210001310132-2130213302001330-2223003312120111-1210122100133133"></a>

## name property — response_headers_to_add / 200332122220 / 5

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

- [secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032): complete subsection reference.

<a id="canonical-2130101231322032-0321232233331222-2220020211122113-0001102323020020-2120300000231113-3220020011231211-1333133030213011-1003033130310120"></a>

<a id="canonical-1131203103221003-2322302322233221-3030233201130330-0223323330123301-3120223313002030-3300120133310010-0220330132013101-3123220121132102"></a>

## value property — response_headers_to_add / 200332122220 / 6

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

<a id="canonical-0002203120000331-3302102130200300-2333020112331103-3203110331330300-1322313011001030-2132012130333223-1211333002010111-1323010323012220"></a>

## Next pages — response_headers_to_add / 200332122220 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213133310300322-1023120310010022-1101122202003010-3321011202003202-1300023130232200-1333000101302321-2020231313031201-0131211301020103"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value — secret_value / 023003023013 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-2100222111320221-1333312123133233-2303301010221011-1223021003231021-3330021032021021-0010112103110323-0313231210122311-0023011233310133"></a>

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

<a id="canonical-2231332312322110-3021121313211223-1212333320121333-2312121322332033-2223230021310223-1221022201113023-3311331321212120-2213112121301320"></a>

## Direct properties — secret_value / 023003023013 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0212110001223211-0300000202210110-0301131022002003-1312113223132130-1302230123003220-0330333230132100-3122323210013103-0311023230310213): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-3323333001111203-3112003312110201-0321323211220001-0332310021002201-0100011123330111-1221223210031130-0120203031202212-1132120031223013): complete subsection reference.

<a id="canonical-1313123013330323-2301021201323123-1013003301110313-0333301323000320-3132232210210023-2310233223133013-0233323023330102-0231302012323112"></a>

## Next pages — secret_value / 023003023013 / 4

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-0212110001223211-0300000202210110-0301131022002003-1312113223132130-1302230123003220-0330333230132100-3122323210013103-0311023230310213)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-3323333001111203-3112003312110201-0321323211220001-0332310021002201-0100011123330111-1221223210031130-0120203031202212-1132120031223013)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0212110001223211-0300000202210110-0301131022002003-1312113223132130-1302230123003220-0330333230132100-3122323210013103-0311023230310213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302300202323321-0113031331010303-3231121322112031-3000303323333122-0120110023230101-0223323222101300-2231330212222012-0003231321203303"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 212230333320 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2003233321123031-3012301212203311-3033330033123331-1000030311233000-0313313133330030-2222322321203102-3130123020132011-2221020002303113"></a>

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

<a id="canonical-1220013003112321-3313311122212200-2302223320022011-1313010120230101-0322101101310133-0023301113311222-1130212213120331-1013211330231132"></a>

## Direct properties — blindfold_secret_info / 212230333320 / 3

<a id="canonical-3013113120323133-2232022010013102-0200121231020022-3003020233321311-1212013322320122-2011011233000203-3300201022001100-0133003123032202"></a>

<a id="canonical-3133210021300330-1101011220200303-0132302022313101-1321132211313323-3231230302023001-1111011332200100-0332111311031120-0332310012121113"></a>

## decryption_provider property — blindfold_secret_info / 212230333320 / 4

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

<a id="canonical-1230011002021323-1103233102112000-0101031123023023-3202111321330231-3200321302101122-3333333111101203-0203203031000313-3102210211100020"></a>

<a id="canonical-1000223123331110-3223113032222220-0010331112302312-2233212121120102-3003020111230322-1310010303210112-3111231021103010-1301111331023123"></a>

## location property — blindfold_secret_info / 212230333320 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-3201112332231131-3123113303111200-3313233112021133-3113322012010121-3003313331012212-1222221310001332-0321101100123302-0213112131103232"></a>

<a id="canonical-0120333123031203-0022021311322200-2233132302200312-0201013002020003-1213131311201130-1201002021020111-0313230323130113-3313213002302003"></a>

## store_provider property — blindfold_secret_info / 212230333320 / 6

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

<a id="canonical-3020332222001120-3001001220102112-2013030320333222-2312030303333031-2213203330302120-3212010211012012-2122201101220002-3301213000131132"></a>

## Next pages — blindfold_secret_info / 212230333320 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3323333001111203-3112003312110201-0321323211220001-0332310021002201-0100011123330111-1221223210031130-0120203031202212-1132120031223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133220023311320-2030120300123121-0123330331211020-1311322230220211-1321333222101121-2032010131221200-2121020120311303-3323102233310222"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 022211102323 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-0110200032010011-3011112021020002-2303221132212003-1221233110311022-1311303233301030-3013130311231121-3333213331121100-1000331331113310)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0332223311302233-3302313000111223-0311223120003201-2210033322331222-3223130311331012-3133110203231213-3332110031021320-2132111112303312"></a>

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

<a id="canonical-1112303202320100-1210201102230323-2310033301131330-1302101211110322-2123313333022323-0022323003330111-3212000333000111-1211001121121330"></a>

## Direct properties — clear_secret_info / 022211102323 / 3

<a id="canonical-2333011103100102-3200121312301022-2321130001333031-3003023032001312-2010031020021322-0000330111203320-3210130211001302-1233032332121232"></a>

<a id="canonical-1212222113221211-0000103311312213-0212103320233131-0020221303321023-1102232101230230-0331103213000133-0103032232030300-0321302100313032"></a>

## provider_ref property — clear_secret_info / 022211102323 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3321231202133103-1023233330103133-0022321012110012-1100120213320032-2333320000202003-1011123103232101-1333112233100113-3123031233200223"></a>

<a id="canonical-3112111211012113-1223032001211130-3112200211013000-3201120123203122-2312112012232302-1310103012110330-3220330123110130-0311221210302210"></a>

## URL property — clear_secret_info / 022211102323 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-3303301232032113-2322230200201301-0233002022303122-0200121220211222-1103333213223300-0211112303313000-3330221020013121-2321130133211121"></a>

## Next pages — clear_secret_info / 022211102323 / 6

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-3331330121302201-2301220332313111-2121202131221201-1131130102221212-0220032101202321-3332231033112200-0010210333133002-3010101213032032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120303020020122-1213332003323030-1202231231001321-2331321222032110-0101302021121003-3221011033200322-2232031323211130-0000122002122120"></a>

## dynamic_proxy.https_proxy.tls_params — tls_params / 122033210001 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- dynamic_proxy.https_proxy.tls_params

<a id="canonical-0201002202231223-2330122223021330-3010321311130202-2221331101011331-0031101033333202-2313100200131213-1302213121122201-3203001001203013"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122132001011312-2231320310130023-1032033012312211-2220012010302200-1113132302322202-1311031220131223-0000011011232003-0022110313110131"></a>

## Direct properties — tls_params / 122033210001 / 3

- [no_mtls](resources--proxy--reference--group-003.md#canonical-3221310223000132-2023121223223310-1231213203103323-2323213320322111-2110230332300023-2002201030213122-0021303300301232-3112120032110030): complete subsection reference.

- [tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021): complete subsection reference.

- [tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130): complete subsection reference.

- [use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032): complete subsection reference.

<a id="canonical-2133122321130212-3322113020201110-3030233230333013-0201023111020213-3233110110321302-1102213313110231-0222231020003131-2102002331132231"></a>

## Next pages — tls_params / 122033210001 / 4

- [dynamic_proxy.https_proxy.tls_params.no_mtls](resources--proxy--reference--group-003.md#canonical-3221310223000132-2023121223223310-1231213203103323-2323213320322111-2110230332300023-2002201030213122-0021303300301232-3112120032110030)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3221310223000132-2023121223223310-1231213203103323-2323213320322111-2110230332300023-2002201030213122-0021303300301232-3112120032110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122313202222211-1331313010123322-1201231033132022-0231000023322111-2111131020121233-3100211111021220-3211121111111133-1020033221302001"></a>

## dynamic_proxy.https_proxy.tls_params.no_mtls — no_mtls / 310101211203 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- dynamic_proxy.https_proxy.tls_params.no_mtls

<a id="canonical-3110323201120332-2313011321232032-0321031121303201-1202230302302310-1300031030210202-2130321210311330-1333033103230300-2212320102203110"></a>

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
no_mtls = {}
```

<a id="canonical-0022321132100330-0101103332213233-0202032211003310-3302211211110120-3131020033312112-2313112000223022-0303110001122003-2100232021111110"></a>

## Direct properties — no_mtls / 310101211203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202101230013002-0313001313031022-1002212313100121-2132322033323020-0203231233021031-2222200102032102-1220002223020331-1222310131011322"></a>

## Next pages — no_mtls / 310101211203 / 4

- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122101231211320-3330012000101020-0032330111303033-0121231201003302-1212231012231232-3231303333310313-0220323320032013-0322231301333300"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates — tls_certificates / 122122031021 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- dynamic_proxy.https_proxy.tls_params.tls_certificates

<a id="canonical-2031120113301231-2000031022332001-1022302121112230-0030201113230120-0130310002102030-2213013201220022-0303112131331320-2323322111003301"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300320210332133-1230321032230022-0031231013100232-1131113031223200-3010233312330123-0113131333131133-1002220012133001-2232020111020021"></a>

## Direct properties — tls_certificates / 122122031021 / 3

<a id="canonical-2323332021332122-3032132312220122-0021301003312032-0120323002120121-3101333201111221-0113200100120131-3231112210133213-3123323102213231"></a>

<a id="canonical-1103302111222100-1110330301202111-0202200103111211-3120011301021213-0101121010022032-3131112212030231-2022220132301320-3321102202030223"></a>

## certificate_url property — tls_certificates / 122122031021 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-1333111333313230-3121323312102302-1312320210203101-1321110330113322-0323123321002010-3111101110023221-0112322321000011-1000311011002201): complete subsection reference.

<a id="canonical-1233032203323230-2121123121301102-1102203133223212-1002120302000000-0200233211020032-1322001213303030-3321131012302222-0000113022220210"></a>

<a id="canonical-0111032012013110-1031320112332310-1231132202120113-3203221122021332-1000031220303103-0121100101332132-2102122123221212-0323133322221323"></a>

## description_spec property — tls_certificates / 122122031021 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-0121121323230130-2231000011111132-3111102133300011-2123022232202211-1120220101023320-2331310121303213-1021202320202131-1221301021210121): complete subsection reference.

- [private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032): complete subsection reference.

- [use_system_defaults](resources--proxy--reference--group-003.md#canonical-1230013300213220-3213200230111020-1103222220212222-1123123333210012-2203033210111223-1303310100021100-1210201013211021-0113320131102322): complete subsection reference.

<a id="canonical-1232021012220231-0122121332020313-1232113222003300-3021111301232122-2323111231322330-1320010003113200-2000023132032312-1220000313320102"></a>

## Next pages — tls_certificates / 122122031021 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-1333111333313230-3121323312102302-1312320210203101-1321110330113322-0323123321002010-3111101110023221-0112322321000011-1000311011002201)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-0121121323230130-2231000011111132-3111102133300011-2123022232202211-1120220101023320-2331310121303213-1021202320202131-1221301021210121)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](resources--proxy--reference--group-003.md#canonical-1230013300213220-3213200230111020-1103222220212222-1123123333210012-2203033210111223-1303310100021100-1210201013211021-0113320131102322)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1333111333313230-3121323312102302-1312320210203101-1321110330113322-0323123321002010-3111101110023221-0112322321000011-1000311011002201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303333113332032-1311201033220330-0013110103211101-2100202103000301-2232332133001113-2303121112020311-0130000311231323-0232321322023003"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 020130010000 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

<a id="canonical-3133010331132323-3100230200101231-2003233212011100-2010012212133233-1332113111100023-2002203322210110-2033320021110310-0202210313011331"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313031313112231-1303130032302031-2113003113302122-2232300120203221-3220100001223321-3211122123021133-1110132310022123-1213233120311321"></a>

## Direct properties — custom_hash_algorithms / 020130010000 / 3

<a id="canonical-1203112202302002-0133332013010003-1321322021312221-0303012201311200-2020101021300131-0311023111023102-3031232223012121-3333030123130231"></a>

<a id="canonical-2122021203030033-1031011103013132-3213011312212231-3232303001000102-0000030313333132-2321131020130302-2120111202210002-2202230303331013"></a>

## hash_algorithms property — custom_hash_algorithms / 020130010000 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-1210201312112001-1003313301231012-1003313020323302-3330201100003101-1301103203332223-0200313213303202-3230330021331310-2021223233012130"></a>

## Next pages — custom_hash_algorithms / 020130010000 / 5

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0121121323230130-2231000011111132-3111102133300011-2123022232202211-1120220101023320-2331310121303213-1021202320202131-1221301021210121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033302132223103-2023303131213022-0012301313113230-1013302322112013-1033111323030203-2231321112223321-1013230132102103-0001322010303312"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 221020201311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-2302023312121112-2211301132203021-3131311123221312-0303220132023132-0222010110330221-2231303022120331-0323033232133021-3301230011212332"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-0131233201232020-1303000031213313-2013223131332230-3132003033331322-2332202113001200-1033300002322303-1303101023223022-2121322120100101"></a>

## Direct properties — disable_ocsp_stapling / 221020201311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130321000330021-1101213210033133-1232300331010221-1112221122011210-2321112233103313-0102003131021132-2200233021210313-2212131011312002"></a>

## Next pages — disable_ocsp_stapling / 221020201311 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211102302313213-1320200130231023-0313111300010301-1312020023200132-2000021313030313-1332333100030113-3212332231213320-1001131000022001"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key — private_key / 310032213010 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

<a id="canonical-3130303022320203-0010103300122112-0020313132130223-1213200333310201-2300301133313320-3222303320323113-1223201120300112-2331120323323323"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000031220220211-2301321132322101-1310103110320120-1031031123001221-2131121130122131-0003301300012013-3123103323303122-0000131310232210"></a>

## Direct properties — private_key / 310032213010 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-1231131331331220-0332121111102020-3121031030102201-0133200202301000-3033023330021123-0122033322213203-0010023103002202-0023022121133232): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-0003010213222102-1131020230321001-0300320031133312-1303322132303103-2131100323132020-0003013212021003-3031011303203221-2002021310112213): complete subsection reference.

<a id="canonical-0022030131002300-2232303013310010-0130212113230011-2330331012022230-1221232300301313-1310023331113321-0200021321310011-3313132001121302"></a>

## Next pages — private_key / 310032213010 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-1231131331331220-0332121111102020-3121031030102201-0133200202301000-3033023330021123-0122033322213203-0010023103002202-0023022121133232)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](resources--proxy--reference--group-003.md#canonical-0003010213222102-1131020230321001-0300320031133312-1303322132303103-2131100323132020-0003013212021003-3031011303203221-2002021310112213)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1231131331331220-0332121111102020-3121031030102201-0133200202301000-3033023330021123-0122033322213203-0010023103002202-0023022121133232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333212013311032-1323030012110222-1310001232302303-0123033022333131-1022013330221112-0003130203203220-1230233111333200-2021301231013010"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 220030323300 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1022212211203011-3100313310201130-3121031100212132-3200330301222013-1112220032120222-3203202200020133-2132310012222110-1121020103123331"></a>

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

<a id="canonical-2203032113131322-3130131102121331-1220031212211031-1132021131302130-3310211003121330-3312122121233121-1102030132223310-1112110211030021"></a>

## Direct properties — blindfold_secret_info / 220030323300 / 3

<a id="canonical-2013032021321303-0320331131310021-0000100310020132-2002223332301131-0102120312230321-2320021022132221-1200113230011121-1111212331213310"></a>

<a id="canonical-2020212123030121-2121303013103022-1300010210022233-3313221032310333-3101102203031221-3320220220111200-1112210332232200-2330113320100221"></a>

## decryption_provider property — blindfold_secret_info / 220030323300 / 4

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

<a id="canonical-2212321321032031-0311313212033211-0031102203310011-0320232013131020-3110222303003100-2001010130221223-2313103220201233-3122113123232312"></a>

<a id="canonical-0002021030100023-0132111300111001-3110003221232113-1101231223030322-2210222301020330-3303101212121303-2222222221312000-2230030321023212"></a>

## location property — blindfold_secret_info / 220030323300 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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

<a id="canonical-2012220303321212-0023123010222310-3303211223103332-2023113022122300-0111132313220221-0122312133333122-1210033033222133-2232011210213112"></a>

<a id="canonical-3032010302013330-1332321132110333-0000300002200030-1133331000002310-2233002310323002-2001200120132320-2202000302003323-2210310202201212"></a>

## store_provider property — blindfold_secret_info / 220030323300 / 6

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

<a id="canonical-0133221132132130-3122231213301210-2201213133031131-1312231221331230-3330231112331213-0300001312330213-0300231110112012-2221031012202330"></a>

## Next pages — blindfold_secret_info / 220030323300 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0003010213222102-1131020230321001-0300320031133312-1303322132303103-2131100323132020-0003013212021003-3031011303203221-2002021310112213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102111032001032-2202222122001001-3232203233103330-1220033130233003-2321010033131331-3122123131111123-0023002230311233-3020321101301121"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 333011021021 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3132322130131110-1210001320210331-2021202123111033-3132223321121030-2113133211030302-3022030310212301-0320210202233323-1233113313012311"></a>

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

<a id="canonical-1102320102113310-2203030120333000-1122001122113332-0233222201213203-0121013300301202-3322112013230032-1331012213310022-2222230132330231"></a>

## Direct properties — clear_secret_info / 333011021021 / 3

<a id="canonical-1230322311311232-3003312200311330-2002111321022301-0233312002033033-2123011310112000-1200321001112030-1213031101302121-3232101021130230"></a>

<a id="canonical-3012322210233130-0302000120031232-3111123311323232-1031211311302003-1020322103021002-2011212302313321-1020000133103101-0010212220002220"></a>

## provider_ref property — clear_secret_info / 333011021021 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1003023330110121-2102020312232021-1022331001003300-2213320130221021-0103022231020200-2312321230010322-3321312201103330-3132333331113210"></a>

<a id="canonical-2010132022002133-2231030000203301-0023031013003031-2301003013020311-0132110131330122-2332000120102120-3110301100001322-1200110131033021"></a>

## URL property — clear_secret_info / 333011021021 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-1101212213003313-3332220231021231-2100021003023012-0023003331120113-0200221122200013-3021131021211030-3123100321011203-0102133333113101"></a>

## Next pages — clear_secret_info / 333011021021 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-2132213022011331-1332022303000320-0302031031322210-1133230310331331-2030230110322320-3223033012111202-0323302330100121-1103331321313032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1230013300213220-3213200230111020-1103222220212222-1123123333210012-2203033210111223-1303310100021100-1210201013211021-0113320131102322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120332322223120-1222222200311232-1123121003011110-3103011310020101-3301321133300011-1320020323122130-0201330113201010-0123233003113333"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults — use_system_defaults / 011113022031 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults

<a id="canonical-0332122102233300-3331002010330122-1010131302320320-2312300020321102-2313010120131300-0200013030112300-0133230310000122-0332222300321322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-1112020101022311-0000321130030213-2303221231120312-3310013100133310-1010212033101303-3330101033002002-3101120232032331-0301111213312300"></a>

## Direct properties — use_system_defaults / 011113022031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320111220303300-1233302322203321-1033310011221211-2310332330020220-3132323213022320-1212330233201022-2133021231213000-2212212202332323"></a>

## Next pages — use_system_defaults / 011113022031 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-2123003001001330-3032122020130321-2211132130321232-1122221122023013-2000320310332313-0001203031303101-1331202203221001-3330303321130021)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111202322301210-2133023200332023-0313122032032312-0111030022021220-3000001020312120-3203031333331212-3120223213000330-1303122200320212"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config — tls_config / 203232302303 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- dynamic_proxy.https_proxy.tls_params.tls_config

<a id="canonical-1231102332233303-1002131231012201-1222213131131203-2213033133313332-1202333030332210-2301022102113303-0120203220131132-0120132312233031"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330232320012310-2320023030000033-2021230030003331-2222313311133222-2033111322030002-3231020301021330-3013013213312311-3330231232022123"></a>

## Direct properties — tls_config / 203232302303 / 3

- [custom_security](resources--proxy--reference--group-003.md#canonical-3021212103003301-0303230022302100-0220330203003210-1011003323333320-1233122202322332-0233101111313030-3230120313001111-1323320331133032): complete subsection reference.

- [default_security](resources--proxy--reference--group-003.md#canonical-2303111020211121-0101013330113231-3033021201102303-3120030313020200-3111031001321221-1203123223123320-1101031302321330-3330110313102133): complete subsection reference.

- [low_security](resources--proxy--reference--group-003.md#canonical-3011003303023213-3101012101323022-2222023002202302-2120112002211000-0312033223010311-1210202313313100-1233120213030110-3022200303310211): complete subsection reference.

- [medium_security](resources--proxy--reference--group-003.md#canonical-2223320103030103-1031020121003012-1001102230213101-3012220301121021-1213302212333020-0210031030032332-3302002133001231-3113033230231211): complete subsection reference.

<a id="canonical-3212213200221102-1023233120032203-1230100101131323-3113322212030221-0111002031312103-3321003301230013-0102003120313123-0023033133130302"></a>

## Next pages — tls_config / 203232302303 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](resources--proxy--reference--group-003.md#canonical-3021212103003301-0303230022302100-0220330203003210-1011003323333320-1233122202322332-0233101111313030-3230120313001111-1323320331133032)
- [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](resources--proxy--reference--group-003.md#canonical-2303111020211121-0101013330113231-3033021201102303-3120030313020200-3111031001321221-1203123223123320-1101031302321330-3330110313102133)
- [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](resources--proxy--reference--group-003.md#canonical-3011003303023213-3101012101323022-2222023002202302-2120112002211000-0312033223010311-1210202313313100-1233120213030110-3022200303310211)
- [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](resources--proxy--reference--group-003.md#canonical-2223320103030103-1031020121003012-1001102230213101-3012220301121021-1213302212333020-0210031030032332-3302002133001231-3113033230231211)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3021212103003301-0303230022302100-0220330203003210-1011003323333320-1233122202322332-0233101111313030-3230120313001111-1323320331133032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321311120003321-3100003332123201-1203012320031302-3001031013221232-0020303221120101-1112330132232030-3122300221030321-0122200103210103"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.custom_security — custom_security / 120023302103 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- dynamic_proxy.https_proxy.tls_params.tls_config.custom_security

<a id="canonical-0122203202132121-3102133100012102-3033333030010120-1210322232123303-2033133030222302-2221233233013110-0030030213000033-3221121133011312"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223330120230233-3120013301100100-2023300032201212-2021121202212221-0211120000032130-3300321032211321-2100002103233003-2020230212211100"></a>

## Direct properties — custom_security / 120023302103 / 3

<a id="canonical-3202110222133320-1231121103021103-2021130200121110-1213223033131101-0333210300010003-0112231100033122-1311030331120302-0020121110032011"></a>

<a id="canonical-2232022203112203-1330200203110012-0112001302213333-1220020003000212-3212221012010303-1231131213233230-3233203122032100-1202002103130303"></a>

## cipher_suites property — custom_security / 120023302103 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-2001303222012110-1030312002011333-1322212222001321-0201022311303100-1020212120023332-1200013313020302-0130332133201103-3202323213023123"></a>

<a id="canonical-2101213113122032-3322203121231320-2103120230302311-0002011101113203-2222100231132133-3321010300232211-0000221332120302-0110200223232123"></a>

## max_version property — custom_security / 120023302103 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-0131113013211331-2130331233213031-3232213301032122-3233133121323031-1231130310131223-0022020032032031-0020032300102232-2102320130233310"></a>

<a id="canonical-2231300213201122-2000321001321211-1200013302020032-3302130101023303-1000133210032031-0312013030011302-0200030033331101-1233233233120211"></a>

## min_version property — custom_security / 120023302103 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-3200131000223101-2220213220320101-1030301210013233-3022000021322213-0130332003213022-3313300210101202-2013230332201032-3313031323111111"></a>

## Next pages — custom_security / 120023302103 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2303111020211121-0101013330113231-3033021201102303-3120030313020200-3111031001321221-1203123223123320-1101031302321330-3330110313102133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130301000013310-0211021023210122-1222111123232010-0302321113123003-1232022032203110-2212312102102201-0213131310321132-0002011101032130"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.default_security — default_security / 002221011311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- dynamic_proxy.https_proxy.tls_params.tls_config.default_security

<a id="canonical-3022110123203303-0330221223113221-0200222033313103-3320231223031032-0313310113031103-1232021230333222-0111201101112211-2030221200113303"></a>

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
default_security = {}
```

<a id="canonical-2120311223021130-2003101210300112-0003033100112332-3312000231023010-3311302211210230-0130113310333321-3323320332032133-0222212202313120"></a>

## Direct properties — default_security / 002221011311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030002013223121-2031001200203002-2211222112031022-3001210311101112-0003110110312120-3213212131311221-0113221101123131-2223120302100130"></a>

## Next pages — default_security / 002221011311 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3011003303023213-3101012101323022-2222023002202302-2120112002211000-0312033223010311-1210202313313100-1233120213030110-3022200303310211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000200020220310-3031312002332131-0011001103123122-3201321020012302-0322013133232001-0223220210133022-3320031202320102-0111010223320221"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.low_security — low_security / 221023323231 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- dynamic_proxy.https_proxy.tls_params.tls_config.low_security

<a id="canonical-3010021210222012-1031100021203223-2101112110133132-3323332130120203-2331332202301021-2032122212101032-2333212002202330-1221022013021101"></a>

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
low_security = {}
```

<a id="canonical-0211222221001332-3332323110301200-2003031130010202-1300301301030000-2213203311310020-3023110303232210-3000222211112210-1322113111111021"></a>

## Direct properties — low_security / 221023323231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000232311101113-2202203023011032-1203021323110232-1003030001221223-2101221333312200-3333200001333000-1122010113203131-2013221022111103"></a>

## Next pages — low_security / 221023323231 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2223320103030103-1031020121003012-1001102230213101-3012220301121021-1213302212333020-0210031030032332-3302002133001231-3113033230231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230121123103322-1012011112000023-2310101120023231-1220030002333213-3022321200030313-1011110212022010-2303233212210033-3032302300312101"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.medium_security — medium_security / 230022000210 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- dynamic_proxy.https_proxy.tls_params.tls_config.medium_security

<a id="canonical-2230000203213000-1120212323122122-3131222220031213-2122333132112022-3033123113032102-1103102203301100-1133032112312220-0312222132210211"></a>

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
medium_security = {}
```

<a id="canonical-1301101001311321-1132222213122200-1212302203312002-2113021223232033-0002111211030010-3210201222031111-3110230313231120-3011302322232011"></a>

## Direct properties — medium_security / 230022000210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102002232231311-2033122233013121-0303021312210210-2101010212320321-1330122130330033-2312130202320023-2122232032023120-2220110201333021"></a>

## Next pages — medium_security / 230022000210 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-3211212103103301-2032200002231203-1320011112122112-0020023320131033-3212030120201213-1323102131133332-2210110022023301-1232133120323130)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1120302120201130-2021020112211312-2012121003331033-1202131121023223-2102330022132231-0123033302022200-3110300012022001-3113320030210032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033032033121230-0312320223110301-1112032003101213-1201212130333313-0013312312010001-1021303020121121-0121103202320001-0220321200013110"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls — use_mtls / 003012323023 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220)
- dynamic_proxy.https_proxy.tls_params.use_mtls

<a id="canonical-1213230003223333-2310120333333212-3222201000132322-0211123031123233-2300311201230113-2300132032310132-0323103122310011-3211110111022030"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121220033102100-3311111312033021-2332032113211203-2032032332123122-0102231300000131-2130201110330101-2112133023103002-1003031111213001"></a>

## Direct properties — use_mtls / 003012323023 / 3

<a id="canonical-2213331233003232-1123001122303213-1302120303331310-2100102113001111-2113211130023010-3120111112220200-3033210300220122-0100012102221023"></a>

<a id="canonical-1130001311221323-3331323203323111-2300010001111321-2311222031323103-3323113113002011-3131232220313021-3233001301223323-2203331211231030"></a>

## client_certificate_optional property — use_mtls / 003012323023 / 4

Type: `"bool"`. Optional.

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

- [crl](resources--proxy--reference--group-004.md#canonical-0001322300020012-2202212022120010-1131232201213100-0000301003323111-1123330313013320-1110121300013301-1020000312223323-2200211201330030): complete subsection reference.

- [no_crl](resources--proxy--reference--group-004.md#canonical-2212123203301223-2303123302332301-1033020013321211-0212311111123033-2100000200020300-1120330133022321-2030023121323133-1302003230302120): complete subsection reference.

- [trusted_ca](resources--proxy--reference--group-004.md#canonical-2230303001100132-2100001122031031-1021322122310121-2330022311303113-3230213033230130-2302031221001030-2120331111023010-0120103221232032): complete subsection reference.

<a id="canonical-2211221112201113-0130230100233301-2332111312231113-1330122333302003-1122111001202302-1211111012300233-0303213023202221-1333000330033201"></a>

<a id="canonical-0021111020232021-1212000123123133-2031010112220123-0013211021013220-2311331021131003-1312101030312020-2111030013100033-2333000322332112"></a>

## trusted_ca_url property — use_mtls / 003012323023 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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

- [xfcc_disabled](resources--proxy--reference--group-004.md#canonical-3301301013003012-1010130000202121-0230031303003100-2111102033131312-0200320003300232-1300031313132231-1320223300221001-0012021223131322): complete subsection reference.

- [xfcc_options](resources--proxy--reference--group-004.md#canonical-0012033023300202-1223030021113200-1023321133332103-1233011321201233-1231122211102010-1031000133330202-2021121032011301-0233203202103022): complete subsection reference.
