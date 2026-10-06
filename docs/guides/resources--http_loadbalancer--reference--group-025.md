---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1102322301312031-3201231211122112-2101022200020331-1021321333212213-3103022322333203-3313100032331102-3033023233111210-1130311231303313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_none

<a id="canonical-3202023323212002-3032033213100112-1123312131202001-2131312331303001-0113230000100210-3001210211011302-0111113312301003-2322003301000330"></a>

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
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332013231031311-3111321203331230-1210323300201203-0103111103210012-1133012210301011-1203213110013110-1012302001330233-3330321002132301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ring_hash.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330)
- [ring_hash.hash_policy](resources--http_loadbalancer--reference--group-024.md#canonical-0211321300323030-2031222300103231-1302110213110202-2232223101103220-1013221212330133-3222213311123033-2320311023010012-3100103311110313)
- [ring_hash.hash_policy.cookie](resources--http_loadbalancer--reference--group-024.md#canonical-3131013300110333-2102200210330133-2102221222232012-3000213212101123-0202022013221110-1322110330301102-3003222330332021-2332003223213022)
- ring_hash.hash_policy.cookie.samesite_strict

<a id="canonical-2221232320232210-3102000210122201-0013322231302022-1213121230230323-0210210232213131-3323121133012011-3022332100130213-3032122111300113"></a>

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
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120201302000120-1101032023332332-1203301030103211-2301112131023100-3323333313021213-2102103120123013-3211232030113002-1310102202033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `round_robin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- round_robin

<a id="canonical-1120231012232110-3001211103003332-2110332213132110-2233201023123010-3331031002232323-2201313122330202-3321233220103133-3330032203222203"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for round robin. Defaults to \`map\[\]\`. Server applies default when
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
round_robin = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- routes

<a id="canonical-1022112302210201-1010011312111323-0033100111313112-0212233102202102-1021022230231123-1121103333323120-1231212033301110-1123111133332122"></a>

Type: `"object"`. list nested block, Optional.

Routes allow users to define match condition on a path and/or HTTP method to either forward matching
traffic to origin pool or redirect matching traffic to a different URL or respond directly to
matching traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("route_state_disabled",
    "route_state_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302333330103332-3000231123012200-1223313032013033-3110312130233300-3201000202121113-1221011200232010-0313220022331010-2221323130111313"></a>

### Direct properties for `routes`

- [custom_route_object](resources--http_loadbalancer--reference--group-025.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130): complete subsection reference.

- [direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311): complete subsection reference.

- [redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223): complete subsection reference.

- [route_state_disabled](resources--http_loadbalancer--reference--group-025.md#canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303): complete subsection reference.

- [route_state_enabled](resources--http_loadbalancer--reference--group-025.md#canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203): complete subsection reference.

- [simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232): complete subsection reference.

<a id="canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.custom_route_object

<a id="canonical-2132120132332033-0111230301202203-1332310120131330-3101230211330320-1302310110310033-2231032202200231-0302213013300201-0020203132023302"></a>

Type: `"object"`. single nested block, Optional.

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310333203312222-2032203322200333-3103123323230211-1320102023312023-0013322211322022-0302203320221221-1122002313223023-1000303311113320"></a>

### Direct properties for `routes.custom_route_object`

- [caching_disable](resources--http_loadbalancer--reference--group-025.md#canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-025.md#canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110): complete subsection reference.

- [route_ref](resources--http_loadbalancer--reference--group-025.md#canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223): complete subsection reference.

<a id="canonical-1311233302110112-2123122312121123-3301300201233322-3000022223123013-2103011011100200-0220322320333122-3102212130300133-0111330221321312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_disable` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-025.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_disable

<a id="canonical-1300132233003321-2310012321003102-3313000230313012-0113022222000330-3212023132023010-0023031201321130-3322202333303202-1100023203132133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.caching_inherit` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-025.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.caching_inherit

<a id="canonical-0023203021031101-3102012330230032-2322300303011202-1102031110221332-2122231020331130-0013003023132321-3233331212131000-3111121233333103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032301330003202-1101233001033313-2120303011100012-0322132022031210-2133201012222110-0102201103210232-1001021232132210-0010021002202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom_route_object.route_ref` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.custom_route_object](resources--http_loadbalancer--reference--group-025.md#canonical-0113030231032003-0103322323331330-1221133321102302-1223101002331311-1112100110201231-3133032002221223-1012023210323121-1113121000321130)
- routes.custom_route_object.route_ref

<a id="canonical-1123021100312002-1020112010103001-2313033300010013-1303122122012223-0322203113312222-3001003121332000-0101330122123222-0033201313203300"></a>

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102303303101220-2220313232201222-2311212111013010-1023203313233111-0333303322101003-3011203321211000-0022233012120313-3313232320032320"></a>

### Direct properties for `routes.custom_route_object.route_ref`

<a id="canonical-2323000200333123-0132121222212333-2210330112033001-0112112303232223-1213013021101003-1213132222113321-2330222001330012-2320213231123322"></a>

#### `routes.custom_route_object.route_ref.name` property

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

<a id="canonical-2201320121001121-0103233133130321-2101212203012032-2033101021032220-1230300223120133-1111112032301030-3312332122301100-0220322033203122"></a>

<a id="canonical-1311321212322103-3101311301011101-2031320300213121-3003320022201310-1210130033302301-3202212201010300-2321101213220000-1222103300233331"></a>

#### `routes.custom_route_object.route_ref.namespace` property

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

<a id="canonical-0221030111013133-1101230221300312-1320131302031110-3310103212313113-0130122232133121-3313211110000220-0120311133031233-3233210101331211"></a>

<a id="canonical-3103132333232223-0210101111301211-2313102222332013-0112132211230132-3103123031231303-2321313013123321-2002130113113013-0022123222000133"></a>

#### `routes.custom_route_object.route_ref.tenant` property

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

<a id="canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.direct_response_route

<a id="canonical-2212332230001201-2200303013102020-2331330332220323-1212131312310033-0303100201123223-3332120313321120-3323111102000031-3120322323221310"></a>

Type: `"object"`. single nested block, Optional.

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
direct_response_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231203213112001-1101032030331021-1020030011212321-1300000102330231-3010001302223211-0111013003103303-2211200202013321-3211210323332211"></a>

### Direct properties for `routes.direct_response_route`

- [headers](resources--http_loadbalancer--reference--group-025.md#canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133): complete subsection reference.

<a id="canonical-1122320301230101-1201300111222321-3200131200000200-0011323233232012-2001002210321013-1223010111113202-3302221330003001-2012320102223311"></a>

<a id="canonical-3303333213133003-2121001302000121-3330230011310322-1130010131201222-3120003332332031-0100023121321202-2320310321322300-2033202022113323"></a>

#### `routes.direct_response_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-025.md#canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233): complete subsection reference.

- [route_direct_response](resources--http_loadbalancer--reference--group-025.md#canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320): complete subsection reference.

<a id="canonical-2003132221231231-0202223123102110-2223232312130233-1130122213330332-2111222333122211-3113210320101301-2200223233123233-2232210003013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.headers

<a id="canonical-0012332300232233-2300123121101322-0230222021131133-1331332110303131-2000331133321110-0120322233122301-0033120130031021-0332222011302133"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013322223213310-2021313020321210-2012032031210100-0113001031130131-3311010233312212-3303322213331130-3312313030321202-1311100301303133"></a>

### Direct properties for `routes.direct_response_route.headers`

<a id="canonical-2013030023011233-0003132110221320-0233300111103123-2213132100123213-0303220033321100-0322200212300331-1133002002203333-1202033030221101"></a>

#### `routes.direct_response_route.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0030210001022320-2130302220112322-1133000102110022-2231303320003321-3132310003201332-0023001033223110-2331233332121103-0303221031322323"></a>

<a id="canonical-0122021221020332-1111130033232312-1330001012102123-0000331013220202-0010103223113133-2311233121002322-1222300001031200-0033222011321002"></a>

#### `routes.direct_response_route.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3013220100011213-3203332322113000-2021133221211121-0123301033332013-3121000110102013-0022330003303012-3321112022031132-0220231123103000"></a>

<a id="canonical-0101213020110003-2131331222302130-2133330010203210-3303030113000220-3013322233012230-2212233032312332-0231031133102113-2123110232211032"></a>

#### `routes.direct_response_route.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1133322321301032-2133311133200113-0200200133231332-1112230202233003-1031302310303012-3030123200330211-2112312022010103-1310202030232213"></a>

<a id="canonical-2120322223313323-1320313303311203-1321300201330110-3211021001212031-1203320331200011-2121010332222201-2032102111013101-0213202201332111"></a>

#### `routes.direct_response_route.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123212112231001-1301302003301210-2211323232003121-3031231333030220-2101023012222322-0130122112101310-3103031133112002-0200131221312123"></a>

<a id="canonical-3211222113013123-2013101001321232-1311231231031011-1320301032132113-1302303211330201-3202312000300110-2023211111032123-3201313013213010"></a>

#### `routes.direct_response_route.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.incoming_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.incoming_port

<a id="canonical-0010030313231201-0201232322323211-0220321212230231-3011233202110100-3033311203002302-0303131010110200-2130333232032332-2130032130012302"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003122032113122-1121123110231012-2031232131133222-2210233100201232-2223103132133330-0120312101013322-1000102130033021-0133311212123012"></a>

### Direct properties for `routes.direct_response_route.incoming_port`

- [no_port_match](resources--http_loadbalancer--reference--group-025.md#canonical-1331102132201121-0032301222132323-1023203312102303-1220103202030130-2320013100232330-0232333120031311-3300203113320032-2113010021211202): complete subsection reference.

<a id="canonical-2310311221220313-3020200222211300-1102313020123210-0013232320300033-2312201321013310-2211230300001231-1330211300023002-2233221330002213"></a>

<a id="canonical-1223223003321303-3212101021322213-0020132322001013-0011100321120111-0101123022031323-1221101101011313-0332112221122121-2031323002123233"></a>

#### `routes.direct_response_route.incoming_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3022020013313131-1201213001231210-1302003100120331-1122113120300221-3302211012230010-3032102211020311-1030120101223100-2303312232022033"></a>

<a id="canonical-1132110121300121-2002211001132222-2021023210322111-2120120032313012-3301230320130232-1202211110010323-2301031103123101-0212232013122113"></a>

#### `routes.direct_response_route.incoming_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-1331102132201121-0032301222132323-1023203312102303-1220103202030130-2320013100232330-0232333120031311-3300203113320032-2113010021211202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- [routes.direct_response_route.incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-3132030132301031-1130201213001202-0131102232000013-1131320220133200-2002301020330312-2202202221323301-3001110121211010-1030223232222221)
- routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-1120100133113130-0230002230001113-3032202022101320-0300101221113200-1123113203302110-3231320202013320-2330113103213220-1321133332012022"></a>

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100012011013133-1121321200112202-3220212333201120-3333120030120203-2110221022101301-3333021223010130-0110202231111013-1113321313212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.path

<a id="canonical-2123030001220310-0032112112231312-1323221203002000-1333210210232011-3032101112112232-0330203300123220-0310221000213310-1323001232013132"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310132333010120-1323213310011123-2002321121022032-0023322020030010-3132202302102131-3233311311111223-3133221131323330-2232332031210332"></a>

### Direct properties for `routes.direct_response_route.path`

<a id="canonical-1210112120201033-0003132032113310-3330302021130333-3023000200111020-0323130123132213-2012301112222202-1130201023302312-1223223200130210"></a>

#### `routes.direct_response_route.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3203200023320320-3022231201103301-2031322231130112-0032101323120120-2213311112211202-2003233011001002-1232220203213221-1221221010000230"></a>

<a id="canonical-2231332102202001-1100111202201301-0302310012303321-0222200111132101-3001001131303020-2300123103203221-2013311221100230-0013103122013301"></a>

#### `routes.direct_response_route.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2113310112230021-2032110332210233-0311321031322301-3021232120221003-0210023230020023-2301231332100230-2131022113131331-0303133222311220"></a>

<a id="canonical-3000312330103302-0202023220002212-3012322110132213-3020103310000312-3021201032112001-1212331231333310-3333321133332203-1001230102100330"></a>

#### `routes.direct_response_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2321321013121302-1022302203221120-1002102000313123-1122302221233033-0003031122132211-2332231022130103-2310213100023113-0321200220232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.direct_response_route.route_direct_response` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.direct_response_route](resources--http_loadbalancer--reference--group-025.md#canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311)
- routes.direct_response_route.route_direct_response

<a id="canonical-3213211133213302-0333001303300200-3003221012221322-2123301223023211-1023301101321300-0230210032320111-1132301103022012-0022032231023331"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132133112100133-2201100022203202-3331023133003231-3232131131110013-1233302322223310-0100203113232201-0321021111002011-2300131211131121"></a>

### Direct properties for `routes.direct_response_route.route_direct_response`

<a id="canonical-2121031131100330-2321212011300132-2023000001220011-0311001220122300-0303001021211010-2010210002220002-1130311211321311-1110223012303333"></a>

#### `routes.direct_response_route.route_direct_response.response_body_encoded` property

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

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
    "byteLength": {
      "max": 65536
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2301221330322221-2031020121230122-0030032113320211-3322101011131320-1330322111032223-2012132120231001-2210202103333123-0212222200103031"></a>

<a id="canonical-0301220030030031-0121221333012021-1233100233230113-1310313303313132-1023021313311312-1200322023320331-2100323031200012-2133113231113130"></a>

#### `routes.direct_response_route.route_direct_response.response_code` property

Type: `"number"`. Optional.

Response Code. Response code to send.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.redirect_route

<a id="canonical-0311011002232332-1323022032033321-2322303332331220-2232311213123212-0102222022201230-2133233221120032-2010221020312131-3130020003032202"></a>

Type: `"object"`. single nested block, Optional.

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
redirect_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300123002010103-2323232011201201-2003210012102023-1122221203130102-0113110331133123-1310211232120102-1032132330220301-0222103210303110"></a>

### Direct properties for `routes.redirect_route`

- [headers](resources--http_loadbalancer--reference--group-025.md#canonical-1012101232011000-1123332000231212-2010303222022313-1111112132300001-3101031010322232-1321223232223321-2231031213011001-2202132202020132): complete subsection reference.

<a id="canonical-1330213120223223-0221132001132120-3302123321100321-0313011332203101-0233221232033022-3121322102330102-1202122122121012-0020220133103320"></a>

<a id="canonical-0220232021313312-2301021202123131-0000031130011022-2303120033223300-3310232012230000-2210231023020101-0323311220113120-1131210312212023"></a>

#### `routes.redirect_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-025.md#canonical-2023020103020223-0112101233123130-1311203130032022-0133100102302010-0002330321210002-2113313131300313-1110223311232213-0310323102121021): complete subsection reference.

- [route_redirect](resources--http_loadbalancer--reference--group-025.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011): complete subsection reference.

<a id="canonical-1012101232011000-1123332000231212-2010303222022313-1111112132300001-3101031010322232-1321223232223321-2231031213011001-2202132202020132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.headers

<a id="canonical-0001220321021231-2131102303311121-3330031003120230-3110003303320203-0132101112110031-0211333331131021-1103011332310131-3133111310211330"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300103223223023-2020123223130212-0330112302203303-1123213103131310-1030133010033233-0322233000323022-2103322302130220-2321333301301133"></a>

### Direct properties for `routes.redirect_route.headers`

<a id="canonical-0320131032231131-3221100221130023-2301002022033320-0023233121113222-0201333112312130-0102232000120211-3110222323222103-3013030223130102"></a>

#### `routes.redirect_route.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3030332321020301-3102320203103330-1013001030103212-0000133330332233-0010211230013130-0132333201311200-0131003210330201-1331030002302111"></a>

<a id="canonical-1133111311021212-2001323231202330-0311100310323122-1023300201100012-1231100123120030-2102320323230321-1111120333110221-0211023233021222"></a>

#### `routes.redirect_route.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132202002310202-0323230320130033-2112211000000032-1131323113122121-3120122131020130-1312311210110000-1120310030221233-3322103231112002"></a>

<a id="canonical-0111031220332123-1103032201231102-2112323110300022-2201011232100001-0232312221102302-2200330302121222-2012301333211021-2121222301310021"></a>

#### `routes.redirect_route.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1322302131111011-0001013131300021-0322333130201232-3310313301321101-0222212303120302-2021333033110212-3023002131131313-2130310232223220"></a>

<a id="canonical-1131132233322112-1131123223121220-3001202023232032-3013331211201033-2121000230033230-2023033313310100-3330102332003320-0200222230001021"></a>

#### `routes.redirect_route.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1302300231213323-2310023133103302-3102220220131013-2301233013103212-0300133313211313-2231131310020320-1131120331100001-2122000222203100"></a>

<a id="canonical-2331031211001012-3132302110320231-3213101203101120-1202013302023010-1320120102200100-3013203223323331-3033012132220000-3101312212101221"></a>

#### `routes.redirect_route.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.incoming_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.incoming_port

<a id="canonical-1130011123033212-3233313012011202-0210023031211022-3232122110203303-1001130102031321-1300323032023212-0120322121103003-1203212130310220"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000320213312332-1033123223003232-0110013303223112-0100230201132223-1102232300113103-0123232123101320-2103133200000031-2201303133030311"></a>

### Direct properties for `routes.redirect_route.incoming_port`

- [no_port_match](resources--http_loadbalancer--reference--group-025.md#canonical-2101022032022113-0000033222103332-2120031101001221-0132231222332011-2022301023332230-0220203133310123-1200221332020011-1023312233230223): complete subsection reference.

<a id="canonical-2322220232233332-0200213123011321-2120201200211002-0222222332002212-0223223022303102-2321331101132003-2020320021120113-0213102121000102"></a>

<a id="canonical-1101313131000122-2311200112001220-3100213220120203-1032313231101310-3203003211233302-0311233121113121-0200100013012113-1121222231012100"></a>

#### `routes.redirect_route.incoming_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1033031321020310-0230103121032333-1231031011332332-1031013221330211-3222203332101133-0311023202000021-1101323030223330-2000123123211000"></a>

<a id="canonical-3212312130121021-1301002321321200-0001301311100302-3012023122031031-2331310031313123-1033212100222001-0010232221311330-0303213120200302"></a>

#### `routes.redirect_route.incoming_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-2101022032022113-0000033222103332-2120031101001221-0132231222332011-2022301023332230-0220203133310123-1200221332020011-1023312233230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.incoming_port](resources--http_loadbalancer--reference--group-025.md#canonical-1312331230210132-0233112212120302-2003110121211311-0213112102301232-1001121130320222-0221012010330101-3112012310102303-3222032301302101)
- routes.redirect_route.incoming_port.no_port_match

<a id="canonical-0232321023233331-3022002131110110-3032233213031333-0113211111200323-0203001110201012-1101012003333200-3201001312033211-3330020212031210"></a>

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023020103020223-0112101233123130-1311203130032022-0133100102302010-0002330321210002-2113313131300313-1110223311232213-0310323102121021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.path

<a id="canonical-0322332033000233-1023101112021013-0121233000230002-1211220012032102-2130203030130033-2332000313011330-2231131130020131-1132322312122321"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113201032002123-0130213022003301-0301111221130000-0330231332313211-1210131022103112-2133020301133201-0312120033220220-0313131311330312"></a>

### Direct properties for `routes.redirect_route.path`

<a id="canonical-3312211100121222-1023203120303312-2112110013130231-3101230003333112-3333133103033033-0321210112010301-0333012302313201-1202200231011311"></a>

#### `routes.redirect_route.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1003033223102311-1023231112032030-2021000320231213-3321030301002301-1321022002311030-1222211313221030-0012322032113322-3211010220000001"></a>

<a id="canonical-3213222312020320-1222311031212202-2023010322201221-3310000011222333-1031110300321123-0001122001023122-2323323303203322-2312210310011320"></a>

#### `routes.redirect_route.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3323000211323231-0130020103111211-3323210211001112-3020333010331122-0002113203113023-2010313310222001-1133203311223030-3301011021211021"></a>

<a id="canonical-3130310210312130-2121031020222222-2201330311023132-3211022323332132-3311212030333030-0310011222133010-3013021301110303-1011211223131231"></a>

#### `routes.redirect_route.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.route_redirect` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- routes.redirect_route.route_redirect

<a id="canonical-2320211100222130-1302322000233133-3031100330010023-3331301300332310-1221102132021100-2023133010131232-0333121022000330-1100131312102212"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022231230022023-3121121000000211-2321102221131101-3031001120333123-1212200132321302-3301220101303223-3102101012313211-2121032003323102"></a>

### Direct properties for `routes.redirect_route.route_redirect`

<a id="canonical-3120120112132302-2123113031300233-1101322202101303-0122132203131023-0200112203133233-2213332223132023-3321200211302131-2320323112323020"></a>

#### `routes.redirect_route.route_redirect.host_redirect` property

Type: `"string"`. Optional.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-1331100133113301-1323312211331300-2230331002021033-2221222320131200-1113211313122301-2210022032033332-3233321331300223-3222232013213330"></a>

<a id="canonical-3002013120231133-1123101232222113-1011332113002013-3120020332232032-3002000112233200-1103220333300010-3110111022330131-3230102032003303"></a>

#### `routes.redirect_route.route_redirect.path_redirect` property

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2212003322032101-3323223302333102-0001220132002003-3321303022000113-1301213002012322-0002011133321100-0112023023202111-1021121030303312"></a>

<a id="canonical-0022300231303230-3232300213011333-0131013303133033-1100131123033331-3012002022101222-1312302133210300-2223121030220020-2003233130202013"></a>

#### `routes.redirect_route.route_redirect.prefix_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3102311302013113-0122122301020032-0331221200123121-3330220113203100-0333000203223013-3010310100131012-3110230230213331-1111020210011003"></a>

<a id="canonical-1033110222133211-2120220210332311-1112230110320211-3132221130333313-2002033331230321-2032330233121002-3121000330223200-2022132313011212"></a>

#### `routes.redirect_route.route_redirect.proto_redirect` property

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["http","https","incoming-proto"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](resources--http_loadbalancer--reference--group-025.md#canonical-1000301113020132-0332010000322121-3312012101100332-3010130312230213-1112330312023330-3112111232332130-1211323030323323-1123232123230311): complete subsection reference.

<a id="canonical-1023222221310302-2102203233103330-0323200223311010-0212110220023120-1033230310031202-3112332211111011-1130332033002022-3233210312010331"></a>

<a id="canonical-3330203112101300-2312203122212030-2223110201331330-2101131222011212-2113310013102112-3302033113320231-0221010022302132-0132333122003102"></a>

#### `routes.redirect_route.route_redirect.replace_params` property

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0323203311130202-0222003013111303-2001231330113223-1333133011232001-3233131332013212-0211120033113321-0332302021000112-2300132011032111"></a>

<a id="canonical-2312311210310132-0130012230120322-3220201312013012-3212200033203331-3210123030202223-0012230223310331-3121003301121220-3301202003312010"></a>

#### `routes.redirect_route.route_redirect.response_code` property

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](resources--http_loadbalancer--reference--group-025.md#canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020): complete subsection reference.

<a id="canonical-1000301113020132-0332010000322121-3312012101100332-3010130312230213-1112330312023330-3112111232332130-1211323030323323-1123232123230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-025.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-2233231301221232-2231323032022230-2310032121130031-3212033001331233-0212033220332231-2322220323010233-0112321312211321-1001211312131201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.redirect_route.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.redirect_route](resources--http_loadbalancer--reference--group-025.md#canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--reference--group-025.md#canonical-2010330103103331-3221112302111121-3011221301002302-3202020300230333-3101100312032213-3003200212130023-2032031332133123-1011310331031011)
- routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-0030323002110330-2012303220102223-2001303002122122-2131131021311312-1131313200202323-3310323321332222-1310100010323123-3010233222321222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103212330200010-1002200200101111-3203321323032332-2322002023022130-2011310020111101-3122113102023032-0332312121303202-3320010033212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_state_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.route_state_disabled

<a id="canonical-3230110113113210-3221321311303201-3311022102121122-0111130013321032-1012101230021333-0002000231013000-2303303123123213-0220133122111332"></a>

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
route_state_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233030011110210-1112111102103333-2313130301030100-2022033113121211-0031211131223103-2133000303130212-3221221230133232-2113131210231203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_state_enabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.route_state_enabled

<a id="canonical-1310132031110303-0300023231112102-1321100032211310-2200031220223023-3002331202033120-1010111121212223-1320113121302113-0222203230210311"></a>

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
route_state_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- routes.simple_route

<a id="canonical-3222110311023220-2231131202300221-3113002110020033-0131333113222113-3312223323223331-2020231333222223-0311232301133203-1211320210323332"></a>

Type: `"object"`. single nested block, Optional.

A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_pools"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033323230021223-3231020102311122-3001112132211232-1020230020113130-2132033201010022-0031001103120001-2223132022032212-2233122031132000"></a>

### Direct properties for `routes.simple_route`

- [advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311): complete subsection reference.

- [auto_host_rewrite](resources--http_loadbalancer--reference--group-027.md#canonical-2301322211233333-2220203210213133-1223123100323130-0320033102000201-0232100103031301-3131223110111100-0211212220020322-2122231131101230): complete subsection reference.

- [caching_disable](resources--http_loadbalancer--reference--group-027.md#canonical-0202222022121123-0202330031011312-2121330022012301-3130103232111220-2101020132131122-0302212100030102-3222333202231011-0111032222221202): complete subsection reference.

- [caching_inherit](resources--http_loadbalancer--reference--group-027.md#canonical-3030300201301203-0333111031333222-0020021113110221-1230233101000133-3230333211121131-3112030300111101-3132120113232101-3332020012003023): complete subsection reference.

- [disable_host_rewrite](resources--http_loadbalancer--reference--group-027.md#canonical-0302130301011103-0313021230021220-0323232031210001-2220233310233323-1010133000201122-1213112103132032-0112010123003213-3123321131203022): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-027.md#canonical-3031023200230323-1323321113310303-2300020320132130-1030230022312313-0003202303001303-0002123020321212-3203111333003001-1323231012303111): complete subsection reference.

<a id="canonical-2022011301231211-3231123220011312-3110303300310302-1133203231130111-2212333333130313-0220202331130322-3232123330102010-1211203312131202"></a>

<a id="canonical-2012123000302111-0030300303002112-0311311223211320-1031112300131103-2132223212121220-2021032310103110-3231011001313200-2200211230221320"></a>

#### `routes.simple_route.host_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3001120032033210-0220211231011001-2100222231322013-3302013033103131-0130203113210303-3313032331002022-3302023133310032-1102102003222223"></a>

<a id="canonical-1330000111233212-0122310320332311-1222330022302200-2200202202023231-3101321101020111-1002033130133122-0310330313100021-0200333011230232"></a>

#### `routes.simple_route.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--http_loadbalancer--reference--group-027.md#canonical-2021321033003230-1330132333221000-2002303112001020-2221033213330222-2123322230122213-3213310100232300-3232300311302131-2032221321131323): complete subsection reference.

- [origin_pools](resources--http_loadbalancer--reference--group-027.md#canonical-2130022000022203-0210102312021103-2132200221323000-1322100313301310-0001331313213000-0231331120331111-0230300123302313-2012003132322232): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-027.md#canonical-3313101103010332-3202233212103301-1302311031332233-0323313323113103-1302101301011133-3132022320301012-2320322112030300-0130020323203110): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-027.md#canonical-1321000220322201-0133323231302021-1011033331200313-1120000203111001-2032133211021021-1010223322202133-2132131102032110-2130300320231101): complete subsection reference.

<a id="canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- routes.simple_route.advanced_options

<a id="canonical-3112132230020132-0112021320130032-2322112130310310-2302312100111023-3123231311312103-3010303303021113-1013311031322010-3312032000310133"></a>

Type: `"object"`. single nested block, Optional.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingObjectAttributes("buffer_policy",
    "common_buffering"),
  validators.ConflictingObjectAttributes("common_hash_policy",
    "specific_hash_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "no_retry_policy"),
  validators.ConflictingObjectAttributes("default_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("disable_mirroring",
    "mirror_policy"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("disable_prefix_rewrite",
    "regex_rewrite"),
  validators.ConflictingObjectAttributes("disable_spdy",
    "enable_spdy"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherited_waf"),
  validators.ConflictingObjectAttributes("disable_web_socket_config",
    "web_socket_config"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingObjectAttributes("no_retry_policy",
    "retry_policy"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030002220330232-2030332021332122-0310021123300320-0231033031222313-2002201203222230-0131012230231221-1233031021012301-3110103330133311"></a>

### Direct properties for `routes.simple_route.advanced_options`

- [app_firewall](resources--http_loadbalancer--reference--group-025.md#canonical-2001003120333132-1102300103312302-3101121230003012-2130113322120203-1111311330221133-1333123230101210-0333012200031023-2230100113201111): complete subsection reference.

- [bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-026.md#canonical-1103203303113331-2302123302001100-1121020032121231-2222122210210133-2002002122010003-2232003301121213-3201221312313321-0202203023112121): complete subsection reference.

- [buffer_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1211001130101123-0202021333221102-1013312002223231-0033202013223031-2322012302201213-0221112223001103-3230112102023130-3102232202213203): complete subsection reference.

- [common_buffering](resources--http_loadbalancer--reference--group-026.md#canonical-3020133322310122-3120022303032302-1333013033231033-0203223021012000-1333112131131133-1101022331131132-0202323320200021-2220033102132032): complete subsection reference.

- [common_hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1220202221213121-2003330033111010-3112031032321133-0331322330011032-2100203210203032-1312312203321022-3331311132203101-1323310311023321): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--reference--group-026.md#canonical-3222022011013211-0231322200331030-3003132300320323-0223133302201212-3011220203110302-1311320103313321-2022323330221202-0131123122131320): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2021122303132221-1311202102331022-3001123032201311-3131110333002321-0130223331320121-2133220111323313-0200111210122312-0202031101030311): complete subsection reference.

- [default_retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-3013002323211230-2000321010322031-0202220333302003-3320223120203023-3311023301310130-1123220212320222-2303233031330210-0222030031022311): complete subsection reference.

<a id="canonical-1020001231020323-2032232202330032-1212122323221322-3001000133101120-2320330013320230-1023131330100220-1233110321231330-3231112002200232"></a>

<a id="canonical-1133213132231302-3111302300301320-2031101231221330-1303213210000322-1200112100230203-2003211210110303-0320223001032232-1331301000002001"></a>

#### `routes.simple_route.advanced_options.disable_location_add` property

Type: `"bool"`. Optional.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [disable_mirroring](resources--http_loadbalancer--reference--group-026.md#canonical-1223210220132003-1032031311131022-0221333010313232-3300232111321333-3020023100102303-3000311333203311-2323022130322120-1301313300201232): complete subsection reference.

- [disable_prefix_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-3101222121331231-0002101210102120-1000231102102220-2110231230123212-2111001231312132-2121211202233201-1300220103223003-2023201222120123): complete subsection reference.

- [disable_spdy](resources--http_loadbalancer--reference--group-026.md#canonical-1333112331031332-3110230102112221-1123102212221222-1320333331030110-2102300032332222-3010102120200200-0200212202113120-3221312122220012): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--reference--group-026.md#canonical-2012322210102001-1000202101032101-0301012012020133-2222012131120302-2033130323220023-0131111302321232-1103220102022322-0130330011322111): complete subsection reference.

- [disable_web_socket_config](resources--http_loadbalancer--reference--group-026.md#canonical-1123013100222203-3012313301032221-2333300120200121-2031023303323312-2330102330202121-3313212031211202-1300121232122232-3230133031002030): complete subsection reference.

- [do_not_retract_cluster](resources--http_loadbalancer--reference--group-026.md#canonical-3013100112201010-3320010213003210-2110103020232131-1233302023023333-1320111303223010-1220201100133030-2322330330102010-1111303301333333): complete subsection reference.

- [enable_spdy](resources--http_loadbalancer--reference--group-026.md#canonical-3220221103013030-1033023211020212-2313000220302030-0112212123203222-2223110003210120-1330110301122333-0120013212320200-2011130012212031): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-026.md#canonical-2123231212310000-0130220212030233-2332222131111223-3330003103011023-0211122003332012-0200233023011101-3232021110002213-0302122311123131): complete subsection reference.

- [inherited_bot_defense_javascript_injection](resources--http_loadbalancer--reference--group-026.md#canonical-3101301002211102-3101130203012322-2333001122113100-1010021103332111-0203220210230210-1230033102313122-1102112111223121-2210120101222000): complete subsection reference.

- [inherited_waf](resources--http_loadbalancer--reference--group-026.md#canonical-0213201113012322-1102122130221322-3332120321331122-3312320023133031-1013323333130130-0033111303102122-3133232312300201-0023120200310323): complete subsection reference.

- [inherited_waf_exclusion](resources--http_loadbalancer--reference--group-026.md#canonical-2002012001013000-3310311312101011-3313033113321211-2131033121023113-1002001131311300-2001101231211002-3100031020332322-2132012212030231): complete subsection reference.

- [mirror_policy](resources--http_loadbalancer--reference--group-026.md#canonical-1013001221331220-1013310121201113-3023002030203320-3030001023331313-3101303133113133-2112000032110210-2021222012332022-3303330132233023): complete subsection reference.

- [no_retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2220333121303102-3132022333332033-2111322310103130-0230222232313212-1102113200030202-3202032303023333-2120330013220013-0120312220300111): complete subsection reference.

<a id="canonical-0202322133003233-2323130221211011-0331013321202102-3001310231320022-3020301123122111-2201333333123123-0011330321020221-2001302303112231"></a>

<a id="canonical-2212000212200030-3312330103231101-3212201231202222-1031233013202332-1232303022222123-3003103022012230-3302010331100331-1013100220300212"></a>

#### `routes.simple_route.advanced_options.prefix_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[disable\_prefix\_rewrite regular expression\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regular expression path
matching, the entire path (not including the query string) will be swapped with this value.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0110321121101133-0021201220321312-0111110232023013-1112120320000010-1310323201022303-1002021330102010-3022330211023330-0303323131312031"></a>

<a id="canonical-0011200312121123-2033031121131330-1013210112112130-0033131022120212-1310131330231013-2303101112003311-0323030213011000-0233323203033322"></a>

#### `routes.simple_route.advanced_options.priority` property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Additional upstream details:

Priority routing for each request. Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DEFAULT","HIGH"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [regex_rewrite](resources--http_loadbalancer--reference--group-026.md#canonical-0320010111202212-1023320033013103-0030311122023131-1323130311311223-3130123101023222-0210202202123330-0003303022231213-3303032312010332): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-2023223212320223-0320123102221000-0122230301222332-1212213120231312-2022132310113113-1202103101033202-1130222000320313-0331210201110321): complete subsection reference.

<a id="canonical-3003321220100110-0303003032212122-1111131323131123-2301133021010232-3233102102201030-1113103310301201-2211113220032032-1320320233220122"></a>

<a id="canonical-1233001313000133-3230101231220232-2020011333233132-2311102000311000-1120211020113113-3023313130330000-1130113000110131-2232202331210233"></a>

#### `routes.simple_route.advanced_options.request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-3332100030233323-1210323031233010-0221332300111122-3113321100000111-2213201102010221-1201111320321312-2022100223303033-3210211102121313): complete subsection reference.

<a id="canonical-2103233003313121-3222000123220312-2101323120100300-2320100202000003-0201221132200213-2002032213002230-0321321121021002-1230130213102030"></a>

<a id="canonical-0321130310323300-3310232201231112-1131202000313112-0121312030213200-1123100031301033-1032303200320230-3012103322123030-1201011203102331"></a>

#### `routes.simple_route.advanced_options.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-0321201123011332-0122110021123303-3123103012221211-1322301201120121-1030212012333220-0232133121332002-3213022033303112-3232001023233020): complete subsection reference.

<a id="canonical-3131112013031110-1121310200030031-3331302220312312-3001332110312231-2320310001313021-1311102032312211-2013102033323303-0132320333132331"></a>

<a id="canonical-1100122123130212-2302021030110233-1303022231311232-1133331303000200-1121011113332331-3230321023021012-3102110130021013-0003311202333212"></a>

#### `routes.simple_route.advanced_options.response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--http_loadbalancer--reference--group-026.md#canonical-1302033232022313-1021121312312133-3011110120013233-0011303102032230-1312323112002300-2310303030311201-2003333031333222-2030332033023230): complete subsection reference.

<a id="canonical-0213021013133312-0203013313301033-0331312100000003-3103111103102000-3130030023222222-1223031020110301-2320330112323122-0302110112300313"></a>

<a id="canonical-2233121130121333-2122333232221031-2311222303113203-3221302102102112-1013330103221323-1200320301303123-2022221203313030-1221032020301111"></a>

#### `routes.simple_route.advanced_options.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](resources--http_loadbalancer--reference--group-026.md#canonical-1001202133030320-0013120123332230-2010322030001110-0213023101000331-3211331232101131-1213100001010311-0221112031110323-2311021231132111): complete subsection reference.

- [retry_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2031213021312332-2113132313130100-3222333013111133-0122010300022110-1323302212201331-0133200321222002-3221100221013001-0011102330211230): complete subsection reference.

- [specific_hash_policy](resources--http_loadbalancer--reference--group-026.md#canonical-2332103311111231-3001202122300101-0003322101013201-2300220101202301-0330222321301302-3233120102200011-0001011001021100-2330220212112132): complete subsection reference.

<a id="canonical-0132110101311213-3022130021113002-0331332231002321-0022322011301222-1102233001102311-0333023130223021-0331031112021222-1111300122223322"></a>

<a id="canonical-2230320203312200-3211103212002021-1231331023003002-0020013132010022-1011131230111212-1230212213200322-1001122210201121-1011222231123230"></a>

#### `routes.simple_route.advanced_options.timeout` property

Type: `"number"`. Optional.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](resources--http_loadbalancer--reference--group-027.md#canonical-3123100221313210-3130201120132022-0032331322102220-1220110221031321-3132211020111030-3012322332111331-0210011333220010-2333010312023201): complete subsection reference.

- [web_socket_config](resources--http_loadbalancer--reference--group-027.md#canonical-2221130310300232-3333311201210030-2220032332203101-0111312233003310-0330020033223201-2312212013003202-2330320321012310-3132220122331020): complete subsection reference.

<a id="canonical-2001003120333132-1102300103312302-3101121230003012-2130113322120203-1111311330221133-1333123230101210-0333012200031023-2230100113201111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.simple_route.advanced_options.app_firewall` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100)
- [routes.simple_route](resources--http_loadbalancer--reference--group-025.md#canonical-1130313311103011-3030331213003000-2111122030020101-0203200030102222-1013001322010120-3333112332001321-3313210011003030-1201120031233232)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--reference--group-025.md#canonical-3033033103020202-1232223012232131-1121301112112030-2332010303303022-3101022302330020-0320200112212131-1130103223120322-2333302301310311)
- routes.simple_route.advanced_options.app_firewall

<a id="canonical-1000212213232202-2100301113231300-2120101013333023-3331233232132321-0311303111131302-1212002120023023-3331301001120301-3012202322222021"></a>

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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121010303031121-0103131211112010-2123101011030232-0332003101132021-1111133310200320-3023320133312210-2221231011232101-3000333221202210"></a>

### Direct properties for `routes.simple_route.advanced_options.app_firewall`

<a id="canonical-0111132101120012-1223113220311220-0022113222213022-3231103021001013-1011300213303331-3220013012331202-1130222023001310-2320213331113220"></a>

#### `routes.simple_route.advanced_options.app_firewall.name` property

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

<a id="canonical-2121202211030001-0212121210021011-2332322233333113-1112203310200223-3201013232011122-0032131333121110-0020310113320132-0311313103013113"></a>

<a id="canonical-0233310011231200-0233003231330023-1011031201123121-0211030003321033-1131133013212200-3320212012131003-1030111322121331-3201023110311031"></a>

#### `routes.simple_route.advanced_options.app_firewall.namespace` property

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

<a id="canonical-1210330122100320-1202022131230121-3033220032101300-2023233120323223-1013333221331013-3120133021200033-0203233313222022-0323103113121130"></a>

<a id="canonical-3332201333001012-3221103132312132-3100201123212211-0111010132332113-0000021003303323-0323101302201111-0330201333133221-0121333233013120"></a>

#### `routes.simple_route.advanced_options.app_firewall.tenant` property

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
