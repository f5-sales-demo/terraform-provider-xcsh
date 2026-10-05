---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1000012313203211-3311022000102313-3330113333200022-0013220113310322-2111132333303022-3301130133312121-3202111122302301-2210320032000013"></a>

## https.tls_parameters.use_mtls.no_crl — no_crl / 002302130223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-1230002332301133-0313013212112020-0323120312100111-0310012323202031-1112132103230200-1321203313230310-0213113231131302-1231323322211312"></a>

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

<a id="canonical-3000033120121212-2201113312111123-0010023103331331-0313313002103023-2303020331301210-3111012131011232-0133230303010131-1123312203012111"></a>

## Direct properties — no_crl / 002302130223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302023201203120-0311321010112301-3133020301100122-0121221030323031-2130300330133220-3133102131122312-0102200122132031-2112112231031120"></a>

## Next pages — no_crl / 002302130223 / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121212130323023-0101122011111203-3321231021310302-0020303121102233-1021001202033022-0120331301101131-2330211002201122-3220320130202313"></a>

## https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 212320010010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2122333012202002-3210202000122111-2111210002202102-1002320330333210-2200110010013233-2310133101231233-0221320203011313-2221020100310333"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0320210211100310-0111100032003203-0331133020302300-3323012030223033-2230112121030323-2232031030030130-2310002100023013-1301110320321130"></a>

## Direct properties — trusted_ca / 212320010010 / 3

<a id="canonical-1223133011301133-0120233310020313-0122030003032231-1123032031101223-0100001212120012-1103311333101000-1310020032121123-1331131122303102"></a>

<a id="canonical-1122323132322330-3013112213022010-2321111233122113-3210120223203223-0301311330121031-3022001310010323-2303001332100010-2023322230223011"></a>

## name property — trusted_ca / 212320010010 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1211333230221201-2320332332322311-2331010033133022-3122020323223032-0213013030132032-0210122111303321-2120023203210002-2010120011331321"></a>

<a id="canonical-2032002210012021-2030222012330201-3001332023112122-0011202012212111-1010212230312332-0023131012311020-1132222033113302-2122102132233110"></a>

## namespace property — trusted_ca / 212320010010 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3120200012201003-1112212113021130-3212303133303103-0333031112000223-2002212020300233-0102323200212233-0323212300300113-2310320122310302"></a>

<a id="canonical-3121230000120021-2200331210000231-3221011300122122-1132120022220220-0010213223012032-0333231320310321-2221000232223233-0220310021320220"></a>

## tenant property — trusted_ca / 212320010010 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1300110120102033-0323330303303112-3302321003001130-0203301002012131-2203213200112102-3303133111100000-2031232320121132-1301231030312022"></a>

## Next pages — trusted_ca / 212320010010 / 7

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302123001113332-3230001222023103-3110113000322213-1231000013133322-1133231230022113-1112003023032011-2233202112231200-0111001032333333"></a>

## https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 120001233100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2001330313213112-2003121033010131-1233202331001132-1031030001011100-3332012002200003-0310121321203212-1121303033022021-1113231103133003"></a>

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

<a id="canonical-0120031033322300-2303200222121212-3131323022120020-1331112113102312-0230103121221331-1220112311201313-1322131023133300-1100331121101310"></a>

## Direct properties — xfcc_disabled / 120001233100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132321011232113-2310213033331113-2113200121221322-3330101333132332-1030332110102312-3313111302313032-2113122012010220-1213033223010213"></a>

## Next pages — xfcc_disabled / 120001233100 / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101123200330222-1033130332223212-3110001021211222-1213003121333011-1203121131123100-3013300211021023-2020331333032112-1233113110132022"></a>

## https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 220103310323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-019.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2300200302102222-2312213000310030-2331231223103023-2013230302023320-0232001013322012-0121010030132101-2222100220202321-2100231300102313"></a>

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

<a id="canonical-1211201320201230-2033211001032221-3030132332112031-1213010223102332-3102122231120312-3000223133010212-3010030110322310-2023002020020113"></a>

## Direct properties — xfcc_options / 220103310323 / 3

<a id="canonical-2103202323130012-0131122233030220-0321022231131220-2122013031002230-1220021330011210-2003111113210102-0111003120222223-3102120320233331"></a>

<a id="canonical-1120323120002000-0101023303220233-0233301323303210-1231000200131101-0220231000333333-0213132332131201-3203103000230120-1300103311230300"></a>

## xfcc_header_elements property — xfcc_options / 220103310323 / 4

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

<a id="canonical-1033111201131031-2031122101111032-2232302300101332-0332123103333110-0221302012333102-0121232003001322-3201030232022130-0221321211321303"></a>

## Next pages — xfcc_options / 220103310323 / 5

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212233130303102-0211032023003220-1330010201111021-2030111110113301-1120330113022112-1020032101331231-1112030202303213-3111303311000323"></a>

## https_auto_cert — https_auto_cert / 133011330031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- https_auto_cert

<a id="canonical-1223123300112032-0222133020221001-3201331012031023-0133020113303300-1120222301233102-1120232121100332-2220301031031210-0023330100002100"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

Upstream description:

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

<a id="canonical-1001100232130231-1001110203232212-1212101300330122-1121301331231033-1233310320200030-2212221012103010-3323200113132311-1002022122031230"></a>

## Direct properties — https_auto_cert / 133011330031 / 3

<a id="canonical-3223321310121123-2301223021013011-2111300211130232-0122000300102301-3333333133201231-1010132220131130-1303010200233013-3010322133001333"></a>

<a id="canonical-0111333230002111-2101133003323202-3320031233202301-2133313001213221-0211332112211300-3220122102021323-2000023111002210-1131331202132303"></a>

## add_hsts property — https_auto_cert / 133011330031 / 4

Type: `"bool"`. Optional, Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Upstream description:

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

<a id="canonical-3031310103320212-0110301112112033-2022100102210201-0222311200101133-0311111221202023-2132303131333320-3233213120130133-0323322103002022"></a>

<a id="canonical-3203333030333211-2003132223020311-0012330203101012-1002133001002122-0223210201300201-2031032303122031-0311220231032233-1321323100230303"></a>

## append_server_name property — https_auto_cert / 133011330031 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103): complete subsection reference.

<a id="canonical-0102002212110200-1321022333101111-3313032013110313-0200122112001310-3012223110030220-0112302110111312-0211132011302230-3333313201301130"></a>

<a id="canonical-1101201322323321-2300331220002202-3322312331221202-3000011301320220-3120021212122122-0120221201112032-0323211122212330-1012331321212230"></a>

## connection_idle_timeout property — https_auto_cert / 133011330031 / 6

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--http_loadbalancer--reference--group-020.md#canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-020.md#canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122): complete subsection reference.

<a id="canonical-0323010101202030-1232230010101110-0301331011022103-2100331221100330-0032022001012220-2222002120303332-1322322020331233-1112132333110203"></a>

<a id="canonical-1232133303302003-2013103311313102-3311202011100320-2311011220323231-3211223231323203-3122021132100011-2230323232233121-2033220230130301"></a>

## http_redirect property — https_auto_cert / 133011330031 / 7

Type: `"bool"`. Optional, Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332): complete subsection reference.

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-020.md#canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-020.md#canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312): complete subsection reference.

<a id="canonical-3102122321001103-0203131322222310-0201221031133023-2212331332021130-3320031012033222-3323112322220220-2232021011310222-1232131121311112"></a>

<a id="canonical-0113300213233320-1030330110032300-3231102113312232-1103333030330320-0232020010021002-3030320213032111-2323132022102032-1122212002232333"></a>

## port property — https_auto_cert / 133011330031 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2013121000300130-2202001203102032-1313012203022002-0302230002020312-3221021222033303-2202301323301033-0320111021020130-1123131122103321"></a>

<a id="canonical-2323102320312112-0333320123220313-3311333201322003-0002123020123312-3213313030113211-3222032132220032-1301210211213012-3313120211333003"></a>

## port_ranges property — https_auto_cert / 133011330031 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0310020220232133-0223222101221100-1230232033110333-0333310311322030-0030221331331110-2200233121213321-0133212311133101-2230303123022223"></a>

<a id="canonical-2232110031300123-2330012131210313-3111111311201213-2111310320302103-1100233100121132-2003120200222032-1222131232133303-0310113203030212"></a>

## server_name property — https_auto_cert / 133011330031 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131): complete subsection reference.

<a id="canonical-1101333332003310-3110123121101232-2133000132123021-3022131022302132-1113003322023220-3223100313120131-0133102111133312-3201002310322230"></a>

## Next pages — https_auto_cert / 133011330031 / 11

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- [https_auto_cert.default_header](resources--http_loadbalancer--reference--group-020.md#canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120)
- [https_auto_cert.default_loadbalancer](resources--http_loadbalancer--reference--group-020.md#canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213)
- [https_auto_cert.disable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330)
- [https_auto_cert.enable_path_normalize](resources--http_loadbalancer--reference--group-020.md#canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.no_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332)
- [https_auto_cert.non_default_loadbalancer](resources--http_loadbalancer--reference--group-020.md#canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003)
- [https_auto_cert.pass_through](resources--http_loadbalancer--reference--group-020.md#canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120322221033330-0132332223122213-2133022113102222-0233222200330113-0310231213200031-1302001100101323-1200032323031203-3220320202013031"></a>

## https_auto_cert.coalescing_options — coalescing_options / 133021030121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.coalescing_options

<a id="canonical-3001231212332201-0313132220012130-0030022212322000-3112123233103301-0003021133311302-3323021100230320-1022133233231222-0310213323221033"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

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

<a id="canonical-1010221022003123-1303010020010213-2233101133311031-3031322013220031-2003132210112211-0120102132310023-3011220300211332-0203201033332000"></a>

## Direct properties — coalescing_options / 133021030121 / 3

- [default_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033): complete subsection reference.

<a id="canonical-1321010020230331-1100011011200001-0123300100311331-3013320101031122-0113013310311121-2012221220121020-1101012131030222-0021231310303031"></a>

## Next pages — coalescing_options / 133021030121 / 4

- [https_auto_cert.coalescing_options.default_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300)
- [https_auto_cert.coalescing_options.strict_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013022113020203-3131323210311113-2120203211130110-2102221323133121-3101310011020203-0323121310332213-1021121203313132-2211321213123011"></a>

## https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 233232131222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1101100331131002-3203323132211211-1313320203333302-1133030220030012-2331011320023013-0313322303300333-3231113222300000-1013220231003331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-0301320221031323-0012100130020200-1333003333322120-1121003102133213-3221003331233321-3022010121322021-3100033232220311-2202013223230212"></a>

## Direct properties — default_coalescing / 233232131222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221201320310221-0201101320330010-0320323212133103-3333233213330333-2132323010223010-3323332311103300-1230011211013230-3323300202203012"></a>

## Next pages — default_coalescing / 233232131222 / 4

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000130310232300-2121222320022300-0111000022122311-3030101001111031-0011231221020121-3202202202132220-1232211323300233-2231230330203303"></a>

## https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 313230223103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3200023331202132-2301112331203032-1112102033020122-0211031100203321-3101130023332222-2300210313110212-2111302020301211-1202323210220222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-2022213200012003-0320321122322223-0311110333323221-0002311023123333-2013033112220330-3023002103200210-3020213212320112-3312313103222232"></a>

## Direct properties — strict_coalescing / 313230223103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012330220003021-2021030100131003-3202123100032322-0323101032322103-3202011321302330-3213203203122133-3022030322211023-1132000010223233"></a>

## Next pages — strict_coalescing / 313230223103 / 4

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320312020301121-1030113132332333-1110211132122012-0103230332220203-2122320310231331-0031012100100113-2313130212330111-2123212021303321"></a>

## https_auto_cert.default_header — default_header / 132332001321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_header

<a id="canonical-2100311123231131-0133311110202100-2332023030121232-0023223011001303-2031232033201002-3302130331132312-3030110102022311-2200021303333233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-1312011200121101-0332122212303111-3331231003201201-1011232112220221-1222220032321033-1201230021222211-3330323100120110-2211032311012023"></a>

## Direct properties — default_header / 132332001321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111233321302223-1210322212200301-2111112101002212-2321203132003333-0021322132111210-0113330003131032-0233103101031232-0331210312230330"></a>

## Next pages — default_header / 132332001321 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002201003012231-3231310000020220-0212310223202312-0011201023311230-0220002223000113-0113012310022332-2321110010303033-1020002210231113"></a>

## https_auto_cert.default_loadbalancer — default_loadbalancer / 303211032330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_loadbalancer

<a id="canonical-3001000020233233-3001321322312021-2322123012222021-0133201022333112-3323333211203103-0330302331102301-1232203033312201-0302322313113110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-2200111003300103-3011131300231132-0230322210021223-3310030300210030-3200002030300121-0323030010003210-1230321021302201-0033321130023301"></a>

## Direct properties — default_loadbalancer / 303211032330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300102110022132-0301210131132201-1310310022331121-1312232101022131-0111222002131100-1013212002220230-3120220021231032-0331132012011011"></a>

## Next pages — default_loadbalancer / 303211032330 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323211223121021-3111001210012312-3112200123212020-2303003312220223-1332321221221323-2022122020000320-0330113302301031-2013133301111001"></a>

## https_auto_cert.disable_path_normalize — disable_path_normalize / 030210321320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.disable_path_normalize

<a id="canonical-0333213122002122-2030101001233332-2123333323000121-1110122231213323-3030331031030011-1223332223313103-3230211212032010-1302311321022120"></a>

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
disable_path_normalize = {}
```

<a id="canonical-3332323130320231-3301101321311013-0201210313233231-2021102110120221-1201223123101222-0120031211010231-1000103223013133-3223120020000312"></a>

## Direct properties — disable_path_normalize / 030210321320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112330002330022-2311033221301210-0021033030312120-3211222230302221-2032103310212230-3102230023221102-1231330120002131-1111322220220312"></a>

## Next pages — disable_path_normalize / 030210321320 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220113213322022-0111032330020322-3312313232230031-0232102122210101-1031200312122202-2021013000210321-0023003203223300-1333303122123323"></a>

## https_auto_cert.enable_path_normalize — enable_path_normalize / 112102313300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.enable_path_normalize

<a id="canonical-2223013223201010-3303013330023003-3300321112131202-0222013313100103-1320203101220202-2130020300130020-0313322332022230-2101230131000111"></a>

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
enable_path_normalize = {}
```

<a id="canonical-3330222013301020-3033321223223230-1023023310213110-1313123200320213-2323021221203120-1111312113023112-2233002223021103-2133311311232133"></a>

## Direct properties — enable_path_normalize / 112102313300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122233130030312-0213301301323033-0010200210311313-0331020203230303-0122133022132201-3210101210100330-3213130311113222-3231203323333310"></a>

## Next pages — enable_path_normalize / 112102313300 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201211232312221-0313321232012333-1321101010111011-1203230113200033-2002232122233211-2321021130321220-1010202121310210-3210003303021102"></a>

## https_auto_cert.http_protocol_options — http_protocol_options / 033113323012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.http_protocol_options

<a id="canonical-2011113132131102-1110223121331230-3310333210200303-2221103221112122-2231100322130011-2202011222323313-3020013102332103-0013323210332230"></a>

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

<a id="canonical-1331121031311213-1012023211323131-1233303333300333-1102211203330332-1011122023321211-3230310101130311-1320223221202013-3331021323003220"></a>

## Direct properties — http_protocol_options / 033113323012 / 3

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-020.md#canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-020.md#canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002): complete subsection reference.

<a id="canonical-2032102313102110-3213332123201110-2310011301210102-2231330310132132-3313000013323032-0033100122120333-3021010123121233-0320010312330003"></a>

## Next pages — http_protocol_options / 033113323012 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-020.md#canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-020.md#canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302013103010301-3000331031202230-2303132032222032-0210212012210320-1031122101112212-3233032003303323-1230131321331311-3200200020313232"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 031222022223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-2031312311322022-2323332220121203-1103310231121222-3201301221130333-0101003230132100-2012103000330332-2201203132212101-1121303022130303"></a>

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

<a id="canonical-1012223331210122-0110301302220123-1011212233022112-2200001120010331-2202302001223223-3311002003300201-1112010022221131-0300212110113302"></a>

## Direct properties — http_protocol_enable_v1_only / 031222022223 / 3

- [header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223): complete subsection reference.

<a id="canonical-2001101200233111-2130032231031300-0303120132301132-0102113133332332-0310031012022223-2203213000131203-2232232022232311-0123302021330001"></a>

## Next pages — http_protocol_enable_v1_only / 031222022223 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020010330230120-1310110000020231-0010231003133211-0032103231320200-2001212200313222-1123220211231123-3303121200220020-3122213212303211"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 121210212220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3223012123130213-3232203130132033-3233200120000330-2301100133230320-1123122133303311-2121332123111230-1310121310032223-0120132231012113"></a>

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

<a id="canonical-3220301021100101-1330021131213033-1023033000102313-0203003132011010-1101030032032323-3231122201310333-2211112323130332-3101322211302112"></a>

## Direct properties — header_transformation / 121210212220 / 3

- [default_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013): complete subsection reference.

<a id="canonical-0220131000012131-3203003331312223-2012320102330010-0022131101021130-2320301120120120-1300332302223230-2330111012111331-0023221231333130"></a>

## Next pages — header_transformation / 121210212220 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231321331013110-0003223333222210-1111101201121200-1002121330103130-2012112111223222-2322221332011101-1200313132101113-3012021221110202"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 300133122121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-2222233231313110-0113330311010200-3323002031102213-3132300131203323-2201023000323030-2333310321111220-0130122301213330-2103003001331332"></a>

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

<a id="canonical-3220210201223030-1000132032211122-1302323212331122-2100220011220012-2020032201122121-3231303013013202-1122012223130122-1132000331120020"></a>

## Direct properties — default_header_transformation / 300133122121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021103010303200-1021302230123201-1332011102331321-1213113330233301-0300202200021010-3322312322222003-2112232232232322-3333232121303321"></a>

## Next pages — default_header_transformation / 300133122121 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000220003302201-2023230031101023-1233013310313130-3302313300010112-1331011302121000-1222213303100110-0323301001300202-3010323201023010"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 313001211001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2330011023030002-3100012200112020-2211222321100330-2103322013300303-2003000212023231-1211311132121102-1113331102012121-1231112132102123"></a>

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

<a id="canonical-1201031232112203-3321330010130023-0202101221111323-3123022111311311-0320333331120121-3323233112331302-0121110220331011-3012123222201233"></a>

## Direct properties — preserve_case_header_transformation / 313001211001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100312011130013-3122133231311223-2032023111011033-2021132010222021-2313322002330101-2023112333130200-0121232331231103-0221021312010312"></a>

## Next pages — preserve_case_header_transformation / 313001211001 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313033320122020-2102121013230212-2333021113121312-0233133332313310-1313121020132101-0310302220213330-0033121211120101-2313000102210002"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 220303032312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1333101202011333-0032333210231203-2110202312123120-1010110033030321-3300330311001321-0130002033221101-2320311333213133-0333102113130001"></a>

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

<a id="canonical-0111021032130111-0133212112231323-1012131111233120-3012010121023321-2030112330011220-3121310320102133-2300330132302310-2212212003132033"></a>

## Direct properties — proper_case_header_transformation / 220303032312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230002112320023-2101031203001123-3322311201002023-3231311331033032-3322312111103021-2233302013332200-1213212032123122-3121310112221233"></a>

## Next pages — proper_case_header_transformation / 220303032312 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021323131113020-1323121332220203-2131021013101113-0211210111330002-2010323121300110-2230003102102030-3332301201011002-2311002320323310"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 212111322312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2123230010031123-1020303333310032-3131112012111011-2333133011021320-2203123130002302-1312012022032231-1321203033101131-1323131121133302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-3203233023221311-1013031211200212-2312033131213102-0201000030112311-3030203203122110-1110331122200033-1331201211231300-1013322331230023"></a>

## Direct properties — http_protocol_enable_v1_v2 / 212111322312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202002002301103-2033023222310031-3111331022011333-2100200223110311-0013311022222113-3032202213230202-0122302023021232-2223221222232213"></a>

## Next pages — http_protocol_enable_v1_v2 / 212111322312 / 4

- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031131133220303-0231000011223201-1320033233001302-1002131010133121-0131313232232202-2001001001000130-2012210000130002-1320211333100100"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 133322331121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3122203002332132-0123021301103122-3103030331002102-2300331311323121-2330030312032202-2120020210212101-2133132222031032-1220013010123022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-2221203010202321-0001220033132221-3100132010312020-0121013302112332-1100031233313003-1211211031301331-3021123000001110-3220023322330120"></a>

## Direct properties — http_protocol_enable_v2_only / 133322331121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003001310311323-3223223233212023-3231312021011300-2110000211103332-3320102220322113-0102021313220031-2111131111110232-3013333211130022"></a>

## Next pages — http_protocol_enable_v2_only / 133322331121 / 4

- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002011331322321-1313100102333132-2023312122100322-1030123111223220-3131300122102203-3132001132320023-1101131033022221-0202112002300002"></a>

## https_auto_cert.no_mtls — no_mtls / 012023111111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.no_mtls

<a id="canonical-2220332020001020-2302321002301012-2311000132102222-1122221231100012-3312231011131012-2303331110220020-3002122100311031-1211320332103221"></a>

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
no_mtls = {}
```

<a id="canonical-1331131203002210-0211001013110131-2213133322300010-2131211001202202-0232231000332000-0112213302113313-0003121222103233-2202031102122211"></a>

## Direct properties — no_mtls / 012023111111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013111100102033-1212030221022020-2321002323203221-1303320223011021-3020103032333002-3233121202300330-2211031322121023-1122313110033121"></a>

## Next pages — no_mtls / 012023111111 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332220122223303-0332210110010201-2313031111231131-3330303301302101-0001210322113121-3320131130010303-2302320200203211-3001321211112003"></a>

## https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 323233023312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-1001221102303011-2220021100220012-0230032100202011-1301332010210020-2030022302012303-3302231211303103-2211120212331103-2333133230111310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-2133212121223120-0010131123203031-3111022112121001-2203310213022310-2200123231103012-0200123213331230-0301003201000002-1020323103131012"></a>

## Direct properties — non_default_loadbalancer / 323233023312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231313320122231-1000101310222313-1313331230203332-1230123233010201-1220202102333002-3320013320100331-3333133323010332-1103210030001303"></a>

## Next pages — non_default_loadbalancer / 323233023312 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131013021002001-3213330102322133-1211102312232100-1101020231202311-0301102113211212-1023030320103102-0122013201232301-1232030103020001"></a>

## https_auto_cert.pass_through — pass_through / 333031032023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.pass_through

<a id="canonical-3311221012012030-3213230332021221-3212331303133020-0333212210303111-1122021210220031-2131202102133211-0213003033202211-0220210023210101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-2230333310110002-2020233312031130-0000313030233013-1032033330330312-2001001021110003-1131322112211021-3031310132213031-2003132332223020"></a>

## Direct properties — pass_through / 333031032023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300113311203200-0130020212332102-2033230133333330-0131011123120301-1233023032322331-3033103310032133-2110233030231331-1030021010302010"></a>

## Next pages — pass_through / 333031032023 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013311133002230-1233330322211320-1012013032323222-2313301320102232-3000101021131210-3102331332322313-2111020210230230-0131012132300222"></a>

## https_auto_cert.tls_config — tls_config / 211113210311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.tls_config

<a id="canonical-0233120200312332-1211103300220120-2332300211130200-2113122123122012-0231332031100303-1132222200303331-2330333330121112-0033311121331021"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

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

<a id="canonical-2310211201133020-0332232011331101-2001231023333010-2031112120203003-1003130021103320-3102131001320300-1321032130133220-0312220110100033"></a>

## Direct properties — tls_config / 211113210311 / 3

- [custom_security](resources--http_loadbalancer--reference--group-020.md#canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-020.md#canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-020.md#canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-020.md#canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320): complete subsection reference.

<a id="canonical-1032121002022203-3013221012313130-0031231301111131-3221132302212332-1222333213022020-3233002112220020-3330120221121012-2321233303110210"></a>

## Next pages — tls_config / 211113210311 / 4

- [https_auto_cert.tls_config.custom_security](resources--http_loadbalancer--reference--group-020.md#canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220)
- [https_auto_cert.tls_config.default_security](resources--http_loadbalancer--reference--group-020.md#canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303)
- [https_auto_cert.tls_config.low_security](resources--http_loadbalancer--reference--group-020.md#canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312)
- [https_auto_cert.tls_config.medium_security](resources--http_loadbalancer--reference--group-020.md#canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310000212323301-3230210312213230-2011211322322332-1221001332030311-2232000032310022-0212112112012220-0331301220221133-0332123122121013"></a>

## https_auto_cert.tls_config.custom_security — custom_security / 101312031011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.custom_security

<a id="canonical-1113313033212323-1021103010103233-0322232311101112-3022122122120223-1311212203012101-1101220012302303-1313212022223100-2101310330031212"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

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

<a id="canonical-1020310330320330-1010210023230002-1310312023201223-1331032301201201-0321033003002223-2032010000101200-0203301221322132-0002310233322332"></a>

## Direct properties — custom_security / 101312031011 / 3

<a id="canonical-2011223032011232-3122112313013030-2330002032002232-0332023111103022-2300103322120111-2110102321111022-0232020310121312-1102332032220030"></a>

<a id="canonical-3303213202012220-1332013112221031-2003021112013323-3123131101110012-0123022012123202-0211211312220132-2202321320231003-2000213010020213"></a>

## cipher_suites property — custom_security / 101312031011 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0000312322102231-1030221133131220-0133321201312000-2132002003332112-1001013233213100-2202201031131022-1311202020320021-1222333312301030"></a>

<a id="canonical-3233323000322321-1001233131102203-0023132323233011-2103322312332333-1300222012100110-2111213110330111-2122312220223132-1312103123110331"></a>

## max_version property — custom_security / 101312031011 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-0101311222222313-3100120110113331-1313221201011301-2011123033301001-3311331222011023-3012311323321133-1230012122223321-2311222101031312"></a>

<a id="canonical-3331132120120010-1203010030022121-2233031101122122-2312120303000120-0222310130331102-2112231001221000-0112302013120112-1133013231203300"></a>

## min_version property — custom_security / 101312031011 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-2132022103202310-0321223010111233-3232303100030211-1123300303103111-2312021123230031-1302230332132103-3223322113333131-0320311010231203"></a>

## Next pages — custom_security / 101312031011 / 7

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132321332233130-1330323132102103-3211022111113301-0133210231033330-1032200112100001-3123121321221020-3232100001230222-3003221002210011"></a>

## https_auto_cert.tls_config.default_security — default_security / 100113000321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.default_security

<a id="canonical-2202001103333021-2002220110001221-2202203001012330-3023320211102210-0300012131312102-1113032320203000-0211230312211022-1101212230013110"></a>

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

<a id="canonical-3103103233112020-0032020033103200-0203021102001003-0002123101131031-1001231201121223-2220000113100001-3212202031300132-1312002221303023"></a>

## Direct properties — default_security / 100113000321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300003032113301-0320121311011332-1101331022033133-1023312020022113-3113133323330101-1220113033020013-1101020330212132-0032131332303120"></a>

## Next pages — default_security / 100113000321 / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230202101113203-2101333312200302-2033131302201232-1311331313220322-0012131231112133-0331001102231122-1013310023200113-3111102211202100"></a>

## https_auto_cert.tls_config.low_security — low_security / 321302223130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.low_security

<a id="canonical-0201122001030021-2100130002313122-2332212021333030-1110221320113022-0032232203231303-3113312031202010-3100323021020020-2213332210201103"></a>

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

<a id="canonical-0011300302200223-2100202122203310-3331210011010331-1010031000211232-0201111211301221-1022213113131201-3112121231301333-2001320112113132"></a>

## Direct properties — low_security / 321302223130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012212023233020-0133203131003122-1102333120303232-2330301322010031-0022313230201323-1212331133003121-3302030231021032-2010320310322033"></a>

## Next pages — low_security / 321302223130 / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231023010201120-3012002213212120-3223000103333020-2202331101011031-3222233031323100-1303013110212133-2311233232030111-2231003020002013"></a>

## https_auto_cert.tls_config.medium_security — medium_security / 130100031011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.medium_security

<a id="canonical-0201130300131301-0013202212221303-0131133200011200-2310022310002311-3212332223112012-2331311203120202-1022110313303203-2113110320323011"></a>

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

<a id="canonical-2311201222313123-2110030002230313-1311303233313110-2311012003133333-0220333221210312-2310132103013203-0103331302302231-1100200220030212"></a>

## Direct properties — medium_security / 130100031011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123310221303330-1110322002131210-1001211332232033-0330013300323300-2022320133122322-0110213122301300-0001213002220130-1331302103300120"></a>

## Next pages — medium_security / 130100031011 / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322101310220102-2111000222222213-3113123330213121-0001111122300311-2033123330130023-1011203221220122-1313002232332121-0133101300121121"></a>

## https_auto_cert.use_mtls — use_mtls / 101321033112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.use_mtls

<a id="canonical-2311300312121330-0000303231320033-3211203320333001-0221203223003110-3103112313220101-1230001333223313-3202003200101201-1223332333221203"></a>

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

<a id="canonical-3311310303230321-0033000101132113-2002322120133130-1301033002212102-1300113301203132-3313222010231320-2303300332020001-0313132122330000"></a>

## Direct properties — use_mtls / 101321033112 / 3

<a id="canonical-0201023200032302-0302323332120111-0002222331301011-3122322300233201-0133311121022333-2303103230031033-1231312132021220-2220303103131101"></a>

<a id="canonical-0211212301233100-3231100101012310-0013013220233123-3223333101203211-1103030132122322-1310332220002320-1310113002202203-3121112212301021"></a>

## client_certificate_optional property — use_mtls / 101321033112 / 4

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

- [crl](resources--http_loadbalancer--reference--group-020.md#canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-020.md#canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-020.md#canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303): complete subsection reference.

<a id="canonical-3120331123332031-3333303223310123-1322133112210000-3031331210032330-0120320102000102-0023231210331030-1302010310211210-0030113301111000"></a>

<a id="canonical-1312322333022222-0010332331203121-0322332300123220-1202022213123312-2221000223312133-0320320121032100-1030133012232021-3102211030211130"></a>

## trusted_ca_url property — use_mtls / 101321033112 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-020.md#canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-020.md#canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110): complete subsection reference.

<a id="canonical-0011232230232000-0331003022300320-3321011033220212-2312023330311301-2300032022103002-1232020000000132-3003220220120021-3203020103221002"></a>

## Next pages — use_mtls / 101321033112 / 6

- [https_auto_cert.use_mtls.crl](resources--http_loadbalancer--reference--group-020.md#canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300)
- [https_auto_cert.use_mtls.no_crl](resources--http_loadbalancer--reference--group-020.md#canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022)
- [https_auto_cert.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-020.md#canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303)
- [https_auto_cert.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-020.md#canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032)
- [https_auto_cert.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-020.md#canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001031303332221-2130122210020131-2212232230223210-0310210013212303-1302312103033300-2001333100011100-3030203213200222-3003122130021310"></a>

## https_auto_cert.use_mtls.crl — crl / 202002021321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.crl

<a id="canonical-2130220123003332-3011211033010333-2003321000303101-0013030121103030-2021033201303122-0312101300313202-3201003333232301-2320302030310201"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0231032200013121-0033230211101132-3010012021230332-2020010110002002-2113020100232210-2233302210232312-0003222132330010-1223113021120230"></a>

## Direct properties — crl / 202002021321 / 3

<a id="canonical-1022302021023111-2023222221100000-3001102201131211-3100201331321132-0330302133310230-1200002220001001-1213021323110322-3033010021302233"></a>

<a id="canonical-0222332333120211-2032133330101112-2312323110032222-1221012302222032-3211213001022311-1031121100300113-1203201230303300-2330112200021131"></a>

## name property — crl / 202002021321 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2101313001311330-0330103313032121-3313323202313310-0233230013312022-3230132322130320-0002011033011132-2023201320002120-3111122313300312"></a>

<a id="canonical-2102223001230312-1022010303033331-0322302332331003-3000233210302310-3102322212100203-0112132223332033-1031113020031203-0021331132212310"></a>

## namespace property — crl / 202002021321 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1021332322113212-0331312301312011-0122133221301331-2213222201310321-2232301220111002-1113001332100120-1023313012323120-0230123233133011"></a>

<a id="canonical-0333102012201211-3121300103201002-0033033112302032-2002131001120111-0122020233333112-3122202311220302-2220000121132332-3313003032012332"></a>

## tenant property — crl / 202002021321 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1020311220302302-3303321323112210-3332202301313032-1020201003203110-2201100222303331-2132013303030231-3003220102032110-1230023100200213"></a>

## Next pages — crl / 202002021321 / 7

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312300230332131-3210212020300133-1132330201310030-3122102330030303-3301331033321122-3030131022312110-1331103012230331-2203103231021012"></a>

## https_auto_cert.use_mtls.no_crl — no_crl / 332301311313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-0013033311012301-1100112101011321-0300033202102030-3110021030101001-3231333032001011-3112322023223121-0213331013311021-1231222110100223"></a>

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

<a id="canonical-3003103003020323-3313002301023312-1110233102120301-2230100020233222-2230202130122230-3203030012022120-3012111023120332-1123132222300232"></a>

## Direct properties — no_crl / 332301311313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020102131330103-3112003110003132-2023022300012320-2313301132333030-3111033310133230-1330320010032323-2112201010020201-2330112030223330"></a>

## Next pages — no_crl / 332301311313 / 4

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010111133022120-0302231230003323-0133331301301033-3230103101120310-1222012001131210-2303223220230323-2220313013233012-0323103313033032"></a>

## https_auto_cert.use_mtls.trusted_ca — trusted_ca / 010222301233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-3201313232021021-2033101212131332-0101000120333313-2323310301323211-1131013133031003-2011132131022322-3321302030002301-1312102230102233"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2001011110201013-2212213313300331-0012132323001313-1300030112333001-2333032000301301-1011200030120003-3321312022203021-3231222103232331"></a>

## Direct properties — trusted_ca / 010222301233 / 3

<a id="canonical-2230101211212131-3200321203012220-1222322021231300-3023000201023021-3221302322300131-0101320102303020-3200021311313111-0212221030300133"></a>

<a id="canonical-0012322110022031-0232120313201103-3233202121130300-1201003221301230-0232223233233203-0233033130213021-2121013022221112-2230201210022301"></a>

## name property — trusted_ca / 010222301233 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1232311330133202-2213203011330210-0132030311213101-3033001031023333-0002323013323330-3103332003213022-0200300203333203-1212032213103302"></a>

<a id="canonical-2000021310223313-3022322131222300-2111032020233002-0132032220333100-0012103111211103-2120333301031323-1200002031321312-1023010020122113"></a>

## namespace property — trusted_ca / 010222301233 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3001021131220101-3303123013013031-2110231313220111-1000331101302103-3013213221021201-3033321221132032-1201321301332032-1103221312133301"></a>

<a id="canonical-1112201031012102-2100032212113203-0033213311301030-3133310301232220-2033033013131330-0120221031232133-0100321001100301-0333102112010302"></a>

## tenant property — trusted_ca / 010222301233 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3010102231021202-3311002322012203-0300020302202202-2012001102331100-3310032132132000-1230113201213112-3010120130233230-3221003203213031"></a>

## Next pages — trusted_ca / 010222301233 / 7

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031020223201301-3210031203202000-1100031031333010-1103022122233001-0222312200301323-2233113000313103-2102213020232101-2230211210330213"></a>

## https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 020020230110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3232020310031110-1031223133130020-3020032100101310-3020231123323112-1113303103032311-0031013310030113-0331310230030122-2022003110102122"></a>

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

<a id="canonical-2231300113112212-0012112220121202-0000332021312032-2013000113113220-0300202231202222-1111122303131021-1223003000303121-1031023201302001"></a>

## Direct properties — xfcc_disabled / 020020230110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331133323223021-2211121222131111-2320332000223311-0113111130213000-0312122321123322-0202102313113311-2012331103311331-3001020300012322"></a>

## Next pages — xfcc_disabled / 020020230110 / 4

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013311331211022-0122313121000221-1232012201200103-2002000012021130-2131012112222212-0131011330311321-2323211010133100-0111133300031322"></a>

## https_auto_cert.use_mtls.xfcc_options — xfcc_options / 222212133012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3312303100000321-0213221220310013-3320132300000302-2002022002003013-0232201032331021-1122033120112010-0302200220013131-0213302010332032"></a>

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

<a id="canonical-1230030223202222-2123111302002310-3132002111021102-3321310333313321-1323223131010222-1122120200100022-1032012133303202-3303121131113121"></a>

## Direct properties — xfcc_options / 222212133012 / 3

<a id="canonical-2303012132330203-2221121002213302-0223201011320131-2112012033300121-0130011223023300-3202101333220332-1221102202203103-2010102020100023"></a>

<a id="canonical-1322001102223203-0332300321020100-0213023132330021-1203122322001123-2022001332033311-1122121111211200-2211131000323303-2100131333231220"></a>

## xfcc_header_elements property — xfcc_options / 222212133012 / 4

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

<a id="canonical-1331013333000102-3130313023230112-0111223122302211-2302110202100301-3031133000302203-3131231011021030-3101101330133212-1122022220232200"></a>

## Next pages — xfcc_options / 222212133012 / 5

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0122213100023133-3102301113333103-2221320301323033-3302232320022112-0012001332132331-1022120132002212-0201313110302323-3320131033130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210212111030230-1212022231030101-2010302221133200-3031010311103230-3303100223102130-1313311032323231-1213233211210321-3322222130310321"></a>

## js_challenge — js_challenge / 003221011232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- js_challenge

<a id="canonical-1011111331300221-3213330020020313-1033022332302021-1111331213030331-0032003020022230-3000023321331133-2002032233231133-0233011032133100"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332123223022102-3310323103303022-2331323322102033-3230211210021003-0223112133121212-3213113332203212-3233321120130132-3332102332210322"></a>

## Direct properties — js_challenge / 003221011232 / 3

<a id="canonical-3032132210021332-3300311000232001-0232220012102222-2311003102301333-2102321332033110-0001302132210111-2113300310302130-3022121102323112"></a>

<a id="canonical-2221203130302023-0322031111201010-3200312303110001-2210123321120203-0020113011133220-0230233003001330-1310120132312311-2113130203013033"></a>

## cookie_expiry property — js_challenge / 003221011232 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0101103200132210-0120122330012132-3100330021112013-2002303232121031-1131003232003111-3302321333011231-3132101303002322-3201230010233003"></a>

<a id="canonical-1303032102102302-0302022333113112-0222332332300012-3123131200302220-1302312111031331-2221232113330210-1211002130110310-3230231120110102"></a>

## custom_page property — js_challenge / 003221011232 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2112211211300000-2003223012330332-2013201023130220-1303111122312102-0330012020111010-1102212333200113-2303210213101022-2321120220022002"></a>

<a id="canonical-0121113110003100-0300303311002203-2101302331103122-2000302321112021-3023223232301221-3321202031002010-1302221103121012-2231300332033223"></a>

## js_script_delay property — js_challenge / 003221011232 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1332012323020003-0100111110231031-3123311020313000-2330000033313200-1133213302312110-0003330021330122-0030103133311002-1221103311331300"></a>

## Next pages — js_challenge / 003221011232 / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031303210200111-2010021320231012-1233202322023100-2111230002121201-1120201011230022-3222301031111121-1120003330233033-2222221123222131"></a>

## jwt_validation — jwt_validation / 110211022223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- jwt_validation

<a id="canonical-3232231033300133-0123002110102230-1310030120201003-2321203312311032-2323123120301112-1300312222020221-1323200030221031-0110022312130022"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("authorization_server",
    "jwks_config")}
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
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333231121101223-3232002103113011-0133103131120230-2230222230101300-3212122130132220-1003330202321102-1332222022110003-1111011313333330"></a>

## Direct properties — jwt_validation / 110211022223 / 3

- [action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333): complete subsection reference.

- [authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310): complete subsection reference.

- [jwks_config](resources--http_loadbalancer--reference--group-020.md#canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221): complete subsection reference.

- [mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030): complete subsection reference.

- [reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222): complete subsection reference.

- [target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221): complete subsection reference.

- [token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120): complete subsection reference.

<a id="canonical-2211103233003020-2313021202130023-2212312221332110-2022220332211021-0201200032033220-1011210210323102-0201303110101222-3200301333122133"></a>

## Next pages — jwt_validation / 110211022223 / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310)
- [jwt_validation.jwks_config](resources--http_loadbalancer--reference--group-020.md#canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221)
- [jwt_validation.mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320133001300322-1303100311010220-3312310111322230-0021312332200103-0130133020100031-3311320001112003-1230033303003201-2211202111312313"></a>

## jwt_validation.action — action / 201221021323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.action

<a id="canonical-1312230111022013-2110100222223100-1202112333132220-3101322110023303-0210131312132213-2300313002201200-0301133130031330-3301213310111220"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313111202313122-1202010011212121-2131111112123021-3101222302330201-3111332313112110-3100333101022100-0030101110330100-1303210322100220"></a>

## Direct properties — action / 201221021323 / 3

- [block](resources--http_loadbalancer--reference--group-020.md#canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-020.md#canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301): complete subsection reference.

<a id="canonical-3100311001123320-0012112100132122-3210013002203321-3121020021202322-0133011123232021-1302203032223133-2013022011110003-2312022223202110"></a>

## Next pages — action / 201221021323 / 4

- [jwt_validation.action.block](resources--http_loadbalancer--reference--group-020.md#canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220)
- [jwt_validation.action.report](resources--http_loadbalancer--reference--group-020.md#canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001102323220110-1330213213220220-3231012312230101-2023331330300122-3323202231113302-1110233131300003-1230131322133021-3021011311231103"></a>

## jwt_validation.action.block — block / 122110223222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.block

<a id="canonical-1312210302123110-2323210031232032-2012000111011300-0231331012131101-2331220021201230-0300201313220333-1132322000112232-3300210102110320"></a>

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
block = {}
```

<a id="canonical-2230311002333231-1002002011322032-2101221123332131-2300123203103122-0133221021333302-0230302102323303-1000222132133003-3103313102100133"></a>

## Direct properties — block / 122110223222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330312203120302-2232131100203211-1022013302233210-2012012133302300-3201113310122330-0222200101301310-0010320123003011-2100103030223332"></a>

## Next pages — block / 122110223222 / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322131331130132-0302101320001110-2013003122011222-3030010002121002-1012331130113032-3222001221100113-3113320311320230-3312200301121333"></a>

## jwt_validation.action.report — report / 000213210001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.report

<a id="canonical-1110333320121232-2011200133213320-2132302111321203-1112111321311301-3123300231013311-1303220332210100-3311033300322002-0222333321303331"></a>

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
report = {}
```

<a id="canonical-0203113332330203-0221130311220001-3301310113311010-3232000331120322-0103021130110123-0313211330221231-0211011133002010-3103001101203010"></a>

## Direct properties — report / 000213210001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321330302220121-3122222002300201-3013133100312210-0102230222022012-0121002321022103-1220123220311321-0233103322033230-1332313023231111"></a>

## Next pages — report / 000213210001 / 4

- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121112323010212-3110200100210030-1103301320022220-1310131030002010-1222313011211223-0303012233323302-0000023332300311-3120100001220330"></a>

## jwt_validation.authorization_server — authorization_server / 010321201130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.authorization_server

<a id="canonical-2312131033133200-1002110221220223-0320033000332210-0103320210201311-1311323022322231-0331232102312121-3033110202310123-1313031201301020"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
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
authorization_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031311303333322-2132232221201020-3101032010332300-0001033213312121-2303330113012330-1002133010331101-2230133001110120-1000223302113112"></a>

## Direct properties — authorization_server / 010321201130 / 3

- [authorization_servers](resources--http_loadbalancer--reference--group-020.md#canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122): complete subsection reference.

<a id="canonical-0002221032110012-0133233221333120-3121310210331210-1010102330302321-1303022103033301-1302230110212131-3333132123221232-0222010330111133"></a>

## Next pages — authorization_server / 010321201130 / 4

- [jwt_validation.authorization_server.authorization_servers](resources--http_loadbalancer--reference--group-020.md#canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202312003202113-1030120323233321-2331131013112022-2312203231020001-3201230131123333-1220200221212222-3223103223303001-2113202312120202"></a>

## jwt_validation.authorization_server.authorization_servers — authorization_servers / 120013000301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-1131313331231012-1112000110030112-1320220301022233-0210023011220120-1332320203021203-0302231132031310-2333332232032113-3123310032000303"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Upstream description:

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
authorization_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111210302311002-0003130101220113-1101103031303312-1032211102003310-2012031022322230-2120200003011200-2121232222100002-3221220220202031"></a>

## Direct properties — authorization_servers / 120013000301 / 3

<a id="canonical-0331312002231320-1212331102222000-0111100333110002-3212231223022122-2223030023232020-1231210032002132-3200220303202113-0032311233023011"></a>

<a id="canonical-3013321110202021-0032200310012130-1132230002203223-0103312303332301-3022022200223331-1113133330113110-3000032302231230-0023311012220213"></a>

## name property — authorization_servers / 120013000301 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3300130003020321-2021000110200122-1123021313010321-3320112303112011-3012022300030310-2323323231123211-2011132001320211-2000331312220120"></a>

<a id="canonical-1131320322323312-0002313010012012-1120212322331310-2110011100032100-0220003210130313-1210110003012323-1222113032003133-1211230130203302"></a>

## namespace property — authorization_servers / 120013000301 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0111133121303211-1201122020121010-3210030222311000-0232021211222023-0333032313333333-1320121200123302-0311102132233033-1121023013120232"></a>

<a id="canonical-1310222231131013-0301210300311121-0233131322200201-1202300213332323-1133031212300233-0323210120123011-0122101333212312-0111131130321133"></a>

## tenant property — authorization_servers / 120013000301 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1122330322322333-1002223020203111-1013200212100230-3302313122220333-3012103333121210-3032301020313132-0213202031002000-3111311320133023"></a>

## Next pages — authorization_servers / 120013000301 / 7

- [jwt_validation.authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003301120023220-1303231233112102-1323320022333322-0120112322031313-1112321202031223-1202201313220303-2120213300323120-2010333102133300"></a>

## jwt_validation.jwks_config — jwks_config / 311013111103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.jwks_config

<a id="canonical-0212201230302231-0210300010032010-3322300020112131-1313032133011021-1121111131300211-0003313032231310-0311011031132020-0203003131022003"></a>

Type: `"object"`. single nested block, Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
jwks_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220203031230012-2232330033212331-2103123022330230-3201300312002301-3101332200030201-1322220002202212-1223013323130211-0222022113302311"></a>

## Direct properties — jwks_config / 311013111103 / 3

<a id="canonical-2323213101210100-2200301001300110-3111223023032003-1313330323001301-3303131012103233-2030021221202223-1312211030033201-1233302111232130"></a>

<a id="canonical-0002320313112212-1121131011110022-0021013312330331-0310133033012121-3212030231022020-0200313300230333-3022302130223032-0000202111223033"></a>

## cleartext property — jwks_config / 311013111103 / 4

Type: `"string"`. Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="canonical-0302011100210302-1032021033100212-1132223023033212-3101201010300100-3232333102003010-3011330332030210-0200312102322313-2322223102212103"></a>

## Next pages — jwks_config / 311013111103 / 5

- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000213231012221-3100223102110101-1112103211001203-3010213032033102-2312211032200302-0123313011212303-1131102301033021-0022131132231202"></a>

## jwt_validation.mandatory_claims — mandatory_claims / 231032123111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.mandatory_claims

<a id="canonical-1123110330302331-3321302021330023-1012003122201101-0201333230323021-2223002322300331-3012323113210122-1203230200221033-2100020333312130"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of mandatory Claims.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
mandatory_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002030301313033-1121320210121202-0003130120001312-2220220203132300-0221330032303300-1300321330130303-3220001003313021-1032110223122021"></a>

## Direct properties — mandatory_claims / 231032123111 / 3

<a id="canonical-0321231312330202-2212101003111100-2030332201212103-0320223101203032-2311021321233201-2110022112130331-2332133332133102-2320320223130220"></a>

<a id="canonical-0000311013001313-0023120121031201-1301332012322211-2112000220201310-3301321120100000-1202003123213300-3302102032133113-2102130100222000"></a>

## claim_names property — mandatory_claims / 231032123111 / 4

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022123000203213-2003321233203101-1111110103012032-0230332023221223-3220020321211300-2030213320201121-3123012120012032-3302303130000331"></a>

## Next pages — mandatory_claims / 231032123111 / 5

- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032023032321212-2012031323211202-3011133003112013-0313113211020011-1111300032020012-0133300211212001-2022120023223222-2333310322312110"></a>

## jwt_validation.reserved_claims — reserved_claims / 130121201233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.reserved_claims

<a id="canonical-2012031101213202-3131103003211202-2031333000203220-1012202130112203-0220322132003212-3332212202130231-0311010033333121-2121322020201032"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("audience",
    "audience_disable"),
  validators.ConflictingObjectAttributes("issuer",
    "issuer_disable"),
  validators.ConflictingObjectAttributes("validate_period_disable",
    "validate_period_enable")}
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
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

Terraform syntax:

```terraform
reserved_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023321010221003-1013100221113321-0112321023211223-0203222221333021-0310321322102203-3223213022231122-0113301230110133-0011021133131313"></a>

## Direct properties — reserved_claims / 130121201233 / 3

- [audience](resources--http_loadbalancer--reference--group-020.md#canonical-2300303331000122-3121101203310331-1101010002013011-3212310311202230-1120312322133332-0103002110110220-2011100122032132-2222021332021310): complete subsection reference.

- [audience_disable](resources--http_loadbalancer--reference--group-020.md#canonical-2223212220013230-3111321103313220-0222122120332033-0010012000300010-2232103201330012-0212200021021033-2213220202003330-0010230213010222): complete subsection reference.

<a id="canonical-0133300020002310-0112322033132310-3323120231311313-1303020100031303-0012203112320310-2002332313201131-1300100100012120-1210101123203310"></a>

<a id="canonical-3211131210310312-0210102302302313-3100132111313310-0201011021010120-1222311221012130-2110010123211322-2330011131203212-3002113121000120"></a>

## issuer property — reserved_claims / 130121201233 / 4

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

Upstream description:

Exclusive with \[issuer\_disable\]

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

- [issuer_disable](resources--http_loadbalancer--reference--group-020.md#canonical-0020321310110001-0231212322321023-2210012110110103-2220201210230313-1001330321023313-0123021031323331-0131330022220103-0012131111000002): complete subsection reference.

- [validate_period_disable](resources--http_loadbalancer--reference--group-020.md#canonical-1212310211203322-2132113320213210-0032103232222000-0003112132223212-3111011312130223-1122212110003000-1023011020311220-0032333301011220): complete subsection reference.

- [validate_period_enable](resources--http_loadbalancer--reference--group-020.md#canonical-1023300203001213-2322122200323300-1330123303122313-0100003200130110-3300321100201032-2123103033121103-0322122213033102-1212003013030233): complete subsection reference.

<a id="canonical-3332300233322202-3313022122100220-1102010232320013-1012122300202220-3103303303213121-2202110331120213-2131023202033322-2201311333201310"></a>

## Next pages — reserved_claims / 130121201233 / 5

- [jwt_validation.reserved_claims.audience](resources--http_loadbalancer--reference--group-020.md#canonical-2300303331000122-3121101203310331-1101010002013011-3212310311202230-1120312322133332-0103002110110220-2011100122032132-2222021332021310)
- [jwt_validation.reserved_claims.audience_disable](resources--http_loadbalancer--reference--group-020.md#canonical-2223212220013230-3111321103313220-0222122120332033-0010012000300010-2232103201330012-0212200021021033-2213220202003330-0010230213010222)
- [jwt_validation.reserved_claims.issuer_disable](resources--http_loadbalancer--reference--group-020.md#canonical-0020321310110001-0231212322321023-2210012110110103-2220201210230313-1001330321023313-0123021031323331-0131330022220103-0012131111000002)
- [jwt_validation.reserved_claims.validate_period_disable](resources--http_loadbalancer--reference--group-020.md#canonical-1212310211203322-2132113320213210-0032103232222000-0003112132223212-3111011312130223-1122212110003000-1023011020311220-0032333301011220)
- [jwt_validation.reserved_claims.validate_period_enable](resources--http_loadbalancer--reference--group-020.md#canonical-1023300203001213-2322122200323300-1330123303122313-0100003200130110-3300321100201032-2123103033121103-0322122213033102-1212003013030233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300303331000122-3121101203310331-1101010002013011-3212310311202230-1120312322133332-0103002110110220-2011100122032132-2222021332021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302302321200010-0120213130100200-2210030010022201-3300133122200003-1132322021203312-1322303332330131-2320221331312033-1110321011213223"></a>

## jwt_validation.reserved_claims.audience — audience / 121210232103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.audience

<a id="canonical-0303221123212101-1321312210210330-2212003130023321-1210001122100002-2323023203331223-2031112210021322-2331131321130013-3211101320313323"></a>

Type: `"object"`. single nested block, Optional.

Audiences

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("audiences")}
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
audience {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013012201220011-2021312202010310-3023012022120312-3013010130200323-0010113101202100-1020112332100131-3010303211302021-3011001120000333"></a>

## Direct properties — audience / 121210232103 / 3

<a id="canonical-2022031120030233-1010313220013012-3110013101021231-1031333003211333-0313232223202033-1121230132012131-0212030011112032-0110130111002001"></a>

<a id="canonical-0222120232032221-0223003121002020-2131133010033210-0111020320313023-0020313211133112-1002202111000222-0201213311213000-0033230023133311"></a>

## audiences property — audience / 121210232103 / 4

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300030311322330-3010011222112233-1033112330312310-3131030012313101-2332010233010321-3010022022010333-3000023122130221-0322133230231100"></a>

## Next pages — audience / 121210232103 / 5

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2223212220013230-3111321103313220-0222122120332033-0010012000300010-2232103201330012-0212200021021033-2213220202003330-0010230213010222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213001302323000-2303322023101022-3330321300312021-3330031120223212-1021231110132123-3031120332233003-2031320120323011-1230101300103000"></a>

## jwt_validation.reserved_claims.audience_disable — audience_disable / 001130331323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-3212210311213201-0122130111203312-2202302021032112-2313022230320110-3333330131123111-0202323111200310-2300113213301330-2132332120023001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

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
audience_disable = {}
```

<a id="canonical-1202100201003013-0320311222313330-1121200032031321-2210322222222312-1331301213233032-3133020133220001-0201000011000332-3222322310022203"></a>

## Direct properties — audience_disable / 001130331323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230020122030033-2023333220330013-0010222201332130-1231200223233321-3112230300103310-2111131000210121-2223003222331223-0122001223112020"></a>

## Next pages — audience_disable / 001130331323 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020321310110001-0231212322321023-2210012110110103-2220201210230313-1001330321023313-0123021031323331-0131330022220103-0012131111000002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321203113010001-3011100232201102-0333012223311121-3123013111333110-3301000113100020-3011112030011323-3133112230232113-0021303132222332"></a>

## jwt_validation.reserved_claims.issuer_disable — issuer_disable / 013302023331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-1210322203203110-2330001213300031-0201132302122232-2221322202211233-0300022030230131-2023230003312132-0012323003020012-0121103212330223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

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
issuer_disable = {}
```

<a id="canonical-2313232133011230-2033120023130012-3101303312320230-1112332023010120-3331003031222201-0012221103003312-3310000230202332-1200102031322010"></a>

## Direct properties — issuer_disable / 013302023331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220031311200321-3103120000102133-2220322001003020-1222202112322312-0111223310300321-0200111332103133-1201201211230103-3103233233220023"></a>

## Next pages — issuer_disable / 013302023331 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1212310211203322-2132113320213210-0032103232222000-0003112132223212-3111011312130223-1122212110003000-1023011020311220-0032333301011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221011322012001-3321320331013020-2222023101000212-1111211210233233-2022112301130301-0311122311332032-1002230320313120-2320321330233331"></a>

## jwt_validation.reserved_claims.validate_period_disable — validate_period_disable / 010331222111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-0113232213113110-1003232330313312-3103100310220103-3011233131010033-2303233101020100-1321120032013313-3013000202111033-0232020322333012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

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
validate_period_disable = {}
```

<a id="canonical-3012222320113032-1102313010330213-0001301132130032-1212201003233123-2100201203001232-0001002301333133-1312300330330201-2100022101033001"></a>

## Direct properties — validate_period_disable / 010331222111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123330032220223-3123331120012310-1000211230130312-3130001033212213-0222222201033012-2112032010101031-2122332233230211-2111213233033013"></a>

## Next pages — validate_period_disable / 010331222111 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1023300203001213-2322122200323300-1330123303122313-0100003200130110-3300321100201032-2123103033121103-0322122213033102-1212003013030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123232331200100-2033101103223301-0311222320202313-1331312012120032-1133002003221011-1200200123011110-2303102123313112-1303110210313233"></a>

## jwt_validation.reserved_claims.validate_period_enable — validate_period_enable / 331101230213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-0021020103322113-0021232223323323-3232131213211023-0131031303313321-1302021230200121-3313300320333222-0200111202330223-2332322000232310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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
validate_period_enable = {}
```

<a id="canonical-1032202300122300-0220022303011331-2112000122211203-3312320222132011-2003130032222002-1322120323133123-1123123201130220-0230032021132010"></a>

## Direct properties — validate_period_enable / 331101230213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101213123311331-3222131131003321-3000022220232122-3132023012200001-3103031023133110-0122311011322133-1022110102122103-2223222112003232"></a>

## Next pages — validate_period_enable / 331101230213 / 4

- [jwt_validation.reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000211333001111-3103220211322313-0330101103330311-1223100312111113-2202330023011202-0031223311300022-2030121022213230-0100121101132332"></a>

## jwt_validation.target — target / 332123122112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.target

<a id="canonical-3330011301202031-0013332012111323-2102010332303011-2313131321311302-2131130112220321-1132010120311113-1022013002132311-0302230301032132"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_endpoint",
    "api_groups"),
  validators.ConflictingObjectAttributes("all_endpoint",
    "base_paths"),
  validators.ConflictingObjectAttributes("api_groups",
    "base_paths")}
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
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

Terraform syntax:

```terraform
target {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013221323311313-3333021013131013-0332203220120031-2310321302101102-1113212020123312-2002331311311101-2323133233202330-2022133223113223"></a>

## Direct properties — target / 332123122112 / 3

- [all_endpoint](resources--http_loadbalancer--reference--group-020.md#canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211): complete subsection reference.

- [api_groups](resources--http_loadbalancer--reference--group-020.md#canonical-1222112033113103-0212223103001302-3322022133003223-2323202221310130-2133022333111311-2320312303300010-2023130011233101-3133212112101102): complete subsection reference.

- [base_paths](resources--http_loadbalancer--reference--group-020.md#canonical-1211200313211220-2030022210023120-0330203001332121-0220311123001323-2112313323311233-2013210013013230-3202001132002122-1233232212322110): complete subsection reference.

<a id="canonical-0313133233212300-0221302103310232-2020033322112102-3020333133123210-3231111321231032-0022232131133022-2111122121320002-2223020231103331"></a>

## Next pages — target / 332123122112 / 4

- [jwt_validation.target.all_endpoint](resources--http_loadbalancer--reference--group-020.md#canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211)
- [jwt_validation.target.api_groups](resources--http_loadbalancer--reference--group-020.md#canonical-1222112033113103-0212223103001302-3322022133003223-2323202221310130-2133022333111311-2320312303300010-2023130011233101-3133212112101102)
- [jwt_validation.target.base_paths](resources--http_loadbalancer--reference--group-020.md#canonical-1211200313211220-2030022210023120-0330203001332121-0220311123001323-2112313323311233-2013210013013230-3202001132002122-1233232212322110)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1020320031023321-0320202313120203-1020333130031033-0231120033031022-3001322011013132-1031131331321201-2333121233322230-3033322013201211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010130233131320-0003023031110321-0012302030231300-0332210000001211-0020112230013222-2211132223111333-1301012222103323-3023033322202122"></a>

## jwt_validation.target.all_endpoint — all_endpoint / 011323213313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- jwt_validation.target.all_endpoint

<a id="canonical-3021120023122321-0213210011132330-0022200203222301-3333110121020311-1010131201232030-3111302321013302-3312003223023131-0320132011202331"></a>

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
all_endpoint = {}
```

<a id="canonical-2001211130112310-2222123220221300-3201131101113322-3301330220033122-0111110121323301-0302301030220301-2003011012221230-0210313013203001"></a>

## Direct properties — all_endpoint / 011323213313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023030111030102-2311301123111111-1313203000222112-0221112211203101-2002300222101323-3012023030231131-0001331311310100-3203113323023331"></a>

## Next pages — all_endpoint / 011323213313 / 4

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1222112033113103-0212223103001302-3322022133003223-2323202221310130-2133022333111311-2320312303300010-2023130011233101-3133212112101102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033200220313201-2320213113121333-1022022001333130-0002100103132112-0033003312021321-2233201010033120-1203013122221123-3323211331020132"></a>

## jwt_validation.target.api_groups — api_groups / 320233310300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- jwt_validation.target.api_groups

<a id="canonical-3311233101020321-3002223100030313-1221311332220213-0130122220030230-0033121321001303-1003300232213223-2020303020031232-2131112223303231"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003030133111100-3020130000223321-3003132011033130-3120102310231030-3111210231031013-2121001233001032-2121132103211022-0223223032301120"></a>

## Direct properties — api_groups / 320233310300 / 3

<a id="canonical-0201030213123102-0213312032123202-2100121100332101-2001020212003133-3130113313132332-3003132332202302-1310321010320112-0021212133001101"></a>

<a id="canonical-3103131233020212-2123023101003202-2112010010302311-0232033330303110-3230200132322321-3301100101303002-3221020101202301-3301323332202220"></a>

## api_groups property — api_groups / 320233310300 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232102033102022-0032020012200032-1310230021322233-3102023223320321-3310123112320022-1211203233102012-3010011322010222-3022122233021222"></a>

## Next pages — api_groups / 320233310300 / 5

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211200313211220-2030022210023120-0330203001332121-0220311123001323-2112313323311233-2013210013013230-3202001132002122-1233232212322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221122302313122-1222301310133002-2303322213012322-1301032001330000-1300212033123223-3130332123210000-2122210001120101-2002000212231210"></a>

## jwt_validation.target.base_paths — base_paths / 233303033202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- jwt_validation.target.base_paths

<a id="canonical-3000130122102012-3210121222023023-2200201131202233-0012213310201332-0302321320121100-2103211000201332-0113221133211020-3033131113301233"></a>

Type: `"object"`. single nested block, Optional.

Base Paths.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("base_paths")}
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
base_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200013110011302-3100333302232121-2023310032313012-3100012110131323-2211223031100031-1001233231203113-0032312202133112-0312011101221022"></a>

## Direct properties — base_paths / 233303033202 / 3

<a id="canonical-3223211203303010-0010123313313033-3333221001101111-0001122103123310-2233102023131131-0211333311010133-2230301102023021-1022311310302301"></a>

<a id="canonical-0132313013103000-3030301011310221-0023030113323203-0201222132110210-3031102122111301-1202031033000011-2022212210203311-3202023110103223"></a>

## base_paths property — base_paths / 233303033202 / 4

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2332013031000331-2232133201102131-0210000333110231-1010212131213131-1212110230100322-1220133001012003-2312032033001213-0112123123231203"></a>

## Next pages — base_paths / 233303033202 / 5

- [jwt_validation.target](resources--http_loadbalancer--reference--group-020.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202220000300021-0311100310210101-3310320131301122-0120103200310301-0203202330311122-3020003301021232-3003110031020031-2222221110233110"></a>

## jwt_validation.token_location — token_location / 200320011202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.token_location

<a id="canonical-0110122110101020-2213022213101220-3001021132321201-1101100011110303-3110023220212110-3302233031320030-0133200002223203-2231132023300230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122020131211020-0120330223232000-2331130021333000-2311212323323201-2103210200330300-0130013011103032-0203031300332321-2203213132221103"></a>

## Direct properties — token_location / 200320011202 / 3

- [bearer_token](resources--http_loadbalancer--reference--group-020.md#canonical-3321120100031033-3323212200212200-0123023031330003-1203231013131122-2011301000100231-3131211333010001-3010002000331213-3110110300331302): complete subsection reference.

<a id="canonical-2231203021201312-1101212030220303-2033102100212210-3300203320131213-2101123313213102-2233101122212330-3210120012013130-2302233303123121"></a>

## Next pages — token_location / 200320011202 / 4

- [jwt_validation.token_location.bearer_token](resources--http_loadbalancer--reference--group-020.md#canonical-3321120100031033-3323212200212200-0123023031330003-1203231013131122-2011301000100231-3131211333010001-3010002000331213-3110110300331302)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3321120100031033-3323212200212200-0123023031330003-1203231013131122-2011301000100231-3131211333010001-3010002000331213-3110110300331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312200113131003-2320133333001131-3231230002023021-0303103012230101-0220303011110021-0302301033112013-2311300300002002-2221300002121203"></a>

## jwt_validation.token_location.bearer_token — bearer_token / 331231322030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120)
- jwt_validation.token_location.bearer_token

<a id="canonical-0211221203032321-1212111002310033-1001220110022211-1001330230131100-2020103120323113-1020311300200203-3213000230222222-1030331313231012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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
bearer_token {}
```

<a id="canonical-0223330103203212-3320223331231203-1113222322001122-0010023031021013-1303321032102020-0000201320221123-3223023302322300-0300232232021301"></a>

## Direct properties — bearer_token / 331231322030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001232302000011-1031101233333322-2232213310302020-0223123033002112-1200300230321321-3320110111310222-2103322003201001-2010323021211110"></a>

## Next pages — bearer_token / 331231322030 / 4

- [jwt_validation.token_location](resources--http_loadbalancer--reference--group-020.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1320110132020001-0022333031222332-0332001223201230-0000302202010121-2213203130001312-3221132020131213-3301310301230123-2220221132033300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103033011332120-1022321003112122-1112030101132223-1122002331333001-1103302310230133-3110322203123202-3331011000033313-3222213213132011"></a>

## l7_ddos_action_block — l7_ddos_action_block / 220121103212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_block

<a id="canonical-0033232131220130-3322300300000223-1323101032222302-2111321001222301-0312212303003130-3002120021323011-2320000031121121-3130133311130202"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](resources--http_loadbalancer--reference--group-020.md#canonical-0033232131220130-3322300300000223-1323101032222302-2111321001222301-0312212303003130-3002120021323011-2320000031121121-3130133311130202)
- [l7_ddos_action_default](resources--http_loadbalancer--reference--group-020.md#canonical-0310222221220320-0030213003223012-0003112031003100-3000300312023111-1011003103012311-3231220231213221-2220232130310122-0120103032031120)
- [l7_ddos_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-3201213003032333-3301120031030220-0131331332333232-1133310130122013-1002021113313300-0022133200331212-2232013002123020-1201313313200322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

<a id="canonical-3233012232132202-2231213002131013-2023310103032212-2113021230202110-0212232332210202-2001021203303220-1000013333231033-3330210132031320"></a>

## Direct properties — l7_ddos_action_block / 220121103212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300302131210002-0131300010030112-1031302022303302-3111021021012002-2023200210101013-2310303232000133-1102102030023321-3131210320201332"></a>

## Next pages — l7_ddos_action_block / 220121103212 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122303101123332-1003322121300112-3112102023031013-1001133230202001-1121101002133313-0121023032001313-1200303023002210-0310210303322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022102220001310-3302002101311131-2320102313302003-0112113000112221-2121213133022201-2303102301120101-1232221322013023-1113010101131131"></a>

## l7_ddos_action_default — l7_ddos_action_default / 202003120012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_default

<a id="canonical-0310222221220320-0030213003223012-0003112031003100-3000300312023111-1011003103012311-3231220231213221-2220232130310122-0120103032031120"></a>

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
l7_ddos_action_default = {}
```

<a id="canonical-3113210233103213-3000202300301010-3132010203020312-1210223131322311-0213033223221133-2312202122320012-0212122220121223-3012010320112220"></a>

## Direct properties — l7_ddos_action_default / 202003120012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211203320331000-2003203320102002-1123201010023310-2330101321321203-0333211002001031-2032033100203123-3220020301102011-3013023212132202"></a>

## Next pages — l7_ddos_action_default / 202003120012 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130022200313200-0012221203233323-2031132210123022-0122213131223223-0033223300233010-3030023013112322-3321030210121223-0130102321222031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313122011032321-1332310221022022-0330302212112302-0201332330113031-2232133313030223-1030113003020012-1220332202011202-1123032122012303"></a>

## l7_ddos_action_js_challenge — l7_ddos_action_js_challenge / 001123312310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- l7_ddos_action_js_challenge

<a id="canonical-3201213003032333-3301120031030220-0131331332333232-1133310130122013-1002021113313300-0022133200331212-2232013002123020-1201313313200322"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
l7_ddos_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221310110312322-0131302002123110-0321131022310312-0312331210232021-3213022310000123-1222102033031100-0232133222022310-0102231212303220"></a>

## Direct properties — l7_ddos_action_js_challenge / 001123312310 / 3

<a id="canonical-3131023030010021-2210303302333303-0031031011112301-0003211203002000-3201031203101110-2021130310210002-2110222203101022-0232212001213110"></a>

<a id="canonical-3101120320330203-1202322132032123-0123131303211230-0020112113031132-2130203233001220-2021101010210200-3030100331020211-3011013213012123"></a>

## cookie_expiry property — l7_ddos_action_js_challenge / 001123312310 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-2233230030312121-1013201232133133-0112131203311002-3030200210000101-1331013333232003-3023220310320112-1301332121213311-1313021013112213"></a>

<a id="canonical-0221023210310103-0013010301113323-3120003123123101-1002001000213011-3300331011333033-1202123030002333-1021122302221313-1231023032203301"></a>

## custom_page property — l7_ddos_action_js_challenge / 001123312310 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1013022110121332-2011211120010123-3103032203022330-0202233000033112-1323030300001033-1022010231133330-3010212311023232-3112212323303110"></a>
