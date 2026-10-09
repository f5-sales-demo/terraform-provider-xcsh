---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1313320312301230-1321010323120200-0002332120223322-2133010302312133-3312210332020100-2211121311010221-2323003003120221-1322330302012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-3122233020332300-1011323220102121-3213212323131332-2333212110211303-3203132100121101-2323211300311023-3200123030222131-0021311131330021"></a>

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

<a id="canonical-1203201320130323-1302320132132110-0032113120132211-0323100112223112-0002323201220000-1122103001023320-2200233202221232-2001321112110112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-1310201203210320-3212123002132201-3311120222033220-1211003021103131-0233322212003321-2231310022001131-2023122323232123-1320232213200223)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-3200132312110323-0001300123213120-0202213001322210-0123033133100021-3232313132011133-1230131312021120-0130300310110202-1101130201121211"></a>

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

<a id="canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-3301111132132111-1131322222100133-0100230010002201-0103112322233200-0131300202121123-2023232122012102-2333220220333311-2220023120011223"></a>

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

<a id="canonical-1320101222200032-1133311121232122-3333202321110230-1320310102132310-2313233213202011-1132113210000322-1012123331022103-2021233022321312"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls`

<a id="canonical-3223323321033131-3110331101202113-2312232311100131-3003003331231320-1302121320202012-3021002202003100-1111110311320012-1212023312121133"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.client_certificate_optional` property

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

- [crl](resources--workload--reference--group-007.md#canonical-3330320011212133-0132212322323321-3111033020332330-3120021223213203-1130322223011331-0120222032003120-0021111220013002-2110132100001201): complete subsection reference.

- [no_crl](resources--workload--reference--group-007.md#canonical-3132232133022013-0021101022010033-1023113222130002-1023230121303133-1030320221221321-3030110332130323-2110003122131202-2231113210212233): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-007.md#canonical-2102323330320121-0322003333331211-1312032013333003-3202230030023233-3131111310133110-0000321133122312-3303301110000232-2013002001331313): complete subsection reference.

<a id="canonical-0032032011222211-2100232033030233-2001122321201333-1210231033120002-0110033111322013-2100322110022012-3012233010332123-3332010021202012"></a>

<a id="canonical-1123312202312100-2321333003021010-2330303321302230-0000233131112020-3022001332332103-1213322111011330-1132330323013210-3030102220331032"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca_url` property

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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-2333232120322321-2130320212013102-0032111131233131-2021102130103132-1021322121001131-3211301012132310-1233002102010132-1200020200232221): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-0002132221020332-1110113000033220-3131330132113303-0331022100321200-1221302021011012-3300210130222002-3021131021010232-1031120312220302): complete subsection reference.

<a id="canonical-3330320011212133-0132212322323321-3111033020332330-3120021223213203-1130322223011331-0120222032003120-0021111220013002-2110132100001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-007.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-1021320010031111-1111302122122003-2123201022322211-1302030003330121-1003301311303032-3003020221312010-2030121330013001-1130203231201121"></a>

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

<a id="canonical-2012330100030110-1330013022233233-0221200123130323-3201112111110222-1021312023323022-0313133221021330-2232200120230302-1211130013003223"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl`

<a id="canonical-1333032231131301-0022333331023030-1220212310001110-1212232033011210-3111013010230003-1133111011223131-1012110232131010-1013210331111312"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.name` property

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

<a id="canonical-0200220011113023-2101110001311210-0300122100031333-1201122200113002-1300200100223010-1223021022122130-2330201000211202-0330323332223102"></a>

<a id="canonical-0233203033002223-0030132332011002-2111231000003111-3221023222110113-3310103103312233-0110323222120221-0020223333103300-2301110132223211"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.namespace` property

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

<a id="canonical-0223330311203131-0113322012200000-0331100033200233-2123311023203030-1132002032031212-0332111202311221-2110313201313212-2210330133323302"></a>

<a id="canonical-1300331233231322-0011031110010021-2233103221112331-1312321130130220-1230210323213003-1013010103212202-1021112122230021-1202331321130310"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl.tenant` property

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

<a id="canonical-3132232133022013-0021101022010033-1023113222130002-1023230121303133-1030320221221321-3030110332130323-2110003122131202-2231113210212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-007.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1000111120300012-3100232222021002-1202211133301211-2300313021000310-1121203330203230-0122102203000111-1132220233002010-3203203313132110"></a>

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

<a id="canonical-2102323330320121-0322003333331211-1312032013333003-3202230030023233-3131111310133110-0000321133122312-3303301110000232-2013002001331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-007.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1232030221300303-0133332310002011-1302320110311201-1201220323230230-0313300100222330-2302202311002133-3131030111332001-0333020031003111"></a>

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

<a id="canonical-1003020300321132-1230020010020311-3221023003103232-3323030031102321-3323312323112221-3012231232020022-3322021013330012-2023331101212033"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-3012312130000001-0202201232020112-0223113023030300-2001121320111030-3102022003102301-2311210020020302-2133003120020012-2103132303002300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.name` property

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

<a id="canonical-3103300333020212-0200003201220212-0303220101331002-2312121001201333-2311122031331131-0220311020121203-0110120011023131-3122030211011030"></a>

<a id="canonical-0231100020202133-3231310133002110-2330332013121002-3003030302133202-2110300200301222-0120021111332023-2333021102313222-3112003130122201"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-0133130001101122-3323100013020112-3203220122031333-1032212022201113-0123233230022221-1030021003013010-0121122131211112-3332112232010312"></a>

<a id="canonical-1033102131203020-1132131111211023-1031302132031130-2211120131112102-1322200201201111-2131102302020123-1013103202001330-3222102201122320"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-2333232120322321-2130320212013102-0032111131233131-2021102130103132-1021322121001131-3211301012132310-1233002102010132-1200020200232221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-007.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-0222132333331121-0010012010230031-1230201012133123-0000312031332032-2103111131130011-1231222313030010-2320300013033120-1130232303002220"></a>

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

<a id="canonical-0002132221020332-1110113000033220-3131330132113303-0331022100321200-1221302021011012-3300210130222002-3021131021010232-1031120312220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-006.md#canonical-0103133333202221-1131021113321323-2122021221232201-1013312322320320-2313333220100330-3011001333202103-0121120020303322-3332000000101010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-007.md#canonical-1101333300120120-3120122223232231-1210100331033331-3010100230302312-0312021231303102-3112111210033200-3310010333223022-3120311223122211)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2130133020111311-0303020021310112-2312333003023103-2233000222320312-0222310132323121-1333212110011201-0231301123233113-2013102221322320"></a>

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

<a id="canonical-0322112322321302-0201021233231230-0032122330033332-0020203011333223-1311013223320210-1321331310203230-1210022111301132-0202210102022223"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-2130130231311312-2222012232300120-2222011103313000-2210301102123113-2322301302300022-1333333321313313-2033211212303211-3102111033222122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-0303010130312230-3021223130110321-0021221232100311-2302032023332130-1302330232313002-1110301000333030-0113030001002110-2032302021021100"></a>

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

<a id="canonical-0103331231200020-2233312333121302-1303320103012322-2111130000323012-2001000130313200-3332331132011021-0101020102220201-0002221231000302"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters`

- [no_mtls](resources--workload--reference--group-007.md#canonical-1222103023212012-1003000121120210-3010300321130303-2020133030103020-0012133331000201-3033100121103313-3021202002302001-2031200021023212): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322): complete subsection reference.

- [tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231): complete subsection reference.

- [use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123): complete subsection reference.

<a id="canonical-1222103023212012-1003000121120210-3010300321130303-2020133030103020-0012133331000201-3033100121103313-3021202002302001-2031200021023212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-1312321001311222-2303333000302303-2312120103110312-0222200011221232-2133123101010231-1010223221213211-1330232002321201-0010132032031221"></a>

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

<a id="canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1303222310312331-1213111002223313-3012320120012111-3002031221013202-1120012100003010-2313301232113010-3013132310003200-0022230211022323"></a>

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

<a id="canonical-0122111332203003-1133022203001132-2021131030301330-0003102230120132-1133321111323321-3223223130200020-1122012111121132-2203331011101322"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates`

- [blindfold](resources--workload--reference--group-007.md#canonical-2201300111232030-2000213230220101-3320220231010011-0001111230300101-2101122311223021-1232100113232301-0232200130230012-3022332213112111): complete subsection reference.

<a id="canonical-1020103102133130-3300223001322123-0323123031022110-2122000202211320-3300303330333222-3132031123131332-2232332320130311-0200021210123122"></a>

<a id="canonical-3313200022332211-1022313030032120-0030001333123310-1133223320102211-3210221321102123-0111301031310313-0232232102213130-3103213020033010"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

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

- [custom_hash_algorithms](resources--workload--reference--group-007.md#canonical-2231110321222212-0201111132112130-0013003032121100-2012323122331020-1223032133221000-1031230012022333-1331123330031200-2032123003310201): complete subsection reference.

<a id="canonical-3023210213100013-3011032201000131-1331203203110203-0302222100102032-0230331123023113-2102221201132202-2312130212212223-1220311230231300"></a>

<a id="canonical-3212002001012332-2323330221132220-1120320310030120-2230210001323010-3030010210112221-3102003233120103-3112023300131231-0111122032223020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-007.md#canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321): complete subsection reference.

- [private_key](resources--workload--reference--group-007.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-007.md#canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310): complete subsection reference.

<a id="canonical-2201300111232030-2000213230220101-3320220231010011-0001111230300101-2101122311223021-1232100113232301-0232200130230012-3022332213112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-3310202211320211-1212320331212103-2322220310110321-3210132113312013-1231223011132023-2211110333332000-2212330110321010-3332222021100001"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-3123303122231331-0012222300300200-1110021321230233-1302020331322022-3223032000122132-2320212232223111-1330000222210202-2111301003130202"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-1302331022000230-0300231302201202-1222003012100010-2023122321331300-1223122222023332-1032101221133323-2000203101132103-1101232233320330"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2301010131011103-1331111000012130-0200133020212210-0210101112002212-2320320103120113-0212330102303303-2032222103302213-1103312020202021"></a>

<a id="canonical-0133322323221212-2212310323110320-3222101201221010-2231233212320203-0333220122103332-0133130310213030-3022100003131331-3230232103301132"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1130300313201301-0230023332100100-0031200131202110-0132331012200000-0221203031112223-0203312330123321-2300131033132112-3111033313031103"></a>

<a id="canonical-1333130001000201-3010033012032001-0032101003003302-1031023213033103-3320312031210023-3320302223011201-3323221220212232-2102333002330312"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-1231220322201311-2232033100222200-3112312311122312-3323223221311133-1302120033030320-0202302021103011-3321323222023213-2131102312303120"></a>

<a id="canonical-1133302203300103-0002230123311301-3022223132313000-0311230231302110-3310100130030203-3201313120101320-2311331213331221-2011000010223220"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-0021033233111332-0013220301232112-3321133011121012-0011311221233213-1330323330002232-3312021200323321-2013322023222201-3302120301213203"></a>

<a id="canonical-1000210312110320-2111133331232012-1132320003002310-3103213013010012-1231013122323110-0122300202130311-1332222103001202-1233131022231021"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-0102233031303112-0100331010233220-1013303113233322-2100120112001321-1231210122311313-3202230310002321-3100110230211003-3020100130020221"></a>

<a id="canonical-2120000310213313-1122213133202203-3222120100111311-1230300233232300-3020022200323033-2103311313131123-1000032131133020-2210311121021113"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-3220113120312323-3133133302203131-2201303013010311-2131123320003223-1320130301123322-0112113212120133-1022033003122103-1220122101021031"></a>

<a id="canonical-1130200112120300-3002131310230003-1032201013033312-0031230020010331-2301123300010223-0020221300233012-0101111230001113-2232331003321210"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3233200311201200-2133000222033200-3301232332032330-3121022221221212-1113322013223200-3223103220301320-2202200030032331-1223230303220130"></a>

<a id="canonical-1222233202211100-1303031221030010-0020330002210233-0100332201120303-2032123103012323-0313113333331100-1031022212103032-3301113003130322"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-2112231012122300-2223312021131132-0223333222000011-0132131031213033-1332201203203112-0220032013011101-2022311332003010-2001212322203221"></a>

<a id="canonical-0032231103011311-3213111233132010-0210230203030033-1032211233211303-3223103210202321-0010201102013013-2320003212120230-1311123032311233"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-1200320221121122-2112010232310000-1131101233111031-3133330123210011-1232103320130233-3323200212321201-2013112001333131-3201313103002121"></a>

<a id="canonical-2002303102011100-1133111123002111-1031020332301003-1300021332133321-1013123333031101-1233212233322303-2103112000313330-1031220323201011"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-2303220002302023-3122130212133123-1231210211133120-2023020312113330-2321130220220003-1111212011302230-3120013130211200-2013020032213023"></a>

<a id="canonical-0201201212322332-3110031012002130-3103203211230103-0003003201320103-1230231112332232-3111013200202103-0002311022023030-1333312113000120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-0233120223311212-1102112210232133-2113123200310012-1211220300312302-0301110132020020-2311223311030020-0221130213323013-1111000302100332"></a>

<a id="canonical-3210210332310213-1313232113300231-1012301201113212-0013010122033100-1132023210131210-3000210331200222-2300003030031333-3033313122011001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2211223000123123-1230313100132320-0220002031002223-3113110320100101-2331020031113113-1232102230013011-3020321010300022-0033213023203223"></a>

<a id="canonical-1013210303003210-2301200003322010-2022302213130023-3012001132012202-1221300223121220-2230313323203310-2002131033231332-2123120033223120"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-1030022230322300-0032132322312010-3122033032233212-3313212202312031-0230101230333102-1132302001001212-0333100210231110-0320230002101212"></a>

<a id="canonical-0320122003203222-0203312012331103-0121023232111221-2000020322320023-1211011313123023-2210101102311123-0320213220132122-1110220330331131"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0113210120220330-0121022112233310-1223021200220331-1131321311233303-3021010133111033-0233321223213330-2103020031023102-3132001023321101"></a>

<a id="canonical-0101301232101211-0101111203022233-0123121000133010-1002022300003302-2323111020131302-3301313322202123-0102123310223113-2031131233230122"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-3221123300333311-2231330130312010-2102013230030222-2130222021203313-0120322332031121-3330202211233003-3011122022210230-1321212122322332"></a>

<a id="canonical-2000011211012110-1132030233212330-0130312102210333-3203232202332323-3030331300201103-2102230121332132-3201301032330310-0310122020233133"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-3303021010121010-3131323221000332-3022203202322200-3310212112031130-2123321202011233-3102030010030030-0122113210223010-2231230012023100"></a>

<a id="canonical-3003122201301322-2221132330020120-2121212232120310-2003021221210110-1230113113111302-0013232101301000-0311012313322023-1021113100200300"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-0022333310221233-2220203313110203-1100012130103232-1121102221230311-3220032213122130-1203112111113000-2211120102020001-0220321033233301"></a>

<a id="canonical-0032201111210133-3132111233310102-0133010202123220-2021001300113010-2313021211132320-3213233311231101-2121322022123332-0130302113132013"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0203130121220031-0221101103233103-0311130202231111-1302320112002302-1322013311221320-3113000200311303-3221232202322323-1302000132112032"></a>

<a id="canonical-0300301022001012-3003111211312112-2101130032331322-0131203101201212-1212302111023121-2333223323001332-2200323032233123-2122110323220100"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-2231110321222212-0201111132112130-0013003032121100-2012323122331020-1223032133221000-1031230012022333-1331123330031200-2032123003310201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3110233122021230-1023230203012113-2213003011302032-2213012301133122-0121032112000233-2323013302231232-0013330013331022-2132131202220023"></a>

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

<a id="canonical-2120303203123020-0110121000133132-1212110310223313-2032020010030120-0021012221002232-0303323313020123-2323333302000232-0222101010201012"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-3231200230011312-0303101230120330-0133022312213233-0133021032231031-0332001020233201-3322010212231202-0233001121213000-0321020201322322"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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

<a id="canonical-3330313200200023-1001120330111313-1100202000301111-3101101000000111-1231221223021002-0210132300321003-0112033022321102-3011223211131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-3032331132111332-3110123301301303-2330021323310103-1033321122112002-3220032321111023-0032100310102303-2233112220131101-2331000020032120"></a>

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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-007.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3201301112233220-2332003210011130-0011021112111211-3002111010223003-1130013110130232-2031333133113222-2033322000313002-2223030233131331"></a>

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

<a id="canonical-1001223012313030-1302132323133100-2302320310013220-3033312113000233-1333333223002322-0002000011033122-2333210102233200-3333012312101213"></a>

<a id="canonical-0302320102210321-2211311203113330-2310020103021131-3200102133300233-0312330001132023-1310013201221322-3022322303313231-0002011123233023"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-3130130300133132-2121122113102301-0310101220022230-0100033212022002-2023020101311120-0231132130033203-1013311133332302-2003112103132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-007.md#canonical-3031113033003222-0302203112110233-3013320030113033-2320330012301030-0323112210133122-1323133221322000-1301022232122020-1023332303223000)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3311001211021322-0330301013021103-1223122111101101-1323032032110211-0023021321301032-2031302200320033-2202330130130201-3300123210030120"></a>

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

<a id="canonical-3130200102313022-1012211213023033-1232022333113120-0320033120300203-0012330030113331-1331132310230211-2022322233100320-3323012232012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-007.md#canonical-2023030202313320-2120223312023223-2011313021311023-1032020310121131-3203020102320011-3102231323232010-2320332011330312-2102320000023322)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-0111012211012203-2321213132300010-2221312320221011-0311121122320331-1330100230031101-2210130232223103-2322032012022323-1323223230332312"></a>

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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-007.md#canonical-0131213121022010-2302331123122311-0333232020222202-2110030111033001-1010012332232032-2230033302023213-0102002320332232-1002130120300231)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-1110103000221223-2121210100122110-3001133220000121-3132023001201301-3303332230200333-1012300020232000-0003203331222013-2330303011020321"></a>

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

<a id="canonical-0222233211103301-2023021131110221-1131132211001103-3233212033320000-0101212211330221-1323233233022303-0200223310321033-0320011122312311"></a>

<a id="canonical-1000112120121113-2100311131221003-3230322012300221-3113032210111330-3331322030102021-2121021022013131-0121002231023033-1323203313012001"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-1210112122202103-1030122112310233-1220330113301232-1202222113210112-2213101331200232-0310022321320230-2303212201012213-2201231320320322"></a>

<a id="canonical-2312222033002021-3330332302323112-3330201002010100-0122102302003312-3301031102033222-2223023312231333-1031102222201031-1023000313221131"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-2011233321321130-3201322122230032-0002032010203313-0313232221033211-2000230233313320-1331303012122032-2010300331001013-3220222100022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-1121012203011210-3333202300113201-0012031100103303-2201301001222123-1013002103022211-1031111023310030-1223311113203020-3123203331330213"></a>

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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303): complete subsection reference.

<a id="canonical-3300311222020333-1030030320231030-3301031313003330-1200212122310110-0210201220111113-3322220310010031-2211000131231301-2011123001002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3313200202322302-2110333300320130-3011011122312030-2231112130020323-3131010212102120-3113332331121231-2201111333021101-0031223123201222"></a>

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

<a id="canonical-1332201330132021-1131002300122102-3331131202122302-0130220220003122-1023031330200110-0330202232100033-2303003211223210-0131331332322023"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-2020010122010222-1000311321012023-3221001220003013-0103300333223203-0101102000112200-0221213002302120-3230110102111132-2012111132301303"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

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

<a id="canonical-1111220130132322-1202231111300102-1211220130203000-3013231033311031-0232000311023022-3010301211121102-0301313300103232-3002211310133303"></a>

<a id="canonical-3333012032322220-1011001100200322-3013020220210300-3232303200331122-3023103202221320-2332310213002032-0333300110303000-2003231022031230"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

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

<a id="canonical-2213203132022002-1121333013201200-2231132103120001-1013001210332122-0303222111333310-3001010100322232-2111230201133022-2100211001302301"></a>

<a id="canonical-2122102003120332-3202022123211201-0222133333130212-0233333310302133-0213220233312122-1111100321123303-3111001333232003-3131103012211211"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

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

<a id="canonical-0101031013110210-3301211201131132-2332013211122123-1203010200202230-0302122311231131-0000302303011203-1032133113001312-1102111020133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0230223311012223-2220223212212231-0010330210301133-3011222311112222-1233110102130100-3102032012020332-1321001020103330-3123321320231032"></a>

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

<a id="canonical-2321030321303331-2100120302001222-1220131301031303-2211122021013110-0302231301230110-2321010223333012-0231332313131012-3203331000120212"></a>

### Direct properties for `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-2011221001201222-2233231031202211-3112103013222202-2322031210313131-1213321103323300-1102031200021013-0112320202331011-1320122021311323"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

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

<a id="canonical-1102312333200312-2033102230013122-2202000103323132-1303210121202300-3201213100103033-1303220103011221-1232220120012331-0113220023103201"></a>

<a id="canonical-2233023323120100-2101210323212112-0301111300313010-0033020330122121-3303031313232201-1132303011013231-3022113011112210-0301013111212020"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

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

<a id="canonical-1232011102210301-0210202231320021-1132322221122133-1020211302130331-1311013222222323-1210322023133020-3201233020030320-2021332113030222"></a>

<a id="canonical-2230223331122012-0023121021222102-0300111211232002-3203231010200303-0101012233300132-3113222212023223-1323200130312311-3000310122011011"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

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

<a id="canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-006.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-007.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-007.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3033121121130101-0301331231112001-2130120032022331-0232112200131030-1001222001133232-0002200031321201-1321023032121001-3202022321321131"></a>

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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1123333213031103-0323222223113223-1300011122313222-0320302331210200-2222213323221312-1010222233323310-1303021110130013-3100032113003123"></a>

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

- [coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130): complete subsection reference.

<a id="canonical-1200233332130313-0222221300300001-1031122032303033-3133132220232302-2321203300200111-3010321003320000-3300322121331320-2321011133010230"></a>

<a id="canonical-1221011221221110-3031323010000302-0212331231231033-0330031010232213-2221020323013301-3031331023030133-2112000120011203-0231122130311202"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

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

- [no_mtls](resources--workload--reference--group-008.md#canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-008.md#canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211): complete subsection reference.

- [pass_through](resources--workload--reference--group-008.md#canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203): complete subsection reference.

<a id="canonical-2332322302112100-1231203022302023-1213302002322110-0233010100300100-2033312132311020-0010111222032010-0312132030203112-3120330023032101"></a>

<a id="canonical-1010300021331132-0320002000202332-0321030322021123-1032123313023032-2302120223023012-3233020203312313-2231112033221111-0311031021110002"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port` property

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

<a id="canonical-1120231103123203-0221112020222020-3310313233233312-3321211303100230-2122300230120122-3233002121032331-0210113330010233-0200221011330311"></a>

<a id="canonical-0101223123333010-3003301333132001-2332223223003212-0132033003301011-0301330230212003-1203100203232311-2220320131010023-0102011213112320"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.port_ranges` property

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

<a id="canonical-2123312131300210-2210311301101010-0211022101331300-1122302100322200-0123323320032200-0122110022023010-2233203302030223-2132103200220311"></a>

<a id="canonical-2221013330032111-2001332022112201-3020020120311311-2302001010033321-1313121200232221-3131113131001102-1203022112023303-1313033200030302"></a>

#### `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.server_name` property

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

- [tls_config](resources--workload--reference--group-008.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310): complete subsection reference.

- [use_mtls](resources--workload--reference--group-008.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230): complete subsection reference.

<a id="canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0111010302130232-2002313003101133-1113322030010102-0030013330233233-2202312032331010-3123213210002210-2320333031223112-3313301103101000"></a>

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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2131031222333210-2010000212032131-2021000311130101-3233132320202020-3013312211312033-1212010323221212-1100102130333112-1001301021221202"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

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

- [http_protocol_enable_v2_only](resources--workload--reference--group-008.md#canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331): complete subsection reference.

<a id="canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-006.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-006.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2213122002000331-1301030100232220-1212321121331102-3232310230103000-3230232122200201-3213121222133102-0102120232231121-2223322210000030"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-006.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
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
