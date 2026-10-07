---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-2313211122003111-0133311312310203-1110230130303301-3320333120121122-1021300210311203-0201330003131332-3020021303223201-0232311033013211"></a>

## Direct properties for `kafka_receiver.use_tls.mtls_enable`

<a id="canonical-0110231120301200-3031202311211211-3121103021011312-0233313210312232-0321211331201002-0332220300311323-0321010001203313-3311113003223112"></a>

### `kafka_receiver.use_tls.mtls_enable.certificate` property

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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

- [key_url](resources--global_log_receiver--reference--group-004.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100): complete subsection reference.

<a id="canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- kafka_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3300203201300012-0012033001230122-2303010112222211-3230021003303320-1103330030000313-1203131002321323-1322103102223033-1302302333013333"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132031133132132-1131001220032233-1301232213120010-1301120111130222-1022103013033300-3222112130002223-3003133310333123-3221030010300021"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2323123022132202-0223213323323201-3313313233120121-1221021100320302-2211311020332132-1113210232233133-2010222012020023-0331201000333023): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-1212021111021103-0213311023223223-2121122231113102-2203321212030000-2022033313220132-3320211002201033-1302120231102331-2211223100002011): complete subsection reference.

<a id="canonical-2323123022132202-0223213323323201-3313313233120121-1221021100320302-2211311020332132-1113210232233133-2010222012020023-0331201000333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0330202331311013-2221000013002320-1023002333321333-2202331030103120-2001020120113123-3312311100010030-2331110012223112-3123311313311003"></a>

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

<a id="canonical-1112312313330223-0033300021103202-3310213213011321-1310310201032103-3012030011332030-1012200233012231-1131330300010302-1323100220230330"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-0303101220303220-1002313302300221-0110000011310300-0220202210101020-1012300221223320-2003112303021120-1022222001122222-3121130201121212"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1333303221122020-0331310322311000-1110331223312311-3102003020312220-3302001201000022-3302230223123323-0032302320101210-3323200323130103"></a>

<a id="canonical-0002021020221230-2220222201302103-3022332010303303-2311213302231130-3213111100023101-2101131032030220-1203220001213333-3133013201000321"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3111032233112222-2033302111212203-1313210103322003-3213313112231010-0221030013103322-0033110231011012-2010022022200131-2032322112010212"></a>

<a id="canonical-0100330103203223-2103020322223302-2330031023121323-0103303313200312-0313223002333032-2131230002231123-3031232212023022-1031101130322301"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1212021111021103-0213311023223223-2121122231113102-2203321212030000-2022033313220132-3320211002201033-1302120231102331-2211223100002011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2100030213311013-1230320030301120-2303323230120331-0012202212031132-1122110120111110-2210232122110320-1000112302220233-3213322001230320"></a>

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

<a id="canonical-1003320322231231-0313203233330333-1113010121010131-0000130013031110-2223011302301301-1032013223303321-0200112032331312-3010333221101122"></a>

### Direct properties for `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-0030112321202221-3202023221313301-2221201132220210-1322231101113203-0132223230302123-1010312200113020-2232122321322033-3302322233330002"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0002000202030112-0110002223313130-1213033302323321-0033323211220122-3130212123322023-0011111130111330-1203231021000110-0033113220332122"></a>

<a id="canonical-0012030301110031-2003123033203103-3202221303231222-2220323333323312-2013012200113023-2330020231111131-3020022031021031-3311130213212120"></a>

#### `kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1003222302113131-3020001323322130-2003312222122100-2203013200001021-3223010100323333-2303133002013033-1011312223111120-0302222010210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.no_ca

<a id="canonical-0311121221322223-3110013021312110-2102323332312233-3012123011123310-1020331301231130-1200103120013310-0000233021213232-0122002222203331"></a>

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
no_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- new_relic_receiver

<a id="canonical-1230012130102110-2311120330023112-0022233130211001-0010032110111023-1303133132230232-3310210232122022-0323333321323230-1021002003003111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for new relic receiver.

Additional upstream details:

Configuration for NewRelic endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("eu",
    "us")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"eu\",\"us\"]"
}
```

Terraform syntax:

```terraform
new_relic_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200231031221120-3213122133312313-1112211212233102-3333212020221111-1203001230213201-1002222331013010-0031222232033313-3103120223132000"></a>

### Direct properties for `new_relic_receiver`

- [api_key](resources--global_log_receiver--reference--group-004.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300): complete subsection reference.

- [eu](resources--global_log_receiver--reference--group-004.md#canonical-0111332302103311-2231321210233231-3312031102230011-1222123203220300-3022101201022312-3033002110232322-1212003333010131-0331232213123100): complete subsection reference.

- [us](resources--global_log_receiver--reference--group-004.md#canonical-1012122103003011-0312131130200120-1312311010011133-1000220032230110-2330033223201123-0121312311122030-1103322213300021-3102320222222220): complete subsection reference.

<a id="canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.api_key

<a id="canonical-0223021210302310-2012203211031313-1322122223013330-3333232121232301-0213001011330131-1222123201101313-1133230002211110-1110322300302202"></a>

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013023023220300-0101211210200110-2332303210122023-3021130332230022-2212332313232301-2311221011003130-2022333322231011-1100221101011001"></a>

### Direct properties for `new_relic_receiver.api_key`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-0312211120000312-0322322322111220-1103110331000011-0123030020200031-1030213100103200-1033200303231210-0021220300121233-3313021331211210): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2033113111202231-2211303110023112-0011123201132123-0222303230312333-2131100001321221-1301313200201112-0131031333212002-2210301321033121): complete subsection reference.

<a id="canonical-0312211120000312-0322322322111220-1103110331000011-0123030020200031-1030213100103200-1033200303231210-0021220300121233-3313021331211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-004.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- new_relic_receiver.api_key.blindfold_secret_info

<a id="canonical-2032120110010311-0032132333203322-1132112202331322-2220130001023332-3213032222031103-0223321221232213-0033211222213300-3303200123321320"></a>

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

<a id="canonical-0202313203303333-3100333320130213-1201332002101310-3001000101011312-0232102333033312-2223003130133202-3023310310010203-1112213121120231"></a>

### Direct properties for `new_relic_receiver.api_key.blindfold_secret_info`

<a id="canonical-1310031222201233-1102211023320132-1201213110213311-1102002120320332-1100303320321332-2003302203112310-1220323111202123-2022010231212230"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2223200113202210-3130331131313233-0223133322100133-0200013231321032-0231232231203311-3230011230103011-1301013203120111-2321323001131203"></a>

<a id="canonical-1110132220012021-2110222103020011-3333223330330321-1320100131133100-1111222032033201-1002200112110333-2323332013130323-2323112322202210"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1133112221202003-3322132333313232-0210001032313002-2010103110102120-1123113211023121-0020310303210310-0030131031110202-0232023301112113"></a>

<a id="canonical-1003233131332231-1022133320121200-3203013210113131-2002333021120033-3002003201020230-2233323120323310-3331231303302021-3313303331132103"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2033113111202231-2211303110023112-0011123201132123-0222303230312333-2131100001321221-1301313200201112-0131031333212002-2210301321033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-004.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- new_relic_receiver.api_key.clear_secret_info

<a id="canonical-0030311021310330-0131113210013213-1232010222121301-1122320200202300-0312220001232330-3100302211131301-0113022101000101-3011220003031232"></a>

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

<a id="canonical-2311311003321022-3322212310220201-0023312321211000-1302223223202133-1131220122121211-0312101223200323-0112212023030030-3323132100012003"></a>

### Direct properties for `new_relic_receiver.api_key.clear_secret_info`

<a id="canonical-0000022200121113-0101123102113002-0030331121313122-1103103311133202-1010213101000300-1321321132131131-0212213131212033-3102120200303022"></a>

#### `new_relic_receiver.api_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2131333100232312-0221013031230030-0311001311122100-1133113130310013-0212310311112031-1031131203320301-3101021331230132-3222333311200332"></a>

<a id="canonical-2120131131222221-1222022233212213-0332333322333121-1120003233000201-1003300210132111-2233223133132130-1002020322112233-1120132201211011"></a>

#### `new_relic_receiver.api_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0111332302103311-2231321210233231-3312031102230011-1222123203220300-3022101201022312-3033002110232322-1212003333010131-0331232213123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.eu` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.eu

<a id="canonical-2102232032231100-0233300020131310-1112120231023302-1031010013300202-1330132220031312-3013320300021031-3302312320130021-2011103002202231"></a>

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
eu = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012122103003011-0312131130200120-1312311010011133-1000220032230110-2330033223201123-0121312311122030-1103322213300021-3102320222222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.us` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-004.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.us

<a id="canonical-1003110133033023-0101310221011312-2323123212111201-1332232223123003-3003100133102013-3000223032322323-0211131211032032-3300012321103200"></a>

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
us = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110223021133332-2001211101313323-3012032332230023-3123201031102303-3033233222010310-0322300301310121-1032103033320210-2132202001210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_all` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_all

<a id="canonical-3002133321321130-0230232203010223-1110130322321023-0302022123020231-1112103320022121-3221331330103213-2132031121012132-0211011322121200"></a>

Type: `["object", {}]`. Optional.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

- [ns_all](resources--global_log_receiver--reference--group-004.md#canonical-3002133321321130-0230232203010223-1110130322321023-0302022123020231-1112103320022121-3221331330103213-2132031121012132-0211011322121200)
- [ns_current](resources--global_log_receiver--reference--group-004.md#canonical-0301322210023112-2021310211001030-3312222220121313-1132211011001223-2002213301202033-3323203023001212-3233001223111202-3221030231123021)
- [ns_list](resources--global_log_receiver--reference--group-004.md#canonical-0331201313123231-3032231333310001-3022103022000000-3122210132332000-1233123202200013-1012100010010213-1321022201103201-3120010330312100)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ns_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131102221232000-0300222020002322-0302133211001210-2101203323112122-0311121123102201-1122110013312213-1303121113303221-1111202320322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_current` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_current

<a id="canonical-0301322210023112-2021310211001030-3312222220121313-1132211011001223-2002213301202033-3323203023001212-3233001223111202-3221030231123021"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
ns_current = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323013202223030-2213130333303312-1001111102133131-1001220102221312-3002003012031130-3113012231221100-0102230132130332-0202133321133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_list` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_list

<a id="canonical-0331201313123231-3032231333310001-3022103022000000-3122210132332000-1233123202200013-1012100010010213-1321022201103201-3120010330312100"></a>

Type: `"object"`. single nested block, Optional.

Namespace List. Namespace List.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("namespaces")}
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
ns_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313111113201030-0133311132101103-3322303023320232-1002211330201121-2123221001221003-2210332201232102-2121333032221230-1132022211120120"></a>

### Direct properties for `ns_list`

<a id="canonical-3311133003311012-2112320033113111-1003020223221332-0023033020003022-1012111233103312-0120211012200222-3302322101001330-1112303001300201"></a>

#### `ns_list.namespaces` property

Type: `["list", "string"]`. Optional.

Namespaces. List of namespaces to stream logs for.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- qradar_receiver

<a id="canonical-3323203113301331-2013010003331233-3133212102210323-2032121213300323-3201011213001320-1331212033033120-2133211210012323-2133022310331130"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for qradar receiver.

Additional upstream details:

Configuration for IBM QRadar endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("uri"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
qradar_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002012330022033-0113301030001303-1001110031010311-2021200322312113-2033330102020111-0302200021330203-2003323020132200-2233123130232102"></a>

### Direct properties for `qradar_receiver`

- [batch](resources--global_log_receiver--reference--group-004.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-004.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332): complete subsection reference.

- [no_tls](resources--global_log_receiver--reference--group-004.md#canonical-0333030003010101-1102331120011301-3023322233120221-1332221330232103-0110203322122232-0233200210022200-0211212203313232-2301331121322003): complete subsection reference.

<a id="canonical-0011323232033200-1111202030133200-2301112132001000-1121023022001212-2201102222023303-2000003311222021-2031022101301322-3113123212002213"></a>

<a id="canonical-3313013021310303-0131130233113000-0112000220122012-2233331121133130-3100222323221322-3020010210233131-3330021130021102-3132200113110102"></a>

#### `qradar_receiver.uri` property

Type: `"string"`. Optional.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133): complete subsection reference.

<a id="canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- qradar_receiver.batch

<a id="canonical-2021331223032232-2031131311130103-3301232121321201-1332130100002112-1021223231300331-1012221203001221-3200312233310110-2012312002332330"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123320323132121-0033311311011122-0013032133202213-0212323230102322-3323203331231131-3000233201002323-3213310102121231-3323330333100332"></a>

### Direct properties for `qradar_receiver.batch`

<a id="canonical-2231321230023331-1121002301100021-0322132002100120-2020330002111131-2330022213303333-2201030021011220-0112100013113211-0120220123220003"></a>

#### `qradar_receiver.batch.max_bytes` property

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(4096, 10485760),
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](resources--global_log_receiver--reference--group-004.md#canonical-3330220031010112-1100331333200220-2202201021110333-2223131121033302-0011021222010200-0021322123202133-0130032213120222-0023203323022333): complete subsection reference.

<a id="canonical-0221001313121110-0023321013210211-2032322313000310-3123120023022101-1220312302130231-2132203230000030-1313000031021022-1322310123000023"></a>

<a id="canonical-3030033033330231-3112132030303012-0321112203111320-3330131222213210-1310003120130312-3313123332320311-3023222102030001-1232033300312101"></a>

#### `qradar_receiver.batch.max_events` property

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(32, 2000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0231123133131030-1120332010312103-1311321020022210-0321021200012200-1003323330323211-0322202210013001-1332013132130003-1111101101123031): complete subsection reference.

<a id="canonical-2331032123013232-3003003033001121-0103210000003232-0200310101230221-2210321100211313-3233002122232231-3123112222112103-3101231010010222"></a>

<a id="canonical-0032321023213010-2103332210330211-0100011022021210-2233331331120231-3103010202033323-3312023313332001-0122000321030220-3031003200122312"></a>

#### `qradar_receiver.batch.timeout_seconds` property

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-004.md#canonical-0221001212321231-3231100121130222-0012231333122312-3010022032021101-3232100330213302-2220020300102210-3322001231033302-0310300313133301): complete subsection reference.

<a id="canonical-3330220031010112-1100331333200220-2202201021110333-2223131121033302-0011021222010200-0021322123202133-0130032213120222-0023203323022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.max_bytes_disabled

<a id="canonical-2103021033123232-0320313103132311-2322133023112320-3121310223102001-1120102103223233-2210311211012223-3302003113103103-2211200031001120"></a>

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
max_bytes_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231123133131030-1120332010312103-1311321020022210-0321021200012200-1003323330323211-0322202210013001-1332013132130003-1111101101123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.max_events_disabled

<a id="canonical-2133021212221123-3010200123200301-3203012132120100-2212111332331322-2122200012111002-0102223202230033-0113232031322223-1221333203320033"></a>

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
max_events_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221001212321231-3231100121130222-0012231333122312-3010022032021101-3232100330213302-2220020300102210-3322001231033302-0310300313133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-004.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.timeout_seconds_default

<a id="canonical-1232330321113311-1212002031333133-1322323221022321-3132001303132001-3221310210013310-0230013132300213-1221312100201230-0112022330011132"></a>

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
timeout_seconds_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- qradar_receiver.compression

<a id="canonical-0300123130230212-0210212003303202-3200011320133021-0102110230222223-3030323320223300-0030201030311113-0031112321032233-0032000210221331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031310313332332-0013213312111003-3212111223022003-2033002210310113-2131223302012103-0110213221322323-3332113131332330-3123220001302000"></a>

### Direct properties for `qradar_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-004.md#canonical-2023010203320113-1101133302213122-2103033320220301-0022312303013010-3220223133220330-1320310321112312-1310331310220133-3132301032203110): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-004.md#canonical-1300332030231231-3321313200311020-3012231130131233-3333210332302203-3222313113213321-2121110030311330-1201303020301220-3021111103333011): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-004.md#canonical-1220131201021112-1303302203110223-3303230233131122-0013032331223111-0012230010212130-1010311132222223-1200130323213000-0001322221200002): complete subsection reference.

<a id="canonical-2023010203320113-1101133302213122-2103033320220301-0022312303013010-3220223133220330-1320310321112312-1310331310220133-3132301032203110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332)
- qradar_receiver.compression.compression_default

<a id="canonical-2222302002300322-1300312202111210-1330022331023221-3332012210021132-0023331300303200-1221133232321121-1120213331010002-0002222101001232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300332030231231-3321313200311020-3012231130131233-3333210332302203-3222313113213321-2121110030311330-1201303020301220-3021111103333011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332)
- qradar_receiver.compression.compression_gzip

<a id="canonical-3001210031003331-1120010333231111-2230100202002110-3233022230230201-2333303212222232-2133301011121123-0233131003303211-2202330010302200"></a>

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
compression_gzip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220131201021112-1303302203110223-3303230233131122-0013032331223111-0012230010212130-1010311132222223-1200130323213000-0001322221200002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-004.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332)
- qradar_receiver.compression.compression_none

<a id="canonical-0002020103233212-0203321120022222-1101230000013021-1210000210221231-0313110201001101-1330022333110020-0002023102313020-3011103322131212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333030003010101-1102331120011301-3023322233120221-1332221330232103-0110203322122232-0233200210022200-0211212203313232-2301331121322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- qradar_receiver.no_tls

<a id="canonical-2011101121131031-2021332220303111-3233201131210220-2333212131300233-1222222330201202-3101133121102322-2202021000020111-1232000333200131"></a>

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
no_tls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- qradar_receiver.use_tls

<a id="canonical-0233303223131023-0121300303011220-0100201323130200-1300032001302221-3020012001133222-1210003303303120-0323113312302012-3111113130232112"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303003230131221-1023122110210223-3003011203213323-2310001130003321-3210330120333320-2310003133123310-2312120111322222-2220313023330012"></a>

### Direct properties for `qradar_receiver.use_tls`

- [disable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-2333120331333210-3311110122102111-3023223323023311-2301120221203023-0223220232230133-3200301011013300-0000232301112110-3121001001213230): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-0102232122102220-0201011030200002-0222223111213123-3013233002212320-1330131302231033-1002113333111131-3303300221311223-0121323310030013): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-004.md#canonical-2033003220021233-1330221230121232-1323010030200330-1323303123223021-0031312031230332-3323131130232123-0201200013113003-3331010301212132): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-004.md#canonical-2122213023332330-0202131022100311-0032013001310020-3112202103301211-2233223121000000-0013023221121301-1330221213120020-3331223220113030): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-004.md#canonical-0231303000231222-0012321302102313-3320232202102121-1230321331103302-3332013001012123-0322021212130322-2133312302100130-1211020201302123): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-3000333320303323-0311232200022323-1011202110313110-1000221212122102-2111100022132032-0323023330020012-3130231202332110-2220202031131130): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-004.md#canonical-1312130321213230-0132213121032011-1330313022131132-0010313132020220-0011000010113303-1020122333111230-3200210122313121-1113132023223333): complete subsection reference.

<a id="canonical-3203122322221121-2201321003222222-1230012021220301-2120232002222320-0021310330311223-2012303332311020-3302203320221322-0211331323033011"></a>

<a id="canonical-3020211312233011-3210123021212203-0301233220303102-2200232203102111-1101022031113323-1012132223223032-3203021321311320-3100331200311322"></a>

#### `qradar_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2333120331333210-3311110122102111-3023223323023311-2301120221203023-0223220232230133-3200301011013300-0000232301112110-3121001001213230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.disable_verify_certificate

<a id="canonical-3132221033311332-2323012221020000-2013321311102230-1220011322023111-0312312321031013-1223332101310200-1012201010031103-1120133111212213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102232122102220-0201011030200002-0222223111213123-3013233002212320-1330131302231033-1002113333111131-3303300221311223-0121323310030013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.disable_verify_hostname

<a id="canonical-1033212320030331-3201021332321121-2302020221233001-2310321212020230-0231003101110212-3330222021122123-2030102302232313-0222020110102031"></a>

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
disable_verify_hostname = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033003220021233-1330221230121232-1323010030200330-1323303123223021-0031312031230332-3323131130232123-0201200013113003-3331010301212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.enable_verify_certificate

<a id="canonical-3103021300113221-1201123113100233-2023110312001011-2132002123200300-1033111222002122-0213210100230131-3203313001132003-3123301001121300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122213023332330-0202131022100311-0032013001310020-3112202103301211-2233223121000000-0013023221121301-1330221213120020-3331223220113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.enable_verify_hostname

<a id="canonical-3112203002022122-0301133021332332-0101300002121221-2010120311333130-3102203310230002-3113333021331332-3121230320032102-2220020032200330"></a>

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
enable_verify_hostname = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231303000231222-0012321302102313-3320232202102121-1230321331103302-3332013001012123-0322021212130322-2133312302100130-1211020201302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.mtls_disabled

<a id="canonical-0230303311011320-0020100002030302-2112211030302123-3221232200303131-1231103303320133-3120132313120331-0222031031000303-1030022003031001"></a>

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
mtls_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000333320303323-0311232200022323-1011202110313110-1000221212122102-2111100022132032-0323023330020012-3130231202332110-2220202031131130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.mtls_enable

<a id="canonical-1033311010032220-2002033203232031-0231230101001123-1332321122011200-1133023113122322-2103000101313132-0300230022201012-0322001220210311"></a>

Type: `"object"`. single nested block, Optional.

MTLS Client config allows configuration of mTLS client OPTIONS.

Receipt-pinned upstream constraints:

```json
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
mtls_enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103100131210030-1313013311013121-2320321013232010-1302220300112221-1202322110032300-1120121110132011-1300003102131220-3023312001213232"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable`

<a id="canonical-0220203200212013-1200132030213013-3210011303023001-3021010332011131-3312233122321001-1202120221113013-3101333210111132-0210203023120001"></a>

#### `qradar_receiver.use_tls.mtls_enable.certificate` property

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(100, 131072),
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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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

- [key_url](resources--global_log_receiver--reference--group-004.md#canonical-0332232320211011-1232213231321002-1111122113121030-0022220122322131-2012020022033220-3321222021032131-2033001001030102-0301333100030330): complete subsection reference.

<a id="canonical-0332232320211011-1232213231321002-1111122113121030-0022220122322131-2012020022033220-3321222021032131-2033001001030102-0301333100030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-3000333320303323-0311232200022323-1011202110313110-1000221212122102-2111100022132032-0323023330020012-3130231202332110-2220202031131130)
- qradar_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3131312021101220-3322023330132313-3213111223310020-3313223203311331-0012122022210103-3322313302210001-1211302031020211-3013233102201133"></a>

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
key_url {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011121113321113-1200023211122121-2231130331121100-1333212330212310-1321202032121112-1123332320101001-0020013103213012-1300313222301322"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2323221000322222-3033312100330310-1103211023102022-0011021003311203-3032012033000012-0130131023022100-3200323331323002-2302322010200232): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-004.md#canonical-2022312021032211-3333213111301231-0130111113321001-2222032122300130-0003121003223301-2000122202202202-3222023132210333-0110023213202232): complete subsection reference.

<a id="canonical-2323221000322222-3033312100330310-1103211023102022-0011021003311203-3032012033000012-0130131023022100-3200323331323002-2302322010200232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-3000333320303323-0311232200022323-1011202110313110-1000221212122102-2111100022132032-0323023330020012-3130231202332110-2220202031131130)
- [qradar_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-0332232320211011-1232213231321002-1111122113121030-0022220122322131-2012020022033220-3321222021032131-2033001001030102-0301333100030330)
- qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-3201330322110011-1122120022000322-1011330312003332-2012313211322231-3331301001213102-2220121100311011-0000032013210220-2323002210202103"></a>

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

<a id="canonical-1002220121012232-2013203200021330-3010230113122122-3323021211201201-1021223101120213-1100210311302210-3103301113120102-3220230133132230"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-3303212203111202-2102121000011223-3002213231303111-0310213010112313-1312000112002030-2032232210223201-2222110331102231-2030033313031101"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0232031330230232-0121032122111003-3113320202001120-0233202323123011-0111030210212103-3221230233230010-0201030132303121-0010111301131100"></a>

<a id="canonical-0121131000320003-2232130130020311-1302300031223031-2333130002310330-0120202023100113-1032002313211103-3313312230112013-0122313020232130"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1333010102230020-2030223210133111-1311300201320301-0230203011222320-1321001013121312-3210233220110021-2323100203312031-1121301030110113"></a>

<a id="canonical-3122130200010300-0112020320132133-2013112033033111-1201033201210031-1202000300003101-1212002210211301-0013030233103020-3131331322231231"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022312021032211-3333213111301231-0130111113321001-2222032122300130-0003121003223301-2000122202202202-3222023132210333-0110023213202232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- [qradar_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-004.md#canonical-3000333320303323-0311232200022323-1011202110313110-1000221212122102-2111100022132032-0323023330020012-3130231202332110-2220202031131130)
- [qradar_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-004.md#canonical-0332232320211011-1232213231321002-1111122113121030-0022220122322131-2012020022033220-3321222021032131-2033001001030102-0301333100030330)
- qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2322131313322332-2332013213120223-2223120133323332-3003330232130120-2120031001013012-3320301311330030-3232321232033132-1023321231033102"></a>

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

<a id="canonical-0111002221330210-3310011103121030-3222120023310022-1320101303201013-3112232102000120-1022121221300300-3300002312103022-3130101010223301"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-3213321001322213-1132332323301232-2320230022013020-2232131032130001-3120202211300101-3023311022031013-0113100133310323-2012131233033123"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1032221303020130-2323222203002113-2032002023121003-0030233233012332-3130223123210121-1210033102112101-1103020031221303-3203203221121121"></a>

<a id="canonical-1032213301120212-1322022322202103-0312032202111121-2300222203001230-2123323011101000-0321013213301322-3111102322233201-1302321111100022"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1312130321213230-0132213121032011-1330313022131132-0010313132020220-0011000010113303-1020122333111230-3200210122313121-1113132023223333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- qradar_receiver.use_tls.no_ca

<a id="canonical-1020301303303302-2101011112310022-3232101302031310-3232331233033113-2030101222310231-2233200312001313-1210000202203110-1322021030101110"></a>

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
no_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- request_logs

<a id="canonical-2112333332233110-1101303330101330-2203303321021123-0311130301133120-3113001330201322-1301233113213302-1011110010211003-1302210033220211"></a>

Type: `"object"`. single nested block, Optional.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("sampled",
    "unsampled")}
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
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

Terraform syntax:

```terraform
request_logs {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101013130110311-3121310221233100-1023330032121122-2011222210333332-2112121111022101-3020221320323321-0211120103322320-2020222232232302"></a>

### Direct properties for `request_logs`

- [sampled](resources--global_log_receiver--reference--group-004.md#canonical-1313122112022211-0110312121002231-1113111231112011-2202302021120223-2313213313013211-2312111122001332-0322233203002201-3202032030230231): complete subsection reference.

- [unsampled](resources--global_log_receiver--reference--group-004.md#canonical-3301312020133323-3210212220301032-0202023302301323-2103210011100233-2033320323101032-1030203023011323-2110002132211220-3123103212202202): complete subsection reference.

<a id="canonical-1313122112022211-0110312121002231-1113111231112011-2202302021120223-2313213313013211-2312111122001332-0322233203002201-3202032030230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs.sampled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211)
- request_logs.sampled

<a id="canonical-1021022020311001-3323233000000011-2103013131202320-1031110313222320-3233311331223111-1312102200230323-0012232110003332-0021002211001220"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
sampled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301312020133323-3210212220301032-0202023302301323-2103210011100233-2033320323101032-1030203023011323-2110002132211220-3123103212202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs.unsampled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [request_logs](resources--global_log_receiver--reference--group-004.md#canonical-1232021303231213-0303210110003213-3111010032312021-1233200302331232-1321133012013331-0030313121010031-3321002030223202-3213303122202211)
- request_logs.unsampled

<a id="canonical-1103211013210211-1302101123022230-1332332100130233-2221202000223123-1301322302322023-3122210110131101-1233121232030020-0112222322202020"></a>

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
unsampled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- s3_receiver

<a id="canonical-2120303131032111-1131001123333020-1133113103233030-3000301211202000-1233230031302020-3210203032233122-1203232331022303-3313332301220301"></a>

Type: `"object"`. single nested block, Optional.

S3 Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "bucket")}
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
s3_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103312030102000-0331230003011131-2232221202022320-1010313003201212-0321033221220212-1323333010323001-2001130211000131-3221200002021230"></a>

### Direct properties for `s3_receiver`

- [aws_cred](resources--global_log_receiver--reference--group-004.md#canonical-1233211220220100-0001322211012003-0320332330012233-2033003202130223-1320213011230001-1323200211322210-3012030221102032-3113222122322122): complete subsection reference.

<a id="canonical-1231020123112000-3323111222222231-2320212313010223-2130131013320113-1300100130331210-3311203012031022-0232212100010122-0320003113120012"></a>

<a id="canonical-2200130311320201-1300123313023220-0220120201103123-2233232222332231-1211001011021121-3130303232302122-3133101203313110-2113131102110202"></a>

#### `s3_receiver.aws_region` property

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](resources--global_log_receiver--reference--group-004.md#canonical-2212222233301133-1023002210322033-0231130330222102-0011130133000313-3300323323033222-2033232313203232-3013223202013210-1211323111232000): complete subsection reference.

<a id="canonical-2213211332001011-0210130101011103-3333113031120333-0330212030023231-3210021231310231-3333232301221232-0213000032320201-2122211111302231"></a>

<a id="canonical-1322301120211130-3112320013222313-0130303222322212-0333232203131011-2322120212013000-2132020031321013-1211233230021331-2232012012212113"></a>

#### `s3_receiver.bucket` property

Type: `"string"`. Optional.

S3 Bucket Name. S3 Bucket Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](resources--global_log_receiver--reference--group-005.md#canonical-2111330022112232-1022203012213032-2120111001213003-2321001220303202-0130300211313331-3300322120302121-3012320000202120-1022011011101303): complete subsection reference.

- [filename_options](resources--global_log_receiver--reference--group-005.md#canonical-0310300212301021-1030032312012230-0231233231132003-0310010220111031-0000031333323200-1103211313320310-2031132312223122-1323220303231200): complete subsection reference.

<a id="canonical-1233211220220100-0001322211012003-0320332330012233-2033003202130223-1320213011230001-1323200211322210-3012030221102032-3113222122322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.aws_cred` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- s3_receiver.aws_cred

<a id="canonical-0312003122213131-1120331302323313-1001103331301332-0012113311200000-3132021331332000-2221313211321311-0002300202023010-0130102311133032"></a>

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100030303301210-1103313232101032-2232202222111232-1122000123313011-2330231221333011-2031313023310203-0013100011210022-3300113113010213"></a>

### Direct properties for `s3_receiver.aws_cred`

<a id="canonical-2320023220212200-2201030100111002-2032130021232223-2331120031212231-1201332210300230-3221022220021223-2212100301330210-1130130003133121"></a>

#### `s3_receiver.aws_cred.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3121302031103131-2112223002112321-1102123021302003-0221033031223021-3032020203001311-2220122220302200-3121003331020211-3222203132331312"></a>

<a id="canonical-0323320022313033-1322113101003012-3230203322000200-1102102222002101-2231231313113003-2022310011000233-0300020133301023-3302110111021233"></a>

#### `s3_receiver.aws_cred.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3001233221333213-2002323321311022-2023202233032312-0023122213320032-0321132120212133-1203203210002031-2231302112121222-3103113301220212"></a>

<a id="canonical-2103022112122333-2302233100000313-0120131321232212-3222333320233313-0100120013013130-1212121332001111-3100100112301010-0221111013123033"></a>

#### `s3_receiver.aws_cred.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2212222233301133-1023002210322033-0231130330222102-0011130133000313-3300323323033222-2033232313203232-3013223202013210-1211323111232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [s3_receiver](resources--global_log_receiver--reference--group-004.md#canonical-2333111230121012-1211132001232103-0131221233103011-2311033122130021-2111100213333231-0022302131303130-2130300100030312-1232123130321122)
- s3_receiver.batch

<a id="canonical-2331332031200122-2130322330312232-2000302331232330-2033223120331123-2320231200033210-3123231100333112-1012133110021323-2203300101210020"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("max_bytes",
    "max_bytes_disabled"),
  validators.ConflictingObjectAttributes("max_events",
    "max_events_disabled"),
  validators.ConflictingObjectAttributes("timeout_seconds",
    "timeout_seconds_default")}
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
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

Terraform syntax:

```terraform
batch {
  # Configure direct properties listed below.
}
```
