---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-2013301200211112-1100003032311320-0023112202120310-2332113221123213-0021032130023101-0103210313310210-1003120310012133-0111202121002221"></a>

## admin_username property — auto_setup / 021133313232 / 4

Type: `"string"`. Optional.

Firewall Admin Username. Firewall Admin Username.

Upstream description:

Firewall Admin Username.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201): complete subsection reference.

<a id="canonical-0030230322312203-2330322303002013-2213213031103003-1311131321021131-0010231110132230-3300233333320233-0133210320102122-1211002301300033"></a>

## Next pages — auto_setup / 021133313232 / 5

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032112012322123-2331100212120133-0310101320323312-3120221102113323-0231322220223103-3233111100101001-3111002100233302-3321202020312223"></a>

## palo_alto_fw_service.auto_setup.admin_password — admin_password / 302131300330 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- palo_alto_fw_service.auto_setup.admin_password

<a id="canonical-2133330101012003-3220110310231233-0121210023200111-1002312202212001-3001330023110033-2330311203202123-3221100300133330-0231000331331011"></a>

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
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333323231010112-0132211221220131-2100201011312013-1223100112031031-3001331031112100-2222021222133031-0102032011012011-2113322301131330"></a>

## Direct properties — admin_password / 302131300330 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-1030002023230000-3131033200222002-1230310102122013-1130333310302011-0233320133212311-2032133033331213-2103031210002132-1200100212021001): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-2103032110002003-0101102132021023-0233002333110332-3200220133331320-0303102011202331-0222021301312211-0122031202123233-3133112322202230): complete subsection reference.

<a id="canonical-1232020233210001-1330301322322121-1233233120212311-0232112203222100-0231231033010020-3123211033201100-0110111031323020-1032023130222031"></a>

## Next pages — admin_password / 302131300330 / 4

- [palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-1030002023230000-3131033200222002-1230310102122013-1130333310302011-0233320133212311-2032133033331213-2103031210002132-1200100212021001)
- [palo_alto_fw_service.auto_setup.admin_password.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-2103032110002003-0101102132021023-0233002333110332-3200220133331320-0303102011202331-0222021301312211-0122031202123233-3133112322202230)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1030002023230000-3131033200222002-1230310102122013-1130333310302011-0233320133212311-2032133033331213-2103031210002132-1200100212021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112302003010123-0020002100111222-1130212323221032-1322323313111210-3211232111321330-1023213101230332-0211110123311232-1120133230232023"></a>

## palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info — blindfold_secret_info / 111001220103 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211)
- palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info

<a id="canonical-3312321002331313-0321112300220131-1222100300200220-0131201330212003-2312113203330022-3220333111110133-3232102011322122-0001001301230310"></a>

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

<a id="canonical-2310113211311100-3003300221133323-1100221330310321-2022110020311031-0121331201230331-1230023320131113-2322222212121110-3122231021013032"></a>

## Direct properties — blindfold_secret_info / 111001220103 / 3

<a id="canonical-3322323302230030-3023223200122013-1110200033331311-2030321202301232-3300213330133200-1302232202313101-1021112213231303-0200310310311332"></a>

<a id="canonical-2332330011202303-0312301311230010-2133200232003321-1330012011330032-1320101003122112-0011323202012231-1211022232301020-2213231320311320"></a>

## decryption_provider property — blindfold_secret_info / 111001220103 / 4

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

<a id="canonical-3330312333010233-0120002102020002-0212330222003303-1013221130232232-1003222033102032-0132110022030013-1312303102022103-3322310000133003"></a>

<a id="canonical-3223103212032013-1330022011321010-2120332100020232-1122030321202133-2300221320113011-3201211002113212-3003200210000232-3003321121022113"></a>

## location property — blindfold_secret_info / 111001220103 / 5

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

<a id="canonical-2023110012031213-1010322130130203-3021220321011331-2120222032202211-3111132321212100-2233132230231333-2211231321211100-2311313302122002"></a>

<a id="canonical-1013202120211020-0012231132330102-1230023200213033-1310221020021131-0002230003120210-3032121103113312-3122112020200130-0012101023202301"></a>

## store_provider property — blindfold_secret_info / 111001220103 / 6

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

<a id="canonical-1313310312110222-2303222301323211-1230100110330112-2332322102313221-2310321203113100-1322101303020010-0333331230110101-2302113322202313"></a>

## Next pages — blindfold_secret_info / 111001220103 / 7

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2103032110002003-0101102132021023-0233002333110332-3200220133331320-0303102011202331-0222021301312211-0122031202123233-3133112322202230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222023110323302-3123131030033301-0230003030023002-0231002312123021-1032012131212130-1212012131221100-0110221230300203-1323021332113213"></a>

## palo_alto_fw_service.auto_setup.admin_password.clear_secret_info — clear_secret_info / 002102321000 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211)
- palo_alto_fw_service.auto_setup.admin_password.clear_secret_info

<a id="canonical-3230020201102003-2021032330122033-2311223332112133-2200023321023203-3201300223320020-2201113113022210-0313223121122011-2030300031323033"></a>

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

<a id="canonical-3103003111001222-1023301221132322-0122212022000131-2231103322011100-3000110312213121-3021100111110331-2200302101211333-2111000331032322"></a>

## Direct properties — clear_secret_info / 002102321000 / 3

<a id="canonical-3211013132311310-1000323123220322-3313231010331110-3330310111113311-2223310303022001-2213102321110222-2012101231313301-1222230211131132"></a>

<a id="canonical-0202022011022021-2331130010232332-2212100133000213-0031003333111002-1203321301001033-1101123012120101-3200102301232003-2200032001203010"></a>

## provider_ref property — clear_secret_info / 002102321000 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0132311130023313-0333011032233031-2212112101110311-1221320102203033-2130300103212203-2003033101103231-3103331122101111-0233121120222032"></a>

<a id="canonical-0213310223010001-1112310123221022-0113230203111102-3233312210130121-3300121210300310-3122321210020022-2103211013000131-1010231233012110"></a>

## URL property — clear_secret_info / 002102321000 / 5

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

<a id="canonical-3331202320203300-2020133233021131-3100010011101310-2200223122311222-1302103312133212-2203320103230103-1221011302030201-3222121220313032"></a>

## Next pages — clear_secret_info / 002102321000 / 6

- [palo_alto_fw_service.auto_setup.admin_password](resources--nfv_service--reference--group-004.md#canonical-2021013202121333-0332132013033030-0130230331300102-0033212321212113-3001333312121030-0303203233003111-2123010300222022-0131120132110211)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323203302132101-2000330031102101-0013201302002201-2033121303012101-3021212012013202-0123300031303321-1131333233231033-0310313131223202"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys — manual_ssh_keys / 100323203221 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- palo_alto_fw_service.auto_setup.manual_ssh_keys

<a id="canonical-0233122221112122-2121000312332323-0210133332122101-3131322212331211-3322102301130031-1333313202212332-1001210012012320-1022303302302113"></a>

Type: `"object"`. single nested block, Optional.

SSH Key includes both public and private key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("public_key")}
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
manual_ssh_keys {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333010203302230-1232203230031002-2302203003102101-2310111213233201-2310222332031113-2121120023323032-3200210221013312-3020113003300032"></a>

## Direct properties — manual_ssh_keys / 100323203221 / 3

- [private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111): complete subsection reference.

<a id="canonical-3131111213120100-3221233311012020-3203302111201311-0030020012213132-2323311130232003-0231031010212130-3312322203310221-1322032301133321"></a>

<a id="canonical-1020033101233033-2000102320200111-0100011223202331-3003033213331031-2111323300002002-0322220132032112-0210133111001112-2101221020101222"></a>

## public_key property — manual_ssh_keys / 100323203221 / 4

Type: `"string"`. Optional.

Authorized Public SSH key which will be programmed on the node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----\\n.*\\n-----END (RSA |EC |)?(PRIVATE |PUBLIC )?KEY-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0310320302121301-2232313320011030-1213122221002100-3312231331300201-3230333132303211-3230013130213001-3000322321032330-1210130312333231"></a>

## Next pages — manual_ssh_keys / 100323203221 / 5

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213001032122013-2101220332131312-0103223330032113-3332221101100100-1301220320030022-1102303221310333-3221002003231321-3103113211111232"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key — private_key / 222312213020 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key

<a id="canonical-2332213020122031-2210303020333031-0231323012200101-1013100333121130-2123022022301132-3312311231223300-0122213032331233-0221101031102201"></a>

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

<a id="canonical-2033210210222312-3123233212202331-2103110233300212-3320101120320132-0023023101031203-1303110002103230-1301210132200120-1023301233132310"></a>

## Direct properties — private_key / 222312213020 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-0230022231221210-2121313321111113-2231032300322003-0322312022123123-3232232313300301-0301133301102210-1333032310111200-2100330230113222): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-3001322133322103-0330213200003321-1220031001131331-0223110202003031-1032003201130301-2101233322212100-0320323100301301-2113313013130310): complete subsection reference.

<a id="canonical-1310230231322033-0323001330012203-1122231112302222-0012221001003220-0323100222010133-3032333333013003-3221312112313100-0021330133210331"></a>

## Next pages — private_key / 222312213020 / 4

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-0230022231221210-2121313321111113-2231032300322003-0322312022123123-3232232313300301-0301133301102210-1333032310111200-2100330230113222)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-3001322133322103-0330213200003321-1220031001131331-0223110202003031-1032003201130301-2101233322212100-0320323100301301-2113313013130310)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0230022231221210-2121313321111113-2231032300322003-0322312022123123-3232232313300301-0301133301102210-1333032310111200-2100330230113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000133130212331-1303320031033222-0101211101021211-3113313110223211-3232132301212021-2210110211130002-3001010003230110-3330333120122000"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info — blindfold_secret_info / 322102001203 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info

<a id="canonical-1223221033213103-0320031310222133-3233330101313111-1023120120332102-1320303212210223-3100031013232012-1332230300100003-0021230222311211"></a>

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

<a id="canonical-3221112311221303-2333321130300333-2023003001123213-1010132001301030-2031111301213331-0010012111200332-2322312002231321-3133120121220312"></a>

## Direct properties — blindfold_secret_info / 322102001203 / 3

<a id="canonical-1013230101201333-0331303122032312-3322011202230123-1222222313122133-3133222303231021-3322033013320033-3313233110100202-0120010203033133"></a>

<a id="canonical-1130001112132132-3320132313221223-2203133033211302-3213031101021203-2310322301301133-0132003112233001-1103033110303232-3100113303310300"></a>

## decryption_provider property — blindfold_secret_info / 322102001203 / 4

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

<a id="canonical-1310011200102123-0222123131313103-1131321311322302-3330213023031130-2123112233311001-3320002231031221-2310312113000101-1110330201100330"></a>

<a id="canonical-0301210300102321-1023032031221030-0123302333220012-0301112023201212-1303012312003110-3230120031323132-0321120332022322-0321322200131233"></a>

## location property — blindfold_secret_info / 322102001203 / 5

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

<a id="canonical-3003230130321023-2030030102121000-2230200133001202-1330131211233213-3022300310130120-2221012131233232-1312200310030010-3102023013100330"></a>

<a id="canonical-1002311301110220-1221123120021222-3021103302200231-0233312132223131-1333032030102123-0121201100133201-1203321310013223-0002232301120033"></a>

## store_provider property — blindfold_secret_info / 322102001203 / 6

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

<a id="canonical-2333022002023113-2122031030232131-2110223232000001-2322101033221102-2112101132121223-2102223320123333-3331102220200302-3120233210212111"></a>

## Next pages — blindfold_secret_info / 322102001203 / 7

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3001322133322103-0330213200003321-1220031001131331-0223110202003031-1032003201130301-2101233322212100-0320323100301301-2113313013130310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030011122131130-2103131321312220-1003132222310320-3021300312321010-0101310331123012-2101303321102302-1303313002131122-0122012232312211"></a>

## palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info — clear_secret_info / 300010002100 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.auto_setup](resources--nfv_service--reference--group-003.md#canonical-1323331203203220-3311331231103031-1001233001210020-3313130122302010-2233001013101302-3333333033103312-2000002012201020-3033003110033123)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys](resources--nfv_service--reference--group-004.md#canonical-0033320030010322-1202102231033331-2100333321203211-1002102220220222-2103103103300010-1013013232210101-3211202330031320-2120100320132201)
- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111)
- palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.clear_secret_info

<a id="canonical-3112020301013232-0033122013032003-2002011010110333-2010120211101130-0132023022030032-2310322323012211-1022031130333130-1221222333230033"></a>

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

<a id="canonical-1010332120223021-1312102233023201-0310132130000203-2202223303013201-0200311210000203-2332223221033021-0322300100313311-1012120001121213"></a>

## Direct properties — clear_secret_info / 300010002100 / 3

<a id="canonical-1022223202331020-3101321032221302-1000221220113003-2110122311233131-2103212010001222-1231303322332112-3313032330001100-3333113222203102"></a>

<a id="canonical-1213100030231013-0311112110010331-2021323223221132-3332000010313020-1020323103330031-1031323312011333-0212111101102120-1023000001313311"></a>

## provider_ref property — clear_secret_info / 300010002100 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3032102313230201-2212220322013032-0312113122003212-1213302120031203-2032320110312132-1002121301220101-1002002020333123-3120303100133223"></a>

<a id="canonical-3310222113203021-0233331102213210-2120233202132201-0212220331201201-3130122103331321-1112020333131223-1102323110131202-3233102102012112"></a>

## URL property — clear_secret_info / 300010002100 / 5

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

<a id="canonical-2120322112322300-3321002010113001-1001203210003301-0320021301131100-3232000212103101-3301001311232103-3003123230021002-2210101302010301"></a>

## Next pages — clear_secret_info / 300010002100 / 6

- [palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key](resources--nfv_service--reference--group-004.md#canonical-0010113000030201-1103121313002011-2023000321132022-3212203332030010-3112211031112232-3000333012101202-1330130222121223-3023233011002111)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1301111030110103-0221103130313220-0310020330231021-3213300230211120-0300001131331322-1302213200103213-2300100123201131-0022123322310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203122110020113-0100300131223102-3121212002102333-3332301222210031-1202211033222322-1201103013200322-3332223300323213-0212231120002130"></a>

## palo_alto_fw_service.aws_tgw_site — aws_tgw_site / 013311110203 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.aws_tgw_site

<a id="canonical-0012332321113213-2222123132313211-3032030210131202-2223301310111330-2321130132211231-2323012213123330-2132001100232011-1310003010212222"></a>

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300311121212312-2212132211201321-2013121203213320-2303302102101213-1210023222211131-3300221233300301-0333102001133111-0012112231313213"></a>

## Direct properties — aws_tgw_site / 013311110203 / 3

<a id="canonical-1230133202013231-1200313212023030-2012011113113222-0231101010023130-0121113223230231-1322301130323101-3020002200231331-1132302332033311"></a>

<a id="canonical-1312212022203000-1223321020230310-3131320302020031-3211021201220122-2010011000131203-2312003232320033-2102212210012311-2111232022033202"></a>

## name property — aws_tgw_site / 013311110203 / 4

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

<a id="canonical-3122310033303300-2130021021220211-3331120031233000-3231023212221200-2210213110100121-1032121012110103-2221100123231321-0101002330032112"></a>

<a id="canonical-2010123320201332-2011030130022120-0130233231330133-3333210020320230-3133011033223232-1232323123231330-0232300111202023-1133222312300333"></a>

## namespace property — aws_tgw_site / 013311110203 / 5

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

<a id="canonical-2122313220212011-1122203033130303-3001030133311020-3032212213333113-2323011312302010-2310122003122122-3022313123300033-3101112313021121"></a>

<a id="canonical-2231133020120000-0310103000203003-2211231332210021-3220212222332133-1332122212321130-1031303223320021-3303122023021002-0022010303201222"></a>

## tenant property — aws_tgw_site / 013311110203 / 6

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

<a id="canonical-3230213311010333-0312033031321132-1312111120001130-3230211023233110-0212000212321220-2002230223103032-3330233130323022-0322220130211221"></a>

## Next pages — aws_tgw_site / 013311110203 / 7

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1123032201111330-0000230312030121-2011301213013022-0012110231212311-2323223023121032-2121213231232023-1132132013020021-3132002201011002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003312110210321-2133123223233231-0322020123201013-1223213232132330-1201111222330210-2223130113123112-1232213120033331-0313002232121200"></a>

## palo_alto_fw_service.disable_panaroma — disable_panaroma / 330002033130 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.disable_panaroma

<a id="canonical-3002013203212303-1231321022023203-1113203102303113-2220333121002030-3102123003003201-3102033131320002-0302310223121031-3000223212121300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable panaroma.

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
disable_panaroma = {}
```

<a id="canonical-1312112310331100-3311323322001203-1313003122001133-3223332130001103-0133203010220001-0302001100313012-0201230332030101-0002200221011000"></a>

## Direct properties — disable_panaroma / 330002033130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213212033120323-2022130233011230-2122231201011112-0000122321010223-0101103112331321-0133033111101331-0223200302122231-2112311310001133"></a>

## Next pages — disable_panaroma / 330002033130 / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3033330222311121-1332010330203303-1002302201120120-0021101102310211-3031013112310322-1220320330331121-3212331103011223-3332101111321030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332033103301122-2322020121102112-2123201213331022-0312111220112321-0301110011020003-0233303130333210-3030203003302022-2210212122032003"></a>

## palo_alto_fw_service.pan_ami_bundle1 — pan_ami_bundle1 / 313323122012 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.pan_ami_bundle1

<a id="canonical-1322111223223220-1200112101001320-3213321113212023-2111020123103132-0110211002031300-0321112200201223-0003223230011200-3113033102023213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pan ami bundle1.

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
pan_ami_bundle1 = {}
```

<a id="canonical-3021201303023230-1100110001213222-0312333220012213-2033330003001233-0133020302213332-2312231010111332-2231221111332231-2133200223013201"></a>

## Direct properties — pan_ami_bundle1 / 313323122012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130113111333003-1331010211323212-1112323323230023-3002003330322012-1223030220233103-0222001033333233-0012230010302322-0121332013302233"></a>

## Next pages — pan_ami_bundle1 / 313323122012 / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1211032002030203-2321110130000111-0100333311120113-3102313030201023-0131001310210103-3201210103110232-2302102303231003-3203101301001000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323213321102121-1311123202321223-0032023130120133-2223113220123031-0201010022103103-0133030021021303-2132003012032212-2133302313221201"></a>

## palo_alto_fw_service.pan_ami_bundle2 — pan_ami_bundle2 / 131032203022 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.pan_ami_bundle2

<a id="canonical-1122011022100102-3223030111002132-0211100321303211-0212120111031321-2331113023322213-0002303202323120-1012022233331310-3123312112302312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pan ami bundle2.

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
pan_ami_bundle2 = {}
```

<a id="canonical-0230322212333111-3302102300011332-3301211100110123-1331200001002311-0110320101213321-0322101232333321-3022101213213122-0200010300120323"></a>

## Direct properties — pan_ami_bundle2 / 131032203022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111032131111123-1000310131200021-2333221232202031-2223030301333300-2012332311233020-2111031022201330-3330321202203023-1111202122321011"></a>

## Next pages — pan_ami_bundle2 / 131032203022 / 4

- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122032302211111-2130103222101320-0020101320112022-3132032300003002-0220223213222310-3233200112010031-3113222303301330-0102201302121303"></a>

## palo_alto_fw_service.panorama_server — panorama_server / 000323203222 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.panorama_server

<a id="canonical-3103121332133303-2123033121010032-2301130103230033-3210111033023111-1130030220000033-2321220010112202-1313313032031300-1012210333022122"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for panorama server.

Upstream description:

Panorama Server Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server")}
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
panorama_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111123002200332-2121333223300101-1211302331123003-0222212222112201-1020022112102121-3303202213110323-3333330300331002-2133232133012012"></a>

## Direct properties — panorama_server / 000323203222 / 3

- [authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002): complete subsection reference.

<a id="canonical-3103003113222211-3320101331021320-1201131103103013-0103330000212130-2321233113202033-2130221312310210-2002233130210120-0312032133001302"></a>

<a id="canonical-0200133231100321-1031230212032110-1312301032322323-0223232131210033-2232201300201303-0333022030032202-2132211011002122-3322202222031131"></a>

## device_group_name property — panorama_server / 000323203222 / 4

Type: `"string"`. Optional.

Device Group Name. Device Group Name.

Upstream description:

Device Group Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1232210103223223-3200213021033113-0012022103122311-3322032322212330-0311132201233101-3322203230212210-0212320120313231-3131321211020320"></a>

<a id="canonical-1031021302230023-2321113233011023-1132110203030120-3223323020130331-2332131203233232-1330313210213311-1223333001211223-0231033233311322"></a>

## server property — panorama_server / 000323203222 / 5

Type: `"string"`. Optional.

Panorama Server Address to which the firewall should connect to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0023310330302002-0230322131303210-3002331023223211-3122231321010002-1021321000121001-3233031321110112-1102131221101031-2022133030233130"></a>

<a id="canonical-2102031031102222-0123303322223021-2001201111020303-1130122323112122-2310031213123202-1000112222212120-3012211231220011-3232023031331223"></a>

## template_stack_name property — panorama_server / 000323203222 / 6

Type: `"string"`. Optional.

Template stack name. Template Stack Name.

Upstream description:

Template Stack Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2211333133123332-2030233211310201-1303113103330023-3000010031230113-2230312203202132-1203011311020311-0221221221220233-1220223001231130"></a>

## Next pages — panorama_server / 000323203222 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320003301331120-0322312310131130-0013313103302123-0320000230301022-1200202121103311-0302310231113110-1102211001332132-3310011020032023"></a>

## palo_alto_fw_service.panorama_server.authorization_key — authorization_key / 230203211023 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111)
- palo_alto_fw_service.panorama_server.authorization_key

<a id="canonical-1030003121212220-1203003232200103-1332300033221223-2312320130233133-3010010111213111-2100321003110122-0121111202200313-2303032003221202"></a>

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
authorization_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332312012200101-3021221133302102-1120202021230013-3211030333020132-3112233110020201-3220222003323122-3222232222111002-0313102020112032"></a>

## Direct properties — authorization_key / 230203211023 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-3333101311100210-3130312213200232-0121033000230001-3022210101021002-0023210120301103-0221301321333002-3100322002300033-2223311333312310): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-1231012311220010-1321203302330110-1033201303113100-2121021131121011-2212010333331122-2123320021223113-3311311121220110-2333001022130221): complete subsection reference.

<a id="canonical-3120331013210132-1012030222102113-2001100203222320-1330313123011000-3120113120031021-3102123020230020-0032300113133030-1111030231301121"></a>

## Next pages — authorization_key / 230203211023 / 4

- [palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info](resources--nfv_service--reference--group-004.md#canonical-3333101311100210-3130312213200232-0121033000230001-3022210101021002-0023210120301103-0221301321333002-3100322002300033-2223311333312310)
- [palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info](resources--nfv_service--reference--group-004.md#canonical-1231012311220010-1321203302330110-1033201303113100-2121021131121011-2212010333331122-2123320021223113-3311311121220110-2333001022130221)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3333101311100210-3130312213200232-0121033000230001-3022210101021002-0023210120301103-0221301321333002-3100322002300033-2223311333312310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320231333023220-0202313203101210-2211012131301122-2011022031332301-1022311333001323-3102223102113233-3030131210312021-2132323231130323"></a>

## palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info — blindfold_secret_info / 333222033202 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111)
- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002)
- palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info

<a id="canonical-1210022023113033-0330333212220010-1303232121201320-3313102123300333-3300311333100330-0233201110011101-0203323210123121-0320110202313021"></a>

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

<a id="canonical-2113102131212030-0323132012231221-0200132013212201-3311030200322013-2211231112130330-0303131021030233-1323201111011113-3210233303032232"></a>

## Direct properties — blindfold_secret_info / 333222033202 / 3

<a id="canonical-2113011001033121-2132111311022122-3322102330100110-3312033220120032-1231331123323100-3103000223122120-0000320323010233-0101012331321221"></a>

<a id="canonical-2130002311310212-2333033230131320-2120000201301231-2022230133300121-0130031100121132-3230011210300002-2331212113122132-2031332300122000"></a>

## decryption_provider property — blindfold_secret_info / 333222033202 / 4

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

<a id="canonical-2001101231131023-3031022011203101-3122231120100302-3003030110100133-1320023000103030-3303021003220032-1203123021122320-1320030230200310"></a>

<a id="canonical-3213201211332230-1023303311110120-3133101123111013-0131021201101301-3130023032020313-3131320100320212-0120211213312313-2203310200130302"></a>

## location property — blindfold_secret_info / 333222033202 / 5

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

<a id="canonical-2210103213020213-1221213112110312-1101131223123323-0022012330012033-2203212100021112-0210103123322123-2120012231230321-0211232202223210"></a>

<a id="canonical-3313133212331110-3023030221133311-3120102102130111-1022002332230223-1032323130312023-0013301033203003-1332231331022120-2212033310112332"></a>

## store_provider property — blindfold_secret_info / 333222033202 / 6

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

<a id="canonical-0130000102000310-2200310332311110-2201102130331332-0002223123201230-1323123320002102-3331002231003022-3032221121300312-2100002100131213"></a>

## Next pages — blindfold_secret_info / 333222033202 / 7

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-1231012311220010-1321203302330110-1033201303113100-2121021131121011-2212010333331122-2123320021223113-3311311121220110-2333001022130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301300132213022-3110003221001131-1222032123023333-3201121123331102-1010323210331220-1212023030202103-0010230213230032-3233120032020121"></a>

## palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info — clear_secret_info / 202310023101 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.panorama_server](resources--nfv_service--reference--group-004.md#canonical-1011323001103321-2030100221203210-2122232000230032-2223123210312101-0202020002101023-0101121233000113-3313223221110312-1303032221032111)
- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002)
- palo_alto_fw_service.panorama_server.authorization_key.clear_secret_info

<a id="canonical-0232201002030230-0011311321110321-0030011010121223-3311000203000223-1303233320122233-3131022210012030-1023010221321311-0330002122031300"></a>

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

<a id="canonical-2032002032130303-2302120011121003-3011230321133102-3323020231000132-0201111310032012-3130033222232222-2023000002221212-3300321132113302"></a>

## Direct properties — clear_secret_info / 202310023101 / 3

<a id="canonical-2013233103320313-2231202111101332-3310023303011232-0102033032121222-2013031102202030-0310201133120110-0322111302021120-1011330120223233"></a>

<a id="canonical-0032221022223300-0231110302011013-1233101113033222-1223300112231132-0230323023211310-3112102013120333-0000221021030003-0202013133202020"></a>

## provider_ref property — clear_secret_info / 202310023101 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0010212223120021-0301133330003002-0202300011200312-2121301213323210-1233102121213031-2002203123210033-3023000213330302-2213030233013110"></a>

<a id="canonical-3331002322001203-1223322121023313-0103120002030302-2213232132102022-0021110300112312-2131032023012232-0233311030013130-0323111301030020"></a>

## URL property — clear_secret_info / 202310023101 / 5

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

<a id="canonical-3231323003321100-0301203230313233-0232130031201332-1001302313320122-2213011220123133-2111113323032210-3031321232233000-2331011030023222"></a>

## Next pages — clear_secret_info / 202310023101 / 6

- [palo_alto_fw_service.panorama_server.authorization_key](resources--nfv_service--reference--group-004.md#canonical-3223012010132102-1113310123002301-2301030310303202-2221312222001131-2203011311322021-3001320231223022-2012210021000320-2201201311033002)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030112312013312-0002221021302320-0301223103310300-3210232030202331-0020200300120020-1022022032302223-0203230202303101-3322221123102231"></a>

## palo_alto_fw_service.service_nodes — service_nodes / 200112011023 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- palo_alto_fw_service.service_nodes

<a id="canonical-1302321210221003-3323232100312333-0233111302033230-0212301310130300-1122133130031320-3202102310221101-2302120203200030-3220011203130022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for service nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nodes")}
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
service_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232020201313302-2211113022122302-1030200113100110-1001010320200331-3321130022323010-1013102030233332-0212022010302003-3003201100121013"></a>

## Direct properties — service_nodes / 200112011023 / 3

- [nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112): complete subsection reference.

<a id="canonical-2220002212032302-0002212113120032-3002032210001102-0211020211332013-3213213030122222-1112113131030031-2110101132102221-0323113022211310"></a>

## Next pages — service_nodes / 200112011023 / 4

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003202000330302-1200031031021001-0212010210001221-3001001023220300-1203313021103033-0032331222120201-0312311003213023-0111302132330300"></a>

## palo_alto_fw_service.service_nodes.nodes — nodes / 013020230201 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- palo_alto_fw_service.service_nodes.nodes

<a id="canonical-2210110321103121-1011222232120021-3130222320021302-3222211233202210-3133330213101103-1032200120002023-0000001131032233-2120301200330233"></a>

Type: `"object"`. list nested block, Optional.

Palo Alto Networks AZ Nodes. Configuration parameter for nodes

Upstream description:

Configuration parameter for nodes

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name",
    "node_name"),
  validators.ConflictingListObjectAttributes("mgmt_subnet",
    "reserved_mgmt_subnet")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333020001220000-3203032032020003-1121023010300020-0210303131211031-2323003213123202-3233102232030322-1210202230211230-1130122100331030"></a>

## Direct properties — nodes / 013020230201 / 3

<a id="canonical-0310331101023321-1020203311312220-2302333233311323-1200310302321103-3033111000201032-1300123111222023-1121102231211320-2023032232113021"></a>

<a id="canonical-0010023211331102-0230300112223112-1213001032102002-2032220003123221-3132100213013221-1222002032122002-3212210111031311-0102323123310132"></a>

## aws_az_name property — nodes / 013020230201 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region. It is recommended that AZ is
one of the AZ for sites.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-0222303002200312-3130302121233131-2332100130121300-0133023030330132-0212222011321323-3033211102023111-3011002002221003-2310321222323120): complete subsection reference.

<a id="canonical-2321023313313102-3323210311131300-2212100330310232-0011111301010312-2112222101213213-1122130030200212-2311110011022222-2131130010323213"></a>

<a id="canonical-2332210333302100-0211300210031311-2103210123233103-0012030011111031-3212312031010130-3132002230233013-3120112031123102-1300200231212032"></a>

## node_name property — nodes / 013020230201 / 5

Type: `"string"`. Optional.

Node Name will be used to assign as hostname to the service.

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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2010020001033131-3003213032203223-0321221213303100-2233121110220121-0101223323203303-3231000132311120-3121300001230222-2303032301221021): complete subsection reference.

<a id="canonical-0330210203321211-3330131232103111-0102103131030300-3321231200020032-0000001223303310-3031021111212201-1310031302012010-3221200131313232"></a>

## Next pages — nodes / 013020230201 / 6

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-0222303002200312-3130302121233131-2332100130121300-0133023030330132-0212222011321323-3033211102023111-3011002002221003-2310321222323120)
- [palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-2010020001033131-3003213032203223-0321221213303100-2233121110220121-0101223323203303-3231000132311120-3121300001230222-2303032301221021)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0222303002200312-3130302121233131-2332100130121300-0133023030330132-0212222011321323-3033211102023111-3011002002221003-2310321222323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200011310120030-0233112321313200-2213132301230201-0300001320212113-3011011122030103-0330001222211321-3203203312123313-3133113300013120"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet — mgmt_subnet / 331121102131 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet

<a id="canonical-3300303212211100-3223222323322012-2221330302221020-1010123222300000-2103132230233001-2303221320100120-3223300103122320-2032122021331300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
mgmt_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102110033021220-3313300100231222-2331301323130020-3130332133001202-0210011332332311-0221332313222223-1201101000103131-2003201222333012"></a>

## Direct properties — mgmt_subnet / 331121102131 / 3

<a id="canonical-0330302021233022-2203233100301221-3021221103123010-0200022311221300-1121001221013203-0231130111021222-1020020312022033-0313203302233131"></a>

<a id="canonical-0231321030203021-1122022223320120-3201202313121303-3332021200031030-3312030223130133-1013113230331321-1130003130230123-3330131033220000"></a>

## existing_subnet_id property — mgmt_subnet / 331121102131 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--nfv_service--reference--group-004.md#canonical-3011013000210133-3030322003221003-2131101001232013-0212130132212102-1223333301100230-1102023210031231-0131122302222201-3322013303230033): complete subsection reference.

<a id="canonical-1100011310132003-3300232133310212-3132023311130032-1323102212331031-1011230020222211-1223203101200310-0021223123302020-1232030232001101"></a>

## Next pages — mgmt_subnet / 331121102131 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-004.md#canonical-3011013000210133-3030322003221003-2131101001232013-0212130132212102-1223333301100230-1102023210031231-0131122302222201-3322013303230033)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-3011013000210133-3030322003221003-2131101001232013-0212130132212102-1223333301100230-1102023210031231-0131122302222201-3322013303230033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111200210121021-0021003121031123-3321012332303320-1031113221223330-3331331332010302-3101230010011021-3230131223131313-3030321210112301"></a>

## palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param — subnet_param / 020311132210 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-0222303002200312-3130302121233131-2332100130121300-0133023030330132-0212222011321323-3033211102023111-3011002002221003-2310321222323120)
- palo_alto_fw_service.service_nodes.nodes.mgmt_subnet.subnet_param

<a id="canonical-1231310021213332-0211120101313002-2012001021112220-1320120103310100-3030003233311020-3122313321221222-0113210223313010-2220122033200233"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020230021131102-2222223302232301-1203121232331303-1230032210233213-2101203202111022-0312303311211212-0001112200113132-3112313031131123"></a>

## Direct properties — subnet_param / 020311132210 / 3

<a id="canonical-2231301202103133-2302103323312212-2230003210333200-3310001200132120-1122102101133210-1203120333111320-0213111002011020-2033333003311322"></a>

<a id="canonical-1102212302210221-0002223222002003-0330111212133323-1120303131030310-3133010002230232-0313122233032233-0011121002031312-3002230203221221"></a>

## IPv4 property — subnet_param / 020311132210 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3120113103013021-0321121122200310-0013221020121231-1002333011010311-1322030120303330-0220103303102103-2130232303030333-0132210103232212"></a>

## Next pages — subnet_param / 020311132210 / 5

- [palo_alto_fw_service.service_nodes.nodes.mgmt_subnet](resources--nfv_service--reference--group-004.md#canonical-0222303002200312-3130302121233131-2332100130121300-0133023030330132-0212222011321323-3033211102023111-3011002002221003-2310321222323120)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-2010020001033131-3003213032203223-0321221213303100-2233121110220121-0101223323203303-3231000132311120-3121300001230222-2303032301221021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102022003220100-1000203113130212-2110302213021013-0123131000110332-3213203310202322-0012303013001132-1230031313232032-1213312223101123"></a>

## palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet — reserved_mgmt_subnet / 311212332310 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [palo_alto_fw_service](resources--nfv_service--reference--group-003.md#canonical-1220032300320010-0322310310320112-3102020032000031-0230002231133113-1321103031132221-2112211003001211-2001103020013300-3102213200033033)
- [palo_alto_fw_service.service_nodes](resources--nfv_service--reference--group-004.md#canonical-2031200013222301-1113112221323113-0100320310111020-2012332001313133-1212000001301222-1002020103233310-3330332120300223-1122101300020031)
- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- palo_alto_fw_service.service_nodes.nodes.reserved_mgmt_subnet

<a id="canonical-3213003231222112-2333030130012231-0233132333121220-1202133030113230-1102002001023102-2120121322111033-2032313103103121-3310203032113132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

<a id="canonical-0010123230000030-0112103220320332-0310023210203020-3333022001000330-2100311122010031-1033031132130023-2232320000023313-2012013320232113"></a>

## Direct properties — reserved_mgmt_subnet / 311212332310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331022012133030-0312311231010330-2023122032312210-3122022302122021-0331023313310331-0002233201221003-2232320312122002-3012332021210102"></a>

## Next pages — reserved_mgmt_subnet / 311212332310 / 4

- [palo_alto_fw_service.service_nodes.nodes](resources--nfv_service--reference--group-004.md#canonical-0212200110000201-0021202333121313-1302110113231122-3312003003110212-3130310221310333-0020313333101102-2013102120011021-2011112121002112)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)

<a id="canonical-0202232021221122-0000133111310212-2020203112302000-1210111320033230-0001031130310130-3231013003311031-0221323010300330-3130003223303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212120203030301-0200221010101321-2022321302233122-2213330122031131-0110101031101011-2332023112300103-0301020302300010-3001010321101212"></a>

## timeouts — timeouts / 002102221000 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- timeouts

<a id="canonical-2232213200300031-0231011230212012-2030223133101200-3212212323011330-2132013202023300-3200300330113012-3333321303101212-3332213233102122"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202330312111212-1032331201221300-2202022133120111-0002120132221220-0332333320323110-2300023223233230-0032120321100202-0132333013303202"></a>

## Direct properties — timeouts / 002102221000 / 3

<a id="canonical-3222232213020000-2113233003111233-0132003233332011-1003013312011030-2000121310103022-1331033313331031-2330020320223021-2100322221223012"></a>

<a id="canonical-2303323222002021-0021011313102021-2301021123330010-3113330203100110-0101100003303233-3032331200031330-1001332031323203-1133001002230032"></a>

## create property — timeouts / 002102221000 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2000013102223211-2110323122200231-3332122233130011-1113222222023000-0120010230202233-2331332033022313-2230303023200321-3133313220110321"></a>

<a id="canonical-0123212223111222-1221000103210120-3333320311011202-0310011001123311-2200132011203202-3010211111123301-2000211012133202-2013012122302203"></a>

## delete property — timeouts / 002102221000 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1303103300312222-2332033231033211-0011323303023333-3230002311302303-3001110303202133-1201113231302103-0022123020301201-2122022232232323"></a>

<a id="canonical-2031130021111003-2202103310331130-0110321303022210-2130133030010022-2030022033301323-3013030120120123-0232121233213232-2021213021201201"></a>

## read property — timeouts / 002102221000 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2232323323112011-2100000222330312-2220130221023302-0101231132103333-3022311232331110-2230112110031201-3023120213302111-3023232122133100"></a>

<a id="canonical-0001130101032103-1232321111133331-0023012211303310-0101031322123132-1021113123313231-3010221330321100-0032122231210010-0233121333220102"></a>

## update property — timeouts / 002102221000 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1133310331033012-3200112302101023-0101013033030211-2231011322132010-3313303303321220-2121231313113023-1220221132120230-1010102303001022"></a>

## Next pages — timeouts / 002102221000 / 8

- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
