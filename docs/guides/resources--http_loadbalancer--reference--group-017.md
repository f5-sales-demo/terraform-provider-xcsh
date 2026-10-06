---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-016.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-016.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0133000112321230-1110222120112232-0333021301203130-3323020310312330-0023121000131003-1102101110023131-3330122333222230-0223021310123001"></a>

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

<a id="canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-0022112112130111-3011023102110131-0333031030310303-2030022132112002-3133330123223031-1132202201000033-0130031010113333-0033302321210101"></a>

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231122113302332-2002303233330021-0332111332303312-3023103303132330-2332103233301201-3220121103302022-1120330231012032-0320120231010122"></a>

### Direct properties for `default_pool.use_tls.use_mtls_obj`

<a id="canonical-0121211103102032-0330313001322120-1112233233332320-0130021120223001-0203302301023132-1303301312221132-1221121130032311-2110223112313130"></a>

#### `default_pool.use_tls.use_mtls_obj.name` property

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

<a id="canonical-0123332020332101-0232133120232030-2000000233320132-1320110222321033-0220020120231321-0321213213001001-3211301002323303-1330020100320110"></a>

<a id="canonical-0302202001310110-1101211010303021-1032211223322000-1232121312203003-3332330010311011-0321132021331220-0003133013313231-2013033123223330"></a>

#### `default_pool.use_tls.use_mtls_obj.namespace` property

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

<a id="canonical-2020133233123002-2322012220031112-1200111211102032-2113130302302200-2010102122302000-0310102312220002-2030202200222210-3221003330003323"></a>

<a id="canonical-3102322103321232-0122203111103223-2102232232331100-2032233303011301-2022033303310230-3002023012311031-3011112002202211-1301201333322133"></a>

#### `default_pool.use_tls.use_mtls_obj.tenant` property

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

<a id="canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_server_verification

<a id="canonical-2130103020232302-1031121000213331-1200111111332320-1123033212132121-1232002101101333-3211110303301331-3210133303122301-3102331023033332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332002310031003-3001331101220003-3212223023322323-0121333111101322-0000211211232121-2311233003220121-1032333303122010-2031313303131031"></a>

### Direct properties for `default_pool.use_tls.use_server_verification`

- [trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000): complete subsection reference.

<a id="canonical-3302121201002223-1132313201012120-1122021012131202-2321102221122231-2311231033221310-1033030303311121-2313113220020103-2230223112333331"></a>

<a id="canonical-3321330102033021-2122301031113312-0311012031222230-2121230320311030-1222222133230230-1300113301122133-0311122210322212-2203322332333311"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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

<a id="canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-1231200013223113-3313130001203221-0300003313303101-2321020102203333-3002301211023301-0111230320300103-1111010320123121-3333222302202002"></a>

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

<a id="canonical-1231001132223201-3000000230221201-3230301222103130-0111320221302321-0030221130133001-3003331233013102-2011232310231012-3302310102230223"></a>

### Direct properties for `default_pool.use_tls.use_server_verification.trusted_ca`

<a id="canonical-3213210320020231-0003132323301200-3232213232232323-3302123312201310-2032212230132002-1333320233001021-1313033203123313-3203312333330102"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.name` property

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

<a id="canonical-3202102011303222-0022021021120200-3313011233333120-0022133020230332-3023001112110013-3123131200213303-0211321023130233-1331322321231010"></a>

<a id="canonical-1222321332201113-1102203230332110-0233103200220002-0001321230302221-1032233212332323-3021332200322330-1000210222110302-3200312033033032"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.namespace` property

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

<a id="canonical-1211213032201123-1101110211320311-3002013002111020-2023201011122120-3122022312332313-0231013113112212-0332301133201321-2203322031021012"></a>

<a id="canonical-0133100203011322-1032013323121132-2012222331222131-1332113030212203-1100112231213000-3020003322203031-2223200131311022-3122323211010102"></a>

#### `default_pool.use_tls.use_server_verification.trusted_ca.tenant` property

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

<a id="canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-016.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-1123131102231031-3303002311113322-2223231122012003-1221030332331211-3223111220030330-1032310301010110-3231233101113000-2110012102333313"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool.view_internal` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.view_internal

<a id="canonical-3111311300001101-1221223112021201-0221010001311121-2322133311013113-0202010001021132-0303200203011113-2102003220222010-2213223113000030"></a>

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
view_internal {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332302222233002-1022201332120301-1130021200330033-2321122302112121-2312313112303302-2001301003102032-1032000212211130-3031303000121202"></a>

### Direct properties for `default_pool.view_internal`

<a id="canonical-2330201021132320-2202103302033221-1223110301030212-1210130012230213-3210132102120002-1001323103002030-2220033103133223-2131111310322330"></a>

#### `default_pool.view_internal.name` property

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

<a id="canonical-1021013023031201-3020331300312210-1023213331112123-0032310012033303-0212302203133210-0123200110130131-2001122112230101-1023220200023222"></a>

<a id="canonical-0321012230120312-3030331101102322-0213222012012323-3121111131332033-0211023011120020-1101001302000322-3030132201202212-0132023333312133"></a>

#### `default_pool.view_internal.namespace` property

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

<a id="canonical-3300212223020031-1202212011022000-2031210322333130-2033121132021000-3230132232123100-2211220301323322-0321113101221012-1130213312120220"></a>

<a id="canonical-2133223222112030-1320112311020112-1322022131230120-1221330231310223-1330010330210302-1133120002210301-3120230310123131-2110031003121202"></a>

#### `default_pool.view_internal.tenant` property

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

<a id="canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_pool_list

<a id="canonical-0130032320003112-3013231322132011-2010113111012220-3312013033110232-0201300301333010-2200230132012313-2122033102033333-3101112132302330"></a>

Type: `"object"`. single nested block, Optional.

Origin Pool List Type. List of Origin Pools.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_pool_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333223003212333-2032110013333332-0132202330213022-0021102200222000-1122122323130323-2112303013332332-0001213202032131-0210212101121222"></a>

### Direct properties for `default_pool_list`

- [pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112): complete subsection reference.

<a id="canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- default_pool_list.pools

<a id="canonical-1130133122303033-3113103332212302-2223111002322122-3312003023012013-1321033110333130-2213233123210322-2330020110210033-0103212032210123"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320223320331321-3221011123300323-3302000000033320-3113130111031300-0323033222220323-1321310012301331-3231033022120221-3220032030010132"></a>

### Direct properties for `default_pool_list.pools`

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130): complete subsection reference.

<a id="canonical-3220013120332110-2330332100201203-3201012203133303-2311130120312132-1233130333001130-1120103003303110-1120202112311201-0132333330222330"></a>

<a id="canonical-0113203302233112-1001122031303022-2313321023320320-1032313021123112-2012022312020033-1310112020012032-3232321102213121-0001122123222010"></a>

#### `default_pool_list.pools.priority` property

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0213013110313133-0003022203030213-3033000010320023-2113231020313113-0020330133323230-3203123222100322-3031010223100220-0020323202001132"></a>

<a id="canonical-3300133133230030-1101000330222113-2132032031211202-3110133033312000-2120232030002232-0201312300133321-2303201000032001-1231230012322112"></a>

#### `default_pool_list.pools.weight` property

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.cluster

<a id="canonical-2022032001300210-0100113302101222-3222032123312312-2103201330210120-2101032102101110-2020300132231333-1131210122301102-1320322013203231"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032300223120013-0121112210213021-1031132233222331-2200211321102301-0102313223032300-2022010200130031-2322220132210302-0102100311111312"></a>

### Direct properties for `default_pool_list.pools.cluster`

<a id="canonical-0311122122232003-0001330212231032-0323030213332321-0012221211302122-2003033001030210-3023310203222202-1121310000101322-0002111102232213"></a>

#### `default_pool_list.pools.cluster.name` property

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

<a id="canonical-0031320031332122-3030010133313031-0033331201001223-2330020223303233-3211232012020303-1100302133331102-2221200102013223-1333202223010221"></a>

<a id="canonical-0010331000011002-2223303223033332-1233332311032113-2323323113320301-1303111002111311-2301302300012322-2221203212303231-2000331110020212"></a>

#### `default_pool_list.pools.cluster.namespace` property

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

<a id="canonical-2203132230103200-3030001133021013-0313332320330322-3221131102121132-2203131102233131-3211013012332223-2010111123133012-0311030233332021"></a>

<a id="canonical-1310302033320232-0021130020221332-2222331200333301-0323033233332122-2100000030123033-1012012110201121-1103233123033122-3320211021012331"></a>

#### `default_pool_list.pools.cluster.tenant` property

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

<a id="canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-1323010032123123-0013103233303002-1223231222330131-1111223120220111-2332113311112011-3202211000212220-3321213312023021-0302111200223122"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_pool_list.pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.pool

<a id="canonical-1321233303001203-2111130000110322-0211222220122211-2123023121132032-1032320012321220-0131303212011332-1101301023112013-0003013300110320"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302123202312111-3301030131020023-3202233131312000-2222221101000120-1322230133012023-1002033133020302-1320031011320232-0301210013321231"></a>

### Direct properties for `default_pool_list.pools.pool`

<a id="canonical-1103222010120231-0100200233110002-1220013113330203-3333210122000023-1221312202221212-3032031003201223-3302103010103000-3223001323303211"></a>

#### `default_pool_list.pools.pool.name` property

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

<a id="canonical-2100233002120313-3201030132212123-0133331233101330-3022112132212221-0231332211031121-2323002013110211-0113131301203220-0330011120102133"></a>

<a id="canonical-1133000313130031-0212102100203203-3120023023113011-2023121032133012-3022021222032220-3333020030002311-0120222221330223-1132232031010011"></a>

#### `default_pool_list.pools.pool.namespace` property

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

<a id="canonical-3310123313033112-2332030321211130-2022230010133023-2313000120231121-2332002201223320-0113030131130032-2310231032122020-3222230320203322"></a>

<a id="canonical-3130310303233111-1113032000331331-3331100000103113-1002121331300021-0123210033021010-3033100122310201-1303132121000203-0021211101322100"></a>

#### `default_pool_list.pools.pool.tenant` property

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

<a id="canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_route_pools

<a id="canonical-2030002000113130-1013213300122022-0332132120103002-3113212002131130-0011121312031211-0300001031000233-1333033103010131-3022331121021300"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools used when no route is specified (default route).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_route_pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111211333211322-2011031300210312-2200120222113021-3221311301201202-1320123221120321-1300131112132100-1012310023303323-1021123303322101"></a>

### Direct properties for `default_route_pools`

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000): complete subsection reference.

<a id="canonical-0330313201303120-1120121031310101-2210120002330330-1110232011202121-2330200231230010-1130003330010010-3120322130131301-3033000121220021"></a>

<a id="canonical-2323132220313230-2322132032303220-1033113033001003-3121011010333303-0231332223302213-1011112021203130-2132313321122023-1113221203202302"></a>

#### `default_route_pools.priority` property

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3133330301132011-3110200033223102-3003102020133323-0222120231200131-3033032131111222-3032200323331021-2100103232101121-0022232200012203"></a>

<a id="canonical-3200132213031302-1111031322331322-2013231313021023-0313312213302002-2210222033311202-1310000121302312-1130033202230230-3131313330010313"></a>

#### `default_route_pools.weight` property

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022310213023222-3210202333110033-3221111331321321-0023021200321202-3231012111100012-3231303220210100-0203203303300122-2202110232103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.cluster` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.cluster

<a id="canonical-2002221002211212-1223131201203133-2220103133110301-1121220022322222-2023301320102130-2132122333112203-0113111212121101-2311200332220212"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022133022132112-0332300221313020-3010002331322323-3200212330312322-0321223003220202-2021202131300321-2300022010221231-2322223101012130"></a>

### Direct properties for `default_route_pools.cluster`

<a id="canonical-2321130310233212-3011120101131101-0122311321202322-1021012232310321-1313330103231211-2212121312003322-1331100331302010-3111232302131313"></a>

#### `default_route_pools.cluster.name` property

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

<a id="canonical-2132021200000231-0203102312233100-3021313333113112-2202332213211223-3101300232322222-0332332222011122-2000003003323323-0303033001301231"></a>

<a id="canonical-2001222232323231-2133210032020233-0023010132120010-3230333223102321-1023113322211122-3301201121012313-2311103023030331-0322311201031202"></a>

#### `default_route_pools.cluster.namespace` property

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

<a id="canonical-2102211032310201-1130032233102003-1021203102210223-1112333023310123-1133012101230322-1323311102023132-1023102312213322-3133110103122021"></a>

<a id="canonical-3123223223301132-0120001023122331-0320112213011320-2231022223300100-0300323311311232-3022032202021312-2330312021311023-3130011331020102"></a>

#### `default_route_pools.cluster.tenant` property

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

<a id="canonical-2122023231002110-2012202011030013-1132311102333232-0002323211230012-0321001232122003-1133213213111110-1221003221320031-2001301333112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.endpoint_subsets

<a id="canonical-0112120313012032-0001303222003221-0211201103201130-2122120211000011-0323110300131202-0213103210100230-1120231223120312-1102301031011133"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023320030321121-2112102000220202-0120020303101031-3101321002131113-3302203333030331-0031013322012221-1013012120003122-0122223312101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_route_pools.pool` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133)
- default_route_pools.pool

<a id="canonical-1203032021222111-2113022103331021-1300131313111112-1232202130003023-0110330312210131-0133132312322110-0333230121100002-2101201013033011"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103203221110330-3232222320030333-1032132220311011-1322222231230311-3110003113321132-3330333133133332-1201120132203133-2021103200002021"></a>

### Direct properties for `default_route_pools.pool`

<a id="canonical-3122123130000212-3202202332332100-0012310000222010-0231000023303120-1110023001033323-3221023310100320-1222311003133300-3201200231001333"></a>

#### `default_route_pools.pool.name` property

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

<a id="canonical-1311111200130213-0111212202112123-3003121023302121-1220103230113311-2121002110301013-1320222032032311-1123111310200212-0021023223231212"></a>

<a id="canonical-3210301233203132-3220201221301321-2002111111022233-0013313332123000-2013103312202312-2122231312023020-0133332110302232-0001320100003312"></a>

#### `default_route_pools.pool.namespace` property

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

<a id="canonical-1322332320222033-3223230221233020-0112021211033100-2100203221010300-3103122211123110-3312121002123322-1222131123201132-2213003001211233"></a>

<a id="canonical-1333200201100110-3223023202112330-1011111322210113-0101130301102003-1203322112211122-2323323223213030-3222301132333102-1021333002331002"></a>

#### `default_route_pools.pool.tenant` property

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

<a id="canonical-2012113030000132-3012100211231221-3130003303031233-0223013133101100-2030232213320012-3201201203311111-0130303112131112-1022213320131303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_sensitive_data_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_sensitive_data_policy

<a id="canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-017.md#canonical-2033022123213123-2323113031001012-2201021021002231-3221313211000130-1202101233120223-2020021110211031-2103021121302100-2001220200021113)
- [sensitive_data_policy](resources--http_loadbalancer--reference--group-027.md#canonical-2331322013011313-3111202121321023-1113002312200000-0132100333303032-2311313300200313-3003330023023213-0133322032323223-3212012023320222)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123111133030011-2123003110002011-1031032130100110-1022210200133330-3303330101000301-0203011300213011-1020303111133311-2132323123000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_definition` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_definition

<a id="canonical-0121212322222130-3103021101033321-1123000003031210-3103120000300132-1002112302102003-3020032021001211-1330203332020133-1300330021233313"></a>

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
disable_api_definition = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232001202112132-3122220002210323-2123112210120123-2111123031002213-0132120133332022-2111330310202011-2220013001013012-0203213300003333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_discovery

<a id="canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3000030203212331-0031103221201012-0231200113003012-3001100132332313-3331212332021012-3322030122330202-3001012130221221-0001222023213222)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313202103132030-0013213023113203-2231222203010002-3132330201211211-1031002112330010-2320302022023023-1021321113132320-2221311103102213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_api_testing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_api_testing

<a id="canonical-2102333133332200-1121021111010220-1122302223121000-2100012203322001-3000221313331130-1113021300210303-2123021303330312-2201022000213330"></a>

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
disable_api_testing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123003300002010-0033323131033110-3101003330221223-1232122203203021-0032123103302012-3321133121110332-3110131110231100-2002011203003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_bot_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_bot_defense

<a id="canonical-0031322221201123-0210020223332001-1212312133213033-3030023312222333-1332200121023202-1033110032330120-1010322121131033-3331133222011133"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable bot defense. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
disable_bot_defense = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321101032330312-2021001302001033-0303012101301113-3221010231232001-2031231033003121-0102030231321121-2103303220300112-2032023320022021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_caching` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_caching

<a id="canonical-2310010330303022-3212102030230223-3311321133320300-3300013001030002-1102021010321320-0321310020220321-0011203222103312-1122212300121322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable caching.

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
disable_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122321323123120-0131330211013033-0223101332002110-1133012200030110-0300010032111010-3133032130220300-0032001323100111-3300101020203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_client_side_defense` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_client_side_defense

<a id="canonical-1313000133203032-2310331201223313-3311201202200022-2303113031201211-1010211132031221-2313112333213303-2001211021211300-1211030023323302"></a>

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
disable_client_side_defense = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202122330212221-3201331122203123-3033330230333010-1333110133101010-1102330300133111-1223133210101012-2101013113121223-0330312222221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ip_reputation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_ip_reputation

<a id="canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-017.md#canonical-1300101232320303-3120302303031113-3023013133003230-2333310323023013-1321321023310323-0310123021012021-2232302023102200-1011120203321300)
- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-1331031122011033-2302131121012122-3313303233203321-2313300032330123-2331111321311122-1221001122320101-3033133302202020-0230023201131132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012231223303232-2102332203003323-3102212131101310-2012202210031310-2221201202003333-1132132220302101-0321000202112303-1321313331130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malicious_user_detection

<a id="canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.
Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-017.md#canonical-0131101221211313-1111300023003030-2102333222020213-1220113103132330-2110321310303202-1230232021202133-2200013231320023-2210202210021231)
- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-1200132303110020-2331223220030122-3300330122310330-2033201303222103-2311130312023020-2001011122320233-2033303101211233-0313110100132203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303333110313121-3231031113100003-0101230011321330-0011203322101330-0102021301032021-1333303102122201-3123112221102102-3010103301030101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_malware_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_malware_protection

<a id="canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_malware\_protection, malware\_protection\_settings; Default:
disable\_malware\_protection\] Configuration parameter for disable malware protection. Defaults to
\`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_malware_protection](resources--http_loadbalancer--reference--group-017.md#canonical-2102310133222323-3303112320121111-1102223222210312-0130011112231230-2320133120213321-1320112231313022-3302310213131033-2003201033332322)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2131132110220032-0231001231223002-1030113021221223-1220220323132330-0320302100310030-1022231121013021-1003021302003300-2223310300300221)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malware_protection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010203201310323-0211110203000202-0211202211033123-0200222331221023-2010100321330113-1113001031021000-2100021313301232-0222321320030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_rate_limit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_rate_limit

<a id="canonical-1013223133201022-1313332221032232-1331200200123112-3000301121301112-0033001302203101-2312333301123320-2111320003320121-2311310003323022"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable rate limit. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_rate_limit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201101022220023-1103210213332120-0012201103310011-2012220223033011-2311212133223002-2010030131033033-1301101200102210-0132000311201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_threat_mesh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_threat_mesh

<a id="canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-017.md#canonical-2020310210321000-1320011300000131-3121003230133311-3120332120300122-0323030233301302-1132133023102100-3031111331113233-0231102311232332)
- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-2100022020020011-0200313113330200-3010002022110312-1310112023231102-1110222331121122-3303232030030220-1330000212210230-0030303210201210)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010130233211201-0102210200231211-2331021322013011-3201112123233301-1203313132231022-1130220320000222-0131212013022231-0330001200132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_trust_client_ip_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_trust_client_ip_headers

<a id="canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: disable\_trust\_client\_ip\_headers, enable\_trust\_client\_ip\_headers; Default:
disable\_trust\_client\_ip\_headers\] Enable this option. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

OneOf alternatives in this subsection:

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-017.md#canonical-1023002121132310-3320330110101320-2303202132230113-1310320101111023-0012122130303312-2130100033130133-1323012232230013-1203210220020211)
- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-3331120100110120-1303333032232002-0301220313322023-3200013000131023-3322323222012333-3110100300213120-3312203133332021-1101320122300121)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_trust_client_ip_headers = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231033130322231-2203223313010330-1200320322113021-1333023133103333-3331123032011122-3320211321333012-0333310010101210-0312032310201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_waf` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- disable_waf

<a id="canonical-0021311132330031-3013103020211321-1223133002212101-2010122201101202-3101313331110313-2013021300302031-2213220033313203-3000113332031001"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for disable waf. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302213303120331-0322020022021133-2103031201103012-0013213333030332-3331330132313320-2112113133220233-3320130323210022-3120123311111331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `do_not_advertise` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- do_not_advertise

<a id="canonical-3021311322031332-3001130120301333-0101131021301003-2310200320010222-0133031223311213-1001132011030231-1000103013131203-3110003031031311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- enable_api_discovery

<a id="canonical-3332100020132100-1000303313230033-0200301222233203-1122131021321120-2033312022300121-1323331001230322-1320111223102220-3211202332310210"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030002131303001-0023303102231103-2212131100113301-3101321331002112-3303200131002010-3333200332003322-3013033322133100-2200231130120112"></a>

### Direct properties for `enable_api_discovery`

- [api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023): complete subsection reference.

- [api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221): complete subsection reference.

- [custom_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-0000012321113332-1333302321200013-0013010310122302-2330031222212000-0110030002020101-3212323100133110-1003113113210123-3333201032123220): complete subsection reference.

- [default_api_auth_discovery](resources--http_loadbalancer--reference--group-018.md#canonical-2111032232010102-2001332210110221-3230210002330121-0101133121102002-0323332000113323-1033002020031111-0300232133222113-3321010221300221): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-3303311120333123-3232301013211201-0312120323310102-2211133020013121-3002223311123231-0120220320333311-2120022120333300-3201333130101313): complete subsection reference.

- [discovered_api_settings](resources--http_loadbalancer--reference--group-018.md#canonical-3211321230303003-0331201322102020-1311001123210021-3120000101132231-1300020013122112-1222332100212100-3131201021231312-1022133323122023): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--http_loadbalancer--reference--group-018.md#canonical-0330033332332131-3300000123331232-3002221001230011-1010202232113010-0222112121000100-1313311123011300-3110331031201133-0301111301201030): complete subsection reference.

<a id="canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.api_crawler

<a id="canonical-3313213101132122-1202132221130300-1302030331330003-0311102333000311-0023322113031310-0303110020123220-2331222223222032-1200233000212320"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210212201020102-0220031023312332-0122121322322100-2133122102321131-2031230221030112-3132233110231131-0313131333010023-0223331032322322"></a>

### Direct properties for `enable_api_discovery.api_crawler`

- [api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310): complete subsection reference.

- [disable_api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210): complete subsection reference.

<a id="canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-0011013001010323-0102201332220003-3312322223210210-1000003121102120-2110103300130133-0222112132001333-3313031321030031-3123020010111021"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103321121120320-0223203021001320-0032031100023110-1201102311323111-2201202003322010-2031010302303311-1222003120010312-1300002220023001"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config`

- [domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002): complete subsection reference.

<a id="canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1000021102003321-3312202002202311-1232002300103200-0303230203313302-1222132301110311-3203011300032320-2232110133033022-3100220102020121"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223212023130233-1030031111233223-1101320131330101-0222231323330033-3102113321031221-3303133333010110-2230211133001233-0332231233102301"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains`

<a id="canonical-3231210223231332-2322303133210301-2113102310231310-1010033112023331-2310102001302003-0110102010333310-1011322123333203-3020211010113302"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.domain` property

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302): complete subsection reference.

<a id="canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-2231030121113112-0333331210332332-2033232111011202-3221331312310032-0221133131211003-0201231321332311-2303130001132121-0331103222231223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300103322200223-3300132020200032-0302323210320033-0301321231311112-3322001101311302-2003123300112101-3020320001310301-0203113112322312"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login`

- [password](resources--http_loadbalancer--reference--group-017.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021): complete subsection reference.

<a id="canonical-2211221111221131-3022220213010230-2121111111123201-3310220122311230-1232121000100223-1320032010100013-1233021331203020-3002211202131223"></a>

<a id="canonical-2303101311223013-1202320110213302-0311112123003011-3102323130000010-3021130311020111-0333220330103201-0113320332301321-0231333002211233"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.user` property

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-0013221312120121-0121102003001202-0310012321221120-3320101300202223-0030113212200113-0330321012013030-0331302223133203-1201101320032231"></a>

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

<a id="canonical-1031233220310211-2321023021233100-3230123111310031-1000113030232231-0131321202132002-0003100301332133-3112010100331100-0213102102120313"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password`

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021): complete subsection reference.

<a id="canonical-3323233012331102-0210001301121113-1121332230211211-2111000000202033-2313313231002121-0110230011132333-0312110220111301-3012133330122133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-1131032213202212-3233101101203313-2201300100101322-3323113312122311-3021123003223220-2120222122121201-0012032131012321-0311200132222132"></a>

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

<a id="canonical-2001321030313203-1011000331232220-0030323302103230-0122231123230220-1001001201313333-3121122230022311-2002211221112212-2322203203031321"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info`

<a id="canonical-3222202303113312-2003211322313021-1213323210211310-2021211030312112-0031121313200000-3030313202313133-0013312130301313-3202013313001010"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0303130123333311-1122031200201332-2002211102311220-1321301032212100-2130112023032011-0200321112132021-2131010331232300-0001302311212020"></a>

<a id="canonical-3012302311322313-2120100011121200-0121330231213310-2021003033002333-1020232130133313-0010212213020120-2102133212020232-3002010023331023"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.location` property

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

<a id="canonical-3300313000211221-1202020320112203-0022312132111110-0212012031332300-2013212102230320-0332300313311300-3011221231222311-1301210323230113"></a>

<a id="canonical-1101022212032211-1322000200200013-2332112213102331-2013010110113030-1220013123112131-3302122320312030-3021211203202223-2231133333121011"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-1200020222210303-3300121022123201-1022330201232230-3212111103021323-2111332321032021-0130222012011322-1231220123003301-0320102102202021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--http_loadbalancer--reference--group-017.md#canonical-1130100001303210-2312223211203321-3131101301223302-0320012230001002-2220032132200013-1120320231233011-0030032120211123-3210220222213310)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--http_loadbalancer--reference--group-017.md#canonical-0331333232210132-0212230223100131-3301012002322212-3002120201303130-3033301303300102-2021303003223131-3211102121312313-2312103232322002)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--http_loadbalancer--reference--group-017.md#canonical-1122331133203032-1120013303232001-1000132231113220-1022232221100203-2100330222212113-3021221231121331-0200032200310230-0330312322321302)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--http_loadbalancer--reference--group-017.md#canonical-0012010023323012-2110330100100131-2300112220200111-2330121011231202-1230012000332310-3223331231032021-0233102333030123-3001010310120021)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-0312021033113131-3203323223210212-1222201323010321-0110230211033123-0122013132221032-0021301133103122-1003102232332121-2031232321102300"></a>

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

<a id="canonical-3302332122000322-2310031023103322-1203213021222011-0213212300232212-2102200020210332-0000011312111120-2013302111302230-0133301011320323"></a>

### Direct properties for `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info`

<a id="canonical-0313023203130033-1233203102211122-0333010232111310-3213100031102221-0021310022220220-3123111302331110-3221232230210113-1233202323202212"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0111233133113321-0302020023211100-3301231210121012-1120101331222113-3021122011313230-3212212102302301-0022321222100201-1023332200121103"></a>

<a id="canonical-0001121030133213-0122311233220110-2210110322023223-3031122000103021-3021013113231011-2201302021233200-0020131210321222-0210101022230310"></a>

#### `enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info.url` property

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

<a id="canonical-2130122330020020-0031312021130033-1110002033303211-2002113130032011-2312312132210322-0332022003213222-0313300212332231-1033111032302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_crawler.disable_api_crawler` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_crawler](resources--http_loadbalancer--reference--group-017.md#canonical-0020300211332303-2011201120121121-1032111002000323-1221100301132122-2212010023203333-1312321212331323-3202032233121031-3301111300232023)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-2213321311201302-2201122132100103-1002220003122122-0100222311301301-2031111322103202-0111030333301223-3100132223221232-3302010023331113"></a>

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
disable_api_crawler = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-1202202222012133-2133031222200200-2002201220002100-0331102113323201-2230220211230012-2122231101330113-2000323332001112-3130200223213121"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033221303322210-0332221131333123-3021023010001312-1132333220222002-2303101102312232-1031302012210201-1120021131013223-3323131301022003"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan`

- [code_base_integrations](resources--http_loadbalancer--reference--group-017.md#canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303): complete subsection reference.

<a id="canonical-0012121010330132-3131012021212021-1001102030023303-0113323023032000-3011323312122223-3102330233001320-1233020102001223-3321001231132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321)
- [enable_api_discovery.api_discovery_from_code_scan](resources--http_loadbalancer--reference--group-017.md#canonical-2022211120020101-1303213211101122-3331012220102332-2002132023213223-2013303102010122-0111321021212300-2033123331121313-1123111003011221)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-1223202010302331-0233102300033320-2300200202321111-3232012201030330-0013133001323003-2330103023103010-2001311313030110-1031313010013103"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220213323011302-1033203303300211-3023112323033131-0000313121221103-3232300322011121-2100222230100013-3111100131031212-3133012112023202"></a>

### Direct properties for `enable_api_discovery.api_discovery_from_code_scan.code_base_integrations`

- [all_repos](resources--http_loadbalancer--reference--group-018.md#canonical-2012231012033323-2313120330300030-3113220001001310-2310312201210333-0330011223232101-2212201331321002-1003002130002031-1100322103203101): complete subsection reference.

- [code_base_integration](resources--http_loadbalancer--reference--group-018.md#canonical-3011020133023203-1020010031211020-3210220300221021-1030003032211012-3320321132101110-2200123112011310-3123013030010000-1112331223011302): complete subsection reference.

- [selected_repos](resources--http_loadbalancer--reference--group-018.md#canonical-1330221102102213-0130023021302212-1210323220223000-0200101320031231-2211330313210013-3330033102122132-1011102002321103-3013033213013332): complete subsection reference.
