---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-0200110323233011-0233030120033302-1313010002310022-0313302201331212-3012203320303110-0210023133213010-2033300020213111-3213201101320132"></a>

## http_receiver.auth_token — auth_token / 032220001222 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.auth_token

<a id="canonical-2012332322000301-3321203200031301-3321100111121022-2020101301201213-1001121221231012-0220022312301333-3111310220310033-0133203012101113"></a>

Type: `"object"`. single nested block, Optional.

Access Token. Authentication Token for access.

Upstream description:

Authentication Token for access.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0000310223010133-2222001000331010-2322133202323111-1223020310213210-1301320022131103-3302133203030320-2131011002132302-2232010223102230"></a>

## Direct properties — auth_token / 032220001222 / 3

- [token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303): complete subsection reference.

<a id="canonical-0101221203301122-3311113233322231-3010101313311011-0321002020021133-0001032020223210-2133210200133202-1310222013313302-1221123132203213"></a>

## Next pages — auth_token / 032220001222 / 4

- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100221031302002-3103000020331023-3211120030332310-1103021303002301-2213230010203302-2331202020311211-1001333201002211-0320221230223331"></a>

## http_receiver.auth_token.token — token / 202122312202 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- http_receiver.auth_token.token

<a id="canonical-0131230233221123-0210000220011330-2200211322003100-2010223233331112-1011123203021122-1330203220222010-3012113320120003-1100102030031312"></a>

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
token {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231020132030131-0003303110313203-1103111101303221-0203320300221302-1313201013320103-1303012011030133-1101120103102122-2032132120220121"></a>

## Direct properties — token / 202122312202 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0123330131101110-0103033001032200-1320230012330300-0013302203332113-1132001201020112-1322112031203010-2130110012222121-1200220212133301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0213303202300301-1320003013221212-1110021133210231-2011110012312300-1311220332311202-2131222032312013-0033110312311330-0203003002220332): complete subsection reference.

<a id="canonical-3321022300031012-0323111012112132-1301332320020022-0033312131203212-3102103101012123-1010232313020032-3312223013301013-2000310123022113"></a>

## Next pages — token / 202122312202 / 4

- [http_receiver.auth_token.token.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0123330131101110-0103033001032200-1320230012330300-0013302203332113-1132001201020112-1322112031203010-2130110012222121-1200220212133301)
- [http_receiver.auth_token.token.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0213303202300301-1320003013221212-1110021133210231-2011110012312300-1311220332311202-2131222032312013-0033110312311330-0203003002220332)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0123330131101110-0103033001032200-1320230012330300-0013302203332113-1132001201020112-1322112031203010-2130110012222121-1200220212133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211321321121310-2133213123113203-0013300212223123-3223001232000312-1030100032111133-0033010232031002-2131222110331333-3212330332300121"></a>

## http_receiver.auth_token.token.blindfold_secret_info — blindfold_secret_info / 021020023221 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- http_receiver.auth_token.token.blindfold_secret_info

<a id="canonical-2112223021201212-1010230300212213-3020312332111121-1001112113012121-3103031331131232-3002303311132311-2023030133030010-2333122031033021"></a>

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

<a id="canonical-2120320220021030-3012310123011321-3313010332103233-1321231001101102-3112033031120232-3312131123131230-3002212101003020-0033123122223230"></a>

## Direct properties — blindfold_secret_info / 021020023221 / 3

<a id="canonical-3123332030303322-0333332022221001-3010012223232020-1331112312031212-0223112013020021-0101011220302200-2212233222223002-1312100321120202"></a>

<a id="canonical-2102000222012201-3322301321133021-1203223212031310-1230202221101122-0313233011231003-2030130302311222-2131022202002320-0231320001223031"></a>

## decryption_provider property — blindfold_secret_info / 021020023221 / 4

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

<a id="canonical-3310223222121231-1131132202102230-1310202233321101-0213310220323010-1113300313331220-1322311301012033-1002033302020012-2233033112133330"></a>

<a id="canonical-2333010233121110-2012123110033201-3003223011013301-0111022130310322-1033203113233133-3103201223023111-1210112012021222-0011132322010021"></a>

## location property — blindfold_secret_info / 021020023221 / 5

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

<a id="canonical-2001221103021232-0011021201002002-2031102023312131-1233123122033321-2031101111202322-3311200211111010-1312030003021200-1223013333302112"></a>

<a id="canonical-3201120320131102-1221011112302320-0233313331102123-3230321033110302-0321211310321022-0121333320033100-3330211131130310-3011311312001010"></a>

## store_provider property — blindfold_secret_info / 021020023221 / 6

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

<a id="canonical-2011000002331202-0222120030301331-1001322333021230-2301323232122220-0133233300101130-0120132201303113-2120002113033200-1001011103312132"></a>

## Next pages — blindfold_secret_info / 021020023221 / 7

- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0213303202300301-1320003013221212-1110021133210231-2011110012312300-1311220332311202-2131222032312013-0033110312311330-0203003002220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323110202333231-2021000111210102-1232031021113232-3210313310232122-3332101100333003-3001320320032200-2032222201210310-0100032032031302"></a>

## http_receiver.auth_token.token.clear_secret_info — clear_secret_info / 132323323323 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.auth_token](resources--global_log_receiver--reference--group-002.md#canonical-3211110223222312-1122132110030223-3113232133213102-3013031131221331-1303210011032221-3303131311330213-2103031022323113-1102013133333303)
- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- http_receiver.auth_token.token.clear_secret_info

<a id="canonical-1230002020101230-3213130320231220-3023330210001232-0100311230200232-2330020313302103-2033101013203312-1320330133301222-3021022012232101"></a>

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

<a id="canonical-1113133213300032-1103233332011313-2332333120121011-0123132112010202-0321321203333001-3302323130232310-1222023310223231-0013310112112111"></a>

## Direct properties — clear_secret_info / 132323323323 / 3

<a id="canonical-2130322330011222-0322330300210311-0010330132222131-0232303131001220-1233012313233320-3133201010232233-2321332302221110-0110010300322100"></a>

<a id="canonical-1311233113300222-3112003221000123-0002202031102113-3321231012022101-1311032112221303-3210331201330122-2033111201312133-0333301333333003"></a>

## provider_ref property — clear_secret_info / 132323323323 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1202222213300333-0303331323101133-3220120201323023-1022130301020213-3001211023030301-2311210223121302-3323332221201313-3033201201032123"></a>

<a id="canonical-0002010130313001-2331302121313013-3232301323332202-0230120223321221-2331302201330220-3322023332123333-2210321200132032-2333000102032200"></a>

## URL property — clear_secret_info / 132323323323 / 5

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

<a id="canonical-1113323003100210-0201213331031010-1121223301212021-1112001032133123-3232102133300001-0023301131302112-0301322022012100-3022200013220121"></a>

## Next pages — clear_secret_info / 132323323323 / 6

- [http_receiver.auth_token.token](resources--global_log_receiver--reference--group-003.md#canonical-3331131010132202-3331021223111020-0031221332212333-3202200200110312-2230332233321023-3133302002230120-1102321011201130-3332210213223303)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103123202002112-0312111031011030-1003301001033020-0030202023121321-1221020232111231-3100031102130013-3233000023100232-1302313300313331"></a>

## http_receiver.batch — batch / 012320032121 / 2

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

<a id="canonical-1212232102010231-2213123231111212-3010101031232121-2212211333211213-2220231033201231-2012013033131212-3200123231212331-0322003300023132"></a>

## Direct properties — batch / 012320032121 / 3

<a id="canonical-1011130200212131-1031002221122100-2333331203030202-1203003301221012-1210202213303201-0022220303001111-2133022002121120-0211030300113303"></a>

<a id="canonical-1202000001312212-1200102002310133-1232133333003323-1121322011223322-3122210011003321-1311230010133120-3233231301011201-2001323101030321"></a>

## max_bytes property — batch / 012320032121 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212320011002120-1300202330033132-2303323031311023-3023231312200333-0222101222213200-0301310013111101-1121103110203212-0230321233122231"></a>

## max_events property — batch / 012320032121 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3102232112321332-3012130323200131-3020222112002132-1021022231030331-1303113000300013-2022200133032132-1020100201302110-2303301320100131"></a>

## timeout_seconds property — batch / 012320032121 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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

<a id="canonical-3123322310312320-0122213321312131-2020112203303222-1102122231113213-3303310331223033-2333020303033120-2030120133222100-3320321232323313"></a>

## Next pages — batch / 012320032121 / 7

- [http_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2322030211000003-1133331020121313-1121222323220301-3130300131123301-3032122333332210-3112230013110020-2112123233131011-3130320110302101)
- [http_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-3220021331302101-0030133213132003-1021212332230013-3031333022301012-0101300112312312-0010133220330101-2030232333313023-0131102031131132)
- [http_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-3300203101200123-1312323331101210-1323030333211102-1201111030322333-1130001230332211-0010232203023321-3300120322112101-3001022231301002)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2322030211000003-1133331020121313-1121222323220301-3130300131123301-3032122333332210-3112230013110020-2112123233131011-3130320110302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010220011320300-3102332120013233-2101102123332212-2020121013010003-1311323003232033-2311312120113231-0232122333010233-1003032032022101"></a>

## http_receiver.batch.max_bytes_disabled — max_bytes_disabled / 103120203333 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.max_bytes_disabled

<a id="canonical-0220301012131310-2301132300103330-2312200312101001-0122123022220320-1322302021332103-0231112011112032-3020211031113210-1030010001310003"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-1211312333211012-3113001231011213-2031000323202122-0033112312312201-1110330033310103-3320111303313101-3003222233302032-3333301121311220"></a>

## Direct properties — max_bytes_disabled / 103120203333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010000001311120-0120010001103212-2233003101332231-1201231021200000-2101323012013122-3123331213232303-1213202110303021-1331202232133032"></a>

## Next pages — max_bytes_disabled / 103120203333 / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3220021331302101-0030133213132003-1021212332230013-3031333022301012-0101300112312312-0010133220330101-2030232333313023-0131102031131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300232323222300-0222220201113230-3130202210312333-1022302112010111-2300213120013123-3013121113001003-2312313101030323-3233112212132032"></a>

## http_receiver.batch.max_events_disabled — max_events_disabled / 200232032003 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.max_events_disabled

<a id="canonical-0310223123233113-1323110321230333-3030023020121023-1121322110323203-2222130230222300-0211131330110102-0231311120221212-1123102303202300"></a>

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
max_events_disabled = {}
```

<a id="canonical-3003013121103113-3100321323202230-0100202202021002-2231101011231110-1022032111322330-1023332330313213-2301030102003111-1032312011202120"></a>

## Direct properties — max_events_disabled / 200232032003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023220001301032-0330022133112333-3102111312333220-1103111330112010-1222003320331303-2310203230122312-0021221323002202-2231131103030133"></a>

## Next pages — max_events_disabled / 200232032003 / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3300203101200123-1312323331101210-1323030333211102-1201111030322333-1130001230332211-0010232203023321-3300120322112101-3001022231301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331121001122333-0330030302223201-3203012303332113-3030112001210011-2203121121010230-2111000002032310-0300112012310001-1300203333222330"></a>

## http_receiver.batch.timeout_seconds_default — timeout_seconds_default / 300210233301 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- http_receiver.batch.timeout_seconds_default

<a id="canonical-0003000020012230-3121131132031020-0330320302100101-2331132010303000-1013332000120020-1200233231002330-0013311030322021-1210220012111202"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-2323230212220012-1231003331221211-3333223100311201-3300331303000311-3310311123112111-2202100330301031-3210321132213121-0333123220231022"></a>

## Direct properties — timeout_seconds_default / 300210233301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232102003301111-0210311133303011-2112031101110200-1000131113301032-2011200132002021-3313220332022130-1130130310023113-1130000012113201"></a>

## Next pages — timeout_seconds_default / 300210233301 / 4

- [http_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-0102310011303233-1212220212302323-3313000020023020-1100320133210211-2333121320223233-2313312313131202-1200320323030001-1320032012112211)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113032231001320-1210000123302021-1011021020330323-1333121122013220-0201323232123301-1223221031323212-2011223001101302-3033312030001023"></a>

## http_receiver.compression — compression / 223323322310 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.compression

<a id="canonical-2320032320122303-0303033220030200-0133322101222122-2332112103133323-1031320012123010-3021022303023200-3033100330132223-1013313201331232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0311202103303221-0321302222033030-0112103301303002-0131310213023110-3132302032111133-0112003302323222-2103231310113223-2103333323301312"></a>

## Direct properties — compression / 223323322310 / 3

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-3121131032201023-3120023111002323-2021100132311133-2300031310212020-3322101102033132-1021330020321021-0003130332231200-2220021100001321): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3111112222223130-0013112302123330-0330123230332132-3222102103213230-2200232102200200-2313022032003010-3030133212112021-0331310013122101): complete subsection reference.

<a id="canonical-2101223233213030-3123210210321110-0000233222202322-1033101102032320-1131323113031003-3333130132303301-1020030203133001-3121200202300331"></a>

## Next pages — compression / 223323322310 / 4

- [http_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-3121131032201023-3120023111002323-2021100132311133-2300031310212020-3322101102033132-1021330020321021-0003130332231200-2220021100001321)
- [http_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111)
- [http_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3111112222223130-0013112302123330-0330123230332132-3222102103213230-2200232102200200-2313022032003010-3030133212112021-0331310013122101)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3121131032201023-3120023111002323-2021100132311133-2300031310212020-3322101102033132-1021330020321021-0003130332231200-2220021100001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211000331010110-3233230302023013-1020030220311013-3123221113232121-1001300233321312-3220231303111023-2013033130011211-0012122110320332"></a>

## http_receiver.compression.compression_default — compression_default / 210011330333 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_default

<a id="canonical-0133101033000002-0330202101311321-3213323312101301-1222030100133200-0210101213000132-0213222333131120-2210132333020301-3033320000133230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-1231232230111112-3023023223030333-1232131231212322-0333011013122202-1231320121221013-3222321112003200-2202101222222202-1033100301312322"></a>

## Direct properties — compression_default / 210011330333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002000221300111-3003130322322112-3120322321001320-1222100212230201-0212303231000210-1300301323121331-0111033000313110-0100322321301110"></a>

## Next pages — compression_default / 210011330333 / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0213103022202120-0223301323013023-2102120011131232-1000330133303213-1123032122212312-1222122300200201-0013211023200232-3030311212132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113313012231101-2303122110202303-0001010111310323-3222101122232013-2010300333323000-1000132323022303-3020203301303233-2211212010123230"></a>

## http_receiver.compression.compression_gzip — compression_gzip / 231002222121 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_gzip

<a id="canonical-3331323200213231-0120103201102010-3030010331020322-3301000120331132-1333300211330202-1000021222121300-0123220232212310-3321102011200311"></a>

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
compression_gzip = {}
```

<a id="canonical-1022030233313103-1020221302000312-3012110231101001-3003111223032003-0020301202112323-1032132101322221-3010312123222322-1021220101133000"></a>

## Direct properties — compression_gzip / 231002222121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230000120010032-1032100002213113-0120303001312210-0310211203330030-3213323231210131-3030101310003123-1323021031011223-0331201030020002"></a>

## Next pages — compression_gzip / 231002222121 / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3111112222223130-0013112302123330-0330123230332132-3222102103213230-2200232102200200-2313022032003010-3030133212112021-0331310013122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202111112200121-1201000000032121-1210000203323300-2230212013100320-0301103120131133-2331310310331102-0010133221023313-1113323203111202"></a>

## http_receiver.compression.compression_none — compression_none / 222113010203 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- http_receiver.compression.compression_none

<a id="canonical-0330111322330010-3200220201223002-0220333222222333-0033310230333203-1300333130231331-2320021113221330-3210100322310122-0231133303220313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-3201133001021222-0210033132010300-3010321100130233-3001232010301031-3111333123021113-2331203213031312-1022031300002231-0121323103121301"></a>

## Direct properties — compression_none / 222113010203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332032011302113-2210321021220211-0203230003110210-0231231222131032-3301010103323100-1231202120301033-3100300330122101-0010111301233211"></a>

## Next pages — compression_none / 222113010203 / 4

- [http_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3133203112210102-2303223113321220-1030012202132221-1231310113101233-1233330221020123-1232300100103012-1321231320233133-2103012323303331)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3233100233103022-2230123213321331-0131001032100301-2102322211220020-3112230003112133-0301020101101012-1200122011110232-2212133013103130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121120232022202-2012010012313221-1322223033201122-1110302113311220-3022223103213220-0201331011221110-1103302001200132-2220022313221311"></a>

## http_receiver.no_tls — no_tls / 231201123311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- http_receiver.no_tls

<a id="canonical-0021010300200121-0020302101130303-0001312133311002-3122123120132133-2221222132121130-0320300113233120-0302113221030121-1311233111321202"></a>

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
no_tls = {}
```

<a id="canonical-1020033102221133-3003031113030202-3300211112020221-3020310102110213-3002212111210023-1121310033021211-2320100002301331-0221031321013222"></a>

## Direct properties — no_tls / 231201123311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332031200311322-3233223212213301-0102222203200230-2220013130020333-0230203210321323-3003210220230213-1231033012131212-2000103030200331"></a>

## Next pages — no_tls / 231201123311 / 4

- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023003020333200-1230121113120032-2300033000331232-2000001023102222-2203220132231100-1202203323010110-1301200011230200-1311100223200133"></a>

## http_receiver.use_tls — use_tls / 133103100022 / 2

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

<a id="canonical-3313011121202232-2303110200201132-0011011211033331-0212322133320030-2000210131132010-0211211120201030-1122233320111102-2212310230022011"></a>

## Direct properties — use_tls / 133103100022 / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-2110120122310202-1332022131132123-1131211233223300-1222312231230211-2111323011102032-3133021030310103-2301122300231111-2203131131213122): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-1231001233202010-1210001002202323-3033112010133000-0121002321001212-1120203110200212-0320020031131021-3311023231013002-0232221133323013): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1133202021221033-0230201020322130-3021222322011200-2002111022000311-2323031330013132-1333000323213301-3112033113010021-2301213003223133): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0000331020321020-0222333100313211-0202111130321023-3002212033113023-0111232022012231-0331202121233212-3320110001122311-2120300002033210): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2113133323211020-0012300213221300-3100102220320103-2313201312310031-3123311022110100-2223121313133033-0321211200233002-0133223110003223): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-003.md#canonical-2120210232111220-1103201201203332-3232310103113313-1130232323321000-3013313330012101-1033233021113012-2131223202233103-2303310321302000): complete subsection reference.

<a id="canonical-0322301231001023-0020030312030110-3333311010310111-2011233120031230-0312112212300303-1212120203222010-0210231301320032-1311300003232233"></a>

<a id="canonical-0123032120022001-0000303222331002-3222301202131322-3322320333223123-0330213232012303-3111103133011301-3020301011101232-3001020333011110"></a>

## trusted_ca_url property — use_tls / 133103100022 / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3021211311022123-1300320100300123-0010233212332021-0123030001003002-1131133033212233-0222022120323213-2012101000031020-1102232310112300"></a>

## Next pages — use_tls / 133103100022 / 5

- [http_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-2110120122310202-1332022131132123-1131211233223300-1222312231230211-2111323011102032-3133021030310103-2301122300231111-2203131131213122)
- [http_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-1231001233202010-1210001002202323-3033112010133000-0121002321001212-1120203110200212-0320020031131021-3311023231013002-0232221133323013)
- [http_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1133202021221033-0230201020322130-3021222322011200-2002111022000311-2323031330013132-1333000323213301-3112033113010021-2301213003223133)
- [http_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0000331020321020-0222333100313211-0202111130321023-3002212033113023-0111232022012231-0331202121233212-3320110001122311-2120300002033210)
- [http_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2113133323211020-0012300213221300-3100102220320103-2313201312310031-3123311022110100-2223121313133033-0321211200233002-0133223110003223)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210)
- [http_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-2120210232111220-1103201201203332-3232310103113313-1130232323321000-3013313330012101-1033233021113012-2131223202233103-2303310321302000)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2110120122310202-1332022131132123-1131211233223300-1222312231230211-2111323011102032-3133021030310103-2301122300231111-2203131131213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302010232021220-0123032230202321-3120130133303033-0030303102331200-1130103321231110-3220221122121313-2131123110133223-1211322033323303"></a>

## http_receiver.use_tls.disable_verify_certificate — disable_verify_certificate / 313212012122 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.disable_verify_certificate

<a id="canonical-1230011120133333-2210331123002101-0120223303323213-3030111023010000-2302201202310220-0020102232202211-0302212032330112-3132101131322313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-0233002322220221-1010103201322223-1302232100031122-2201101221213002-0323332130000013-2330210110222230-2321212200212232-1231212031213123"></a>

## Direct properties — disable_verify_certificate / 313212012122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301203311201032-2113220230213013-0323211233000130-0213300132113022-1013223112033220-2020312002333203-0302011333301130-0010223223311003"></a>

## Next pages — disable_verify_certificate / 313212012122 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1231001233202010-1210001002202323-3033112010133000-0121002321001212-1120203110200212-0320020031131021-3311023231013002-0232221133323013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033021112020002-2301120311221112-1312202030131113-2311002331213311-2103302120012323-3113201033333120-0320322332323023-2013323230331011"></a>

## http_receiver.use_tls.disable_verify_hostname — disable_verify_hostname / 022311121322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.disable_verify_hostname

<a id="canonical-0022121122120312-0122300132022022-2330301012310123-1221232131121101-1111300323132131-0023130331302132-2210012200321021-1021301323222102"></a>

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
disable_verify_hostname = {}
```

<a id="canonical-1012310021032030-3131331220221111-1013021302130330-1031303030312231-3003020131031020-0113211221001322-1310132332102200-3022221013311002"></a>

## Direct properties — disable_verify_hostname / 022311121322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330212123212221-2131121322222222-2321320302102232-1301000001312203-3302202123021223-3102232320021022-3330002123132011-1120311131300332"></a>

## Next pages — disable_verify_hostname / 022311121322 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1133202021221033-0230201020322130-3021222322011200-2002111022000311-2323031330013132-1333000323213301-3112033113010021-2301213003223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233011212120231-3313233313333031-0311212011010301-0020022230311001-2022302330002210-1010313201022332-2123030123010121-2233311100113230"></a>

## http_receiver.use_tls.enable_verify_certificate — enable_verify_certificate / 233110200010 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.enable_verify_certificate

<a id="canonical-3213311211122021-1320312311322031-2301211010222032-0102000201011000-3203133233332110-3302200313110033-0330323133000012-3303030102000300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

<a id="canonical-0321132201313000-0130113230303031-2201112231131012-1132123312030321-0321030202120330-2010223102133322-1103331213110202-1312021300311213"></a>

## Direct properties — enable_verify_certificate / 233110200010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102131212113320-0201300203032301-1310332011033221-3003211110112031-2220303010220030-1311121303020133-1332031021112001-2131122210301322"></a>

## Next pages — enable_verify_certificate / 233110200010 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0000331020321020-0222333100313211-0202111130321023-3002212033113023-0111232022012231-0331202121233212-3320110001122311-2120300002033210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211023123321300-2223310331302000-2103111221302231-1321212032230020-0201001233133321-0302331231313302-0100023230213312-1323032132331312"></a>

## http_receiver.use_tls.enable_verify_hostname — enable_verify_hostname / 233012123011 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.enable_verify_hostname

<a id="canonical-2120111331123120-1002033032002023-0131233322331330-0313020103112031-2002032103321023-2001311012321123-3110020202323100-3331203102300232"></a>

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
enable_verify_hostname = {}
```

<a id="canonical-3300232033130210-3100311102121023-3031022031330003-3101202132203012-3113103212103030-3103323101220312-2132233311101203-2033121321211231"></a>

## Direct properties — enable_verify_hostname / 233012123011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120102313001122-3021111131010002-1230232032213121-0022102212321012-1210010003113323-1122212202000000-2003012300322330-3212200131210133"></a>

## Next pages — enable_verify_hostname / 233012123011 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2113133323211020-0012300213221300-3100102220320103-2313201312310031-3123311022110100-2223121313133033-0321211200233002-0133223110003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300320030111302-0002322102212030-1120301221322232-3332330131103222-3101213332332222-3031321210232030-0111031001031132-2123132003332321"></a>

## http_receiver.use_tls.mtls_disabled — mtls_disabled / 203012032233 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.mtls_disabled

<a id="canonical-3331313133020210-2300121120320213-2012301031302101-1322313203221300-3011302131322310-3212233300222201-3210302323022321-0031130012301300"></a>

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
mtls_disabled = {}
```

<a id="canonical-2112202203230100-1012331101313113-1003333132333112-3003023010202312-0323200020021300-3003232010210222-1300330322200211-0211223113111320"></a>

## Direct properties — mtls_disabled / 203012032233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013001003033112-1233220001022113-3132102201333130-0333002200023011-1132110212213030-3321123310223120-3323101321303013-0233323323122101"></a>

## Next pages — mtls_disabled / 203012032233 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323323020003211-2010230102200002-1022233000103211-3020312311233021-0202122230103032-1021213110110201-3201100103100301-2333130010113200"></a>

## http_receiver.use_tls.mtls_enable — mtls_enable / 300301111323 / 2

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

<a id="canonical-3322120303331332-0033121123122232-0023313310210021-3003310012213222-1311111300021322-3333223302223323-3233213221020103-3021210222120201"></a>

## Direct properties — mtls_enable / 300301111323 / 3

<a id="canonical-0130333200031221-0322220233213003-1133311112130313-0133212012212003-0320123202121101-3202300113203131-2310203202012121-0121112210310123"></a>

<a id="canonical-0333020001103213-3010112031210230-0031033011213100-1021120220220000-1022122322333022-0011022030303032-2012310320102131-0333223031011100"></a>

## certificate property — mtls_enable / 300301111323 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1233031030103131-2331313030121231-1111301211320210-3102133122003220-1123331110200132-0122332121333111-0003331131200022-0203032011132221"></a>

## Next pages — mtls_enable / 300301111323 / 5

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100101311130323-1022032131000111-3032112030020121-3000123310313130-3203012222203112-3201302211223123-0320131133010313-1103323102012310"></a>

## http_receiver.use_tls.mtls_enable.key_url — key_url / 013313312201 / 2

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

<a id="canonical-1032222122300300-3333010101203003-0101121132000133-0103311220203010-2201313323030201-0311200320200202-2332301030300320-2333020210313113"></a>

## Direct properties — key_url / 013313312201 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2030301030133032-3212313313313131-0102303132011211-3211031113330322-0002232331331111-2200323102110311-0010223001300213-1023002100111301): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-3203132230022020-1103113201022030-0103231013222221-3023113032323020-1223023231303112-0320133132211312-3233300033200132-0313130222311332): complete subsection reference.

<a id="canonical-2122012102133033-1110233201200113-1303013031120103-3200110030311320-1123303230301120-2110012111130000-3201032211032303-0010122030111111"></a>

## Next pages — key_url / 013313312201 / 4

- [http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2030301030133032-3212313313313131-0102303132011211-3211031113330322-0002232331331111-2200323102110311-0010223001300213-1023002100111301)
- [http_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-3203132230022020-1103113201022030-0103231013222221-3023113032323020-1223023231303112-0320133132211312-3233300033200132-0313130222311332)
- [http_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-0312002222213210-3223033110133313-2030021303102230-2013000101131032-2311220002203102-3003101311000303-3003132331210031-3113310100130210)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2030301030133032-3212313313313131-0102303132011211-3211031113330322-0002232331331111-2200323102110311-0010223001300213-1023002100111301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223223231020113-2112113220111221-3323220112333032-3113122123202231-1130121201332012-3033303231031011-1132312312322022-0012311122133012"></a>

## http_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 131301230330 / 2

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

<a id="canonical-1232213101303002-0202301111231213-0223133102320130-3010202310113200-1302101202000103-3213110010323002-2103032132112103-0103012013010032"></a>

## Direct properties — blindfold_secret_info / 131301230330 / 3

<a id="canonical-0112232210120133-3213301232011013-3231220220331231-3030331132002303-0102012303313130-3131221023011121-3313203031023111-2300130230310220"></a>

<a id="canonical-2302320101210100-2123120130012131-3032222121220121-3101311213120003-1131301312100331-1220023030033013-1310300013323131-0011221201120013"></a>

## decryption_provider property — blindfold_secret_info / 131301230330 / 4

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

<a id="canonical-3120011313320303-3003212131102012-1013001303032302-3122102110110310-0333022132313302-2000031312001100-0103303130323313-0021333300032101"></a>

<a id="canonical-3211203132202230-3103020112200331-3000013223222000-3010220210222202-2330200310002122-0030101022321112-0031132023123110-1122130302230113"></a>

## location property — blindfold_secret_info / 131301230330 / 5

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

<a id="canonical-1111000231221000-1311332101130212-1302033300010021-2001110120022230-3220032310323123-1020222213200311-3333103301022103-0020231211201231"></a>

<a id="canonical-2330002322222233-0012303310111312-0312010021302130-0012312111013111-1303021302311233-1101003321302103-3200303330210220-2220203312320323"></a>

## store_provider property — blindfold_secret_info / 131301230330 / 6

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

<a id="canonical-0212313213301232-2033213133013211-1311101010013312-0232230320102131-1133231312023312-2033311101232103-0223122202230012-2111000030132212"></a>

## Next pages — blindfold_secret_info / 131301230330 / 7

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3203132230022020-1103113201022030-0103231013222221-3023113032323020-1223023231303112-0320133132211312-3233300033200132-0313130222311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300122000332230-2112102112231302-3303002100133123-2301233031333130-2023111132222303-2023130101331121-2202321123132011-2033312132122231"></a>

## http_receiver.use_tls.mtls_enable.key_url.clear_secret_info — clear_secret_info / 210203033200 / 2

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

<a id="canonical-3121302131321221-0030100311110213-2301233020030032-0132231220211312-0131032213332302-2132221002010102-3211002133001023-0203003121311233"></a>

## Direct properties — clear_secret_info / 210203033200 / 3

<a id="canonical-0031130110321102-2022101230002202-2233030121102322-3332312301231202-2221231021010223-1330311113100312-2113233123330231-1330220202100231"></a>

<a id="canonical-2031112032203203-2132020013112003-3111233311203200-2000101013300103-2023320003003313-3331200313232033-2022223112133130-3321110300200102"></a>

## provider_ref property — clear_secret_info / 210203033200 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0212300023232033-2132302113320130-2303011312333112-2101330321230200-1031122030131022-1110202122133312-0333211003203013-1110123330022321"></a>

<a id="canonical-3311221230202011-0333110020113322-1303211331110122-3332100322310330-0220310200102322-1312210233101132-1001011111210001-2332103123302322"></a>

## URL property — clear_secret_info / 210203033200 / 5

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

<a id="canonical-0201101331000010-3223221030313220-1102303112303303-2312012212302030-1320321321112030-2313321310310333-2033333302003232-2211101110023323"></a>

## Next pages — clear_secret_info / 210203033200 / 6

- [http_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-0322200101111022-3312322200113313-0003312032113332-3310112300030021-2322101032013231-1000311330012110-0100023213131132-1023201201202002)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2120210232111220-1103201201203332-3232310103113313-1130232323321000-3013313330012101-1033233021113012-2131223202233103-2303310321302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103210202112202-1220131113302113-1233323121311201-2130233123023323-3230230023211230-0103323301023213-2002123123003120-1102130101203113"></a>

## http_receiver.use_tls.no_ca — no_ca / 121211112100 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [http_receiver](resources--global_log_receiver--reference--group-002.md#canonical-2333020321103033-1131201001210213-1030333120020331-3102120203022030-2020212132223122-3212110301033113-2301112212113303-0200101300212333)
- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- http_receiver.use_tls.no_ca

<a id="canonical-0320303010223313-3103100300103013-0020011220310123-0110033013323222-3233313331203132-2013021030122303-1332023323110030-1211131230331202"></a>

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
no_ca = {}
```

<a id="canonical-3201113310323003-2020003230202031-3213313131100123-2230310122232223-0110201031112121-2222321320310320-3331231212333112-3022332103222000"></a>

## Direct properties — no_ca / 121211112100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203103210012003-0230121003023232-1311321033032212-0121323103212232-3210113103301233-2213213200031322-2110310001031231-3222301122300223"></a>

## Next pages — no_ca / 121211112100 / 4

- [http_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2120003110030212-2012220302210213-2011103122223333-3310303122000230-2122231230010301-3002011330030112-0222012333322310-0013332200211123)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300311301010212-2202213301230001-3121310223022110-3030013310003320-0232122230002003-0031222310100321-3333120300310222-2230120312000202"></a>

## kafka_receiver — kafka_receiver / 000132213003 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- kafka_receiver

<a id="canonical-0300211310300033-3303232313232332-0103302032102312-3331310202130303-2112110331221010-0120101202030011-0132333301330000-1003221223001012"></a>

Type: `"object"`. single nested block, Optional.

Kafka Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1121333112311200-1100032003223212-2230030323032133-3032102212022220-3331011110212000-0211013213231000-1320112000033230-2211033000331220"></a>

## Direct properties — kafka_receiver / 000132213003 / 3

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311): complete subsection reference.

<a id="canonical-1221210133303301-3131320221033213-0113123330213102-3023200111131233-3121001032200013-3332111022231122-1212311222202203-3203321121301323"></a>

<a id="canonical-0203223120213122-0311112101322220-2201210223030023-1020113201233331-0232021331312021-2313221123000222-0323210110202030-0222220101002112"></a>

## bootstrap_servers property — kafka_receiver / 000132213003 / 4

Type: `["list", "string"]`. Optional.

List of host:port pairs of the Kafka brokers.

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

<a id="canonical-0120120020223303-2100312332020223-1231322023313303-2201221202233020-3233001232302333-1201102102013023-2013330120303012-3031320022030220"></a>

## kafka_topic property — kafka_receiver / 000132213003 / 5

Type: `"string"`. Optional.

The Kafka topic name to write events to.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1311331201213230-3001310323000223-1231233122300031-0332113330120012-2311120111133211-3331200000012311-2201101001113000-2221223322202113"></a>

## Next pages — kafka_receiver / 000132213003 / 6

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- [kafka_receiver.no_tls](resources--global_log_receiver--reference--group-003.md#canonical-1330000103221001-0032201311231232-0021202030102312-1103323120300010-3132201012031130-2133020310001133-2300230132200301-1121212323312103)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133001333222301-1020302302212033-3331203310230333-1012220323310113-1311322310300131-3213011113011232-2330303332203100-3213132310112223"></a>

## kafka_receiver.batch — batch / 210033110233 / 2

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

<a id="canonical-2111300031103311-2330313311013313-2031332200113231-2310102030220013-2103132101010311-1000002210012213-1130212103030012-3112231120002312"></a>

## Direct properties — batch / 210033110233 / 3

<a id="canonical-0101133130201200-1121011031221020-0200210112230102-3110312333233130-2200301302311333-2003223010022113-0113132102123121-2021313123302121"></a>

<a id="canonical-0023210122023220-3301220010223121-3112001032010120-0201211132332020-1233010232023010-2020302033312202-0232113313022311-3000020030313301"></a>

## max_bytes property — batch / 210033110233 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0103130121223310-2220030213033202-2223030230222020-2123322011312012-3110130112111222-2030221101030121-1320113302010112-3130331100120020"></a>

## max_events property — batch / 210033110233 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1133123200102021-1310000110330020-0121122202332032-2003011320102223-0121000011230233-0221332030031133-0311301130110323-1201123323333101"></a>

## timeout_seconds property — batch / 210033110233 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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

<a id="canonical-0303120032131133-0231130113122011-1132103322100310-3033110313103021-1110031202122302-1321233031332223-1032131310321121-1011212222233110"></a>

## Next pages — batch / 210033110233 / 7

- [kafka_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0003310101230200-2130120313100023-2211313123000121-2330130111033320-3322020122222130-1000110000230021-2101203033300220-2212100131003233)
- [kafka_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-1230133223101300-2113013232301220-2233322030222130-1023301322013120-3331302032323220-0021120203310212-1302001220033202-3221213232022100)
- [kafka_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-0230322320212102-1231230032123311-2012002232222313-3231003111003330-1012202300222321-2210102220210002-3023113333111130-0000302233010212)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0003310101230200-2130120313100023-2211313123000121-2330130111033320-3322020122222130-1000110000230021-2101203033300220-2212100131003233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212130223232302-1323220211221331-2000021132130330-1312010330110323-3213012021300113-1133213310021011-3333021001211223-3102110121003210"></a>

## kafka_receiver.batch.max_bytes_disabled — max_bytes_disabled / 131120031221 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.max_bytes_disabled

<a id="canonical-1112221213102031-1313102331032021-1233312122213201-3130030222333303-2231321313213010-3311023220100012-2333313321220221-1132100101030221"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-1331022020110212-1030231312212321-0323020102233333-2130120223313130-3101112121022323-2020021112332203-0312103313021321-1011113221222330"></a>

## Direct properties — max_bytes_disabled / 131120031221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333302212002222-1112031200231033-3131230133220230-1333332202211223-3032122232033122-3021101301121332-2313212120321031-0302323222332021"></a>

## Next pages — max_bytes_disabled / 131120031221 / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1230133223101300-2113013232301220-2233322030222130-1023301322013120-3331302032323220-0021120203310212-1302001220033202-3221213232022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203003111200201-0332033303032032-2002110220020000-0320023323212321-0301201200010230-3313131321333302-1031311131022312-0202101223113203"></a>

## kafka_receiver.batch.max_events_disabled — max_events_disabled / 212133213322 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.max_events_disabled

<a id="canonical-0222332301112112-1323203030202220-3103013022033203-2203320200333023-2002213231012111-0103303121230030-2030021311013231-1120011333102213"></a>

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
max_events_disabled = {}
```

<a id="canonical-1201203102313132-1130030232313013-0202021333201232-2211021321120011-3200011021330111-1001311112103331-3010112002202302-0112113011331110"></a>

## Direct properties — max_events_disabled / 212133213322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203331221220121-3123310031103123-3000211013121300-0300002001023211-1122310012112300-3301313301002120-3000331202131230-3002220123130211"></a>

## Next pages — max_events_disabled / 212133213322 / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0230322320212102-1231230032123311-2012002232222313-3231003111003330-1012202300222321-2210102220210002-3023113333111130-0000302233010212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031313010111210-3300000010113122-1312001323323012-3013311132333201-3133313210132012-3202102312231320-1001110011013220-1013212103300113"></a>

## kafka_receiver.batch.timeout_seconds_default — timeout_seconds_default / 023032013331 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- kafka_receiver.batch.timeout_seconds_default

<a id="canonical-2133331331213111-0313110103103201-3312202113031130-3010311000320231-3303103330103232-0330022211321202-1033202123303312-3202000032132202"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-0033002331030021-3103132203013132-0023132322233210-2002112231132322-2110010022302111-0110210111103222-1330311220012011-0010313313112211"></a>

## Direct properties — timeout_seconds_default / 023032013331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200203123202231-2233022202212300-3222211333310102-1200313130310231-3103023113112121-3012301203303220-1321112310211311-2021023322303320"></a>

## Next pages — timeout_seconds_default / 023032013331 / 4

- [kafka_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2011131303210003-0323131003232022-3100313331203311-0101212031323012-1122330320311300-1311032121013131-1133210003213201-0022010120212311)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033023233200111-1201120310113013-1233133010123102-3231011031223222-2013210230230101-2103222212223310-1033120132200101-0220020200202012"></a>

## kafka_receiver.compression — compression / 323012130312 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.compression

<a id="canonical-2131322222331221-2110213002112031-2300320001322210-2100201113112232-3020311002230033-0331111311303331-2102230230213332-3232323222033100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1200021323323310-0122321010132231-2032112101222220-0312311313331123-1020323223101302-3323210221030112-1132103332000230-2320023110231022"></a>

## Direct properties — compression / 323012130312 / 3

- [compression_default](resources--global_log_receiver--reference--group-003.md#canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-3031323100032313-2023232013230033-2100002230032211-2020003211230133-2321121233030020-0220203330013322-1011332031230001-3001333311223033): complete subsection reference.

- [compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3110103223213301-2003030032302313-1022210232002202-2133102320302211-0201003122230031-1001320300222100-3121022001220103-2021113120311332): complete subsection reference.

<a id="canonical-1330200221033333-2300131330331320-1310223113000233-1020223031330321-3133230211013001-3333110223312332-3013203111022300-2000313102002001"></a>

## Next pages — compression / 323012130312 / 4

- [kafka_receiver.compression.compression_default](resources--global_log_receiver--reference--group-003.md#canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330)
- [kafka_receiver.compression.compression_gzip](resources--global_log_receiver--reference--group-003.md#canonical-3031323100032313-2023232013230033-2100002230032211-2020003211230133-2321121233030020-0220203330013322-1011332031230001-3001333311223033)
- [kafka_receiver.compression.compression_none](resources--global_log_receiver--reference--group-003.md#canonical-3110103223213301-2003030032302313-1022210232002202-2133102320302211-0201003122230031-1001320300222100-3121022001220103-2021113120311332)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331222311203212-0233130103132103-0200220021003222-1002002011123332-3001122221322331-3113331233230301-3021122303301232-1121213000220022"></a>

## kafka_receiver.compression.compression_default — compression_default / 032113120123 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_default

<a id="canonical-1013110231132100-3313101233013302-1003030230003122-2320120111230223-3011321321222110-0231312100130222-1302320101103312-2200021321003113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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
compression_default = {}
```

<a id="canonical-3013312232022311-2023320001230231-2331102013200020-1013031123033030-1223303012220200-0310200031113001-0231122223001003-2301013013231131"></a>

## Direct properties — compression_default / 032113120123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300222301021122-1222121231310213-3321230231131033-1131201203002122-1203110322201312-3033211211313300-1122213021231012-1110032313221211"></a>

## Next pages — compression_default / 032113120123 / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3031323100032313-2023232013230033-2100002230032211-2020003211230133-2321121233030020-0220203330013322-1011332031230001-3001333311223033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203110210111101-1023313011200011-1313333311122221-0202233120212120-3332212201210223-3030230011320211-2330132212033232-0120230303032020"></a>

## kafka_receiver.compression.compression_gzip — compression_gzip / 011333211120 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_gzip

<a id="canonical-2210200010322033-3031130130333300-1313021023220310-3230222301121132-2002200331333212-1212210012013101-2332333102132003-3131100231103133"></a>

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
compression_gzip = {}
```

<a id="canonical-2233222131233123-1201030211002020-0113200013130200-2210120320201323-1001333023110000-0131110130330232-1233023103133012-1323003131203320"></a>

## Direct properties — compression_gzip / 011333211120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310300333023201-3200321310010200-0211113331230111-2102101231001221-3010021130231322-1310133330132222-0322121110331010-3022223220102001"></a>

## Next pages — compression_gzip / 011333211120 / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3110103223213301-2003030032302313-1022210232002202-2133102320302211-0201003122230031-1001320300222100-3121022001220103-2021113120311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230322123133002-1023203303101131-3121332133203202-1300133111220122-1222130121132232-2110013002323021-2202210123322203-2020132132311212"></a>

## kafka_receiver.compression.compression_none — compression_none / 322102003102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- kafka_receiver.compression.compression_none

<a id="canonical-3300301221313200-1310321202302222-3132011210312001-2113201212123201-0301003020122321-0323123132021133-2103330202211001-2233301023303031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression none.

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
compression_none = {}
```

<a id="canonical-0002031121301000-1232222233103220-3003301203213103-3010011122201031-1201332313211212-2201012101310313-0230311021020330-1131021332201030"></a>

## Direct properties — compression_none / 322102003102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012012011320232-3122312002310120-0133032302200123-3020133331021032-2211220333210113-3020120003122032-1313011100331102-0131022133322113"></a>

## Next pages — compression_none / 322102003102 / 4

- [kafka_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-0130222110211022-2132123223022011-2323232022112322-2310103231011332-0011031212202010-2221322211231101-3120232330230330-3232120302032332)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1330000103221001-0032201311231232-0021202030102312-1103323120300010-3132201012031130-2133020310001133-2300230132200301-1121212323312103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320300123303200-3210232323033030-2202031303003011-3333003111132230-0231231200332311-1222000030201102-1300233131321033-2132322303230300"></a>

## kafka_receiver.no_tls — no_tls / 212312100200 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- kafka_receiver.no_tls

<a id="canonical-0003333312013332-3312332113032023-2100333310320231-2212012030331222-0332123013120001-3011300232011122-3313133201230202-1010223222000111"></a>

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
no_tls = {}
```

<a id="canonical-2320013331232132-1110220022322100-2101103231011203-3331331000321120-1133113130110100-2020001330123032-3121121110313200-0100312222021121"></a>

## Direct properties — no_tls / 212312100200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031001022303301-2010301223033231-0113231032201211-2021013132123131-3303211331132101-2030012311312131-3231002232311331-1320232101300222"></a>

## Next pages — no_tls / 212312100200 / 4

- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011113323123111-3202022222012312-2322110032213010-0231122203102221-2213023131000322-2331220321132131-3101130212012233-0222300332330303"></a>

## kafka_receiver.use_tls — use_tls / 030121012321 / 2

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

<a id="canonical-1130230033013033-1311100210102223-0311121001231213-1021323333113122-0003110231032022-0323002110330302-0230001303033303-1133202312213221"></a>

## Direct properties — use_tls / 030121012321 / 3

- [disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1301020211212020-3131111102133200-3310333122011211-1033221303012201-3232233001223023-3233032333010022-1300012102303230-3300310032101103): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-3133203010121333-0000113013133311-2323220112113300-2211322131132233-1100300223320011-2123321300011212-0203202131313121-2321021132002121): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-3230223102001002-2231333030231132-3300011300021002-2310030320032020-2103331312002231-1002213323221322-1322331011101233-1203101002021100): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2232220303301202-0101201302302110-0110333020303032-1101131001003010-3330111311001030-3110312323332233-3110221302232133-2031012223010222): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003): complete subsection reference.

- [no_ca](resources--global_log_receiver--reference--group-003.md#canonical-1003222302113131-3020001323322130-2003312222122100-2203013200001021-3223010100323333-2303133002013033-1011312223111120-0302222010210022): complete subsection reference.

<a id="canonical-1112112120233110-3303111231021013-2200102232013332-1032310111313300-1032112202230210-2011203310021331-1010323311201201-0203220001310113"></a>

<a id="canonical-1313321010132202-1230012212200201-1200130310100222-2001011321103333-1010221302122132-1000102002013103-3111320321103113-1020322330222021"></a>

## trusted_ca_url property — use_tls / 030121012321 / 4

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3002030111333312-1011000112320121-1333022212233323-0200213103223301-0000030102132023-1130003023012311-3103121012222311-2100333022213221"></a>

## Next pages — use_tls / 030121012321 / 5

- [kafka_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-1301020211212020-3131111102133200-3310333122011211-1033221303012201-3232233001223023-3233032333010022-1300012102303230-3300310032101103)
- [kafka_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-3133203010121333-0000113013133311-2323220112113300-2211322131132233-1100300223320011-2123321300011212-0203202131313121-2321021132002121)
- [kafka_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--reference--group-003.md#canonical-3230223102001002-2231333030231132-3300011300021002-2310030320032020-2103331312002231-1002213323221322-1322331011101233-1203101002021100)
- [kafka_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--reference--group-003.md#canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310)
- [kafka_receiver.use_tls.mtls_disabled](resources--global_log_receiver--reference--group-003.md#canonical-2232220303301202-0101201302302110-0110333020303032-1101131001003010-3330111311001030-3110312323332233-3110221302232133-2031012223010222)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [kafka_receiver.use_tls.no_ca](resources--global_log_receiver--reference--group-003.md#canonical-1003222302113131-3020001323322130-2003312222122100-2203013200001021-3223010100323333-2303133002013033-1011312223111120-0302222010210022)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1301020211212020-3131111102133200-3310333122011211-1033221303012201-3232233001223023-3233032333010022-1300012102303230-3300310032101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002122301220002-1131203130323322-3033311122333130-0132303320132221-2333003131023021-0130210012100203-2220131003333312-1303021132022300"></a>

## kafka_receiver.use_tls.disable_verify_certificate — disable_verify_certificate / 213322203311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.disable_verify_certificate

<a id="canonical-1202023110110032-3221310313122100-0212010202123132-0122012331332122-0210030212120221-3320302320312132-1202313320013200-3223321213333002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable verify certificate.

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
disable_verify_certificate = {}
```

<a id="canonical-2031311200222010-2331301220202323-3010020210000030-1222012232310313-2223110312321311-3010321303033220-2201031211032001-0133023033230312"></a>

## Direct properties — disable_verify_certificate / 213322203311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220203231200013-0111000030030113-3330123231021233-0221102213002133-1232323001132000-3013310301120213-0121120133232301-0221303123233303"></a>

## Next pages — disable_verify_certificate / 213322203311 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3133203010121333-0000113013133311-2323220112113300-2211322131132233-1100300223320011-2123321300011212-0203202131313121-2321021132002121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030130310010233-0310102102120213-1202130121103020-0023002031311013-0023102113123130-3022030303321220-0233332030132233-2311210233123233"></a>

## kafka_receiver.use_tls.disable_verify_hostname — disable_verify_hostname / 033112322101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.disable_verify_hostname

<a id="canonical-2133213212331130-1121230221302202-2111100110221222-2200233210033212-1213321202112010-3132332201333301-2000101001220323-0023332022211023"></a>

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
disable_verify_hostname = {}
```

<a id="canonical-2221021301331211-2112301123011131-0302313013210102-1221122313223223-2101013303020113-0013021032001210-3303201131330010-1303120022303101"></a>

## Direct properties — disable_verify_hostname / 033112322101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031221200001031-0120033323201312-3003032013021110-0201121222023302-1223120002123313-0332120323102011-3223120030113221-1120012121003223"></a>

## Next pages — disable_verify_hostname / 033112322101 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3230223102001002-2231333030231132-3300011300021002-2310030320032020-2103331312002231-1002213323221322-1322331011101233-1203101002021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021212320233003-0031212113213013-3333211233133230-0210203102120012-3133333103333022-0200110012332121-3033011001130000-2312313132033313"></a>

## kafka_receiver.use_tls.enable_verify_certificate — enable_verify_certificate / 312133021300 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.enable_verify_certificate

<a id="canonical-3111112213202003-3033300013212303-1103310011310113-0122130302033030-0123111122013221-3332120120023322-1310202231133012-2121012330013210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable verify certificate.

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
enable_verify_certificate = {}
```

<a id="canonical-3020013200323032-0133322213202012-0131120121011023-0131211120033312-2000330300103110-3333211112220023-0131102332121310-1001212013123030"></a>

## Direct properties — enable_verify_certificate / 312133021300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231311323322203-0301213202330231-3220231223022310-0321323222310010-1013210221232313-2321001222133123-0022113023102011-0332231321203022"></a>

## Next pages — enable_verify_certificate / 312133021300 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0202222011220001-3010030131003222-2221003210311023-1330030111201211-1011121213132032-2202332321012002-3101201213231331-1211200123312310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011303103101220-1210231123223102-1210230103011231-2232222112022301-2112011233003130-2113333330111201-2033022113301333-3112113012031000"></a>

## kafka_receiver.use_tls.enable_verify_hostname — enable_verify_hostname / 033303112111 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.enable_verify_hostname

<a id="canonical-1101213130230231-0301030231020130-1330011000310103-2302233213111010-0033311330330220-0201111213202032-3102003032230123-3123121232303120"></a>

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
enable_verify_hostname = {}
```

<a id="canonical-1331322123232020-3221221132131220-0012022100223312-2220201301030213-3000323003212321-3110331110022012-3103011100001020-0203211330322031"></a>

## Direct properties — enable_verify_hostname / 033303112111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212033211231310-2033221331333302-3312213230302230-1120200211012001-0232031113220313-2003223331330233-3013020233310331-3111021333122322"></a>

## Next pages — enable_verify_hostname / 033303112111 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2232220303301202-0101201302302110-0110333020303032-1101131001003010-3330111311001030-3110312323332233-3110221302232133-2031012223010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032020333321301-2200231013132113-2322220021130002-0120101132102203-1022023101201033-1010101121100202-0031103201230112-2220203220132110"></a>

## kafka_receiver.use_tls.mtls_disabled — mtls_disabled / 200013201030 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.mtls_disabled

<a id="canonical-1130120021222200-3202130213203011-0112112223132001-3002320033321222-2032311310211330-1010231033001120-0021320233030133-3111230222120010"></a>

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
mtls_disabled = {}
```

<a id="canonical-2030003110030222-0122201331330123-1212000112222132-0012012330101130-2130321320212000-2321020320110321-1203303013101120-0322110032133013"></a>

## Direct properties — mtls_disabled / 200013201030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020010333333100-3013231033323131-0002322103021020-0222333310233203-3133103321231123-1222111211220202-0112203332132123-3300323322032232"></a>

## Next pages — mtls_disabled / 200013201030 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313211122003111-0133311312310203-1110230130303301-3320333120121122-1021300210311203-0201330003131332-3020021303223201-0232311033013211"></a>

## kafka_receiver.use_tls.mtls_enable — mtls_enable / 033332020102 / 2

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

<a id="canonical-0121113232023210-2311332203332000-0332003103122020-0201212230113010-2010300032032111-1202002232231000-2201131100011200-1012112013333310"></a>

## Direct properties — mtls_enable / 033332020102 / 3

<a id="canonical-0110231120301200-3031202311211211-3121103021011312-0233313210312232-0321211331201002-0332220300311323-0321010001203313-3311113003223112"></a>

<a id="canonical-2102112233220333-0030113212000200-0213311222120320-3123303012221111-0312111103021120-1301303202220000-1221312132202021-2233133001133100"></a>

## certificate property — mtls_enable / 033332020102 / 4

Type: `"string"`. Optional.

Client certificate is PEM-encoded certificate or certificate-chain.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100): complete subsection reference.

<a id="canonical-3130022110130131-3302121233113112-2013103133212120-0110202100012231-0221302032000210-1033132213321310-2101222122111300-2332320111002201"></a>

## Next pages — mtls_enable / 033332020102 / 5

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132031133132132-1131001220032233-1301232213120010-1301120111130222-1022103013033300-3222112130002223-3003133310333123-3221030010300021"></a>

## kafka_receiver.use_tls.mtls_enable.key_url — key_url / 022123120300 / 2

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

<a id="canonical-1030203030032002-2012300030311121-0232101012113013-3032310012230101-1211202310101020-2213303033313313-3122101300220320-1013200103120230"></a>

## Direct properties — key_url / 022123120300 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2323123022132202-0223213323323201-3313313233120121-1221021100320302-2211311020332132-1113210232233133-2010222012020023-0331201000333023): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-1212021111021103-0213311023223223-2121122231113102-2203321212030000-2022033313220132-3320211002201033-1302120231102331-2211223100002011): complete subsection reference.

<a id="canonical-2120012111122212-1223133121300333-1330111212120000-1300332031113013-2332330000333323-0210330123123100-2022300303310100-3321123220313121"></a>

## Next pages — key_url / 022123120300 / 4

- [kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2323123022132202-0223213323323201-3313313233120121-1221021100320302-2211311020332132-1113210232233133-2010222012020023-0331201000333023)
- [kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-1212021111021103-0213311023223223-2121122231113102-2203321212030000-2022033313220132-3320211002201033-1302120231102331-2211223100002011)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2323123022132202-0223213323323201-3313313233120121-1221021100320302-2211311020332132-1113210232233133-2010222012020023-0331201000333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112312313330223-0033300021103202-3310213213011321-1310310201032103-3012030011332030-1012200233012231-1131330300010302-1323100220230330"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 013110010321 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- kafka_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0330202331311013-2221000013002320-1023002333321333-2202331030103120-2001020120113123-3312311100010030-2331110012223112-3123311313311003"></a>

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

<a id="canonical-0002021020221230-2220222201302103-3022332010303303-2311213302231130-3213111100023101-2101131032030220-1203220001213333-3133013201000321"></a>

## Direct properties — blindfold_secret_info / 013110010321 / 3

<a id="canonical-0303101220303220-1002313302300221-0110000011310300-0220202210101020-1012300221223320-2003112303021120-1022222001122222-3121130201121212"></a>

<a id="canonical-0100330103203223-2103020322223302-2330031023121323-0103303313200312-0313223002333032-2131230002231123-3031232212023022-1031101130322301"></a>

## decryption_provider property — blindfold_secret_info / 013110010321 / 4

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

<a id="canonical-1333303221122020-0331310322311000-1110331223312311-3102003020312220-3302001201000022-3302230223123323-0032302320101210-3323200323130103"></a>

<a id="canonical-2322210313030330-0321020101313220-2133101323123202-1301000011202032-2230203131333222-1010120101110113-0101112001123230-2331311221112233"></a>

## location property — blindfold_secret_info / 013110010321 / 5

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

<a id="canonical-3111032233112222-2033302111212203-1313210103322003-3213313112231010-0221030013103322-0033110231011012-2010022022200131-2032322112010212"></a>

<a id="canonical-2132102000002222-3003313020202023-0333230033131021-0231021110332023-2322122023020130-0200213220132100-2111210212321112-2020021112301332"></a>

## store_provider property — blindfold_secret_info / 013110010321 / 6

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

<a id="canonical-0303333101111133-0200120212211203-3113130220321120-0311002313310002-0120331322003211-1121312122100010-3032333203012311-3300331311101000"></a>

## Next pages — blindfold_secret_info / 013110010321 / 7

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1212021111021103-0213311023223223-2121122231113102-2203321212030000-2022033313220132-3320211002201033-1302120231102331-2211223100002011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003320322231231-0313203233330333-1113010121010131-0000130013031110-2223011302301301-1032013223303321-0200112032331312-3010333221101122"></a>

## kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info — clear_secret_info / 002211312010 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [kafka_receiver.use_tls.mtls_enable](resources--global_log_receiver--reference--group-003.md#canonical-3132320030321203-1123223203110100-2110010332210110-2233312333320111-2322231202202331-0200002312211010-3011030230230303-1303103101133003)
- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- kafka_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2100030213311013-1230320030301120-2303323230120331-0012202212031132-1122110120111110-2210232122110320-1000112302220233-3213322001230320"></a>

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

<a id="canonical-0012030301110031-2003123033203103-3202221303231222-2220323333323312-2013012200113023-2330020231111131-3020022031021031-3311130213212120"></a>

## Direct properties — clear_secret_info / 002211312010 / 3

<a id="canonical-0030112321202221-3202023221313301-2221201132220210-1322231101113203-0132223230302123-1010312200113020-2232122321322033-3302322233330002"></a>

<a id="canonical-3311110101323333-2121100310233213-0100332223132131-2012233111311122-0112111111013213-3103021330002223-1201130222121112-3013331133000022"></a>

## provider_ref property — clear_secret_info / 002211312010 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0002000202030112-0110002223313130-1213033302323321-0033323211220122-3130212123322023-0011111130111330-1203231021000110-0033113220332122"></a>

<a id="canonical-1113313102232332-1003201031003202-1212002301332131-2211231233033032-1301223123113230-0320311100211302-1322023133331231-0322211031033300"></a>

## URL property — clear_secret_info / 002211312010 / 5

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

<a id="canonical-2031211103333202-3113212202022022-2110030102000123-0231201021232121-0313110220201321-1011013002332121-0301233032213033-1303133333103233"></a>

## Next pages — clear_secret_info / 002211312010 / 6

- [kafka_receiver.use_tls.mtls_enable.key_url](resources--global_log_receiver--reference--group-003.md#canonical-1203010003330300-2010101220321003-0021123021012301-2331213102112133-1130002332012220-3233203321210031-2012131213331222-1011100231010100)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1003222302113131-3020001323322130-2003312222122100-2203013200001021-3223010100323333-2303133002013033-1011312223111120-0302222010210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302023102200001-3003013003330123-3201300203121111-1232211310020233-3331101101321012-1330103221231301-3023022332100320-1302211302212332"></a>

## kafka_receiver.use_tls.no_ca — no_ca / 310123221313 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [kafka_receiver](resources--global_log_receiver--reference--group-003.md#canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112)
- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- kafka_receiver.use_tls.no_ca

<a id="canonical-0311121221322223-3110013021312110-2102323332312233-3012123011123310-1020331301231130-1200103120013310-0000233021213232-0122002222203331"></a>

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
no_ca = {}
```

<a id="canonical-0230122122111130-2200210021000000-3301020233322322-0020013020003110-2310331130322300-0311310202200110-1222320300312022-3222223320301120"></a>

## Direct properties — no_ca / 310123221313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021302312323303-2010133010021230-2301333101021301-3333323002102003-3311030233120001-1232131033331100-3000110200011330-3220210300032120"></a>

## Next pages — no_ca / 310123221313 / 4

- [kafka_receiver.use_tls](resources--global_log_receiver--reference--group-003.md#canonical-2220323000001131-2120013000010332-3033311211303213-2121200233222201-3000321311001101-3311030112130303-2131310002103131-2001311222022301)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200231031221120-3213122133312313-1112211212233102-3333212020221111-1203001230213201-1002222331013010-0031222232033313-3103120223132000"></a>

## new_relic_receiver — new_relic_receiver / 023032123102 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- new_relic_receiver

<a id="canonical-1230012130102110-2311120330023112-0022233130211001-0010032110111023-1303133132230232-3310210232122022-0323333321323230-1021002003003111"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for new relic receiver.

Upstream description:

Configuration for NewRelic endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0012121113133123-2222211220023302-0201211200331220-0003313231133230-0011200210021012-2323020003001001-0012220220103011-0323233333102313"></a>

## Direct properties — new_relic_receiver / 023032123102 / 3

- [api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300): complete subsection reference.

- [eu](resources--global_log_receiver--reference--group-003.md#canonical-0111332302103311-2231321210233231-3312031102230011-1222123203220300-3022101201022312-3033002110232322-1212003333010131-0331232213123100): complete subsection reference.

- [us](resources--global_log_receiver--reference--group-003.md#canonical-1012122103003011-0312131130200120-1312311010011133-1000220032230110-2330033223201123-0121312311122030-1103322213300021-3102320222222220): complete subsection reference.

<a id="canonical-0311033021222013-0232031212020102-0002220002302001-3323323200133301-1230020332213002-0111001202031032-1220321202231012-0120031112210230"></a>

## Next pages — new_relic_receiver / 023032123102 / 4

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- [new_relic_receiver.eu](resources--global_log_receiver--reference--group-003.md#canonical-0111332302103311-2231321210233231-3312031102230011-1222123203220300-3022101201022312-3033002110232322-1212003333010131-0331232213123100)
- [new_relic_receiver.us](resources--global_log_receiver--reference--group-003.md#canonical-1012122103003011-0312131130200120-1312311010011133-1000220032230110-2330033223201123-0121312311122030-1103322213300021-3102320222222220)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013023023220300-0101211210200110-2332303210122023-3021130332230022-2212332313232301-2311221011003130-2022333322231011-1100221101011001"></a>

## new_relic_receiver.api_key — api_key / 331012012311 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.api_key

<a id="canonical-0223021210302310-2012203211031313-1322122223013330-3333232121232301-0213001011330131-1222123201101313-1133230002211110-1110322300302202"></a>

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
api_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102200230103130-1022030310130002-2220022301221223-2130031312323333-2122203130310310-2303121230233223-0123012233301132-0322010330012123"></a>

## Direct properties — api_key / 331012012311 / 3

- [blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0312211120000312-0322322322111220-1103110331000011-0123030020200031-1030213100103200-1033200303231210-0021220300121233-3313021331211210): complete subsection reference.

- [clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2033113111202231-2211303110023112-0011123201132123-0222303230312333-2131100001321221-1301313200201112-0131031333212002-2210301321033121): complete subsection reference.

<a id="canonical-2030021123221320-1330213212330111-0131121220332110-0211131112321222-3022211300313213-1033123131212132-3213200222020111-2221232302300213"></a>

## Next pages — api_key / 331012012311 / 4

- [new_relic_receiver.api_key.blindfold_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-0312211120000312-0322322322111220-1103110331000011-0123030020200031-1030213100103200-1033200303231210-0021220300121233-3313021331211210)
- [new_relic_receiver.api_key.clear_secret_info](resources--global_log_receiver--reference--group-003.md#canonical-2033113111202231-2211303110023112-0011123201132123-0222303230312333-2131100001321221-1301313200201112-0131031333212002-2210301321033121)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0312211120000312-0322322322111220-1103110331000011-0123030020200031-1030213100103200-1033200303231210-0021220300121233-3313021331211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202313203303333-3100333320130213-1201332002101310-3001000101011312-0232102333033312-2223003130133202-3023310310010203-1112213121120231"></a>

## new_relic_receiver.api_key.blindfold_secret_info — blindfold_secret_info / 103033203321 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- new_relic_receiver.api_key.blindfold_secret_info

<a id="canonical-2032120110010311-0032132333203322-1132112202331322-2220130001023332-3213032222031103-0223321221232213-0033211222213300-3303200123321320"></a>

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

<a id="canonical-1110132220012021-2110222103020011-3333223330330321-1320100131133100-1111222032033201-1002200112110333-2323332013130323-2323112322202210"></a>

## Direct properties — blindfold_secret_info / 103033203321 / 3

<a id="canonical-1310031222201233-1102211023320132-1201213110213311-1102002120320332-1100303320321332-2003302203112310-1220323111202123-2022010231212230"></a>

<a id="canonical-1003233131332231-1022133320121200-3203013210113131-2002333021120033-3002003201020230-2233323120323310-3331231303302021-3313303331132103"></a>

## decryption_provider property — blindfold_secret_info / 103033203321 / 4

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

<a id="canonical-2223200113202210-3130331131313233-0223133322100133-0200013231321032-0231232231203311-3230011230103011-1301013203120111-2321323001131203"></a>

<a id="canonical-1211302010000320-2303310102013112-3201000103220020-0002332132002121-1101030001010210-0023201311331120-0221211113131103-0310001222202021"></a>

## location property — blindfold_secret_info / 103033203321 / 5

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

<a id="canonical-1133112221202003-3322132333313232-0210001032313002-2010103110102120-1123113211023121-0020310303210310-0030131031110202-0232023301112113"></a>

<a id="canonical-3333132033112002-0223332301322310-0220321113132202-0112201233001023-2030012021122211-1002022313332302-3230231213021312-0321133310210322"></a>

## store_provider property — blindfold_secret_info / 103033203321 / 6

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

<a id="canonical-1001213332211132-1311022322013031-2220013113202300-1233001201213021-3032113110321033-0202131023121023-0211313223221131-3202220333323010"></a>

## Next pages — blindfold_secret_info / 103033203321 / 7

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2033113111202231-2211303110023112-0011123201132123-0222303230312333-2131100001321221-1301313200201112-0131031333212002-2210301321033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311311003321022-3322212310220201-0023312321211000-1302223223202133-1131220122121211-0312101223200323-0112212023030030-3323132100012003"></a>

## new_relic_receiver.api_key.clear_secret_info — clear_secret_info / 031021100223 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- new_relic_receiver.api_key.clear_secret_info

<a id="canonical-0030311021310330-0131113210013213-1232010222121301-1122320200202300-0312220001232330-3100302211131301-0113022101000101-3011220003031232"></a>

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

<a id="canonical-2120131131222221-1222022233212213-0332333322333121-1120003233000201-1003300210132111-2233223133132130-1002020322112233-1120132201211011"></a>

## Direct properties — clear_secret_info / 031021100223 / 3

<a id="canonical-0000022200121113-0101123102113002-0030331121313122-1103103311133202-1010213101000300-1321321132131131-0212213131212033-3102120200303022"></a>

<a id="canonical-3101230213210331-2023120012102220-1211121111212210-2302212131000021-2101133210231233-0013212121333101-2211012013122230-0101123233332130"></a>

## provider_ref property — clear_secret_info / 031021100223 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2131333100232312-0221013031230030-0311001311122100-1133113130310013-0212310311112031-1031131203320301-3101021331230132-3222333311200332"></a>

<a id="canonical-1033210100122121-0012031032031310-2002221103010121-1231233201033103-1133002213212303-0333201120100200-1211130122310303-2230133211010230"></a>

## URL property — clear_secret_info / 031021100223 / 5

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

<a id="canonical-0122311033310132-2333111201213301-2022232011123111-0010120220130210-1233332002323131-0021002231021310-3132022123321001-1232222232220333"></a>

## Next pages — clear_secret_info / 031021100223 / 6

- [new_relic_receiver.api_key](resources--global_log_receiver--reference--group-003.md#canonical-0201310023231230-2120100303231131-3301021302020311-1211131102102120-0322221232320303-1001013121310000-2233031021323020-3231333132321300)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0111332302103311-2231321210233231-3312031102230011-1222123203220300-3022101201022312-3033002110232322-1212003333010131-0331232213123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130321023030120-3233302130133321-1020013320213313-0101333330001323-1110321012223320-0012202210233330-0222013022133002-2311031131121231"></a>

## new_relic_receiver.eu — eu / 210223132021 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.eu

<a id="canonical-2102232032231100-0233300020131310-1112120231023302-1031010013300202-1330132220031312-3013320300021031-3302312320130021-2011103002202231"></a>

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
eu = {}
```

<a id="canonical-0032130111121210-1000322000003133-0013301032303012-0023302203323130-2321300012212303-3330000032131131-3321022201302000-3110022230300232"></a>

## Direct properties — eu / 210223132021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223201302233120-3003021001131303-2302031330220231-1132230101333220-1233331023230203-1113010030220301-1113113312133200-3023231010323123"></a>

## Next pages — eu / 210223132021 / 4

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1012122103003011-0312131130200120-1312311010011133-1000220032230110-2330033223201123-0121312311122030-1103322213300021-3102320222222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202133112210011-1102003120001000-2033120100202233-2123023322030333-2323110200211003-1300330220311313-0013313332101001-2312102330103002"></a>

## new_relic_receiver.us — us / 221323322220 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- new_relic_receiver.us

<a id="canonical-1003110133033023-0101310221011312-2323123212111201-1332232223123003-3003100133102013-3000223032322323-0211131211032032-3300012321103200"></a>

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
us = {}
```

<a id="canonical-1320002130213323-1330333222322110-3202211020123102-2032013111230101-3230313203232021-3323322011322020-0322032122101222-3201023130212231"></a>

## Direct properties — us / 221323322220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220320300122300-1331132121320101-1103312200311330-0331331012213202-3031232302232131-1032013210312201-3102310210000331-1122120113213113"></a>

## Next pages — us / 221323322220 / 4

- [new_relic_receiver](resources--global_log_receiver--reference--group-003.md#canonical-1100211030222001-0303013102201103-2231212213022013-2332111223302311-3300333000130122-3030322200000211-3002103202031202-2313122113032020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2110223021133332-2001211101313323-3012032332230023-3123201031102303-3033233222010310-0322300301310121-1032103033320210-2132202001210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313122213320232-1201312020201131-1112203003201302-1102200032022220-1312022300113322-3230323001223302-3113023023030023-1311112320323210"></a>

## ns_all — ns_all / 312000203330 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_all

<a id="canonical-3002133321321130-0230232203010223-1110130322321023-0302022123020231-1112103320022121-3221331330103213-2132031121012132-0211011322121200"></a>

Type: `["object", {}]`. Optional.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

- [ns_all](resources--global_log_receiver--reference--group-003.md#canonical-3002133321321130-0230232203010223-1110130322321023-0302022123020231-1112103320022121-3221331330103213-2132031121012132-0211011322121200)
- [ns_current](resources--global_log_receiver--reference--group-003.md#canonical-0301322210023112-2021310211001030-3312222220121313-1132211011001223-2002213301202033-3323203023001212-3233001223111202-3221030231123021)
- [ns_list](resources--global_log_receiver--reference--group-003.md#canonical-0331201313123231-3032231333310001-3022103022000000-3122210132332000-1233123202200013-1012100010010213-1321022201103201-3120010330312100)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ns_all = {}
```

<a id="canonical-0111112013302032-1332011223021220-1201231201030122-2113113302233012-3323000120120013-3101213302220220-3002020021010301-1313103113331313"></a>

## Direct properties — ns_all / 312000203330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320113322011123-2223030202113232-1222312212332223-1220301101030102-3133202322030001-0020201311213211-3323222131233122-3213012213022012"></a>

## Next pages — ns_all / 312000203330 / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1131102221232000-0300222020002322-0302133211001210-2101203323112122-0311121123102201-1122110013312213-1303121113303221-1111202320322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310232130201020-1020000201123223-2121200101212030-0122333210232332-2020020301320133-2300213232120002-0311110210001323-2030131323310121"></a>

## ns_current — ns_current / 310113232300 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_current

<a id="canonical-0301322210023112-2021310211001030-3312222220121313-1132211011001223-2002213301202033-3323203023001212-3233001223111202-3221030231123021"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
ns_current = {}
```

<a id="canonical-1313212013310121-1301121312002221-3130101311013113-2002023310221210-3003122321320020-1231010300222000-2233220310010000-2231012321201000"></a>

## Direct properties — ns_current / 310113232300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201033330012221-1301103313123320-0021131010033103-3201321020010112-1112101311330102-2220333002210121-1111231230123202-3000223321231002"></a>

## Next pages — ns_current / 310113232300 / 4

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-1323013202223030-2213130333303312-1001111102133131-1001220102221312-3002003012031130-3113012231221100-0102230132130332-0202133321133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313111113201030-0133311132101103-3322303023320232-1002211330201121-2123221001221003-2210332201232102-2121333032221230-1132022211120120"></a>

## ns_list — ns_list / 130313011101 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- ns_list

<a id="canonical-0331201313123231-3032231333310001-3022103022000000-3122210132332000-1233123202200013-1012100010010213-1321022201103201-3120010330312100"></a>

Type: `"object"`. single nested block, Optional.

Namespace List. Namespace List.

Upstream description:

Namespace List.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3221010020202313-1011111122311031-1332313233011312-1222211322301012-1322230023103032-3320112321303103-2200113310221022-0110231102311000"></a>

## Direct properties — ns_list / 130313011101 / 3

<a id="canonical-3311133003311012-2112320033113111-1003020223221332-0023033020003022-1012111233103312-0120211012200222-3302322101001330-1112303001300201"></a>

<a id="canonical-3013132221210013-2020321101211210-3023231213021300-2301112301230231-3231121112100100-0211033230122233-2100122100000010-1331123211003031"></a>

## namespaces property — ns_list / 130313011101 / 4

Type: `["list", "string"]`. Optional.

Namespaces. List of namespaces to stream logs for.

Upstream description:

List of namespaces to stream logs for.

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

<a id="canonical-1120032230331300-3200130330112010-1132112213131100-2020101202301000-1230013231302000-1133301131212112-2221311103102030-0223301223221322"></a>

## Next pages — ns_list / 130313011101 / 5

- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002012330022033-0113301030001303-1001110031010311-2021200322312113-2033330102020111-0302200021330203-2003323020132200-2233123130232102"></a>

## qradar_receiver — qradar_receiver / 123233132233 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- qradar_receiver

<a id="canonical-3323203113301331-2013010003331233-3133212102210323-2032121213300323-3201011213001320-1331212033033120-2133211210012323-2133022310331130"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for qradar receiver.

Upstream description:

Configuration for IBM QRadar endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3313013021310303-0131130233113000-0112000220122012-2233331121133130-3100222323221322-3020010210233131-3330021130021102-3132200113110102"></a>

## Direct properties — qradar_receiver / 123233132233 / 3

- [batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011): complete subsection reference.

- [compression](resources--global_log_receiver--reference--group-003.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332): complete subsection reference.

- [no_tls](resources--global_log_receiver--reference--group-004.md#canonical-0333030003010101-1102331120011301-3023322233120221-1332221330232103-0110203322122232-0233200210022200-0211212203313232-2301331121322003): complete subsection reference.

<a id="canonical-0011323232033200-1111202030133200-2301112132001000-1121023022001212-2201102222023303-2000003311222021-2031022101301322-3113123212002213"></a>

<a id="canonical-0220222021013232-1101333013230123-0232300320003122-0100300223011031-1101310132033012-2300102330100202-1010031012132313-0003330031032021"></a>

## uri property — qradar_receiver / 123233132233 / 4

Type: `"string"`. Optional.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

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
    "format": "uri",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133): complete subsection reference.

<a id="canonical-3102221230012300-3231013320202333-3310011121101311-0102103202133013-0003101102131322-3311302230233320-3323003131023232-0101122331111332"></a>

## Next pages — qradar_receiver / 123233132233 / 5

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- [qradar_receiver.compression](resources--global_log_receiver--reference--group-003.md#canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332)
- [qradar_receiver.no_tls](resources--global_log_receiver--reference--group-004.md#canonical-0333030003010101-1102331120011301-3023322233120221-1332221330232103-0110203322122232-0233200210022200-0211212203313232-2301331121322003)
- [qradar_receiver.use_tls](resources--global_log_receiver--reference--group-004.md#canonical-2003200023211303-0121331320031300-1320122332032121-0113312022330020-1223311122011320-0123233330300233-2101122222133012-2020331322012133)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123320323132121-0033311311011122-0013032133202213-0212323230102322-3323203331231131-3000233201002323-3213310102121231-3323330333100332"></a>

## qradar_receiver.batch — batch / 331020013120 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- qradar_receiver.batch

<a id="canonical-2021331223032232-2031131311130103-3301232121321201-1332130100002112-1021223231300331-1012221203001221-3200312233310110-2012312002332330"></a>

Type: `"object"`. single nested block, Optional.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3030033033330231-3112132030303012-0321112203111320-3330131222213210-1310003120130312-3313123332320311-3023222102030001-1232033300312101"></a>

## Direct properties — batch / 331020013120 / 3

<a id="canonical-2231321230023331-1121002301100021-0322132002100120-2020330002111131-2330022213303333-2201030021011220-0112100013113211-0120220123220003"></a>

<a id="canonical-0032321023213010-2103332210330211-0100011022021210-2233331331120231-3103010202033323-3312023313332001-0122000321030220-3031003200122312"></a>

## max_bytes property — batch / 331020013120 / 4

Type: `"number"`. Optional.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-3330220031010112-1100331333200220-2202201021110333-2223131121033302-0011021222010200-0021322123202133-0130032213120222-0023203323022333): complete subsection reference.

<a id="canonical-0221001313121110-0023321013210211-2032322313000310-3123120023022101-1220312302130231-2132203230000030-1313000031021022-1322310123000023"></a>

<a id="canonical-0312232300003111-0212310102022121-3000003230211000-1212003103133320-0130010110332333-3100121300232133-2113102010222200-1033333323302132"></a>

## max_events property — batch / 331020013120 / 5

Type: `"number"`. Optional.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0231123133131030-1120332010312103-1311321020022210-0321021200012200-1003323330323211-0322202210013001-1332013132130003-1111101101123031): complete subsection reference.

<a id="canonical-2331032123013232-3003003033001121-0103210000003232-0200310101230221-2210321100211313-3233002122232231-3123112222112103-3101231010010222"></a>

<a id="canonical-0130123212031023-2301201332321311-2300311000003230-0300312013333101-2313110123002321-3212123112002001-0232210103331222-2110011321202300"></a>

## timeout_seconds property — batch / 331020013120 / 6

Type: `"string"`. Optional.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-0221001212321231-3231100121130222-0012231333122312-3010022032021101-3232100330213302-2220020300102210-3322001231033302-0310300313133301): complete subsection reference.

<a id="canonical-0231310330332321-1102021221311313-3211323211123110-2033001331120311-3230002303110210-2130020232202012-1221302212211312-2103201301203231"></a>

## Next pages — batch / 331020013120 / 7

- [qradar_receiver.batch.max_bytes_disabled](resources--global_log_receiver--reference--group-003.md#canonical-3330220031010112-1100331333200220-2202201021110333-2223131121033302-0011021222010200-0021322123202133-0130032213120222-0023203323022333)
- [qradar_receiver.batch.max_events_disabled](resources--global_log_receiver--reference--group-003.md#canonical-0231123133131030-1120332010312103-1311321020022210-0321021200012200-1003323330323211-0322202210013001-1332013132130003-1111101101123031)
- [qradar_receiver.batch.timeout_seconds_default](resources--global_log_receiver--reference--group-003.md#canonical-0221001212321231-3231100121130222-0012231333122312-3010022032021101-3232100330213302-2220020300102210-3322001231033302-0310300313133301)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3330220031010112-1100331333200220-2202201021110333-2223131121033302-0011021222010200-0021322123202133-0130032213120222-0023203323022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023310331002320-0321133021323300-2312001212331121-3103221003313213-1212003212202203-1230302121333310-3330300030032303-2020200311302212"></a>

## qradar_receiver.batch.max_bytes_disabled — max_bytes_disabled / 211101013010 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.max_bytes_disabled

<a id="canonical-2103021033123232-0320313103132311-2322133023112320-3121310223102001-1120102103223233-2210311211012223-3302003113103103-2211200031001120"></a>

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
max_bytes_disabled = {}
```

<a id="canonical-0321221111331110-1013101120320323-3300210333120233-2120112321301120-0122320102001021-3213120022233202-2101030112203320-2122231300202212"></a>

## Direct properties — max_bytes_disabled / 211101013010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212001112210033-1233023232002032-3121123310212000-3123000232221312-2110233221213200-2310311230230030-2022212031201300-3122311031131003"></a>

## Next pages — max_bytes_disabled / 211101013010 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0231123133131030-1120332010312103-1311321020022210-0321021200012200-1003323330323211-0322202210013001-1332013132130003-1111101101123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310222010202220-2303222321312022-1310132103323323-2000230130233131-0221223203310121-0033013330120103-1123213133311323-3211312313021323"></a>

## qradar_receiver.batch.max_events_disabled — max_events_disabled / 300212020202 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.max_events_disabled

<a id="canonical-2133021212221123-3010200123200301-3203012132120100-2212111332331322-2122200012111002-0102223202230033-0113232031322223-1221333203320033"></a>

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
max_events_disabled = {}
```

<a id="canonical-0122323222113200-0210132333311000-0033333200301202-2003123333122103-1110310032333011-1230333322202101-3031111122131112-1022111212202120"></a>

## Direct properties — max_events_disabled / 300212020202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130313002113323-0222221300012113-2203013301321111-2001310222320113-1211133111232032-0202132130023020-1011213120330103-0220102012000210"></a>

## Next pages — max_events_disabled / 300212020202 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-0221001212321231-3231100121130222-0012231333122312-3010022032021101-3232100330213302-2220020300102210-3322001231033302-0310300313133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111210230020132-2000332301001102-2132100312222110-3230112132020330-0031223313113332-3111311322310011-0302330231102102-3232213023122103"></a>

## qradar_receiver.batch.timeout_seconds_default — timeout_seconds_default / 302311322121 / 2

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)
- [Property reference](resources--global_log_receiver--reference--group-001.md#canonical-1322330330000012-0203110000231213-0031203023120323-1120213213023112-1123022010033101-0301112213201020-2213113031120201-3132330010200020)
- [qradar_receiver](resources--global_log_receiver--reference--group-003.md#canonical-2233230320023023-0332103111321202-3233300302103030-0310133012303210-0303100301030123-1001211010333123-0231302220101210-1330301133202332)
- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- qradar_receiver.batch.timeout_seconds_default

<a id="canonical-1232330321113311-1212002031333133-1322323221022321-3132001303132001-3221310210013310-0230013132300213-1221312100201230-0112022330011132"></a>

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
timeout_seconds_default = {}
```

<a id="canonical-1122021010213321-1313203303111200-0102302113023022-0323312121031121-3001300120300203-3110030002113300-2323322120300321-0312021200100101"></a>

## Direct properties — timeout_seconds_default / 302311322121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222312101300001-0311212233112323-2003301131231101-1020032102010332-3232212302122323-0300212223321012-0303033002011331-2130111300332223"></a>

## Next pages — timeout_seconds_default / 302311322121 / 4

- [qradar_receiver.batch](resources--global_log_receiver--reference--group-003.md#canonical-2012101000101222-0131312232201201-1131331132111333-2032013223002221-3322201011301201-0021101300330031-1221211203330320-1210300330022011)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md#canonical-0132202000321330-1323313301300101-2310111111111122-2110210002233013-1320333203212032-1331313321331132-2201202321220210-2322330330113211)

<a id="canonical-3301131220021212-2003230020312213-3212012223100002-2102020333300000-1121103102303322-0002301332101111-3120020101112320-0023331201213332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
