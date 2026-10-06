---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-001.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value

<a id="canonical-0001301310201112-3222103103101101-2020122030230212-3313120311010010-3331210011130011-0132302031130302-2112301113003330-3303011221023112"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3311130202331003-0323233311000321-1233300021221021-0210230133322322-2230230320222231-2133130113333211-2322322202021033-2321310002201301"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2003123130303002-0033123013233203-1333010122322230-0020001200312011-0120032321011102-0232230103330111-3313230132322113-3303312020021233): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-2233302133310312-1003211200232032-1301322322202333-1203333211210312-1130021211002320-1323232223300313-1223312230022030-3312332213130331): complete subsection reference.

<a id="canonical-2003123130303002-0033123013233203-1333010122322230-0020001200312011-0120032321011102-0232230103330111-3313230132322113-3303312020021233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-001.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3223110031103030-0322023113110213-3212113131221203-1221233231310202-2203122131232030-1323311101110031-0111133101330202-1320120031332222"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3121320333223203-1112101213303230-3033223320002311-2000001313310111-0330232211332301-3010322010230230-3210320133000102-0102011302021332"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1310103132001013-0331220123000320-0232300020230321-3213032021103012-0311233033323311-1230223230321222-3012130333300232-0311233121231330"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3013220111003220-2112002213301031-1120223012111323-0101121320031000-1223110233000331-3311202231201121-1023210332323031-1320021233013313"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0210131320200000-0302121331121023-3011322101031100-3130032312000302-0123213232020121-3211330033132221-2133123020231032-0311021111001130"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2233302133310312-1003211200232032-1301322322202333-1203333211210312-1130021211002320-1323232223300313-1223312230022030-3312332213130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.http_proxy](resources--proxy--reference--group-001.md#canonical-3010030230300233-0022012030133013-0221201031310302-0303120102300030-2123221020311232-1011333012001321-3003022113003223-0232330232302103)
- [dynamic_proxy.http_proxy.more_option](resources--proxy--reference--group-001.md#canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-001.md#canonical-2202201333112001-1113233232320310-1221001002302311-0312120031201222-3230230022213231-3133322233323010-3332110221230120-0230320130222220)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-2221121133321303-1301031031112023-1111231110021110-2223211301300001-1113002112011222-1211111333212102-1212331102122331-2003210210030301)
- dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2221300200230203-1103230202232212-1120100321231101-3103120231310311-3002322311120212-3230100002333023-0003131213111120-1131313112321032"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3022000213011021-3122323100313001-3323033232022311-1002313102113111-0100121031010011-2331301212013203-2221023222112303-0033231013102220"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-1213310200312330-3303111332132013-0220312223110131-3023102111320121-3212212202313033-3323330101121120-0221012313032322-2333023133211132"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2330132102030303-2122021321132333-0003131032302103-2231033323102211-2021301013300031-2133133120222330-1130300332313022-0333332010332121"></a>

<a id="canonical-3200321001203220-2000300230222202-2213020020331102-2333221323232313-3202133101011031-2212332211011330-1311113213031212-0311201123303101"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1002022321310321-1002330233323203-2202030110100111-2310321303122121-2011321201120202-2233101100223123-3311230010321231-3331122113232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303100311333213-3010122133221302-0103010320222030-3333221032021231-0212001100101131-2301000112130213-2220010032120011-0222232021211031"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add`

<a id="canonical-2210013321203123-1223301312033202-1322303202011203-3101030303300022-2132232022120202-3331220333313313-0311013003021131-0132023130212330"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` property

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

<a id="canonical-0303322301202131-3322330220302012-1330123000223011-2313230210003230-1122303022202211-3222010003121111-2303032132032333-2202032122322020"></a>

<a id="canonical-2103010013031013-2202203122120012-1200332003213222-3100333031010222-2132200132011103-0230112200203031-3323100132321010-2302120212220310"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2022212030301123-0303212221202300-1111202220302332-2323230230030333-1220000110033201-1020033332213303-2212302111213012-3300113120220230"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0101300221220033-0130233333213131-2032112233120323-3211123010322113-3212310103220212-0122231331222211-1312213323032222-0110312232013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` properties

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
EnumExtractionComplete: false
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

<a id="canonical-3222203223011220-3023200113121333-2110311310000012-3323322201132232-3133131220120200-3132230113322331-3001200102321231-3122021311132133"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-2211130000033232-2003133020123311-3203103301100322-2031223301033121-1133223321112330-1100012022221111-1220030321300323-3232201323220102): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-0023100032221102-3000313313001020-0122312201033220-2231022212120022-1011130320311233-3320122110322133-2302122220211112-0003010130203001): complete subsection reference.

<a id="canonical-2211130000033232-2003133020123311-3203103301100322-2031223301033121-1133223321112330-1100012022221111-1220030321300323-3232201323220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-2213113320003003-0010321021300133-3210130022111022-0330222113113313-1121301111322122-1131123013200300-1330202012022303-0021233222322300"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3313331002100310-3203222233132122-3022322110130231-1230222333223010-1031233332333101-0101221211000132-2003220122310110-2331122102322032"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0332033013133313-1133133121022220-3312132130310302-2223101002230222-2003010221212110-2100103113222130-1303133101120321-1030122321000210"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3012002100312022-2002200023001131-1012000021311113-1223022113012032-2002201113012321-3300112301212230-3312032100011010-3030323030233312"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0023100032221102-3000313313001020-0122312201033220-2231022212120022-1011130320311233-3320122110322133-2302122220211112-0003010130203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1320101122213132-2300303310122323-3023101033231012-3203111002213202-2130210032103223-1211022231300031-1210211333002120-1000120330023001"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1020322220112012-1331103312102013-2112030213211033-0121011330123020-0221323101212101-0302120110123311-1012300221310331-1222002130133310"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3203022131030022-1023231022332103-2121110300010113-0320030311100123-0230110331113313-1233311230022033-2020031021001221-3311331130103320"></a>

<a id="canonical-1132111233230331-0021003100311130-3301231223012113-1110303003121122-2233331220322110-1330323032131022-3122213320212212-0101220312133311"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1133232023231313-1010322232321011-0023101201231200-0032321000233200-0311010212003323-1221233300102121-2322121101300113-0222132121112002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2121000110023231-0032220111311233-0103101132321000-0210120211131111-3111330300213322-0010012220113033-1301211033131122-2201330330132312"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add`

<a id="canonical-2100313122110033-1010230312313003-0012110330301001-1001211032320223-3001103223230303-1100310111110023-2131103230213112-1023321333303223"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1300010233222230-3033032020312130-1003311310231222-0202123232123032-2202211230333102-3300222120222303-3121132301133303-2031020200321122"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1121102231130012-3110233223313303-1232300222001030-2000103232321210-0122133301310333-3211101003033020-0203310102320122-1312130122210013"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1311132110202311-3010120230010312-2130110211332201-0201020312212032-0302112230313300-0301100022032110-0331202232011022-3221021113332330"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2333212122322323-3232110233322322-1210130022031013-3022211212100110-0223303110301001-1303020022133310-0200101220102313-2002123102322301"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0221333222033331-3022100211222330-3030332111033003-2330002322310201-1223023212011012-2330010230323102-2001002320202320-0031010033310031"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` property

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

- [samesite_lax](resources--proxy--reference--group-002.md#canonical-0200102201011213-2222230221032313-2311211112321022-3111131322023132-1303311311100312-0013320103303100-1033321320113313-1210233221123032): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-002.md#canonical-3202033321232232-1231123230032120-0011320312320010-0300221231211121-3200303303303331-0021020201102203-3111231121002330-1323013020310132): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-002.md#canonical-2001200323200002-0100021031222112-1113010212212110-1211123102320312-2032011110231213-1333033233201023-1330300201023123-0112001320101131): complete subsection reference.

- [secret_value](resources--proxy--reference--group-002.md#canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333): complete subsection reference.

<a id="canonical-3033331112111032-0030310131023121-2333232311231221-0200000323033201-2310113121310201-0320013201302332-2101113012222202-1313103231001231"></a>

<a id="canonical-2131111200012300-2001222220121213-3022321122002210-2302230302122023-2303010301223001-1123222223231201-0112300323202233-3230333310221001"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3033213011100232-3233233212212022-1300303300033113-2330000333112330-1002000230212220-1131302000110101-3103222123201021-2003012322032232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` properties

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

<a id="canonical-2331122031133303-1013033330113203-3212131002231210-0321211300331221-3101323302000213-1302303001223023-0131100333021111-1232300230300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` properties

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

<a id="canonical-3022132322222221-2113323112133200-1102231302131321-1233003303030320-2010310111120033-2031113023100231-1220013033133102-2220132120032000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` properties

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

<a id="canonical-2232322033322210-2123311102020333-1212111303033133-3202301213220233-1223012313320321-3111221333210203-1013333130213223-1123122212113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` properties

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

<a id="canonical-0201000303101102-0323111232010312-0123131231100212-0101212232200003-0020103212102232-3321311322310313-3023110302223123-2232303021013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

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

<a id="canonical-0323112031212001-3031321031312312-3311323122100110-0113103323111012-3132001022332103-1133202320233230-1233020020000100-1332012132030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

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

<a id="canonical-0232130310110203-3213130300000330-0031101022100221-1111100010120002-2333232001103113-2313031033202112-0002130210011002-0131101301200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

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

<a id="canonical-2220010022130303-2223300332110113-3223232303213321-1123000032003130-1123010131112100-3032323221011212-3222222113122000-1211212032223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

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

<a id="canonical-0101303332312031-3212110122321000-1031230003032201-2220032323110012-0102312203231032-3123230321211000-2030020110130233-2011332003031133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` properties

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

<a id="canonical-0020202232330223-3120103211031311-1133322110201000-0203122103013021-1023032320013130-0313021012010030-1320323110120112-2011121012101233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

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

<a id="canonical-2322330201013310-2122032230131032-1122220231333301-0033033113231232-3233123212011000-0101121312100323-0220201223202113-0023233221100102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` properties

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

<a id="canonical-0322020233022323-3301022331322223-1100330022133330-1010023101322120-0032330133302213-3312331313233201-2000332130131302-1001023102213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` properties

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

<a id="canonical-0200102201011213-2222230221032313-2311211112321022-3111131322023132-1303311311100312-0013320103303100-1033321320113313-1210233221123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` properties

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

<a id="canonical-3202033321232232-1231123230032120-0011320312320010-0300221231211121-3200303303303331-0021020201102203-3111231121002330-1323013020310132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` properties

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

<a id="canonical-2001200323200002-0100021031222112-1113010212212110-1211123102320312-2032011110231213-1333033233201023-1330300201023123-0112001320101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` properties

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

<a id="canonical-3213320213100323-1032103221231202-0133221030112021-3231311111102131-1223320103023110-2301030130332202-1331203110021211-0222320020011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0310303303010100-1130222311130121-3121021233133032-2313123012003221-2122330232130221-1011120010221232-3201112323331113-2121330203132112"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0022310021300010-2002213123221021-1001012203233313-2130010100230112-0312302011203101-0121333223102332-3030033321101032-2033313220230311): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-1311101122321023-1021222121333111-0323323122313323-0232233202323111-0211331223233002-0000031332320020-2330323311321320-1033323213111111): complete subsection reference.

<a id="canonical-0022310021300010-2002213123221021-1001012203233313-2130010100230112-0312302011203101-0121333223102332-3030033321101032-2033313220230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-3200000301322021-1130333131120231-3122100231123210-0130000130332132-3023320033302113-3200300133230300-3132111312100112-0003100223132322"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3200330221230330-0303003013011101-3201233120132000-0111201003030010-1000003030202301-0012101021320321-0032220321103120-2230312033022203"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3022312023131132-2230002201023132-3113000312303320-1021322202121312-1202133003001030-1330222300003120-2003223113132223-0330232222110201"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1231232023033333-1111021110032232-2323100213020231-2021121222230101-2312011301332320-3133330303200012-3000103232013211-3221211223313211"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1311101122321023-1021222121333111-0323323122313323-0232233202323111-0211331223233002-0000031332320020-2330323311321320-1033323213111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1311333331130010-3211001222322333-1330221202132301-1310211003001303-0123102201132130-2122110003312332-3321213303011331-1131112101330221"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0123003323021323-0110002311323103-2122021320231112-2310202322131033-0011003010013003-0232000013000311-2203232110333231-1311212023230013"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0201123322102021-0332103110120030-0303131202223122-3010101022030332-3022321301233220-2131221001201222-2111101201101301-0330033203113222"></a>

<a id="canonical-1130003221312013-1222311210132212-1102210230112023-2333001300210100-0112030313032033-3302101123212013-1203022113001031-1313020211220022"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1212230120213332-0202331232012331-1300210112213133-3300202222202130-1033203202131323-2032233032303103-2021201331031231-3300133123003000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0000122310002303-2000223311103021-0311120301130332-3123323001010230-0101210102023031-3002330313020213-1122103111012120-1333100120210331"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add`

<a id="canonical-0313113011332211-1211103020323313-1103320000202200-2311201322321131-1013122100003131-0131131130221133-3220231302031120-0002020010113121"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` property

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

<a id="canonical-1001332203111311-3201330300033222-2220330221012232-3003132020323121-0223302002212001-0122033233232232-3120001132233100-2133312213322230"></a>

<a id="canonical-3011000032110133-3103331113123323-3110033133321103-3201111033220302-3122230331002222-1313000320012321-3322213300020121-3112232311211300"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1033130201032221-1232132001021313-2022302010120333-3231223111320311-1022230203011200-0330020212011112-3021111220003321-0212002022122221"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1100320102001211-2122102000322023-2123231230220300-2330123323131200-2103123230013202-2303122131100120-1121000100020203-2022120000233112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1111132111122203-3033101022301111-1302332300111033-3031310020203201-1110301232332313-2032331121203121-2321320310322321-2000302303220312"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3102132133120001-2322032200131131-3102200312123023-2111333031002203-1012032100010100-2010312231001232-1320103210121032-2211030310011200): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-1202133032020033-3232201002223003-2332303300323100-0232010323212231-0320330110012230-3300030113122103-3113013333231102-1103323223122200): complete subsection reference.

<a id="canonical-3102132133120001-2322032200131131-3102200312123023-2111333031002203-1012032100010100-2010312231001232-1320103210121032-2211030310011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1210003031013121-3031113132023300-1302201110202303-3113201211133200-0133001220013333-3101121103213103-1112312221311032-1020123122020323"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0111312023100012-0201030133111300-2131321030212110-2010221001021223-1110103121103301-0112321333220132-0221021002230111-2222133112121320"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3203331031132122-2302311322022331-3130031200023330-2023312331321013-1303103121211033-3203302322311230-2310010230222200-0130333302100303"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2301121110123202-3122021321113203-1011002110003330-1003121002302101-3003203120200313-2212131120112130-0033320011221133-3302033232120102"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1202133032020033-3232201002223003-2332303300323100-0232010323212231-0320330110012230-3300030113122103-3113013333231102-1103323223122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0120203220033320-0132020333003011-1322132332133121-0312111010321030-2011033010020101-0001313002202120-1212202312223330-1321321220113301"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0001323233103000-2131102212210230-0021113103200010-1312233010221122-0030120202322203-1310111302111023-2220311202212032-2201232223022120"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1303101323232033-2311301322330003-3111330303000301-3111332002311010-1230120122211013-2203320132112002-0232002301012013-0110122233321010"></a>

<a id="canonical-2013220233322021-0330202021011022-1032111003300021-0312103323303100-1003011012311000-0021010331113220-3202212121100301-2322102001121110"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- dynamic_proxy.https_proxy

<a id="canonical-2100012333203313-0133032221230031-0122133000311232-3230212201200101-1112310123103300-1212322311331203-1321131033021110-2123132102220131"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https proxy.

Additional upstream details:

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

<a id="canonical-3211300100131300-3311111303123221-1030321233331213-3323311213121120-2221233311331232-3221212110030031-0311323303223112-3230200033102300"></a>

### Direct properties for `dynamic_proxy.https_proxy`

- [more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202): complete subsection reference.

- [tls_params](resources--proxy--reference--group-003.md#canonical-0202230303213331-0000212221213101-0112112302233312-1211230320301111-2131110223221220-0201320302231332-3120220330312330-1222032000312220): complete subsection reference.

<a id="canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- dynamic_proxy.https_proxy.more_option

<a id="canonical-2233030200133233-3231230013030010-3020020220201332-2131311031203311-3130100232221111-2112323321112031-3220313030202200-2202302301103331"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1022031102201001-2102122233120302-2312120230300013-1101212133132002-3321200121312023-0131010121111320-3101333031233030-3310321101320022"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option`

- [buffer_policy](resources--proxy--reference--group-002.md#canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123): complete subsection reference.

- [compression_params](resources--proxy--reference--group-002.md#canonical-1222132113133101-3012113301330313-1001231201322310-3102232230330302-2311333002202122-3101002112021213-0301331130312113-1223113123330210): complete subsection reference.

<a id="canonical-0120301120131103-2132222303132210-2230010301111310-0101301232231221-1230223112031231-1320120223021012-1112221211020313-0203203121120311"></a>

<a id="canonical-2123200103311002-2032032130233030-1232102031013322-3302300320001222-0133100220000332-1123300011212121-2231133222031311-1230201211021013"></a>

#### `dynamic_proxy.https_proxy.more_option.custom_errors` property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"ranges\":[[3,3],[4,4],[5,5],[300,599]],\"type\":\"uint32-string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.uint32.ranges\":\"3,4,5,300-599\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"65536\",\"ves.io.schema.rules.map.values.string.uri_ref\":\"true\"},\"values\":{\"format\":\"uri-reference\",\"maxLength\":65536,\"type\":\"string\"}}")}
```

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
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
      "type": "string"
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

<a id="canonical-0130103131212132-3201322200300301-1121101102301032-1112001220113120-1101122133331223-1232330332230111-1213311020000112-1333202322003221"></a>

#### `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` property

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

<a id="canonical-1102232333313101-3112102103301212-1322212020003013-1212103033011121-1301120333220203-0331011010332320-0320000013122323-0021232022332030"></a>

#### `dynamic_proxy.https_proxy.more_option.idle_timeout` property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3300220320312221-2210302032021223-3011210102020132-2311222332231330-0312031030123132-2113332100200303-2100211312132302-3330120001012003"></a>

#### `dynamic_proxy.https_proxy.more_option.max_request_header_size` property

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2231231310322212-2113021331331301-0131023233030330-3200112012211122-1323020133232302-3211233112333231-0010100033130012-2003130212220001"></a>

#### `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2213321322100022-0021321303033203-0011233002333333-1313001213331003-1223333220000202-0312210321021100-2330011211212131-0001120320202020"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [request_headers_to_add](resources--proxy--reference--group-002.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121): complete subsection reference.

<a id="canonical-0333222211023020-0202111230101111-0013132323101331-0333002103200233-3113331100111002-0330101021222100-1311231023010320-1030133022011333"></a>

<a id="canonical-2233200031330121-2202203012200033-1022312212012333-0221122300001212-3023310310333210-0223021020121331-1131031031123213-0232033303110300"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310): complete subsection reference.

<a id="canonical-1333121220313131-1002010123102033-0023123312002331-2320312030221110-1302103230333132-0313031310100112-1302321211021203-1102211302130302"></a>

<a id="canonical-1223131302020133-0200323011000021-0003302332021102-0333223330132102-0300020221333231-3310033023331203-2011123011330012-3300313320222130"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1331020020033222-1200233121033330-2202203333113203-2230230300013300-2032332130223330-1313001123210100-3233322101002230-3200322021020032"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.buffer_policy` properties

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

<a id="canonical-3323100112320331-1213200102322210-0013113223210102-0210112120311331-0133311023011230-2030021203230203-2320031231013320-0323201301302113"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.buffer_policy`

<a id="canonical-2002021103022022-3202332002103332-1030101320100133-2112023231303200-1202300313232133-0130201001313220-3102322112132010-2101302313221133"></a>

#### `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-0133202101231020-1130122130310221-1103222113022201-2201300310012222-1030230000013120-0232012233013332-2102310200031233-2223212010132132"></a>

#### `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1222132113133101-3012113301330313-1001231201322310-3102232230330302-2311333002202122-3101002112021213-0301331130312113-1223113123330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.compression_params` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1223031031202133-3120301111021021-0213113231321312-2022101301001011-1130001220233032-2323021010110100-1331211322030100-2300323121121333"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.compression_params`

<a id="canonical-2112113302012012-3121020300113223-2321013212003013-3000322232122323-2110022112000032-1310333131231321-0002301202111200-2331102331333002"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.content_length` property

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3132222222223010-3133122120121103-2103232212332132-2223212103112321-1022212030322303-2002100233033002-1003131120331300-2330220013200120"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.content_type` property

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2133102121022210-3030333222203333-0012031213123100-0002023000132312-2211003000321023-3322121311221101-3221313020112002-1111133113101100"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` property

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

<a id="canonical-0131212132302121-0123131023033013-1123023000110200-3303301002211331-3222130313123131-3200321012112303-1303232021222033-1322233131211330"></a>

<a id="canonical-0033011201113112-0121032013110230-3113130231132323-3321103300002020-3313300232230303-2123302323010110-3202011223111323-3312002123221102"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` property

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

<a id="canonical-2012321031232122-3103310110123120-2212322223233021-3221100003023213-3103320011002121-1310230233220211-0201013323111120-2310322133221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.disable_path_normalize` properties

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

<a id="canonical-2011301323311002-0100322122311323-0033200002112312-2130133101033220-1100122111003300-0332213002323011-3222122013313333-2003100011133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.enable_path_normalize` properties

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

<a id="canonical-2210210100110313-0022103330230330-2121233301303310-0222013023232232-1203321013323021-0311212202110200-2133301313121203-1110330330320131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` properties

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

<a id="canonical-3322322130003100-1322233333210102-1211131211102121-0331020230033122-1203110233313310-3033030113221232-2030321102311131-1221221000123203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1210322320020020-0112000233122230-1222201331133302-0033012021002010-1010111100211311-2321203122300322-2212231121323132-0323121121000123"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add`

<a id="canonical-0233121311213230-2320122132202220-3310111002302213-0030320122311312-1320111112110000-2203112313300130-1101100233322031-0211212311210330"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0202001123110010-0203111100231211-2331333110123333-2113131002001132-2101003221011313-1303120121201001-2133221201323230-2313112033323232"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` property

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

- [secret_value](resources--proxy--reference--group-002.md#canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000): complete subsection reference.

<a id="canonical-3301311032222000-2132003130312102-0222222001302200-2023201332221203-0203000113010010-2210121303320223-2123221123132133-3101311130220000"></a>

<a id="canonical-3320130122320223-1233131323112301-1113212010112302-3011133122201111-1003023131322222-1223110232310322-3122011023011300-2020203002002320"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1201100231202023-0100112011213022-0120103013023121-1230222333002210-1013332322111130-1122020020210230-0032113111100322-1330212032033000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0301100002130013-1012123133131013-3110010232322231-3202301312002013-2010303021022220-1031101300302012-1302132033012002-0101003322000101"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-0131201303010331-2131112100321201-0111200122301000-2302302023001302-2012302030011111-1322110221100122-2010013001323123-0332022032332131): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-2212311010210312-3001323222221101-1032000012131230-2310303323220330-0312213122313330-3300131220211130-0311331021321313-0103011210120000): complete subsection reference.

<a id="canonical-0131201303010331-2131112100321201-0111200122301000-2302302023001302-2012302030011111-1322110221100122-2010013001323123-0332022032332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0210333232223313-3120223320112332-3322330212111323-1331132332213002-2320121112003332-1002310223122102-2021023021031331-2031320002230303"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3122102131003202-0121230121032320-2101111100300120-3021202322122002-1333211111312322-0012321301120022-1233121320011110-0231302011032010"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1301131230033112-2312130001010322-1200012032022113-3102223002221200-3120213030100222-3002030112312001-1200133213132333-3321233111212230"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3332033012023222-3131312033113213-0211221033221123-1213011011023310-1320332321333222-2133322102230331-1112311133331112-0230012122123023"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2212311010210312-3001323222221101-1032000012131230-2310303323220330-0312213122313330-3300131220211130-0311331021321313-0103011210120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

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
EnumExtractionComplete: false
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

<a id="canonical-2221022100311112-0333132111211322-3221330301332313-1030001021010000-1123321003232202-1213113100303011-3223312311123121-3031121211330132"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-1010222013311110-3013312213202221-0010103320222100-2021001133331221-1212111221030030-3310111033102323-3131221122320202-0120213103031011"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2231312211132020-1312013002030110-2332023221121312-3133210232322231-0100332300110131-0033231133322021-0203323230101133-3000000102211231"></a>

<a id="canonical-1003213313131110-1313111232220300-1021333030210203-3030302011331312-0231310321012223-2113230130320100-3121133113113310-1222202212020203"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1303203110103322-1213031132130120-1013321001133330-2210022311333032-1102213010200303-3210311212311132-2103200121123010-3111020120110211"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add`

<a id="canonical-3233301321211031-2013200220102310-1312112023102103-3011011010013121-1023013130332311-2223311313330303-1122333330323220-0311300031123102"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` property

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

<a id="canonical-3120321211211202-2211101113213230-1322010210302122-3333113012021100-2032222311002333-0233100220320010-2312120201020121-1013232303321021"></a>

<a id="canonical-0221102133333002-1211312331313302-0210030213301231-2121011222021310-3023311123213200-2010101223030231-0021323112202330-0031011003100331"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [secret_value](resources--proxy--reference--group-002.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133): complete subsection reference.

<a id="canonical-0032021220222300-1223123002201013-0310321012201013-1302001133331103-3103201101323013-0211331220000001-2221332013231033-1203223223303021"></a>

<a id="canonical-1223113211100311-2210333300320312-0222211000202312-0203000331200022-0123033132002311-1100230223312302-2112022010103012-1203301010112132"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-1303010212333201-0312112222010130-3120112023110313-0111300333032303-1303110323331133-2112203122322323-2322020222113001-2310023120202323"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3333233123300302-3312102313331323-2202131310101330-0032232201033131-0113021102230333-0233030232001321-3320333233310132-2212023111303231"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-002.md#canonical-3000332013332130-3203012020121100-3013131023233003-3000321012022221-1022211012121102-2100131301022002-0310322101102330-1121322120133011): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-002.md#canonical-3013211221333301-0102123101313322-2212001003211300-0333303021031220-2120231322223102-0022310211133032-1202121330320100-3120232322000131): complete subsection reference.

<a id="canonical-3000332013332130-3203012020121100-3013131023233003-3000321012022221-1022211012121102-2100131301022002-0310322101102330-1121322120133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3132331131232130-2101121322323000-2133310112310031-1001212003223330-3321013212313120-3301331230321203-3223030323201323-3230003121001321"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1302003021131021-2110312023131213-1220321220222212-0030212233020131-3232331002003000-3011220331233330-1031210000020113-2020032130232033"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3131320232011121-1103213112011320-1023012021310212-1102312211233320-1320313313330212-2300213220301013-1321123012033100-3113330233303332"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2311321122100221-2003233332003133-1101211130311213-2213233032211102-2231012021301221-2103210232001100-0100231130012311-3000203330212022"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2220321312311231-3231202032010310-0203333331300130-3220201120000300-3221030201302311-2232321311102022-3223010321031100-1122331213221033"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3013211221333301-0102123101313322-2212001003211300-0333303021031220-2120231322223102-0022310211133032-1202121330320100-3120232322000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-002.md#canonical-2332131110221310-0131223230021220-3210111333022122-0020030210233000-3110000232030310-3022021323022132-0331013111332321-0033103233003121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-1113023023131221-3033003303233213-2032331001111220-3212121321321120-1023100221013323-2220133222000301-3211120013201211-0220002131113133)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0310112231113133-2020130021131213-2322210232013012-0221312302133303-1121013010113222-2030022331030323-1131330111233220-2003023313023022"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3223130110220203-1231330123222313-2021311301030130-3221122133303011-1221333310111110-1023232103122002-0102013013132133-0233011223221220"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-2303020312220032-2030321103210031-3221323111000222-1111301302223122-3120332101103212-1331311203033300-2233030323311112-1322300302210013"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2103001302111231-3120010233233213-0010232311133301-1323102122312231-1032132331201232-1332231320131312-3101133301323112-1330013111301221"></a>

<a id="canonical-0300132000333312-3002003200023201-3102022301002312-0022023123123331-3122033111301312-3203032312333011-2020320102031330-0102313031332100"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` properties

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2103323303202300-2023231123200223-0013330032130102-3112330202003312-3100110203111002-2003202302333033-0102001022232313-0312222133110232"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_cookies_to_add`

<a id="canonical-3312132021303220-3012010312122123-3121230220022103-0313223031003232-0020013312033112-3122201031312331-3201321203300333-3302021323313103"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1320001321232232-1100321331310131-3330303202000131-3100021200222322-0023021302020102-3230321223023031-2211312300211013-2012322222132330"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [add_httponly](resources--proxy--reference--group-002.md#canonical-0322010232101122-0220323003101213-3323131000321030-0300010230213023-1331022000133201-3201200302000110-3021322321022201-1202123022311221): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-002.md#canonical-2212200310330103-1032232331211021-2113122202302111-3011332310020302-0322000231121330-2200301223120322-0302331030022301-3123112112020102): complete subsection reference.

<a id="canonical-0321121023130330-0101333112232023-2302301202323130-2333210020122303-1010111313310300-1133230111330113-0033203203302001-0200211333102031"></a>

<a id="canonical-3133030113202000-1121113003112012-3122331100010231-2220302200133313-2103122010131333-1320012002003101-1311122211032012-3221210331032013"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [add_secure](resources--proxy--reference--group-002.md#canonical-0001010013121313-3303320320020010-2101110003333123-2232132113101033-2302203033311322-1110222122010213-2232121111202200-1322020110121301): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-002.md#canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-002.md#canonical-1102103003333103-0031133300111212-2030023201032333-0123220310012322-1300131203233230-3212111012333230-3013201033020311-2032132232031022): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-002.md#canonical-3321312012221102-2011213032302220-0111223111010211-3032230011320113-3310202223120232-0132133130131112-0130021113003300-1303222003311323): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-002.md#canonical-2200202100213111-0132121020222001-1203333223222122-0220200332022213-2003210022021220-3311003320233333-2233131101110012-0103013203133013): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-002.md#canonical-3010103012002231-1230233022121112-0132103120103221-2101131130231201-0332012312311031-1022022021313131-2313101101103111-0213323330103331): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-002.md#canonical-2033011031112123-1001200301010122-2102002131220130-2300002033112303-0122231133323312-0213001310121011-1122322312200210-0300020030310111): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-002.md#canonical-1211213001210301-2322033012322331-1103132223010023-0321332020332102-3330322330101330-2133023313300333-1213101111220203-1310233003200121): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-002.md#canonical-2213321311302100-3220123332103320-0211203033210101-2323003122302332-3320323021310000-3331321200102210-2303200132132310-0300032002002013): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-002.md#canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200): complete subsection reference.

<a id="canonical-0330131020302331-0032121212210102-0012320311313033-1210032030322203-3021012300033202-2311021312302011-0010110133021010-0220130013221313"></a>

<a id="canonical-2222000320321313-1002133100103101-2132032102332222-1133013202113221-2111202330213312-0100100003221221-2231301023323020-3110220222300100"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` property

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0302321303033020-0222331323012020-3210121020101303-2102230110212211-3132321103331010-0313131333101020-0310201331032112-1300003103122213"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1232020301203221-2132332212312030-1321013120123021-3313230100333201-2323103132230110-3211130033320221-0133112300312300-3311202021231211"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` property

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

- [samesite_lax](resources--proxy--reference--group-003.md#canonical-0130331210231000-1311132112001330-2100133002022130-0003332021222230-3032210131202201-1203330013213010-2013000302222323-1131300123132013): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-003.md#canonical-0211210313001301-3213002022100221-1031002001112022-0232231100321333-2030212220310303-2320203300312102-0221301031213023-0102223331033320): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-003.md#canonical-3211313331201302-2103100000320112-3123223001012122-1003210203003032-2000002323123311-1202010022011332-0101100012211022-0002011001123110): complete subsection reference.

- [secret_value](resources--proxy--reference--group-003.md#canonical-0031210303110330-2221030212103133-1031130122330200-0112331110330330-1212002302131203-3203201221211210-2220103200223100-0032303130013011): complete subsection reference.

<a id="canonical-1323110321001120-2110022133101211-3102123123300333-3111123130313331-2023020331210132-0230010210203323-2210322131222213-1022112330132010"></a>

<a id="canonical-2113321303123322-3010330311210210-3310200031122232-2302322332233021-2031313331302301-3133210222311303-3320113322230123-3232222020221331"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0322010232101122-0220323003101213-3323131000321030-0300010230213023-1331022000133201-3201200302000110-3021322321022201-1202123022311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-3331321132111110-1303201230300302-0133132303130120-1301301310003332-3030312203032331-1230323220232211-3201201113312202-3302003011000303"></a>

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

<a id="canonical-2212200310330103-1032232331211021-2113122202302111-3011332310020302-0322000231121330-2200301223120322-0302331030022301-3123112112020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-1022300231122233-2300033323102210-3210210032322033-0103000321223020-0003332020021001-2033302332220231-1233303333300203-3330021000233031"></a>

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

<a id="canonical-0001010013121313-3303320320020010-2101110003333123-2232132113101033-2302203033311322-1110222122010213-2232121111202200-1322020110121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-0320101201012130-3320000121222322-0031313133130131-0102313220030212-2221321203303301-2221132310233213-0120211032210332-0310102022311012"></a>

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

<a id="canonical-2130212002130011-3311123202110101-1333211303332300-2131100332110020-1022320212311020-3331310330033132-3111021202100330-3032323232312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1102102133022220-2301331301132113-2122123132110202-0101023032010320-2101011033101130-1320110001122221-2033332310203100-2131112031122131"></a>

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

<a id="canonical-1102103003333103-0031133300111212-2030023201032333-0123220310012322-1300131203233230-3212111012333230-3013201033020311-2032132232031022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-1311232102322203-1001312301000031-2020313310322201-3002221320223020-0130032101211000-2010202011303310-3121132222001120-2120230312230221"></a>

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

<a id="canonical-3321312012221102-2011213032302220-0111223111010211-3032230011320113-3310202223120232-0132133130131112-0130021113003300-1303222003311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-0331322312001122-3332031002303031-1012002121213213-0302133020212111-1221110222013022-2002032021221011-2211023131100310-1221102211323200"></a>

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

<a id="canonical-2200202100213111-0132121020222001-1203333223222122-0220200332022213-2003210022021220-3311003320233333-2233131101110012-0103013203133013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0021033021202111-0202011120112302-0320213102111313-3001303133012031-2331211301200103-0311322321202030-1120312230312322-3321202312011030"></a>

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

<a id="canonical-3010103012002231-1230233022121112-0132103120103221-2101131130231201-0332012312311031-1022022021313131-2313101101103111-0213323330103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-2330223031212303-2210033121111302-2330203200131033-2013001021133011-2211132331013130-1311213321021200-0303332103030133-1220121321013321"></a>

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

<a id="canonical-2033011031112123-1001200301010122-2102002131220130-2300002033112303-0122231133323312-0213001310121011-1122322312200210-0300020030310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-2210222201222220-3322123223010022-2320221122133010-0022221332023210-1221230022310211-3313301302313232-1011223331131213-2120030132122021"></a>

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

<a id="canonical-1211213001210301-2322033012322331-1103132223010023-0321332020332102-3330322330101330-2133023313300333-1213101111220203-1310233003200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-1132103301333331-2100322212221230-3303203330321333-3232023010201303-1230300230322233-2000300022011223-1302133020123210-3201302300202320"></a>

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

<a id="canonical-2213321311302100-3220123332103320-0211203033210101-2323003122302332-3320323021310000-3331321200102210-2303200132132310-0300032002002013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-3022331022100322-2320213332033311-2211032211200223-1201013302021000-2200012201012321-3220210333333301-0112310322101133-3121201211011300"></a>

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

<a id="canonical-1313012130201032-2222121001201222-2300123313203311-0221001020222032-0033020033131021-0231111013123123-2100123102122101-0233003112002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-2113202112212121-1011022102113232-3230331132220202-1123220031210121-3101202003212120-3201030233332031-2311232301212001-3100202032113330)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0033132111320233-2011322231113202-2020323111021022-1301011120122032-0130023120000002-1231101010311312-1231330202000102-3301003110212203)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-1100230213313000-0130201002302032-2021202213303231-2021102323011312-0123330102131332-3113321301230331-3020211222211311-1223311310330202)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-002.md#canonical-0332321332102213-2113202202211022-0131202101220121-3302231222320201-2200202021300121-0000022100331131-1232120022001101-0100132301321310)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-0201122322233323-0301231210021300-0301220222232332-3310031130132213-3103120322000122-2021323310123102-0303333313212213-0303022313321112"></a>

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
