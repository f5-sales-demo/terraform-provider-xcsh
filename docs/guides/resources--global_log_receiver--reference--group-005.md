---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-3022033313323111-0210301021020213-1220202220003122-3030002012120131-2231223033132333-2322333110210032-0011313010022113-2021301332200310"></a>

## location property — blindfold_secret_info / 222020333133 / 5

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

<a id="canonical-3222013330321333-2212200112223103-3121022313300311-2110223113000330-0222232202303221-3001313023112003-0330031133001303-3031200232122321"></a>

<a id="canonical-2112121201312202-3132200211013033-2001121322213211-2022211201032110-3130011031131200-3301111133210320-1333201312030221-3002231313312032"></a>

## store_provider property — blindfold_secret_info / 222020333133 / 6

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

<a id="canonical-1233030111200133-2332303123330001-1101211130130223-3223113003030110-1311121030111011-3200223333202012-0121002201303030-0331223023222210"></a>

## Next pages — blindfold_secret_info / 222020333133 / 7

- [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-004.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3121032322130121-1230231123003111-3311322132332221-3210310302030210-2213011003012100-0311331110013130-2130103210211231-1201203322103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311230233203022-2120102310031011-2200033311132302-2333300100031011-1123200210033200-2322122011112302-3102222200112121-1132032001002002"></a>

## sumo_logic_receiver.URL.clear_secret_info — clear_secret_info / 023201002312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [sumo_logic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-3133110202111331-3110012132230231-1313311200132200-3313312202001233-0330332232223132-0100032000033300-1111310013112123-0303121102330130)
- [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-004.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030)
- sumo_logic_receiver.URL.clear_secret_info

<a id="canonical-3230133031132200-3013323300021221-0112022130230010-3202231331223233-1020232013320220-1302102100022000-3020202200110211-0120121310313130"></a>

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

<a id="canonical-0322011210201100-3031103320332313-0031222030213100-1331323321030102-3022102102100101-2032133103302033-0031202102311312-2222211110312311"></a>

## Direct properties — clear_secret_info / 023201002312 / 3

<a id="canonical-2311220113323202-3332332033030300-1330231021010113-0320111133110010-0112013000310000-3010220131131323-1000313020310201-1200201322000011"></a>

<a id="canonical-1221211113031202-0002311130301201-3133230222300211-1003321302300121-3322120200330100-0311231122133213-2132022123322003-2130112322033303"></a>

## provider_ref property — clear_secret_info / 023201002312 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2321122321002030-1210232310222233-0110123110130110-1302010200301031-1201130100021020-0300011301033012-2000332231330311-3123203312321002"></a>

<a id="canonical-1233200233033021-2030320121331020-3131231210211331-1122110122021230-3302012322023003-0021302312321221-3312212100220231-3110231231103033"></a>

## URL property — clear_secret_info / 023201002312 / 5

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

<a id="canonical-1020111311021020-3320121031200030-2031130132323220-2112121311331010-0312303123301321-3021200123110213-0332102121130323-2233300322131133"></a>

## Next pages — clear_secret_info / 023201002312 / 6

- [sumo_logic_receiver.url](resources--global_log_receiver--reference--group-004.md#canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2300030330132202-1012220213322303-1300301203131302-1312202110301333-1300333322313010-2203010102023132-0210122130033122-1033121302122132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030220010312322-3101120012303230-3312001101200101-2003022113331220-0313121223131001-3113320112331321-0201121031121333-1113320323310102"></a>

## timeouts — timeouts / 113200330221 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- timeouts

<a id="canonical-2323101322201020-2201013121231321-3110201100222113-1023233112330212-2113123032321331-2031001100133233-0232113100210011-1212031323321012"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133322001001321-3233130010301033-0213100202110021-0332233013231232-1311011131230133-1222321001131121-1200301200020330-0211320231322111"></a>

## Direct properties — timeouts / 113200330221 / 3

<a id="canonical-2130100022000302-0233020112231100-1203103031203200-0221232133223120-1010202210130203-1301002211010230-0133223110303203-0330233303323013"></a>

<a id="canonical-1313030310320030-1121132132103311-2222332113300230-2302030002121103-0301030322001100-2200101301233232-1211301212231130-3021021213223132"></a>

## create property — timeouts / 113200330221 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3121121313100001-0210133021230333-2012103322333331-1311011223032023-3311113113131130-2203312331030031-3233013333221131-2103111212303002"></a>

<a id="canonical-3030130120321110-2310311232211223-3112233012130133-3033132313333211-1231030211003122-0000000211200301-0202011031030313-2022030020123323"></a>

## delete property — timeouts / 113200330221 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1221022123333101-3003311211233102-1033203310301211-3111100313022012-2231100203331300-2033122101323220-1211312002313233-3001222231021020"></a>

<a id="canonical-1033212231103123-0001321121023102-2310330230303122-3020311003012100-2001302132001102-2211031333132200-2123010011121331-2133002323122221"></a>

## read property — timeouts / 113200330221 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0313011203010113-3312002121230013-0210100012300332-3111003220030121-3033102121222311-1321312032232010-3002210210020333-2102131202223102"></a>

<a id="canonical-1210031103223033-1002322023300300-0030323211133012-0231123130332231-0230221222113000-2221200111102231-1222021221321220-0330002213032233"></a>

## update property — timeouts / 113200330221 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3132202303233301-2332123013331302-0120032021133113-0001120110202013-2102323303111213-3201322201220133-0023213001023003-1231202200313330"></a>

## Next pages — timeouts / 113200330221 / 8

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
