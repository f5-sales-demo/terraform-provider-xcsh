---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-0022222203032012-1211203122122220-3001323100103221-0211202220012030-3231131312133230-0112212103313022-3323101021302213-2231101332223032"></a>

## `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-2302201103321222-2232302000123302-0223031030121111-3212223103132320-0130111302302100-1302203200002113-3010120122303201-0003302102233031"></a>

<a id="canonical-2102313020230232-0022332331112313-3122312300222220-1301020010331010-1020003300212232-1033212210033032-3212033103213020-2302221300030322"></a>

## `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0332020022011013-0132312011111202-2313031230311313-3200013110320001-3202013302013121-0330000023331320-0222300302312031-1103301330312022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-1222311103011223-2023210230230212-3221022312223230-2011000303333233-2201120130230001-0322033000011011-0012020012001121-1313001212301020)
- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-0131021101021111-2020001221122310-0130013213011131-2010320233202120-0230331100012010-0223112033311130-3122300203102331-3031102030131230)
- routes.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-1302010322033002-0002023011311213-3003132213212323-2023130302230230-1102201002332113-1230000100121202-1011322023000122-3022311123223123"></a>

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

<a id="canonical-0312132021210013-0332210023101010-3023011023012002-3232122201112001-1311222233132113-3000111013313312-1122330323022213-0332220312303303"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0301322020023310-0332102310303010-0202011321331203-2133321000312211-3100220122223323-3020312100013103-2332330020113112-2232103011312203"></a>

#### `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3202200100333113-1210221333322311-0210331330100200-2331332211120320-0230102210130301-0221210122111221-0223000213112332-3000101033131221"></a>

<a id="canonical-3220301130313323-3131311203210331-1010011123112201-2032223202201030-3102002312223333-1010222111020300-0110230202001111-0301202212332001"></a>

#### `routes.request_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-3110130201033312-1313023321233221-3201132303201121-3020210012322131-2100011331031103-0301020201111323-0032131300201003-1021223233210133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.request_headers_to_add

<a id="canonical-1132203010013323-3020310323133333-0001320310213001-0031303131120100-0302301011123303-3313311112123211-3131210123100211-0312001013002032"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333101132012121-3311102300030111-2212311233130322-2031233333331113-3223123222110223-2222130133110023-0232112131232102-3123031022112113"></a>

### Direct properties for `routes.request_headers_to_add`

<a id="canonical-1231110302120301-2213203232101212-2102222111312110-2313330002003003-1131323013322310-0222123110231133-3321112323113032-1320121020301333"></a>

#### `routes.request_headers_to_add.append` property

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

<a id="canonical-0331233131112102-0233330301212023-0300020101030103-1332221133212303-0313200121312032-1100231232322130-2122313212321111-2323120023002301"></a>

<a id="canonical-3323232020012103-0312030312101302-0330221131122032-1233313003310130-0303102331000102-2120211132010213-2103113331223002-0211220322113303"></a>

#### `routes.request_headers_to_add.name` property

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

- [secret_value](resources--route--reference--group-002.md#canonical-3223121100010231-3022121123131313-2010112103000010-0312121303131300-1312010132332302-0302033003332022-1113231302232222-2311300200203311): complete subsection reference.

<a id="canonical-0332212132012103-2100320021230031-1111011011101123-0331311213220302-2102100013302310-1003113213022310-2332222313311032-3323320120022121"></a>

<a id="canonical-3121002122211010-2122012020003120-0133030131111133-3312330021321023-0331220132310030-0323001123321323-3212310003032201-3113120111000232"></a>

#### `routes.request_headers_to_add.value` property

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

<a id="canonical-3223121100010231-3022121123131313-2010112103000010-0312121303131300-1312010132332302-0302033003332022-1113231302232222-2311300200203311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_headers_to_add](resources--route--reference--group-002.md#canonical-3110130201033312-1313023321233221-3201132303201121-3020210012322131-2100011331031103-0301020201111323-0032131300201003-1021223233210133)
- routes.request_headers_to_add.secret_value

<a id="canonical-1102202003112002-0303233213222201-0231003212122100-2133112130320232-0333332013010200-0201030201023002-3131112211310112-1130121110131012"></a>

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

<a id="canonical-0203213020330021-2211232032232313-0212020000022023-0001323301311300-2322303230110003-1000222020232231-2012223213322032-3222122012301231"></a>

### Direct properties for `routes.request_headers_to_add.secret_value`

- [blindfold_secret_info](resources--route--reference--group-002.md#canonical-0022010210312122-3102223122222003-3232212203313020-1321013300221120-2201221232313300-2002332122223201-3200121000310222-0131302310302023): complete subsection reference.

- [clear_secret_info](resources--route--reference--group-002.md#canonical-0103123010331121-3033320122301333-0303110232213203-3012312232302232-2203020331233302-0212321001023202-2001301220221222-2312003100213232): complete subsection reference.

<a id="canonical-0022010210312122-3102223122222003-3232212203313020-1321013300221120-2201221232313300-2002332122223201-3200121000310222-0131302310302023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_headers_to_add](resources--route--reference--group-002.md#canonical-3110130201033312-1313023321233221-3201132303201121-3020210012322131-2100011331031103-0301020201111323-0032131300201003-1021223233210133)
- [routes.request_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-3223121100010231-3022121123131313-2010112103000010-0312121303131300-1312010132332302-0302033003332022-1113231302232222-2311300200203311)
- routes.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0021322313122300-0321132321320300-0020313230232201-3331033012103101-3320011121023112-2111110232031330-1133102012111011-0031001110000201"></a>

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

<a id="canonical-0323013123132120-1321302322110100-1123330320022022-1331020322301322-3322010030022000-0031002023302113-0122220101221030-2031122220321323"></a>

### Direct properties for `routes.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1022103121000233-2113100133221211-1230301211132003-0220321002133311-0011332011330122-2312112332110012-1102203200333223-1311110302223010"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2202023223011203-0300033102221001-1212210230300210-3033221120233001-3030231223112300-1122020312102130-0121311122222223-1230003010320111"></a>

<a id="canonical-0031103013113223-0000013223310121-1033230002020111-0230321333130030-3002120031331022-1332321031021023-3102211032233121-1213232302120311"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3313021302012222-1332203013331332-0333002110322013-2311113122300030-3121121010100321-2313133002021320-2311131112332013-3312230011213323"></a>

<a id="canonical-0132311200323003-2031200320220300-1002113010121133-0121321310302011-0202213123323032-2300133103100331-2110231123102210-3003232100033211"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0103123010331121-3033320122301333-0303110232213203-3012312232302232-2203020331233302-0212321001023202-2001301220221222-2312003100213232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_headers_to_add](resources--route--reference--group-002.md#canonical-3110130201033312-1313023321233221-3201132303201121-3020210012322131-2100011331031103-0301020201111323-0032131300201003-1021223233210133)
- [routes.request_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-3223121100010231-3022121123131313-2010112103000010-0312121303131300-1312010132332302-0302033003332022-1113231302232222-2311300200203311)
- routes.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2200231113331102-2320203122111000-2112201132211231-0110032320110223-3210033112102020-1202113221210201-0113120101300221-3231002333021303"></a>

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

<a id="canonical-2332310203110312-1110321201200211-3132322202223021-1301010310130302-3021222000123122-1130323013220132-3022202121102333-1301033003310123"></a>

### Direct properties for `routes.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1323023311113212-1330013111012322-3312322202033110-3201110322233301-0300012301333020-3130212310303220-2032203311220012-1012322333212001"></a>

#### `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3230023032133231-2233013032333331-0320332201203330-0020000010012131-0023000033002100-1121320030012231-2211313221230012-1210233011103113"></a>

<a id="canonical-2132133302002201-1033001001023320-1203330123100232-0330003203303330-3020330033011331-0311302101023331-2012210022223020-1010302032133110"></a>

#### `routes.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.response_cookies_to_add

<a id="canonical-0013333121022032-2312111032311133-0220132002003233-1123020133213121-1010302332221333-0213200110232021-2231311110031332-0101201033113330"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream.

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

<a id="canonical-0100321023332011-2011301120303313-2213033221113333-3331232302223303-3121330331330323-0131203111122221-0031130220021022-1220122110033220"></a>

### Direct properties for `routes.response_cookies_to_add`

<a id="canonical-1230301332112030-3323133021030012-0223202322110103-2033022130131332-2203220102022011-1103220023320032-2322321210330302-2122220323333131"></a>

#### `routes.response_cookies_to_add.add_domain` property

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

<a id="canonical-2311312331013221-0010302332221202-0031211210303111-1021100023221321-3020313231031200-1310112130201302-0023201231232203-2103123231323020"></a>

<a id="canonical-3032102233211032-2210130032331332-2012232131301000-0313022100130223-0322030120332031-1112011310302212-3213302033121110-0322013133012222"></a>

#### `routes.response_cookies_to_add.add_expiry` property

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

- [add_httponly](resources--route--reference--group-002.md#canonical-1103221012331203-0122010210313203-3100132001130311-2320012210012033-2130100302232200-1203000300120310-2130101020231122-1203133333321221): complete subsection reference.

- [add_partitioned](resources--route--reference--group-002.md#canonical-0333013013231223-0032201201001000-3103131330223020-0021203131010320-1121223010033203-0313233101012002-1222300213021003-0303233202212210): complete subsection reference.

<a id="canonical-2210220010112201-3300213012203231-0211333312112032-0222122321312012-0103121222331313-2230212311003123-3203230030130032-0311033212032123"></a>

<a id="canonical-1111321201330230-3301000130032131-3010100032221110-1310311302211133-2011221331122033-2223202213033022-0133130300132201-0231223132232302"></a>

#### `routes.response_cookies_to_add.add_path` property

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

- [add_secure](resources--route--reference--group-002.md#canonical-3133223103001203-3011023331133320-2222201122330001-1331313311301031-1022122132310201-2220322211230132-2301032022133320-1202020210023003): complete subsection reference.

- [ignore_domain](resources--route--reference--group-002.md#canonical-2221323220301102-2322102303133302-2112332210121321-1201020132132122-0101221001000313-2211211002102201-1101020333230313-3003012012232300): complete subsection reference.

- [ignore_expiry](resources--route--reference--group-002.md#canonical-0213330323021101-2110302131013102-1232011130102302-3211102330223103-0103121113203002-3013133123231332-0023101011311132-3231103213013133): complete subsection reference.

- [ignore_httponly](resources--route--reference--group-002.md#canonical-0000211202200123-0231132233013113-3220031031130001-3012312023221010-1011210121113131-1320211203321302-3110133313122202-2121023323110030): complete subsection reference.

- [ignore_max_age](resources--route--reference--group-002.md#canonical-0033212230013331-1102230233220332-1212320231303132-1322221130323000-1103222010031332-1202321113113033-2331101210223311-2202002112202201): complete subsection reference.

- [ignore_partitioned](resources--route--reference--group-002.md#canonical-0100221321223333-2030102321133131-2320233110231023-0232103323133110-1021120212332011-3320100101300330-2332301233200331-1221112322300221): complete subsection reference.

- [ignore_path](resources--route--reference--group-002.md#canonical-0032022323323102-1203111102122010-0111010121230223-1302233033130310-1033010101231222-0313020111311212-2013023200122211-3313320231123021): complete subsection reference.

- [ignore_samesite](resources--route--reference--group-002.md#canonical-2001111123100221-1302133330212132-3011003033303201-2033332212303031-1221000223230313-1322223123031302-0231331011110000-2310010333212111): complete subsection reference.

- [ignore_secure](resources--route--reference--group-002.md#canonical-3100023311333323-3213312201000321-2120003332232211-1213302020012020-3033132213101220-1012032103212011-2033311211012100-1020333301322121): complete subsection reference.

- [ignore_value](resources--route--reference--group-002.md#canonical-3033122223002313-1330020030130120-2223131110212100-2220201010103320-3213031123230021-2320003302203212-0200320133012320-0331023323123222): complete subsection reference.

<a id="canonical-2322130201023212-2301122001100031-2103100220020101-1212111300211311-3132110332002330-1013210333301322-2201122200003032-0213303220022122"></a>

<a id="canonical-0333013210331303-3133102030023331-2221202120103302-2003213303220332-1011210031333311-0131022201001002-0021120112200211-3021310002032201"></a>

#### `routes.response_cookies_to_add.max_age_value` property

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

<a id="canonical-2202011030200302-0022031032312002-2100203101101030-1011123220031213-3013312131203122-1120323322223113-2333111232122013-1010113232222132"></a>

<a id="canonical-0023233032233021-2232220112322120-1103310233311101-0301003210110030-3001101033223231-0103310020311033-3230100100100202-2002301032321221"></a>

#### `routes.response_cookies_to_add.name` property

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

<a id="canonical-2131230100112111-1223133320131030-3220222201212303-1231103121113322-1132312112213202-1003020200203002-2001311302231003-1301021011021302"></a>

<a id="canonical-3132232023232333-1012213031212322-3112220331323022-3320302200333010-0001203011302333-2301010012030322-3001221220002032-2220321031003112"></a>

#### `routes.response_cookies_to_add.overwrite` property

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

- [samesite_lax](resources--route--reference--group-002.md#canonical-2302021110101102-1132000320120333-0213312131331100-1230300030001333-0132012033302312-2100320011100330-2030312311222121-3333002131112212): complete subsection reference.

- [samesite_none](resources--route--reference--group-002.md#canonical-2330133002220000-3003313022331120-0033310120120320-1122000301101120-1001103202102212-2020203302112022-1011220300331333-3302322131202121): complete subsection reference.

- [samesite_strict](resources--route--reference--group-002.md#canonical-3313220212220022-0321001311132031-2230003232320310-3223320111023203-2311132313123202-2133231203123223-1111032230223302-2033220323321321): complete subsection reference.

- [secret_value](resources--route--reference--group-002.md#canonical-0103203110121200-2121220322032313-0233301211311330-2311322312210122-3313130023102202-2101033321222231-0202232321331000-3331200020123000): complete subsection reference.

<a id="canonical-1113130012332221-0201333220100130-3031212102223232-0031332010031031-3310211011231132-2231201321303003-0100221120220302-2210332332030133"></a>

<a id="canonical-3112222123122201-0222201332233101-2120122113101112-1320002233021101-2301102333132200-0321233000310013-0113003311132022-0133030230121332"></a>

#### `routes.response_cookies_to_add.value` property

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

<a id="canonical-1103221012331203-0122010210313203-3100132001130311-2320012210012033-2130100302232200-1203000300120310-2130101020231122-1203133333321221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.add_httponly

<a id="canonical-3313121310131313-1200232200002320-3010311302311013-2122300200132233-3303133131303232-1130131310212211-0332311211221100-3201310101003223"></a>

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

<a id="canonical-0333013013231223-0032201201001000-3103131330223020-0021203131010320-1121223010033203-0313233101012002-1222300213021003-0303233202212210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.add_partitioned

<a id="canonical-3202102133211301-0232112122021123-2120010110011222-0030232300021313-1332102101311330-0023100303232031-2033120102210110-2230303000202322"></a>

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

<a id="canonical-3133223103001203-3011023331133320-2222201122330001-1331313311301031-1022122132310201-2220322211230132-2301032022133320-1202020210023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.add_secure

<a id="canonical-2020200201210230-2101132301023322-3021310033200022-3012210121020011-3331330033132012-3133023231120033-0211312033023300-0332321220103220"></a>

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

<a id="canonical-2221323220301102-2322102303133302-2112332210121321-1201020132132122-0101221001000313-2211211002102201-1101020333230313-3003012012232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_domain

<a id="canonical-0322223013231320-0112121303012213-2033222310022210-3220320033300333-3303222121233100-3011031111131100-1230212322003001-3333201222212323"></a>

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

<a id="canonical-0213330323021101-2110302131013102-1232011130102302-3211102330223103-0103121113203002-3013133123231332-0023101011311132-3231103213013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_expiry

<a id="canonical-0223311001021311-3201101231101121-2123120013223303-1122232102103232-3310102311000001-1320100011330130-3222121133222001-0100131101111233"></a>

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

<a id="canonical-0000211202200123-0231132233013113-3220031031130001-3012312023221010-1011210121113131-1320211203321302-3110133313122202-2121023323110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_httponly

<a id="canonical-1010300020330100-2003321200002321-2333210303203230-0313223021022011-0330330203323022-1302212210130212-2131312331320002-1120221002201310"></a>

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

<a id="canonical-0033212230013331-1102230233220332-1212320231303132-1322221130323000-1103222010031332-1202321113113033-2331101210223311-2202002112202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_max_age

<a id="canonical-3311130300323131-2302231310000213-0132310110303020-2302122010120332-0230332103213320-1230231211010212-1101110010012130-1121221312300001"></a>

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

<a id="canonical-0100221321223333-2030102321133131-2320233110231023-0232103323133110-1021120212332011-3320100101300330-2332301233200331-1221112322300221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_partitioned

<a id="canonical-1302212122310333-3231222220112232-0231100213102021-2102211133203000-2313310020310032-1303021323000312-2113202123021233-3320201131113302"></a>

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

<a id="canonical-0032022323323102-1203111102122010-0111010121230223-1302233033130310-1033010101231222-0313020111311212-2013023200122211-3313320231123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_path

<a id="canonical-3333332031312010-2100013032012213-3121312000030000-0130201031012103-2003011303111202-1301222101002223-1230310123330003-1320002102332312"></a>

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

<a id="canonical-2001111123100221-1302133330212132-3011003033303201-2033332212303031-1221000223230313-1322223123031302-0231331011110000-2310010333212111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_samesite

<a id="canonical-1103213222110232-3130201200113123-1233330132111030-1310030303322311-1230231322033333-1320113203200021-1111120132330030-1323000230311022"></a>

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

<a id="canonical-3100023311333323-3213312201000321-2120003332232211-1213302020012020-3033132213101220-1012032103212011-2033311211012100-1020333301322121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_secure

<a id="canonical-3020212013112210-2331303102032023-3102132112200031-3233311121300111-3032111233223102-2331001120332101-2221132202000111-0000310020011030"></a>

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

<a id="canonical-3033122223002313-1330020030130120-2223131110212100-2220201010103320-3213031123230021-2320003302203212-0200320133012320-0331023323123222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.ignore_value

<a id="canonical-3103133231111101-0121120331100222-0221023311211012-1231330311332331-1131002321023101-1110330321111032-1301211310231300-2321100301301000"></a>

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

<a id="canonical-2302021110101102-1132000320120333-0213312131331100-1230300030001333-0132012033302312-2100320011100330-2030312311222121-3333002131112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.samesite_lax

<a id="canonical-0112001302022023-2323303332111330-0002201230102212-1320000221312213-2011131012232202-2010010010101311-1212121230220012-1223212110131333"></a>

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

<a id="canonical-2330133002220000-3003313022331120-0033310120120320-1122000301101120-1001103202102212-2020203302112022-1011220300331333-3302322131202121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.samesite_none

<a id="canonical-3321302310202120-1122120133321121-1123033221223100-1122123010020133-2213003320313332-2201022130213003-1123013230223303-3222332211120031"></a>

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

<a id="canonical-3313220212220022-0321001311132031-2230003232320310-3223320111023203-2311132313123202-2133231203123223-1111032230223302-2033220323321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.samesite_strict

<a id="canonical-0130210212200013-3201301113032300-1323301330210223-1323001233330321-1311113121231213-1110100031032202-2020112132000122-2132121122333102"></a>

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

<a id="canonical-0103203110121200-2121220322032313-0233301211311330-2311322312210122-3313130023102202-2101033321222231-0202232321331000-3331200020123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- routes.response_cookies_to_add.secret_value

<a id="canonical-3333012202231210-1213231003113113-3013110202200021-3130213311331223-1302020012113132-1201311223130132-2212110302232231-0222130212021331"></a>

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

<a id="canonical-3313122203232221-2323131323321313-3213011102222301-3003131022102102-0001020102333320-3121022111023002-1322033022311131-2331020301321111"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--route--reference--group-002.md#canonical-1033322232113112-2203302230313223-2001020032101213-1032022133223301-2121211203213003-0112333333100332-1021012112112210-0202133133101300): complete subsection reference.

- [clear_secret_info](resources--route--reference--group-002.md#canonical-0303322133000203-2133233013312300-0103023230133112-2231220232232002-0110031213312230-1030203320222103-0001220032301212-0212221211201133): complete subsection reference.

<a id="canonical-1033322232113112-2203302230313223-2001020032101213-1032022133223301-2121211203213003-0112333333100332-1021012112112210-0202133133101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- [routes.response_cookies_to_add.secret_value](resources--route--reference--group-002.md#canonical-0103203110121200-2121220322032313-0233301211311330-2311322312210122-3313130023102202-2101033321222231-0202232321331000-3331200020123000)
- routes.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0303113223102022-1013222213232312-0033331310201200-0033112110002322-3223123233032110-1200133110123002-0012323302012302-3130221113333333"></a>

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

<a id="canonical-1002223222001233-2210333322332011-0323122333332202-3230332330120123-3023001313203030-0130103203123011-1321023013120102-2203333022013302"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1320200232112221-0310101311132113-2322003033123213-0322221132003033-0123112030201200-1213120122120301-1033031003103310-3102020010312201"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3332321000030013-0030022113000333-1333311132312303-0232023031120032-1231311100001033-1033002232322033-2130113231021213-2200122120313132"></a>

<a id="canonical-3110212120330322-1000323110332211-0023313023320210-1001103122210020-0211130210201210-3131020102313012-2122120311322211-0010130130312312"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-0203311012021020-1201000121031203-1200332103120323-1203232112100133-2100211210103213-3032021203110000-3331122100100030-0010002102030023"></a>

<a id="canonical-3111130301033012-3010330220021222-2330220002103330-2130103130322022-0213101300101110-2313133303022030-2233233132221000-2121313100023212"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0303322133000203-2133233013312300-0103023230133112-2231220232232002-0110031213312230-1030203320222103-0001220032301212-0212221211201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100)
- [routes.response_cookies_to_add.secret_value](resources--route--reference--group-002.md#canonical-0103203110121200-2121220322032313-0233301211311330-2311322312210122-3313130023102202-2101033321222231-0202232321331000-3331200020123000)
- routes.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0100330301122212-0332002202023200-1120220132321303-1310102230331011-1212333232121103-2130322020000333-2301100033200120-3122301303010320"></a>

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

<a id="canonical-3023203130313322-1211122132210031-3230013001133232-0003321000000212-2221322210323332-0000230300120112-3023001020323031-2210003010331322"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0222311132220211-0323221102202232-2032203211123012-1220013113333130-1313213133223213-3103003212332220-2030111111332122-1020123323102010"></a>

#### `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1211100222030223-0303102213221322-1002013303011131-1033012210112222-0212001320021131-3201011233003132-1131031300002120-2300102103302212"></a>

<a id="canonical-2010322131211211-1321030020130230-3021233021131230-3001322330311032-1311233001030221-3132003211000031-3033301200012313-3213113212233123"></a>

#### `routes.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1033033030111003-2022222120022003-1123301032332002-1313030103020021-3112122133001003-0321311203111100-2123212000030110-2132031210122220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.response_headers_to_add

<a id="canonical-3012321031120202-3202102101021221-0203033222130113-3111023000220022-1031202222231233-0033100130113123-1033113333112101-2231222023203000"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223213003300003-0213023302233302-0111323002121212-0001233222110211-0113312332011222-3230223131003002-3123302122302022-0321112222322201"></a>

### Direct properties for `routes.response_headers_to_add`

<a id="canonical-0323103003213112-2030130102321021-1012120012203301-2300223320203020-0022202013333102-0000201032211230-2033100102312111-2222322003021312"></a>

#### `routes.response_headers_to_add.append` property

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

<a id="canonical-1121120323000111-2321222301100202-0332221131102001-0132122300111123-0131302113302303-2332201111231203-3332030102010331-1321103131013100"></a>

<a id="canonical-2231201213333233-0012220230322230-0232303121232221-2030102102020330-3122330213100110-1013013002213232-2233010330130133-1331320131102221"></a>

#### `routes.response_headers_to_add.name` property

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

- [secret_value](resources--route--reference--group-002.md#canonical-3010111232312331-2003220221321131-0322032331220012-3211121312110333-1002012232113320-1010130133031122-0013112202333213-0232202132220103): complete subsection reference.

<a id="canonical-0331310331200301-1301313020030203-2131321332222313-1312210200103233-0321133212123303-2103101120133330-2213023130013113-0202322201100130"></a>

<a id="canonical-2103001220031100-0203020011100331-0130331212021013-0000011311303133-0022313113221320-1001121303220300-2133033031133111-1211330312033200"></a>

#### `routes.response_headers_to_add.value` property

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

<a id="canonical-3010111232312331-2003220221321131-0322032331220012-3211121312110333-1002012232113320-1010130133031122-0013112202333213-0232202132220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-1033033030111003-2022222120022003-1123301032332002-1313030103020021-3112122133001003-0321311203111100-2123212000030110-2132031210122220)
- routes.response_headers_to_add.secret_value

<a id="canonical-0000022131220300-0121321101113311-3213311012133020-1033322222111313-3022201321012111-0202323131303130-0323232032320023-2322012002000131"></a>

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
