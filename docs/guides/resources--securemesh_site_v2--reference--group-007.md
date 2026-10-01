---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1121121110102310-2222113010011120-2312032302103031-2203320110212223-0032322010133001-0313121111213220-1002313311021222-1133302133332312"></a>

## custom_proxy.password.blindfold_secret_info — blindfold_secret_info / 022233122123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- custom_proxy.password.blindfold_secret_info

<a id="canonical-3213132010110031-3033323101123123-0203300322101322-3200012000223111-3030033331222011-2311112033123332-1322231231121032-3313322323033221"></a>

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

<a id="canonical-1130113320022310-1320001022322312-2022212112011132-0032111302313110-3010230102213321-1221201101033100-0103211210033211-1132321020120302"></a>

## Direct properties — blindfold_secret_info / 022233122123 / 3

<a id="canonical-2113201021003313-2023323221131131-1020131300122022-2111231030122210-2321022033030113-1120030020303302-2121012223203221-1011223032011002"></a>

<a id="canonical-3122312311203312-0030333013231221-3300311220222301-2220132311021213-0203231310030120-2232200012001231-1011311220031112-1121111200202310"></a>

## decryption_provider property — blindfold_secret_info / 022233122123 / 4

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

<a id="canonical-0000012230300123-0101333033030111-2232203222223233-1032202120200330-3220120112101131-0130311131212120-3332113011201301-1230012001103001"></a>

<a id="canonical-1113221030322322-2013223012311012-1002020133200231-0121330211003022-2133313221001032-1312313313002001-0302133012122133-2232310133331220"></a>

## location property — blindfold_secret_info / 022233122123 / 5

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

<a id="canonical-3011321001133133-0033201011210002-2332303212111201-1322121113031212-2131100112003313-3320112333203001-3121102200032021-3312203323120301"></a>

<a id="canonical-3313202322222320-1313311031321110-1113103300332031-0002032303100131-0311331122310013-3112123233310221-0211011320022030-1101203203303330"></a>

## store_provider property — blindfold_secret_info / 022233122123 / 6

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

<a id="canonical-0023332210031202-1311020110213330-2130322222213102-3102031333300131-3000320231230222-3023313222022203-2210010120023213-0312221221312220"></a>

## Next pages — blindfold_secret_info / 022233122123 / 7

- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3223030303103201-2313210133333313-1120221112033100-1332302313313131-2023322120302103-1310021221110300-2130010231132310-0131332332031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021012112202220-3311323222220021-3303301311002231-0212031330331030-1012230321121210-2230021001322213-2131323202312123-2321220223020233"></a>

## custom_proxy.password.clear_secret_info — clear_secret_info / 113311230130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131)
- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- custom_proxy.password.clear_secret_info

<a id="canonical-1232210121101021-0213110130101122-2232202030310230-3330330020113033-3333312333033232-0213312133100203-1131003031221312-0322022320132023"></a>

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

<a id="canonical-3020210031210120-3211013331111123-1012220102202031-0101010022123203-3020220122212133-1321022221100220-1311221022331220-2101103203033002"></a>

## Direct properties — clear_secret_info / 113311230130 / 3

<a id="canonical-3102321103020112-1011031211001022-3201332310323320-2100203300101023-0301330121020320-1230002003003313-1202333233200133-3131332222333111"></a>

<a id="canonical-1311020000322333-1201303301310302-2031032210300130-3003002031331230-3032221101311230-1313220101101021-0333321302103130-2032101232121312"></a>

## provider_ref property — clear_secret_info / 113311230130 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0223032110210203-1301020112322212-0332131212310112-1313313003301311-3100132332100303-1210001023111112-2131012202122012-2013203221020103"></a>

<a id="canonical-2030220133223000-1011113322210302-0030013233231213-1002220102212210-1332320003300023-0233331132101200-3212002001012101-2200300333132223"></a>

## URL property — clear_secret_info / 113311230130 / 5

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

<a id="canonical-0031322233322330-3221303230023321-2321030222223310-2331012331110332-1013110203102333-3020233030313123-2220133120221112-0321221101332110"></a>

## Next pages — clear_secret_info / 113311230130 / 6

- [custom_proxy.password](resources--securemesh_site_v2--reference--group-006.md#canonical-3001030210312201-2020010031322113-3001013121310303-2030021121311010-1300301111130220-0312313102131210-3231122222311300-2113320131233010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1201120122021230-3013113003303010-0312222210310120-1032233221132010-0003011010201111-3212223322123210-1233330103221220-3111220102212230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112131010001000-3033221133301010-3322111110120103-3320120122100201-0231100003201213-3200332233122021-1301230313123130-0200131112320201"></a>

## custom_proxy_bypass — custom_proxy_bypass / 101122131300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- custom_proxy_bypass

<a id="canonical-0321121310031302-3001323333210001-0332113021331022-3202330233302121-1322302022300111-3020133013030000-1333032113321213-2300231103033123"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_proxy\_bypass, no\_proxy\_bypass; Default: no\_proxy\_bypass\] Configuration
parameter for custom proxy bypass.

Upstream description:

List of domains to bypass the proxy.

Receipt-pinned upstream constraints:

```json
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

- [custom_proxy_bypass](resources--securemesh_site_v2--reference--group-007.md#canonical-0321121310031302-3001323333210001-0332113021331022-3202330233302121-1322302022300111-3020133013030000-1333032113321213-2300231103033123)
- [no_proxy_bypass](resources--securemesh_site_v2--reference--group-012.md#canonical-3022023002231032-3101230312322203-1210102303113020-3133121101220113-3332021101223101-0031033231310313-2221010131313222-1131033300021022)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_proxy_bypass {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102200010312303-0031312020223203-1323330303132121-3001110110130003-2303332013311103-0022122320301223-3030231100220022-1232201302000233"></a>

## Direct properties — custom_proxy_bypass / 101122131300 / 3

<a id="canonical-1133113202202210-3220332300222203-0103312001213200-2310121320113322-3022312001302021-3120002212002331-0231013010032012-2011102213323313"></a>

<a id="canonical-2031032233200321-2003302120123303-3220201012101212-2331001211333000-2020111130233112-2121320123233301-3211111321333322-3321300001320113"></a>

## proxy_bypass property — custom_proxy_bypass / 101122131300 / 4

Type: `["list", "string"]`. Optional.

Proxy Bypass. List of domains to bypass the proxy.

Upstream description:

List of domains to bypass the proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname_or_ip": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312021022000022-1201311033321301-2120023202322130-2303132213011332-2322231020012310-3030102201232213-3133123111011212-0203302103001112"></a>

## Next pages — custom_proxy_bypass / 101122131300 / 5

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1020321133311233-2110133311213121-2123333100012012-3200301200123213-1020103130321222-1021002013333211-1002020122222102-3103333110002323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133332103021202-3113322212201121-3120122323302031-3110011130220331-3200210330312233-3322320012112320-0231213001012300-1332313113123133"></a>

## dc_cluster_group_sli — dc_cluster_group_sli / 023320120203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dc_cluster_group_sli

<a id="canonical-1020300212301120-1320101102310232-1010130313301102-3202331322202133-1101131123210202-1101022221021110-0300103310332221-1303203313130333"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group\_sli, no\_s2s\_connectivity\_sli; Default: no\_s2s\_connectivity\_sli\]
Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [dc_cluster_group_sli](resources--securemesh_site_v2--reference--group-007.md#canonical-1020300212301120-1320101102310232-1010130313301102-3202331322202133-1101131123210202-1101022221021110-0300103310332221-1303203313130333)
- [no_s2s_connectivity_sli](resources--securemesh_site_v2--reference--group-012.md#canonical-0330303002222222-1213002233022103-0000311222000331-1230302330001022-0022102302021300-2031322212021101-0113011320113023-0212000011230312)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312310313321031-0202000302233013-0312221301202003-2201100310310321-2322300020233211-2022103020203123-0301011333210232-1020012220333232"></a>

## Direct properties — dc_cluster_group_sli / 023320120203 / 3

<a id="canonical-1211122111111332-3221001220320302-3213333311000133-2310131333003312-3303000222022000-0000211312311312-2303320022012233-2313111312232323"></a>

<a id="canonical-3012022013303003-0011021023201302-3031213321030313-2133123003312032-3313331023312023-3021132300021102-0333203102121011-0210002300331123"></a>

## name property — dc_cluster_group_sli / 023320120203 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1031313320003220-1201021202120131-0033123132102103-3123003122012033-1312103330221301-2332121302323213-1002003220101022-0132211312221030"></a>

<a id="canonical-1121203310221030-2011133030300322-1132132023322202-2333103021320112-2313000033210011-1110332130232230-3312222311310200-0001103020302131"></a>

## namespace property — dc_cluster_group_sli / 023320120203 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3133103022012023-0031122133021303-1033322023302320-0312323030100222-3322010301333023-2112320210300031-0202203213310213-3221030202001211"></a>

<a id="canonical-2330302232100120-2001320310312003-0232203011100223-3002202102220302-3111213302232130-2020322313120011-1322101320222333-0032033221311103"></a>

## tenant property — dc_cluster_group_sli / 023320120203 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0300232301021131-3133320322012121-2230022131333022-3132132302111020-2033313200000020-3330331030220012-2231100020102131-1333331320133001"></a>

## Next pages — dc_cluster_group_sli / 023320120203 / 7

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3100323033003130-1302331203022103-1233101210123010-3220222322211011-0332110232310011-1003002331302211-2010313222320033-2132103002200131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232033113111320-1011133201032110-2012312231223321-2200333230233332-1230202100233333-0121323211120210-3223333321002302-0213221002333231"></a>

## dc_cluster_group_slo — dc_cluster_group_slo / 012121212002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dc_cluster_group_slo

<a id="canonical-3120230010200133-0213301102030000-1103030023223132-1122222023103023-1312200222303103-3312320310320112-2131211200103212-1021200322100033"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group\_slo, no\_s2s\_connectivity\_slo, site\_mesh\_group\_on\_slo; Default:
no\_s2s\_connectivity\_slo\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [dc_cluster_group_slo](resources--securemesh_site_v2--reference--group-007.md#canonical-3120230010200133-0213301102030000-1103030023223132-1122222023103023-1312200222303103-3312320310320112-2131211200103212-1021200322100033)
- [no_s2s_connectivity_slo](resources--securemesh_site_v2--reference--group-012.md#canonical-1030303003312203-3112321311031221-2310332332000120-1233313121311313-2111010213232133-2302111301212233-2211333200111312-1310302310103121)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-3212222333120111-2230203121211023-0313133130013232-2310301323102000-1222112001023100-3302312232123100-3211022311121112-3323310000112231)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group_slo {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210021333201121-2302003101010111-3320230313233023-2210030120102212-0103001211032011-2131310330100003-2010311112330301-2211011021120332"></a>

## Direct properties — dc_cluster_group_slo / 012121212002 / 3

<a id="canonical-1302220003223223-2222313131020001-0103202120003031-2321110213333020-2321133011101020-2313101210022103-1012201123030121-0022132312101322"></a>

<a id="canonical-0230332013233331-0022120010222011-1123032322312113-3010200331102022-3223322303013110-1212222232101013-0330122220223332-1221210302310200"></a>

## name property — dc_cluster_group_slo / 012121212002 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1310300103011313-2321012221001100-1113330113131202-0320101301002321-3332210102200110-3030231313320121-2033301122200100-2022231310023322"></a>

<a id="canonical-3121210030220003-3233213221031210-1130023032313102-2233132310113222-0111323312333203-2110111320302132-0100301122113302-1322300232113032"></a>

## namespace property — dc_cluster_group_slo / 012121212002 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0231101301113003-3101022323111210-0031132310312132-1321013231010303-3100313210033311-3213000212210230-3201333231321010-1003322112122232"></a>

<a id="canonical-1230203201111330-3232210121200022-3013102120021301-2023031302103023-3312132112313303-2101221312001300-1312033112330212-0310202113312203"></a>

## tenant property — dc_cluster_group_slo / 012121212002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2022222321103120-0023230133111311-1313132223312200-3013211221200031-1222312111121111-1213003112121331-3230022003033203-3332330020322322"></a>

## Next pages — dc_cluster_group_slo / 012121212002 / 7

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1112203010312210-0132130311102202-0022222202032000-2131322332010323-3210102110223000-3200001233331101-1223311301211212-3200100111113210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320313010302312-2333122111101232-0112120211132300-3120133303321300-1332102213201331-0202132333000222-1123020231112102-2030312022320031"></a>

## disable_advanced_delivery — disable_advanced_delivery / 021232000130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_advanced_delivery

<a id="canonical-2122201323223101-2113232031003001-2120123002330002-2031202123102313-0110121120013130-3201031113231212-0023001220102221-0001120132122013"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_advanced\_delivery, enable\_advanced\_delivery; Default:
disable\_advanced\_delivery\] Configuration parameter for disable advanced delivery.

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

OneOf alternatives in this subsection:

- [disable_advanced_delivery](resources--securemesh_site_v2--reference--group-007.md#canonical-2122201323223101-2113232031003001-2120123002330002-2031202123102313-0110121120013130-3201031113231212-0023001220102221-0001120132122013)
- [enable_advanced_delivery](resources--securemesh_site_v2--reference--group-008.md#canonical-1220131302000301-1122033110230313-0322113111321333-3220130110202322-2023020022301130-2201202233132330-3211202131020111-2200131101320320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_advanced_delivery = {}
```

<a id="canonical-2310103303030113-3110311132222130-3332133023123033-3310301211313033-0022001012200011-3110011102213332-1322202112110232-0230213133002033"></a>

## Direct properties — disable_advanced_delivery / 021232000130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213312231220322-1020313020000033-1300103333202303-1212220302233120-2132130230302010-2202033323310333-1201012010310113-3323222112313032"></a>

## Next pages — disable_advanced_delivery / 021232000130 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3231002302323332-2013010223311332-0021120301210112-2232331120122110-0301210303200310-2002220333321202-2022103131122023-3101003100301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110222100132301-0000300011231201-1001031330310223-1312330101022103-1013122021220211-0323020013220302-2033131320032201-3203322013321002"></a>

## disable_ha — disable_ha / 310131200330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_ha

<a id="canonical-2320203132102000-0303312322222212-1301210010311110-2023032100003222-1021311001131211-3210210003302032-1322200012330113-2132221002021220"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ha, enable\_ha; Default: disable\_ha\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_ha](resources--securemesh_site_v2--reference--group-007.md#canonical-2320203132102000-0303312322222212-1301210010311110-2023032100003222-1021311001131211-3210210003302032-1322200012330113-2132221002021220)
- [enable_ha](resources--securemesh_site_v2--reference--group-008.md#canonical-2330001010130310-3032020220122111-2022223300332103-0310000223001331-1202013231320033-2100001200103323-1020221333133201-2231302313002023)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ha = {}
```

<a id="canonical-1211131110012323-3113003233310121-1131031020011113-0321212330220020-0111003123202212-2101101220012013-1200033012211230-3320202031321311"></a>

## Direct properties — disable_ha / 310131200330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312121221011133-1220101300123100-2033001301310110-1331020033012032-1213122110132112-0002231010333112-1213021131233313-3013210112311221"></a>

## Next pages — disable_ha / 310131200330 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3203112302132303-1313202311322332-0222023201233202-3221021303032331-1111112120320030-1100333003113200-0133320331321111-1203121122310001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031021230103311-3232030202210332-2310332032001021-2032333332313223-1012212120233121-0013131110231212-1020001132003101-2220133122222012"></a>

## disable_log_anonymization — disable_log_anonymization / 010032311100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_log_anonymization

<a id="canonical-0103233332011021-3130012122032333-3333010121032130-3010112020123223-2103103122113023-1013312110211302-2223333230330322-1031302213200033"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

OneOf alternatives in this subsection:

- [disable_log_anonymization](resources--securemesh_site_v2--reference--group-007.md#canonical-0103233332011021-3130012122032333-3333010121032130-3010112020123223-2103103122113023-1013312110211302-2223333230330322-1031302213200033)
- [enable_log_anonymization](resources--securemesh_site_v2--reference--group-008.md#canonical-1132023020201001-3210133312102302-2003130003203131-1011133331132330-0002300031323000-3311011021302323-3122332102101131-1011013211131323)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_log_anonymization = {}
```

<a id="canonical-1231033010020131-1112123111123202-1331233001032331-2211202033302001-0213011000003021-2122230132220012-0303202311120022-3032330022030123"></a>

## Direct properties — disable_log_anonymization / 010032311100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230020122213110-2122101010102323-0102223002302132-1323211322230012-0313133100120133-1100200021003021-0310022332130231-0012213112333333"></a>

## Next pages — disable_log_anonymization / 010032311100 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3131220221232201-2023022102131221-2033023020002031-0101312232133333-3212110223303200-0212322023231221-2122222012123200-2002031123021212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303211122111320-1330022123112221-0220331123022210-3333222023120102-3000130222221310-1310221322211303-2121332210100333-2201311020013322"></a>

## disable_management_network — disable_management_network / 111100112202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_management_network

<a id="canonical-0123200302122213-3223003012220322-1003113202213003-1331312133130231-1100212321202101-0000102330231113-2230230101320011-3133303132010323"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_management\_network, enable\_management\_network; Default:
disable\_management\_network\] Configuration parameter for disable management network.

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

OneOf alternatives in this subsection:

- [disable_management_network](resources--securemesh_site_v2--reference--group-007.md#canonical-0123200302122213-3223003012220322-1003113202213003-1331312133130231-1100212321202101-0000102330231113-2230230101320011-3133303132010323)
- [enable_management_network](resources--securemesh_site_v2--reference--group-008.md#canonical-0312212202231003-0331322130101202-3220012301310333-1330003212202002-3111120113022031-3333231233121120-0000311222301201-3130230321003022)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_management_network = {}
```

<a id="canonical-2002121031220102-2332231310121310-0033021113032033-2310301222103103-0020022102211030-0331003332212131-0131223102123030-3213003313212320"></a>

## Direct properties — disable_management_network / 111100112202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122013323002122-3012202300211013-3103202132313001-1132310332313020-3131332122020303-1110333231300122-2032020210213011-2210230300001001"></a>

## Next pages — disable_management_network / 111100112202 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0302003210031020-0002223310123233-0230133013212201-1212020012311333-0112301130131323-0032203021212031-1033011301012302-3123112120202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201001212303032-2132333203311303-0211110110033222-0321323330203212-3030010332311210-3123033201112302-3022003020023223-2220000112312233"></a>

## disable_url_categorization — disable_url_categorization / 132303031333 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- disable_url_categorization

<a id="canonical-3211111131131233-0121311023002102-0103031231100113-0101230013202231-2220222302212232-3111201103133120-2012213203113232-3210330222130300"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_url\_categorization, enable\_url\_categorization; Default:
disable\_url\_categorization\] Enable this option

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

OneOf alternatives in this subsection:

- [disable_url_categorization](resources--securemesh_site_v2--reference--group-007.md#canonical-3211111131131233-0121311023002102-0103031231100113-0101230013202231-2220222302212232-3111201103133120-2012213203113232-3210330222130300)
- [enable_url_categorization](resources--securemesh_site_v2--reference--group-008.md#canonical-0322300022333021-3031321220220003-3121001123003331-3210300311220231-2113221133000010-1111302322003320-3311113331332300-0113223000130223)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_url_categorization = {}
```

<a id="canonical-1310012122123213-2100331302123312-2002033013220200-3111213220310221-2333012331003312-3121212132211323-2311033302003122-1232333101112033"></a>

## Direct properties — disable_url_categorization / 132303031333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212120231222312-2000131222030301-3112121121200001-1022211331312232-1233213122032022-3212211220022311-3212231102032100-0131023332011100"></a>

## Next pages — disable_url_categorization / 132303031333 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023213300122130-1133322203203031-2031222030032120-1122332122310111-1212211122302121-2321230221200011-0312221022333302-0232122011231313"></a>

## dns_ntp_config — dns_ntp_config / 032221023200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- dns_ntp_config

<a id="canonical-0223123033322020-3331101211232032-2003212113133320-2330131311332313-3002122310030211-3302232302212222-2020003121330012-1010330002011332"></a>

Type: `"object"`. single nested block, Optional.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_dns",
    "f5_dns_default"),
  validators.ConflictingObjectAttributes("custom_ntp",
    "f5_ntp_default")}
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
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

Terraform syntax:

```terraform
dns_ntp_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113031133301220-2232302002301313-0100230110000203-1021320033013323-2120301230001332-1033013033031010-0033111323031120-3233222310133033"></a>

## Direct properties — dns_ntp_config / 032221023200 / 3

- [custom_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031): complete subsection reference.

- [custom_ntp](resources--securemesh_site_v2--reference--group-007.md#canonical-0223311330322023-1031030023000003-2101020110023233-2222130102033223-2332103012302011-0020210112211022-1220232303021232-3032220330030110): complete subsection reference.

- [f5_dns_default](resources--securemesh_site_v2--reference--group-007.md#canonical-3013002232311000-0033131110202020-1213021032031033-1303213220200223-2332212011310211-1133223021130120-3333210012011000-2002210132030321): complete subsection reference.

- [f5_ntp_default](resources--securemesh_site_v2--reference--group-007.md#canonical-1013321102112123-3101002330233212-0213101211010332-1323311033231203-1022110213210023-1033323013110330-0000013023222121-0330331202122223): complete subsection reference.

<a id="canonical-3211122323201332-2332211202231110-3202002331121203-2001130012213302-1300202301330311-3021310022131110-3122023223200233-3131232002223000"></a>

## Next pages — dns_ntp_config / 032221023200 / 4

- [dns_ntp_config.custom_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031)
- [dns_ntp_config.custom_ntp](resources--securemesh_site_v2--reference--group-007.md#canonical-0223311330322023-1031030023000003-2101020110023233-2222130102033223-2332103012302011-0020210112211022-1220232303021232-3032220330030110)
- [dns_ntp_config.f5_dns_default](resources--securemesh_site_v2--reference--group-007.md#canonical-3013002232311000-0033131110202020-1213021032031033-1303213220200223-2332212011310211-1133223021130120-3333210012011000-2002210132030321)
- [dns_ntp_config.f5_ntp_default](resources--securemesh_site_v2--reference--group-007.md#canonical-1013321102112123-3101002330233212-0213101211010332-1323311033231203-1022110213210023-1033323013110330-0000013023222121-0330331202122223)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3012001213011312-2021230120013323-2020301303130133-2103101312332311-3022013310032001-1213000122300303-1311133100123222-3132220000230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102302121101023-3132303313323112-3303331310102102-0002120303202222-1122031232223303-1013213303121202-2112131331131101-1121230003213330"></a>

## dns_ntp_config.custom_dns — custom_dns / 322200230220 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.custom_dns

<a id="canonical-2112010003113112-1023301323001130-3121213111201030-2013130203000301-3002030020010112-0022102210132012-3222102032112100-3121210102210231"></a>

Type: `"object"`. single nested block, Optional.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

Receipt-pinned upstream constraints:

```json
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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112033312030210-1031121322103132-3133023303333113-0221333110201202-2212233023011332-2220101203123302-3022100011330132-1133221132202101"></a>

## Direct properties — custom_dns / 322200230220 / 3

<a id="canonical-0221120211322322-2123130332010213-2323313322210120-2130022031301112-0123213332000113-2200333200322030-3001310103011221-0010002132012130"></a>

<a id="canonical-2133000302031303-0000032123200220-1333322001103332-2020211100302131-1130010222203102-2221201201232233-0021020002303201-1011332032321100"></a>

## dns_servers property — custom_dns / 322200230220 / 4

Type: `["list", "string"]`. Optional.

DNS Servers. DNS Servers.

Upstream description:

DNS Servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2301021131103000-3020013120201211-1113312223021223-1100311311011001-0232010111332013-1001120002021301-2023123000210221-1210330112213102"></a>

## Next pages — custom_dns / 322200230220 / 5

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0223311330322023-1031030023000003-2101020110023233-2222130102033223-2332103012302011-0020210112211022-1220232303021232-3032220330030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113032123102311-1001321020221113-3003321210020312-3010133031221123-1031202002221112-1301112202233202-2133112003102223-3022313311203102"></a>

## dns_ntp_config.custom_ntp — custom_ntp / 010322002002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.custom_ntp

<a id="canonical-0231122333013022-1132212111001021-3313311021020221-3112100311013331-1203221122220013-0200023321212013-3333020232212123-3101320302013002"></a>

Type: `"object"`. single nested block, Optional.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

Receipt-pinned upstream constraints:

```json
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
custom_ntp {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013120211230021-0233103030121000-0020103223112232-2012223123110032-3221123232121001-3120031222022302-1032320212233201-3302230130331113"></a>

## Direct properties — custom_ntp / 010322002002 / 3

<a id="canonical-1330103310103032-3202002020101031-1010111032122310-2230033001301202-1233023202300003-3001123313300020-2331030031013022-3213002100202100"></a>

<a id="canonical-3002202100011321-1221103023232033-2310221201123232-3213213112211302-3310123112210001-3113310212320100-2023310301022131-0230310120023312"></a>

## ntp_servers property — custom_ntp / 010322002002 / 4

Type: `["list", "string"]`. Optional.

NTP Servers. NTP Servers.

Upstream description:

NTP Servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1111102000310203-3321230312132020-0000121210002210-1012310100003300-3313320133213100-3011213330111120-3211000232303032-1223021221011122"></a>

## Next pages — custom_ntp / 010322002002 / 5

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3013002232311000-0033131110202020-1213021032031033-1303213220200223-2332212011310211-1133223021130120-3333210012011000-2002210132030321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011030010012223-1001303122201312-1312311110200112-3110011133213133-3300330121233221-1101223323321221-2313032331313332-0220320022121332"></a>

## dns_ntp_config.f5_dns_default — f5_dns_default / 022322002033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.f5_dns_default

<a id="canonical-1213221331011030-3021211101023201-3230303030132122-3302202321031202-1320132120320100-0322020303230101-3101101302032101-2123033331312113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for f5 DNS default.

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
f5_dns_default = {}
```

<a id="canonical-0223012001100102-2230200212330312-2021303023322211-3303312020013230-3103022303232312-0131113133213000-1323323012122132-3313303220200210"></a>

## Direct properties — f5_dns_default / 022322002033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103213132332311-0231100301223221-2201130122201220-2010331003212003-3130311230313120-1223112120003330-3213213302310200-2222122132003203"></a>

## Next pages — f5_dns_default / 022322002033 / 4

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1013321102112123-3101002330233212-0213101211010332-1323311033231203-1022110213210023-1033323013110330-0000013023222121-0330331202122223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222022021220013-1200111311222332-0300102201002120-2203000000302132-3323313110231001-2230021233303213-3210333303132101-0100322202021012"></a>

## dns_ntp_config.f5_ntp_default — f5_ntp_default / 103110301001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- dns_ntp_config.f5_ntp_default

<a id="canonical-2300033312132033-2122000300211013-0331323213022010-3023213210321133-2223213123111111-3111122220211212-1320002132321311-3120231102130003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for f5 ntp default.

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
f5_ntp_default = {}
```

<a id="canonical-1202100232232111-3020313011323110-2322131101122020-2003023023220101-3203113333100101-2022000033131210-3211212311212011-0332130201100003"></a>

## Direct properties — f5_ntp_default / 103110301001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111033231121213-2311111323022333-1201032032013211-1213010101022030-3130020211220311-1112032201311133-2110121123223212-3032202121120130"></a>

## Next pages — f5_ntp_default / 103110301001 / 4

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-007.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002231131011001-2133302312133200-2211003220120110-0112200303001123-3103303112013202-2101020103010023-1131020112210112-0333102121331331"></a>

## eks_k8s — eks_k8s / 023112132322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- eks_k8s

<a id="canonical-3210101303003113-2321231211232111-1200012021132302-2232223222232303-0301021301220332-2322201322313001-0323302010313111-3120122330030323"></a>

Type: `"object"`. single nested block, Optional.

Kubernetes Provider Type. Kubernetes Provider Type.

Upstream description:

Kubernetes Provider Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_anti_affinity",
    "enable_anti_affinity")}
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
  "x-ves-oneof-field-anti_affinity_choice": "[\"disable_anti_affinity\",\"enable_anti_affinity\"]"
}
```

Terraform syntax:

```terraform
eks_k8s {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310203023023121-1131023011121211-3301300211121231-2332001102222121-3121313203133012-3003001233031122-1201002200001212-1221100212313003"></a>

## Direct properties — eks_k8s / 023112132322 / 3

<a id="canonical-3021030020131220-2101123101201122-0030312122113230-0233303123321302-0333032233331010-3112121032002121-1033201220330111-0233220322022301"></a>

<a id="canonical-3101111022023302-3231002112002211-1023111032201122-1110302233210203-0033002011020211-0033122323001101-1221303011230231-2221321203323102"></a>

## deployment_size property — eks_k8s / 023112132322 / 4

Type: `"string"`. Optional.

\[Enum: KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM|KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\] Enum for
Kubernetes deployment size OPTIONS - KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium Medium deployment
size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most deployments. -
KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large Large deployment size with higher resource.. Possible
values are \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`, \`KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE\`.
Defaults to \`KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM\`.

Upstream description:

Enum for Kubernetes deployment size OPTIONS

&#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_MEDIUM: Medium

Medium deployment size with moderate resource requirements (8 vCPU, 32 GB memory). Suitable for most
deployments. &#8203;- KUBERNETES\_DEPLOYMENT\_SIZE\_LARGE: Large

Large deployment size with higher resource requirements (16 vCPU, 64 GB memory) for demanding
workloads requiring additional performance and capacity.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
  "enum": [
    "KUBERNETES_DEPLOYMENT_SIZE_MEDIUM",
    "KUBERNETES_DEPLOYMENT_SIZE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020): complete subsection reference.

- [enable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010): complete subsection reference.

<a id="canonical-1233223212313102-2112232031112231-3322120010131331-0000300213221111-0332200002020200-1010021222332203-3303131300120321-3213010031033330"></a>

<a id="canonical-3030030231112131-0023002121220210-1110300330112121-0131203130033001-3132010203300210-2002221203003201-1220303321033012-3231133113302203"></a>

## labels property — eks_k8s / 023112132322 / 5

Type: `["map", "string"]`. Optional.

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Upstream description:

Add labels to control which Kubernetes nodes the VPM and related pods (etcd, VER, prometheus) are
deployed to. Specify label key-value pairs that match the labels on your Kubernetes nodes. This uses
Kubernetes nodeSelector to schedule pods only on nodes with matching labels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "253",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "63",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010): complete subsection reference.

<a id="canonical-0312122313323320-3333001221131020-2110322321211302-0022113000303000-3100333032133320-0020222022312313-3213111022002303-2012210031011202"></a>

## Next pages — eks_k8s / 023112132322 / 6

- [eks_k8s.disable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020)
- [eks_k8s.enable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2302123313002100-2233333033311201-1310100113000030-3230101310313120-2023123110023111-3310203313322012-2221021100323013-3032312333231020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121222032231003-1123013210213121-3021302012223103-1323103313021013-1312001111230030-2122123101212221-2012100101001230-1211311023133013"></a>

## eks_k8s.disable_anti_affinity — disable_anti_affinity / 230030013222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.disable_anti_affinity

<a id="canonical-2221002233013323-1013300303313310-0213001331100022-1013213321030200-2221032231101303-0031020321332120-2232202021332210-3232313130022213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable anti affinity.

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
disable_anti_affinity = {}
```

<a id="canonical-1010013022221112-2000130323333223-3320301323303110-3000301233133332-0213311002001001-3131303133012002-3011223333223032-3112330221232220"></a>

## Direct properties — disable_anti_affinity / 230030013222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311232100322013-0033121330002201-1033210100112022-2201230310310322-1301022212320131-1012103203001313-3301001113321111-3233133121020323"></a>

## Next pages — disable_anti_affinity / 230030013222 / 4

- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300330003331000-2303302221120332-0000030222000202-1302031021303131-1231211212310123-1112120231332132-0010111002211231-0103020222103223"></a>

## eks_k8s.enable_anti_affinity — enable_anti_affinity / 322211302330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.enable_anti_affinity

<a id="canonical-2110220032230201-2223313320133330-2322130200321311-1130003031202122-2123111001031302-3022030332112010-2233021022101121-2313311201220320"></a>

Type: `"object"`. single nested block, Optional.

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Upstream description:

Configuration for pod anti-affinity scheduling rules. Define multiple rules to control how different
applications/components are distributed across your Kubernetes cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
enable_anti_affinity {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001031332032120-2001303031130232-3313330113210231-0330331032331221-3202301111232311-2210212000220132-1310311001201321-1000120210113312"></a>

## Direct properties — enable_anti_affinity / 322211302330 / 3

- [rules](resources--securemesh_site_v2--reference--group-007.md#canonical-0002210012021123-3003312201220103-0201333310103102-2102301312332131-0111231332122200-0321113022232322-1133333131302323-3210232203133131): complete subsection reference.

<a id="canonical-2322112302302023-3110210011011203-0100233323131323-0020120003310302-3203223313133133-1002132230030222-0223332100330121-0123103000200130"></a>

## Next pages — enable_anti_affinity / 322211302330 / 4

- [eks_k8s.enable_anti_affinity.rules](resources--securemesh_site_v2--reference--group-007.md#canonical-0002210012021123-3003312201220103-0201333310103102-2102301312332131-0111231332122200-0321113022232322-1133333131302323-3210232203133131)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0002210012021123-3003312201220103-0201333310103102-2102301312332131-0111231332122200-0321113022232322-1133333131302323-3210232203133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121312103232000-3001311331121323-1103103012331130-3011222231311023-0121033201123320-2210023220233221-0100123211022213-2122310022103121"></a>

## eks_k8s.enable_anti_affinity.rules — rules / 230123213132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.enable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010)
- eks_k8s.enable_anti_affinity.rules

<a id="canonical-2232011021032131-0211121302022333-3032000111100332-3033223122331020-1031220220130321-1213331313121002-2122022123203031-1223122033020313"></a>

Type: `"object"`. list nested block, Optional.

Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be
distributed across which topology domains.

Upstream description:

Define one or more anti-affinity rules. Each rule specifies which pods (by labels) should be
distributed across which topology domains. Example: Rule 1 - Distribute VPM pods across nodes, Rule
2 - Distribute Prometheus pods across zones.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("label_key",
    "label_value",
    "topology_keys")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121000312031222-1002312222310203-0113131033013232-1013210203302122-2131113032111013-0300011110111032-2323331230331003-3130232202333233"></a>

## Direct properties — rules / 230123213132 / 3

<a id="canonical-3301120121212013-0202212210123121-0130011312102220-0233313230313020-3313110022032313-3003302223030333-0030032132331212-0030233121132002"></a>

<a id="canonical-3120000303100023-1231322231213123-2103010303122001-1230210222201313-0322112130130131-1100100112203312-2000100103333022-3001021002230332"></a>

## label_key property — rules / 230123213132 / 4

Type: `"string"`. Optional.

Specify the label key of the customer pods that CE pods should avoid being co-scheduled with.
Combined with the label value below, this identifies the target pods.

Upstream description:

Specify the label key of the customer pods that CE pods should avoid being co-scheduled with.
Combined with the label value below, this identifies the target pods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 253),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 253,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 253,
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
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "253",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3210013103021113-0202112131113131-2303210003300232-3321000212333013-1231020203223030-2300220123021233-1012013022233031-2323231033103332"></a>

<a id="canonical-0303210010133003-2031220112330200-2121010200031212-0310101111101011-3331313030301010-1233131223300122-0302030303000231-0023210032220221"></a>

## label_value property — rules / 230123213132 / 5

Type: `"string"`. Optional.

Specify the label value that, together with the label key, identifies the customer pods to avoid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0201211232020233-2123211230232103-2011011030303111-3131322111222302-1232223031321102-2233131322031330-3223232313013203-0001331120113101"></a>

<a id="canonical-3000000332000120-2303322000330033-1221130113010310-2321313323102331-0011032202020022-3101310002211103-1210120122203233-2210021202001211"></a>

## topology_keys property — rules / 230123213132 / 6

Type: `["list", "string"]`. Optional.

Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g.,
Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already
running a pod with the above specified label.

Upstream description:

Specify one or more node label keys that define the scope of avoidance. For each topology key (e.g.,
Kubernetes.I/O/hostname), CE pods will avoid nodes whose topology value matches a node already
running a pod with the above specified label. Example: with Kubernetes.I/O/hostname, CE pods are
kept off any node running the matching customer pod.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "253",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0230203011133301-3332313022202311-3330113013102123-0001221100213022-2002312220032022-2322331331102121-2212001010311021-1120012130321030"></a>

## Next pages — rules / 230123213132 / 7

- [eks_k8s.enable_anti_affinity](resources--securemesh_site_v2--reference--group-007.md#canonical-1313030010022021-0221332100201013-0020030200213230-0021013230101031-1131032330001021-0111013012301323-1310322112213013-0102113300201010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003131113220230-3313212011232303-3211321013012322-2121223030120211-3132021032320113-1233333230232103-3011113031103032-1201031231102200"></a>

## eks_k8s.not_managed — not_managed / 332102231201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- eks_k8s.not_managed

<a id="canonical-3130113003303110-2033010301323132-3110311003011313-1130302131021033-1223020332122222-0223032333031232-2132323111120201-1302010201330310"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102132013330000-3213231322120301-3032203200122312-1312033220311321-3123033210213131-3332201010320010-2022012022222130-2030113101033132"></a>

## Direct properties — not_managed / 332102231201 / 3

- [node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231): complete subsection reference.

<a id="canonical-2023111330221310-2002300310213333-0120120020120031-1303201330300301-1112303302120130-0013133113320203-3222322021131023-1101033200212010"></a>

## Next pages — not_managed / 332102231201 / 4

- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212222012323213-1111120103202002-3222032000123130-3003032021131031-3113102301131311-3203311130121201-0023221320102310-0033322112320310"></a>

## eks_k8s.not_managed.node_list — node_list / 132020101121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- eks_k8s.not_managed.node_list

<a id="canonical-2122122333322033-3100113021312120-0210110100033002-1333003020021012-3020030223001111-0121002331320332-1030100303213302-1113222113030110"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313103212303232-3203133111211001-1031222331101222-3230102130033222-2123012221313201-2130230112310201-1010303332012021-3221213223032022"></a>

## Direct properties — node_list / 132020101121 / 3

<a id="canonical-1013123311312332-0332021321113202-1022331102201200-0303333010302202-0211313031201221-1021011003120130-0221331332222012-1030310120202121"></a>

<a id="canonical-0321013300232110-2101333030110310-2233310003310323-0003001130022233-2323101113331213-0012132010202033-1322210322201110-2330131003233200"></a>

## hostname property — node_list / 132020101121 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210): complete subsection reference.

<a id="canonical-2333101123320133-2122103110232332-3232110303303330-0112202332322101-3301021122000131-3133030233132221-1131031112001120-3010220123210023"></a>

<a id="canonical-1102132301333211-2010222033131232-2222032330302233-0023232203203032-2020303201320132-1112012222301020-3030323013112103-2120020130020201"></a>

## public_ip property — node_list / 132020101121 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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

<a id="canonical-1013211111013203-3031113031113313-1332002303201120-2331233311113020-3322121013003110-0301023020032022-0021001021312301-0222230301020222"></a>

<a id="canonical-2213312031011210-0303011211130232-2112211122120330-0123300120212110-1230103020011302-3232110003223022-3221033021133023-3333121310132100"></a>

## type property — node_list / 132020101121 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3130031312331210-0320030303333013-1200302300021230-0203202211222223-0223320332221322-0200300210102102-3210021202212330-1201331223001210"></a>

## Next pages — node_list / 132020101121 / 7

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313100320203012-0230002221000033-3203132031302113-3330300302101103-2320213013303000-3312302001323303-0132001233212033-0020102112213320"></a>

## eks_k8s.not_managed.node_list.interface_list — interface_list / 113033110112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- eks_k8s.not_managed.node_list.interface_list

<a id="canonical-1213312202021031-3002103122300032-3303202132111103-0323331330231311-1103123011021111-0012232113202222-2302200112122320-1313232012013121"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210001101003012-0111113010300031-2221110010032122-3323103130112002-0331321311000330-2032030011120320-1011030231102333-1112322131003310"></a>

## Direct properties — interface_list / 113033110112 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331): complete subsection reference.

<a id="canonical-2231213231231020-3201113010022121-1103331331121023-2332030200332231-0113020332210120-3131003133320020-1022231113322310-1022003133132322"></a>

<a id="canonical-3222111302021222-3100010330002133-0112021311112130-1313031012331113-3222330123131001-0333212031313023-1000122110132310-3002220030101203"></a>

## description_spec property — interface_list / 113033110112 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-007.md#canonical-3012211120012130-2120220310021002-0011001021013213-1201330020123013-2321130021032122-1223300321221131-2310211131333311-0310233021221020): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-1332210212311330-2022022112011131-1010231320331201-2022030010222231-3332011213032033-0302230033003033-0112221032213112-1223000100220010): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222): complete subsection reference.

<a id="canonical-0330221100330133-2330201002202021-1020021311111200-3013201210033003-3221121101123030-3023011312210233-3321020100213102-1313032302232122"></a>

<a id="canonical-1310010003111200-3032012011203112-2032302030331033-3003132221110001-0203230322012101-0111132123122201-2312322210301031-0302313203003010"></a>

## is_management property — interface_list / 113033110112 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0101321101302313-3311010012322323-3222232321002222-1310222112033112-0323302112332202-3132000121211231-2000232130211012-0331332021311031"></a>

<a id="canonical-3231112302022120-2231012011323311-0102332111101222-3032230133103110-1122331111212113-3110110300233311-0012300133211203-3023122110230313"></a>

## is_primary property — interface_list / 113033110112 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1123103321332322-0320210301132223-3332113001103202-2012111310113013-0013111030113202-2331210011332221-1132330020131232-1300202312220202"></a>

<a id="canonical-0012123331220211-1331310002203123-2122333000300333-3232000100131221-0210300311230232-1023321010210231-3231231102121110-0201222222232312"></a>

## labels property — interface_list / 113033110112 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-008.md#canonical-0011102111120011-1022033123221032-1331321200332203-2301111122122233-3303333213003203-2021123332202300-1022321002112333-0023222221123022): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-008.md#canonical-3203001200120111-0313132010231330-3001221200201310-3122003203211002-0032332311222332-1021320120202100-3231020003010120-3230012212102130): complete subsection reference.

<a id="canonical-0301030223013230-0121223030131313-3123231100031000-0222112012031012-3220211322111231-2233210300033111-2012001321022020-3102132321102022"></a>

<a id="canonical-2333203003110132-2122030232130030-0122022112023230-0011020000203302-2110122200011313-1023221123110003-0023331310230013-3223102122023200"></a>

## mtu property — interface_list / 113033110112 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-1320213111202200-3223222232333220-3310033133012301-3200103100012113-1130332200033123-3110002222310212-2232312022122322-1221200223121303"></a>

<a id="canonical-2332030330301333-3201233300031002-1103321011310302-0031032130131231-0120121310220101-1110333220000132-2102030201313223-2303202312213300"></a>

## name property — interface_list / 113033110112 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

- [network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2301031122023323-2312030233133232-0003013033030323-3223231102220222-3121003131102022-1010213001332131-0132020321212203-3121010111120023): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-1130210113011010-0201323203331131-2111202121232233-0002232003002002-1113221321222220-3032221120013120-1301312312021220-3211333120012012): complete subsection reference.

<a id="canonical-0231332122113033-0023022123003230-1130213133212122-1020213220303111-2030323333210223-1312023322312323-3133133330012323-3133210130300330"></a>

<a id="canonical-2230330332331132-2200001010010112-2222201121020311-0003302023330232-1332110303302223-1002011103102211-1122013300203230-2032000002121312"></a>

## priority property — interface_list / 113033110112 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-008.md#canonical-0103332210102222-0101313012123113-3110223010330231-1233210031332122-1220201120130302-1030220311033000-1331310033002133-3213231231323310): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-008.md#canonical-2031202033300012-1103300021230131-2131322123021211-3111003121233331-0121331212100212-2222021220102331-2311321321112233-0313323220023112): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-3030211032102301-0013032213201203-2331222321333320-0011121300110003-0303323230202321-1230002033332221-3333131203032011-3301013013201112): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-0301032123232223-1100031103321323-0133220303233001-0223330231112002-3230302111330102-2033000033312333-2302233132112031-3122210031202101): complete subsection reference.

<a id="canonical-1201131232232002-0333111122310022-2102101211312212-0031011223322022-3031013223112031-2311213223033200-3023102323010202-2331312133323130"></a>

## Next pages — interface_list / 113033110112 / 11

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-007.md#canonical-3012211120012130-2120220310021002-0011001021013213-1201330020123013-2321130021032122-1223300321221131-2310211131333311-0310233021221020)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-1332210212311330-2022022112011131-1010231320331201-2022030010222231-3332011213032033-0302230033003033-0112221032213112-1223000100220010)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-008.md#canonical-0011102111120011-1022033123221032-1331321200332203-2301111122122233-3303333213003203-2021123332202300-1022321002112333-0023222221123022)
- [eks_k8s.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-008.md#canonical-3203001200120111-0313132010231330-3001221200201310-3122003203211002-0032332311222332-1021320120202100-3231020003010120-3230012212102130)
- [eks_k8s.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-008.md#canonical-1012031123332102-1310222131203300-2100231031013313-0112133021032111-2001121102220030-2221212132103031-3003320131021203-3320102321032133)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2301031122023323-2312030233133232-0003013033030323-3223231102220222-3121003131102022-1010213001332131-0132020321212203-3121010111120023)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-1130210113011010-0201323203331131-2111202121232233-0002232003002002-1113221321222220-3032221120013120-1301312312021220-3211333120012012)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-008.md#canonical-0103332210102222-0101313012123113-3110223010330231-1233210031332122-1220201120130302-1030220311033000-1331310033002133-3213231231323310)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-008.md#canonical-2031202033300012-1103300021230131-2131322123021211-3111003121233331-0121331212100212-2222021220102331-2311321321112233-0313323220023112)
- [eks_k8s.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-008.md#canonical-3030211032102301-0013032213201203-2331222321333320-0011121300110003-0303323230202321-1230002033332221-3333131203032011-3301013013201112)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-008.md#canonical-2122021200012213-0332230132132110-0131200032202221-0033201320313133-0022021231202330-2201100002332233-2302222300130131-3133203210310221)
- [eks_k8s.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-008.md#canonical-0301032123232223-1100031103321323-0133220303233001-0223330231112002-3230302111330102-2033000033312333-2302233132112031-3122210031202101)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013200021321120-1233320021110331-0102130001332222-1332100130203210-0332101011001023-0222033221203001-1200223330220332-1320132322222202"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface — bond_interface / 012302222102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2222203202312133-0332331232200033-3301122010223112-0300330110312113-1203210321311032-2122133113022100-0013303010131221-0312032031001011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120102213311000-3330232322103231-1003111122022032-1100021223211233-2310012013312033-0202320203123303-3111100202313022-2213311012231120"></a>

## Direct properties — bond_interface / 012302222102 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-007.md#canonical-2123232033001200-0323200020121300-2303121110330103-1333122230220202-0101123010100133-2023003022010201-3101310213003032-3110302321023122): complete subsection reference.

<a id="canonical-0001201102330303-0100012320323213-2131211333201000-1331131230212300-0020001310123220-2212230001201313-2223032130022030-2222221110201120"></a>

<a id="canonical-3201111213303110-3230231200022231-0201203023333101-0323323331023201-2100313232130301-3302013223020220-2001322130001222-3301320120033123"></a>

## devices property — bond_interface / 012302222102 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-007.md#canonical-2321301010010310-2102312313102310-1000313020113320-2121212002000110-0301221231033132-3310310130130203-2003301113001120-2133031010100130): complete subsection reference.

<a id="canonical-3130301331102120-0020311003032131-1012210121333333-1011032022221103-3013321030133221-1130233013121022-3302131202303111-1221322221102320"></a>

<a id="canonical-2120303333113303-1233212030103100-1031131013311220-2111203010312200-3213031232222022-3233102001320221-1001111221201010-2331311133221231"></a>

## link_polling_interval property — bond_interface / 012302222102 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1201211213102112-3222120120230032-2220031112111003-0013333000023110-2322030220203101-3230030012213322-1303000332030002-1313013130201010"></a>

<a id="canonical-1331103223111013-0303120112311131-1203213332033333-3331213001121310-1321101312221130-3102332232103303-0112220330021123-2232223332101220"></a>

## link_up_delay property — bond_interface / 012302222102 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-0232233100121233-1103123111232132-1311120333232332-1110111203013102-2101021030203122-0110323122203012-1001011123223110-1301333223132030"></a>

<a id="canonical-2033200220310332-0302300213320130-2332220121220220-1313120102121020-1101023001113110-0130020300101130-0200130030200121-2102132113112012"></a>

## name property — bond_interface / 012302222102 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1030021302320231-2122013012203100-0222003332203102-0020020021011220-3113021230321302-3220230312031330-1201322111000211-0133303130301320"></a>

## Next pages — bond_interface / 012302222102 / 8

- [eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-007.md#canonical-2123232033001200-0323200020121300-2303121110330103-1333122230220202-0101123010100133-2023003022010201-3101310213003032-3110302321023122)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-007.md#canonical-2321301010010310-2102312313102310-1000313020113320-2121212002000110-0301221231033132-3310310130130203-2003301113001120-2133031010100130)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2123232033001200-0323200020121300-2303121110330103-1333122230220202-0101123010100133-2023003022010201-3101310213003032-3110302321023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132010313013011-2020133120021321-1233312033202320-0233003023013111-2313031220122033-2031222203232012-1301022300113001-3212130320030031"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 333031110333 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-3220211020330112-2110011132230130-0010212301201313-0323202211103231-3013101030002330-0111203132012312-2330233110132231-0213231303211332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-2333102001003031-0030211211000102-2120231130222011-2020102131010210-2322322220213330-1002211201121101-1302033221121122-3123023220120210"></a>

## Direct properties — active_backup / 333031110333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331222232021020-1010110102333301-0211223331022001-1302022112301010-2123210200120032-3130213312023220-0120032212022131-3002101201303112"></a>

## Next pages — active_backup / 333031110333 / 4

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2321301010010310-2102312313102310-1000313020113320-2121212002000110-0301221231033132-3310310130130203-2003301113001120-2133031010100130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321203203030003-2030221103222223-0221020223322300-1111013213003332-3213133331312000-1232223230330233-1233111122012013-1231003102112010"></a>

## eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 102302222332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- eks_k8s.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1210323321310103-0023110103003103-1000100202103103-3113112021001021-0200223212230120-0120330223300132-1020103330313233-0231221232003202"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130022132211013-1201320003331333-1000010321220321-1203013012223212-0111020131220013-3120312222030331-0331012322202332-2002202000112111"></a>

## Direct properties — lacp / 102302222332 / 3

<a id="canonical-3203030330231030-1111231012023301-1120131332100002-3312301200001333-0333122022121313-0230223123101000-0200130023121013-1023330320311210"></a>

<a id="canonical-2212223302321320-0320121021111230-1311303021012232-3221211301320221-1003220312313310-0330321311001121-3022123123200320-2222313122230013"></a>

## rate property — lacp / 102302222332 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1001211033012320-3101033131030200-3132213320303000-0301103302311220-3203012030200313-3212102232030100-0122231203223201-0011303030321203"></a>

## Next pages — lacp / 102302222332 / 5

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-007.md#canonical-3212321310022010-2311313233011120-0211230212110331-3013121331230112-3103201022013300-2110131302223120-3203330202013103-2120130211131331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3012211120012130-2120220310021002-0011001021013213-1201330020123013-2321130021032122-1223300321221131-2310211131333311-0310233021221020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313232103200202-0131003020011133-1323332211032320-2302200002323122-0200022101300032-0202220131310212-3220133002320013-3133131303103322"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 231012102201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-1330233210022311-1132132111111112-3010210332131033-1003222332320003-0311210330321002-3103222320230120-1111021002321200-0110311132103120"></a>

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
dhcp_client = {}
```

<a id="canonical-2102020120131201-2130010003330113-1110310122303133-0220013330333213-2110112010003210-2322132103112113-3320310110213102-1230230321330200"></a>

## Direct properties — dhcp_client / 231012102201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220221122131303-2202123123021112-1102100222301133-0333130201131133-1100221211130123-3033112122222330-0231301110201201-2322111112333022"></a>

## Next pages — dhcp_client / 231012102201 / 4

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012333102200032-2231302213322113-1121122103211202-3010100313022131-0311103011333033-2213301110002030-3101111123022102-0101210000232011"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 110111100131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-3310202212001323-3302310311001201-3322102003122011-1222102233202300-0101100122023332-3112201203120013-0311132233013233-1020211310132112"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321011012201021-3311030213003011-1302203320031231-0331321000033210-0210201231223002-1110020101010323-0211211132123103-2333102203031111"></a>

## Direct properties — dhcp_server / 110111100131 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-007.md#canonical-0220100333223003-2111303331103023-0111313021333210-0212302311320110-3200102011323320-3313221010202202-0213311230130320-3313200121322030): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-007.md#canonical-2132110321321323-2332002331132213-3300210233310101-0023012003303313-1032311202200131-0032131110030211-1233013233233303-0212120112232020): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002): complete subsection reference.

<a id="canonical-1130211101231002-0301330333223133-0310300230320231-1333012222203323-0322130220223101-2032120310223232-3102212130133032-3210023101311002"></a>

<a id="canonical-2232101110102032-0030102333020210-1021320001130222-3023112300213210-1100013220303032-0003102233020013-1211123002230103-3002212001110103"></a>

## dhcp_option82_tag property — dhcp_server / 110111100131 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-1122030221132223-2031112012130213-1211032220222313-2013112320102121-3101113010201002-2320102332213010-0121303233202013-1211102131002131"></a>

<a id="canonical-3200032101030122-2033220111113010-0030102210020123-1221111120233220-0110330013031211-0310133221033011-3302311203120001-1032201233221200"></a>

## fixed_ip_map property — dhcp_server / 110111100131 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-007.md#canonical-3303302301230200-0011330321312322-1022200030121323-1132113323220100-3121232202222222-2233333120113121-3000322210011300-0330300230101301): complete subsection reference.

<a id="canonical-1312233301130101-1323320132223323-3211132213220330-3330010200211100-1133232000220110-0203001032213213-0331230132302002-0101233221131121"></a>

## Next pages — dhcp_server / 110111100131 / 6

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-007.md#canonical-0220100333223003-2111303331103023-0111313021333210-0212302311320110-3200102011323320-3313221010202202-0213311230130320-3313200121322030)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-007.md#canonical-2132110321321323-2332002331132213-3300210233310101-0023012003303313-1032311202200131-0032131110030211-1233013233233303-0212120112232020)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-007.md#canonical-3303302301230200-0011330321312322-1022200030121323-1132113323220100-3121232202222222-2233333120113121-3000322210011300-0330300230101301)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0220100333223003-2111303331103023-0111313021333210-0212302311320110-3200102011323320-3313221010202202-0213311230130320-3313200121322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121010303132102-3220130301212010-2222220220120033-2113313332320022-3222302121321301-3333302232100000-0113102113020030-1220323132013320"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 320013320002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1020101330310103-1100302302201002-3302012010201322-0323003231302303-0101321032310102-1213123113303003-1020300311101031-2310220323120023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-1023211330233323-3210310121210223-2302033002212210-2113201001113220-2203000321213202-1312232302232313-2031303222303012-1023001320110332"></a>

## Direct properties — automatic_from_end / 320013320002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223203330021211-2123200210233000-0222220202303030-1223110203100203-1133301322213220-0103121231033111-3020112323301232-0323030330120010"></a>

## Next pages — automatic_from_end / 320013320002 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2132110321321323-2332002331132213-3300210233310101-0023012003303313-1032311202200131-0032131110030211-1233013233233303-0212120112232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102313303131120-2113233132222130-2330323310121122-3312010302131023-0223101212320133-1321031101332333-0311313330233032-3212002201231202"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 313132323200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-1300102333323221-2031332312232212-0300202021113132-0201221022210311-3020220220013301-2112023011230311-2110010011101213-3320213333301201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-0311033110112313-1333011200113201-1223321222323310-3201132330021231-1321011332212103-2112112222110332-2331121131133223-2122202132203213"></a>

## Direct properties — automatic_from_start / 313132323200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301032232312320-2021220123312101-2113112103113200-3012001030313200-1002130101302103-1232110031130331-1023100332302130-0230321322330020"></a>

## Next pages — automatic_from_start / 313132323200 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333031323312311-2223202322201203-1231321323213100-1213122112110022-0132003232113032-3233010012123111-2121112133322012-3302013131133112"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 313223132300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2321101212212010-0212301333302333-1332133231203300-0331232022023210-1003230101003223-1332021311102110-0331121200103201-3312311330232020"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320332112112213-2232313111122133-1230323222030120-2002122103322000-1210210102320133-2310102131023112-2300230333011010-3303200030320301"></a>

## Direct properties — dhcp_networks / 313223132300 / 3

<a id="canonical-3300033311201211-2013110312120131-3211331212103222-3103231233120202-2200001113332203-1130132331101033-3113123221101121-1201033030321333"></a>

<a id="canonical-2023030223310110-1200231222301023-3131220000302302-0330133030012323-1331000010130313-3330223021230332-1030013330001220-2213210312002330"></a>

## dgw_address property — dhcp_networks / 313223132300 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3312102012211201-1120022202132132-3320102112202210-1023230233301300-3302031110203003-1133221231000200-1322112321233031-1311302032102301"></a>

<a id="canonical-0013122311323322-3201331130023203-3210202131112322-3133201020202122-2202000311031031-2002032101221002-0003222300103020-1121310013122131"></a>

## dns_address property — dhcp_networks / 313223132300 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-2030103303132012-2013131220223312-3303133321333100-1203121131211200-2030130000113211-0320010212133000-3201032210323020-2012102032021001): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0021303331303301-3031221212331311-1100310330232131-1303312212002221-1020110333313012-3102031331122112-3111113310000103-3133313132221002): complete subsection reference.

<a id="canonical-3010010223021030-0331302130021131-0131110023120323-0120320323203222-3323312310033102-2030300032233002-0300021222230200-3332201103213020"></a>

<a id="canonical-2021201112130003-0330231032211332-2213100300213323-3110200023222221-2013011313220120-2113020123011221-2103223330001232-0321021232221310"></a>

## network_prefix property — dhcp_networks / 313223132300 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-1120100023302103-0031021233232222-3332300333130210-3223011002000103-2030002121120023-1123133100032120-3103330113130110-2023303222222322"></a>

<a id="canonical-0113102132100031-0322232132003322-3211300333321302-1011220322311003-0211023122212103-3113310311310310-1003011033100021-3303020021311122"></a>

## pool_settings property — dhcp_networks / 313223132300 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-007.md#canonical-2223322130221332-3301201323222230-3221311211100132-2013033103333321-1103010331130331-1310211301002312-3333232223112233-3023020132031121): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-007.md#canonical-0132200210332303-1211133112023112-2300232030020112-1231130312023013-2300100032223203-2020203223000223-0000232123222003-0331210123232222): complete subsection reference.

<a id="canonical-2102122333302321-2111000003233321-3121202200210113-2201311102323113-2330230201100213-0103130233321220-0221321131323002-1203200013120101"></a>

## Next pages — dhcp_networks / 313223132300 / 8

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-2030103303132012-2013131220223312-3303133321333100-1203121131211200-2030130000113211-0320010212133000-3201032210323020-2012102032021001)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0021303331303301-3031221212331311-1100310330232131-1303312212002221-1020110333313012-3102031331122112-3111113310000103-3133313132221002)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-007.md#canonical-2223322130221332-3301201323222230-3221311211100132-2013033103333321-1103010331130331-1310211301002312-3333232223112233-3023020132031121)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-007.md#canonical-0132200210332303-1211133112023112-2300232030020112-1231130312023013-2300100032223203-2020203223000223-0000232123222003-0331210123232222)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2030103303132012-2013131220223312-3303133321333100-1203121131211200-2030130000113211-0320010212133000-3201032210323020-2012102032021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132032032222020-1103212112112133-2020102323221131-2021223122030000-0302030333310100-2000111021221230-3003223333030131-2120321222220100"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 311230120202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1001030323120133-0133011102000233-2231203030112211-0212223202111301-3031333131020200-1103313022220121-1102202033322202-0113211301032330"></a>

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
first_address = {}
```

<a id="canonical-2023013021323201-2301030220232022-3203200301013201-0010212313312230-3320031231221210-2033321331310312-0303011121130231-0201110311101231"></a>

## Direct properties — first_address / 311230120202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012013021211110-1320323233321000-2313221121221221-1132233123313203-3212221031130013-1032121133130111-0320100222211303-2101322110120031"></a>

## Next pages — first_address / 311230120202 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0021303331303301-3031221212331311-1100310330232131-1303312212002221-1020110333313012-3102031331122112-3111113310000103-3133313132221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130133233231211-0112211313002112-0202021323233331-1221321201013331-1011101122120212-3020300032112120-3312001212300022-3120200122003302"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 310312302330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1310000001101032-3130023311320230-0132301000023230-0301030203221202-1002000113021130-3122331322320132-3010211111113020-2212222200113320"></a>

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
last_address = {}
```

<a id="canonical-0313031000100310-2301022320133101-1022311111101322-2100321131122233-2133231212023021-0112032120312102-1123331301330021-2101101112103011"></a>

## Direct properties — last_address / 310312302330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020212222331302-1220202310212223-1023031233110001-2121220200001133-2210110012021111-3301013010101223-1223111311320121-3120331310121032"></a>

## Next pages — last_address / 310312302330 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2223322130221332-3301201323222230-3221311211100132-2013033103333321-1103010331130331-1310211301002312-3333232223112233-3023020132031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201011200002320-3332232312230023-3312000302301131-0221121020003013-1230221213123113-1323031121121132-1223132120302133-0011120130331203"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 223001330033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2203223133310332-0313321233310021-1121021212330000-3233232200020233-0101010212332132-2131321321110222-2021000221300313-2013011213003231"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330303110130011-3113310131320330-2120121212120023-3302212132221110-3230020221331022-3312023101112100-2001012131312100-0020313210200112"></a>

## Direct properties — pools / 223001330033 / 3

<a id="canonical-0021200131101133-3022013231032010-3212011320233000-1233103002311310-2321303130212311-1231212321021221-1132030102133102-3033322322223232"></a>

<a id="canonical-1230322203013101-2022121203203332-1211230031332300-3112003100331120-2302230223032003-2210123133313230-1031201013133011-3020123312311201"></a>

## end_ip property — pools / 223001330033 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1002101232032003-1332021002012102-3020132133213100-0121023202022301-3133200131322200-2110210112223113-3023001101300201-3223030001110022"></a>

<a id="canonical-2213210130112200-2311210030121101-3100321011110231-0022323100001233-0021002203211333-2101302331112232-3222212020001130-3131333331201023"></a>

## exclude property — pools / 223001330033 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-0131001300212220-0223330212322301-1322221131020012-2122100111000230-2000303302120021-1021112301120031-3303333100323232-0332022012022213"></a>

<a id="canonical-3132231013330001-1032200312231302-2231131112333030-1102222202121303-1102233001022200-2222310110302202-2131123001000310-1012110112111232"></a>

## start_ip property — pools / 223001330033 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3231210212122202-0312311322022001-0130133230110233-2111012303223000-2001221133232302-3123331102230010-2221102132320101-0231200021100102"></a>

## Next pages — pools / 223001330033 / 7

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132200210332303-1211133112023112-2300232030020112-1231130312023013-2300100032223203-2020203223000223-0000232123222003-0331210123232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302130230133213-0000031010210312-0311332122223221-1101220023101230-3003310112330131-1010330202133231-2102223012103200-3111301322032203"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 313332120203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-0111023323231233-1031312030320130-0001311013032110-3120102232132313-0103011223333220-1102121021323103-0233121120331132-1112002113010011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-1332132312111332-0012000110332332-1120020000200233-0101220202011201-2212212323110330-1110102223220321-3121332023201322-2333021330100033"></a>

## Direct properties — same_as_dgw / 313332120203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320232332121213-3003023333301232-3030123001111023-3210020313333132-1202130223332300-1323213130301333-1011312313202022-2330101032212212"></a>

## Next pages — same_as_dgw / 313332120203 / 4

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-007.md#canonical-2201333032221120-2121130202331100-0001033200333212-0322101311321223-0113211133111002-0131122203110011-1020220111320321-3322230203002002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3303302301230200-0011330321312322-1022200030121323-1132113323220100-3121232202222222-2233333120113121-3000322210011300-0330300230101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112302313003110-1113221133322110-3020122023232133-3003130203100323-1110113201230103-1330310132223033-1100321221312322-3021001212213013"></a>

## eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 303212123113 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- eks_k8s.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-1132130013121312-3313320310022310-3302313111121213-0111321303011100-1333111212112230-2301220022202022-1001003233013333-3330022333201213"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

Receipt-pinned upstream constraints:

```json
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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300102030113231-3122201131013111-0101102212012230-1021302030020202-1130023322210210-3112300330311203-2302232332021203-1323211132323221"></a>

## Direct properties — interface_ip_map / 303212123113 / 3

<a id="canonical-3032232333223021-2003002000120002-0131102003331211-1122211230121323-3032312133221300-3202031033133001-3222330233100320-3233011321333230"></a>

<a id="canonical-3133002110212133-3313112012120002-2023132200130222-3010203013223102-3302300101321121-3301022331103332-1021203333231102-1212122301112223"></a>

## interface_ip_map property — interface_ip_map / 303212123113 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-0023123131212232-2310233233221122-1003233031203013-0210311300201311-0223103120201231-3031213331021310-0332032130123020-2330110230211103"></a>

## Next pages — interface_ip_map / 303212123113 / 5

- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-007.md#canonical-1320102013013311-3031131230311233-2113002322022003-1222111220232233-3321222103011230-2210022232031002-2203020012303101-1120303111033202)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1332210212311330-2022022112011131-1010231320331201-2022030010222231-3332011213032033-0302230033003033-0112221032213112-1223000100220010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301131100221230-1002022331230031-2123300222320303-3100213211322222-3312212120300012-3230003223203132-1300123223330121-1301212313203313"></a>

## eks_k8s.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 022332301313 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0130110330021022-2223312311132201-3111323210330032-1220002213013301-0133123032122201-3133310013312222-1322032132103022-3000022232022323"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211123230202231-2122013032132101-3200212100031310-3221220201001312-3011103220213033-2320323233020131-1232012021020210-2320302210321222"></a>

## Direct properties — ethernet_interface / 022332301313 / 3

<a id="canonical-1232012212122011-1203003131301320-2331323131331003-3030220011113320-0113303321201203-2033200200030300-1222012312322032-0213311230111103"></a>

<a id="canonical-3013311102131011-3100120023331313-2102303030331010-2202122022221022-3011001110113112-3330231233010133-0230122333032313-1201100302131213"></a>

## device property — ethernet_interface / 022332301313 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1131020101222301-2233230231331332-2120033012203033-0002023312333101-3131333213123313-2223013013332121-2012221200213333-3220133012222102"></a>

<a id="canonical-2232302112132020-1331233310131310-3120113131110222-0210221301102211-1210232213133302-0003333221110222-0003230222232201-0113112230110203"></a>

## mac property — ethernet_interface / 022332301313 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-2233110203110322-0210331031011101-0223302100202022-2221320311210320-2132201033300201-1112121223110231-1100131023202023-1031112100202113"></a>

## Next pages — ethernet_interface / 022332301313 / 6

- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233022132121112-0311322322322101-2311313121011311-2132033300231130-0221012011232032-1023203130122310-3213010302321112-2032132033201311"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 132311301103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0323323312231303-3112111130123133-1123221310031112-2132222032030101-3301303110330133-2002120312303033-3022303202210000-1320130212323300"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123023033101322-3223331102012201-2000101021323312-3233100233103111-1232310230223103-0133222013100102-1000013322102321-0323222320231110"></a>

## Direct properties — ipv6_auto_config / 132311301103 / 3

- [host](resources--securemesh_site_v2--reference--group-007.md#canonical-3210110102113333-1210220300312303-2102123210011322-3211201311123203-3013323200023210-2000130102311123-0213000301220301-1230221231101130): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213): complete subsection reference.

<a id="canonical-3302111321331200-0200320210212001-0012113311332220-3201313320001030-0332301032133133-3033103313103301-1002102301322121-2113100210232301"></a>

## Next pages — ipv6_auto_config / 132311301103 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-007.md#canonical-3210110102113333-1210220300312303-2102123210011322-3211201311123203-3013323200023210-2000130102311123-0213000301220301-1230221231101130)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3210110102113333-1210220300312303-2102123210011322-3211201311123203-3013323200023210-2000130102311123-0213000301220301-1230221231101130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202120132011132-0020301302110122-2033023220321123-0002313103010300-2031113110132310-0000031021211003-1020101220221110-3031030323110302"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 103310230003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2001223103233312-0332131201322031-0320323233000031-2211100033132232-2213022010231102-3022113311130200-2003222011330322-0123313120220032"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-2232010100210112-3312130130111330-1200310320122321-0233100000233200-3131302003033312-2032002333200300-2213102201032003-1012131321103310"></a>

## Direct properties — host / 103310230003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322102013320121-3312101110110231-3203231030133012-2123123113002323-3210200103202310-0103022130320201-1210321033030203-2030111310200030"></a>

## Next pages — host / 103310230003 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200311230212100-0303202323322033-1131230230230022-3120020220233221-3012031020322313-2101323111102011-0231231203033110-1013011233323303"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 121221220102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2201312203012123-1331333003133111-2120220033200323-2330312102210002-1322113331200230-2231331020000202-3231022001121202-3132220130003222"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131323132320303-3203312130311002-3120120113011021-0010020333021032-2330201320122330-0011220031011002-2220332131020100-2010230132310200"></a>

## Direct properties — router / 121221220102 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132): complete subsection reference.

<a id="canonical-0023112320201230-2233013333033101-3232132032002110-1232022320030220-3333111123121212-2200010031032003-0111132321121333-2012033130202013"></a>

<a id="canonical-2203010310013033-1322303312312010-1322203121003011-2133031332300320-1211133111333302-1002112301232310-0222010102022100-0110110012002022"></a>

## network_prefix property — router / 121221220102 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333): complete subsection reference.

<a id="canonical-3332032110211003-1322112311030331-3331122031302200-1132131211322021-0203200010102001-2310031020300113-2311323232201012-2320003300112103"></a>

## Next pages — router / 121221220102 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-007.md#canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020323223213000-2133001211301303-3110300200002001-2112020211002233-1231102101200333-0001323203301311-3210000132213320-2201100130220330"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 310111233031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3000222303312220-2021333323121002-0013223010210013-0000203012031301-1032030322301000-1023223023022112-2212010331231302-3311212013032103"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110031333232300-2221321131013031-0310301333033031-0010311023011331-3032030202113030-2202013113013021-1211312101201021-2023100101012013"></a>

## Direct properties — dns_config / 310111233031 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-007.md#canonical-0301133201300303-0002301023011110-3300133332313333-0021101333021112-0311201222203131-0032301032123223-3302112000333120-1001322302013211): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113): complete subsection reference.

<a id="canonical-2133020200023022-0313213332021231-0220033323033121-2333232130031220-0220030230310311-1002031331011323-1223003000120021-0002223202132031"></a>

## Next pages — dns_config / 310111233031 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-007.md#canonical-0301133201300303-0002301023011110-3300133332313333-0021101333021112-0311201222203131-0032301032123223-3302112000333120-1001322302013211)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0301133201300303-0002301023011110-3300133332313333-0021101333021112-0311201222203131-0032301032123223-3302112000333120-1001322302013211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113201000121113-0001110112221232-3013030003020230-3000013001223000-0131013012033321-0112033101322323-0121101032113333-2100102211323100"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 023332200132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1230030202311210-1013122031230322-2301033232033330-2220230011132202-1221130301233032-1102113121013232-2312110031331023-1023312222123123"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303032210313320-2003111330131113-1320232000022112-0100323231010112-2011330213130302-3333032323301112-3331032013013000-2030322023012212"></a>

## Direct properties — configured_list / 023332200132 / 3

<a id="canonical-1313201313331020-3102022210220322-2101303002220311-3001120302112323-0011100113002112-1111123302222223-2212203010213133-3233011031311100"></a>

<a id="canonical-2230101123331223-3020221012220102-1112122130211113-0213210013020222-2320012100212123-0002201111013030-2300222200332220-0113022113001023"></a>

## dns_list property — configured_list / 023332200132 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0122032012023022-2311030112211230-2111132233020203-0101313003220130-2121031232312010-3133223100001233-2112332323112222-2032002331211112"></a>

## Next pages — configured_list / 023332200132 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213323301112123-2223012312220311-2321330013302111-1001311001301300-2321203120112120-3203201112103120-2201102231032133-1113112310320201"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 123221000230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2303031310223303-2021102133020203-0112303202033210-1321012122101112-3322113010332003-1302202031112310-1323200132220221-2230000102113210"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322133131323200-0123313113133223-0303213213110200-0132033022232132-0111003020221002-2211000110300323-1203130201321310-1002232313232030"></a>

## Direct properties — local_dns / 123221000230 / 3

<a id="canonical-2201312001211202-0132322001223130-3012333031123213-1133320330002203-3132202213100233-2013031030001022-0121202212312300-1012131312013031"></a>

<a id="canonical-3210200000322333-0300123031313111-3300311223112311-2313011102010313-1002120301202320-1211021031032233-3131010232121101-3032213331121123"></a>

## configured_address property — local_dns / 123221000230 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0131022310030131-3333020223333011-1230003233111300-3031222033302211-1333133132132301-2133201232130223-3232221221021313-1123010122311332): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-3121131031322320-0033222122012233-1233012103023202-0001131112000212-2212222210103130-3010022222203320-1002303321232003-1301012211110221): complete subsection reference.

<a id="canonical-0231202320011132-0303200321210232-1003000102313023-1113233330210312-1213131301201112-3100201101120302-2003212223231232-1002203123123200"></a>

## Next pages — local_dns / 123221000230 / 5

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-007.md#canonical-0131022310030131-3333020223333011-1230003233111300-3031222033302211-1333133132132301-2133201232130223-3232221221021313-1123010122311332)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-007.md#canonical-3121131031322320-0033222122012233-1233012103023202-0001131112000212-2212222210103130-3010022222203320-1002303321232003-1301012211110221)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0131022310030131-3333020223333011-1230003233111300-3031222033302211-1333133132132301-2133201232130223-3232221221021313-1123010122311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312201233212120-0002233313110320-1103121131113021-1002313121202100-1233300330330232-3213222202122012-2100122212032313-3301301111022001"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 332003112032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1231001132102110-1003310130010213-1313202010132021-0100123222202202-3233021210021320-3013121130313033-2233300010320331-3331112013122030"></a>

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
first_address = {}
```

<a id="canonical-3302301032212223-0113132332322132-1113301210311001-2301122222110302-3012202111210003-1101002110310330-2123233031332120-2032323123023030"></a>

## Direct properties — first_address / 332003112032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030232230302220-2111020020200323-2333211110321313-3103233233332200-0332312133222310-3213232311311011-3031320233230122-1002111123132211"></a>

## Next pages — first_address / 332003112032 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3121131031322320-0033222122012233-1233012103023202-0001131112000212-2212222210103130-3010022222203320-1002303321232003-1301012211110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322301013130332-3302002320330222-1000113312333100-0201112121210032-2002020332131320-0223301332202113-2003113302103220-2310120133320323"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 003230201011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-007.md#canonical-3101203112302012-0101233311220020-2101330333131010-3322013120030133-2311313312323111-1030030023103101-0013030030011223-2230020020103132)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-1212110222100113-2023233133211302-1221000311202233-3001211101110232-0333211200223212-1222111122331202-0232020101322323-3003121301221031"></a>

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
last_address = {}
```

<a id="canonical-0303302023133012-3010033031331013-2220031132020322-0233202100321202-3310002231222021-3332323310020230-3130120300000303-3003303332123200"></a>

## Direct properties — last_address / 003230201011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232211233303003-1032332300111321-3300321303333113-3012211103111001-1101223102022010-3322211020002333-3023310201020101-0032020010312010"></a>

## Next pages — last_address / 003230201011 / 4

- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-007.md#canonical-2103023300310122-0020130221210301-2202201333333003-3002330213233132-2003312311023012-0222030200012310-3013011312333312-3100002303131113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1221331232200310-3211220011101133-3311230011000332-3122031332220323-2133302330031331-2313202323211303-3220201230013000-1121032230200333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202100102310313-1223333300102200-0230002032133221-3230330013333122-3201032103220310-1300021330233301-3103323120123320-0111013210300202"></a>

## eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 312323223223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [eks_k8s](resources--securemesh_site_v2--reference--group-007.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100)
- [eks_k8s.not_managed](resources--securemesh_site_v2--reference--group-007.md#canonical-1110112122310000-1333332210131211-0233031011331213-0020022112032211-3102103310003112-0302021001133131-2202131332021122-0222311013020010)
- [eks_k8s.not_managed.node_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2121032313322223-3333000313313333-3110022111111033-0101201201203110-1101302202200200-3330001130302311-0011302121022331-2122121230101231)
- [eks_k8s.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-007.md#canonical-2220203201033202-1121022310000310-3002113303130320-3120122021013313-0003130111303222-0330101230013032-1302121101030103-2013232013323210)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-007.md#canonical-2200333202012013-1032020123110312-1103003022101211-3120230223332300-1330200101333321-0322311021020320-2101112231223302-3130303130132222)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-007.md#canonical-1133022021222230-1131112113010300-1203023211101211-2200113222223331-2210233123121111-2302101220203100-2010002301320020-3330011013303213)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3312101300231202-2100001031211330-0123103002333311-3011312101332111-3132111232030033-2110112230330320-0323331321310130-0133120112121210"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122022323320222-1311303330303111-1030102100103210-2231310233333110-2313231213322031-2132302011322313-2231030210203303-0103213121212213"></a>

## Direct properties — stateful / 312323223223 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-008.md#canonical-1310223221212020-2203212021200230-0221303023101330-1301310133310311-3032121213033212-1202203201003011-3220012231311123-1010330330003120): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-008.md#canonical-2211100321120010-2330031313123202-0132303113010022-1320222311120030-0033133130012002-2302131000322013-1300321113021231-1030002012203322): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-0333212222022210-3220201220203020-2311212122313333-0211213312102121-1112301232200330-0113200130212002-1012121112033333-2101323302230330): complete subsection reference.

<a id="canonical-3301100301232221-3220000311010211-0010120231203111-0023210203223221-0033102012101210-2202123302010122-0031013001330312-1320233120002021"></a>
