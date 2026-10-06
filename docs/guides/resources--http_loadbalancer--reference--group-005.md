---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0101310122001121-3031003300313023-2332320232123211-0310330303310223-0330313122101312-3312031031303221-1201300202330022-1101130001332233"></a>

## `advertise_on_public.public_ip.namespace` property

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

<a id="canonical-1210312210100213-0032333112033311-3030013112131033-3322032013213202-1300112023231001-0012330130132132-2011121112131023-2010202332310133"></a>

<a id="canonical-2012013022031303-3311231202323012-2123110323003011-2231203332013320-1023311111022211-0101113221013302-1022221313203010-3122301220012223"></a>

## `advertise_on_public.public_ip.tenant` property

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

<a id="canonical-3122032220203002-1320231213301220-1303120120111230-2231321223321013-0331311110130203-3101201330021212-3102330002010010-0011210003233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- advertise_on_public_default_vip

<a id="canonical-3201330121022122-3323220020131223-3101123310000020-1020300323303132-3021003123202123-3011313330020133-1101012321201313-2230032300223013"></a>

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
advertise_on_public_default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133220122231331-1012333101303110-2323230033102213-1132133010132311-1200323020332232-0211032223320321-2020233002320131-3320110032220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- advertise_v6_on_public

<a id="canonical-1121221002220212-0021120313311013-2132132013231303-1020103320301022-0302011322201310-2312232122212000-1221031020022111-3101323212220030"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121230111102021-2200001002001003-0311211130103320-3110322313233002-0100211101133111-1333213221031121-2100000323133202-3320223003221120"></a>

### Direct properties for `advertise_v6_on_public`

- [public_ip](resources--http_loadbalancer--reference--group-005.md#canonical-1023331000331232-3133332320001113-1230223213202213-3233001331122221-0322330220012301-1201023020300200-1210100221333232-1001302202000201): complete subsection reference.

<a id="canonical-1023331000331232-3133332320001113-1230223213202213-3233001331122221-0322330220012301-1201023020300200-1210100221333232-1001302202000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [advertise_v6_on_public](resources--http_loadbalancer--reference--group-005.md#canonical-2133220122231331-1012333101303110-2323230033102213-1132133010132311-1200323020332232-0211032223320321-2020233002320131-3320110032220210)
- advertise_v6_on_public.public_ip

<a id="canonical-3311233111031313-0222202130032031-2123331300130032-0002311211311220-3010231130011023-0220301102211230-1203330122100321-1130323032230321"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203033013131123-2233011312231122-3112331212122200-3030130002201330-1323010123320212-0213103202231113-0223212313212103-1232210331122201"></a>

### Direct properties for `advertise_v6_on_public.public_ip`

<a id="canonical-0333132222133222-0023200110213130-1223320210203131-1022111203002323-0033212222312203-2220333001222001-0003232112300312-0012102132132112"></a>

#### `advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-3220202321221310-0322220002202112-0101202230013112-1112133010302130-1002113223311030-2302132320321210-0313301123001300-0203100020133201"></a>

<a id="canonical-0330130311222330-3310130203010210-3300021010020311-2022323331003000-2122022003202131-0001110311322323-1001101303022222-1031203323101330"></a>

#### `advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-3300130233231121-2032022020222030-0022103211020212-2212111311110120-1321102132033321-3222221110321202-0033202202202311-2130212103233213"></a>

<a id="canonical-2323303122000223-0130102120333322-2113131031311133-3213021323330001-0122030132311000-2130221232303033-2130223302220303-0030313211121112"></a>

#### `advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- api_protection_rules

<a id="canonical-1203122010313032-1012233110133012-0022303112312103-0323232033202231-0000332132021030-2132322122213133-3211220200212131-0120202010130313"></a>

Type: `"object"`. single nested block, Optional.

API Protection Rules. API Protection Rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_protection_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030302101011311-1323211022100031-0013012322103313-0023022023313010-0210232310100330-2032122012023211-3001312331313211-2033320022301300"></a>

### Direct properties for `api_protection_rules`

- [api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313): complete subsection reference.

- [api_groups_rules](resources--http_loadbalancer--reference--group-006.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011): complete subsection reference.

<a id="canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- api_protection_rules.api_endpoint_rules

<a id="canonical-0233211332333132-2113310333233200-1103330000332330-1302303123203021-2323001102100010-2203103323003220-0022223030111031-2123113110103230"></a>

Type: `"object"`. list nested block, Optional.

This category defines specific rules per API endpoints. If request matches any of these rules,
skipping second category rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331102332131221-0333200230323022-0010323200330003-1123332012130021-1001003102303100-2210010200001330-1320012111301332-1233012312012013"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules`

- [action](resources--http_loadbalancer--reference--group-005.md#canonical-1003013021023300-0013111220323212-0300112011211302-1331201321101223-0012102211102100-1321223202213033-0222213102310301-2231323301323210): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-005.md#canonical-0132110303131122-3311111320212011-3312230310110302-1033123013133203-1223122231332101-0220111012311200-2130221133213202-3323111121020033): complete subsection reference.

- [api_endpoint_method](resources--http_loadbalancer--reference--group-005.md#canonical-2133333030311010-0133313102301221-2231102001223222-2200223012233111-1111210202001232-2012232323023022-0011202220233333-2221212320332001): complete subsection reference.

<a id="canonical-2313120312201210-0313233121111133-3032201201311313-3230200002210232-0100212113010132-2301212032022322-3332003022210100-1000212120233200"></a>

<a id="canonical-2222123333313302-3333030320131123-2222013222233012-0321321300221323-2130103121222320-0012230121200331-1201223112020002-2021301131313222"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-005.md#canonical-3030001110033000-3131310301001101-1113220222331313-0003111330222130-1113022020310230-0321012302132102-3000133202210303-2332233100122302): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012): complete subsection reference.

<a id="canonical-2002121132012112-2231203331231010-0302022030320133-0223123120133101-1332120233021223-2212123320312131-2322222222131220-3123132230321210"></a>

<a id="canonical-3111131002301202-0033333131021022-1001003301322122-0300010103322300-1033123303020303-1302211022323110-1033020022322133-2200110333322300"></a>

#### `api_protection_rules.api_endpoint_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-1003013021023300-0013111220323212-0300112011211302-1331201321101223-0012102211102100-1321223202213033-0222213102310301-2231323301323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.action

<a id="canonical-1010221131022223-3202110111301232-0012121032311203-2021023313233321-1230133222132311-1012331321320321-1022113320200121-3132220323201321"></a>

Type: `"object"`. single nested block, Optional.

The action to take if the input request matches the rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
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
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113333331313111-2220003221002230-0133020211330100-3113230120020212-3201000133223230-2213313332211200-3030003120100331-2323202020201333"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.action`

- [allow](resources--http_loadbalancer--reference--group-005.md#canonical-0320213103301101-3121222001310312-0333020200033323-0021120330120323-3302221002021133-3312103223323032-2110112231022113-0012110032310223): complete subsection reference.

- [deny](resources--http_loadbalancer--reference--group-005.md#canonical-1113222113302312-2313111022101103-1021202102123323-0112000032231220-1321000223003233-3221023213330312-0121222022020232-2232011000123330): complete subsection reference.

<a id="canonical-0320213103301101-3121222001310312-0333020200033323-0021120330120323-3302221002021133-3312103223323032-2110112231022113-0012110032310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action.allow` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1003013021023300-0013111220323212-0300112011211302-1331201321101223-0012102211102100-1321223202213033-0222213102310301-2231323301323210)
- api_protection_rules.api_endpoint_rules.action.allow

<a id="canonical-2012321231001032-1122020310323321-1031223302022023-0120320012210211-0112312031201123-0311121321313323-3131303310020202-1322010331201221"></a>

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
allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113222113302312-2313111022101103-1021202102123323-0112000032231220-1321000223003233-3221023213330312-0121222022020232-2232011000123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.action.deny` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-1003013021023300-0013111220323212-0300112011211302-1331201321101223-0012102211102100-1321223202213033-0222213102310301-2231323301323210)
- api_protection_rules.api_endpoint_rules.action.deny

<a id="canonical-3302013001002212-1130210033300333-3302033002321033-2010000123123202-3001131322211333-0020332102201031-2003303302112333-1200331003330331"></a>

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
deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132110303131122-3311111320212011-3312230310110302-1033123013133203-1223122231332101-0220111012311200-2130221133213202-3323111121020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.any_domain

<a id="canonical-3101030322332330-3233123023202220-1322112331302022-3211121030202033-0131000031320313-3110233113223103-1210101110000311-2020323233300111"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133333030311010-0133313102301221-2231102001223222-2200223012233111-1111210202001232-2012232323023022-0011202220233333-2221212320332001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.api_endpoint_method

<a id="canonical-1230012111310022-0321212023030212-2320201023022103-1331330122311322-3023033132302130-2332102210102202-3221211112122111-1103310020300311"></a>

Type: `"object"`. single nested block, Optional.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
api_endpoint_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132031103032123-1100020023301322-3203030332231221-2103322311202212-0100303002031110-0023230321323123-0322201200023212-1223323112223310"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.api_endpoint_method`

<a id="canonical-1013001132020120-3313322332012133-1211010221230210-2200121102100231-3033033200302020-0033130022033023-1312222212230100-3232033330110131"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_method.invert_matcher` property

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1003030030210210-1232313333101103-3232100211333110-0210300020033011-0202203311001100-0031002201221330-0011002232311110-3011211023200221"></a>

<a id="canonical-0010331100313031-0120300202233133-2223203032030230-2003132310020223-1310330033233203-1003200300002122-1033220131010210-0310322000020123"></a>

#### `api_protection_rules.api_endpoint_rules.api_endpoint_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.client_matcher

<a id="canonical-2233103103112321-3010222031131111-1123111013002003-3220202010203111-1110311303003131-2123101301021233-1122101003133013-1201000200032010"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122202013301312-2202102020210121-1222101202101023-3222212113313013-3130300231323110-2231200003112031-3332103011322303-1220222210313132"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher`

- [any_client](resources--http_loadbalancer--reference--group-005.md#canonical-2020112123211133-1023000111210131-0210323220032233-0320021121233113-0211210123123323-0332121320002033-2123030313023102-2331133203112001): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-005.md#canonical-0301101110310301-2100023033303011-3000032120233130-1232013222023103-0003000030002011-2212302221213122-2022311131031000-2012132302012100): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-005.md#canonical-0202132123233123-2323223223222103-3031333211132130-2101222132133323-2221111032321223-3220230132112233-0312302321223103-3110321303011201): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-005.md#canonical-1131312220313133-0012332222221011-2103002010311130-0331210300030032-0122301313332302-2111022302232122-3102002331002032-3332302033220321): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-005.md#canonical-1133120210023121-0002310111331330-2020101032033222-1121331003220303-1313232312313110-0110223113221122-0321132113112223-1222031231002301): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-005.md#canonical-0123120001322103-2331123203130211-0113111100023101-3012213212000010-2222131223033123-0123212321303310-3101133011113133-1120011212302030): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0130103020102113-3201320201003312-2033202202012021-1303212231222321-2231121221332103-1331011032100122-3110323111300113-2003100223313321): complete subsection reference.

<a id="canonical-2020112123211133-1023000111210131-0210323220032233-0320021121233113-0211210123123323-0332121320002033-2123030313023102-2331133203112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.any_client

<a id="canonical-3031103002200100-1311331012331303-0320322210010301-1112303312111201-3210220312110113-1330323033213103-2121013021003130-1233030111311331"></a>

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301101110310301-2100023033303011-3000032120233130-1232013222023103-0003000030002011-2212302221213122-2022311131031000-2012132302012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-2201121301222003-1010010011102103-1312323212210021-1032112121101111-2012303103202230-1021013302022331-2101310202131003-0213021222022123"></a>

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202132123233123-2323223223222103-3031333211132130-2101222132133323-2221111032321223-3220230132112233-0312302321223103-3110321303011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-2332111321002113-1000030003333131-2132032213233101-1202221212021213-3021313130333212-1201000030203322-0032313220103003-0310122102230133"></a>

Type: `"object"`. single nested block, Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302301300121001-0001212233230031-2203300311002331-0310001002310213-2000011020203301-1100232211231101-0022112221131301-2201133032233220"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-2331031203312321-3102032102322331-1021330303311032-2011110121020133-0130303112332320-3322302011312003-1123232020223333-3303220311231330"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-3301031000222302-2231013312113202-2031010211100133-2313020332310001-1321213212032310-0302110000002331-0230012233232201-0103203331012212"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010300312322320-0033102232302012-3112223001222030-2032032030332312-3031012321103200-0120032002031031-3231020211032100-0011210301300133"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2112310301111211-2102332321023101-0013312220203213-3003101033322210-2322232103012200-3023221012330200-0321113201301201-2033232032212122): complete subsection reference.

<a id="canonical-2112310301111211-2102332321023101-0013312220203213-3003101033322210-2322232103012200-3023221012330200-0321113201301201-2033232032212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2313020333100121-3211012003302233-2132200312112133-1232130330000111-2233301100221102-2202012132220231-1110333232230203-2003332133302331"></a>

Type: `"object"`. list nested block, Optional.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222300121131021-2020003110231211-1320131311120102-2311300121111031-0121120310102302-1212302311330002-0211220200130032-1012213230010332"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-0010023200320321-2131101122013313-2121103032112132-2113310233220210-3322132331211033-3000101302323201-0210210321322023-0121322320101203"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-2211100321031222-1232201200120333-1230103102100300-1302031312023331-1120033210231310-3322030311110221-1033222130010003-2002331103101032"></a>

<a id="canonical-1012012232213312-2022110120031000-3012121303220310-3010320331211323-2320310323320302-0313032022012021-1222322132301320-1003021131320331"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2113231310221003-3301302102331133-3110122220022012-1213202330013313-3030031130221013-3320023232323311-3212132122013210-1301331013230012"></a>

<a id="canonical-2012200230312232-1131232131212201-1133322331103120-1010023032023201-1022230212011123-1332102033301110-1013100321121111-1330101121231202"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-3131013032033312-3300303110313133-0023310333101200-0220010333013332-2101121123132132-3223013211030131-0133100020312210-3000020230232210"></a>

<a id="canonical-1130323000102131-1330013001321021-3122130123212311-2113103133023103-2332022131310021-3023001212330203-3220300311202320-3231013300102000"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2330333113232333-3332300200030230-1300010201101301-0332220302230021-2130222202102030-0321022003200213-1213013122031211-0310232202031333"></a>

<a id="canonical-1112100333312202-3113202011213121-0323012203230302-0113231203103110-1312122000020232-2130200313323231-2302301013110331-2321121000220212"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1131312220313133-0012332222221011-2103002010311130-0331210300030032-0122301313332302-2111022302232122-3102002331002032-3332302033220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-3022222323110220-3200031330320231-0031211202011120-3321022321012022-3322222133121013-0302000331333332-0310132323203331-1012031020023111"></a>

Type: `"object"`. single nested block, Optional.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323221100033233-1320121310002012-3312331231212023-3012212301332031-3123323110202132-3032302121220220-0022313203100323-0312133332333131"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-2230210121001100-1203130203320233-2223332303013300-1223020103313030-1201011130113231-3132230012001321-1320002002221031-2010221101201021"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.client_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-2023223100033321-1201032330310033-0012000112020210-1003100212011221-0203303322223110-3201221130110120-0003323031103013-1210020130032332"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120021123331331-1032323202303133-1022332033331001-2310023202011121-2212230302303001-3321210020120221-1120122320032211-3130202213020102"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-0130301313110102-2103303311102330-3011232322102033-0121102232320102-1302110212212031-1300032223000103-1220330223101003-3000212131120332"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2221103031301100-2211323222101322-3031312320131102-3030033232122310-1132001032023320-3200003303130130-1130032101123233-2232233310122030): complete subsection reference.

<a id="canonical-2221103031301100-2211323222101322-3031312320131102-3030033232122310-1132001032023320-3200003303130130-1130032101123233-2232233310122030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0023021231110102-2131131232131030-2321220303001211-1023111330132130-1023120123112021-3313331001131312-2201103200301102-1002320303301021"></a>

Type: `"object"`. list nested block, Optional.

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230230210003030-0233312113200001-3220300211320011-0230103233212022-0020220223103020-1312233012033020-3321331031100313-0122131323201331"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3210121103230132-1321220101013312-0122121003331233-1200111013002330-0331322111003101-0132321112213112-3312031222301023-3301200122020110"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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

<a id="canonical-0000312320332133-1111221133102211-2203323100132223-1122202223233233-0302023203002321-2100301113013113-3303203132101122-2103221312011333"></a>

<a id="canonical-1323321100332323-1311031211100111-3220333002300031-2033111301010310-3031221212032311-3012023312021331-1030133101331210-2022301110120331"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3310101210121132-2220022122022322-2021130312100212-2031312033020000-1032010030201323-0101123011102323-0031112330120010-0101320011210222"></a>

<a id="canonical-0020301023021132-0202131012022112-3320010022011110-3001200021211311-0003122020023231-0210112321132102-1120311221222030-3233112323201120"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-3232122302211223-1222101221231102-2012130202111310-0313332132230030-0022120202121132-0303312010020123-1311211001122000-1000110233012203"></a>

<a id="canonical-3101112122002031-2320110231033000-0001111100301111-2113113010212232-0112012302112212-1212022010030021-1032310322033121-1301012023221001"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0132300210132013-2131330031030121-2032110132131123-1023223102123220-2133213233320301-0220222211133000-0310003200010200-3123210130130321"></a>

<a id="canonical-3301333321133112-0332232031233022-1001020220312022-1210033330230321-1100020311133101-0131130213303112-0112230201133203-2023013332312321"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-1133120210023121-0002310111331330-2020101032033222-1121331003220303-1313232312313110-0110223113221122-0321132113112223-1222031231002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-1213313213202111-2112231220020213-3103032112113201-2120233322230123-0030021110322203-1010112010310210-2121222330232321-2330322223011311"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011201011102120-0101310322230101-1332223130120210-2031211032101232-2332011033033323-1302303123230012-3313203003133321-0101020001220113"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-1111011111122311-1223000123011011-0321131212103222-1211103002101113-3111112203313133-3003033201011103-2202331220210131-3301010323101022"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0323232012023210-3320333133003221-2123111113030231-0320202200211203-2010002030100003-2103330130031003-2331121012302201-0222012111001121"></a>

<a id="canonical-2003103023130023-0101001033330012-3103003010002101-3300313103321331-2210212233222302-3333013101030031-0122112212002223-2320201231213001"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0123120001322103-2331123203130211-0113111100023101-3012213212000010-2222131223033123-0123212321303310-3101133011113133-1120011212302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-2102001233030121-1220220232203000-0303232312230031-3331033221120031-1112203023120101-2313302302200311-3210103021003211-3323202102310013"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201211212002233-1321123311031032-1321301100302333-1123102110112103-1320201012310233-2132011111332212-3321312112120322-2330321022321220"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-1010123030000012-1321330122302303-1011323201130133-1230011200320330-2233310202131020-1131031232231000-1222123321132030-2131223000131031"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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

<a id="canonical-0130103020102113-3201320201003312-2033202202012021-1303212231222321-2231121221332103-1331011032100122-3110323111300113-2003100223313321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3003010103031122-0133303310112333-3110312110333000-2331211002103131-0323233332200311-1211130031302223-1013213003022321-3022213232122010"></a>

Type: `"object"`. single nested block, Optional.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001131032010113-3202300111331203-2001202210323221-0132300313233103-3303230303032320-2000101130021322-0103310022221233-0112102203121322"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-0200310333303230-1022300313333130-0321331210333020-0010132113110120-3122130201313332-1021322233201112-2201200322222213-2212200233102022"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-3132331233222303-3031331301121032-0033122333203330-1230202233132132-2130103210231112-1303033220323112-0300110333000232-2201200030123303"></a>

<a id="canonical-2012111102223212-2131102012030320-0312110223133301-3323210012023122-1033231301020131-3212212310333122-1100332223011012-3302233112003320"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222302231313031-1013232232133100-3122000023310323-3213132110332130-3230122300011102-1302030330212032-3232031331102311-1113010013203313"></a>

<a id="canonical-2203112022032322-1030023221112222-0323012102113301-2113311000322133-1011302011102120-0320122001000322-3300101120202201-1211303110311320"></a>

#### `api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3030001110033000-3131310301001101-1113220222331313-0003111330222130-1113022020310230-0321012302132102-3000133202210303-2332233100122302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.metadata

<a id="canonical-3303332130023002-1210023233020120-0332012032220231-0123022213003113-0211023101212320-0312300211102212-1200033332132120-2210200030110213"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113311220033133-1020102310011302-1331320201030323-0033202223011221-2220021130222001-1010312133020012-2230330103130132-2132111033111213"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.metadata`

<a id="canonical-0211121322110210-1330302113001113-2120220110103121-3132100331121030-1322012123221023-3231113310032222-2232223300313121-2012330233100120"></a>

#### `api_protection_rules.api_endpoint_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1011322212223232-2220213120232231-1033320200223123-1301130323213233-0103201323320223-1222001021021201-1323303003230010-2212000011000131"></a>

<a id="canonical-0031013033210031-0001222122202203-3011113122233223-1031330321131322-0221330331331333-1312231311330231-0321303030213213-3223000212101112"></a>

#### `api_protection_rules.api_endpoint_rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="canonical-2233000333303300-0301120300310222-0132201110222011-2202021231010323-3121323112213221-0112231031331121-3112330331221321-0012120203112000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Additional upstream details:

Request conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232103210022201-3202323201131113-3333103330022132-2030312320100232-2223032233002001-1230123321233312-1100313313132213-0112221130101130"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher`

- [cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-006.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332): complete subsection reference.

<a id="canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-2020113010323330-2202100122302102-0131221301333213-2033211023223023-1012112321003100-0320232123022211-1323313310212102-2211300003132333"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313310332033120-2020311333202021-3323231202121121-0333313032010210-2121100132010203-1010131033330012-3201033010313301-1322300130212321"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-3210232211323311-0023211102021213-1311222011113111-2222132232210022-1212023013010212-1223233112323033-3232131023020133-1300313010311103): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-3220223121012213-1031000033003030-1020331010300303-0223030311111023-3123131132203210-3012102330223330-0103203101020022-2023201331302030): complete subsection reference.

<a id="canonical-0011030021121320-1101130201231113-3130320102101323-3031200233302222-3210302310130221-1300132320203303-2002000212330022-0130101020303211"></a>

<a id="canonical-0023112003031221-0002003002331132-3121332213201301-3321103010332210-2333301202310101-0120103331010211-2331222110323210-2112003222033301"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-1002133022121033-0201122130133302-1123103100203310-1030022303002202-1123330300122213-2320111112202230-2113233131230131-3301011321013233): complete subsection reference.

<a id="canonical-0313203100320011-1322210333130301-2310220331110311-0230132203201323-0332031123122033-1122232322000212-1110021231222320-1321121321331130"></a>

<a id="canonical-3311021220222210-0013303012300202-1132323122312120-1332123133302022-3210232112102020-3312123013130203-0013110232313321-1030323221231301"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

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
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3210232211323311-0023211102021213-1311222011113111-2222132232210022-1212023013010212-1223233112323033-3232131023020133-1300313010311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-1213233301123201-2311033223131311-1002300330110212-2002010133013100-3313113322001203-0033101112110002-0010323313332330-0002011232332003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220223121012213-1031000033003030-1020331010300303-0223030311111023-3123131132203210-3012102330223330-0103203101020022-2023201331302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3120102301323031-2231320202031312-2300132322032322-0321002011321202-3320232300323303-3022302311210021-0301113323330013-3131031003312313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002133022121033-0201122130133302-1123103100203310-1030022303002202-1123330300122213-2320111112202230-2113233131230131-3301011321013233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-2033221102320210-1031212223002101-0320112131312012-1123322112312011-0011101101212120-2220301101202010-0202212011220303-1230323022203030"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221112013011231-0021002202210132-2300300110310302-1313022331310021-2222220323321030-0021211023030133-3203321211011112-1113330101121213"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-3301200000132301-0011122023212320-2332003321022223-2201022212112233-3110101331331300-3300302231010201-0223222022201321-2221000202213112"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0233213212330312-2011021020323311-0222103312021022-0131011222323001-1210023113331100-0023122033313220-2022023210223133-1201212031133233"></a>

<a id="canonical-3013122100110322-0031110030101120-0002222331233103-2003111111303320-2302022300103220-0321032012002032-3130122103010100-2020303013203232"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1223030110321032-1300313300322313-3233100132011131-0233013133223121-3210201131001212-2202323221002103-0201212121312231-3300120121133130"></a>

<a id="canonical-2323122322122233-2101133132110301-3030221033300131-2013300213312330-0322000231001321-2000222321000231-1132020301130113-1033001200323312"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.headers

<a id="canonical-0222330302012330-1221032300020231-0003201103223231-1123333330330003-1201212020333332-2122120203003303-0123331203010223-3021131331010113"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103122233202102-3102120221320010-2231133100012333-1102221033102003-2100202113221311-1232331321013232-3130030333111022-3133322132103332"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-1131212122002031-3301103113313301-3123113013200331-3210203100230333-3101332233223223-0002000322110020-2303130211303320-3113123122202203): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-0303123022031230-0210331320202313-3011120333111210-2311131130302031-1322331131013123-0302110202112301-2311331021211312-2102000333231113): complete subsection reference.

<a id="canonical-0130222202023013-1211212211003132-0202203101112313-3310221012023020-3330233333213031-3310031110302212-1302201320131103-0133220110231210"></a>

<a id="canonical-2201110020301231-2100311331231303-3001233132103121-3311133232230002-0300323213113233-2120003021312230-3120021311002001-1330031322131121"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-0222202333231133-3130233100122132-2001300113002031-1222000230031121-0033300020310102-1102201133102220-2001302332231031-1201030012220110): complete subsection reference.

<a id="canonical-1232102201302331-0011012103220103-3020112300313311-0132332021023102-1211210000101122-3200323101002220-0212122333031212-3113310123333020"></a>

<a id="canonical-2023111012302212-1230021123220131-3230002021032110-0131131101033322-0331203331030321-1033231320302330-3112031002221310-0110033122121010"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

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
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1131212122002031-3301103113313301-3123113013200331-3210203100230333-3101332233223223-0002000322110020-2303130211303320-3113123122202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0310313222023133-3322122230030131-2330312120112022-3021131030302312-2301021321332032-0021203230310002-0013013311323000-0120230210310032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303123022031230-0210331320202313-3011120333111210-2311131130302031-1322331131013123-0302110202112301-2311331021211312-2102000333231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3222113130313113-3321021112013222-1100121100200103-2003022010002330-0320120310331310-3023101300121323-2313121011302012-1302223201310100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222202333231133-3130233100122132-2001300113002031-1222000230031121-0033300020310102-1102201133102220-2001302332231031-1201030012220110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-0313310030132200-0201130200111123-1311113023120103-1013313123322033-3123202132222202-2021012221100232-0230220111130102-0013023002020132"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232132132331311-1023002212010032-0021221213222223-1322112220213302-3211320123022100-1113301331330312-3030123311332212-1102203021203011"></a>

### Direct properties for `api_protection_rules.api_endpoint_rules.request_matcher.headers.item`

<a id="canonical-0222223123231233-1012223010120003-3223210300231133-0211102031021213-0110223310010132-3302303212010312-1101000311311201-3110223320300133"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2030011133231112-2120331323310222-2201233302022112-2222322021010223-2122321010201233-2030013321102002-1331301102331302-0211303010003010"></a>

<a id="canonical-0033231200101102-3000113331321322-1211113302230223-1201003100122120-2233023320021213-2233122021000330-1132233213301120-3013002211101300"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1122001001230103-1010200323210300-0212130123322032-1303230313321022-2312303033202101-3032330133010033-0033202021022330-0113313002020303"></a>

<a id="canonical-0331023031302332-3130122013033222-2320021011011121-1121233132012030-1031133323310223-0210133031110000-1002200230113222-2320101231022223"></a>

#### `api_protection_rules.api_endpoint_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0021100202000323-3013102011221022-2231210200230102-0002113030002003-2302111033300120-3322032011213103-3300221111211102-2201313132333023"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```
