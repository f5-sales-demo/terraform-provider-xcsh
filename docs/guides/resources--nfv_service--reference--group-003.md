---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-2310030231113303-2033101102332013-1331110233332322-0310323000231011-3303212131313201-2003113011031122-1130002302333131-1132332122011112"></a>

## `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.max_version` property

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

<a id="canonical-1030303032300000-2113132130001330-2110210002203021-0102132031212110-2222130101203120-3032320030232212-3301331330010330-2200332131330003"></a>

<a id="canonical-0301131203322012-2223212120223320-1021200330010102-3001031102122000-2122111030120000-0300002113323023-0321331130100010-0101022200333103"></a>

## `https_management.advertise_on_slo_internet_vip.tls_config.custom_security.min_version` property

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

<a id="canonical-1110320013323011-1221022301022012-1322313110032031-0113212112202223-2232123020330103-2103233202220200-1033132313332211-0020110123232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-1213132313200213-3123233220332100-3110233321120302-3302330011200013-1020000021301021-2311231102233101-3132011331023103-0020311123001231)
- https_management.advertise_on_slo_internet_vip.tls_config.default_security

<a id="canonical-2111021302020010-1212233132030233-3033202030220223-1123310233232202-2133023300010023-0233111032020020-3212130022102102-0220103103013121"></a>

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

<a id="canonical-1013003000033020-1113232000210310-2321203001320001-1221311323300021-2120320323231233-0022203313121011-1013222232133321-1122203123230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-1213132313200213-3123233220332100-3110233321120302-3302330011200013-1020000021301021-2311231102233101-3132011331023103-0020311123001231)
- https_management.advertise_on_slo_internet_vip.tls_config.low_security

<a id="canonical-1233103000320302-0301003332100021-2000232131123320-3333132020020002-3221202311310333-3002301221010220-0303103131123132-0311321102000223"></a>

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

<a id="canonical-3120302233232220-1010103303200001-2103100232033012-3213231322030023-1321011220230003-2011220221130113-0101211112103321-2032030013113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-1213132313200213-3123233220332100-3110233321120302-3302330011200013-1020000021301021-2311231102233101-3132011331023103-0020311123001231)
- https_management.advertise_on_slo_internet_vip.tls_config.medium_security

<a id="canonical-0232022330032012-1320012332332331-3202210121333312-3230011333201001-1303210011231011-3331333033113001-3021033312312003-2031223103221202"></a>

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

<a id="canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- https_management.advertise_on_slo_internet_vip.use_mtls

<a id="canonical-0203200203212321-2132002301201203-0313000213332021-2310211333332113-3201321010032332-2111233132100032-0030012210220233-2313210003013012"></a>

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

<a id="canonical-0130222310012303-1101022032222120-2123303331103223-2201332022232210-3012313000013122-3001111031332322-3211123103331002-0313022010221332"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls`

<a id="canonical-2211002013323210-0013323113233123-1300332312231102-2132323331213003-2303223101100311-0103321113231302-1221123101111300-3112331321012032"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.client_certificate_optional` property

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

- [crl](resources--nfv_service--reference--group-003.md#canonical-2312311003223132-2222120000020203-2321301121120103-2231120101321310-2231233000333100-3011031112113230-2020022222311131-3121301202311230): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-3003031320320332-0321333001323110-1022203213110311-1323101231313031-3013220113030013-2111132110031203-0102112031312132-0230332133312303): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-0300010231301102-1300013232001021-3130003100132310-0320033103133301-2103113201012311-0100313003232233-0300003132301301-3022202303302030): complete subsection reference.

<a id="canonical-1210221023303021-1333310000102202-0012003221023203-0302201332013210-1021223300210313-1100311023103210-0020112302210330-1121122220201302"></a>

<a id="canonical-3333002323113022-3302302213231313-2310230012233130-2110300222103202-1222023020031233-3202322112332200-1132300333213122-2123301000231030"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-3110102211112330-1123023120010002-1033312223202210-0123212101233111-2300030310312033-2030132333310031-2100100100010222-0213013313320231): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-2212213112313311-3122110131300103-3021131321320332-2131122231132020-2111013001221220-1310212122100202-1200033003222330-1203302030311303): complete subsection reference.

<a id="canonical-2312311003223132-2222120000020203-2321301121120103-2231120101321310-2231233000333100-3011031112113230-2020022222311131-3121301202311230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.crl

<a id="canonical-3030101113312111-2003312332110222-0313222332010033-3202333010003123-2003132310023121-1012121032302320-0101001103221112-3322013030332022"></a>

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

<a id="canonical-2230212102013321-1231202322322230-2220231023311222-0223310033032133-1210333322222110-3213010313222212-1133011032013023-1313223311122313"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.crl`

<a id="canonical-0030133212103232-2211333103032311-1321102002312000-3313312033130220-0301202023223203-1020010233111130-0220032201111011-0131101211222331"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1033110012030301-1330212313012300-0202123011030022-2123212320200231-1011303322121003-1331323233301303-1012000320103321-0233031212331202"></a>

<a id="canonical-3320111022013321-1003021300303302-0123012003330231-0200302110212333-0122111011212120-2122323102211010-2332132021000013-0230232122123002"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1102002020022313-3002213222313030-2122200132033223-2020322321000131-2302231201313313-1221233032013320-3103322231322132-2110312132131301"></a>

<a id="canonical-1212000312212332-2020331112012122-2000030100132303-2211033131332231-1032112303321332-0023132000002302-2031102130223120-3130220301132130"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.crl.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3003031320320332-0321333001323110-1022203213110311-1323101231313031-3013220113030013-2111132110031203-0102112031312132-0230332133312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.no_crl

<a id="canonical-0131033311220010-0111221133010001-3013200000202122-2103123033323133-3200032303132331-3102220030203012-3020231213211231-2111122120210021"></a>

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

<a id="canonical-0300010231301102-1300013232001021-3130003100132310-0320033103133301-2103113201012311-0100313003232233-0300003132301301-3022202303302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca

<a id="canonical-3212200110222312-2321131210222313-0321011000130032-0132302100212101-2332033032103311-1311011031113112-1220032102033323-2130201311123212"></a>

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

<a id="canonical-0321011223101321-2100001313012030-0012033131113031-1332121212333001-1103222220022210-0130213113320031-2003122032031210-3201231102130201"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca`

<a id="canonical-3212120022223222-3323031213022133-2330220203100002-0332121310012230-3122133121100003-3003222331313311-2331321300310332-1132031122320200"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2030202333213210-1003313111010211-3322033331322103-1330321133302020-1101112213111001-0113001121213312-1211100212211121-2131030200121321"></a>

<a id="canonical-1200223222023002-2003101032330030-3020013301302300-3210313121303223-0132032111231012-0301210100320200-1232222203221103-1320023332220033"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3030112230010031-2102221330322332-1331103202310132-1202210200111200-3112222202203120-3310100031220003-1331221212102212-1200331033122301"></a>

<a id="canonical-2313220133110301-2113133232033203-1200110310302112-0013310213033310-2312331022311030-3330012101323330-0111132302001302-2032300332000223"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3110102211112330-1123023120010002-1033312223202210-0123212101233111-2300030310312033-2030132333310031-2100100100010222-0213013313320231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled

<a id="canonical-3132303122120001-2322032211030112-2133320303120223-1113203331002120-2033212003130223-2210110300121130-0330022321112210-2122012223331113"></a>

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

<a id="canonical-2212213112313311-3122110131300103-3021131321320332-2131122231132020-2111013001221220-1310212122100202-1200033003222330-1203302030311303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-3031132200011022-1313221310121032-3200012331302121-0120313212301311-1101211132233033-2113333100000023-0121110203223213-0303202010232130)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-003.md#canonical-2220303111010322-0112002023201232-0133012112231002-3103111320133100-1120331002010022-2320111201201332-2112303320000210-2130023033333112)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options

<a id="canonical-3212211231012230-0230223010133130-2030210000303021-2300303312011203-0132123010231213-0031113032111122-1013031112222301-0210023022120133"></a>

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

<a id="canonical-1233021033312020-1001311131303012-3231001032213123-2301110200222333-2202220032012000-3003331012222020-3112202323203213-2022013211202020"></a>

### Direct properties for `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options`

<a id="canonical-2110300210103203-1200312202131301-1030000130232011-1130002233201310-2210033223021203-1313032132011213-3003010333022333-1032102101221011"></a>

#### `https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- https_management.advertise_on_slo_sli

<a id="canonical-1222310230103003-2030221101312331-3000330233111031-2200303020312322-0110001110321021-2333201123200123-1332223311201110-2330030101210230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for advertise on slo sli.

Additional upstream details:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
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
advertise_on_slo_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122030120301322-0000133033222003-0303220131332212-2320322331233211-3133331012312320-3022221313110130-1003303031111033-2101123003230020"></a>

### Direct properties for `https_management.advertise_on_slo_sli`

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-1130313331121202-0111232220003011-0103033303320223-2120021203213211-0121112223201032-3110331102023222-3212212010031302-0310211322013020): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310): complete subsection reference.

<a id="canonical-1130313331121202-0111232220003011-0103033303320223-2120021203213211-0121112223201032-3110331102023222-3212212010031302-0310211322013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.no_mtls

<a id="canonical-0322033010311130-0022320003121201-3113011020223313-3232231003321320-0300230030123202-2203122112030232-0323133000022110-0022203221323023"></a>

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

<a id="canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.tls_certificates

<a id="canonical-2222132112223302-0332312122221233-0320110121020203-1031001022030002-3123021201112213-3020300223003103-0003103121332201-1223202131101201"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1223033330121120-0012332331212110-1110312011122322-2032320301002103-1321233013323212-2132113312333321-2110123113120003-1100310120102223"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates`

<a id="canonical-0110213100123121-1000322120212120-0321033020211011-2203113110023200-3003321002200333-3012330102212333-1000121101130321-0330221303001100"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-2303321112330003-2223311011030233-2332321211031112-1201311231223013-3222201133030120-3230230230032332-1233230330110210-3213012313331132): complete subsection reference.

<a id="canonical-1010202211022311-0211131011331123-1002311102113312-2001110233223000-0201022233222112-3310013103000230-1333211321312202-0110111232332013"></a>

<a id="canonical-2130113221003031-2333212331202131-3031322300001011-3113121223231313-0102100130021101-3102020202312310-1320032023122300-3102203202301211"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1230130333103303-1031122031320031-0120200122220002-0022303132211113-2032232123131120-3302222223003330-3112110311213311-2301321032011110): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-003.md#canonical-1032013113221221-2330201313131201-2012232102321120-2021021033220122-0023000111021213-0131113010312113-1231110201003300-1333012213221332): complete subsection reference.

<a id="canonical-2303321112330003-2223311011030233-2332321211031112-1201311231223013-3222201133030120-3230230230032332-1233230330110210-3213012313331132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms

<a id="canonical-0023013102032200-0233131301222211-1200232321210020-1320032131133020-2133332103012210-1200220200200212-0122333030002123-2030200312110010"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122333100330212-0123020031233202-1223013313101103-1133001221223121-2100221132212210-2000121330102201-1003202212312211-1211120202020031"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms`

<a id="canonical-1003032001010322-2110330013122202-2113312222103000-1321220133331130-3013311000212331-1020122231221201-0113000212023313-3131301132321233"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-1230130333103303-1031122031320031-0120200122220002-0022303132211113-2032232123131120-3302222223003330-3112110311213311-2301321032011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.disable_ocsp_stapling

<a id="canonical-3011233100130321-3013222003232231-0230100121203322-1110102211221320-0033322213102203-2333210010333202-1012231213200221-3121103121133003"></a>

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

<a id="canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.private_key

<a id="canonical-0211221232320211-2013103032113323-3103010030102020-2220102022233032-3200110030000211-3113311111231130-1203213300323010-1202312333113123"></a>

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

<a id="canonical-2130132110221020-2020321230011111-0223120131233102-3221301000303210-1032300333010013-3013123020121312-3212110300112122-1101001113031311"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key`

- [blindfold_secret_info](resources--nfv_service--reference--group-003.md#canonical-2100111301202103-0330110312010320-3330323203130132-1012211113222112-2131203211031121-0102023122232121-0210211010102122-2110133113203230): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-003.md#canonical-0231010001133320-3230032322031013-1321332103323213-0120120300231233-1221100220101331-2011202010330002-1220120223303312-0132300132110233): complete subsection reference.

<a id="canonical-2100111301202103-0330110312010320-3330323203130132-1012211113222112-2131203211031121-0102023122232121-0210211010102122-2110133113203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0122301330013003-3232321333221120-1303310131300323-0233113032301133-0013112010302101-3222110012001311-0333020021210111-3220003330122300"></a>

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

<a id="canonical-1233131221332220-3201120231123221-2230130121201001-2123212233321321-2302202131103101-0101330310313222-2101023211320233-1130001300201331"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-3222003112330000-1031331012330322-3122120131022013-0323322032120322-1102222200231101-0033123130322010-1100332322322301-1202200320002020"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2232312211233331-3100112233111212-0111103012123122-3332012230001002-2020001032202102-0010122320002030-1001301223102123-1322233211313210"></a>

<a id="canonical-2321210313210230-0002312120113313-3102313222023131-2331010000131133-2231230212111313-0303021203230323-3023331222131100-3102133101113202"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-1023100103112010-0003032120211000-0113100122223320-1222121132113200-0103221232233022-0110322100211102-0113022000131302-2023212010112312"></a>

<a id="canonical-0230020320132222-0200332232102322-0220102013033101-2133110103213212-1333030221211322-1212212131123112-3031333133113033-0021201312002320"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0231010001133320-3230032322031013-1321332103323213-0120120300231233-1221100220101331-2011202010330002-1220120223303312-0132300132110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- [https_management.advertise_on_slo_sli.tls_certificates.private_key](resources--nfv_service--reference--group-003.md#canonical-2220121131203232-3202100013123232-2221330330202323-3131100122221013-3221230232121021-3133003111020322-0103003210222311-3230101011023000)
- https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info

<a id="canonical-0223110131031313-3121233223320320-1310233230301303-2212301110220013-2120201122333033-2322323210022020-2220100202302223-1122200130231300"></a>

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

<a id="canonical-0002012001111003-0232233023303310-2033331131212120-1132022130332020-1323220323101322-0023002123210133-2103230333233011-2212221022113313"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3132220103031220-3020022000132030-2212330212120103-1321113013023323-1332002320122100-0201302300201110-1310130301323231-0123101300033121"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0212022103102203-1321011132010213-3201211200330313-3223303100001031-1303001322303232-1012303303002020-3321321202031131-2111113002033333"></a>

<a id="canonical-2100202222322010-2233000231130310-2033012330131031-0310133111313003-1330021232023211-2231103201121322-1013131131310221-2311000202221312"></a>

#### `https_management.advertise_on_slo_sli.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-1032013113221221-2330201313131201-2012232102321120-2021021033220122-0023000111021213-0131113010312113-1231110201003300-1333012213221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-2111112012021311-1120230232310333-2311202212021032-0010013131233301-2323132231213030-2330330333013322-2222003003003212-0023112031021113)
- https_management.advertise_on_slo_sli.tls_certificates.use_system_defaults

<a id="canonical-2200301201111133-2313022022013102-3113110020000000-2131032112121023-2222223210213302-1311210202322013-1130322332032001-0033303321223130"></a>

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

<a id="canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.tls_config

<a id="canonical-3031201110103220-2332033131020233-3100133000113300-0220032103110233-3333333321111132-0201113321230023-0232030123131331-3113333120110222"></a>

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

<a id="canonical-3000310100021320-2120233003033322-0100312100212320-2232000223300113-0332202233111013-3020230000013121-1132202303220010-1233120023030021"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_config`

- [custom_security](resources--nfv_service--reference--group-003.md#canonical-3032330203113133-0300213012202320-2003232020002100-0203330112213222-2201131230311213-1021301332011022-3202111330233322-1003031130032302): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-003.md#canonical-1002021221233202-3130021310012220-1310110003230101-3003323320213322-3310110222322330-3130121121120100-1201313002332131-3131200121123002): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-003.md#canonical-1313030333021133-1013033023021223-0033130302313320-2212101202312033-0311010232220003-0113012233000221-2302032020012303-2033302131321302): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-003.md#canonical-2112302102010123-1033302111322023-1121001202132030-0321211301021213-1302021203010030-2111330211332132-2210102210033100-1011303231030233): complete subsection reference.

<a id="canonical-3032330203113133-0300213012202320-2003232020002100-0203330112213222-2201131230311213-1021301332011022-3202111330233322-1003031130032302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.custom_security

<a id="canonical-3000120202122123-3002200103132220-3221200331221031-3213122123223111-0201231213103211-2223130310123132-2002210210302001-2300033213311313"></a>

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

<a id="canonical-2110223301201231-3220110210313212-0210011300023213-0100200113120023-2021033311223320-0130000130300303-2210302310202332-3023023311220230"></a>

### Direct properties for `https_management.advertise_on_slo_sli.tls_config.custom_security`

<a id="canonical-1302021011020110-0121132011323321-0203323033133333-0102003022301201-2231021221111012-1030210210300110-3121213131013310-0320332220322223"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2023000203003103-0230132300331311-3001013112300212-3110031132030132-2223123101301031-3221220223300210-1210013303311131-3131121331103132"></a>

<a id="canonical-1103122000331201-2333100323230000-2312033220121103-1220233122131123-1312202301202132-2220330233003013-2100301102112000-0300332311331012"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.max_version` property

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

<a id="canonical-2331130001230231-0001212313112222-0302001110322100-1200121003210133-1313133012113233-0303313010133222-3122221313111200-3000320023131321"></a>

<a id="canonical-2313130132033123-1112031111102031-1223212210311102-2130001131323301-3131211123300222-2322332131121112-1310000100002011-2101033212002231"></a>

#### `https_management.advertise_on_slo_sli.tls_config.custom_security.min_version` property

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

<a id="canonical-1002021221233202-3130021310012220-1310110003230101-3003323320213322-3310110222322330-3130121121120100-1201313002332131-3131200121123002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.default_security

<a id="canonical-2211111111233321-0331013132321302-3320211030003020-1223022020000203-1100120100000120-0123112203222110-0332211033110130-2103220123030321"></a>

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

<a id="canonical-1313030333021133-1013033023021223-0033130302313320-2212101202312033-0311010232220003-0113012233000221-2302032020012303-2033302131321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.low_security

<a id="canonical-3011310322301103-3000300013103011-3101221220033103-3223112202023023-3322113113133100-0211312112220302-0131210110130212-1033311012332032"></a>

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

<a id="canonical-2112302102010123-1033302111322023-1121001202132030-0321211301021213-1302021203010030-2111330211332132-2210102210033100-1011303231030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.tls_config](resources--nfv_service--reference--group-003.md#canonical-3220330322200213-1131332232110011-2033122101021203-3133303310201321-1100333102323010-2210233023331021-2221332001303332-0331211032201031)
- https_management.advertise_on_slo_sli.tls_config.medium_security

<a id="canonical-2022100312330102-2103131100103111-2202030202121132-2303013301020101-1311101102203012-2122011310111232-1013212131110203-3332011332302023"></a>

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

<a id="canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- https_management.advertise_on_slo_sli.use_mtls

<a id="canonical-0101002012210331-3130133103321132-0122211223123002-3131001203213102-0123232111233223-0032233021132033-1101330321232312-3321311203121303"></a>

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

<a id="canonical-2313021011221113-2000020100101001-1120102013020120-0113031332303233-0221311303031321-3110213113012031-2311323130102333-1101312012111233"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls`

<a id="canonical-1202210002310331-1110132200031013-2021320321222332-1300301232123000-0022022022230212-0311032010312010-0003130130122203-3231223002012032"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.client_certificate_optional` property

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

- [crl](resources--nfv_service--reference--group-003.md#canonical-2122323020322313-3112221322132323-3230213310332002-2122231223123100-1213323123202002-2122210023333030-3210132311133003-1010300231330103): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-003.md#canonical-1203023000000003-2302301231310103-2023100211121011-0201303020031000-1022023201333321-1220102231102200-3131222002220233-1133322300102301): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-003.md#canonical-0300022102301312-1131213301311333-1313011200032130-1033221120132133-1331001133010231-2223002012321110-2200122122311223-1131233322333203): complete subsection reference.

<a id="canonical-0300102112310320-0012011321030133-2103130122102211-1002202332202301-2121321330312133-3323213030132013-0221200312122120-3010221100103013"></a>

<a id="canonical-0322112323020322-1122200000132000-2022330000310332-1111210030203122-3201300310133233-0302112201221003-3001133020312211-3233321233321200"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [xfcc_disabled](resources--nfv_service--reference--group-003.md#canonical-3303321210232100-3033322310101121-1131030131110032-2003321300131230-3320222310111123-1320012022120011-0212313223103013-3130123033102210): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-2202121020010233-0110202113211132-2332301321323120-0021021112223301-3210202210232031-3233002032120033-0222202210211133-1010223320333201): complete subsection reference.

<a id="canonical-2122323020322313-3112221322132323-3230213310332002-2122231223123100-1213323123202002-2122210023333030-3210132311133003-1010300231330103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.crl

<a id="canonical-2102223303121033-3003323311233023-1310300103003310-2311001233230132-0333231323232111-1102111212221010-1130033012313322-2222002311022231"></a>

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

<a id="canonical-1110001302132310-0012110113200212-1130121202210320-0100222011210230-0031111113111213-0302111333012003-0212032000020222-3221001022002021"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.crl`

<a id="canonical-3011022203222132-1132111132033011-1130022201303223-2331312302203300-2313023311231233-2320012223203120-1331203130111020-2300131302121003"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230033001232033-2120033231223231-3332201020213230-3332311002030110-3002310202032301-3332100332320202-0133302222320203-3330031322301103"></a>

<a id="canonical-2300103032130313-1202021333310112-2120330330111231-3000230012333101-0232301112123300-3321200321001231-3113331303203122-1133313203201033"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0332303302220003-0332222112201002-2310131133010221-2302103302302223-2301223300013003-3233013122002010-2221310212032123-2103302300030311"></a>

<a id="canonical-3022022010201310-1332111333330233-0023113312311000-1101030112311321-2001113102030233-3123213103120230-2031002131022032-2020033100311011"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.crl.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1203023000000003-2302301231310103-2023100211121011-0201303020031000-1022023201333321-1220102231102200-3131222002220233-1133322300102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.no_crl

<a id="canonical-0030012230222331-2002311130032222-0011012220130033-1030323331011220-0223000113223232-2231303130322210-3201213101333312-2220132130033011"></a>

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

<a id="canonical-0300022102301312-1131213301311333-1313011200032130-1033221120132133-1331001133010231-2223002012321110-2200122122311223-1131233322333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.trusted_ca

<a id="canonical-3131023010230002-0231002203322301-1132012311333303-1310010312023212-0320131120220230-2221212300110100-3011231323110122-1312223033212110"></a>

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

<a id="canonical-0020231112213130-0200033230010112-3331213012010213-1321031022230023-0120200230011123-3121222002113200-3103210200020001-1132322212013213"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.trusted_ca`

<a id="canonical-1301301030001303-0020323202023312-1313333113033231-2110323101021310-3022100012321201-1322201303111133-1231030012311131-0123301111030233"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1302130320313333-2012313032001211-1200221321210311-2102031331103321-0213303110013013-1203310110112031-3220120122320332-3312200113132232"></a>

<a id="canonical-2132112333030201-1302301021213313-2221321200212130-0100021001123031-1022102130332100-1333313032013233-2001222200323000-3310222330001112"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0221210212212310-2103332213201110-3032201023222313-3030120103322020-2110210300201123-0013123203310002-0003223203322030-0111101023331002"></a>

<a id="canonical-3021323002022333-2332320311133200-3230311121030322-0012203202120211-0310321212222000-3100301310203003-3131122220000211-1133022031131211"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.trusted_ca.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3303321210232100-3033322310101121-1131030131110032-2003321300131230-3320222310111123-1320012022120011-0212313223103013-3130123033102210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_disabled

<a id="canonical-0230021031001003-3311031210100201-2100011213310322-2032231013301012-3033303313211333-0113110111312300-1030312301333101-1132202020211222"></a>

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

<a id="canonical-2202121020010233-0110202113211132-2332301321323120-0021021112223301-3210202210232031-3233002032120033-0222202210211133-1010223320333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_sli.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-0300000202230320-1330311113123300-3013200231333023-2210113332222133-1202010133200331-1103103321101332-1311010232212102-3002120020221012)
- [https_management.advertise_on_slo_sli.use_mtls](resources--nfv_service--reference--group-003.md#canonical-0102111312001022-1020221220200232-1320312232121311-0003320013313230-0032133233002223-0200032003101212-3120000213032311-3002030123333310)
- https_management.advertise_on_slo_sli.use_mtls.xfcc_options

<a id="canonical-0320223201130322-1301030333320222-2120313130300112-2112210121030100-0321020002131301-1311112010311130-3323330122211021-2111211212202230"></a>

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

<a id="canonical-2023131202322322-1301002131021113-2102033203011133-1020103123130332-3212200013233130-1112101212223210-3133302322222123-2001122333321020"></a>

### Direct properties for `https_management.advertise_on_slo_sli.use_mtls.xfcc_options`

<a id="canonical-1232002303002201-3220113203023223-3122213121331031-3021320312302003-3121131223011212-1222310221312021-1103200021011313-1201031011203021"></a>

#### `https_management.advertise_on_slo_sli.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- https_management.advertise_on_slo_vip

<a id="canonical-0311030123113020-0310323023030223-2132032033103233-3000210331321312-1131200021222000-1232022213100001-1221331332332222-3313121002001232"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
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
advertise_on_slo_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222331132031103-2131232030022312-1231313020022313-2103023310300201-2320011203101302-0222032022222002-0221310333003020-3111230300030221"></a>

### Direct properties for `https_management.advertise_on_slo_vip`

- [no_mtls](resources--nfv_service--reference--group-003.md#canonical-0211010101302002-1033333201312123-0303321221111102-0233212330023110-1221033301032313-3130023100303333-1320012331023231-3010321120110301): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-004.md#canonical-1220000301313030-3210123130033303-0000003031222101-0230222200021002-2212330020213020-3221320323031031-1203232101302221-3230202011332313): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-004.md#canonical-0013001011312211-0302330103310232-2330013030103010-2111232121111121-1132202200220233-3303310322202002-2110202013220100-2223302323332223): complete subsection reference.

<a id="canonical-0211010101302002-1033333201312123-0303321221111102-0233212330023110-1221033301032313-3130023100303333-1320012331023231-3010321120110301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.no_mtls` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.no_mtls

<a id="canonical-0223102102213323-3120131112012331-2011112320231311-1021011120131321-0223330102011211-1303132001130100-2120020021323301-0033120103310330"></a>

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

<a id="canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="canonical-2033311030122032-3110102122003230-0011002203212233-1032211121000322-2122131133333230-0200313213020211-1202331022100100-3301023220233000"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333231212312301-3201130222232021-1032001021111200-0020030203103113-2302212120300311-2302032011010202-1121203313330312-1033333310030102"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates`

<a id="canonical-1120032033333030-3203101020110111-2300121111222022-0213213231333133-3222331220330300-2303112003300321-1310021131201312-1200123211122223"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [custom_hash_algorithms](resources--nfv_service--reference--group-003.md#canonical-3332021021113313-0310032003132330-0222331210123333-3322003030331020-0210101130020113-1321321012202011-1101101010333033-1032011210212011): complete subsection reference.

<a id="canonical-1301110230122310-3000311233031332-3121111031002000-3010230001301101-1101203131321121-0213110012203131-2323323311220321-3023322223020220"></a>

<a id="canonical-0323310303312200-3103213120332101-3230233013330132-3311020302222133-2210112032032211-3213311003101120-0102211012123032-2132020103133311"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-003.md#canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-004.md#canonical-0022012123311000-1212203230321020-1330231021210332-3201020233133002-2323100303020232-0312203302233123-3000300312202110-2223011033132013): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-004.md#canonical-1011101211010132-3312322032312331-3113201301221030-0122312102211300-0221000121010231-3102333110002230-2312323213001223-1111002232331221): complete subsection reference.

<a id="canonical-3332021021113313-0310032003132330-0222331210123333-3322003030331020-0210101130020113-1321321012202011-1101101010333033-1032011210212011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-1123112301132312-1301222223010313-1013321112310331-2130133001312332-2332222000032303-1301231010301310-3120203211332011-2231222100321310"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002010233311010-3321233120122021-2332301231010320-2101233211212033-2301012330033133-3002022212020023-0213122020023203-0312012021333210"></a>

### Direct properties for `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms`

<a id="canonical-3213230122223201-0300123032111130-1022332131222302-0101131123311331-1302131210333213-2031331003312132-1220010222210302-1032110020322333"></a>

#### `https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-1223312232131130-3323102031233232-3103230230232313-0023320033320202-0210000233222211-3012010210210210-2201000102130223-2220232000002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-0223321113210321-0013210302112221-0313201132021001-1102330023301310-1131322132000230-3303202123330303-3122032322302202-2202303201230222)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-1033200313012130-2303320020000220-3032203302003212-3030323031320111-0322032022323123-0203102131300133-1232321110231101-1022222010201320)
- [https_management](resources--nfv_service--reference--group-001.md#canonical-2012312013030011-0102312120121131-2122112110311110-1102330303132211-1031011331133302-3311032130133131-3022001211230302-3310030220221101)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-3123232331313231-1003300030030012-0013233023203203-0221011121303133-2002120310330210-0313012223300221-2030131203030100-3230303033112000)
- [https_management.advertise_on_slo_vip.tls_certificates](resources--nfv_service--reference--group-003.md#canonical-0201100000120203-2021020301330032-0310123232303233-0102332130311132-1233200003321230-1112232313301301-3210030321330101-0101132333333033)
- https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-0221121130213003-0221302100231011-1332330232032013-0000321233021313-2212103103133210-1013212211333213-1330112012111233-0323300330223112"></a>

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
