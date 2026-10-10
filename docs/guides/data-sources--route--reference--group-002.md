---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-0020121133012230-3021000322312112-2000301212110101-0232120123123121-1233000233203221-2213312130233232-3120100112111032-1021020312231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.request_headers_to_add

<a id="canonical-3003023133133313-2311002201102312-1303313131123220-2130233111202333-0211202201202300-2020102303201032-1320112201113131-2312221010222200"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP requests being sent towards upstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

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

<a id="canonical-0310222233010111-0201233333032313-1321303102301000-0333130232210230-1321010331231100-1222113331230320-1233030110200203-2000210301133211"></a>

### Direct properties for `routes.request_headers_to_add`

<a id="canonical-3220133111130133-2202301000232330-3133221102102320-2010002232311312-0203330322113312-3111202032123102-2022201323011331-1320120232011301"></a>

#### `routes.request_headers_to_add.append` property

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

<a id="canonical-1033321013023132-1120002130323302-0312303112200131-2021322002022003-0032113320010030-3203113101213322-3331301001003001-2300301103300032"></a>

<a id="canonical-3222323023113121-2022133231012223-2302130132103312-3310010303112030-0032311332331013-1321202122223331-0121101123310333-1203130033031030"></a>

#### `routes.request_headers_to_add.name` property

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

- [secret_value](data-sources--route--reference--group-002.md#canonical-0211000021002133-2200333200112330-1003321112000133-3112000303313301-1022100211101110-1233121020213122-0231032100231211-3133302022010000): complete subsection reference.

<a id="canonical-2000130020132212-3001011003110202-2130322010023023-2030011012302002-0023301011122333-2330112011122231-1322230123320302-1312231313110131"></a>

<a id="canonical-0131003023000030-2303212012120023-1233311032221002-2221013011001223-0321010320111001-0120112232003323-0001012311001100-2131100331300232"></a>

#### `routes.request_headers_to_add.value` property

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

<a id="canonical-0211000021002133-2200333200112330-1003321112000133-3112000303313301-1022100211101110-1233121020213122-0231032100231211-3133302022010000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_headers_to_add](data-sources--route--reference--group-002.md#canonical-0020121133012230-3021000322312112-2000301212110101-0232120123123121-1233000233203221-2213312130233232-3120100112111032-1021020312231033)
- routes.request_headers_to_add.secret_value

<a id="canonical-3320233012100023-2300202310101300-2032102123001211-2313002301123002-2230122230221131-2022032013301032-3123331033013103-2031113033232321"></a>

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

<a id="canonical-2212020021313133-1122310132113102-1333300202333010-2312122322030113-1112310333231031-3301231030312013-1121203201330001-2303031022130132"></a>

### Direct properties for `routes.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-2321012022211010-1002212103223111-1332213023212032-2302313123022102-2301312233302130-2000301000110111-3223312210201011-3131203121132331): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-002.md#canonical-3031111202131023-1231111311220330-2001103301321201-1221233120030321-0311210111021012-3110110022002023-0322221213312323-2120102002120003): complete subsection reference.

<a id="canonical-2321012022211010-1002212103223111-1332213023212032-2302313123022102-2301312233302130-2000301000110111-3223312210201011-3131203121132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_headers_to_add](data-sources--route--reference--group-002.md#canonical-0020121133012230-3021000322312112-2000301212110101-0232120123123121-1233000233203221-2213312130233232-3120100112111032-1021020312231033)
- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-0211000021002133-2200333200112330-1003321112000133-3112000303313301-1022100211101110-1233121020213122-0231032100231211-3133302022010000)
- routes.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1000121202221302-0230031232230112-3200331313130111-0121100022303033-3110311122102002-2001000122231101-3030101101130123-0013110111212030"></a>

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

<a id="canonical-2210021023030100-1120322002030031-1011021111103203-1210330322210001-1130200202230303-1320303003130200-2003332033112302-0111031303332330"></a>

### Direct properties for `routes.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3231023210130201-1123003312301331-2321021031031001-0111202002310321-0002033011131133-0222133323030200-0200211123132020-1303023332100001"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1132313222210321-3230200220313111-0310030022312313-3111311220333031-0210322213222110-0031333003230031-2201303330100303-0100031130123301"></a>

<a id="canonical-3220133231033021-0313232130200121-3111303333330301-1230101322201033-0203313002030120-2213011013322201-3121030211233101-1100330112300221"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-0231021011133121-2201002100121032-3212123011313323-1131331233312031-0100333330102133-1333023210200132-2013300111201313-3311131322001312"></a>

<a id="canonical-0212313331003103-3203302201010013-0100103103300320-0212230101303130-2320210010222211-0132131103302020-1323321120333033-3302121031211112"></a>

#### `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3031111202131023-1231111311220330-2001103301321201-1221233120030321-0311210111021012-3110110022002023-0322221213312323-2120102002120003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.request_headers_to_add](data-sources--route--reference--group-002.md#canonical-0020121133012230-3021000322312112-2000301212110101-0232120123123121-1233000233203221-2213312130233232-3120100112111032-1021020312231033)
- [routes.request_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-0211000021002133-2200333200112330-1003321112000133-3112000303313301-1022100211101110-1233121020213122-0231032100231211-3133302022010000)
- routes.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3221123313112322-2112011033201222-1231132311322211-0222001110303032-1032023130121023-3022312310203102-0102312221331131-2210323302211011"></a>

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

<a id="canonical-0221332203301330-1232120212220313-0102323302233121-3033232313011103-2332112110212023-0031300123020300-2131302322231003-2103323221203013"></a>

### Direct properties for `routes.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3220203131213120-1211302213220111-1130211303110022-2301100231311102-3311201113123200-3102123002113200-3101021320001311-1323003223121102"></a>

#### `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3030202021321213-0300031113130021-2023321012311013-1323023123021013-3223001332010121-0231200332220020-1203122231201202-0103331033031302"></a>

<a id="canonical-2330222010220302-3221001133003120-1003120313001031-0202023232130130-1300011131122202-0111213112301133-3332322021311333-0303213121200330"></a>

#### `routes.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.response_cookies_to_add

<a id="canonical-3202023130203030-3021232130210002-3010232023121333-1310232103101313-1321310113112223-3210010210212202-1020203032301213-2120313322120233"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream.

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

<a id="canonical-3301301313020030-1013313211323033-0023002122132311-0230211110000322-0210222032112232-1322121010220233-1230333200233232-0302223102102030"></a>

### Direct properties for `routes.response_cookies_to_add`

<a id="canonical-3100022203120300-0132303223300112-2330032112001102-3203310332131202-1300121222130201-0233200112032000-3002333031300321-3023021110103202"></a>

#### `routes.response_cookies_to_add.add_domain` property

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

<a id="canonical-3221323212232013-3300331313322302-3101100101111223-3311321000310301-3122312331121322-1221021203123201-1311233303213003-0302222100113030"></a>

<a id="canonical-1030201223131003-1003133001200222-2203132013213123-0011101300310301-3133332201131123-1100000012330303-2102313002012321-2130121001030132"></a>

#### `routes.response_cookies_to_add.add_expiry` property

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

- [add_httponly](data-sources--route--reference--group-002.md#canonical-1202211010312303-2310011020212210-1121312302220221-0233303200310102-0103231322322102-1030032301312300-3323011213010123-1111310203312203): complete subsection reference.

- [add_partitioned](data-sources--route--reference--group-002.md#canonical-3121111210011022-0323101203033301-2313313320030121-0312231112012203-1200021000210222-2232000320211032-1303022212013221-0320210111111302): complete subsection reference.

<a id="canonical-0330102120000112-1212001202001313-2023312113013232-0233123003201030-1012013210033011-3331311321223021-3231321021203223-2003212110211212"></a>

<a id="canonical-2021101333032031-2333320222221311-1312222233203133-1113111133123311-0101333303132113-2223132312302002-3122213003033231-0230221303323213"></a>

#### `routes.response_cookies_to_add.add_path` property

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

- [add_secure](data-sources--route--reference--group-002.md#canonical-2002130122230223-1111011302313321-0031223100321013-0202210322332230-0110103312323113-0120233231133221-2330201321220300-2013033023003311): complete subsection reference.

- [ignore_domain](data-sources--route--reference--group-002.md#canonical-3102121031321222-0031131130021121-0220220103100321-3020331110020031-2220032130312103-0223112133213303-0310110301120102-3031132001212001): complete subsection reference.

- [ignore_expiry](data-sources--route--reference--group-002.md#canonical-0220313101112103-2211213002100113-2130031103211303-1301111003312123-1002101322300332-3113320230220223-1330212013331022-1131311133131321): complete subsection reference.

- [ignore_httponly](data-sources--route--reference--group-002.md#canonical-0013022112323203-0331000012320222-1230131310223302-1201121001001320-1202012012123203-3123220310001200-1011210313111203-1133001113223013): complete subsection reference.

- [ignore_max_age](data-sources--route--reference--group-002.md#canonical-3203321220330020-3321221220310032-2112023333110120-1101313311220321-2003021111332311-3022201023110223-3111122330310223-1003131111021331): complete subsection reference.

- [ignore_partitioned](data-sources--route--reference--group-002.md#canonical-2232021210122201-3001033322020112-0012331323102112-0320100003303321-3120321011132331-2300023111320100-3303132020123023-0123132031312313): complete subsection reference.

- [ignore_path](data-sources--route--reference--group-002.md#canonical-3030211112231030-1032002301101013-1303313212112233-2121233022022030-0301230032313231-3321201202120032-1211013133013121-1020000021320211): complete subsection reference.

- [ignore_samesite](data-sources--route--reference--group-002.md#canonical-1222213112100231-2323103213120232-3023321021200220-3321120212331022-0323111231110312-3300130122321012-3312302232010013-2132133200002103): complete subsection reference.

- [ignore_secure](data-sources--route--reference--group-002.md#canonical-1302031231221032-1123331002113303-1332232132323311-0212132211211133-2020111320200020-1220121320013321-2020133210111320-2223121003022201): complete subsection reference.

- [ignore_value](data-sources--route--reference--group-002.md#canonical-3333332022113210-2132312032332302-0233013330212002-0312323211020302-0200300330010301-3121031233332210-2300333112313120-3013202132032220): complete subsection reference.

<a id="canonical-3312220203003201-1202332322121102-3101202332310032-0033020112330130-2113322331022211-2012013121130311-2313032132203030-3321213130032303"></a>

<a id="canonical-1031323002210313-3122030232100121-2210333132001101-3003111102311130-0320111002220313-2123211110323323-3102213120332012-3332001012203323"></a>

#### `routes.response_cookies_to_add.max_age_value` property

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

<a id="canonical-3133133202210310-0333131000220312-0130002212230020-2202010212010310-0103311003301333-2213012011122003-3122220031100310-1033323331233121"></a>

<a id="canonical-0031110233001010-3010303101200021-3230310023221231-3230330303333311-2103321331312000-2002032322010011-2221110031001231-0213000103112010"></a>

#### `routes.response_cookies_to_add.name` property

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

<a id="canonical-1110300321121313-3321333231101102-3112202102312222-2101333031130333-0233320332121122-3103333122033011-2101122320303122-3221022022321320"></a>

<a id="canonical-1013113301102031-1030103102221313-3301201312302130-0101133332323323-1133102012111102-2113211030012220-3300211130112002-3230123003312331"></a>

#### `routes.response_cookies_to_add.overwrite` property

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

- [samesite_lax](data-sources--route--reference--group-002.md#canonical-1201102313111332-0020133130332023-3213221030323222-0211230312232331-2121211031000210-1002203131321103-0200213021322202-3103123122011101): complete subsection reference.

- [samesite_none](data-sources--route--reference--group-002.md#canonical-0333103202002222-3323120112303010-0101013322233220-3300312020020112-2030103021200220-2132300320210120-0231033210001212-3003001132120121): complete subsection reference.

- [samesite_strict](data-sources--route--reference--group-002.md#canonical-0003213132003311-1021100110210010-0103303303031320-0312022121102112-0011221132333111-3233032331112032-3001001322203222-1011300000231100): complete subsection reference.

- [secret_value](data-sources--route--reference--group-002.md#canonical-0332003212021212-3110133212331132-2321222001321322-3111222103011113-1101031321312032-2021110110231013-3121220001233302-0020013302332222): complete subsection reference.

<a id="canonical-2302233101300301-0011022330231212-0101010113211333-2112211032210201-3102130313203131-3103100120101003-1131013202200120-2022223220302211"></a>

<a id="canonical-2223333132323131-3333300110212101-0223112233203310-1303212103312203-1211132133331301-0032200203301310-2022131122111132-2131102320001330"></a>

#### `routes.response_cookies_to_add.value` property

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

<a id="canonical-1202211010312303-2310011020212210-1121312302220221-0233303200310102-0103231322322102-1030032301312300-3323011213010123-1111310203312203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.add_httponly

<a id="canonical-1303132013102032-3023023130201123-0030012332200113-3131212310111303-2103121300122101-3202231021220000-3030103021330100-3020131103111020"></a>

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

<a id="canonical-3121111210011022-0323101203033301-2313313320030121-0312231112012203-1200021000210222-2232000320211032-1303022212013221-0320210111111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.add_partitioned

<a id="canonical-1112333111223210-3003113110300011-3210332220201322-2111112213310310-1022221313002333-1201220023222323-2103122033201112-1320220033212120"></a>

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

<a id="canonical-2002130122230223-1111011302313321-0031223100321013-0202210322332230-0110103312323113-0120233231133221-2330201321220300-2013033023003311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.add_secure

<a id="canonical-0101013013130113-2031003112130123-1112121203323322-0332020223001301-2333103131300321-3000311320331030-3220000121212202-1230211220003333"></a>

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

<a id="canonical-3102121031321222-0031131130021121-0220220103100321-3020331110020031-2220032130312103-0223112133213303-0310110301120102-3031132001212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_domain

<a id="canonical-3031021330200310-1012200132012200-3313320213033223-3221222033233322-0100123202310121-1302132233221310-3011333230022321-0301012121100321"></a>

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

<a id="canonical-0220313101112103-2211213002100113-2130031103211303-1301111003312123-1002101322300332-3113320230220223-1330212013331022-1131311133131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_expiry

<a id="canonical-0033313321123231-3233032313223320-2203010120112122-2003132131212001-3002202102210220-0133130013012000-0121212233132233-1333332023110202"></a>

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

<a id="canonical-0013022112323203-0331000012320222-1230131310223302-1201121001001320-1202012012123203-3123220310001200-1011210313111203-1133001113223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_httponly

<a id="canonical-3003111333130300-1311201021212100-2101020002003302-0333210212032001-1132003013311111-2112322330323113-0210313330032302-0033032321203332"></a>

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

<a id="canonical-3203321220330020-3321221220310032-2112023333110120-1101313311220321-2003021111332311-3022201023110223-3111122330310223-1003131111021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_max_age

<a id="canonical-3120223322221221-0113223033222003-1021222201111020-2301102331110320-2130101313030300-0331210212003030-1012022022031320-1032300123130222"></a>

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

<a id="canonical-2232021210122201-3001033322020112-0012331323102112-0320100003303321-3120321011132331-2300023111320100-3303132020123023-0123132031312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_partitioned

<a id="canonical-2300132020121310-3312110102021333-0000321013333322-0332021131201303-2313133113032310-2210023103321131-0103321111011123-2313033032122113"></a>

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

<a id="canonical-3030211112231030-1032002301101013-1303313212112233-2121233022022030-0301230032313231-3321201202120032-1211013133013121-1020000021320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_path

<a id="canonical-0002212130310110-2121323102232332-1100220310201323-2121220122231210-0331003211113302-0230031010311103-1233210322100332-1033121221203111"></a>

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

<a id="canonical-1222213112100231-2323103213120232-3023321021200220-3321120212331022-0323111231110312-3300130122321012-3312302232010013-2132133200002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_samesite

<a id="canonical-2031011013010021-2030221023333133-1001012031030023-0213320222223310-3303320011233332-0231222112122322-2001030311021133-3130111303323023"></a>

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

<a id="canonical-1302031231221032-1123331002113303-1332232132323311-0212132211211133-2020111320200020-1220121320013321-2020133210111320-2223121003022201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_secure

<a id="canonical-3232003310300232-3013031310031211-2321101201200032-1232133332300230-3133332110103003-3221001233302331-2320231323002120-1010312322013103"></a>

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

<a id="canonical-3333332022113210-2132312032332302-0233013330212002-0312323211020302-0200300330010301-3121031233332210-2300333112313120-3013202132032220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.ignore_value

<a id="canonical-1212131111303233-1100331201120033-0123332320112020-0322011131220313-2001120011113010-0220012312232011-0332013223232100-2312332021011223"></a>

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

<a id="canonical-1201102313111332-0020133130332023-3213221030323222-0211230312232331-2121211031000210-1002203131321103-0200213021322202-3103123122011101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.samesite_lax

<a id="canonical-3302021321231001-2322002000002223-0333113312101322-2030233101232012-1320030300223312-0120000033232303-2232032332220010-2011212011303222"></a>

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

<a id="canonical-0333103202002222-3323120112303010-0101013322233220-3300312020020112-2030103021200220-2132300320210120-0231033210001212-3003001132120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.samesite_none

<a id="canonical-2130231100213011-1111031200220311-2020210122112130-2211302003310301-2330230023313102-2020001203203212-3321312211131313-1233310333130232"></a>

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

<a id="canonical-0003213132003311-1021100110210010-0103303303031320-0312022121102112-0011221132333111-3233032331112032-3001001322203222-1011300000231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.samesite_strict

<a id="canonical-1131213011031031-0220033311302000-1011301203100032-2013320001033031-0121030011301331-2210213200311212-3331130112232232-1311131113333323"></a>

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

<a id="canonical-0332003212021212-3110133212331132-2321222001321322-3111222103011113-1101031321312032-2021110110231013-3121220001233302-0020013302332222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- routes.response_cookies_to_add.secret_value

<a id="canonical-0013232213233331-2300123023120132-1303001112122331-2120313301321033-2312202313003230-0111233302000113-2120222213301221-0200002130233302"></a>

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

<a id="canonical-0020022130212002-2101110203003101-2300133232213201-1032030022020303-2010132020113211-2332213233012302-2220100333020132-2303322332111210"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-0012213111332130-3110303133200103-2011000133333223-3321111032213110-0330032233023313-1301013011203013-3331200211322022-2020122013103232): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-002.md#canonical-0202212013331133-3031103330131320-0110023102313333-1030001003001313-3120102001300002-0223132000321002-1211323311003120-3332310132203312): complete subsection reference.

<a id="canonical-0012213111332130-3110303133200103-2011000133333223-3321111032213110-0330032233023313-1301013011203013-3331200211322022-2020122013103232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- [routes.response_cookies_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-0332003212021212-3110133212331132-2321222001321322-3111222103011113-1101031321312032-2021110110231013-3121220001233302-0020013302332222)
- routes.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3331330211103113-0001221300202302-1123130230131013-3211030121123111-3331220220331313-0022130322100001-2002211022312310-2312132320323131"></a>

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

<a id="canonical-0103112201030212-1012302333312320-1001000311003320-1330021231232130-1000320130011110-0313031211122332-1010101021013022-3202130102110212"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3113320100320202-2311133301331303-2323330031222011-1300110121321001-1330320031211301-0032203231322112-3103331120212001-2030030332121300"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2031112221010300-0130330131122123-1230030100301333-2120301111020010-3020200202033231-0320002020230230-1033122311133130-1130030103202333"></a>

<a id="canonical-2312120302320000-0031331213313110-3011231121311202-0200012200201220-1310210103200030-0001000323132031-0002202111213030-2002200021122000"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-1102300031110132-1331122223233331-1222333100320003-2022212032311130-0031103122321102-3013122231303030-3122333311232033-2100032212232020"></a>

<a id="canonical-2121221120011201-2221321312122331-1233330303110123-2202012220001000-0100322120031003-3213303212231300-3202211122132200-0012123200100012"></a>

#### `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0202212013331133-3031103330131320-0110023102313333-1030001003001313-3120102001300002-0223132000321002-1211323311003120-3332310132203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_cookies_to_add](data-sources--route--reference--group-002.md#canonical-3200032101012132-0202320312021201-2130132030201331-2222113332221302-3301232301021101-3303122120113233-1331330031110220-2203203221211031)
- [routes.response_cookies_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-0332003212021212-3110133212331132-2321222001321322-3111222103011113-1101031321312032-2021110110231013-3121220001233302-0020013302332222)
- routes.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3213102213132123-1201133011001322-3023130222130000-0330020113333021-3022112322210322-3032320000211231-2112302202212202-2033221201133332"></a>

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

<a id="canonical-1312022322131322-3221133222020323-3031101320201110-1003232321333311-1332022313223013-2222233301121330-2312223311313020-0312012301230200"></a>

### Direct properties for `routes.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-2023320231323210-1132133032201102-2011320132212223-2021100321132013-3111222011320111-1000211312222210-3111322330232032-2320030022332020"></a>

#### `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3021333022202231-2102111001221113-2120103330100000-3300321010333002-1303122322120110-3303301313233301-1303032000122320-0001130100133313"></a>

<a id="canonical-1031313231101332-0111100101113102-2001013312233230-1322022000131002-3321312203110020-1332132030230312-2232013103203201-0021231220221001"></a>

#### `routes.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-2031132012221313-2231011210002013-3003320011222122-2113303301333231-3130312023212100-1221120320210322-1210200320102232-3230120313103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.response_headers_to_add

<a id="canonical-0203001121301223-3212032101312230-1230230200012310-3233330320211312-0120323333312032-1111130030201210-3323122202212311-0002200220031312"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied before headers from the enclosing VirtualHost object level.

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

<a id="canonical-3113001202122200-0120031200202111-0101222123200013-2131021123311101-2221232202032333-0033332011223320-2213110002200103-3330003011212001"></a>

### Direct properties for `routes.response_headers_to_add`

<a id="canonical-0100303221213101-1021103202122123-3020321120121213-1021130302231322-0230130323321131-3203220220120121-0320302130220222-0222001131130210"></a>

#### `routes.response_headers_to_add.append` property

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

<a id="canonical-0020010132023210-0132320303312221-2322100130300100-2000120230023121-2223300012211020-0220020030311320-2231021010133231-2010330300330110"></a>

<a id="canonical-3013121133320203-1220021122320110-0001211120120002-3031223212132310-3123220302321221-0211301201330003-1113330203313101-0333030100320031"></a>

#### `routes.response_headers_to_add.name` property

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

- [secret_value](data-sources--route--reference--group-002.md#canonical-1320012322111033-0033000321210120-2210011202301233-2221013023232001-0000100313113020-0200222322232033-0203213003003031-0000021000001212): complete subsection reference.

<a id="canonical-1101223032332231-3200231002331032-2130120310233232-1222010322113232-3001332303300011-2103302032200211-1012003120323202-3200301110300131"></a>

<a id="canonical-2220010222333203-2321011000001301-3233203003110221-0032211120031110-1120320321202023-3130321001111030-3113120320221201-3100200202002321"></a>

#### `routes.response_headers_to_add.value` property

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

<a id="canonical-1320012322111033-0033000321210120-2210011202301233-2221013023232001-0000100313113020-0200222322232033-0203213003003031-0000021000001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-2031132012221313-2231011210002013-3003320011222122-2113303301333231-3130312023212100-1221120320210322-1210200320102232-3230120313103301)
- routes.response_headers_to_add.secret_value

<a id="canonical-2111300222103120-3330313130331002-0011033102011312-1102321002021202-2002101212303231-2221121012222121-1232211223103010-0020233300021021"></a>

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

<a id="canonical-2112223330021011-3131300001133202-2021003303103220-3101102121021122-1303011203020220-1233021300111113-2221322002222331-3202012220001211"></a>

### Direct properties for `routes.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--route--reference--group-002.md#canonical-2102121132131333-3203301320001120-2032223231130211-1312230300023122-3322013112201320-0231331013301100-3011021331333221-0123321232010301): complete subsection reference.

- [clear_secret_info](data-sources--route--reference--group-002.md#canonical-3002302100123031-0333211013120001-3320001321213032-2120030111132102-0210222020323003-2301110211131232-0212001131122133-0330110011201221): complete subsection reference.

<a id="canonical-2102121132131333-3203301320001120-2032223231130211-1312230300023122-3322013112201320-0231331013301100-3011021331333221-0123321232010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-2031132012221313-2231011210002013-3003320011222122-2113303301333231-3130312023212100-1221120320210322-1210200320102232-3230120313103301)
- [routes.response_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-1320012322111033-0033000321210120-2210011202301233-2221013023232001-0000100313113020-0200222322232033-0203213003003031-0000021000001212)
- routes.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3303001330031120-0023013233031032-0011302310320131-0013122121223331-2333321310230013-2202110020220303-2130021203013303-2233103221221223"></a>

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

<a id="canonical-0201012001311031-1020222112330102-2133012001213210-3102101002133122-3110212330131300-3010300220222330-0032202121303031-3112023323310032"></a>

### Direct properties for `routes.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3331203103223013-1303100323220122-2101333312211330-2001303033200202-1321313233101111-1332021310121330-1210010313001011-0130032332200223"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3113320010231333-1233131031100202-0213220030113112-1100022321022113-1220011331002231-2223332123323032-1223201323123302-3133102312003010"></a>

<a id="canonical-2321001131020011-0031120110232220-2223303021321202-0022303210200001-1003110222213212-2003000230331132-0030110303223213-2132101223033120"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-2201222211030213-0121012221212013-1101210333002122-0222023002031300-0012000233332021-1233011201201112-3120020222210001-2011333123131313"></a>

<a id="canonical-1232020313301202-3101211313303211-3301322122303312-1301003133120213-3031323303003230-3112310003112302-1322012123322312-0312231032220110"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3002302100123031-0333211013120001-3320001321213032-2120030111132102-0210222020323003-2301110211131232-0212001131122133-0330110011201221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.response_headers_to_add](data-sources--route--reference--group-002.md#canonical-2031132012221313-2231011210002013-3003320011222122-2113303301333231-3130312023212100-1221120320210322-1210200320102232-3230120313103301)
- [routes.response_headers_to_add.secret_value](data-sources--route--reference--group-002.md#canonical-1320012322111033-0033000321210120-2210011202301233-2221013023232001-0000100313113020-0200222322232033-0203213003003031-0000021000001212)
- routes.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3223231310032112-2113032003332221-0203233330221030-3101311120301312-0311121331120023-2211223132003122-1321233130111131-3302232001101310"></a>

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

<a id="canonical-2331133023332122-0302301123201213-0020002102201203-2230200113122131-0301013130112013-2032233031100123-2221131233113000-1333220222030003"></a>

### Direct properties for `routes.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3233102233322112-3210021220103123-2103003031003200-0223113001021233-0002231012100100-2320333023211332-1032212103310021-2103020330122121"></a>

#### `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0030231031000102-1233131012031030-3003312013320022-1320323212012113-1022121132133022-0332010130000013-2013320223210322-1101231203311011"></a>

<a id="canonical-3203032210231033-0231333211212131-2021020303223321-3213331010110211-1221123102003212-0013101021133202-0330230123120013-0132123212102001"></a>

#### `routes.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1202012112011321-0121001021100300-0222122031112103-2301031033232202-2111211202231213-2113103132202333-3320120123000013-2322213321313230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination` properties

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.route_destination

<a id="canonical-2012333103302102-1200112011010103-2321032030203303-0230113222313323-1003211133101332-3200122302022000-3111311320333223-1221111012203323"></a>

Type: `"single"`. Computed.

List of destination to choose if the route is match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"host_rewrite\"]",
  "x-ves-oneof-field-route_destination_rewrite": "[\"prefix_rewrite\",\"regex_rewrite\"]"
}
```
