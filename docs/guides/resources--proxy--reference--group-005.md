---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1210113233221230-1203323122300112-2012110322330310-2123210112123301-2110120032232111-0200223332121310-3122221101333302-1132231101001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3333123130100000-2100030320313133-1023322130300210-2110122230231023-1003110310131130-2231022121323302-0201202132100000-0313013132223032"></a>

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

<a id="canonical-0100330312301200-0222330323230300-1110223111232103-0232200313002022-1033230022202312-0322333130130121-2131201103022023-2200211122303322"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1322101213233323-0021003222021021-0313103233202203-0312133311320113-0131331130102022-0121103111122102-0001233021330003-3030033331111011"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3310112200003233-0330310212312012-1112233102021320-3111222012311013-0301010202320300-1023021312220311-0032032112202232-0312221232123101"></a>

<a id="canonical-1220000310211323-1030302032122301-2212200211001012-0213310302210002-2012131200003203-3330331232012013-0301102033013220-3013122103100210"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-0121031200311302-0221100023112130-1130220012202102-0122110212010003-3322131220212233-1210200213131113-1113030212001213-2303100300012211"></a>

<a id="canonical-0032321033231031-0110000000012123-3311123220203111-1302312023013230-3111130330233102-2102302230122133-0203000200003122-1232313000232232"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-2111320201010202-1001012323322201-0233312100110032-2123030212002221-3332123331003122-0122100302330333-2233101000313130-2011200001003302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-004.md#canonical-0000001231023213-2333232322200002-1101223312110200-1332032022023223-3210130302001101-3210211200123011-3203100230022300-3110000311031311)
- [http_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-004.md#canonical-0032223220331321-2122103333322322-0112302321320233-1222120311123120-3122320100100323-2000101331022231-0320013223311032-0110033301203222)
- http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2100103200111022-1232311232303131-3302200200002321-0030031313122311-0130110020122031-3223311213201232-1103222003022021-3112313133223303"></a>

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

<a id="canonical-3003111303301111-1313123220300123-3301121123120001-3302121220313130-2031333002302002-0121223022321111-3310130312221033-1301121101303203"></a>

### Direct properties for `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0332120021321111-2101201130210302-1221133321311313-1120312111011021-0302231010213031-0211111202131032-2210310133311312-1322333013301002"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1232330320220220-0010100010231021-1003301110333302-2103220211230100-2203023302020221-0020320032110313-2301113230320311-1110331111003023"></a>

<a id="canonical-3202130313032021-1031000111001131-3130112300111121-3102131121012303-1010131331121310-0300230100030303-0302010023213212-2321322200212213"></a>

#### `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- http_proxy.more_option.response_headers_to_add

<a id="canonical-1333331213213220-1313002212301213-0311101201203133-0102130131011013-3101203013301303-1021301033120000-3213210013232103-1221322223212331"></a>

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

<a id="canonical-1310323122233211-0000102322010201-1332032101323203-2210313231021022-1120013021012001-1322032102202333-1130333210312230-1312020110300230"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add`

<a id="canonical-3303320120322021-1301322213313021-1233212312311013-0323322032320113-1212323001011330-0302021030210220-2201132011212212-1312031111030303"></a>

#### `http_proxy.more_option.response_headers_to_add.append` property

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

<a id="canonical-3313221033320223-3302221133110232-3022030031120130-3320111101111302-3122332312323210-1122212303010201-0120121100301133-2101033112331321"></a>

<a id="canonical-3102213121020230-3323311203011121-1331302033012112-1123102011022112-3013331021123202-3102221212312001-2113033112002111-0013220213031211"></a>

#### `http_proxy.more_option.response_headers_to_add.name` property

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

- [secret_value](resources--proxy--reference--group-005.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013): complete subsection reference.

<a id="canonical-2230222201003303-1330320021123112-0020112223121310-0231110111313132-3022211301313222-3131033301313020-0031121301023313-2013321012033022"></a>

<a id="canonical-2030100032301000-2203311221310321-3233321333312223-0203132112132132-2220120020310100-2122032002132232-1213021132121221-2030322333112000"></a>

#### `http_proxy.more_option.response_headers_to_add.value` property

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

<a id="canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-005.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- http_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-3101123131023332-1232313122120131-2221032013303333-0122202011000103-3213012213321111-1320331103202313-1202132103012100-2111002322232122"></a>

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

<a id="canonical-1033322311233332-0313002332110220-2123110123202121-0000212031302330-1003333312131010-3212023313210001-0033101222102210-2023212000202122"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-1001121112013303-2001222331111110-0131023301312020-0200231122202010-1223312200222032-3230330332320230-2311000203321210-2132122001323023): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-005.md#canonical-1010013020132002-1033212202221011-2332220021033303-0313002000023320-3113323013011221-1103111030031202-3233202312232001-3330221032020030): complete subsection reference.

<a id="canonical-1001121112013303-2001222331111110-0131023301312020-0200231122202010-1223312200222032-3230330332320230-2311000203321210-2132122001323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-005.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-005.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0121133002111022-1002000221032031-2000113021022120-3102302022222122-1120232323303023-1333312301201303-1230031121122202-1210320112231313"></a>

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

<a id="canonical-3312233110300301-3311223000012321-2321010121032022-1132131021212130-1331132123133303-0130133221031303-1012330231030033-3033011002023210"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2110212123001032-1101302133002011-2102223322230132-0323011323112000-3312012012213031-2230213230111003-3221111023332321-0310132232232210"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3231122103112003-3313302033100331-1321100220000233-1013220000002203-1331121230000032-0110231310013311-0002211203010200-3310312003113222"></a>

<a id="canonical-3120200300013223-2013131113000110-3130022000201210-3123011311320113-2003321123232033-0123121011300020-0033333221013201-1003220001122113"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3013103000030102-0023123021032023-1321310333023132-3122032233111331-2202023032013132-3212233000312313-2230002200220123-0223012100001311"></a>

<a id="canonical-0113321321330232-3311011332001131-3233332232232133-2302000030123122-3230211113120312-0330010132231330-1110223223131210-2130222100102322"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-1010013020132002-1033212202221011-2332220021033303-0313002000023320-3113323013011221-1103111030031202-3233202312232001-3330221032020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [http_proxy](resources--proxy--reference--group-003.md#canonical-3302133012220212-1002331310020302-1203301311211210-0210121333012131-2110013230033031-1011330131123021-1320232033123130-1032121203032331)
- [http_proxy.more_option](resources--proxy--reference--group-004.md#canonical-3001012213010301-3221113322132133-1003230211323111-2221122020301212-1212323320201323-2103212032121003-0332130232112122-0312011030301003)
- [http_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-005.md#canonical-0321302212120332-1003310003320101-2331211322200330-1231022123102030-0001322133103022-1202202120331013-0312112103131001-3233331200001032)
- [http_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-005.md#canonical-2300011033210230-3200023212122302-3133032202031123-3031312001002303-2211020320303111-3210011300022320-2331101223122002-1130322131232013)
- http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2200213112232300-3313331032330221-1202320001001221-0020102133011231-1300131333022001-0220201012233223-1332132221012203-3003301310020330"></a>

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

<a id="canonical-0221033320231023-3010201310033232-3201130312233133-0031010233301012-0220312312121230-0102200031003121-2001001130021312-1133220032201020"></a>

### Direct properties for `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1020230230001113-2222202233133031-3223103120332120-3002322321003020-2102310212230102-3213032033033320-3232300030202000-0122013112300020"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0330010312031012-1031211200003032-1113101120121000-0103322101000003-1123231100201303-1330122212321002-2013223333021300-3311313330112122"></a>

<a id="canonical-0101113223010001-3021102101313232-2123212021030310-0112313120222100-2231011003023211-3020031022331312-0223011001131232-2033033031331001"></a>

#### `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-3333131321121110-3221311322032300-1111103312333233-0321012201211000-1102333201111020-0020303330301101-2313211321302213-0221203120310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_forward_proxy_policy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- no_forward_proxy_policy

<a id="canonical-1302311121130102-2131203200103320-1303323313202210-0113023213112321-0121211321110201-2311112312303120-1201103232201010-1112001213313132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_forward_proxy_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001311021210113-1203322120301211-2003112333102203-1131010301220310-1212123211110331-3201330202033133-1001002102301300-0023323331213323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `no_interception` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- no_interception

<a id="canonical-3332023313312012-3012310221032321-3101200211301200-2222112211033303-1033311031230203-0102333033323201-3101321210231120-3233331221000130"></a>

Type: `["object", {}]`. Optional.

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

- [no_interception](resources--proxy--reference--group-005.md#canonical-3332023313312012-3012310221032321-3101200211301200-2222112211033303-1033311031230203-0102333033323201-3101321210231120-3233331221000130)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-2300212110133011-2303221131021333-1221330301300022-0231220012001111-0213123123321222-0000030203023101-1203222021230313-0321231013011230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300220130222220-0321012100023110-3312120100301031-2332310123110321-0222103331130303-0122320111223322-2100122213201320-3222032202331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_inside_network` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- site_local_inside_network

<a id="canonical-3121123313321101-0001303002003321-2112130332133322-1102033212123313-1303202113031023-1012333112101221-2212211032231111-0323301233320300"></a>

Type: `["object", {}]`. Optional.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

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

- [site_local_inside_network](resources--proxy--reference--group-005.md#canonical-3121123313321101-0001303002003321-2112130332133322-1102033212123313-1303202113031023-1012333112101221-2212211032231111-0323301233320300)
- [site_local_network](resources--proxy--reference--group-005.md#canonical-3220322333321001-2032010012232201-0231323320030211-1123130120121020-2021030112012033-2213000303013330-0001000213111211-3213131133222023)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303132001322003-3222233301002011-1312320023221130-0032320032223121-2231221200221010-2211212211123222-1202133222332122-0032211113003013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_local_network` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- site_local_network

<a id="canonical-3220322333321001-2032010012232201-0231323320030211-1123130120121020-2021030112012033-2213000303013330-0001000213111211-3213131133222023"></a>

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
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- site_virtual_sites

<a id="canonical-0333011221000001-0111132130112111-0231020012230322-3301122013323311-0123011330302122-0120213032111302-3301221321211312-3112303021033110"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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
site_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323022202332022-3031023120231021-0003033031001021-1122020301302033-3320310220112312-0111323200133230-1231202211130133-3213300132201102"></a>

### Direct properties for `site_virtual_sites`

- [advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002): complete subsection reference.

<a id="canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- site_virtual_sites.advertise_where

<a id="canonical-2100110201320221-3330213111022300-3111200010313311-0331110213021031-2123233022023023-3121012211012322-3222213131320121-0203133310233020"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130122003113103-3330210333021220-0221223003112203-0011321211203331-2301023222021300-1330312121133000-1002312021111102-3130010310133003"></a>

### Direct properties for `site_virtual_sites.advertise_where`

<a id="canonical-2313331113002231-0223313032111320-2121123130113111-2112002101331020-3013211030130110-0320211001020010-0312303030130123-3100331002223113"></a>

#### `site_virtual_sites.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013): complete subsection reference.

- [use_default_port](resources--proxy--reference--group-005.md#canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310): complete subsection reference.

- [virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110): complete subsection reference.

<a id="canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.site` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.site

<a id="canonical-3322321212133033-3111230213233223-2022230112131332-3202102321032031-0110133333110320-3302032210313332-0303331212200032-3310010110232220"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322021000330210-0332101121133223-0202232222321221-2321220300101001-2000013121313121-2232012231203230-0223102323121202-1000321310331133"></a>

### Direct properties for `site_virtual_sites.advertise_where.site`

<a id="canonical-3020123232013323-3210201103102201-2103012332120020-3002111323022032-2000010312231311-3232103100233123-1120022203113223-0110122221133311"></a>

#### `site_virtual_sites.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3332301132001121-2330200212011001-1133312132310230-2123020113231313-3210101012031113-1322322130033302-3210222023131233-1023120311110100"></a>

<a id="canonical-1103323002332032-3023030130113210-3120330320203102-2301210030010120-2102110101012332-0302110120312332-0321000120213102-2101321121301320"></a>

#### `site_virtual_sites.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--proxy--reference--group-005.md#canonical-0022310220020012-1031101221022300-3210031221310121-3032030331230003-0221201323303013-3013333303211123-3130211030221323-3233002033123330): complete subsection reference.

<a id="canonical-0022310220020012-1031101221022300-3210031221310121-3032030331230003-0221201323303013-3013333303211123-3130211030221323-3233002033123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013)
- site_virtual_sites.advertise_where.site.site

<a id="canonical-2103321233003120-2201200120122210-3131312013011000-2031220232012110-0201110232321202-3133001211302030-3132303203030301-1123130121013213"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030122203100232-2300022313100313-3222113130211301-3100102203110113-1311333003132032-2013023032103133-1320203013201213-0300300000302113"></a>

### Direct properties for `site_virtual_sites.advertise_where.site.site`

<a id="canonical-3121310320130220-2331301102112020-1220113322032121-0031302332011122-0020033121333001-0212302110201103-2231100123111333-0232333022132113"></a>

#### `site_virtual_sites.advertise_where.site.site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1010221100200201-2213001313323102-2121021201110331-3202330100130002-3311023221213023-3111003032100031-2330030133310310-1111130201031233"></a>

<a id="canonical-2012022313033300-1111013112123120-2302001002020101-3120321122130310-3010203201021311-3330332113110322-0211002022230222-1331120201121233"></a>

#### `site_virtual_sites.advertise_where.site.site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3133131122003012-1302032100213122-0310302322101230-1222213220330320-1110231123323312-2230211231012031-0120313203331100-2112002210310032"></a>

<a id="canonical-1311201010012323-3002233233212211-3101211023210033-3120011123012122-3213303002303212-2111101100210103-3131213021122301-1233101103323130"></a>

#### `site_virtual_sites.advertise_where.site.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.use_default_port

<a id="canonical-0122120302113102-2312133200030003-0320023221321311-1211202100300300-3232310132131300-3003202022122110-0012312121103301-1103211230011310"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.virtual_site

<a id="canonical-0212121332300203-2110123100212012-1102113201011032-2333332232213032-3333021310230323-0012103113020323-1322021231120310-2011323023331022"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320033202000200-0213131232021113-2022132213111101-3021333222321133-3312012000101013-2023102113310203-0302101230032110-2211000302231212"></a>

### Direct properties for `site_virtual_sites.advertise_where.virtual_site`

<a id="canonical-1023312001111300-1100323013130010-3132301100320232-2101021300021111-2000311030233303-3122230221012130-3023003123201320-2330313321010010"></a>

#### `site_virtual_sites.advertise_where.virtual_site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--proxy--reference--group-005.md#canonical-0210101233323000-1323230321122112-2322203003111110-0221010023212220-2231220113112322-2333133212033103-1212013232103323-3020232130133321): complete subsection reference.

<a id="canonical-0210101233323000-1323230321122112-2322203003111110-0221010023212220-2231220113112322-2333133212033103-1212013232103323-3020232130133321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_virtual_sites.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-0221220331332212-2030311010233103-3111323303330033-1303122020131120-3012203103332302-1021100230031222-2303032031301320-0012210232123021"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031230022203202-3112013130301000-1100332231132320-0032021210200032-1032031033020322-3300303030103233-1203023222030313-1322320320212122"></a>

### Direct properties for `site_virtual_sites.advertise_where.virtual_site.virtual_site`

<a id="canonical-0333011010230103-0210213231232323-3232310311111101-0013203301102123-2231332122102002-1113123133120112-3300230110310210-3311023102322321"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3010310031323201-1202312201220132-3102011121220122-1000310331002102-2111221323121221-3200312222231333-3230132003111322-1213110232113303"></a>

<a id="canonical-0020112021313310-1102220221011011-1201101300223113-0220033132300200-3132020320131312-0012000020200231-3203211020321121-1321322322031031"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1232220200003130-1331033010203112-1233030100221122-0301122111200133-0110300231213002-1203020023003011-0203111220123131-3223201233230230"></a>

<a id="canonical-2002333210113331-2201330213320022-2230301313323013-3202013001103001-0220010000222032-2322221122011312-3313010012212001-3001013230232312"></a>

#### `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2113110122012210-0300103202002221-2012320121222331-2230022312003312-3203002101213002-1312002203112032-1202231032011220-3312233322001132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- timeouts

<a id="canonical-3100221132113223-0221031321312001-3112300110111021-3322203020203310-3012010232232321-0222122130021011-3230012320013131-0303023223020122"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233031011202310-1331021310220132-3012010211322222-0223220230321230-0023220211022201-0130000102121021-2331130200221211-0133133132322132"></a>

### Direct properties for `timeouts`

<a id="canonical-1103120221321111-3002033010233002-3110210222233021-0231311012103102-1123033223001123-3001110112012100-2021320110011201-0300330321102300"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1221303323102020-1110010213210303-2331222111130233-2013220310022102-2323220301220031-0022100223230213-3211022101003231-2032033220301300"></a>

<a id="canonical-3313212000201331-1111011133332023-1303233222330112-1130031311333111-2112002010200010-1122221323011333-0223303131112203-1022132033130212"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2320022221231000-1231021201033201-1023021213332301-1230312111300001-0323033312020123-1023201001033331-1133021011021300-2032110033212110"></a>

<a id="canonical-0022020130320303-1101330230000231-0331231030313333-1021120232233330-3013120300012322-2232321333131222-0000220120203332-2302133310211323"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2232312313013221-0311031132223112-0030011111200223-1103230231301131-3200311300003133-1212112321110313-3121322030010203-3133001203330223"></a>

<a id="canonical-0213210023313011-0132111030310010-1201021210010220-3020003321001320-2302330010313232-2312033003000201-2333320210022301-0201333101021332"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- tls_intercept

<a id="canonical-2300212110133011-2303221131021333-1221330301300022-0231220012001111-0213123123321222-0000030203023101-1203222021230313-0321231013011230"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322103130100103-1122010031120010-0023311333322203-1110310301011131-0120132113000101-0310311023130313-0123230021000031-1200113310101122"></a>

### Direct properties for `tls_intercept`

- [custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223): complete subsection reference.

- [enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102): complete subsection reference.

- [policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103): complete subsection reference.

<a id="canonical-2302310120010303-0110132203322000-0001021231100100-2330023331030131-0023031130300232-3200023320232110-1113013233133323-1310310320311202"></a>

<a id="canonical-1113133301102202-1322030002301202-0110331202111330-1130121212232022-2012101332330023-2213233113121110-1223212300230213-1012130321222120"></a>

#### `tls_intercept.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--proxy--reference--group-005.md#canonical-0322201212300210-3122231010311203-0031222013232003-3121000333001303-2020203321212103-0110333232012202-2000012101030312-2210122003113011): complete subsection reference.

- [volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-0221120012112232-3222102003023211-2332123231332121-1022321130233300-3132233121322211-0210112032300212-1100120323003230-3333123211101001): complete subsection reference.

<a id="canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.custom_certificate

<a id="canonical-0130022112311203-0002331003312132-1033021013303233-2012221301021003-0123002202003302-3320330131212202-3310321213331132-0333330300202030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Additional upstream details:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232021000201100-0301220310032113-0121222121203222-1322133101200030-3022123233002220-3121220120330031-3300313122313133-3332032020113320"></a>

### Direct properties for `tls_intercept.custom_certificate`

<a id="canonical-2201022013301311-0213100320031220-1002032113200202-1310102202002013-2330221000130002-3323111100231222-0301212033232303-1222032322333202"></a>

#### `tls_intercept.custom_certificate.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-3032013211022031-2003020322332212-2232310313202130-1212131023200112-2311110231233322-0213332121331132-2131203121303020-3333330323131203): complete subsection reference.

<a id="canonical-3133130011332203-2030002102133223-2312210333013330-3000211323231223-2002111120220221-1101123210002112-0201001310000103-1232113321303100"></a>

<a id="canonical-2033001111101210-3133133011132120-1220301201301230-1210011132101300-1021023102233231-2010303122130231-2232213230223212-2201223120312311"></a>

#### `tls_intercept.custom_certificate.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-3313223130002100-2133323221002213-2313132023211030-0111231131322301-3031302232031101-1030231322202333-2022231030132030-1201002010230031): complete subsection reference.

- [private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221): complete subsection reference.

- [use_system_defaults](resources--proxy--reference--group-005.md#canonical-0203212312012200-0022331233322101-3312232012110231-2001030223033320-3211223032133022-0222102213210232-0213123221113002-2200213003332120): complete subsection reference.

<a id="canonical-3032013211022031-2003020322332212-2232310313202130-1212131023200112-2311110231233322-0213332121331132-2131203121303020-3333330323131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-0010221002113203-1320123333021111-2131113310302131-2223121311002101-0211030313032012-1213232231303030-0332013011332201-2012220031313001"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2330010231031133-3020122122020232-2213032121233103-1211113302001121-0103311203211321-0311310122320212-3022031023331031-0032020101130111"></a>

### Direct properties for `tls_intercept.custom_certificate.custom_hash_algorithms`

<a id="canonical-1123210200321201-0103130010111311-3011202013331200-2330210113233011-2232220100323300-0310023113112211-2121012020332202-2231231231203112"></a>

#### `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3313223130002100-2133323221002213-2313132023211030-0111231131322301-3031302232031101-1030231322202333-2022231030132030-1201002010230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-0211213312302113-3322133023012121-0200203231303133-3323110331113331-3302203033022223-1013333112320101-2000123033001203-1130200012111122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.private_key

<a id="canonical-2200200320331121-1333131312130012-3322323323023132-0203231323032321-3111332333001312-1230102301032332-3202021313322200-0330202120220030"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232220032100121-2121001102003033-2022103331120100-2203303113232300-2331232110121232-2111302101303122-2231000002310112-3311100201100313"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key`

- [blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-3023221303101103-1211031223130023-0132333022302321-2023220133310312-1021103200000333-2310003100323222-2121103122121012-0230102020230211): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-005.md#canonical-3201000030300312-0322013220130313-0013123021000321-0030321111111122-2212302321121233-2302202131031011-1103212303112102-1223020013122013): complete subsection reference.

<a id="canonical-3023221303101103-1211031223130023-0132333022302321-2023220133310312-1021103200000333-2310003100323222-2121103122121012-0230102020230211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-1303000210122130-0033031020131202-2223232022111303-0101023310321310-0103201102201023-2200101021211100-3303302032331312-2300211330320103"></a>

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

<a id="canonical-0021112131033110-3322211010302012-3103112003333101-1121322010022012-0203201200133020-1022323130220003-1022020023300322-1022332302113012"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key.blindfold_secret_info`

<a id="canonical-2311302231311013-1313022212223011-3320300010200122-2103232330000310-0211000232310332-0220232233231302-2110310213001311-0310203102110013"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2302122000031322-1230230313132022-1102003222122101-2002320132312312-2002103012330122-3310002312111221-1031233102023120-1020121323132022"></a>

<a id="canonical-3310120021211111-1013030301032023-1320323133302213-1230133300223330-1102202330303312-0202233032300322-0130020023230010-3233201033331011"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` property

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

<a id="canonical-3310333120321022-1021132212211032-2331030010132110-0030113122233032-3102312133000121-1111213102300120-0012303330102203-2333331113101201"></a>

<a id="canonical-3233122022310003-3200332122312303-1001232213212100-3011223033320132-0110021111210002-1101012102022201-3133302210220220-1131201332021133"></a>

#### `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3201000030300312-0322013220130313-0013123021000321-0030321111111122-2212302321121233-2302202131031011-1103212303112102-1223020013122013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-3211200113130130-2323212132131212-2010223130002213-1210211300333223-2300220012101110-2313201212010302-2010300002320333-3221010210032210"></a>

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

<a id="canonical-0013001211110031-0203231301133233-2000311033112102-3333001310220031-1320122213221120-3313033131011312-3012031100033210-3311310230032213"></a>

### Direct properties for `tls_intercept.custom_certificate.private_key.clear_secret_info`

<a id="canonical-3220233012210213-0003323133131222-2021100300232131-3201323313110031-0330213312331311-0113212001233312-3233213220032102-3010222023312203"></a>

#### `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2213022103112120-1123211032221101-0033212213001120-1021332101023210-2031003323031010-1220310103320001-0333000210100022-0231120000011000"></a>

<a id="canonical-0321302311203201-1013301023122300-2332322103323301-0130003033202201-2221031331030220-3103030311310301-2023320131333321-1211120322012202"></a>

#### `tls_intercept.custom_certificate.private_key.clear_secret_info.url` property

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

<a id="canonical-0203212312012200-0022331233322101-3312232012110231-2001030223033320-3211223032133022-0222102213210232-0213123221113002-2200213003332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.custom_certificate.use_system_defaults` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-1110203332211003-3020032332111020-3013002222220323-1032201121213320-0022011103023313-3231220113303122-2211023121012311-2000111032121001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.enable_for_all_domains` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.enable_for_all_domains

<a id="canonical-3303033331133302-0302031220133100-3101211302231333-0133120110321330-2112331230232232-2002001213311121-2122333311103001-3031221331320202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.policy

<a id="canonical-3011100031221211-1203111210222303-3201101010103122-1012213130310031-1131102203013211-1200222033133233-3123020100000232-3001100100320300"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
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
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131003201201102-3112222303112120-1212132210233302-3010031011230320-1201303022222201-2010002111000013-0031022123002230-0101232313013320"></a>

### Direct properties for `tls_intercept.policy`

- [interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032): complete subsection reference.

<a id="canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- tls_intercept.policy.interception_rules

<a id="canonical-3312120023110110-1021033230023313-3012032301110231-0312332203030120-1302020330230020-2013110232000331-3202313011323212-0202011032302323"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023100312203222-0200212322012032-3203020002013330-2110120210300132-3323000231322233-0200021222033322-3122221311011103-1102102000103211"></a>

### Direct properties for `tls_intercept.policy.interception_rules`

- [disable_interception](resources--proxy--reference--group-005.md#canonical-1000112331102113-3212211312313230-0302331033011221-1311123131101303-0211200322130033-1332120220320010-2313223123033301-0312113131110321): complete subsection reference.

- [domain_match](resources--proxy--reference--group-005.md#canonical-1023120230321222-1033211013013300-0023103003301010-2100102002020300-0332210232003222-1301032200113301-3222220323213330-2221221100213320): complete subsection reference.

- [enable_interception](resources--proxy--reference--group-005.md#canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033): complete subsection reference.

<a id="canonical-1000112331102113-3212211312313230-0302331033011221-1311123131101303-0211200322130033-1332120220320010-2313223123033301-0312113131110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.disable_interception` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-0131220000203202-2012032110322310-1130312113203001-3202200031223021-2111112001303320-3031213121103212-2212320133201212-1022020211033011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable interception.

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
disable_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023120230321222-1033211013013300-0023103003301010-2100102002020300-0332210232003222-1301032200113301-3222220323213330-2221221100213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.domain_match` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-0222232012012321-1012022110323022-2130020212210131-3222232100323101-0110033332333101-1213131322200231-2331333113132111-1121000230312002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for domain match.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322000310032110-2031000322100331-2031010121030003-2301330011203100-3011003321132131-3303212123311203-1023322233103222-2303300300113101"></a>

### Direct properties for `tls_intercept.policy.interception_rules.domain_match`

<a id="canonical-3201000310002110-3111121001031331-1220221031121301-1313022123030033-0020000223300020-3023033002213233-3030301112233131-0123033223222203"></a>

#### `tls_intercept.policy.interception_rules.domain_match.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-1232322302030023-2132213220131130-0032303303133001-2311010030202200-1301230211311323-2222010132121103-1023231021323013-2031332131031030"></a>

<a id="canonical-3303012203023011-3201120223112203-1132301213232311-3231231130332222-3001023012202130-2230302230133022-0123021111333020-1303330221123201"></a>

#### `tls_intercept.policy.interception_rules.domain_match.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0230330021220112-1120022222202102-1300102121033322-0231001100213030-1111320121100221-3120221233230210-0011102121321221-3002121303032233"></a>

<a id="canonical-2322121323313103-2221321022213030-3001101323311212-0203322133223123-0031301231022103-0103310132102313-1222110211130322-3312121230231002"></a>

#### `tls_intercept.policy.interception_rules.domain_match.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.policy.interception_rules.enable_interception` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-3013000203331130-3003021302221020-2112333323201001-0011311321113103-0002221001030022-2312213230213010-0312201033202321-3231302301120002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322201212300210-3122231010311203-0031222013232003-3121000333001303-2020203321212103-0110333232012202-2000012101030312-2210122003113011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.volterra_certificate` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.volterra_certificate

<a id="canonical-2121122002311101-2001321331222012-1332131220210100-0323202103032100-2011030332222001-3112331133023222-1302120010030002-0311103030022013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra certificate.

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
volterra_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221120012112232-3222102003023211-2332123231332121-1022321130233300-3132233121322211-0210112032300212-1100120323003230-3333123211101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_intercept.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.volterra_trusted_ca

<a id="canonical-1013120230302223-0033003022013000-2223010222321303-2021123331333122-1122330221121323-1302122111221013-3302102203012022-0003002030322113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.
