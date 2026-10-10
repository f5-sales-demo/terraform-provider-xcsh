---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-0313013020021030-2002113111010212-2210121310232331-1100011000010003-3223202131130123-0021001311223120-3102322223312111-0100120033123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-001.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-001.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
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

<a id="canonical-2212223032212133-2103033011231100-1010231203113111-0121112033012311-1223200220021330-1131001210331221-0202003312230212-0123312211220330"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3303311232111323-3222113001012120-3032223032120130-3212102210010203-0220203220211303-0131100132233333-0212100021322100-1312302103002131"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1223030023101231-3112022210032323-0102122112113111-0200221231203023-0320230011231121-3100300030120021-2031302313331323-2200130002221321"></a>

<a id="canonical-1222320222302022-2231102323310012-2230212113331201-2333113030020001-2121103323033332-1201232001203103-3333003030000312-1102320333202023"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3210113310033300-2322002110333102-0022212201033331-0102133331203003-1010023310233312-2303210113322233-0131101332023020-2031030211000200"></a>

<a id="canonical-1101113233311032-1232032100303013-3032131320120201-1212022110233210-1313221323131031-1212302231102113-0313020120211312-2310201320332031"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0232103032301203-0201202110301122-2213321000022112-2323020312201033-0023333220302310-3003330233110230-0003031003202111-3032121231102013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.http_proxy](data-sources--proxy--reference--group-001.md#canonical-3210231023333002-2031133312213030-0233021110220001-2013230121333232-3013102330022101-0300001012212232-0101023013121321-0123233212111000)
- [dynamic_proxy.http_proxy.more_option](data-sources--proxy--reference--group-001.md#canonical-0102312311322132-1200123220023330-3100003031330102-1223131331022222-3022232021330103-3012221011303302-0123320201113203-1001321310031201)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--reference--group-001.md#canonical-2303301033010202-3133022000003110-1031331200010303-2230302030022310-0100230303310021-0120221220002211-3121010321000033-1002103011232120)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--reference--group-001.md#canonical-3121223110303320-0020230331202132-1033113120102020-3221031201133132-0230123011323313-2310131201213313-2311321311102301-2223333233002032)
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

<a id="canonical-3020002203000100-3230131211112002-2000330230023300-0130011131212211-2313030301003001-3230330013221213-1232333000213112-3230201333321221"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-0312331223033230-0211003030101323-3131213131002022-2222333230102330-3313231300230201-3103333121201112-1301202210332330-2131103103310111"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2323300032213010-3121211031212221-3120330123120222-1100002232200201-3002331033103313-2013110012232003-1122023222312003-1001002320200312"></a>

<a id="canonical-3001102000003300-1302203323133123-1201110321012011-3033201101023023-0221100223311033-0110323331103222-3322032203231021-0023231031031210"></a>

#### `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1003213031100030-3111200032003222-2231310222011322-3102330002003302-2310213002030322-0102220300201000-1311312001301103-1113302021310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add` properties

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

<a id="canonical-3010203021031300-0000110310313022-0310012202210313-1310112232213122-0323301230221202-3302212320320011-0122011222100133-3002320321333033"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add`

<a id="canonical-2321122302230302-2013133233200101-1133331021120212-2230323313212122-0031213222121201-1231321303101310-3003003030102231-0113021320311011"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` property

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

<a id="canonical-2320132322211121-2102333302210233-3323020111200302-2221023103003132-0223110202030021-0310133113131321-0003022223030302-3333323203303331"></a>

<a id="canonical-0322133100322131-0130221031130100-3023331010331200-2120030023331332-0012323331302223-3300113230332000-1131100100003123-2131303322302002"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` property

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012): complete subsection reference.

<a id="canonical-1223303110101311-2313013223303000-1121030211303130-1203110311322321-3003021220110333-3110301222221223-0131331131332302-0021122331201302"></a>

<a id="canonical-3100323010123323-2320301020220223-2130023020201220-2111132121203033-3232310130310321-1113332021210331-0323310312102312-0130113101331202"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` property

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

<a id="canonical-0300221020113123-3123131313231130-0220210030112022-1033132101232313-0200213331112110-2213113233133222-2220232320301111-2013203310221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` properties

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

<a id="canonical-0013231030322001-2232221302232233-2031011130332311-1313123331310332-3221022010321221-2103121122130323-3012323303200122-2313033021311011"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1030023133302332-3102110121220301-3012111312210130-3203020223200132-2001033331112200-2131221333003032-0032013322131222-2112002033123313): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-2100310300023120-2120103332133310-3132022220322302-1202113210122131-1001330031131210-1121012021303110-1132130300022231-2112202122302212): complete subsection reference.

<a id="canonical-1030023133302332-3102110121220301-3012111312210130-3203020223200132-2001033331112200-2131221333003032-0032013322131222-2112002033123313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

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

<a id="canonical-0222210000210113-2002000122332003-0013213121233220-2201300032110120-3322213321032200-2300320331111233-1221312030032022-0112122013301301"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3001023131031010-2231330200312133-2322101230312331-3011320223022223-0000030033333020-0120001031010132-3320333303121212-0132312230200200"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1311311131030100-1210002313201333-0112333221213020-0210112201211020-3321312032333331-2012130112011321-2033102121330320-1100202011331101"></a>

<a id="canonical-2201031230302121-3120123130003313-0120210322133320-2110000320212130-2033003103013021-1332203120222221-3312310310301320-2111030133123020"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-0202003202012110-0213223010031310-2213300200100313-3310302010322121-1231232130010303-3201110010122131-0002121023213222-0003210021231212"></a>

<a id="canonical-3302013123012200-0113200212322123-0221220310030111-0030303200002132-2011122113223122-3110221033223133-0321121131223302-3102002200302103"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-2100310300023120-2120103332133310-3132022220322302-1202113210122131-1001330031131210-1121012021303110-1132130300022231-2112202122302212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

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

<a id="canonical-1032021120112302-1132211213020220-2221331020023103-3131232113331131-3132330032320320-2321300221303031-3013130000220310-2111333100301330"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-1312023111303332-0313231320121023-0031221020031331-0332001120210012-1312122113220000-0013030221311233-3103203123002232-3112200113321200"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3330131200223123-0333131311021013-0111121223230110-2200201312111333-1113331120203120-2033302312131222-0311321303202233-2033320011212321"></a>

<a id="canonical-2230013111023232-0220033320211331-3233221000201010-3200122010210030-2220211021321112-1230303303021011-3332011132202032-0201112212313311"></a>

#### `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1001121231112113-2321013313123233-3121011120032112-2230103211203211-2220002211200301-0323022031220331-2232320202211030-3233021220031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` properties

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

<a id="canonical-1111132200222230-3120023132201022-3000123213032032-3101230110331330-3010203300133011-1113230233132322-2131211320023333-0101132333000200"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add`

<a id="canonical-3023013221131033-3202201310301300-2332010032102102-1233031200300201-1023132003100011-1100022102322232-2121201331101322-2132221033322223"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` property

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

<a id="canonical-1033211023122220-0131120210001303-2322021230222232-0303111203013313-1231321031332130-3131313023012223-1203312000221322-3100331313110031"></a>

<a id="canonical-1001100321301310-1122101000220111-3013331110030311-2110211013302010-0112333013222120-1123323111032133-0011331103231032-1033230023033222"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` property

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

- [add_httponly](data-sources--proxy--reference--group-002.md#canonical-1330320023303103-2302022130231131-2333232331213210-0002003331120312-0201333020223020-1232133321122300-1122113013020030-1220003331100012): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-002.md#canonical-3001022031232111-0122001320123233-1310332333101030-1320233302212010-3012310113033010-0122210321033202-0312311303112102-1112021130023003): complete subsection reference.

<a id="canonical-2003222133030332-1123332022110010-3112233213021133-0123312131313012-3121212120131112-3002201333103213-2232000001100111-0020003212303110"></a>

<a id="canonical-0301131103110103-3123302230100322-0202330230313201-3203320001332320-2131220111113310-0022313201120212-1203310313001020-3312211033012103"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` property

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

<a id="canonical-0102013311012221-3013312103011123-1000002022102012-1131023211121113-1311222200123123-1121331331102313-3122331231310032-0330032323033312"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` property

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

<a id="canonical-3313201223020301-0203303212322000-1113201022203120-1130012021020013-0123312312002312-0130012321230022-0231121311001222-1100023122333101"></a>

<a id="canonical-0321232330100031-1133121323023203-1132220323020122-0301022032301123-3321010312012320-3232012022101011-3202111103031233-0121302233133030"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` property

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

<a id="canonical-0110222213003303-2020220123133310-2210311110223312-2013032112032001-1002110300330130-3323031320230111-2320202113021320-3300222311203012"></a>

<a id="canonical-0031202020210103-2013323211232101-3333100210020113-0030111000032020-1230220010123300-2100300022233123-2020333131201331-0231130322311002"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` property

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

- [samesite_lax](data-sources--proxy--reference--group-002.md#canonical-1321101233023113-2230303111330010-1322303301031233-3302120010112303-2213220032222020-0212033123332010-2201112211201312-2330013331202030): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-002.md#canonical-3000332003321310-2222200123021310-0022123022030213-2300033332121010-1232020131322111-0023111020201203-1013103313120120-1213331301010331): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-002.md#canonical-2302100211010213-1221110132232310-2212203011100023-1232220102103322-3120330203030100-3002113110022323-2221303300000213-0230130110100330): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203): complete subsection reference.

<a id="canonical-1230232323110101-1102031022210022-1021312321323213-0300001323213303-2112213330203121-2302321323030303-0220131203011102-1121012011203323"></a>

<a id="canonical-0313000000101211-3331211333233132-2210333032320132-0133130201312230-1220030301313331-0222131102013113-1200013230313201-2312122223222122"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` property

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

<a id="canonical-1330320023303103-2302022130231131-2333232331213210-0002003331120312-0201333020223020-1232133321122300-1122113013020030-1220003331100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` properties

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

<a id="canonical-3001022031232111-0122001320123233-1310332333101030-1320233302212010-3012310113033010-0122210321033202-0312311303112102-1112021130023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` properties

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

<a id="canonical-0011102001032030-2331301112031120-1300022220021201-3320121023213032-2302001023033003-1322213002221113-1223023120131120-1012201320030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` properties

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

<a id="canonical-2000320130332032-1233020232300330-2002133203003311-3312200230233212-2122220303123211-3320330322222230-1300010131133120-3210223023123231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` properties

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

<a id="canonical-1210112000100300-0332300323233332-2212132223111012-0310310211302021-0110320131032113-1220113113132301-2310132120112302-3033222131021333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

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

<a id="canonical-3232020031003300-0031310323013133-0022131022233013-1231210333020221-3323231302203112-1133002233213031-2022111003222120-0013203333303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

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

<a id="canonical-3203111331101100-1133310211331120-1103032202331311-1213221210300333-3131001031221321-2200003330023222-0033023003200020-1120202002030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

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

<a id="canonical-0233030130130312-3322213313223010-2321000120031233-1003130313222231-1020300333211303-3031311032310112-3212202220120120-3221011311100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

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

<a id="canonical-1113133120202223-1022303120003311-3112222232223321-3232211120002030-0103200122122123-3022000201310203-0110121311103213-1003231022022322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` properties

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

<a id="canonical-2103030113132301-0301222332011111-1121233230131310-2123221130233021-0332220000130213-1221012000111321-2300122220133232-2233223111131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

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

<a id="canonical-3233220130010331-2032012212000020-2313232310022133-3133230003311210-0000310201313221-2220112312012011-0012121021113322-2312030230120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` properties

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

<a id="canonical-0022102133201032-3321113032333101-3331132223220132-1302203111101301-3032120221020231-3000321310112230-3220320323312202-2200033122032220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` properties

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

<a id="canonical-1321101233023113-2230303111330010-1322303301031233-3302120010112303-2213220032222020-0212033123332010-2201112211201312-2330013331202030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` properties

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

<a id="canonical-3000332003321310-2222200123021310-0022123022030213-2300033332121010-1232020131322111-0023111020201203-1013103313120120-1213331301010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` properties

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

<a id="canonical-2302100211010213-1221110132232310-2212203011100023-1232220102103322-3120330203030100-3002113110022323-2221303300000213-0230130110100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` properties

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

<a id="canonical-1000020120232101-2030110031031121-3030111111001301-1002231333222012-2332210130333100-3130120123310313-0222002231213031-0122222322212203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` properties

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

<a id="canonical-2020123233023101-2230130011033221-3033131313202111-0310303102022031-0323203012133232-1322333020322201-2222332222012010-2333233300111032"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3200001130303132-3013213021102132-1123000331211133-0313021202000333-1213101310333201-1221333111130233-2121323122111112-0100333023201101): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0103203201123021-2323131121130323-1110013233332102-2112103023202201-1232030123210122-3301032300123331-2100001310120210-3330011230230312): complete subsection reference.

<a id="canonical-3200001130303132-3013213021102132-1123000331211133-0313021202000333-1213101310333201-1221333111130233-2121323122111112-0100333023201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

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

<a id="canonical-2303032113112131-1110022103132000-0010200303231023-2012020221013312-2320302230021300-2232121331303033-3031000310310331-0001200020123332"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0321222212300231-2212221100131023-1233200021032232-0102231033011212-0113122232020030-2320020003123132-3022200320233231-0011312231213003"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1000033231102332-3331101220100303-2231220202321201-2312310233111211-0002023333111332-1133101321323022-3312313133310231-1201112101020100"></a>

<a id="canonical-0321012103223010-2000021310101223-0301113103123031-1003133102131330-1010001011000010-1010102223111122-0101202110221323-0120333212113221"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3222223232013300-1011122202303010-1113102100300303-2010201110333222-1223221013011323-3220023221321200-2123131310033301-0302233330000111"></a>

<a id="canonical-1200210003023220-1232002233221331-1303102103110202-2033322131311201-1113310302010300-1030222022310101-2301233103331230-3032133201101121"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0103203201123021-2323131121130323-1110013233332102-2112103023202201-1232030123210122-3301032300123331-2100001310120210-3330011230230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` properties

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

<a id="canonical-3221111023121311-3000021110301222-1213122100011310-0213012202120103-1110221232013101-3023312130213122-3233230333130011-1120231311301111"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3303201130223220-0203303223302211-3001021132200222-1222203131300213-1222111000010232-1103003211020300-2132223000223002-2121332230303121"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0331003312223320-1111131030123311-0200321322003221-1321210120132220-2131022001331130-3003000113011312-2211001203111333-3223133323310012"></a>

<a id="canonical-2212203133233031-3201320230333111-0230203133322333-3000213133233203-2320003222110222-2210202131323301-2032203010133032-2220100033032303"></a>

#### `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-1203020210021221-3100021212030223-0112223000011011-3030123133121320-1131332121102312-0101112333012222-1201023112032320-3300013210320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add` properties

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

<a id="canonical-0220102322121130-1123232112111100-3201112232302010-1131212032033000-1331212233103012-0022101302021311-3111210220302000-1132312033200230"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add`

<a id="canonical-1323013212323222-0133321011100133-2333213202320102-1223333010032010-2223003322201213-3032033112221002-2303013121220003-2300013210310010"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` property

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

<a id="canonical-1012103112123222-1301212121322131-2311113300033002-1002302330012130-3233130120211102-2020022212321211-1302220202233010-3010312200232002"></a>

<a id="canonical-2331002131113232-0022020012132333-2012202121113323-1300000201312032-1231130211212201-0312132333013023-1202330201110132-0332130123331003"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` property

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123): complete subsection reference.

<a id="canonical-1230313323202122-2020220200321131-3322123220303122-1232202302011313-2301030330311020-3233312120032133-0111213333132301-1022121123202203"></a>

<a id="canonical-1200133121312032-0000102231022222-2120130013011130-2302232122330323-0133312023033132-3222111102120220-3000020110132230-0310101321023310"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` property

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

<a id="canonical-3000130031033120-3310322321101133-2203122303210320-2202323201112322-1030120132130030-2111333222202232-2321132032112132-0122021012333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` properties

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

<a id="canonical-2200203133033313-1020002223001101-3001032313013322-0003002222332023-2322313320132312-1121220033013210-2200302113001130-0322121202131032"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-3131013011323133-2030001313001023-2033103301222120-0001023000020223-3322030121211210-2223310202330202-3112213313113220-3110303030312102): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-3003011133003003-0230110223002131-0000130130221012-0320203010112020-2222332312333231-2013021021033332-3322130120200021-2101013222223203): complete subsection reference.

<a id="canonical-3131013011323133-2030001313001023-2033103301222120-0001023000020223-3322030121211210-2223310202330202-3112213313113220-3110303030312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` properties

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

<a id="canonical-0002002011211031-3313131311010221-0022223322121302-0110020201012333-1233010111130002-0102303230220222-3232232011320222-3020001020311100"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-2203122130223323-3003210121230023-1021220133033131-3320003312111331-3330030010033020-0113112013231021-3232011211001010-3133000233233300"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0101123010030322-0033102211302131-2000200321033131-1322131201111022-3032102003220133-3011213101231032-2002232201130011-3212312321220301"></a>

<a id="canonical-1223110221002102-1200333220212330-3022011302302322-2323333310121332-0232310233032300-1220222000201121-3120200303313233-2021133123130202"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3133303311122322-1310223103032302-2013010233303330-1312322111100110-3223200101223111-1020300112103021-0210103310023021-0121221010030221"></a>

<a id="canonical-0000100133130013-3333032303303331-1322313111202202-2323231310110332-3301232211332012-2211003023012232-3021020303322133-0301302222222311"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3003011133003003-0230110223002131-0000130130221012-0320203010112020-2222332312333231-2013021021033332-3322130120200021-2101013222223203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` properties

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

<a id="canonical-3210021123312020-1131202230031220-0010302112333213-3103001333220320-1120211223212203-0221132323111210-3100301202113003-2321122110023132"></a>

### Direct properties for `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3102201322001013-3203201013031202-3003333023302012-3131310322031333-2103002121302113-3322232313022210-0300102131110032-1133032300032321"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202110312310200-1031310233010010-3110311022210110-2133120010113003-0132021131000010-2320003022322121-1013121223031113-0120221031331221"></a>

<a id="canonical-1120121311213220-1331022100011200-3020030100101301-0111113220212323-2011131233022333-3002120233232320-3030011030323323-3200131002031222"></a>

#### `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- dynamic_proxy.https_proxy

<a id="canonical-1222311003032122-1302001312111013-2110112322002123-3320333003001322-0310220221331213-3032032120010010-1010200111113120-0300001233302320"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1211100010310322-3013313121213021-1320103312202010-0322202323222300-1112323221122012-2310233300011230-1100113031301132-2220003201013202"></a>

### Direct properties for `dynamic_proxy.https_proxy`

- [more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121): complete subsection reference.

- [tls_params](data-sources--proxy--reference--group-003.md#canonical-0033001023212230-3221103130030023-0130000002100302-3132121111030230-0113213112323330-0013122132112221-1203201333333221-2322013012200133): complete subsection reference.

<a id="canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- dynamic_proxy.https_proxy.more_option

<a id="canonical-0330330032102302-0232000313232301-3130033233130020-3331220030003003-0300111231313010-0220033130201110-3321033000213100-3331323221103331"></a>

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

<a id="canonical-2203023111122002-0233123333310313-2203312103212201-2012000313122011-3031113100013101-2121113231031210-2320100002210033-1102120323121133"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option`

- [buffer_policy](data-sources--proxy--reference--group-002.md#canonical-1303102010212200-2123321202223330-2000222103103030-2333113231120301-2311212220101110-3310112001323202-1011333131023301-0003122130300303): complete subsection reference.

- [compression_params](data-sources--proxy--reference--group-002.md#canonical-1300300133233300-1223221223133013-1022301301030100-3330231313221320-0113302330220102-2031110231311202-2122210033231312-2221213333100110): complete subsection reference.

<a id="canonical-1302320102022003-1333333000110020-1012331230132103-1222010333210330-2202321101220122-3232101221031000-0102132312131302-2111301003330301"></a>

<a id="canonical-1301312121203032-0121110022103212-0010121221132232-0113312020000001-0131121220002113-1210131231020001-1020130303110223-1113122130233320"></a>

#### `dynamic_proxy.https_proxy.more_option.custom_errors` property

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

<a id="canonical-0223222221302010-0121102012211230-0121112100120322-0123230212103030-0002101113013302-0030330321232012-3001300102033103-1133102033002212"></a>

<a id="canonical-2030233311221333-3033022301111301-1131020030133202-1122310212221323-0023013132023313-1332002321030022-1230120122320301-3200300222102222"></a>

#### `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` property

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

<a id="canonical-0330111222120001-1121011220011021-3332102321000210-3120101230211103-3310032110222220-3100223020302300-2202333200312133-1322103301331221"></a>

#### `dynamic_proxy.https_proxy.more_option.idle_timeout` property

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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-0002213032113133-1200220100003330-0231031223001001-0330113303132221-0202112113332311-1112120120223211-1213301131311022-2111121131311332"></a>

<a id="canonical-3113123030012111-3232013030322100-1231133232031021-1211202020010013-1321020230022000-2320211232002223-0121021310112200-0313222200232230"></a>

#### `dynamic_proxy.https_proxy.more_option.max_request_header_size` property

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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-2332110313100232-1131222210031023-0303012112310203-0323102202300320-3303200030222001-1001023231031120-3011000223203020-1032200300223022"></a>

<a id="canonical-1002201033111103-1322011332311200-3030303213030112-0223031123312312-1201211111011212-1221010120310301-3321023030112113-0020213203221231"></a>

#### `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3103301222021123-1311002313233330-1221222130111000-1212312111101001-0033322032303130-3223123022001103-3200320003331321-1110131123313002"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` property

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

- [request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230): complete subsection reference.

<a id="canonical-1212333223303331-3022333132132112-1203003331023323-3032033020311210-2331301132021322-3211312001213031-3303322313130032-0303311302102031"></a>

<a id="canonical-2103211310323102-0322102333031230-1221100021231123-3032331000020211-1032120131032320-1333113311331011-0200221230223212-1100113223222032"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` property

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

- [response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332): complete subsection reference.

<a id="canonical-0313202333230232-0103223103102030-2202010300303120-1313132013333101-2332322122312001-1122100232002210-1311122231003100-3313300213211212"></a>

<a id="canonical-3111123323002311-2102021002011102-2303300310220230-3103331021200003-0232131221331200-1210230030321130-0332311133221312-0310132311230010"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` property

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

<a id="canonical-1301110100030200-3013121121302020-0032330101133120-2000130311321121-1130102020031010-0130002312323110-2202100110313332-2331213020112202"></a>

#### `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` property

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

<a id="canonical-1303102010212200-2123321202223330-2000222103103030-2333113231120301-2311212220101110-3310112001323202-1011333131023301-0003122130300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.buffer_policy` properties

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

<a id="canonical-3231213131333233-1313311113113002-0310021000300122-1131331121202113-0021001113001302-2033200110211112-1022323322321122-0310211323103322"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.buffer_policy`

<a id="canonical-1203121332032200-1223221101111131-1310112222102202-3133233200132133-2313110030313103-3011123213332230-2130123013213113-0010002233313120"></a>

#### `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` property

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

<a id="canonical-1021131103002220-0300022023301210-3200010011320302-3120120332300130-0211122032211102-0113003022022230-1001322131012303-2131230101131102"></a>

<a id="canonical-0222131121220300-2013002302011301-3200122133112300-1320331212031220-1122210022133002-0302033322323033-2013222212222003-0101231202011220"></a>

#### `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` property

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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-1300300133233300-1223221223133013-1022301301030100-3330231313221320-0113302330220102-2031110231311202-2122210033231312-2221213333100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.compression_params` properties

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

<a id="canonical-1003222131223103-2113032033303120-1202312131303021-0302221323210113-0101313001013202-3300033002031030-1101322121212220-3200113323123023"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.compression_params`

<a id="canonical-1311303312130301-1023002133002212-3010220210302301-2103223022001132-0022032231332033-3112010132021102-0031312001202101-1232331232003220"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.content_length` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2101222333130201-0323031013231101-2013031123231113-1210203311021123-2130332031112202-3110132220213103-1331231312312212-1103222210301113"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.content_type` property

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

<a id="canonical-1332112112121113-0310310122023123-1020110120223311-1322322233203210-0002000133103120-3201020331002101-0313203001312303-2231133313110130"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` property

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

<a id="canonical-1010221101013213-2032320220222230-1102100011002203-3311200031013012-2230101231012133-1033130330112113-1202231112020231-1000111302331112"></a>

<a id="canonical-2310230012100123-0310321330133210-1301102303100003-0023131232311110-0301033201020331-3323203021030103-1301301321010311-1013131203332121"></a>

#### `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` property

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

<a id="canonical-2021330300201200-1321031132112312-3210210020302020-1201320002200120-0222123032233313-0030333330021332-0301022221032211-1131210120113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.disable_path_normalize` properties

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

<a id="canonical-2223113323111313-1020333300201031-0322032313321010-1012200231231223-0123200303221002-3233003213023203-3303020003213301-3011033222002303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.enable_path_normalize` properties

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

<a id="canonical-2221323221232001-2003200310122001-0121102203111132-2312131321132231-1022000002200020-3333232001203113-3320123302202300-0300313200012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` properties

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

<a id="canonical-2331010312013203-1212022320011332-2123223230102130-2211033330023211-0020232220200331-0211110112013300-2011131011312101-0001213133022210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` properties

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

<a id="canonical-0110323233012113-1122113021223100-1113321230232302-2132303233023011-0320333133131210-3130120132113132-0133000001321000-3330000203122010"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add`

<a id="canonical-1100023031122013-1331003032000011-2011322302303301-3203312230121030-2233200223203200-0300320203032202-1222001122113000-2020031023003013"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` property

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

<a id="canonical-1130010121100103-3011233012203111-2110212110322131-0010312120010233-2321110312001212-3333001021001011-1020331130031320-1300002110101210"></a>

<a id="canonical-1112132001010210-2021123303021011-3203321220331203-2103202122200113-2130301302132303-1311223203031231-3031103221130313-0212330201002202"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` property

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003): complete subsection reference.

<a id="canonical-3021100312103311-2102033330200311-0013322122301123-0122110333211210-3323112320311102-3310332223011132-1000001312233221-0231020102213213"></a>

<a id="canonical-2032230023011302-1231131032313212-3132000303033032-3201303312132103-2323311133331131-2130200310100323-0221100221210313-3022021023213200"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` property

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

<a id="canonical-2123303102012010-3131311120021120-1102021232230023-3112122132012001-3111010021312102-1332111120220002-1210221302100312-0111302302023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` properties

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

<a id="canonical-1320030221123132-3131133020211110-2012121100221123-0113221330313123-1101031133131220-2211310221231202-2323231211033012-2112010123313310"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1301330002000203-2311011011030020-3202222132233330-0210220323030111-0201130301212230-0013202310131120-1102320300210211-2030312223223132): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-0002213211333012-0301110220233033-1022321002103003-3230303322201203-1102022101031132-1230012130231310-0021333120201223-3131321021013010): complete subsection reference.

<a id="canonical-1301330002000203-2311011011030020-3202222132233330-0210220323030111-0201130301212230-0013202310131120-1102320300210211-2030312223223132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

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

<a id="canonical-1132022301130302-0021322123120020-3300333132210211-1120222202230033-2221302131213010-3130202311131001-0130300122322022-1333201310131030"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3103002003010201-3021132302230011-3300010130223332-2101111230202210-0111103311123100-0111311311020023-2103010130023331-1220221010320302"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3132202111311010-3311010320310033-1122322303312122-1203012001322312-0210313310102133-3130203121113012-0013131301013223-2022311320301310"></a>

<a id="canonical-2230000133333211-3311233113100013-0013321320112322-0211201302002313-0203130133303210-0100033123300020-3023201313120120-2101132000232002"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-3001003132301223-1232013310232003-0133100211221012-1131300331221312-1110222122002303-2011332001312321-3233001331213330-0323313322212331"></a>

<a id="canonical-1220201332330311-2310002323311310-3223000031230100-2230300020222233-3132222201222323-2132123330313310-2301113121221332-3202311310002222"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-0002213211333012-0301110220233033-1022321002103003-3230303322201203-1102022101031132-1230012130231310-0021333120201223-3131321021013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

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

<a id="canonical-1221320221131021-1110321300020103-1113100121002030-0112311320101002-2112021030302002-3132333301313100-0110211323032123-2032132212130111"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3111300220203102-3031202013232321-2312111032100332-3001100302231212-1101020300111320-0302322211132302-1203010011220012-3301030131300131"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0021121303003233-3200122100110013-1133113020133311-0001003012301202-1102113310312132-2313320320203010-0233100132023220-2331203032301301"></a>

<a id="canonical-0021012003331033-1003203220133132-0330112233003222-1321032322223120-3121012332213000-2210322331101101-3112000032320101-3002113231032030"></a>

#### `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add` properties

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

<a id="canonical-2100030200122331-0121031300121213-2210001323032010-1233103230312212-1323013332030131-3213312332331111-2113001232223011-3112333013111000"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add`

<a id="canonical-2212000203223010-2011131120233121-3202112123223212-1202111112133130-2231013131032122-0310131023332023-3010222311220312-2102212020231131"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` property

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

<a id="canonical-2210312100202211-2013002111300220-0020103200032232-2200022311110202-1210002020220311-1012010313120200-2002231012021120-3013201032032113"></a>

<a id="canonical-0231123121010212-1031200332231302-1231220301123321-3212332223201331-1111132022200131-2201112323320233-3020011012011001-3333322201233310"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` property

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

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223): complete subsection reference.

<a id="canonical-3101202010010102-2013122031202021-0320233111012113-2202211220032300-3021322310300301-0213202323222031-1011311031121332-1321233313221013"></a>

<a id="canonical-3303133111101000-0232121123332233-2222012022310031-2030132203223333-2021312110103230-3313311301200130-0021322011032122-2110300212333103"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` property

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

<a id="canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
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

<a id="canonical-2312203131033322-2322313302303311-1102003021012221-1033022022212113-3102223223310322-3122032003222133-1110002002332220-3302103113031031"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-2031112112320023-2113300013121202-2023001332113201-1002000123330210-0121032211312110-3032133300101303-3300000213021300-2312330303231233): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-002.md#canonical-3112222122223320-0030311322322011-2101000223213320-3210330331031200-3311323031220032-3302233113130331-3312120333001332-3213031331112311): complete subsection reference.

<a id="canonical-2031112112320023-2113300013121202-2023001332113201-1002000123330210-0121032211312110-3032133300101303-3300000213021300-2312330303231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
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

<a id="canonical-0130000321120022-0102232222012233-2223300120020133-3230203013033121-1311121220101321-3300220130023001-3120230233301021-2123303101210203"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0322132111223111-2202133330222211-1221230231001330-3231222132300210-3010112231231102-2123131010101110-2202002103210011-3222132012300221"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2320313302110102-2210321030231001-2101100211020111-2230323322223010-0310213100130211-1000210020132333-0322013331220231-3120313220311312"></a>

<a id="canonical-2012023021212101-0213322123300010-0223302010010103-2013322231101300-3230202202231303-3201121110311020-2220003301330132-1010120111310132"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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

<a id="canonical-1123323133232202-3211333100130330-0020002032203111-1322101201000131-1220020331212222-3032331030312213-0321111332111123-2331200003212000"></a>

<a id="canonical-1111121203021312-3300113213020030-0202210132233123-3123203010310332-0301202330030312-2031333010301212-1003122230231331-2233302233330003"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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

<a id="canonical-3112222122223320-0030311322322011-2101000223213320-3210330331031200-3311323031220032-3302233113130331-3312120333001332-3213031331112311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--reference--group-002.md#canonical-3201322021120111-2230010312130312-1212001001230213-0331310010220210-2310322322312011-2233032200130302-2232011330313311-0021221110112230)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-1301221223211212-3330010320031222-0221213121222230-0333311320223133-3231212332201333-3221211221322320-1312220100003131-1311103232020223)
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

<a id="canonical-1123000302230123-1131211201003131-3013230132202331-3331200103323130-3021123210203311-2213010322023010-2112130312232102-3200311201312322"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-3123032203120132-1120012321001000-0003211232211102-3312131213333323-3032311131330303-2023032022123231-0003031300032202-1123200033130021"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2122312013300023-1221200303212211-1102101332313232-1332332111322121-2203113320303313-2012103022033212-1020333120132211-1133323213322121"></a>

<a id="canonical-2013002212302320-0110301222331231-2011131033222212-2132112032302212-2233022312332330-0103312233303033-1120102322001333-1210320203012302"></a>

#### `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

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

<a id="canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` properties

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

<a id="canonical-1022230331121330-3002221311131022-2331323102301103-3213010000313102-0202022003310032-1132132232130222-0020210200110002-0331001331121201"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_cookies_to_add`

<a id="canonical-2022112002030110-1222131032112301-1233132032302020-0120322100003302-2312102032222000-1332312211010230-3323001133230011-2232010333322202"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` property

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

<a id="canonical-1212032033300313-1301132213333302-3030222031003203-1322202300112311-2022221101100221-1110322100031020-3322033223130021-0122232100120202"></a>

<a id="canonical-1013013102233211-3312321330201301-1101132320010210-3012130132311001-1002000030310221-0331110131211231-3210021231231323-3301033220022200"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` property

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

- [add_httponly](data-sources--proxy--reference--group-002.md#canonical-3122230200320020-0203101011212112-1122120213302300-2310030121010020-2113103310313322-0030213121220120-1130321033312301-0201011231131013): complete subsection reference.

- [add_partitioned](data-sources--proxy--reference--group-002.md#canonical-2230311202132222-0220021001101113-1313002110233213-1303322302232111-3323121233121200-1313302123003032-0002313033020311-2001213003322312): complete subsection reference.

<a id="canonical-1210103000201213-3220112232320100-3330002123110310-2300011203210321-2210123033311210-3303213100123121-3000000213330323-1112320221323100"></a>

<a id="canonical-3001033010332220-1123123232030030-2300022332310233-3332033113322013-1112113321133020-1312110323021110-2021102301222131-3020013011212001"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` property

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

- [add_secure](data-sources--proxy--reference--group-002.md#canonical-1122110102003113-0201231213102303-0020300201220112-1222102303312001-2100031202323000-0000032132220301-3311200133123333-1121002013201130): complete subsection reference.

- [ignore_domain](data-sources--proxy--reference--group-002.md#canonical-1011223022130232-2213000332120032-2222200123303221-2122101103123003-2303123221003011-2300302132220023-1100123022122301-2333001300020220): complete subsection reference.

- [ignore_expiry](data-sources--proxy--reference--group-002.md#canonical-2100131121013311-2320001110331133-3102001101312223-1130310331113211-2112313320212311-0312012120323111-0020121300103212-3220033122233032): complete subsection reference.

- [ignore_httponly](data-sources--proxy--reference--group-002.md#canonical-3332320023202320-0023212233213011-2232311332000001-1311001012012110-2131000311011132-3000031312120200-2211003110213123-2031323332113112): complete subsection reference.

- [ignore_max_age](data-sources--proxy--reference--group-002.md#canonical-1021003112130310-1300012012002110-3321101000010322-3333030103303322-0330001321102311-3333111211302130-0310000210303330-1020232122301032): complete subsection reference.

- [ignore_partitioned](data-sources--proxy--reference--group-002.md#canonical-3120201003310332-1302111023203023-1301200121132200-3201021200113313-1210132121030221-3102312000200222-2032230323130000-0111220213101202): complete subsection reference.

- [ignore_path](data-sources--proxy--reference--group-002.md#canonical-2101230011122320-2101010310122300-0011120122303032-1111032230002223-2020201231332123-0031102301003313-3301111122131121-0030212001123201): complete subsection reference.

- [ignore_samesite](data-sources--proxy--reference--group-002.md#canonical-0301322122323102-1001130011132022-1213212213313111-0102220011102322-3210103332132032-1313111101213331-2322230101310332-3220312210231223): complete subsection reference.

- [ignore_secure](data-sources--proxy--reference--group-002.md#canonical-1030330302332322-0233031021220000-0100332101130121-1233000123120220-0203102112111100-2323322101123323-2003300203100223-2101002213311112): complete subsection reference.

- [ignore_value](data-sources--proxy--reference--group-002.md#canonical-1332223103223212-3322102032333023-3332322110323033-0103203131103223-0121033323102203-1303320032031201-3203102002320120-2303233220233220): complete subsection reference.

<a id="canonical-1212020221200133-2323012123101123-3213132102310110-2233230111231022-3102103021301310-2002221321223101-2021020021202002-1110002211310103"></a>

<a id="canonical-1211203123030113-3120121231311331-0302221131301313-0232120000210311-2223222100320131-3213233310133310-0230102022223012-2111013111133032"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` property

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

<a id="canonical-0110202313010001-1121133213333113-0101222132333212-3200022023312210-2310131213301221-1331200000123030-0323333323032232-3223210130201322"></a>

<a id="canonical-0310101220220211-1322311110131023-3231032200302032-0112122301111220-2021111331312002-2321100102013103-1332001122220200-2022310020103331"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` property

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

<a id="canonical-3000122130200330-3231321103321023-3331223223203131-1120233000131120-3333013101301112-0210113230232133-3200330101230123-0323223310210211"></a>

<a id="canonical-0123100000221231-2232301122022313-1212230202123012-2201211011322220-2013320111201223-0120213322231230-0200331323013001-2132221323211111"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` property

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

- [samesite_lax](data-sources--proxy--reference--group-002.md#canonical-3102032203001213-3032011002021310-1032011123113110-0102132012023003-2232232132122312-1101322232111312-3321312303222022-2232233320002123): complete subsection reference.

- [samesite_none](data-sources--proxy--reference--group-002.md#canonical-2212222330313101-2322111231313030-2332023121212222-3030001220232220-1130332111302120-3311322001200001-2122312100001220-0102023303131222): complete subsection reference.

- [samesite_strict](data-sources--proxy--reference--group-002.md#canonical-1321202301011021-2333122122030010-2212121313020312-0133303302201213-0023000333131221-0313003231300312-0011313331231210-0120322232202212): complete subsection reference.

- [secret_value](data-sources--proxy--reference--group-002.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130): complete subsection reference.

<a id="canonical-3032110323023033-2002113132223120-1200213300220020-0023233022202212-0012023221111023-0130313101113113-3001233302302130-2331033323302301"></a>

<a id="canonical-3013102233201132-1001331012333123-2013111010301313-2321030133200031-3112220032100122-2313032002110201-1211033121233333-0322311303123032"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` property

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

<a id="canonical-3122230200320020-0203101011212112-1122120213302300-2310030121010020-2113103310313322-0030213121220120-1130321033312301-0201011231131013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-1030022032310100-2200333123033030-1020121311301202-3232002113321223-3210202201221232-1012003202213320-3233013010301320-1210031031113033"></a>

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

<a id="canonical-2230311202132222-0220021001101113-1313002110233213-1303322302232111-3323121233121200-1313302123003032-0002313033020311-2001213003322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-2011313122003311-0223031331331003-2301310011200011-3112222330233212-1131120132013220-3123003001333333-3132230233110333-0311330131203110"></a>

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

<a id="canonical-1122110102003113-0201231213102303-0020300201220112-1222102303312001-2100031202323000-0000032132220301-3311200133123333-1121002013201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-0321121222223030-1311333233112223-2103133110020100-0201112112220123-2223311032212113-2120131010310131-1012112301223313-3212002002312110"></a>

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

<a id="canonical-1011223022130232-2213000332120032-2222200123303221-2122101103123003-2303123221003011-2300302132220023-1100123022122301-2333001300020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-1231030112203133-1010332211102021-3210120233303311-3320121112233023-2123300013011131-2002103011020201-1232121131310112-1320221000011130"></a>

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

<a id="canonical-2100131121013311-2320001110331133-3102001101312223-1130310331113211-2112313320212311-0312012120323111-0020121300103212-3220033122233032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-0233113013210303-2013112211312213-3113312111130111-2202032233310023-3213131133321033-2113332020222301-1331220220232222-0033110133120232"></a>

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

<a id="canonical-3332320023202320-0023212233213011-2232311332000001-1311001012012110-2131000311011132-3000031312120200-2211003110213123-2031323332113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-0230123113032310-2322121012122312-3031113000122300-2221222231011322-1311103002101030-1013101312023002-1230013310330002-1331021203130223"></a>

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

<a id="canonical-1021003112130310-1300012012002110-3321101000010322-3333030103303322-0330001321102311-3333111211302130-0310000210303330-1020232122301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-0131313331203112-2100101212131003-2020203023333111-0310100020310331-0310210320213022-2313111203333013-3003102002231012-0231303203003233"></a>

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

<a id="canonical-3120201003310332-1302111023203023-1301200121132200-3201021200113313-1210132121030221-3102312000200222-2032230323130000-0111220213101202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-0231110023133002-3110330031210010-0320233231333020-0302020122033101-0330232221310331-0103321113031111-1131131223313020-2310010123333320"></a>

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

<a id="canonical-2101230011122320-2101010310122300-0011120122303032-1111032230002223-2020201231332123-0031102301003313-3301111122131121-0030212001123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-2030032232231213-3130313302001010-2300312000022300-1132120112323211-0231113112103320-0302023001012301-0322012212132303-3013222333100010"></a>

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

<a id="canonical-0301322122323102-1001130011132022-1213212213313111-0102220011102322-3210103332132032-1313111101213331-2322230101310332-3220312210231223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-2203303222021322-0201021023033333-2123103311320302-1331302322030200-3301022300222132-0103101123022023-1001321330113020-3321332123013221"></a>

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

<a id="canonical-1030330302332322-0233031021220000-0100332101130121-1233000123120220-0203102112111100-2323322101123323-2003300203100223-2101002213311112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-2200022321221103-2123330123332320-2311233131232032-2023203213210230-2101311320111103-2332003031300032-2100012103012022-2203222011323222"></a>

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

<a id="canonical-1332223103223212-3322102032333023-3332322110323033-0103203131103223-0121033323102203-1303320032031201-3203102002320120-2303233220233220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-3020130010203313-2210332222101212-1322222213313300-2101011023020002-2202022313232223-2223301103300232-2232102331233001-1020331111212232"></a>

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

<a id="canonical-3102032203001213-3032011002021310-1032011123113110-0102132012023003-2232232132122312-1101322232111312-3321312303222022-2232233320002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2222110232022130-1222222010113003-0200310333010113-1220212012200300-0123132111000303-1001223013233230-1210320030233102-1211002312322133"></a>

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

<a id="canonical-2212222330313101-2322111231313030-2332023121212222-3030001220232220-1130332111302120-3311322001200001-2122312100001220-0102023303131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-0312001230223020-1002031033023122-1100130220123213-1310030333202222-1003211123331333-0331112021100212-1002312132332031-3203220321233001"></a>

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

<a id="canonical-1321202301011021-2333122122030010-2212121313020312-0133303302201213-0023000333131221-0313003231300312-0011313331231210-0120322232202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1320223211213201-1223210201030313-1102123301201132-2100323030100020-0232022333320012-3200033230230033-0032330000330111-3213232300121022"></a>

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

<a id="canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
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

<a id="canonical-1333200003120203-2222111013010110-0201021122233300-2313201133000020-1003222333313000-0000333130013211-0122210020322323-0201003223123221"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--proxy--reference--group-002.md#canonical-1320221103311302-1333211323122013-2321233100010221-1021310323032121-0232220333021210-1323100023120001-1220033213233333-0312032322111111): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-003.md#canonical-0111102122200200-0231211332230122-0002101102313301-2021122122203210-0203231033302312-3222322203312130-2020310123200302-2303323322131300): complete subsection reference.

<a id="canonical-1320221103311302-1333211323122013-2321233100010221-1021310323032121-0232220333021210-1323100023120001-1220033213233333-0312032322111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [dynamic_proxy](data-sources--proxy--reference--group-001.md#canonical-3120113220030001-1232203001002230-2202312203031103-2000132310120011-1311003013300333-0133221002300322-2112330320310321-2100002331020313)
- [dynamic_proxy.https_proxy](data-sources--proxy--reference--group-002.md#canonical-0201021123310321-0330023030131223-0202200010121321-3112301033310021-2303020312222211-0302223310002100-3322211113021112-0122332132001220)
- [dynamic_proxy.https_proxy.more_option](data-sources--proxy--reference--group-002.md#canonical-3301230313010011-2110321222023223-0220203332233312-2133302331020323-2201021021122033-1000121032013102-3123100333023003-1110113102021121)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--reference--group-002.md#canonical-2023113131120033-0301233010122110-1221200023021113-3131221123032331-3231122000022200-1022012031321210-2122000323221010-1001012312311332)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--reference--group-002.md#canonical-3030320030322220-3302000303100200-2210230212332331-1200210023310332-1331020331110131-3331120111310113-1213011203312300-1330330203132130)
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

<a id="canonical-1210320330331021-1003331132133230-1221211130333203-3303111221031201-0333301210003313-1002333301122022-3131321010230021-3231032320332002"></a>

### Direct properties for `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0111120030232022-1323011110111102-3033202332200303-2212102123010221-1211132233012111-0203330322001021-0302132221231323-3030113031020103"></a>

#### `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3020112322000122-1122021322330203-1320311223332311-0331220232120013-2213313102123132-2000202211132200-1120321122023313-0013101230310330"></a>
