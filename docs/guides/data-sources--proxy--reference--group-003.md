---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-1032010333203031-0311210232233323-3113033322010220-0323133011033301-3222010211301030-1003031331031202-3121111112313123-2331130313011021"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3232222122232311-1303120010233130-2230232120103321-2023120003321320-1200211100322123-1133121230202310-3211203200201212-2010131310302021"></a>

<a id="canonical-1302030020123110-1332130233101023-3120032112331213-1203330011220303-3123331122230302-3130311323012302-1332022113020333-1020021210331023"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0111102122200200-0231211332230122-0002101102313301-2021122122203210-0203231033302312-3222322203312130-2020310123200302-2303323322131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
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

<a id="canonical-3313212111230123-0012020021112203-3330032121110313-0202030022220110-3113302320232002-1110213333020311-0130110333111211-2030303113301112"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0033313300123220-0303310302311121-3233300133011331-2303023003323011-2313313003332201-1000333331122312-1303033221011200-0321333230020320"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2021332301232030-3332130002023033-0213332130212210-3323202321013330-1003020310333300-0203131021231232-1330130101230111-2132011002112003"></a>

<a id="canonical-0202112332131303-2011312302013211-1201103301222330-0122003333030010-3300110233332000-2222210003323123-3332321011113130-0101132112112220"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-2032131130021113-3112221331102131-2330120323033003-1211212001130203-2230120310022031-0233122110233300-1322203220120310-3211333211112032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_headers_to_add` properties

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

<a id="canonical-1320002031233322-3212323023103123-1201220013122320-3131322103003331-2313233133301133-3021000213131132-0110000321113331-3221320213222103"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_headers_to_add`

<a id="canonical-1203032111112212-1020332110123231-3130231133112021-0131313031133202-1321301012112002-3212120301020021-3231122210222203-1212212233102003"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` property

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

<a id="canonical-3011221200301132-0212202310111112-0303101213203330-2211100330222323-1201320103233113-0310121302112332-1220031203302303-2032112113121230"></a>

<a id="canonical-0023001121011230-1311320110201330-1223133010020220-2133331221002233-3202322132000033-1223221310122030-0011123230123231-2312120222312023"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` property

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

- [secret_value](data-sources--proxy--reference--group-003.md#canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331): complete subsection reference.

<a id="canonical-1012013020301023-2200130313023231-0212323002033231-3023001233120132-1120300321332100-3011223232231003-3001301013222003-1022111320101103"></a>

<a id="canonical-1313200022301201-0323002103303123-3103133013202020-2110202003312022-3033213033333002-2230113230133323-3220120002332103-2022213221112201"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` property

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

<a id="canonical-3232003231221231-3113233001000203-1221301113222132-1300331001112111-0222310013110303-2003231120132333-0332033122322132-0322123101112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` properties

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

<a id="canonical-2330312122121231-1122020131121123-0322120031132222-1322101221033213-2101032203330111-1033101203211330-3322032110313121-3003031011203333"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-2003012131112120-3303100100311231-1030111010120121-1300333220012030-0023121100122323-3130010332130313-2303100311223230-0003232123222222): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-1102122120002100-1201303223210221-0233320231232230-1321212120103212-2130002101011111-0010031200301122-2301221310332231-3000312110121100): complete subsection reference.

<a id="canonical-2003012131112120-3303100100311231-1030111010120121-1300333220012030-0023121100122323-3130010332130313-2303100311223230-0003232123222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

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

<a id="canonical-3333111233310223-0201123230003121-3203203211230003-1202231121001103-0201302310212022-0333222100311213-1321123022323323-3101012220111132"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1333311221313012-0030120122211311-1133310333000323-1330123010022211-1113302230030213-1130310111220103-3212022001202221-0002211203133121"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3310030111011022-3321011130112133-3113132022130011-3233200231011133-1220312210201110-0032212231001233-2112111300221030-3110021021200320"></a>

<a id="canonical-0322320301130013-1313232223332132-1200201111311111-1310200000131100-2321312301022333-2123011001211132-3331333221010301-2313203321010311"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-0222300231300331-0132020303031330-3131022120033021-2013110032131232-1332220232013012-1211113200120330-0010022120030132-0001313023101200"></a>

<a id="canonical-1330100201231311-2131110000323221-2030303233331022-0010233203123230-1201222211231023-3122101333232033-0211000010133033-1213032130112020"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-1102122120002100-1201303223210221-0233320231232230-1321212120103212-2130002101011111-0010031200301122-2301221310332231-3000312110121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` properties

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

<a id="canonical-3013233103103001-1302231202320212-0222131000033210-3002333202133130-0020021311001021-0203131010010131-3231013301020120-2330322000123022"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1333200201101002-1033310131011013-0321333022133222-0010322133221113-1202013000131213-1333203313123211-0333322113023222-1221302001133332"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2032211333012200-0220231303113202-2311000022011121-1102123302100200-1000332131033000-1231001213202310-0332003312123230-2013230032302311"></a>

<a id="canonical-2021222031133103-1110322300202330-3303223013130130-0021233021013320-1301023031201310-2212220031010323-0131232211132233-2103010020320033"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- dynamic_proxy.https_proxy.tls_params

<a id="canonical-2300121300020131-3132302333231323-3132222033011130-0022032313222002-2010322122031022-0022320002210033-3111311330003131-1101302211212013"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

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

<a id="canonical-2022303120223303-0122313013020221-2213103331111301-0321130331200111-2030100012213331-0013110312002332-0312010230201312-0122130230023002"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params`

- [no_mtls](data-sources--proxy--reference--group-003.md#canonical-1023112302312003-1312011202020002-0322211203230223-3231321001022232-2221312333131312-3220322311022112-2120203313101113-0311322212001220): complete subsection reference.

- [tls_certificates](data-sources--proxy--reference--group-003.md#canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001): complete subsection reference.

- [tls_config](data-sources--proxy--reference--group-003.md#canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001): complete subsection reference.

- [use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131): complete subsection reference.

<a id="canonical-1023112302312003-1312011202020002-0322211203230223-3231321001022232-2221312333131312-3220322311022112-2120203313101113-0311322212001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.no_mtls` properties

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

<a id="canonical-3003233102322130-3313112202233223-3232201232213203-2102301120330200-2101210331212121-3011232211322003-1320310003231203-0131133003012001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates` properties

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2103301230022313-1012103302203333-2001200110030302-0021003033010223-3322013133012311-0021113010203321-2021323231002001-0200233013221131"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_certificates`

<a id="canonical-1300132222232231-1223212123003021-0111212022130003-3323320311212010-3300001320121301-3100232323012031-1010023213011301-1113232112020030"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101202231331212-0032003122132333-1312311330321210-2223323212310303-1133202221303103-3023331013010120-3221302122100121-0300103323302330"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--reference--group-003.md#canonical-2101033220232001-1100001001011103-2233202122213332-0232301311211332-3123231020112122-1133123101332221-0102211102111112-0223212113212131): complete subsection reference.

- [private_key](data-sources--proxy--reference--group-003.md#canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100): complete subsection reference.

- [use_system_defaults](data-sources--proxy--reference--group-003.md#canonical-0210120212131231-3021303200313100-3212202103233020-1011221011112310-2110003203323012-1310103203022122-2230011320133221-2011331231201332): complete subsection reference.

<a id="canonical-0231003113201323-1020120131300322-3230321110203300-1102022300311322-0110022231221221-0131330230232000-3023032332311103-3032202212033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` properties

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

<a id="canonical-0303220230100201-3312203323213120-2203121330302331-2212231220010132-2313130201221232-1230302213003001-1330001200132110-1300001033330111"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-2023110202032210-1030213310201112-3122032123133210-2011232302131231-1023103201122233-1311100200330100-3131203213111303-2310332331120323"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101033220232001-1100001001011103-2233202122213332-0232301311211332-3123231020112122-1133123101332221-0102211102111112-0223212113212131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` properties

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

<a id="canonical-0013323221130130-1121330221102010-3332030201231000-1103110310003210-0231322201212100-0321212312231201-2110203012033033-2122313013313100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` properties

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

<a id="canonical-3222210212201000-3303123022131020-1233312310132322-0202123311312210-2020022330202233-3233002331302312-2031002212001203-3313110301003110"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--proxy--reference--group-003.md#canonical-1302122310302002-0120221303120332-3233300233112211-1113333001111131-0232313021110311-0020131020103302-2030301211033310-2122113320302021): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-3303200310023122-1032010022210020-3021121023320000-0112130121302001-3221302313312310-3322301321312233-2031200011331012-1200210233221132): complete subsection reference.

<a id="canonical-1302122310302002-0120221303120332-3233300233112211-1113333001111131-0232313021110311-0020131020103302-2030301211033310-2122113320302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` properties

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

<a id="canonical-0123202101000031-0010000101323130-1133131211000332-2232331232201322-1030103221211000-2133032220031122-2301210233300132-2030301233122101"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0113332310103101-2201013210101201-0322100201210010-0130031001220132-1101031000312111-0101223011302321-0100312211232230-1120120103220211"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3311220132230002-2320322001203333-1233002013210001-2132331211212302-2000010110112303-0021123101023301-1103200010131302-1112001132333203"></a>

<a id="canonical-0211301212003110-0301023020200310-2221133122213223-1023320302313120-2200110332333110-0321001330002001-1132203223300032-3233101033312001"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-1211132311312320-3000011013212332-2031301121321132-0331222330001310-1321220321232222-0333220001221023-1301002123310200-2311231002100201"></a>

<a id="canonical-2130303302331103-3310330131223301-0221033320232133-3310222301201000-2322331120221200-3333330322122001-1220122032333232-2031203200200202"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3303200310023122-1032010022210020-3021121023320000-0112130121302001-3221302313312310-3322301321312233-2031200011331012-1200210233221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` properties

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

<a id="canonical-0320313333313031-2231331023223112-3302100202322013-1123202130001330-3323122323211113-2113331013310103-2321001022322022-3001002220032230"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3320300320001002-3130031302021133-1313031321123203-2322022303002203-2332123220023202-0222101023033101-2213230201032202-1112112133200033"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1020100020033023-1120223332312312-3100031003333000-1023013321312232-1023031311213011-2030022002003200-3132232223321303-0311003013111311"></a>

<a id="canonical-2213210001122112-1011122312301220-3021331220212122-1211232001003100-3122002121333000-1010203300033123-2210122101221322-2013110223002012"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-0210120212131231-3021303200313100-3212202103233020-1011221011112310-2110003203323012-1310103203022122-2230011320133221-2011331231201332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` properties

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

<a id="canonical-2000031332101132-2110322102330011-2031301032313101-3001223022231332-3213012100013112-1002101132320102-0311231303202231-3202020123113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_config` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- dynamic_proxy.https_proxy.tls_params.tls_config

<a id="canonical-2233112202122010-3010112132002021-3210133102201233-0102331302303331-1203210212030220-2213011200032022-2112200013302032-2331033311323313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3011202122021232-0220302232120030-0322131001120301-1313300221111130-2220300300332310-1113030010222032-2323120033002030-1213101212000132"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_config`

- [custom_security](data-sources--proxy--reference--group-003.md#canonical-2211120102100322-3102312113211200-2310020331210123-2110111320121202-0002013212020332-3223301032211222-3023203212010123-1023312211211213): complete subsection reference.

- [default_security](data-sources--proxy--reference--group-003.md#canonical-2003120012313300-2200013003312231-0231010021103102-0222300032223222-1320123321121121-3202002011130101-2011312223313133-3132102111003023): complete subsection reference.

- [low_security](data-sources--proxy--reference--group-003.md#canonical-1322232001111230-3033322212200012-0223101130132111-3222202020301011-0031003132111201-3030013311200103-2131102320001111-3330323312313210): complete subsection reference.

- [medium_security](data-sources--proxy--reference--group-003.md#canonical-2210312203012321-3301212233210203-1331333021200202-1033303202302320-1121210033321312-0103311220213133-1122313020201213-2211110021222103): complete subsection reference.

<a id="canonical-2211120102100322-3102312113211200-2310020331210123-2110111320121202-0002013212020332-3223301032211222-3023203212010123-1023312211211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` properties

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

<a id="canonical-2312223310123333-3132232221302133-1330033022231223-3012113110323220-1320133021320201-0111133201221310-3131130313301033-3212320022123220"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security`

<a id="canonical-2003332020112002-3210312220301022-3001002301303232-0131013312213112-1302033021123101-1000322131321013-1123301100302321-0113312212001110"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0021203212322311-1203001102222310-1102003033111303-0111301323312130-0222310103103120-1233001323120032-2120110022033223-2302110102213011"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-1032133211212032-3223103011123201-0021311122130310-1223201131101220-1201320001320102-2111212201011331-3310323022102332-1010312230200200"></a>

#### `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2003120012313300-2200013003312231-0231010021103102-0222300032223222-1320123321121121-3202002011130101-2011312223313133-3132102111003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` properties

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

<a id="canonical-1322232001111230-3033322212200012-0223101130132111-3222202020301011-0031003132111201-3030013311200103-2131102320001111-3330323312313210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` properties

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

<a id="canonical-2210312203012321-3301212233210203-1331333021200202-1033303202302320-1121210033321312-0103311220213133-1122313020201213-2211110021222103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` properties

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

<a id="canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls` properties

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

<a id="canonical-3132011332221321-1332333303223013-2320030300331101-1321130300201330-1302323113322300-0123023033113203-0113000131222220-3022212310331230"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.use_mtls`

<a id="canonical-1103113010020312-3300002120333031-2211001101000133-0312213313003010-2011223103131322-0111321133213011-3021310013032100-0321210202313300"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

<a id="canonical-3223210031202121-0221331202322120-2203321031002231-0031221123033333-2302300212322320-0213020102012013-0103111220132333-1222130211000031"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

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

- [xfcc_disabled](data-sources--proxy--reference--group-003.md#canonical-2010122003302033-1301211033330203-3022022121101331-3100102323332312-3002231211130301-2332311313323130-1232011003111010-0213113320211312): complete subsection reference.

- [xfcc_options](data-sources--proxy--reference--group-003.md#canonical-3121122030011130-1120332230321011-1210123321213212-1333110133301201-0312133122322311-0000300211111022-1133132110103131-1011321303323000): complete subsection reference.

<a id="canonical-3232321001130202-1301031200001033-3220030003032210-0313323223012122-3330110213230122-1201302333323310-3001302213000023-1021112101033001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` properties

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

<a id="canonical-1203322011110232-2033312110102311-2302033122033312-1210212001313113-0023201123033023-2000211012322112-1021210213100320-1133323303201113"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.use_mtls.crl`

<a id="canonical-0133030130112123-1212102000133321-0311213121323301-2323033000122222-3211023210313102-0203200033231223-1231313331030133-0210213233223112"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2033231112112220-3130211130122022-2110130123002013-3202120020330330-3010031123132311-1130101220112211-2111033321310112-3130332101300231"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2200012223100311-2302120320122121-2020013023112332-0033002113100120-2302331232132002-0123121122201102-3332020021301112-0022320230311323"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2302132330013031-1101202110332220-2320320230323133-3210130122131121-3303232022221020-2223302103232311-2233222021332233-1212230303313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` properties

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

<a id="canonical-0113322102331021-0102312331132203-1131012310302113-3003012200201322-0012303213001132-2323010231111133-2130300201322033-1031010133201001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` properties

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

<a id="canonical-3332030232013021-0333113120132130-3232300332203002-3212300101211333-1301123010133333-3331302200233222-3320201213033310-0110113322131032"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca`

<a id="canonical-1103202323220033-2101211321111112-2121113333300023-2110222103032012-3203100213120033-3011113322000003-3020302310021231-0021210332111221"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230231023320011-1020032202212203-2000120110103022-1102022023013220-0331010221211103-2021113101032111-3201020213322121-1330213233013131"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Computed.

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

<a id="canonical-2133212200030023-2113233102210120-1303030231130132-1013320230222312-2102332221220020-0312220230101131-1020002122113102-3030220232333013"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2010122003302033-1301211033330203-3022022121101331-3100102323332312-3002231211130301-2332311313323130-1232011003111010-0213113320211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled

<a id="canonical-1202022302133120-0300000122233013-3220023003210312-1111200222033232-3121000202130220-0122123310002213-1001121220103201-2022203130321011"></a>

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

<a id="canonical-3121122030011130-1120332230321011-1210123321213212-1333110133301201-0312133122322311-0000300211111022-1133132110103131-1011321303323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--reference--group-003.md#canonical-0331322111111130-0031322311000132-0321311001001330-2022120120201310-0333202311302212-2330213112101021-1332110320230030-1210313210233131)
- dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options

<a id="canonical-2310130021133233-1332332030232213-1300031130301012-1331213332102230-1033302320022221-0031213302121323-1321231031322020-1112233103330111"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1310022123100311-1202221100001232-0022230020312302-3320332312312310-0221203023112310-1310310111330232-2023300113220300-2201212120101232"></a>

### Direct properties for `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options`

<a id="canonical-2203001132031022-2000203311031220-1322123101332333-3313300201210223-3133303120300111-2022100200122011-3330333231123333-3202301102212033"></a>

#### `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3201203013323203-2030201011013320-2110012213112021-0333333012303323-1113131112130002-1333330220013232-0212100100120203-2000200313330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.sni_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.sni_proxy

<a id="canonical-3032311002003003-0031320000113303-0231310022213220-1000120202113332-0322013233223033-3222133013023011-3303303230201203-2301013113121233"></a>

Type: `"single"`. Computed.

Dynamic SNI Proxy Type. Parameters for dynamic SNI proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3120112311131303-2001231201321233-3021002233220312-0012100223301131-0100103333102321-2333312311100302-0010012001022231-1213100131011310"></a>

### Direct properties for `dynamic_proxy.sni_proxy`

<a id="canonical-1011010201103311-1212012221033122-0130321032331332-3001112123001233-3300202030322131-1303013320000301-2312033331002021-3223312233001103"></a>

#### `dynamic_proxy.sni_proxy.idle_timeout` property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400000,
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
    "ves.io.schema.rules.uint32.lte": "86400000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "86400000"
  }
}
```

<a id="canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- http_proxy

<a id="canonical-2121010022301312-1300233121020222-1323121211031111-0131133120001130-0000032031101232-1000113012222203-2213302130100231-2210233233221031"></a>

Type: `"single"`. Computed.

HTTP Connect Proxy. Parameters for HTTP Connect Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_https_choice": "[\"enable_http\"]"
}
```

<a id="canonical-3132211202231013-3020012132131212-0103331221133230-3030200233101331-1213223133012331-3220112013103130-3001013311331003-3221001122203322"></a>

### Direct properties for `http_proxy`

- [enable_http](data-sources--proxy--reference--group-003.md#canonical-1100110320112012-2233332002233000-0003202002200303-2002322321311233-3320033322101132-1122110111123021-0000013321103311-2122032320223100): complete subsection reference.

- [more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233): complete subsection reference.

<a id="canonical-1100110320112012-2233332002233000-0003202002200303-2002322321311233-3320033322101132-1122110111123021-0000013321103311-2122032320223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.enable_http` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- http_proxy.enable_http

<a id="canonical-0230111231103311-1312220320210302-2330212201112022-3310011330222110-3132131100012230-2000310112111311-0030002102130202-2303211213010213"></a>

Type: `"single"`. Computed.

Configuration parameter for enable http.

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

<a id="canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- http_proxy.more_option

<a id="canonical-1201330330310212-2100230231110200-3132112310002331-0010303310030303-1330321022302012-2110220102020001-0331300200111120-0131133122200222"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0133003301013001-1220212321303012-3131222131030202-0221021033120313-2030130020302131-1123310100230002-1110113212310003-0001102303322222"></a>

### Direct properties for `http_proxy.more_option`

- [buffer_policy](data-sources--proxy--reference--group-003.md#canonical-0001222302033103-3032012113233132-1010312221011333-1320323333301003-3222110010100100-1033330233020033-0100200111320301-2121313011312222): complete subsection reference.

- [compression_params](data-sources--proxy--reference--group-004.md#canonical-2212011320322023-1101022131130001-0130300302013222-0330232202110332-2200130332111302-2202221323012132-2210301212211213-2020022130110130): complete subsection reference.

<a id="canonical-2333002231021331-1223330222233313-0101223222300120-2201313302230100-2033113301212220-0111120302312132-1111211110003010-2001023212110120"></a>

<a id="canonical-1133331001113321-1302110121011231-0333322200222123-2302101302200231-2003222132032220-1022133212002032-1112111103203122-2330113333333131"></a>

#### `http_proxy.more_option.custom_errors` property

Type: `["map", "string"]`. Computed.

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

<a id="canonical-0322331033031011-3133000320003201-3122310323213210-0230212123310233-1022330210020122-1132200321001012-2301221000221230-0222232003013103"></a>

<a id="canonical-0301033302001213-0233221001311333-1232133101003212-0301120112020231-1102322331123013-2213012030103121-0123001013300333-0032322120120021"></a>

#### `http_proxy.more_option.disable_default_error_pages` property

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

- [disable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-1003210011230331-2323213330220012-2132303121131103-2001310000033121-2320121032032302-0201100310313031-0202303321302310-3110013120231120): complete subsection reference.

- [enable_path_normalize](data-sources--proxy--reference--group-004.md#canonical-1222133101111030-2231302230300323-1002110332333200-1212330122133323-2100222233003123-1011133001010013-3121103331320101-0331233320102230): complete subsection reference.

<a id="canonical-3230100301210323-2331103232032032-2312021301210102-2303231322000030-3013231113010120-0333120203020233-0310002123112021-2132021030330003"></a>

<a id="canonical-1100220233011020-1320103230303120-0320333012131232-0131301312030131-0101313210021102-0030320313101221-3032113202100320-2112210202002313"></a>

#### `http_proxy.more_option.idle_timeout` property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

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

<a id="canonical-1212311023133021-0210021031311231-3300310003332302-2210023032233010-0333113122010121-3223223022301331-3113223303333233-2122131203321233"></a>

<a id="canonical-2203230210230330-2212023322010120-1130332231212130-0301133101130321-1131222232002133-0330211320021332-0212233332312133-1003012323303313"></a>

#### `http_proxy.more_option.max_request_header_size` property

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
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

<a id="canonical-0031233121333221-3320313202222003-1232130012101101-2210201002133311-1303010102031032-3201233201123011-3010003302100032-1231030102331231"></a>

<a id="canonical-3003301031012020-1132121332230012-3020201132112223-2212013003300120-3330131021303210-2200031021302223-2102322002122323-0013300332123313"></a>

#### `http_proxy.more_option.max_requests_per_connection` property

Type: `"number"`. Computed.

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

- [no_request_limit_per_connection](data-sources--proxy--reference--group-004.md#canonical-1202323110231213-1113003323031203-2323321001333030-3110301331112323-0100221331123110-1311231300211313-0300021121133003-2122123200122000): complete subsection reference.

- [request_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-2333120002103000-0021233110122311-0220031112322320-0123003233011103-0001231311223111-0232001220300220-1213321001012222-2003330220332213): complete subsection reference.

<a id="canonical-1333112201131100-0013301330011210-1210310230113332-1132000232012332-2320103102023022-2133201100202132-1013213203301113-1301200322132320"></a>

<a id="canonical-2200230113101332-0311303121101030-3231331310310312-3120212322020312-0300112111332002-3331020101333230-3133031123211120-1103202123131123"></a>

#### `http_proxy.more_option.request_cookies_to_remove` property

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

- [request_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-2213012201232033-0321330202023321-2330031333303321-2113210102231203-3232031210002002-2201101202033323-0031022310111323-3332302313022022): complete subsection reference.

<a id="canonical-3002330133023103-1221321012010332-2102210101113202-1003001322212101-0303003133302003-1213301311010022-2213031032010312-3133101201020200"></a>

<a id="canonical-3230120311122311-0213323132222012-1012332022101131-2230300122111133-1100102002220110-3300203232110010-1332303012123203-0200030031103030"></a>

#### `http_proxy.more_option.request_headers_to_remove` property

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

- [response_cookies_to_add](data-sources--proxy--reference--group-004.md#canonical-0012320332123001-0332101010113232-3001021111233230-0013220321102002-3000021323121100-3001031330121312-0320133111212300-3100221200222322): complete subsection reference.

<a id="canonical-2113122013300121-1023223313203123-3103111120130023-3313031223211330-2133223003000001-3331301000020003-1332321203031002-2110103013230132"></a>

<a id="canonical-2121000022231221-3303120312001300-0101311111110331-2130232112223122-3012330001122002-0020010100001010-2210303102023031-1301120031100111"></a>

#### `http_proxy.more_option.response_cookies_to_remove` property

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

- [response_headers_to_add](data-sources--proxy--reference--group-004.md#canonical-3220133020330022-1232303102011310-1113033131110013-3100220200031212-3021203311201313-2122021321110320-0110232121302032-2311132322032113): complete subsection reference.

<a id="canonical-0202123213213210-2301333312111103-3013031003323121-3301331221010203-1022120101001100-0003032011000001-1321102223223012-0301231030211322"></a>

<a id="canonical-2033120121200023-0232233301201003-2103031111232202-0310112323232332-1310131323303131-1032010210010001-1202002010321231-2312020232120232"></a>

#### `http_proxy.more_option.response_headers_to_remove` property

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

<a id="canonical-0001222302033103-3032012113233132-1010312221011333-1320323333301003-3222110010100100-1033330233020033-0100200111320301-2121313011312222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_proxy.more_option.buffer_policy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [http_proxy](data-sources--proxy--reference--group-003.md#canonical-1322223011311300-2113232222300030-2322033332222010-0213033232301212-1001222321013113-0033123000331123-3101222202133301-0011020023112302)
- [http_proxy.more_option](data-sources--proxy--reference--group-003.md#canonical-0021103321131333-0122023312212232-0212301101302001-0322110000213333-0001213120113031-1332312213132223-3321030310223032-3321120122023233)
- http_proxy.more_option.buffer_policy

<a id="canonical-1321132332021200-2120310123310102-0103322000100121-3000030121130002-2223320110121200-3112131000011300-2302222120113103-0113223023320013"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0332222103312310-2130333110000101-2221312211013012-3310222230300310-2301003121213300-2321310033230000-0322023221031100-0222030332202322"></a>

### Direct properties for `http_proxy.more_option.buffer_policy`

<a id="canonical-0313221210102222-3210003110231013-3300331310032013-3133022030230113-3130303231202203-0230032231133111-3002212000013123-0212133113301220"></a>

#### `http_proxy.more_option.buffer_policy.disabled` property

Type: `"bool"`. Computed.

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

<a id="canonical-1022132332103130-2100113231102003-2233023212231200-1330201321202132-2301302100131211-0101312020230220-1230012310100323-1313000000313122"></a>

<a id="canonical-0303121310322311-1111212212110010-3312310200122022-2011221321110111-0300331021330113-3300023123133213-3200132333120113-2020122210123210"></a>

#### `http_proxy.more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Computed.

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
