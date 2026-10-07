---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1232102002022032-2021232312323333-3213000012102220-2231332002220320-3120002010313001-1120323111120100-0303101131312321-3301100110102221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3032331132111332-3110123301301303-2330021323310103-1033321122112002-3220032321111023-0032100310102303-2233112220131101-2331000020032120"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300312320112010-3113310031210101-1032111110102221-3232111131131311-2020213100200010-3033221030032132-3301210210323001-1012301023112020"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--workload--reference--group-007.md#canonical-3132102231030032-3330311201300002-0000231202313011-2001232130310101-3030010303022222-2202323223113203-3103110310312321-1230020033301332): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-007.md#canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002): complete subsection reference.

<a id="canonical-3132102231030032-3330311201300002-0000231202313011-2001232130310101-3030010303022222-2202323223113203-3103110310312321-1230020033301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-007.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3201301112233220-2332003210011130-0011021112111211-3002111010223003-1130013110130232-2031333133113222-2033322000313002-2223030233131331"></a>

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

<a id="canonical-3212232103200201-0200111302302132-1012233333223122-0232130233100221-0113211222123031-0013112001312230-3312201102013213-0102313133033102"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3010320011131100-0300311220302012-0330013030301212-0031013113332231-2231223020301220-2111203323320130-2212313300232121-0303112120201122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1001223012313030-1302132323133100-2302320310013220-3033312113000233-1333333223002322-0002000011033122-2333210102233200-3333012312101213"></a>

<a id="canonical-0302320102210321-2211311203113330-2310020103021131-3200102133300233-0312330001132023-1310013201221322-3022322303313231-0002011123233023"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-2132111202033221-3210223321303303-1011311132022112-1320301223123100-1130111233312013-2111201012312020-1323332213010131-3203203112002313"></a>

<a id="canonical-2132201102131103-3332132220203111-2032221221331333-3001232311311211-0332021123110030-2200211321320012-2332031230202321-3121220211132111"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-007.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3311001211021322-0330301013021103-1223122111101101-1323032032110211-0023021321301032-2031302200320033-2202330130130201-3300123210030120"></a>

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

<a id="canonical-1303001202012221-3232302212031011-1213331033030030-1131022320300220-1010021122010110-2113201031322322-0332211233012131-0013310122033312"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0002030111112112-0132300301311222-0103120211133132-0020312221231130-0130222202301200-2232233213222322-2111133113112001-1022110300303320"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1021211320210000-2231310021023122-2220011011120203-1120231212230002-1110211012121320-2302310120310122-3102203121200001-0111002202332020"></a>

<a id="canonical-3213023003220230-1131333003131330-1132000111131022-2223302031023220-3021232222202100-2232021133113130-1103002300232032-0233120120012003"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0212312021133232-3122103133221113-0200111202001131-0201223213231331-2223132121233232-1230222320323211-0303033110121101-1023212132032110"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-0111012211012203-2321213132300010-2221312320221011-0311121122320331-1330100230031101-2210130232223103-2322032012022323-1323223230332312"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010322103321022-1111323110031131-0131110022033201-2022302121000233-3313121031301232-1312330010222200-3111201110210013-0022231333130131"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](resources--workload--reference--group-007.md#canonical-3300213330112100-0310203322220231-0033200003022313-1100221202220000-1331112022112312-0020220210300331-2311111202203022-1000130320221013): complete subsection reference.

- [default_security](resources--workload--reference--group-007.md#canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011): complete subsection reference.

- [low_security](resources--workload--reference--group-007.md#canonical-2012232121033232-1130011220011211-2332303232301121-0200323121313003-3332332103310330-1332020100210301-0322101200333112-1032002123211110): complete subsection reference.

- [medium_security](resources--workload--reference--group-007.md#canonical-0003320232020300-1300231123101010-0003320113021101-2223321032212013-1103032100102230-2000203322111131-1301130203232031-3122230202303302): complete subsection reference.

<a id="canonical-3300213330112100-0310203322220231-0033200003022313-1100221202220000-1331112022112312-0020220210300331-2311111202203022-1000130320221013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1110103000221223-2121210100122110-3001133220000121-3132023001201301-3303332230200333-1012300020232000-0003203331222013-2330303011020321"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303210202131003-2212332103021113-0210333220130320-1102033122133023-3321200331132303-3013301211000111-0300330133311030-3300000233312201"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-1112011223031100-1111000300130133-3121021002011110-3212231230132012-1112000232322220-2021302200211101-0011012232300302-1310022232200013"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0222233211103301-2023021131110221-1131132211001103-3233212033320000-0101212211330221-1323233233022303-0200223310321033-0320011122312311"></a>

<a id="canonical-1000112120121113-2100311131221003-3230322012300221-3113032210111330-3331322030102021-2121021022013131-0121002231023033-1323203313012001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-1210112122202103-1030122112310233-1220330113301232-1202222113210112-2213101331200232-0310022321320230-2303212201012213-2201231320320322"></a>

<a id="canonical-2312222033002021-3330332302323112-3330201002010100-0122102302003312-3301031102033222-2223023312231333-1031102222201031-1023000313221131"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3211121100023102-3203103102232301-1200221330120000-2133311031220031-0010032212313311-0300122113331022-3011322020313001-1012130012021312"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012232121033232-1130011220011211-2332303232301121-0200323121313003-3332332103310330-1332020100210301-0322101200333112-1032002123211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-2302223013322231-1122200010000322-1313102031331123-0321303123102203-3223301221033231-2012020100031013-2120200103031302-2330231321321010"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003320232020300-1300231123101010-0003320113021101-2223321032212013-1103032100102230-2000203322111131-1301130203232031-3122230202303302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-2010223101120332-2120331023133311-2112320303210101-3133232221333121-3201302333111111-1211030213100331-2100000022330010-2011120002110122"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1121012203011210-3333202300113201-0012031100103303-2201301001222123-1013002103022211-1031111023310030-1223311113203020-3123203331330213"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333030021202133-2133220110203203-3130220131203122-1133200120002333-0013311323110121-2201110003132103-2132001003003112-0312131110032203"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-1031022011120122-0113133132231320-1221003323110113-3123123212032232-2012213022312032-2331301301322133-2310203321221032-2232231103001101"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-007.md#canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023): complete subsection reference.

- [no_crl](resources--workload--reference--group-007.md#canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-007.md#canonical-3012021133200132-0121233331132323-2001013132311003-1303021013222301-0222013312123101-0011132232102131-3221210133123022-0300110000120121): complete subsection reference.

<a id="canonical-3212011011113133-2220012321321223-2003020103102002-2030202132203222-2130033220313013-3200320303031000-1021300000330332-3131303112300012"></a>

<a id="canonical-2102100031023323-0300300333010310-3232103102212212-1330232112020020-3230123033012122-0331213102212200-1333031022011222-0311233031312310"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303): complete subsection reference.

<a id="canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3313200202322302-2110333300320130-3011011122312030-2231112130020323-3131010212102120-3113332331121231-2201111333021101-0031223123201222"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332201330132021-1131002300122102-3331131202122302-0130220220003122-1023031330200110-0330202232100033-2303003211223210-0131331332322023"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-2020010122010222-1000311321012023-3221001220003013-0103300333223203-0101102000112200-0221213002302120-3230110102111132-2012111132301303"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-1111220130132322-1202231111300102-1211220130203000-3013231033311031-0232000311023022-3010301211121102-0301313300103232-3002211310133303"></a>

<a id="canonical-3333012032322220-1011001100200322-3013020220210300-3232303200331122-3023103202221320-2332310213002032-0333300110303000-2003231022031230"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-2213203132022002-1121333013201200-2231132103120001-1013001210332122-0303222111333310-3001010100322232-2111230201133022-2100211001302301"></a>

<a id="canonical-2122102003120332-3202022123211201-0222133333130212-0233333310302133-0213220233312122-1111100321123303-3111001333232003-3131103012211211"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-0002213023230000-1232330203322310-1012131302020033-2303301010033000-2201131000031123-0203303210303013-1300030133120311-1020332223200001"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012021133200132-0121233331132323-2001013132311003-1303021013222301-0222013312123101-0011132232102131-3221210133123022-0300110000120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0230223311012223-2220223212212231-0010330210301133-3011222311112222-1233110102130100-3102032012020332-1321001020103330-3123321320231032"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321030321303331-2100120302001222-1220131301031303-2211122021013110-0302231301230110-2321010223333012-0231332313131012-3203331000120212"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-2011221001201222-2233231031202211-3112103013222202-2322031210313131-1213321103323300-1102031200021013-0112320202331011-1320122021311323"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-1102312333200312-2033102230013122-2202000103323132-1303210121202300-3201213100103033-1303220103011221-1232220120012331-0113220023103201"></a>

<a id="canonical-2233023323120100-2101210323212112-0301111300313010-0033020330122121-3303031313232201-1132303011013231-3022113011112210-0301013111212020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1232011102210301-0210202231320021-1132322221122133-1020211302130331-1311013222222323-1210322023133020-3201233020030320-2021332113030222"></a>

<a id="canonical-2230223331122012-0023121021222102-0300111211232002-3203231010200303-0101012233300132-3113222212023223-1323200130312311-3000310122011011"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2202112102302132-0332113313012333-0120330020131030-2131332323301100-0032332332201100-0111232011030002-1000231213010100-0332212331211300"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3033121121130101-0301331231112001-2130120032022331-0232112200131030-1001222001133232-0002200031321201-1321023032121001-3202022321321131"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032123110330301-2200220223110103-1222001310122123-0201100202033103-1123100212301223-3313022101112131-1103320133230033-0110201032232320"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-3012103332330331-1323133201302212-1231010321232132-0131111110023322-2100220003202321-0130232120033211-2111331013112103-1031123123113132"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

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

<a id="canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1123333213031103-0323222223113223-1300011122313222-0320302331210200-2222213323221312-1010222233323310-1303021110130013-3100032113003123"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300021231003002-2300111103320323-2100200313001013-1132300232212000-0120202320130211-1011212220203231-3202000102013222-1122003213002111"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert`

<a id="canonical-1110220120202122-1123002323020321-0210110122120210-0022112001013322-1202212102320021-0311223232130233-1111211323223133-0303333023331330"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.add_hsts` property

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1133230122301300-1322220020221222-0203333011320203-3221112323222030-1333311300321020-1222100002103333-1012030330311223-2331030213011110"></a>

<a id="canonical-3032112303320221-2100231110200202-2122221101130032-0110212132113122-2101011033320203-1212102131233210-1320201333311212-3023133010333013"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130): complete subsection reference.

<a id="canonical-1200233332130313-0222221300300001-1031122032303033-3133132220232302-2321203300200111-3010321003320000-3300322121331320-2321011133010230"></a>

<a id="canonical-1221011221221110-3031323010000302-0212331231231033-0330031010232213-2221020323013301-3031331023030133-2112000120011203-0231122130311202"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-007.md#canonical-0000231210020333-3302022311033333-0031231210123232-1203021320310130-3112121213201103-0010302000101103-3022103033320101-1023211322212211): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-007.md#canonical-0020020301313213-3202033111301102-3330131131101221-2333123113210202-1000201032331321-3010202212333313-1221010200321030-1313031222112030): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-007.md#canonical-2003030013221102-1101111123001021-3301102031002322-2131202112003232-0002112021331221-3332211211210100-2131003132001310-1332001113100131): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-007.md#canonical-1211320032121322-3012022032220200-3130312112122033-3023221212122220-1323331123100310-1133232013113110-3221303012333303-3131223001201200): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100): complete subsection reference.

<a id="canonical-3213030000201020-1010222013032203-2102330031120220-3101231203101102-0322013320320003-1223230320310312-1113312313003103-0103133302320322"></a>

<a id="canonical-1200232303011030-0001211011203022-2233132322130302-0122211323330032-0300103001230221-2102210332200322-2300200100312201-3110120023021203"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](resources--workload--reference--group-007.md#canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-007.md#canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211): complete subsection reference.

- [pass_through](resources--workload--reference--group-007.md#canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203): complete subsection reference.

<a id="canonical-2332322302112100-1231203022302023-1213302002322110-0233010100300100-2033312132311020-0010111222032010-0312132030203112-3120330023032101"></a>

<a id="canonical-1010300021331132-0320002000202332-0321030322021123-1032123313023032-2302120223023012-3233020203312313-2231112033221111-0311031021110002"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1120231103123203-0221112020222020-3310313233233312-3321211303100230-2122300230120122-3233002121032331-0210113330010233-0200221011330311"></a>

<a id="canonical-0101223123333010-3003301333132001-2332223223003212-0132033003301011-0301330230212003-1203100203232311-2220320131010023-0102011213112320"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2123312131300210-2210311301101010-0211022101331300-1122302100322200-0123323320032200-0122110022023010-2233203302030223-2132103200220311"></a>

<a id="canonical-2221013330032111-2001332022112201-3020020120311311-2302001010033321-1313121200232221-3131113131001102-1203022112023303-1313033200030302"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310): complete subsection reference.

- [use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230): complete subsection reference.

<a id="canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0111010302130232-2002313003101133-1113322030010102-0030013330233233-2202312032331010-3123213210002210-2320333031223112-3313301103101000"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212221233231221-1130232321213303-3023000213130322-3323322112000201-1120332030120033-3121211102021101-0230100330132102-0133020323120320"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](resources--workload--reference--group-007.md#canonical-2033100212322113-0030113210213312-3133203001101010-3012133101203130-1131220111120022-1312121312100300-3021123213311110-1001321022032212): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-007.md#canonical-2231332113233303-2301213033203312-0320332322103222-0033131331302213-3012321233100333-1133121231020310-0302311220010310-1211113312320303): complete subsection reference.

<a id="canonical-2033100212322113-0030113210213312-3133203001101010-3012133101203130-1131220111120022-1312121312100300-3021123213311110-1001321022032212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-0120110123010023-3101331200220210-2131220213111101-2021020302123302-2121001100202330-3120022022232023-1202303031112332-0100311121031032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231332113233303-2301213033203312-0320332322103222-0033131331302213-3012321233100333-1133121231020310-0302311220010310-1211113312320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3022000202010221-3301122122311230-0222230033012120-2231300011123022-0313332131130331-0133131102231023-2003311033300331-3113330202230302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000231210020333-3302022311033333-0031231210123232-1203021320310130-3112121213201103-0010302000101103-3022103033320101-1023211322212211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1012132203023033-3030011310223012-3013031010233332-1201120332311332-2232020233012113-3110130121010202-1023230221300011-1220120122310013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020020301313213-3202033111301102-3330131131101221-2333123113210202-1000201032331321-3010202212333313-1221010200321030-1313031222112030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2310230321323202-3132011231201302-2321121300310222-1313002110020332-2301330010130113-1302201303203030-0333302123012223-0232110301231300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003030013221102-1101111123001021-3301102031002322-2131202112003232-0002112021331221-3332211211210100-2131003132001310-1332001113100131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-0011301330220003-2302013222103110-3021301303200313-1011231010123002-1200320111002122-1301203222330102-0313020111031130-1132323020002111"></a>

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
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211320032121322-3012022032220200-3130312112122033-3023221212122220-1323331123100310-1133232013113110-3221303012333303-3131223001201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-0233333212123021-0101303232130232-3302122320111102-3320120133131000-3003101033013220-1301201131223101-1020221121112033-3310031302223303"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2131031222333210-2010000212032131-2021000311130101-3233132320202020-3013312211312033-1212010323221212-1100102130333112-1001301021221202"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032022311021313-2212233130122121-0012320131001210-3300102102330101-2321002222032132-3103001133000210-3032103110000101-3320010211312212"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-007.md#canonical-3032211223130332-0332212123020232-1221300020022302-1212302001302122-0120100223310333-0113322103011110-1132321201023300-2331203211323221): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-007.md#canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331): complete subsection reference.

<a id="canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0313110302122332-2223130123102232-2330230022202031-2021221110012033-2331003010300030-2130123321123223-3131312023330302-3133213313133030"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003032300220010-0120011102322221-1220310202103222-3023213200222102-0023020013303020-2213311132130212-0232300133030210-1031220110213333"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212): complete subsection reference.

<a id="canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2213122002000331-1301030100232220-1212321121331102-3232310230103000-3230232122200201-3213121222133102-0102120232231121-2223322210000030"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200230102300211-2001313210203303-2231332033013013-2230012231001132-3200321003123321-2310230220212010-0321113010111013-0132313310201323"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-007.md#canonical-0203312031021013-1002020000101122-0200200102200011-0131210323221213-1022220113201212-2220013211230023-3223033030111321-0013332333032310): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-007.md#canonical-0120031012131011-3003023003031321-1001023020122332-1131101210103120-2113100303031201-1132201102030321-2010230033030133-1101101312121113): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-007.md#canonical-3302121110313213-0133003200112310-2030321302223310-1331310102232010-1232122230011133-3111223321202010-0231102300321212-3013201220202211): complete subsection reference.

<a id="canonical-0203312031021013-1002020000101122-0200200102200011-0131210323221213-1022220113201212-2220013211230023-3223033030111321-0013332333032310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1301023201132132-2023311113002312-1303300013102113-3030333202030003-3333013212120232-2013311131200032-2113003202030321-2000202011002202"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
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
default_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120031012131011-3003023003031321-1001023020122332-1131101210103120-2113100303031201-1132201102030321-2010230033030133-1101101312121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1132221200201220-3223211132231000-3233110001202222-2102202203003023-0302133213102000-3322212230012310-1030331023300222-1321003033311103"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
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
preserve_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302121110313213-0133003200112310-2030321302223310-1331310102232010-1232122230011133-3111223321202010-0231102300321212-3013201220202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1233112022312323-0131131312102030-1022020023301300-0321310330222210-1023002233133230-2220102231121222-2220200003200013-2223100201013211"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
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
proper_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032211223130332-0332212123020232-1221300020022302-1212302001302122-0120100223310333-0113322103011110-1132321201023300-2331203211323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2223323212133232-1021210321121312-3231032322022223-0111102201330032-0310201130210311-2312103003200101-1011233203030012-1031032200011003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0133030122320133-1302001233320010-3112220233333101-1112332002323212-3132121103313333-1211110200210313-0022320303100313-3310133131022213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-3221030000033213-3302230112301212-3132313212122213-1202201313230211-1120013003331002-0220010112323000-0020322110021110-0000310012002303"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-1021301102331030-2103303130301321-3032313312100011-1013302133312003-1013313323322233-2300300220120221-2120021122130013-0123231312302131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0010130331201022-2223011210220302-1211303222113202-3013230233132112-1120022112310331-3320120302320331-2232023033211310-2023222010223103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0322213321332121-2110123101020020-3332101300211100-1230111011013120-3110222331102301-2320001010033013-0220112032321332-2320123213101223"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130101002033201-2322002031323000-1221010000332231-3122212030231203-0232213333322333-1131301200201102-0011121230022102-2330333211230320"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](resources--workload--reference--group-007.md#canonical-2210213031030220-0331231123310231-1202102311020323-1200032103333323-3313210101213210-3200101102302323-3030323130322020-1033033033223102): complete subsection reference.

- [default_security](resources--workload--reference--group-007.md#canonical-0303012013132120-2031131031120213-0123302322222320-0323212310001233-0131210112032312-3311031330123311-0331310232231111-3221302030110232): complete subsection reference.

- [low_security](resources--workload--reference--group-007.md#canonical-0312213103031030-1103023303301112-1030123112332331-0210233330010203-3213130011120110-1212030232012012-0100232011300333-0200200100320120): complete subsection reference.

- [medium_security](resources--workload--reference--group-007.md#canonical-3213020000102211-2312312032233321-3130231023301302-2320001201111131-3001221010231300-3231001303001133-3030301330133032-2021102012000320): complete subsection reference.

<a id="canonical-2210213031030220-0331231123310231-1202102311020323-1200032103333323-3313210101213210-3200101102302323-3030323130322020-1033033033223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-1020300302123203-0300133101321211-1302132302301100-3011311222002120-1023230023122313-1301213000103210-3133311130231100-1010122310200201"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131120032303302-3013121000131223-3320013301131110-1113333032322231-0120121103230013-2233323003030010-1201201221003110-1133021013002300"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-1031131322332121-0123232220131320-0213212023101232-1212111203021110-1102303332012021-1321301121012023-3123111322032123-1302110132230322"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3203103011233330-1322311001323323-0120020002212133-3212111200003223-2000311001002131-2031023020020011-1201320131133121-1010331330121223"></a>

<a id="canonical-3131003133000023-1202132000333230-2332132130131023-2201001003331203-1020033203133122-1010230331133013-0032323310120212-0323032313201232"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3131330200032010-0202110223102101-2030311232033010-3012233022223222-2313212212202030-0111031300233331-1310012011321032-2232000123130123"></a>

<a id="canonical-1003003323103030-0333012303131002-3310200333031132-1111121300113233-0233223233231232-2310203323200002-2033222130001031-2231330333330321"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["TLS_AUTO","TLSv1_0","TLSv1_1","TLSv1_2","TLSv1_3"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-0303012013132120-2031131031120213-0123302322222320-0323212310001233-0131210112032312-3311031330123311-0331310232231111-3221302030110232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-3021333200232033-2310010131011120-1200033223003311-0322102223200032-3323103103120310-0123110023302333-3201220013111222-3211220111123213"></a>

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
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312213103031030-1103023303301112-1030123112332331-0210233330010203-3213130011120110-1212030232012012-0100232011300333-0200200100320120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-2232331032032332-0020330311232221-0122200122312110-1302321033100210-2333101202331033-3123301213013123-2030102232021013-2123211221313232"></a>

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
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213020000102211-2312312032233321-3130231023301302-2320001201111131-3001221010231300-3231001303001133-3030301330133032-2021102012000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-3330333002231303-1130200111020221-0132332133301020-3032023233121010-1123131100301030-1011300112013102-0030302323022222-3322332132223230"></a>

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
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0112132223201103-2012210001022101-1202101022210311-0002300202133133-2223011222011021-0030310321030301-0330302031303301-0032333012312233"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213312103113311-2020111022213220-2130030302200123-2230230301120113-0312300130022332-1223201001330311-1322130231023100-1122332310202303"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls`

<a id="canonical-2233113303130031-2232032203322212-0022002030121113-2313023203003220-0022023211111011-1232120130122101-2323312330321020-0123232213112133"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-007.md#canonical-1013302010320000-1011230322021312-3012323210131123-3013312012031311-0022332332220121-3220001222033012-2013223323201213-3301300020232222): complete subsection reference.

- [no_crl](resources--workload--reference--group-007.md#canonical-3302231003012110-2000302332123203-3313302022111322-3230310203330311-1122230012120113-2313102032030111-0213000203031203-0031301310033200): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-007.md#canonical-3312221002221111-0222020203132233-1233200001302233-3002223000011203-3312232220103132-2323300230323022-3230031300330032-3121332103113222): complete subsection reference.

<a id="canonical-0210032111200002-2230323120233123-3110120112131301-2022332013113302-2220111120220221-0310111133023213-3100121122103033-3220123111211112"></a>

<a id="canonical-1121021012130101-3232100223223011-0321200030313211-1323200222020132-0231323013030310-2202200100313322-3231000123100001-3021300333111313"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-1233303310133222-2233301300010300-0330120213311132-3312323203123231-3003031133123110-2230220233101011-0112102000310033-1130101303113222): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010): complete subsection reference.

<a id="canonical-1013302010320000-1011230322021312-3012323210131123-3013312012031311-0022332332220121-3220001222033012-2013223323201213-3301300020232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3302100103232001-1113213111211313-1100221212220202-0130110321130212-1233010013230331-0201001232032221-1211212120020002-3003313330113230"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313332201233111-1323102313120000-2330031110232101-2200120333121122-0103110103032301-2232021033000130-0122311201232313-2320313121310312"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl`

<a id="canonical-3010133121312001-2311230113122301-3322120121200122-1110211033023330-3032130100113112-2321001123331122-2310023111123131-2020303123222011"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.name` property

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

<a id="canonical-3201200331012222-0031122202001301-2021211312112033-1220232013001022-1102030130312031-2202213322310001-1231330303000112-1132120303310301"></a>

<a id="canonical-3223020332221202-1031301001111102-1220133221310223-3102003300022330-2130010310301201-2103223300212013-0233233133320012-1221312300312102"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.namespace` property

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

<a id="canonical-1032332220233002-2033303230022312-2022333132322301-0223032232212301-1031221032113332-3231311002123233-3231300131002100-0302301302131313"></a>

<a id="canonical-2320223111121132-2123113021233023-2311321303113111-1131022200303011-3010231311030030-0123322120111231-2131003122321230-2323311333222123"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl.tenant` property

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

<a id="canonical-3302231003012110-2000302332123203-3313302022111322-3230310203330311-1122230012120113-2313102032030111-0213000203031203-0031301310033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-1003120012203210-0110302111102201-0300223220010010-0211020010000012-1220332133323012-1300202303202032-0003121031033232-2210202123030232"></a>

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
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312221002221111-0222020203132233-1233200001302233-3002223000011203-3312232220103132-2323300230323022-3230031300330032-3121332103113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-0011023310313320-3232231100312221-2313112022121131-0133230202123201-1231302211210033-3210313301322112-2033002011012201-3201000120232230"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231100310313131-3310103301033311-0111211222210110-3311021220021032-1021011223111222-1220203020012000-3311331211310103-2033121310222033"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-1131201020030221-2121132012000113-2120013303131220-1002322300133232-2031210010312010-0110033221122321-3210311213230233-1211303213200121"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.name` property

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

<a id="canonical-3120311003232110-1103313233212022-2333031321223311-0000021030233333-0112012202321320-3323023222123300-2202323021230330-0000332113100310"></a>

<a id="canonical-2322302003011223-3211131121111033-1013132100312103-0003102202001201-2302130203230000-1030320020002330-2233232102133200-3220310301331000"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1320221100331001-2010031132021000-2232302213103202-2133320330332301-2032202212220113-0202333303133010-2112311113320013-0012233030331323"></a>

<a id="canonical-1130000122333003-0130300230321131-3133023021002030-0331120232221321-1123303331122112-2031102310202312-2300330320130313-0311321213011311"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-1233303310133222-2233301300010300-0330120213311132-3312323203123231-3003031133123110-2230220233101011-0112102000310033-1130101303113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3212000331322202-0230222221111110-3112112003133323-1201003233203003-2100010313100202-2021021100122312-3021303020233213-0001013330320212"></a>

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
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-2223222212002233-2310112010203311-1133332112313122-1001013321021330-1022333221301332-0121010211131122-3333211020211030-3212002303331220"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313100112102011-1313320232103232-0220130111102023-1021131331210121-3221013003313313-0311203010312022-1021001031310002-3230110323322303"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-1101111110132231-0330231032110211-1112023223222202-3320212212101131-3232303002320223-0332112001010300-1010120110213222-0012331321003132"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

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
