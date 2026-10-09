---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-0232201331203310-2122202320032030-0223310010112021-0332133222010032-2221123102300201-0232102233220203-3000111320332100-0030002013203202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132201131201221-1003331331012123-2211001203122312-0111033333330113-1223001022312220-1230032120022121-3111021131201022-1130321112333031"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params`

- [certificates](resources--workload--reference--group-026.md#canonical-1320320332332013-3000002322310202-0212222002210131-1200300012013103-1132121032133313-2310321030202102-2101301330012220-1211233231212313): complete subsection reference.

- [no_mtls](resources--workload--reference--group-026.md#canonical-3311301002213232-3032212233333300-1301201000110122-1211001010000011-2233001000223302-1111201011313003-0233200311222201-2203120030313212): complete subsection reference.

- [tls_config](resources--workload--reference--group-026.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011): complete subsection reference.

- [use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223): complete subsection reference.

<a id="canonical-1320320332332013-3000002322310202-0212222002210131-1200300012013103-1132121032133313-2310321030202102-2101301330012220-1211233231212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-0021213320200012-3030203022202131-3211330332001033-3310110312010132-3213310032200302-3201111301200332-3200200220201013-2330310033231023"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302113200202121-3122333120003133-2322322231102223-0021222231032121-0310303221003123-0121001113002310-3213031311023011-3123122032132112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates`

<a id="canonical-3333031100231023-3032331203322202-2100301110221133-1030110003033221-3113220132310202-1103021131000203-3221210001220010-1230003010300213"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211002010220322-2312020112323302-3330123323030020-2012311303133300-1321010230221200-3133223121320233-3103230301312022-2312311113331002"></a>

<a id="canonical-2313202031202031-3313310013221003-0120001112112000-3231112130202223-3020031321113203-0332300032231112-1222321101323302-3022333010032113"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1033303230012311-0212221303211031-2200231210331023-3313102232303310-1030303300300222-0113020121000033-2301203023120303-1101022332303122"></a>

<a id="canonical-3232010003212322-0110002131221300-3213133233122201-0100133211333333-2101122132022023-3132111331022303-0103123113323300-3222330321030312"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3311301002213232-3032212233333300-1301201000110122-1211001010000011-2233001000223302-1111201011313003-0233200311222201-2203120030313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-3112031122103320-0100200120030033-2013332302100330-0002130201002130-2121303300103131-3321230122223003-3333031211202023-1201203322230330"></a>

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

<a id="canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-1101333300002323-3112031032301200-1320313003203320-2221230102000002-0113200310003323-3231202222203131-2132221330203010-1311221023122312"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-1332121303330132-3310033221212200-3013132312331103-1201300203020321-0122201311131202-0300021023201213-2200111201313011-1022320113221311"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config`

- [custom_security](resources--workload--reference--group-026.md#canonical-2112322132321113-1232102323132021-1020022020231213-2110122023233001-1201220132211203-3303131030013123-2000311312312330-1122210312322313): complete subsection reference.

- [default_security](resources--workload--reference--group-026.md#canonical-2002033023012300-2201103310331010-3321121223323030-3300221332201020-1221032330232332-3030300222212231-0123220002311021-3131123133313102): complete subsection reference.

- [low_security](resources--workload--reference--group-026.md#canonical-3320330113031320-3032132133223001-2320100223311202-0100313133112322-1302013323031323-0332210132001002-1112100122032020-0132321130030032): complete subsection reference.

- [medium_security](resources--workload--reference--group-026.md#canonical-2323202010211333-0330330203002032-2112022123112223-1121301321231310-3203321223213302-3311222300113033-2303220103233010-1031231000201111): complete subsection reference.

<a id="canonical-2112322132321113-1232102323132021-1020022020231213-2110122023233001-1201220132211203-3303131030013123-2000311312312330-1122210312322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-026.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-1001131210210310-2203312023020203-2111230312323301-3213302332221000-1233012112313123-3312103311001131-0201201113222003-2211210121201122"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3130203003102110-2001010022110023-2233122033333223-3100022223302230-0201321032130331-3333312210321130-0202111022123130-3130303301111001"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security`

<a id="canonical-3120003201311231-2330321031033312-3223221300200312-3013332212311320-3321110121303110-3012110123120312-1103312021202131-2233310201203012"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2232213312330312-2302102201032323-2002000230212133-2201123001101122-3011311203132210-1230313021013312-0232111123021301-0231312203120200"></a>

<a id="canonical-3101333223322102-2201220010231113-2003102330320111-1021133213230301-0101330313221022-0013132131210003-3123031201023113-0133331000020220"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-1130303310330021-2123210300000300-3031230320133022-0012001300003211-0212213131320233-1220020003102220-2123301332301332-3322222112033013"></a>

<a id="canonical-2132000003323030-0232220211313303-2002321220210033-2012113003012110-2002203133231232-0231133220231220-0333211231220320-3302233210010213"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2002033023012300-2201103310331010-3321121223323030-3300221332201020-1221032330232332-3030300222212231-0123220002311021-3131123133313102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-026.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-1133330002011220-2301010232112020-0023020101233333-0230000010133313-1323012303312101-1013011202121130-3310213021230210-0322313202032030"></a>

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

<a id="canonical-3320330113031320-3032132133223001-2320100223311202-0100313133112322-1302013323031323-0332210132001002-1112100122032020-0132321130030032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-026.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3121101131323330-3123123030233121-3121110210232030-1111013130101230-2333220320122133-3232203001113003-3232123003123031-2113012301213102"></a>

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

<a id="canonical-2323202010211333-0330330203002032-2112022123112223-1121301321231310-3203321223213302-3311222300113033-2303220103233010-1031231000201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-026.md#canonical-3011322100021322-2121021022310000-2011022031301301-3323233310022131-0113312101320323-0211100303013221-0023223321310100-2133210222331011)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2003233330130323-3031310330000310-3002233331211103-1020223031003002-2201020122103233-1023120222032131-0311100110122201-3102000212011033"></a>

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

<a id="canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3313210303323203-0330000312130300-0200030100322103-0011110301211320-3011112323101300-1302232201312133-2122102231211210-3313310011130012"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

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

<a id="canonical-2302033101310120-1010131322120202-0302210013323130-2322310113310313-2303203302202020-2201003201310010-3011012030032301-3200122000210213"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-1203020312313132-1210030303233101-0000200123202031-3012033011101013-0310303222022300-3113212022323131-1111010123000103-0223031000302233"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-026.md#canonical-0200123201103213-1003301322200021-1320312012332001-2121030121131323-0300131210310023-2310030212002320-3202012120203333-3232021020323031): complete subsection reference.

- [no_crl](resources--workload--reference--group-026.md#canonical-3300112301020211-2210203222231212-0132132310103013-1203301331232221-0133001332203333-1303011100201133-2303231103233310-1030022233323333): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-026.md#canonical-0303222332201101-3103233323020220-3030332212202201-3121311320301310-1032020131001030-1220203212111223-3113303122010102-2323202133221113): complete subsection reference.

<a id="canonical-2232230213213303-2300233121301312-2112001023021233-3101123221010202-0322112233201313-1113101230030030-0301333320131213-3221021310002022"></a>

<a id="canonical-3311131022023033-3233021113230312-0330003113003223-0123221012202120-3210220232020021-1210221332232331-1102103201211210-3320132101310331"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-026.md#canonical-3020100130232111-0122013010332131-1311000123113310-2202031223222010-1132311030133111-0013222211331001-1103222233123211-3123030123230302): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-026.md#canonical-1221132012332310-1013300332321132-0031323313202001-0213002031331321-1232213321322120-1020311333231333-0231001200222130-3130102313201210): complete subsection reference.

<a id="canonical-0200123201103213-1003301322200021-1320312012332001-2121030121131323-0300131210310023-2310030212002320-3202012120203333-3232021020323031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-3012303023110023-2022123320220222-1222003031301123-0100221011000221-1310131120031233-0023233120122301-2122010220002302-0003221000023301"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3000003212221123-3301003103201202-2320123110323111-0311113210322233-3110330323110221-1303001231303110-0303311101012113-3130331100022132"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-2000021302003332-2321210232121331-3201101311233002-3131333131213333-1312111103100133-1130002111222002-2020320231023132-1021211220032022"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002001032303333-0310233020011022-3300322212002201-0322302031232311-2123133301312233-2030031123302323-2030133030301023-3102132023021102"></a>

<a id="canonical-3122010103213203-1130233002200300-0103101203100202-3313332121102021-3231112022323112-2210112330233013-3232210230123200-2212110011001001"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3310011102213121-2233311310000123-3021211103330120-1211222221002100-1313202202311233-0011210022001113-3210110022131323-0323330310112300"></a>

<a id="canonical-2331232022202100-3000210330202102-3332302211323023-0032102113000232-0111031120313313-2101031222233323-2220012330303120-0120220333221312"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3300112301020211-2210203222231212-0132132310103013-1203301331232221-0133001332203333-1303011100201133-2303231103233310-1030022233323333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3033013000122202-1012230102021120-2320010102331120-0223103013302231-0300310221002122-2200303022030233-1333323221212302-1021330133001201"></a>

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

<a id="canonical-0303222332201101-3103233323020220-3030332212202201-3121311320301310-1032020131001030-1220203212111223-3113303122010102-2323202133221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0321332231220223-1223121331102221-1131103330233323-3132020132023231-0223112302101003-0112232120100131-3231333330301123-2301030122023332"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-2201123023231020-1013233011233120-2300213203212121-2332300332002002-1020302101100331-3010131310100123-2013203232012002-1000101321020112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3312310000331213-3232101320122113-1322021331023032-3222330223333032-3323113233132323-2310332023032301-0033021111023210-2213202102320231"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0231103331221003-3131102323212331-1213233012300112-1232332231001022-2331302213133301-2231131033132013-1303033100132302-3100312333211322"></a>

<a id="canonical-1112012000333203-3100112221031121-3011232022301012-1111121302103130-1300222203131111-3000203000010310-2033000120030113-2033023203213021"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3123111100320222-3211210031313112-2322310320130330-3122010330122000-3221301333101112-0033023230333232-3300333231011103-0032300131121132"></a>

<a id="canonical-1303133122002312-1030333320303111-3111323310321203-0003212022031033-1313212330033211-3002111020320101-0202212333012203-1302320123232331"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3020100130232111-0122013010332131-1311000123113310-2202031223222010-1132311030133111-0013222211331001-1103222233123211-3123030123230302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1303120311133112-0103330222233033-1300111301121212-2211102200021113-2302132030211022-2221230011212200-1330112222220311-2113121301132222"></a>

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

<a id="canonical-1221132012332310-1013300332321132-0031323313202001-0213002031331321-1232213321322120-1020311333231333-0231001200222130-3130102313201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-026.md#canonical-0330120220221323-3001303122321320-1021021120022323-3033111012020201-2201223300221222-2120132123232130-3021131112021012-3210022322020231)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-026.md#canonical-1023122310132111-2212330202030310-2112313020200023-3101321202230303-1220121020221312-1210310302131011-2132302200011221-1002023020300223)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3330212011213113-3102223330331233-3303203202322020-0021110013123323-1201011031111112-0002133013230312-1333332310221220-0000213302220002"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1121121311202330-2301100333201113-2201132302001020-2030021001311320-2133123012133311-2330123232311003-3120112230023312-2032123321223322"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-3330333100002212-0331001211230103-3300322022130310-3111012100100312-2202233201222021-2023311202112213-3122132000113221-2130111233320300"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters

<a id="canonical-3212102330320302-0201112232131132-2121103320231230-1313300021231002-0110212200132233-0112223122331032-3030231012010010-2010131100103033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211231332310032-3230201132022303-1131100022031112-3210321213022233-1301020020002113-2123012312103330-0303223103112211-2130300130000030"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters`

- [no_mtls](resources--workload--reference--group-026.md#canonical-0333033212102133-3112203130110300-1000013200233233-2103012201313123-2010100010321030-3223031000302001-3032111010103223-2020311022133211): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330): complete subsection reference.

- [tls_config](resources--workload--reference--group-026.md#canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012): complete subsection reference.

- [use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022): complete subsection reference.

<a id="canonical-0333033212102133-3112203130110300-1000013200233233-2103012201313123-2010100010321030-3223031000302001-3032111010103223-2020311022133211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-3033000000123321-3003302021102232-2033111301030103-1112122131122303-0122021132121122-1110332012323003-3102323003332133-1022120121203131"></a>

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

<a id="canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1312112133322113-1113103012013131-1332101321003120-2201133321130320-0110321101332201-2313202033321111-3200312031233132-0103122121230131"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330033100013302-3212303222322301-1300213331313302-1230321323303321-2031131120101222-1100032301212300-3201000033300120-0003301200131330"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates`

- [blindfold](resources--workload--reference--group-026.md#canonical-2322321122311202-3210133123131303-0021232103100031-2200003111112230-1011021012023032-2031121001222313-2032113103221020-2220311223200311): complete subsection reference.

<a id="canonical-2003101122001320-2331103113123223-3001110123211022-1232212032310130-1013002131232212-0221031113233002-2101312333110303-0330213030332121"></a>

<a id="canonical-3202300112010012-3321031023003300-1223112012131123-3123003002123120-2221223102131322-0100313113303322-3220230320303313-1333311112302023"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [custom_hash_algorithms](resources--workload--reference--group-026.md#canonical-3232333211103021-2010300302022030-0330331102120230-1130103311210020-1021030103121303-0212011033201120-2003231200323301-3301233213200013): complete subsection reference.

<a id="canonical-2201202111112331-2020010031112213-2101133222330002-0211230310033201-0033220331331212-3212320220300130-0333230030121312-1333033202230031"></a>

<a id="canonical-3331110320023303-3332330200012211-3030210210322220-0312212320030101-1111222111223311-3313122200212303-1010202332003133-3102001213220311"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-026.md#canonical-1331112120231022-1211222032010120-1120011331331132-1212030300231333-3233111122102333-2131100303302011-1331201103010031-0231013223230020): complete subsection reference.

- [private_key](resources--workload--reference--group-026.md#canonical-1201203231110213-3311320202222111-3300110331300231-0032211222001120-1323302223331031-0300210011300231-3031101101201210-3122203022310303): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-026.md#canonical-0213112032033003-2223112330333022-0131033230202213-1110002321101123-0030303030321123-1023133210023233-2100220302310011-2120102101012223): complete subsection reference.

<a id="canonical-2322321122311202-3210133123131303-0021232103100031-2200003111112230-1011021012023032-2031121001222313-2032113103221020-2220311223200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-0023122301200320-3113332130113202-2013212221010203-0001133102321003-0310132122010010-0211323200313133-0120203221333201-0012312313103030"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-1221111013031332-1120023333312101-1120100200021303-1030003333002322-3223020031211302-2230120220201313-1123223010320022-1020210111131110"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-3020203310031230-0022330011323023-1213030332231220-3211103130010100-1332223300000121-1231120103221222-1332230311332132-3313330121130322"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-0212111111132133-1313203003013000-3310200200013332-0332111102233331-1121310000233123-1031131121321022-1201331320131120-0122203133020203"></a>

<a id="canonical-2233130132203130-1121030322013121-0332232111100011-1013223003320132-2310023100110001-1030220310302103-2221310113320022-0002323211102203"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1033300331110301-1130010123222221-2032111100203230-1303232133321201-2023010113021010-2033323333220220-1203121323320121-0330221032101112"></a>

<a id="canonical-3120302131131111-3121320023100203-3313102113000023-1321111313032330-1023332001221001-2301322133113212-1302200320301122-0321231210001110"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-2033000333131100-1122201213033111-3311232230031130-1010122301213110-3203332120202303-3110310231312122-3211222222100230-0233022201010033"></a>

<a id="canonical-2133011221312311-1030130000110000-2321020222300030-1033220223103331-1311313131131223-2212102301031223-2230011332022121-1012020102020332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-1322010333203322-3200303211332100-2102213231232022-2023301123221003-0321100120123003-1002210101223213-0031111332320211-0122110130111332"></a>

<a id="canonical-3131120101102100-3010021310032321-3121010200312300-2303330111111120-3311032223023102-2310323313202022-0032111321230232-0013220110021231"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-0231031212003130-1121320300300230-2221212333333202-3013102223000232-2123200210112222-2013110223002113-0123003231313213-1030213312221013"></a>

<a id="canonical-1021131103310230-3212002231203322-3210232330300203-2010210322103312-1110221321200123-3123212331203122-2203310230202001-1303302301002033"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-1032300102232130-1021003303113110-1100200202011203-0311202022313010-0032320121232113-0203032020232223-0321022001132331-2323201311322031"></a>

<a id="canonical-2333310111132120-3123231120223331-3002101030021010-2330123121021122-3303032000212200-1011220332222013-1220032331030100-3323300101230122"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-1110210323112223-1001122221300331-2301222211020333-1110003012013131-1330303221133102-0013210030203230-2002020310323012-3203333313212032"></a>

<a id="canonical-3323000033000113-0102221303012322-3202231222110222-1320120032310012-1012330312021003-0233100122132020-2131000000332100-3002330013222211"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-3310030112101030-3030311101300032-1333200313130211-1103011101222331-2313100013030221-2111310103031030-2030123010213330-2023332331303311"></a>

<a id="canonical-1312110021113032-2203323220301100-2100233000131033-3111312202330331-0123313200211123-1313331210133131-1100031132031201-3233220023321310"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3201030313213032-3020200333211103-3100221022031233-2332320003010111-2001133230032110-0203302121301132-3311022330323122-0302222121311033"></a>

<a id="canonical-3200101231233311-3132333212102212-0230003003220030-2212010332000122-0311110013231030-3233123333231333-1303120303123232-3211211132013301"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-0300020022313211-3310323231002023-3131312232001112-1202230120132132-1120330102221213-1200000202233202-2202303022002113-0102302213130312"></a>

<a id="canonical-3101222101001222-1201222111313223-0122020212003011-2232232302012313-1310033100312131-1202321213233203-3121322220210300-0323211102020300"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-2203310011010300-3130233102002223-2321131010032213-3210203010112311-2321111122200110-1300300222221332-0320100302210022-1231232323022003"></a>

<a id="canonical-2103200112320020-1003232102310202-1101221013210330-2222122300130220-0131020300122113-0230133302332332-2321322133320033-3120331033113233"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0222303123313211-1010222130233333-1312130223122301-2133202022310333-3030031100130310-3010010003221022-0032030121112311-2030232300222111"></a>

<a id="canonical-0131210331310123-3100123220330210-0310313333230101-1330201013311003-1101222010123011-0132111102120113-2121103130121211-2231103230022033"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-1023020102003030-2332021201332013-0213213300332200-1222313032333322-0301032333003111-3130221001100021-0320030110110220-2300313331133330"></a>

<a id="canonical-3113012122321020-2110212300111331-3101012222110211-0000330201330033-3321122230232330-1132100300001220-2221122200032113-3202302002032132"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1001133211310310-3022302023332120-2202212232133103-3330030330132333-1303103123222210-2001023301123202-2310301300012303-0133121103000312"></a>

<a id="canonical-1031211022113332-1331330200322321-2332122123103311-2123013302312323-1202001200222033-3331300102003000-3100321231331003-0033111210000311"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-3100000231210211-3010332210133330-1101010011232303-0212310310321210-0210012323030113-1031210020200032-1033031031020203-1131120230110013"></a>

<a id="canonical-1223010212011202-3233023131313020-1303002013032020-2312013101311310-3200320332311230-1222032122010112-1321213111102011-3130213122222310"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0110013011120231-0312022012032002-1230112322300101-2231221232013200-2013131111332012-2301002303030221-2212223031103201-3231011102023002"></a>

<a id="canonical-2101332333003032-0021202102031131-0020002121200222-3332232112121333-1131332030311313-3303102231001113-0212200002311102-1111333000330210"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-1320231120201313-2130112303123002-1020113023010311-1102301222222313-2132032212332010-3200003211112203-1320320233221033-2321103012230123"></a>

<a id="canonical-1123032031032112-2322321231212002-1323202003133212-1321002220020213-0120021300211023-0000021001213332-3301201322301221-3233010313033111"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1122030220202311-2120333311331023-3323130230210223-2110202300300202-0122032022321332-2232331031131120-1213220123303200-2011213213013232"></a>

<a id="canonical-0212020221220232-3211112022323033-1103032101103113-3032110213200211-0120313231111020-0323302122032211-2202313231103002-3331222302233331"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-3232333211103021-2010300302022030-0330331102120230-1130103311210020-1021030103121303-0212011033201120-2003231200323301-3301233213200013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2323202131003220-0222300223031311-1113311003122313-0313111222211123-0013302110200110-2121111011020122-2220312211100310-3322202112232111"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1212132331211011-3020201132002032-3011013133302113-2133323122312311-3033201023333130-2311020013003023-0302101311001313-0122320203323310"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-1313320302232103-2111222133312000-3212301001033321-3332013220231021-3120330110021222-2102133112232301-0201103332110210-1220012203231200"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1331112120231022-1211222032010120-1120011331331132-1212030300231333-3233111122102333-2131100303302011-1331201103010031-0231013223230020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-2010303312112000-1031230321333010-0011221200211023-1303333123220300-1222013303203020-0323300300333002-0302320101331222-2113321023120101"></a>

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

<a id="canonical-1201203231110213-3311320202222111-3300110331300231-0032211222001120-1323302223331031-0300210011300231-3031101101201210-3122203022310303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-2110201130101300-1000200122302223-1103323120131011-1112303032330310-2111011203320103-1110122121100222-3213202120330221-3332230131230120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131311313322020-3022330210000330-3223311202120222-2301032322131001-0220001003132023-0103200332310202-1322110030013311-1322210130102311"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--workload--reference--group-026.md#canonical-1020133212020200-2200211132032222-3210231121301033-2301211321031012-1033301213312233-3202022303223233-1312103331201210-3202301312300030): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-026.md#canonical-3232221321210202-3302300003213133-2123221002233322-1000131121303112-0201002321230223-1000033230300122-3233332330100031-3231032013102232): complete subsection reference.

<a id="canonical-1020133212020200-2200211132032222-3210231121301033-2301211321031012-1033301213312233-3202022303223233-1312103331201210-3202301312300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-026.md#canonical-1201203231110213-3311320202222111-3300110331300231-0032211222001120-1323302223331031-0300210011300231-3031101101201210-3122203022310303)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2330323012111021-3101132132331311-2303003103211103-1122312030202211-1233102221210321-3331103033011220-2100112211330013-0012223110323200"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030113213030013-2213131330003303-3330200031022133-1230202031113000-2301230200203112-3122313331132112-1320010302112031-2313133103302202"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0102103233222102-3313311123233202-3002332321113022-3113102120113230-2221110130030030-0310223121132112-0211100033300330-0000313220322220"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2211131031212100-2200022033300003-0332032023302011-1100023310333312-1310122110112200-2122022132210003-2310303202032321-1102111312130012"></a>

<a id="canonical-2320121200213102-3103121112321033-0021311100012003-0312003030032021-0112321033223301-3023211133133313-2222211303001031-0100030221130120"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110231113310331-1322213022300330-1033302022003013-1301021022023132-0000301021113023-1013122200133121-2130100100202220-3322003133211322"></a>

<a id="canonical-2320023013333031-2131112221303111-2311033031331311-0013300122310010-1002102211122112-1210201012113322-2232133003201020-2203100221023012"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3232221321210202-3302300003213133-2123221002233322-1000131121303112-0201002321230223-1000033230300122-3233332330100031-3231032013102232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-026.md#canonical-1201203231110213-3311320202222111-3300110331300231-0032211222001120-1323302223331031-0300210011300231-3031101101201210-3122203022310303)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-2011123021201010-3030010221111131-2223110133112330-0303031303113122-0110323210100320-0231223120133232-3122301223222111-1210333121300231"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220031300123112-2103100003133210-3232012020220311-0002222110000001-0313310122000313-1022113200131023-2113210213122031-0000132302321112"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3232323131122303-3001112302032133-1020020132221020-1313223221002320-0011131303232011-2233010220020220-1313002331233313-3320133113222323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1300213322023130-0022301003000200-2323201301110232-0300320311020221-3220230313330130-3010010232211302-1033133212112121-2303120333312222"></a>

<a id="canonical-3201132130111012-2330210201000120-3131013331022312-1313111223032122-3022203300102312-0312021232231032-1113103310213103-2233112223322122"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213112032033003-2223112330333022-0131033230202213-1110002321101123-0030303030321123-1023133210023233-2100220302310011-2120102101012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-026.md#canonical-3032233200023022-0132233302102331-2331112100333322-3230001331120013-1032031020223220-3331330220303002-1303212030003222-0013012321030330)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1323012021112200-0003211320313213-0100011022123000-3233133110312110-0003101210030122-3112200010321212-0331212030002203-3221223332213012"></a>

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

<a id="canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-3231113332220302-2032110232320303-1303220233233301-3011033102132132-3112300032030103-1133111012032201-0301220222030120-2302023223332102"></a>

Type: `"object"`. single nested block, Optional.

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-3122323021312230-1120130030212312-1100022032331131-1133302103122330-2023031233312002-3230210110021213-2010001031102311-1101303310221332"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](resources--workload--reference--group-026.md#canonical-0022301212303232-1023020111013313-2222211133331031-1210321321112122-0220221211222321-1303121123311102-3230021033002300-3333020130202213): complete subsection reference.

- [default_security](resources--workload--reference--group-026.md#canonical-0211231202313000-3133321200002233-1233211321000230-2200032211313330-3230133100231300-2032232300021312-0030013010230323-3311003002110110): complete subsection reference.

- [low_security](resources--workload--reference--group-026.md#canonical-0222200010033031-0223021111310023-2211313302210321-3332320301110332-3220312101332201-0203203222311000-1222312122110212-0200211311222030): complete subsection reference.

- [medium_security](resources--workload--reference--group-026.md#canonical-3112200232203001-0101023013202030-2112022312310010-2301102210210302-1220133213303223-2211030312002100-2210223303220223-0033002101002211): complete subsection reference.

<a id="canonical-0022301212303232-1023020111013313-2222211133331031-1210321321112122-0220221211222321-1303121123311102-3230021033002300-3333020130202213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-026.md#canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0223023112132321-2222323111120130-2310302223113000-2021330120003003-3003101021103302-1103331002023200-3021303023123103-0233113010022200"></a>

Type: `"object"`. single nested block, Optional.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3032013310202203-3122030201013031-0222230201113223-0211300012230330-3331011213130312-2100311212003220-2333100212120203-1122132121111130"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-2320112310301330-3221122111331123-2123310312322121-3132111311031323-2212132113122231-0001132020011303-2022330000230213-0203122122301000"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3003122331303221-2233302132202002-1003101110020233-1121233000033210-3212213000333201-0203303101103111-0132322301032112-2022031020011111"></a>

<a id="canonical-2230133101123011-3131211330212122-0223001112322332-2230232020230130-1331310330020022-2111311000132223-0003002021001000-2100311330230231"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0122321012312121-2203222120032202-1002301232032311-0232322030033001-0232030101320013-3120212031001322-1012030231331323-3132333031000110"></a>

<a id="canonical-1223131021110122-3010031023122232-1221100133331220-2231212323023212-3123011210310332-1233231333312200-0302110320331112-1112230211100103"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0211231202313000-3133321200002233-1233211321000230-2200032211313330-3230133100231300-2032232300021312-0030013010230323-3311003002110110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-026.md#canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-3221302120112321-1230200110100130-0132203300331330-0302003131133322-0120120201133321-1300033330233200-2102120010303222-1123223131020012"></a>

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

<a id="canonical-0222200010033031-0223021111310023-2211313302210321-3332320301110332-3220312101332201-0203203222311000-1222312122110212-0200211311222030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-026.md#canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-3013032103010223-2312332103233311-3121131120032231-3123021201113121-0010221330231133-0022001212020010-3111310230311101-1221231101322121"></a>

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

<a id="canonical-3112200232203001-0101023013202030-2112022312310010-2301102210210302-1220133213303223-2211030312002100-2210223303220223-0033002101002211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-026.md#canonical-2221103232120103-0111202103022323-1130211111203010-1221111300232103-3320202130110112-1013130122200313-3333301331003010-1221320111302012)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-3100130201210311-1022102100023033-3300321310301010-1112013012020033-1033012220032221-0111302123112333-0312021103233112-3003123022312103"></a>

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

<a id="canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1022011101012000-2012003231113313-3231223011111130-1002123001323230-1131200321202102-3103001023221330-1101212210222010-0323330201001202"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

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

<a id="canonical-3030012011300010-3313213102311110-2320112023332022-2303232131130032-0222130033231230-2010122001032233-0130223323111332-2322231322222012"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-2110200222112013-0122331133110023-0201232112211211-3011320230321111-1201133100022111-1113130301320121-1100323000300020-0321020031031203"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-026.md#canonical-1333121212222121-1332202023030131-3312101120111021-2103302132033212-1321303020311022-1333101000012210-1001012213012011-3131013122033131): complete subsection reference.

- [no_crl](resources--workload--reference--group-026.md#canonical-2102302022132132-3320211232113320-1100223320320200-3313311121113201-0113110212310013-1030020032330012-3211130202023322-3010333100010330): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-026.md#canonical-0320201110203230-2112313031203010-1320332131102020-0221132120222010-3223031033311332-1122133303331203-0222031310320031-3033312131032133): complete subsection reference.

<a id="canonical-3330031301321321-2011333102300113-0312131200230331-0332011121001233-1230003303310111-2101320201311120-2222223303202221-3312020030311320"></a>

<a id="canonical-3032020103022211-2230232120122301-2113033032022111-0002021221123220-1312201330023022-3210120101112011-1103031230112110-0321221033210202"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-026.md#canonical-2123220132301131-2032003223203132-3222302302031212-3310312220110130-1212222011012231-0132011000010110-3302003131001200-0313212200013000): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-026.md#canonical-1013103030130312-2203321222123003-2110111120011232-3220333131001002-0323300102120310-2132203022000231-2030030322010110-3133303331202000): complete subsection reference.

<a id="canonical-1333121212222121-1332202023030131-3312101120111021-2103302132033212-1321303020311022-1333101000012210-1001012213012011-3131013122033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3120031203330330-1232111121332321-0120133330203330-3103102011033331-2220023032012320-3221332013330302-2113200112323011-3302103020103002"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1113213032031111-3202102003113231-1011221220303111-3321232113331023-3222122310133100-3112113211222301-2033333310000211-3221013302211221"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-0221112012221122-3002013103131110-1102300233133202-2333233131301331-1203303022011220-1210232121213032-3132202020111310-3032321330330112"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0300133132333013-1330033321201022-1202202122310011-2323031131013132-2230111320120010-3333213131030022-2021313031022312-1120110233010110"></a>

<a id="canonical-2122033332223132-1221133033101111-0232202232001010-2033013120210110-0220220310003302-0322231000132111-2101011202202203-1332331313132111"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3210311102023121-1213313321202000-2003000311201203-2100223001120122-3222310132213030-3022022100233201-0011230233233221-3320222221222300"></a>

<a id="canonical-0113230011101022-0130113031113231-2032131323131200-0022210332031103-2300312311013312-2133232031123331-1330121310201322-3211111133230001"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2102302022132132-3320211232113320-1100223320320200-3313311121113201-0113110212310013-1030020032330012-3211130202023322-3010333100010330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-1223333110002203-2032100101313123-1301102212022021-1303322021033203-0013113101213202-0302111203312330-1212123213032231-1010022111112123"></a>

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

<a id="canonical-0320201110203230-2112313031203010-1320332131102020-0221132120222010-3223031033311332-1122133303331203-0222031310320031-3033312131032133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1020101103113212-1132223221033310-2133123300220301-1212121003323230-1113322130033313-0033303203312302-0110131033302020-0021203320311111"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0112023031133311-1312133301302222-2233302002202021-3100300031232210-3330302013233213-2020323313120002-2311320230330103-2201133130021201"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1223033312330320-3231320221303031-2011101113200221-2332323322333003-1221301033320033-3101132120000230-1211023322121033-1332210202213101"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1122032120013320-3323312203012321-2332002010022303-3002131311321311-1310010201110122-3213321331112132-3021311111001220-0320302201103002"></a>

<a id="canonical-3320312311013010-1013030023230213-3201110123333300-3102121111020210-3131323103323312-0112022023131201-1002100132121023-2212211312332313"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0103203131112031-1312300231310022-1313313210301230-0320010323022110-3021222121113201-3111110032113102-2022233001301212-0203123322313221"></a>

<a id="canonical-1000331322221310-2321311122132013-1321020330111022-1000310111233122-1000011221023113-1021021123330020-1111110020200211-3010032302311230"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2123220132301131-2032003223203132-3222302302031212-3310312220110130-1212222011012231-0132011000010110-3302003131001200-0313212200013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3130301123101133-0031321332021121-2300003021121331-3113300000021003-0032200312113220-1222202100102021-1130112330220121-0322113111101001"></a>

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

<a id="canonical-1013103030130312-2203321222123003-2110111120011232-3220333131001002-0323300102120310-2132203022000231-2030030322010110-3133303331202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-025.md#canonical-2320131312110211-0233203100020221-1102210111313012-2101323032112112-2012023213113210-1213111131330011-1020010101131133-0020231202232230)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-026.md#canonical-2220233222302121-1211133331002001-2312102020200100-3222223333223312-3022013120330322-2001232101131301-2201322330231121-2322331210022031)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-026.md#canonical-0332123103101030-2113203021330310-3213223131221310-3033101331311230-0101012111200021-0313111203300011-3300133111331000-3122222220032022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1000122122311302-0210011312110223-3022103111032101-3111201232021010-3331011303330111-2202020303002001-2312331323202212-2321330320320002"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3013320103000031-0120322000032323-1202030323133120-2101020101130101-3123002100111032-2320203112030012-0110111231333232-3203121033330031"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-0323230123003030-1320000200323201-3200100000130320-3202102032132102-0122203201010301-1221110312012132-0233233332320232-1223313313220231"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert

<a id="canonical-3113032122023302-3003233202133001-1320013311101212-0331103331223033-2221102003003111-0200132220231111-2010202010012113-0032322011131220"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

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

<a id="canonical-2113113320132321-2033321012013120-0113333313011103-1231001311133311-3211220202311313-2332131103331133-1303223301122311-2021102021231310"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert`

<a id="canonical-1123310310021300-0001211013122221-1332023033120332-1103030223030220-2103233333130301-1031322133110123-1310120231102033-3332133101020310"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.add_hsts` property

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

<a id="canonical-2232331212123021-3311110300020312-1130110110322112-2003212330131130-0110301312331322-0133122122213203-1131222110321010-0322112023200033"></a>

<a id="canonical-0010333030303112-2310103321312212-3023211230032323-0102102321321222-1100021001320200-0323013110120310-1003321113230130-0010021231201303"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [coalescing_options](resources--workload--reference--group-026.md#canonical-3220203231112321-1031120121300022-2303232301112203-3023111302322320-2022322113112200-1033303000300300-2202312320222003-3131133303310102): complete subsection reference.

<a id="canonical-3103122202103333-0210131021302102-1121331303000130-3022323323301000-1132322122111213-3312311222131310-1103000021210101-1201313200322000"></a>

<a id="canonical-1102103003012022-2200020000311100-0213302003131222-2230303233332102-0211332120320122-0213210033200321-1202020333331220-1310101200333323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [default_header](resources--workload--reference--group-026.md#canonical-1113120313210112-0011123132231001-0102333102302130-2230133113223212-1322112032020131-2100010322033312-3120332310133130-0000203113313120): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-026.md#canonical-3021221320310230-3213211203201332-0220011331112331-2230322300003030-1223120102312023-0311331011103321-3220011021300231-2001132320230201): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-026.md#canonical-3033000223020323-2103331331102210-0011020130223300-2121102320111301-1231002220330013-0110101313031120-1113102002313301-3230320331312022): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-027.md#canonical-1132111112122012-0113301203300322-3331333231323033-0231310321332230-0212003101112300-3103121101301023-1122002031330001-3201222100302001): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-027.md#canonical-3111212033013022-0013130201203321-3112121303030312-1000312300110101-1302000111133002-2123102233332323-3310121102112230-0011132102231313): complete subsection reference.

<a id="canonical-1132031031132120-2333002031003002-2232122333221123-1311132031211312-1210123130010330-3002310211230101-2002321212022202-3101330031322323"></a>

<a id="canonical-2000001022122133-2010301221010302-0202232001122220-1100301011020113-2010002331023302-1231111101120300-3021000201333111-2000212323223123"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.http_redirect` property

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

- [no_mtls](resources--workload--reference--group-027.md#canonical-1101112113010123-0310031231111212-1320223011302301-2010232232211320-2000100120312320-1231222013230310-1013330031330010-2230323021212212): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-027.md#canonical-2213221202020111-1230222220230033-2233220322223202-0002031231323322-2110220031333031-1202223320223203-2023323103203222-1310012303101112): complete subsection reference.

- [pass_through](resources--workload--reference--group-027.md#canonical-1210221231112112-2312202201021312-0031013011120123-2011312030013201-0020033201103032-3310112223133211-0131332330032312-0201010032013222): complete subsection reference.

<a id="canonical-0201320012320223-2301313212122213-3021110013201113-1332123200221010-2220003011022130-0123311112023011-0033032100331033-3203023032022022"></a>

<a id="canonical-0001200313130311-1121031030120301-2212111230130313-0230011023132031-1232301010003201-2213311220132131-1113300100123201-3321122011313323"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1223231033120101-0131221100101032-0332013230001111-1112321302311300-3211030013030210-2102220013122113-3203310021131111-1232331001133103"></a>

<a id="canonical-0100301322113311-0311003223311100-1012300012312000-3011212302030332-1210021300220330-2332321012103232-1323000102122000-2031102322211332"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0003020303001311-2230102113103021-3230221121232030-3102213022220100-3033033103132210-2300003000230101-1130110212201030-3013003230201233"></a>

<a id="canonical-3232132311020210-0323122213102120-3212130122033002-3133211031100121-0122222322101333-1201310213232333-1302030113232111-3222213213012203"></a>

#### `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [tls_config](resources--workload--reference--group-027.md#canonical-0310210102013132-3210033020300001-1012032303002020-3312202003132121-1201330230102320-2300111322301332-2323200221333120-1102213222113300): complete subsection reference.

- [use_mtls](resources--workload--reference--group-027.md#canonical-2333133220013023-2213202103321302-2200200031101200-1110110132122230-2112303211121122-3013330333300103-3000201200111130-2002030023333122): complete subsection reference.

<a id="canonical-3220203231112321-1031120121300022-2303232301112203-3023111302322320-2022322113112200-1033303000300300-2202312320222003-3131133303310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-2020031232022122-3031023021023202-2200323222301002-1220222200231030-0100133022130131-2011132303212121-2002333112211130-3330302130201113"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-1101020132212102-2303211011123312-0323131320330202-0112112000012031-3013313330201223-2301220023213303-2211301300222210-3021302031221213"></a>

### Direct properties for `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](resources--workload--reference--group-026.md#canonical-1120102102202102-3313020200311000-2202332100312333-0200333211030321-0022200320130300-0111210120002310-2003102210210031-0202210212023010): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-026.md#canonical-2302310321312112-0132230200223021-2103231333003101-1221301130320303-1020002302220311-1132211230100131-1133202313010132-3221023210101311): complete subsection reference.

<a id="canonical-1120102102202102-3313020200311000-2202332100312333-0200333211030321-0022200320130300-0111210120002310-2003102210210031-0202210212023010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-026.md#canonical-3220203231112321-1031120121300022-2303232301112203-3023111302322320-2022322113112200-1033303000300300-2202312320222003-3131133303310102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3333313223313202-3002001022022231-0021010013111121-0321033000013201-0032230231330223-1023311122100313-0330100133110131-2202330230221130"></a>

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

<a id="canonical-2302310321312112-0132230200223021-2103231333003101-1221301130320303-1020002302220311-1132211230100131-1133202313010132-3221023210101311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-026.md#canonical-3220203231112321-1031120121300022-2303232301112203-3023111302322320-2022322113112200-1033303000300300-2202312320222003-3131133303310102)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-0002213231333120-1202013033123320-2030330031030110-1320103013022120-3121222100023032-3112332320332331-0001001201200103-0333221110201230"></a>

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

<a id="canonical-1113120313210112-0011123132231001-0102333102302130-2230133113223212-1322112032020131-2100010322033312-3120332310133130-0000203113313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-2123023212123113-2300131322101022-1131321320203221-2133100013200121-0330201002203112-3000313323032300-0203121301213120-1322023222323133"></a>

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

<a id="canonical-3021221320310230-3213211203201332-0220011331112331-2230322300003030-1223120102312023-0311331011103321-3220011021300231-2001132320230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2001223332002011-1030013010112113-1310103020320303-0312332013100200-2330123211001202-1300202223302000-3312223211122211-3010133202320030"></a>

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

<a id="canonical-3033000223020323-2103331331102210-0011020130223300-2121102320111301-1231002220330013-0110101313031120-1113102002313301-3230320331312022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-018.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-018.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-022.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-025.md#canonical-0120012201110120-2320030110233132-3233102311321321-2202232022223030-2133131131013111-3231210131302330-1022020011332122-3111230313032011)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-025.md#canonical-2203002110120031-3203013111221013-1321101003110222-2103121303322022-0021002301311230-0203100203111101-0321302220100020-1320220233023212)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-2220320330100322-1330210101012132-0301302331113332-3313313132203223-2330123211123011-0222332113113122-0203030132101203-1231012010033022)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-1013332311001202-2231030033023323-1032320323202301-3221203021101103-2123013133112012-0000001233213211-0122001220201113-0001023312330101"></a>

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
