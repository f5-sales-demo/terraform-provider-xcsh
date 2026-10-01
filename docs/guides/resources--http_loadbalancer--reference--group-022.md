---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1313230223121022-0122310233030320-2123322112321320-0323311130031133-0032021123303233-2010331321203300-0123012232011032-0110011203313230"></a>

## more_option.response_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 323012010223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2200200122310020-0110310232232021-1222333032121113-2132113310203121-3231301322122303-2021200233110202-0132211223101330-3021123303122012"></a>

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

<a id="canonical-0323101230210203-2101132033331112-3023012200310231-2022303213320022-0023201121321231-3322130300103221-2010111331202112-2201130020002131"></a>

## Direct properties — clear_secret_info / 323012010223 / 3

<a id="canonical-0103101013031023-2223221001221333-2320310332223023-0310103113213321-0001132023021032-0012000302101220-1110213011222113-1202223013221311"></a>

<a id="canonical-3133031321111101-2303201302311323-0023321331221220-3201313002033202-2022313333222002-3220201110303220-0011320230222010-1033223303101233"></a>

## provider_ref property — clear_secret_info / 323012010223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3021311113323321-0302310310303310-3120333223231113-3310133230023021-0100100332123212-2122201132300121-1302123311221211-3302030223320122"></a>

<a id="canonical-0222330111233223-2200001031000320-1212132302203023-1131333002101102-2300030310011122-0023120001330010-2333001330300031-3230301312331302"></a>

## URL property — clear_secret_info / 323012010223 / 5

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

<a id="canonical-0212333221103310-0022001100130130-2133031312011233-3330222011220021-1121320130231101-1330033131111112-3031011301123101-0231301300212133"></a>

## Next pages — clear_secret_info / 323012010223 / 6

- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302222212011212-2032201101300213-2103130300212131-0310331212201000-1102112221303310-3131033133131103-0031201133121111-1020311022333002"></a>

## more_option.response_headers_to_add — response_headers_to_add / 131233021332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.response_headers_to_add

<a id="canonical-1301103200200321-0000011321231201-3311302210031330-2023130120232313-3330022223021122-3120120123122323-3031001001031312-1330002332211003"></a>

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

<a id="canonical-3132201233221033-2223303322023230-3133230212100112-1312013311231000-0221022302131033-3201030311302011-0122223132101001-2132222330123133"></a>

## Direct properties — response_headers_to_add / 131233021332 / 3

<a id="canonical-0203310022313333-2203022002300223-0323300122003123-2301111133311122-2200103322330230-3331132133323123-0221000323101013-1100202200010023"></a>

<a id="canonical-2001012103120310-3230023031130223-3101131321022223-0001223011203110-2013100203203220-1100333130103033-0101120020020120-0121103332302200"></a>

## append property — response_headers_to_add / 131233021332 / 4

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

<a id="canonical-3013113212021310-2110330122331033-2010003233130013-2033121123123012-0303312331030102-3021012112212011-2310231110102231-1310121110200321"></a>

<a id="canonical-2330222031003323-0311020101231122-1012303023132320-2113103012330211-1032213110013220-3320113312320323-2202303231222103-0113122232013212"></a>

## name property — response_headers_to_add / 131233021332 / 5

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

- [secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231): complete subsection reference.

<a id="canonical-0121233013012010-1022322011320020-3212313112132130-2111210020030302-1012223320220102-3322013023202033-2310103130331101-1110231233321231"></a>

<a id="canonical-3101000200111323-2322330111120231-3011003132131013-2203332120332033-2130200331133313-3012111121233110-3231110132321001-0033113031322210"></a>

## value property — response_headers_to_add / 131233021332 / 6

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

<a id="canonical-3311020010220130-1021000120221330-0303213232230311-3201030313020100-3100321232312120-0201303230200022-3233123310312032-1302021012033121"></a>

## Next pages — response_headers_to_add / 131233021332 / 7

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313120003102231-3222103300332101-1101100310300302-1202031332013111-2322333221322233-0310330231033021-0301030322322132-3113301213022022"></a>

## more_option.response_headers_to_add.secret_value — secret_value / 102330303222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- more_option.response_headers_to_add.secret_value

<a id="canonical-1131232202121221-3231323222231302-2103322102032323-0133333312313200-2323132212030032-2332330130303111-0220021002023311-2100301332323122"></a>

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

<a id="canonical-0120011330012121-3123331212312133-3000333233113133-2030202312011032-2220130222033213-3020212323133300-3131331201113310-1012103332033223"></a>

## Direct properties — secret_value / 102330303222 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3030120032231000-2100231032120312-1320230213332233-2212100020132121-3221000312202102-1022120113330231-1330223333003232-2321101101211320): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3021201211021111-2000322133100203-1322031332223201-0321212023100121-3232211031301120-3100222010222211-2210332102032010-0321331313321322): complete subsection reference.

<a id="canonical-2231122311130100-3201311302131003-1120031233302332-1332303101031211-3102122222030123-1012330302022232-0010022132103331-2303222103210013"></a>

## Next pages — secret_value / 102330303222 / 4

- [more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3030120032231000-2100231032120312-1320230213332233-2212100020132121-3221000312202102-1022120113330231-1330223333003232-2321101101211320)
- [more_option.response_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-022.md#canonical-3021201211021111-2000322133100203-1322031332223201-0321212023100121-3232211031301120-3100222010222211-2210332102032010-0321331313321322)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3030120032231000-2100231032120312-1320230213332233-2212100020132121-3221000312202102-1022120113330231-1330223333003232-2321101101211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113222102120003-0001131131233002-2223211033011230-1011230321333222-0201302300313002-1022131333313010-2013223131122312-0233101333230202"></a>

## more_option.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 101102101002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-2202300123231233-0232013101031001-3123212121000011-0020131022111121-2203302300013212-0122023313132112-2313121123111223-1321230322131013"></a>

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

<a id="canonical-2230223200133113-2123310303103211-0333011201130103-2301223021332000-3331100323123332-3200300311030312-0020212213211020-0113212330113110"></a>

## Direct properties — blindfold_secret_info / 101102101002 / 3

<a id="canonical-0310231002322303-1003221333031200-0113132313111300-0112030302100021-1133000103310101-1111122321011333-0212002011031103-3213232130312203"></a>

<a id="canonical-1131111313200202-0132301032033231-2220232221302012-0321130031131202-1101102300333033-1101310220200103-1133002310332211-3223130212222200"></a>

## decryption_provider property — blindfold_secret_info / 101102101002 / 4

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

<a id="canonical-2213311001122111-2331321303030202-2201330000103321-0121311012320330-2131220123132011-1112001123232023-0213033031221330-3103213101021110"></a>

<a id="canonical-2032030122100110-1302132113013102-1303013110122331-3313232132132122-0302303030111003-3010131211012202-1202102110313001-3110331321301033"></a>

## location property — blindfold_secret_info / 101102101002 / 5

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

<a id="canonical-3333103000232120-1311231301113020-2133203001120212-3211011232233301-3113333110303232-2313200323022232-2002200312121212-0320101011330211"></a>

<a id="canonical-3102113102031021-1200020202103101-0331102003111303-2332212201323012-3011322312030311-2302122331213220-2102201000222131-0033001321323323"></a>

## store_provider property — blindfold_secret_info / 101102101002 / 6

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

<a id="canonical-2022322213203121-2130321301131003-1313021320100211-3113323103010311-1321213312222202-1120130211110310-0222123001302100-3121001313003031"></a>

## Next pages — blindfold_secret_info / 101102101002 / 7

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021201211021111-2000322133100203-1322031332223201-0321212023100121-3232211031301120-3100222010222211-2210332102032010-0321331313321322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111310030221113-3000103300023033-3200212202030300-3232332223210330-2002122013230113-3220020003212122-2021232301000102-1213033212020223"></a>

## more_option.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 020320012122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0122121332321113-3332021300302123-0211022203233203-0230202310210031-1231131213011003-1211031133000202-1131013100030221-0022323321311100"></a>

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

<a id="canonical-3012201100122031-0202030132030011-1311112030013112-1032020323232031-3011333012132102-0011003302103221-0231221021313330-3221211220302011"></a>

## Direct properties — clear_secret_info / 020320012122 / 3

<a id="canonical-1113010312122311-2122030130220232-3331320100200100-3312033111322101-0202211322200012-1131003033120020-3010122010111000-1220020332312103"></a>

<a id="canonical-1323000001133120-3102013032230200-3311200311321201-1320110010032320-0231302221020002-2130310231230133-0303133001311212-1133103313300133"></a>

## provider_ref property — clear_secret_info / 020320012122 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3212013113003213-0022020202012330-2202312201331313-2002020121101202-1332033123102003-1010122003200113-0221301223301213-2201023112233103"></a>

<a id="canonical-1220112113321013-3130012201113213-3231013121012111-0200133102132033-3213120232102300-0110322221033200-1102002023313130-0002032101021200"></a>

## URL property — clear_secret_info / 020320012122 / 5

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

<a id="canonical-1120131303302323-1201130020012123-2010120320012100-0212011123331211-3312330203000322-0312211112333303-3232100302221223-3323310131310121"></a>

## Next pages — clear_secret_info / 020320012122 / 6

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-022.md#canonical-3311312100303221-1130212020103202-1012101323003311-3031222011123001-2020202323010222-0011323121001130-2011010111321232-0131300232130231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0022320220230323-0223203133103320-3032023020320010-3000011312022000-3332213021010220-3323011033022102-3132320201321311-0310230130132011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112030220010331-3113023012301012-1011100321232003-0122210331030211-2310130110213330-2200120102321330-3000011033313103-0203121312120033"></a>

## multi_lb_app — multi_lb_app / 113210311022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- multi_lb_app

<a id="canonical-1113010133303031-1030321102031003-3030322233333020-0132110322111130-3302030002322023-3231011110103303-3222330333330311-1020023031330321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

- [multi_lb_app](resources--http_loadbalancer--reference--group-022.md#canonical-1113010133303031-1030321102031003-3030322233333020-0132110322111130-3302030002322023-3231011110103303-3222330333330311-1020023031330321)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-1201323130312033-1332212010203222-1112131130000000-3313131233123132-0112201202300002-2122020103032101-3012032203220001-1123233022300133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
multi_lb_app = {}
```

<a id="canonical-2230303312202200-2003102220032222-3320103112210202-0322302331211202-2012030323222130-0032013030012332-0203203011020310-1121221112200003"></a>

## Direct properties — multi_lb_app / 113210311022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031123132020033-2210131301013130-3010130211211031-3031112022132113-2132321333212223-0131232000002131-2310301332221211-0022022130010221"></a>

## Next pages — multi_lb_app / 113210311022 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101121220322123-0303121132222013-1232012001122310-2310001013023221-1200112133200300-0332032203023122-3310012322111310-2011133110323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100233322201033-2210122033130312-3101022132311231-1011233300031223-3221200023200320-2000111222132200-2032303221022123-1303232132333032"></a>

## no_challenge — no_challenge / 012313330233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- no_challenge

<a id="canonical-3132021103233010-0102003303001330-3110103203321112-0211303310001001-0000320020121320-3033100021101033-2220023223112313-3310221213303023"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no challenge. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_challenge = {}
```

<a id="canonical-3200231320300303-2011221113023121-3200023100023110-3232232323233010-3000121000111013-1002232310330030-3301002012222102-1011331112130003"></a>

## Direct properties — no_challenge / 012313330233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330203030100032-3210030133103311-2012102121321223-3012233303023330-3120101312112100-2223313100320132-3031202202022130-2232103031101332"></a>

## Next pages — no_challenge / 012313330233 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1213221123121332-0130210312020112-0030010203112202-1320321321320032-1222212130023330-2011211132030330-2231013002131030-0013313003020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003032333130220-1001032312202311-0201002200223310-3322011323002120-1023231220031312-2001102011113012-2112130003030030-2002121202110333"></a>

## no_service_policies — no_service_policies / 331010321131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- no_service_policies

<a id="canonical-1101022221112332-2030130211022031-1133103101333213-1101003230013311-3103302223310030-3213132312012201-0000201122030113-1223213310313002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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
no_service_policies = {}
```

<a id="canonical-1302300203211332-0120203333022033-2200133122122123-1213223030003031-0011330113300020-2230111100103133-2302001313311011-0331122011223033"></a>

## Direct properties — no_service_policies / 331010321131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102210001002300-1022102132322131-0102110300200132-0202212310212220-3300331011300323-0012300013122212-1032123223110122-3233000223223320"></a>

## Next pages — no_service_policies / 331010321131 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320132321221023-1012033003330010-1002032213320131-1311333132022102-3032103112322320-0123110222201011-3332201302333111-1202100021121321"></a>

## origin_server_subset_rule_list — origin_server_subset_rule_list / 122323313300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- origin_server_subset_rule_list

<a id="canonical-3011031302203312-1320300230222133-2210031101333001-3012120333010320-1021301101102203-3232213320210222-0233230312223211-3020200330331313"></a>

Type: `"object"`. single nested block, Optional.

Origin Server Subset Rule List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
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
origin_server_subset_rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203030030201003-0223313011300130-2112323320110112-1233121332133002-0301113113002020-3221312332210012-2022203021233202-0113310021331102"></a>

## Direct properties — origin_server_subset_rule_list / 122323313300 / 3

- [origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210): complete subsection reference.

<a id="canonical-3020222330210301-1230101203301203-3113011122013111-3211233222330100-3300101122131331-0313033030323231-1011330233032200-0010212031011210"></a>

## Next pages — origin_server_subset_rule_list / 122323313300 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212302032011001-1323310322311331-1222133113021222-3102313012223113-2001220112021220-0001213121120103-3332021120300022-0333010133220331"></a>

## origin_server_subset_rule_list.origin_server_subset_rules — origin_server_subset_rules / 003311101100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="canonical-1213300313322030-1003111033010333-3131211020312030-3221030020022012-3233021021232302-2230010311030323-2112103131133320-1010110322311213"></a>

Type: `"object"`. list nested block, Optional.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to..

Upstream description:

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("origin_server_subsets_action"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("client_selector",
    "none"),
  validators.ConflictingListObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
origin_server_subset_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020033113003231-3101231201031230-2010312113232221-2301001223130102-3333110120223203-2330330323032010-1130233002131231-2023122220123133"></a>

## Direct properties — origin_server_subset_rules / 003311101100 / 3

- [any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-1302111220302030-3001330021000231-3112112223302013-2220311032232231-2020101312020033-2310233211031132-0131223203311131-1222001331101321): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-2210123021022011-3220321020322313-1232200022032120-0002100200013131-1202331212220302-3210312031113231-0320130320222102-0120112330131132): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-3222323023321133-3012102203110231-3232333031113211-1203011322120130-2013302110020210-2230233122312101-0311121100023013-1331111002121103): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-022.md#canonical-3212132133222002-2131213233133333-0121323232210133-0023102330231222-3110220221030033-1321101302100320-2212002031310232-2021123112120111): complete subsection reference.

<a id="canonical-2211011233311303-1332330221112122-3021130201122320-0102131002101030-2333322100030033-3320121100111310-1210211221300030-2020103311213320"></a>

<a id="canonical-2110112033203022-0032210333102222-1220130030122203-0322021201333101-2110122103101020-1010102200232000-1021101021103212-1223301201000101"></a>

## country_codes property — origin_server_subset_rules / 003311101100 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

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

- [ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-022.md#canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0020222103110101-2120130130101001-2303101303303210-2203023310303011-0101200212201213-1310123333103230-3133121312302332-2300113211122131): complete subsection reference.

- [none](resources--http_loadbalancer--reference--group-022.md#canonical-2211230323332300-3330211101102100-2023012322001003-1112202131233213-2321022211100031-0100321221002031-1333310000231113-0102002232313223): complete subsection reference.

<a id="canonical-1331010112333200-0320033011322030-2202133222113201-1112211202033001-0031031022113003-1213332232330312-0021111131300112-0332222212232031"></a>

<a id="canonical-2322200001110333-0213120121000021-2030232201323302-0230302321300201-0030102131203020-0100313213211023-1123300031132330-1102133123302213"></a>

## origin_server_subsets_action property — origin_server_subset_rules / 003311101100 / 5

Type: `["map", "string"]`. Optional.

Add labels to select one or more origin servers.

Upstream description:

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3320220033301301-0000332222110111-3023100203013113-1111020331203121-3130020233100300-3022330130023113-1011312010210033-2000223331211331"></a>

<a id="canonical-0302130013131330-1111323311031120-2210022111221300-3212003220223322-1122312012322023-2210312310320212-3102121221220123-0231303330000220"></a>

## re_name_list property — origin_server_subset_rules / 003311101100 / 6

Type: `["list", "string"]`. Optional.

RE Names. List of RE names for match.

Upstream description:

List of RE names for match.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030112103233102-0233330122122101-1102333323120012-1202303332133031-3022303110200231-0313112032212303-1323023001001021-3002313210323012"></a>

## Next pages — origin_server_subset_rules / 003311101100 / 7

- [origin_server_subset_rule_list.origin_server_subset_rules.any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-1302111220302030-3001330021000231-3112112223302013-2220311032232231-2020101312020033-2310233211031132-0131223203311131-1222001331101321)
- [origin_server_subset_rule_list.origin_server_subset_rules.any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-2210123021022011-3220321020322313-1232200022032120-0002100200013131-1202331212220302-3210312031113231-0320130320222102-0120112330131132)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-3222323023321133-3012102203110231-3232333031113211-1203011322120130-2013302110020210-2230233122312101-0311121100023013-1331111002121103)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213)
- [origin_server_subset_rule_list.origin_server_subset_rules.client_selector](resources--http_loadbalancer--reference--group-022.md#canonical-3212132133222002-2131213233133333-0121323232210133-0023102330231222-3110220221030033-1321101302100320-2212002031310232-2021123112120111)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list](resources--http_loadbalancer--reference--group-022.md#canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301)
- [origin_server_subset_rule_list.origin_server_subset_rules.metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0020222103110101-2120130130101001-2303101303303210-2203023310303011-0101200212201213-1310123333103230-3133121312302332-2300113211122131)
- [origin_server_subset_rule_list.origin_server_subset_rules.none](resources--http_loadbalancer--reference--group-022.md#canonical-2211230323332300-3330211101102100-2023012322001003-1112202131233213-2321022211100031-0100321221002031-1333310000231113-0102002232313223)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1302111220302030-3001330021000231-3112112223302013-2220311032232231-2020101312020033-2310233211031132-0131223203311131-1222001331101321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232100213303201-3121120021123330-1222133100200103-0102333222103011-2113131230222323-0000130230212133-0231020110002233-1121013223321302"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_asn — any_asn / 333110312220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.any_asn

<a id="canonical-3211310331033002-1022231030332230-3023111101120210-0030031322010331-3010013212100011-3211010123033013-0133020310031010-0013201323010112"></a>

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
any_asn = {}
```

<a id="canonical-1313031333030123-3321200003100131-1132221120221100-0110322310322113-2103030113320011-1311030313030100-0112023131331232-2300202212110323"></a>

## Direct properties — any_asn / 333110312220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032333020231131-3132032010200113-1310320123223013-0223201123202232-3120031223032012-3000110201021110-1030220021230303-1002133233330212"></a>

## Next pages — any_asn / 333110312220 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2210123021022011-3220321020322313-1232200022032120-0002100200013131-1202331212220302-3210312031113231-0320130320222102-0120112330131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310331132021212-3231102201211210-1333013010310020-3302123032002022-0031000332023013-0110213203112030-1233010213033132-0313131011022032"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_ip — any_ip / 021033313001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.any_ip

<a id="canonical-3113211132333302-1212020311323010-0001213121233100-2031202201112121-3212331323010321-0111222131031033-0330100332220202-1001020322331333"></a>

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
any_ip = {}
```

<a id="canonical-3301331323022122-0031322133133103-0002123320021003-3213102120032120-0131301111103032-1210220301303132-0202212231300300-1233021111111303"></a>

## Direct properties — any_ip / 021033313001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012303102022211-0300123030030031-3312033012330233-2220010231331232-2301003331130023-3333010102122313-3032022113112200-2201232102033333"></a>

## Next pages — any_ip / 021033313001 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3222323023321133-3012102203110231-3232333031113211-1203011322120130-2013302110020210-2230233122312101-0311121100023013-1331111002121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202123123021111-2301332013032001-0222101302002103-2203303002021320-1322310030331003-0230012002002030-0232010332303121-3221303210110223"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_list — asn_list / 300222031201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="canonical-3111011320213333-2033230120210312-2311021333320203-3223220222320300-0201202323113212-3132011231120223-2003133021303123-3030111103033211"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033000011333313-3123321111011320-1013231303133333-2031301232220012-1121022020110121-1301103103120111-0313031232022131-2212330222020222"></a>

## Direct properties — asn_list / 300222031201 / 3

<a id="canonical-2332221012013123-3232312030103201-1011331003013103-1301103010320320-0121021021112303-3223320013322010-0322323111012212-2223130102123222"></a>

<a id="canonical-2201130313033210-3013002102233320-3333310131221021-2132322012302031-0113211220202331-3232231131320322-3210202321020103-3133103121030103"></a>

## as_numbers property — asn_list / 300222031201 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0311220312013331-2020303111321023-1031320013202230-2033103220031032-2121321203232103-0333110220023032-3010103020110202-3002223220323133"></a>

## Next pages — asn_list / 300222031201 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130122302102213-0213321300003013-1012221310012100-2321002033021133-3310003112323101-2332133223312330-3033323210110322-1022022211221333"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher — asn_matcher / 133312133022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="canonical-1122010312123323-0220323031202020-2032021011320120-0132333012232321-3312223000122321-2112010100210331-0030122000210132-3232230012222232"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121200313022031-3101101233333133-0130131000331131-0003301031301021-1331100132213232-2000231323001013-1020021213203231-3320311320130131"></a>

## Direct properties — asn_matcher / 133312133022 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-2133331202111302-1223330200330231-1303003211331110-0302111201302110-3220213222231011-1023012032011322-2311313300133130-1011321011233023): complete subsection reference.

<a id="canonical-3000302133011233-2212223302003120-3030110121022302-3321233200000003-2030000302220023-0102011122002333-1100122220013230-0322303332231121"></a>

## Next pages — asn_matcher / 133312133022 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-2133331202111302-1223330200330231-1303003211331110-0302111201302110-3220213222231011-1023012032011322-2311313300133130-1011321011233023)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2133331202111302-1223330200330231-1303003211331110-0302111201302110-3220213222231011-1023012032011322-2311313300133130-1011321011233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031220202303023-2033011020213322-2120310112022011-3203013302111030-1311310201213110-1130033230132102-1022101031100013-2322201223112221"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets — asn_sets / 201110201000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets

<a id="canonical-0031332100011201-1210220132310232-0131021330110001-3112210331002003-2110200200121133-2333121300020131-0011003210032333-3131303231232310"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212033212101113-1212331003321133-2003233021203031-2220211230301222-2011020200031210-0320110023013203-3131221023111213-2123233131323030"></a>

## Direct properties — asn_sets / 201110201000 / 3

<a id="canonical-0021001123330013-2133102310013111-3321302321023010-3223012033332123-1321001321023313-2011221111302130-3313101222133111-1210302100310132"></a>

<a id="canonical-1200330020301022-1030332220023203-2002113323121031-2130111113102012-3110031212311301-1220303202330033-1002031331310303-0313301020313312"></a>

## kind property — asn_sets / 201110201000 / 4

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

<a id="canonical-2033201210130112-0122231311332221-1230323110032330-0003210331002110-2112131310311313-0133222233002200-3113201013323331-1331002300300121"></a>

<a id="canonical-2001002031123211-3002112101102003-3323130321223330-1111100010023322-0210301320223132-0333313321322312-0303020102301012-0232103312210003"></a>

## name property — asn_sets / 201110201000 / 5

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

<a id="canonical-3121133031202300-0033331000331130-0200213201010013-2110112300102023-2013320223333332-2110012103300120-2202132011011312-1110031233220012"></a>

<a id="canonical-2102030033012332-2202132233101310-3221311213220000-0321032211232123-1303220313231322-2000121101032300-3201010223120311-3201111033203321"></a>

## namespace property — asn_sets / 201110201000 / 6

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
  }
}
```

<a id="canonical-0201113112311211-2100011012111222-2103301000222023-0312103312002212-2332122301103002-0100032002222120-2022013331112103-1211333300321232"></a>

<a id="canonical-2111332221313221-2000213013020303-0012103120030311-0223303031212230-2302113210320332-2322320130233000-3303131132323311-3010331011002201"></a>

## tenant property — asn_sets / 201110201000 / 7

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

<a id="canonical-0121033002101020-0020111221211331-2013123232013003-0022002233201230-0212023133110020-2332232303321302-0332012103202200-3030131021031312"></a>

<a id="canonical-3030103033120121-1031333201112233-1000102000221322-2323020310321220-0321322101123101-0221131123132031-0300011220302330-0303300322122331"></a>

## uid property — asn_sets / 201110201000 / 8

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

<a id="canonical-0102000113320312-2312211012300120-0323021330122111-0022222100320332-0023133300232220-3210212213220310-3333212122003131-2210121011303031"></a>

## Next pages — asn_sets / 201110201000 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2131232131220022-3231312032001311-1001123321021221-1010012233233102-3032020321321013-2003031121101100-2013021031132131-2032111122221213)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3212132133222002-2131213233133333-0121323232210133-0023102330231222-3110220221030033-1321101302100320-2212002031310232-2021123112120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133000201210300-1013302030222023-0201022332230011-2020302021210133-2001301113011123-1010220110133220-3002300322123000-1330100311122100"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.client_selector — client_selector / 133130102213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-3312322003112231-3332221132212021-1100221223100220-2231012323101203-0323101112003001-1111012112013203-0021300202210111-1222230212000221"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331313012130221-1103100032331033-2320222121231223-3031011021220112-0013200122301332-3123213103013012-3301321000033332-3001203310332232"></a>

## Direct properties — client_selector / 133130102213 / 3

<a id="canonical-0003103120121013-1203102003121331-0120102031013121-1323103113110222-0323231123130010-3310302120023100-1012323123123331-0210002302212312"></a>

<a id="canonical-0211300223310222-2222202231203303-2221101032010300-2323233212223003-2303102103023032-2220203222032133-2302033212022130-0231000330111102"></a>

## expressions property — client_selector / 133130102213 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0110331033030023-3130100233201302-2201300021010310-3223123122233322-3331133021110231-2220022013212020-3102130010210110-2031133123033123"></a>

## Next pages — client_selector / 133130102213 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212101232021120-2333122132132322-1012103212230002-1112222203200333-2220000130231313-0220023020203011-3111212102032030-0210321233212122"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher — ip_matcher / 302011021233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-0111302330233122-0031221011331101-0331003300113023-2003100302003102-2210130202223001-3022130203330232-0002332313103022-0230312231331303"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202331010220020-0021023310001200-3022231232311222-0032202022212233-0100313000003101-3123103032312131-2222033213223320-0211021013120213"></a>

## Direct properties — ip_matcher / 302011021233 / 3

<a id="canonical-1331310031213312-3130221021100003-1202200000313121-3021123311313023-2000333103312112-1220000211121232-0031013023112100-0312222033221103"></a>

<a id="canonical-0132211232333221-3000222203032223-0101131330323210-1123222333200312-3020000233111133-1132110222032032-2211000300220002-2223311030203122"></a>

## invert_matcher property — ip_matcher / 302011021233 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](resources--http_loadbalancer--reference--group-022.md#canonical-1002210121103330-0232111131223301-2202012213202122-1212001301000302-3312010302120311-1322203332231203-2032221300121331-2332312222132221): complete subsection reference.

<a id="canonical-2331211020213031-1133032210001122-2033232003201113-0312132231232312-1013003112001032-1302210100103133-2231101103320020-2323002100332101"></a>

## Next pages — ip_matcher / 302011021233 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-022.md#canonical-1002210121103330-0232111131223301-2202012213202122-1212001301000302-3312010302120311-1322203332231203-2032221300121331-2332312222132221)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002210121103330-0232111131223301-2202012213202122-1212001301000302-3312010302120311-1322203332231203-2032221300121331-2332312222132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031321222231000-1301212131131132-3220232302011113-2323101221220302-3113223233132220-1312023132230233-3213332320103230-3323000101111131"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets — prefix_sets / 301033101200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-3000103012213031-0000013032122121-1030202112101223-1012030102320210-2001232230210223-1202232203101330-1002031200320301-3200033103310302"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300220331100330-1222333332202302-0100031003112023-3030312210333112-1331321102233230-2331110332211313-1312011020010303-0132323112011011"></a>

## Direct properties — prefix_sets / 301033101200 / 3

<a id="canonical-0001210321320320-0231220212332121-2232222103003131-3033330302331111-2033010200230331-3100233031220002-1300123103313120-2001220201122031"></a>

<a id="canonical-2200311223221003-1301331131133212-0030132112222211-2201013101312210-0020333021003112-3131021233121112-2032011230213011-3303330322021203"></a>

## kind property — prefix_sets / 301033101200 / 4

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

<a id="canonical-1130321112020323-3212312011233033-2300102232001311-3100333131313000-2332211022311311-3332031102101113-2320011102031130-0023022231022302"></a>

<a id="canonical-2033320323220033-1203230203233331-0030110221121212-3331220311131223-1133100213132332-0201302030231301-1132021313232003-0320132220000102"></a>

## name property — prefix_sets / 301033101200 / 5

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

<a id="canonical-1213331223320121-3302133002020132-1102333100233211-0030131000003011-2321133212120210-1010101333222000-1203323023021103-1211230011331033"></a>

<a id="canonical-1303132233322000-1332102311211320-1201103332300003-3231000023212103-3132002301000130-3210020221001220-3020222102310030-0033131231231313"></a>

## namespace property — prefix_sets / 301033101200 / 6

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
  }
}
```

<a id="canonical-1322201323013313-2213320300322300-2301100202123102-1123301302200220-0133222133213321-0111302300112302-2233001320013322-3303310112003201"></a>

<a id="canonical-3330202123210112-2233320110132122-0232002011222020-1331110010300203-0302301101211003-2121103210333101-2322022311211233-1322010132331301"></a>

## tenant property — prefix_sets / 301033101200 / 7

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

<a id="canonical-2110100310223230-0121220011233012-0131303012231222-2002101302333201-2220230123132303-3330232113033302-2103122303321111-0212112112220201"></a>

<a id="canonical-1310231332031202-1323101331210322-3331311312302231-1020032333330032-1013033003300322-0103330112231032-1100320301221020-0001002112201322"></a>

## uid property — prefix_sets / 301033101200 / 8

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

<a id="canonical-1023011223020112-2001323123133333-1311222211311110-0332030222221120-1201003133113011-0321321201232230-3201332133333011-0130110230022230"></a>

## Next pages — prefix_sets / 301033101200 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-0203121202130031-2203220200231131-2001322323233203-0133212303033123-1123320021332113-2112022122112233-2331321110200131-3131031333310133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111102320002302-0210101031332300-3113003023002203-3122211120031033-0021332113133230-0120303321101031-3133132212123213-3011103223302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003322003231333-2003332010031331-0011330033120322-3101111031222332-3232013002103011-3321332221022120-1221311230123102-3201231201232310"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list — ip_prefix_list / 323201023022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-2322031131330030-2011201131110022-2233311200212320-3001202323231202-2322101130330202-0220211332202301-3022233001312301-0303223110213111"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003223122232131-3022223233113113-2231232331002211-3011111000013212-0032101310032210-3202213310012203-0031033031002231-3222102230113203"></a>

## Direct properties — ip_prefix_list / 323201023022 / 3

<a id="canonical-0311010311111232-1012333130130112-3301120111000321-1323032322312202-3002030333220130-2202231232032023-2012312132010021-3210333133302031"></a>

<a id="canonical-2112131210320323-1231011032233121-2121333322123033-1122211333302300-2311211333122330-3200113203313133-0201110003321333-1210023312023033"></a>

## invert_match property — ip_prefix_list / 323201023022 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0101230020031101-0300001020321222-3212123100310312-1110100333310302-0112223000231303-3321223320020000-2320111010203101-1332022313012202"></a>

<a id="canonical-3313110022033221-1031210011333230-2020032111202223-0101311203303011-1000300003301320-1132230220010202-2222303130321200-1002123310120120"></a>

## ip_prefixes property — ip_prefix_list / 323201023022 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1123301130303213-3331301232101102-3122003330223003-2000323310202311-1332131212013002-0311120022200101-0030103210301112-0133220330111301"></a>

## Next pages — ip_prefix_list / 323201023022 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020222103110101-2120130130101001-2303101303303210-2203023310303011-0101200212201213-1310123333103230-3133121312302332-2300113211122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210222033220032-3201213012222011-0132020311302300-1113311132222122-1232023102300200-3122002322322330-0332103122231012-0031022011031221"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.metadata — metadata / 322301311213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-3310232202230200-1301223221221232-0032130002323321-0121200212302101-1202300333111131-2222111323201213-2222310030023011-1200300112132133"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110010212122033-3120111300031133-3133023222031333-1220003323323021-3100330213101322-2311310112210011-3332013100230321-2231222021212310"></a>

## Direct properties — metadata / 322301311213 / 3

<a id="canonical-3033123210121212-2320010113303230-0131103010322330-3131223313133121-1031211232133321-1133310022230001-3332101233203133-0303220022223031"></a>

<a id="canonical-3120000222333321-3110120221320303-3302102112232100-0032233233020133-2232333000022201-1010302031322033-2020212013312122-1231220311013210"></a>

## description_spec property — metadata / 322301311213 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3031123200211102-0130302002033110-3201210122013103-0313120321201310-2323000220110202-0201133030320232-1000302302300113-2113103202030033"></a>

<a id="canonical-1211331201000133-1010120121232332-2231103213022221-0120333023120310-3212120222311331-3020111000331123-1131021313320333-1102123002311300"></a>

## name property — metadata / 322301311213 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0031210210323030-3332121210312010-2023232321333320-1211333120212121-0300312003322022-2010120130202223-3100121302012330-3000201023310301"></a>

## Next pages — metadata / 322301311213 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2211230323332300-3330211101102100-2023012322001003-1112202131233213-2321022211100031-0100321221002031-1333310000231113-0102002232313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223310302231203-1013133101320301-3011130111133103-2120101310322120-0330233320101331-3032220200332132-3301223133103122-2102121001030212"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.none — none / 222012030023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-0230220011112222-0331131123310103-1111230301221011-0010010003111023-1333101100011200-3311211023012032-3303310310320330-3002322011200331"></a>

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
none = {}
```

<a id="canonical-0223211222032230-0021122321122033-2121132231031301-0201332120301130-0322300322332232-1010030022011322-0303201123031211-1222332330203020"></a>

## Direct properties — none / 222012030023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331202102220013-0121100120212323-0120321212103113-2321132313221310-1122003010113030-1200223102301111-2213111210200321-2233133010032012"></a>

## Next pages — none / 222012030023 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-022.md#canonical-0101020121111330-0131112101110030-2303101320302313-2010200231331320-2330032111021330-3332112120230001-3123111322330033-3121130132021210)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113100030000322-1000020320313012-3123001301230221-2310022300233321-2003320223202212-3211200100311000-1021110011223310-3111111112301232"></a>

## policy_based_challenge — policy_based_challenge / 302202100102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- policy_based_challenge

<a id="canonical-0321223103102130-0121132313031223-1220210002011320-2133231330312121-2131132110220002-1313302332130023-1110033030130013-0300230120302232"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322130133111303-0311101322000321-2220220330230220-1122203010130321-1011210110312230-2333300203111012-2330021230011032-1221032003103003"></a>

## Direct properties — policy_based_challenge / 302202100102 / 3

- [always_enable_captcha_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1020103212230210-0302013301230200-1330100003322012-0303133122222112-2120233032330102-2023332131210232-2001103103013232-0223230323101222): complete subsection reference.

- [always_enable_js_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123): complete subsection reference.

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0303231223000210-3231100031100331-2121330111002123-2223112222233123-0023323312332202-1030113002130111-1203331300230133-2320320211010312): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-3100203011131011-3211013133233133-2213000122113301-1021032002123120-2312002312023020-3010133101311113-1122232312121323-1030100121231011): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0103303331213011-2202323031323123-0102330313313310-0232020323231013-3323001032002010-3302121300202230-1123012033300110-2102313203013002): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-022.md#canonical-2032331232010130-3032232001332021-3210130210012020-0000320322331212-0210133323311100-3200011130023303-2331221023110031-0222332120002102): complete subsection reference.

- [default_temporary_blocking_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2021133201102323-2012112211133233-2020010222002310-1013203132233103-3332132002103023-0011012112312331-3021312321312110-3300200102202133): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2231013300332210-2031112112313030-2221032321102312-0132223230313021-1321303100322023-2033103133323020-3122021120230210-1013101123002233): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-022.md#canonical-1322110130002031-0223022033013121-0332110131013120-0000132123321212-2123232030112001-0231012210300222-3202331330011102-2112222303312322): complete subsection reference.

- [no_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2232210000001111-3302330222322313-1131322202232312-3100230033303233-2301233011220321-1313012120021222-3122202003333100-2121101032211221): complete subsection reference.

- [rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033): complete subsection reference.

- [temporary_user_blocking](resources--http_loadbalancer--reference--group-023.md#canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001): complete subsection reference.

<a id="canonical-0302231311331200-3221303113233332-2223103301001113-2323121222131030-3202032201132012-0321201230132110-0212031012111333-3213220112311333"></a>

## Next pages — policy_based_challenge / 302202100102 / 4

- [policy_based_challenge.always_enable_captcha_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1020103212230210-0302013301230200-1330100003322012-0303133122222112-2120233032330102-2023332131210232-2001103103013232-0223230323101222)
- [policy_based_challenge.always_enable_js_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123)
- [policy_based_challenge.captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0303231223000210-3231100031100331-2121330111002123-2223112222233123-0023323312332202-1030113002130111-1203331300230133-2320320211010312)
- [policy_based_challenge.default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-3100203011131011-3211013133233133-2213000122113301-1021032002123120-2312002312023020-3010133101311113-1122232312121323-1030100121231011)
- [policy_based_challenge.default_js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-0103303331213011-2202323031323123-0102330313313310-0232020323231013-3323001032002010-3302121300202230-1123012033300110-2102313203013002)
- [policy_based_challenge.default_mitigation_settings](resources--http_loadbalancer--reference--group-022.md#canonical-2032331232010130-3032232001332021-3210130210012020-0000320322331212-0210133323311100-3200011130023303-2331221023110031-0222332120002102)
- [policy_based_challenge.default_temporary_blocking_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2021133201102323-2012112211133233-2020010222002310-1013203132233103-3332132002103023-0011012112312331-3021312321312110-3300200102202133)
- [policy_based_challenge.js_challenge_parameters](resources--http_loadbalancer--reference--group-022.md#canonical-2231013300332210-2031112112313030-2221032321102312-0132223230313021-1321303100322023-2033103133323020-3122021120230210-1013101123002233)
- [policy_based_challenge.malicious_user_mitigation](resources--http_loadbalancer--reference--group-022.md#canonical-1322110130002031-0223022033013121-0332110131013120-0000132123321212-2123232030112001-0231012210300222-3202331330011102-2112222303312322)
- [policy_based_challenge.no_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-2232210000001111-3302330222322313-1131322202232312-3100230033303233-2301233011220321-1313012120021222-3122202003333100-2121101032211221)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.temporary_user_blocking](resources--http_loadbalancer--reference--group-023.md#canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1020103212230210-0302013301230200-1330100003322012-0303133122222112-2120233032330102-2023332131210232-2001103103013232-0223230323101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122021223012122-2232320323031002-1111123322220003-0312030212301222-3013110132022333-1021201310222322-3101122203231323-2330310021112112"></a>

## policy_based_challenge.always_enable_captcha_challenge — always_enable_captcha_challenge / 000210122101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-3202301202020322-3102312023130303-3211312103313203-0003110221223113-3021101100231301-3323131103010221-3103300012013030-0032203212110233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

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
always_enable_captcha_challenge = {}
```

<a id="canonical-0113202303301231-0231211111000323-1031303300031001-0022223332130023-0323110131221313-3311221231021321-2330002111023323-1220131331200210"></a>

## Direct properties — always_enable_captcha_challenge / 000210122101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013112312301112-2001213022322302-0001001212300113-2210311320333101-0000213223122211-2322311112031221-3211113212331112-0000100102301000"></a>

## Next pages — always_enable_captcha_challenge / 000210122101 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2020101203121101-3122131030121332-0001233010301221-2210011313322031-2230000010103111-2322322320133211-3211101333203323-0113330131100123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003323030330310-3120222323012111-1302020000103132-0133321101310300-1302002302033332-0021220333222223-2130313232330222-3020302113000003"></a>

## policy_based_challenge.always_enable_js_challenge — always_enable_js_challenge / 123332021323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-3202221220113021-3213310202320012-0011232033203303-2203031310130130-2303032300301001-2013003312032103-0101103220320320-0013020231122131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

<a id="canonical-3113020331121130-1100022020312120-2123120021030013-1112202132310322-2122110312232013-2122230223332301-2003121203300223-2310312131110030"></a>

## Direct properties — always_enable_js_challenge / 123332021323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131003011110331-3131333331310210-0022003113100020-1230030320131130-2223320320131012-3023130031013311-0201102032122101-1220012022301321"></a>

## Next pages — always_enable_js_challenge / 123332021323 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303231223000210-3231100031100331-2121330111002123-2223112222233123-0023323312332202-1030113002130111-1203331300230133-2320320211010312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210022332010012-1223231322022233-2330302312023110-2310213222031321-0000011131303313-0232303022221322-3321010000212012-1231133320322120"></a>

## policy_based_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 211332220112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-3021003323123112-0111312010132002-3002120230301131-1122111133331303-0301133202010330-3101121321210232-3232020101213213-2330022321033203"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122230201301320-1120000131311031-3113322321131032-3321003113023310-2212300311020101-2023203033030210-3122201211102333-3131103111233202"></a>

## Direct properties — captcha_challenge_parameters / 211332220112 / 3

<a id="canonical-3030300121113111-0121113012013123-1021213032033200-3120011212311213-0130223233100211-2311120132130100-2222233031011221-0300120310102031"></a>

<a id="canonical-0312201021222302-3100002001001111-1021133210213020-3301233133023200-0012032321031332-0112011221333331-2211030022232321-3101120103133033"></a>

## cookie_expiry property — captcha_challenge_parameters / 211332220112 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2332013001311022-1112323121222103-1333223000230032-3111223021300203-0201200203302311-0201112302220120-2202033222033112-1303313212030321"></a>

<a id="canonical-3023100232332300-1332002002003212-3333102131001121-1303201110120301-3011232221203330-0331010212030112-2331023330330102-2013200101121200"></a>

## custom_page property — captcha_challenge_parameters / 211332220112 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3023212232230202-3132100330212100-2131223321011011-3022113223120332-0102203110303322-1202103002032221-0130322111022031-1021221001212331"></a>

## Next pages — captcha_challenge_parameters / 211332220112 / 6

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100203011131011-3211013133233133-2213000122113301-1021032002123120-2312002312023020-3010133101311113-1122232312121323-1030100121231011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300000322022121-1133232122311313-0223330310023222-1331130122111331-0123112021231031-1002103212201011-3102320231123002-2220210110030233"></a>

## policy_based_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 020020033232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-1022223320333010-1323310210203121-1333211223311011-2132023113222113-1132301212133331-0023200020201331-3222211223310021-2231321033110030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

<a id="canonical-2120012021333303-3321211322033233-2013131121113113-1301200210331212-2202211213313310-0130322201133222-2123013130030021-2210011312220002"></a>

## Direct properties — default_captcha_challenge_parameters / 020020033232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132131223312311-2323302310210333-0320102023230213-1103313001200322-0113002320310220-2221112103123220-2321123202201030-3030330301131000"></a>

## Next pages — default_captcha_challenge_parameters / 020020033232 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103303331213011-2202323031323123-0102330313313310-0232020323231013-3323001032002010-3302121300202230-1123012033300110-2102313203013002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301331020131333-3002202000122220-3122122023113233-1011220112330000-0132123211301012-1230230321323210-1003310011211333-2220202212203022"></a>

## policy_based_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 301002132221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-3323223032321022-0210001331213103-3103333023100313-0021213333232301-1021022111201313-1111012011221013-2332303231133232-2221203213201033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default js challenge parameters.

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
default_js_challenge_parameters = {}
```

<a id="canonical-1131201302110310-2233322013201101-3031122003120201-1211020200321013-0323313230202322-0210203022100301-0311321313002101-0211131003212321"></a>

## Direct properties — default_js_challenge_parameters / 301002132221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032203312320330-3030321110213303-0311230113330233-2111221322210121-2020231311111303-1223010133033300-2032001111310222-0030210213133103"></a>

## Next pages — default_js_challenge_parameters / 301002132221 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2032331232010130-3032232001332021-3210130210012020-0000320322331212-0210133323311100-3200011130023303-2331221023110031-0222332120002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300010113302032-0310031131033231-2203023121223100-0003301132310030-3123130002301031-0013323001033303-3111301322321321-1231213111022210"></a>

## policy_based_challenge.default_mitigation_settings — default_mitigation_settings / 102013011030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0200111120313200-0302001130002103-2230233333331321-0002223101111020-0013010133212303-0020320313223030-2112220331212101-1033233330122203"></a>

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
default_mitigation_settings = {}
```

<a id="canonical-1133101210333002-1301031023133122-2201102233131321-3003001001223202-2131001222222111-2321012202330230-2310022230021113-3010033112122201"></a>

## Direct properties — default_mitigation_settings / 102013011030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310321302102233-1330201231020300-1131322201203330-1321130321321120-0033011103202021-0301011121311230-0313102102322220-3322312103213013"></a>

## Next pages — default_mitigation_settings / 102013011030 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2021133201102323-2012112211133233-2020010222002310-1013203132233103-3332132002103023-0011012112312331-3021312321312110-3300200102202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322233202210010-1112301223300223-3233000321133002-2001130321330321-3233321211202210-0000010032311310-0113131300311022-3032120321203230"></a>

## policy_based_challenge.default_temporary_blocking_parameters — default_temporary_blocking_parameters / 230201333210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-0312312320020110-2333130330210331-0330023302013312-0300000100311213-1311310300110010-1321200303103113-0230001112203122-2213333012312202"></a>

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
default_temporary_blocking_parameters = {}
```

<a id="canonical-2031210330120321-0221310220223102-1010103103211103-0323103332201133-0231300323310312-1312121103002311-2202022022223321-2022333221202021"></a>

## Direct properties — default_temporary_blocking_parameters / 230201333210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210212220100333-0223112132100201-3311220012320333-3301102303103223-0132302213211301-2220321303101112-2033330330130120-3220313122011331"></a>

## Next pages — default_temporary_blocking_parameters / 230201333210 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231013300332210-2031112112313030-2221032321102312-0132223230313021-1321303100322023-2033103133323020-3122021120230210-1013101123002233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200300132320222-0303213032030020-2002222303301232-1303001321113231-1030330031130033-3111223123220313-2003112323033112-0222222120103231"></a>

## policy_based_challenge.js_challenge_parameters — js_challenge_parameters / 202000311121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-1302101310203202-1200232210223303-0013013232101202-0333332022300033-2233030321023030-3023001201302102-3322011222000021-3021331211030030"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020021021201333-0121112020300301-3320031103112310-0003111300132123-1220322300001111-1323313012230231-1012013031232331-2111212320231302"></a>

## Direct properties — js_challenge_parameters / 202000311121 / 3

<a id="canonical-2111211113300310-2120133101120122-3121213202210021-2322132020103301-3211122222300013-3101132130331311-3130221212122021-3033012021203122"></a>

<a id="canonical-3223030110100022-3112031201023230-0203002010220022-1221203231010221-2021212221301011-3231121333303220-0330223120202012-3331132231010101"></a>

## cookie_expiry property — js_challenge_parameters / 202000311121 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3203122313231121-1320222320123202-1202230200030233-3130120123110010-3112031331321311-2020222301002012-1103000020233310-2232320021201023"></a>

<a id="canonical-1300132313032022-1032200323333012-0333322011130003-2002330023321030-3021200111032123-2330032131133333-2123230231102201-3110002223010302"></a>

## custom_page property — js_challenge_parameters / 202000311121 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3020023123302223-3003323111203300-2132133020030023-1121212222322313-0320201021233132-0013030000112001-3320212300332002-2213213231222212"></a>

<a id="canonical-2000313131023302-0122322120221033-1223233102310312-3103311133320010-1333222230322332-3002003112000010-1320303012023221-0202211113110231"></a>

## js_script_delay property — js_challenge_parameters / 202000311121 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2121210211301313-0302332000003312-3331323102121003-0012211102333322-3203020003012101-1122233120013300-0023121112330303-0130121002031122"></a>

## Next pages — js_challenge_parameters / 202000311121 / 7

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322110130002031-0223022033013121-0332110131013120-0000132123321212-2123232030112001-0231012210300222-3202331330011102-2112222303312322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323220220321130-1321222022231302-2211022320233200-3030202223220112-2313320310232133-3230210001232130-1303101011213032-2021213221120012"></a>

## policy_based_challenge.malicious_user_mitigation — malicious_user_mitigation / 102301032133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-0110113330020201-1122230200121211-2230213021322131-2323211120330222-3000331303300013-2213000031021030-0122123023101210-2100132323321312"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212230111330312-3021223133031101-1302301023133232-0310020120310213-1323231320110300-3301131002021200-0302013102033231-1012111331233002"></a>

## Direct properties — malicious_user_mitigation / 102301032133 / 3

<a id="canonical-2030120120023123-1233011320121100-1021313030013320-1123212303113020-1222323131211132-1002031102301200-0230000200212233-2111003011101323"></a>

<a id="canonical-3012320221330113-1311130220020232-0323133313101101-1102133311332220-3303110232132203-0133023022121123-2010203032303230-3230033023123133"></a>

## name property — malicious_user_mitigation / 102301032133 / 4

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

<a id="canonical-2203202101203100-2101010302211213-3213302231331312-1323311202012111-1212100132312101-0212213233031230-0303202032013011-1201102011112002"></a>

<a id="canonical-3300030020231301-3110133131130213-3300131322023121-3001013322313030-0322331101102312-0213233133000123-1313011230031222-1333300321133003"></a>

## namespace property — malicious_user_mitigation / 102301032133 / 5

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

<a id="canonical-1222022003331221-3120132031233233-2032202013321323-0310202032132013-1313312000330012-0313032223231012-0311122122031331-1123131221231101"></a>

<a id="canonical-2001231322022110-2332121210322303-0232002020122120-2133201030032313-2123330121232302-0003230210312332-1122303002031300-3122303132320131"></a>

## tenant property — malicious_user_mitigation / 102301032133 / 6

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

<a id="canonical-3333203020302312-0322121032200233-3003303011011311-2112201031031303-0002131110131330-3112302330030230-3113300031202231-1212232213131101"></a>

## Next pages — malicious_user_mitigation / 102301032133 / 7

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2232210000001111-3302330222322313-1131322202232312-3100230033303233-2301233011220321-1313012120021222-3122202003333100-2121101032211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033032230011302-3113211101000200-3122331112120300-2112213200102231-1332002022221022-0130232223302020-0212220331200032-3313011012230311"></a>

## policy_based_challenge.no_challenge — no_challenge / 022131031133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.no_challenge

<a id="canonical-3322020102033022-2300220312302330-0131122233313010-0330000030030122-2032122320103010-2201203300001313-0212133230003321-3302111031222110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

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
no_challenge = {}
```

<a id="canonical-0303303302213000-3200231322202112-2311131323030112-3021130203323302-2230323331232022-1030132230122101-0312002213213010-1300212031120033"></a>

## Direct properties — no_challenge / 022131031133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220112332201201-0131022332020330-3321332332220222-0200202331000003-2113120213032023-0312221121101203-3013222323032222-1301223321200330"></a>

## Next pages — no_challenge / 022131031133 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231013011211332-0212210130201212-2203010011313103-2322103100220231-2310221220231230-1233201330332231-1233332032322333-0021101132100231"></a>

## policy_based_challenge.rule_list — rule_list / 210332211012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.rule_list

<a id="canonical-1021200113120231-1032213101120203-2022312102303122-3132200101003110-2332320123033011-2202000211133003-0320221133100202-3200333001230010"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

Receipt-pinned upstream constraints:

```json
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331002232302220-2332120201203322-3220233003112013-0012102103200103-2213201301332221-2133223221102333-2323121030000212-0100100103222331"></a>

## Direct properties — rule_list / 210332211012 / 3

- [rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030): complete subsection reference.

<a id="canonical-0333303002133213-0222130200310232-0023102131300111-3122310212331310-1113212201221132-0300223222030132-3322201201222010-2020221030300210"></a>

## Next pages — rule_list / 210332211012 / 4

- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102301313310312-0100211212030221-2112133012321300-2201110031212201-2101133022130033-3322131211010222-0202233112313320-1321002010310102"></a>

## policy_based_challenge.rule_list.rules — rules / 022202103121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- policy_based_challenge.rule_list.rules

<a id="canonical-1021200210032232-3303203123232232-3103201232322213-3200130123021131-2311200233232120-3010200302131222-0300013310133222-2202033022332232"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133133230201001-3230312001320320-0232020100022022-3000033302022130-0123003231113131-1032100031132103-1130103013001032-3112130022002320"></a>

## Direct properties — rules / 022202103121 / 3

- [metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0302122012200302-0210113312111030-0212231000333011-2200033300033333-2220311000021123-2332322330021310-0013231001133031-0020320013212302): complete subsection reference.

- [spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212): complete subsection reference.

<a id="canonical-0301202132230010-0221213031233233-2132211130210323-2232311233002213-1103011113221000-1113230212313223-1310013211010022-0213022302021231"></a>

## Next pages — rules / 022202103121 / 4

- [policy_based_challenge.rule_list.rules.metadata](resources--http_loadbalancer--reference--group-022.md#canonical-0302122012200302-0210113312111030-0212231000333011-2200033300033333-2220311000021123-2332322330021310-0013231001133031-0020320013212302)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0302122012200302-0210113312111030-0212231000333011-2200033300033333-2220311000021123-2332322330021310-0013231001133031-0020320013212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023210210010310-1002032020010001-1310312302310111-0110231003011311-0312100111001130-2331030321331033-3213132332102212-3001311003212111"></a>

## policy_based_challenge.rule_list.rules.metadata — metadata / 000201313121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-3131100303011033-2200033232132011-3031231230200102-2322030133132213-3013301131031102-3231021120332121-0032032322222000-2303000233220113"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213323013100323-2102113102111320-0122030203122113-2321230122120130-3001032112312131-2113213221302312-2201313133001222-0023011012002330"></a>

## Direct properties — metadata / 000201313121 / 3

<a id="canonical-1320232312110130-2102312300220031-1311103222322030-1233121000231001-1013202331212211-2312112201221212-1013322221222332-1003133210031102"></a>

<a id="canonical-2122232330100103-1313113001023210-2233221203200123-0032110100232312-2210312120301033-2100010332020202-2223331101000202-2331101131230330"></a>

## description_spec property — metadata / 000201313121 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3332332130010230-1123002320221000-3232120212013032-3031300131300120-2023213100010210-0202111223011012-3320023221210312-2210010311300233"></a>

<a id="canonical-0320032231233302-1302220310032012-0023023210200012-3310122131131210-1332302232232011-1313000320121211-1302221101001021-3323213120202131"></a>

## name property — metadata / 000201313121 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3101202003331211-1230230010020130-0000101223220310-3020131030221102-0311100123222201-0103321302113013-1213020110112223-1101210013111100"></a>

## Next pages — metadata / 000201313121 / 6

- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121113212202332-3022111213032320-3000332021000223-0022001032121331-2330200023312130-1301221132320311-0230130203210210-1131321111230210"></a>

## policy_based_challenge.rule_list.rules.spec — spec / 203300333132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-2022303111012211-1220213132123231-2022321212300020-1131132033230111-1321111301000001-2331030223131302-3011023300232321-0221120102012232"></a>

Type: `"object"`. single nested block, Optional.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110030122020113-1111323220212023-0322032000231323-2000123212032103-3012122201031231-0211123033103312-0321011300002331-2032222322032230"></a>

## Direct properties — spec / 203300333132 / 3

- [any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-0210300211122310-2122130331131112-2202100211131200-2203112212032321-2320330111111003-1022233132221233-0131330231132000-2210223331000133): complete subsection reference.

- [any_client](resources--http_loadbalancer--reference--group-022.md#canonical-3001013311103320-0133221211231322-0231122010331231-2020231333201210-0112210302022123-3232230023220113-1021012221130132-0031320122021332): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-3002222112023221-1013111330033200-2033000003103010-1313132112333032-3103213131123020-0130332002312010-2013133022111211-0133020111203033): complete subsection reference.

- [arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-0003301213131101-1312033131333213-2101303001113021-1023212221120031-1331213111320113-3131102321210212-0323000011233200-3231112332221032): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120): complete subsection reference.

- [body_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-023.md#canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133): complete subsection reference.

- [cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030): complete subsection reference.

- [disable_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322): complete subsection reference.

- [domain_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030): complete subsection reference.

- [enable_captcha_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121): complete subsection reference.

- [enable_javascript_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023): complete subsection reference.

<a id="canonical-3030123003333323-3321200033223023-0321023003322300-2310122221113131-3223202212223313-0122123320110331-1332103101001331-3212311310012200"></a>

<a id="canonical-1221111030002003-1023000201203201-3202310002103231-2302200000333133-2332121210102320-0312002310003313-1303231031102020-2121122201231300"></a>

## expiration_timestamp property — spec / 203300333132 / 4

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111): complete subsection reference.

- [http_method](resources--http_loadbalancer--reference--group-023.md#canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-023.md#canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-023.md#canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313): complete subsection reference.

<a id="canonical-3003121022001322-3100330212311323-0112012122103311-3032311010011111-0230001113003310-0210023001321012-0331032203032213-0222322113132212"></a>

## Next pages — spec / 203300333132 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](resources--http_loadbalancer--reference--group-022.md#canonical-0210300211122310-2122130331131112-2202100211131200-2203112212032321-2320330111111003-1022233132221233-0131330231132000-2210223331000133)
- [policy_based_challenge.rule_list.rules.spec.any_client](resources--http_loadbalancer--reference--group-022.md#canonical-3001013311103320-0133221211231322-0231122010331231-2020231333201210-0112210302022123-3232230023220113-1021012221130132-0031320122021332)
- [policy_based_challenge.rule_list.rules.spec.any_ip](resources--http_loadbalancer--reference--group-022.md#canonical-3002222112023221-1013111330033200-2033000003103010-1313132112333032-3103213131123020-0130332002312010-2013133022111211-0133020111203033)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- [policy_based_challenge.rule_list.rules.spec.asn_list](resources--http_loadbalancer--reference--group-022.md#canonical-0003301213131101-1312033131333213-2101303001113021-1023212221120031-1331213111320113-3131102321210212-0323000011233200-3231112332221032)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031)
- [policy_based_challenge.rule_list.rules.spec.client_selector](resources--http_loadbalancer--reference--group-023.md#canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](resources--http_loadbalancer--reference--group-023.md#canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- [policy_based_challenge.rule_list.rules.spec.http_method](resources--http_loadbalancer--reference--group-023.md#canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](resources--http_loadbalancer--reference--group-023.md#canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212)
- [policy_based_challenge.rule_list.rules.spec.path](resources--http_loadbalancer--reference--group-023.md#canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0210300211122310-2122130331131112-2202100211131200-2203112212032321-2320330111111003-1022233132221233-0131330231132000-2210223331000133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232230310023031-1020112031031223-3020011321031323-0002223331102101-1222310300302102-1010012000021321-1023330211032333-2231330000222310"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — any_asn / 222302000301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2013222223203003-1010002210001202-0221101112111322-3011022010201213-3203222223030333-2212232131313102-3301211023323030-2220330030202203"></a>

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
any_asn = {}
```

<a id="canonical-3133003302221001-3101120222102230-2322322012203012-0133033230030302-0212020320313030-1210201211101022-3110102200200112-2333310121303130"></a>

## Direct properties — any_asn / 222302000301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321220110021120-3021121302111000-2200131333002303-2101110132030121-1221001123200010-3213103133322122-2030310232121300-2001313200033230"></a>

## Next pages — any_asn / 222302000301 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3001013311103320-0133221211231322-0231122010331231-2020231333201210-0112210302022123-3232230023220113-1021012221130132-0031320122021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123033103103033-0032230111222232-2203011121203123-2021021312202210-0231233113203330-2101332022220101-0331310110110130-0232120202133131"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — any_client / 130101133023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-3000222333000233-0223102010013120-1301003330013123-1010033303310201-3212103301133331-0200332233032123-0200333010322013-0201120201211232"></a>

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
any_client = {}
```

<a id="canonical-1313320211132123-0201200122312303-2011333021333211-2211331101101020-2313001200102012-0212002132330133-0311100012312012-3333133221321011"></a>

## Direct properties — any_client / 130101133023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311120233212210-0100213303220222-0303013101002102-3010120330101323-2311032311332011-1330112032131313-1132221030233303-0323220001131113"></a>

## Next pages — any_client / 130101133023 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3002222112023221-1013111330033200-2033000003103010-1313132112333032-3103213131123020-0130332002312010-2013133022111211-0133020111203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303123301033032-3313310100221233-2010312102032322-2302033131103223-3010203113333011-3200002230301023-1013032312321031-2313233300130331"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — any_ip / 330031232010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0123320101312210-3323010200330110-3120331031000321-1101301302221312-1102131310230303-3002031220301321-0321113031203132-3311110011202033"></a>

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
any_ip = {}
```

<a id="canonical-0013000211130333-3110233132212331-1321313122223111-0130131002200031-0031213210231032-3312210032120222-2110200120102222-3031322030110033"></a>

## Direct properties — any_ip / 330031232010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310100320113002-3030011102313233-2000313121222330-3002333303220022-0213031301103101-0300030001331123-1300200023131210-3201331213100022"></a>

## Next pages — any_ip / 330031232010 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013023210232130-1232121023233100-3102111131110020-0303131122003332-1230030110323230-0320102211132321-3322111123023201-0110311203122030"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — arg_matchers / 120323313002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1001313021020303-2233232132332232-2201323103221012-3133111210202121-3020221223111300-3033323232101202-0203203312332002-2121003213212002"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132023312223130-1212131310031303-3000222200310012-3323321023100322-2023312031133232-3222123210210300-3003331200221003-1103020300030203"></a>

## Direct properties — arg_matchers / 120323313002 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-2303020231320132-3322122201120122-3210131111212200-3120303012212232-2322233112010021-1311130030331303-0012320213231121-1203220020213211): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-022.md#canonical-1131222013031202-3211232223001312-0013120330133031-1110232213332332-2102200023123203-2111103310301333-2323122300323123-0020231033321300): complete subsection reference.

<a id="canonical-1130221332110311-2333233103121133-1233013311131210-1212123111203221-2233022131100120-2032031021023331-3013133121220100-0101121100320032"></a>

<a id="canonical-1033321322003302-1232011331032211-1221300323332331-2012231331331011-0233320012203310-1302201101223311-3100123313300232-3302130123030332"></a>

## invert_matcher property — arg_matchers / 120323313002 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-022.md#canonical-1103010012000013-2110202002230020-2032130233221320-1233032213333113-3202021112010211-0033113003100012-2012222100210311-3211210200221222): complete subsection reference.

<a id="canonical-2231310202311021-0222332222232011-2023002233000023-1211213110302201-1313330001200102-0003230032011032-0303031012110313-2203002020020020"></a>

<a id="canonical-3223222322220032-3133231313013233-3233322333213232-2030200312122000-1221212121333321-2011023230301100-1130101012222223-0000321321232333"></a>

## name property — arg_matchers / 120323313002 / 5

Type: `"string"`. Optional.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1103110213300031-2221330001000322-0013103021321231-1333222202323210-1333103023112123-1212312201001302-0300223102113222-2233231000120213"></a>

## Next pages — arg_matchers / 120323313002 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](resources--http_loadbalancer--reference--group-022.md#canonical-2303020231320132-3322122201120122-3210131111212200-3120303012212232-2322233112010021-1311130030331303-0012320213231121-1203220020213211)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](resources--http_loadbalancer--reference--group-022.md#canonical-1131222013031202-3211232223001312-0013120330133031-1110232213332332-2102200023123203-2111103310301333-2323122300323123-0020231033321300)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](resources--http_loadbalancer--reference--group-022.md#canonical-1103010012000013-2110202002230020-2032130233221320-1233032213333113-3202021112010211-0033113003100012-2012222100210311-3211210200221222)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2303020231320132-3322122201120122-3210131111212200-3120303012212232-2322233112010021-1311130030331303-0012320213231121-1203220020213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133113011031230-3303332110320121-2332321133320331-1033020023313222-1220020010001202-3312301002130132-1030010333100322-2112013023202100"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — check_not_present / 311103203231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-0100121332123113-3331011212220111-3012201313102321-3113000032032110-0220303321331033-1113223122000032-0020201123301003-0133120233112101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3001311113113223-1012121101130103-3102312313203320-0312103323021032-3113111332031103-3323313303001012-2212021032032010-0312313210231333"></a>

## Direct properties — check_not_present / 311103203231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220213210022223-0232221133212033-1033221203330220-3232221213103033-0210110031210322-0222232312121002-3333312302033132-3133002031013132"></a>

## Next pages — check_not_present / 311103203231 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131222013031202-3211232223001312-0013120330133031-1110232213332332-2102200023123203-2111103310301333-2323122300323123-0020231033321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023221100211333-1300213321231222-3330133331101333-2203212021220332-2131221332322011-1030211113312012-0100332032330202-1303012122323131"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — check_present / 102321332222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0031112120332301-3123332121231323-1201312030322202-0111232033010203-3202333011323033-1330011322121331-3033001101031121-3333002300030132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-2212010333221100-0102323102103020-1200033220120313-3111022201300330-3003311333113310-3323311313031132-3023111321322232-3133210132203011"></a>

## Direct properties — check_present / 102321332222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312132231222223-2110020120332200-3001111320320313-0203130232222123-3232303212303132-0010232123133110-1201300231032232-3312011133123032"></a>

## Next pages — check_present / 102321332222 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103010012000013-2110202002230020-2032130233221320-1233032213333113-3202021112010211-0033113003100012-2012222100210311-3211210200221222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223213111201123-3311123121330201-3010133000223003-3012132310320313-0233233323121122-0121031013133222-1313001110033033-2133233211130133"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — item / 311111223000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-2330222211011232-1201201232011221-2131000030021300-2131103230112111-0123213133303311-3121201131132120-1222330222113112-1023133330302013"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231120321322131-2233301312310003-1111111012301312-2020001101302023-0020002010121312-3331030230311012-1101201323203021-1303013203211321"></a>

## Direct properties — item / 311111223000 / 3

<a id="canonical-3022312231102031-3003001123332303-2121230122122222-3200331030120323-0111232200323033-1112303302301213-0101003211333213-0113013003020311"></a>

<a id="canonical-3130301213011102-1122233101311102-0130321201112031-3010300202001330-1222323330203213-3322002333302301-3200331020033103-0210013310303011"></a>

## exact_values property — item / 311111223000 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2322213110203301-0313223010210110-3032211302002120-0133322100303232-0223133102232013-3320221333022001-2112300131030000-0002301113113200"></a>

<a id="canonical-1211012033113201-1202133303321231-0130022212303213-2103313323323211-1023003132312012-2200213132213322-0231032103313223-1001202332030333"></a>

## regex_values property — item / 311111223000 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2132133303101030-1101013003012113-2202113222012033-3201030130031230-0331320320330333-0301321323331321-2133202322300233-2130302220223313"></a>

<a id="canonical-0000102203333003-0012332311301233-1021113003201031-0030211322210133-2133213131321000-0231102003012133-3202122000231220-3223200321333120"></a>

## transformers property — item / 311111223000 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2221332302320322-1321230233213033-3012333220222212-2132011303313101-2003133323030320-3012321131032213-2001132232323131-2130201220211303"></a>

## Next pages — item / 311111223000 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--reference--group-022.md#canonical-1101313011001300-1031201312333332-2301022020220121-1201202331311323-3022013032203330-1000001113112111-2300012330010303-3001011233020023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0003301213131101-1312033131333213-2101303001113021-1023212221120031-1331213111320113-3131102321210212-0323000011233200-3231112332221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122232220022111-2313330200023223-2002212301312012-0233110300122110-1301102030012312-3332321332320002-3130223121102311-2100013200220030"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — asn_list / 121310311303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3313120122012310-1312303301112131-0330231303213310-2123203000010211-1300220010012002-1202311122313213-0122023202310323-1323200031333331"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213003110022222-3132103211211311-1333322022321322-3022222220012002-1222100332311022-1223103210001131-1020021121220322-1120303302312032"></a>

## Direct properties — asn_list / 121310311303 / 3

<a id="canonical-2011021001231323-2000131220333220-1213030231001021-2222211233332011-3200203011230231-1032303202201010-3033311321013300-3033212211221031"></a>

<a id="canonical-1001200020333010-3113233203232001-2023101320033000-2201333011232113-1020333010032002-0122331200132201-0001021102003001-1303013022132301"></a>

## as_numbers property — asn_list / 121310311303 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3103311200320231-3201300033130002-0311110313110013-1210323331320130-0100220332011321-0303303122311020-0032210332011030-0331100330120332"></a>

## Next pages — asn_list / 121310311303 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131123003101103-1001130313231203-0302303201002010-3022100130331111-2133311213231233-1033310021312202-3013012121222013-3332222022130311"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — asn_matcher / 210001223123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3332020131111132-3123312111131233-0221120323023112-0032020302001230-1123232131010332-2012320311231101-1302113001110312-3223032100201333"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231223333033033-3303102333110111-3033120210100120-2211122020101103-0333122220201113-2030222331322030-0301001201200300-1123002102031022"></a>

## Direct properties — asn_matcher / 210001223123 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021): complete subsection reference.

<a id="canonical-1333123310303111-0211231033000302-3030213023233320-0001120123323322-0211031332012220-0013201330102332-1120020203232000-3320103210212330"></a>

## Next pages — asn_matcher / 210001223123 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-022.md#canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332102123323020-3310120000112032-2113123312312323-2101132113131211-0033212110333232-0231330122031101-1111023020011202-2203333010232331"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — asn_sets / 301131010220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-0013021303033210-1310033201100322-3122233210301123-1330330222000221-3003222212323023-1000102221322202-0121222022133110-0330112311003203"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302113112310220-3121121123332011-3311210022200113-2001102001133302-2233123313303221-1212300101301000-1200200123011322-1222302310220213"></a>

## Direct properties — asn_sets / 301131010220 / 3

<a id="canonical-0320221303031113-1201232302222013-2220013000231333-0231101013122131-1232213201030302-3013222323303001-0021302230012310-0332321322121202"></a>

<a id="canonical-3320122012320132-0202320012122131-0222311000000121-3012332300300322-3010221020002101-3121232001032132-1023113201302313-1321301020212313"></a>

## kind property — asn_sets / 301131010220 / 4

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

<a id="canonical-0230311112202302-2301012301101323-1311131202301211-2200212332113200-1313003010320113-0013223202321213-1112332101330033-0032222300101331"></a>

<a id="canonical-0300213200120212-0123302231201203-3233312031002311-0211102003111120-1311030013032310-2202000211131023-0303111203212021-0300131330131301"></a>

## name property — asn_sets / 301131010220 / 5

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

<a id="canonical-3212100010123103-2201132023213020-1211212303303201-2123031302211020-2013030122021222-1201023120030321-3000022303233000-3323331201010001"></a>

<a id="canonical-2312012303230012-1311213303330030-0211111103303233-1220321213132132-0133023121122111-1211011010311111-1223022030120032-2302032203313023"></a>

## namespace property — asn_sets / 301131010220 / 6

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
  }
}
```

<a id="canonical-1220333021333210-3132101033202303-3103321130213330-3233000200202023-0123113103103333-1200202320212323-2130233111210101-2330331333021011"></a>

<a id="canonical-2131000323030221-0011320300231300-1213331113030022-2010300102320112-0023200112021311-2013113302012102-0012023000103002-0231212131023330"></a>

## tenant property — asn_sets / 301131010220 / 7

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

<a id="canonical-0212030212030313-1320332202232022-1320023311300203-1131000132123221-1222323212112202-0001200211302323-0233231320133300-0300232321131022"></a>

<a id="canonical-0100113313330001-2303011300320123-0030202232123221-1333301203011201-1131300321313131-3123303232112222-3111310211022222-1202120023022233"></a>

## uid property — asn_sets / 301131010220 / 8

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

<a id="canonical-3231303300303020-0111003331202311-3113303011120021-2031113323121111-1112122312101300-2031121010022022-2123232002122211-1123001313102131"></a>

## Next pages — asn_sets / 301131010220 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-022.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
