---
page_title: "xcsh_virtual_host reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_virtual_host reference."
---

# xcsh_virtual_host reference

<a id="canonical-0333230131223232-0200103011030200-3131113332320211-0030232211301021-2322030132223221-2201221201221312-3212123000230213-3003300323033303"></a>

## response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 133003133031 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_cookies_to_add](resources--virtual_host--reference--group-002.md#canonical-1032133033112132-0100100123122223-2232322330210321-3113210320300200-3230220212110103-1233220122033110-2010201013002320-1023011120330313)
- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130)
- response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2201120200032020-3231311213302112-1301111132002220-0202333213133220-0021121202211221-3030301001202113-1022222200202312-3003001322031221"></a>

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

<a id="canonical-1012033032130033-2000020030000202-2030210212223012-3021121303213131-0021023102133233-3310003033321011-1122320132112022-2032130202313323"></a>

## Direct properties — clear_secret_info / 133003133031 / 3

<a id="canonical-0320023110210030-0132120000023010-2003233200032232-1033033333102032-2302232311011110-2201231232310331-0220033101211311-1023031221330303"></a>

<a id="canonical-2102000331323031-0321133320003111-0321113103022020-2132001000220220-3101331201113100-2121000000222030-1213301031032023-1131310021201130"></a>

## provider_ref property — clear_secret_info / 133003133031 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0122102121020011-2032120130332333-2200123132330033-0220030130232103-0130313312010111-3220231110302223-3331313222023220-1130131333330003"></a>

<a id="canonical-2323002223201011-2312013221023113-1202001330032313-2123111030231302-3231310332000300-2330003032030131-0102222111023000-2112022100331011"></a>

## URL property — clear_secret_info / 133003133031 / 5

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

<a id="canonical-1112001101132322-0002012201102201-2101210103121130-3313303121303311-0202023013021101-3023332311012311-0221123100232223-1022213202231122"></a>

## Next pages — clear_secret_info / 133003133031 / 6

- [response_cookies_to_add.secret_value](resources--virtual_host--reference--group-002.md#canonical-1010303131122313-3322223020312220-2031002021323312-2321220101102002-2222302113021322-2120000333313022-0320332221303010-0110313302022130)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201220201322331-0103302213323313-1211221122220231-1223323022130130-2200202010331133-2221303032300110-1201230101032033-3021212030120223"></a>

## response_headers_to_add — response_headers_to_add / 201330200200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- response_headers_to_add

<a id="canonical-1332201100013111-2211122001031301-3021122213021002-3103112322011321-1203121332113223-1233013010332233-2313001300333001-1313003003202120"></a>

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
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313223110303201-0200323132031021-1013013132113023-3032201121210210-0032201222003032-2210312213031231-3033322330320112-0310023010200000"></a>

## Direct properties — response_headers_to_add / 201330200200 / 3

<a id="canonical-1002223013013103-2102120223002030-1132001311122000-0310111231212011-0030132021203103-1200302310031220-3320213123333230-2020120032130001"></a>

<a id="canonical-0221311332133221-3302031302010312-0121232211213313-2301330312332213-3201012303131330-1211220313103200-2332103031131131-3220310101033013"></a>

## append property — response_headers_to_add / 201330200200 / 4

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

<a id="canonical-1110020330011320-0301111023102030-0003130000010223-2001003023103133-0000032033331303-1132211333120122-2310221231211001-3013002233202111"></a>

<a id="canonical-3001011130123023-3100122313122102-1132233133323232-1021012020220113-2323210111101202-2011021122032011-3133331330330233-2313323000322231"></a>

## name property — response_headers_to_add / 201330200200 / 5

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

- [secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312): complete subsection reference.

<a id="canonical-1020202333133332-0012231300003001-1303210322300100-2030022011210310-2101011223100332-0321110122110112-0100302200222313-2100011130210121"></a>

<a id="canonical-1102110232110112-1300031123130113-0003212322031303-1002031203330233-0230112222122100-2201310212333032-3220200302023201-1111310111310201"></a>

## value property — response_headers_to_add / 201330200200 / 6

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

<a id="canonical-0321302023220332-1320333212102302-0020221103330021-3311300203233131-3000333112200233-0023000312333023-3031123330313102-1111112122303300"></a>

## Next pages — response_headers_to_add / 201330200200 / 7

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002223211213130-1321032302330003-1200220310001300-1003002313110333-1233230303331020-1100211112230022-3122003311113221-1022221110312231"></a>

## response_headers_to_add.secret_value — secret_value / 032210321111 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-003.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- response_headers_to_add.secret_value

<a id="canonical-1022221223111302-1022022232321022-1120002030211230-1101301311000311-3111213102210232-2132233222132130-1231321212320312-3133213103111301"></a>

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

<a id="canonical-3232330021202001-0320012030110200-3110022132210223-3102033110131311-2130330020020111-0023200322131222-1233303021030021-0000103101331031"></a>

## Direct properties — secret_value / 032210321111 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-1233013201222331-3030133300002211-3201230110320221-1000211222112222-0232333220101033-1322020211330332-0210031213033123-1022332322120233): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-1210303002202023-1322031021321021-3331232302132303-3201130332331023-0303111303030100-0213221021012223-1222230233321230-2123233200312131): complete subsection reference.

<a id="canonical-1122223102223223-2321321233030032-2021232101230311-2011012223203020-2001300312211311-3232030120210202-0220212122130100-0110233301110220"></a>

## Next pages — secret_value / 032210321111 / 4

- [response_headers_to_add.secret_value.blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-1233013201222331-3030133300002211-3201230110320221-1000211222112222-0232333220101033-1322020211330332-0210031213033123-1022332322120233)
- [response_headers_to_add.secret_value.clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-1210303002202023-1322031021321021-3331232302132303-3201130332331023-0303111303030100-0213221021012223-1222230233321230-2123233200312131)
- [response_headers_to_add](resources--virtual_host--reference--group-003.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1233013201222331-3030133300002211-3201230110320221-1000211222112222-0232333220101033-1322020211330332-0210031213033123-1022332322120233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011131223000110-1122333311031200-2100123322332323-0112301312310312-0130122023122033-1123002102221102-1022012303321313-1101332123022212"></a>

## response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 320233311300 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-003.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0130000101120133-3330110310003133-1030200231211102-0332201203201230-3102013302233322-1212300113012200-2213203321130132-2232031310101020"></a>

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

<a id="canonical-0021211222122020-3233002021320231-1222200113331301-2101102032322102-2300120323111323-3323120311201111-3101130321112321-3302230101122001"></a>

## Direct properties — blindfold_secret_info / 320233311300 / 3

<a id="canonical-1111111033210122-1101122203100302-0221001130032300-3102022022333022-2330121333133333-2232210230332203-2113123220013313-3313222123232310"></a>

<a id="canonical-1213031000221201-2111122211222003-2001122020023302-0200323130131320-1223302231100132-0331112203211303-0130013202310000-2203132113220032"></a>

## decryption_provider property — blindfold_secret_info / 320233311300 / 4

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

<a id="canonical-1111130000031302-0133021123222020-2320332132300230-3313201201030202-2221201333312132-2323012303313133-3133323121221232-1232032332310330"></a>

<a id="canonical-0002202232121010-3012012001131133-1100203020100223-0213010210203230-1210112223212300-2033102320132103-1120230130021111-3211030300121102"></a>

## location property — blindfold_secret_info / 320233311300 / 5

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

<a id="canonical-0022001000331301-1333032031222220-3320311321313032-0030313223332031-3003300301330300-1103101031131120-1132111233220002-0112201201312321"></a>

<a id="canonical-1221121111310332-0130030220100232-2111100223330111-3100112013110101-2133013300133113-1120202100211021-1131213303303033-1133232330103031"></a>

## store_provider property — blindfold_secret_info / 320233311300 / 6

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

<a id="canonical-3131021112320113-2201101231132131-1330311313002333-0301121312032202-3233232122231000-1110332010212221-3302332303220303-3011003001012312"></a>

## Next pages — blindfold_secret_info / 320233311300 / 7

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1210303002202023-1322031021321021-3331232302132303-3201130332331023-0303111303030100-0213221021012223-1222230233321230-2123233200312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100102000321121-1201231312112221-0111211333303210-1100131310132233-3302130002022101-2010112133030213-2212003120122113-3011122010332101"></a>

## response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 130120221222 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [response_headers_to_add](resources--virtual_host--reference--group-003.md#canonical-2031023100033202-2330100210303322-3113323303313133-2030111112021212-1302210011133231-0221213311321333-0210033213203311-0133030211030022)
- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3201300033031132-3102213312230110-3200121221130002-1023200120331300-3230013111301132-2030330103012213-0301003133123000-0132311302103310"></a>

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

<a id="canonical-3330122233110102-3233233333223132-1301231120200100-1002223032332122-1312121302103103-1121112221102200-2110232030332130-0332300113330202"></a>

## Direct properties — clear_secret_info / 130120221222 / 3

<a id="canonical-0301030202012021-1031112013213332-0020010013030001-0132201123200003-0213013101010120-3011122021023131-0011023032220332-3011103001102211"></a>

<a id="canonical-1233032023202212-2202030221100120-1323301112021331-2300021333113010-3233121220202233-0323102032010201-3330221103110203-2202132312321112"></a>

## provider_ref property — clear_secret_info / 130120221222 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1320102133121300-3203103220333200-3010331312220321-1133333132303110-1111330002332321-2220023213320031-3003310101013301-0130323201220222"></a>

<a id="canonical-2022000322123101-1013331210130330-2130032211231203-0200301321131111-1130131121110102-1021330310332331-3220221331332210-1203102211223303"></a>

## URL property — clear_secret_info / 130120221222 / 5

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

<a id="canonical-1111120300203231-1212211123131003-1320323130010232-1331231131200021-1333212323003312-2202001132200330-2110033131031333-3122103103012213"></a>

## Next pages — clear_secret_info / 130120221222 / 6

- [response_headers_to_add.secret_value](resources--virtual_host--reference--group-003.md#canonical-2111212032020003-3033322032323303-2321321231321000-1310211200131101-1030333200010020-3101110203003230-1013011320132113-3312210222133312)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103120320021100-3132322131301320-3303002101020223-2203313120301131-2133232032010222-2110101122130112-0122301022110322-1210210230333331"></a>

## retry_policy — retry_policy / 330311130320 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- retry_policy

<a id="canonical-2030200310111133-3030032003003120-3002230020001330-3301223321122321-1312122333313220-0210331300030232-3323011033030330-1311130010323222"></a>

Type: `"object"`. single nested block, Optional.

Retry policy configuration for route destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("retry_condition")}
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
retry_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210212300231223-0202130200102012-2022001032200020-2302031330312231-0101213010112220-0010230323331203-1012120022031110-0313113103122121"></a>

## Direct properties — retry_policy / 330311130320 / 3

- [back_off](resources--virtual_host--reference--group-003.md#canonical-3232203011321132-1200220222313300-0222313113000111-2210230121110011-1033231311222001-3213020130210323-3330003232013232-2301313032100120): complete subsection reference.

<a id="canonical-3321101220111111-0210303233310200-3322021320321323-3330223120110123-1012032303123103-2222330310020030-1321031311333220-2031130212212310"></a>

<a id="canonical-3231103000111123-1120120120302200-3031112331210010-2311311332220023-0300103301203221-0331212022032101-0002111213113230-2333313301310202"></a>

## num_retries property — retry_policy / 330311130320 / 4

Type: `"number"`. Optional.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Upstream description:

Specifies the allowed number of retries. Defaults to 1. Retries can be done any number of times. An
exponential back-off algorithm is used between each retry.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-0330120133021002-2200030001111231-1010010302211330-1110133103332012-2212201003221211-0120331032312010-1021120021131332-3013112231103000"></a>

<a id="canonical-3002033122131031-1313102202113110-2203102223110111-0321001323000312-0203123003010213-0301002123010332-3030012311110011-2133322022101013"></a>

## per_try_timeout property — retry_policy / 330311130320 / 5

Type: `"number"`. Optional.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2103233032202121-2131010003130032-2113013210000222-1111110310120113-0210013031311000-3323331321011201-1030332120032200-3003202201132113"></a>

<a id="canonical-2032030330333200-3313031103131330-2223320111033012-2211232100200202-2320130101323011-0003231132001202-3112010031112001-3201011211132200"></a>

## retriable_status_codes property — retry_policy / 330311130320 / 6

Type: `["list", "number"]`. Optional.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2313310213313002-3233220321023030-0133101030030101-0310300223013100-1123210020031310-0133131313012233-3330200322311330-2011031110102212"></a>

<a id="canonical-0303100321133221-3223331120003301-2030032331203320-1020232033300300-1011112133001230-1313223321000110-0300110203030222-2313202100201313"></a>

## retry_condition property — retry_policy / 330311130320 / 7

Type: `["list", "string"]`. Optional.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Upstream description:

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310130033223020-2121113001202302-1201201000021030-3122023122210111-2330001121003030-1031212103222132-1121021013033332-0133123102113031"></a>

## Next pages — retry_policy / 330311130320 / 8

- [retry_policy.back_off](resources--virtual_host--reference--group-003.md#canonical-3232203011321132-1200220222313300-0222313113000111-2210230121110011-1033231311222001-3213020130210323-3330003232013232-2301313032100120)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3232203011321132-1200220222313300-0222313113000111-2210230121110011-1033231311222001-3213020130210323-3330003232013232-2301313032100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330003022232110-0030132122220001-2302130310332011-0123130131120133-1233112213203332-0030031213203032-1111130000322302-2301102000312102"></a>

## retry_policy.back_off — back_off / 213220020200 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200)
- retry_policy.back_off

<a id="canonical-2022202132020122-0313220232010302-3323220003131000-2031030233333230-3131220223033201-0011222012332112-0112320230011102-2023131003303230"></a>

Type: `"object"`. single nested block, Optional.

Specifies parameters that control retry back off.

Receipt-pinned upstream constraints:

```json
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
back_off {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132332312100003-0111032002223301-1020233210123300-3113302230123030-1022121021220013-0003022232123010-3013300231230120-3323201001131130"></a>

## Direct properties — back_off / 213220020200 / 3

<a id="canonical-0000230233313210-2020102122303230-3321200120031222-1303022020310110-1013313323231311-1333003213221020-3001311030132301-3011222031210332"></a>

<a id="canonical-3213023222022020-3312233101102203-3331313120100132-2020003300122310-0100223202113033-1123322033122300-3332130332033203-3001311113032011"></a>

## base_interval property — back_off / 213220020200 / 4

Type: `"number"`. Optional.

Specifies the base interval between retries in milliseconds.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-3231131112331230-2022002123220210-3113011203301202-2302330310302201-1333332112023330-2310323300012111-1332023211003102-3032211120133331"></a>

<a id="canonical-0330312131333110-1122130220213032-1230111122300132-2020331203232310-3311203303033013-2103011013101133-2122132223300112-2013221002232131"></a>

## max_interval property — back_off / 213220020200 / 5

Type: `"number"`. Optional.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Upstream description:

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The default is 10 times the base\_interval.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2111002103013022-3013133002302012-1113332332021020-1012303330010323-3203031233310211-1032223122331323-1213011011320322-0300103012233103"></a>

## Next pages — back_off / 213220020200 / 6

- [retry_policy](resources--virtual_host--reference--group-003.md#canonical-3120210313011012-3220120012303030-2320221131003020-2003111301323011-0131032003311130-3311130033201323-1332312222032333-0100213232033200)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1230120312310211-1032010120013201-1323220112101321-0101321330220132-1310202220130210-3031233111323031-2123332330332300-2030220320303223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203332011202103-3312112230113330-3022230011132010-3131100222111221-2100031222002021-0121221001223333-2023213131132010-0210103111000213"></a>

## routes — routes / 220001003131 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- routes

<a id="canonical-1001231111002131-3010222122130310-3000212310212132-2332312133330201-3310103133323231-3122230203211313-0231120223222120-1013010013030320"></a>

Type: `"object"`. list nested block, Optional.

HTTP routing rules that match incoming requests based on path, headers, or query parameters and
forward them to appropriate backend origin pools.

Upstream description:

The list of routes that will be matched, in order, for incoming requests. The first route that
matches will be used. Currently route object is redundant in case of TCP proxy but required. For
TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts, the route object only specifies the
cluster/weighted-cluster as route destination without any match condition. In other words, match
condition in route object is ignored for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts.
Routes used for TCP\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY VirtualHosts cannot have DirectResponse
or Redirect as actions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321301102012312-0130312231201111-1211001001321310-0011123122012321-3013110133100300-1303130201220321-0111220010121101-3013031202030200"></a>

## Direct properties — routes / 220001003131 / 3

<a id="canonical-1332302203002203-1323131112012302-3212130321233131-2222020231232012-0112202202120210-0003100001213322-2223322311313002-2012200003301002"></a>

<a id="canonical-0133332003131113-1221032333330000-1111303301311232-1100322231332030-1230133100012312-3233013011232123-1102123023131233-3323212032101031"></a>

## kind property — routes / 220001003131 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3233200101223303-0331333031002211-3222223003320121-1322231321313012-0121030223213223-0322132220323222-3021200203131133-1120233203031200"></a>

<a id="canonical-1323033112303330-3121112113230311-0330312330122233-3001330012202103-0132123333011222-0022130333210112-3130111003322200-1332201100210321"></a>

## name property — routes / 220001003131 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2333233100103023-3300021033112020-3031120032113020-0122110223123013-0002132123121023-3033013230301313-3211201223013131-0220133132133222"></a>

<a id="canonical-0201303320023131-2232010101210301-2110303012122330-0003130032300121-2333313202223122-1213201131100121-0210303011231211-1332310011021213"></a>

## namespace property — routes / 220001003131 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-0103130202322010-0213301300021223-3112001333123233-0103211230330111-1321103321202003-0301023010232321-1303032013202001-2032030210101102"></a>

<a id="canonical-0203023232222200-2020311033001112-1221310021110003-3103022122112100-2003111011330331-2300030301131010-1302100122033112-0120032110100302"></a>

## tenant property — routes / 220001003131 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2102332203002001-0323322233101200-1331003332313110-3223012300133232-1213122100320301-2020213310212303-2200001301312132-1311101333021310"></a>

<a id="canonical-3133223210310032-3322220111032023-2020303220300312-0011223123020331-2233012002023113-3320230231323201-2031220312103333-2303100032110010"></a>

## uid property — routes / 220001003131 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-0022101103221213-2003330300012122-3123302201123201-3223030331321120-3201002320210332-3101003333133301-2113230312002332-2311002121122231"></a>

## Next pages — routes / 220001003131 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1113203023200020-3223013230231323-2212021220030131-0310020233201101-0121323110323203-3122333302133331-0211123332303330-1130022122020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301111313230320-3012201202020311-2230300320120003-0331113130331233-3001332303020033-0023330323321032-2020130012013102-1002011231231223"></a>

## sensitive_data_policy — sensitive_data_policy / 211223020300 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- sensitive_data_policy

<a id="canonical-0222123113222031-2013231001121032-1101203330320312-2232211013122133-3202312112000010-0012121230102013-2101110323013332-3013221200021231"></a>

Type: `"object"`. list nested block, Optional.

Policy configuration for this feature.

Upstream description:

References to sensitive\_data\_policy objects.

Receipt-pinned upstream constraints:

```json
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
sensitive_data_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213312222221233-0210023001231011-2122232120123323-3111200233230032-0020332333132311-3123132200301313-3201021131300222-2220321212010023"></a>

## Direct properties — sensitive_data_policy / 211223020300 / 3

<a id="canonical-0131000121302011-2211210111203202-3202112101330022-3313330121012122-2021031221310010-3111111301310100-0110013310300230-3222333202322002"></a>

<a id="canonical-3030332012323121-1001203320330301-2332302332020222-2031102200203200-1310000303002103-0030233012301021-2232002001202233-0210031231202000"></a>

## kind property — sensitive_data_policy / 211223020300 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3233322232232023-0203210130330331-0303200101211000-2132112013233123-3320231200012123-2210220313133102-2232031332322121-3100231132323232"></a>

<a id="canonical-1321130203331002-0203222233130200-2031120132301123-2323330031211232-1113330111123230-1112131023100303-0000223000031003-3100201222233102"></a>

## name property — sensitive_data_policy / 211223020300 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3110012223312303-1313130030112020-0331100232120323-3321001203301003-0200233222200113-2323133010211300-1320001100121123-2233313303101100"></a>

<a id="canonical-3130331231112132-1130130310120210-1323230223323003-0232102331231312-1111032221222103-1203223300032112-0300300311332020-2202201003010300"></a>

## namespace property — sensitive_data_policy / 211223020300 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-0113201221033133-3130332221120333-1032123330203122-2122101012020032-3103202022233021-1210012302020032-0202132120312133-2231220202012031"></a>

<a id="canonical-3021312233213213-2120311323023002-2123131323000113-1003133210000030-2333023013022221-3132222331233133-1230000321003221-1320001321110120"></a>

## tenant property — sensitive_data_policy / 211223020300 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2110331321310100-1220302312211220-2313213200133332-1303221103013133-1201323220013221-1202021020312132-2200123111300130-0103030110123330"></a>

<a id="canonical-1012012313122012-0220232232302000-1323110220231000-2323131320011100-2033230101121123-1012203011331121-1021032103320033-1330331113020013"></a>

## uid property — sensitive_data_policy / 211223020300 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-0130110013033023-1111001203122220-2100032200310301-3311130033102231-1120101323302301-3123202300020320-0002213203210030-3103011222230033"></a>

## Next pages — sensitive_data_policy / 211223020300 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023200001202302-0103003132331233-1310013232202123-1121131312032122-2011031003022333-0311102200300101-0313122023302130-3131032330331230"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 011122003222 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- slow_ddos_mitigation

<a id="canonical-0213112013313021-0210021331121222-0031003332330302-2013322310222003-2302001202130331-0023232133221022-2330101022202311-2132022201020312"></a>

Type: `"object"`. single nested block, Optional.

'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_headers_timeout"),
  validators.ConflictingObjectAttributes("disable_request_timeout",
    "request_timeout")}
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
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

Terraform syntax:

```terraform
slow_ddos_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320203012320333-1100320111121122-0323210332002103-3010223211003010-3330310113213320-0130333301122213-2010221310003211-0332112022000221"></a>

## Direct properties — slow_ddos_mitigation / 011122003222 / 3

- [disable_request_timeout](resources--virtual_host--reference--group-003.md#canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302): complete subsection reference.

<a id="canonical-0310121203030120-2202111030132111-3230013223020011-0232201133232112-3210322310111002-1020201320121213-2211132302123003-0003003120012112"></a>

<a id="canonical-0102221122133221-1022020211010311-1212310032003123-3302131220313113-2120020013222122-1331120213132312-0230131233100031-0231231023300121"></a>

## request_headers_timeout property — slow_ddos_mitigation / 011122003222 / 4

Type: `"number"`. Optional.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-1010331231022331-1132221033201011-3230133101222221-2013121321230003-3031003203021123-1310310002323232-2012301111230332-3011023231213112"></a>

<a id="canonical-1133310301120320-1302330021000223-3023313201031110-0202111000000230-2233230132133213-0113010200213000-2001221210001122-2100122222320121"></a>

## request_timeout property — slow_ddos_mitigation / 011122003222 / 5

Type: `"number"`. Optional.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2000, 300000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-1301001300002113-0110210332213212-0103132200121003-1103311312332232-2002112023300002-3323110122003101-3210211302301222-0220032033211321"></a>

## Next pages — slow_ddos_mitigation / 011122003222 / 6

- [slow_ddos_mitigation.disable_request_timeout](resources--virtual_host--reference--group-003.md#canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0302112002310101-2021323133201110-2302121112300222-2022200302212321-1020202132131320-1021002113020201-2320232221231103-3212312020311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320023203122311-2022331213113101-1311113111222200-0100101131303113-0201312023232102-0112013203230120-0333200131022012-0310320030133223"></a>

## slow_ddos_mitigation.disable_request_timeout — disable_request_timeout / 110020121302 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-0032120312201033-3200032031011122-0011133112303101-1033331011003302-3232321223332332-3022232232310020-1232233310013321-0201211221133333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable request timeout.

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
disable_request_timeout = {}
```

<a id="canonical-3203220221320331-0013312000002120-3110303013200310-0100230030110000-0123103102022021-1210033332100333-3112023021213102-2323102120303001"></a>

## Direct properties — disable_request_timeout / 110020121302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130310300330010-3110200113232102-3212033120201131-3320113331033203-2221131213020320-0132223101200321-3320121322223233-2111101013033121"></a>

## Next pages — disable_request_timeout / 110020121302 / 4

- [slow_ddos_mitigation](resources--virtual_host--reference--group-003.md#canonical-2333312030122121-1100330320331032-1120232213211221-3131120330330231-2221003131100222-1230211011000221-1033321202131202-1012333211023301)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2133311002322212-0233231012302121-3111302012013321-0301321213223223-0023120022022300-3003323023033302-0131213332220332-1033313111212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120210332031313-0232213211230300-2222320022033110-2113332303321012-2113202310302101-1000100320023030-1230312332333312-3221111310102233"></a>

## timeouts — timeouts / 121233202301 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- timeouts

<a id="canonical-1223303000101110-3011111313112031-1002101201030211-1001300332202033-0331102311120200-0233023330322221-2003333130122211-3310101221230210"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333023101313303-2203310133113000-2210021123222322-2011003133223320-0101203312213210-2110032011101210-0010210020301233-3220020030001231"></a>

## Direct properties — timeouts / 121233202301 / 3

<a id="canonical-1113332021000102-1022012011032030-2313220010330013-3333310312000023-1122113130313220-1013111221012123-0032002330101330-3012130212102132"></a>

<a id="canonical-3020022312213010-3302300230032012-0123210003222313-0033113231003021-0220102303120220-3210313102101133-3223331300311123-0022302232333202"></a>

## create property — timeouts / 121233202301 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3030233233112021-1210231122222201-0022210333223313-0013233323323122-0011321101322012-1120210233132011-3303000120023020-2313103233221213"></a>

<a id="canonical-1211203233201033-3301333112033333-3313100001102231-1210210021132120-0002122012210303-3313310323201331-0202223111331001-0212310120222101"></a>

## delete property — timeouts / 121233202301 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1302101300113101-2232013120221032-2130001111123003-1011023213020302-0311120111212333-1223211000133333-0210003110220230-3230132211103030"></a>

<a id="canonical-0102222030112011-1110310032220110-0230132132032003-1003331212023112-1202102213103312-2021203220223212-0302102211110100-0210332210123130"></a>

## read property — timeouts / 121233202301 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1310103030131222-1332220101011001-1012020331030322-1323110131211203-1231000312020102-1331323032132213-1133012111210022-3001023220230031"></a>

<a id="canonical-3310121322330002-2113031020222001-2200131301213113-3012031033011032-1030233333121223-0332213212310321-3231032133030133-2220022000320203"></a>

## update property — timeouts / 121233202301 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1230232231123231-3211202332001110-0102301121320323-1200312200033132-3121323303012003-2013032231033302-2302210030112301-1012311013130123"></a>

## Next pages — timeouts / 121233202301 / 8

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223313210310003-3130320133220110-0020002300110120-1321100103200302-2020323332123231-0102330022300012-0202210332203312-2022210222212321"></a>

## tls_cert_params — tls_cert_params / 202102211112 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- tls_cert_params

<a id="canonical-2123330110100202-3132023320131302-1212121211011102-3110211230311121-1121003310302011-2120012132033323-2102333031220302-2103213231223331"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: tls\_cert\_params, tls\_parameters\] Certificate Parameters for authentication, TLS
ciphers, and trust store.

Upstream description:

Certificate Parameters for authentication, TLS ciphers, and trust store.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "client_certificate_required"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "no_client_certificate"),
  validators.ConflictingObjectAttributes("client_certificate_required",
    "no_client_certificate")}
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
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

OneOf alternatives in this subsection:

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-2123330110100202-3132023320131302-1212121211011102-3110211230311121-1121003310302011-2120012132033323-2102333031220302-2103213231223331)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0221010032112232-0002311102011103-1002323023221012-1223320011202010-1100310130023232-0301233120212133-0021101133201301-2201231323313333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302331222310233-0002032133320203-2021113202031331-2330211212102221-2301000110130200-2103331301333223-1321303201020101-3311011013301332"></a>

## Direct properties — tls_cert_params / 202102211112 / 3

- [certificates](resources--virtual_host--reference--group-003.md#canonical-1333132230002120-2220113212031031-1100313100101320-2112323110202120-2103200122222302-2101113102020123-3020020213230333-1303213032212222): complete subsection reference.

<a id="canonical-3332232212210020-0203212331130212-1300121313102121-1221223322011312-3333022100321011-2130003130102020-3332020313022023-0321302322211331"></a>

<a id="canonical-2023023111200213-1310102301203203-2122223230321101-2103230020333130-1111020121002010-2121010300303132-0110320022210221-3302032032233331"></a>

## cipher_suites property — tls_cert_params / 202102211112 / 4

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-2220311223310000-1311303333221131-0123212110113222-3113013210012112-3331322002321331-0123212233121233-2003331313111101-0301010133230130): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231): complete subsection reference.

<a id="canonical-3202210313202032-2103210032220213-2333313310011023-3320000022101131-0032313032232223-0101023100203130-2320333202112301-0003222302323300"></a>

<a id="canonical-3102200032010101-3010201303313322-3211330211113300-2333221303110222-0200320020031110-3002212130301300-3321222001232131-0322323223300320"></a>

## maximum_protocol_version property — tls_cert_params / 202102211112 / 5

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

<a id="canonical-3113032321331133-1323121012031212-3213020313001013-3302110032002301-2213313313322121-0230022201023101-2332323023210222-2013122300311010"></a>

<a id="canonical-0322313210000210-0311320131211300-3201213003303113-2102201320210032-3202023020011131-2123031321202023-0130311113001301-3113211221321202"></a>

## minimum_protocol_version property — tls_cert_params / 202102211112 / 6

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

- [no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2101132133101113-1022120223131223-3310112111001032-1310100230310130-3020001203133321-0211202000130300-3012200301033002-3323223231312321): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-003.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112): complete subsection reference.

<a id="canonical-3200233311230203-0032131020020012-0003320011300330-2100210331110120-2023023332113213-2013120223101120-0313131322110120-3022030102321311"></a>

<a id="canonical-1131300120323211-3330103322030200-2033110210230202-0330330323331111-3111302020212320-0120103002030020-1013301213311023-0133221130233213"></a>

## xfcc_header_elements property — tls_cert_params / 202102211112 / 7

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3203220331013213-0133213212211103-1112110020012311-2131031002202301-1322233002320233-0023121321110003-3301220221322030-2103101113232023"></a>

## Next pages — tls_cert_params / 202102211112 / 8

- [tls_cert_params.certificates](resources--virtual_host--reference--group-003.md#canonical-1333132230002120-2220113212031031-1100313100101320-2112323110202120-2103200122222302-2101113102020123-3020020213230333-1303213032212222)
- [tls_cert_params.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-2220311223310000-1311303333221131-0123212110113222-3113013210012112-3331322002321331-0123212233121233-2003331313111101-0301010133230130)
- [tls_cert_params.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231)
- [tls_cert_params.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2101132133101113-1022120223131223-3310112111001032-1310100230310130-3020001203133321-0211202000130300-3012200301033002-3323223231312321)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1333132230002120-2220113212031031-1100313100101320-2112323110202120-2103200122222302-2101113102020123-3020020213230333-1303213032212222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032121102300303-3131122002332201-0333000231202230-0230013103102331-1220321201030021-1022330220010001-0010200331122201-1320301132100233"></a>

## tls_cert_params.certificates — certificates / 120032032320 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.certificates

<a id="canonical-1002330313231012-3132011231022303-2330322112320020-3333223010230220-2230020323030221-3301101101021312-0001221121123213-3222202330100022"></a>

Type: `"object"`. list nested block, Optional.

Certificates. Set of certificates.

Upstream description:

Set of certificates.

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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110311310020331-1103321012101210-3330323100002013-2130313112002110-3121213101330233-1113123033022220-0312111331002010-2000002003000132"></a>

## Direct properties — certificates / 120032032320 / 3

<a id="canonical-2030112331132002-3113310132003032-2331232100230000-0200230011323213-0022103130310033-1103130211221313-2003301233203132-0300301222330133"></a>

<a id="canonical-2330203322303221-1231012133201213-1120213031223302-2130012330110133-1032120030233113-3332211203112130-3032020123111103-3032231111330333"></a>

## kind property — certificates / 120032032320 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2000100313113101-1301313210100011-2200312233122122-2221320303323322-0102032301030311-0030302132331230-2332202022231223-0112023233333321"></a>

<a id="canonical-0311023203332223-0030230110321211-3323121000200023-0023103122030123-1231022312033010-2112332102320330-2102233233300000-2121000000212212"></a>

## name property — certificates / 120032032320 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2320202202302103-3203100131231231-2202033310030023-3323131020333232-2202213202302212-3210222330301000-2320300010333030-1013121122130111"></a>

<a id="canonical-2321313231330003-1322301112013313-3113313222001123-1000221333220313-0203010030330302-0122011121122220-1223230100021333-3203332101131201"></a>

## namespace property — certificates / 120032032320 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-0210103321000312-3101022020303002-0222032131002212-0301330303330320-0110201223120300-2323233303233003-1221103312002120-1200112220101233"></a>

<a id="canonical-1321000220202222-3313222313101202-2110211223033322-3111310111333101-3213132002310331-1300123301321002-2031303022101222-2103121112200313"></a>

## tenant property — certificates / 120032032320 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0033103333210122-3131033102132023-2103000002320311-2201223212010302-2003002203313230-2311333021012310-3010020111030332-3212120020301103"></a>

<a id="canonical-0210223011013333-0312331332130211-3322011303301203-2123131000323210-0331320201130000-1302020333223130-0111113320103133-0110123110223133"></a>

## uid property — certificates / 120032032320 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2332200223203001-3022220031010031-0012212000332111-2102323013111303-0022231200001221-0001022232112210-1033121221001230-2232031021123112"></a>

## Next pages — certificates / 120032032320 / 9

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2220311223310000-1311303333221131-0123212110113222-3113013210012112-3331322002321331-0123212233121233-2003331313111101-0301010133230130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113221233330031-2000310023013211-1111312132110211-1003002331203103-1200133113331112-2011213123221103-3212323133102121-3331020000322202"></a>

## tls_cert_params.client_certificate_optional — client_certificate_optional / 010010230011 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.client_certificate_optional

<a id="canonical-1223302233121202-2213311121111220-0322222120322002-2333330312333021-0301101302320000-1331020233022201-0301232202203023-1322033023121103"></a>

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
client_certificate_optional = {}
```

<a id="canonical-2330112202021221-1033000311002123-3210030212302022-1231223100332331-0313213101123212-0201311203201211-1330031311310222-2022323310022212"></a>

## Direct properties — client_certificate_optional / 010010230011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332000033212330-0012300210320031-2231023310120030-0322233001333301-0323103022233102-2310000023132302-0031101301001233-3023310112121030"></a>

## Next pages — client_certificate_optional / 010010230011 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302112023330212-1230223203222101-2203220232113203-0203231211212322-1103212201123231-2210232332220022-2322032130302131-2113013111101330"></a>

## tls_cert_params.client_certificate_required — client_certificate_required / 233032112210 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.client_certificate_required

<a id="canonical-2102010012200310-3201101313122222-3023032020132310-0123202022323332-0220210320300102-0230223003110030-3100321002312200-0300000220221331"></a>

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
client_certificate_required = {}
```

<a id="canonical-1032332333023212-1020231013112320-2300222330330000-2330121311210202-3321123023111211-0001023322312210-0120200030330123-2021333102312330"></a>

## Direct properties — client_certificate_required / 233032112210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222210311312032-3303212300131201-2213230113331211-1210110210101232-2213031130303332-1332212332130312-1233231020023121-3321230323031231"></a>

## Next pages — client_certificate_required / 233032112210 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2101132133101113-1022120223131223-3310112111001032-1310100230310130-3020001203133321-0211202000130300-3012200301033002-3323223231312321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322223111013022-0002112303123031-0213220110101302-1130201010113203-3203210320221332-2123110013222030-1020003101232001-2310223100131201"></a>

## tls_cert_params.no_client_certificate — no_client_certificate / 212220203201 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.no_client_certificate

<a id="canonical-3310213000123010-3132121002122102-3103230132120112-3301100200202131-3310211001131130-2332131103111121-0010012310330323-1132323202112103"></a>

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
no_client_certificate = {}
```

<a id="canonical-3200132301323331-2330003312031032-0003212030220213-1233223033232210-0133332101200112-3320023011013013-2031102003221310-2110313323011222"></a>

## Direct properties — no_client_certificate / 212220203201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103002212222011-0322021111323000-3223210030201231-1132232311210133-2101312013300213-0202310003101123-3003122131011312-3212001302002120"></a>

## Next pages — no_client_certificate / 212220203201 / 4

- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032013002112111-1202102003233113-1222033130312011-0310303110313100-2233301333213120-3010232113200033-2120313012230031-0110103100102221"></a>

## tls_cert_params.validation_params — validation_params / 220310022203 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- tls_cert_params.validation_params

<a id="canonical-1032202302300211-3001333330201033-2112023021131021-2121022110300031-2233123033220131-3032133101222000-3022121132022232-2331211111102331"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201301313202103-2000001323032032-2000302212302022-2213202201022132-1211313033101301-1130111232013223-2300120123201112-3031102303210222"></a>

## Direct properties — validation_params / 220310022203 / 3

<a id="canonical-0110332132111203-2133212021330330-3212100333320203-0022213121203020-0233000003322130-3211110100121032-0312120230130232-1323222112100233"></a>

<a id="canonical-2102110233011301-2303131202200200-2123332333210323-2230122201233231-3122033020002130-1212033023003031-1201233110011011-1202200003102312"></a>

## skip_hostname_verification property — validation_params / 220310022203 / 4

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](resources--virtual_host--reference--group-003.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131): complete subsection reference.

<a id="canonical-0102232033030023-0113332012201000-1312102323333021-2031133211031031-3313020313320333-1300232331030302-3302013232323323-1011120210323103"></a>

<a id="canonical-3323122103230230-3003222321021230-1013231221212300-2022213212021001-3121100000301013-0313001101101122-3200302321030112-0012001211021332"></a>

## trusted_ca_url property — validation_params / 220310022203 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1330331320133131-3230022131231000-1320110033311002-0013002210032220-3201200210123011-3210101321101131-3020300122020130-3002301301100201"></a>

<a id="canonical-0233031023132210-1132130202130322-0233220110123301-1313330313310130-1032030100030110-1303020021000312-2131130003203333-2203331133322231"></a>

## verify_subject_alt_names property — validation_params / 220310022203 / 6

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1031001130231103-3122102103303330-2313023012120311-3133013302110303-2230233012010333-2303223013211122-1302313330213230-0220130331310112"></a>

## Next pages — validation_params / 220310022203 / 7

- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123033302102100-1220322312103133-1213210230300221-2120200222003132-1333132110230012-0213032120221210-1133132122322131-1210010020311201"></a>

## tls_cert_params.validation_params.trusted_ca — trusted_ca / 033132033321 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- tls_cert_params.validation_params.trusted_ca

<a id="canonical-1002033131101220-0211131331011102-3222212131332030-3030112200103030-2121330011020101-1333203122110120-0322323002313131-1131020202321102"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130100300033123-3332200203032021-0112000002000131-0231003300300013-1201001311010223-2030232033302322-1010320313001103-3330113122112133"></a>

## Direct properties — trusted_ca / 033132033321 / 3

- [trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-0231232332310003-0020000023320100-2002120131220321-2130121301021230-0201332013033300-1130032113023210-1112333030122301-3201313113210030): complete subsection reference.

<a id="canonical-3121303023302302-2233013031221132-2230332222302030-3223000213013021-3013102333223331-0322300001031011-0133223012001231-3031302010203233"></a>

## Next pages — trusted_ca / 033132033321 / 4

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-0231232332310003-0020000023320100-2002120131220321-2130121301021230-0201332013033300-1130032113023210-1112333030122301-3201313113210030)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0231232332310003-0020000023320100-2002120131220321-2130121301021230-0201332013033300-1130032113023210-1112333030122301-3201313113210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021201233310002-1331113033002010-1132320021013012-1303222032202111-3032310100030011-3322100011202312-0230113021010122-1032031021001122"></a>

## tls_cert_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 301000222000 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_cert_params](resources--virtual_host--reference--group-003.md#canonical-3110210131021323-3322110122123101-2111232233031031-2333002311012210-2321323210033002-2221332313013332-2032010210303312-0230121122130331)
- [tls_cert_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-2111212302120220-0101203001130310-1133231301002100-3121101200220032-0321213313100212-2200021303113021-0023002112012032-3022331122002112)
- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131)
- tls_cert_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1111101121102001-0321110223122321-2230131322101220-0303201223120133-3123230230020330-3232201232300022-0001002321302333-1002213333312003"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010131103023031-2122223112000033-1131203112013112-0002100230012322-2003100302230102-2310213203030131-0012330303022311-0211133011201330"></a>

## Direct properties — trusted_ca_list / 301000222000 / 3

<a id="canonical-1333112023322221-3313100101333133-1002320321022103-0220322010211232-3200303231013330-3201333310012122-0331001331123100-0232021113032202"></a>

<a id="canonical-1311221031121233-0331300210000112-0100122231332012-3231203332102203-0111130201231313-3013301032202230-3202200300102332-0000011322131302"></a>

## kind property — trusted_ca_list / 301000222000 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3323131031113030-3113210102313220-3103232022202130-3303030002230221-3323133211013021-3130223000030013-3011223210330211-0123211231100323"></a>

<a id="canonical-1222031132233302-2312010221202012-1021201023330010-0323112211100330-3131023203233011-2023031000221101-1322022012011220-2222322223103310"></a>

## name property — trusted_ca_list / 301000222000 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3121030130012330-2210330211031333-1132310012230213-1103113012210102-2213032002020302-0210123121221030-2110311331000333-1002331332001011"></a>

<a id="canonical-3321210211003110-1001212022213032-1333003112021013-2213120233310312-0302323213220000-1010323013133030-1102010002213330-0120110101200120"></a>

## namespace property — trusted_ca_list / 301000222000 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-1100201012230233-0213301233120301-3033323121112133-1023222100113222-0002201103130021-1100313113220113-1212201220111220-1203033102002311"></a>

<a id="canonical-3103310210213302-3013213303301111-1211210022022100-0030213233121301-0231233201331333-2032200013223101-0303122212032031-2332033020233003"></a>

## tenant property — trusted_ca_list / 301000222000 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0233133023123231-3003132311323220-2110120101023033-3323303310130133-1120030212121311-1232200123221003-1121130203203202-0010210011200322"></a>

<a id="canonical-1312231330211201-1020320002223220-0200000301230313-1003302220213223-0212200033123213-2023123130323102-1102213210111021-3220130102202333"></a>

## uid property — trusted_ca_list / 301000222000 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-0011020203333303-2122332233223023-1221332111123122-3323210301311321-0302000113200120-1321132200023122-3232310223011033-0020022022011313"></a>

## Next pages — trusted_ca_list / 301000222000 / 9

- [tls_cert_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202300301300021-0201230301102010-3133213031032021-1310112212332023-2030020331322110-1001021320131223-2233102230201022-1032130013313230"></a>

## tls_parameters — tls_parameters / 223211033201 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- tls_parameters

<a id="canonical-0221010032112232-0002311102011103-1002323023221012-1223320011202010-1100310130023232-0301233120212133-0021101133201301-2201231323313333"></a>

Type: `"object"`. single nested block, Optional.

TLS configuration for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("client_certificate_optional",
    "client_certificate_required"),
  validators.ConflictingObjectAttributes("client_certificate_optional",
    "no_client_certificate"),
  validators.ConflictingObjectAttributes("client_certificate_required",
    "no_client_certificate")}
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
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300231130210032-0202022323200113-3023213011202030-1102212000231313-2102020002303332-2222231331213230-1212003120013103-0123223030131030"></a>

## Direct properties — tls_parameters / 223211033201 / 3

- [client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-2322133313323222-3132132032131311-2200332013110013-2311323311021120-2130302130010230-0122102013022233-3013312030021222-2121002101111211): complete subsection reference.

- [client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-3202130002100313-3301122123011011-0330322120202303-0333011230321202-1223210103310323-2303032110201303-3012000030223131-3231311133101122): complete subsection reference.

- [common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221): complete subsection reference.

- [no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2213212201113300-0223221321112131-2300013131310300-2330130313033001-3003001223303232-3000310231222000-2101120210302122-2300231020001031): complete subsection reference.

<a id="canonical-0213200310122130-0101300331002010-3102233222333003-2122132330021030-3201110012210001-3122300102120002-0031012231012200-2300100230221000"></a>

<a id="canonical-3132022213022303-0032233311310231-0210000130230211-0022332122020031-1313012313020110-0330200231103321-2111332031320302-3122310132013212"></a>

## xfcc_header_elements property — tls_parameters / 223211033201 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1110303213021031-0220132022102111-2203303033311303-2012331013100223-2210213012010212-2310002222123100-1101133133333202-3112111320022303"></a>

## Next pages — tls_parameters / 223211033201 / 5

- [tls_parameters.client_certificate_optional](resources--virtual_host--reference--group-003.md#canonical-2322133313323222-3132132032131311-2200332013110013-2311323311021120-2130302130010230-0122102013022233-3013312030021222-2121002101111211)
- [tls_parameters.client_certificate_required](resources--virtual_host--reference--group-003.md#canonical-3202130002100313-3301122123011011-0330322120202303-0333011230321202-1223210103310323-2303032110201303-3012000030223131-3231311133101122)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.no_client_certificate](resources--virtual_host--reference--group-003.md#canonical-2213212201113300-0223221321112131-2300013131310300-2330130313033001-3003001223303232-3000310231222000-2101120210302122-2300231020001031)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2322133313323222-3132132032131311-2200332013110013-2311323311021120-2130302130010230-0122102013022233-3013312030021222-2121002101111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130220120131200-0123213322021112-2311211111030122-0212202132003031-2201232113302122-1132113023100030-1102111210011120-3103131302231233"></a>

## tls_parameters.client_certificate_optional — client_certificate_optional / 201331302130 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.client_certificate_optional

<a id="canonical-0333121111220133-0323020310031210-3333011101131031-3201121301201312-3211211020311333-0123102320230100-3032302023333002-3010000221101202"></a>

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
client_certificate_optional = {}
```

<a id="canonical-2321023222323223-0222002313000132-1333101322101111-0210313311223010-2120030010311212-3003121101303011-0010203110311301-0230200003311300"></a>

## Direct properties — client_certificate_optional / 201331302130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310020320221221-3210311123212030-1112111021232020-2232200202303313-0111020112121300-2031002311210132-1200223132100313-3002013210311210"></a>

## Next pages — client_certificate_optional / 201331302130 / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3202130002100313-3301122123011011-0330322120202303-0333011230321202-1223210103310323-2303032110201303-3012000030223131-3231311133101122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110221030102211-3202132122203101-2133313230001011-0332120133021321-3220232000301311-0301123202033013-3012112310221132-3030021230013211"></a>

## tls_parameters.client_certificate_required — client_certificate_required / 303023301000 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.client_certificate_required

<a id="canonical-3130232032303203-0232011333223321-0030113131322311-0313010120132021-1001000100003301-2121210321312131-2013311200100131-1313302322110000"></a>

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
client_certificate_required = {}
```

<a id="canonical-1313330313232030-1233102310112030-0200112002111330-0122133200322333-2323102313222321-0133321113133321-3303202010021123-0223111331201121"></a>

## Direct properties — client_certificate_required / 303023301000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032213212110231-1213203002213100-2123312212000200-0120320012233313-3200112121210232-1122023001223302-1332233001001221-0201321111310101"></a>

## Next pages — client_certificate_required / 303023301000 / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031103330110211-0101210231203021-2223103233122030-3122201003331331-2313123211223010-3233121023000023-0302021323121200-0311100020323201"></a>

## tls_parameters.common_params — common_params / 321013013232 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.common_params

<a id="canonical-0303120133133233-0132220103133322-1022302231330312-1302033300200132-0301031130022023-3013200203121330-3031202120031130-0010023113130121"></a>

Type: `"object"`. single nested block, Optional.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Upstream description:

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

Receipt-pinned upstream constraints:

```json
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
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301300101003101-3322132033112203-2322213321031301-0023103220133113-3023000333131002-2010002202210121-0031013223132312-1032021111003110"></a>

## Direct properties — common_params / 321013013232 / 3

<a id="canonical-2310321300133003-0230000103032103-1231003120202213-3113000213210122-1131113331001110-3222122220323311-1212222013202333-0213312132100201"></a>

<a id="canonical-0222203000002123-2031022220122100-3200233223320113-3101323213310133-0222112202221013-1323311100222311-0312221202331111-3212312131312300"></a>

## cipher_suites property — common_params / 321013013232 / 4

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256..

Upstream description:

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3133332330120023-2322121212320313-2131311120313023-2222221001230020-3221112322200303-3231220021213000-0022100122003132-2333333332323110"></a>

<a id="canonical-0001300222312121-0303110012332102-1220113333233323-0220121031310213-3112303201002333-1103022231133301-2002323223002023-1002011100100222"></a>

## maximum_protocol_version property — common_params / 321013013232 / 5

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

<a id="canonical-0213203310111320-3203301022001113-0203122220203000-3200302123110222-3231321112312320-1303333303122323-3232311230032312-0323311013031321"></a>

<a id="canonical-3301212023032233-2212122233213133-1001101210312213-2203332313212030-0220102022030123-2333131102013220-0013201300311002-1300121303331022"></a>

## minimum_protocol_version property — common_params / 321013013232 / 6

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

- [tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203): complete subsection reference.

- [validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123): complete subsection reference.

<a id="canonical-2300120211312301-3113020021102001-2220210110020023-2032122012013131-0001031233221212-0233103211321232-1103020120300231-1110322321122132"></a>

## Next pages — common_params / 321013013232 / 7

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222312033203012-2310212010220201-2311031210321032-3230113200332301-3301112312323100-0223301230002223-2033311311110022-1222331300301323"></a>

## tls_parameters.common_params.tls_certificates — tls_certificates / 033022322312 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- tls_parameters.common_params.tls_certificates

<a id="canonical-1011323231323101-2001121311000010-1021211103001120-2203300311211232-0231231033120330-0102330323332022-1001313320010300-1222312131021200"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011330000300003-0232002132220312-0323031231210121-1101213331133212-3113010300231333-2022033233121130-3210033132021220-2232232210033310"></a>

## Direct properties — tls_certificates / 033022322312 / 3

<a id="canonical-3302211131233331-1203311302023312-0302121312102232-2033201220322003-0033302011231110-2211111233000312-0111112211001122-2012210123233002"></a>

<a id="canonical-0213003303301332-0023101021001330-3132021101223131-1132323033210231-0220003201311312-0323321111000131-1332023001203002-1200212033202101"></a>

## certificate_url property — tls_certificates / 033022322312 / 4

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

- [custom_hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-1000011133010022-3111113122203210-0111223313002300-0133022021203110-0330111101023322-1230023110021010-1223202212030000-1320021132303212): complete subsection reference.

<a id="canonical-1011010332212323-1313312303231200-0033031122132330-3033330310233021-3030002010001233-2123323233323022-2120221301330301-2322220212223200"></a>

<a id="canonical-0221302232112100-3033132221210220-0211013132022122-3332000120323123-3021220100220210-1033232312230022-1230102310033201-3303330000022331"></a>

## description_spec property — tls_certificates / 033022322312 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--virtual_host--reference--group-003.md#canonical-3230303332000320-0100030312022332-1212333031013131-1303111211023103-2201222130212130-0230030321331231-0003312300322110-0003032221332333): complete subsection reference.

- [private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113): complete subsection reference.

- [use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203): complete subsection reference.

<a id="canonical-2022011121133331-0232321232003013-0210313202332020-1021100221101311-3332110202210221-2131230320033122-1023010313010320-2310230210323330"></a>

## Next pages — tls_certificates / 033022322312 / 6

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--virtual_host--reference--group-003.md#canonical-1000011133010022-3111113122203210-0111223313002300-0133022021203110-0330111101023322-1230023110021010-1223202212030000-1320021132303212)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--virtual_host--reference--group-003.md#canonical-3230303332000320-0100030312022332-1212333031013131-1303111211023103-2201222130212130-0230030321331231-0003312300322110-0003032221332333)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--virtual_host--reference--group-003.md#canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1000011133010022-3111113122203210-0111223313002300-0133022021203110-0330111101023322-1230023110021010-1223202212030000-1320021132303212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323110232232330-3301102323123221-0012210030233031-0312220201333220-1030331202200110-1313221022001022-1113333301110111-3002133012310210"></a>

## tls_parameters.common_params.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 302113023303 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-1301331000212220-0123310303202130-1202122222101313-0021222302101313-3330133233100022-3013103103331121-2002100321130302-1031320201012133"></a>

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

<a id="canonical-2020020021110021-0312100011223001-1101203120012212-1100031023221102-2232330213331330-3031012103022031-0032223012033010-3122130022302113"></a>

## Direct properties — custom_hash_algorithms / 302113023303 / 3

<a id="canonical-2101210120001021-3312301120323023-1100101002233120-3322221303231330-3320000311132031-1330223130103001-3302010013211200-3331031002333311"></a>

<a id="canonical-3133230133020131-0203111201322230-1332132112201203-2002232310202102-1202021100112110-0010330003303330-2121213201033113-1310112012012231"></a>

## hash_algorithms property — custom_hash_algorithms / 302113023303 / 4

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

<a id="canonical-1121123032031011-0012310320112112-0011121122103103-3121030322023112-2323101220200331-0212232321132331-0303022130033323-2200033330312211"></a>

## Next pages — custom_hash_algorithms / 302113023303 / 5

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3230303332000320-0100030312022332-1212333031013131-1303111211023103-2201222130212130-0230030321331231-0003312300322110-0003032221332333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022321131331313-2221130131331221-1202011310333123-1021233312120110-0100133320311231-2211230213200000-3221301122221233-0122002030113130"></a>

## tls_parameters.common_params.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 233302031123 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-0101302010000031-3023011131112013-2233103023132011-2112211322201222-0333030013211102-1112301112212301-1323133003022312-1222332133003033"></a>

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

<a id="canonical-3320232231130223-0003231310112132-1231222201020330-0320001131030333-3232322333000030-0123101113020113-1023310033301332-0203200113221101"></a>

## Direct properties — disable_ocsp_stapling / 233302031123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213011122123121-2333000322130121-2032210120100002-1023232202101111-0311030103020320-2221200332302020-1220322311202310-3212220320323221"></a>

## Next pages — disable_ocsp_stapling / 233302031123 / 4

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201002011011313-2121231322333130-0021212221313000-0001210220203230-3021111102330320-3021320031212323-0310201100131010-3232001212201110"></a>

## tls_parameters.common_params.tls_certificates.private_key — private_key / 211102100210 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-2302212021121312-0000312312213033-0102223312230203-1231320211020201-1302122302102123-3130030311202313-1210120133033102-3012110331121213"></a>

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

<a id="canonical-1133112110211322-3222111112100203-1323023301310233-3121020001311013-2233212123300201-0331031010303002-2133110213001202-1113122132110233"></a>

## Direct properties — private_key / 211102100210 / 3

- [blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333): complete subsection reference.

- [clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-2321012223121233-0320222032102230-2330222220013231-1100112312301103-2312121002131223-3302012231330100-2130300301011211-3120012312133212): complete subsection reference.

<a id="canonical-1002013322312102-2330100032220112-1001133202113210-3033320110312311-1031031303231313-0212310332012020-2221123032022021-3211120002323110"></a>

## Next pages — private_key / 211102100210 / 4

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](resources--virtual_host--reference--group-003.md#canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](resources--virtual_host--reference--group-003.md#canonical-2321012223121233-0320222032102230-2330222220013231-1100112312301103-2312121002131223-3302012231330100-2130300301011211-3120012312133212)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2303223212323122-2111120020132110-2303311132222211-3122020120123201-1301200220013030-0121301312210220-3202033203132021-2010113331023333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002001010000323-0221021330011113-3001300321111120-1020122020201010-2300232030223203-3103030130133030-1123002032011313-2220312320220220"></a>

## tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 310313233121 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3231101123122320-1201100232102023-2010212231312012-1001113320211332-3213000223222222-2013031102003301-2131313102300223-3023321110213033"></a>

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

<a id="canonical-0203003230200123-0030312100320110-2103333010100311-2101112223102203-2120010201130232-0212021300223132-0022202323111130-0333201032001333"></a>

## Direct properties — blindfold_secret_info / 310313233121 / 3

<a id="canonical-3030230232232312-3303112101112320-3212030010310032-1010013100021230-3031122122133033-2220333111231020-2003122101201021-0213213332311220"></a>

<a id="canonical-2233300002303213-1033313331331201-2233331233331213-0222300132110113-0020102112021103-1301033213112220-2123312111330131-0011023303212023"></a>

## decryption_provider property — blindfold_secret_info / 310313233121 / 4

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

<a id="canonical-3100331112122311-3220232233111123-0321301322130101-2002001331113332-2202302003200313-3322032002010310-2322303201033033-2113030033012023"></a>

<a id="canonical-3312211202300221-1303132000210123-3012301002222330-0023211331311332-1133213101013220-2022030033302313-3210122200313331-3020213211320100"></a>

## location property — blindfold_secret_info / 310313233121 / 5

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

<a id="canonical-0303011010303232-3221003323310310-3110111121213110-0123330223321133-2323313211311122-1233331330121300-0032333223022030-0232030103311203"></a>

<a id="canonical-1132301333031232-2113312231303103-2302112032232101-2021230223232302-3033303001310121-1033203201120302-2100331232213301-1031101012020301"></a>

## store_provider property — blindfold_secret_info / 310313233121 / 6

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

<a id="canonical-1101302033310233-2323120121033323-2113131201031232-0301211303122131-0110301231021021-1023210010331021-1132033110133233-3313302201311300"></a>

## Next pages — blindfold_secret_info / 310313233121 / 7

- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2321012223121233-0320222032102230-2330222220013231-1100112312301103-2312121002131223-3302012231330100-2130300301011211-3120012312133212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211031122221310-3131033323020012-1233030000323012-2202300013120112-2212330300220310-1322302223223320-2100012310203112-1212030100130030"></a>

## tls_parameters.common_params.tls_certificates.private_key.clear_secret_info — clear_secret_info / 003301201012 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0301020201000130-1303100322320331-2223120202001311-0022210010221201-3120021103322203-3002120231130230-2002132332210033-0030221312122011"></a>

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

<a id="canonical-2013230120032311-0000310023210312-2000120021231311-1320012302201130-1100333011330310-2301201021033112-2111013212322221-0210110110221022"></a>

## Direct properties — clear_secret_info / 003301201012 / 3

<a id="canonical-3313010330101223-3000011030300100-1313321023120311-1230220310221010-2011111133113232-1330302100232100-0131211033333122-0102000321130332"></a>

<a id="canonical-3011231000010221-1323322323230212-2201031100133311-2313100033100130-1312130310203232-1200313322301210-2031110010212333-1232123121021300"></a>

## provider_ref property — clear_secret_info / 003301201012 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1022201333211033-3102322311210000-3323103323031202-1213302020101211-2330103303103001-0202000111131313-3223221132332232-1031323102102212"></a>

<a id="canonical-0130010212303011-2300201123033100-1123101302132022-2302122232033033-3120031032011111-1333021302232313-0210010023220211-3232033111202013"></a>

## URL property — clear_secret_info / 003301201012 / 5

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

<a id="canonical-3203131223023220-0203111112222212-0132123101110022-1021131223013022-3213320121330312-0102110330110130-2310232232120101-3333110312211323"></a>

## Next pages — clear_secret_info / 003301201012 / 6

- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--reference--group-003.md#canonical-0100101001331011-1232022323322231-2223232000102312-0213110033331032-0321030213022011-3300202030223321-1111110020102320-1230230111233113)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200332221011201-3000020133103221-1223132012302311-2313223200002021-3031011020311200-0230033203101310-0012332223113330-0131032121110231"></a>

## tls_parameters.common_params.tls_certificates.use_system_defaults — use_system_defaults / 020311313302 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-3233303122201020-3302201102012022-2002230132202021-1022313312101123-2320113202122202-1223303101311133-2323210201310212-0132132113032203"></a>

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

<a id="canonical-2000323201100012-2000233013013321-0320121111130111-3210012122000103-0220013213331232-2331321323022032-0213002302111311-3113003111232002"></a>

## Direct properties — use_system_defaults / 020311313302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131231321032331-0023000112221330-3102313312123333-2000302221301303-0011212321212222-3000212213113100-0322122231012131-0333210231313210"></a>

## Next pages — use_system_defaults / 020311313302 / 4

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--reference--group-003.md#canonical-0200003231223022-1321323211123233-0123321102300103-1301210333113212-1111032300001213-3313322300133213-3311332031202333-2110202132233203)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301330301201003-2010333331220223-1122331200131302-0330033230121221-1002231320303311-0232101211331100-2212320022001213-1320013320012103"></a>

## tls_parameters.common_params.validation_params — validation_params / 120100012312 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- tls_parameters.common_params.validation_params

<a id="canonical-2031103232112023-1301131322123321-3322010330112122-2333130021322233-1130323231323001-3333003021210333-3103131313200010-2310321101122310"></a>

Type: `"object"`. single nested block, Optional.

Includes URL for a trust store, whether SAN verification is required and list of Subject Alt Names
for verification.

Upstream description:

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212201002222100-2022321301323131-2123323321301213-1233111111123230-0022100312332211-3330112220122321-0100321101132330-3131320012113120"></a>

## Direct properties — validation_params / 120100012312 / 3

<a id="canonical-2300021102323300-0221100330002100-3333213200321211-2031123132031202-3113301010112003-0223102333201011-3303311102310130-2213332022212001"></a>

<a id="canonical-1112033323203213-1230302022330223-0110030132110023-1111131030022202-3102302332323321-2332003320020303-1311212000203230-3320020023130300"></a>

## skip_hostname_verification property — validation_params / 120100012312 / 4

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Upstream description:

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311): complete subsection reference.

<a id="canonical-2201330001003130-1103011331220233-0222132032022310-0202130110120123-3313013301131232-2100012011220310-2311130330033130-0310300011033013"></a>

<a id="canonical-3330021002001022-1323010220232132-1000233310023133-3122311331231010-1321011030301331-1301233120103223-2323221222122322-2122233002333302"></a>

## trusted_ca_url property — validation_params / 120100012312 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Provider validators and defaults (from schema source):

```go
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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2322030231212021-2033120211323221-0300330333313200-2030313333202300-0101103133202310-2102020323122011-3120212302200022-0213121113110013"></a>

<a id="canonical-1103330330323202-2321010330103022-0132002332330301-1113022110222313-1223030010003113-2333201122200132-3321320122132111-3321220133230100"></a>

## verify_subject_alt_names property — validation_params / 120100012312 / 6

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Upstream description:

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133321220200133-3100312200110120-1221023010003201-0201222010001232-3131310010320130-3303312131322300-1011210132130013-2110011020121103"></a>

## Next pages — validation_params / 120100012312 / 7

- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312013010303301-2311133022100130-1313323120230213-1021302022201120-0002213003322002-1122022203321230-0312323322301001-3212231330212133"></a>

## tls_parameters.common_params.validation_params.trusted_ca — trusted_ca / 102120302221 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-2102002130033323-2330120222232233-0302302223102230-2321122000230332-2002221220220110-0222020122210222-0310213113000300-1221310213031203"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312303322103001-0301112322111120-1201230231230100-0020220101111200-1200203010312300-0101010230023222-1201133000232012-1100023022032231"></a>

## Direct properties — trusted_ca / 102120302221 / 3

- [trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-0003333010322121-3122312223300132-0221230312011321-3031120011203202-2020200121301130-3030230011200310-3113333031131002-3010201203110311): complete subsection reference.

<a id="canonical-3331120231012020-2232301002022132-3322012100010001-1201021011030001-2130102320002130-2231230302232033-1233033013033230-1231223321031202"></a>

## Next pages — trusted_ca / 102120302221 / 4

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--reference--group-003.md#canonical-0003333010322121-3122312223300132-0221230312011321-3031120011203202-2020200121301130-3030230011200310-3113333031131002-3010201203110311)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0003333010322121-3122312223300132-0221230312011321-3031120011203202-2020200121301130-3030230011200310-3113333031131002-3010201203110311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012311322021200-2303100021210313-2103201323121001-0132100022001021-3320023032230102-0030022302113321-1000032333011131-2201330120033110"></a>

## tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list — trusted_ca_list / 322100001100 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [tls_parameters.common_params](resources--virtual_host--reference--group-003.md#canonical-0333211303112320-2232312331020330-3200001211332113-1332003022321233-0123311023033203-1001103201023313-1200233332212321-3130010111203221)
- [tls_parameters.common_params.validation_params](resources--virtual_host--reference--group-003.md#canonical-1100132303312233-2302331223202323-2232231202112111-2233122222201102-3320121233323003-0022333231032202-1210011013012132-1022023001212123)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3012330211123302-1022033223131232-3113120133013000-0000221203203123-3210202130332330-3210013003001012-2202232013131322-1232222121013122"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321120100131033-0312333221001230-1002232130003132-3213212021022001-2003312301003321-1233120021031121-2320203123103233-1313032022111312"></a>

## Direct properties — trusted_ca_list / 322100001100 / 3

<a id="canonical-3212002232220301-3313010002121021-2003113122131223-2111300310233020-0022220303032011-2230022111021101-0120003113132301-2130223012221233"></a>

<a id="canonical-3201122101002120-1223101323123220-2212023131203301-2213133301101033-1323132311210312-0023223123110000-1313010120300000-1330021331321331"></a>

## kind property — trusted_ca_list / 322100001100 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2313110201103221-0131000112102213-1323231013021300-1022212013303212-0201001231110210-1231003332302301-1131020200010130-2012331321330100"></a>

<a id="canonical-1312331202300222-2000020130310300-1132123030003222-2111003032133202-1321213221001102-2033222312123210-0001332033222002-2231330033211232"></a>

## name property — trusted_ca_list / 322100001100 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1023023233122030-3330233203231330-0132100103311113-3031010113003331-0312120100203022-1201031103120330-0111221121300102-2320213000002233"></a>

<a id="canonical-1312333010220131-3212320132013301-3013220330321131-0100031123010022-3100331301333323-1300330230002313-0121210003001223-3120102211022113"></a>

## namespace property — trusted_ca_list / 322100001100 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-3031210003122301-1210322301302202-0212313003122220-2113312132330000-2020123000222303-1121011311110210-2103232030313003-3001122103023322"></a>

<a id="canonical-2201102120221233-0323110323203133-2103311312103300-3023313221232201-1323021031211011-2232332122331200-0230031031002230-0220302332310003"></a>

## tenant property — trusted_ca_list / 322100001100 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-3323031231011300-2221021132130121-0033000333033011-3100330102112110-1332121030013310-3113232123233002-2023332321103121-1120321321121200"></a>

<a id="canonical-2210120020130232-3100212132311033-2002232223323023-3102303231011123-1021123102121132-1123223012002303-0111313320113033-1313203330202132"></a>

## uid property — trusted_ca_list / 322100001100 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-2123333100113112-0020232331021310-0331202133303323-3130013311200222-3122303102123110-2311212001203331-0132033211100302-1232122120120112"></a>

## Next pages — trusted_ca_list / 322100001100 / 9

- [tls_parameters.common_params.validation_params.trusted_ca](resources--virtual_host--reference--group-003.md#canonical-2121201203020333-2300231002003323-3130133031231300-2232003010111033-3233213301303021-1030121101132021-2100030202322021-2020223201311311)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2213212201113300-0223221321112131-2300013131310300-2330130313033001-3003001223303232-3000310231222000-2101120210302122-2300231020001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130310010210023-0202123312313312-1020002222323113-0031322021012130-2130121200110010-3122222133031212-0102133002333000-1313331213120020"></a>

## tls_parameters.no_client_certificate — no_client_certificate / 300213220332 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- tls_parameters.no_client_certificate

<a id="canonical-2323112302311330-1200313313203222-0323031221031333-3111020121011033-0233231302132131-3303333311000121-0000010300221333-2213020222010031"></a>

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
no_client_certificate = {}
```

<a id="canonical-2322132311001322-0331322212232011-1311100211031311-0120102202010212-1231302122221010-2120121111203321-3313112113111311-2230130201322012"></a>

## Direct properties — no_client_certificate / 300213220332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202302322233203-2222203003221023-3023213301221100-2010013033301012-2200200100231312-2333130210301201-0133313022130012-0322001310111313"></a>

## Next pages — no_client_certificate / 300213220332 / 4

- [tls_parameters](resources--virtual_host--reference--group-003.md#canonical-0331211001213000-2022200002023320-0333201001301110-2110111020031121-0112102012231123-2103133133121332-1220021100301010-3112201011023323)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1012201010031002-2232101323213000-1022110300231232-0230000221021101-1203021223001011-2301313113020120-2002021332022003-0311230002232310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332211313222013-0200232320233113-1131232210213221-0000320120132131-1320210130031332-3120001201330233-0332220210200001-0102303032203231"></a>

## user_identification — user_identification / 210312113211 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- user_identification

<a id="canonical-3123231102032223-3202112001213220-2323200123223203-3121212301113113-1210231333210211-1123211210111022-3223021311311112-0221103332030102"></a>

Type: `"object"`. list nested block, Optional.

Reference to user\_identification object. The rules in the user\_identification object are evaluated
to determine the user identifier to be rate limited.

Upstream description:

A reference to user\_identification object. The rules in the user\_identification object are
evaluated to determine the user identifier to be rate limited.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211001322130020-0223112222323112-0322200123203322-1330222020212232-2003210101001103-3211010300123202-3202310121322220-3211100131333301"></a>

## Direct properties — user_identification / 210312113211 / 3

<a id="canonical-0120222111310222-0320333030121310-1102300020313231-1030130301033323-3221101111030120-1212213333213300-1203123010003020-0211031003300033"></a>

<a id="canonical-1121230102113101-0100010110230200-0101003330311023-0123213003302000-3100310300131001-0332311001010201-2003133321321010-1031133323223032"></a>

## kind property — user_identification / 210312113211 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-2120122131223012-0223301301103212-3223131020130310-0232232310300032-2221002122032323-3232211101011112-0231223121212011-3211100213131103"></a>

<a id="canonical-2032123311021011-2210222233223110-2030330120031102-3323102100330200-0232302032023113-2320021331002123-0001133302313211-0123103000333010"></a>

## name property — user_identification / 210312113211 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0203221000130021-1111003301123201-1103331311111102-2111332110120302-3102113223330131-1213123230022303-1312333332002220-3203320022133102"></a>

<a id="canonical-2232233123333021-1331031302121101-1312120123131302-3023000201010203-3201120033023312-0031210311212130-0102233032233230-0132223220032332"></a>

## namespace property — user_identification / 210312113211 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-1101222132001233-2111020200313002-1001210010013223-0010312130211201-0113130311031323-0311031031232233-2133011202013233-2210300023310220"></a>

<a id="canonical-3132203312330221-1010103220013210-3002113000121023-2032322032130131-2313131321000231-3023313133101102-3000233312103323-2003303213320122"></a>

## tenant property — user_identification / 210312113211 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0322121032330201-3201023101112312-2101102133132312-3332212232010300-2201212200202233-1002222210023032-2021100121233103-0332130232003102"></a>

<a id="canonical-2332132311010313-0313330022031022-1032012320211031-1332030021333020-2031033231130220-3203101131332201-1102311001012100-0011011211311212"></a>

## uid property — user_identification / 210312113211 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1301000220110102-2132310121131131-3332103312200210-3031230220000100-1202103231030033-2111100012300313-1323221310123201-2011011213032102"></a>

## Next pages — user_identification / 210312113211 / 9

- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110211020000103-1220010130012003-0021003033221213-0001002220132320-1001303212300320-2201303103222030-1332102333211010-2010020330120322"></a>

## waf_type — waf_type / 322112120232 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- waf_type

<a id="canonical-3230013003023000-1033211112313122-3023101303102010-3313122122303000-1202210222021131-1031330201322010-0303332312111121-3302321220210320"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
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
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302323100102312-0313303103323332-1200311030013021-0113233131110100-1203313031001000-3020103230100222-3021212220122233-0101112100033233"></a>

## Direct properties — waf_type / 322112120232 / 3

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031): complete subsection reference.

- [disable_waf](resources--virtual_host--reference--group-003.md#canonical-0200013230100212-1300201000323321-0100110000112200-1102021123321212-1003031221312331-3213322010301212-3122122332032230-2231003012221103): complete subsection reference.

- [inherit_waf](resources--virtual_host--reference--group-003.md#canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200): complete subsection reference.

<a id="canonical-0122101321332112-3022021012231323-2001203230300103-1120313103233233-0102021232232130-0021330311211212-1101022013010322-1331212020310011"></a>

## Next pages — waf_type / 322112120232 / 4

- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031)
- [waf_type.disable_waf](resources--virtual_host--reference--group-003.md#canonical-0200013230100212-1300201000323321-0100110000112200-1102021123321212-1003031221312331-3213322010301212-3122122332032230-2231003012221103)
- [waf_type.inherit_waf](resources--virtual_host--reference--group-003.md#canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123120203232203-2210223222021301-1012023303112211-0210320321301031-0321022130322213-0333201023100002-0011033330020111-2023211002211332"></a>

## waf_type.app_firewall — app_firewall / 012211133033 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.app_firewall

<a id="canonical-3213022012121222-0002330131002030-3310033233203333-1033330211110330-2120011000001330-2201021020201233-0332123000320320-2212020330301333"></a>

Type: `"object"`. single nested block, Optional.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221222303212113-3121013113013133-3233120000020113-3011202111030330-2310303121321332-0222312321021130-2311120111120313-2133312210133103"></a>

## Direct properties — app_firewall / 012211133033 / 3

- [app_firewall](resources--virtual_host--reference--group-003.md#canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203): complete subsection reference.

<a id="canonical-1121102133231231-1313013120113003-1023212333323321-2330312100322003-2032212010023022-0312130303000222-3300321331112210-0322013203030222"></a>

## Next pages — app_firewall / 012211133033 / 4

- [waf_type.app_firewall.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-3002131032323211-1313323003223001-1030222330120001-3220101010122031-0213113222211303-3312312200323101-1132121102031200-3312230200131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211133231112321-0023232023013213-0003300030311201-2133301310033213-0300111113123023-3223212030320322-2233211000220001-0200300223222222"></a>

## waf_type.app_firewall.app_firewall — app_firewall / 031230322302 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031)
- waf_type.app_firewall.app_firewall

<a id="canonical-2123223303123331-0321031133303331-2022221030110212-0301330033003311-1012333203333123-3230110301020030-1221011102120001-0210201321321333"></a>

Type: `"object"`. list nested block, Optional.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202330021332300-2021330221011131-1312012033001001-1313113022331131-2113203223003011-3221032011221231-1231131230022213-0110000021032112"></a>

## Direct properties — app_firewall / 031230322302 / 3

<a id="canonical-0010032012013331-0303102231333212-3201133131010023-3133001021110200-2202200133000320-0130230102311131-1323223223130000-2101322301322200"></a>

<a id="canonical-1331033322311300-0102023130121032-0223211231033011-1221200121310123-3010300220323230-3311313013232112-1220320213323303-2222112221110232"></a>

## kind property — app_firewall / 031230322302 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1200031333031012-0023302213113102-1121212310312203-3310320333231022-0321110113222010-2030302020120332-1303232230333001-1010322323123100"></a>

<a id="canonical-0123232222000021-1101213101210211-3020211313231100-1202131101203221-1111223031200311-3103011131223222-2320331003331033-1113010233023200"></a>

## name property — app_firewall / 031230322302 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2031123111321331-3332320323132012-3232332013331211-1212331111033220-0003212320322001-1121312120331231-1222212232220321-2313321203022032"></a>

<a id="canonical-1322030031123122-3101301231102223-3123113011120003-3110030103013133-3232301220002010-1003033301020311-3010230333222332-3312322033122331"></a>

## namespace property — app_firewall / 031230322302 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-3101322222023123-0001313303001031-0033203333000220-0003233001000020-2212202003310100-1312320202023303-1201312112203202-0131222003332121"></a>

<a id="canonical-1100120113203021-1200310010011021-1101213011203022-3322031222003110-0123230312332223-3111230022102122-2122333002222011-1320111323031012"></a>

## tenant property — app_firewall / 031230322302 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-1123303212111101-0231123003302020-1031321303000110-3122101313300322-0030111332130222-1221123331013331-2012120220012013-3332131100221002"></a>

<a id="canonical-3103133110320023-0210131022132331-2222002031231013-0103233001302300-3321302323220312-1310231133011021-0121232231003110-3202021331200232"></a>

## uid property — app_firewall / 031230322302 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-3201102332111102-1200103111210220-2113212033012112-1303030211320333-0310231132312211-3202233300221230-3330233032310112-2303100103020322"></a>

## Next pages — app_firewall / 031230322302 / 9

- [waf_type.app_firewall](resources--virtual_host--reference--group-003.md#canonical-3013121233031313-3131021232212222-3222231213232301-1333021332220013-1002103010033211-0110331022300302-1330321011210331-2131112331011031)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-0200013230100212-1300201000323321-0100110000112200-1102021123321212-1003031221312331-3213322010301212-3122122332032230-2231003012221103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230033210312010-0103003103300103-1020310132133220-1020001221010003-1013001311320221-2331103230020331-2212111332232112-0301232132121201"></a>

## waf_type.disable_waf — disable_waf / 103200122213 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.disable_waf

<a id="canonical-0302310212030013-0311333130001103-0122210310011201-1201023311102112-0230033112133120-3001212100112231-3123023021221320-0220210212310210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

<a id="canonical-0303201113233103-2202323321232032-2121031300031323-0131113031123222-1001121020300200-3120022031311321-2012230230220022-0020011323330113"></a>

## Direct properties — disable_waf / 103200122213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031023022130122-0221110232233111-2030003133003231-2121112102131121-0023313212021032-2000021312010322-0023303300130012-2000222333232121"></a>

## Next pages — disable_waf / 103200122213 / 4

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)

<a id="canonical-1310320202133212-2221033303020302-0011200030000030-2010112132322220-1232023133223203-1113301212012300-1132200001133223-1303021333223200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300301003313120-3332213212022112-0323011133300231-2032210202002130-1321310310021213-0332313232110222-3130223123022102-0233032230113030"></a>

## waf_type.inherit_waf — inherit_waf / 312131220033 / 2

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
- [Property reference](resources--virtual_host--reference--group-001.md#canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220)
- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- waf_type.inherit_waf

<a id="canonical-0011131120230110-2322330020322312-1011301301130303-1223322210131003-0221300000132021-1302213300322220-0231320323010231-2321222020130331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

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
inherit_waf = {}
```

<a id="canonical-3202003320112033-2312100031102033-0101020010130020-1233221013201003-1011132103002030-3322011010121113-2213123220110002-2101211322323301"></a>

## Direct properties — inherit_waf / 312131220033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331011232000303-3020003111101303-0223321030020213-2032312233013120-3011133310220333-2100001222230102-3211303322332102-1111333302113021"></a>

## Next pages — inherit_waf / 312131220033 / 4

- [waf_type](resources--virtual_host--reference--group-003.md#canonical-2322100101202101-2331131111122123-0200313013302211-3200032010130132-1310103000220023-2103022322122223-2311300301002030-1022212031023331)
- [xcsh_virtual_host](../resources/virtual_host.md#canonical-3223131132111310-0302331213231313-3010210001002120-2000312103313323-2221222232122213-2203113131033213-2103120313003230-3112232203320222)
