---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0222010121130210-2230331320103033-1003331213102030-1231111322120121-1013331102203000-3000212313333123-3110020103321020-0000222200021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.crl

<a id="canonical-2331030020221031-1231123300213211-1012322130331331-0101122012120301-2003231113020203-3221230223233303-3332100111121312-3102300203320223"></a>

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

<a id="canonical-0302122203021023-0010322301021102-3303011001110013-2103011002312131-2210003120033302-2131113310230033-1303101023330123-1112002022301110"></a>

### Direct properties for `https.tls_parameters.use_mtls.crl`

<a id="canonical-2212122201022003-2210020203120233-0132023213233133-1313101202002320-3203031031321323-1130200013021203-1100322100022103-0123123031112023"></a>

#### `https.tls_parameters.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2011322131222112-3321132200030332-1312322011110303-3032022232111231-0323213130102233-3130302131213123-3313300012332021-1001112223310022"></a>

<a id="canonical-2303011120122132-3330113223011031-2332333012130030-2221002111023031-0221220113213233-3231000133022101-0011303203011311-3100021223113212"></a>

#### `https.tls_parameters.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3332010102322023-1130112031100211-3133131210210121-2032033211200310-3032212313032231-0201030223320303-3003023312331123-0002301120311201"></a>

<a id="canonical-1222312011220331-0130210021200112-1130013213322021-1120123101102213-3332030101131300-3302102001130233-3022201232202230-1011313023201021"></a>

#### `https.tls_parameters.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1030111002012120-3333122220103311-3320230001320012-1232020311120002-0232303302022021-1113121110000110-0020100233000213-2102200103102111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-1230002332301133-0313013212112020-0323120312100111-0310012323202031-1112132103230200-1321203313230310-0213113231131302-1231323322211312"></a>

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

<a id="canonical-2322231221033132-3033011200232320-1122002210223211-3311212321020032-2112201023110101-2023322002222213-0110121233231212-1322321203331110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2122333012202002-3210202000122111-2111210002202102-1002320330333210-2200110010013233-2310133101231233-0221320203011313-2221020100310333"></a>

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

<a id="canonical-0121212130323023-0101122011111203-3321231021310302-0020303121102233-1021001202033022-0120331301101131-2330211002201122-3220320130202313"></a>

### Direct properties for `https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-1223133011301133-0120233310020313-0122030003032231-1123032031101223-0100001212120012-1103311333101000-1310020032121123-1331131122303102"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0320210211100310-0111100032003203-0331133020302300-3323012030223033-2230112121030323-2232031030030130-2310002100023013-1301110320321130"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1122323132322330-3013112213022010-2321111233122113-3210120223203223-0301311330121031-3022001310010323-2303001332100010-2023322230223011"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2322002130002320-1330120021101020-1200130232312020-0133311022301011-3202331222300110-3322232232122000-1021003111002231-2232030010220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2001330313213112-2003121033010131-1233202331001132-1031030001011100-3332012002200003-0310121321203212-1121303033022021-1113231103133003"></a>

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

<a id="canonical-0133133030323002-3012021101322010-0220200122033323-3222120033110132-3121211121220001-0131113203131111-0020130103231221-1131302303211333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-1231113013303201-3233222000300212-2211131130013302-2322211003010221-0320233103301130-2102232111212102-2321212333113100-3130203302022212)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-1321332101331311-1023112313030210-2312030203110123-1202033022020030-0101101100102110-2022201111233012-3210333032232312-3311120222312230)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2300200302102222-2312213000310030-2331231223103023-2013230302023320-0232001013322012-0121010030132101-2222100220202321-2100231300102313"></a>

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

<a id="canonical-3101123200330222-1033130332223212-3110001021211222-1213003121333011-1203121131123100-3013300211021023-2020331333032112-1233113110132022"></a>

### Direct properties for `https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-2103202323130012-0131122233030220-0321022231131220-2122013031002230-1220021330011210-2003111113210102-0111003120222223-3102120320233331"></a>

#### `https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- https_auto_cert

<a id="canonical-1223123300112032-0222133020221001-3201331012031023-0133020113303300-1120222301233102-1120232121100332-2220301031031210-0023330100002100"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

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

<a id="canonical-2212233130303102-0211032023003220-1330010201111021-2030111110113301-1120330113022112-1020032101331231-1112030202303213-3111303311000323"></a>

### Direct properties for `https_auto_cert`

<a id="canonical-3223321310121123-2301223021013011-2111300211130232-0122000300102301-3333333133201231-1010132220131130-1303010200233013-3010322133001333"></a>

#### `https_auto_cert.add_hsts` property

Type: `"bool"`. Optional, Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1001100232130231-1001110203232212-1212101300330122-1121301331231033-1233310320200030-2212221012103010-3323200113132311-1002022122031230"></a>

#### `https_auto_cert.append_server_name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0111333230002111-2101133003323202-3320031233202301-2133313001213221-0211332112211300-3220122102021323-2000023111002210-1131331202132303"></a>

#### `https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Additional upstream details:

Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203333030333211-2003132223020311-0012330203101012-1002133001002122-0223210201300201-2031032303122031-0311220231032233-1321323100230303"></a>

#### `https_auto_cert.http_redirect` property

Type: `"bool"`. Optional, Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1101201322323321-2300331220002202-3322312331221202-3000011301320220-3120021212122122-0120221201112032-0323211122212330-1012331321212230"></a>

#### `https_auto_cert.port` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1232133303302003-2013103311313102-3311202011100320-2311011220323231-3211223231323203-3122021132100011-2230323232233121-2033220230130301"></a>

#### `https_auto_cert.port_ranges` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0113300213233320-1030330110032300-3231102113312232-1103333030330320-0232020010021002-3030320213032111-2323132022102032-1122212002232333"></a>

#### `https_auto_cert.server_name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.coalescing_options

<a id="canonical-3001231212332201-0313132220012130-0030022212322000-3112123233103301-0003021133311302-3323021100230320-1022133233231222-0310213323221033"></a>

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

<a id="canonical-0120322221033330-0132332223122213-2133022113102222-0233222200330113-0310231213200031-1302001100101323-1200032323031203-3220320202013031"></a>

### Direct properties for `https_auto_cert.coalescing_options`

- [default_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-020.md#canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033): complete subsection reference.

<a id="canonical-3311102022220123-1330323002111233-0312000201022002-0303301222233222-2120113311313320-1203233212222320-2233310233102210-0110101121022300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1101100331131002-3203323132211211-1313320203333302-1133030220030012-2331011320023013-0313322303300333-3231113222300000-1013220231003331"></a>

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

<a id="canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-020.md#canonical-2320110111023121-1030011122221033-2223122301021032-1210022123121010-3201011210020220-1301333221333211-3200002223112333-1332023101103103)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3200023331202132-2301112331203032-1112102033020122-0211031100203321-3101130023332222-2300210313110212-2111302020301211-1202323210220222"></a>

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

<a id="canonical-2113022110102112-1011013213132131-0223312310010333-2123303311003111-1223330331230103-1131321102330330-3101022002113023-3102133322022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_header

<a id="canonical-2100311123231131-0133311110202100-2332023030121232-0023223011001303-2031232033201002-3302130331132312-3030110102022311-2200021303333233"></a>

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

<a id="canonical-2310002132131131-2223213002312301-1111021023011313-3200130033301120-2302233110201031-0130100200312202-2011013130032133-3231300330203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.default_loadbalancer

<a id="canonical-3001000020233233-3001321322312021-2322123012222021-0133201022333112-3323333211203103-0330302331102301-1232203033312201-0302322313113110"></a>

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

<a id="canonical-0110302231311201-1000101323303201-0311032011203113-0121202021212111-2200212000022211-3222111100133033-3222222222301302-3210030312200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.disable_path_normalize

<a id="canonical-0333213122002122-2030101001233332-2123333323000121-1110122231213323-3030331031030011-1223332223313103-3230211212032010-1302311321022120"></a>

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

<a id="canonical-2310221300113300-3121200223230020-0330303032331230-2212300030130223-3011211311022112-0320032111303201-3100313311202101-1312323302330131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.enable_path_normalize

<a id="canonical-2223013223201010-3303013330023003-3300321112131202-0222013313100103-1320203101220202-2130020300130020-0313322332022230-2101230131000111"></a>

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
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.http_protocol_options

<a id="canonical-2011113132131102-1110223121331230-3310333210200303-2221103221112122-2231100322130011-2202011222323313-3020013102332103-0013323210332230"></a>

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

<a id="canonical-1201211232312221-0313321232012333-1321101010111011-1203230113200033-2002232122233211-2321021130321220-1010202121310210-3210003303021102"></a>

### Direct properties for `https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-020.md#canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-020.md#canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-020.md#canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002): complete subsection reference.

<a id="canonical-1130221103222100-1130010220032331-3312221100201222-0011331013220033-3302131032302202-2311130110321232-3112010210322221-3113103330210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

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

<a id="canonical-1302013103010301-3000331031202230-2303132032222032-0210212012210320-1031122101112212-3233032003303323-1230131321331311-3200200020313232"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223): complete subsection reference.

<a id="canonical-1013321322330012-3032313200312222-3231313032322330-2212030033322012-1331010010103102-1130221301130020-0222130233000011-2030113332100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

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

<a id="canonical-2020010330230120-1310110000020231-0010231003133211-0032103231320200-2001212200313222-1123220211231123-3303121200220020-3122213212303211"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-020.md#canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013): complete subsection reference.

<a id="canonical-1111333113221110-1311210232302022-1021332121102320-3021331120301230-1113312012233323-0203120202120312-1120203112130123-1132222202012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312122212023330-0232021131232132-1020130112133013-3003030123001020-3011321232321002-2101330312323222-1333312300111200-3101230233131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202300322132111-1110000030230211-3203200132321310-1220110113103030-2213301323301132-2310221103012111-3211221003023032-1313311000202013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032210133003231-3032132203012020-1303023210001022-2303222230012301-1100232311012011-2001203131022033-2310210000212122-2010111323123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2123230010031123-1020303333310032-3131112012111011-2333133011021320-2203123130002302-1312012022032231-1321203033101131-1323131121133302"></a>

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

<a id="canonical-0210331120231100-1310212233122110-2120131132011221-3211121110031123-3300300031033202-2100222100231302-3331130100023231-1220010221201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-020.md#canonical-2221233103301012-1100320102122132-1231130032133121-1332333103312332-2331021123003200-0213032321013301-3112202102011331-0220300233200122)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-3122203002332132-0123021301103122-3103030331002102-2300331311323121-2330030312032202-2120020210212101-2133132222031032-1220013010123022"></a>

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

<a id="canonical-1210120101233222-0023010312023312-0330303012201233-3233301232011011-0120033202101122-1132300330001120-3230110110030212-3200212233021332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.no_mtls

<a id="canonical-2220332020001020-2302321002301012-2311000132102222-1122221231100012-3312231011131012-2303331110220020-3002122100311031-1211320332103221"></a>

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
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102022210321000-1330022312123302-0203001131320130-0102013232033211-0202012012300202-3032100100201010-3302333323021200-1010231100131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-1001221102303011-2220021100220012-0230032100202011-1301332010210020-2030022302012303-3302231211303103-2211120212331103-2333133230111310"></a>

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

<a id="canonical-2131303122323030-0001331033213121-0333112221110231-0133000022320231-3211302113323332-2301333223033313-2232110302101320-3202300101332312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.pass_through

<a id="canonical-3311221012012030-3213230332021221-3212331303133020-0333212210303111-1122021210220031-2131202102133211-0213003033202211-0220210023210101"></a>

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

<a id="canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.tls_config

<a id="canonical-0233120200312332-1211103300220120-2332300211130200-2113122123122012-0231332031100303-1132222200303331-2330333330121112-0033311121331021"></a>

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

<a id="canonical-3013311133002230-1233330322211320-1012013032323222-2313301320102232-3000101021131210-3102331332322313-2111020210230230-0131012132300222"></a>

### Direct properties for `https_auto_cert.tls_config`

- [custom_security](resources--http_loadbalancer--reference--group-020.md#canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-020.md#canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-020.md#canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-020.md#canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320): complete subsection reference.

<a id="canonical-1200020020221222-1310122212212011-3101132221330213-2202331222221212-2021331100112233-2130302020213131-3021332310221032-2113132221113220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.custom_security

<a id="canonical-1113313033212323-1021103010103233-0322232311101112-3022122122120223-1311212203012101-1101220012302303-1313212022223100-2101310330031212"></a>

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

<a id="canonical-2310000212323301-3230210312213230-2011211322322332-1221001332030311-2232000032310022-0212112112012220-0331301220221133-0332123122121013"></a>

### Direct properties for `https_auto_cert.tls_config.custom_security`

<a id="canonical-2011223032011232-3122112313013030-2330002032002232-0332023111103022-2300103322120111-2110102321111022-0232020310121312-1102332032220030"></a>

#### `https_auto_cert.tls_config.custom_security.cipher_suites` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1020310330320330-1010210023230002-1310312023201223-1331032301201201-0321033003002223-2032010000101200-0203301221322132-0002310233322332"></a>

#### `https_auto_cert.tls_config.custom_security.max_version` property

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

<a id="canonical-0101311222222313-3100120110113331-1313221201011301-2011123033301001-3311331222011023-3012311323321133-1230012122223321-2311222101031312"></a>

<a id="canonical-3303213202012220-1332013112221031-2003021112013323-3123131101110012-0123022012123202-0211211312220132-2202321320231003-2000213010020213"></a>

#### `https_auto_cert.tls_config.custom_security.min_version` property

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

<a id="canonical-3223202003303223-2033112010233130-1300011333021230-0033120002012013-3203210120313021-2202032331002232-3011212300023201-3222120202213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.default_security

<a id="canonical-2202001103333021-2002220110001221-2202203001012330-3023320211102210-0300012131312102-1113032320203000-0211230312211022-1101212230013110"></a>

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

<a id="canonical-3101310223111320-3303103201323333-3332221200132012-0002132210230210-1310200102002300-0222131300103312-0200121111013023-2112333101021312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.low_security

<a id="canonical-0201122001030021-2100130002313122-2332212021333030-1110221320113022-0032232203231303-3113312031202010-3100323021020020-2213332210201103"></a>

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

<a id="canonical-3011231211020313-3200233133300122-1003113222023330-0020322123321333-3113110211322030-1111200310331301-0301123320220231-2020233321202320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-020.md#canonical-3023222011130302-1021031110023033-0023230113010303-3211301301032003-0000213031223032-3021010301122222-1031233002032001-2130201233320022)
- https_auto_cert.tls_config.medium_security

<a id="canonical-0201130300131301-0013202212221303-0131133200011200-2310022310002311-3212332223112012-2331311203120202-1022110313303203-2113110320323011"></a>

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

<a id="canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- https_auto_cert.use_mtls

<a id="canonical-2311300312121330-0000303231320033-3211203320333001-0221203223003110-3103112313220101-1230001333223313-3202003200101201-1223332333221203"></a>

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

<a id="canonical-2322101310220102-2111000222222213-3113123330213121-0001111122300311-2033123330130023-1011203221220122-1313002232332121-0133101300121121"></a>

### Direct properties for `https_auto_cert.use_mtls`

<a id="canonical-0201023200032302-0302323332120111-0002222331301011-3122322300233201-0133311121022333-2303103230031033-1231312132021220-2220303103131101"></a>

#### `https_auto_cert.use_mtls.client_certificate_optional` property

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

- [crl](resources--http_loadbalancer--reference--group-020.md#canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-020.md#canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-020.md#canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303): complete subsection reference.

<a id="canonical-3120331123332031-3333303223310123-1322133112210000-3031331210032330-0120320102000102-0023231210331030-1302010310211210-0030113301111000"></a>

<a id="canonical-3311310303230321-0033000101132113-2002322120133130-1301033002212102-1300113301203132-3313222010231320-2303300332020001-0313132122330000"></a>

#### `https_auto_cert.use_mtls.trusted_ca_url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1210133002013033-1103131000011220-2302030011330102-0300132202211303-1122223022303332-1332003033321233-1202213102001221-2022011123112300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.crl

<a id="canonical-2130220123003332-3011211033010333-2003321000303101-0013030121103030-2021033201303122-0312101300313202-3201003333232301-2320302030310201"></a>

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

<a id="canonical-1001031303332221-2130122210020131-2212232230223210-0310210013212303-1302312103033300-2001333100011100-3030203213200222-3003122130021310"></a>

### Direct properties for `https_auto_cert.use_mtls.crl`

<a id="canonical-1022302021023111-2023222221100000-3001102201131211-3100201331321132-0330302133310230-1200002220001001-1213021323110322-3033010021302233"></a>

#### `https_auto_cert.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0231032200013121-0033230211101132-3010012021230332-2020010110002002-2113020100232210-2233302210232312-0003222132330010-1223113021120230"></a>

#### `https_auto_cert.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0222332333120211-2032133330101112-2312323110032222-1221012302222032-3211213001022311-1031121100300113-1203201230303300-2330112200021131"></a>

#### `https_auto_cert.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0312200200300303-2013120220223101-1310330133001210-1221232100221012-3300223122023203-2233331112113111-2213102311310310-1232222023020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-0013033311012301-1100112101011321-0300033202102030-3110021030101001-3231333032001011-3112322023223121-0213331013311021-1231222110100223"></a>

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

<a id="canonical-0122332201002231-1031200230212110-3332000013303332-0221012200223100-1231022113301023-0230211001322303-1200133000103101-3001130102330303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-3201313232021021-2033101212131332-0101000120333313-2323310301323211-1131013133031003-2011132131022322-3321302030002301-1312102230102233"></a>

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

<a id="canonical-0010111133022120-0302231230003323-0133331301301033-3230103101120310-1222012001131210-2303223220230323-2220313013233012-0323103313033032"></a>

### Direct properties for `https_auto_cert.use_mtls.trusted_ca`

<a id="canonical-2230101211212131-3200321203012220-1222322021231300-3023000201023021-3221302322300131-0101320102303020-3200021311313111-0212221030300133"></a>

#### `https_auto_cert.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2001011110201013-2212213313300331-0012132323001313-1300030112333001-2333032000301301-1011200030120003-3321312022203021-3231222103232331"></a>

#### `https_auto_cert.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0012322110022031-0232120313201103-3233202121130300-1201003221301230-0232223233233203-0233033130213021-2121013022221112-2230201210022301"></a>

#### `https_auto_cert.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3203223313321313-1310203230020321-0010112031231321-0112210212013021-1121232000031112-1321201310001132-2202200001103323-0101210330332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3232020310031110-1031223133130020-3020032100101310-3020231123323112-1113303103032311-0031013310030113-0331310230030122-2022003110102122"></a>

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

<a id="canonical-3130113131012000-1210112132131310-0033003132301102-3111330302032002-2211330203322100-1002022211011122-3023010020113102-1100322222003110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [https_auto_cert](resources--http_loadbalancer--reference--group-020.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-020.md#canonical-2123232203322302-1202212121011223-2010012122210331-2211330300002203-3330002222233012-0223332221210232-2320210231110120-2130233121221131)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-3312303100000321-0213221220310013-3320132300000302-2002022002003013-0232201032331021-1122033120112010-0302200220013131-0213302010332032"></a>

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

<a id="canonical-2013311331211022-0122313121000221-1232012201200103-2002000012021130-2131012112222212-0131011330311321-2323211010133100-0111133300031322"></a>

### Direct properties for `https_auto_cert.use_mtls.xfcc_options`

<a id="canonical-2303012132330203-2221121002213302-0223201011320131-2112012033300121-0130011223023300-3202101333220332-1221102202203103-2010102020100023"></a>

#### `https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` property

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

<a id="canonical-0122213100023133-3102301113333103-2221320301323033-3302232320022112-0012001332132331-1022120132002212-0201313110302323-3320131033130221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- js_challenge

<a id="canonical-1011111331300221-3213330020020313-1033022332302021-1111331213030331-0032003020022230-3000023321331133-2002032233231133-0233011032133100"></a>

Type: `"object"`. single nested block, Optional.

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

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-1210212111030230-1212022231030101-2010302221133200-3031010311103230-3303100223102130-1313311032323231-1213233211210321-3322222130310321"></a>

### Direct properties for `js_challenge`

<a id="canonical-3032132210021332-3300311000232001-0232220012102222-2311003102301333-2102321332033110-0001302132210111-2113300310302130-3022121102323112"></a>

#### `js_challenge.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1332123223022102-3310323103303022-2331323322102033-3230211210021003-0223112133121212-3213113332203212-3233321120130132-3332102332210322"></a>

#### `js_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2221203130302023-0322031111201010-3200312303110001-2210123321120203-0020113011133220-0230233003001330-1310120132312311-2113130203013033"></a>

#### `js_challenge.js_script_delay` property

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- jwt_validation

<a id="canonical-3232231033300133-0123002110102230-1310030120201003-2321203312311032-2323123120301112-1300312222020221-1323200030221031-0110022312130022"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

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

<a id="canonical-0031303210200111-2010021320231012-1233202322023100-2111230002121201-1120201011230022-3222301031111121-1120003330233033-2222221123222131"></a>

### Direct properties for `jwt_validation`

- [action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333): complete subsection reference.

- [authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310): complete subsection reference.

- [jwks_config](resources--http_loadbalancer--reference--group-020.md#canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221): complete subsection reference.

- [mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030): complete subsection reference.

- [reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222): complete subsection reference.

- [target](resources--http_loadbalancer--reference--group-021.md#canonical-3203010010312213-3220111213032103-0132012333011222-1022230022121330-0102222111013230-1112030311103110-2131210033203332-1122013310120221): complete subsection reference.

- [token_location](resources--http_loadbalancer--reference--group-021.md#canonical-2100130323133112-2010031310200203-1331011110302222-2122110331203111-1230200312133021-0102032323223320-3132123010020310-1130031012332120): complete subsection reference.

<a id="canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.action

<a id="canonical-1312230111022013-2110100222223100-1202112333132220-3101322110023303-0210131312132213-2300313002201200-0301133130031330-3301213310111220"></a>

Type: `"object"`. single nested block, Optional.

Action

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

<a id="canonical-2320133001300322-1303100311010220-3312310111322230-0021312332200103-0130133020100031-3311320001112003-1230033303003201-2211202111312313"></a>

### Direct properties for `jwt_validation.action`

- [block](resources--http_loadbalancer--reference--group-020.md#canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-020.md#canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301): complete subsection reference.

<a id="canonical-2011211312131013-2130223032323023-1102121322310103-3221230021031333-3313000330212202-3100203312323010-0231111022202333-3010202300333220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.block

<a id="canonical-1312210302123110-2323210031232032-2012000111011300-0231331012131101-2331220021201230-0300201313220333-1132322000112232-3300210102110320"></a>

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
block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133210212101102-3012122320001322-0203222020200310-3333003211010032-1211031310102232-3313130212032001-2033201121231101-3312010031220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.action.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- [jwt_validation.action](resources--http_loadbalancer--reference--group-020.md#canonical-2321020001303302-1010303302133003-2132222300302133-0101111210023030-1010200323302010-0333222210103130-2303201122032012-2333102130111333)
- jwt_validation.action.report

<a id="canonical-1110333320121232-2011200133213320-2132302111321203-1112111321311301-3123300231013311-1303220332210100-3311033300322002-0222333321303331"></a>

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
report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022003113111113-0032002130102133-2220130121300000-0122202022230311-2300211112222121-1303223203111120-1213002002313001-1311023131121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.authorization_server

<a id="canonical-2312131033133200-1002110221220223-0320033000332210-0103320210201311-1311323022322231-0331232102312121-3033110202310123-1313031201301020"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
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

<a id="canonical-0121112323010212-3110200100210030-1103301320022220-1310131030002010-1222313011211223-0303012233323302-0000023332300311-3120100001220330"></a>

### Direct properties for `jwt_validation.authorization_server`

- [authorization_servers](resources--http_loadbalancer--reference--group-020.md#canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122): complete subsection reference.

<a id="canonical-2312313023300223-0233212301100131-0001232313021131-3030310002010121-1301020231203223-1321201011310323-2113211301232133-1101023131330122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.authorization_server.authorization_servers` properties

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

<a id="canonical-0202312003202113-1030120323233321-2331131013112022-2312203231020001-3201230131123333-1220200221212222-3223103223303001-2113202312120202"></a>

### Direct properties for `jwt_validation.authorization_server.authorization_servers`

<a id="canonical-0331312002231320-1212331102222000-0111100333110002-3212231223022122-2223030023232020-1231210032002132-3200220303202113-0032311233023011"></a>

#### `jwt_validation.authorization_server.authorization_servers.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0111210302311002-0003130101220113-1101103031303312-1032211102003310-2012031022322230-2120200003011200-2121232222100002-3221220220202031"></a>

#### `jwt_validation.authorization_server.authorization_servers.namespace` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3013321110202021-0032200310012130-1132230002203223-0103312303332301-3022022200223331-1113133330113110-3000032302231230-0023311012220213"></a>

#### `jwt_validation.authorization_server.authorization_servers.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0200132101022323-3022233122203202-0231000123332030-2202231321102000-3220301013030332-1301102133332120-1312321003301313-3320311010330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.jwks_config` properties

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

<a id="canonical-3003301120023220-1303231233112102-1323320022333322-0120112322031313-1112321202031223-1202201313220303-2120213300323120-2010333102133300"></a>

### Direct properties for `jwt_validation.jwks_config`

<a id="canonical-2323213101210100-2200301001300110-3111223023032003-1313330323001301-3303131012103233-2030021221202223-1312211030033201-1233302111232130"></a>

#### `jwt_validation.jwks_config.cleartext` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3021210032222001-3201003210113000-3310200202320013-3233011322102311-1213220210003332-1232213333010022-3011313231331002-1232320210131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.mandatory_claims` properties

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

<a id="canonical-0000213231012221-3100223102110101-1112103211001203-3010213032033102-2312211032200302-0123313011212303-1131102301033021-0022131132231202"></a>

### Direct properties for `jwt_validation.mandatory_claims`

<a id="canonical-0321231312330202-2212101003111100-2030332201212103-0320223101203032-2311021321233201-2110022112130331-2332133332133102-2320320223130220"></a>

#### `jwt_validation.mandatory_claims.claim_names` property

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0322002001030322-2030222000232333-2221321013013010-0231231101303331-0021222031123230-2201102331302120-1012123020213302-3022000100011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `jwt_validation.reserved_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [jwt_validation](resources--http_loadbalancer--reference--group-020.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333)
- jwt_validation.reserved_claims

<a id="canonical-2012031101213202-3131103003211202-2031333000203220-1012202130112203-0220322132003212-3332212202130231-0311010033333121-2121322020201032"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

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
