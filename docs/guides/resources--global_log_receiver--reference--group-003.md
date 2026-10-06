---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- http_receiver.auth_basic.password

<a id="canonical-2230311112201102-3323031123201210-3131022111310113-0230111200302000-0102332121021221-3033302311313132-2103000101122122-0302230223110221"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203220313323012-1033203012123232-1332303202313330-3020132313130112-0010213032213103-2323030113210330-3301132103213200-1211331233033123"></a>

### Direct properties for `http_receiver.auth_basic.password`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0113112210111023-0200001333020232-3103313213132311-1113313220011133-2001121131112113-0123330111113002-3121122212223132-2331311211102223): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2033102031312132-3201120320220222-1001232213300230-1220223033333221-3130212312220010-0000222323110231-1201311013023031-1133302032233013): complete subsection reference.

<a id="canonical-0113112210111023-0200001333020232-3103313213132311-1113313220011133-2001121131112113-0123330111113002-3121122212223132-2331311211102223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-003.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- http_receiver.auth_basic.password.blindfold_secret_info

<a id="canonical-2021002230020210-3213022213201321-3122200021310323-1300123213003330-1031230133320031-1113102101322022-3311320011101222-2201201031200102"></a>

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

<a id="canonical-1003313111201220-0201323210320001-1003011123113302-3331200203100323-1101002131201003-0002321011030312-0022222120111300-0021323032221210"></a>

### Direct properties for `http_receiver.auth_basic.password.blindfold_secret_info`

<a id="canonical-0131101320302333-2033321030133302-2000122020012211-3203313102003303-0231332112302220-1202002000111002-0003132033201322-3123331203122103"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1210110320220213-3323201211311201-2310103123313231-3212301312020313-2303121021122321-2223203320012320-0033111313020231-0023112302212023"></a>

<a id="canonical-3303211301030111-2010002103320113-1310110323000113-2220001120012112-2002200213003320-1320301131012102-2302101000001031-3033123113131001"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.location` property

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

<a id="canonical-1311131111222112-0111123320012221-1331021102010331-1023121223022232-3013023233223233-3102233033202330-3332010321333311-3213020003323220"></a>

<a id="canonical-0131302300200133-2133232033002020-1021213313331200-3330000330201330-1202113331210101-0112003031123330-1201130010120302-3303223203002032"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-2033102031312132-3201120320220222-1001232213300230-1220223033333221-3130212312220010-0000222323110231-1201311013023031-1133302032233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_basic](resources--global_log_receiver--reference--group-002.md#canonical-1022130202001203-2123222000302011-0303000113321310-1202322232132113-3001121022210122-3232110303312101-0210032303121020-2131003101302233)
- [http_receiver.auth_basic.password](resources--global_log_receiver--reference--group-003.md#canonical-1033323112022211-3003201200310311-3010321003132310-1213101100232030-1130010321211202-3321121122113013-0110212000322033-2023303300123230)
- http_receiver.auth_basic.password.clear_secret_info

<a id="canonical-1211313223030302-2022231333303013-0313323321330023-1133333313300103-1203022230012210-0333200030333101-1131012131020200-2020310322000222"></a>

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

<a id="canonical-3123133033020131-0320003331100201-0110211220001303-3023000212322112-3121130102310003-1121232121311223-2032300022013001-2113012313111120"></a>

### Direct properties for `http_receiver.auth_basic.password.clear_secret_info`

<a id="canonical-2132111303200032-0332221123110312-0012220203030100-2310230010032203-0202311203302321-2132123331113100-2332320211112320-2311111032131001"></a>

#### `http_receiver.auth_basic.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1002101213231003-0333103100122231-0020001021011201-0303001003133230-3122132211032330-2132213113132321-3122221003313332-0211012210113013"></a>

<a id="canonical-3031112021030222-3100323133111113-0010202220222330-1201132123332212-1022023310333030-1330013222010012-0102310202032313-2111023111323102"></a>

#### `http_receiver.auth_basic.password.clear_secret_info.url` property

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

<a id="canonical-3331330032301033-3102210321202122-3212212113000230-1203031121110232-3301023022023030-1033012130200332-3101222021333330-1303103201031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.auth_none

<a id="canonical-0330311131133321-1321020021102231-3100320111330201-3120000200123211-2311302122121303-0123202330213223-1101203311333222-2323220133111231"></a>

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
auth_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.auth_token

<a id="canonical-2012332322000301-3321203200031301-3321100111121022-2020101301201213-1001121221231012-0220022312301333-3111310220310033-0133203012101113"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Receipt-pinned upstream constraints:

```json
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
auth_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200110323233011-0233030120033302-1313010002310022-0313302201331212-3012203320303110-0210023133213010-2033300020213111-3213201101320132"></a>

### Direct properties for `http_receiver.auth_token`

- [token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303): complete subsection reference.

<a id="canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-003.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- http_receiver.auth_token.token

<a id="canonical-0131230233221123-0210000220011330-2200211322003100-2010223233331112-1011123203021122-1330203220222010-3012113320120003-1100102030031312"></a>

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
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100221031302002-3103000020331023-3211120030332310-1103021303002301-2213230010203302-2331202020311211-1001333201002211-0320221230223331"></a>

### Direct properties for `http_receiver.auth_token.token`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0123330131101110-0103033001032200-1320230012330300-0013302203332113-1132001201020112-1322112031203010-2130110012222121-1200220212133301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0213303202300301-1320003013221212-1110021133210231-2011110012312300-1311220332311202-2131222032312013-0033110312311330-0203003002220332): complete subsection reference.

<a id="canonical-0123330131101110-0103033001032200-1320230012330300-0013302203332113-1132001201020112-1322112031203010-2130110012222121-1200220212133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-003.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- http_receiver.auth_token.token.blindfold_secret_info

<a id="canonical-2112223021201212-1010230300212213-3020312332111121-1001112113012121-3103031331131232-3002303311132311-2023030133030010-2333122031033021"></a>

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

<a id="canonical-0211321321121310-2133213123113203-0013300212223123-3223001232000312-1030100032111133-0033010232031002-2131222110331333-3212330332300121"></a>

### Direct properties for `http_receiver.auth_token.token.blindfold_secret_info`

<a id="canonical-3123332030303322-0333332022221001-3010012223232020-1331112312031212-0223112013020021-0101011220302200-2212233222223002-1312100321120202"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3310223222121231-1131132202102230-1310202233321101-0213310220323010-1113300313331220-1322311301012033-1002033302020012-2233033112133330"></a>

<a id="canonical-2120320220021030-3012310123011321-3313010332103233-1321231001101102-3112033031120232-3312131123131230-3002212101003020-0033123122223230"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.location` property

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

<a id="canonical-2001221103021232-0011021201002002-2031102023312131-1233123122033321-2031101111202322-3311200211111010-1312030003021200-1223013333302112"></a>

<a id="canonical-2102000222012201-3322301321133021-1203223212031310-1230202221101122-0313233011231003-2030130302311222-2131022202002320-0231320001223031"></a>

#### `http_receiver.auth_token.token.blindfold_secret_info.store_provider` property

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

<a id="canonical-0213303202300301-1320003013221212-1110021133210231-2011110012312300-1311220332311202-2131222032312013-0033110312311330-0203003002220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-003.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- http_receiver.auth_token.token.clear_secret_info

<a id="canonical-1230002020101230-3213130320231220-3023330210001232-0100311230200232-2330020313302103-2033101013203312-1320330133301222-3021022012232101"></a>

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

<a id="canonical-3323110202333231-2021000111210102-1232031021113232-3210313310232122-3332101100333003-3001320320032200-2032222201210310-0100032032031302"></a>

### Direct properties for `http_receiver.auth_token.token.clear_secret_info`

<a id="canonical-2130322330011222-0322330300210311-0010330132222131-0232303131001220-1233012313233320-3133201010232233-2321332302221110-0110010300322100"></a>

#### `http_receiver.auth_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202222213300333-0303331323101133-3220120201323023-1022130301020213-3001211023030301-2311210223121302-3323332221201313-3033201201032123"></a>

<a id="canonical-1113133213300032-1103233332011313-2332333120121011-0123132112010202-0321321203333001-3302323130232310-1222023310223231-0013310112112111"></a>

#### `http_receiver.auth_token.token.clear_secret_info.url` property

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

<a id="canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.batch

<a id="canonical-0212110113100101-1302102130100131-1303323102101031-3232110012311112-0332330121310131-0012223323323212-3000230100311332-1102312133320213"></a>

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

<a id="canonical-2103123202002112-0312111031011030-1003301001033020-0030202023121321-1221020232111231-3100031102130013-3233000023100232-1302313300313331"></a>

### Direct properties for `http_receiver.batch`

<a id="canonical-1011130200212131-1031002221122100-2333331203030202-1203003301221012-1210202213303201-0022220303001111-2133022002121120-0211030300113303"></a>

#### `http_receiver.batch.max_bytes` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2322030211000003-1133331020121313-1121222323220301-3130300131123301-3032122333332210-3112230013110020-2112123233131011-3130320110302101): complete subsection reference.

<a id="canonical-2301003313322031-0102013323111303-0232301233031212-3021103130331010-0030121121110133-0212002122230202-3313322010330111-2031202222032001"></a>

<a id="canonical-1212232102010231-2213123231111212-3010101031232121-2212211333211213-2220231033201231-2012013033131212-3200123231212331-0322003300023132"></a>

#### `http_receiver.batch.max_events` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-3220021331302101-0030133213132003-1021212332230013-3031333022301012-0101300112312312-0010133220330101-2030232333313023-0131102031131132): complete subsection reference.

<a id="canonical-3100133332212011-2213223013320023-1130102000301002-1032212303000021-3230311100310022-0301330113000323-1130010203123120-0332202030320101"></a>

<a id="canonical-1202000001312212-1200102002310133-1232133333003323-1121322011223322-3122210011003321-1311230010133120-3233231301011201-2001323101030321"></a>

#### `http_receiver.batch.timeout_seconds` property

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-3300203101200123-1312323331101210-1323030333211102-1201111030322333-1130001230332211-0010232203023321-3300120322112101-3001022231301002): complete subsection reference.

<a id="canonical-2322030211000003-1133331020121313-1121222323220301-3130300131123301-3032122333332210-3112230013110020-2112123233131011-3130320110302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.max_bytes_disabled

<a id="canonical-0220301012131310-2301132300103330-2312200312101001-0122123022220320-1322302021332103-0231112011112032-3020211031113210-1030010001310003"></a>

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

<a id="canonical-3220021331302101-0030133213132003-1021212332230013-3031333022301012-0101300112312312-0010133220330101-2030232333313023-0131102031131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.max_events_disabled

<a id="canonical-0310223123233113-1323110321230333-3030023020121023-1121322110323203-2222130230222300-0211131330110102-0231311120221212-1123102303202300"></a>

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

<a id="canonical-3300203101200123-1312323331101210-1323030333211102-1201111030322333-1130001230332211-0010232203023321-3300120322112101-3001022231301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.timeout_seconds_default

<a id="canonical-0003000020012230-3121131132031020-0330320302100101-2331132010303000-1013332000120020-1200233231002330-0013311030322021-1210220012111202"></a>

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

<a id="canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.compression

<a id="canonical-2320032320122303-0303033220030200-0133322101222122-2332112103133323-1031320012123010-3021022303023200-3033100330132223-1013313201331232"></a>

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

<a id="canonical-0113032231001320-1210000123302021-1011021020330323-1333121122013220-0201323232123301-1223221031323212-2011223001101302-3033312030001023"></a>

### Direct properties for `http_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-3121131032201023-3120023111002323-2021100132311133-2300031310212020-3322101102033132-1021330020321021-0003130332231200-2220021100001321): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3111112222223130-0013112302123330-0330123230332132-3222102103213230-2200232102200200-2313022032003010-3030133212112021-0331310013122101): complete subsection reference.

<a id="canonical-3121131032201023-3120023111002323-2021100132311133-2300031310212020-3322101102033132-1021330020321021-0003130332231200-2220021100001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_default

<a id="canonical-0133101033000002-0330202101311321-3213323312101301-1222030100133200-0210101213000132-0213222333131120-2210132333020301-3033320000133230"></a>

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

<a id="canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_gzip

<a id="canonical-3331323200213231-0120103201102010-3030010331020322-3301000120331132-1333300211330202-1000021222121300-0123220232212310-3321102011200311"></a>

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

<a id="canonical-3111112222223130-0013112302123330-0330123230332132-3222102103213230-2200232102200200-2313022032003010-3030133212112021-0331310013122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_none

<a id="canonical-0330111322330010-3200220201223002-0220333222222333-0033310230333203-1300333130231331-2320021113221330-3210100322310122-0231133303220313"></a>

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

<a id="canonical-3233100233103022-2230123213321331-0131001032100301-2102322211220020-3112230003112133-0301020101101012-1200122011110232-2212133013103130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.no_tls

<a id="canonical-0021010300200121-0020302101130303-0001312133311002-3122123120132133-2221222132121130-0320300113233120-0302113221030121-1311233111321202"></a>

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

<a id="canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.use_tls

<a id="canonical-3122200203231003-0022232232103320-0303132131122332-2323323332112310-2013232023300220-2100003022111130-0030001321112012-3333303112121312"></a>

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

<a id="canonical-1023003020333200-1230121113120032-2300033000331232-2000001023102222-2203220132231100-1202203323010110-1301200011230200-1311100223200133"></a>

### Direct properties for `http_receiver.use_tls`

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-2110120122310202-1332022131132123-1131211233223300-1222312231230211-2111323011102032-3133021030310103-2301122300231111-2203131131213122): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-1231001233202010-1210001002202323-3033112010133000-0121002321001212-1120203110200212-0320020031131021-3311023231013002-0232221133323013): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1133202021221033-0230201020322130-3021222322011200-2002111022000311-2323031330013132-1333000323213301-3112033113010021-2301213003223133): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0000331020321020-0222333100313211-0202111130321023-3002212033113023-0111232022012231-0331202121233212-3320110001122311-2120300002033210): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2113133323211020-0012300213221300-3100102220320103-2313201312310031-3123311022110100-2223121313133033-0321211200233002-0133223110003223): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-003.md#canonical-2120210232111220-1103201201203332-3232310103113313-1130232323321000-3013313330012101-1033233021113012-2131223202233103-2303310321302000): complete subsection reference.

<a id="canonical-0322301231001023-0020030312030110-3333311010310111-2011233120031230-0312112212300303-1212120203222010-0210231301320032-1311300003232233"></a>

<a id="canonical-3313011121202232-2303110200201132-0011011211033331-0212322133320030-2000210131132010-0211211120201030-1122233320111102-2212310230022011"></a>

#### `http_receiver.use_tls.trusted_ca_url` property

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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2110120122310202-1332022131132123-1131211233223300-1222312231230211-2111323011102032-3133021030310103-2301122300231111-2203131131213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.disable_verify_certificate

<a id="canonical-1230011120133333-2210331123002101-0120223303323213-3030111023010000-2302201202310220-0020102232202211-0302212032330112-3132101131322313"></a>

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

<a id="canonical-1231001233202010-1210001002202323-3033112010133000-0121002321001212-1120203110200212-0320020031131021-3311023231013002-0232221133323013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.disable_verify_hostname

<a id="canonical-0022121122120312-0122300132022022-2330301012310123-1221232131121101-1111300323132131-0023130331302132-2210012200321021-1021301323222102"></a>

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

<a id="canonical-1133202021221033-0230201020322130-3021222322011200-2002111022000311-2323031330013132-1333000323213301-3112033113010021-2301213003223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.enable_verify_certificate

<a id="canonical-3213311211122021-1320312311322031-2301211010222032-0102000201011000-3203133233332110-3302200313110033-0330323133000012-3303030102000300"></a>

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

<a id="canonical-0000331020321020-0222333100313211-0202111130321023-3002212033113023-0111232022012231-0331202121233212-3320110001122311-2120300002033210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.enable_verify_hostname

<a id="canonical-2120111331123120-1002033032002023-0131233322331330-0313020103112031-2002032103321023-2001311012321123-3110020202323100-3331203102300232"></a>

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

<a id="canonical-2113133323211020-0012300213221300-3100102220320103-2313201312310031-3123311022110100-2223121313133033-0321211200233002-0133223110003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.mtls_disabled

<a id="canonical-3331313133020210-2300121120320213-2012301031302101-1322313203221300-3011302131322310-3212233300222201-3210302323022321-0031130012301300"></a>

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

<a id="canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.mtls_enable

<a id="canonical-3231003303012101-2233300303223221-1121012330210330-2201112223110213-0003120200131103-0321211230222330-2303322010110330-0033120010203102"></a>

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

<a id="canonical-2323323020003211-2010230102200002-1022233000103211-3020312311233021-0202122230103032-1021213110110201-3201100103100301-2333130010113200"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable`

<a id="canonical-0130333200031221-0322220233213003-1133311112130313-0133212012212003-0320123202121101-3202300113203131-2310203202012121-0121112210310123"></a>

#### `http_receiver.use_tls.mtls_enable.certificate` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002): complete subsection reference.

<a id="canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210)
- http_receiver.use_tls.mtls_enable.key_url

<a id="canonical-1203310121332211-2031210130122101-2203021230311110-1032303331233321-2231121333000000-3033203211001032-3110133020022220-0210222320222230"></a>

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

<a id="canonical-3100101311130323-1022032131000111-3032112030020121-3000123310313130-3203012222203112-3201302211223123-0320131133010313-1103323102012310"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2030301030133032-3212313313313131-0102303132011211-3211031113330322-0002232331331111-2200323102110311-0010223001300213-1023002100111301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-3203132230022020-1103113201022030-0103231013222221-3023113032323020-1223023231303112-0320133132211312-3233300033200132-0313130222311332): complete subsection reference.

<a id="canonical-2030301030133032-3212313313313131-0102303132011211-3211031113330322-0002232331331111-2200323102110311-0010223001300213-1023002100111301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210)
- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002)
- http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0022203101211020-0200100002133300-0221010012112322-1011202322012233-0110212003132312-1121001113230230-1323301132022113-0111001211100333"></a>

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

<a id="canonical-1223223231020113-2112113220111221-3323220112333032-3113122123202231-1130121201332012-3033303231031011-1132312312322022-0012311122133012"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-0112232210120133-3213301232011013-3231220220331231-3030331132002303-0102012303313130-3131221023011121-3313203031023111-2300130230310220"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3120011313320303-3003212131102012-1013001303032302-3122102110110310-0333022132313302-2000031312001100-0103303130323313-0021333300032101"></a>

<a id="canonical-1232213101303002-0202301111231213-0223133102320130-3010202310113200-1302101202000103-3213110010323002-2103032132112103-0103012013010032"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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

<a id="canonical-1111000231221000-1311332101130212-1302033300010021-2001110120022230-3220032310323123-1020222213200311-3333103301022103-0020231211201231"></a>

<a id="canonical-2302320101210100-2123120130012131-3032222121220121-3101311213120003-1131301312100331-1220023030033013-1310300013323131-0011221201120013"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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

<a id="canonical-3203132230022020-1103113201022030-0103231013222221-3023113032323020-1223023231303112-0320133132211312-3233300033200132-0313130222311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210)
- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002)
- http_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-0021311210103002-0323230322231121-2120030013223311-1013221000330112-0332013301220101-2002023333032331-0132021210201233-2310230233212032"></a>

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

<a id="canonical-0300122000332230-2112102112231302-3303002100133123-2301233031333130-2023111132222303-2023130101331121-2202321123132011-2033312132122231"></a>

### Direct properties for `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-0031130110321102-2022101230002202-2233030121102322-3332312301231202-2221231021010223-1330311113100312-2113233123330231-1330220202100231"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0212300023232033-2132302113320130-2303011312333112-2101330321230200-1031122030131022-1110202122133312-0333211003203013-1110123330022321"></a>

<a id="canonical-3121302131321221-0030100311110213-2301233020030032-0132231220211312-0131032213332302-2132221002010102-3211002133001023-0203003121311233"></a>

#### `http_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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

<a id="canonical-2120210232111220-1103201201203332-3232310103113313-1130232323321000-3013313330012101-1033233021113012-2131223202233103-2303310321302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.no_ca

<a id="canonical-0320303010223313-3103100300103013-0020011220310123-0110033013323222-3233313331203132-2013021030122303-1332023323110030-1211131230331202"></a>

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

<a id="canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- kafka_receiver

<a id="canonical-0300211310300033-3303232313232332-0103302032102312-3331310202130303-2112110331221010-0120101202030011-0132333301330000-1003221223001012"></a>

Type: `"object"`. single nested block, Optional.

Kafka Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("bootstrap_servers",
    "kafka_topic"),
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
kafka_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300311301010212-2202213301230001-3121310223022110-3030013310003320-0232122230002003-0031222310100321-3333120300310222-2230120312000202"></a>

### Direct properties for `kafka_receiver`

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311): complete subsection reference.

<a id="canonical-1221210133303301-3131320221033213-0113123330213102-3023200111131233-3121001032200013-3332111022231122-1212311222202203-3203321121301323"></a>

<a id="canonical-1121333112311200-1100032003223212-2230030323032133-3032102212022220-3331011110212000-0211013213231000-1320112000033230-2211033000331220"></a>

#### `kafka_receiver.bootstrap_servers` property

Type: `["list", "string"]`. Optional.

List of host:port pairs of the Kafka brokers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332): complete subsection reference.

<a id="canonical-0202200230302332-0301131301312301-1223322023220123-3203110321321211-1011233122201132-3133111111300230-1213001210221303-2311112233131130"></a>

<a id="canonical-0203223120213122-0311112101322220-2201210223030023-1020113201233331-0232021331312021-2313221123000222-0323210110202030-0222220101002112"></a>

#### `kafka_receiver.kafka_topic` property

Type: `"string"`. Optional.

The Kafka topic name to write events to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  }
}
```

- [no_tls](resources--global_log_receiver--reference--group-003.md#canonical-1330000103221001-0032201311231232-0021202030102312-1103323120300010-3132201012031130-2133020310001133-2300230132200301-1121212323312103): complete subsection reference.

- [use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301): complete subsection reference.

<a id="canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.batch

<a id="canonical-2112032212232221-1030333002202212-3302323323000212-3223031130103302-0320112232021330-2031202301002133-2023013122212303-2203133212301110"></a>

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

<a id="canonical-0133001333222301-1020302302212033-3331203310230333-1012220323310113-1311322310300131-3213011113011232-2330303332203100-3213132310112223"></a>

### Direct properties for `kafka_receiver.batch`

<a id="canonical-0101133130201200-1121011031221020-0200210112230102-3110312333233130-2200301302311333-2003223010022113-0113132102123121-2021313123302121"></a>

#### `kafka_receiver.batch.max_bytes` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0003310101230200-2130120313100023-2211313123000121-2330130111033320-3322020122222130-1000110000230021-2101203033300220-2212100131003233): complete subsection reference.

<a id="canonical-0010331020112202-1032322032022223-1210221330203203-0001310323002311-3020201023311000-3122001301122012-0102223123101210-2213212330031211"></a>

<a id="canonical-2111300031103311-2330313311013313-2031332200113231-2310102030220013-2103132101010311-1000002210012213-1130212103030012-3112231120002312"></a>

#### `kafka_receiver.batch.max_events` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-1230133223101300-2113013232301220-2233322030222130-1023301322013120-3331302032323220-0021120203310212-1302001220033202-3221213232022100): complete subsection reference.

<a id="canonical-1013211011223101-2233003023220001-0321233112232011-3003130032212003-3032213022121203-2212323200133201-1313111212102331-1110232013333200"></a>

<a id="canonical-0023210122023220-3301220010223121-3112001032010120-0201211132332020-1233010232023010-2020302033312202-0232113313022311-3000020030313301"></a>

#### `kafka_receiver.batch.timeout_seconds` property

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-0230322320212102-1231230032123311-2012002232222313-3231003111003330-1012202300222321-2210102220210002-3023113333111130-0000302233010212): complete subsection reference.

<a id="canonical-0003310101230200-2130120313100023-2211313123000121-2330130111033320-3322020122222130-1000110000230021-2101203033300220-2212100131003233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.max_bytes_disabled

<a id="canonical-1112221213102031-1313102331032021-1233312122213201-3130030222333303-2231321313213010-3311023220100012-2333313321220221-1132100101030221"></a>

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

<a id="canonical-1230133223101300-2113013232301220-2233322030222130-1023301322013120-3331302032323220-0021120203310212-1302001220033202-3221213232022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.max_events_disabled

<a id="canonical-0222332301112112-1323203030202220-3103013022033203-2203320200333023-2002213231012111-0103303121230030-2030021311013231-1120011333102213"></a>

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

<a id="canonical-0230322320212102-1231230032123311-2012002232222313-3231003111003330-1012202300222321-2210102220210002-3023113333111130-0000302233010212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.timeout_seconds_default

<a id="canonical-2133331331213111-0313110103103201-3312202113031130-3010311000320231-3303103330103232-0330022211321202-1033202123303312-3202000032132202"></a>

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

<a id="canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.compression

<a id="canonical-2131322222331221-2110213002112031-2300320001322210-2100201113112232-3020311002230033-0331111311303331-2102230230213332-3232323222033100"></a>

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

<a id="canonical-0033023233200111-1201120310113013-1233133010123102-3231011031223222-2013210230230101-2103222212223310-1033120132200101-0220020200202012"></a>

### Direct properties for `kafka_receiver.compression`

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-3031323100032313-2023232013230033-2100002230032211-2020003211230133-2321121233030020-0220203330013322-1011332031230001-3001333311223033): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3110103223213301-2003030032302313-1022210232002202-2133102320302211-0201003122230031-1001320300222100-3121022001220103-2021113120311332): complete subsection reference.

<a id="canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_default

<a id="canonical-1013110231132100-3313101233013302-1003030230003122-2320120111230223-3011321321222110-0231312100130222-1302320101103312-2200021321003113"></a>

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

<a id="canonical-3031323100032313-2023232013230033-2100002230032211-2020003211230133-2321121233030020-0220203330013322-1011332031230001-3001333311223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_gzip

<a id="canonical-2210200010322033-3031130130333300-1313021023220310-3230222301121132-2002200331333212-1212210012013101-2332333102132003-3131100231103133"></a>

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

<a id="canonical-3110103223213301-2003030032302313-1022210232002202-2133102320302211-0201003122230031-1001320300222100-3121022001220103-2021113120311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_none

<a id="canonical-3300301221313200-1310321202302222-3132011210312001-2113201212123201-0301003020122321-0323123132021133-2103330202211001-2233301023303031"></a>

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

<a id="canonical-1330000103221001-0032201311231232-0021202030102312-1103323120300010-3132201012031130-2133020310001133-2300230132200301-1121212323312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.no_tls

<a id="canonical-0003333312013332-3312332113032023-2100333310320231-2212012030331222-0332123013120001-3011300232011122-3313133201230202-1010223222000111"></a>

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

<a id="canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.use_tls

<a id="canonical-3121332102100013-3310211230003011-3310331311200112-1211202111323012-2222010132122201-3023133030110011-0231031202120021-1333213123313031"></a>

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

<a id="canonical-1011113323123111-3202022222012312-2322110032213010-0231122203102221-2213023131000322-2331220321132131-3101130212012233-0222300332330303"></a>

### Direct properties for `kafka_receiver.use_tls`

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1301020211212020-3131111102133200-3310333122011211-1033221303012201-3232233001223023-3233032333010022-1300012102303230-3300310032101103): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-3133203010121333-0000113013133311-2323220112113300-2211322131132233-1100300223320011-2123321300011212-0203202131313121-2321021132002121): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-3230223102001002-2231333030231132-3300011300021002-2310030320032020-2103331312002231-1002213323221322-1322331011101233-1203101002021100): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2232220303301202-0101201302302110-0110333020303032-1101131001003010-3330111311001030-3110312323332233-3110221302232133-2031012223010222): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-004.md#canonical-1003222302113131-3020001323322130-2003312222122100-2203013200001021-3223010100323333-2303133002013033-1011312223111120-0302222010210022): complete subsection reference.

<a id="canonical-1112112120233110-3303111231021013-2200102232013332-1032310111313300-1032112202230210-2011203310021331-1010323311201201-0203220001310113"></a>

<a id="canonical-1130230033013033-1311100210102223-0311121001231213-1021323333113122-0003110231032022-0323002110330302-0230001303033303-1133202312213221"></a>

#### `kafka_receiver.use_tls.trusted_ca_url` property

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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1301020211212020-3131111102133200-3310333122011211-1033221303012201-3232233001223023-3233032333010022-1300012102303230-3300310032101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.disable_verify_certificate

<a id="canonical-1202023110110032-3221310313122100-0212010202123132-0122012331332122-0210030212120221-3320302320312132-1202313320013200-3223321213333002"></a>

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

<a id="canonical-3133203010121333-0000113013133311-2323220112113300-2211322131132233-1100300223320011-2123321300011212-0203202131313121-2321021132002121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.disable_verify_hostname

<a id="canonical-2133213212331130-1121230221302202-2111100110221222-2200233210033212-1213321202112010-3132332201333301-2000101001220323-0023332022211023"></a>

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

<a id="canonical-3230223102001002-2231333030231132-3300011300021002-2310030320032020-2103331312002231-1002213323221322-1322331011101233-1203101002021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="canonical-3111112213202003-3033300013212303-1103310011310113-0122130302033030-0123111122013221-3332120120023322-1310202231133012-2121012330013210"></a>

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

<a id="canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.enable_verify_hostname

<a id="canonical-1101213130230231-0301030231020130-1330011000310103-2302233213111010-0033311330330220-0201111213202032-3102003032230123-3123121232303120"></a>

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

<a id="canonical-2232220303301202-0101201302302110-0110333020303032-1101131001003010-3330111311001030-3110312323332233-3110221302232133-2031012223010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.mtls_disabled

<a id="canonical-1130120021222200-3202130213203011-0112112223132001-3002320033321222-2032311310211330-1010231033001120-0021320233030133-3111230222120010"></a>

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

<a id="canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kafka_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.mtls_enable

<a id="canonical-1131130022323103-1203321331130201-1310103333213203-3002020102120011-2001002122211211-0233313013200330-0223023330323000-3113200013003032"></a>

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
