---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-2221301000131000-0310010010313313-1103002201211110-1030322101221103-0331103103033121-2230122113030212-1021102310303012-0312011020020110"></a>

## Direct properties — trusted_ca / 332322112100 / 3

<a id="canonical-2212212023121202-2313023330321321-0231321121110211-3123301023001322-3130220200231021-3212313322310221-3313202210101101-1232031212223231"></a>

<a id="canonical-0322022222301102-0031100211113212-2103313120302220-1012212002131133-1113130112021300-0330130201131020-1011203022030213-3131223100233231"></a>

## name property — trusted_ca / 332322112100 / 4

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

<a id="canonical-2230130301232230-3001011121023102-2302011103023023-2003121212313023-1110102030102030-2003031323222233-0120002222012302-1220033330133113"></a>

<a id="canonical-3212322031313302-3320133022331231-0001331002331013-1301033100021200-2113121002202112-1233031320011022-1020231300110102-3332102130031031"></a>

## namespace property — trusted_ca / 332322112100 / 5

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

<a id="canonical-0130123031211303-0212032231230031-3231110103101211-0001310013220021-3032222201300001-2001313132300123-2200003031320000-3221223311320333"></a>

<a id="canonical-3321311321212132-2300212323221011-3312123222012010-3133101232012232-2220111120122310-0211001002233112-1012211022223321-1332112012321033"></a>

## tenant property — trusted_ca / 332322112100 / 6

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

<a id="canonical-1133010021312300-0313331221200030-0331123333313211-1002031123103201-2331331300300212-3120223123011121-3323022121331102-1331231110121033"></a>

## Next pages — trusted_ca / 332322112100 / 7

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3221120302133200-2013113331332203-0323332322332331-2113030121313132-3201101323233220-1002123231331100-3030233310332120-3033231201103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031023023003312-3213000121013211-3000200101100100-2300031233113012-3003021133213013-0222103103320221-2202332332312330-3313311013223132"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 233031221131 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1301301113111110-2233321110301200-3201113132333301-3123023200120333-2130030330321230-3130101213131301-3320033210001100-2302023131330223"></a>

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
xfcc_disabled = {}
```

<a id="canonical-1032020221032210-0012233320030002-3203023322223130-3233212333311133-2020121202112122-0100112301311101-3100011331131323-2112100123201113"></a>

## Direct properties — xfcc_disabled / 233031221131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322001021300233-0032333203130311-0233310020111302-0101110131132131-2201310223321202-2331211312023130-1232120320100200-2122030103313122"></a>

## Next pages — xfcc_disabled / 233031221131 / 4

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0000312003001302-1122232200100212-2311230011331001-2201013210031300-0230220320310321-3013203311333033-3322330201331111-2213231322023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303331130132023-0322103302211320-0103003011133022-3211031233221103-1232130203330322-3122232123213030-3312013233102202-0120212031100112"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_options — xfcc_options / 323003032202 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-3202033002211233-1301122221120120-1033012321102210-2321013132013033-0311312210012210-2023111231301233-1132102221100102-1112232122033120)
- [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-0121111300120302-2230031133133130-2123131213221130-1002012101111221-2123331230122003-0131120123223023-3320211032330010-1110023113101132)
- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1232011003223131-1231312010203002-0320300222000333-2020333200020212-0013010110113301-0321220123313213-3331330222001311-0121200312220102"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0011310213133232-0232102321333332-1113212201022220-0132122311101200-0331002322303121-2223231202033013-3233233102201033-1303121203110112"></a>

## Direct properties — xfcc_options / 323003032202 / 3

<a id="canonical-1231010103323111-0133022222121131-3031321113032032-2113331320323301-0010101132333022-1312231130002100-1331200300323310-3001332033133021"></a>

<a id="canonical-2330301030210003-2211211133113313-1003211032103330-1131330122310212-2323131122200123-0300223102202210-0232330001112201-0031130222231011"></a>

## xfcc_header_elements property — xfcc_options / 323003032202 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-1100311230203233-2102333201330103-2210100031132131-3032011301013010-0000120012130111-3033032001230331-2110212120221000-3310203010221222"></a>

## Next pages — xfcc_options / 323003032202 / 5

- [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-3212302133121022-3003310220033333-2103011321231131-1323011000210123-1202121231231222-3333333012323210-1130103201331103-3103101011323000)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301332223213023-2323011010011022-2133100013322011-1312230132311011-3212030120332121-0110022011222231-1300013003302110-3033222102321220"></a>

## tls_tcp_auto_cert — tls_tcp_auto_cert / 023300230130 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- tls_tcp_auto_cert

<a id="canonical-1021002212100302-1311013322023113-0033102003303321-1102223200321033-3111102121313232-0020131233022011-3121111200002222-0320203122133203"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting TLS over TCP proxy with automatic certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_tcp_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101223212312112-3123302012003323-1130210300233121-2102333330200113-1112300113323022-3211331311232200-3213113102323012-1030302210133232"></a>

## Direct properties — tls_tcp_auto_cert / 023300230130 / 3

- [no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033): complete subsection reference.

- [tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301): complete subsection reference.

- [use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330): complete subsection reference.

<a id="canonical-1321123211212302-1101032310310211-3101112322001033-1123201333202323-2133330300023100-3220111000102221-3000211201032021-0010011303310310"></a>

## Next pages — tls_tcp_auto_cert / 023300230130 / 4

- [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2332303230031210-0302010003210303-0310320200311033-1033332033110010-0002031132312001-3011301111002122-3230213021313202-1110210010312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211221313003030-1310113321111133-0101110133302232-0311122231312301-3321013203331210-1100230022131132-1231231121013323-0320121221300211"></a>

## tls_tcp_auto_cert.no_mtls — no_mtls / 121003203210 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-2313120012003131-0311100200113110-0123033033123200-3003200132300220-1021000011003221-3200000101330201-2110012131301202-3132133220230013"></a>

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
no_mtls = {}
```

<a id="canonical-2100322213013232-0330002311302203-2023333223203233-3301230122110303-1320232221322130-2120302332322303-0302130202012322-1222103202231320"></a>

## Direct properties — no_mtls / 121003203210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100131311302311-2201233312030032-1100320202202312-0032123311110231-1110320131302231-2012313301002002-0111123233021303-0323113331300030"></a>

## Next pages — no_mtls / 121003203210 / 4

- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322031011113223-1021330102311222-2013112020130210-3033133323321302-3000230321102233-1232103112233223-2321330112333111-0233032210322131"></a>

## tls_tcp_auto_cert.tls_config — tls_config / 033230200311 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.tls_config

<a id="canonical-3321310201213312-1102103322302112-2213212202200032-3302231101100330-0213221333111211-1023311110133123-0020313120313112-1300130202230102"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0132303233012112-2200113013003021-3213200303010101-3301333210232131-1100001312020100-2222030000021202-0231200303313211-1211013231133311"></a>

## Direct properties — tls_config / 033230200311 / 3

- [custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2313013020003011-3223231210320332-2110311102211120-3033132323001101-0011210322131033-0132220020103033-0003120020233123-2220023113001312): complete subsection reference.

- [default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220): complete subsection reference.

- [low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2201112121202132-0000011113332030-3232020202200320-2131120132123300-0331302111333333-2313220001121022-3111212221132303-1112311301131322): complete subsection reference.

- [medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3000323120232101-0221013010300110-2213311012323201-0213132003100230-2121333300332311-1101023310121302-3121012021212101-3330311103103320): complete subsection reference.

<a id="canonical-1023321133201313-0213311013300220-0032103322130301-2012122332023113-3313103120033322-3001221230003101-2312202101023322-2301200233113221"></a>

## Next pages — tls_config / 033230200311 / 4

- [tls_tcp_auto_cert.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2313013020003011-3223231210320332-2110311102211120-3033132323001101-0011210322131033-0132220020103033-0003120020233123-2220023113001312)
- [tls_tcp_auto_cert.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220)
- [tls_tcp_auto_cert.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-2201112121202132-0000011113332030-3232020202200320-2131120132123300-0331302111333333-2313220001121022-3111212221132303-1112311301131322)
- [tls_tcp_auto_cert.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-3000323120232101-0221013010300110-2213311012323201-0213132003100230-2121333300332311-1101023310121302-3121012021212101-3330311103103320)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2313013020003011-3223231210320332-2110311102211120-3033132323001101-0011210322131033-0132220020103033-0003120020233123-2220023113001312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030232132203002-0001202221202120-1300000010223123-2311112301121220-1300330323022332-1232213030231230-0302030301303213-0331202333222320"></a>

## tls_tcp_auto_cert.tls_config.custom_security — custom_security / 302312302203 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.custom_security

<a id="canonical-3122320210213120-3132310300203321-0023203133132221-2012132122221032-1132021010222230-3110333322211003-1200110210233200-3202113110210333"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1200111332222101-0133033130231233-3000312010123303-3102022103131022-1201302121312102-3222203003032301-2010233102021002-1013133131022013"></a>

## Direct properties — custom_security / 302312302203 / 3

<a id="canonical-3110230210310221-1011323300032322-0113000130312010-2022231123113023-1022302023122221-3111300133033221-3101321133322100-3233221210211200"></a>

<a id="canonical-0113001330012013-1301121111210002-3013130131310121-2230210222222310-1330032311200131-1010122210232001-3301023021133233-2230300322221110"></a>

## cipher_suites property — custom_security / 302312302203 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3302230203202020-1020012130110030-2020113300330120-3012211330212032-2212103232133231-1230220210132333-3301023232002323-1023002132323012"></a>

<a id="canonical-2101003132321113-1201232310332102-3020220223311021-2310212200230201-0122322001131302-3122100220020101-2102122311132210-1221001030221110"></a>

## max_version property — custom_security / 302312302203 / 5

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

<a id="canonical-2331012331102330-2222311331122201-3102103312030101-0320011333220032-0103132130020020-1210103003231030-0300130312113113-3122112322122300"></a>

<a id="canonical-0113322012122210-2320313222131022-2021233103300302-1312031010210203-2231131130330103-2001112203130030-3310232002322311-1320310012103133"></a>

## min_version property — custom_security / 302312302203 / 6

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

<a id="canonical-3203300310330201-2312023011023302-3222322022002311-1310001111110302-1021131123233121-3120013211132321-1110232030111220-0020113313022021"></a>

## Next pages — custom_security / 302312302203 / 7

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1200213310211002-1313301303222331-2312132112130122-1300233203000230-3222000322331122-1023012300132233-2011131302132312-2322000302023220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121303332211210-2223011211231213-3101211131111111-3033121302321201-0212100222322322-1001320303110320-0303200120200231-1201130200030322"></a>

## tls_tcp_auto_cert.tls_config.default_security — default_security / 211232203330 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.default_security

<a id="canonical-2010010302012332-2230323211012202-0312230201203102-3300001130302322-1330311131021002-0113302200111113-2122213232321110-2211120031233012"></a>

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
default_security = {}
```

<a id="canonical-2032131120220122-2331013013322123-0002012032301313-3313112202131302-0321332103101102-1003102122133121-0311033111123032-3130300100222122"></a>

## Direct properties — default_security / 211232203330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210023121023003-2322112100000122-2220303320203011-1022331130033201-2131301330033102-2000000000202312-3233221031202032-2031020022300003"></a>

## Next pages — default_security / 211232203330 / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2201112121202132-0000011113332030-3232020202200320-2131120132123300-0331302111333333-2313220001121022-3111212221132303-1112311301131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103220222333121-1022333201223210-1113023001301102-3213121202031301-2011013003300100-2321102123213013-1022002310102211-0030312222220103"></a>

## tls_tcp_auto_cert.tls_config.low_security — low_security / 332113213231 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.low_security

<a id="canonical-0012331301303113-0122013031120210-3111133330002022-1223232130010220-0213102121003312-3333303311320231-2330320030021323-1132012003211232"></a>

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
low_security = {}
```

<a id="canonical-0022133311332221-2113030022202023-1200231313211332-0220201122211231-2131032030212112-3330321312213013-0333032221321222-2033120033021021"></a>

## Direct properties — low_security / 332113213231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000030210220011-2312023110331321-0122110131103213-2321223022202102-1312100322020210-3320333200101310-0130013122112202-3032213030010001"></a>

## Next pages — low_security / 332113213231 / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3000323120232101-0221013010300110-2213311012323201-0213132003100230-2121333300332311-1101023310121302-3121012021212101-3330311103103320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313232130011311-0300330232310331-0013201331330221-1202330301302320-3012030231011301-2023010331332110-3111332202000311-1130331121210221"></a>

## tls_tcp_auto_cert.tls_config.medium_security — medium_security / 121130023333 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- tls_tcp_auto_cert.tls_config.medium_security

<a id="canonical-2203331123101331-3200220000230021-0112033122002013-0333232221023131-1003012123322012-1001322221220201-2000000113100103-2111033323100312"></a>

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
medium_security = {}
```

<a id="canonical-1120013221332322-3131103011331111-0312011011032033-2233220213010120-3121113223230302-0301331030211013-0233113001333120-1303212212220011"></a>

## Direct properties — medium_security / 121130023333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201010301032002-1022030231121213-0120200232222223-1020310112332030-2221311211232022-1002023002113020-2030020213302122-3003032022233302"></a>

## Next pages — medium_security / 121130023333 / 4

- [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-2022320330100222-1002022331311031-2300333321023112-2213021113122210-3001322002201333-3021221123102103-0203321133220210-0203000303130301)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011033001030203-0123300321112112-2203333031332221-2002331230032310-2002131100101113-1330103113100010-1001313310222210-2230333032213002"></a>

## tls_tcp_auto_cert.use_mtls — use_mtls / 112110003210 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- tls_tcp_auto_cert.use_mtls

<a id="canonical-2122123310312001-1332302032133120-0112220022001113-0332021122030022-0331230110000112-3330102331101112-1023032300031022-0133033333131032"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1020221220303300-3321131202233221-1313021322320002-0230322013013130-0322131302303101-2323321323312033-2323202232102300-1113310313011012"></a>

## Direct properties — use_mtls / 112110003210 / 3

<a id="canonical-1021332202221031-0011322033110003-0000322201110132-2312203201221221-1310102321101111-2323221212331321-3000212001013023-3303203103330332"></a>

<a id="canonical-1321323000200233-3121120001322311-0002012231332320-1001211313312030-3311013012312113-3302122333123031-2021231113112321-0120001120332132"></a>

## client_certificate_optional property — use_mtls / 112110003210 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

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

- [crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-1211131111020302-1000023331233231-2300202002002211-2312321110020022-2222322300110003-2030101231311323-2120313123031031-3123103313130203): complete subsection reference.

- [no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0221331323323322-2233111213111121-2103120211102130-2112133132101331-2100213203012113-1213303221011011-2112132230030001-2022231220113302): complete subsection reference.

- [trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-1313322021032031-2133113312113013-0313333312300101-0000313030010021-2101100033003031-2103231100121022-1311023121103301-0312120231221130): complete subsection reference.

<a id="canonical-1301233321023301-3221301123322013-3120111200100111-3213131121132220-1322131002130010-1021133020131331-2222210230130311-1313302230033120"></a>

<a id="canonical-1302332331322110-3001000310212121-1131111323203110-0320200003230130-1220310230212120-0201203121331213-2023020133110031-2300010333321130"></a>

## trusted_ca_url property — use_mtls / 112110003210 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3332232033013000-2032222110022033-0001300033312133-0211031021233110-0110331002111312-1311101022331130-3223133323323212-3031212323210310): complete subsection reference.

- [xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300): complete subsection reference.

<a id="canonical-1131013130033133-2110133100221203-1212202230223003-1302003200223212-3313221213201131-1101032033023333-0130021203003030-3232000101322232"></a>

## Next pages — use_mtls / 112110003210 / 6

- [tls_tcp_auto_cert.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-1211131111020302-1000023331233231-2300202002002211-2312321110020022-2222322300110003-2030101231311323-2120313123031031-3123103313130203)
- [tls_tcp_auto_cert.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-0221331323323322-2233111213111121-2103120211102130-2112133132101331-2100213203012113-1213303221011011-2112132230030001-2022231220113302)
- [tls_tcp_auto_cert.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-1313322021032031-2133113312113013-0313333312300101-0000313030010021-2101100033003031-2103231100121022-1311023121103301-0312120231221130)
- [tls_tcp_auto_cert.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-3332232033013000-2032222110022033-0001300033312133-0211031021233110-0110331002111312-1311101022331130-3223133323323212-3031212323210310)
- [tls_tcp_auto_cert.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1211131111020302-1000023331233231-2300202002002211-2312321110020022-2222322300110003-2030101231311323-2120313123031031-3123103313130203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022030212303100-2333232200111031-2131311232023312-0301021330031221-3030010133111021-0023323122020221-2102310222033120-1102131213301300"></a>

## tls_tcp_auto_cert.use_mtls.crl — crl / 233231023320 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-1020103222230313-0322312213133322-0121301323231031-0302030102111222-0300002203120120-0010103310112202-3130101220001322-3000220213230313"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013320201303231-1232322313001231-0021101100113011-3001200012100030-1220121212221112-1333223121031121-1010312022302112-2211013103131110"></a>

## Direct properties — crl / 233231023320 / 3

<a id="canonical-2022111101313010-0313322233121013-1331001202300120-1223110330130112-1030131213103012-2132110103021201-3013121013320201-0010010332220303"></a>

<a id="canonical-1000000121011311-0002133320130131-3021031132122332-2331322113220001-1303213112131320-1112312220222333-2133033223223320-3223031122031121"></a>

## name property — crl / 233231023320 / 4

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

<a id="canonical-2123213131033113-0101303223000211-1112011200112121-3110332222321100-3213022001232231-1300122121202230-0012311230210033-3333321322223111"></a>

<a id="canonical-1121311030010031-2210103101132202-1133212101330300-2113112232123131-0120223101210321-2123030003102222-2021221310331130-3110302201232120"></a>

## namespace property — crl / 233231023320 / 5

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

<a id="canonical-3023300012323011-3112212113020112-3203202020332121-3202100302232031-3000323012331313-1211002031231033-1323023311223000-3001221322120300"></a>

<a id="canonical-1003312313212000-0100001320221001-0322003202231333-3020101201303130-2233232130101122-1112101303213120-3121103313212332-3323003112002313"></a>

## tenant property — crl / 233231023320 / 6

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

<a id="canonical-0121000011302003-0223232031223333-2322003133310113-1330310003202220-3322310333201331-2312303221213122-0302121030022101-2303113200220303"></a>

## Next pages — crl / 233231023320 / 7

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-0221331323323322-2233111213111121-2103120211102130-2112133132101331-2100213203012113-1213303221011011-2112132230030001-2022231220113302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233201003332221-2312100132000301-2303320123132003-1111100322221111-3312003201310131-0132133131221330-3033030221112123-0203230023231200"></a>

## tls_tcp_auto_cert.use_mtls.no_crl — no_crl / 231330020310 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-1331203211100310-2320012102132103-0030010000223320-3300333111012011-2010310131300000-1010323123033112-1010210100203101-2310320223223302"></a>

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
no_crl = {}
```

<a id="canonical-2222102203032211-3101232121023211-3000133032333020-0130311301033120-0200012210030101-3220310022100321-1103301110111003-1331003300010310"></a>

## Direct properties — no_crl / 231330020310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012211123333312-2011323001212001-1301303201312133-0112220221313213-2302122303301320-0311110333121000-0320032120310010-1221012332022133"></a>

## Next pages — no_crl / 231330020310 / 4

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1313322021032031-2133113312113013-0313333312300101-0000313030010021-2101100033003031-2103231100121022-1311023121103301-0312120231221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331230310010010-1022010332200010-3233301010013112-3230123031313013-1200032013021011-3001100330221212-0131103322110102-1320131123110133"></a>

## tls_tcp_auto_cert.use_mtls.trusted_ca — trusted_ca / 331103020110 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-0310300202211220-2130103020100302-0102210122321330-2030310221102210-3211111311311132-3303103211003233-1323202302101323-1103012012131011"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122010000201231-2123022320210111-2201331213103123-0321002231211102-1330221211122201-2332202023321302-2112103213310310-2002023310121111"></a>

## Direct properties — trusted_ca / 331103020110 / 3

<a id="canonical-0101112133310323-1100032200311133-2011111120122012-0120033102303020-1103302330232031-0323101331133123-0210322302301331-0113113013302222"></a>

<a id="canonical-0103310301311323-1312232123120103-3102120301203303-3003232112330000-1000111113131032-0302010332330100-1133012000313220-2330013131221200"></a>

## name property — trusted_ca / 331103020110 / 4

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

<a id="canonical-3100333102033203-1230122020112301-3010023213303000-1120310003030230-3131012112101221-3333023301131013-0101220110332130-0130312003203320"></a>

<a id="canonical-3333212033203121-0031123222331001-3101211213033113-3302222233020102-3021030223221032-3011223023132213-1310300000032303-0033101222101110"></a>

## namespace property — trusted_ca / 331103020110 / 5

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

<a id="canonical-3132002233303301-3102011330212212-1201132302202201-0213033231101332-0111010120002031-0330003031010021-3213110011101303-1011031101020301"></a>

<a id="canonical-0010330111202223-3203222113332223-3013133132302223-0002002013223212-2012210330310233-2211101130202110-1313133311231212-3301200100322212"></a>

## tenant property — trusted_ca / 331103020110 / 6

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

<a id="canonical-2133311101100311-3121003213212131-0301223111121133-3321000013102003-0310201310203033-1322000020002123-1033032300232232-2123130221231033"></a>

## Next pages — trusted_ca / 331103020110 / 7

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-3332232033013000-2032222110022033-0001300033312133-0211031021233110-0110331002111312-1311101022331130-3223133323323212-3031212323210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013230331121023-0030310320101202-0003302330321232-2211013202111121-1232023023131131-1100112210002213-2202310322320210-2232220220220122"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 201100300112 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0013103230120133-2032003020333120-1113103020030311-3300020222213333-1202230333110320-1210013302101223-1323322211113010-3321320103021301"></a>

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
xfcc_disabled = {}
```

<a id="canonical-3111311033031110-3212313333300232-2231122233233031-1323030330201120-3302312131103203-3302120333002101-2021032232130123-3103313330000213"></a>

## Direct properties — xfcc_disabled / 201100300112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122022001100132-1330301101332210-0202102332113033-3003232101212320-3011231122303333-1031210330101223-2321221122303111-2122020103202301"></a>

## Next pages — xfcc_disabled / 201100300112 / 4

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)

<a id="canonical-1130133122020122-1203201011332202-0113220130221321-3031122013001103-1211212012303103-2210110030111113-0133120302322110-1110212200100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202011012321113-2231330323231232-3131311021233020-0233012302320002-0112030002002303-2302211012131010-3112022310131011-1313132232231101"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_options — xfcc_options / 021101331101 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-2012213313231232-3113222123311030-1312023102102220-1333030303022022-3132112211312313-0310133313220331-3313212213313330-2132122131111001)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-003.md#canonical-1020000333221312-0212323132120330-0112301302023312-1010022111201312-0212120203321302-3011031322120320-0022031132030332-3001230220212331)
- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- tls_tcp_auto_cert.use_mtls.xfcc_options

<a id="canonical-2213032130123000-1133101312211231-0130100322231020-3133332321101210-2202113113101213-2133221323233300-3123203031201133-1322320303110102"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2011013313300032-1120132311220121-1111102112013232-3313333313122131-0201332130000130-2032301233011300-1032111133023221-3232013100322001"></a>

## Direct properties — xfcc_options / 021101331101 / 3

<a id="canonical-1122302311011201-3233002131113212-0111022023223311-3002001100333310-2132111310012313-0201110102321230-2332103233021031-1231211232113132"></a>

<a id="canonical-1311012032301200-0211033013110000-1213120020023230-1010111331111120-0133210231212201-3213210110102302-3122111233212031-0020320212213001"></a>

## xfcc_header_elements property — xfcc_options / 021101331101 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-0211122110331331-2131313123000023-2213001131313333-0122203101223003-1231011022132013-1233332312010220-2221332000220110-0331332210203320"></a>

## Next pages — xfcc_options / 021101331101 / 5

- [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-2122133202102320-2301201232130231-3003320030220113-3103213110221101-3011102202003030-0301303131230113-2232330000320302-3111231012201330)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
